// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package compoundcontract

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

// CErc20ABI is the input ABI used to generate the binding from.
const CErc20ABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying_\",\"type\":\"address\"},{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialExchangeRateMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"contractCToken\",\"name\":\"cTokenCollateral\",\"type\":\"address\"}],\"name\":\"liquidateBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeem\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"}],\"name\":\"redeemUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowBehalf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"underlying\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// CErc20FuncSigs maps the 4-byte function signature to its string representation.
var CErc20FuncSigs = map[string]string{
	"e9c714f2": "_acceptAdmin()",
	"601a0bf1": "_reduceReserves(uint256)",
	"4576b5db": "_setComptroller(address)",
	"f2b3abbd": "_setInterestRateModel(address)",
	"b71d1a0c": "_setPendingAdmin(address)",
	"fca7820b": "_setReserveFactor(uint256)",
	"6c540baf": "accrualBlockNumber()",
	"a6afed95": "accrueInterest()",
	"f851a440": "admin()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"3af9e669": "balanceOfUnderlying(address)",
	"c5ebeaec": "borrow(uint256)",
	"17bfdfbc": "borrowBalanceCurrent(address)",
	"95dd9193": "borrowBalanceStored(address)",
	"aa5af0fd": "borrowIndex()",
	"f8f9da28": "borrowRatePerBlock()",
	"5fe3b567": "comptroller()",
	"313ce567": "decimals()",
	"bd6d894d": "exchangeRateCurrent()",
	"182df0f5": "exchangeRateStored()",
	"c37f68e2": "getAccountSnapshot(address)",
	"3b1d21a2": "getCash()",
	"675d972c": "initialExchangeRateMantissa()",
	"f3fdb15a": "interestRateModel()",
	"fe9c44ae": "isCToken()",
	"f5e3c462": "liquidateBorrow(address,uint256,address)",
	"a0712d68": "mint(uint256)",
	"06fdde03": "name()",
	"26782247": "pendingAdmin()",
	"db006a75": "redeem(uint256)",
	"852a12e3": "redeemUnderlying(uint256)",
	"0e752702": "repayBorrow(uint256)",
	"2608f818": "repayBorrowBehalf(address,uint256)",
	"173b9904": "reserveFactorMantissa()",
	"b2a02ff1": "seize(address,address,uint256)",
	"ae9d70b0": "supplyRatePerBlock()",
	"95d89b41": "symbol()",
	"47bd3718": "totalBorrows()",
	"73acee98": "totalBorrowsCurrent()",
	"8f840ddd": "totalReserves()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"6f307dc3": "underlying()",
}

// CErc20Bin is the compiled bytecode used for deploying new contracts.
var CErc20Bin = "0x60806040523480156200001157600080fd5b506040516200519c3803806200519c833981810160405260e08110156200003757600080fd5b81516020830151604080850151606086015160808701805193519597949692959194919392820192846401000000008211156200007357600080fd5b9083019060208201858111156200008957600080fd5b8251640100000000811182820188101715620000a457600080fd5b82525081516020918201929091019080838360005b83811015620000d3578181015183820152602001620000b9565b50505050905090810190601f168015620001015780820380516001836020036101000a031916815260200191505b50604052602001805160405193929190846401000000008211156200012557600080fd5b9083019060208201858111156200013b57600080fd5b82516401000000008111828201881017156200015657600080fd5b82525081516020918201929091019080838360005b83811015620001855781810151838201526020016200016b565b50505050905090810190601f168015620001b35780820380516001836020036101000a031916815260200191505b50604052602001516001600055600480546001600160a01b03191633179055600886905591508690508585858585836200021f5760405162461bcd60e51b81526004018080602001828103825260308152602001806200516c6030913960400191505060405180910390fd5b600062000235876001600160e01b03620003ce16565b905080156200028b576040805162461bcd60e51b815260206004820152601a60248201527f53657474696e6720636f6d7074726f6c6c6572206661696c6564000000000000604482015290519081900360640190fd5b6200029e6001600160e01b036200053116565b600a55670de0b6b3a7640000600b55620002c1866001600160e01b036200053616565b90508015620003025760405162461bcd60e51b81526004018080602001828103825260228152602001806200514a6022913960400191505060405180910390fd5b83516200031790600190602087019062000747565b5082516200032d90600290602086019062000747565b50506003555050601280546001600160a01b0319166001600160a01b038c81169190911791829055604080516318160ddd60e01b815290519290911694506318160ddd93506004808201935060209291829003018186803b1580156200039257600080fd5b505afa158015620003a7573d6000803e3d6000fd5b505050506040513d6020811015620003be57600080fd5b50620007e9975050505050505050565b6004546000906001600160a01b031633146200040357620003fb6001603f6001600160e01b03620006d716565b90506200052c565b60065460408051623f1ee960e11b815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b1580156200044957600080fd5b505afa1580156200045e573d6000803e3d6000fd5b505050506040513d60208110156200047557600080fd5b5051620004c9576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160005b9150505b919050565b435b90565b60045460009081906001600160a01b031633146200056e5762000565600160426001600160e01b03620006d716565b9150506200052c565b620005816001600160e01b036200053116565b600a5414620005a15762000565600a60416001600160e01b03620006d716565b600760009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b158015620005f357600080fd5b505afa15801562000608573d6000803e3d6000fd5b505050506040513d60208110156200061f57600080fd5b505162000673576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600780546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a1600062000528565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa08360108111156200070757fe5b83604d8111156200071457fe5b604080519283526020830191909152600082820152519081900360600190a18260108111156200074057fe5b9392505050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106200078a57805160ff1916838001178555620007ba565b82800160010185558215620007ba579182015b82811115620007ba5782518255916020019190600101906200079d565b50620007c8929150620007cc565b5090565b6200053391905b80821115620007c85760008155600101620007d3565b61495180620007f96000396000f3fe608060405234801561001057600080fd5b506004361061028a5760003560e01c80638f840ddd1161015c578063c37f68e2116100ce578063f3fdb15a11610087578063f3fdb15a14610708578063f5e3c46214610710578063f851a44014610746578063f8f9da281461074e578063fca7820b14610756578063fe9c44ae146107735761028a565b8063c37f68e214610626578063c5ebeaec14610672578063db006a751461068f578063dd62ed3e146106ac578063e9c714f2146106da578063f2b3abbd146106e25761028a565b8063a9059cbb11610120578063a9059cbb14610586578063aa5af0fd146105b2578063ae9d70b0146105ba578063b2a02ff1146105c2578063b71d1a0c146105f8578063bd6d894d1461061e5761028a565b80638f840ddd1461052b57806395d89b411461053357806395dd91931461053b578063a0712d6814610561578063a6afed951461057e5761028a565b80633af9e66911610200578063675d972c116101b9578063675d972c146104c85780636c540baf146104d05780636f307dc3146104d857806370a08231146104e057806373acee9814610506578063852a12e31461050e5761028a565b80633af9e669146104475780633b1d21a21461046d5780634576b5db1461047557806347bd37181461049b5780635fe3b567146104a3578063601a0bf1146104ab5761028a565b806318160ddd1161025257806318160ddd146103a9578063182df0f5146103b157806323b872dd146103b95780632608f818146103ef578063267822471461041b578063313ce5671461043f5761028a565b806306fdde031461028f578063095ea7b31461030c5780630e7527021461034c578063173b99041461037b57806317bfdfbc14610383575b600080fd5b61029761077b565b6040805160208082528351818301528351919283929083019185019080838360005b838110156102d15781810151838201526020016102b9565b50505050905090810190601f1680156102fe5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6103386004803603604081101561032257600080fd5b506001600160a01b038135169060200135610808565b604080519115158252519081900360200190f35b6103696004803603602081101561036257600080fd5b5035610875565b60408051918252519081900360200190f35b610369610888565b6103696004803603602081101561039957600080fd5b50356001600160a01b031661088e565b610369610941565b610369610947565b610338600480360360608110156103cf57600080fd5b506001600160a01b038135811691602081013590911690604001356109aa565b6103696004803603604081101561040557600080fd5b506001600160a01b038135169060200135610a10565b610423610a23565b604080516001600160a01b039092168252519081900360200190f35b610369610a32565b6103696004803603602081101561045d57600080fd5b50356001600160a01b0316610a38565b610369610aa6565b6103696004803603602081101561048b57600080fd5b50356001600160a01b0316610ab5565b610369610c04565b610423610c0a565b610369600480360360208110156104c157600080fd5b5035610c19565b610369610ca1565b610369610ca7565b610423610cad565b610369600480360360208110156104f657600080fd5b50356001600160a01b0316610cbc565b610369610cd7565b6103696004803603602081101561052457600080fd5b5035610d81565b610369610d8c565b610297610d92565b6103696004803603602081101561055157600080fd5b50356001600160a01b0316610dea565b6103696004803603602081101561057757600080fd5b5035610e47565b610369610e52565b6103386004803603604081101561059c57600080fd5b506001600160a01b03813516906020013561124b565b6103696112b0565b6103696112b6565b610369600480360360608110156105d857600080fd5b506001600160a01b03813581169160208101359091169060400135611581565b6103696004803603602081101561060e57600080fd5b50356001600160a01b0316611832565b6103696118b9565b61064c6004803603602081101561063c57600080fd5b50356001600160a01b0316611964565b604080519485526020850193909352838301919091526060830152519081900360800190f35b6103696004803603602081101561068857600080fd5b50356119f9565b610369600480360360208110156106a557600080fd5b5035611a04565b610369600480360360408110156106c257600080fd5b506001600160a01b0381358116916020013516611a0f565b610369611a3a565b610369600480360360208110156106f857600080fd5b50356001600160a01b0316611b29565b610423611b63565b6103696004803603606081101561072657600080fd5b506001600160a01b03813581169160208101359160409091013516611b72565b610423611b7f565b610369611b8e565b6103696004803603602081101561076c57600080fd5b5035611c6a565b610338611ca4565b60018054604080516020600284861615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156108005780601f106107d557610100808354040283529160200191610800565b820191906000526020600020905b8154815290600101906020018083116107e357829003601f168201915b505050505081565b3360008181526010602090815260408083206001600160a01b03871680855290835281842086905581518681529151939493909284927f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925929081900390910190a360019150505b92915050565b600061088082611ca9565b90505b919050565b60095481565b60008054600101808255816108a1610e52565b146108ec576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b6108f583610dea565b91505b600054811461093b576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b50919050565b600e5481565b6000806000610954611ce5565b9092509050600082600381111561096757fe5b146109a35760405162461bcd60e51b81526004018080602001828103825260358152602001806148906035913960400191505060405180910390fd5b9150505b90565b60008054600101808255816109c133878787611d93565b1491505b6000548114610a08576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b509392505050565b6000610a1c83836120a1565b9392505050565b6005546001600160a01b031681565b60035481565b6000610a426145e9565b6040518060200160405280610a556118b9565b90526001600160a01b0384166000908152600f6020526040812054919250908190610a8190849061212b565b90925090506000826003811115610a9457fe5b14610a9e57600080fd5b949350505050565b6000610ab061217f565b905090565b6004546000906001600160a01b03163314610add57610ad66001603f6121ff565b9050610883565b60065460408051623f1ee960e11b815290516001600160a01b0392831692851691627e3dd2916004808301926020929190829003018186803b158015610b2257600080fd5b505afa158015610b36573d6000803e3d6000fd5b505050506040513d6020811015610b4c57600080fd5b5051610b9f576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600680546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517f7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d9281900390910190a160009392505050565b600c5481565b6006546001600160a01b031681565b6000805460010180825581610c2c610e52565b90508015610c5257610c4a816010811115610c4357fe5b60306121ff565b9250506108f8565b610c5b84612265565b925050600054811461093b576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b60085481565b600a5481565b6012546001600160a01b031681565b6001600160a01b03166000908152600f602052604090205490565b6000805460010180825581610cea610e52565b14610d35576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b600c5491506000548114610d7d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b5090565b6000610880826123e3565b600d5481565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156108005780601f106107d557610100808354040283529160200191610800565b6000806000610df884612420565b90925090506000826003811115610e0b57fe5b14610a1c5760405162461bcd60e51b81526004018080602001828103825260378152602001806147646037913960400191505060405180910390fd5b6000610880826124d4565b6000610e5c6145fc565b6007546001600160a01b03166315f24053610e7561217f565b600c54600d546040518463ffffffff1660e01b8152600401808481526020018381526020018281526020019350505050604080518083038186803b158015610ebc57600080fd5b505afa158015610ed0573d6000803e3d6000fd5b505050506040513d6040811015610ee657600080fd5b50805160209182015160408401819052918301526601c6bf526340001015610f55576040805162461bcd60e51b815260206004820152601c60248201527f626f72726f772072617465206973206162737572646c79206869676800000000604482015290519081900360640190fd5b602081015115610f7857610f7060056002836020015161250f565b9150506109a7565b610f80612575565b60608201819052600a54610f949190612579565b6080830181905282826003811115610fa857fe5b6003811115610fb357fe5b9052506000905081516003811115610fc757fe5b14610fce57fe5b610fee60405180602001604052808360400151815250826080015161259c565b60a083018190528282600381111561100257fe5b600381111561100d57fe5b905250600090508151600381111561102157fe5b1461104257610f70600960068360000151600381111561103d57fe5b61250f565b6110528160a00151600c5461212b565b60c083018190528282600381111561106657fe5b600381111561107157fe5b905250600090508151600381111561108557fe5b146110a157610f70600960018360000151600381111561103d57fe5b6110b18160c00151600c54612604565b60e08301819052828260038111156110c557fe5b60038111156110d057fe5b90525060009050815160038111156110e457fe5b1461110057610f70600960048360000151600381111561103d57fe5b61112160405180602001604052806009548152508260c00151600d5461262a565b61010083018190528282600381111561113657fe5b600381111561114157fe5b905250600090508151600381111561115557fe5b1461117157610f70600960058360000151600381111561103d57fe5b6111848160a00151600b54600b5461262a565b61012083018190528282600381111561119957fe5b60038111156111a457fe5b90525060009050815160038111156111b857fe5b146111d457610f70600960038360000151600381111561103d57fe5b606080820151600a55610120820151600b81905560e0830151600c819055610100840151600d5560c08401516040805191825260208201939093528083019190915290517f875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9929181900390910190a1600091505090565b600080546001018082558161126233338787611d93565b1491505b60005481146112a9576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b5092915050565b600b5481565b6000806112c1610947565b60075490915060009081906001600160a01b03166315f240536112e261217f565b600c54600d546040518463ffffffff1660e01b8152600401808481526020018381526020018281526020019350505050604080518083038186803b15801561132957600080fd5b505afa15801561133d573d6000803e3d6000fd5b505050506040513d604081101561135357600080fd5b5080516020909101519092509050811561139e5760405162461bcd60e51b81526004018080602001828103825260318152602001806148036031913960400191505060405180910390fd5b60006113a86145e9565b6113c2604051806020016040528087815250600e5461259c565b909250905060008260038111156113d557fe5b146114115760405162461bcd60e51b815260040180806020018281038252603181526020018061479b6031913960400191505060405180910390fd5b600061141b6145e9565b611427600c5484612686565b9092509050600082600381111561143a57fe5b146114765760405162461bcd60e51b81526004018080602001828103825260318152602001806146df6031913960400191505060405180910390fd5b60006114806145e9565b6114b06040518060200160405280670de0b6b3a764000081525060405180602001604052806009548152506126e5565b909250905060008260038111156114c357fe5b146114ff5760405162461bcd60e51b815260040180806020018281038252603c815260200180614854603c913960400191505060405180910390fd5b60006115096145e9565b61152260405180602001604052808b815250848761271f565b9092509050600082600381111561153557fe5b146115715760405162461bcd60e51b81526004018080602001828103825260318152602001806147336031913960400191505060405180910390fd5b519a505050505050505050505090565b600080546001018082556006546040805163d02f735160e01b81523060048201523360248201526001600160a01b03888116604483015287811660648301526084820187905291518593929092169163d02f73519160a48082019260209290919082900301818787803b1580156115f757600080fd5b505af115801561160b573d6000803e3d6000fd5b505050506040513d602081101561162157600080fd5b505190508015611640576116386003601b8361250f565b9250506109c5565b856001600160a01b0316856001600160a01b03161415611666576116386006601c6121ff565b6001600160a01b0385166000908152600f60205260408120548190819061168d9088612579565b909350915060008360038111156116a057fe5b146116c3576116b86009601a85600381111561103d57fe5b9550505050506109c5565b6001600160a01b0389166000908152600f60205260409020546116e69088612604565b909350905060008360038111156116f957fe5b14611711576116b86009601985600381111561103d57fe5b6001600160a01b038089166000818152600f60209081526040808320879055938d168083529184902085905583518b815293519193600080516020614834833981519152929081900390910190a360065460408051636d35bf9160e01b81523060048201523360248201526001600160a01b038c811660448301528b81166064830152608482018b905291519190921691636d35bf919160a480830192600092919082900301818387803b1580156117c857600080fd5b505af11580156117dc573d6000803e3d6000fd5b50600092506117e9915050565b9550505050506000548114610a08576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6004546000906001600160a01b0316331461185357610ad6600160456121ff565b600580546001600160a01b038481166001600160a01b0319831681179093556040805191909216808252602082019390935281517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9929181900390910190a16000610a1c565b60008054600101808255816118cc610e52565b14611917576040805162461bcd60e51b81526020600482015260166024820152751858d8dc9d59481a5b9d195c995cdd0819985a5b195960521b604482015290519081900360640190fd5b61191f610947565b91506000548114610d7d576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6001600160a01b0381166000908152600f602052604081205481908190819081808061198f89612420565b9350905060008160038111156119a157fe5b146119bf5760095b9750600096508695508594506119f29350505050565b6119c7611ce5565b9250905060008160038111156119d957fe5b146119e55760096119a9565b5060009650919450925090505b9193509193565b600061088082612769565b6000610880826127a4565b6001600160a01b03918216600090815260106020908152604080832093909416825291909152205490565b6005546000906001600160a01b031633141580611a55575033155b15611a6d57611a66600160006121ff565b90506109a7565b60048054600580546001600160a01b038082166001600160a01b031980861682179687905590921690925560408051938316808552949092166020840152815190927ff9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc92908290030190a1600554604080516001600160a01b038085168252909216602083015280517fca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a99281900390910190a160009250505090565b600080611b34610e52565b90508015611b5a57611b52816010811115611b4b57fe5b60406121ff565b915050610883565b610a1c836127da565b6007546001600160a01b031681565b6000610a9e84848461294a565b6004546001600160a01b031681565b600754600090819081906001600160a01b03166315f24053611bae61217f565b600c54600d546040518463ffffffff1660e01b8152600401808481526020018381526020018281526020019350505050604080518083038186803b158015611bf557600080fd5b505afa158015611c09573d6000803e3d6000fd5b505050506040513d6040811015611c1f57600080fd5b508051602090910151909250905081156109a35760405162461bcd60e51b81526004018080602001828103825260378152602001806147cc6037913960400191505060405180910390fd5b6000805460010180825581611c7d610e52565b90508015611c9b57610c4a816010811115611c9457fe5b60466121ff565b610c5b84612a52565b600181565b6000805460010180825581611cbc610e52565b90508015611cda57610c4a816010811115611cd357fe5b60366121ff565b610c5b333386612af5565b600080600e5460001415611d00575050600854600090611d8f565b6000611d0a61217f565b90506000611d166145e9565b6000611d2784600c54600d54612f48565b935090506000816003811115611d3957fe5b14611d4d57945060009350611d8f92505050565b611d5983600e54612f86565b925090506000816003811115611d6b57fe5b14611d7f57945060009350611d8f92505050565b5051600094509250611d8f915050565b9091565b600654604080516317b9b84b60e31b81523060048201526001600160a01b03868116602483015285811660448301526064820185905291516000938493169163bdcdc25891608480830192602092919082900301818787803b158015611df857600080fd5b505af1158015611e0c573d6000803e3d6000fd5b505050506040513d6020811015611e2257600080fd5b505190508015611e4157611e396003604a8361250f565b915050610a9e565b836001600160a01b0316856001600160a01b03161415611e6757611e396002604b6121ff565b60006001600160a01b038781169087161415611e865750600019611eae565b506001600160a01b038086166000908152601060209081526040808320938a16835292905220545b600080600080611ebe8589612579565b90945092506000846003811115611ed157fe5b14611eef57611ee26009604b6121ff565b9650505050505050610a9e565b6001600160a01b038a166000908152600f6020526040902054611f129089612579565b90945091506000846003811115611f2557fe5b14611f3657611ee26009604c6121ff565b6001600160a01b0389166000908152600f6020526040902054611f599089612604565b90945090506000846003811115611f6c57fe5b14611f7d57611ee26009604d6121ff565b6001600160a01b03808b166000908152600f6020526040808220859055918b168152208190556000198514611fd5576001600160a01b03808b166000908152601060209081526040808320938f168352929052208390555b886001600160a01b03168a6001600160a01b03166000805160206148348339815191528a6040518082815260200191505060405180910390a36006546040805163352b4a3f60e11b81523060048201526001600160a01b038d811660248301528c81166044830152606482018c905291519190921691636a56947e91608480830192600092919082900301818387803b15801561207157600080fd5b505af1158015612085573d6000803e3d6000fd5b5060009250612092915050565b9b9a5050505050505050505050565b60008054600101808255816120b4610e52565b905080156120da576120d28160108111156120cb57fe5b60356121ff565b925050611266565b6120e5338686612af5565b92505060005481146112a9576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b60008060006121386145e9565b612142868661259c565b9092509050600082600381111561215557fe5b146121665750915060009050612178565b600061217182613036565b9350935050505b9250929050565b601254604080516370a0823160e01b815230600482015290516000926001600160a01b03169182916370a0823191602480820192602092909190829003018186803b1580156121cd57600080fd5b505afa1580156121e1573d6000803e3d6000fd5b505050506040513d60208110156121f757600080fd5b505191505090565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa083601081111561222e57fe5b83604d81111561223a57fe5b604080519283526020830191909152600082820152519081900360600190a1826010811115610a1c57fe5b600454600090819081906001600160a01b031633146122935761228a600160316121ff565b92505050610883565b61229b612575565b600a54146122af5761228a600a60336121ff565b836122b861217f565b10156122ca5761228a600e60326121ff565b600d548411156122e05761228a600260346121ff565b50600d54838103908111156123265760405162461bcd60e51b81526004018080602001828103825260248152602001806148f96024913960400191505060405180910390fd5b600d819055600454612341906001600160a01b031685613045565b9150600082601081111561235157fe5b1461238d5760405162461bcd60e51b81526004018080602001828103825260238152602001806147106023913960400191505060405180910390fd5b600454604080516001600160a01b03909216825260208201869052818101839052517f3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e9181900360600190a16000949350505050565b60008054600101808255816123f6610e52565b9050801561241457610c4a81601081111561240d57fe5b60276121ff565b610c5b33600086613101565b6001600160a01b0381166000908152601160205260408120805482918291829182916124575750600094508493506124cf92505050565b6124678160000154600b5461360a565b9094509250600084600381111561247a57fe5b1461248f5750919350600092506124cf915050565b61249d838260010154613649565b909450915060008460038111156124b057fe5b146124c55750919350600092506124cf915050565b5060009450925050505b915091565b60008054600101808255816124e7610e52565b9050801561250557610c4a8160108111156124fe57fe5b601e6121ff565b610c5b3385613674565b60007f45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa084601081111561253e57fe5b84604d81111561254a57fe5b604080519283526020830191909152818101859052519081900360600190a1836010811115610a9e57fe5b4390565b600080838311612590575060009050818303612178565b50600390506000612178565b60006125a66145e9565b6000806125b786600001518661360a565b909250905060008260038111156125ca57fe5b146125e957506040805160208101909152600081529092509050612178565b60408051602081019091529081526000969095509350505050565b60008083830184811061261c57600092509050612178565b506002915060009050612178565b60008060006126376145e9565b612641878761259c565b9092509050600082600381111561265457fe5b14612665575091506000905061267e565b61267761267182613036565b86612604565b9350935050505b935093915050565b60006126906145e9565b6000806126a5670de0b6b3a76400008761360a565b909250905060008260038111156126b857fe5b146126d757506040805160208101909152600081529092509050612178565b612171818660000151612f86565b60006126ef6145e9565b60008061270486600001518660000151612579565b60408051602081019091529081529097909650945050505050565b60006127296145e9565b60006127336145e9565b61273d8787613abc565b9092509050600082600381111561275057fe5b1461275f57909250905061267e565b6126778186613abc565b600080546001018082558161277c610e52565b9050801561279a57610c4a81601081111561279357fe5b60086121ff565b610c5b3385613ba5565b60008054600101808255816127b7610e52565b905080156127ce57610c4a81601081111561240d57fe5b610c5b33856000613101565b60045460009081906001600160a01b031633146127fd57611b52600160426121ff565b612805612575565b600a541461281957611b52600a60416121ff565b600760009054906101000a90046001600160a01b03169050826001600160a01b0316632191f92a6040518163ffffffff1660e01b815260040160206040518083038186803b15801561286a57600080fd5b505afa15801561287e573d6000803e3d6000fd5b505050506040513d602081101561289457600080fd5b50516128e7576040805162461bcd60e51b815260206004820152601c60248201527f6d61726b6572206d6574686f642072657475726e65642066616c736500000000604482015290519081900360640190fd5b600780546001600160a01b0319166001600160a01b03858116918217909255604080519284168352602083019190915280517fedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f9269281900390910190a16000610a1c565b600080546001018082558161295d610e52565b9050801561297b5761163881601081111561297457fe5b600f6121ff565b836001600160a01b031663a6afed956040518163ffffffff1660e01b8152600401602060405180830381600087803b1580156129b657600080fd5b505af11580156129ca573d6000803e3d6000fd5b505050506040513d60208110156129e057600080fd5b505190508015612a00576116388160108111156129f957fe5b60106121ff565b612a0c33878787613f0b565b9250506000548114610a08576040805162461bcd60e51b815260206004820152600a6024820152691c994b595b9d195c995960b21b604482015290519081900360640190fd5b6004546000906001600160a01b03163314612a7357610ad6600160476121ff565b612a7b612575565b600a5414612a8f57610ad6600a60486121ff565b670de0b6b3a7640000821115612aab57610ad6600260496121ff565b6009805490839055604080518281526020810185905281517faaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460929181900390910190a16000610a1c565b60065460408051631200453160e11b81523060048201526001600160a01b0386811660248301528581166044830152606482018590529151600093849316916324008a6291608480830192602092919082900301818787803b158015612b5a57600080fd5b505af1158015612b6e573d6000803e3d6000fd5b505050506040513d6020811015612b8457600080fd5b505190508015612ba357612b9b600360388361250f565b915050610a1c565b612bab612575565b600a5414612bbf57612b9b600a60396121ff565b612bc7614656565b6001600160a01b0385166000908152601160205260409020600101546060820152612bf185612420565b6080830181905260208301826003811115612c0857fe5b6003811115612c1357fe5b9052506000905081602001516003811115612c2a57fe5b14612c4f57612c46600960378360200151600381111561103d57fe5b92505050610a1c565b600019841415612c685760808101516040820152612c70565b604081018490525b612c7e8682604001516143e7565b81906010811115612c8b57fe5b90816010811115612c9857fe5b905250600081516010811115612caa57fe5b14612cbc578051612c4690603c6121ff565b612cce81608001518260400151612579565b60a0830181905260208301826003811115612ce557fe5b6003811115612cf057fe5b9052506000905081602001516003811115612d0757fe5b14612d2357612c466009603a8360200151600381111561103d57fe5b612d33600c548260400151612579565b60c0830181905260208301826003811115612d4a57fe5b6003811115612d5557fe5b9052506000905081602001516003811115612d6c57fe5b14612d8857612c466009603b8360200151600381111561103d57fe5b612d9686826040015161451b565b81906010811115612da357fe5b90816010811115612db057fe5b905250600081516010811115612dc257fe5b14612e14576040805162461bcd60e51b815260206004820152601f60248201527f726570617920626f72726f77207472616e7366657220696e206661696c656400604482015290519081900360640190fd5b60a080820180516001600160a01b03808916600081815260116020908152604091829020948555600b5460019095019490945560c0870151600c8190558188015195518251948e16855294840192909252828101949094526060820192909252608081019190915290517f1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1929181900390910190a160065460408083015160608401518251631ededc9160e01b81523060048201526001600160a01b038b811660248301528a81166044830152606482019390935260848101919091529151921691631ededc919160a48082019260009290919082900301818387803b158015612f1d57600080fd5b505af1158015612f31573d6000803e3d6000fd5b5060009250612f3e915050565b9695505050505050565b600080600080612f588787612604565b90925090506000826003811115612f6b57fe5b14612f7c575091506000905061267e565b6126778186612579565b6000612f906145e9565b600080612fa586670de0b6b3a764000061360a565b90925090506000826003811115612fb857fe5b14612fd757506040805160208101909152600081529092509050612178565b600080612fe48388613649565b90925090506000826003811115612ff757fe5b1461301957506040805160208101909152600081529094509250612178915050565b604080516020810190915290815260009890975095505050505050565b51670de0b6b3a7640000900490565b6012546040805163a9059cbb60e01b81526001600160a01b03858116600483015260248201859052915160009392909216918391839163a9059cbb91604480820192869290919082900301818387803b1580156130a157600080fd5b505af11580156130b5573d6000803e3d6000fd5b505050503d600081146130cf57602081146130d957600080fd5b60001991506130e5565b60206000803e60005191505b50806130f65760109250505061086f565b506000949350505050565b600082158061310e575081155b6131495760405162461bcd60e51b81526004018080602001828103825260348152602001806148c56034913960400191505060405180910390fd5b613151614656565b613159611ce5565b604083018190526020830182600381111561317057fe5b600381111561317b57fe5b905250600090508160200151600381111561319257fe5b146131ae57612b9b6009602b8360200151600381111561103d57fe5b831561322f5760608101849052604080516020810182529082015181526131d5908561212b565b60808301819052602083018260038111156131ec57fe5b60038111156131f757fe5b905250600090508160200151600381111561320e57fe5b1461322a57612b9b600960298360200151600381111561103d57fe5b6132a8565b61324b83604051806020016040528084604001518152506145d2565b606083018190526020830182600381111561326257fe5b600381111561326d57fe5b905250600090508160200151600381111561328457fe5b146132a057612b9b6009602a8360200151600381111561103d57fe5b608081018390525b60065460608201516040805163eabe7d9160e01b81523060048201526001600160a01b03898116602483015260448201939093529051600093929092169163eabe7d919160648082019260209290919082900301818787803b15801561330d57600080fd5b505af1158015613321573d6000803e3d6000fd5b505050506040513d602081101561333757600080fd5b50519050801561334e57612c46600360288361250f565b613356612575565b600a541461336a57612c46600a602c6121ff565b61337a600e548360600151612579565b60a084018190526020840182600381111561339157fe5b600381111561339c57fe5b90525060009050826020015160038111156133b357fe5b146133cf57612c466009602e8460200151600381111561103d57fe5b6001600160a01b0386166000908152600f602052604090205460608301516133f79190612579565b60c084018190526020840182600381111561340e57fe5b600381111561341957fe5b905250600090508260200151600381111561343057fe5b1461344c57612c466009602d8460200151600381111561103d57fe5b816080015161345961217f565b101561346b57612c46600e602f6121ff565b613479868360800151613045565b8290601081111561348657fe5b9081601081111561349357fe5b9052506000825160108111156134a557fe5b146134f7576040805162461bcd60e51b815260206004820152601a60248201527f72656465656d207472616e73666572206f7574206661696c6564000000000000604482015290519081900360640190fd5b60a0820151600e5560c08201516001600160a01b0387166000818152600f6020908152604091829020939093556060850151815190815290513093600080516020614834833981519152928290030190a36080820151606080840151604080516001600160a01b038b168152602081019490945283810191909152517fe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a9299281900390910190a160065460808301516060840151604080516351dff98960e01b81523060048201526001600160a01b038b81166024830152604482019490945260648101929092525191909216916351dff98991608480830192600092919082900301818387803b158015612f1d57600080fd5b6000808361361d57506000905080612178565b8383028385828161362a57fe5b041461363e57506002915060009050612178565b600092509050612178565b6000808261365d5750600190506000612178565b600083858161366857fe5b04915091509250929050565b60065460408051634ef4c3e160e01b81523060048201526001600160a01b03858116602483015260448201859052915160009384931691634ef4c3e191606480830192602092919082900301818787803b1580156136d157600080fd5b505af11580156136e5573d6000803e3d6000fd5b505050506040513d60208110156136fb57600080fd5b50519050801561371a576137126003601f8361250f565b91505061086f565b613722612575565b600a541461373657613712600a60226121ff565b61373e614694565b61374885856143e7565b8190601081111561375557fe5b9081601081111561376257fe5b90525060008151601081111561377457fe5b1461378f5780516137869060266121ff565b9250505061086f565b613797611ce5565b60408301819052602083018260038111156137ae57fe5b60038111156137b957fe5b90525060009050816020015160038111156137d057fe5b146137ec57613786600960218360200151600381111561103d57fe5b61380884604051806020016040528084604001518152506145d2565b606083018190526020830182600381111561381f57fe5b600381111561382a57fe5b905250600090508160200151600381111561384157fe5b1461385d57613786600960208360200151600381111561103d57fe5b61386d600e548260600151612604565b608083018190526020830182600381111561388457fe5b600381111561388f57fe5b90525060009050816020015160038111156138a657fe5b146138c257613786600960248360200151600381111561103d57fe5b6001600160a01b0385166000908152600f602052604090205460608201516138ea9190612604565b60a083018190526020830182600381111561390157fe5b600381111561390c57fe5b905250600090508160200151600381111561392357fe5b1461393f57613786600960238360200151600381111561103d57fe5b613949858561451b565b8190601081111561395657fe5b9081601081111561396357fe5b90525060008151601081111561397557fe5b146139875780516137869060256121ff565b6080810151600e5560a08101516001600160a01b0386166000818152600f602090815260409182902093909355606080850151825193845293830188905282820193909352517f4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f929181900390910190a1606081015160408051918252516001600160a01b0387169130916000805160206148348339815191529181900360200190a36006546060820151604080516341c728b960e01b81523060048201526001600160a01b038981166024830152604482018990526064820193909352905191909216916341c728b991608480830192600092919082900301818387803b158015613a9257600080fd5b505af1158015613aa6573d6000803e3d6000fd5b5060009250613ab3915050565b95945050505050565b6000613ac66145e9565b600080613adb8660000151866000015161360a565b90925090506000826003811115613aee57fe5b14613b0d57506040805160208101909152600081529092509050612178565b600080613b226706f05b59d3b2000084612604565b90925090506000826003811115613b3557fe5b14613b5757506040805160208101909152600081529094509250612178915050565b600080613b6c83670de0b6b3a7640000613649565b90925090506000826003811115613b7f57fe5b14613b8657fe5b604080516020810190915290815260009a909950975050505050505050565b6006546040805163368f515360e21b81523060048201526001600160a01b0385811660248301526044820185905291516000938493169163da3d454c91606480830192602092919082900301818787803b158015613c0257600080fd5b505af1158015613c16573d6000803e3d6000fd5b505050506040513d6020811015613c2c57600080fd5b505190508015613c43576137126003600e8361250f565b613c4b612575565b600a5414613c5e57613712600a806121ff565b82613c6761217f565b1015613c7957613712600e60096121ff565b613c816146ae565b613c8a85612420565b6040830181905260208301826003811115613ca157fe5b6003811115613cac57fe5b9052506000905081602001516003811115613cc357fe5b14613cdf57613786600960078360200151600381111561103d57fe5b613ced816040015185612604565b6060830181905260208301826003811115613d0457fe5b6003811115613d0f57fe5b9052506000905081602001516003811115613d2657fe5b14613d42576137866009600c8360200151600381111561103d57fe5b613d4e600c5485612604565b6080830181905260208301826003811115613d6557fe5b6003811115613d7057fe5b9052506000905081602001516003811115613d8757fe5b14613da3576137866009600b8360200151600381111561103d57fe5b613dad8585613045565b81906010811115613dba57fe5b90816010811115613dc757fe5b905250600081516010811115613dd957fe5b14613e2b576040805162461bcd60e51b815260206004820152601a60248201527f626f72726f77207472616e73666572206f7574206661696c6564000000000000604482015290519081900360640190fd5b606080820180516001600160a01b038816600081815260116020908152604091829020938455600b54600190940193909355608080870151600c819055945182519384529383018a9052828201939093529381019290925291517f13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80929181900390910190a160065460408051635c77860560e01b81523060048201526001600160a01b0388811660248301526044820188905291519190921691635c77860591606480830192600092919082900301818387803b158015613a9257600080fd5b60065460408051632fe3f38f60e11b81523060048201526001600160a01b0384811660248301528781166044830152868116606483015260848201869052915160009384931691635fc7e71e9160a480830192602092919082900301818787803b158015613f7857600080fd5b505af1158015613f8c573d6000803e3d6000fd5b505050506040513d6020811015613fa257600080fd5b505190508015613fb957611e39600360128361250f565b613fc1612575565b600a5414613fd557611e39600a60166121ff565b613fdd612575565b836001600160a01b0316636c540baf6040518163ffffffff1660e01b815260040160206040518083038186803b15801561401657600080fd5b505afa15801561402a573d6000803e3d6000fd5b505050506040513d602081101561404057600080fd5b50511461405357611e39600a60116121ff565b856001600160a01b0316856001600160a01b0316141561407957611e39600660176121ff565b8361408a57611e39600760156121ff565b6000198414156140a057611e39600760146121ff565b6006546040805163c488847b60e01b81523060048201526001600160a01b038681166024830152604482018890528251600094859492169263c488847b926064808301939192829003018186803b1580156140fa57600080fd5b505afa15801561410e573d6000803e3d6000fd5b505050506040513d604081101561412457600080fd5b5080516020909101519092509050811561414f57614145600460138461250f565b9350505050610a9e565b846001600160a01b03166370a08231886040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156141a557600080fd5b505afa1580156141b9573d6000803e3d6000fd5b505050506040513d60208110156141cf57600080fd5b50518111156141e457614145600d601d6121ff565b60006141f1898989612af5565b9050801561421a5761420f81601081111561420857fe5b60186121ff565b945050505050610a9e565b6040805163b2a02ff160e01b81526001600160a01b038b811660048301528a8116602483015260448201859052915160009289169163b2a02ff191606480830192602092919082900301818787803b15801561427557600080fd5b505af1158015614289573d6000803e3d6000fd5b505050506040513d602081101561429f57600080fd5b5051905080156142ed576040805162461bcd60e51b81526020600482015260146024820152731d1bdad95b881cd95a5e9d5c994819985a5b195960621b604482015290519081900360640190fd5b604080516001600160a01b03808d168252808c1660208301528183018b9052891660608201526080810185905290517f298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb529181900360a00190a1600654604080516347ef3b3b60e01b81523060048201526001600160a01b038a811660248301528d811660448301528c81166064830152608482018c905260a48201879052915191909216916347ef3b3b9160c480830192600092919082900301818387803b1580156143b857600080fd5b505af11580156143cc573d6000803e3d6000fd5b50600092506143d9915050565b9a9950505050505050505050565b60125460408051636eb1769f60e11b81526001600160a01b038581166004830152306024830152915160009392909216918491839163dd62ed3e91604480820192602092909190829003018186803b15801561444257600080fd5b505afa158015614456573d6000803e3d6000fd5b505050506040513d602081101561446c57600080fd5b5051101561447e57600c91505061086f565b82816001600160a01b03166370a08231866040518263ffffffff1660e01b815260040180826001600160a01b03166001600160a01b0316815260200191505060206040518083038186803b1580156144d557600080fd5b505afa1580156144e9573d6000803e3d6000fd5b505050506040513d60208110156144ff57600080fd5b5051101561451157600d91505061086f565b5060009392505050565b601254604080516323b872dd60e01b81526001600160a01b0385811660048301523060248301526044820185905291516000939290921691839183916323b872dd91606480820192869290919082900301818387803b15801561457d57600080fd5b505af1158015614591573d6000803e3d6000fd5b505050503d600081146145ab57602081146145b557600080fd5b60001991506145c1565b60206000803e60005191505b50806130f657600f9250505061086f565b60008060006145df6145e9565b6121428686612686565b6040518060200160405280600081525090565b6040805161014081019091528060008152602001600081526020016000815260200160008152602001600081526020016146346145e9565b8152602001600081526020016000815260200160008152602001600081525090565b6040805160e0810190915280600081526020016000815260200160008152602001600081526020016000815260200160008152602001600081525090565b6040805160c0810190915280600081526020016000614634565b6040805160a08101909152806000815260200160008152602001600081526020016000815260200160008152509056fe737570706c7952617465506572426c6f636b3a2063616c63756c6174696e6720626f72726f7773506572206661696c6564726564756365207265736572766573207472616e73666572206f7574206661696c6564737570706c7952617465506572426c6f636b3a2063616c63756c6174696e6720737570706c7952617465206661696c6564626f72726f7742616c616e636553746f7265643a20626f72726f7742616c616e636553746f726564496e7465726e616c206661696c6564737570706c7952617465506572426c6f636b3a2063616c63756c6174696e6720756e6465726c79696e67206661696c6564626f72726f7752617465506572426c6f636b3a20696e746572657374526174654d6f64656c2e626f72726f7752617465206661696c6564737570706c7952617465506572426c6f636b3a2063616c63756c6174696e6720626f72726f7752617465206661696c6564ddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef737570706c7952617465506572426c6f636b3a2063616c63756c6174696e67206f6e654d696e757352657365727665466163746f72206661696c656465786368616e67655261746553746f7265643a2065786368616e67655261746553746f726564496e7465726e616c206661696c65646f6e65206f662072656465656d546f6b656e73496e206f722072656465656d416d6f756e74496e206d757374206265207a65726f72656475636520726573657276657320756e657870656374656420756e646572666c6f77a265627a7a7231582024fab8635706da943a105c2409c5f0bc9267f0e4f1072b04135ee9c87521d40c64736f6c6343000511003253657474696e6720696e7465726573742072617465206d6f64656c206661696c6564496e697469616c2065786368616e67652072617465206d7573742062652067726561746572207468616e207a65726f2e"

// DeployCErc20 deploys a new Ethereum contract, binding an instance of CErc20 to it.
func DeployCErc20(auth *bind.TransactOpts, backend bind.ContractBackend, underlying_ common.Address, comptroller_ common.Address, interestRateModel_ common.Address, initialExchangeRateMantissa_ *big.Int, name_ string, symbol_ string, decimals_ *big.Int) (common.Address, *types.Transaction, *CErc20, error) {
	parsed, err := abi.JSON(strings.NewReader(CErc20ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CErc20Bin), backend, underlying_, comptroller_, interestRateModel_, initialExchangeRateMantissa_, name_, symbol_, decimals_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CErc20{CErc20Caller: CErc20Caller{contract: contract}, CErc20Transactor: CErc20Transactor{contract: contract}, CErc20Filterer: CErc20Filterer{contract: contract}}, nil
}

// CErc20 is an auto generated Go binding around an Ethereum contract.
type CErc20 struct {
	CErc20Caller     // Read-only binding to the contract
	CErc20Transactor // Write-only binding to the contract
	CErc20Filterer   // Log filterer for contract events
}

// CErc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type CErc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type CErc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CErc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CErc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CErc20Session struct {
	Contract     *CErc20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CErc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CErc20CallerSession struct {
	Contract *CErc20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CErc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CErc20TransactorSession struct {
	Contract     *CErc20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CErc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type CErc20Raw struct {
	Contract *CErc20 // Generic contract binding to access the raw methods on
}

// CErc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CErc20CallerRaw struct {
	Contract *CErc20Caller // Generic read-only contract binding to access the raw methods on
}

// CErc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CErc20TransactorRaw struct {
	Contract *CErc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewCErc20 creates a new instance of CErc20, bound to a specific deployed contract.
func NewCErc20(address common.Address, backend bind.ContractBackend) (*CErc20, error) {
	contract, err := bindCErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CErc20{CErc20Caller: CErc20Caller{contract: contract}, CErc20Transactor: CErc20Transactor{contract: contract}, CErc20Filterer: CErc20Filterer{contract: contract}}, nil
}

// NewCErc20Caller creates a new read-only instance of CErc20, bound to a specific deployed contract.
func NewCErc20Caller(address common.Address, caller bind.ContractCaller) (*CErc20Caller, error) {
	contract, err := bindCErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CErc20Caller{contract: contract}, nil
}

// NewCErc20Transactor creates a new write-only instance of CErc20, bound to a specific deployed contract.
func NewCErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*CErc20Transactor, error) {
	contract, err := bindCErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CErc20Transactor{contract: contract}, nil
}

// NewCErc20Filterer creates a new log filterer instance of CErc20, bound to a specific deployed contract.
func NewCErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*CErc20Filterer, error) {
	contract, err := bindCErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CErc20Filterer{contract: contract}, nil
}

// bindCErc20 binds a generic wrapper to an already deployed contract.
func bindCErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CErc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CErc20 *CErc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CErc20.Contract.CErc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CErc20 *CErc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.Contract.CErc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CErc20 *CErc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CErc20.Contract.CErc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CErc20 *CErc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CErc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CErc20 *CErc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CErc20 *CErc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CErc20.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CErc20 *CErc20Caller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CErc20 *CErc20Session) AccrualBlockNumber() (*big.Int, error) {
	return _CErc20.Contract.AccrualBlockNumber(&_CErc20.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CErc20 *CErc20CallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _CErc20.Contract.AccrualBlockNumber(&_CErc20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CErc20 *CErc20Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CErc20 *CErc20Session) Admin() (common.Address, error) {
	return _CErc20.Contract.Admin(&_CErc20.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CErc20 *CErc20CallerSession) Admin() (common.Address, error) {
	return _CErc20.Contract.Admin(&_CErc20.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CErc20 *CErc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CErc20 *CErc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CErc20.Contract.Allowance(&_CErc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CErc20 *CErc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CErc20.Contract.Allowance(&_CErc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CErc20 *CErc20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CErc20 *CErc20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CErc20.Contract.BalanceOf(&_CErc20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CErc20 *CErc20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CErc20.Contract.BalanceOf(&_CErc20.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CErc20 *CErc20Caller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CErc20 *CErc20Session) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CErc20.Contract.BorrowBalanceStored(&_CErc20.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CErc20 *CErc20CallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CErc20.Contract.BorrowBalanceStored(&_CErc20.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CErc20 *CErc20Caller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CErc20 *CErc20Session) BorrowIndex() (*big.Int, error) {
	return _CErc20.Contract.BorrowIndex(&_CErc20.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CErc20 *CErc20CallerSession) BorrowIndex() (*big.Int, error) {
	return _CErc20.Contract.BorrowIndex(&_CErc20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Caller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Session) BorrowRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.BorrowRatePerBlock(&_CErc20.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20CallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.BorrowRatePerBlock(&_CErc20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CErc20 *CErc20Caller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CErc20 *CErc20Session) Comptroller() (common.Address, error) {
	return _CErc20.Contract.Comptroller(&_CErc20.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CErc20 *CErc20CallerSession) Comptroller() (common.Address, error) {
	return _CErc20.Contract.Comptroller(&_CErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CErc20 *CErc20Caller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CErc20 *CErc20Session) Decimals() (*big.Int, error) {
	return _CErc20.Contract.Decimals(&_CErc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CErc20 *CErc20CallerSession) Decimals() (*big.Int, error) {
	return _CErc20.Contract.Decimals(&_CErc20.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CErc20 *CErc20Caller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CErc20 *CErc20Session) ExchangeRateStored() (*big.Int, error) {
	return _CErc20.Contract.ExchangeRateStored(&_CErc20.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CErc20 *CErc20CallerSession) ExchangeRateStored() (*big.Int, error) {
	return _CErc20.Contract.ExchangeRateStored(&_CErc20.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CErc20 *CErc20Caller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CErc20 *CErc20Session) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CErc20.Contract.GetAccountSnapshot(&_CErc20.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CErc20 *CErc20CallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CErc20.Contract.GetAccountSnapshot(&_CErc20.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CErc20 *CErc20Caller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CErc20 *CErc20Session) GetCash() (*big.Int, error) {
	return _CErc20.Contract.GetCash(&_CErc20.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CErc20 *CErc20CallerSession) GetCash() (*big.Int, error) {
	return _CErc20.Contract.GetCash(&_CErc20.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CErc20 *CErc20Caller) InitialExchangeRateMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "initialExchangeRateMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CErc20 *CErc20Session) InitialExchangeRateMantissa() (*big.Int, error) {
	return _CErc20.Contract.InitialExchangeRateMantissa(&_CErc20.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CErc20 *CErc20CallerSession) InitialExchangeRateMantissa() (*big.Int, error) {
	return _CErc20.Contract.InitialExchangeRateMantissa(&_CErc20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CErc20 *CErc20Caller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CErc20 *CErc20Session) InterestRateModel() (common.Address, error) {
	return _CErc20.Contract.InterestRateModel(&_CErc20.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CErc20 *CErc20CallerSession) InterestRateModel() (common.Address, error) {
	return _CErc20.Contract.InterestRateModel(&_CErc20.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CErc20 *CErc20Caller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CErc20 *CErc20Session) IsCToken() (bool, error) {
	return _CErc20.Contract.IsCToken(&_CErc20.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CErc20 *CErc20CallerSession) IsCToken() (bool, error) {
	return _CErc20.Contract.IsCToken(&_CErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CErc20 *CErc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CErc20 *CErc20Session) Name() (string, error) {
	return _CErc20.Contract.Name(&_CErc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CErc20 *CErc20CallerSession) Name() (string, error) {
	return _CErc20.Contract.Name(&_CErc20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CErc20 *CErc20Caller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CErc20 *CErc20Session) PendingAdmin() (common.Address, error) {
	return _CErc20.Contract.PendingAdmin(&_CErc20.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CErc20 *CErc20CallerSession) PendingAdmin() (common.Address, error) {
	return _CErc20.Contract.PendingAdmin(&_CErc20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CErc20 *CErc20Caller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CErc20 *CErc20Session) ReserveFactorMantissa() (*big.Int, error) {
	return _CErc20.Contract.ReserveFactorMantissa(&_CErc20.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CErc20 *CErc20CallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _CErc20.Contract.ReserveFactorMantissa(&_CErc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Caller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20Session) SupplyRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.SupplyRatePerBlock(&_CErc20.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CErc20 *CErc20CallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CErc20.Contract.SupplyRatePerBlock(&_CErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CErc20 *CErc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CErc20 *CErc20Session) Symbol() (string, error) {
	return _CErc20.Contract.Symbol(&_CErc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CErc20 *CErc20CallerSession) Symbol() (string, error) {
	return _CErc20.Contract.Symbol(&_CErc20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CErc20 *CErc20Caller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CErc20 *CErc20Session) TotalBorrows() (*big.Int, error) {
	return _CErc20.Contract.TotalBorrows(&_CErc20.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CErc20 *CErc20CallerSession) TotalBorrows() (*big.Int, error) {
	return _CErc20.Contract.TotalBorrows(&_CErc20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CErc20 *CErc20Caller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CErc20 *CErc20Session) TotalReserves() (*big.Int, error) {
	return _CErc20.Contract.TotalReserves(&_CErc20.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CErc20 *CErc20CallerSession) TotalReserves() (*big.Int, error) {
	return _CErc20.Contract.TotalReserves(&_CErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CErc20 *CErc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CErc20 *CErc20Session) TotalSupply() (*big.Int, error) {
	return _CErc20.Contract.TotalSupply(&_CErc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CErc20 *CErc20CallerSession) TotalSupply() (*big.Int, error) {
	return _CErc20.Contract.TotalSupply(&_CErc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20Caller) Underlying(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CErc20.contract.Call(opts, &out, "underlying")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20Session) Underlying() (common.Address, error) {
	return _CErc20.Contract.Underlying(&_CErc20.CallOpts)
}

// Underlying is a free data retrieval call binding the contract method 0x6f307dc3.
//
// Solidity: function underlying() view returns(address)
func (_CErc20 *CErc20CallerSession) Underlying() (common.Address, error) {
	return _CErc20.Contract.Underlying(&_CErc20.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CErc20 *CErc20Transactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CErc20 *CErc20Session) AcceptAdmin() (*types.Transaction, error) {
	return _CErc20.Contract.AcceptAdmin(&_CErc20.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CErc20 *CErc20TransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _CErc20.Contract.AcceptAdmin(&_CErc20.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CErc20 *CErc20Session) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.ReduceReserves(&_CErc20.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.ReduceReserves(&_CErc20.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CErc20 *CErc20Transactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CErc20 *CErc20Session) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetComptroller(&_CErc20.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CErc20 *CErc20TransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetComptroller(&_CErc20.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CErc20 *CErc20Transactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CErc20 *CErc20Session) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetInterestRateModel(&_CErc20.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CErc20 *CErc20TransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetInterestRateModel(&_CErc20.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CErc20 *CErc20Transactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CErc20 *CErc20Session) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetPendingAdmin(&_CErc20.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CErc20 *CErc20TransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.SetPendingAdmin(&_CErc20.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CErc20 *CErc20Transactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CErc20 *CErc20Session) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.SetReserveFactor(&_CErc20.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CErc20 *CErc20TransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.SetReserveFactor(&_CErc20.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CErc20 *CErc20Transactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CErc20 *CErc20Session) AccrueInterest() (*types.Transaction, error) {
	return _CErc20.Contract.AccrueInterest(&_CErc20.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CErc20 *CErc20TransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _CErc20.Contract.AccrueInterest(&_CErc20.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CErc20 *CErc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CErc20 *CErc20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Approve(&_CErc20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CErc20 *CErc20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Approve(&_CErc20.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CErc20 *CErc20Transactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CErc20 *CErc20Session) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.BalanceOfUnderlying(&_CErc20.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CErc20 *CErc20TransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.BalanceOfUnderlying(&_CErc20.TransactOpts, owner)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) Borrow(opts *bind.TransactOpts, borrowAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "borrow", borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CErc20 *CErc20Session) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Borrow(&_CErc20.TransactOpts, borrowAmount)
}

// Borrow is a paid mutator transaction binding the contract method 0xc5ebeaec.
//
// Solidity: function borrow(uint256 borrowAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) Borrow(borrowAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Borrow(&_CErc20.TransactOpts, borrowAmount)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CErc20 *CErc20Transactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CErc20 *CErc20Session) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.BorrowBalanceCurrent(&_CErc20.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CErc20 *CErc20TransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.BorrowBalanceCurrent(&_CErc20.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CErc20 *CErc20Transactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CErc20 *CErc20Session) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CErc20.Contract.ExchangeRateCurrent(&_CErc20.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CErc20 *CErc20TransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CErc20.Contract.ExchangeRateCurrent(&_CErc20.TransactOpts)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CErc20 *CErc20Transactor) LiquidateBorrow(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "liquidateBorrow", borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CErc20 *CErc20Session) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.LiquidateBorrow(&_CErc20.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// LiquidateBorrow is a paid mutator transaction binding the contract method 0xf5e3c462.
//
// Solidity: function liquidateBorrow(address borrower, uint256 repayAmount, address cTokenCollateral) returns(uint256)
func (_CErc20 *CErc20TransactorSession) LiquidateBorrow(borrower common.Address, repayAmount *big.Int, cTokenCollateral common.Address) (*types.Transaction, error) {
	return _CErc20.Contract.LiquidateBorrow(&_CErc20.TransactOpts, borrower, repayAmount, cTokenCollateral)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) Mint(opts *bind.TransactOpts, mintAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "mint", mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CErc20 *CErc20Session) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Mint(&_CErc20.TransactOpts, mintAmount)
}

// Mint is a paid mutator transaction binding the contract method 0xa0712d68.
//
// Solidity: function mint(uint256 mintAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) Mint(mintAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Mint(&_CErc20.TransactOpts, mintAmount)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CErc20 *CErc20Transactor) Redeem(opts *bind.TransactOpts, redeemTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "redeem", redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CErc20 *CErc20Session) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Redeem(&_CErc20.TransactOpts, redeemTokens)
}

// Redeem is a paid mutator transaction binding the contract method 0xdb006a75.
//
// Solidity: function redeem(uint256 redeemTokens) returns(uint256)
func (_CErc20 *CErc20TransactorSession) Redeem(redeemTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Redeem(&_CErc20.TransactOpts, redeemTokens)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) RedeemUnderlying(opts *bind.TransactOpts, redeemAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "redeemUnderlying", redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CErc20 *CErc20Session) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlying(&_CErc20.TransactOpts, redeemAmount)
}

// RedeemUnderlying is a paid mutator transaction binding the contract method 0x852a12e3.
//
// Solidity: function redeemUnderlying(uint256 redeemAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) RedeemUnderlying(redeemAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RedeemUnderlying(&_CErc20.TransactOpts, redeemAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) RepayBorrow(opts *bind.TransactOpts, repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "repayBorrow", repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20Session) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RepayBorrow(&_CErc20.TransactOpts, repayAmount)
}

// RepayBorrow is a paid mutator transaction binding the contract method 0x0e752702.
//
// Solidity: function repayBorrow(uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) RepayBorrow(repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RepayBorrow(&_CErc20.TransactOpts, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20Transactor) RepayBorrowBehalf(opts *bind.TransactOpts, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "repayBorrowBehalf", borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20Session) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RepayBorrowBehalf(&_CErc20.TransactOpts, borrower, repayAmount)
}

// RepayBorrowBehalf is a paid mutator transaction binding the contract method 0x2608f818.
//
// Solidity: function repayBorrowBehalf(address borrower, uint256 repayAmount) returns(uint256)
func (_CErc20 *CErc20TransactorSession) RepayBorrowBehalf(borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.RepayBorrowBehalf(&_CErc20.TransactOpts, borrower, repayAmount)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CErc20 *CErc20Transactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CErc20 *CErc20Session) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Seize(&_CErc20.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CErc20 *CErc20TransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Seize(&_CErc20.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CErc20 *CErc20Transactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CErc20 *CErc20Session) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CErc20.Contract.TotalBorrowsCurrent(&_CErc20.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CErc20 *CErc20TransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CErc20.Contract.TotalBorrowsCurrent(&_CErc20.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20Transactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20Session) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Transfer(&_CErc20.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20TransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.Transfer(&_CErc20.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20Transactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20Session) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.TransferFrom(&_CErc20.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CErc20 *CErc20TransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CErc20.Contract.TransferFrom(&_CErc20.TransactOpts, src, dst, amount)
}

// CErc20AccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the CErc20 contract.
type CErc20AccrueInterestIterator struct {
	Event *CErc20AccrueInterest // Event containing the contract specifics and raw log

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
func (it *CErc20AccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20AccrueInterest)
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
		it.Event = new(CErc20AccrueInterest)
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
func (it *CErc20AccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20AccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20AccrueInterest represents a AccrueInterest event raised by the CErc20 contract.
type CErc20AccrueInterest struct {
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) FilterAccrueInterest(opts *bind.FilterOpts) (*CErc20AccrueInterestIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &CErc20AccrueInterestIterator{contract: _CErc20.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *CErc20AccrueInterest) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20AccrueInterest)
				if err := _CErc20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) ParseAccrueInterest(log types.Log) (*CErc20AccrueInterest, error) {
	event := new(CErc20AccrueInterest)
	if err := _CErc20.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CErc20 contract.
type CErc20ApprovalIterator struct {
	Event *CErc20Approval // Event containing the contract specifics and raw log

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
func (it *CErc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Approval)
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
		it.Event = new(CErc20Approval)
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
func (it *CErc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Approval represents a Approval event raised by the CErc20 contract.
type CErc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CErc20 *CErc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CErc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CErc20ApprovalIterator{contract: _CErc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CErc20 *CErc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CErc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Approval)
				if err := _CErc20.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CErc20 *CErc20Filterer) ParseApproval(log types.Log) (*CErc20Approval, error) {
	event := new(CErc20Approval)
	if err := _CErc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the CErc20 contract.
type CErc20BorrowIterator struct {
	Event *CErc20Borrow // Event containing the contract specifics and raw log

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
func (it *CErc20BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Borrow)
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
		it.Event = new(CErc20Borrow)
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
func (it *CErc20BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Borrow represents a Borrow event raised by the CErc20 contract.
type CErc20Borrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) FilterBorrow(opts *bind.FilterOpts) (*CErc20BorrowIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &CErc20BorrowIterator{contract: _CErc20.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *CErc20Borrow) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Borrow)
				if err := _CErc20.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) ParseBorrow(log types.Log) (*CErc20Borrow, error) {
	event := new(CErc20Borrow)
	if err := _CErc20.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20FailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the CErc20 contract.
type CErc20FailureIterator struct {
	Event *CErc20Failure // Event containing the contract specifics and raw log

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
func (it *CErc20FailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Failure)
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
		it.Event = new(CErc20Failure)
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
func (it *CErc20FailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20FailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Failure represents a Failure event raised by the CErc20 contract.
type CErc20Failure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CErc20 *CErc20Filterer) FilterFailure(opts *bind.FilterOpts) (*CErc20FailureIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &CErc20FailureIterator{contract: _CErc20.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CErc20 *CErc20Filterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *CErc20Failure) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Failure)
				if err := _CErc20.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CErc20 *CErc20Filterer) ParseFailure(log types.Log) (*CErc20Failure, error) {
	event := new(CErc20Failure)
	if err := _CErc20.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20LiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the CErc20 contract.
type CErc20LiquidateBorrowIterator struct {
	Event *CErc20LiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *CErc20LiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20LiquidateBorrow)
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
		it.Event = new(CErc20LiquidateBorrow)
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
func (it *CErc20LiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20LiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20LiquidateBorrow represents a LiquidateBorrow event raised by the CErc20 contract.
type CErc20LiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CErc20 *CErc20Filterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*CErc20LiquidateBorrowIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &CErc20LiquidateBorrowIterator{contract: _CErc20.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CErc20 *CErc20Filterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *CErc20LiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20LiquidateBorrow)
				if err := _CErc20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CErc20 *CErc20Filterer) ParseLiquidateBorrow(log types.Log) (*CErc20LiquidateBorrow, error) {
	event := new(CErc20LiquidateBorrow)
	if err := _CErc20.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CErc20 contract.
type CErc20MintIterator struct {
	Event *CErc20Mint // Event containing the contract specifics and raw log

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
func (it *CErc20MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Mint)
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
		it.Event = new(CErc20Mint)
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
func (it *CErc20MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Mint represents a Mint event raised by the CErc20 contract.
type CErc20Mint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CErc20 *CErc20Filterer) FilterMint(opts *bind.FilterOpts) (*CErc20MintIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &CErc20MintIterator{contract: _CErc20.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CErc20 *CErc20Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CErc20Mint) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Mint)
				if err := _CErc20.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CErc20 *CErc20Filterer) ParseMint(log types.Log) (*CErc20Mint, error) {
	event := new(CErc20Mint)
	if err := _CErc20.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20NewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CErc20 contract.
type CErc20NewAdminIterator struct {
	Event *CErc20NewAdmin // Event containing the contract specifics and raw log

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
func (it *CErc20NewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20NewAdmin)
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
		it.Event = new(CErc20NewAdmin)
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
func (it *CErc20NewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20NewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20NewAdmin represents a NewAdmin event raised by the CErc20 contract.
type CErc20NewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CErc20 *CErc20Filterer) FilterNewAdmin(opts *bind.FilterOpts) (*CErc20NewAdminIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &CErc20NewAdminIterator{contract: _CErc20.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CErc20 *CErc20Filterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CErc20NewAdmin) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20NewAdmin)
				if err := _CErc20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CErc20 *CErc20Filterer) ParseNewAdmin(log types.Log) (*CErc20NewAdmin, error) {
	event := new(CErc20NewAdmin)
	if err := _CErc20.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20NewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the CErc20 contract.
type CErc20NewComptrollerIterator struct {
	Event *CErc20NewComptroller // Event containing the contract specifics and raw log

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
func (it *CErc20NewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20NewComptroller)
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
		it.Event = new(CErc20NewComptroller)
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
func (it *CErc20NewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20NewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20NewComptroller represents a NewComptroller event raised by the CErc20 contract.
type CErc20NewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CErc20 *CErc20Filterer) FilterNewComptroller(opts *bind.FilterOpts) (*CErc20NewComptrollerIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &CErc20NewComptrollerIterator{contract: _CErc20.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CErc20 *CErc20Filterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *CErc20NewComptroller) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20NewComptroller)
				if err := _CErc20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CErc20 *CErc20Filterer) ParseNewComptroller(log types.Log) (*CErc20NewComptroller, error) {
	event := new(CErc20NewComptroller)
	if err := _CErc20.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20NewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the CErc20 contract.
type CErc20NewMarketInterestRateModelIterator struct {
	Event *CErc20NewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *CErc20NewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20NewMarketInterestRateModel)
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
		it.Event = new(CErc20NewMarketInterestRateModel)
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
func (it *CErc20NewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20NewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20NewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the CErc20 contract.
type CErc20NewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CErc20 *CErc20Filterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*CErc20NewMarketInterestRateModelIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &CErc20NewMarketInterestRateModelIterator{contract: _CErc20.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CErc20 *CErc20Filterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *CErc20NewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20NewMarketInterestRateModel)
				if err := _CErc20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CErc20 *CErc20Filterer) ParseNewMarketInterestRateModel(log types.Log) (*CErc20NewMarketInterestRateModel, error) {
	event := new(CErc20NewMarketInterestRateModel)
	if err := _CErc20.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20NewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the CErc20 contract.
type CErc20NewPendingAdminIterator struct {
	Event *CErc20NewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *CErc20NewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20NewPendingAdmin)
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
		it.Event = new(CErc20NewPendingAdmin)
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
func (it *CErc20NewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20NewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20NewPendingAdmin represents a NewPendingAdmin event raised by the CErc20 contract.
type CErc20NewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CErc20 *CErc20Filterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*CErc20NewPendingAdminIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &CErc20NewPendingAdminIterator{contract: _CErc20.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CErc20 *CErc20Filterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *CErc20NewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20NewPendingAdmin)
				if err := _CErc20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CErc20 *CErc20Filterer) ParseNewPendingAdmin(log types.Log) (*CErc20NewPendingAdmin, error) {
	event := new(CErc20NewPendingAdmin)
	if err := _CErc20.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20NewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the CErc20 contract.
type CErc20NewReserveFactorIterator struct {
	Event *CErc20NewReserveFactor // Event containing the contract specifics and raw log

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
func (it *CErc20NewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20NewReserveFactor)
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
		it.Event = new(CErc20NewReserveFactor)
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
func (it *CErc20NewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20NewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20NewReserveFactor represents a NewReserveFactor event raised by the CErc20 contract.
type CErc20NewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CErc20 *CErc20Filterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*CErc20NewReserveFactorIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &CErc20NewReserveFactorIterator{contract: _CErc20.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CErc20 *CErc20Filterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *CErc20NewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20NewReserveFactor)
				if err := _CErc20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CErc20 *CErc20Filterer) ParseNewReserveFactor(log types.Log) (*CErc20NewReserveFactor, error) {
	event := new(CErc20NewReserveFactor)
	if err := _CErc20.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20RedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the CErc20 contract.
type CErc20RedeemIterator struct {
	Event *CErc20Redeem // Event containing the contract specifics and raw log

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
func (it *CErc20RedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Redeem)
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
		it.Event = new(CErc20Redeem)
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
func (it *CErc20RedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20RedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Redeem represents a Redeem event raised by the CErc20 contract.
type CErc20Redeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CErc20 *CErc20Filterer) FilterRedeem(opts *bind.FilterOpts) (*CErc20RedeemIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &CErc20RedeemIterator{contract: _CErc20.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CErc20 *CErc20Filterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *CErc20Redeem) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Redeem)
				if err := _CErc20.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CErc20 *CErc20Filterer) ParseRedeem(log types.Log) (*CErc20Redeem, error) {
	event := new(CErc20Redeem)
	if err := _CErc20.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20RepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the CErc20 contract.
type CErc20RepayBorrowIterator struct {
	Event *CErc20RepayBorrow // Event containing the contract specifics and raw log

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
func (it *CErc20RepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20RepayBorrow)
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
		it.Event = new(CErc20RepayBorrow)
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
func (it *CErc20RepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20RepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20RepayBorrow represents a RepayBorrow event raised by the CErc20 contract.
type CErc20RepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) FilterRepayBorrow(opts *bind.FilterOpts) (*CErc20RepayBorrowIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &CErc20RepayBorrowIterator{contract: _CErc20.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *CErc20RepayBorrow) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20RepayBorrow)
				if err := _CErc20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CErc20 *CErc20Filterer) ParseRepayBorrow(log types.Log) (*CErc20RepayBorrow, error) {
	event := new(CErc20RepayBorrow)
	if err := _CErc20.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20ReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the CErc20 contract.
type CErc20ReservesReducedIterator struct {
	Event *CErc20ReservesReduced // Event containing the contract specifics and raw log

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
func (it *CErc20ReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20ReservesReduced)
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
		it.Event = new(CErc20ReservesReduced)
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
func (it *CErc20ReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20ReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20ReservesReduced represents a ReservesReduced event raised by the CErc20 contract.
type CErc20ReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CErc20 *CErc20Filterer) FilterReservesReduced(opts *bind.FilterOpts) (*CErc20ReservesReducedIterator, error) {

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &CErc20ReservesReducedIterator{contract: _CErc20.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CErc20 *CErc20Filterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *CErc20ReservesReduced) (event.Subscription, error) {

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20ReservesReduced)
				if err := _CErc20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CErc20 *CErc20Filterer) ParseReservesReduced(log types.Log) (*CErc20ReservesReduced, error) {
	event := new(CErc20ReservesReduced)
	if err := _CErc20.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CErc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CErc20 contract.
type CErc20TransferIterator struct {
	Event *CErc20Transfer // Event containing the contract specifics and raw log

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
func (it *CErc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CErc20Transfer)
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
		it.Event = new(CErc20Transfer)
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
func (it *CErc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CErc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CErc20Transfer represents a Transfer event raised by the CErc20 contract.
type CErc20Transfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CErc20 *CErc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CErc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CErc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CErc20TransferIterator{contract: _CErc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CErc20 *CErc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CErc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CErc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CErc20Transfer)
				if err := _CErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CErc20 *CErc20Filterer) ParseTransfer(log types.Log) (*CErc20Transfer, error) {
	event := new(CErc20Transfer)
	if err := _CErc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenABI is the input ABI used to generate the binding from.
const CTokenABI = "[{\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"comptroller_\",\"type\":\"address\"},{\"internalType\":\"contractInterestRateModel\",\"name\":\"interestRateModel_\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"initialExchangeRateMantissa_\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"decimals_\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestAccumulated\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"AccrueInterest\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"LiquidateBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"NewAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"oldComptroller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"NewComptroller\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"oldInterestRateModel\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"NewMarketInterestRateModel\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldPendingAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"NewPendingAdmin\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldReserveFactorMantissa\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"NewReserveFactor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"Redeem\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"accountBorrows\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalBorrows\",\"type\":\"uint256\"}],\"name\":\"RepayBorrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalReserves\",\"type\":\"uint256\"}],\"name\":\"ReservesReduced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[],\"name\":\"_acceptAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"_reduceReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"newComptroller\",\"type\":\"address\"}],\"name\":\"_setComptroller\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"newInterestRateModel\",\"type\":\"address\"}],\"name\":\"_setInterestRateModel\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"addresspayable\",\"name\":\"newPendingAdmin\",\"type\":\"address\"}],\"name\":\"_setPendingAdmin\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newReserveFactorMantissa\",\"type\":\"uint256\"}],\"name\":\"_setReserveFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"accrualBlockNumber\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"accrueInterest\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOfUnderlying\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"borrowBalanceStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"comptroller\",\"outputs\":[{\"internalType\":\"contractComptrollerInterface\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"exchangeRateCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exchangeRateStored\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountSnapshot\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCash\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialExchangeRateMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"interestRateModel\",\"outputs\":[{\"internalType\":\"contractInterestRateModel\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isCToken\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pendingAdmin\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveFactorMantissa\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyRatePerBlock\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalBorrows\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"totalBorrowsCurrent\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// CTokenFuncSigs maps the 4-byte function signature to its string representation.
var CTokenFuncSigs = map[string]string{
	"e9c714f2": "_acceptAdmin()",
	"601a0bf1": "_reduceReserves(uint256)",
	"4576b5db": "_setComptroller(address)",
	"f2b3abbd": "_setInterestRateModel(address)",
	"b71d1a0c": "_setPendingAdmin(address)",
	"fca7820b": "_setReserveFactor(uint256)",
	"6c540baf": "accrualBlockNumber()",
	"a6afed95": "accrueInterest()",
	"f851a440": "admin()",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"3af9e669": "balanceOfUnderlying(address)",
	"17bfdfbc": "borrowBalanceCurrent(address)",
	"95dd9193": "borrowBalanceStored(address)",
	"aa5af0fd": "borrowIndex()",
	"f8f9da28": "borrowRatePerBlock()",
	"5fe3b567": "comptroller()",
	"313ce567": "decimals()",
	"bd6d894d": "exchangeRateCurrent()",
	"182df0f5": "exchangeRateStored()",
	"c37f68e2": "getAccountSnapshot(address)",
	"3b1d21a2": "getCash()",
	"675d972c": "initialExchangeRateMantissa()",
	"f3fdb15a": "interestRateModel()",
	"fe9c44ae": "isCToken()",
	"06fdde03": "name()",
	"26782247": "pendingAdmin()",
	"173b9904": "reserveFactorMantissa()",
	"b2a02ff1": "seize(address,address,uint256)",
	"ae9d70b0": "supplyRatePerBlock()",
	"95d89b41": "symbol()",
	"47bd3718": "totalBorrows()",
	"73acee98": "totalBorrowsCurrent()",
	"8f840ddd": "totalReserves()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// CToken is an auto generated Go binding around an Ethereum contract.
type CToken struct {
	CTokenCaller     // Read-only binding to the contract
	CTokenTransactor // Write-only binding to the contract
	CTokenFilterer   // Log filterer for contract events
}

// CTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type CTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CTokenSession struct {
	Contract     *CToken           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CTokenCallerSession struct {
	Contract *CTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// CTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CTokenTransactorSession struct {
	Contract     *CTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type CTokenRaw struct {
	Contract *CToken // Generic contract binding to access the raw methods on
}

// CTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CTokenCallerRaw struct {
	Contract *CTokenCaller // Generic read-only contract binding to access the raw methods on
}

// CTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CTokenTransactorRaw struct {
	Contract *CTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCToken creates a new instance of CToken, bound to a specific deployed contract.
func NewCToken(address common.Address, backend bind.ContractBackend) (*CToken, error) {
	contract, err := bindCToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CToken{CTokenCaller: CTokenCaller{contract: contract}, CTokenTransactor: CTokenTransactor{contract: contract}, CTokenFilterer: CTokenFilterer{contract: contract}}, nil
}

// NewCTokenCaller creates a new read-only instance of CToken, bound to a specific deployed contract.
func NewCTokenCaller(address common.Address, caller bind.ContractCaller) (*CTokenCaller, error) {
	contract, err := bindCToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CTokenCaller{contract: contract}, nil
}

// NewCTokenTransactor creates a new write-only instance of CToken, bound to a specific deployed contract.
func NewCTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*CTokenTransactor, error) {
	contract, err := bindCToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CTokenTransactor{contract: contract}, nil
}

// NewCTokenFilterer creates a new log filterer instance of CToken, bound to a specific deployed contract.
func NewCTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*CTokenFilterer, error) {
	contract, err := bindCToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CTokenFilterer{contract: contract}, nil
}

// bindCToken binds a generic wrapper to an already deployed contract.
func bindCToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CToken *CTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CToken.Contract.CTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CToken *CTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.Contract.CTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CToken *CTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CToken.Contract.CTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CToken *CTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CToken *CTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CToken *CTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CToken.Contract.contract.Transact(opts, method, params...)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CToken *CTokenCaller) AccrualBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "accrualBlockNumber")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CToken *CTokenSession) AccrualBlockNumber() (*big.Int, error) {
	return _CToken.Contract.AccrualBlockNumber(&_CToken.CallOpts)
}

// AccrualBlockNumber is a free data retrieval call binding the contract method 0x6c540baf.
//
// Solidity: function accrualBlockNumber() view returns(uint256)
func (_CToken *CTokenCallerSession) AccrualBlockNumber() (*big.Int, error) {
	return _CToken.Contract.AccrualBlockNumber(&_CToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CToken *CTokenCaller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CToken *CTokenSession) Admin() (common.Address, error) {
	return _CToken.Contract.Admin(&_CToken.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_CToken *CTokenCallerSession) Admin() (common.Address, error) {
	return _CToken.Contract.Admin(&_CToken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CToken *CTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CToken *CTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CToken.Contract.Allowance(&_CToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_CToken *CTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _CToken.Contract.Allowance(&_CToken.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CToken *CTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CToken *CTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CToken.Contract.BalanceOf(&_CToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_CToken *CTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _CToken.Contract.BalanceOf(&_CToken.CallOpts, owner)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CToken *CTokenCaller) BorrowBalanceStored(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "borrowBalanceStored", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CToken *CTokenSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CToken.Contract.BorrowBalanceStored(&_CToken.CallOpts, account)
}

// BorrowBalanceStored is a free data retrieval call binding the contract method 0x95dd9193.
//
// Solidity: function borrowBalanceStored(address account) view returns(uint256)
func (_CToken *CTokenCallerSession) BorrowBalanceStored(account common.Address) (*big.Int, error) {
	return _CToken.Contract.BorrowBalanceStored(&_CToken.CallOpts, account)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CToken *CTokenCaller) BorrowIndex(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "borrowIndex")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CToken *CTokenSession) BorrowIndex() (*big.Int, error) {
	return _CToken.Contract.BorrowIndex(&_CToken.CallOpts)
}

// BorrowIndex is a free data retrieval call binding the contract method 0xaa5af0fd.
//
// Solidity: function borrowIndex() view returns(uint256)
func (_CToken *CTokenCallerSession) BorrowIndex() (*big.Int, error) {
	return _CToken.Contract.BorrowIndex(&_CToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CToken *CTokenCaller) BorrowRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "borrowRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CToken *CTokenSession) BorrowRatePerBlock() (*big.Int, error) {
	return _CToken.Contract.BorrowRatePerBlock(&_CToken.CallOpts)
}

// BorrowRatePerBlock is a free data retrieval call binding the contract method 0xf8f9da28.
//
// Solidity: function borrowRatePerBlock() view returns(uint256)
func (_CToken *CTokenCallerSession) BorrowRatePerBlock() (*big.Int, error) {
	return _CToken.Contract.BorrowRatePerBlock(&_CToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CToken *CTokenCaller) Comptroller(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "comptroller")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CToken *CTokenSession) Comptroller() (common.Address, error) {
	return _CToken.Contract.Comptroller(&_CToken.CallOpts)
}

// Comptroller is a free data retrieval call binding the contract method 0x5fe3b567.
//
// Solidity: function comptroller() view returns(address)
func (_CToken *CTokenCallerSession) Comptroller() (common.Address, error) {
	return _CToken.Contract.Comptroller(&_CToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CToken *CTokenCaller) Decimals(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CToken *CTokenSession) Decimals() (*big.Int, error) {
	return _CToken.Contract.Decimals(&_CToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint256)
func (_CToken *CTokenCallerSession) Decimals() (*big.Int, error) {
	return _CToken.Contract.Decimals(&_CToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CToken *CTokenCaller) ExchangeRateStored(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "exchangeRateStored")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CToken *CTokenSession) ExchangeRateStored() (*big.Int, error) {
	return _CToken.Contract.ExchangeRateStored(&_CToken.CallOpts)
}

// ExchangeRateStored is a free data retrieval call binding the contract method 0x182df0f5.
//
// Solidity: function exchangeRateStored() view returns(uint256)
func (_CToken *CTokenCallerSession) ExchangeRateStored() (*big.Int, error) {
	return _CToken.Contract.ExchangeRateStored(&_CToken.CallOpts)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CToken *CTokenCaller) GetAccountSnapshot(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "getAccountSnapshot", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CToken *CTokenSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CToken.Contract.GetAccountSnapshot(&_CToken.CallOpts, account)
}

// GetAccountSnapshot is a free data retrieval call binding the contract method 0xc37f68e2.
//
// Solidity: function getAccountSnapshot(address account) view returns(uint256, uint256, uint256, uint256)
func (_CToken *CTokenCallerSession) GetAccountSnapshot(account common.Address) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _CToken.Contract.GetAccountSnapshot(&_CToken.CallOpts, account)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CToken *CTokenCaller) GetCash(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "getCash")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CToken *CTokenSession) GetCash() (*big.Int, error) {
	return _CToken.Contract.GetCash(&_CToken.CallOpts)
}

// GetCash is a free data retrieval call binding the contract method 0x3b1d21a2.
//
// Solidity: function getCash() view returns(uint256)
func (_CToken *CTokenCallerSession) GetCash() (*big.Int, error) {
	return _CToken.Contract.GetCash(&_CToken.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CToken *CTokenCaller) InitialExchangeRateMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "initialExchangeRateMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CToken *CTokenSession) InitialExchangeRateMantissa() (*big.Int, error) {
	return _CToken.Contract.InitialExchangeRateMantissa(&_CToken.CallOpts)
}

// InitialExchangeRateMantissa is a free data retrieval call binding the contract method 0x675d972c.
//
// Solidity: function initialExchangeRateMantissa() view returns(uint256)
func (_CToken *CTokenCallerSession) InitialExchangeRateMantissa() (*big.Int, error) {
	return _CToken.Contract.InitialExchangeRateMantissa(&_CToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CToken *CTokenCaller) InterestRateModel(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "interestRateModel")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CToken *CTokenSession) InterestRateModel() (common.Address, error) {
	return _CToken.Contract.InterestRateModel(&_CToken.CallOpts)
}

// InterestRateModel is a free data retrieval call binding the contract method 0xf3fdb15a.
//
// Solidity: function interestRateModel() view returns(address)
func (_CToken *CTokenCallerSession) InterestRateModel() (common.Address, error) {
	return _CToken.Contract.InterestRateModel(&_CToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CToken *CTokenCaller) IsCToken(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "isCToken")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CToken *CTokenSession) IsCToken() (bool, error) {
	return _CToken.Contract.IsCToken(&_CToken.CallOpts)
}

// IsCToken is a free data retrieval call binding the contract method 0xfe9c44ae.
//
// Solidity: function isCToken() view returns(bool)
func (_CToken *CTokenCallerSession) IsCToken() (bool, error) {
	return _CToken.Contract.IsCToken(&_CToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CToken *CTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CToken *CTokenSession) Name() (string, error) {
	return _CToken.Contract.Name(&_CToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CToken *CTokenCallerSession) Name() (string, error) {
	return _CToken.Contract.Name(&_CToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CToken *CTokenCaller) PendingAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "pendingAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CToken *CTokenSession) PendingAdmin() (common.Address, error) {
	return _CToken.Contract.PendingAdmin(&_CToken.CallOpts)
}

// PendingAdmin is a free data retrieval call binding the contract method 0x26782247.
//
// Solidity: function pendingAdmin() view returns(address)
func (_CToken *CTokenCallerSession) PendingAdmin() (common.Address, error) {
	return _CToken.Contract.PendingAdmin(&_CToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CToken *CTokenCaller) ReserveFactorMantissa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "reserveFactorMantissa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CToken *CTokenSession) ReserveFactorMantissa() (*big.Int, error) {
	return _CToken.Contract.ReserveFactorMantissa(&_CToken.CallOpts)
}

// ReserveFactorMantissa is a free data retrieval call binding the contract method 0x173b9904.
//
// Solidity: function reserveFactorMantissa() view returns(uint256)
func (_CToken *CTokenCallerSession) ReserveFactorMantissa() (*big.Int, error) {
	return _CToken.Contract.ReserveFactorMantissa(&_CToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CToken *CTokenCaller) SupplyRatePerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "supplyRatePerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CToken *CTokenSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CToken.Contract.SupplyRatePerBlock(&_CToken.CallOpts)
}

// SupplyRatePerBlock is a free data retrieval call binding the contract method 0xae9d70b0.
//
// Solidity: function supplyRatePerBlock() view returns(uint256)
func (_CToken *CTokenCallerSession) SupplyRatePerBlock() (*big.Int, error) {
	return _CToken.Contract.SupplyRatePerBlock(&_CToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CToken *CTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CToken *CTokenSession) Symbol() (string, error) {
	return _CToken.Contract.Symbol(&_CToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CToken *CTokenCallerSession) Symbol() (string, error) {
	return _CToken.Contract.Symbol(&_CToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CToken *CTokenCaller) TotalBorrows(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "totalBorrows")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CToken *CTokenSession) TotalBorrows() (*big.Int, error) {
	return _CToken.Contract.TotalBorrows(&_CToken.CallOpts)
}

// TotalBorrows is a free data retrieval call binding the contract method 0x47bd3718.
//
// Solidity: function totalBorrows() view returns(uint256)
func (_CToken *CTokenCallerSession) TotalBorrows() (*big.Int, error) {
	return _CToken.Contract.TotalBorrows(&_CToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CToken *CTokenCaller) TotalReserves(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "totalReserves")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CToken *CTokenSession) TotalReserves() (*big.Int, error) {
	return _CToken.Contract.TotalReserves(&_CToken.CallOpts)
}

// TotalReserves is a free data retrieval call binding the contract method 0x8f840ddd.
//
// Solidity: function totalReserves() view returns(uint256)
func (_CToken *CTokenCallerSession) TotalReserves() (*big.Int, error) {
	return _CToken.Contract.TotalReserves(&_CToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CToken *CTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CToken *CTokenSession) TotalSupply() (*big.Int, error) {
	return _CToken.Contract.TotalSupply(&_CToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CToken *CTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _CToken.Contract.TotalSupply(&_CToken.CallOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CToken *CTokenTransactor) AcceptAdmin(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_acceptAdmin")
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CToken *CTokenSession) AcceptAdmin() (*types.Transaction, error) {
	return _CToken.Contract.AcceptAdmin(&_CToken.TransactOpts)
}

// AcceptAdmin is a paid mutator transaction binding the contract method 0xe9c714f2.
//
// Solidity: function _acceptAdmin() returns(uint256)
func (_CToken *CTokenTransactorSession) AcceptAdmin() (*types.Transaction, error) {
	return _CToken.Contract.AcceptAdmin(&_CToken.TransactOpts)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CToken *CTokenTransactor) ReduceReserves(opts *bind.TransactOpts, reduceAmount *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_reduceReserves", reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CToken *CTokenSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.ReduceReserves(&_CToken.TransactOpts, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0x601a0bf1.
//
// Solidity: function _reduceReserves(uint256 reduceAmount) returns(uint256)
func (_CToken *CTokenTransactorSession) ReduceReserves(reduceAmount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.ReduceReserves(&_CToken.TransactOpts, reduceAmount)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CToken *CTokenTransactor) SetComptroller(opts *bind.TransactOpts, newComptroller common.Address) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_setComptroller", newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CToken *CTokenSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetComptroller(&_CToken.TransactOpts, newComptroller)
}

// SetComptroller is a paid mutator transaction binding the contract method 0x4576b5db.
//
// Solidity: function _setComptroller(address newComptroller) returns(uint256)
func (_CToken *CTokenTransactorSession) SetComptroller(newComptroller common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetComptroller(&_CToken.TransactOpts, newComptroller)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CToken *CTokenTransactor) SetInterestRateModel(opts *bind.TransactOpts, newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_setInterestRateModel", newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CToken *CTokenSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetInterestRateModel(&_CToken.TransactOpts, newInterestRateModel)
}

// SetInterestRateModel is a paid mutator transaction binding the contract method 0xf2b3abbd.
//
// Solidity: function _setInterestRateModel(address newInterestRateModel) returns(uint256)
func (_CToken *CTokenTransactorSession) SetInterestRateModel(newInterestRateModel common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetInterestRateModel(&_CToken.TransactOpts, newInterestRateModel)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CToken *CTokenTransactor) SetPendingAdmin(opts *bind.TransactOpts, newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_setPendingAdmin", newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CToken *CTokenSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetPendingAdmin(&_CToken.TransactOpts, newPendingAdmin)
}

// SetPendingAdmin is a paid mutator transaction binding the contract method 0xb71d1a0c.
//
// Solidity: function _setPendingAdmin(address newPendingAdmin) returns(uint256)
func (_CToken *CTokenTransactorSession) SetPendingAdmin(newPendingAdmin common.Address) (*types.Transaction, error) {
	return _CToken.Contract.SetPendingAdmin(&_CToken.TransactOpts, newPendingAdmin)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CToken *CTokenTransactor) SetReserveFactor(opts *bind.TransactOpts, newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "_setReserveFactor", newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CToken *CTokenSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.SetReserveFactor(&_CToken.TransactOpts, newReserveFactorMantissa)
}

// SetReserveFactor is a paid mutator transaction binding the contract method 0xfca7820b.
//
// Solidity: function _setReserveFactor(uint256 newReserveFactorMantissa) returns(uint256)
func (_CToken *CTokenTransactorSession) SetReserveFactor(newReserveFactorMantissa *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.SetReserveFactor(&_CToken.TransactOpts, newReserveFactorMantissa)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CToken *CTokenTransactor) AccrueInterest(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "accrueInterest")
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CToken *CTokenSession) AccrueInterest() (*types.Transaction, error) {
	return _CToken.Contract.AccrueInterest(&_CToken.TransactOpts)
}

// AccrueInterest is a paid mutator transaction binding the contract method 0xa6afed95.
//
// Solidity: function accrueInterest() returns(uint256)
func (_CToken *CTokenTransactorSession) AccrueInterest() (*types.Transaction, error) {
	return _CToken.Contract.AccrueInterest(&_CToken.TransactOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CToken *CTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CToken *CTokenSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Approve(&_CToken.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_CToken *CTokenTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Approve(&_CToken.TransactOpts, spender, amount)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CToken *CTokenTransactor) BalanceOfUnderlying(opts *bind.TransactOpts, owner common.Address) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "balanceOfUnderlying", owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CToken *CTokenSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CToken.Contract.BalanceOfUnderlying(&_CToken.TransactOpts, owner)
}

// BalanceOfUnderlying is a paid mutator transaction binding the contract method 0x3af9e669.
//
// Solidity: function balanceOfUnderlying(address owner) returns(uint256)
func (_CToken *CTokenTransactorSession) BalanceOfUnderlying(owner common.Address) (*types.Transaction, error) {
	return _CToken.Contract.BalanceOfUnderlying(&_CToken.TransactOpts, owner)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CToken *CTokenTransactor) BorrowBalanceCurrent(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "borrowBalanceCurrent", account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CToken *CTokenSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CToken.Contract.BorrowBalanceCurrent(&_CToken.TransactOpts, account)
}

// BorrowBalanceCurrent is a paid mutator transaction binding the contract method 0x17bfdfbc.
//
// Solidity: function borrowBalanceCurrent(address account) returns(uint256)
func (_CToken *CTokenTransactorSession) BorrowBalanceCurrent(account common.Address) (*types.Transaction, error) {
	return _CToken.Contract.BorrowBalanceCurrent(&_CToken.TransactOpts, account)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CToken *CTokenTransactor) ExchangeRateCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "exchangeRateCurrent")
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CToken *CTokenSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CToken.Contract.ExchangeRateCurrent(&_CToken.TransactOpts)
}

// ExchangeRateCurrent is a paid mutator transaction binding the contract method 0xbd6d894d.
//
// Solidity: function exchangeRateCurrent() returns(uint256)
func (_CToken *CTokenTransactorSession) ExchangeRateCurrent() (*types.Transaction, error) {
	return _CToken.Contract.ExchangeRateCurrent(&_CToken.TransactOpts)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CToken *CTokenTransactor) Seize(opts *bind.TransactOpts, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "seize", liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CToken *CTokenSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Seize(&_CToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// Seize is a paid mutator transaction binding the contract method 0xb2a02ff1.
//
// Solidity: function seize(address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_CToken *CTokenTransactorSession) Seize(liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Seize(&_CToken.TransactOpts, liquidator, borrower, seizeTokens)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CToken *CTokenTransactor) TotalBorrowsCurrent(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "totalBorrowsCurrent")
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CToken *CTokenSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CToken.Contract.TotalBorrowsCurrent(&_CToken.TransactOpts)
}

// TotalBorrowsCurrent is a paid mutator transaction binding the contract method 0x73acee98.
//
// Solidity: function totalBorrowsCurrent() returns(uint256)
func (_CToken *CTokenTransactorSession) TotalBorrowsCurrent() (*types.Transaction, error) {
	return _CToken.Contract.TotalBorrowsCurrent(&_CToken.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CToken *CTokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CToken *CTokenSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Transfer(&_CToken.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool)
func (_CToken *CTokenTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.Transfer(&_CToken.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CToken *CTokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CToken *CTokenSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.TransferFrom(&_CToken.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool)
func (_CToken *CTokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _CToken.Contract.TransferFrom(&_CToken.TransactOpts, src, dst, amount)
}

// CTokenAccrueInterestIterator is returned from FilterAccrueInterest and is used to iterate over the raw logs and unpacked data for AccrueInterest events raised by the CToken contract.
type CTokenAccrueInterestIterator struct {
	Event *CTokenAccrueInterest // Event containing the contract specifics and raw log

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
func (it *CTokenAccrueInterestIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenAccrueInterest)
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
		it.Event = new(CTokenAccrueInterest)
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
func (it *CTokenAccrueInterestIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenAccrueInterestIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenAccrueInterest represents a AccrueInterest event raised by the CToken contract.
type CTokenAccrueInterest struct {
	InterestAccumulated *big.Int
	BorrowIndex         *big.Int
	TotalBorrows        *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterAccrueInterest is a free log retrieval operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CToken *CTokenFilterer) FilterAccrueInterest(opts *bind.FilterOpts) (*CTokenAccrueInterestIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return &CTokenAccrueInterestIterator{contract: _CToken.contract, event: "AccrueInterest", logs: logs, sub: sub}, nil
}

// WatchAccrueInterest is a free log subscription operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CToken *CTokenFilterer) WatchAccrueInterest(opts *bind.WatchOpts, sink chan<- *CTokenAccrueInterest) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "AccrueInterest")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenAccrueInterest)
				if err := _CToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
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

// ParseAccrueInterest is a log parse operation binding the contract event 0x875352fb3fadeb8c0be7cbbe8ff761b308fa7033470cd0287f02f3436fd76cb9.
//
// Solidity: event AccrueInterest(uint256 interestAccumulated, uint256 borrowIndex, uint256 totalBorrows)
func (_CToken *CTokenFilterer) ParseAccrueInterest(log types.Log) (*CTokenAccrueInterest, error) {
	event := new(CTokenAccrueInterest)
	if err := _CToken.contract.UnpackLog(event, "AccrueInterest", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the CToken contract.
type CTokenApprovalIterator struct {
	Event *CTokenApproval // Event containing the contract specifics and raw log

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
func (it *CTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenApproval)
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
		it.Event = new(CTokenApproval)
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
func (it *CTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenApproval represents a Approval event raised by the CToken contract.
type CTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CToken *CTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*CTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &CTokenApprovalIterator{contract: _CToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CToken *CTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *CTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenApproval)
				if err := _CToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_CToken *CTokenFilterer) ParseApproval(log types.Log) (*CTokenApproval, error) {
	event := new(CTokenApproval)
	if err := _CToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the CToken contract.
type CTokenBorrowIterator struct {
	Event *CTokenBorrow // Event containing the contract specifics and raw log

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
func (it *CTokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenBorrow)
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
		it.Event = new(CTokenBorrow)
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
func (it *CTokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenBorrow represents a Borrow event raised by the CToken contract.
type CTokenBorrow struct {
	Borrower       common.Address
	BorrowAmount   *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) FilterBorrow(opts *bind.FilterOpts) (*CTokenBorrowIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return &CTokenBorrowIterator{contract: _CToken.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *CTokenBorrow) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Borrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenBorrow)
				if err := _CToken.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x13ed6866d4e1ee6da46f845c46d7e54120883d75c5ea9a2dacc1c4ca8984ab80.
//
// Solidity: event Borrow(address borrower, uint256 borrowAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) ParseBorrow(log types.Log) (*CTokenBorrow, error) {
	event := new(CTokenBorrow)
	if err := _CToken.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the CToken contract.
type CTokenFailureIterator struct {
	Event *CTokenFailure // Event containing the contract specifics and raw log

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
func (it *CTokenFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenFailure)
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
		it.Event = new(CTokenFailure)
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
func (it *CTokenFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenFailure represents a Failure event raised by the CToken contract.
type CTokenFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CToken *CTokenFilterer) FilterFailure(opts *bind.FilterOpts) (*CTokenFailureIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &CTokenFailureIterator{contract: _CToken.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CToken *CTokenFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *CTokenFailure) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenFailure)
				if err := _CToken.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_CToken *CTokenFilterer) ParseFailure(log types.Log) (*CTokenFailure, error) {
	event := new(CTokenFailure)
	if err := _CToken.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenLiquidateBorrowIterator is returned from FilterLiquidateBorrow and is used to iterate over the raw logs and unpacked data for LiquidateBorrow events raised by the CToken contract.
type CTokenLiquidateBorrowIterator struct {
	Event *CTokenLiquidateBorrow // Event containing the contract specifics and raw log

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
func (it *CTokenLiquidateBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenLiquidateBorrow)
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
		it.Event = new(CTokenLiquidateBorrow)
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
func (it *CTokenLiquidateBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenLiquidateBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenLiquidateBorrow represents a LiquidateBorrow event raised by the CToken contract.
type CTokenLiquidateBorrow struct {
	Liquidator       common.Address
	Borrower         common.Address
	RepayAmount      *big.Int
	CTokenCollateral common.Address
	SeizeTokens      *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidateBorrow is a free log retrieval operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CToken *CTokenFilterer) FilterLiquidateBorrow(opts *bind.FilterOpts) (*CTokenLiquidateBorrowIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return &CTokenLiquidateBorrowIterator{contract: _CToken.contract, event: "LiquidateBorrow", logs: logs, sub: sub}, nil
}

// WatchLiquidateBorrow is a free log subscription operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CToken *CTokenFilterer) WatchLiquidateBorrow(opts *bind.WatchOpts, sink chan<- *CTokenLiquidateBorrow) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "LiquidateBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenLiquidateBorrow)
				if err := _CToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
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

// ParseLiquidateBorrow is a log parse operation binding the contract event 0x298637f684da70674f26509b10f07ec2fbc77a335ab1e7d6215a4b2484d8bb52.
//
// Solidity: event LiquidateBorrow(address liquidator, address borrower, uint256 repayAmount, address cTokenCollateral, uint256 seizeTokens)
func (_CToken *CTokenFilterer) ParseLiquidateBorrow(log types.Log) (*CTokenLiquidateBorrow, error) {
	event := new(CTokenLiquidateBorrow)
	if err := _CToken.contract.UnpackLog(event, "LiquidateBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the CToken contract.
type CTokenMintIterator struct {
	Event *CTokenMint // Event containing the contract specifics and raw log

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
func (it *CTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenMint)
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
		it.Event = new(CTokenMint)
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
func (it *CTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenMint represents a Mint event raised by the CToken contract.
type CTokenMint struct {
	Minter     common.Address
	MintAmount *big.Int
	MintTokens *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CToken *CTokenFilterer) FilterMint(opts *bind.FilterOpts) (*CTokenMintIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return &CTokenMintIterator{contract: _CToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CToken *CTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *CTokenMint) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Mint")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenMint)
				if err := _CToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x4c209b5fc8ad50758f13e2e1088ba56a560dff690a1c6fef26394f4c03821c4f.
//
// Solidity: event Mint(address minter, uint256 mintAmount, uint256 mintTokens)
func (_CToken *CTokenFilterer) ParseMint(log types.Log) (*CTokenMint, error) {
	event := new(CTokenMint)
	if err := _CToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenNewAdminIterator is returned from FilterNewAdmin and is used to iterate over the raw logs and unpacked data for NewAdmin events raised by the CToken contract.
type CTokenNewAdminIterator struct {
	Event *CTokenNewAdmin // Event containing the contract specifics and raw log

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
func (it *CTokenNewAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenNewAdmin)
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
		it.Event = new(CTokenNewAdmin)
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
func (it *CTokenNewAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenNewAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenNewAdmin represents a NewAdmin event raised by the CToken contract.
type CTokenNewAdmin struct {
	OldAdmin common.Address
	NewAdmin common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewAdmin is a free log retrieval operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CToken *CTokenFilterer) FilterNewAdmin(opts *bind.FilterOpts) (*CTokenNewAdminIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return &CTokenNewAdminIterator{contract: _CToken.contract, event: "NewAdmin", logs: logs, sub: sub}, nil
}

// WatchNewAdmin is a free log subscription operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CToken *CTokenFilterer) WatchNewAdmin(opts *bind.WatchOpts, sink chan<- *CTokenNewAdmin) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "NewAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenNewAdmin)
				if err := _CToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
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

// ParseNewAdmin is a log parse operation binding the contract event 0xf9ffabca9c8276e99321725bcb43fb076a6c66a54b7f21c4e8146d8519b417dc.
//
// Solidity: event NewAdmin(address oldAdmin, address newAdmin)
func (_CToken *CTokenFilterer) ParseNewAdmin(log types.Log) (*CTokenNewAdmin, error) {
	event := new(CTokenNewAdmin)
	if err := _CToken.contract.UnpackLog(event, "NewAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenNewComptrollerIterator is returned from FilterNewComptroller and is used to iterate over the raw logs and unpacked data for NewComptroller events raised by the CToken contract.
type CTokenNewComptrollerIterator struct {
	Event *CTokenNewComptroller // Event containing the contract specifics and raw log

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
func (it *CTokenNewComptrollerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenNewComptroller)
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
		it.Event = new(CTokenNewComptroller)
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
func (it *CTokenNewComptrollerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenNewComptrollerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenNewComptroller represents a NewComptroller event raised by the CToken contract.
type CTokenNewComptroller struct {
	OldComptroller common.Address
	NewComptroller common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterNewComptroller is a free log retrieval operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CToken *CTokenFilterer) FilterNewComptroller(opts *bind.FilterOpts) (*CTokenNewComptrollerIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return &CTokenNewComptrollerIterator{contract: _CToken.contract, event: "NewComptroller", logs: logs, sub: sub}, nil
}

// WatchNewComptroller is a free log subscription operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CToken *CTokenFilterer) WatchNewComptroller(opts *bind.WatchOpts, sink chan<- *CTokenNewComptroller) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "NewComptroller")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenNewComptroller)
				if err := _CToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
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

// ParseNewComptroller is a log parse operation binding the contract event 0x7ac369dbd14fa5ea3f473ed67cc9d598964a77501540ba6751eb0b3decf5870d.
//
// Solidity: event NewComptroller(address oldComptroller, address newComptroller)
func (_CToken *CTokenFilterer) ParseNewComptroller(log types.Log) (*CTokenNewComptroller, error) {
	event := new(CTokenNewComptroller)
	if err := _CToken.contract.UnpackLog(event, "NewComptroller", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenNewMarketInterestRateModelIterator is returned from FilterNewMarketInterestRateModel and is used to iterate over the raw logs and unpacked data for NewMarketInterestRateModel events raised by the CToken contract.
type CTokenNewMarketInterestRateModelIterator struct {
	Event *CTokenNewMarketInterestRateModel // Event containing the contract specifics and raw log

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
func (it *CTokenNewMarketInterestRateModelIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenNewMarketInterestRateModel)
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
		it.Event = new(CTokenNewMarketInterestRateModel)
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
func (it *CTokenNewMarketInterestRateModelIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenNewMarketInterestRateModelIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenNewMarketInterestRateModel represents a NewMarketInterestRateModel event raised by the CToken contract.
type CTokenNewMarketInterestRateModel struct {
	OldInterestRateModel common.Address
	NewInterestRateModel common.Address
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterNewMarketInterestRateModel is a free log retrieval operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CToken *CTokenFilterer) FilterNewMarketInterestRateModel(opts *bind.FilterOpts) (*CTokenNewMarketInterestRateModelIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return &CTokenNewMarketInterestRateModelIterator{contract: _CToken.contract, event: "NewMarketInterestRateModel", logs: logs, sub: sub}, nil
}

// WatchNewMarketInterestRateModel is a free log subscription operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CToken *CTokenFilterer) WatchNewMarketInterestRateModel(opts *bind.WatchOpts, sink chan<- *CTokenNewMarketInterestRateModel) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "NewMarketInterestRateModel")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenNewMarketInterestRateModel)
				if err := _CToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
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

// ParseNewMarketInterestRateModel is a log parse operation binding the contract event 0xedffc32e068c7c95dfd4bdfd5c4d939a084d6b11c4199eac8436ed234d72f926.
//
// Solidity: event NewMarketInterestRateModel(address oldInterestRateModel, address newInterestRateModel)
func (_CToken *CTokenFilterer) ParseNewMarketInterestRateModel(log types.Log) (*CTokenNewMarketInterestRateModel, error) {
	event := new(CTokenNewMarketInterestRateModel)
	if err := _CToken.contract.UnpackLog(event, "NewMarketInterestRateModel", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenNewPendingAdminIterator is returned from FilterNewPendingAdmin and is used to iterate over the raw logs and unpacked data for NewPendingAdmin events raised by the CToken contract.
type CTokenNewPendingAdminIterator struct {
	Event *CTokenNewPendingAdmin // Event containing the contract specifics and raw log

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
func (it *CTokenNewPendingAdminIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenNewPendingAdmin)
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
		it.Event = new(CTokenNewPendingAdmin)
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
func (it *CTokenNewPendingAdminIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenNewPendingAdminIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenNewPendingAdmin represents a NewPendingAdmin event raised by the CToken contract.
type CTokenNewPendingAdmin struct {
	OldPendingAdmin common.Address
	NewPendingAdmin common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterNewPendingAdmin is a free log retrieval operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CToken *CTokenFilterer) FilterNewPendingAdmin(opts *bind.FilterOpts) (*CTokenNewPendingAdminIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return &CTokenNewPendingAdminIterator{contract: _CToken.contract, event: "NewPendingAdmin", logs: logs, sub: sub}, nil
}

// WatchNewPendingAdmin is a free log subscription operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CToken *CTokenFilterer) WatchNewPendingAdmin(opts *bind.WatchOpts, sink chan<- *CTokenNewPendingAdmin) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "NewPendingAdmin")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenNewPendingAdmin)
				if err := _CToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
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

// ParseNewPendingAdmin is a log parse operation binding the contract event 0xca4f2f25d0898edd99413412fb94012f9e54ec8142f9b093e7720646a95b16a9.
//
// Solidity: event NewPendingAdmin(address oldPendingAdmin, address newPendingAdmin)
func (_CToken *CTokenFilterer) ParseNewPendingAdmin(log types.Log) (*CTokenNewPendingAdmin, error) {
	event := new(CTokenNewPendingAdmin)
	if err := _CToken.contract.UnpackLog(event, "NewPendingAdmin", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenNewReserveFactorIterator is returned from FilterNewReserveFactor and is used to iterate over the raw logs and unpacked data for NewReserveFactor events raised by the CToken contract.
type CTokenNewReserveFactorIterator struct {
	Event *CTokenNewReserveFactor // Event containing the contract specifics and raw log

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
func (it *CTokenNewReserveFactorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenNewReserveFactor)
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
		it.Event = new(CTokenNewReserveFactor)
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
func (it *CTokenNewReserveFactorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenNewReserveFactorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenNewReserveFactor represents a NewReserveFactor event raised by the CToken contract.
type CTokenNewReserveFactor struct {
	OldReserveFactorMantissa *big.Int
	NewReserveFactorMantissa *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterNewReserveFactor is a free log retrieval operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CToken *CTokenFilterer) FilterNewReserveFactor(opts *bind.FilterOpts) (*CTokenNewReserveFactorIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return &CTokenNewReserveFactorIterator{contract: _CToken.contract, event: "NewReserveFactor", logs: logs, sub: sub}, nil
}

// WatchNewReserveFactor is a free log subscription operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CToken *CTokenFilterer) WatchNewReserveFactor(opts *bind.WatchOpts, sink chan<- *CTokenNewReserveFactor) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "NewReserveFactor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenNewReserveFactor)
				if err := _CToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
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

// ParseNewReserveFactor is a log parse operation binding the contract event 0xaaa68312e2ea9d50e16af5068410ab56e1a1fd06037b1a35664812c30f821460.
//
// Solidity: event NewReserveFactor(uint256 oldReserveFactorMantissa, uint256 newReserveFactorMantissa)
func (_CToken *CTokenFilterer) ParseNewReserveFactor(log types.Log) (*CTokenNewReserveFactor, error) {
	event := new(CTokenNewReserveFactor)
	if err := _CToken.contract.UnpackLog(event, "NewReserveFactor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenRedeemIterator is returned from FilterRedeem and is used to iterate over the raw logs and unpacked data for Redeem events raised by the CToken contract.
type CTokenRedeemIterator struct {
	Event *CTokenRedeem // Event containing the contract specifics and raw log

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
func (it *CTokenRedeemIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenRedeem)
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
		it.Event = new(CTokenRedeem)
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
func (it *CTokenRedeemIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenRedeemIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenRedeem represents a Redeem event raised by the CToken contract.
type CTokenRedeem struct {
	Redeemer     common.Address
	RedeemAmount *big.Int
	RedeemTokens *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterRedeem is a free log retrieval operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CToken *CTokenFilterer) FilterRedeem(opts *bind.FilterOpts) (*CTokenRedeemIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return &CTokenRedeemIterator{contract: _CToken.contract, event: "Redeem", logs: logs, sub: sub}, nil
}

// WatchRedeem is a free log subscription operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CToken *CTokenFilterer) WatchRedeem(opts *bind.WatchOpts, sink chan<- *CTokenRedeem) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Redeem")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenRedeem)
				if err := _CToken.contract.UnpackLog(event, "Redeem", log); err != nil {
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

// ParseRedeem is a log parse operation binding the contract event 0xe5b754fb1abb7f01b499791d0b820ae3b6af3424ac1c59768edb53f4ec31a929.
//
// Solidity: event Redeem(address redeemer, uint256 redeemAmount, uint256 redeemTokens)
func (_CToken *CTokenFilterer) ParseRedeem(log types.Log) (*CTokenRedeem, error) {
	event := new(CTokenRedeem)
	if err := _CToken.contract.UnpackLog(event, "Redeem", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenRepayBorrowIterator is returned from FilterRepayBorrow and is used to iterate over the raw logs and unpacked data for RepayBorrow events raised by the CToken contract.
type CTokenRepayBorrowIterator struct {
	Event *CTokenRepayBorrow // Event containing the contract specifics and raw log

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
func (it *CTokenRepayBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenRepayBorrow)
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
		it.Event = new(CTokenRepayBorrow)
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
func (it *CTokenRepayBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenRepayBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenRepayBorrow represents a RepayBorrow event raised by the CToken contract.
type CTokenRepayBorrow struct {
	Payer          common.Address
	Borrower       common.Address
	RepayAmount    *big.Int
	AccountBorrows *big.Int
	TotalBorrows   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterRepayBorrow is a free log retrieval operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) FilterRepayBorrow(opts *bind.FilterOpts) (*CTokenRepayBorrowIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return &CTokenRepayBorrowIterator{contract: _CToken.contract, event: "RepayBorrow", logs: logs, sub: sub}, nil
}

// WatchRepayBorrow is a free log subscription operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) WatchRepayBorrow(opts *bind.WatchOpts, sink chan<- *CTokenRepayBorrow) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "RepayBorrow")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenRepayBorrow)
				if err := _CToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
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

// ParseRepayBorrow is a log parse operation binding the contract event 0x1a2a22cb034d26d1854bdc6666a5b91fe25efbbb5dcad3b0355478d6f5c362a1.
//
// Solidity: event RepayBorrow(address payer, address borrower, uint256 repayAmount, uint256 accountBorrows, uint256 totalBorrows)
func (_CToken *CTokenFilterer) ParseRepayBorrow(log types.Log) (*CTokenRepayBorrow, error) {
	event := new(CTokenRepayBorrow)
	if err := _CToken.contract.UnpackLog(event, "RepayBorrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenReservesReducedIterator is returned from FilterReservesReduced and is used to iterate over the raw logs and unpacked data for ReservesReduced events raised by the CToken contract.
type CTokenReservesReducedIterator struct {
	Event *CTokenReservesReduced // Event containing the contract specifics and raw log

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
func (it *CTokenReservesReducedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenReservesReduced)
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
		it.Event = new(CTokenReservesReduced)
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
func (it *CTokenReservesReducedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenReservesReducedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenReservesReduced represents a ReservesReduced event raised by the CToken contract.
type CTokenReservesReduced struct {
	Admin            common.Address
	ReduceAmount     *big.Int
	NewTotalReserves *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterReservesReduced is a free log retrieval operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CToken *CTokenFilterer) FilterReservesReduced(opts *bind.FilterOpts) (*CTokenReservesReducedIterator, error) {

	logs, sub, err := _CToken.contract.FilterLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return &CTokenReservesReducedIterator{contract: _CToken.contract, event: "ReservesReduced", logs: logs, sub: sub}, nil
}

// WatchReservesReduced is a free log subscription operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CToken *CTokenFilterer) WatchReservesReduced(opts *bind.WatchOpts, sink chan<- *CTokenReservesReduced) (event.Subscription, error) {

	logs, sub, err := _CToken.contract.WatchLogs(opts, "ReservesReduced")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenReservesReduced)
				if err := _CToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
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

// ParseReservesReduced is a log parse operation binding the contract event 0x3bad0c59cf2f06e7314077049f48a93578cd16f5ef92329f1dab1420a99c177e.
//
// Solidity: event ReservesReduced(address admin, uint256 reduceAmount, uint256 newTotalReserves)
func (_CToken *CTokenFilterer) ParseReservesReduced(log types.Log) (*CTokenReservesReduced, error) {
	event := new(CTokenReservesReduced)
	if err := _CToken.contract.UnpackLog(event, "ReservesReduced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CToken contract.
type CTokenTransferIterator struct {
	Event *CTokenTransfer // Event containing the contract specifics and raw log

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
func (it *CTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CTokenTransfer)
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
		it.Event = new(CTokenTransfer)
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
func (it *CTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CTokenTransfer represents a Transfer event raised by the CToken contract.
type CTokenTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CToken *CTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CTokenTransferIterator{contract: _CToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CToken *CTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CTokenTransfer)
				if err := _CToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_CToken *CTokenFilterer) ParseTransfer(log types.Log) (*CTokenTransfer, error) {
	event := new(CTokenTransfer)
	if err := _CToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CarefulMathABI is the input ABI used to generate the binding from.
const CarefulMathABI = "[]"

// CarefulMathBin is the compiled bytecode used for deploying new contracts.
var CarefulMathBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a723158206eadab0baa369d2b59af4d1f5e7e567463a8641ae696f8b011129aa528cb256064736f6c63430005110032"

// DeployCarefulMath deploys a new Ethereum contract, binding an instance of CarefulMath to it.
func DeployCarefulMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CarefulMath, error) {
	parsed, err := abi.JSON(strings.NewReader(CarefulMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CarefulMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CarefulMath{CarefulMathCaller: CarefulMathCaller{contract: contract}, CarefulMathTransactor: CarefulMathTransactor{contract: contract}, CarefulMathFilterer: CarefulMathFilterer{contract: contract}}, nil
}

// CarefulMath is an auto generated Go binding around an Ethereum contract.
type CarefulMath struct {
	CarefulMathCaller     // Read-only binding to the contract
	CarefulMathTransactor // Write-only binding to the contract
	CarefulMathFilterer   // Log filterer for contract events
}

// CarefulMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type CarefulMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarefulMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CarefulMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarefulMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CarefulMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CarefulMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CarefulMathSession struct {
	Contract     *CarefulMath      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CarefulMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CarefulMathCallerSession struct {
	Contract *CarefulMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// CarefulMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CarefulMathTransactorSession struct {
	Contract     *CarefulMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// CarefulMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type CarefulMathRaw struct {
	Contract *CarefulMath // Generic contract binding to access the raw methods on
}

// CarefulMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CarefulMathCallerRaw struct {
	Contract *CarefulMathCaller // Generic read-only contract binding to access the raw methods on
}

// CarefulMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CarefulMathTransactorRaw struct {
	Contract *CarefulMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCarefulMath creates a new instance of CarefulMath, bound to a specific deployed contract.
func NewCarefulMath(address common.Address, backend bind.ContractBackend) (*CarefulMath, error) {
	contract, err := bindCarefulMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CarefulMath{CarefulMathCaller: CarefulMathCaller{contract: contract}, CarefulMathTransactor: CarefulMathTransactor{contract: contract}, CarefulMathFilterer: CarefulMathFilterer{contract: contract}}, nil
}

// NewCarefulMathCaller creates a new read-only instance of CarefulMath, bound to a specific deployed contract.
func NewCarefulMathCaller(address common.Address, caller bind.ContractCaller) (*CarefulMathCaller, error) {
	contract, err := bindCarefulMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CarefulMathCaller{contract: contract}, nil
}

// NewCarefulMathTransactor creates a new write-only instance of CarefulMath, bound to a specific deployed contract.
func NewCarefulMathTransactor(address common.Address, transactor bind.ContractTransactor) (*CarefulMathTransactor, error) {
	contract, err := bindCarefulMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CarefulMathTransactor{contract: contract}, nil
}

// NewCarefulMathFilterer creates a new log filterer instance of CarefulMath, bound to a specific deployed contract.
func NewCarefulMathFilterer(address common.Address, filterer bind.ContractFilterer) (*CarefulMathFilterer, error) {
	contract, err := bindCarefulMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CarefulMathFilterer{contract: contract}, nil
}

// bindCarefulMath binds a generic wrapper to an already deployed contract.
func bindCarefulMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CarefulMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarefulMath *CarefulMathRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarefulMath.Contract.CarefulMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarefulMath *CarefulMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarefulMath.Contract.CarefulMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarefulMath *CarefulMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarefulMath.Contract.CarefulMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CarefulMath *CarefulMathCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CarefulMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CarefulMath *CarefulMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CarefulMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CarefulMath *CarefulMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CarefulMath.Contract.contract.Transact(opts, method, params...)
}

// ComptrollerErrorReporterABI is the input ABI used to generate the binding from.
const ComptrollerErrorReporterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"}]"

// ComptrollerErrorReporterBin is the compiled bytecode used for deploying new contracts.
var ComptrollerErrorReporterBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820488a0dda88815ac55c6a9f0fb74073fb5f49b895cb1ef0bdc6320bbf8529668164736f6c63430005110032"

// DeployComptrollerErrorReporter deploys a new Ethereum contract, binding an instance of ComptrollerErrorReporter to it.
func DeployComptrollerErrorReporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ComptrollerErrorReporter, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerErrorReporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ComptrollerErrorReporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ComptrollerErrorReporter{ComptrollerErrorReporterCaller: ComptrollerErrorReporterCaller{contract: contract}, ComptrollerErrorReporterTransactor: ComptrollerErrorReporterTransactor{contract: contract}, ComptrollerErrorReporterFilterer: ComptrollerErrorReporterFilterer{contract: contract}}, nil
}

// ComptrollerErrorReporter is an auto generated Go binding around an Ethereum contract.
type ComptrollerErrorReporter struct {
	ComptrollerErrorReporterCaller     // Read-only binding to the contract
	ComptrollerErrorReporterTransactor // Write-only binding to the contract
	ComptrollerErrorReporterFilterer   // Log filterer for contract events
}

// ComptrollerErrorReporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerErrorReporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerErrorReporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerErrorReporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerErrorReporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerErrorReporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerErrorReporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerErrorReporterSession struct {
	Contract     *ComptrollerErrorReporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ComptrollerErrorReporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerErrorReporterCallerSession struct {
	Contract *ComptrollerErrorReporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// ComptrollerErrorReporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerErrorReporterTransactorSession struct {
	Contract     *ComptrollerErrorReporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// ComptrollerErrorReporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerErrorReporterRaw struct {
	Contract *ComptrollerErrorReporter // Generic contract binding to access the raw methods on
}

// ComptrollerErrorReporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerErrorReporterCallerRaw struct {
	Contract *ComptrollerErrorReporterCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerErrorReporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerErrorReporterTransactorRaw struct {
	Contract *ComptrollerErrorReporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptrollerErrorReporter creates a new instance of ComptrollerErrorReporter, bound to a specific deployed contract.
func NewComptrollerErrorReporter(address common.Address, backend bind.ContractBackend) (*ComptrollerErrorReporter, error) {
	contract, err := bindComptrollerErrorReporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComptrollerErrorReporter{ComptrollerErrorReporterCaller: ComptrollerErrorReporterCaller{contract: contract}, ComptrollerErrorReporterTransactor: ComptrollerErrorReporterTransactor{contract: contract}, ComptrollerErrorReporterFilterer: ComptrollerErrorReporterFilterer{contract: contract}}, nil
}

// NewComptrollerErrorReporterCaller creates a new read-only instance of ComptrollerErrorReporter, bound to a specific deployed contract.
func NewComptrollerErrorReporterCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerErrorReporterCaller, error) {
	contract, err := bindComptrollerErrorReporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerErrorReporterCaller{contract: contract}, nil
}

// NewComptrollerErrorReporterTransactor creates a new write-only instance of ComptrollerErrorReporter, bound to a specific deployed contract.
func NewComptrollerErrorReporterTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerErrorReporterTransactor, error) {
	contract, err := bindComptrollerErrorReporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerErrorReporterTransactor{contract: contract}, nil
}

// NewComptrollerErrorReporterFilterer creates a new log filterer instance of ComptrollerErrorReporter, bound to a specific deployed contract.
func NewComptrollerErrorReporterFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerErrorReporterFilterer, error) {
	contract, err := bindComptrollerErrorReporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerErrorReporterFilterer{contract: contract}, nil
}

// bindComptrollerErrorReporter binds a generic wrapper to an already deployed contract.
func bindComptrollerErrorReporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerErrorReporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerErrorReporter.Contract.ComptrollerErrorReporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerErrorReporter.Contract.ComptrollerErrorReporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerErrorReporter.Contract.ComptrollerErrorReporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerErrorReporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerErrorReporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerErrorReporter *ComptrollerErrorReporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerErrorReporter.Contract.contract.Transact(opts, method, params...)
}

// ComptrollerErrorReporterFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the ComptrollerErrorReporter contract.
type ComptrollerErrorReporterFailureIterator struct {
	Event *ComptrollerErrorReporterFailure // Event containing the contract specifics and raw log

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
func (it *ComptrollerErrorReporterFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ComptrollerErrorReporterFailure)
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
		it.Event = new(ComptrollerErrorReporterFailure)
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
func (it *ComptrollerErrorReporterFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ComptrollerErrorReporterFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ComptrollerErrorReporterFailure represents a Failure event raised by the ComptrollerErrorReporter contract.
type ComptrollerErrorReporterFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_ComptrollerErrorReporter *ComptrollerErrorReporterFilterer) FilterFailure(opts *bind.FilterOpts) (*ComptrollerErrorReporterFailureIterator, error) {

	logs, sub, err := _ComptrollerErrorReporter.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &ComptrollerErrorReporterFailureIterator{contract: _ComptrollerErrorReporter.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_ComptrollerErrorReporter *ComptrollerErrorReporterFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *ComptrollerErrorReporterFailure) (event.Subscription, error) {

	logs, sub, err := _ComptrollerErrorReporter.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ComptrollerErrorReporterFailure)
				if err := _ComptrollerErrorReporter.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_ComptrollerErrorReporter *ComptrollerErrorReporterFilterer) ParseFailure(log types.Log) (*ComptrollerErrorReporterFailure, error) {
	event := new(ComptrollerErrorReporterFailure)
	if err := _ComptrollerErrorReporter.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ComptrollerInterfaceABI is the input ABI used to generate the binding from.
const ComptrollerInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"cTokens\",\"type\":\"address[]\"}],\"name\":\"enterMarkets\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"}],\"name\":\"exitMarket\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isComptroller\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateCalculateSeizeTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"name\":\"mintAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"mintTokens\",\"type\":\"uint256\"}],\"name\":\"mintVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"redeemer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"redeemAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"redeemTokens\",\"type\":\"uint256\"}],\"name\":\"redeemVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"repayBorrowAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"payer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowerIndex\",\"type\":\"uint256\"}],\"name\":\"repayBorrowVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"seizeTokens\",\"type\":\"uint256\"}],\"name\":\"seizeVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferAllowed\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"cToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferVerify\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ComptrollerInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var ComptrollerInterfaceFuncSigs = map[string]string{
	"da3d454c": "borrowAllowed(address,address,uint256)",
	"5c778605": "borrowVerify(address,address,uint256)",
	"c2998238": "enterMarkets(address[])",
	"ede4edd0": "exitMarket(address)",
	"007e3dd2": "isComptroller()",
	"5fc7e71e": "liquidateBorrowAllowed(address,address,address,address,uint256)",
	"47ef3b3b": "liquidateBorrowVerify(address,address,address,address,uint256,uint256)",
	"c488847b": "liquidateCalculateSeizeTokens(address,address,uint256)",
	"4ef4c3e1": "mintAllowed(address,address,uint256)",
	"41c728b9": "mintVerify(address,address,uint256,uint256)",
	"eabe7d91": "redeemAllowed(address,address,uint256)",
	"51dff989": "redeemVerify(address,address,uint256,uint256)",
	"24008a62": "repayBorrowAllowed(address,address,address,uint256)",
	"1ededc91": "repayBorrowVerify(address,address,address,uint256,uint256)",
	"d02f7351": "seizeAllowed(address,address,address,address,uint256)",
	"6d35bf91": "seizeVerify(address,address,address,address,uint256)",
	"bdcdc258": "transferAllowed(address,address,address,uint256)",
	"6a56947e": "transferVerify(address,address,address,uint256)",
}

// ComptrollerInterface is an auto generated Go binding around an Ethereum contract.
type ComptrollerInterface struct {
	ComptrollerInterfaceCaller     // Read-only binding to the contract
	ComptrollerInterfaceTransactor // Write-only binding to the contract
	ComptrollerInterfaceFilterer   // Log filterer for contract events
}

// ComptrollerInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComptrollerInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComptrollerInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComptrollerInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComptrollerInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComptrollerInterfaceSession struct {
	Contract     *ComptrollerInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ComptrollerInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComptrollerInterfaceCallerSession struct {
	Contract *ComptrollerInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ComptrollerInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComptrollerInterfaceTransactorSession struct {
	Contract     *ComptrollerInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ComptrollerInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComptrollerInterfaceRaw struct {
	Contract *ComptrollerInterface // Generic contract binding to access the raw methods on
}

// ComptrollerInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComptrollerInterfaceCallerRaw struct {
	Contract *ComptrollerInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// ComptrollerInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComptrollerInterfaceTransactorRaw struct {
	Contract *ComptrollerInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComptrollerInterface creates a new instance of ComptrollerInterface, bound to a specific deployed contract.
func NewComptrollerInterface(address common.Address, backend bind.ContractBackend) (*ComptrollerInterface, error) {
	contract, err := bindComptrollerInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ComptrollerInterface{ComptrollerInterfaceCaller: ComptrollerInterfaceCaller{contract: contract}, ComptrollerInterfaceTransactor: ComptrollerInterfaceTransactor{contract: contract}, ComptrollerInterfaceFilterer: ComptrollerInterfaceFilterer{contract: contract}}, nil
}

// NewComptrollerInterfaceCaller creates a new read-only instance of ComptrollerInterface, bound to a specific deployed contract.
func NewComptrollerInterfaceCaller(address common.Address, caller bind.ContractCaller) (*ComptrollerInterfaceCaller, error) {
	contract, err := bindComptrollerInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerInterfaceCaller{contract: contract}, nil
}

// NewComptrollerInterfaceTransactor creates a new write-only instance of ComptrollerInterface, bound to a specific deployed contract.
func NewComptrollerInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*ComptrollerInterfaceTransactor, error) {
	contract, err := bindComptrollerInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComptrollerInterfaceTransactor{contract: contract}, nil
}

// NewComptrollerInterfaceFilterer creates a new log filterer instance of ComptrollerInterface, bound to a specific deployed contract.
func NewComptrollerInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*ComptrollerInterfaceFilterer, error) {
	contract, err := bindComptrollerInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComptrollerInterfaceFilterer{contract: contract}, nil
}

// bindComptrollerInterface binds a generic wrapper to an already deployed contract.
func bindComptrollerInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComptrollerInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerInterface *ComptrollerInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerInterface.Contract.ComptrollerInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerInterface *ComptrollerInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.ComptrollerInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerInterface *ComptrollerInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.ComptrollerInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ComptrollerInterface *ComptrollerInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ComptrollerInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ComptrollerInterface *ComptrollerInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ComptrollerInterface *ComptrollerInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.contract.Transact(opts, method, params...)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerInterface *ComptrollerInterfaceCaller) IsComptroller(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ComptrollerInterface.contract.Call(opts, &out, "isComptroller")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerInterface *ComptrollerInterfaceSession) IsComptroller() (bool, error) {
	return _ComptrollerInterface.Contract.IsComptroller(&_ComptrollerInterface.CallOpts)
}

// IsComptroller is a free data retrieval call binding the contract method 0x007e3dd2.
//
// Solidity: function isComptroller() view returns(bool)
func (_ComptrollerInterface *ComptrollerInterfaceCallerSession) IsComptroller() (bool, error) {
	return _ComptrollerInterface.Contract.IsComptroller(&_ComptrollerInterface.CallOpts)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_ComptrollerInterface *ComptrollerInterfaceCaller) LiquidateCalculateSeizeTokens(opts *bind.CallOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ComptrollerInterface.contract.Call(opts, &out, "liquidateCalculateSeizeTokens", cTokenBorrowed, cTokenCollateral, repayAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ComptrollerInterface.Contract.LiquidateCalculateSeizeTokens(&_ComptrollerInterface.CallOpts, cTokenBorrowed, cTokenCollateral, repayAmount)
}

// LiquidateCalculateSeizeTokens is a free data retrieval call binding the contract method 0xc488847b.
//
// Solidity: function liquidateCalculateSeizeTokens(address cTokenBorrowed, address cTokenCollateral, uint256 repayAmount) view returns(uint256, uint256)
func (_ComptrollerInterface *ComptrollerInterfaceCallerSession) LiquidateCalculateSeizeTokens(cTokenBorrowed common.Address, cTokenCollateral common.Address, repayAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ComptrollerInterface.Contract.LiquidateCalculateSeizeTokens(&_ComptrollerInterface.CallOpts, cTokenBorrowed, cTokenCollateral, repayAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) BorrowAllowed(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "borrowAllowed", cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.BorrowAllowed(&_ComptrollerInterface.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowAllowed is a paid mutator transaction binding the contract method 0xda3d454c.
//
// Solidity: function borrowAllowed(address cToken, address borrower, uint256 borrowAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) BorrowAllowed(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.BorrowAllowed(&_ComptrollerInterface.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) BorrowVerify(opts *bind.TransactOpts, cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "borrowVerify", cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.BorrowVerify(&_ComptrollerInterface.TransactOpts, cToken, borrower, borrowAmount)
}

// BorrowVerify is a paid mutator transaction binding the contract method 0x5c778605.
//
// Solidity: function borrowVerify(address cToken, address borrower, uint256 borrowAmount) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) BorrowVerify(cToken common.Address, borrower common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.BorrowVerify(&_ComptrollerInterface.TransactOpts, cToken, borrower, borrowAmount)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) EnterMarkets(opts *bind.TransactOpts, cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "enterMarkets", cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerInterface *ComptrollerInterfaceSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.EnterMarkets(&_ComptrollerInterface.TransactOpts, cTokens)
}

// EnterMarkets is a paid mutator transaction binding the contract method 0xc2998238.
//
// Solidity: function enterMarkets(address[] cTokens) returns(uint256[])
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) EnterMarkets(cTokens []common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.EnterMarkets(&_ComptrollerInterface.TransactOpts, cTokens)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cToken) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) ExitMarket(opts *bind.TransactOpts, cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "exitMarket", cToken)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cToken) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) ExitMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.ExitMarket(&_ComptrollerInterface.TransactOpts, cToken)
}

// ExitMarket is a paid mutator transaction binding the contract method 0xede4edd0.
//
// Solidity: function exitMarket(address cToken) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) ExitMarket(cToken common.Address) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.ExitMarket(&_ComptrollerInterface.TransactOpts, cToken)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) LiquidateBorrowAllowed(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "liquidateBorrowAllowed", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.LiquidateBorrowAllowed(&_ComptrollerInterface.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowAllowed is a paid mutator transaction binding the contract method 0x5fc7e71e.
//
// Solidity: function liquidateBorrowAllowed(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) LiquidateBorrowAllowed(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.LiquidateBorrowAllowed(&_ComptrollerInterface.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) LiquidateBorrowVerify(opts *bind.TransactOpts, cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "liquidateBorrowVerify", cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.LiquidateBorrowVerify(&_ComptrollerInterface.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// LiquidateBorrowVerify is a paid mutator transaction binding the contract method 0x47ef3b3b.
//
// Solidity: function liquidateBorrowVerify(address cTokenBorrowed, address cTokenCollateral, address liquidator, address borrower, uint256 repayAmount, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) LiquidateBorrowVerify(cTokenBorrowed common.Address, cTokenCollateral common.Address, liquidator common.Address, borrower common.Address, repayAmount *big.Int, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.LiquidateBorrowVerify(&_ComptrollerInterface.TransactOpts, cTokenBorrowed, cTokenCollateral, liquidator, borrower, repayAmount, seizeTokens)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) MintAllowed(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "mintAllowed", cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.MintAllowed(&_ComptrollerInterface.TransactOpts, cToken, minter, mintAmount)
}

// MintAllowed is a paid mutator transaction binding the contract method 0x4ef4c3e1.
//
// Solidity: function mintAllowed(address cToken, address minter, uint256 mintAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) MintAllowed(cToken common.Address, minter common.Address, mintAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.MintAllowed(&_ComptrollerInterface.TransactOpts, cToken, minter, mintAmount)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) MintVerify(opts *bind.TransactOpts, cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "mintVerify", cToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) MintVerify(cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.MintVerify(&_ComptrollerInterface.TransactOpts, cToken, minter, mintAmount, mintTokens)
}

// MintVerify is a paid mutator transaction binding the contract method 0x41c728b9.
//
// Solidity: function mintVerify(address cToken, address minter, uint256 mintAmount, uint256 mintTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) MintVerify(cToken common.Address, minter common.Address, mintAmount *big.Int, mintTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.MintVerify(&_ComptrollerInterface.TransactOpts, cToken, minter, mintAmount, mintTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) RedeemAllowed(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "redeemAllowed", cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RedeemAllowed(&_ComptrollerInterface.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemAllowed is a paid mutator transaction binding the contract method 0xeabe7d91.
//
// Solidity: function redeemAllowed(address cToken, address redeemer, uint256 redeemTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) RedeemAllowed(cToken common.Address, redeemer common.Address, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RedeemAllowed(&_ComptrollerInterface.TransactOpts, cToken, redeemer, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) RedeemVerify(opts *bind.TransactOpts, cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "redeemVerify", cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RedeemVerify(&_ComptrollerInterface.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RedeemVerify is a paid mutator transaction binding the contract method 0x51dff989.
//
// Solidity: function redeemVerify(address cToken, address redeemer, uint256 redeemAmount, uint256 redeemTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) RedeemVerify(cToken common.Address, redeemer common.Address, redeemAmount *big.Int, redeemTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RedeemVerify(&_ComptrollerInterface.TransactOpts, cToken, redeemer, redeemAmount, redeemTokens)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) RepayBorrowAllowed(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "repayBorrowAllowed", cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RepayBorrowAllowed(&_ComptrollerInterface.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowAllowed is a paid mutator transaction binding the contract method 0x24008a62.
//
// Solidity: function repayBorrowAllowed(address cToken, address payer, address borrower, uint256 repayAmount) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) RepayBorrowAllowed(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RepayBorrowAllowed(&_ComptrollerInterface.TransactOpts, cToken, payer, borrower, repayAmount)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) RepayBorrowVerify(opts *bind.TransactOpts, cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "repayBorrowVerify", cToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RepayBorrowVerify(&_ComptrollerInterface.TransactOpts, cToken, payer, borrower, repayAmount, borrowerIndex)
}

// RepayBorrowVerify is a paid mutator transaction binding the contract method 0x1ededc91.
//
// Solidity: function repayBorrowVerify(address cToken, address payer, address borrower, uint256 repayAmount, uint256 borrowerIndex) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) RepayBorrowVerify(cToken common.Address, payer common.Address, borrower common.Address, repayAmount *big.Int, borrowerIndex *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.RepayBorrowVerify(&_ComptrollerInterface.TransactOpts, cToken, payer, borrower, repayAmount, borrowerIndex)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) SeizeAllowed(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "seizeAllowed", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.SeizeAllowed(&_ComptrollerInterface.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeAllowed is a paid mutator transaction binding the contract method 0xd02f7351.
//
// Solidity: function seizeAllowed(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) SeizeAllowed(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.SeizeAllowed(&_ComptrollerInterface.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) SeizeVerify(opts *bind.TransactOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "seizeVerify", cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.SeizeVerify(&_ComptrollerInterface.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// SeizeVerify is a paid mutator transaction binding the contract method 0x6d35bf91.
//
// Solidity: function seizeVerify(address cTokenCollateral, address cTokenBorrowed, address liquidator, address borrower, uint256 seizeTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) SeizeVerify(cTokenCollateral common.Address, cTokenBorrowed common.Address, liquidator common.Address, borrower common.Address, seizeTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.SeizeVerify(&_ComptrollerInterface.TransactOpts, cTokenCollateral, cTokenBorrowed, liquidator, borrower, seizeTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) TransferAllowed(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "transferAllowed", cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.TransferAllowed(&_ComptrollerInterface.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferAllowed is a paid mutator transaction binding the contract method 0xbdcdc258.
//
// Solidity: function transferAllowed(address cToken, address src, address dst, uint256 transferTokens) returns(uint256)
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) TransferAllowed(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.TransferAllowed(&_ComptrollerInterface.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactor) TransferVerify(opts *bind.TransactOpts, cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.contract.Transact(opts, "transferVerify", cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.TransferVerify(&_ComptrollerInterface.TransactOpts, cToken, src, dst, transferTokens)
}

// TransferVerify is a paid mutator transaction binding the contract method 0x6a56947e.
//
// Solidity: function transferVerify(address cToken, address src, address dst, uint256 transferTokens) returns()
func (_ComptrollerInterface *ComptrollerInterfaceTransactorSession) TransferVerify(cToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _ComptrollerInterface.Contract.TransferVerify(&_ComptrollerInterface.TransactOpts, cToken, src, dst, transferTokens)
}

// EIP20InterfaceABI is the input ABI used to generate the binding from.
const EIP20InterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EIP20InterfaceFuncSigs maps the 4-byte function signature to its string representation.
var EIP20InterfaceFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// EIP20Interface is an auto generated Go binding around an Ethereum contract.
type EIP20Interface struct {
	EIP20InterfaceCaller     // Read-only binding to the contract
	EIP20InterfaceTransactor // Write-only binding to the contract
	EIP20InterfaceFilterer   // Log filterer for contract events
}

// EIP20InterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP20InterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20InterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP20InterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20InterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP20InterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20InterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP20InterfaceSession struct {
	Contract     *EIP20Interface   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP20InterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP20InterfaceCallerSession struct {
	Contract *EIP20InterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// EIP20InterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP20InterfaceTransactorSession struct {
	Contract     *EIP20InterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// EIP20InterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP20InterfaceRaw struct {
	Contract *EIP20Interface // Generic contract binding to access the raw methods on
}

// EIP20InterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP20InterfaceCallerRaw struct {
	Contract *EIP20InterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// EIP20InterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP20InterfaceTransactorRaw struct {
	Contract *EIP20InterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP20Interface creates a new instance of EIP20Interface, bound to a specific deployed contract.
func NewEIP20Interface(address common.Address, backend bind.ContractBackend) (*EIP20Interface, error) {
	contract, err := bindEIP20Interface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP20Interface{EIP20InterfaceCaller: EIP20InterfaceCaller{contract: contract}, EIP20InterfaceTransactor: EIP20InterfaceTransactor{contract: contract}, EIP20InterfaceFilterer: EIP20InterfaceFilterer{contract: contract}}, nil
}

// NewEIP20InterfaceCaller creates a new read-only instance of EIP20Interface, bound to a specific deployed contract.
func NewEIP20InterfaceCaller(address common.Address, caller bind.ContractCaller) (*EIP20InterfaceCaller, error) {
	contract, err := bindEIP20Interface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20InterfaceCaller{contract: contract}, nil
}

// NewEIP20InterfaceTransactor creates a new write-only instance of EIP20Interface, bound to a specific deployed contract.
func NewEIP20InterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP20InterfaceTransactor, error) {
	contract, err := bindEIP20Interface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20InterfaceTransactor{contract: contract}, nil
}

// NewEIP20InterfaceFilterer creates a new log filterer instance of EIP20Interface, bound to a specific deployed contract.
func NewEIP20InterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP20InterfaceFilterer, error) {
	contract, err := bindEIP20Interface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP20InterfaceFilterer{contract: contract}, nil
}

// bindEIP20Interface binds a generic wrapper to an already deployed contract.
func bindEIP20Interface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP20InterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20Interface *EIP20InterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP20Interface.Contract.EIP20InterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20Interface *EIP20InterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20Interface.Contract.EIP20InterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20Interface *EIP20InterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20Interface.Contract.EIP20InterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20Interface *EIP20InterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP20Interface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20Interface *EIP20InterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20Interface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20Interface *EIP20InterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20Interface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20Interface *EIP20InterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EIP20Interface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20Interface *EIP20InterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EIP20Interface.Contract.Allowance(&_EIP20Interface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20Interface *EIP20InterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EIP20Interface.Contract.Allowance(&_EIP20Interface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20Interface *EIP20InterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EIP20Interface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20Interface *EIP20InterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EIP20Interface.Contract.BalanceOf(&_EIP20Interface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20Interface *EIP20InterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EIP20Interface.Contract.BalanceOf(&_EIP20Interface.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20Interface *EIP20InterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIP20Interface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20Interface *EIP20InterfaceSession) TotalSupply() (*big.Int, error) {
	return _EIP20Interface.Contract.TotalSupply(&_EIP20Interface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20Interface *EIP20InterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _EIP20Interface.Contract.TotalSupply(&_EIP20Interface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.Approve(&_EIP20Interface.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.Approve(&_EIP20Interface.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.Transfer(&_EIP20Interface.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.Transfer(&_EIP20Interface.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.TransferFrom(&_EIP20Interface.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns(bool success)
func (_EIP20Interface *EIP20InterfaceTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20Interface.Contract.TransferFrom(&_EIP20Interface.TransactOpts, src, dst, amount)
}

// EIP20InterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EIP20Interface contract.
type EIP20InterfaceApprovalIterator struct {
	Event *EIP20InterfaceApproval // Event containing the contract specifics and raw log

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
func (it *EIP20InterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20InterfaceApproval)
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
		it.Event = new(EIP20InterfaceApproval)
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
func (it *EIP20InterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20InterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20InterfaceApproval represents a Approval event raised by the EIP20Interface contract.
type EIP20InterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EIP20InterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20Interface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EIP20InterfaceApprovalIterator{contract: _EIP20Interface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EIP20InterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20Interface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20InterfaceApproval)
				if err := _EIP20Interface.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) ParseApproval(log types.Log) (*EIP20InterfaceApproval, error) {
	event := new(EIP20InterfaceApproval)
	if err := _EIP20Interface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EIP20InterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EIP20Interface contract.
type EIP20InterfaceTransferIterator struct {
	Event *EIP20InterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *EIP20InterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20InterfaceTransfer)
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
		it.Event = new(EIP20InterfaceTransfer)
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
func (it *EIP20InterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20InterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20InterfaceTransfer represents a Transfer event raised by the EIP20Interface contract.
type EIP20InterfaceTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EIP20InterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20Interface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EIP20InterfaceTransferIterator{contract: _EIP20Interface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EIP20InterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20Interface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20InterfaceTransfer)
				if err := _EIP20Interface.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20Interface *EIP20InterfaceFilterer) ParseTransfer(log types.Log) (*EIP20InterfaceTransfer, error) {
	event := new(EIP20InterfaceTransfer)
	if err := _EIP20Interface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EIP20NonStandardInterfaceABI is the input ABI used to generate the binding from.
const EIP20NonStandardInterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EIP20NonStandardInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var EIP20NonStandardInterfaceFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// EIP20NonStandardInterface is an auto generated Go binding around an Ethereum contract.
type EIP20NonStandardInterface struct {
	EIP20NonStandardInterfaceCaller     // Read-only binding to the contract
	EIP20NonStandardInterfaceTransactor // Write-only binding to the contract
	EIP20NonStandardInterfaceFilterer   // Log filterer for contract events
}

// EIP20NonStandardInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type EIP20NonStandardInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20NonStandardInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP20NonStandardInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20NonStandardInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP20NonStandardInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20NonStandardInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP20NonStandardInterfaceSession struct {
	Contract     *EIP20NonStandardInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts              // Call options to use throughout this session
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// EIP20NonStandardInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP20NonStandardInterfaceCallerSession struct {
	Contract *EIP20NonStandardInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                    // Call options to use throughout this session
}

// EIP20NonStandardInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP20NonStandardInterfaceTransactorSession struct {
	Contract     *EIP20NonStandardInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                    // Transaction auth options to use throughout this session
}

// EIP20NonStandardInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type EIP20NonStandardInterfaceRaw struct {
	Contract *EIP20NonStandardInterface // Generic contract binding to access the raw methods on
}

// EIP20NonStandardInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP20NonStandardInterfaceCallerRaw struct {
	Contract *EIP20NonStandardInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// EIP20NonStandardInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP20NonStandardInterfaceTransactorRaw struct {
	Contract *EIP20NonStandardInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP20NonStandardInterface creates a new instance of EIP20NonStandardInterface, bound to a specific deployed contract.
func NewEIP20NonStandardInterface(address common.Address, backend bind.ContractBackend) (*EIP20NonStandardInterface, error) {
	contract, err := bindEIP20NonStandardInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterface{EIP20NonStandardInterfaceCaller: EIP20NonStandardInterfaceCaller{contract: contract}, EIP20NonStandardInterfaceTransactor: EIP20NonStandardInterfaceTransactor{contract: contract}, EIP20NonStandardInterfaceFilterer: EIP20NonStandardInterfaceFilterer{contract: contract}}, nil
}

// NewEIP20NonStandardInterfaceCaller creates a new read-only instance of EIP20NonStandardInterface, bound to a specific deployed contract.
func NewEIP20NonStandardInterfaceCaller(address common.Address, caller bind.ContractCaller) (*EIP20NonStandardInterfaceCaller, error) {
	contract, err := bindEIP20NonStandardInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterfaceCaller{contract: contract}, nil
}

// NewEIP20NonStandardInterfaceTransactor creates a new write-only instance of EIP20NonStandardInterface, bound to a specific deployed contract.
func NewEIP20NonStandardInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*EIP20NonStandardInterfaceTransactor, error) {
	contract, err := bindEIP20NonStandardInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterfaceTransactor{contract: contract}, nil
}

// NewEIP20NonStandardInterfaceFilterer creates a new log filterer instance of EIP20NonStandardInterface, bound to a specific deployed contract.
func NewEIP20NonStandardInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*EIP20NonStandardInterfaceFilterer, error) {
	contract, err := bindEIP20NonStandardInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterfaceFilterer{contract: contract}, nil
}

// bindEIP20NonStandardInterface binds a generic wrapper to an already deployed contract.
func bindEIP20NonStandardInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP20NonStandardInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP20NonStandardInterface.Contract.EIP20NonStandardInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.EIP20NonStandardInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.EIP20NonStandardInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EIP20NonStandardInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EIP20NonStandardInterface.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.Allowance(&_EIP20NonStandardInterface.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.Allowance(&_EIP20NonStandardInterface.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _EIP20NonStandardInterface.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.BalanceOf(&_EIP20NonStandardInterface.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.BalanceOf(&_EIP20NonStandardInterface.CallOpts, owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _EIP20NonStandardInterface.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) TotalSupply() (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.TotalSupply(&_EIP20NonStandardInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _EIP20NonStandardInterface.Contract.TotalSupply(&_EIP20NonStandardInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.Approve(&_EIP20NonStandardInterface.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool success)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.Approve(&_EIP20NonStandardInterface.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.contract.Transact(opts, "transfer", dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.Transfer(&_EIP20NonStandardInterface.TransactOpts, dst, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactorSession) Transfer(dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.Transfer(&_EIP20NonStandardInterface.TransactOpts, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.contract.Transact(opts, "transferFrom", src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.TransferFrom(&_EIP20NonStandardInterface.TransactOpts, src, dst, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amount) returns()
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceTransactorSession) TransferFrom(src common.Address, dst common.Address, amount *big.Int) (*types.Transaction, error) {
	return _EIP20NonStandardInterface.Contract.TransferFrom(&_EIP20NonStandardInterface.TransactOpts, src, dst, amount)
}

// EIP20NonStandardInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EIP20NonStandardInterface contract.
type EIP20NonStandardInterfaceApprovalIterator struct {
	Event *EIP20NonStandardInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *EIP20NonStandardInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20NonStandardInterfaceApproval)
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
		it.Event = new(EIP20NonStandardInterfaceApproval)
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
func (it *EIP20NonStandardInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20NonStandardInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20NonStandardInterfaceApproval represents a Approval event raised by the EIP20NonStandardInterface contract.
type EIP20NonStandardInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EIP20NonStandardInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20NonStandardInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterfaceApprovalIterator{contract: _EIP20NonStandardInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EIP20NonStandardInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20NonStandardInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20NonStandardInterfaceApproval)
				if err := _EIP20NonStandardInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) ParseApproval(log types.Log) (*EIP20NonStandardInterfaceApproval, error) {
	event := new(EIP20NonStandardInterfaceApproval)
	if err := _EIP20NonStandardInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// EIP20NonStandardInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EIP20NonStandardInterface contract.
type EIP20NonStandardInterfaceTransferIterator struct {
	Event *EIP20NonStandardInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *EIP20NonStandardInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20NonStandardInterfaceTransfer)
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
		it.Event = new(EIP20NonStandardInterfaceTransfer)
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
func (it *EIP20NonStandardInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20NonStandardInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20NonStandardInterfaceTransfer represents a Transfer event raised by the EIP20NonStandardInterface contract.
type EIP20NonStandardInterfaceTransfer struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EIP20NonStandardInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20NonStandardInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EIP20NonStandardInterfaceTransferIterator{contract: _EIP20NonStandardInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EIP20NonStandardInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20NonStandardInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20NonStandardInterfaceTransfer)
				if err := _EIP20NonStandardInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed from, address indexed to, uint256 amount)
func (_EIP20NonStandardInterface *EIP20NonStandardInterfaceFilterer) ParseTransfer(log types.Log) (*EIP20NonStandardInterfaceTransfer, error) {
	event := new(EIP20NonStandardInterfaceTransfer)
	if err := _EIP20NonStandardInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ExponentialABI is the input ABI used to generate the binding from.
const ExponentialABI = "[]"

// ExponentialBin is the compiled bytecode used for deploying new contracts.
var ExponentialBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820e2661a48a1147c3815e7681c95eca8eb0639373da7bb76023894e2f4e96390a264736f6c63430005110032"

// DeployExponential deploys a new Ethereum contract, binding an instance of Exponential to it.
func DeployExponential(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Exponential, error) {
	parsed, err := abi.JSON(strings.NewReader(ExponentialABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ExponentialBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Exponential{ExponentialCaller: ExponentialCaller{contract: contract}, ExponentialTransactor: ExponentialTransactor{contract: contract}, ExponentialFilterer: ExponentialFilterer{contract: contract}}, nil
}

// Exponential is an auto generated Go binding around an Ethereum contract.
type Exponential struct {
	ExponentialCaller     // Read-only binding to the contract
	ExponentialTransactor // Write-only binding to the contract
	ExponentialFilterer   // Log filterer for contract events
}

// ExponentialCaller is an auto generated read-only Go binding around an Ethereum contract.
type ExponentialCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExponentialTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ExponentialTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExponentialFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ExponentialFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ExponentialSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ExponentialSession struct {
	Contract     *Exponential      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ExponentialCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ExponentialCallerSession struct {
	Contract *ExponentialCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ExponentialTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ExponentialTransactorSession struct {
	Contract     *ExponentialTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ExponentialRaw is an auto generated low-level Go binding around an Ethereum contract.
type ExponentialRaw struct {
	Contract *Exponential // Generic contract binding to access the raw methods on
}

// ExponentialCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ExponentialCallerRaw struct {
	Contract *ExponentialCaller // Generic read-only contract binding to access the raw methods on
}

// ExponentialTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ExponentialTransactorRaw struct {
	Contract *ExponentialTransactor // Generic write-only contract binding to access the raw methods on
}

// NewExponential creates a new instance of Exponential, bound to a specific deployed contract.
func NewExponential(address common.Address, backend bind.ContractBackend) (*Exponential, error) {
	contract, err := bindExponential(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Exponential{ExponentialCaller: ExponentialCaller{contract: contract}, ExponentialTransactor: ExponentialTransactor{contract: contract}, ExponentialFilterer: ExponentialFilterer{contract: contract}}, nil
}

// NewExponentialCaller creates a new read-only instance of Exponential, bound to a specific deployed contract.
func NewExponentialCaller(address common.Address, caller bind.ContractCaller) (*ExponentialCaller, error) {
	contract, err := bindExponential(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ExponentialCaller{contract: contract}, nil
}

// NewExponentialTransactor creates a new write-only instance of Exponential, bound to a specific deployed contract.
func NewExponentialTransactor(address common.Address, transactor bind.ContractTransactor) (*ExponentialTransactor, error) {
	contract, err := bindExponential(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ExponentialTransactor{contract: contract}, nil
}

// NewExponentialFilterer creates a new log filterer instance of Exponential, bound to a specific deployed contract.
func NewExponentialFilterer(address common.Address, filterer bind.ContractFilterer) (*ExponentialFilterer, error) {
	contract, err := bindExponential(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ExponentialFilterer{contract: contract}, nil
}

// bindExponential binds a generic wrapper to an already deployed contract.
func bindExponential(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ExponentialABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exponential *ExponentialRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exponential.Contract.ExponentialCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exponential *ExponentialRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exponential.Contract.ExponentialTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exponential *ExponentialRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exponential.Contract.ExponentialTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Exponential *ExponentialCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Exponential.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Exponential *ExponentialTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Exponential.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Exponential *ExponentialTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Exponential.Contract.contract.Transact(opts, method, params...)
}

// InterestRateModelABI is the input ABI used to generate the binding from.
const InterestRateModelABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"cash\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrows\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserves\",\"type\":\"uint256\"}],\"name\":\"getBorrowRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isInterestRateModel\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// InterestRateModelFuncSigs maps the 4-byte function signature to its string representation.
var InterestRateModelFuncSigs = map[string]string{
	"15f24053": "getBorrowRate(uint256,uint256,uint256)",
	"2191f92a": "isInterestRateModel()",
}

// InterestRateModel is an auto generated Go binding around an Ethereum contract.
type InterestRateModel struct {
	InterestRateModelCaller     // Read-only binding to the contract
	InterestRateModelTransactor // Write-only binding to the contract
	InterestRateModelFilterer   // Log filterer for contract events
}

// InterestRateModelCaller is an auto generated read-only Go binding around an Ethereum contract.
type InterestRateModelCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelTransactor is an auto generated write-only Go binding around an Ethereum contract.
type InterestRateModelTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type InterestRateModelFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// InterestRateModelSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type InterestRateModelSession struct {
	Contract     *InterestRateModel // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// InterestRateModelCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type InterestRateModelCallerSession struct {
	Contract *InterestRateModelCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// InterestRateModelTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type InterestRateModelTransactorSession struct {
	Contract     *InterestRateModelTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// InterestRateModelRaw is an auto generated low-level Go binding around an Ethereum contract.
type InterestRateModelRaw struct {
	Contract *InterestRateModel // Generic contract binding to access the raw methods on
}

// InterestRateModelCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type InterestRateModelCallerRaw struct {
	Contract *InterestRateModelCaller // Generic read-only contract binding to access the raw methods on
}

// InterestRateModelTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type InterestRateModelTransactorRaw struct {
	Contract *InterestRateModelTransactor // Generic write-only contract binding to access the raw methods on
}

// NewInterestRateModel creates a new instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModel(address common.Address, backend bind.ContractBackend) (*InterestRateModel, error) {
	contract, err := bindInterestRateModel(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &InterestRateModel{InterestRateModelCaller: InterestRateModelCaller{contract: contract}, InterestRateModelTransactor: InterestRateModelTransactor{contract: contract}, InterestRateModelFilterer: InterestRateModelFilterer{contract: contract}}, nil
}

// NewInterestRateModelCaller creates a new read-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelCaller(address common.Address, caller bind.ContractCaller) (*InterestRateModelCaller, error) {
	contract, err := bindInterestRateModel(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelCaller{contract: contract}, nil
}

// NewInterestRateModelTransactor creates a new write-only instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelTransactor(address common.Address, transactor bind.ContractTransactor) (*InterestRateModelTransactor, error) {
	contract, err := bindInterestRateModel(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelTransactor{contract: contract}, nil
}

// NewInterestRateModelFilterer creates a new log filterer instance of InterestRateModel, bound to a specific deployed contract.
func NewInterestRateModelFilterer(address common.Address, filterer bind.ContractFilterer) (*InterestRateModelFilterer, error) {
	contract, err := bindInterestRateModel(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &InterestRateModelFilterer{contract: contract}, nil
}

// bindInterestRateModel binds a generic wrapper to an already deployed contract.
func bindInterestRateModel(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(InterestRateModelABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.InterestRateModelCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.InterestRateModelTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_InterestRateModel *InterestRateModelCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _InterestRateModel.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_InterestRateModel *InterestRateModelTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _InterestRateModel.Contract.contract.Transact(opts, method, params...)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256, uint256)
func (_InterestRateModel *InterestRateModelCaller) GetBorrowRate(opts *bind.CallOpts, cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "getBorrowRate", cash, borrows, reserves)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256, uint256)
func (_InterestRateModel *InterestRateModelSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, *big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// GetBorrowRate is a free data retrieval call binding the contract method 0x15f24053.
//
// Solidity: function getBorrowRate(uint256 cash, uint256 borrows, uint256 reserves) view returns(uint256, uint256)
func (_InterestRateModel *InterestRateModelCallerSession) GetBorrowRate(cash *big.Int, borrows *big.Int, reserves *big.Int) (*big.Int, *big.Int, error) {
	return _InterestRateModel.Contract.GetBorrowRate(&_InterestRateModel.CallOpts, cash, borrows, reserves)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_InterestRateModel *InterestRateModelCaller) IsInterestRateModel(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _InterestRateModel.contract.Call(opts, &out, "isInterestRateModel")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_InterestRateModel *InterestRateModelSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// IsInterestRateModel is a free data retrieval call binding the contract method 0x2191f92a.
//
// Solidity: function isInterestRateModel() view returns(bool)
func (_InterestRateModel *InterestRateModelCallerSession) IsInterestRateModel() (bool, error) {
	return _InterestRateModel.Contract.IsInterestRateModel(&_InterestRateModel.CallOpts)
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// TokenErrorReporterABI is the input ABI used to generate the binding from.
const TokenErrorReporterABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"error\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"info\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"detail\",\"type\":\"uint256\"}],\"name\":\"Failure\",\"type\":\"event\"}]"

// TokenErrorReporterBin is the compiled bytecode used for deploying new contracts.
var TokenErrorReporterBin = "0x6080604052348015600f57600080fd5b50603e80601d6000396000f3fe6080604052600080fdfea265627a7a72315820f036ff9175528ee0696cf4099e57d0f10d1bcc5eb05cc40a9611951785cc1f9364736f6c63430005110032"

// DeployTokenErrorReporter deploys a new Ethereum contract, binding an instance of TokenErrorReporter to it.
func DeployTokenErrorReporter(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *TokenErrorReporter, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenErrorReporterABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(TokenErrorReporterBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &TokenErrorReporter{TokenErrorReporterCaller: TokenErrorReporterCaller{contract: contract}, TokenErrorReporterTransactor: TokenErrorReporterTransactor{contract: contract}, TokenErrorReporterFilterer: TokenErrorReporterFilterer{contract: contract}}, nil
}

// TokenErrorReporter is an auto generated Go binding around an Ethereum contract.
type TokenErrorReporter struct {
	TokenErrorReporterCaller     // Read-only binding to the contract
	TokenErrorReporterTransactor // Write-only binding to the contract
	TokenErrorReporterFilterer   // Log filterer for contract events
}

// TokenErrorReporterCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenErrorReporterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErrorReporterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenErrorReporterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErrorReporterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenErrorReporterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenErrorReporterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenErrorReporterSession struct {
	Contract     *TokenErrorReporter // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TokenErrorReporterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenErrorReporterCallerSession struct {
	Contract *TokenErrorReporterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// TokenErrorReporterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenErrorReporterTransactorSession struct {
	Contract     *TokenErrorReporterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// TokenErrorReporterRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenErrorReporterRaw struct {
	Contract *TokenErrorReporter // Generic contract binding to access the raw methods on
}

// TokenErrorReporterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenErrorReporterCallerRaw struct {
	Contract *TokenErrorReporterCaller // Generic read-only contract binding to access the raw methods on
}

// TokenErrorReporterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenErrorReporterTransactorRaw struct {
	Contract *TokenErrorReporterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenErrorReporter creates a new instance of TokenErrorReporter, bound to a specific deployed contract.
func NewTokenErrorReporter(address common.Address, backend bind.ContractBackend) (*TokenErrorReporter, error) {
	contract, err := bindTokenErrorReporter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenErrorReporter{TokenErrorReporterCaller: TokenErrorReporterCaller{contract: contract}, TokenErrorReporterTransactor: TokenErrorReporterTransactor{contract: contract}, TokenErrorReporterFilterer: TokenErrorReporterFilterer{contract: contract}}, nil
}

// NewTokenErrorReporterCaller creates a new read-only instance of TokenErrorReporter, bound to a specific deployed contract.
func NewTokenErrorReporterCaller(address common.Address, caller bind.ContractCaller) (*TokenErrorReporterCaller, error) {
	contract, err := bindTokenErrorReporter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenErrorReporterCaller{contract: contract}, nil
}

// NewTokenErrorReporterTransactor creates a new write-only instance of TokenErrorReporter, bound to a specific deployed contract.
func NewTokenErrorReporterTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenErrorReporterTransactor, error) {
	contract, err := bindTokenErrorReporter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenErrorReporterTransactor{contract: contract}, nil
}

// NewTokenErrorReporterFilterer creates a new log filterer instance of TokenErrorReporter, bound to a specific deployed contract.
func NewTokenErrorReporterFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenErrorReporterFilterer, error) {
	contract, err := bindTokenErrorReporter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenErrorReporterFilterer{contract: contract}, nil
}

// bindTokenErrorReporter binds a generic wrapper to an already deployed contract.
func bindTokenErrorReporter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenErrorReporterABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenErrorReporter *TokenErrorReporterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenErrorReporter.Contract.TokenErrorReporterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenErrorReporter *TokenErrorReporterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenErrorReporter.Contract.TokenErrorReporterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenErrorReporter *TokenErrorReporterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenErrorReporter.Contract.TokenErrorReporterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenErrorReporter *TokenErrorReporterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenErrorReporter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenErrorReporter *TokenErrorReporterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenErrorReporter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenErrorReporter *TokenErrorReporterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenErrorReporter.Contract.contract.Transact(opts, method, params...)
}

// TokenErrorReporterFailureIterator is returned from FilterFailure and is used to iterate over the raw logs and unpacked data for Failure events raised by the TokenErrorReporter contract.
type TokenErrorReporterFailureIterator struct {
	Event *TokenErrorReporterFailure // Event containing the contract specifics and raw log

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
func (it *TokenErrorReporterFailureIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TokenErrorReporterFailure)
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
		it.Event = new(TokenErrorReporterFailure)
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
func (it *TokenErrorReporterFailureIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TokenErrorReporterFailureIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TokenErrorReporterFailure represents a Failure event raised by the TokenErrorReporter contract.
type TokenErrorReporterFailure struct {
	Error  *big.Int
	Info   *big.Int
	Detail *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFailure is a free log retrieval operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_TokenErrorReporter *TokenErrorReporterFilterer) FilterFailure(opts *bind.FilterOpts) (*TokenErrorReporterFailureIterator, error) {

	logs, sub, err := _TokenErrorReporter.contract.FilterLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return &TokenErrorReporterFailureIterator{contract: _TokenErrorReporter.contract, event: "Failure", logs: logs, sub: sub}, nil
}

// WatchFailure is a free log subscription operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_TokenErrorReporter *TokenErrorReporterFilterer) WatchFailure(opts *bind.WatchOpts, sink chan<- *TokenErrorReporterFailure) (event.Subscription, error) {

	logs, sub, err := _TokenErrorReporter.contract.WatchLogs(opts, "Failure")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TokenErrorReporterFailure)
				if err := _TokenErrorReporter.contract.UnpackLog(event, "Failure", log); err != nil {
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

// ParseFailure is a log parse operation binding the contract event 0x45b96fe442630264581b197e84bbada861235052c5a1aadfff9ea4e40a969aa0.
//
// Solidity: event Failure(uint256 error, uint256 info, uint256 detail)
func (_TokenErrorReporter *TokenErrorReporterFilterer) ParseFailure(log types.Log) (*TokenErrorReporterFailure, error) {
	event := new(TokenErrorReporterFailure)
	if err := _TokenErrorReporter.contract.UnpackLog(event, "Failure", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
