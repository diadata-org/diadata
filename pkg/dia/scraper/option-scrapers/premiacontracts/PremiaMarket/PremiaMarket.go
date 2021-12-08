// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package PremiaMarket

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

// IPremiaOptionOptionWriteArgs is an auto generated low-level Go binding around an user-defined struct.
type IPremiaOptionOptionWriteArgs struct {
	Token       common.Address
	Amount      *big.Int
	StrikePrice *big.Int
	Expiration  *big.Int
	IsCall      bool
}

// PremiaMarketOption is an auto generated low-level Go binding around an user-defined struct.
type PremiaMarketOption struct {
	Token       common.Address
	Expiration  *big.Int
	StrikePrice *big.Int
	IsCall      bool
}

// PremiaMarketOrder is an auto generated low-level Go binding around an user-defined struct.
type PremiaMarketOrder struct {
	Maker            common.Address
	Side             uint8
	IsDelayedWriting bool
	OptionContract   common.Address
	OptionId         *big.Int
	PaymentToken     common.Address
	PricePerUnit     *big.Int
	ExpirationTime   *big.Int
	Salt             *big.Int
	Decimals         uint8
}

// PremiaMarketABI is the input ABI used to generate the binding from.
const PremiaMarketABI = "[{\"inputs\":[{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"_uPremia\",\"type\":\"address\"},{\"internalType\":\"contractIFeeCalculator\",\"name\":\"_feeCalculator\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"},{\"internalType\":\"contractIPremiaReferral\",\"name\":\"_premiaReferral\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"OrderCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"}],\"name\":\"OrderFilled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedOptionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"}],\"name\":\"addWhitelistedPaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"amounts\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"}],\"name\":\"areOrdersValid\",\"outputs\":[{\"internalType\":\"bool[]\",\"name\":\"\",\"type\":\"bool[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimUPremia\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"createOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orderCandidates\",\"type\":\"tuple[]\"},{\"internalType\":\"bool\",\"name\":\"_writeOnBuyFill\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"createOrderAndTryToFill\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"internalType\":\"structPremiaMarket.Option\",\"name\":\"_option\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"createOrderForNewOption\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"_amounts\",\"type\":\"uint256[]\"}],\"name\":\"createOrders\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeCalculator\",\"outputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"_amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_writeOnBuyFill\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"fillOrder\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"_maxAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"_writeOnBuyFill\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"fillOrders\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"_orderIds\",\"type\":\"bytes32[]\"}],\"name\":\"getAmountsBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"getOrderHash\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order[]\",\"name\":\"_orders\",\"type\":\"tuple[]\"}],\"name\":\"getOrderHashBatch\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedOptionContracts\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getWhitelistedPaymentTokens\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isDelayedWritingEnabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"}],\"name\":\"isOrderValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"paymentTokenDecimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"premiaReferral\",\"outputs\":[{\"internalType\":\"contractIPremiaReferral\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedOptionContracts\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"_addr\",\"type\":\"address[]\"}],\"name\":\"removeWhitelistedPaymentTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"_state\",\"type\":\"bool\"}],\"name\":\"setDelayedWritingEnabled\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFeeCalculator\",\"name\":\"_feeCalculator\",\"type\":\"address\"}],\"name\":\"setFeeCalculator\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"_uPremia\",\"type\":\"address\"}],\"name\":\"setPremiaUncutErc20\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"uPremia\",\"outputs\":[{\"internalType\":\"contractIPremiaUncutErc20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"uPremiaBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"strikePrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isCall\",\"type\":\"bool\"}],\"internalType\":\"structIPremiaOption.OptionWriteArgs\",\"name\":\"_option\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"internalType\":\"enumPremiaMarket.SaleSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bool\",\"name\":\"isDelayedWriting\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"optionContract\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"optionId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"pricePerUnit\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"decimals\",\"type\":\"uint8\"}],\"internalType\":\"structPremiaMarket.Order\",\"name\":\"_order\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"_referrer\",\"type\":\"address\"}],\"name\":\"writeAndCreateOrder\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PremiaMarket is an auto generated Go binding around an Ethereum contract.
type PremiaMarket struct {
	PremiaMarketCaller     // Read-only binding to the contract
	PremiaMarketTransactor // Write-only binding to the contract
	PremiaMarketFilterer   // Log filterer for contract events
}

// PremiaMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type PremiaMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PremiaMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PremiaMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PremiaMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PremiaMarketSession struct {
	Contract     *PremiaMarket     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PremiaMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PremiaMarketCallerSession struct {
	Contract *PremiaMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PremiaMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PremiaMarketTransactorSession struct {
	Contract     *PremiaMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PremiaMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type PremiaMarketRaw struct {
	Contract *PremiaMarket // Generic contract binding to access the raw methods on
}

// PremiaMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PremiaMarketCallerRaw struct {
	Contract *PremiaMarketCaller // Generic read-only contract binding to access the raw methods on
}

// PremiaMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PremiaMarketTransactorRaw struct {
	Contract *PremiaMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPremiaMarket creates a new instance of PremiaMarket, bound to a specific deployed contract.
func NewPremiaMarket(address common.Address, backend bind.ContractBackend) (*PremiaMarket, error) {
	contract, err := bindPremiaMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PremiaMarket{PremiaMarketCaller: PremiaMarketCaller{contract: contract}, PremiaMarketTransactor: PremiaMarketTransactor{contract: contract}, PremiaMarketFilterer: PremiaMarketFilterer{contract: contract}}, nil
}

// NewPremiaMarketCaller creates a new read-only instance of PremiaMarket, bound to a specific deployed contract.
func NewPremiaMarketCaller(address common.Address, caller bind.ContractCaller) (*PremiaMarketCaller, error) {
	contract, err := bindPremiaMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketCaller{contract: contract}, nil
}

// NewPremiaMarketTransactor creates a new write-only instance of PremiaMarket, bound to a specific deployed contract.
func NewPremiaMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*PremiaMarketTransactor, error) {
	contract, err := bindPremiaMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketTransactor{contract: contract}, nil
}

// NewPremiaMarketFilterer creates a new log filterer instance of PremiaMarket, bound to a specific deployed contract.
func NewPremiaMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*PremiaMarketFilterer, error) {
	contract, err := bindPremiaMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketFilterer{contract: contract}, nil
}

// bindPremiaMarket binds a generic wrapper to an already deployed contract.
func bindPremiaMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PremiaMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PremiaMarket *PremiaMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PremiaMarket.Contract.PremiaMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PremiaMarket *PremiaMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaMarket.Contract.PremiaMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PremiaMarket *PremiaMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PremiaMarket.Contract.PremiaMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PremiaMarket *PremiaMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PremiaMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PremiaMarket *PremiaMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PremiaMarket *PremiaMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PremiaMarket.Contract.contract.Transact(opts, method, params...)
}

// Amounts is a free data retrieval call binding the contract method 0xc8662d95.
//
// Solidity: function amounts(bytes32 ) view returns(uint256)
func (_PremiaMarket *PremiaMarketCaller) Amounts(opts *bind.CallOpts, arg0 [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "amounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Amounts is a free data retrieval call binding the contract method 0xc8662d95.
//
// Solidity: function amounts(bytes32 ) view returns(uint256)
func (_PremiaMarket *PremiaMarketSession) Amounts(arg0 [32]byte) (*big.Int, error) {
	return _PremiaMarket.Contract.Amounts(&_PremiaMarket.CallOpts, arg0)
}

// Amounts is a free data retrieval call binding the contract method 0xc8662d95.
//
// Solidity: function amounts(bytes32 ) view returns(uint256)
func (_PremiaMarket *PremiaMarketCallerSession) Amounts(arg0 [32]byte) (*big.Int, error) {
	return _PremiaMarket.Contract.Amounts(&_PremiaMarket.CallOpts, arg0)
}

// AreOrdersValid is a free data retrieval call binding the contract method 0x9e06c6d9.
//
// Solidity: function areOrdersValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) view returns(bool[])
func (_PremiaMarket *PremiaMarketCaller) AreOrdersValid(opts *bind.CallOpts, _orders []PremiaMarketOrder) ([]bool, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "areOrdersValid", _orders)

	if err != nil {
		return *new([]bool), err
	}

	out0 := *abi.ConvertType(out[0], new([]bool)).(*[]bool)

	return out0, err

}

// AreOrdersValid is a free data retrieval call binding the contract method 0x9e06c6d9.
//
// Solidity: function areOrdersValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) view returns(bool[])
func (_PremiaMarket *PremiaMarketSession) AreOrdersValid(_orders []PremiaMarketOrder) ([]bool, error) {
	return _PremiaMarket.Contract.AreOrdersValid(&_PremiaMarket.CallOpts, _orders)
}

// AreOrdersValid is a free data retrieval call binding the contract method 0x9e06c6d9.
//
// Solidity: function areOrdersValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) view returns(bool[])
func (_PremiaMarket *PremiaMarketCallerSession) AreOrdersValid(_orders []PremiaMarketOrder) ([]bool, error) {
	return _PremiaMarket.Contract.AreOrdersValid(&_PremiaMarket.CallOpts, _orders)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaMarket *PremiaMarketCaller) FeeCalculator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "feeCalculator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaMarket *PremiaMarketSession) FeeCalculator() (common.Address, error) {
	return _PremiaMarket.Contract.FeeCalculator(&_PremiaMarket.CallOpts)
}

// FeeCalculator is a free data retrieval call binding the contract method 0xb00eb9fe.
//
// Solidity: function feeCalculator() view returns(address)
func (_PremiaMarket *PremiaMarketCallerSession) FeeCalculator() (common.Address, error) {
	return _PremiaMarket.Contract.FeeCalculator(&_PremiaMarket.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaMarket *PremiaMarketCaller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaMarket *PremiaMarketSession) FeeRecipient() (common.Address, error) {
	return _PremiaMarket.Contract.FeeRecipient(&_PremiaMarket.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_PremiaMarket *PremiaMarketCallerSession) FeeRecipient() (common.Address, error) {
	return _PremiaMarket.Contract.FeeRecipient(&_PremiaMarket.CallOpts)
}

// GetAmountsBatch is a free data retrieval call binding the contract method 0x1ef6ab30.
//
// Solidity: function getAmountsBatch(bytes32[] _orderIds) view returns(uint256[])
func (_PremiaMarket *PremiaMarketCaller) GetAmountsBatch(opts *bind.CallOpts, _orderIds [][32]byte) ([]*big.Int, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "getAmountsBatch", _orderIds)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAmountsBatch is a free data retrieval call binding the contract method 0x1ef6ab30.
//
// Solidity: function getAmountsBatch(bytes32[] _orderIds) view returns(uint256[])
func (_PremiaMarket *PremiaMarketSession) GetAmountsBatch(_orderIds [][32]byte) ([]*big.Int, error) {
	return _PremiaMarket.Contract.GetAmountsBatch(&_PremiaMarket.CallOpts, _orderIds)
}

// GetAmountsBatch is a free data retrieval call binding the contract method 0x1ef6ab30.
//
// Solidity: function getAmountsBatch(bytes32[] _orderIds) view returns(uint256[])
func (_PremiaMarket *PremiaMarketCallerSession) GetAmountsBatch(_orderIds [][32]byte) ([]*big.Int, error) {
	return _PremiaMarket.Contract.GetAmountsBatch(&_PremiaMarket.CallOpts, _orderIds)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xfb3f9661.
//
// Solidity: function getOrderHash((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) pure returns(bytes32)
func (_PremiaMarket *PremiaMarketCaller) GetOrderHash(opts *bind.CallOpts, _order PremiaMarketOrder) ([32]byte, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "getOrderHash", _order)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetOrderHash is a free data retrieval call binding the contract method 0xfb3f9661.
//
// Solidity: function getOrderHash((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) pure returns(bytes32)
func (_PremiaMarket *PremiaMarketSession) GetOrderHash(_order PremiaMarketOrder) ([32]byte, error) {
	return _PremiaMarket.Contract.GetOrderHash(&_PremiaMarket.CallOpts, _order)
}

// GetOrderHash is a free data retrieval call binding the contract method 0xfb3f9661.
//
// Solidity: function getOrderHash((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) pure returns(bytes32)
func (_PremiaMarket *PremiaMarketCallerSession) GetOrderHash(_order PremiaMarketOrder) ([32]byte, error) {
	return _PremiaMarket.Contract.GetOrderHash(&_PremiaMarket.CallOpts, _order)
}

// GetOrderHashBatch is a free data retrieval call binding the contract method 0x783fb8ef.
//
// Solidity: function getOrderHashBatch((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) pure returns(bytes32[])
func (_PremiaMarket *PremiaMarketCaller) GetOrderHashBatch(opts *bind.CallOpts, _orders []PremiaMarketOrder) ([][32]byte, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "getOrderHashBatch", _orders)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// GetOrderHashBatch is a free data retrieval call binding the contract method 0x783fb8ef.
//
// Solidity: function getOrderHashBatch((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) pure returns(bytes32[])
func (_PremiaMarket *PremiaMarketSession) GetOrderHashBatch(_orders []PremiaMarketOrder) ([][32]byte, error) {
	return _PremiaMarket.Contract.GetOrderHashBatch(&_PremiaMarket.CallOpts, _orders)
}

// GetOrderHashBatch is a free data retrieval call binding the contract method 0x783fb8ef.
//
// Solidity: function getOrderHashBatch((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) pure returns(bytes32[])
func (_PremiaMarket *PremiaMarketCallerSession) GetOrderHashBatch(_orders []PremiaMarketOrder) ([][32]byte, error) {
	return _PremiaMarket.Contract.GetOrderHashBatch(&_PremiaMarket.CallOpts, _orders)
}

// GetWhitelistedOptionContracts is a free data retrieval call binding the contract method 0x3baa38ef.
//
// Solidity: function getWhitelistedOptionContracts() view returns(address[])
func (_PremiaMarket *PremiaMarketCaller) GetWhitelistedOptionContracts(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "getWhitelistedOptionContracts")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedOptionContracts is a free data retrieval call binding the contract method 0x3baa38ef.
//
// Solidity: function getWhitelistedOptionContracts() view returns(address[])
func (_PremiaMarket *PremiaMarketSession) GetWhitelistedOptionContracts() ([]common.Address, error) {
	return _PremiaMarket.Contract.GetWhitelistedOptionContracts(&_PremiaMarket.CallOpts)
}

// GetWhitelistedOptionContracts is a free data retrieval call binding the contract method 0x3baa38ef.
//
// Solidity: function getWhitelistedOptionContracts() view returns(address[])
func (_PremiaMarket *PremiaMarketCallerSession) GetWhitelistedOptionContracts() ([]common.Address, error) {
	return _PremiaMarket.Contract.GetWhitelistedOptionContracts(&_PremiaMarket.CallOpts)
}

// GetWhitelistedPaymentTokens is a free data retrieval call binding the contract method 0xe6f073e0.
//
// Solidity: function getWhitelistedPaymentTokens() view returns(address[])
func (_PremiaMarket *PremiaMarketCaller) GetWhitelistedPaymentTokens(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "getWhitelistedPaymentTokens")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetWhitelistedPaymentTokens is a free data retrieval call binding the contract method 0xe6f073e0.
//
// Solidity: function getWhitelistedPaymentTokens() view returns(address[])
func (_PremiaMarket *PremiaMarketSession) GetWhitelistedPaymentTokens() ([]common.Address, error) {
	return _PremiaMarket.Contract.GetWhitelistedPaymentTokens(&_PremiaMarket.CallOpts)
}

// GetWhitelistedPaymentTokens is a free data retrieval call binding the contract method 0xe6f073e0.
//
// Solidity: function getWhitelistedPaymentTokens() view returns(address[])
func (_PremiaMarket *PremiaMarketCallerSession) GetWhitelistedPaymentTokens() ([]common.Address, error) {
	return _PremiaMarket.Contract.GetWhitelistedPaymentTokens(&_PremiaMarket.CallOpts)
}

// IsDelayedWritingEnabled is a free data retrieval call binding the contract method 0x4c878a41.
//
// Solidity: function isDelayedWritingEnabled() view returns(bool)
func (_PremiaMarket *PremiaMarketCaller) IsDelayedWritingEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "isDelayedWritingEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsDelayedWritingEnabled is a free data retrieval call binding the contract method 0x4c878a41.
//
// Solidity: function isDelayedWritingEnabled() view returns(bool)
func (_PremiaMarket *PremiaMarketSession) IsDelayedWritingEnabled() (bool, error) {
	return _PremiaMarket.Contract.IsDelayedWritingEnabled(&_PremiaMarket.CallOpts)
}

// IsDelayedWritingEnabled is a free data retrieval call binding the contract method 0x4c878a41.
//
// Solidity: function isDelayedWritingEnabled() view returns(bool)
func (_PremiaMarket *PremiaMarketCallerSession) IsDelayedWritingEnabled() (bool, error) {
	return _PremiaMarket.Contract.IsDelayedWritingEnabled(&_PremiaMarket.CallOpts)
}

// IsOrderValid is a free data retrieval call binding the contract method 0x4969f871.
//
// Solidity: function isOrderValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) view returns(bool)
func (_PremiaMarket *PremiaMarketCaller) IsOrderValid(opts *bind.CallOpts, _order PremiaMarketOrder) (bool, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "isOrderValid", _order)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOrderValid is a free data retrieval call binding the contract method 0x4969f871.
//
// Solidity: function isOrderValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) view returns(bool)
func (_PremiaMarket *PremiaMarketSession) IsOrderValid(_order PremiaMarketOrder) (bool, error) {
	return _PremiaMarket.Contract.IsOrderValid(&_PremiaMarket.CallOpts, _order)
}

// IsOrderValid is a free data retrieval call binding the contract method 0x4969f871.
//
// Solidity: function isOrderValid((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) view returns(bool)
func (_PremiaMarket *PremiaMarketCallerSession) IsOrderValid(_order PremiaMarketOrder) (bool, error) {
	return _PremiaMarket.Contract.IsOrderValid(&_PremiaMarket.CallOpts, _order)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaMarket *PremiaMarketCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaMarket *PremiaMarketSession) Owner() (common.Address, error) {
	return _PremiaMarket.Contract.Owner(&_PremiaMarket.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PremiaMarket *PremiaMarketCallerSession) Owner() (common.Address, error) {
	return _PremiaMarket.Contract.Owner(&_PremiaMarket.CallOpts)
}

// PaymentTokenDecimals is a free data retrieval call binding the contract method 0xd481515a.
//
// Solidity: function paymentTokenDecimals(address ) view returns(uint8)
func (_PremiaMarket *PremiaMarketCaller) PaymentTokenDecimals(opts *bind.CallOpts, arg0 common.Address) (uint8, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "paymentTokenDecimals", arg0)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// PaymentTokenDecimals is a free data retrieval call binding the contract method 0xd481515a.
//
// Solidity: function paymentTokenDecimals(address ) view returns(uint8)
func (_PremiaMarket *PremiaMarketSession) PaymentTokenDecimals(arg0 common.Address) (uint8, error) {
	return _PremiaMarket.Contract.PaymentTokenDecimals(&_PremiaMarket.CallOpts, arg0)
}

// PaymentTokenDecimals is a free data retrieval call binding the contract method 0xd481515a.
//
// Solidity: function paymentTokenDecimals(address ) view returns(uint8)
func (_PremiaMarket *PremiaMarketCallerSession) PaymentTokenDecimals(arg0 common.Address) (uint8, error) {
	return _PremiaMarket.Contract.PaymentTokenDecimals(&_PremiaMarket.CallOpts, arg0)
}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaMarket *PremiaMarketCaller) PremiaReferral(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "premiaReferral")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaMarket *PremiaMarketSession) PremiaReferral() (common.Address, error) {
	return _PremiaMarket.Contract.PremiaReferral(&_PremiaMarket.CallOpts)
}

// PremiaReferral is a free data retrieval call binding the contract method 0x4ebd2287.
//
// Solidity: function premiaReferral() view returns(address)
func (_PremiaMarket *PremiaMarketCallerSession) PremiaReferral() (common.Address, error) {
	return _PremiaMarket.Contract.PremiaReferral(&_PremiaMarket.CallOpts)
}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaMarket *PremiaMarketCaller) UPremia(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "uPremia")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaMarket *PremiaMarketSession) UPremia() (common.Address, error) {
	return _PremiaMarket.Contract.UPremia(&_PremiaMarket.CallOpts)
}

// UPremia is a free data retrieval call binding the contract method 0xb9a8e9d3.
//
// Solidity: function uPremia() view returns(address)
func (_PremiaMarket *PremiaMarketCallerSession) UPremia() (common.Address, error) {
	return _PremiaMarket.Contract.UPremia(&_PremiaMarket.CallOpts)
}

// UPremiaBalance is a free data retrieval call binding the contract method 0x929945d9.
//
// Solidity: function uPremiaBalance(address ) view returns(uint256)
func (_PremiaMarket *PremiaMarketCaller) UPremiaBalance(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _PremiaMarket.contract.Call(opts, &out, "uPremiaBalance", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// UPremiaBalance is a free data retrieval call binding the contract method 0x929945d9.
//
// Solidity: function uPremiaBalance(address ) view returns(uint256)
func (_PremiaMarket *PremiaMarketSession) UPremiaBalance(arg0 common.Address) (*big.Int, error) {
	return _PremiaMarket.Contract.UPremiaBalance(&_PremiaMarket.CallOpts, arg0)
}

// UPremiaBalance is a free data retrieval call binding the contract method 0x929945d9.
//
// Solidity: function uPremiaBalance(address ) view returns(uint256)
func (_PremiaMarket *PremiaMarketCallerSession) UPremiaBalance(arg0 common.Address) (*big.Int, error) {
	return _PremiaMarket.Contract.UPremiaBalance(&_PremiaMarket.CallOpts, arg0)
}

// AddWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0x8ea19f9b.
//
// Solidity: function addWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactor) AddWhitelistedOptionContracts(opts *bind.TransactOpts, _addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "addWhitelistedOptionContracts", _addr)
}

// AddWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0x8ea19f9b.
//
// Solidity: function addWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketSession) AddWhitelistedOptionContracts(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.AddWhitelistedOptionContracts(&_PremiaMarket.TransactOpts, _addr)
}

// AddWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0x8ea19f9b.
//
// Solidity: function addWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) AddWhitelistedOptionContracts(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.AddWhitelistedOptionContracts(&_PremiaMarket.TransactOpts, _addr)
}

// AddWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0xf4084cee.
//
// Solidity: function addWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactor) AddWhitelistedPaymentTokens(opts *bind.TransactOpts, _addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "addWhitelistedPaymentTokens", _addr)
}

// AddWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0xf4084cee.
//
// Solidity: function addWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketSession) AddWhitelistedPaymentTokens(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.AddWhitelistedPaymentTokens(&_PremiaMarket.TransactOpts, _addr)
}

// AddWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0xf4084cee.
//
// Solidity: function addWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) AddWhitelistedPaymentTokens(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.AddWhitelistedPaymentTokens(&_PremiaMarket.TransactOpts, _addr)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x3eeba766.
//
// Solidity: function cancelOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) returns()
func (_PremiaMarket *PremiaMarketTransactor) CancelOrder(opts *bind.TransactOpts, _order PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "cancelOrder", _order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x3eeba766.
//
// Solidity: function cancelOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) returns()
func (_PremiaMarket *PremiaMarketSession) CancelOrder(_order PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CancelOrder(&_PremiaMarket.TransactOpts, _order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0x3eeba766.
//
// Solidity: function cancelOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) CancelOrder(_order PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CancelOrder(&_PremiaMarket.TransactOpts, _order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x306de2df.
//
// Solidity: function cancelOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) returns()
func (_PremiaMarket *PremiaMarketTransactor) CancelOrders(opts *bind.TransactOpts, _orders []PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "cancelOrders", _orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x306de2df.
//
// Solidity: function cancelOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) returns()
func (_PremiaMarket *PremiaMarketSession) CancelOrders(_orders []PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CancelOrders(&_PremiaMarket.TransactOpts, _orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0x306de2df.
//
// Solidity: function cancelOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) CancelOrders(_orders []PremiaMarketOrder) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CancelOrders(&_PremiaMarket.TransactOpts, _orders)
}

// ClaimUPremia is a paid mutator transaction binding the contract method 0x6c5c662e.
//
// Solidity: function claimUPremia() returns()
func (_PremiaMarket *PremiaMarketTransactor) ClaimUPremia(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "claimUPremia")
}

// ClaimUPremia is a paid mutator transaction binding the contract method 0x6c5c662e.
//
// Solidity: function claimUPremia() returns()
func (_PremiaMarket *PremiaMarketSession) ClaimUPremia() (*types.Transaction, error) {
	return _PremiaMarket.Contract.ClaimUPremia(&_PremiaMarket.TransactOpts)
}

// ClaimUPremia is a paid mutator transaction binding the contract method 0x6c5c662e.
//
// Solidity: function claimUPremia() returns()
func (_PremiaMarket *PremiaMarketTransactorSession) ClaimUPremia() (*types.Transaction, error) {
	return _PremiaMarket.Contract.ClaimUPremia(&_PremiaMarket.TransactOpts)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x11e4c3bc.
//
// Solidity: function createOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactor) CreateOrder(opts *bind.TransactOpts, _order PremiaMarketOrder, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "createOrder", _order, _amount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x11e4c3bc.
//
// Solidity: function createOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount) returns(bytes32)
func (_PremiaMarket *PremiaMarketSession) CreateOrder(_order PremiaMarketOrder, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrder(&_PremiaMarket.TransactOpts, _order, _amount)
}

// CreateOrder is a paid mutator transaction binding the contract method 0x11e4c3bc.
//
// Solidity: function createOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactorSession) CreateOrder(_order PremiaMarketOrder, _amount *big.Int) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrder(&_PremiaMarket.TransactOpts, _order, _amount)
}

// CreateOrderAndTryToFill is a paid mutator transaction binding the contract method 0x4f66efd9.
//
// Solidity: function createOrderAndTryToFill((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orderCandidates, bool _writeOnBuyFill, address _referrer) returns()
func (_PremiaMarket *PremiaMarketTransactor) CreateOrderAndTryToFill(opts *bind.TransactOpts, _order PremiaMarketOrder, _amount *big.Int, _orderCandidates []PremiaMarketOrder, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "createOrderAndTryToFill", _order, _amount, _orderCandidates, _writeOnBuyFill, _referrer)
}

// CreateOrderAndTryToFill is a paid mutator transaction binding the contract method 0x4f66efd9.
//
// Solidity: function createOrderAndTryToFill((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orderCandidates, bool _writeOnBuyFill, address _referrer) returns()
func (_PremiaMarket *PremiaMarketSession) CreateOrderAndTryToFill(_order PremiaMarketOrder, _amount *big.Int, _orderCandidates []PremiaMarketOrder, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrderAndTryToFill(&_PremiaMarket.TransactOpts, _order, _amount, _orderCandidates, _writeOnBuyFill, _referrer)
}

// CreateOrderAndTryToFill is a paid mutator transaction binding the contract method 0x4f66efd9.
//
// Solidity: function createOrderAndTryToFill((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orderCandidates, bool _writeOnBuyFill, address _referrer) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) CreateOrderAndTryToFill(_order PremiaMarketOrder, _amount *big.Int, _orderCandidates []PremiaMarketOrder, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrderAndTryToFill(&_PremiaMarket.TransactOpts, _order, _amount, _orderCandidates, _writeOnBuyFill, _referrer)
}

// CreateOrderForNewOption is a paid mutator transaction binding the contract method 0x947f75d0.
//
// Solidity: function createOrderForNewOption((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint256,uint256,bool) _option, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactor) CreateOrderForNewOption(opts *bind.TransactOpts, _order PremiaMarketOrder, _amount *big.Int, _option PremiaMarketOption, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "createOrderForNewOption", _order, _amount, _option, _referrer)
}

// CreateOrderForNewOption is a paid mutator transaction binding the contract method 0x947f75d0.
//
// Solidity: function createOrderForNewOption((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint256,uint256,bool) _option, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketSession) CreateOrderForNewOption(_order PremiaMarketOrder, _amount *big.Int, _option PremiaMarketOption, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrderForNewOption(&_PremiaMarket.TransactOpts, _order, _amount, _option, _referrer)
}

// CreateOrderForNewOption is a paid mutator transaction binding the contract method 0x947f75d0.
//
// Solidity: function createOrderForNewOption((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, (address,uint256,uint256,bool) _option, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactorSession) CreateOrderForNewOption(_order PremiaMarketOrder, _amount *big.Int, _option PremiaMarketOption, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrderForNewOption(&_PremiaMarket.TransactOpts, _order, _amount, _option, _referrer)
}

// CreateOrders is a paid mutator transaction binding the contract method 0x6fc7aafc.
//
// Solidity: function createOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256[] _amounts) returns(bytes32[])
func (_PremiaMarket *PremiaMarketTransactor) CreateOrders(opts *bind.TransactOpts, _orders []PremiaMarketOrder, _amounts []*big.Int) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "createOrders", _orders, _amounts)
}

// CreateOrders is a paid mutator transaction binding the contract method 0x6fc7aafc.
//
// Solidity: function createOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256[] _amounts) returns(bytes32[])
func (_PremiaMarket *PremiaMarketSession) CreateOrders(_orders []PremiaMarketOrder, _amounts []*big.Int) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrders(&_PremiaMarket.TransactOpts, _orders, _amounts)
}

// CreateOrders is a paid mutator transaction binding the contract method 0x6fc7aafc.
//
// Solidity: function createOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256[] _amounts) returns(bytes32[])
func (_PremiaMarket *PremiaMarketTransactorSession) CreateOrders(_orders []PremiaMarketOrder, _amounts []*big.Int) (*types.Transaction, error) {
	return _PremiaMarket.Contract.CreateOrders(&_PremiaMarket.TransactOpts, _orders, _amounts)
}

// FillOrder is a paid mutator transaction binding the contract method 0x668dd276.
//
// Solidity: function fillOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketTransactor) FillOrder(opts *bind.TransactOpts, _order PremiaMarketOrder, _amount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "fillOrder", _order, _amount, _writeOnBuyFill, _referrer)
}

// FillOrder is a paid mutator transaction binding the contract method 0x668dd276.
//
// Solidity: function fillOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketSession) FillOrder(_order PremiaMarketOrder, _amount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.FillOrder(&_PremiaMarket.TransactOpts, _order, _amount, _writeOnBuyFill, _referrer)
}

// FillOrder is a paid mutator transaction binding the contract method 0x668dd276.
//
// Solidity: function fillOrder((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, uint256 _amount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketTransactorSession) FillOrder(_order PremiaMarketOrder, _amount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.FillOrder(&_PremiaMarket.TransactOpts, _order, _amount, _writeOnBuyFill, _referrer)
}

// FillOrders is a paid mutator transaction binding the contract method 0x097d29a7.
//
// Solidity: function fillOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256 _maxAmount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketTransactor) FillOrders(opts *bind.TransactOpts, _orders []PremiaMarketOrder, _maxAmount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "fillOrders", _orders, _maxAmount, _writeOnBuyFill, _referrer)
}

// FillOrders is a paid mutator transaction binding the contract method 0x097d29a7.
//
// Solidity: function fillOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256 _maxAmount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketSession) FillOrders(_orders []PremiaMarketOrder, _maxAmount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.FillOrders(&_PremiaMarket.TransactOpts, _orders, _maxAmount, _writeOnBuyFill, _referrer)
}

// FillOrders is a paid mutator transaction binding the contract method 0x097d29a7.
//
// Solidity: function fillOrders((address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8)[] _orders, uint256 _maxAmount, bool _writeOnBuyFill, address _referrer) returns(uint256)
func (_PremiaMarket *PremiaMarketTransactorSession) FillOrders(_orders []PremiaMarketOrder, _maxAmount *big.Int, _writeOnBuyFill bool, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.FillOrders(&_PremiaMarket.TransactOpts, _orders, _maxAmount, _writeOnBuyFill, _referrer)
}

// RemoveWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0xd906ba58.
//
// Solidity: function removeWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactor) RemoveWhitelistedOptionContracts(opts *bind.TransactOpts, _addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "removeWhitelistedOptionContracts", _addr)
}

// RemoveWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0xd906ba58.
//
// Solidity: function removeWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketSession) RemoveWhitelistedOptionContracts(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.RemoveWhitelistedOptionContracts(&_PremiaMarket.TransactOpts, _addr)
}

// RemoveWhitelistedOptionContracts is a paid mutator transaction binding the contract method 0xd906ba58.
//
// Solidity: function removeWhitelistedOptionContracts(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) RemoveWhitelistedOptionContracts(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.RemoveWhitelistedOptionContracts(&_PremiaMarket.TransactOpts, _addr)
}

// RemoveWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0x79c19c83.
//
// Solidity: function removeWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactor) RemoveWhitelistedPaymentTokens(opts *bind.TransactOpts, _addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "removeWhitelistedPaymentTokens", _addr)
}

// RemoveWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0x79c19c83.
//
// Solidity: function removeWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketSession) RemoveWhitelistedPaymentTokens(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.RemoveWhitelistedPaymentTokens(&_PremiaMarket.TransactOpts, _addr)
}

// RemoveWhitelistedPaymentTokens is a paid mutator transaction binding the contract method 0x79c19c83.
//
// Solidity: function removeWhitelistedPaymentTokens(address[] _addr) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) RemoveWhitelistedPaymentTokens(_addr []common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.RemoveWhitelistedPaymentTokens(&_PremiaMarket.TransactOpts, _addr)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaMarket *PremiaMarketTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaMarket *PremiaMarketSession) RenounceOwnership() (*types.Transaction, error) {
	return _PremiaMarket.Contract.RenounceOwnership(&_PremiaMarket.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PremiaMarket *PremiaMarketTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PremiaMarket.Contract.RenounceOwnership(&_PremiaMarket.TransactOpts)
}

// SetDelayedWritingEnabled is a paid mutator transaction binding the contract method 0x9a52b73e.
//
// Solidity: function setDelayedWritingEnabled(bool _state) returns()
func (_PremiaMarket *PremiaMarketTransactor) SetDelayedWritingEnabled(opts *bind.TransactOpts, _state bool) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "setDelayedWritingEnabled", _state)
}

// SetDelayedWritingEnabled is a paid mutator transaction binding the contract method 0x9a52b73e.
//
// Solidity: function setDelayedWritingEnabled(bool _state) returns()
func (_PremiaMarket *PremiaMarketSession) SetDelayedWritingEnabled(_state bool) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetDelayedWritingEnabled(&_PremiaMarket.TransactOpts, _state)
}

// SetDelayedWritingEnabled is a paid mutator transaction binding the contract method 0x9a52b73e.
//
// Solidity: function setDelayedWritingEnabled(bool _state) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) SetDelayedWritingEnabled(_state bool) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetDelayedWritingEnabled(&_PremiaMarket.TransactOpts, _state)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaMarket *PremiaMarketTransactor) SetFeeCalculator(opts *bind.TransactOpts, _feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "setFeeCalculator", _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaMarket *PremiaMarketSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetFeeCalculator(&_PremiaMarket.TransactOpts, _feeCalculator)
}

// SetFeeCalculator is a paid mutator transaction binding the contract method 0x8c66d04f.
//
// Solidity: function setFeeCalculator(address _feeCalculator) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) SetFeeCalculator(_feeCalculator common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetFeeCalculator(&_PremiaMarket.TransactOpts, _feeCalculator)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaMarket *PremiaMarketTransactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaMarket *PremiaMarketSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetFeeRecipient(&_PremiaMarket.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetFeeRecipient(&_PremiaMarket.TransactOpts, _feeRecipient)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaMarket *PremiaMarketTransactor) SetPremiaUncutErc20(opts *bind.TransactOpts, _uPremia common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "setPremiaUncutErc20", _uPremia)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaMarket *PremiaMarketSession) SetPremiaUncutErc20(_uPremia common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetPremiaUncutErc20(&_PremiaMarket.TransactOpts, _uPremia)
}

// SetPremiaUncutErc20 is a paid mutator transaction binding the contract method 0xc0b5f2cc.
//
// Solidity: function setPremiaUncutErc20(address _uPremia) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) SetPremiaUncutErc20(_uPremia common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.SetPremiaUncutErc20(&_PremiaMarket.TransactOpts, _uPremia)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaMarket *PremiaMarketTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaMarket *PremiaMarketSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.TransferOwnership(&_PremiaMarket.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PremiaMarket *PremiaMarketTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.TransferOwnership(&_PremiaMarket.TransactOpts, newOwner)
}

// WriteAndCreateOrder is a paid mutator transaction binding the contract method 0xf4738c13.
//
// Solidity: function writeAndCreateOrder((address,uint256,uint256,uint256,bool) _option, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactor) WriteAndCreateOrder(opts *bind.TransactOpts, _option IPremiaOptionOptionWriteArgs, _order PremiaMarketOrder, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.contract.Transact(opts, "writeAndCreateOrder", _option, _order, _referrer)
}

// WriteAndCreateOrder is a paid mutator transaction binding the contract method 0xf4738c13.
//
// Solidity: function writeAndCreateOrder((address,uint256,uint256,uint256,bool) _option, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketSession) WriteAndCreateOrder(_option IPremiaOptionOptionWriteArgs, _order PremiaMarketOrder, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.WriteAndCreateOrder(&_PremiaMarket.TransactOpts, _option, _order, _referrer)
}

// WriteAndCreateOrder is a paid mutator transaction binding the contract method 0xf4738c13.
//
// Solidity: function writeAndCreateOrder((address,uint256,uint256,uint256,bool) _option, (address,uint8,bool,address,uint256,address,uint256,uint256,uint256,uint8) _order, address _referrer) returns(bytes32)
func (_PremiaMarket *PremiaMarketTransactorSession) WriteAndCreateOrder(_option IPremiaOptionOptionWriteArgs, _order PremiaMarketOrder, _referrer common.Address) (*types.Transaction, error) {
	return _PremiaMarket.Contract.WriteAndCreateOrder(&_PremiaMarket.TransactOpts, _option, _order, _referrer)
}

// PremiaMarketOrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the PremiaMarket contract.
type PremiaMarketOrderCancelledIterator struct {
	Event *PremiaMarketOrderCancelled // Event containing the contract specifics and raw log

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
func (it *PremiaMarketOrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaMarketOrderCancelled)
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
		it.Event = new(PremiaMarketOrderCancelled)
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
func (it *PremiaMarketOrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaMarketOrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaMarketOrderCancelled represents a OrderCancelled event raised by the PremiaMarket contract.
type PremiaMarketOrderCancelled struct {
	Hash           [32]byte
	Maker          common.Address
	OptionContract common.Address
	PaymentToken   common.Address
	Amount         *big.Int
	PricePerUnit   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5d22ca4d7d4d28cd5651ddf3defbaab145d6146acb2b899eb0698144fd9d5e12.
//
// Solidity: event OrderCancelled(bytes32 indexed hash, address indexed maker, address indexed optionContract, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) FilterOrderCancelled(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, optionContract []common.Address) (*PremiaMarketOrderCancelledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.FilterLogs(opts, "OrderCancelled", hashRule, makerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketOrderCancelledIterator{contract: _PremiaMarket.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5d22ca4d7d4d28cd5651ddf3defbaab145d6146acb2b899eb0698144fd9d5e12.
//
// Solidity: event OrderCancelled(bytes32 indexed hash, address indexed maker, address indexed optionContract, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *PremiaMarketOrderCancelled, hash [][32]byte, maker []common.Address, optionContract []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.WatchLogs(opts, "OrderCancelled", hashRule, makerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaMarketOrderCancelled)
				if err := _PremiaMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5d22ca4d7d4d28cd5651ddf3defbaab145d6146acb2b899eb0698144fd9d5e12.
//
// Solidity: event OrderCancelled(bytes32 indexed hash, address indexed maker, address indexed optionContract, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) ParseOrderCancelled(log types.Log) (*PremiaMarketOrderCancelled, error) {
	event := new(PremiaMarketOrderCancelled)
	if err := _PremiaMarket.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaMarketOrderCreatedIterator is returned from FilterOrderCreated and is used to iterate over the raw logs and unpacked data for OrderCreated events raised by the PremiaMarket contract.
type PremiaMarketOrderCreatedIterator struct {
	Event *PremiaMarketOrderCreated // Event containing the contract specifics and raw log

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
func (it *PremiaMarketOrderCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaMarketOrderCreated)
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
		it.Event = new(PremiaMarketOrderCreated)
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
func (it *PremiaMarketOrderCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaMarketOrderCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaMarketOrderCreated represents a OrderCreated event raised by the PremiaMarket contract.
type PremiaMarketOrderCreated struct {
	Hash             [32]byte
	Maker            common.Address
	OptionContract   common.Address
	Side             uint8
	IsDelayedWriting bool
	OptionId         *big.Int
	PaymentToken     common.Address
	PricePerUnit     *big.Int
	ExpirationTime   *big.Int
	Salt             *big.Int
	Amount           *big.Int
	Decimals         uint8
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterOrderCreated is a free log retrieval operation binding the contract event 0x7b1bb3d3aaec5962def20e424565d7eee44dd6e74694eb87d96f5415ebd57dd3.
//
// Solidity: event OrderCreated(bytes32 indexed hash, address indexed maker, address indexed optionContract, uint8 side, bool isDelayedWriting, uint256 optionId, address paymentToken, uint256 pricePerUnit, uint256 expirationTime, uint256 salt, uint256 amount, uint8 decimals)
func (_PremiaMarket *PremiaMarketFilterer) FilterOrderCreated(opts *bind.FilterOpts, hash [][32]byte, maker []common.Address, optionContract []common.Address) (*PremiaMarketOrderCreatedIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.FilterLogs(opts, "OrderCreated", hashRule, makerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketOrderCreatedIterator{contract: _PremiaMarket.contract, event: "OrderCreated", logs: logs, sub: sub}, nil
}

// WatchOrderCreated is a free log subscription operation binding the contract event 0x7b1bb3d3aaec5962def20e424565d7eee44dd6e74694eb87d96f5415ebd57dd3.
//
// Solidity: event OrderCreated(bytes32 indexed hash, address indexed maker, address indexed optionContract, uint8 side, bool isDelayedWriting, uint256 optionId, address paymentToken, uint256 pricePerUnit, uint256 expirationTime, uint256 salt, uint256 amount, uint8 decimals)
func (_PremiaMarket *PremiaMarketFilterer) WatchOrderCreated(opts *bind.WatchOpts, sink chan<- *PremiaMarketOrderCreated, hash [][32]byte, maker []common.Address, optionContract []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.WatchLogs(opts, "OrderCreated", hashRule, makerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaMarketOrderCreated)
				if err := _PremiaMarket.contract.UnpackLog(event, "OrderCreated", log); err != nil {
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

// ParseOrderCreated is a log parse operation binding the contract event 0x7b1bb3d3aaec5962def20e424565d7eee44dd6e74694eb87d96f5415ebd57dd3.
//
// Solidity: event OrderCreated(bytes32 indexed hash, address indexed maker, address indexed optionContract, uint8 side, bool isDelayedWriting, uint256 optionId, address paymentToken, uint256 pricePerUnit, uint256 expirationTime, uint256 salt, uint256 amount, uint8 decimals)
func (_PremiaMarket *PremiaMarketFilterer) ParseOrderCreated(log types.Log) (*PremiaMarketOrderCreated, error) {
	event := new(PremiaMarketOrderCreated)
	if err := _PremiaMarket.contract.UnpackLog(event, "OrderCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaMarketOrderFilledIterator is returned from FilterOrderFilled and is used to iterate over the raw logs and unpacked data for OrderFilled events raised by the PremiaMarket contract.
type PremiaMarketOrderFilledIterator struct {
	Event *PremiaMarketOrderFilled // Event containing the contract specifics and raw log

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
func (it *PremiaMarketOrderFilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaMarketOrderFilled)
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
		it.Event = new(PremiaMarketOrderFilled)
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
func (it *PremiaMarketOrderFilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaMarketOrderFilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaMarketOrderFilled represents a OrderFilled event raised by the PremiaMarket contract.
type PremiaMarketOrderFilled struct {
	Hash           [32]byte
	Taker          common.Address
	OptionContract common.Address
	Maker          common.Address
	PaymentToken   common.Address
	Amount         *big.Int
	PricePerUnit   *big.Int
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterOrderFilled is a free log retrieval operation binding the contract event 0x0a7de5ac65db42a71326c71bc572df49b4eaa934a4fdd044ea7fefb03b621280.
//
// Solidity: event OrderFilled(bytes32 indexed hash, address indexed taker, address indexed optionContract, address maker, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) FilterOrderFilled(opts *bind.FilterOpts, hash [][32]byte, taker []common.Address, optionContract []common.Address) (*PremiaMarketOrderFilledIterator, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.FilterLogs(opts, "OrderFilled", hashRule, takerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketOrderFilledIterator{contract: _PremiaMarket.contract, event: "OrderFilled", logs: logs, sub: sub}, nil
}

// WatchOrderFilled is a free log subscription operation binding the contract event 0x0a7de5ac65db42a71326c71bc572df49b4eaa934a4fdd044ea7fefb03b621280.
//
// Solidity: event OrderFilled(bytes32 indexed hash, address indexed taker, address indexed optionContract, address maker, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) WatchOrderFilled(opts *bind.WatchOpts, sink chan<- *PremiaMarketOrderFilled, hash [][32]byte, taker []common.Address, optionContract []common.Address) (event.Subscription, error) {

	var hashRule []interface{}
	for _, hashItem := range hash {
		hashRule = append(hashRule, hashItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}
	var optionContractRule []interface{}
	for _, optionContractItem := range optionContract {
		optionContractRule = append(optionContractRule, optionContractItem)
	}

	logs, sub, err := _PremiaMarket.contract.WatchLogs(opts, "OrderFilled", hashRule, takerRule, optionContractRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaMarketOrderFilled)
				if err := _PremiaMarket.contract.UnpackLog(event, "OrderFilled", log); err != nil {
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

// ParseOrderFilled is a log parse operation binding the contract event 0x0a7de5ac65db42a71326c71bc572df49b4eaa934a4fdd044ea7fefb03b621280.
//
// Solidity: event OrderFilled(bytes32 indexed hash, address indexed taker, address indexed optionContract, address maker, address paymentToken, uint256 amount, uint256 pricePerUnit)
func (_PremiaMarket *PremiaMarketFilterer) ParseOrderFilled(log types.Log) (*PremiaMarketOrderFilled, error) {
	event := new(PremiaMarketOrderFilled)
	if err := _PremiaMarket.contract.UnpackLog(event, "OrderFilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PremiaMarketOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PremiaMarket contract.
type PremiaMarketOwnershipTransferredIterator struct {
	Event *PremiaMarketOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PremiaMarketOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PremiaMarketOwnershipTransferred)
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
		it.Event = new(PremiaMarketOwnershipTransferred)
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
func (it *PremiaMarketOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PremiaMarketOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PremiaMarketOwnershipTransferred represents a OwnershipTransferred event raised by the PremiaMarket contract.
type PremiaMarketOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PremiaMarket *PremiaMarketFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PremiaMarketOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PremiaMarket.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PremiaMarketOwnershipTransferredIterator{contract: _PremiaMarket.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PremiaMarket *PremiaMarketFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PremiaMarketOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PremiaMarket.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PremiaMarketOwnershipTransferred)
				if err := _PremiaMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PremiaMarket *PremiaMarketFilterer) ParseOwnershipTransferred(log types.Log) (*PremiaMarketOwnershipTransferred, error) {
	event := new(PremiaMarketOwnershipTransferred)
	if err := _PremiaMarket.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
