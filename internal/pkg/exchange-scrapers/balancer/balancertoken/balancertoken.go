// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package balancertoken

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

// BalancertokenABI is the input ABI used to generate the binding from.
const BalancertokenABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LOG_CALL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPriceSansFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"rebind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"public_\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Balancertoken is an auto generated Go binding around an Ethereum contract.
type Balancertoken struct {
	BalancertokenCaller     // Read-only binding to the contract
	BalancertokenTransactor // Write-only binding to the contract
	BalancertokenFilterer   // Log filterer for contract events
}

// BalancertokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancertokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancertokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancertokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancertokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancertokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancertokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancertokenSession struct {
	Contract     *Balancertoken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BalancertokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancertokenCallerSession struct {
	Contract *BalancertokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// BalancertokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancertokenTransactorSession struct {
	Contract     *BalancertokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// BalancertokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancertokenRaw struct {
	Contract *Balancertoken // Generic contract binding to access the raw methods on
}

// BalancertokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancertokenCallerRaw struct {
	Contract *BalancertokenCaller // Generic read-only contract binding to access the raw methods on
}

// BalancertokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancertokenTransactorRaw struct {
	Contract *BalancertokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancertoken creates a new instance of Balancertoken, bound to a specific deployed contract.
func NewBalancertoken(address common.Address, backend bind.ContractBackend) (*Balancertoken, error) {
	contract, err := bindBalancertoken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Balancertoken{BalancertokenCaller: BalancertokenCaller{contract: contract}, BalancertokenTransactor: BalancertokenTransactor{contract: contract}, BalancertokenFilterer: BalancertokenFilterer{contract: contract}}, nil
}

// NewBalancertokenCaller creates a new read-only instance of Balancertoken, bound to a specific deployed contract.
func NewBalancertokenCaller(address common.Address, caller bind.ContractCaller) (*BalancertokenCaller, error) {
	contract, err := bindBalancertoken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancertokenCaller{contract: contract}, nil
}

// NewBalancertokenTransactor creates a new write-only instance of Balancertoken, bound to a specific deployed contract.
func NewBalancertokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancertokenTransactor, error) {
	contract, err := bindBalancertoken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancertokenTransactor{contract: contract}, nil
}

// NewBalancertokenFilterer creates a new log filterer instance of Balancertoken, bound to a specific deployed contract.
func NewBalancertokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancertokenFilterer, error) {
	contract, err := bindBalancertoken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancertokenFilterer{contract: contract}, nil
}

// bindBalancertoken binds a generic wrapper to an already deployed contract.
func bindBalancertoken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancertokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancertoken *BalancertokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Balancertoken.Contract.BalancertokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancertoken *BalancertokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancertoken.Contract.BalancertokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancertoken *BalancertokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancertoken.Contract.BalancertokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Balancertoken *BalancertokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Balancertoken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Balancertoken *BalancertokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancertoken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Balancertoken *BalancertokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Balancertoken.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "BONE")
	return *ret0, err
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) BONE() (*big.Int, error) {
	return _Balancertoken.Contract.BONE(&_Balancertoken.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) BONE() (*big.Int, error) {
	return _Balancertoken.Contract.BONE(&_Balancertoken.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "BPOW_PRECISION")
	return *ret0, err
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancertoken *BalancertokenSession) BPOWPRECISION() (*big.Int, error) {
	return _Balancertoken.Contract.BPOWPRECISION(&_Balancertoken.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _Balancertoken.Contract.BPOWPRECISION(&_Balancertoken.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "EXIT_FEE")
	return *ret0, err
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) EXITFEE() (*big.Int, error) {
	return _Balancertoken.Contract.EXITFEE(&_Balancertoken.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) EXITFEE() (*big.Int, error) {
	return _Balancertoken.Contract.EXITFEE(&_Balancertoken.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "INIT_POOL_SUPPLY")
	return *ret0, err
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancertoken *BalancertokenSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Balancertoken.Contract.INITPOOLSUPPLY(&_Balancertoken.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _Balancertoken.Contract.INITPOOLSUPPLY(&_Balancertoken.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_BOUND_TOKENS")
	return *ret0, err
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Balancertoken.Contract.MAXBOUNDTOKENS(&_Balancertoken.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _Balancertoken.Contract.MAXBOUNDTOKENS(&_Balancertoken.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_BPOW_BASE")
	return *ret0, err
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXBPOWBASE() (*big.Int, error) {
	return _Balancertoken.Contract.MAXBPOWBASE(&_Balancertoken.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _Balancertoken.Contract.MAXBPOWBASE(&_Balancertoken.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_FEE")
	return *ret0, err
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXFEE() (*big.Int, error) {
	return _Balancertoken.Contract.MAXFEE(&_Balancertoken.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXFEE() (*big.Int, error) {
	return _Balancertoken.Contract.MAXFEE(&_Balancertoken.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_IN_RATIO")
	return *ret0, err
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXINRATIO() (*big.Int, error) {
	return _Balancertoken.Contract.MAXINRATIO(&_Balancertoken.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXINRATIO() (*big.Int, error) {
	return _Balancertoken.Contract.MAXINRATIO(&_Balancertoken.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_OUT_RATIO")
	return *ret0, err
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXOUTRATIO() (*big.Int, error) {
	return _Balancertoken.Contract.MAXOUTRATIO(&_Balancertoken.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _Balancertoken.Contract.MAXOUTRATIO(&_Balancertoken.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_TOTAL_WEIGHT")
	return *ret0, err
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MAXTOTALWEIGHT(&_Balancertoken.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MAXTOTALWEIGHT(&_Balancertoken.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MAX_WEIGHT")
	return *ret0, err
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MAXWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MAXWEIGHT(&_Balancertoken.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MAXWEIGHT(&_Balancertoken.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MIN_BALANCE")
	return *ret0, err
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MINBALANCE() (*big.Int, error) {
	return _Balancertoken.Contract.MINBALANCE(&_Balancertoken.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MINBALANCE() (*big.Int, error) {
	return _Balancertoken.Contract.MINBALANCE(&_Balancertoken.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MIN_BOUND_TOKENS")
	return *ret0, err
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Balancertoken.Contract.MINBOUNDTOKENS(&_Balancertoken.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _Balancertoken.Contract.MINBOUNDTOKENS(&_Balancertoken.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MIN_BPOW_BASE")
	return *ret0, err
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MINBPOWBASE() (*big.Int, error) {
	return _Balancertoken.Contract.MINBPOWBASE(&_Balancertoken.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _Balancertoken.Contract.MINBPOWBASE(&_Balancertoken.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MIN_FEE")
	return *ret0, err
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MINFEE() (*big.Int, error) {
	return _Balancertoken.Contract.MINFEE(&_Balancertoken.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MINFEE() (*big.Int, error) {
	return _Balancertoken.Contract.MINFEE(&_Balancertoken.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "MIN_WEIGHT")
	return *ret0, err
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenSession) MINWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MINWEIGHT(&_Balancertoken.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) MINWEIGHT() (*big.Int, error) {
	return _Balancertoken.Contract.MINWEIGHT(&_Balancertoken.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancertoken *BalancertokenCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "allowance", src, dst)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancertoken *BalancertokenSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.Allowance(&_Balancertoken.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.Allowance(&_Balancertoken.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancertoken *BalancertokenCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "balanceOf", whom)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancertoken *BalancertokenSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.BalanceOf(&_Balancertoken.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.BalanceOf(&_Balancertoken.CallOpts, whom)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
	return *ret0, err
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcInGivenOut(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcInGivenOut(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
	return *ret0, err
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcOutGivenIn(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcOutGivenIn(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
	return *ret0, err
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcPoolInGivenSingleOut(&_Balancertoken.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcPoolInGivenSingleOut(&_Balancertoken.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
	return *ret0, err
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcPoolOutGivenSingleIn(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcPoolOutGivenSingleIn(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
	return *ret0, err
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSingleInGivenPoolOut(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSingleInGivenPoolOut(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
	return *ret0, err
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSingleOutGivenPoolIn(&_Balancertoken.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSingleOutGivenPoolIn(&_Balancertoken.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
	return *ret0, err
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSpotPrice(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _Balancertoken.Contract.CalcSpotPrice(&_Balancertoken.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancertoken *BalancertokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancertoken *BalancertokenSession) Decimals() (uint8, error) {
	return _Balancertoken.Contract.Decimals(&_Balancertoken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Balancertoken *BalancertokenCallerSession) Decimals() (uint8, error) {
	return _Balancertoken.Contract.Decimals(&_Balancertoken.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getBalance", token)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetBalance(&_Balancertoken.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetBalance(&_Balancertoken.CallOpts, token)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancertoken *BalancertokenCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancertoken *BalancertokenSession) GetColor() ([32]byte, error) {
	return _Balancertoken.Contract.GetColor(&_Balancertoken.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_Balancertoken *BalancertokenCallerSession) GetColor() ([32]byte, error) {
	return _Balancertoken.Contract.GetColor(&_Balancertoken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancertoken *BalancertokenCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancertoken *BalancertokenSession) GetController() (common.Address, error) {
	return _Balancertoken.Contract.GetController(&_Balancertoken.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_Balancertoken *BalancertokenCallerSession) GetController() (common.Address, error) {
	return _Balancertoken.Contract.GetController(&_Balancertoken.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getCurrentTokens")
	return *ret0, err
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenSession) GetCurrentTokens() ([]common.Address, error) {
	return _Balancertoken.Contract.GetCurrentTokens(&_Balancertoken.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _Balancertoken.Contract.GetCurrentTokens(&_Balancertoken.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getDenormalizedWeight", token)
	return *ret0, err
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetDenormalizedWeight(&_Balancertoken.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetDenormalizedWeight(&_Balancertoken.CallOpts, token)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenCaller) GetFinalTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getFinalTokens")
	return *ret0, err
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenSession) GetFinalTokens() ([]common.Address, error) {
	return _Balancertoken.Contract.GetFinalTokens(&_Balancertoken.CallOpts)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_Balancertoken *BalancertokenCallerSession) GetFinalTokens() ([]common.Address, error) {
	return _Balancertoken.Contract.GetFinalTokens(&_Balancertoken.CallOpts)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetNormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getNormalizedWeight", token)
	return *ret0, err
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetNormalizedWeight(&_Balancertoken.CallOpts, token)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetNormalizedWeight(&_Balancertoken.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getNumTokens")
	return *ret0, err
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetNumTokens() (*big.Int, error) {
	return _Balancertoken.Contract.GetNumTokens(&_Balancertoken.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetNumTokens() (*big.Int, error) {
	return _Balancertoken.Contract.GetNumTokens(&_Balancertoken.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getSpotPrice", tokenIn, tokenOut)
	return *ret0, err
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetSpotPrice(&_Balancertoken.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetSpotPrice(&_Balancertoken.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCaller) GetSpotPriceSansFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getSpotPriceSansFee", tokenIn, tokenOut)
	return *ret0, err
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetSpotPriceSansFee(&_Balancertoken.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_Balancertoken *BalancertokenCallerSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _Balancertoken.Contract.GetSpotPriceSansFee(&_Balancertoken.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getSwapFee")
	return *ret0, err
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetSwapFee() (*big.Int, error) {
	return _Balancertoken.Contract.GetSwapFee(&_Balancertoken.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetSwapFee() (*big.Int, error) {
	return _Balancertoken.Contract.GetSwapFee(&_Balancertoken.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "getTotalDenormalizedWeight")
	return *ret0, err
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancertoken *BalancertokenSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Balancertoken.Contract.GetTotalDenormalizedWeight(&_Balancertoken.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _Balancertoken.Contract.GetTotalDenormalizedWeight(&_Balancertoken.CallOpts)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancertoken *BalancertokenCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "isBound", t)
	return *ret0, err
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancertoken *BalancertokenSession) IsBound(t common.Address) (bool, error) {
	return _Balancertoken.Contract.IsBound(&_Balancertoken.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_Balancertoken *BalancertokenCallerSession) IsBound(t common.Address) (bool, error) {
	return _Balancertoken.Contract.IsBound(&_Balancertoken.CallOpts, t)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancertoken *BalancertokenCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancertoken *BalancertokenSession) IsFinalized() (bool, error) {
	return _Balancertoken.Contract.IsFinalized(&_Balancertoken.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_Balancertoken *BalancertokenCallerSession) IsFinalized() (bool, error) {
	return _Balancertoken.Contract.IsFinalized(&_Balancertoken.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancertoken *BalancertokenCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "isPublicSwap")
	return *ret0, err
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancertoken *BalancertokenSession) IsPublicSwap() (bool, error) {
	return _Balancertoken.Contract.IsPublicSwap(&_Balancertoken.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_Balancertoken *BalancertokenCallerSession) IsPublicSwap() (bool, error) {
	return _Balancertoken.Contract.IsPublicSwap(&_Balancertoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancertoken *BalancertokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancertoken *BalancertokenSession) Name() (string, error) {
	return _Balancertoken.Contract.Name(&_Balancertoken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Balancertoken *BalancertokenCallerSession) Name() (string, error) {
	return _Balancertoken.Contract.Name(&_Balancertoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancertoken *BalancertokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancertoken *BalancertokenSession) Symbol() (string, error) {
	return _Balancertoken.Contract.Symbol(&_Balancertoken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Balancertoken *BalancertokenCallerSession) Symbol() (string, error) {
	return _Balancertoken.Contract.Symbol(&_Balancertoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancertoken *BalancertokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Balancertoken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancertoken *BalancertokenSession) TotalSupply() (*big.Int, error) {
	return _Balancertoken.Contract.TotalSupply(&_Balancertoken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Balancertoken *BalancertokenCallerSession) TotalSupply() (*big.Int, error) {
	return _Balancertoken.Contract.TotalSupply(&_Balancertoken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Approve(&_Balancertoken.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Approve(&_Balancertoken.TransactOpts, dst, amt)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenTransactor) Bind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "bind", token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Bind(&_Balancertoken.TransactOpts, token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenTransactorSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Bind(&_Balancertoken.TransactOpts, token, balance, denorm)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.DecreaseApproval(&_Balancertoken.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.DecreaseApproval(&_Balancertoken.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancertoken *BalancertokenTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancertoken *BalancertokenSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitPool(&_Balancertoken.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_Balancertoken *BalancertokenTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitPool(&_Balancertoken.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitswapExternAmountOut(&_Balancertoken.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_Balancertoken *BalancertokenTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitswapExternAmountOut(&_Balancertoken.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitswapPoolAmountIn(&_Balancertoken.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.ExitswapPoolAmountIn(&_Balancertoken.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancertoken *BalancertokenTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancertoken *BalancertokenSession) Finalize() (*types.Transaction, error) {
	return _Balancertoken.Contract.Finalize(&_Balancertoken.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_Balancertoken *BalancertokenTransactorSession) Finalize() (*types.Transaction, error) {
	return _Balancertoken.Contract.Finalize(&_Balancertoken.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancertoken *BalancertokenTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancertoken *BalancertokenSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.Gulp(&_Balancertoken.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_Balancertoken *BalancertokenTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.Gulp(&_Balancertoken.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.IncreaseApproval(&_Balancertoken.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.IncreaseApproval(&_Balancertoken.TransactOpts, dst, amt)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancertoken *BalancertokenTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancertoken *BalancertokenSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinPool(&_Balancertoken.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_Balancertoken *BalancertokenTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinPool(&_Balancertoken.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinswapExternAmountIn(&_Balancertoken.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_Balancertoken *BalancertokenTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinswapExternAmountIn(&_Balancertoken.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinswapPoolAmountOut(&_Balancertoken.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.JoinswapPoolAmountOut(&_Balancertoken.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenTransactor) Rebind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "rebind", token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Rebind(&_Balancertoken.TransactOpts, token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_Balancertoken *BalancertokenTransactorSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Rebind(&_Balancertoken.TransactOpts, token, balance, denorm)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancertoken *BalancertokenTransactor) SetController(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "setController", manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancertoken *BalancertokenSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetController(&_Balancertoken.TransactOpts, manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_Balancertoken *BalancertokenTransactorSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetController(&_Balancertoken.TransactOpts, manager)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancertoken *BalancertokenTransactor) SetPublicSwap(opts *bind.TransactOpts, public_ bool) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "setPublicSwap", public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancertoken *BalancertokenSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetPublicSwap(&_Balancertoken.TransactOpts, public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_Balancertoken *BalancertokenTransactorSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetPublicSwap(&_Balancertoken.TransactOpts, public_)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancertoken *BalancertokenTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancertoken *BalancertokenSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetSwapFee(&_Balancertoken.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_Balancertoken *BalancertokenTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SetSwapFee(&_Balancertoken.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SwapExactAmountIn(&_Balancertoken.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SwapExactAmountIn(&_Balancertoken.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SwapExactAmountOut(&_Balancertoken.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_Balancertoken *BalancertokenTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.SwapExactAmountOut(&_Balancertoken.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Transfer(&_Balancertoken.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.Transfer(&_Balancertoken.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.TransferFrom(&_Balancertoken.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_Balancertoken *BalancertokenTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _Balancertoken.Contract.TransferFrom(&_Balancertoken.TransactOpts, src, dst, amt)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancertoken *BalancertokenTransactor) Unbind(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _Balancertoken.contract.Transact(opts, "unbind", token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancertoken *BalancertokenSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.Unbind(&_Balancertoken.TransactOpts, token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_Balancertoken *BalancertokenTransactorSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _Balancertoken.Contract.Unbind(&_Balancertoken.TransactOpts, token)
}

// BalancertokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Balancertoken contract.
type BalancertokenApprovalIterator struct {
	Event *BalancertokenApproval // Event containing the contract specifics and raw log

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
func (it *BalancertokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancertokenApproval)
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
		it.Event = new(BalancertokenApproval)
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
func (it *BalancertokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancertokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancertokenApproval represents a Approval event raised by the Balancertoken contract.
type BalancertokenApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Balancertoken *BalancertokenFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancertokenApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancertoken.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancertokenApprovalIterator{contract: _Balancertoken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_Balancertoken *BalancertokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancertokenApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancertoken.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancertokenApproval)
				if err := _Balancertoken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Balancertoken *BalancertokenFilterer) ParseApproval(log types.Log) (*BalancertokenApproval, error) {
	event := new(BalancertokenApproval)
	if err := _Balancertoken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancertokenLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the Balancertoken contract.
type BalancertokenLOGEXITIterator struct {
	Event *BalancertokenLOGEXIT // Event containing the contract specifics and raw log

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
func (it *BalancertokenLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancertokenLOGEXIT)
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
		it.Event = new(BalancertokenLOGEXIT)
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
func (it *BalancertokenLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancertokenLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancertokenLOGEXIT represents a LOGEXIT event raised by the Balancertoken contract.
type BalancertokenLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*BalancertokenLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancertoken.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancertokenLOGEXITIterator{contract: _Balancertoken.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *BalancertokenLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Balancertoken.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancertokenLOGEXIT)
				if err := _Balancertoken.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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
func (_Balancertoken *BalancertokenFilterer) ParseLOGEXIT(log types.Log) (*BalancertokenLOGEXIT, error) {
	event := new(BalancertokenLOGEXIT)
	if err := _Balancertoken.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancertokenLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the Balancertoken contract.
type BalancertokenLOGJOINIterator struct {
	Event *BalancertokenLOGJOIN // Event containing the contract specifics and raw log

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
func (it *BalancertokenLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancertokenLOGJOIN)
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
		it.Event = new(BalancertokenLOGJOIN)
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
func (it *BalancertokenLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancertokenLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancertokenLOGJOIN represents a LOGJOIN event raised by the Balancertoken contract.
type BalancertokenLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*BalancertokenLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Balancertoken.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &BalancertokenLOGJOINIterator{contract: _Balancertoken.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_Balancertoken *BalancertokenFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *BalancertokenLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _Balancertoken.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancertokenLOGJOIN)
				if err := _Balancertoken.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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
func (_Balancertoken *BalancertokenFilterer) ParseLOGJOIN(log types.Log) (*BalancertokenLOGJOIN, error) {
	event := new(BalancertokenLOGJOIN)
	if err := _Balancertoken.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancertokenLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the Balancertoken contract.
type BalancertokenLOGSWAPIterator struct {
	Event *BalancertokenLOGSWAP // Event containing the contract specifics and raw log

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
func (it *BalancertokenLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancertokenLOGSWAP)
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
		it.Event = new(BalancertokenLOGSWAP)
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
func (it *BalancertokenLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancertokenLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancertokenLOGSWAP represents a LOGSWAP event raised by the Balancertoken contract.
type BalancertokenLOGSWAP struct {
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
func (_Balancertoken *BalancertokenFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*BalancertokenLOGSWAPIterator, error) {

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

	logs, sub, err := _Balancertoken.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancertokenLOGSWAPIterator{contract: _Balancertoken.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_Balancertoken *BalancertokenFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *BalancertokenLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Balancertoken.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancertokenLOGSWAP)
				if err := _Balancertoken.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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
func (_Balancertoken *BalancertokenFilterer) ParseLOGSWAP(log types.Log) (*BalancertokenLOGSWAP, error) {
	event := new(BalancertokenLOGSWAP)
	if err := _Balancertoken.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancertokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Balancertoken contract.
type BalancertokenTransferIterator struct {
	Event *BalancertokenTransfer // Event containing the contract specifics and raw log

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
func (it *BalancertokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancertokenTransfer)
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
		it.Event = new(BalancertokenTransfer)
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
func (it *BalancertokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancertokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancertokenTransfer represents a Transfer event raised by the Balancertoken contract.
type BalancertokenTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Balancertoken *BalancertokenFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancertokenTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancertoken.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancertokenTransferIterator{contract: _Balancertoken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_Balancertoken *BalancertokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancertokenTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _Balancertoken.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancertokenTransfer)
				if err := _Balancertoken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Balancertoken *BalancertokenFilterer) ParseTransfer(log types.Log) (*BalancertokenTransfer, error) {
	event := new(BalancertokenTransfer)
	if err := _Balancertoken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

