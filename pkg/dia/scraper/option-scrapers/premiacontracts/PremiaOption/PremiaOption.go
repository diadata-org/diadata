// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PremiaOption

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

// PremiaOptionOptionData is an auto generated low-level Go binding around an user-defined struct.
type PremiaOptionOptionData struct {
	Token         common.Address
	StrikePrice   *big.Int
	Expiration    *big.Int
	IsCall        bool
	ClaimsPreExp  *big.Int
	ClaimsPostExp *big.Int
	Exercised     *big.Int
	Supply        *big.Int
	Decimals      uint8
}

// PremiaOptionOptionWriteArgs is an auto generated low-level Go binding around an user-defined struct.
type PremiaOptionOptionWriteArgs struct {
	Token       common.Address
	Amount      *big.Int
	StrikePrice *big.Int
	Expiration  *big.Int
	IsCall      bool
}

// PremiaOptionQuoteExercise is an auto generated low-level Go binding around an user-defined struct.
type PremiaOptionQuoteExercise struct {
	InputToken     common.Address
	Input          *big.Int
	InputDecimals  uint8
	OutputToken    common.Address
	Output         *big.Int
	OutputDecimals uint8
	Fee            *big.Int
	FeeReferrer    *big.Int
}

// PremiaOptionQuoteWrite is an auto generated low-level Go binding around an user-defined struct.
type PremiaOptionQuoteWrite struct {
	CollateralToken    common.Address
	Collateral         *big.Int
	CollateralDecimals uint8
	Fee                *big.Int
	FeeReferrer        *big.Int
}

// PremiaOptionABI is the input ABI used to generate the binding from.
const PremiaOptionABI = "[{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_uri\",\"type\":\"string\"},{\"internalType\":\"contractIERC20\",\"name\":\"_denominator\",\"type\":\"address\"},{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"_uPremia\",\"type\":\"address\"},{\"internalType\":\"contractIFeeCalculator\",\"name\":\"_feeCalculator\",\"type\":\"address\"},{\"internalType\":\"contractIPremiaReferral\",\"name\":\"_premiaReferral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"referrer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeProtocol\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeReferrer\",\"type\":\"uint256\"}],\"name\":\"FeePaid\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OptionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OptionExercised\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"OptionIdCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"OptionWritten\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"strikePriceIncrement\",\"type\":\"uint256\"}],\"name\":\"SetToken\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"values\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"TransferSingle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"value\",\"type\":\"string\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"URI\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cancelOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"cancelOptionFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominator\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"denominatorDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"exerciseOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"exerciseOptionFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCalculator\",\"outputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"flashExerciseOption\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"contractIUniswapV2Router02\",\"name\":\"_router\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amountInMax\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"_path\",\"type\":\"address[]\"}],\"name\":\"flashExerciseOptionFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_tokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"contractIFlashLoanReceiver\",\"name\":\"_receiver\",\"type\":\"address\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"claimsPreExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsPostExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exercised\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaOption.OptionData\",\"name\":\"_option\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"getExerciseQuote\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"inputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"input\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"inputDecimals\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"outputToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"output\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"outputDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReferrer\",\"type\":\"uint256\"}],\"internalType\":\"structPremiaOption.QuoteExercise\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isCall\",\"type\":\"bool\"}],\"name\":\"getOptionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_isCall\",\"type\":\"bool\"}],\"name\":\"getOptionIdOrCreate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"internalType\":\"structPremiaOption.OptionWriteArgs\",\"name\":\"_option\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"_decimals\",\"type\":\"uint8\"}],\"name\":\"getWriteQuote\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"collateralToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateral\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"collateralDecimals\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"feeReferrer\",\"type\":\"uint256\"}],\"internalType\":\"structPremiaOption.QuoteWrite\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"maxExpiration\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nbWritten\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"nextOptionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"optionData\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"claimsPreExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"claimsPostExp\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"exercised\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"name\":\"options\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"pools\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denominatorAmount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"premiaReferral\",\"outputs\":[{\"internalType\":\"contractIPremiaReferral\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeBatchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"_feeCalculator\",\"type\":\"address\"}],\"name\":\"setFeeCalculator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_max\",\"type\":\"uint256\"}],\"name\":\"setMaxExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPremiaReferral\",\"name\":\"_premiaReferral\",\"type\":\"address\"}],\"name\":\"setPremiaReferral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"_uPremia\",\"type\":\"address\"}],\"name\":\"setPremiaUncutErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_strikePriceIncrement\",\"type\":\"uint256[]\"}],\"name\":\"setTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_newUri\",\"type\":\"string\"}],\"name\":\"setURI\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addrList\",\"type\":\"address[]\"}],\"name\":\"setWhitelistedUniswapRouters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"tokenStrikeIncrement\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokensLength\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uPremia\",\"outputs\":[{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"uri\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"whitelistedUniswapRouters\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"}],\"name\":\"withdrawFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPreExpiration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawPreExpirationFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"internalType\":\"structPremiaOption.OptionWriteArgs\",\"name\":\"_option\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"writeOption\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"internalType\":\"structPremiaOption.OptionWriteArgs\",\"name\":\"_option\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"writeOptionFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_optionId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"writeOptionWithIdFrom\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PremiaOption is an auto generated Go binding around an Ethereum contract.
type PremiaOption struct {
	PremiaOptionCaller     // Read-only binding to the contract
	PremiaOptionTransactor // Write-only binding to the contract
	PremiaOptionFilterer   // Log filterer for contract events
}

// PremiaOptionCaller is an auto generated read-only Go binding around an Ethereum contract.
type PremiaOptionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaOptionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PremiaOptionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaOptionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PremiaOptionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaOptionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PremiaOptionSession struct {
	Contract     *PremiaOption     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PremiaOptionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PremiaOptionCallerSession struct {
	Contract *PremiaOptionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PremiaOptionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PremiaOptionTransactorSession struct {
	Contract     *PremiaOptionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PremiaOptionRaw is an auto generated low-level Go binding around an Ethereum contract.
type PremiaOptionRaw struct {
	Contract *PremiaOption // Generic contract binding to access the raw methods on
}

// PremiaOptionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PremiaOptionCallerRaw struct {
	Contract *PremiaOptionCaller // Generic read-only contract binding to access the raw methods on
}

// PremiaOptionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PremiaOptionTransactorRaw struct {
	Contract *PremiaOptionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPremiaOption creates a new instance of PremiaOption, bound to a specific deployed contract.
func NewPremiaOption(address common.Address, backend bind.ContractBackend) (*PremiaOption, error) {
	contract, err := bindPremiaOption(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PremiaOption{PremiaOptionCaller: PremiaOptionCaller{contract: contract}, PremiaOptionTransactor: PremiaOptionTransactor{contract: contract}, PremiaOptionFilterer: PremiaOptionFilterer{contract: contract}}, nil
}

// NewPremiaOptionCaller creates a new read-only instance of PremiaOption, bound to a specific deployed contract.
func NewPremiaOptionCaller(address common.Address, caller bind.ContractCaller) (*PremiaOptionCaller, error) {
	contract, err := bindPremiaOption(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionCaller{contract: contract}, nil
}

// NewPremiaOptionTransactor creates a new write-only instance of PremiaOption, bound to a specific deployed contract.
func NewPremiaOptionTransactor(address common.Address, transactor bind.ContractTransactor) (*PremiaOptionTransactor, error) {
	contract, err := bindPremiaOption(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionTransactor{contract: contract}, nil
}

// NewPremiaOptionFilterer creates a new log filterer instance of PremiaOption, bound to a specific deployed contract.
func NewPremiaOptionFilterer(address common.Address, filterer bind.ContractFilterer) (*PremiaOptionFilterer, error) {
	contract, err := bindPremiaOption(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionFilterer{contract: contract}, nil
}

// bindPremiaOption binds a generic wrapper to an already deployed contract.
func bindPremiaOption(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PremiaOptionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PremiaOption *PremiaOptionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PremiaOption.Contract.PremiaOptionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PremiaOption *PremiaOptionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaOption.Contract.PremiaOptionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PremiaOption *PremiaOptionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PremiaOption.Contract.PremiaOptionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PremiaOption *PremiaOptionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PremiaOption.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PremiaOption *PremiaOptionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaOption.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PremiaOption *PremiaOptionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PremiaOption.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PremiaOption *PremiaOptionSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PremiaOption.Contract.BalanceOf(&_PremiaOption.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _PremiaOption.Contract.BalanceOf(&_PremiaOption.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PremiaOption *PremiaOptionCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PremiaOption *PremiaOptionSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PremiaOption.Contract.BalanceOfBatch(&_PremiaOption.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_PremiaOption *PremiaOptionCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _PremiaOption.Contract.BalanceOfBatch(&_PremiaOption.CallOpts, accounts, ids)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(address)
func (_PremiaOption *PremiaOptionCaller) Denominator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "denominator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(address)
func (_PremiaOption *PremiaOptionSession) Denominator() (common.Address, error) {
	return _PremiaOption.Contract.Denominator(&_PremiaOption.CallOpts)
}

// Denominator is a free data retrieval call binding the contract method 0x96ce0795.
//
// Solidity: function denominator() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) Denominator() (common.Address, error) {
	return _PremiaOption.Contract.Denominator(&_PremiaOption.CallOpts)
}

// DenominatorDecimals is a free data retrieval call binding the contract method 0x4cc47c0d.
//
// Solidity: function denominatorDecimals() view returns(uint8)
func (_PremiaOption *PremiaOptionCaller) DenominatorDecimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "denominatorDecimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DenominatorDecimals is a free data retrieval call binding the contract method 0x4cc47c0d.
//
// Solidity: function denominatorDecimals() view returns(uint8)
func (_PremiaOption *PremiaOptionSession) DenominatorDecimals() (uint8, error) {
	return _PremiaOption.Contract.DenominatorDecimals(&_PremiaOption.CallOpts)
}

// DenominatorDecimals is a free data retrieval call binding the contract method 0x4cc47c0d.
//
// Solidity: function denominatorDecimals() view returns(uint8)
func (_PremiaOption *PremiaOptionCallerSession) DenominatorDecimals() (uint8, error) {
	return _PremiaOption.Contract.DenominatorDecimals(&_PremiaOption.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaOption *PremiaOptionCaller) FeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "feeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaOption *PremiaOptionSession) FeeCalculator() (common.Address, error) {
	return _PremiaOption.Contract.FeeCalculator(&_PremiaOption.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) FeeCalculator() (common.Address, error) {
	return _PremiaOption.Contract.FeeCalculator(&_PremiaOption.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaOption *PremiaOptionCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaOption *PremiaOptionSession) FeeRecipient() (common.Address, error) {
	return _PremiaOption.Contract.FeeRecipient(&_PremiaOption.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) FeeRecipient() (common.Address, error) {
	return _PremiaOption.Contract.FeeRecipient(&_PremiaOption.CallOpts)
}

// GetExerciseQuote is a free data retrieval call binding the contract method 0x3e998172.
//
// Solidity: function getExerciseQuote(address _from, (address,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint8) _option, uint256 _amount, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionCaller) GetExerciseQuote(opts *bind.CallOpts, _from common.Address, _option PremiaOptionOptionData, _amount *big.Int, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteExercise, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "getExerciseQuote", _from, _option, _amount, _referrer, _decimals)

	if err != nil {
		return *new(PremiaOptionQuoteExercise), err
	}

	out0 := *abi.ConvertType(out[0], new(PremiaOptionQuoteExercise)).(*PremiaOptionQuoteExercise)

	return out0, err

}

// GetExerciseQuote is a free data retrieval call binding the contract method 0x3e998172.
//
// Solidity: function getExerciseQuote(address _from, (address,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint8) _option, uint256 _amount, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionSession) GetExerciseQuote(_from common.Address, _option PremiaOptionOptionData, _amount *big.Int, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteExercise, error) {
	return _PremiaOption.Contract.GetExerciseQuote(&_PremiaOption.CallOpts, _from, _option, _amount, _referrer, _decimals)
}

// GetExerciseQuote is a free data retrieval call binding the contract method 0x3e998172.
//
// Solidity: function getExerciseQuote(address _from, (address,uint256,uint256,bool,uint256,uint256,uint256,uint256,uint8) _option, uint256 _amount, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionCallerSession) GetExerciseQuote(_from common.Address, _option PremiaOptionOptionData, _amount *big.Int, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteExercise, error) {
	return _PremiaOption.Contract.GetExerciseQuote(&_PremiaOption.CallOpts, _from, _option, _amount, _referrer, _decimals)
}

// GetOptionId is a free data retrieval call binding the contract method 0x263621f3.
//
// Solidity: function getOptionId(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) GetOptionId(opts *bind.CallOpts, _token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "getOptionId", _token, _expiration, _strikePrice, _isCall)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOptionId is a free data retrieval call binding the contract method 0x263621f3.
//
// Solidity: function getOptionId(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) view returns(uint256)
func (_PremiaOption *PremiaOptionSession) GetOptionId(_token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*big.Int, error) {
	return _PremiaOption.Contract.GetOptionId(&_PremiaOption.CallOpts, _token, _expiration, _strikePrice, _isCall)
}

// GetOptionId is a free data retrieval call binding the contract method 0x263621f3.
//
// Solidity: function getOptionId(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) GetOptionId(_token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*big.Int, error) {
	return _PremiaOption.Contract.GetOptionId(&_PremiaOption.CallOpts, _token, _expiration, _strikePrice, _isCall)
}

// GetWriteQuote is a free data retrieval call binding the contract method 0xd35daf2b.
//
// Solidity: function getWriteQuote(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionCaller) GetWriteQuote(opts *bind.CallOpts, _from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteWrite, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "getWriteQuote", _from, _option, _referrer, _decimals)

	if err != nil {
		return *new(PremiaOptionQuoteWrite), err
	}

	out0 := *abi.ConvertType(out[0], new(PremiaOptionQuoteWrite)).(*PremiaOptionQuoteWrite)

	return out0, err

}

// GetWriteQuote is a free data retrieval call binding the contract method 0xd35daf2b.
//
// Solidity: function getWriteQuote(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionSession) GetWriteQuote(_from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteWrite, error) {
	return _PremiaOption.Contract.GetWriteQuote(&_PremiaOption.CallOpts, _from, _option, _referrer, _decimals)
}

// GetWriteQuote is a free data retrieval call binding the contract method 0xd35daf2b.
//
// Solidity: function getWriteQuote(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer, uint8 _decimals) view returns((address,uint256,uint8,uint256,uint256))
func (_PremiaOption *PremiaOptionCallerSession) GetWriteQuote(_from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address, _decimals uint8) (PremiaOptionQuoteWrite, error) {
	return _PremiaOption.Contract.GetWriteQuote(&_PremiaOption.CallOpts, _from, _option, _referrer, _decimals)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PremiaOption *PremiaOptionCaller) IsApprovedForAll(opts *bind.CallOpts, account common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "isApprovedForAll", account, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PremiaOption *PremiaOptionSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PremiaOption.Contract.IsApprovedForAll(&_PremiaOption.CallOpts, account, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address account, address operator) view returns(bool)
func (_PremiaOption *PremiaOptionCallerSession) IsApprovedForAll(account common.Address, operator common.Address) (bool, error) {
	return _PremiaOption.Contract.IsApprovedForAll(&_PremiaOption.CallOpts, account, operator)
}

// MaxExpiration is a free data retrieval call binding the contract method 0x25bee3e9.
//
// Solidity: function maxExpiration() view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) MaxExpiration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "maxExpiration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxExpiration is a free data retrieval call binding the contract method 0x25bee3e9.
//
// Solidity: function maxExpiration() view returns(uint256)
func (_PremiaOption *PremiaOptionSession) MaxExpiration() (*big.Int, error) {
	return _PremiaOption.Contract.MaxExpiration(&_PremiaOption.CallOpts)
}

// MaxExpiration is a free data retrieval call binding the contract method 0x25bee3e9.
//
// Solidity: function maxExpiration() view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) MaxExpiration() (*big.Int, error) {
	return _PremiaOption.Contract.MaxExpiration(&_PremiaOption.CallOpts)
}

// NbWritten is a free data retrieval call binding the contract method 0x410e4e26.
//
// Solidity: function nbWritten(address , uint256 ) view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) NbWritten(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "nbWritten", arg0, arg1)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NbWritten is a free data retrieval call binding the contract method 0x410e4e26.
//
// Solidity: function nbWritten(address , uint256 ) view returns(uint256)
func (_PremiaOption *PremiaOptionSession) NbWritten(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PremiaOption.Contract.NbWritten(&_PremiaOption.CallOpts, arg0, arg1)
}

// NbWritten is a free data retrieval call binding the contract method 0x410e4e26.
//
// Solidity: function nbWritten(address , uint256 ) view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) NbWritten(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PremiaOption.Contract.NbWritten(&_PremiaOption.CallOpts, arg0, arg1)
}

// NextOptionId is a free data retrieval call binding the contract method 0x9215abb0.
//
// Solidity: function nextOptionId() view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) NextOptionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "nextOptionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextOptionId is a free data retrieval call binding the contract method 0x9215abb0.
//
// Solidity: function nextOptionId() view returns(uint256)
func (_PremiaOption *PremiaOptionSession) NextOptionId() (*big.Int, error) {
	return _PremiaOption.Contract.NextOptionId(&_PremiaOption.CallOpts)
}

// NextOptionId is a free data retrieval call binding the contract method 0x9215abb0.
//
// Solidity: function nextOptionId() view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) NextOptionId() (*big.Int, error) {
	return _PremiaOption.Contract.NextOptionId(&_PremiaOption.CallOpts)
}

// OptionData is a free data retrieval call binding the contract method 0x1f01a794.
//
// Solidity: function optionData(uint256 ) view returns(address token, uint256 strikePrice, uint256 expiration, bool isCall, uint256 claimsPreExp, uint256 claimsPostExp, uint256 exercised, uint256 supply, uint8 decimals)
func (_PremiaOption *PremiaOptionCaller) OptionData(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Token         common.Address
	StrikePrice   *big.Int
	Expiration    *big.Int
	IsCall        bool
	ClaimsPreExp  *big.Int
	ClaimsPostExp *big.Int
	Exercised     *big.Int
	Supply        *big.Int
	Decimals      uint8
}, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "optionData", arg0)

	outstruct := new(struct {
		Token         common.Address
		StrikePrice   *big.Int
		Expiration    *big.Int
		IsCall        bool
		ClaimsPreExp  *big.Int
		ClaimsPostExp *big.Int
		Exercised     *big.Int
		Supply        *big.Int
		Decimals      uint8
	})

	outstruct.Token = out[0].(common.Address)
	outstruct.StrikePrice = out[1].(*big.Int)
	outstruct.Expiration = out[2].(*big.Int)
	outstruct.IsCall = out[3].(bool)
	outstruct.ClaimsPreExp = out[4].(*big.Int)
	outstruct.ClaimsPostExp = out[5].(*big.Int)
	outstruct.Exercised = out[6].(*big.Int)
	outstruct.Supply = out[7].(*big.Int)
	outstruct.Decimals = out[8].(uint8)

	return *outstruct, err

}

// OptionData is a free data retrieval call binding the contract method 0x1f01a794.
//
// Solidity: function optionData(uint256 ) view returns(address token, uint256 strikePrice, uint256 expiration, bool isCall, uint256 claimsPreExp, uint256 claimsPostExp, uint256 exercised, uint256 supply, uint8 decimals)
func (_PremiaOption *PremiaOptionSession) OptionData(arg0 *big.Int) (struct {
	Token         common.Address
	StrikePrice   *big.Int
	Expiration    *big.Int
	IsCall        bool
	ClaimsPreExp  *big.Int
	ClaimsPostExp *big.Int
	Exercised     *big.Int
	Supply        *big.Int
	Decimals      uint8
}, error) {
	return _PremiaOption.Contract.OptionData(&_PremiaOption.CallOpts, arg0)
}

// OptionData is a free data retrieval call binding the contract method 0x1f01a794.
//
// Solidity: function optionData(uint256 ) view returns(address token, uint256 strikePrice, uint256 expiration, bool isCall, uint256 claimsPreExp, uint256 claimsPostExp, uint256 exercised, uint256 supply, uint8 decimals)
func (_PremiaOption *PremiaOptionCallerSession) OptionData(arg0 *big.Int) (struct {
	Token         common.Address
	StrikePrice   *big.Int
	Expiration    *big.Int
	IsCall        bool
	ClaimsPreExp  *big.Int
	ClaimsPostExp *big.Int
	Exercised     *big.Int
	Supply        *big.Int
	Decimals      uint8
}, error) {
	return _PremiaOption.Contract.OptionData(&_PremiaOption.CallOpts, arg0)
}

// Options is a free data retrieval call binding the contract method 0xb4ac9460.
//
// Solidity: function options(address , uint256 , uint256 , bool ) view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) Options(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 bool) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "options", arg0, arg1, arg2, arg3)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Options is a free data retrieval call binding the contract method 0xb4ac9460.
//
// Solidity: function options(address , uint256 , uint256 , bool ) view returns(uint256)
func (_PremiaOption *PremiaOptionSession) Options(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 bool) (*big.Int, error) {
	return _PremiaOption.Contract.Options(&_PremiaOption.CallOpts, arg0, arg1, arg2, arg3)
}

// Options is a free data retrieval call binding the contract method 0xb4ac9460.
//
// Solidity: function options(address , uint256 , uint256 , bool ) view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) Options(arg0 common.Address, arg1 *big.Int, arg2 *big.Int, arg3 bool) (*big.Int, error) {
	return _PremiaOption.Contract.Options(&_PremiaOption.CallOpts, arg0, arg1, arg2, arg3)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaOption *PremiaOptionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaOption *PremiaOptionSession) Owner() (common.Address, error) {
	return _PremiaOption.Contract.Owner(&_PremiaOption.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) Owner() (common.Address, error) {
	return _PremiaOption.Contract.Owner(&_PremiaOption.CallOpts)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 tokenAmount, uint256 denominatorAmount)
func (_PremiaOption *PremiaOptionCaller) Pools(opts *bind.CallOpts, arg0 *big.Int) (struct {
	TokenAmount       *big.Int
	DenominatorAmount *big.Int
}, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "pools", arg0)

	outstruct := new(struct {
		TokenAmount       *big.Int
		DenominatorAmount *big.Int
	})

	outstruct.TokenAmount = out[0].(*big.Int)
	outstruct.DenominatorAmount = out[1].(*big.Int)

	return *outstruct, err

}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 tokenAmount, uint256 denominatorAmount)
func (_PremiaOption *PremiaOptionSession) Pools(arg0 *big.Int) (struct {
	TokenAmount       *big.Int
	DenominatorAmount *big.Int
}, error) {
	return _PremiaOption.Contract.Pools(&_PremiaOption.CallOpts, arg0)
}

// Pools is a free data retrieval call binding the contract method 0xac4afa38.
//
// Solidity: function pools(uint256 ) view returns(uint256 tokenAmount, uint256 denominatorAmount)
func (_PremiaOption *PremiaOptionCallerSession) Pools(arg0 *big.Int) (struct {
	TokenAmount       *big.Int
	DenominatorAmount *big.Int
}, error) {
	return _PremiaOption.Contract.Pools(&_PremiaOption.CallOpts, arg0)
}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaOption *PremiaOptionCaller) PremiaReferral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "premiaReferral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaOption *PremiaOptionSession) PremiaReferral() (common.Address, error) {
	return _PremiaOption.Contract.PremiaReferral(&_PremiaOption.CallOpts)
}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) PremiaReferral() (common.Address, error) {
	return _PremiaOption.Contract.PremiaReferral(&_PremiaOption.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PremiaOption *PremiaOptionCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PremiaOption *PremiaOptionSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PremiaOption.Contract.SupportsInterface(&_PremiaOption.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_PremiaOption *PremiaOptionCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _PremiaOption.Contract.SupportsInterface(&_PremiaOption.CallOpts, interfaceId)
}

// TokenStrikeIncrement is a free data retrieval call binding the contract method 0x48af1968.
//
// Solidity: function tokenStrikeIncrement(address ) view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) TokenStrikeIncrement(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "tokenStrikeIncrement", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenStrikeIncrement is a free data retrieval call binding the contract method 0x48af1968.
//
// Solidity: function tokenStrikeIncrement(address ) view returns(uint256)
func (_PremiaOption *PremiaOptionSession) TokenStrikeIncrement(arg0 common.Address) (*big.Int, error) {
	return _PremiaOption.Contract.TokenStrikeIncrement(&_PremiaOption.CallOpts, arg0)
}

// TokenStrikeIncrement is a free data retrieval call binding the contract method 0x48af1968.
//
// Solidity: function tokenStrikeIncrement(address ) view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) TokenStrikeIncrement(arg0 common.Address) (*big.Int, error) {
	return _PremiaOption.Contract.TokenStrikeIncrement(&_PremiaOption.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionCaller) Tokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "tokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _PremiaOption.Contract.Tokens(&_PremiaOption.CallOpts, arg0)
}

// Tokens is a free data retrieval call binding the contract method 0x4f64b2be.
//
// Solidity: function tokens(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) Tokens(arg0 *big.Int) (common.Address, error) {
	return _PremiaOption.Contract.Tokens(&_PremiaOption.CallOpts, arg0)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PremiaOption *PremiaOptionCaller) TokensLength(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "tokensLength")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PremiaOption *PremiaOptionSession) TokensLength() (*big.Int, error) {
	return _PremiaOption.Contract.TokensLength(&_PremiaOption.CallOpts)
}

// TokensLength is a free data retrieval call binding the contract method 0xd92fc67b.
//
// Solidity: function tokensLength() view returns(uint256)
func (_PremiaOption *PremiaOptionCallerSession) TokensLength() (*big.Int, error) {
	return _PremiaOption.Contract.TokensLength(&_PremiaOption.CallOpts)
}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaOption *PremiaOptionCaller) UPremia(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "uPremia")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaOption *PremiaOptionSession) UPremia() (common.Address, error) {
	return _PremiaOption.Contract.UPremia(&_PremiaOption.CallOpts)
}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) UPremia() (common.Address, error) {
	return _PremiaOption.Contract.UPremia(&_PremiaOption.CallOpts)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PremiaOption *PremiaOptionCaller) Uri(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "uri", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PremiaOption *PremiaOptionSession) Uri(arg0 *big.Int) (string, error) {
	return _PremiaOption.Contract.Uri(&_PremiaOption.CallOpts, arg0)
}

// Uri is a free data retrieval call binding the contract method 0x0e89341c.
//
// Solidity: function uri(uint256 ) view returns(string)
func (_PremiaOption *PremiaOptionCallerSession) Uri(arg0 *big.Int) (string, error) {
	return _PremiaOption.Contract.Uri(&_PremiaOption.CallOpts, arg0)
}

// WhitelistedUniswapRouters is a free data retrieval call binding the contract method 0xdbe5f0eb.
//
// Solidity: function whitelistedUniswapRouters(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionCaller) WhitelistedUniswapRouters(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _PremiaOption.contract.Call(opts, &out, "whitelistedUniswapRouters", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WhitelistedUniswapRouters is a free data retrieval call binding the contract method 0xdbe5f0eb.
//
// Solidity: function whitelistedUniswapRouters(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionSession) WhitelistedUniswapRouters(arg0 *big.Int) (common.Address, error) {
	return _PremiaOption.Contract.WhitelistedUniswapRouters(&_PremiaOption.CallOpts, arg0)
}

// WhitelistedUniswapRouters is a free data retrieval call binding the contract method 0xdbe5f0eb.
//
// Solidity: function whitelistedUniswapRouters(uint256 ) view returns(address)
func (_PremiaOption *PremiaOptionCallerSession) WhitelistedUniswapRouters(arg0 *big.Int) (common.Address, error) {
	return _PremiaOption.Contract.WhitelistedUniswapRouters(&_PremiaOption.CallOpts, arg0)
}

// CancelOption is a paid mutator transaction binding the contract method 0x30b025eb.
//
// Solidity: function cancelOption(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactor) CancelOption(opts *bind.TransactOpts, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "cancelOption", _optionId, _amount)
}

// CancelOption is a paid mutator transaction binding the contract method 0x30b025eb.
//
// Solidity: function cancelOption(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionSession) CancelOption(_optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.CancelOption(&_PremiaOption.TransactOpts, _optionId, _amount)
}

// CancelOption is a paid mutator transaction binding the contract method 0x30b025eb.
//
// Solidity: function cancelOption(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactorSession) CancelOption(_optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.CancelOption(&_PremiaOption.TransactOpts, _optionId, _amount)
}

// CancelOptionFrom is a paid mutator transaction binding the contract method 0xf3944b27.
//
// Solidity: function cancelOptionFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactor) CancelOptionFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "cancelOptionFrom", _from, _optionId, _amount)
}

// CancelOptionFrom is a paid mutator transaction binding the contract method 0xf3944b27.
//
// Solidity: function cancelOptionFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionSession) CancelOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.CancelOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount)
}

// CancelOptionFrom is a paid mutator transaction binding the contract method 0xf3944b27.
//
// Solidity: function cancelOptionFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactorSession) CancelOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.CancelOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount)
}

// ExerciseOption is a paid mutator transaction binding the contract method 0x3a5b7c82.
//
// Solidity: function exerciseOption(uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionTransactor) ExerciseOption(opts *bind.TransactOpts, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "exerciseOption", _optionId, _amount, _referrer)
}

// ExerciseOption is a paid mutator transaction binding the contract method 0x3a5b7c82.
//
// Solidity: function exerciseOption(uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionSession) ExerciseOption(_optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.ExerciseOption(&_PremiaOption.TransactOpts, _optionId, _amount, _referrer)
}

// ExerciseOption is a paid mutator transaction binding the contract method 0x3a5b7c82.
//
// Solidity: function exerciseOption(uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionTransactorSession) ExerciseOption(_optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.ExerciseOption(&_PremiaOption.TransactOpts, _optionId, _amount, _referrer)
}

// ExerciseOptionFrom is a paid mutator transaction binding the contract method 0x2216ef60.
//
// Solidity: function exerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionTransactor) ExerciseOptionFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "exerciseOptionFrom", _from, _optionId, _amount, _referrer)
}

// ExerciseOptionFrom is a paid mutator transaction binding the contract method 0x2216ef60.
//
// Solidity: function exerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionSession) ExerciseOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.ExerciseOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer)
}

// ExerciseOptionFrom is a paid mutator transaction binding the contract method 0x2216ef60.
//
// Solidity: function exerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns()
func (_PremiaOption *PremiaOptionTransactorSession) ExerciseOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.ExerciseOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer)
}

// FlashExerciseOption is a paid mutator transaction binding the contract method 0xb7f1dd2e.
//
// Solidity: function flashExerciseOption(uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionTransactor) FlashExerciseOption(opts *bind.TransactOpts, _optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "flashExerciseOption", _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashExerciseOption is a paid mutator transaction binding the contract method 0xb7f1dd2e.
//
// Solidity: function flashExerciseOption(uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionSession) FlashExerciseOption(_optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashExerciseOption(&_PremiaOption.TransactOpts, _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashExerciseOption is a paid mutator transaction binding the contract method 0xb7f1dd2e.
//
// Solidity: function flashExerciseOption(uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionTransactorSession) FlashExerciseOption(_optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashExerciseOption(&_PremiaOption.TransactOpts, _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashExerciseOptionFrom is a paid mutator transaction binding the contract method 0x5c6a945e.
//
// Solidity: function flashExerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionTransactor) FlashExerciseOptionFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "flashExerciseOptionFrom", _from, _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashExerciseOptionFrom is a paid mutator transaction binding the contract method 0x5c6a945e.
//
// Solidity: function flashExerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionSession) FlashExerciseOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashExerciseOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashExerciseOptionFrom is a paid mutator transaction binding the contract method 0x5c6a945e.
//
// Solidity: function flashExerciseOptionFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer, address _router, uint256 _amountInMax, address[] _path) returns()
func (_PremiaOption *PremiaOptionTransactorSession) FlashExerciseOptionFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address, _router common.Address, _amountInMax *big.Int, _path []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashExerciseOptionFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer, _router, _amountInMax, _path)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x02c25b54.
//
// Solidity: function flashLoan(address _tokenAddress, uint256 _amount, address _receiver) returns()
func (_PremiaOption *PremiaOptionTransactor) FlashLoan(opts *bind.TransactOpts, _tokenAddress common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "flashLoan", _tokenAddress, _amount, _receiver)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x02c25b54.
//
// Solidity: function flashLoan(address _tokenAddress, uint256 _amount, address _receiver) returns()
func (_PremiaOption *PremiaOptionSession) FlashLoan(_tokenAddress common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashLoan(&_PremiaOption.TransactOpts, _tokenAddress, _amount, _receiver)
}

// FlashLoan is a paid mutator transaction binding the contract method 0x02c25b54.
//
// Solidity: function flashLoan(address _tokenAddress, uint256 _amount, address _receiver) returns()
func (_PremiaOption *PremiaOptionTransactorSession) FlashLoan(_tokenAddress common.Address, _amount *big.Int, _receiver common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.FlashLoan(&_PremiaOption.TransactOpts, _tokenAddress, _amount, _receiver)
}

// GetOptionIdOrCreate is a paid mutator transaction binding the contract method 0xbf9092ce.
//
// Solidity: function getOptionIdOrCreate(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) returns(uint256)
func (_PremiaOption *PremiaOptionTransactor) GetOptionIdOrCreate(opts *bind.TransactOpts, _token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "getOptionIdOrCreate", _token, _expiration, _strikePrice, _isCall)
}

// GetOptionIdOrCreate is a paid mutator transaction binding the contract method 0xbf9092ce.
//
// Solidity: function getOptionIdOrCreate(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) returns(uint256)
func (_PremiaOption *PremiaOptionSession) GetOptionIdOrCreate(_token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*types.Transaction, error) {
	return _PremiaOption.Contract.GetOptionIdOrCreate(&_PremiaOption.TransactOpts, _token, _expiration, _strikePrice, _isCall)
}

// GetOptionIdOrCreate is a paid mutator transaction binding the contract method 0xbf9092ce.
//
// Solidity: function getOptionIdOrCreate(address _token, uint256 _expiration, uint256 _strikePrice, bool _isCall) returns(uint256)
func (_PremiaOption *PremiaOptionTransactorSession) GetOptionIdOrCreate(_token common.Address, _expiration *big.Int, _strikePrice *big.Int, _isCall bool) (*types.Transaction, error) {
	return _PremiaOption.Contract.GetOptionIdOrCreate(&_PremiaOption.TransactOpts, _token, _expiration, _strikePrice, _isCall)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaOption *PremiaOptionTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaOption *PremiaOptionSession) RenounceOwnership() (*types.Transaction, error) {
	return _PremiaOption.Contract.RenounceOwnership(&_PremiaOption.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaOption *PremiaOptionTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PremiaOption.Contract.RenounceOwnership(&_PremiaOption.TransactOpts)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PremiaOption *PremiaOptionTransactor) SafeBatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "safeBatchTransferFrom", from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PremiaOption *PremiaOptionSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.Contract.SafeBatchTransferFrom(&_PremiaOption.TransactOpts, from, to, ids, amounts, data)
}

// SafeBatchTransferFrom is a paid mutator transaction binding the contract method 0x2eb2c2d6.
//
// Solidity: function safeBatchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts, bytes data) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SafeBatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.Contract.SafeBatchTransferFrom(&_PremiaOption.TransactOpts, from, to, ids, amounts, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PremiaOption *PremiaOptionTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "safeTransferFrom", from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PremiaOption *PremiaOptionSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.Contract.SafeTransferFrom(&_PremiaOption.TransactOpts, from, to, id, amount, data)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0xf242432a.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 id, uint256 amount, bytes data) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SafeTransferFrom(from common.Address, to common.Address, id *big.Int, amount *big.Int, data []byte) (*types.Transaction, error) {
	return _PremiaOption.Contract.SafeTransferFrom(&_PremiaOption.TransactOpts, from, to, id, amount, data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PremiaOption *PremiaOptionTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PremiaOption *PremiaOptionSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetApprovalForAll(&_PremiaOption.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetApprovalForAll(&_PremiaOption.TransactOpts, operator, approved)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaOption *PremiaOptionTransactor) SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setFeeCalculator", _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaOption *PremiaOptionSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetFeeCalculator(&_PremiaOption.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetFeeCalculator(&_PremiaOption.TransactOpts, _feeCalculator)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaOption *PremiaOptionTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaOption *PremiaOptionSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetFeeRecipient(&_PremiaOption.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetFeeRecipient(&_PremiaOption.TransactOpts, _feeRecipient)
}

// SetMaxExpiration is a paid mutator transaction binding the contract method 0x0676695b.
//
// Solidity: function setMaxExpiration(uint256 _max) returns()
func (_PremiaOption *PremiaOptionTransactor) SetMaxExpiration(opts *bind.TransactOpts, _max *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setMaxExpiration", _max)
}

// SetMaxExpiration is a paid mutator transaction binding the contract method 0x0676695b.
//
// Solidity: function setMaxExpiration(uint256 _max) returns()
func (_PremiaOption *PremiaOptionSession) SetMaxExpiration(_max *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetMaxExpiration(&_PremiaOption.TransactOpts, _max)
}

// SetMaxExpiration is a paid mutator transaction binding the contract method 0x0676695b.
//
// Solidity: function setMaxExpiration(uint256 _max) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetMaxExpiration(_max *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetMaxExpiration(&_PremiaOption.TransactOpts, _max)
}

// SetPremiaReferral is a paid mutator transaction binding the contract method 0xf18e1a94.
//
// Solidity: function setPremiaReferral(address _premiaReferral) returns()
func (_PremiaOption *PremiaOptionTransactor) SetPremiaReferral(opts *bind.TransactOpts, _premiaReferral common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setPremiaReferral", _premiaReferral)
}

// SetPremiaReferral is a paid mutator transaction binding the contract method 0xf18e1a94.
//
// Solidity: function setPremiaReferral(address _premiaReferral) returns()
func (_PremiaOption *PremiaOptionSession) SetPremiaReferral(_premiaReferral common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetPremiaReferral(&_PremiaOption.TransactOpts, _premiaReferral)
}

// SetPremiaReferral is a paid mutator transaction binding the contract method 0xf18e1a94.
//
// Solidity: function setPremiaReferral(address _premiaReferral) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetPremiaReferral(_premiaReferral common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetPremiaReferral(&_PremiaOption.TransactOpts, _premiaReferral)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaOption *PremiaOptionTransactor) SetPremiaUncutErc20(opts *bind.TransactOpts, _uPremia common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setPremiaUncutErc20", _uPremia)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaOption *PremiaOptionSession) SetPremiaUncutErc20(_uPremia common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetPremiaUncutErc20(&_PremiaOption.TransactOpts, _uPremia)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetPremiaUncutErc20(_uPremia common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetPremiaUncutErc20(&_PremiaOption.TransactOpts, _uPremia)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _strikePriceIncrement) returns()
func (_PremiaOption *PremiaOptionTransactor) SetTokens(opts *bind.TransactOpts, _tokens []common.Address, _strikePriceIncrement []*big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setTokens", _tokens, _strikePriceIncrement)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _strikePriceIncrement) returns()
func (_PremiaOption *PremiaOptionSession) SetTokens(_tokens []common.Address, _strikePriceIncrement []*big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetTokens(&_PremiaOption.TransactOpts, _tokens, _strikePriceIncrement)
}

// SetTokens is a paid mutator transaction binding the contract method 0xc8390a48.
//
// Solidity: function setTokens(address[] _tokens, uint256[] _strikePriceIncrement) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetTokens(_tokens []common.Address, _strikePriceIncrement []*big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetTokens(&_PremiaOption.TransactOpts, _tokens, _strikePriceIncrement)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newUri) returns()
func (_PremiaOption *PremiaOptionTransactor) SetURI(opts *bind.TransactOpts, _newUri string) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setURI", _newUri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newUri) returns()
func (_PremiaOption *PremiaOptionSession) SetURI(_newUri string) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetURI(&_PremiaOption.TransactOpts, _newUri)
}

// SetURI is a paid mutator transaction binding the contract method 0x02fe5305.
//
// Solidity: function setURI(string _newUri) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetURI(_newUri string) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetURI(&_PremiaOption.TransactOpts, _newUri)
}

// SetWhitelistedUniswapRouters is a paid mutator transaction binding the contract method 0x9fe9f848.
//
// Solidity: function setWhitelistedUniswapRouters(address[] _addrList) returns()
func (_PremiaOption *PremiaOptionTransactor) SetWhitelistedUniswapRouters(opts *bind.TransactOpts, _addrList []common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "setWhitelistedUniswapRouters", _addrList)
}

// SetWhitelistedUniswapRouters is a paid mutator transaction binding the contract method 0x9fe9f848.
//
// Solidity: function setWhitelistedUniswapRouters(address[] _addrList) returns()
func (_PremiaOption *PremiaOptionSession) SetWhitelistedUniswapRouters(_addrList []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetWhitelistedUniswapRouters(&_PremiaOption.TransactOpts, _addrList)
}

// SetWhitelistedUniswapRouters is a paid mutator transaction binding the contract method 0x9fe9f848.
//
// Solidity: function setWhitelistedUniswapRouters(address[] _addrList) returns()
func (_PremiaOption *PremiaOptionTransactorSession) SetWhitelistedUniswapRouters(_addrList []common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.SetWhitelistedUniswapRouters(&_PremiaOption.TransactOpts, _addrList)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaOption *PremiaOptionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaOption *PremiaOptionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.TransferOwnership(&_PremiaOption.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaOption *PremiaOptionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.TransferOwnership(&_PremiaOption.TransactOpts, newOwner)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionTransactor) Withdraw(opts *bind.TransactOpts, _optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "withdraw", _optionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionSession) Withdraw(_optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.Withdraw(&_PremiaOption.TransactOpts, _optionId)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionTransactorSession) Withdraw(_optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.Withdraw(&_PremiaOption.TransactOpts, _optionId)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address _from, uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionTransactor) WithdrawFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "withdrawFrom", _from, _optionId)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address _from, uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionSession) WithdrawFrom(_from common.Address, _optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawFrom(&_PremiaOption.TransactOpts, _from, _optionId)
}

// WithdrawFrom is a paid mutator transaction binding the contract method 0x9470b0bd.
//
// Solidity: function withdrawFrom(address _from, uint256 _optionId) returns()
func (_PremiaOption *PremiaOptionTransactorSession) WithdrawFrom(_from common.Address, _optionId *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawFrom(&_PremiaOption.TransactOpts, _from, _optionId)
}

// WithdrawPreExpiration is a paid mutator transaction binding the contract method 0x5c070b65.
//
// Solidity: function withdrawPreExpiration(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactor) WithdrawPreExpiration(opts *bind.TransactOpts, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "withdrawPreExpiration", _optionId, _amount)
}

// WithdrawPreExpiration is a paid mutator transaction binding the contract method 0x5c070b65.
//
// Solidity: function withdrawPreExpiration(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionSession) WithdrawPreExpiration(_optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawPreExpiration(&_PremiaOption.TransactOpts, _optionId, _amount)
}

// WithdrawPreExpiration is a paid mutator transaction binding the contract method 0x5c070b65.
//
// Solidity: function withdrawPreExpiration(uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactorSession) WithdrawPreExpiration(_optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawPreExpiration(&_PremiaOption.TransactOpts, _optionId, _amount)
}

// WithdrawPreExpirationFrom is a paid mutator transaction binding the contract method 0xa852c7e7.
//
// Solidity: function withdrawPreExpirationFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactor) WithdrawPreExpirationFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "withdrawPreExpirationFrom", _from, _optionId, _amount)
}

// WithdrawPreExpirationFrom is a paid mutator transaction binding the contract method 0xa852c7e7.
//
// Solidity: function withdrawPreExpirationFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionSession) WithdrawPreExpirationFrom(_from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawPreExpirationFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount)
}

// WithdrawPreExpirationFrom is a paid mutator transaction binding the contract method 0xa852c7e7.
//
// Solidity: function withdrawPreExpirationFrom(address _from, uint256 _optionId, uint256 _amount) returns()
func (_PremiaOption *PremiaOptionTransactorSession) WithdrawPreExpirationFrom(_from common.Address, _optionId *big.Int, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaOption.Contract.WithdrawPreExpirationFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount)
}

// WriteOption is a paid mutator transaction binding the contract method 0x0e2e8c54.
//
// Solidity: function writeOption((address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactor) WriteOption(opts *bind.TransactOpts, _option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "writeOption", _option, _referrer)
}

// WriteOption is a paid mutator transaction binding the contract method 0x0e2e8c54.
//
// Solidity: function writeOption((address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionSession) WriteOption(_option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOption(&_PremiaOption.TransactOpts, _option, _referrer)
}

// WriteOption is a paid mutator transaction binding the contract method 0x0e2e8c54.
//
// Solidity: function writeOption((address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactorSession) WriteOption(_option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOption(&_PremiaOption.TransactOpts, _option, _referrer)
}

// WriteOptionFrom is a paid mutator transaction binding the contract method 0x8e5a2892.
//
// Solidity: function writeOptionFrom(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactor) WriteOptionFrom(opts *bind.TransactOpts, _from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "writeOptionFrom", _from, _option, _referrer)
}

// WriteOptionFrom is a paid mutator transaction binding the contract method 0x8e5a2892.
//
// Solidity: function writeOptionFrom(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionSession) WriteOptionFrom(_from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOptionFrom(&_PremiaOption.TransactOpts, _from, _option, _referrer)
}

// WriteOptionFrom is a paid mutator transaction binding the contract method 0x8e5a2892.
//
// Solidity: function writeOptionFrom(address _from, (address,uint256,uint256,uint256,bool) _option, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactorSession) WriteOptionFrom(_from common.Address, _option PremiaOptionOptionWriteArgs, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOptionFrom(&_PremiaOption.TransactOpts, _from, _option, _referrer)
}

// WriteOptionWithIdFrom is a paid mutator transaction binding the contract method 0xd48b0e52.
//
// Solidity: function writeOptionWithIdFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactor) WriteOptionWithIdFrom(opts *bind.TransactOpts, _from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.contract.Transact(opts, "writeOptionWithIdFrom", _from, _optionId, _amount, _referrer)
}

// WriteOptionWithIdFrom is a paid mutator transaction binding the contract method 0xd48b0e52.
//
// Solidity: function writeOptionWithIdFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionSession) WriteOptionWithIdFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOptionWithIdFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer)
}

// WriteOptionWithIdFrom is a paid mutator transaction binding the contract method 0xd48b0e52.
//
// Solidity: function writeOptionWithIdFrom(address _from, uint256 _optionId, uint256 _amount, address _referrer) returns(uint256)
func (_PremiaOption *PremiaOptionTransactorSession) WriteOptionWithIdFrom(_from common.Address, _optionId *big.Int, _amount *big.Int, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaOption.Contract.WriteOptionWithIdFrom(&_PremiaOption.TransactOpts, _from, _optionId, _amount, _referrer)
}

// PremiaOptionApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the PremiaOption contract.
type PremiaOptionApprovalForAllIterator struct {
	Event *PremiaOptionApprovalForAll // Event containing the contract specifics and raw log

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
func (it *PremiaOptionApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionApprovalForAll)
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
		it.Event = new(PremiaOptionApprovalForAll)
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
func (it *PremiaOptionApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionApprovalForAll represents a ApprovalForAll event raised by the PremiaOption contract.
type PremiaOptionApprovalForAll struct {
	Account  common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PremiaOption *PremiaOptionFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, operator []common.Address) (*PremiaOptionApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionApprovalForAllIterator{contract: _PremiaOption.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PremiaOption *PremiaOptionFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *PremiaOptionApprovalForAll, account []common.Address, operator []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "ApprovalForAll", accountRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionApprovalForAll)
				if err := _PremiaOption.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed operator, bool approved)
func (_PremiaOption *PremiaOptionFilterer) ParseApprovalForAll(log types.Log) (*PremiaOptionApprovalForAll, error) {
	event := new(PremiaOptionApprovalForAll)
	if err := _PremiaOption.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionFeePaidIterator is returned from FilterFeePaid and is used to iterate over the raw logs and unpacked data for FeePaid events raised by the PremiaOption contract.
type PremiaOptionFeePaidIterator struct {
	Event *PremiaOptionFeePaid // Event containing the contract specifics and raw log

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
func (it *PremiaOptionFeePaidIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionFeePaid)
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
		it.Event = new(PremiaOptionFeePaid)
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
func (it *PremiaOptionFeePaidIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionFeePaidIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionFeePaid represents a FeePaid event raised by the PremiaOption contract.
type PremiaOptionFeePaid struct {
	User        common.Address
	Token       common.Address
	Referrer    common.Address
	FeeProtocol *big.Int
	FeeReferrer *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFeePaid is a free log retrieval operation binding the contract event 0x7b025f69f5843a875988abc4e248350e4e5d0ec876e181973034e7a5c5bad884.
//
// Solidity: event FeePaid(address indexed user, address indexed token, address indexed referrer, uint256 feeProtocol, uint256 feeReferrer)
func (_PremiaOption *PremiaOptionFilterer) FilterFeePaid(opts *bind.FilterOpts, user []common.Address, token []common.Address, referrer []common.Address) (*PremiaOptionFeePaidIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "FeePaid", userRule, tokenRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionFeePaidIterator{contract: _PremiaOption.contract, event: "FeePaid", logs: logs, sub: sub}, nil
}

// WatchFeePaid is a free log subscription operation binding the contract event 0x7b025f69f5843a875988abc4e248350e4e5d0ec876e181973034e7a5c5bad884.
//
// Solidity: event FeePaid(address indexed user, address indexed token, address indexed referrer, uint256 feeProtocol, uint256 feeReferrer)
func (_PremiaOption *PremiaOptionFilterer) WatchFeePaid(opts *bind.WatchOpts, sink chan<- *PremiaOptionFeePaid, user []common.Address, token []common.Address, referrer []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var referrerRule []interface{}
	for _, referrerItem := range referrer {
		referrerRule = append(referrerRule, referrerItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "FeePaid", userRule, tokenRule, referrerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionFeePaid)
				if err := _PremiaOption.contract.UnpackLog(event, "FeePaid", log); err != nil {
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

// ParseFeePaid is a log parse operation binding the contract event 0x7b025f69f5843a875988abc4e248350e4e5d0ec876e181973034e7a5c5bad884.
//
// Solidity: event FeePaid(address indexed user, address indexed token, address indexed referrer, uint256 feeProtocol, uint256 feeReferrer)
func (_PremiaOption *PremiaOptionFilterer) ParseFeePaid(log types.Log) (*PremiaOptionFeePaid, error) {
	event := new(PremiaOptionFeePaid)
	if err := _PremiaOption.contract.UnpackLog(event, "FeePaid", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionOptionCancelledIterator is returned from FilterOptionCancelled and is used to iterate over the raw logs and unpacked data for OptionCancelled events raised by the PremiaOption contract.
type PremiaOptionOptionCancelledIterator struct {
	Event *PremiaOptionOptionCancelled // Event containing the contract specifics and raw log

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
func (it *PremiaOptionOptionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionOptionCancelled)
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
		it.Event = new(PremiaOptionOptionCancelled)
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
func (it *PremiaOptionOptionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionOptionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionOptionCancelled represents a OptionCancelled event raised by the PremiaOption contract.
type PremiaOptionOptionCancelled struct {
	Owner    common.Address
	OptionId *big.Int
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOptionCancelled is a free log retrieval operation binding the contract event 0x9035f481e4b25fdafc60883ab62c86fec841bb1c318f5b189ebee1b10ced360b.
//
// Solidity: event OptionCancelled(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) FilterOptionCancelled(opts *bind.FilterOpts, owner []common.Address, optionId []*big.Int, token []common.Address) (*PremiaOptionOptionCancelledIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "OptionCancelled", ownerRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionOptionCancelledIterator{contract: _PremiaOption.contract, event: "OptionCancelled", logs: logs, sub: sub}, nil
}

// WatchOptionCancelled is a free log subscription operation binding the contract event 0x9035f481e4b25fdafc60883ab62c86fec841bb1c318f5b189ebee1b10ced360b.
//
// Solidity: event OptionCancelled(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) WatchOptionCancelled(opts *bind.WatchOpts, sink chan<- *PremiaOptionOptionCancelled, owner []common.Address, optionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "OptionCancelled", ownerRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionOptionCancelled)
				if err := _PremiaOption.contract.UnpackLog(event, "OptionCancelled", log); err != nil {
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

// ParseOptionCancelled is a log parse operation binding the contract event 0x9035f481e4b25fdafc60883ab62c86fec841bb1c318f5b189ebee1b10ced360b.
//
// Solidity: event OptionCancelled(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) ParseOptionCancelled(log types.Log) (*PremiaOptionOptionCancelled, error) {
	event := new(PremiaOptionOptionCancelled)
	if err := _PremiaOption.contract.UnpackLog(event, "OptionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionOptionExercisedIterator is returned from FilterOptionExercised and is used to iterate over the raw logs and unpacked data for OptionExercised events raised by the PremiaOption contract.
type PremiaOptionOptionExercisedIterator struct {
	Event *PremiaOptionOptionExercised // Event containing the contract specifics and raw log

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
func (it *PremiaOptionOptionExercisedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionOptionExercised)
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
		it.Event = new(PremiaOptionOptionExercised)
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
func (it *PremiaOptionOptionExercisedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionOptionExercisedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionOptionExercised represents a OptionExercised event raised by the PremiaOption contract.
type PremiaOptionOptionExercised struct {
	User     common.Address
	OptionId *big.Int
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOptionExercised is a free log retrieval operation binding the contract event 0xa9149acb233878ee232a8279e24b74301705dcfd2319e4bf98bda4b04040afd6.
//
// Solidity: event OptionExercised(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) FilterOptionExercised(opts *bind.FilterOpts, user []common.Address, optionId []*big.Int, token []common.Address) (*PremiaOptionOptionExercisedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "OptionExercised", userRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionOptionExercisedIterator{contract: _PremiaOption.contract, event: "OptionExercised", logs: logs, sub: sub}, nil
}

// WatchOptionExercised is a free log subscription operation binding the contract event 0xa9149acb233878ee232a8279e24b74301705dcfd2319e4bf98bda4b04040afd6.
//
// Solidity: event OptionExercised(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) WatchOptionExercised(opts *bind.WatchOpts, sink chan<- *PremiaOptionOptionExercised, user []common.Address, optionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "OptionExercised", userRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionOptionExercised)
				if err := _PremiaOption.contract.UnpackLog(event, "OptionExercised", log); err != nil {
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

// ParseOptionExercised is a log parse operation binding the contract event 0xa9149acb233878ee232a8279e24b74301705dcfd2319e4bf98bda4b04040afd6.
//
// Solidity: event OptionExercised(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) ParseOptionExercised(log types.Log) (*PremiaOptionOptionExercised, error) {
	event := new(PremiaOptionOptionExercised)
	if err := _PremiaOption.contract.UnpackLog(event, "OptionExercised", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionOptionIdCreatedIterator is returned from FilterOptionIdCreated and is used to iterate over the raw logs and unpacked data for OptionIdCreated events raised by the PremiaOption contract.
type PremiaOptionOptionIdCreatedIterator struct {
	Event *PremiaOptionOptionIdCreated // Event containing the contract specifics and raw log

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
func (it *PremiaOptionOptionIdCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionOptionIdCreated)
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
		it.Event = new(PremiaOptionOptionIdCreated)
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
func (it *PremiaOptionOptionIdCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionOptionIdCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionOptionIdCreated represents a OptionIdCreated event raised by the PremiaOption contract.
type PremiaOptionOptionIdCreated struct {
	OptionId *big.Int
	Token    common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOptionIdCreated is a free log retrieval operation binding the contract event 0xfde1d545a79e30139419cf7880b5c46d455dd9f1dfd88cc6cbc898d348af72ff.
//
// Solidity: event OptionIdCreated(uint256 indexed optionId, address indexed token)
func (_PremiaOption *PremiaOptionFilterer) FilterOptionIdCreated(opts *bind.FilterOpts, optionId []*big.Int, token []common.Address) (*PremiaOptionOptionIdCreatedIterator, error) {

	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "OptionIdCreated", optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionOptionIdCreatedIterator{contract: _PremiaOption.contract, event: "OptionIdCreated", logs: logs, sub: sub}, nil
}

// WatchOptionIdCreated is a free log subscription operation binding the contract event 0xfde1d545a79e30139419cf7880b5c46d455dd9f1dfd88cc6cbc898d348af72ff.
//
// Solidity: event OptionIdCreated(uint256 indexed optionId, address indexed token)
func (_PremiaOption *PremiaOptionFilterer) WatchOptionIdCreated(opts *bind.WatchOpts, sink chan<- *PremiaOptionOptionIdCreated, optionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "OptionIdCreated", optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionOptionIdCreated)
				if err := _PremiaOption.contract.UnpackLog(event, "OptionIdCreated", log); err != nil {
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

// ParseOptionIdCreated is a log parse operation binding the contract event 0xfde1d545a79e30139419cf7880b5c46d455dd9f1dfd88cc6cbc898d348af72ff.
//
// Solidity: event OptionIdCreated(uint256 indexed optionId, address indexed token)
func (_PremiaOption *PremiaOptionFilterer) ParseOptionIdCreated(log types.Log) (*PremiaOptionOptionIdCreated, error) {
	event := new(PremiaOptionOptionIdCreated)
	if err := _PremiaOption.contract.UnpackLog(event, "OptionIdCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionOptionWrittenIterator is returned from FilterOptionWritten and is used to iterate over the raw logs and unpacked data for OptionWritten events raised by the PremiaOption contract.
type PremiaOptionOptionWrittenIterator struct {
	Event *PremiaOptionOptionWritten // Event containing the contract specifics and raw log

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
func (it *PremiaOptionOptionWrittenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionOptionWritten)
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
		it.Event = new(PremiaOptionOptionWritten)
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
func (it *PremiaOptionOptionWrittenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionOptionWrittenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionOptionWritten represents a OptionWritten event raised by the PremiaOption contract.
type PremiaOptionOptionWritten struct {
	Owner    common.Address
	OptionId *big.Int
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOptionWritten is a free log retrieval operation binding the contract event 0xb621f8758571be4a426feacc13e4bb7a476c12cda0805ba65f14f89d1a0a05c3.
//
// Solidity: event OptionWritten(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) FilterOptionWritten(opts *bind.FilterOpts, owner []common.Address, optionId []*big.Int, token []common.Address) (*PremiaOptionOptionWrittenIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "OptionWritten", ownerRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionOptionWrittenIterator{contract: _PremiaOption.contract, event: "OptionWritten", logs: logs, sub: sub}, nil
}

// WatchOptionWritten is a free log subscription operation binding the contract event 0xb621f8758571be4a426feacc13e4bb7a476c12cda0805ba65f14f89d1a0a05c3.
//
// Solidity: event OptionWritten(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) WatchOptionWritten(opts *bind.WatchOpts, sink chan<- *PremiaOptionOptionWritten, owner []common.Address, optionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "OptionWritten", ownerRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionOptionWritten)
				if err := _PremiaOption.contract.UnpackLog(event, "OptionWritten", log); err != nil {
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

// ParseOptionWritten is a log parse operation binding the contract event 0xb621f8758571be4a426feacc13e4bb7a476c12cda0805ba65f14f89d1a0a05c3.
//
// Solidity: event OptionWritten(address indexed owner, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) ParseOptionWritten(log types.Log) (*PremiaOptionOptionWritten, error) {
	event := new(PremiaOptionOptionWritten)
	if err := _PremiaOption.contract.UnpackLog(event, "OptionWritten", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PremiaOption contract.
type PremiaOptionOwnershipTransferredIterator struct {
	Event *PremiaOptionOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PremiaOptionOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionOwnershipTransferred)
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
		it.Event = new(PremiaOptionOwnershipTransferred)
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
func (it *PremiaOptionOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionOwnershipTransferred represents a OwnershipTransferred event raised by the PremiaOption contract.
type PremiaOptionOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PremiaOption *PremiaOptionFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PremiaOptionOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionOwnershipTransferredIterator{contract: _PremiaOption.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PremiaOption *PremiaOptionFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PremiaOptionOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionOwnershipTransferred)
				if err := _PremiaOption.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PremiaOption *PremiaOptionFilterer) ParseOwnershipTransferred(log types.Log) (*PremiaOptionOwnershipTransferred, error) {
	event := new(PremiaOptionOwnershipTransferred)
	if err := _PremiaOption.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionSetTokenIterator is returned from FilterSetToken and is used to iterate over the raw logs and unpacked data for SetToken events raised by the PremiaOption contract.
type PremiaOptionSetTokenIterator struct {
	Event *PremiaOptionSetToken // Event containing the contract specifics and raw log

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
func (it *PremiaOptionSetTokenIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionSetToken)
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
		it.Event = new(PremiaOptionSetToken)
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
func (it *PremiaOptionSetTokenIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionSetTokenIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionSetToken represents a SetToken event raised by the PremiaOption contract.
type PremiaOptionSetToken struct {
	Token                common.Address
	StrikePriceIncrement *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterSetToken is a free log retrieval operation binding the contract event 0x720764556647dd167f4229d6a4255ac86018e302a50fc29dd67a70edb7b314d0.
//
// Solidity: event SetToken(address indexed token, uint256 strikePriceIncrement)
func (_PremiaOption *PremiaOptionFilterer) FilterSetToken(opts *bind.FilterOpts, token []common.Address) (*PremiaOptionSetTokenIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "SetToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionSetTokenIterator{contract: _PremiaOption.contract, event: "SetToken", logs: logs, sub: sub}, nil
}

// WatchSetToken is a free log subscription operation binding the contract event 0x720764556647dd167f4229d6a4255ac86018e302a50fc29dd67a70edb7b314d0.
//
// Solidity: event SetToken(address indexed token, uint256 strikePriceIncrement)
func (_PremiaOption *PremiaOptionFilterer) WatchSetToken(opts *bind.WatchOpts, sink chan<- *PremiaOptionSetToken, token []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "SetToken", tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionSetToken)
				if err := _PremiaOption.contract.UnpackLog(event, "SetToken", log); err != nil {
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

// ParseSetToken is a log parse operation binding the contract event 0x720764556647dd167f4229d6a4255ac86018e302a50fc29dd67a70edb7b314d0.
//
// Solidity: event SetToken(address indexed token, uint256 strikePriceIncrement)
func (_PremiaOption *PremiaOptionFilterer) ParseSetToken(log types.Log) (*PremiaOptionSetToken, error) {
	event := new(PremiaOptionSetToken)
	if err := _PremiaOption.contract.UnpackLog(event, "SetToken", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the PremiaOption contract.
type PremiaOptionTransferBatchIterator struct {
	Event *PremiaOptionTransferBatch // Event containing the contract specifics and raw log

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
func (it *PremiaOptionTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionTransferBatch)
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
		it.Event = new(PremiaOptionTransferBatch)
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
func (it *PremiaOptionTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionTransferBatch represents a TransferBatch event raised by the PremiaOption contract.
type PremiaOptionTransferBatch struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Ids      []*big.Int
	Values   []*big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PremiaOption *PremiaOptionFilterer) FilterTransferBatch(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PremiaOptionTransferBatchIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionTransferBatchIterator{contract: _PremiaOption.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PremiaOption *PremiaOptionFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *PremiaOptionTransferBatch, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "TransferBatch", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionTransferBatch)
				if err := _PremiaOption.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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

// ParseTransferBatch is a log parse operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed operator, address indexed from, address indexed to, uint256[] ids, uint256[] values)
func (_PremiaOption *PremiaOptionFilterer) ParseTransferBatch(log types.Log) (*PremiaOptionTransferBatch, error) {
	event := new(PremiaOptionTransferBatch)
	if err := _PremiaOption.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionTransferSingleIterator is returned from FilterTransferSingle and is used to iterate over the raw logs and unpacked data for TransferSingle events raised by the PremiaOption contract.
type PremiaOptionTransferSingleIterator struct {
	Event *PremiaOptionTransferSingle // Event containing the contract specifics and raw log

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
func (it *PremiaOptionTransferSingleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionTransferSingle)
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
		it.Event = new(PremiaOptionTransferSingle)
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
func (it *PremiaOptionTransferSingleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionTransferSingleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionTransferSingle represents a TransferSingle event raised by the PremiaOption contract.
type PremiaOptionTransferSingle struct {
	Operator common.Address
	From     common.Address
	To       common.Address
	Id       *big.Int
	Value    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferSingle is a free log retrieval operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PremiaOption *PremiaOptionFilterer) FilterTransferSingle(opts *bind.FilterOpts, operator []common.Address, from []common.Address, to []common.Address) (*PremiaOptionTransferSingleIterator, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionTransferSingleIterator{contract: _PremiaOption.contract, event: "TransferSingle", logs: logs, sub: sub}, nil
}

// WatchTransferSingle is a free log subscription operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PremiaOption *PremiaOptionFilterer) WatchTransferSingle(opts *bind.WatchOpts, sink chan<- *PremiaOptionTransferSingle, operator []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "TransferSingle", operatorRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionTransferSingle)
				if err := _PremiaOption.contract.UnpackLog(event, "TransferSingle", log); err != nil {
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

// ParseTransferSingle is a log parse operation binding the contract event 0xc3d58168c5ae7397731d063d5bbf3d657854427343f4c083240f7aacaa2d0f62.
//
// Solidity: event TransferSingle(address indexed operator, address indexed from, address indexed to, uint256 id, uint256 value)
func (_PremiaOption *PremiaOptionFilterer) ParseTransferSingle(log types.Log) (*PremiaOptionTransferSingle, error) {
	event := new(PremiaOptionTransferSingle)
	if err := _PremiaOption.contract.UnpackLog(event, "TransferSingle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionURIIterator is returned from FilterURI and is used to iterate over the raw logs and unpacked data for URI events raised by the PremiaOption contract.
type PremiaOptionURIIterator struct {
	Event *PremiaOptionURI // Event containing the contract specifics and raw log

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
func (it *PremiaOptionURIIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionURI)
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
		it.Event = new(PremiaOptionURI)
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
func (it *PremiaOptionURIIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionURIIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionURI represents a URI event raised by the PremiaOption contract.
type PremiaOptionURI struct {
	Value string
	Id    *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterURI is a free log retrieval operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PremiaOption *PremiaOptionFilterer) FilterURI(opts *bind.FilterOpts, id []*big.Int) (*PremiaOptionURIIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionURIIterator{contract: _PremiaOption.contract, event: "URI", logs: logs, sub: sub}, nil
}

// WatchURI is a free log subscription operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PremiaOption *PremiaOptionFilterer) WatchURI(opts *bind.WatchOpts, sink chan<- *PremiaOptionURI, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "URI", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionURI)
				if err := _PremiaOption.contract.UnpackLog(event, "URI", log); err != nil {
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

// ParseURI is a log parse operation binding the contract event 0x6bb7ff708619ba0610cba295a58592e0451dee2622938c8755667688daf3529b.
//
// Solidity: event URI(string value, uint256 indexed id)
func (_PremiaOption *PremiaOptionFilterer) ParseURI(log types.Log) (*PremiaOptionURI, error) {
	event := new(PremiaOptionURI)
	if err := _PremiaOption.contract.UnpackLog(event, "URI", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaOptionWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the PremiaOption contract.
type PremiaOptionWithdrawIterator struct {
	Event *PremiaOptionWithdraw // Event containing the contract specifics and raw log

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
func (it *PremiaOptionWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaOptionWithdraw)
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
		it.Event = new(PremiaOptionWithdraw)
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
func (it *PremiaOptionWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaOptionWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaOptionWithdraw represents a Withdraw event raised by the PremiaOption contract.
type PremiaOptionWithdraw struct {
	User     common.Address
	OptionId *big.Int
	Token    common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x457f950b75085c30ff780acd57bde642ff1316cc4aad9f286af2c1ffc4163a78.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) FilterWithdraw(opts *bind.FilterOpts, user []common.Address, optionId []*big.Int, token []common.Address) (*PremiaOptionWithdrawIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.FilterLogs(opts, "Withdraw", userRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return &PremiaOptionWithdrawIterator{contract: _PremiaOption.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x457f950b75085c30ff780acd57bde642ff1316cc4aad9f286af2c1ffc4163a78.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *PremiaOptionWithdraw, user []common.Address, optionId []*big.Int, token []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var optionIdRule []interface{}
	for _, optionIdItem := range optionId {
		optionIdRule = append(optionIdRule, optionIdItem)
	}
	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}

	logs, sub, err := _PremiaOption.contract.WatchLogs(opts, "Withdraw", userRule, optionIdRule, tokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaOptionWithdraw)
				if err := _PremiaOption.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x457f950b75085c30ff780acd57bde642ff1316cc4aad9f286af2c1ffc4163a78.
//
// Solidity: event Withdraw(address indexed user, uint256 indexed optionId, address indexed token, uint256 amount)
func (_PremiaOption *PremiaOptionFilterer) ParseWithdraw(log types.Log) (*PremiaOptionWithdraw, error) {
	event := new(PremiaOptionWithdraw)
	if err := _PremiaOption.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
