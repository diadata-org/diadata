// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ConverterTypeFour

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

// ConverterTypeFourABI is the input ABI used to generate the binding from.
const ConverterTypeFourABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_UPGRADER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"_virtualBalance\",\"type\":\"uint256\"}],\"name\":\"updateConnector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"connectors\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"},{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bancorX\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"}],\"name\":\"getReserveBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_index\",\"type\":\"uint256\"}],\"name\":\"connectorTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BNT_TOKEN\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allowRegistryUpdate\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_block\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"quickConvertPrioritized\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableConversions\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_REGISTRY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convertInternal\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"uint256[]\"}],\"name\":\"completeXConversion2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptTokenOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_magnitude\",\"type\":\"uint8\"}],\"name\":\"getFinalAmount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"converterType\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_weight\",\"type\":\"uint32\"},{\"name\":\"\",\"type\":\"bool\"}],\"name\":\"addConnector\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"liquidate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawFromToken\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newManager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"updateRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_whitelist\",\"type\":\"address\"}],\"name\":\"setConversionWhitelist\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_conversionId\",\"type\":\"uint256\"},{\"name\":\"_block\",\"type\":\"uint256\"},{\"name\":\"_v\",\"type\":\"uint8\"},{\"name\":\"_r\",\"type\":\"bytes32\"},{\"name\":\"_s\",\"type\":\"bytes32\"}],\"name\":\"completeXConversion\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"version\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_CONVERTER_FACTORY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"withdrawTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"change\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"prevRegistry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_scaleFactor\",\"type\":\"uint16\"}],\"name\":\"enableVirtualBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_ratio\",\"type\":\"uint32\"}],\"name\":\"addReserve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_FORMULA\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"convert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"connectorTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_sellAmount\",\"type\":\"uint256\"}],\"name\":\"getSaleReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_fromToken\",\"type\":\"address\"},{\"name\":\"_toToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"convert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"registry\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONTRACT_FEATURES\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromConnectorToken\",\"type\":\"address\"},{\"name\":\"_toConnectorToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getCrossConnectorReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_NETWORK\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_GAS_PRICE_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"CONVERTER_CONVERSION_WHITELIST\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_virtualBalance\",\"type\":\"uint256\"}],\"name\":\"updateReserveVirtualBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"maxConversionFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"reserveTokenCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint16\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"},{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableConnectorSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_depositAmount\",\"type\":\"uint256\"}],\"name\":\"getPurchaseReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableReserveSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_signature\",\"type\":\"uint256[]\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"quickConvertPrioritized2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"restoreRegistry\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionsEnabled\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"conversionWhitelist\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_X\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"acceptManagement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"fund\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"BANCOR_X_UPGRADER\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_fromReserveToken\",\"type\":\"address\"},{\"name\":\"_toReserveToken\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"getCrossReserveReturn\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"reserveTokens\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"upgrade\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"reserves\",\"outputs\":[{\"name\":\"virtualBalance\",\"type\":\"uint256\"},{\"name\":\"ratio\",\"type\":\"uint32\"},{\"name\":\"isVirtualBalanceEnabled\",\"type\":\"bool\"},{\"name\":\"isSaleEnabled\",\"type\":\"bool\"},{\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_connectorToken\",\"type\":\"address\"}],\"name\":\"getConnectorBalance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_bancorX\",\"type\":\"address\"}],\"name\":\"setBancorX\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"},{\"name\":\"_affiliateAccount\",\"type\":\"address\"},{\"name\":\"_affiliateFee\",\"type\":\"uint256\"}],\"name\":\"quickConvert2\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"transferManagement\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_conversionFee\",\"type\":\"uint32\"}],\"name\":\"setConversionFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_path\",\"type\":\"address[]\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_minReturn\",\"type\":\"uint256\"}],\"name\":\"quickConvert\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"NON_STANDARD_TOKEN_REGISTRY\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_disable\",\"type\":\"bool\"}],\"name\":\"disableRegistryUpdate\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"token\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"claimTokens\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_registry\",\"type\":\"address\"},{\"name\":\"_maxConversionFee\",\"type\":\"uint32\"},{\"name\":\"_reserveToken\",\"type\":\"address\"},{\"name\":\"_reserveRatio\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_fromToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_toToken\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_trader\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_return\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_conversionFee\",\"type\":\"int256\"}],\"name\":\"Conversion\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_connectorToken\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_tokenSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"_connectorWeight\",\"type\":\"uint32\"}],\"name\":\"PriceDataUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_prevFee\",\"type\":\"uint32\"},{\"indexed\":false,\"name\":\"_newFee\",\"type\":\"uint32\"}],\"name\":\"ConversionFeeUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_conversionsEnabled\",\"type\":\"bool\"}],\"name\":\"ConversionsEnable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"_enabled\",\"type\":\"bool\"}],\"name\":\"VirtualBalancesEnable\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevManager\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newManager\",\"type\":\"address\"}],\"name\":\"ManagerUpdate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_prevOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"OwnerUpdate\",\"type\":\"event\"}]"

// ConverterTypeFour is an auto generated Go binding around an Ethereum contract.
type ConverterTypeFour struct {
	ConverterTypeFourCaller     // Read-only binding to the contract
	ConverterTypeFourTransactor // Write-only binding to the contract
	ConverterTypeFourFilterer   // Log filterer for contract events
}

// ConverterTypeFourCaller is an auto generated read-only Go binding around an Ethereum contract.
type ConverterTypeFourCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeFourTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ConverterTypeFourTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeFourFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ConverterTypeFourFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ConverterTypeFourSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ConverterTypeFourSession struct {
	Contract     *ConverterTypeFour // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ConverterTypeFourCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ConverterTypeFourCallerSession struct {
	Contract *ConverterTypeFourCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// ConverterTypeFourTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ConverterTypeFourTransactorSession struct {
	Contract     *ConverterTypeFourTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// ConverterTypeFourRaw is an auto generated low-level Go binding around an Ethereum contract.
type ConverterTypeFourRaw struct {
	Contract *ConverterTypeFour // Generic contract binding to access the raw methods on
}

// ConverterTypeFourCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ConverterTypeFourCallerRaw struct {
	Contract *ConverterTypeFourCaller // Generic read-only contract binding to access the raw methods on
}

// ConverterTypeFourTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ConverterTypeFourTransactorRaw struct {
	Contract *ConverterTypeFourTransactor // Generic write-only contract binding to access the raw methods on
}

// NewConverterTypeFour creates a new instance of ConverterTypeFour, bound to a specific deployed contract.
func NewConverterTypeFour(address common.Address, backend bind.ContractBackend) (*ConverterTypeFour, error) {
	contract, err := bindConverterTypeFour(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFour{ConverterTypeFourCaller: ConverterTypeFourCaller{contract: contract}, ConverterTypeFourTransactor: ConverterTypeFourTransactor{contract: contract}, ConverterTypeFourFilterer: ConverterTypeFourFilterer{contract: contract}}, nil
}

// NewConverterTypeFourCaller creates a new read-only instance of ConverterTypeFour, bound to a specific deployed contract.
func NewConverterTypeFourCaller(address common.Address, caller bind.ContractCaller) (*ConverterTypeFourCaller, error) {
	contract, err := bindConverterTypeFour(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourCaller{contract: contract}, nil
}

// NewConverterTypeFourTransactor creates a new write-only instance of ConverterTypeFour, bound to a specific deployed contract.
func NewConverterTypeFourTransactor(address common.Address, transactor bind.ContractTransactor) (*ConverterTypeFourTransactor, error) {
	contract, err := bindConverterTypeFour(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourTransactor{contract: contract}, nil
}

// NewConverterTypeFourFilterer creates a new log filterer instance of ConverterTypeFour, bound to a specific deployed contract.
func NewConverterTypeFourFilterer(address common.Address, filterer bind.ContractFilterer) (*ConverterTypeFourFilterer, error) {
	contract, err := bindConverterTypeFour(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourFilterer{contract: contract}, nil
}

// bindConverterTypeFour binds a generic wrapper to an already deployed contract.
func bindConverterTypeFour(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ConverterTypeFourABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeFour *ConverterTypeFourRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeFour.Contract.ConverterTypeFourCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeFour *ConverterTypeFourRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ConverterTypeFourTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeFour *ConverterTypeFourRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ConverterTypeFourTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ConverterTypeFour *ConverterTypeFourCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ConverterTypeFour.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ConverterTypeFour *ConverterTypeFourTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ConverterTypeFour *ConverterTypeFourTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.contract.Transact(opts, method, params...)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORCONVERTERFACTORY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_CONVERTER_FACTORY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORCONVERTERFACTORY(&_ConverterTypeFour.CallOpts)
}

// BANCORCONVERTERFACTORY is a free data retrieval call binding the contract method 0x5a46f06c.
//
// Solidity: function BANCOR_CONVERTER_FACTORY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORCONVERTERFACTORY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORCONVERTERFACTORY(&_ConverterTypeFour.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORCONVERTERUPGRADER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_CONVERTER_UPGRADER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORCONVERTERUPGRADER(&_ConverterTypeFour.CallOpts)
}

// BANCORCONVERTERUPGRADER is a free data retrieval call binding the contract method 0x0c87355e.
//
// Solidity: function BANCOR_CONVERTER_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORCONVERTERUPGRADER() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORCONVERTERUPGRADER(&_ConverterTypeFour.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORFORMULA(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_FORMULA")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORFORMULA() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORFORMULA(&_ConverterTypeFour.CallOpts)
}

// BANCORFORMULA is a free data retrieval call binding the contract method 0x6d7bd3fc.
//
// Solidity: function BANCOR_FORMULA() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORFORMULA() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORFORMULA(&_ConverterTypeFour.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORGASPRICELIMIT(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_GAS_PRICE_LIMIT")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORGASPRICELIMIT(&_ConverterTypeFour.CallOpts)
}

// BANCORGASPRICELIMIT is a free data retrieval call binding the contract method 0x9249993a.
//
// Solidity: function BANCOR_GAS_PRICE_LIMIT() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORGASPRICELIMIT() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORGASPRICELIMIT(&_ConverterTypeFour.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORNETWORK(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_NETWORK")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORNETWORK() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORNETWORK(&_ConverterTypeFour.CallOpts)
}

// BANCORNETWORK is a free data retrieval call binding the contract method 0x9232494e.
//
// Solidity: function BANCOR_NETWORK() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORNETWORK() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORNETWORK(&_ConverterTypeFour.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORX(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_X")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORX() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORX(&_ConverterTypeFour.CallOpts)
}

// BANCORX is a free data retrieval call binding the contract method 0xc4a8598e.
//
// Solidity: function BANCOR_X() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORX() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORX(&_ConverterTypeFour.CallOpts)
}

// BANCORXUPGRADER is a free data retrieval call binding the contract method 0xcc97b38f.
//
// Solidity: function BANCOR_X_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BANCORXUPGRADER(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BANCOR_X_UPGRADER")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BANCORXUPGRADER is a free data retrieval call binding the contract method 0xcc97b38f.
//
// Solidity: function BANCOR_X_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BANCORXUPGRADER() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORXUPGRADER(&_ConverterTypeFour.CallOpts)
}

// BANCORXUPGRADER is a free data retrieval call binding the contract method 0xcc97b38f.
//
// Solidity: function BANCOR_X_UPGRADER() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BANCORXUPGRADER() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BANCORXUPGRADER(&_ConverterTypeFour.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) BNTTOKEN(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "BNT_TOKEN")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) BNTTOKEN() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BNTTOKEN(&_ConverterTypeFour.CallOpts)
}

// BNTTOKEN is a free data retrieval call binding the contract method 0x1d000b61.
//
// Solidity: function BNT_TOKEN() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BNTTOKEN() ([32]byte, error) {
	return _ConverterTypeFour.Contract.BNTTOKEN(&_ConverterTypeFour.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) CONTRACTFEATURES(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "CONTRACT_FEATURES")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) CONTRACTFEATURES() ([32]byte, error) {
	return _ConverterTypeFour.Contract.CONTRACTFEATURES(&_ConverterTypeFour.CallOpts)
}

// CONTRACTFEATURES is a free data retrieval call binding the contract method 0x83315b6e.
//
// Solidity: function CONTRACT_FEATURES() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) CONTRACTFEATURES() ([32]byte, error) {
	return _ConverterTypeFour.Contract.CONTRACTFEATURES(&_ConverterTypeFour.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) CONTRACTREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "CONTRACT_REGISTRY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) CONTRACTREGISTRY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.CONTRACTREGISTRY(&_ConverterTypeFour.CallOpts)
}

// CONTRACTREGISTRY is a free data retrieval call binding the contract method 0x25f9bfef.
//
// Solidity: function CONTRACT_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) CONTRACTREGISTRY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.CONTRACTREGISTRY(&_ConverterTypeFour.CallOpts)
}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) CONVERTERCONVERSIONWHITELIST(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "CONVERTER_CONVERSION_WHITELIST")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) CONVERTERCONVERSIONWHITELIST() (*big.Int, error) {
	return _ConverterTypeFour.Contract.CONVERTERCONVERSIONWHITELIST(&_ConverterTypeFour.CallOpts)
}

// CONVERTERCONVERSIONWHITELIST is a free data retrieval call binding the contract method 0x92d1abb7.
//
// Solidity: function CONVERTER_CONVERSION_WHITELIST() view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) CONVERTERCONVERSIONWHITELIST() (*big.Int, error) {
	return _ConverterTypeFour.Contract.CONVERTERCONVERSIONWHITELIST(&_ConverterTypeFour.CallOpts)
}

// NONSTANDARDTOKENREGISTRY is a free data retrieval call binding the contract method 0xf5286b9c.
//
// Solidity: function NON_STANDARD_TOKEN_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCaller) NONSTANDARDTOKENREGISTRY(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "NON_STANDARD_TOKEN_REGISTRY")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// NONSTANDARDTOKENREGISTRY is a free data retrieval call binding the contract method 0xf5286b9c.
//
// Solidity: function NON_STANDARD_TOKEN_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourSession) NONSTANDARDTOKENREGISTRY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.NONSTANDARDTOKENREGISTRY(&_ConverterTypeFour.CallOpts)
}

// NONSTANDARDTOKENREGISTRY is a free data retrieval call binding the contract method 0xf5286b9c.
//
// Solidity: function NON_STANDARD_TOKEN_REGISTRY() view returns(bytes32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) NONSTANDARDTOKENREGISTRY() ([32]byte, error) {
	return _ConverterTypeFour.Contract.NONSTANDARDTOKENREGISTRY(&_ConverterTypeFour.CallOpts)
}

// AllowRegistryUpdate is a free data retrieval call binding the contract method 0x20d7d367.
//
// Solidity: function allowRegistryUpdate() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourCaller) AllowRegistryUpdate(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "allowRegistryUpdate")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowRegistryUpdate is a free data retrieval call binding the contract method 0x20d7d367.
//
// Solidity: function allowRegistryUpdate() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourSession) AllowRegistryUpdate() (bool, error) {
	return _ConverterTypeFour.Contract.AllowRegistryUpdate(&_ConverterTypeFour.CallOpts)
}

// AllowRegistryUpdate is a free data retrieval call binding the contract method 0x20d7d367.
//
// Solidity: function allowRegistryUpdate() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) AllowRegistryUpdate() (bool, error) {
	return _ConverterTypeFour.Contract.AllowRegistryUpdate(&_ConverterTypeFour.CallOpts)
}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) BancorX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "bancorX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) BancorX() (common.Address, error) {
	return _ConverterTypeFour.Contract.BancorX(&_ConverterTypeFour.CallOpts)
}

// BancorX is a free data retrieval call binding the contract method 0x1120a776.
//
// Solidity: function bancorX() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) BancorX() (common.Address, error) {
	return _ConverterTypeFour.Contract.BancorX(&_ConverterTypeFour.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConnectorTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "connectorTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeFour.Contract.ConnectorTokenCount(&_ConverterTypeFour.CallOpts)
}

// ConnectorTokenCount is a free data retrieval call binding the contract method 0x71f52bf3.
//
// Solidity: function connectorTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConnectorTokenCount() (uint16, error) {
	return _ConverterTypeFour.Contract.ConnectorTokenCount(&_ConverterTypeFour.CallOpts)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConnectorTokens(opts *bind.CallOpts, _index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "connectorTokens", _index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeFour.Contract.ConnectorTokens(&_ConverterTypeFour.CallOpts, _index)
}

// ConnectorTokens is a free data retrieval call binding the contract method 0x19b64015.
//
// Solidity: function connectorTokens(uint256 _index) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConnectorTokens(_index *big.Int) (common.Address, error) {
	return _ConverterTypeFour.Contract.ConnectorTokens(&_ConverterTypeFour.CallOpts, _index)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeFour *ConverterTypeFourCaller) Connectors(opts *bind.CallOpts, _address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "connectors", _address)

	if err != nil {
		return *new(*big.Int), *new(uint32), *new(bool), *new(bool), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)
	out3 := *abi.ConvertType(out[3], new(bool)).(*bool)
	out4 := *abi.ConvertType(out[4], new(bool)).(*bool)

	return out0, out1, out2, out3, out4, err

}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeFour *ConverterTypeFourSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeFour.Contract.Connectors(&_ConverterTypeFour.CallOpts, _address)
}

// Connectors is a free data retrieval call binding the contract method 0x0e53aae9.
//
// Solidity: function connectors(address _address) view returns(uint256, uint32, bool, bool, bool)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Connectors(_address common.Address) (*big.Int, uint32, bool, bool, bool, error) {
	return _ConverterTypeFour.Contract.Connectors(&_ConverterTypeFour.CallOpts, _address)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "conversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourSession) ConversionFee() (uint32, error) {
	return _ConverterTypeFour.Contract.ConversionFee(&_ConverterTypeFour.CallOpts)
}

// ConversionFee is a free data retrieval call binding the contract method 0x579cd3ca.
//
// Solidity: function conversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConversionFee() (uint32, error) {
	return _ConverterTypeFour.Contract.ConversionFee(&_ConverterTypeFour.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConversionWhitelist(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "conversionWhitelist")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) ConversionWhitelist() (common.Address, error) {
	return _ConverterTypeFour.Contract.ConversionWhitelist(&_ConverterTypeFour.CallOpts)
}

// ConversionWhitelist is a free data retrieval call binding the contract method 0xc45d3d92.
//
// Solidity: function conversionWhitelist() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConversionWhitelist() (common.Address, error) {
	return _ConverterTypeFour.Contract.ConversionWhitelist(&_ConverterTypeFour.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConversionsEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "conversionsEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourSession) ConversionsEnabled() (bool, error) {
	return _ConverterTypeFour.Contract.ConversionsEnabled(&_ConverterTypeFour.CallOpts)
}

// ConversionsEnabled is a free data retrieval call binding the contract method 0xbf754558.
//
// Solidity: function conversionsEnabled() view returns(bool)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConversionsEnabled() (bool, error) {
	return _ConverterTypeFour.Contract.ConversionsEnabled(&_ConverterTypeFour.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() view returns(string)
func (_ConverterTypeFour *ConverterTypeFourCaller) ConverterType(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "converterType")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() view returns(string)
func (_ConverterTypeFour *ConverterTypeFourSession) ConverterType() (string, error) {
	return _ConverterTypeFour.Contract.ConverterType(&_ConverterTypeFour.CallOpts)
}

// ConverterType is a free data retrieval call binding the contract method 0x3e8ff43f.
//
// Solidity: function converterType() view returns(string)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ConverterType() (string, error) {
	return _ConverterTypeFour.Contract.ConverterType(&_ConverterTypeFour.CallOpts)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetConnectorBalance(opts *bind.CallOpts, _connectorToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getConnectorBalance", _connectorToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetConnectorBalance(&_ConverterTypeFour.CallOpts, _connectorToken)
}

// GetConnectorBalance is a free data retrieval call binding the contract method 0xd8959512.
//
// Solidity: function getConnectorBalance(address _connectorToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetConnectorBalance(_connectorToken common.Address) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetConnectorBalance(&_ConverterTypeFour.CallOpts, _connectorToken)
}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetCrossConnectorReturn(opts *bind.CallOpts, _fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getCrossConnectorReturn", _fromConnectorToken, _toConnectorToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetCrossConnectorReturn(_fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetCrossConnectorReturn(&_ConverterTypeFour.CallOpts, _fromConnectorToken, _toConnectorToken, _amount)
}

// GetCrossConnectorReturn is a free data retrieval call binding the contract method 0x8e3047e0.
//
// Solidity: function getCrossConnectorReturn(address _fromConnectorToken, address _toConnectorToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetCrossConnectorReturn(_fromConnectorToken common.Address, _toConnectorToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetCrossConnectorReturn(&_ConverterTypeFour.CallOpts, _fromConnectorToken, _toConnectorToken, _amount)
}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetCrossReserveReturn(opts *bind.CallOpts, _fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getCrossReserveReturn", _fromReserveToken, _toReserveToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetCrossReserveReturn(_fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetCrossReserveReturn(&_ConverterTypeFour.CallOpts, _fromReserveToken, _toReserveToken, _amount)
}

// GetCrossReserveReturn is a free data retrieval call binding the contract method 0xcf73266a.
//
// Solidity: function getCrossReserveReturn(address _fromReserveToken, address _toReserveToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetCrossReserveReturn(_fromReserveToken common.Address, _toReserveToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetCrossReserveReturn(&_ConverterTypeFour.CallOpts, _fromReserveToken, _toReserveToken, _amount)
}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetFinalAmount(opts *bind.CallOpts, _amount *big.Int, _magnitude uint8) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getFinalAmount", _amount, _magnitude)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetFinalAmount(_amount *big.Int, _magnitude uint8) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetFinalAmount(&_ConverterTypeFour.CallOpts, _amount, _magnitude)
}

// GetFinalAmount is a free data retrieval call binding the contract method 0x3aa0145a.
//
// Solidity: function getFinalAmount(uint256 _amount, uint8 _magnitude) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetFinalAmount(_amount *big.Int, _magnitude uint8) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetFinalAmount(&_ConverterTypeFour.CallOpts, _amount, _magnitude)
}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetPurchaseReturn(opts *bind.CallOpts, _reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getPurchaseReturn", _reserveToken, _depositAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetPurchaseReturn(_reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetPurchaseReturn(&_ConverterTypeFour.CallOpts, _reserveToken, _depositAmount)
}

// GetPurchaseReturn is a free data retrieval call binding the contract method 0xa2c4c336.
//
// Solidity: function getPurchaseReturn(address _reserveToken, uint256 _depositAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetPurchaseReturn(_reserveToken common.Address, _depositAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetPurchaseReturn(&_ConverterTypeFour.CallOpts, _reserveToken, _depositAmount)
}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetReserveBalance(opts *bind.CallOpts, _reserveToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getReserveBalance", _reserveToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetReserveBalance(&_ConverterTypeFour.CallOpts, _reserveToken)
}

// GetReserveBalance is a free data retrieval call binding the contract method 0x15226b54.
//
// Solidity: function getReserveBalance(address _reserveToken) view returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetReserveBalance(_reserveToken common.Address) (*big.Int, error) {
	return _ConverterTypeFour.Contract.GetReserveBalance(&_ConverterTypeFour.CallOpts, _reserveToken)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetReturn(opts *bind.CallOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getReturn", _fromToken, _toToken, _amount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetReturn(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetReturn(&_ConverterTypeFour.CallOpts, _fromToken, _toToken, _amount)
}

// GetReturn is a free data retrieval call binding the contract method 0x1e1401f8.
//
// Solidity: function getReturn(address _fromToken, address _toToken, uint256 _amount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetReturn(_fromToken common.Address, _toToken common.Address, _amount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetReturn(&_ConverterTypeFour.CallOpts, _fromToken, _toToken, _amount)
}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCaller) GetSaleReturn(opts *bind.CallOpts, _reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "getSaleReturn", _reserveToken, _sellAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) GetSaleReturn(_reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetSaleReturn(&_ConverterTypeFour.CallOpts, _reserveToken, _sellAmount)
}

// GetSaleReturn is a free data retrieval call binding the contract method 0x72b44b2c.
//
// Solidity: function getSaleReturn(address _reserveToken, uint256 _sellAmount) view returns(uint256, uint256)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) GetSaleReturn(_reserveToken common.Address, _sellAmount *big.Int) (*big.Int, *big.Int, error) {
	return _ConverterTypeFour.Contract.GetSaleReturn(&_ConverterTypeFour.CallOpts, _reserveToken, _sellAmount)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) Manager() (common.Address, error) {
	return _ConverterTypeFour.Contract.Manager(&_ConverterTypeFour.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Manager() (common.Address, error) {
	return _ConverterTypeFour.Contract.Manager(&_ConverterTypeFour.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourCaller) MaxConversionFee(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "maxConversionFee")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeFour.Contract.MaxConversionFee(&_ConverterTypeFour.CallOpts)
}

// MaxConversionFee is a free data retrieval call binding the contract method 0x94c275ad.
//
// Solidity: function maxConversionFee() view returns(uint32)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) MaxConversionFee() (uint32, error) {
	return _ConverterTypeFour.Contract.MaxConversionFee(&_ConverterTypeFour.CallOpts)
}

// NewManager is a free data retrieval call binding the contract method 0x42906029.
//
// Solidity: function newManager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) NewManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "newManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewManager is a free data retrieval call binding the contract method 0x42906029.
//
// Solidity: function newManager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) NewManager() (common.Address, error) {
	return _ConverterTypeFour.Contract.NewManager(&_ConverterTypeFour.CallOpts)
}

// NewManager is a free data retrieval call binding the contract method 0x42906029.
//
// Solidity: function newManager() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) NewManager() (common.Address, error) {
	return _ConverterTypeFour.Contract.NewManager(&_ConverterTypeFour.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) NewOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "newOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) NewOwner() (common.Address, error) {
	return _ConverterTypeFour.Contract.NewOwner(&_ConverterTypeFour.CallOpts)
}

// NewOwner is a free data retrieval call binding the contract method 0xd4ee1d90.
//
// Solidity: function newOwner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) NewOwner() (common.Address, error) {
	return _ConverterTypeFour.Contract.NewOwner(&_ConverterTypeFour.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) Owner() (common.Address, error) {
	return _ConverterTypeFour.Contract.Owner(&_ConverterTypeFour.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Owner() (common.Address, error) {
	return _ConverterTypeFour.Contract.Owner(&_ConverterTypeFour.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) PrevRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "prevRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeFour.Contract.PrevRegistry(&_ConverterTypeFour.CallOpts)
}

// PrevRegistry is a free data retrieval call binding the contract method 0x61cd756e.
//
// Solidity: function prevRegistry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) PrevRegistry() (common.Address, error) {
	return _ConverterTypeFour.Contract.PrevRegistry(&_ConverterTypeFour.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) Registry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "registry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) Registry() (common.Address, error) {
	return _ConverterTypeFour.Contract.Registry(&_ConverterTypeFour.CallOpts)
}

// Registry is a free data retrieval call binding the contract method 0x7b103999.
//
// Solidity: function registry() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Registry() (common.Address, error) {
	return _ConverterTypeFour.Contract.Registry(&_ConverterTypeFour.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCaller) ReserveTokenCount(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "reserveTokenCount")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeFour.Contract.ReserveTokenCount(&_ConverterTypeFour.CallOpts)
}

// ReserveTokenCount is a free data retrieval call binding the contract method 0x9b99a8e2.
//
// Solidity: function reserveTokenCount() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ReserveTokenCount() (uint16, error) {
	return _ConverterTypeFour.Contract.ReserveTokenCount(&_ConverterTypeFour.CallOpts)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) ReserveTokens(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "reserveTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConverterTypeFour.Contract.ReserveTokens(&_ConverterTypeFour.CallOpts, arg0)
}

// ReserveTokens is a free data retrieval call binding the contract method 0xd031370b.
//
// Solidity: function reserveTokens(uint256 ) view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) ReserveTokens(arg0 *big.Int) (common.Address, error) {
	return _ConverterTypeFour.Contract.ReserveTokens(&_ConverterTypeFour.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeFour *ConverterTypeFourCaller) Reserves(opts *bind.CallOpts, arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "reserves", arg0)

	outstruct := new(struct {
		VirtualBalance          *big.Int
		Ratio                   uint32
		IsVirtualBalanceEnabled bool
		IsSaleEnabled           bool
		IsSet                   bool
	})

	outstruct.VirtualBalance = out[0].(*big.Int)
	outstruct.Ratio = out[1].(uint32)
	outstruct.IsVirtualBalanceEnabled = out[2].(bool)
	outstruct.IsSaleEnabled = out[3].(bool)
	outstruct.IsSet = out[4].(bool)

	return *outstruct, err

}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeFour *ConverterTypeFourSession) Reserves(arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	return _ConverterTypeFour.Contract.Reserves(&_ConverterTypeFour.CallOpts, arg0)
}

// Reserves is a free data retrieval call binding the contract method 0xd66bd524.
//
// Solidity: function reserves(address ) view returns(uint256 virtualBalance, uint32 ratio, bool isVirtualBalanceEnabled, bool isSaleEnabled, bool isSet)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Reserves(arg0 common.Address) (struct {
	VirtualBalance          *big.Int
	Ratio                   uint32
	IsVirtualBalanceEnabled bool
	IsSaleEnabled           bool
	IsSet                   bool
}, error) {
	return _ConverterTypeFour.Contract.Reserves(&_ConverterTypeFour.CallOpts, arg0)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourSession) Token() (common.Address, error) {
	return _ConverterTypeFour.Contract.Token(&_ConverterTypeFour.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Token() (common.Address, error) {
	return _ConverterTypeFour.Contract.Token(&_ConverterTypeFour.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCaller) Version(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ConverterTypeFour.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourSession) Version() (uint16, error) {
	return _ConverterTypeFour.Contract.Version(&_ConverterTypeFour.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() view returns(uint16)
func (_ConverterTypeFour *ConverterTypeFourCallerSession) Version() (uint16, error) {
	return _ConverterTypeFour.Contract.Version(&_ConverterTypeFour.CallOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) AcceptManagement(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "acceptManagement")
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) AcceptManagement() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptManagement(&_ConverterTypeFour.TransactOpts)
}

// AcceptManagement is a paid mutator transaction binding the contract method 0xc8c2fe6c.
//
// Solidity: function acceptManagement() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) AcceptManagement() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptManagement(&_ConverterTypeFour.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) AcceptOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "acceptOwnership")
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptOwnership(&_ConverterTypeFour.TransactOpts)
}

// AcceptOwnership is a paid mutator transaction binding the contract method 0x79ba5097.
//
// Solidity: function acceptOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) AcceptOwnership() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptOwnership(&_ConverterTypeFour.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) AcceptTokenOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "acceptTokenOwnership")
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptTokenOwnership(&_ConverterTypeFour.TransactOpts)
}

// AcceptTokenOwnership is a paid mutator transaction binding the contract method 0x38a5e016.
//
// Solidity: function acceptTokenOwnership() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) AcceptTokenOwnership() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AcceptTokenOwnership(&_ConverterTypeFour.TransactOpts)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) AddConnector(opts *bind.TransactOpts, _token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "addConnector", _token, _weight, arg2)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) AddConnector(_token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AddConnector(&_ConverterTypeFour.TransactOpts, _token, _weight, arg2)
}

// AddConnector is a paid mutator transaction binding the contract method 0x3f4d2fc2.
//
// Solidity: function addConnector(address _token, uint32 _weight, bool ) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) AddConnector(_token common.Address, _weight uint32, arg2 bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AddConnector(&_ConverterTypeFour.TransactOpts, _token, _weight, arg2)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) AddReserve(opts *bind.TransactOpts, _token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "addReserve", _token, _ratio)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) AddReserve(_token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AddReserve(&_ConverterTypeFour.TransactOpts, _token, _ratio)
}

// AddReserve is a paid mutator transaction binding the contract method 0x6a49d2c4.
//
// Solidity: function addReserve(address _token, uint32 _ratio) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) AddReserve(_token common.Address, _ratio uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.AddReserve(&_ConverterTypeFour.TransactOpts, _token, _ratio)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) Change(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "change", _fromToken, _toToken, _amount, _minReturn)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) Change(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Change(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Change is a paid mutator transaction binding the contract method 0x5e5144eb.
//
// Solidity: function change(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Change(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Change(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) ClaimTokens(opts *bind.TransactOpts, _from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "claimTokens", _from, _amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) ClaimTokens(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ClaimTokens(&_ConverterTypeFour.TransactOpts, _from, _amount)
}

// ClaimTokens is a paid mutator transaction binding the contract method 0xfe417fa5.
//
// Solidity: function claimTokens(address _from, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) ClaimTokens(_from common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ClaimTokens(&_ConverterTypeFour.TransactOpts, _from, _amount)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) CompleteXConversion(opts *bind.TransactOpts, _path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "completeXConversion", _path, _minReturn, _conversionId, _block, _v, _r, _s)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) CompleteXConversion(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.CompleteXConversion(&_ConverterTypeFour.TransactOpts, _path, _minReturn, _conversionId, _block, _v, _r, _s)
}

// CompleteXConversion is a paid mutator transaction binding the contract method 0x50057351.
//
// Solidity: function completeXConversion(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) CompleteXConversion(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.CompleteXConversion(&_ConverterTypeFour.TransactOpts, _path, _minReturn, _conversionId, _block, _v, _r, _s)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0x2cc1cd65.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256[] _signature) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) CompleteXConversion2(opts *bind.TransactOpts, _path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _signature []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "completeXConversion2", _path, _minReturn, _conversionId, _signature)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0x2cc1cd65.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256[] _signature) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) CompleteXConversion2(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _signature []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.CompleteXConversion2(&_ConverterTypeFour.TransactOpts, _path, _minReturn, _conversionId, _signature)
}

// CompleteXConversion2 is a paid mutator transaction binding the contract method 0x2cc1cd65.
//
// Solidity: function completeXConversion2(address[] _path, uint256 _minReturn, uint256 _conversionId, uint256[] _signature) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) CompleteXConversion2(_path []common.Address, _minReturn *big.Int, _conversionId *big.Int, _signature []*big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.CompleteXConversion2(&_ConverterTypeFour.TransactOpts, _path, _minReturn, _conversionId, _signature)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) Convert(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "convert", _fromToken, _toToken, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) Convert(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Convert(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Convert is a paid mutator transaction binding the contract method 0x75892cf1.
//
// Solidity: function convert(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Convert(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Convert(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) Convert2(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "convert2", _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) Convert2(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Convert2(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// Convert2 is a paid mutator transaction binding the contract method 0x6ebf36c0.
//
// Solidity: function convert2(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Convert2(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Convert2(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) ConvertInternal(opts *bind.TransactOpts, _fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "convertInternal", _fromToken, _toToken, _amount, _minReturn)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) ConvertInternal(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ConvertInternal(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// ConvertInternal is a paid mutator transaction binding the contract method 0x2a2e2f0c.
//
// Solidity: function convertInternal(address _fromToken, address _toToken, uint256 _amount, uint256 _minReturn) returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) ConvertInternal(_fromToken common.Address, _toToken common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.ConvertInternal(&_ConverterTypeFour.TransactOpts, _fromToken, _toToken, _amount, _minReturn)
}

// DisableConnectorSale is a paid mutator transaction binding the contract method 0x9e568553.
//
// Solidity: function disableConnectorSale(address _connectorToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) DisableConnectorSale(opts *bind.TransactOpts, _connectorToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "disableConnectorSale", _connectorToken, _disable)
}

// DisableConnectorSale is a paid mutator transaction binding the contract method 0x9e568553.
//
// Solidity: function disableConnectorSale(address _connectorToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) DisableConnectorSale(_connectorToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableConnectorSale(&_ConverterTypeFour.TransactOpts, _connectorToken, _disable)
}

// DisableConnectorSale is a paid mutator transaction binding the contract method 0x9e568553.
//
// Solidity: function disableConnectorSale(address _connectorToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) DisableConnectorSale(_connectorToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableConnectorSale(&_ConverterTypeFour.TransactOpts, _connectorToken, _disable)
}

// DisableConversions is a paid mutator transaction binding the contract method 0x228d2820.
//
// Solidity: function disableConversions(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) DisableConversions(opts *bind.TransactOpts, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "disableConversions", _disable)
}

// DisableConversions is a paid mutator transaction binding the contract method 0x228d2820.
//
// Solidity: function disableConversions(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) DisableConversions(_disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableConversions(&_ConverterTypeFour.TransactOpts, _disable)
}

// DisableConversions is a paid mutator transaction binding the contract method 0x228d2820.
//
// Solidity: function disableConversions(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) DisableConversions(_disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableConversions(&_ConverterTypeFour.TransactOpts, _disable)
}

// DisableRegistryUpdate is a paid mutator transaction binding the contract method 0xfa1c594e.
//
// Solidity: function disableRegistryUpdate(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) DisableRegistryUpdate(opts *bind.TransactOpts, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "disableRegistryUpdate", _disable)
}

// DisableRegistryUpdate is a paid mutator transaction binding the contract method 0xfa1c594e.
//
// Solidity: function disableRegistryUpdate(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) DisableRegistryUpdate(_disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableRegistryUpdate(&_ConverterTypeFour.TransactOpts, _disable)
}

// DisableRegistryUpdate is a paid mutator transaction binding the contract method 0xfa1c594e.
//
// Solidity: function disableRegistryUpdate(bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) DisableRegistryUpdate(_disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableRegistryUpdate(&_ConverterTypeFour.TransactOpts, _disable)
}

// DisableReserveSale is a paid mutator transaction binding the contract method 0xa6a11c71.
//
// Solidity: function disableReserveSale(address _reserveToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) DisableReserveSale(opts *bind.TransactOpts, _reserveToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "disableReserveSale", _reserveToken, _disable)
}

// DisableReserveSale is a paid mutator transaction binding the contract method 0xa6a11c71.
//
// Solidity: function disableReserveSale(address _reserveToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) DisableReserveSale(_reserveToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableReserveSale(&_ConverterTypeFour.TransactOpts, _reserveToken, _disable)
}

// DisableReserveSale is a paid mutator transaction binding the contract method 0xa6a11c71.
//
// Solidity: function disableReserveSale(address _reserveToken, bool _disable) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) DisableReserveSale(_reserveToken common.Address, _disable bool) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.DisableReserveSale(&_ConverterTypeFour.TransactOpts, _reserveToken, _disable)
}

// EnableVirtualBalances is a paid mutator transaction binding the contract method 0x677c0812.
//
// Solidity: function enableVirtualBalances(uint16 _scaleFactor) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) EnableVirtualBalances(opts *bind.TransactOpts, _scaleFactor uint16) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "enableVirtualBalances", _scaleFactor)
}

// EnableVirtualBalances is a paid mutator transaction binding the contract method 0x677c0812.
//
// Solidity: function enableVirtualBalances(uint16 _scaleFactor) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) EnableVirtualBalances(_scaleFactor uint16) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.EnableVirtualBalances(&_ConverterTypeFour.TransactOpts, _scaleFactor)
}

// EnableVirtualBalances is a paid mutator transaction binding the contract method 0x677c0812.
//
// Solidity: function enableVirtualBalances(uint16 _scaleFactor) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) EnableVirtualBalances(_scaleFactor uint16) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.EnableVirtualBalances(&_ConverterTypeFour.TransactOpts, _scaleFactor)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) Fund(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "fund", _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Fund(&_ConverterTypeFour.TransactOpts, _amount)
}

// Fund is a paid mutator transaction binding the contract method 0xca1d209d.
//
// Solidity: function fund(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Fund(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Fund(&_ConverterTypeFour.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) Liquidate(opts *bind.TransactOpts, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "liquidate", _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Liquidate(&_ConverterTypeFour.TransactOpts, _amount)
}

// Liquidate is a paid mutator transaction binding the contract method 0x415f1240.
//
// Solidity: function liquidate(uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Liquidate(_amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Liquidate(&_ConverterTypeFour.TransactOpts, _amount)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) QuickConvert(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "quickConvert", _path, _amount, _minReturn)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) QuickConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvert(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn)
}

// QuickConvert is a paid mutator transaction binding the contract method 0xf0843ba9.
//
// Solidity: function quickConvert(address[] _path, uint256 _amount, uint256 _minReturn) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) QuickConvert(_path []common.Address, _amount *big.Int, _minReturn *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvert(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) QuickConvert2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "quickConvert2", _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) QuickConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvert2(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvert2 is a paid mutator transaction binding the contract method 0xe4dd22f6.
//
// Solidity: function quickConvert2(address[] _path, uint256 _amount, uint256 _minReturn, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) QuickConvert2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvert2(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) QuickConvertPrioritized(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "quickConvertPrioritized", _path, _amount, _minReturn, _block, _v, _r, _s)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) QuickConvertPrioritized(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvertPrioritized(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _block, _v, _r, _s)
}

// QuickConvertPrioritized is a paid mutator transaction binding the contract method 0x22742564.
//
// Solidity: function quickConvertPrioritized(address[] _path, uint256 _amount, uint256 _minReturn, uint256 _block, uint8 _v, bytes32 _r, bytes32 _s) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) QuickConvertPrioritized(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _block *big.Int, _v uint8, _r [32]byte, _s [32]byte) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvertPrioritized(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _block, _v, _r, _s)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] _signature, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactor) QuickConvertPrioritized2(opts *bind.TransactOpts, _path []common.Address, _amount *big.Int, _minReturn *big.Int, _signature []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "quickConvertPrioritized2", _path, _amount, _minReturn, _signature, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] _signature, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourSession) QuickConvertPrioritized2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _signature []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvertPrioritized2(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _signature, _affiliateAccount, _affiliateFee)
}

// QuickConvertPrioritized2 is a paid mutator transaction binding the contract method 0xb3a426d5.
//
// Solidity: function quickConvertPrioritized2(address[] _path, uint256 _amount, uint256 _minReturn, uint256[] _signature, address _affiliateAccount, uint256 _affiliateFee) payable returns(uint256)
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) QuickConvertPrioritized2(_path []common.Address, _amount *big.Int, _minReturn *big.Int, _signature []*big.Int, _affiliateAccount common.Address, _affiliateFee *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.QuickConvertPrioritized2(&_ConverterTypeFour.TransactOpts, _path, _amount, _minReturn, _signature, _affiliateAccount, _affiliateFee)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) RestoreRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "restoreRegistry")
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.RestoreRegistry(&_ConverterTypeFour.TransactOpts)
}

// RestoreRegistry is a paid mutator transaction binding the contract method 0xb4a176d3.
//
// Solidity: function restoreRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) RestoreRegistry() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.RestoreRegistry(&_ConverterTypeFour.TransactOpts)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) SetBancorX(opts *bind.TransactOpts, _bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "setBancorX", _bancorX)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) SetBancorX(_bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetBancorX(&_ConverterTypeFour.TransactOpts, _bancorX)
}

// SetBancorX is a paid mutator transaction binding the contract method 0xd924f0c3.
//
// Solidity: function setBancorX(address _bancorX) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) SetBancorX(_bancorX common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetBancorX(&_ConverterTypeFour.TransactOpts, _bancorX)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) SetConversionFee(opts *bind.TransactOpts, _conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "setConversionFee", _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetConversionFee(&_ConverterTypeFour.TransactOpts, _conversionFee)
}

// SetConversionFee is a paid mutator transaction binding the contract method 0xecbca55d.
//
// Solidity: function setConversionFee(uint32 _conversionFee) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) SetConversionFee(_conversionFee uint32) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetConversionFee(&_ConverterTypeFour.TransactOpts, _conversionFee)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) SetConversionWhitelist(opts *bind.TransactOpts, _whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "setConversionWhitelist", _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetConversionWhitelist(&_ConverterTypeFour.TransactOpts, _whitelist)
}

// SetConversionWhitelist is a paid mutator transaction binding the contract method 0x4af80f0e.
//
// Solidity: function setConversionWhitelist(address _whitelist) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) SetConversionWhitelist(_whitelist common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.SetConversionWhitelist(&_ConverterTypeFour.TransactOpts, _whitelist)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(address _newManager) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) TransferManagement(opts *bind.TransactOpts, _newManager common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "transferManagement", _newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(address _newManager) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) TransferManagement(_newManager common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferManagement(&_ConverterTypeFour.TransactOpts, _newManager)
}

// TransferManagement is a paid mutator transaction binding the contract method 0xe4edf852.
//
// Solidity: function transferManagement(address _newManager) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) TransferManagement(_newManager common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferManagement(&_ConverterTypeFour.TransactOpts, _newManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferOwnership(&_ConverterTypeFour.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferOwnership(&_ConverterTypeFour.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) TransferTokenOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "transferTokenOwnership", _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferTokenOwnership(&_ConverterTypeFour.TransactOpts, _newOwner)
}

// TransferTokenOwnership is a paid mutator transaction binding the contract method 0x21e6b53d.
//
// Solidity: function transferTokenOwnership(address _newOwner) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) TransferTokenOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.TransferTokenOwnership(&_ConverterTypeFour.TransactOpts, _newOwner)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) UpdateConnector(opts *bind.TransactOpts, _connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "updateConnector", _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) UpdateConnector(_connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateConnector(&_ConverterTypeFour.TransactOpts, _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateConnector is a paid mutator transaction binding the contract method 0x0ca78923.
//
// Solidity: function updateConnector(address _connectorToken, uint32 , bool , uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) UpdateConnector(_connectorToken common.Address, arg1 uint32, arg2 bool, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateConnector(&_ConverterTypeFour.TransactOpts, _connectorToken, arg1, arg2, _virtualBalance)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) UpdateRegistry(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "updateRegistry")
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateRegistry(&_ConverterTypeFour.TransactOpts)
}

// UpdateRegistry is a paid mutator transaction binding the contract method 0x49d10b64.
//
// Solidity: function updateRegistry() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) UpdateRegistry() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateRegistry(&_ConverterTypeFour.TransactOpts)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) UpdateReserveVirtualBalance(opts *bind.TransactOpts, _reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "updateReserveVirtualBalance", _reserveToken, _virtualBalance)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) UpdateReserveVirtualBalance(_reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateReserveVirtualBalance(&_ConverterTypeFour.TransactOpts, _reserveToken, _virtualBalance)
}

// UpdateReserveVirtualBalance is a paid mutator transaction binding the contract method 0x935e2ae1.
//
// Solidity: function updateReserveVirtualBalance(address _reserveToken, uint256 _virtualBalance) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) UpdateReserveVirtualBalance(_reserveToken common.Address, _virtualBalance *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.UpdateReserveVirtualBalance(&_ConverterTypeFour.TransactOpts, _reserveToken, _virtualBalance)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) Upgrade(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "upgrade")
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeFour *ConverterTypeFourSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Upgrade(&_ConverterTypeFour.TransactOpts)
}

// Upgrade is a paid mutator transaction binding the contract method 0xd55ec697.
//
// Solidity: function upgrade() returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) Upgrade() (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.Upgrade(&_ConverterTypeFour.TransactOpts)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) WithdrawFromToken(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "withdrawFromToken", _token, _to, _amount)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) WithdrawFromToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.WithdrawFromToken(&_ConverterTypeFour.TransactOpts, _token, _to, _amount)
}

// WithdrawFromToken is a paid mutator transaction binding the contract method 0x41a5b33d.
//
// Solidity: function withdrawFromToken(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) WithdrawFromToken(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.WithdrawFromToken(&_ConverterTypeFour.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactor) WithdrawTokens(opts *bind.TransactOpts, _token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.contract.Transact(opts, "withdrawTokens", _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.WithdrawTokens(&_ConverterTypeFour.TransactOpts, _token, _to, _amount)
}

// WithdrawTokens is a paid mutator transaction binding the contract method 0x5e35359e.
//
// Solidity: function withdrawTokens(address _token, address _to, uint256 _amount) returns()
func (_ConverterTypeFour *ConverterTypeFourTransactorSession) WithdrawTokens(_token common.Address, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _ConverterTypeFour.Contract.WithdrawTokens(&_ConverterTypeFour.TransactOpts, _token, _to, _amount)
}

// ConverterTypeFourConversionIterator is returned from FilterConversion and is used to iterate over the raw logs and unpacked data for Conversion events raised by the ConverterTypeFour contract.
type ConverterTypeFourConversionIterator struct {
	Event *ConverterTypeFourConversion // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourConversionIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourConversion)
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
		it.Event = new(ConverterTypeFourConversion)
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
func (it *ConverterTypeFourConversionIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourConversionIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourConversion represents a Conversion event raised by the ConverterTypeFour contract.
type ConverterTypeFourConversion struct {
	FromToken     common.Address
	ToToken       common.Address
	Trader        common.Address
	Amount        *big.Int
	Return        *big.Int
	ConversionFee *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterConversion is a free log retrieval operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterConversion(opts *bind.FilterOpts, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (*ConverterTypeFourConversionIterator, error) {

	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}
	var _traderRule []interface{}
	for _, _traderItem := range _trader {
		_traderRule = append(_traderRule, _traderItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourConversionIterator{contract: _ConverterTypeFour.contract, event: "Conversion", logs: logs, sub: sub}, nil
}

// WatchConversion is a free log subscription operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchConversion(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourConversion, _fromToken []common.Address, _toToken []common.Address, _trader []common.Address) (event.Subscription, error) {

	var _fromTokenRule []interface{}
	for _, _fromTokenItem := range _fromToken {
		_fromTokenRule = append(_fromTokenRule, _fromTokenItem)
	}
	var _toTokenRule []interface{}
	for _, _toTokenItem := range _toToken {
		_toTokenRule = append(_toTokenRule, _toTokenItem)
	}
	var _traderRule []interface{}
	for _, _traderItem := range _trader {
		_traderRule = append(_traderRule, _traderItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "Conversion", _fromTokenRule, _toTokenRule, _traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourConversion)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "Conversion", log); err != nil {
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

// ParseConversion is a log parse operation binding the contract event 0x276856b36cbc45526a0ba64f44611557a2a8b68662c5388e9fe6d72e86e1c8cb.
//
// Solidity: event Conversion(address indexed _fromToken, address indexed _toToken, address indexed _trader, uint256 _amount, uint256 _return, int256 _conversionFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseConversion(log types.Log) (*ConverterTypeFourConversion, error) {
	event := new(ConverterTypeFourConversion)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "Conversion", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourConversionFeeUpdateIterator is returned from FilterConversionFeeUpdate and is used to iterate over the raw logs and unpacked data for ConversionFeeUpdate events raised by the ConverterTypeFour contract.
type ConverterTypeFourConversionFeeUpdateIterator struct {
	Event *ConverterTypeFourConversionFeeUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourConversionFeeUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourConversionFeeUpdate)
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
		it.Event = new(ConverterTypeFourConversionFeeUpdate)
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
func (it *ConverterTypeFourConversionFeeUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourConversionFeeUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourConversionFeeUpdate represents a ConversionFeeUpdate event raised by the ConverterTypeFour contract.
type ConverterTypeFourConversionFeeUpdate struct {
	PrevFee uint32
	NewFee  uint32
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterConversionFeeUpdate is a free log retrieval operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterConversionFeeUpdate(opts *bind.FilterOpts) (*ConverterTypeFourConversionFeeUpdateIterator, error) {

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourConversionFeeUpdateIterator{contract: _ConverterTypeFour.contract, event: "ConversionFeeUpdate", logs: logs, sub: sub}, nil
}

// WatchConversionFeeUpdate is a free log subscription operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchConversionFeeUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourConversionFeeUpdate) (event.Subscription, error) {

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "ConversionFeeUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourConversionFeeUpdate)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
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

// ParseConversionFeeUpdate is a log parse operation binding the contract event 0x81cd2ffb37dd237c0e4e2a3de5265fcf9deb43d3e7801e80db9f1ccfba7ee600.
//
// Solidity: event ConversionFeeUpdate(uint32 _prevFee, uint32 _newFee)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseConversionFeeUpdate(log types.Log) (*ConverterTypeFourConversionFeeUpdate, error) {
	event := new(ConverterTypeFourConversionFeeUpdate)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "ConversionFeeUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourConversionsEnableIterator is returned from FilterConversionsEnable and is used to iterate over the raw logs and unpacked data for ConversionsEnable events raised by the ConverterTypeFour contract.
type ConverterTypeFourConversionsEnableIterator struct {
	Event *ConverterTypeFourConversionsEnable // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourConversionsEnableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourConversionsEnable)
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
		it.Event = new(ConverterTypeFourConversionsEnable)
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
func (it *ConverterTypeFourConversionsEnableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourConversionsEnableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourConversionsEnable represents a ConversionsEnable event raised by the ConverterTypeFour contract.
type ConverterTypeFourConversionsEnable struct {
	ConversionsEnabled bool
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterConversionsEnable is a free log retrieval operation binding the contract event 0xb8e670608a57255ce4f35952b324cba70211a4200a91ce81d26e06d488c1f66b.
//
// Solidity: event ConversionsEnable(bool _conversionsEnabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterConversionsEnable(opts *bind.FilterOpts) (*ConverterTypeFourConversionsEnableIterator, error) {

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "ConversionsEnable")
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourConversionsEnableIterator{contract: _ConverterTypeFour.contract, event: "ConversionsEnable", logs: logs, sub: sub}, nil
}

// WatchConversionsEnable is a free log subscription operation binding the contract event 0xb8e670608a57255ce4f35952b324cba70211a4200a91ce81d26e06d488c1f66b.
//
// Solidity: event ConversionsEnable(bool _conversionsEnabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchConversionsEnable(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourConversionsEnable) (event.Subscription, error) {

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "ConversionsEnable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourConversionsEnable)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "ConversionsEnable", log); err != nil {
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

// ParseConversionsEnable is a log parse operation binding the contract event 0xb8e670608a57255ce4f35952b324cba70211a4200a91ce81d26e06d488c1f66b.
//
// Solidity: event ConversionsEnable(bool _conversionsEnabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseConversionsEnable(log types.Log) (*ConverterTypeFourConversionsEnable, error) {
	event := new(ConverterTypeFourConversionsEnable)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "ConversionsEnable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourManagerUpdateIterator is returned from FilterManagerUpdate and is used to iterate over the raw logs and unpacked data for ManagerUpdate events raised by the ConverterTypeFour contract.
type ConverterTypeFourManagerUpdateIterator struct {
	Event *ConverterTypeFourManagerUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourManagerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourManagerUpdate)
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
		it.Event = new(ConverterTypeFourManagerUpdate)
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
func (it *ConverterTypeFourManagerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourManagerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourManagerUpdate represents a ManagerUpdate event raised by the ConverterTypeFour contract.
type ConverterTypeFourManagerUpdate struct {
	PrevManager common.Address
	NewManager  common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterManagerUpdate is a free log retrieval operation binding the contract event 0xbe4cc281795971a471c980e842627a7f1ea3892ddfce8c5b6357cd2611c19732.
//
// Solidity: event ManagerUpdate(address indexed _prevManager, address indexed _newManager)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterManagerUpdate(opts *bind.FilterOpts, _prevManager []common.Address, _newManager []common.Address) (*ConverterTypeFourManagerUpdateIterator, error) {

	var _prevManagerRule []interface{}
	for _, _prevManagerItem := range _prevManager {
		_prevManagerRule = append(_prevManagerRule, _prevManagerItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "ManagerUpdate", _prevManagerRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourManagerUpdateIterator{contract: _ConverterTypeFour.contract, event: "ManagerUpdate", logs: logs, sub: sub}, nil
}

// WatchManagerUpdate is a free log subscription operation binding the contract event 0xbe4cc281795971a471c980e842627a7f1ea3892ddfce8c5b6357cd2611c19732.
//
// Solidity: event ManagerUpdate(address indexed _prevManager, address indexed _newManager)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchManagerUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourManagerUpdate, _prevManager []common.Address, _newManager []common.Address) (event.Subscription, error) {

	var _prevManagerRule []interface{}
	for _, _prevManagerItem := range _prevManager {
		_prevManagerRule = append(_prevManagerRule, _prevManagerItem)
	}
	var _newManagerRule []interface{}
	for _, _newManagerItem := range _newManager {
		_newManagerRule = append(_newManagerRule, _newManagerItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "ManagerUpdate", _prevManagerRule, _newManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourManagerUpdate)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "ManagerUpdate", log); err != nil {
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

// ParseManagerUpdate is a log parse operation binding the contract event 0xbe4cc281795971a471c980e842627a7f1ea3892ddfce8c5b6357cd2611c19732.
//
// Solidity: event ManagerUpdate(address indexed _prevManager, address indexed _newManager)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseManagerUpdate(log types.Log) (*ConverterTypeFourManagerUpdate, error) {
	event := new(ConverterTypeFourManagerUpdate)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "ManagerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourOwnerUpdateIterator is returned from FilterOwnerUpdate and is used to iterate over the raw logs and unpacked data for OwnerUpdate events raised by the ConverterTypeFour contract.
type ConverterTypeFourOwnerUpdateIterator struct {
	Event *ConverterTypeFourOwnerUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourOwnerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourOwnerUpdate)
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
		it.Event = new(ConverterTypeFourOwnerUpdate)
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
func (it *ConverterTypeFourOwnerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourOwnerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourOwnerUpdate represents a OwnerUpdate event raised by the ConverterTypeFour contract.
type ConverterTypeFourOwnerUpdate struct {
	PrevOwner common.Address
	NewOwner  common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterOwnerUpdate is a free log retrieval operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterOwnerUpdate(opts *bind.FilterOpts, _prevOwner []common.Address, _newOwner []common.Address) (*ConverterTypeFourOwnerUpdateIterator, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourOwnerUpdateIterator{contract: _ConverterTypeFour.contract, event: "OwnerUpdate", logs: logs, sub: sub}, nil
}

// WatchOwnerUpdate is a free log subscription operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchOwnerUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourOwnerUpdate, _prevOwner []common.Address, _newOwner []common.Address) (event.Subscription, error) {

	var _prevOwnerRule []interface{}
	for _, _prevOwnerItem := range _prevOwner {
		_prevOwnerRule = append(_prevOwnerRule, _prevOwnerItem)
	}
	var _newOwnerRule []interface{}
	for _, _newOwnerItem := range _newOwner {
		_newOwnerRule = append(_newOwnerRule, _newOwnerItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "OwnerUpdate", _prevOwnerRule, _newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourOwnerUpdate)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
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

// ParseOwnerUpdate is a log parse operation binding the contract event 0x343765429aea5a34b3ff6a3785a98a5abb2597aca87bfbb58632c173d585373a.
//
// Solidity: event OwnerUpdate(address indexed _prevOwner, address indexed _newOwner)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseOwnerUpdate(log types.Log) (*ConverterTypeFourOwnerUpdate, error) {
	event := new(ConverterTypeFourOwnerUpdate)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "OwnerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourPriceDataUpdateIterator is returned from FilterPriceDataUpdate and is used to iterate over the raw logs and unpacked data for PriceDataUpdate events raised by the ConverterTypeFour contract.
type ConverterTypeFourPriceDataUpdateIterator struct {
	Event *ConverterTypeFourPriceDataUpdate // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourPriceDataUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourPriceDataUpdate)
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
		it.Event = new(ConverterTypeFourPriceDataUpdate)
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
func (it *ConverterTypeFourPriceDataUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourPriceDataUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourPriceDataUpdate represents a PriceDataUpdate event raised by the ConverterTypeFour contract.
type ConverterTypeFourPriceDataUpdate struct {
	ConnectorToken   common.Address
	TokenSupply      *big.Int
	ConnectorBalance *big.Int
	ConnectorWeight  uint32
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPriceDataUpdate is a free log retrieval operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterPriceDataUpdate(opts *bind.FilterOpts, _connectorToken []common.Address) (*ConverterTypeFourPriceDataUpdateIterator, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourPriceDataUpdateIterator{contract: _ConverterTypeFour.contract, event: "PriceDataUpdate", logs: logs, sub: sub}, nil
}

// WatchPriceDataUpdate is a free log subscription operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchPriceDataUpdate(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourPriceDataUpdate, _connectorToken []common.Address) (event.Subscription, error) {

	var _connectorTokenRule []interface{}
	for _, _connectorTokenItem := range _connectorToken {
		_connectorTokenRule = append(_connectorTokenRule, _connectorTokenItem)
	}

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "PriceDataUpdate", _connectorTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourPriceDataUpdate)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
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

// ParsePriceDataUpdate is a log parse operation binding the contract event 0x8a6a7f53b3c8fa1dc4b83e3f1be668c1b251ff8d44cdcb83eb3acec3fec6a788.
//
// Solidity: event PriceDataUpdate(address indexed _connectorToken, uint256 _tokenSupply, uint256 _connectorBalance, uint32 _connectorWeight)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParsePriceDataUpdate(log types.Log) (*ConverterTypeFourPriceDataUpdate, error) {
	event := new(ConverterTypeFourPriceDataUpdate)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "PriceDataUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ConverterTypeFourVirtualBalancesEnableIterator is returned from FilterVirtualBalancesEnable and is used to iterate over the raw logs and unpacked data for VirtualBalancesEnable events raised by the ConverterTypeFour contract.
type ConverterTypeFourVirtualBalancesEnableIterator struct {
	Event *ConverterTypeFourVirtualBalancesEnable // Event containing the contract specifics and raw log

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
func (it *ConverterTypeFourVirtualBalancesEnableIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ConverterTypeFourVirtualBalancesEnable)
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
		it.Event = new(ConverterTypeFourVirtualBalancesEnable)
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
func (it *ConverterTypeFourVirtualBalancesEnableIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ConverterTypeFourVirtualBalancesEnableIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ConverterTypeFourVirtualBalancesEnable represents a VirtualBalancesEnable event raised by the ConverterTypeFour contract.
type ConverterTypeFourVirtualBalancesEnable struct {
	Enabled bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterVirtualBalancesEnable is a free log retrieval operation binding the contract event 0x64622fbd54039f76d87a876ecaea9bdb6b9b493d7a35ca38ae82b53dcddbe2e4.
//
// Solidity: event VirtualBalancesEnable(bool _enabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) FilterVirtualBalancesEnable(opts *bind.FilterOpts) (*ConverterTypeFourVirtualBalancesEnableIterator, error) {

	logs, sub, err := _ConverterTypeFour.contract.FilterLogs(opts, "VirtualBalancesEnable")
	if err != nil {
		return nil, err
	}
	return &ConverterTypeFourVirtualBalancesEnableIterator{contract: _ConverterTypeFour.contract, event: "VirtualBalancesEnable", logs: logs, sub: sub}, nil
}

// WatchVirtualBalancesEnable is a free log subscription operation binding the contract event 0x64622fbd54039f76d87a876ecaea9bdb6b9b493d7a35ca38ae82b53dcddbe2e4.
//
// Solidity: event VirtualBalancesEnable(bool _enabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) WatchVirtualBalancesEnable(opts *bind.WatchOpts, sink chan<- *ConverterTypeFourVirtualBalancesEnable) (event.Subscription, error) {

	logs, sub, err := _ConverterTypeFour.contract.WatchLogs(opts, "VirtualBalancesEnable")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ConverterTypeFourVirtualBalancesEnable)
				if err := _ConverterTypeFour.contract.UnpackLog(event, "VirtualBalancesEnable", log); err != nil {
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

// ParseVirtualBalancesEnable is a log parse operation binding the contract event 0x64622fbd54039f76d87a876ecaea9bdb6b9b493d7a35ca38ae82b53dcddbe2e4.
//
// Solidity: event VirtualBalancesEnable(bool _enabled)
func (_ConverterTypeFour *ConverterTypeFourFilterer) ParseVirtualBalancesEnable(log types.Log) (*ConverterTypeFourVirtualBalancesEnable, error) {
	event := new(ConverterTypeFourVirtualBalancesEnable)
	if err := _ConverterTypeFour.contract.UnpackLog(event, "VirtualBalancesEnable", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
