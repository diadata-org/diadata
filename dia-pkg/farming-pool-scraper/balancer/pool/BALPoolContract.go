// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancerpoolcontract

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

// BalancerpoolcontractABI is the input ABI used to generate the binding from.
const BalancerpoolcontractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LOG_CALL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPriceSansFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"rebind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"public_\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Balancerpoolcontract is an auto generated Go binding around an Ethereum contract.
type Balancerpoolcontract struct {
	BalancerpoolcontractCaller     // Read-only binding to the contract
	BalancerpoolcontractTransactor // Write-only binding to the contract
	BalancerpoolcontractFilterer   // Log filterer for contract events
}

// BalancerpoolcontractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerpoolcontractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolcontractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerpoolcontractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolcontractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerpoolcontractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerpoolcontractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerpoolcontractSession struct {
	Contract     *Balancerpoolcontract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalancerpoolcontractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerpoolcontractCallerSession struct {
	Contract *BalancerpoolcontractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BalancerpoolcontractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerpoolcontractTransactorSession struct {
	Contract     *BalancerpoolcontractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BalancerpoolcontractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerpoolcontractRaw struct {
	Contract *Balancerpoolcontract // Generic contract binding to access the raw methods on
}

// BalancerpoolcontractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerpoolcontractCallerRaw struct {
	Contract *BalancerpoolcontractCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerpoolcontractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerpoolcontractTransactorRaw struct {
	Contract *BalancerpoolcontractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerpoolcontract creates a new instance of Balancerpoolcontract, bound to a specific deployed contract.
func NewBalancerpoolcontract(address common.Address, backend bind.ContractBackend) (*Balancerpoolcontract, error) {
	contract, err := bindBalancerpoolcontract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancerpoolcontract{BalancerpoolcontractCaller: BalancerpoolcontractCaller{contract: contract}, BalancerpoolcontractTransactor: BalancerpoolcontractTransactor{contract: contract}, BalancerpoolcontractFilterer: BalancerpoolcontractFilterer{contract: contract}}, nil
}

// NewBalancerpoolcontractCaller creates a new read-only instance of Balancerpoolcontract, bound to a specific deployed contract.
func NewBalancerpoolcontractCaller(address common.Address, caller bind.ContractCaller) (*BalancerpoolcontractCaller, error) {
	contract, err := bindBalancerpoolcontract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractCaller{contract: contract}, nil
}

// NewBalancerpoolcontractTransactor creates a new write-only instance of Balancerpoolcontract, bound to a specific deployed contract.
func NewBalancerpoolcontractTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerpoolcontractTransactor, error) {
	contract, err := bindBalancerpoolcontract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractTransactor{contract: contract}, nil
}

// NewBalancerpoolcontractFilterer creates a new log filterer instance of Balancerpoolcontract, bound to a specific deployed contract.
func NewBalancerpoolcontractFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerpoolcontractFilterer, error) {
	contract, err := bindBalancerpoolcontract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractFilterer{contract: contract}, nil
}

// bindBalancerpoolcontract binds a generic wrapper to an already deployed contract.
func bindBalancerpoolcontract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerpoolcontractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerpoolcontract *BalancerpoolcontractRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerpoolcontract.Contract.BalancerpoolcontractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerpoolcontract *BalancerpoolcontractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.BalancerpoolcontractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerpoolcontract *BalancerpoolcontractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.BalancerpoolcontractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancerpoolcontract *BalancerpoolcontractCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Balancerpoolcontract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancerpoolcontract *BalancerpoolcontractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancerpoolcontract *BalancerpoolcontractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "BONE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) BONE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BONE(&_Balancerpoolcontract.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) BONE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BONE(&_Balancerpoolcontract.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "BPOW_PRECISION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) BPOWPRECISION() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BPOWPRECISION(&_Balancerpoolcontract.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BPOWPRECISION(&_Balancerpoolcontract.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "EXIT_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) EXITFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.EXITFEE(&_Balancerpoolcontract.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) EXITFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.EXITFEE(&_Balancerpoolcontract.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "INIT_POOL_SUPPLY")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.INITPOOLSUPPLY(&_Balancerpoolcontract.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.INITPOOLSUPPLY(&_Balancerpoolcontract.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXBOUNDTOKENS(&_Balancerpoolcontract.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXBOUNDTOKENS(&_Balancerpoolcontract.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXBPOWBASE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXBPOWBASE(&_Balancerpoolcontract.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXBPOWBASE(&_Balancerpoolcontract.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXFEE(&_Balancerpoolcontract.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXFEE(&_Balancerpoolcontract.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_IN_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXINRATIO() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXINRATIO(&_Balancerpoolcontract.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXINRATIO() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXINRATIO(&_Balancerpoolcontract.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_OUT_RATIO")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXOUTRATIO() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXOUTRATIO(&_Balancerpoolcontract.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXOUTRATIO(&_Balancerpoolcontract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_TOTAL_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXTOTALWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXTOTALWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MAX_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MAXWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MAXWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MIN_BALANCE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MINBALANCE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBALANCE(&_Balancerpoolcontract.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MINBALANCE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBALANCE(&_Balancerpoolcontract.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MIN_BOUND_TOKENS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBOUNDTOKENS(&_Balancerpoolcontract.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBOUNDTOKENS(&_Balancerpoolcontract.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MIN_BPOW_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MINBPOWBASE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBPOWBASE(&_Balancerpoolcontract.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINBPOWBASE(&_Balancerpoolcontract.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MIN_FEE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MINFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINFEE(&_Balancerpoolcontract.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MINFEE() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINFEE(&_Balancerpoolcontract.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "MIN_WEIGHT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) MINWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) MINWEIGHT() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.MINWEIGHT(&_Balancerpoolcontract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "allowance", src, dst)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.Allowance(&_Balancerpoolcontract.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.Allowance(&_Balancerpoolcontract.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "balanceOf", whom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BalanceOf(&_Balancerpoolcontract.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.BalanceOf(&_Balancerpoolcontract.CallOpts, whom)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcInGivenOut(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcInGivenOut(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcOutGivenIn(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcOutGivenIn(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcPoolInGivenSingleOut(&_Balancerpoolcontract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcPoolInGivenSingleOut(&_Balancerpoolcontract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcPoolOutGivenSingleIn(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcPoolOutGivenSingleIn(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSingleInGivenPoolOut(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSingleInGivenPoolOut(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSingleOutGivenPoolIn(&_Balancerpoolcontract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSingleOutGivenPoolIn(&_Balancerpoolcontract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSpotPrice(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.CalcSpotPrice(&_Balancerpoolcontract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Decimals() (uint8, error) {
	return _Balancerpoolcontract.Contract.Decimals(&_Balancerpoolcontract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) Decimals() (uint8, error) {
	return _Balancerpoolcontract.Contract.Decimals(&_Balancerpoolcontract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getBalance", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetBalance(&_Balancerpoolcontract.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetBalance(&_Balancerpoolcontract.CallOpts, token)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getColor")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetColor() ([32]byte, error) {
	return _Balancerpoolcontract.Contract.GetColor(&_Balancerpoolcontract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetColor() ([32]byte, error) {
	return _Balancerpoolcontract.Contract.GetColor(&_Balancerpoolcontract.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetController() (common.Address, error) {
	return _Balancerpoolcontract.Contract.GetController(&_Balancerpoolcontract.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetController() (common.Address, error) {
	return _Balancerpoolcontract.Contract.GetController(&_Balancerpoolcontract.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getCurrentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetCurrentTokens() ([]common.Address, error) {
	return _Balancerpoolcontract.Contract.GetCurrentTokens(&_Balancerpoolcontract.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _Balancerpoolcontract.Contract.GetCurrentTokens(&_Balancerpoolcontract.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getDenormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetDenormalizedWeight(&_Balancerpoolcontract.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetDenormalizedWeight(&_Balancerpoolcontract.CallOpts, token)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetFinalTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getFinalTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetFinalTokens() ([]common.Address, error) {
	return _Balancerpoolcontract.Contract.GetFinalTokens(&_Balancerpoolcontract.CallOpts)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetFinalTokens() ([]common.Address, error) {
	return _Balancerpoolcontract.Contract.GetFinalTokens(&_Balancerpoolcontract.CallOpts)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetNormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getNormalizedWeight", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetNormalizedWeight(&_Balancerpoolcontract.CallOpts, token)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetNormalizedWeight(&_Balancerpoolcontract.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getNumTokens")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetNumTokens() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetNumTokens(&_Balancerpoolcontract.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetNumTokens() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetNumTokens(&_Balancerpoolcontract.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getSpotPrice", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSpotPrice(&_Balancerpoolcontract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSpotPrice(&_Balancerpoolcontract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetSpotPriceSansFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getSpotPriceSansFee", tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSpotPriceSansFee(&_Balancerpoolcontract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSpotPriceSansFee(&_Balancerpoolcontract.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getSwapFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetSwapFee() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSwapFee(&_Balancerpoolcontract.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetSwapFee() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetSwapFee(&_Balancerpoolcontract.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "getTotalDenormalizedWeight")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetTotalDenormalizedWeight(&_Balancerpoolcontract.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.GetTotalDenormalizedWeight(&_Balancerpoolcontract.CallOpts)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "isBound", t)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) IsBound(t common.Address) (bool, error) {
	return _Balancerpoolcontract.Contract.IsBound(&_Balancerpoolcontract.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) IsBound(t common.Address) (bool, error) {
	return _Balancerpoolcontract.Contract.IsBound(&_Balancerpoolcontract.CallOpts, t)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "isFinalized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) IsFinalized() (bool, error) {
	return _Balancerpoolcontract.Contract.IsFinalized(&_Balancerpoolcontract.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) IsFinalized() (bool, error) {
	return _Balancerpoolcontract.Contract.IsFinalized(&_Balancerpoolcontract.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "isPublicSwap")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) IsPublicSwap() (bool, error) {
	return _Balancerpoolcontract.Contract.IsPublicSwap(&_Balancerpoolcontract.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) IsPublicSwap() (bool, error) {
	return _Balancerpoolcontract.Contract.IsPublicSwap(&_Balancerpoolcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Name() (string, error) {
	return _Balancerpoolcontract.Contract.Name(&_Balancerpoolcontract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) Name() (string, error) {
	return _Balancerpoolcontract.Contract.Name(&_Balancerpoolcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Symbol() (string, error) {
	return _Balancerpoolcontract.Contract.Symbol(&_Balancerpoolcontract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) Symbol() (string, error) {
	return _Balancerpoolcontract.Contract.Symbol(&_Balancerpoolcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Balancerpoolcontract.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractSession) TotalSupply() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.TotalSupply(&_Balancerpoolcontract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancerpoolcontract *BalancerpoolcontractCallerSession) TotalSupply() (*big.Int, error) {
	return _Balancerpoolcontract.Contract.TotalSupply(&_Balancerpoolcontract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Approve(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Approve(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Bind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "bind", token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Bind(&_Balancerpoolcontract.TransactOpts, token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Bind(&_Balancerpoolcontract.TransactOpts, token, balance, denorm)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.DecreaseApproval(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.DecreaseApproval(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitPool(&_Balancerpoolcontract.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitPool(&_Balancerpoolcontract.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitswapExternAmountOut(&_Balancerpoolcontract.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitswapExternAmountOut(&_Balancerpoolcontract.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitswapPoolAmountIn(&_Balancerpoolcontract.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.ExitswapPoolAmountIn(&_Balancerpoolcontract.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) Finalize() (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Finalize(&_Balancerpoolcontract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Finalize() (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Finalize(&_Balancerpoolcontract.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Gulp(&_Balancerpoolcontract.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Gulp(&_Balancerpoolcontract.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.IncreaseApproval(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.IncreaseApproval(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinPool(&_Balancerpoolcontract.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinPool(&_Balancerpoolcontract.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinswapExternAmountIn(&_Balancerpoolcontract.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinswapExternAmountIn(&_Balancerpoolcontract.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinswapPoolAmountOut(&_Balancerpoolcontract.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.JoinswapPoolAmountOut(&_Balancerpoolcontract.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Rebind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "rebind", token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Rebind(&_Balancerpoolcontract.TransactOpts, token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Rebind(&_Balancerpoolcontract.TransactOpts, token, balance, denorm)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) SetController(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "setController", manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetController(&_Balancerpoolcontract.TransactOpts, manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetController(&_Balancerpoolcontract.TransactOpts, manager)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) SetPublicSwap(opts *bind.TransactOpts, public_ bool) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "setPublicSwap", public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetPublicSwap(&_Balancerpoolcontract.TransactOpts, public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetPublicSwap(&_Balancerpoolcontract.TransactOpts, public_)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetSwapFee(&_Balancerpoolcontract.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SetSwapFee(&_Balancerpoolcontract.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SwapExactAmountIn(&_Balancerpoolcontract.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SwapExactAmountIn(&_Balancerpoolcontract.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SwapExactAmountOut(&_Balancerpoolcontract.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.SwapExactAmountOut(&_Balancerpoolcontract.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Transfer(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Transfer(&_Balancerpoolcontract.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.TransferFrom(&_Balancerpoolcontract.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.TransferFrom(&_Balancerpoolcontract.TransactOpts, src, dst, amt)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactor) Unbind(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.contract.Transact(opts, "unbind", token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Unbind(&_Balancerpoolcontract.TransactOpts, token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancerpoolcontract *BalancerpoolcontractTransactorSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _Balancerpoolcontract.Contract.Unbind(&_Balancerpoolcontract.TransactOpts, token)
}

// BalancerpoolcontractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Balancerpoolcontract contract.
type BalancerpoolcontractApprovalIterator struct {
	Event *BalancerpoolcontractApproval // Event containing the contract specifics and raw log

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
func (it *BalancerpoolcontractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolcontractApproval)
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
		it.Event = new(BalancerpoolcontractApproval)
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
func (it *BalancerpoolcontractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolcontractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolcontractApproval represents a Approval event raised by the Balancerpoolcontract contract.
type BalancerpoolcontractApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerpoolcontractApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractApprovalIterator{contract: _Balancerpoolcontract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerpoolcontractApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolcontractApproval)
				if err := _Balancerpoolcontract.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) ParseApproval(log types.Log) (*BalancerpoolcontractApproval, error) {
	event := new(BalancerpoolcontractApproval)
	if err := _Balancerpoolcontract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerpoolcontractLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGEXITIterator struct {
	Event *BalancerpoolcontractLOGEXIT // Event containing the contract specifics and raw log

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
func (it *BalancerpoolcontractLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolcontractLOGEXIT)
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
		it.Event = new(BalancerpoolcontractLOGEXIT)
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
func (it *BalancerpoolcontractLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolcontractLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolcontractLOGEXIT represents a LOGEXIT event raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*BalancerpoolcontractLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractLOGEXITIterator{contract: _Balancerpoolcontract.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *BalancerpoolcontractLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolcontractLOGEXIT)
				if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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

// ParseLOGEXIT is a log parse operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) ParseLOGEXIT(log types.Log) (*BalancerpoolcontractLOGEXIT, error) {
	event := new(BalancerpoolcontractLOGEXIT)
	if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerpoolcontractLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGJOINIterator struct {
	Event *BalancerpoolcontractLOGJOIN // Event containing the contract specifics and raw log

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
func (it *BalancerpoolcontractLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolcontractLOGJOIN)
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
		it.Event = new(BalancerpoolcontractLOGJOIN)
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
func (it *BalancerpoolcontractLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolcontractLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolcontractLOGJOIN represents a LOGJOIN event raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*BalancerpoolcontractLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractLOGJOINIterator{contract: _Balancerpoolcontract.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *BalancerpoolcontractLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolcontractLOGJOIN)
				if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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

// ParseLOGJOIN is a log parse operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) ParseLOGJOIN(log types.Log) (*BalancerpoolcontractLOGJOIN, error) {
	event := new(BalancerpoolcontractLOGJOIN)
	if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerpoolcontractLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGSWAPIterator struct {
	Event *BalancerpoolcontractLOGSWAP // Event containing the contract specifics and raw log

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
func (it *BalancerpoolcontractLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolcontractLOGSWAP)
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
		it.Event = new(BalancerpoolcontractLOGSWAP)
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
func (it *BalancerpoolcontractLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolcontractLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolcontractLOGSWAP represents a LOGSWAP event raised by the Balancerpoolcontract contract.
type BalancerpoolcontractLOGSWAP struct {
	Caller         common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	TokenAmountIn  *big.Int
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGSWAP is a free log retrieval operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*BalancerpoolcontractLOGSWAPIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractLOGSWAPIterator{contract: _Balancerpoolcontract.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *BalancerpoolcontractLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolcontractLOGSWAP)
				if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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

// ParseLOGSWAP is a log parse operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) ParseLOGSWAP(log types.Log) (*BalancerpoolcontractLOGSWAP, error) {
	event := new(BalancerpoolcontractLOGSWAP)
	if err := _Balancerpoolcontract.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerpoolcontractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Balancerpoolcontract contract.
type BalancerpoolcontractTransferIterator struct {
	Event *BalancerpoolcontractTransfer // Event containing the contract specifics and raw log

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
func (it *BalancerpoolcontractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerpoolcontractTransfer)
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
		it.Event = new(BalancerpoolcontractTransfer)
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
func (it *BalancerpoolcontractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerpoolcontractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerpoolcontractTransfer represents a Transfer event raised by the Balancerpoolcontract contract.
type BalancerpoolcontractTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerpoolcontractTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerpoolcontractTransferIterator{contract: _Balancerpoolcontract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerpoolcontractTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancerpoolcontract.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerpoolcontractTransfer)
				if err := _Balancerpoolcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Balancerpoolcontract *BalancerpoolcontractFilterer) ParseTransfer(log types.Log) (*BalancerpoolcontractTransfer, error) {
	event := new(BalancerpoolcontractTransfer)
	if err := _Balancerpoolcontract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
