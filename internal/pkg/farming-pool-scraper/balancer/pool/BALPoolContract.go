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

// BalancerPoolContractABI is the input ABI used to generate the binding from.
const BalancerPoolContractABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":true,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes4\",\"name\":\"sig\",\"type\":\"bytes4\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"LOG_CALL\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_EXIT\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"name\":\"LOG_JOIN\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"name\":\"LOG_SWAP\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"BONE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BPOW_PRECISION\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"EXIT_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"INIT_POOL_SUPPLY\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_IN_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_OUT_RATIO\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_TOTAL_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MAX_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BALANCE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BOUND_TOKENS\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_BPOW_BASE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_FEE\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"MIN_WEIGHT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"whom\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"bind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcInGivenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcOutGivenIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolInGivenSingleOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcPoolOutGivenSingleIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleInGivenPoolOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolSupply\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalWeight\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSingleOutGivenPoolIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenBalanceIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBalanceOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenWeightOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"calcSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"exitPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPoolAmountIn\",\"type\":\"uint256\"}],\"name\":\"exitswapExternAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"exitswapPoolAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finalize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getColor\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getCurrentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getFinalTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getNormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getNumTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getSpotPriceSansFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"spotPrice\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getSwapFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getTotalDenormalizedWeight\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"gulp\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"t\",\"type\":\"address\"}],\"name\":\"isBound\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isFinalized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPublicSwap\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"}],\"name\":\"joinPool\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minPoolAmountOut\",\"type\":\"uint256\"}],\"name\":\"joinswapExternAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"poolAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"joinswapPoolAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denorm\",\"type\":\"uint256\"}],\"name\":\"rebind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setController\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bool\",\"name\":\"public_\",\"type\":\"bool\"}],\"name\":\"setPublicSwap\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"swapFee\",\"type\":\"uint256\"}],\"name\":\"setSwapFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountIn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxPrice\",\"type\":\"uint256\"}],\"name\":\"swapExactAmountOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"spotPriceAfter\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amt\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"unbind\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// BalancerPoolContract is an auto generated Go binding around an Ethereum contract.
type BalancerPoolContract struct {
	BalancerPoolContractCaller     // Read-only binding to the contract
	BalancerPoolContractTransactor // Write-only binding to the contract
	BalancerPoolContractFilterer   // Log filterer for contract events
}

// BalancerPoolContractCaller is an auto generated read-only Go binding around an Ethereum contract.
type BalancerPoolContractCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerPoolContractTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BalancerPoolContractTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerPoolContractFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BalancerPoolContractFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BalancerPoolContractSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BalancerPoolContractSession struct {
	Contract     *BalancerPoolContract // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BalancerPoolContractCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BalancerPoolContractCallerSession struct {
	Contract *BalancerPoolContractCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// BalancerPoolContractTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BalancerPoolContractTransactorSession struct {
	Contract     *BalancerPoolContractTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// BalancerPoolContractRaw is an auto generated low-level Go binding around an Ethereum contract.
type BalancerPoolContractRaw struct {
	Contract *BalancerPoolContract // Generic contract binding to access the raw methods on
}

// BalancerPoolContractCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BalancerPoolContractCallerRaw struct {
	Contract *BalancerPoolContractCaller // Generic read-only contract binding to access the raw methods on
}

// BalancerPoolContractTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BalancerPoolContractTransactorRaw struct {
	Contract *BalancerPoolContractTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBalancerPoolContract creates a new instance of BalancerPoolContract, bound to a specific deployed contract.
func NewBalancerPoolContract(address common.Address, backend bind.ContractBackend) (*BalancerPoolContract, error) {
	contract, err := bindBalancerPoolContract(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContract{BalancerPoolContractCaller: BalancerPoolContractCaller{contract: contract}, BalancerPoolContractTransactor: BalancerPoolContractTransactor{contract: contract}, BalancerPoolContractFilterer: BalancerPoolContractFilterer{contract: contract}}, nil
}

// NewBalancerPoolContractCaller creates a new read-only instance of BalancerPoolContract, bound to a specific deployed contract.
func NewBalancerPoolContractCaller(address common.Address, caller bind.ContractCaller) (*BalancerPoolContractCaller, error) {
	contract, err := bindBalancerPoolContract(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractCaller{contract: contract}, nil
}

// NewBalancerPoolContractTransactor creates a new write-only instance of BalancerPoolContract, bound to a specific deployed contract.
func NewBalancerPoolContractTransactor(address common.Address, transactor bind.ContractTransactor) (*BalancerPoolContractTransactor, error) {
	contract, err := bindBalancerPoolContract(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractTransactor{contract: contract}, nil
}

// NewBalancerPoolContractFilterer creates a new log filterer instance of BalancerPoolContract, bound to a specific deployed contract.
func NewBalancerPoolContractFilterer(address common.Address, filterer bind.ContractFilterer) (*BalancerPoolContractFilterer, error) {
	contract, err := bindBalancerPoolContract(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractFilterer{contract: contract}, nil
}

// bindBalancerPoolContract binds a generic wrapper to an already deployed contract.
func bindBalancerPoolContract(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BalancerPoolContractABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerPoolContract *BalancerPoolContractRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerPoolContract.Contract.BalancerPoolContractCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerPoolContract *BalancerPoolContractRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.BalancerPoolContractTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerPoolContract *BalancerPoolContractRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.BalancerPoolContractTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BalancerPoolContract *BalancerPoolContractCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BalancerPoolContract.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BalancerPoolContract *BalancerPoolContractTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BalancerPoolContract *BalancerPoolContractTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.contract.Transact(opts, method, params...)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) BONE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "BONE")
	return *ret0, err
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) BONE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.BONE(&_BalancerPoolContract.CallOpts)
}

// BONE is a free data retrieval call binding the contract method 0xc36596a6.
//
// Solidity: function BONE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) BONE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.BONE(&_BalancerPoolContract.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) BPOWPRECISION(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "BPOW_PRECISION")
	return *ret0, err
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) BPOWPRECISION() (*big.Int, error) {
	return _BalancerPoolContract.Contract.BPOWPRECISION(&_BalancerPoolContract.CallOpts)
}

// BPOWPRECISION is a free data retrieval call binding the contract method 0x189d00ca.
//
// Solidity: function BPOW_PRECISION() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) BPOWPRECISION() (*big.Int, error) {
	return _BalancerPoolContract.Contract.BPOWPRECISION(&_BalancerPoolContract.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) EXITFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "EXIT_FEE")
	return *ret0, err
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) EXITFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.EXITFEE(&_BalancerPoolContract.CallOpts)
}

// EXITFEE is a free data retrieval call binding the contract method 0xc6580d12.
//
// Solidity: function EXIT_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) EXITFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.EXITFEE(&_BalancerPoolContract.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) INITPOOLSUPPLY(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "INIT_POOL_SUPPLY")
	return *ret0, err
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BalancerPoolContract.Contract.INITPOOLSUPPLY(&_BalancerPoolContract.CallOpts)
}

// INITPOOLSUPPLY is a free data retrieval call binding the contract method 0x9381cd2b.
//
// Solidity: function INIT_POOL_SUPPLY() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) INITPOOLSUPPLY() (*big.Int, error) {
	return _BalancerPoolContract.Contract.INITPOOLSUPPLY(&_BalancerPoolContract.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_BOUND_TOKENS")
	return *ret0, err
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXBOUNDTOKENS(&_BalancerPoolContract.CallOpts)
}

// MAXBOUNDTOKENS is a free data retrieval call binding the contract method 0xb0e0d136.
//
// Solidity: function MAX_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXBOUNDTOKENS() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXBOUNDTOKENS(&_BalancerPoolContract.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_BPOW_BASE")
	return *ret0, err
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXBPOWBASE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXBPOWBASE(&_BalancerPoolContract.CallOpts)
}

// MAXBPOWBASE is a free data retrieval call binding the contract method 0xbc694ea2.
//
// Solidity: function MAX_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXBPOWBASE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXBPOWBASE(&_BalancerPoolContract.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_FEE")
	return *ret0, err
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXFEE(&_BalancerPoolContract.CallOpts)
}

// MAXFEE is a free data retrieval call binding the contract method 0xbc063e1a.
//
// Solidity: function MAX_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXFEE(&_BalancerPoolContract.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXINRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_IN_RATIO")
	return *ret0, err
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXINRATIO() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXINRATIO(&_BalancerPoolContract.CallOpts)
}

// MAXINRATIO is a free data retrieval call binding the contract method 0xec093021.
//
// Solidity: function MAX_IN_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXINRATIO() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXINRATIO(&_BalancerPoolContract.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXOUTRATIO(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_OUT_RATIO")
	return *ret0, err
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXOUTRATIO() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXOUTRATIO(&_BalancerPoolContract.CallOpts)
}

// MAXOUTRATIO is a free data retrieval call binding the contract method 0x992e2a92.
//
// Solidity: function MAX_OUT_RATIO() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXOUTRATIO() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXOUTRATIO(&_BalancerPoolContract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXTOTALWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_TOTAL_WEIGHT")
	return *ret0, err
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXTOTALWEIGHT(&_BalancerPoolContract.CallOpts)
}

// MAXTOTALWEIGHT is a free data retrieval call binding the contract method 0x09a3bbe4.
//
// Solidity: function MAX_TOTAL_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXTOTALWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXTOTALWEIGHT(&_BalancerPoolContract.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MAXWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MAX_WEIGHT")
	return *ret0, err
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MAXWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXWEIGHT(&_BalancerPoolContract.CallOpts)
}

// MAXWEIGHT is a free data retrieval call binding the contract method 0xe4a28a52.
//
// Solidity: function MAX_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MAXWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MAXWEIGHT(&_BalancerPoolContract.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MINBALANCE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MIN_BALANCE")
	return *ret0, err
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MINBALANCE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBALANCE(&_BalancerPoolContract.CallOpts)
}

// MINBALANCE is a free data retrieval call binding the contract method 0x867378c5.
//
// Solidity: function MIN_BALANCE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MINBALANCE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBALANCE(&_BalancerPoolContract.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MINBOUNDTOKENS(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MIN_BOUND_TOKENS")
	return *ret0, err
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBOUNDTOKENS(&_BalancerPoolContract.CallOpts)
}

// MINBOUNDTOKENS is a free data retrieval call binding the contract method 0xb7b800a4.
//
// Solidity: function MIN_BOUND_TOKENS() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MINBOUNDTOKENS() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBOUNDTOKENS(&_BalancerPoolContract.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MINBPOWBASE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MIN_BPOW_BASE")
	return *ret0, err
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MINBPOWBASE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBPOWBASE(&_BalancerPoolContract.CallOpts)
}

// MINBPOWBASE is a free data retrieval call binding the contract method 0xba019dab.
//
// Solidity: function MIN_BPOW_BASE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MINBPOWBASE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINBPOWBASE(&_BalancerPoolContract.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MINFEE(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MIN_FEE")
	return *ret0, err
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MINFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINFEE(&_BalancerPoolContract.CallOpts)
}

// MINFEE is a free data retrieval call binding the contract method 0x76c7a3c7.
//
// Solidity: function MIN_FEE() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MINFEE() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINFEE(&_BalancerPoolContract.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) MINWEIGHT(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "MIN_WEIGHT")
	return *ret0, err
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) MINWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINWEIGHT(&_BalancerPoolContract.CallOpts)
}

// MINWEIGHT is a free data retrieval call binding the contract method 0x218b5382.
//
// Solidity: function MIN_WEIGHT() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) MINWEIGHT() (*big.Int, error) {
	return _BalancerPoolContract.Contract.MINWEIGHT(&_BalancerPoolContract.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) Allowance(opts *bind.CallOpts, src common.Address, dst common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "allowance", src, dst)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.Allowance(&_BalancerPoolContract.CallOpts, src, dst)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address src, address dst) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) Allowance(src common.Address, dst common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.Allowance(&_BalancerPoolContract.CallOpts, src, dst)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) BalanceOf(opts *bind.CallOpts, whom common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "balanceOf", whom)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.BalanceOf(&_BalancerPoolContract.CallOpts, whom)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address whom) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) BalanceOf(whom common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.BalanceOf(&_BalancerPoolContract.CallOpts, whom)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcInGivenOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcInGivenOut", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
	return *ret0, err
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcInGivenOut(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcInGivenOut is a free data retrieval call binding the contract method 0xf8d6aed4.
//
// Solidity: function calcInGivenOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcInGivenOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcInGivenOut(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountOut, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcOutGivenIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcOutGivenIn", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
	return *ret0, err
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcOutGivenIn(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcOutGivenIn is a free data retrieval call binding the contract method 0xba9530a6.
//
// Solidity: function calcOutGivenIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcOutGivenIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcOutGivenIn(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, tokenAmountIn, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcPoolInGivenSingleOut(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcPoolInGivenSingleOut", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
	return *ret0, err
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcPoolInGivenSingleOut(&_BalancerPoolContract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolInGivenSingleOut is a free data retrieval call binding the contract method 0x82f652ad.
//
// Solidity: function calcPoolInGivenSingleOut(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountOut, uint256 swapFee) pure returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcPoolInGivenSingleOut(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcPoolInGivenSingleOut(&_BalancerPoolContract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, tokenAmountOut, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcPoolOutGivenSingleIn(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcPoolOutGivenSingleIn", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
	return *ret0, err
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcPoolOutGivenSingleIn(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcPoolOutGivenSingleIn is a free data retrieval call binding the contract method 0x8656b653.
//
// Solidity: function calcPoolOutGivenSingleIn(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 tokenAmountIn, uint256 swapFee) pure returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcPoolOutGivenSingleIn(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, tokenAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcPoolOutGivenSingleIn(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, tokenAmountIn, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcSingleInGivenPoolOut(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcSingleInGivenPoolOut", tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
	return *ret0, err
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSingleInGivenPoolOut(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleInGivenPoolOut is a free data retrieval call binding the contract method 0x5c1bbaf7.
//
// Solidity: function calcSingleInGivenPoolOut(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountOut, uint256 swapFee) pure returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcSingleInGivenPoolOut(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSingleInGivenPoolOut(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, poolSupply, totalWeight, poolAmountOut, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcSingleOutGivenPoolIn(opts *bind.CallOpts, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcSingleOutGivenPoolIn", tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
	return *ret0, err
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSingleOutGivenPoolIn(&_BalancerPoolContract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSingleOutGivenPoolIn is a free data retrieval call binding the contract method 0x89298012.
//
// Solidity: function calcSingleOutGivenPoolIn(uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 poolSupply, uint256 totalWeight, uint256 poolAmountIn, uint256 swapFee) pure returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcSingleOutGivenPoolIn(tokenBalanceOut *big.Int, tokenWeightOut *big.Int, poolSupply *big.Int, totalWeight *big.Int, poolAmountIn *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSingleOutGivenPoolIn(&_BalancerPoolContract.CallOpts, tokenBalanceOut, tokenWeightOut, poolSupply, totalWeight, poolAmountIn, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCaller) CalcSpotPrice(opts *bind.CallOpts, tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "calcSpotPrice", tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
	return *ret0, err
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSpotPrice(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// CalcSpotPrice is a free data retrieval call binding the contract method 0xa221ee49.
//
// Solidity: function calcSpotPrice(uint256 tokenBalanceIn, uint256 tokenWeightIn, uint256 tokenBalanceOut, uint256 tokenWeightOut, uint256 swapFee) pure returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) CalcSpotPrice(tokenBalanceIn *big.Int, tokenWeightIn *big.Int, tokenBalanceOut *big.Int, tokenWeightOut *big.Int, swapFee *big.Int) (*big.Int, error) {
	return _BalancerPoolContract.Contract.CalcSpotPrice(&_BalancerPoolContract.CallOpts, tokenBalanceIn, tokenWeightIn, tokenBalanceOut, tokenWeightOut, swapFee)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerPoolContract *BalancerPoolContractCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerPoolContract *BalancerPoolContractSession) Decimals() (uint8, error) {
	return _BalancerPoolContract.Contract.Decimals(&_BalancerPoolContract.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) Decimals() (uint8, error) {
	return _BalancerPoolContract.Contract.Decimals(&_BalancerPoolContract.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetBalance(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getBalance", token)
	return *ret0, err
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetBalance(&_BalancerPoolContract.CallOpts, token)
}

// GetBalance is a free data retrieval call binding the contract method 0xf8b2cb4f.
//
// Solidity: function getBalance(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetBalance(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetBalance(&_BalancerPoolContract.CallOpts, token)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetColor(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getColor")
	return *ret0, err
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerPoolContract *BalancerPoolContractSession) GetColor() ([32]byte, error) {
	return _BalancerPoolContract.Contract.GetColor(&_BalancerPoolContract.CallOpts)
}

// GetColor is a free data retrieval call binding the contract method 0x9a86139b.
//
// Solidity: function getColor() view returns(bytes32)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetColor() ([32]byte, error) {
	return _BalancerPoolContract.Contract.GetColor(&_BalancerPoolContract.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetController(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getController")
	return *ret0, err
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerPoolContract *BalancerPoolContractSession) GetController() (common.Address, error) {
	return _BalancerPoolContract.Contract.GetController(&_BalancerPoolContract.CallOpts)
}

// GetController is a free data retrieval call binding the contract method 0x3018205f.
//
// Solidity: function getController() view returns(address)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetController() (common.Address, error) {
	return _BalancerPoolContract.Contract.GetController(&_BalancerPoolContract.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetCurrentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getCurrentTokens")
	return *ret0, err
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractSession) GetCurrentTokens() ([]common.Address, error) {
	return _BalancerPoolContract.Contract.GetCurrentTokens(&_BalancerPoolContract.CallOpts)
}

// GetCurrentTokens is a free data retrieval call binding the contract method 0xcc77828d.
//
// Solidity: function getCurrentTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetCurrentTokens() ([]common.Address, error) {
	return _BalancerPoolContract.Contract.GetCurrentTokens(&_BalancerPoolContract.CallOpts)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetDenormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getDenormalizedWeight", token)
	return *ret0, err
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetDenormalizedWeight(&_BalancerPoolContract.CallOpts, token)
}

// GetDenormalizedWeight is a free data retrieval call binding the contract method 0x948d8ce6.
//
// Solidity: function getDenormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetDenormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetDenormalizedWeight(&_BalancerPoolContract.CallOpts, token)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetFinalTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var (
		ret0 = new([]common.Address)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getFinalTokens")
	return *ret0, err
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractSession) GetFinalTokens() ([]common.Address, error) {
	return _BalancerPoolContract.Contract.GetFinalTokens(&_BalancerPoolContract.CallOpts)
}

// GetFinalTokens is a free data retrieval call binding the contract method 0xbe3bbd2e.
//
// Solidity: function getFinalTokens() view returns(address[] tokens)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetFinalTokens() ([]common.Address, error) {
	return _BalancerPoolContract.Contract.GetFinalTokens(&_BalancerPoolContract.CallOpts)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetNormalizedWeight(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getNormalizedWeight", token)
	return *ret0, err
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetNormalizedWeight(&_BalancerPoolContract.CallOpts, token)
}

// GetNormalizedWeight is a free data retrieval call binding the contract method 0xf1b8a9b7.
//
// Solidity: function getNormalizedWeight(address token) view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetNormalizedWeight(token common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetNormalizedWeight(&_BalancerPoolContract.CallOpts, token)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetNumTokens(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getNumTokens")
	return *ret0, err
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetNumTokens() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetNumTokens(&_BalancerPoolContract.CallOpts)
}

// GetNumTokens is a free data retrieval call binding the contract method 0xcd2ed8fb.
//
// Solidity: function getNumTokens() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetNumTokens() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetNumTokens(&_BalancerPoolContract.CallOpts)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetSpotPrice(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getSpotPrice", tokenIn, tokenOut)
	return *ret0, err
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSpotPrice(&_BalancerPoolContract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPrice is a free data retrieval call binding the contract method 0x15e84af9.
//
// Solidity: function getSpotPrice(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetSpotPrice(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSpotPrice(&_BalancerPoolContract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetSpotPriceSansFee(opts *bind.CallOpts, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getSpotPriceSansFee", tokenIn, tokenOut)
	return *ret0, err
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSpotPriceSansFee(&_BalancerPoolContract.CallOpts, tokenIn, tokenOut)
}

// GetSpotPriceSansFee is a free data retrieval call binding the contract method 0x1446a7ff.
//
// Solidity: function getSpotPriceSansFee(address tokenIn, address tokenOut) view returns(uint256 spotPrice)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetSpotPriceSansFee(tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSpotPriceSansFee(&_BalancerPoolContract.CallOpts, tokenIn, tokenOut)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetSwapFee(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getSwapFee")
	return *ret0, err
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetSwapFee() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSwapFee(&_BalancerPoolContract.CallOpts)
}

// GetSwapFee is a free data retrieval call binding the contract method 0xd4cadf68.
//
// Solidity: function getSwapFee() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetSwapFee() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetSwapFee(&_BalancerPoolContract.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) GetTotalDenormalizedWeight(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "getTotalDenormalizedWeight")
	return *ret0, err
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetTotalDenormalizedWeight(&_BalancerPoolContract.CallOpts)
}

// GetTotalDenormalizedWeight is a free data retrieval call binding the contract method 0x936c3477.
//
// Solidity: function getTotalDenormalizedWeight() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) GetTotalDenormalizedWeight() (*big.Int, error) {
	return _BalancerPoolContract.Contract.GetTotalDenormalizedWeight(&_BalancerPoolContract.CallOpts)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCaller) IsBound(opts *bind.CallOpts, t common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "isBound", t)
	return *ret0, err
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) IsBound(t common.Address) (bool, error) {
	return _BalancerPoolContract.Contract.IsBound(&_BalancerPoolContract.CallOpts, t)
}

// IsBound is a free data retrieval call binding the contract method 0x2f37b624.
//
// Solidity: function isBound(address t) view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) IsBound(t common.Address) (bool, error) {
	return _BalancerPoolContract.Contract.IsBound(&_BalancerPoolContract.CallOpts, t)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCaller) IsFinalized(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "isFinalized")
	return *ret0, err
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) IsFinalized() (bool, error) {
	return _BalancerPoolContract.Contract.IsFinalized(&_BalancerPoolContract.CallOpts)
}

// IsFinalized is a free data retrieval call binding the contract method 0x8d4e4083.
//
// Solidity: function isFinalized() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) IsFinalized() (bool, error) {
	return _BalancerPoolContract.Contract.IsFinalized(&_BalancerPoolContract.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCaller) IsPublicSwap(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "isPublicSwap")
	return *ret0, err
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) IsPublicSwap() (bool, error) {
	return _BalancerPoolContract.Contract.IsPublicSwap(&_BalancerPoolContract.CallOpts)
}

// IsPublicSwap is a free data retrieval call binding the contract method 0xfde924f7.
//
// Solidity: function isPublicSwap() view returns(bool)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) IsPublicSwap() (bool, error) {
	return _BalancerPoolContract.Contract.IsPublicSwap(&_BalancerPoolContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractSession) Name() (string, error) {
	return _BalancerPoolContract.Contract.Name(&_BalancerPoolContract.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) Name() (string, error) {
	return _BalancerPoolContract.Contract.Name(&_BalancerPoolContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractSession) Symbol() (string, error) {
	return _BalancerPoolContract.Contract.Symbol(&_BalancerPoolContract.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) Symbol() (string, error) {
	return _BalancerPoolContract.Contract.Symbol(&_BalancerPoolContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _BalancerPoolContract.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractSession) TotalSupply() (*big.Int, error) {
	return _BalancerPoolContract.Contract.TotalSupply(&_BalancerPoolContract.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BalancerPoolContract *BalancerPoolContractCallerSession) TotalSupply() (*big.Int, error) {
	return _BalancerPoolContract.Contract.TotalSupply(&_BalancerPoolContract.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactor) Approve(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "approve", dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Approve(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Approve(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Approve(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) Bind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "bind", token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Bind(&_BalancerPoolContract.TransactOpts, token, balance, denorm)
}

// Bind is a paid mutator transaction binding the contract method 0xe4e1e538.
//
// Solidity: function bind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Bind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Bind(&_BalancerPoolContract.TransactOpts, token, balance, denorm)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactor) DecreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "decreaseApproval", dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.DecreaseApproval(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) DecreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.DecreaseApproval(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) ExitPool(opts *bind.TransactOpts, poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "exitPool", poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitPool(&_BalancerPoolContract.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitPool is a paid mutator transaction binding the contract method 0xb02f0b73.
//
// Solidity: function exitPool(uint256 poolAmountIn, uint256[] minAmountsOut) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) ExitPool(poolAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitPool(&_BalancerPoolContract.TransactOpts, poolAmountIn, minAmountsOut)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractTransactor) ExitswapExternAmountOut(opts *bind.TransactOpts, tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "exitswapExternAmountOut", tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitswapExternAmountOut(&_BalancerPoolContract.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapExternAmountOut is a paid mutator transaction binding the contract method 0x02c96748.
//
// Solidity: function exitswapExternAmountOut(address tokenOut, uint256 tokenAmountOut, uint256 maxPoolAmountIn) returns(uint256 poolAmountIn)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) ExitswapExternAmountOut(tokenOut common.Address, tokenAmountOut *big.Int, maxPoolAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitswapExternAmountOut(&_BalancerPoolContract.TransactOpts, tokenOut, tokenAmountOut, maxPoolAmountIn)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractTransactor) ExitswapPoolAmountIn(opts *bind.TransactOpts, tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "exitswapPoolAmountIn", tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitswapPoolAmountIn(&_BalancerPoolContract.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// ExitswapPoolAmountIn is a paid mutator transaction binding the contract method 0x46ab38f1.
//
// Solidity: function exitswapPoolAmountIn(address tokenOut, uint256 poolAmountIn, uint256 minAmountOut) returns(uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) ExitswapPoolAmountIn(tokenOut common.Address, poolAmountIn *big.Int, minAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.ExitswapPoolAmountIn(&_BalancerPoolContract.TransactOpts, tokenOut, poolAmountIn, minAmountOut)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) Finalize(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "finalize")
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerPoolContract *BalancerPoolContractSession) Finalize() (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Finalize(&_BalancerPoolContract.TransactOpts)
}

// Finalize is a paid mutator transaction binding the contract method 0x4bb278f3.
//
// Solidity: function finalize() returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Finalize() (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Finalize(&_BalancerPoolContract.TransactOpts)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) Gulp(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "gulp", token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Gulp(&_BalancerPoolContract.TransactOpts, token)
}

// Gulp is a paid mutator transaction binding the contract method 0x8c28cbe8.
//
// Solidity: function gulp(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Gulp(token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Gulp(&_BalancerPoolContract.TransactOpts, token)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactor) IncreaseApproval(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "increaseApproval", dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.IncreaseApproval(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) IncreaseApproval(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.IncreaseApproval(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) JoinPool(opts *bind.TransactOpts, poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "joinPool", poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinPool(&_BalancerPoolContract.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinPool is a paid mutator transaction binding the contract method 0x4f69c0d4.
//
// Solidity: function joinPool(uint256 poolAmountOut, uint256[] maxAmountsIn) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) JoinPool(poolAmountOut *big.Int, maxAmountsIn []*big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinPool(&_BalancerPoolContract.TransactOpts, poolAmountOut, maxAmountsIn)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractTransactor) JoinswapExternAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "joinswapExternAmountIn", tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinswapExternAmountIn(&_BalancerPoolContract.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapExternAmountIn is a paid mutator transaction binding the contract method 0x5db34277.
//
// Solidity: function joinswapExternAmountIn(address tokenIn, uint256 tokenAmountIn, uint256 minPoolAmountOut) returns(uint256 poolAmountOut)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) JoinswapExternAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, minPoolAmountOut *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinswapExternAmountIn(&_BalancerPoolContract.TransactOpts, tokenIn, tokenAmountIn, minPoolAmountOut)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractTransactor) JoinswapPoolAmountOut(opts *bind.TransactOpts, tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "joinswapPoolAmountOut", tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinswapPoolAmountOut(&_BalancerPoolContract.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// JoinswapPoolAmountOut is a paid mutator transaction binding the contract method 0x6d06dfa0.
//
// Solidity: function joinswapPoolAmountOut(address tokenIn, uint256 poolAmountOut, uint256 maxAmountIn) returns(uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) JoinswapPoolAmountOut(tokenIn common.Address, poolAmountOut *big.Int, maxAmountIn *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.JoinswapPoolAmountOut(&_BalancerPoolContract.TransactOpts, tokenIn, poolAmountOut, maxAmountIn)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) Rebind(opts *bind.TransactOpts, token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "rebind", token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Rebind(&_BalancerPoolContract.TransactOpts, token, balance, denorm)
}

// Rebind is a paid mutator transaction binding the contract method 0x3fdddaa2.
//
// Solidity: function rebind(address token, uint256 balance, uint256 denorm) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Rebind(token common.Address, balance *big.Int, denorm *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Rebind(&_BalancerPoolContract.TransactOpts, token, balance, denorm)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) SetController(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "setController", manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetController(&_BalancerPoolContract.TransactOpts, manager)
}

// SetController is a paid mutator transaction binding the contract method 0x92eefe9b.
//
// Solidity: function setController(address manager) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) SetController(manager common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetController(&_BalancerPoolContract.TransactOpts, manager)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) SetPublicSwap(opts *bind.TransactOpts, public_ bool) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "setPublicSwap", public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetPublicSwap(&_BalancerPoolContract.TransactOpts, public_)
}

// SetPublicSwap is a paid mutator transaction binding the contract method 0x49b59552.
//
// Solidity: function setPublicSwap(bool public_) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) SetPublicSwap(public_ bool) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetPublicSwap(&_BalancerPoolContract.TransactOpts, public_)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) SetSwapFee(opts *bind.TransactOpts, swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "setSwapFee", swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetSwapFee(&_BalancerPoolContract.TransactOpts, swapFee)
}

// SetSwapFee is a paid mutator transaction binding the contract method 0x34e19907.
//
// Solidity: function setSwapFee(uint256 swapFee) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) SetSwapFee(swapFee *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SetSwapFee(&_BalancerPoolContract.TransactOpts, swapFee)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractTransactor) SwapExactAmountIn(opts *bind.TransactOpts, tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "swapExactAmountIn", tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SwapExactAmountIn(&_BalancerPoolContract.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountIn is a paid mutator transaction binding the contract method 0x8201aa3f.
//
// Solidity: function swapExactAmountIn(address tokenIn, uint256 tokenAmountIn, address tokenOut, uint256 minAmountOut, uint256 maxPrice) returns(uint256 tokenAmountOut, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) SwapExactAmountIn(tokenIn common.Address, tokenAmountIn *big.Int, tokenOut common.Address, minAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SwapExactAmountIn(&_BalancerPoolContract.TransactOpts, tokenIn, tokenAmountIn, tokenOut, minAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractTransactor) SwapExactAmountOut(opts *bind.TransactOpts, tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "swapExactAmountOut", tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SwapExactAmountOut(&_BalancerPoolContract.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// SwapExactAmountOut is a paid mutator transaction binding the contract method 0x7c5e9ea4.
//
// Solidity: function swapExactAmountOut(address tokenIn, uint256 maxAmountIn, address tokenOut, uint256 tokenAmountOut, uint256 maxPrice) returns(uint256 tokenAmountIn, uint256 spotPriceAfter)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) SwapExactAmountOut(tokenIn common.Address, maxAmountIn *big.Int, tokenOut common.Address, tokenAmountOut *big.Int, maxPrice *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.SwapExactAmountOut(&_BalancerPoolContract.TransactOpts, tokenIn, maxAmountIn, tokenOut, tokenAmountOut, maxPrice)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactor) Transfer(opts *bind.TransactOpts, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "transfer", dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Transfer(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Transfer(dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Transfer(&_BalancerPoolContract.TransactOpts, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactor) TransferFrom(opts *bind.TransactOpts, src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "transferFrom", src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.TransferFrom(&_BalancerPoolContract.TransactOpts, src, dst, amt)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address src, address dst, uint256 amt) returns(bool)
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) TransferFrom(src common.Address, dst common.Address, amt *big.Int) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.TransferFrom(&_BalancerPoolContract.TransactOpts, src, dst, amt)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactor) Unbind(opts *bind.TransactOpts, token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.contract.Transact(opts, "unbind", token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Unbind(&_BalancerPoolContract.TransactOpts, token)
}

// Unbind is a paid mutator transaction binding the contract method 0xcf5e7bd3.
//
// Solidity: function unbind(address token) returns()
func (_BalancerPoolContract *BalancerPoolContractTransactorSession) Unbind(token common.Address) (*types.Transaction, error) {
	return _BalancerPoolContract.Contract.Unbind(&_BalancerPoolContract.TransactOpts, token)
}

// BalancerPoolContractApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BalancerPoolContract contract.
type BalancerPoolContractApprovalIterator struct {
	Event *BalancerPoolContractApproval // Event containing the contract specifics and raw log

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
func (it *BalancerPoolContractApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerPoolContractApproval)
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
		it.Event = new(BalancerPoolContractApproval)
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
func (it *BalancerPoolContractApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerPoolContractApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerPoolContractApproval represents a Approval event raised by the BalancerPoolContract contract.
type BalancerPoolContractApproval struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BalancerPoolContract *BalancerPoolContractFilterer) FilterApproval(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerPoolContractApprovalIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.FilterLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractApprovalIterator{contract: _BalancerPoolContract.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed src, address indexed dst, uint256 amt)
func (_BalancerPoolContract *BalancerPoolContractFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BalancerPoolContractApproval, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.WatchLogs(opts, "Approval", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerPoolContractApproval)
				if err := _BalancerPoolContract.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) ParseApproval(log types.Log) (*BalancerPoolContractApproval, error) {
	event := new(BalancerPoolContractApproval)
	if err := _BalancerPoolContract.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerPoolContractLOGEXITIterator is returned from FilterLOGEXIT and is used to iterate over the raw logs and unpacked data for LOGEXIT events raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGEXITIterator struct {
	Event *BalancerPoolContractLOGEXIT // Event containing the contract specifics and raw log

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
func (it *BalancerPoolContractLOGEXITIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerPoolContractLOGEXIT)
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
		it.Event = new(BalancerPoolContractLOGEXIT)
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
func (it *BalancerPoolContractLOGEXITIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerPoolContractLOGEXITIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerPoolContractLOGEXIT represents a LOGEXIT event raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGEXIT struct {
	Caller         common.Address
	TokenOut       common.Address
	TokenAmountOut *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterLOGEXIT is a free log retrieval operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractFilterer) FilterLOGEXIT(opts *bind.FilterOpts, caller []common.Address, tokenOut []common.Address) (*BalancerPoolContractLOGEXITIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.FilterLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractLOGEXITIterator{contract: _BalancerPoolContract.contract, event: "LOG_EXIT", logs: logs, sub: sub}, nil
}

// WatchLOGEXIT is a free log subscription operation binding the contract event 0xe74c91552b64c2e2e7bd255639e004e693bd3e1d01cc33e65610b86afcc1ffed.
//
// Solidity: event LOG_EXIT(address indexed caller, address indexed tokenOut, uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractFilterer) WatchLOGEXIT(opts *bind.WatchOpts, sink chan<- *BalancerPoolContractLOGEXIT, caller []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.WatchLogs(opts, "LOG_EXIT", callerRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerPoolContractLOGEXIT)
				if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) ParseLOGEXIT(log types.Log) (*BalancerPoolContractLOGEXIT, error) {
	event := new(BalancerPoolContractLOGEXIT)
	if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_EXIT", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerPoolContractLOGJOINIterator is returned from FilterLOGJOIN and is used to iterate over the raw logs and unpacked data for LOGJOIN events raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGJOINIterator struct {
	Event *BalancerPoolContractLOGJOIN // Event containing the contract specifics and raw log

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
func (it *BalancerPoolContractLOGJOINIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerPoolContractLOGJOIN)
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
		it.Event = new(BalancerPoolContractLOGJOIN)
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
func (it *BalancerPoolContractLOGJOINIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerPoolContractLOGJOINIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerPoolContractLOGJOIN represents a LOGJOIN event raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGJOIN struct {
	Caller        common.Address
	TokenIn       common.Address
	TokenAmountIn *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLOGJOIN is a free log retrieval operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractFilterer) FilterLOGJOIN(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address) (*BalancerPoolContractLOGJOINIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.FilterLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractLOGJOINIterator{contract: _BalancerPoolContract.contract, event: "LOG_JOIN", logs: logs, sub: sub}, nil
}

// WatchLOGJOIN is a free log subscription operation binding the contract event 0x63982df10efd8dfaaaa0fcc7f50b2d93b7cba26ccc48adee2873220d485dc39a.
//
// Solidity: event LOG_JOIN(address indexed caller, address indexed tokenIn, uint256 tokenAmountIn)
func (_BalancerPoolContract *BalancerPoolContractFilterer) WatchLOGJOIN(opts *bind.WatchOpts, sink chan<- *BalancerPoolContractLOGJOIN, caller []common.Address, tokenIn []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.WatchLogs(opts, "LOG_JOIN", callerRule, tokenInRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerPoolContractLOGJOIN)
				if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) ParseLOGJOIN(log types.Log) (*BalancerPoolContractLOGJOIN, error) {
	event := new(BalancerPoolContractLOGJOIN)
	if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_JOIN", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerPoolContractLOGSWAPIterator is returned from FilterLOGSWAP and is used to iterate over the raw logs and unpacked data for LOGSWAP events raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGSWAPIterator struct {
	Event *BalancerPoolContractLOGSWAP // Event containing the contract specifics and raw log

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
func (it *BalancerPoolContractLOGSWAPIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerPoolContractLOGSWAP)
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
		it.Event = new(BalancerPoolContractLOGSWAP)
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
func (it *BalancerPoolContractLOGSWAPIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerPoolContractLOGSWAPIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerPoolContractLOGSWAP represents a LOGSWAP event raised by the BalancerPoolContract contract.
type BalancerPoolContractLOGSWAP struct {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) FilterLOGSWAP(opts *bind.FilterOpts, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*BalancerPoolContractLOGSWAPIterator, error) {

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

	logs, sub, err := _BalancerPoolContract.contract.FilterLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractLOGSWAPIterator{contract: _BalancerPoolContract.contract, event: "LOG_SWAP", logs: logs, sub: sub}, nil
}

// WatchLOGSWAP is a free log subscription operation binding the contract event 0x908fb5ee8f16c6bc9bc3690973819f32a4d4b10188134543c88706e0e1d43378.
//
// Solidity: event LOG_SWAP(address indexed caller, address indexed tokenIn, address indexed tokenOut, uint256 tokenAmountIn, uint256 tokenAmountOut)
func (_BalancerPoolContract *BalancerPoolContractFilterer) WatchLOGSWAP(opts *bind.WatchOpts, sink chan<- *BalancerPoolContractLOGSWAP, caller []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _BalancerPoolContract.contract.WatchLogs(opts, "LOG_SWAP", callerRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerPoolContractLOGSWAP)
				if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) ParseLOGSWAP(log types.Log) (*BalancerPoolContractLOGSWAP, error) {
	event := new(BalancerPoolContractLOGSWAP)
	if err := _BalancerPoolContract.contract.UnpackLog(event, "LOG_SWAP", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BalancerPoolContractTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BalancerPoolContract contract.
type BalancerPoolContractTransferIterator struct {
	Event *BalancerPoolContractTransfer // Event containing the contract specifics and raw log

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
func (it *BalancerPoolContractTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BalancerPoolContractTransfer)
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
		it.Event = new(BalancerPoolContractTransfer)
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
func (it *BalancerPoolContractTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BalancerPoolContractTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BalancerPoolContractTransfer represents a Transfer event raised by the BalancerPoolContract contract.
type BalancerPoolContractTransfer struct {
	Src common.Address
	Dst common.Address
	Amt *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BalancerPoolContract *BalancerPoolContractFilterer) FilterTransfer(opts *bind.FilterOpts, src []common.Address, dst []common.Address) (*BalancerPoolContractTransferIterator, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.FilterLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return &BalancerPoolContractTransferIterator{contract: _BalancerPoolContract.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed src, address indexed dst, uint256 amt)
func (_BalancerPoolContract *BalancerPoolContractFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BalancerPoolContractTransfer, src []common.Address, dst []common.Address) (event.Subscription, error) {

	var srcRule []interface{}
	for _, srcItem := range src {
		srcRule = append(srcRule, srcItem)
	}
	var dstRule []interface{}
	for _, dstItem := range dst {
		dstRule = append(dstRule, dstItem)
	}

	logs, sub, err := _BalancerPoolContract.contract.WatchLogs(opts, "Transfer", srcRule, dstRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BalancerPoolContractTransfer)
				if err := _BalancerPoolContract.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BalancerPoolContract *BalancerPoolContractFilterer) ParseTransfer(log types.Log) (*BalancerPoolContractTransfer, error) {
	event := new(BalancerPoolContractTransfer)
	if err := _BalancerPoolContract.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
