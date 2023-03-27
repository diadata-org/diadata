// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package blurv1

import (
	"errors"
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
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// Execution is an auto generated low-level Go binding around an user-defined struct.
type Execution struct {
	Sell Input
	Buy  Input
}

// Fee is an auto generated low-level Go binding around an user-defined struct.
type Fee struct {
	Rate      uint16
	Recipient common.Address
}

// Input is an auto generated low-level Go binding around an user-defined struct.
type Input struct {
	Order            Order
	V                uint8
	R                [32]byte
	S                [32]byte
	ExtraSignature   []byte
	SignatureVersion uint8
	BlockNumber      *big.Int
}

// Order is an auto generated low-level Go binding around an user-defined struct.
type Order struct {
	Trader         common.Address
	Side           uint8
	MatchingPolicy common.Address
	Collection     common.Address
	TokenId        *big.Int
	Amount         *big.Int
	PaymentToken   common.Address
	Price          *big.Int
	ListingTime    *big.Int
	ExpirationTime *big.Int
	Fees           []Fee
	Salt           *big.Int
	ExtraParams    []byte
}

// Blurv1MetaData contains all meta data concerning the Blurv1 contract.
var Blurv1MetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"previousAdmin\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAdmin\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"beacon\",\"type\":\"address\"}],\"name\":\"BeaconUpgraded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Closed\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"blockRange\",\"type\":\"uint256\"}],\"name\":\"NewBlockRange\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIExecutionDelegate\",\"name\":\"executionDelegate\",\"type\":\"address\"}],\"name\":\"NewExecutionDelegate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"feeRate\",\"type\":\"uint256\"}],\"name\":\"NewFeeRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"NewFeeRecipient\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"governor\",\"type\":\"address\"}],\"name\":\"NewGovernor\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oracle\",\"type\":\"address\"}],\"name\":\"NewOracle\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIPolicyManager\",\"name\":\"policyManager\",\"type\":\"address\"}],\"name\":\"NewPolicyManager\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newNonce\",\"type\":\"uint256\"}],\"name\":\"NonceIncremented\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Opened\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"hash\",\"type\":\"bytes32\"}],\"name\":\"OrderCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"maker\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"taker\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"sell\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"sellHash\",\"type\":\"bytes32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"indexed\":false,\"internalType\":\"structOrder\",\"name\":\"buy\",\"type\":\"tuple\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"buyHash\",\"type\":\"bytes32\"}],\"name\":\"OrdersMatched\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"implementation\",\"type\":\"address\"}],\"name\":\"Upgraded\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"FEE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"INVERSE_BASIS_POINT\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"NAME\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORACLE_ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ORDER_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"POOL\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ROOT_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VERSION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"WETH\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"_execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"blockRange\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"internalType\":\"structExecution[]\",\"name\":\"executions\",\"type\":\"tuple[]\"}],\"name\":\"bulkExecute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"}],\"name\":\"cancelOrder\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder[]\",\"name\":\"orders\",\"type\":\"tuple[]\"}],\"name\":\"cancelOrders\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"cancelledOrFilled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"close\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"sell\",\"type\":\"tuple\"},{\"components\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"enumSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"matchingPolicy\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collection\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"paymentToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"listingTime\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expirationTime\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint16\",\"name\":\"rate\",\"type\":\"uint16\"},{\"internalType\":\"addresspayable\",\"name\":\"recipient\",\"type\":\"address\"}],\"internalType\":\"structFee[]\",\"name\":\"fees\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"salt\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"extraParams\",\"type\":\"bytes\"}],\"internalType\":\"structOrder\",\"name\":\"order\",\"type\":\"tuple\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"extraSignature\",\"type\":\"bytes\"},{\"internalType\":\"enumSignatureVersion\",\"name\":\"signatureVersion\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"internalType\":\"structInput\",\"name\":\"buy\",\"type\":\"tuple\"}],\"name\":\"execute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionDelegate\",\"outputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"feeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"governor\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"incrementNonce\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"},{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInternal\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isOpen\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"nonces\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"open\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"policyManager\",\"outputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proxiableUUID\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"remainingETH\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_blockRange\",\"type\":\"uint256\"}],\"name\":\"setBlockRange\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIExecutionDelegate\",\"name\":\"_executionDelegate\",\"type\":\"address\"}],\"name\":\"setExecutionDelegate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_feeRate\",\"type\":\"uint256\"}],\"name\":\"setFeeRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_governor\",\"type\":\"address\"}],\"name\":\"setGovernor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIPolicyManager\",\"name\":\"_policyManager\",\"type\":\"address\"}],\"name\":\"setPolicyManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"}],\"name\":\"upgradeTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newImplementation\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"upgradeToAndCall\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"}]",
}

// Blurv1ABI is the input ABI used to generate the binding from.
// Deprecated: Use Blurv1MetaData.ABI instead.
var Blurv1ABI = Blurv1MetaData.ABI

// Blurv1 is an auto generated Go binding around an Ethereum contract.
type Blurv1 struct {
	Blurv1Caller     // Read-only binding to the contract
	Blurv1Transactor // Write-only binding to the contract
	Blurv1Filterer   // Log filterer for contract events
}

// Blurv1Caller is an auto generated read-only Go binding around an Ethereum contract.
type Blurv1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Blurv1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Blurv1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Blurv1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Blurv1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Blurv1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Blurv1Session struct {
	Contract     *Blurv1           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Blurv1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Blurv1CallerSession struct {
	Contract *Blurv1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Blurv1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Blurv1TransactorSession struct {
	Contract     *Blurv1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Blurv1Raw is an auto generated low-level Go binding around an Ethereum contract.
type Blurv1Raw struct {
	Contract *Blurv1 // Generic contract binding to access the raw methods on
}

// Blurv1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Blurv1CallerRaw struct {
	Contract *Blurv1Caller // Generic read-only contract binding to access the raw methods on
}

// Blurv1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Blurv1TransactorRaw struct {
	Contract *Blurv1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewBlurv1 creates a new instance of Blurv1, bound to a specific deployed contract.
func NewBlurv1(address common.Address, backend bind.ContractBackend) (*Blurv1, error) {
	contract, err := bindBlurv1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Blurv1{Blurv1Caller: Blurv1Caller{contract: contract}, Blurv1Transactor: Blurv1Transactor{contract: contract}, Blurv1Filterer: Blurv1Filterer{contract: contract}}, nil
}

// NewBlurv1Caller creates a new read-only instance of Blurv1, bound to a specific deployed contract.
func NewBlurv1Caller(address common.Address, caller bind.ContractCaller) (*Blurv1Caller, error) {
	contract, err := bindBlurv1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Blurv1Caller{contract: contract}, nil
}

// NewBlurv1Transactor creates a new write-only instance of Blurv1, bound to a specific deployed contract.
func NewBlurv1Transactor(address common.Address, transactor bind.ContractTransactor) (*Blurv1Transactor, error) {
	contract, err := bindBlurv1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Blurv1Transactor{contract: contract}, nil
}

// NewBlurv1Filterer creates a new log filterer instance of Blurv1, bound to a specific deployed contract.
func NewBlurv1Filterer(address common.Address, filterer bind.ContractFilterer) (*Blurv1Filterer, error) {
	contract, err := bindBlurv1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Blurv1Filterer{contract: contract}, nil
}

// bindBlurv1 binds a generic wrapper to an already deployed contract.
func bindBlurv1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := Blurv1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blurv1 *Blurv1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blurv1.Contract.Blurv1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blurv1 *Blurv1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.Contract.Blurv1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blurv1 *Blurv1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blurv1.Contract.Blurv1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Blurv1 *Blurv1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Blurv1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Blurv1 *Blurv1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Blurv1 *Blurv1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Blurv1.Contract.contract.Transact(opts, method, params...)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Caller) FEETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "FEE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Session) FEETYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.FEETYPEHASH(&_Blurv1.CallOpts)
}

// FEETYPEHASH is a free data retrieval call binding the contract method 0x4832ede1.
//
// Solidity: function FEE_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1CallerSession) FEETYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.FEETYPEHASH(&_Blurv1.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blurv1 *Blurv1Caller) INVERSEBASISPOINT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "INVERSE_BASIS_POINT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blurv1 *Blurv1Session) INVERSEBASISPOINT() (*big.Int, error) {
	return _Blurv1.Contract.INVERSEBASISPOINT(&_Blurv1.CallOpts)
}

// INVERSEBASISPOINT is a free data retrieval call binding the contract method 0xcae6047f.
//
// Solidity: function INVERSE_BASIS_POINT() view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) INVERSEBASISPOINT() (*big.Int, error) {
	return _Blurv1.Contract.INVERSEBASISPOINT(&_Blurv1.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blurv1 *Blurv1Caller) NAME(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "NAME")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blurv1 *Blurv1Session) NAME() (string, error) {
	return _Blurv1.Contract.NAME(&_Blurv1.CallOpts)
}

// NAME is a free data retrieval call binding the contract method 0xa3f4df7e.
//
// Solidity: function NAME() view returns(string)
func (_Blurv1 *Blurv1CallerSession) NAME() (string, error) {
	return _Blurv1.Contract.NAME(&_Blurv1.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Caller) ORACLEORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "ORACLE_ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Session) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ORACLEORDERTYPEHASH(&_Blurv1.CallOpts)
}

// ORACLEORDERTYPEHASH is a free data retrieval call binding the contract method 0x1d97c9bb.
//
// Solidity: function ORACLE_ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1CallerSession) ORACLEORDERTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ORACLEORDERTYPEHASH(&_Blurv1.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Caller) ORDERTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "ORDER_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Session) ORDERTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ORDERTYPEHASH(&_Blurv1.CallOpts)
}

// ORDERTYPEHASH is a free data retrieval call binding the contract method 0xf973a209.
//
// Solidity: function ORDER_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1CallerSession) ORDERTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ORDERTYPEHASH(&_Blurv1.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blurv1 *Blurv1Caller) POOL(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "POOL")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blurv1 *Blurv1Session) POOL() (common.Address, error) {
	return _Blurv1.Contract.POOL(&_Blurv1.CallOpts)
}

// POOL is a free data retrieval call binding the contract method 0x7535d246.
//
// Solidity: function POOL() view returns(address)
func (_Blurv1 *Blurv1CallerSession) POOL() (common.Address, error) {
	return _Blurv1.Contract.POOL(&_Blurv1.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Caller) ROOTTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "ROOT_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1Session) ROOTTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ROOTTYPEHASH(&_Blurv1.CallOpts)
}

// ROOTTYPEHASH is a free data retrieval call binding the contract method 0x31e6d0fe.
//
// Solidity: function ROOT_TYPEHASH() view returns(bytes32)
func (_Blurv1 *Blurv1CallerSession) ROOTTYPEHASH() ([32]byte, error) {
	return _Blurv1.Contract.ROOTTYPEHASH(&_Blurv1.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blurv1 *Blurv1Caller) VERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blurv1 *Blurv1Session) VERSION() (string, error) {
	return _Blurv1.Contract.VERSION(&_Blurv1.CallOpts)
}

// VERSION is a free data retrieval call binding the contract method 0xffa1ad74.
//
// Solidity: function VERSION() view returns(string)
func (_Blurv1 *Blurv1CallerSession) VERSION() (string, error) {
	return _Blurv1.Contract.VERSION(&_Blurv1.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blurv1 *Blurv1Caller) WETH(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "WETH")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blurv1 *Blurv1Session) WETH() (common.Address, error) {
	return _Blurv1.Contract.WETH(&_Blurv1.CallOpts)
}

// WETH is a free data retrieval call binding the contract method 0xad5c4648.
//
// Solidity: function WETH() view returns(address)
func (_Blurv1 *Blurv1CallerSession) WETH() (common.Address, error) {
	return _Blurv1.Contract.WETH(&_Blurv1.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blurv1 *Blurv1Caller) BlockRange(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "blockRange")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blurv1 *Blurv1Session) BlockRange() (*big.Int, error) {
	return _Blurv1.Contract.BlockRange(&_Blurv1.CallOpts)
}

// BlockRange is a free data retrieval call binding the contract method 0xa4b2c674.
//
// Solidity: function blockRange() view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) BlockRange() (*big.Int, error) {
	return _Blurv1.Contract.BlockRange(&_Blurv1.CallOpts)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blurv1 *Blurv1Caller) CancelledOrFilled(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "cancelledOrFilled", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blurv1 *Blurv1Session) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Blurv1.Contract.CancelledOrFilled(&_Blurv1.CallOpts, arg0)
}

// CancelledOrFilled is a free data retrieval call binding the contract method 0x5511f319.
//
// Solidity: function cancelledOrFilled(bytes32 ) view returns(bool)
func (_Blurv1 *Blurv1CallerSession) CancelledOrFilled(arg0 [32]byte) (bool, error) {
	return _Blurv1.Contract.CancelledOrFilled(&_Blurv1.CallOpts, arg0)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blurv1 *Blurv1Caller) ExecutionDelegate(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "executionDelegate")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blurv1 *Blurv1Session) ExecutionDelegate() (common.Address, error) {
	return _Blurv1.Contract.ExecutionDelegate(&_Blurv1.CallOpts)
}

// ExecutionDelegate is a free data retrieval call binding the contract method 0x986c9b20.
//
// Solidity: function executionDelegate() view returns(address)
func (_Blurv1 *Blurv1CallerSession) ExecutionDelegate() (common.Address, error) {
	return _Blurv1.Contract.ExecutionDelegate(&_Blurv1.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blurv1 *Blurv1Caller) FeeRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "feeRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blurv1 *Blurv1Session) FeeRate() (*big.Int, error) {
	return _Blurv1.Contract.FeeRate(&_Blurv1.CallOpts)
}

// FeeRate is a free data retrieval call binding the contract method 0x978bbdb9.
//
// Solidity: function feeRate() view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) FeeRate() (*big.Int, error) {
	return _Blurv1.Contract.FeeRate(&_Blurv1.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blurv1 *Blurv1Caller) FeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "feeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blurv1 *Blurv1Session) FeeRecipient() (common.Address, error) {
	return _Blurv1.Contract.FeeRecipient(&_Blurv1.CallOpts)
}

// FeeRecipient is a free data retrieval call binding the contract method 0x46904840.
//
// Solidity: function feeRecipient() view returns(address)
func (_Blurv1 *Blurv1CallerSession) FeeRecipient() (common.Address, error) {
	return _Blurv1.Contract.FeeRecipient(&_Blurv1.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blurv1 *Blurv1Caller) Governor(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "governor")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blurv1 *Blurv1Session) Governor() (common.Address, error) {
	return _Blurv1.Contract.Governor(&_Blurv1.CallOpts)
}

// Governor is a free data retrieval call binding the contract method 0x0c340a24.
//
// Solidity: function governor() view returns(address)
func (_Blurv1 *Blurv1CallerSession) Governor() (common.Address, error) {
	return _Blurv1.Contract.Governor(&_Blurv1.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blurv1 *Blurv1Caller) IsInternal(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "isInternal")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blurv1 *Blurv1Session) IsInternal() (bool, error) {
	return _Blurv1.Contract.IsInternal(&_Blurv1.CallOpts)
}

// IsInternal is a free data retrieval call binding the contract method 0x16e29d71.
//
// Solidity: function isInternal() view returns(bool)
func (_Blurv1 *Blurv1CallerSession) IsInternal() (bool, error) {
	return _Blurv1.Contract.IsInternal(&_Blurv1.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blurv1 *Blurv1Caller) IsOpen(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "isOpen")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blurv1 *Blurv1Session) IsOpen() (*big.Int, error) {
	return _Blurv1.Contract.IsOpen(&_Blurv1.CallOpts)
}

// IsOpen is a free data retrieval call binding the contract method 0x47535d7b.
//
// Solidity: function isOpen() view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) IsOpen() (*big.Int, error) {
	return _Blurv1.Contract.IsOpen(&_Blurv1.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blurv1 *Blurv1Caller) Nonces(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "nonces", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blurv1 *Blurv1Session) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blurv1.Contract.Nonces(&_Blurv1.CallOpts, arg0)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address ) view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) Nonces(arg0 common.Address) (*big.Int, error) {
	return _Blurv1.Contract.Nonces(&_Blurv1.CallOpts, arg0)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blurv1 *Blurv1Caller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blurv1 *Blurv1Session) Oracle() (common.Address, error) {
	return _Blurv1.Contract.Oracle(&_Blurv1.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Blurv1 *Blurv1CallerSession) Oracle() (common.Address, error) {
	return _Blurv1.Contract.Oracle(&_Blurv1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blurv1 *Blurv1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blurv1 *Blurv1Session) Owner() (common.Address, error) {
	return _Blurv1.Contract.Owner(&_Blurv1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Blurv1 *Blurv1CallerSession) Owner() (common.Address, error) {
	return _Blurv1.Contract.Owner(&_Blurv1.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blurv1 *Blurv1Caller) PolicyManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "policyManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blurv1 *Blurv1Session) PolicyManager() (common.Address, error) {
	return _Blurv1.Contract.PolicyManager(&_Blurv1.CallOpts)
}

// PolicyManager is a free data retrieval call binding the contract method 0xab3dbf3b.
//
// Solidity: function policyManager() view returns(address)
func (_Blurv1 *Blurv1CallerSession) PolicyManager() (common.Address, error) {
	return _Blurv1.Contract.PolicyManager(&_Blurv1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blurv1 *Blurv1Caller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blurv1 *Blurv1Session) ProxiableUUID() ([32]byte, error) {
	return _Blurv1.Contract.ProxiableUUID(&_Blurv1.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_Blurv1 *Blurv1CallerSession) ProxiableUUID() ([32]byte, error) {
	return _Blurv1.Contract.ProxiableUUID(&_Blurv1.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blurv1 *Blurv1Caller) RemainingETH(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Blurv1.contract.Call(opts, &out, "remainingETH")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blurv1 *Blurv1Session) RemainingETH() (*big.Int, error) {
	return _Blurv1.Contract.RemainingETH(&_Blurv1.CallOpts)
}

// RemainingETH is a free data retrieval call binding the contract method 0x2c7acf8c.
//
// Solidity: function remainingETH() view returns(uint256)
func (_Blurv1 *Blurv1CallerSession) RemainingETH() (*big.Int, error) {
	return _Blurv1.Contract.RemainingETH(&_Blurv1.CallOpts)
}

// Execution is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1Transactor) Execution(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "_execute", sell, buy)
}

// Execution is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1Session) Execution(sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.Contract.Execution(&_Blurv1.TransactOpts, sell, buy)
}

// Execution is a paid mutator transaction binding the contract method 0xe04d94ae.
//
// Solidity: function _execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1TransactorSession) Execution(sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.Contract.Execution(&_Blurv1.TransactOpts, sell, buy)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blurv1 *Blurv1Transactor) BulkExecute(opts *bind.TransactOpts, executions []Execution) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "bulkExecute", executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blurv1 *Blurv1Session) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Blurv1.Contract.BulkExecute(&_Blurv1.TransactOpts, executions)
}

// BulkExecute is a paid mutator transaction binding the contract method 0xb3be57f8.
//
// Solidity: function bulkExecute((((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256),((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256))[] executions) payable returns()
func (_Blurv1 *Blurv1TransactorSession) BulkExecute(executions []Execution) (*types.Transaction, error) {
	return _Blurv1.Contract.BulkExecute(&_Blurv1.TransactOpts, executions)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blurv1 *Blurv1Transactor) CancelOrder(opts *bind.TransactOpts, order Order) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "cancelOrder", order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blurv1 *Blurv1Session) CancelOrder(order Order) (*types.Transaction, error) {
	return _Blurv1.Contract.CancelOrder(&_Blurv1.TransactOpts, order)
}

// CancelOrder is a paid mutator transaction binding the contract method 0xf4acd740.
//
// Solidity: function cancelOrder((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) order) returns()
func (_Blurv1 *Blurv1TransactorSession) CancelOrder(order Order) (*types.Transaction, error) {
	return _Blurv1.Contract.CancelOrder(&_Blurv1.TransactOpts, order)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blurv1 *Blurv1Transactor) CancelOrders(opts *bind.TransactOpts, orders []Order) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "cancelOrders", orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blurv1 *Blurv1Session) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Blurv1.Contract.CancelOrders(&_Blurv1.TransactOpts, orders)
}

// CancelOrders is a paid mutator transaction binding the contract method 0xab7e8cba.
//
// Solidity: function cancelOrders((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes)[] orders) returns()
func (_Blurv1 *Blurv1TransactorSession) CancelOrders(orders []Order) (*types.Transaction, error) {
	return _Blurv1.Contract.CancelOrders(&_Blurv1.TransactOpts, orders)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blurv1 *Blurv1Transactor) Close(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "close")
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blurv1 *Blurv1Session) Close() (*types.Transaction, error) {
	return _Blurv1.Contract.Close(&_Blurv1.TransactOpts)
}

// Close is a paid mutator transaction binding the contract method 0x43d726d6.
//
// Solidity: function close() returns()
func (_Blurv1 *Blurv1TransactorSession) Close() (*types.Transaction, error) {
	return _Blurv1.Contract.Close(&_Blurv1.TransactOpts)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1Transactor) Execute(opts *bind.TransactOpts, sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "execute", sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1Session) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.Contract.Execute(&_Blurv1.TransactOpts, sell, buy)
}

// Execute is a paid mutator transaction binding the contract method 0x9a1fc3a7.
//
// Solidity: function execute(((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) sell, ((address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes),uint8,bytes32,bytes32,bytes,uint8,uint256) buy) payable returns()
func (_Blurv1 *Blurv1TransactorSession) Execute(sell Input, buy Input) (*types.Transaction, error) {
	return _Blurv1.Contract.Execute(&_Blurv1.TransactOpts, sell, buy)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blurv1 *Blurv1Transactor) IncrementNonce(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "incrementNonce")
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blurv1 *Blurv1Session) IncrementNonce() (*types.Transaction, error) {
	return _Blurv1.Contract.IncrementNonce(&_Blurv1.TransactOpts)
}

// IncrementNonce is a paid mutator transaction binding the contract method 0x627cdcb9.
//
// Solidity: function incrementNonce() returns()
func (_Blurv1 *Blurv1TransactorSession) IncrementNonce() (*types.Transaction, error) {
	return _Blurv1.Contract.IncrementNonce(&_Blurv1.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blurv1 *Blurv1Transactor) Initialize(opts *bind.TransactOpts, _executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "initialize", _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blurv1 *Blurv1Session) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.Initialize(&_Blurv1.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Initialize is a paid mutator transaction binding the contract method 0xcf756fdf.
//
// Solidity: function initialize(address _executionDelegate, address _policyManager, address _oracle, uint256 _blockRange) returns()
func (_Blurv1 *Blurv1TransactorSession) Initialize(_executionDelegate common.Address, _policyManager common.Address, _oracle common.Address, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.Initialize(&_Blurv1.TransactOpts, _executionDelegate, _policyManager, _oracle, _blockRange)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blurv1 *Blurv1Transactor) Open(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "open")
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blurv1 *Blurv1Session) Open() (*types.Transaction, error) {
	return _Blurv1.Contract.Open(&_Blurv1.TransactOpts)
}

// Open is a paid mutator transaction binding the contract method 0xfcfff16f.
//
// Solidity: function open() returns()
func (_Blurv1 *Blurv1TransactorSession) Open() (*types.Transaction, error) {
	return _Blurv1.Contract.Open(&_Blurv1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blurv1 *Blurv1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blurv1 *Blurv1Session) RenounceOwnership() (*types.Transaction, error) {
	return _Blurv1.Contract.RenounceOwnership(&_Blurv1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Blurv1 *Blurv1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Blurv1.Contract.RenounceOwnership(&_Blurv1.TransactOpts)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blurv1 *Blurv1Transactor) SetBlockRange(opts *bind.TransactOpts, _blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setBlockRange", _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blurv1 *Blurv1Session) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.SetBlockRange(&_Blurv1.TransactOpts, _blockRange)
}

// SetBlockRange is a paid mutator transaction binding the contract method 0x6992aa36.
//
// Solidity: function setBlockRange(uint256 _blockRange) returns()
func (_Blurv1 *Blurv1TransactorSession) SetBlockRange(_blockRange *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.SetBlockRange(&_Blurv1.TransactOpts, _blockRange)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blurv1 *Blurv1Transactor) SetExecutionDelegate(opts *bind.TransactOpts, _executionDelegate common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setExecutionDelegate", _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blurv1 *Blurv1Session) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetExecutionDelegate(&_Blurv1.TransactOpts, _executionDelegate)
}

// SetExecutionDelegate is a paid mutator transaction binding the contract method 0x037c9be2.
//
// Solidity: function setExecutionDelegate(address _executionDelegate) returns()
func (_Blurv1 *Blurv1TransactorSession) SetExecutionDelegate(_executionDelegate common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetExecutionDelegate(&_Blurv1.TransactOpts, _executionDelegate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blurv1 *Blurv1Transactor) SetFeeRate(opts *bind.TransactOpts, _feeRate *big.Int) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setFeeRate", _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blurv1 *Blurv1Session) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.SetFeeRate(&_Blurv1.TransactOpts, _feeRate)
}

// SetFeeRate is a paid mutator transaction binding the contract method 0x45596e2e.
//
// Solidity: function setFeeRate(uint256 _feeRate) returns()
func (_Blurv1 *Blurv1TransactorSession) SetFeeRate(_feeRate *big.Int) (*types.Transaction, error) {
	return _Blurv1.Contract.SetFeeRate(&_Blurv1.TransactOpts, _feeRate)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blurv1 *Blurv1Transactor) SetFeeRecipient(opts *bind.TransactOpts, _feeRecipient common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setFeeRecipient", _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blurv1 *Blurv1Session) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetFeeRecipient(&_Blurv1.TransactOpts, _feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address _feeRecipient) returns()
func (_Blurv1 *Blurv1TransactorSession) SetFeeRecipient(_feeRecipient common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetFeeRecipient(&_Blurv1.TransactOpts, _feeRecipient)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blurv1 *Blurv1Transactor) SetGovernor(opts *bind.TransactOpts, _governor common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setGovernor", _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blurv1 *Blurv1Session) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetGovernor(&_Blurv1.TransactOpts, _governor)
}

// SetGovernor is a paid mutator transaction binding the contract method 0xc42cf535.
//
// Solidity: function setGovernor(address _governor) returns()
func (_Blurv1 *Blurv1TransactorSession) SetGovernor(_governor common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetGovernor(&_Blurv1.TransactOpts, _governor)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blurv1 *Blurv1Transactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blurv1 *Blurv1Session) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetOracle(&_Blurv1.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Blurv1 *Blurv1TransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetOracle(&_Blurv1.TransactOpts, _oracle)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blurv1 *Blurv1Transactor) SetPolicyManager(opts *bind.TransactOpts, _policyManager common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "setPolicyManager", _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blurv1 *Blurv1Session) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetPolicyManager(&_Blurv1.TransactOpts, _policyManager)
}

// SetPolicyManager is a paid mutator transaction binding the contract method 0xadde41e1.
//
// Solidity: function setPolicyManager(address _policyManager) returns()
func (_Blurv1 *Blurv1TransactorSession) SetPolicyManager(_policyManager common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.SetPolicyManager(&_Blurv1.TransactOpts, _policyManager)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blurv1 *Blurv1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blurv1 *Blurv1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.TransferOwnership(&_Blurv1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Blurv1 *Blurv1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.TransferOwnership(&_Blurv1.TransactOpts, newOwner)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blurv1 *Blurv1Transactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blurv1 *Blurv1Session) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.UpgradeTo(&_Blurv1.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_Blurv1 *Blurv1TransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _Blurv1.Contract.UpgradeTo(&_Blurv1.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blurv1 *Blurv1Transactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blurv1.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blurv1 *Blurv1Session) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blurv1.Contract.UpgradeToAndCall(&_Blurv1.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_Blurv1 *Blurv1TransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _Blurv1.Contract.UpgradeToAndCall(&_Blurv1.TransactOpts, newImplementation, data)
}

// Blurv1AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Blurv1 contract.
type Blurv1AdminChangedIterator struct {
	Event *Blurv1AdminChanged // Event containing the contract specifics and raw log

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
func (it *Blurv1AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1AdminChanged)
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
		it.Event = new(Blurv1AdminChanged)
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
func (it *Blurv1AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1AdminChanged represents a AdminChanged event raised by the Blurv1 contract.
type Blurv1AdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blurv1 *Blurv1Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*Blurv1AdminChangedIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Blurv1AdminChangedIterator{contract: _Blurv1.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blurv1 *Blurv1Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Blurv1AdminChanged) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1AdminChanged)
				if err := _Blurv1.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_Blurv1 *Blurv1Filterer) ParseAdminChanged(log types.Log) (*Blurv1AdminChanged, error) {
	event := new(Blurv1AdminChanged)
	if err := _Blurv1.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1BeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the Blurv1 contract.
type Blurv1BeaconUpgradedIterator struct {
	Event *Blurv1BeaconUpgraded // Event containing the contract specifics and raw log

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
func (it *Blurv1BeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1BeaconUpgraded)
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
		it.Event = new(Blurv1BeaconUpgraded)
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
func (it *Blurv1BeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1BeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1BeaconUpgraded represents a BeaconUpgraded event raised by the Blurv1 contract.
type Blurv1BeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blurv1 *Blurv1Filterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*Blurv1BeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1BeaconUpgradedIterator{contract: _Blurv1.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blurv1 *Blurv1Filterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *Blurv1BeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1BeaconUpgraded)
				if err := _Blurv1.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
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

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_Blurv1 *Blurv1Filterer) ParseBeaconUpgraded(log types.Log) (*Blurv1BeaconUpgraded, error) {
	event := new(Blurv1BeaconUpgraded)
	if err := _Blurv1.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1ClosedIterator is returned from FilterClosed and is used to iterate over the raw logs and unpacked data for Closed events raised by the Blurv1 contract.
type Blurv1ClosedIterator struct {
	Event *Blurv1Closed // Event containing the contract specifics and raw log

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
func (it *Blurv1ClosedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1Closed)
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
		it.Event = new(Blurv1Closed)
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
func (it *Blurv1ClosedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1ClosedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1Closed represents a Closed event raised by the Blurv1 contract.
type Blurv1Closed struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClosed is a free log retrieval operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blurv1 *Blurv1Filterer) FilterClosed(opts *bind.FilterOpts) (*Blurv1ClosedIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return &Blurv1ClosedIterator{contract: _Blurv1.contract, event: "Closed", logs: logs, sub: sub}, nil
}

// WatchClosed is a free log subscription operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blurv1 *Blurv1Filterer) WatchClosed(opts *bind.WatchOpts, sink chan<- *Blurv1Closed) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "Closed")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1Closed)
				if err := _Blurv1.contract.UnpackLog(event, "Closed", log); err != nil {
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

// ParseClosed is a log parse operation binding the contract event 0x1cdde67b72a90f19919ac732a437ac2f7a10fc128d28c2a6e525d89ce5cd9d3a.
//
// Solidity: event Closed()
func (_Blurv1 *Blurv1Filterer) ParseClosed(log types.Log) (*Blurv1Closed, error) {
	event := new(Blurv1Closed)
	if err := _Blurv1.contract.UnpackLog(event, "Closed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Blurv1 contract.
type Blurv1InitializedIterator struct {
	Event *Blurv1Initialized // Event containing the contract specifics and raw log

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
func (it *Blurv1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1Initialized)
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
		it.Event = new(Blurv1Initialized)
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
func (it *Blurv1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1Initialized represents a Initialized event raised by the Blurv1 contract.
type Blurv1Initialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blurv1 *Blurv1Filterer) FilterInitialized(opts *bind.FilterOpts) (*Blurv1InitializedIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &Blurv1InitializedIterator{contract: _Blurv1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blurv1 *Blurv1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *Blurv1Initialized) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1Initialized)
				if err := _Blurv1.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_Blurv1 *Blurv1Filterer) ParseInitialized(log types.Log) (*Blurv1Initialized, error) {
	event := new(Blurv1Initialized)
	if err := _Blurv1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewBlockRangeIterator is returned from FilterNewBlockRange and is used to iterate over the raw logs and unpacked data for NewBlockRange events raised by the Blurv1 contract.
type Blurv1NewBlockRangeIterator struct {
	Event *Blurv1NewBlockRange // Event containing the contract specifics and raw log

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
func (it *Blurv1NewBlockRangeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewBlockRange)
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
		it.Event = new(Blurv1NewBlockRange)
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
func (it *Blurv1NewBlockRangeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewBlockRangeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewBlockRange represents a NewBlockRange event raised by the Blurv1 contract.
type Blurv1NewBlockRange struct {
	BlockRange *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterNewBlockRange is a free log retrieval operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blurv1 *Blurv1Filterer) FilterNewBlockRange(opts *bind.FilterOpts) (*Blurv1NewBlockRangeIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return &Blurv1NewBlockRangeIterator{contract: _Blurv1.contract, event: "NewBlockRange", logs: logs, sub: sub}, nil
}

// WatchNewBlockRange is a free log subscription operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blurv1 *Blurv1Filterer) WatchNewBlockRange(opts *bind.WatchOpts, sink chan<- *Blurv1NewBlockRange) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewBlockRange")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewBlockRange)
				if err := _Blurv1.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
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

// ParseNewBlockRange is a log parse operation binding the contract event 0x7706177c541ba1b858371bfc568aa77450b4713bbdbba62c730d4484ab6c1251.
//
// Solidity: event NewBlockRange(uint256 blockRange)
func (_Blurv1 *Blurv1Filterer) ParseNewBlockRange(log types.Log) (*Blurv1NewBlockRange, error) {
	event := new(Blurv1NewBlockRange)
	if err := _Blurv1.contract.UnpackLog(event, "NewBlockRange", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewExecutionDelegateIterator is returned from FilterNewExecutionDelegate and is used to iterate over the raw logs and unpacked data for NewExecutionDelegate events raised by the Blurv1 contract.
type Blurv1NewExecutionDelegateIterator struct {
	Event *Blurv1NewExecutionDelegate // Event containing the contract specifics and raw log

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
func (it *Blurv1NewExecutionDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewExecutionDelegate)
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
		it.Event = new(Blurv1NewExecutionDelegate)
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
func (it *Blurv1NewExecutionDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewExecutionDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewExecutionDelegate represents a NewExecutionDelegate event raised by the Blurv1 contract.
type Blurv1NewExecutionDelegate struct {
	ExecutionDelegate common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterNewExecutionDelegate is a free log retrieval operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blurv1 *Blurv1Filterer) FilterNewExecutionDelegate(opts *bind.FilterOpts, executionDelegate []common.Address) (*Blurv1NewExecutionDelegateIterator, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1NewExecutionDelegateIterator{contract: _Blurv1.contract, event: "NewExecutionDelegate", logs: logs, sub: sub}, nil
}

// WatchNewExecutionDelegate is a free log subscription operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blurv1 *Blurv1Filterer) WatchNewExecutionDelegate(opts *bind.WatchOpts, sink chan<- *Blurv1NewExecutionDelegate, executionDelegate []common.Address) (event.Subscription, error) {

	var executionDelegateRule []interface{}
	for _, executionDelegateItem := range executionDelegate {
		executionDelegateRule = append(executionDelegateRule, executionDelegateItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewExecutionDelegate", executionDelegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewExecutionDelegate)
				if err := _Blurv1.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
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

// ParseNewExecutionDelegate is a log parse operation binding the contract event 0xf9a0f356a7ef079355de09d32ce45cc3cfabc8f118681c19a17501f005a376ac.
//
// Solidity: event NewExecutionDelegate(address indexed executionDelegate)
func (_Blurv1 *Blurv1Filterer) ParseNewExecutionDelegate(log types.Log) (*Blurv1NewExecutionDelegate, error) {
	event := new(Blurv1NewExecutionDelegate)
	if err := _Blurv1.contract.UnpackLog(event, "NewExecutionDelegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewFeeRateIterator is returned from FilterNewFeeRate and is used to iterate over the raw logs and unpacked data for NewFeeRate events raised by the Blurv1 contract.
type Blurv1NewFeeRateIterator struct {
	Event *Blurv1NewFeeRate // Event containing the contract specifics and raw log

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
func (it *Blurv1NewFeeRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewFeeRate)
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
		it.Event = new(Blurv1NewFeeRate)
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
func (it *Blurv1NewFeeRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewFeeRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewFeeRate represents a NewFeeRate event raised by the Blurv1 contract.
type Blurv1NewFeeRate struct {
	FeeRate *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRate is a free log retrieval operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blurv1 *Blurv1Filterer) FilterNewFeeRate(opts *bind.FilterOpts) (*Blurv1NewFeeRateIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return &Blurv1NewFeeRateIterator{contract: _Blurv1.contract, event: "NewFeeRate", logs: logs, sub: sub}, nil
}

// WatchNewFeeRate is a free log subscription operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blurv1 *Blurv1Filterer) WatchNewFeeRate(opts *bind.WatchOpts, sink chan<- *Blurv1NewFeeRate) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewFeeRate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewFeeRate)
				if err := _Blurv1.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
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

// ParseNewFeeRate is a log parse operation binding the contract event 0x788980e82f4651cc86d1cc00916685528f16c9abb21b2afe72325496c18c94ae.
//
// Solidity: event NewFeeRate(uint256 feeRate)
func (_Blurv1 *Blurv1Filterer) ParseNewFeeRate(log types.Log) (*Blurv1NewFeeRate, error) {
	event := new(Blurv1NewFeeRate)
	if err := _Blurv1.contract.UnpackLog(event, "NewFeeRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewFeeRecipientIterator is returned from FilterNewFeeRecipient and is used to iterate over the raw logs and unpacked data for NewFeeRecipient events raised by the Blurv1 contract.
type Blurv1NewFeeRecipientIterator struct {
	Event *Blurv1NewFeeRecipient // Event containing the contract specifics and raw log

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
func (it *Blurv1NewFeeRecipientIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewFeeRecipient)
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
		it.Event = new(Blurv1NewFeeRecipient)
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
func (it *Blurv1NewFeeRecipientIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewFeeRecipientIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewFeeRecipient represents a NewFeeRecipient event raised by the Blurv1 contract.
type Blurv1NewFeeRecipient struct {
	FeeRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterNewFeeRecipient is a free log retrieval operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blurv1 *Blurv1Filterer) FilterNewFeeRecipient(opts *bind.FilterOpts) (*Blurv1NewFeeRecipientIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return &Blurv1NewFeeRecipientIterator{contract: _Blurv1.contract, event: "NewFeeRecipient", logs: logs, sub: sub}, nil
}

// WatchNewFeeRecipient is a free log subscription operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blurv1 *Blurv1Filterer) WatchNewFeeRecipient(opts *bind.WatchOpts, sink chan<- *Blurv1NewFeeRecipient) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewFeeRecipient")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewFeeRecipient)
				if err := _Blurv1.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
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

// ParseNewFeeRecipient is a log parse operation binding the contract event 0x412871529f3cedd6ca6f10784258f4965a5d6e254127593fe354e7a62f6d0a23.
//
// Solidity: event NewFeeRecipient(address feeRecipient)
func (_Blurv1 *Blurv1Filterer) ParseNewFeeRecipient(log types.Log) (*Blurv1NewFeeRecipient, error) {
	event := new(Blurv1NewFeeRecipient)
	if err := _Blurv1.contract.UnpackLog(event, "NewFeeRecipient", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewGovernorIterator is returned from FilterNewGovernor and is used to iterate over the raw logs and unpacked data for NewGovernor events raised by the Blurv1 contract.
type Blurv1NewGovernorIterator struct {
	Event *Blurv1NewGovernor // Event containing the contract specifics and raw log

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
func (it *Blurv1NewGovernorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewGovernor)
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
		it.Event = new(Blurv1NewGovernor)
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
func (it *Blurv1NewGovernorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewGovernorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewGovernor represents a NewGovernor event raised by the Blurv1 contract.
type Blurv1NewGovernor struct {
	Governor common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNewGovernor is a free log retrieval operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blurv1 *Blurv1Filterer) FilterNewGovernor(opts *bind.FilterOpts) (*Blurv1NewGovernorIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return &Blurv1NewGovernorIterator{contract: _Blurv1.contract, event: "NewGovernor", logs: logs, sub: sub}, nil
}

// WatchNewGovernor is a free log subscription operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blurv1 *Blurv1Filterer) WatchNewGovernor(opts *bind.WatchOpts, sink chan<- *Blurv1NewGovernor) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewGovernor")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewGovernor)
				if err := _Blurv1.contract.UnpackLog(event, "NewGovernor", log); err != nil {
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

// ParseNewGovernor is a log parse operation binding the contract event 0x5425363a03f182281120f5919107c49c7a1a623acc1cbc6df468b6f0c11fcf8c.
//
// Solidity: event NewGovernor(address governor)
func (_Blurv1 *Blurv1Filterer) ParseNewGovernor(log types.Log) (*Blurv1NewGovernor, error) {
	event := new(Blurv1NewGovernor)
	if err := _Blurv1.contract.UnpackLog(event, "NewGovernor", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewOracleIterator is returned from FilterNewOracle and is used to iterate over the raw logs and unpacked data for NewOracle events raised by the Blurv1 contract.
type Blurv1NewOracleIterator struct {
	Event *Blurv1NewOracle // Event containing the contract specifics and raw log

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
func (it *Blurv1NewOracleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewOracle)
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
		it.Event = new(Blurv1NewOracle)
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
func (it *Blurv1NewOracleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewOracleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewOracle represents a NewOracle event raised by the Blurv1 contract.
type Blurv1NewOracle struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterNewOracle is a free log retrieval operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blurv1 *Blurv1Filterer) FilterNewOracle(opts *bind.FilterOpts, oracle []common.Address) (*Blurv1NewOracleIterator, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1NewOracleIterator{contract: _Blurv1.contract, event: "NewOracle", logs: logs, sub: sub}, nil
}

// WatchNewOracle is a free log subscription operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blurv1 *Blurv1Filterer) WatchNewOracle(opts *bind.WatchOpts, sink chan<- *Blurv1NewOracle, oracle []common.Address) (event.Subscription, error) {

	var oracleRule []interface{}
	for _, oracleItem := range oracle {
		oracleRule = append(oracleRule, oracleItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewOracle", oracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewOracle)
				if err := _Blurv1.contract.UnpackLog(event, "NewOracle", log); err != nil {
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

// ParseNewOracle is a log parse operation binding the contract event 0xb3eacd0e351fafdfefdec84e1cd19679b38dbcd63ea7c2c24da17fd2bc3b3c0e.
//
// Solidity: event NewOracle(address indexed oracle)
func (_Blurv1 *Blurv1Filterer) ParseNewOracle(log types.Log) (*Blurv1NewOracle, error) {
	event := new(Blurv1NewOracle)
	if err := _Blurv1.contract.UnpackLog(event, "NewOracle", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NewPolicyManagerIterator is returned from FilterNewPolicyManager and is used to iterate over the raw logs and unpacked data for NewPolicyManager events raised by the Blurv1 contract.
type Blurv1NewPolicyManagerIterator struct {
	Event *Blurv1NewPolicyManager // Event containing the contract specifics and raw log

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
func (it *Blurv1NewPolicyManagerIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NewPolicyManager)
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
		it.Event = new(Blurv1NewPolicyManager)
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
func (it *Blurv1NewPolicyManagerIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NewPolicyManagerIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NewPolicyManager represents a NewPolicyManager event raised by the Blurv1 contract.
type Blurv1NewPolicyManager struct {
	PolicyManager common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterNewPolicyManager is a free log retrieval operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blurv1 *Blurv1Filterer) FilterNewPolicyManager(opts *bind.FilterOpts, policyManager []common.Address) (*Blurv1NewPolicyManagerIterator, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1NewPolicyManagerIterator{contract: _Blurv1.contract, event: "NewPolicyManager", logs: logs, sub: sub}, nil
}

// WatchNewPolicyManager is a free log subscription operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blurv1 *Blurv1Filterer) WatchNewPolicyManager(opts *bind.WatchOpts, sink chan<- *Blurv1NewPolicyManager, policyManager []common.Address) (event.Subscription, error) {

	var policyManagerRule []interface{}
	for _, policyManagerItem := range policyManager {
		policyManagerRule = append(policyManagerRule, policyManagerItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NewPolicyManager", policyManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NewPolicyManager)
				if err := _Blurv1.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
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

// ParseNewPolicyManager is a log parse operation binding the contract event 0xdbe18f3fd927cc2aefe380ffd2abfdb8e13f0239c0258ccfc84c3d8fdd8c0418.
//
// Solidity: event NewPolicyManager(address indexed policyManager)
func (_Blurv1 *Blurv1Filterer) ParseNewPolicyManager(log types.Log) (*Blurv1NewPolicyManager, error) {
	event := new(Blurv1NewPolicyManager)
	if err := _Blurv1.contract.UnpackLog(event, "NewPolicyManager", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1NonceIncrementedIterator is returned from FilterNonceIncremented and is used to iterate over the raw logs and unpacked data for NonceIncremented events raised by the Blurv1 contract.
type Blurv1NonceIncrementedIterator struct {
	Event *Blurv1NonceIncremented // Event containing the contract specifics and raw log

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
func (it *Blurv1NonceIncrementedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1NonceIncremented)
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
		it.Event = new(Blurv1NonceIncremented)
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
func (it *Blurv1NonceIncrementedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1NonceIncrementedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1NonceIncremented represents a NonceIncremented event raised by the Blurv1 contract.
type Blurv1NonceIncremented struct {
	Trader   common.Address
	NewNonce *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterNonceIncremented is a free log retrieval operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blurv1 *Blurv1Filterer) FilterNonceIncremented(opts *bind.FilterOpts, trader []common.Address) (*Blurv1NonceIncrementedIterator, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1NonceIncrementedIterator{contract: _Blurv1.contract, event: "NonceIncremented", logs: logs, sub: sub}, nil
}

// WatchNonceIncremented is a free log subscription operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blurv1 *Blurv1Filterer) WatchNonceIncremented(opts *bind.WatchOpts, sink chan<- *Blurv1NonceIncremented, trader []common.Address) (event.Subscription, error) {

	var traderRule []interface{}
	for _, traderItem := range trader {
		traderRule = append(traderRule, traderItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "NonceIncremented", traderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1NonceIncremented)
				if err := _Blurv1.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
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

// ParseNonceIncremented is a log parse operation binding the contract event 0xa82a649bbd060c9099cd7b7326e2b0dc9e9af0836480e0f849dc9eaa79710b3b.
//
// Solidity: event NonceIncremented(address indexed trader, uint256 newNonce)
func (_Blurv1 *Blurv1Filterer) ParseNonceIncremented(log types.Log) (*Blurv1NonceIncremented, error) {
	event := new(Blurv1NonceIncremented)
	if err := _Blurv1.contract.UnpackLog(event, "NonceIncremented", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1OpenedIterator is returned from FilterOpened and is used to iterate over the raw logs and unpacked data for Opened events raised by the Blurv1 contract.
type Blurv1OpenedIterator struct {
	Event *Blurv1Opened // Event containing the contract specifics and raw log

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
func (it *Blurv1OpenedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1Opened)
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
		it.Event = new(Blurv1Opened)
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
func (it *Blurv1OpenedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1OpenedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1Opened represents a Opened event raised by the Blurv1 contract.
type Blurv1Opened struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterOpened is a free log retrieval operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blurv1 *Blurv1Filterer) FilterOpened(opts *bind.FilterOpts) (*Blurv1OpenedIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return &Blurv1OpenedIterator{contract: _Blurv1.contract, event: "Opened", logs: logs, sub: sub}, nil
}

// WatchOpened is a free log subscription operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blurv1 *Blurv1Filterer) WatchOpened(opts *bind.WatchOpts, sink chan<- *Blurv1Opened) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "Opened")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1Opened)
				if err := _Blurv1.contract.UnpackLog(event, "Opened", log); err != nil {
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

// ParseOpened is a log parse operation binding the contract event 0xd1dcd00534373f20882b79e6ab6875a5c358c5bd576448757ed50e63069ab518.
//
// Solidity: event Opened()
func (_Blurv1 *Blurv1Filterer) ParseOpened(log types.Log) (*Blurv1Opened, error) {
	event := new(Blurv1Opened)
	if err := _Blurv1.contract.UnpackLog(event, "Opened", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1OrderCancelledIterator is returned from FilterOrderCancelled and is used to iterate over the raw logs and unpacked data for OrderCancelled events raised by the Blurv1 contract.
type Blurv1OrderCancelledIterator struct {
	Event *Blurv1OrderCancelled // Event containing the contract specifics and raw log

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
func (it *Blurv1OrderCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1OrderCancelled)
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
		it.Event = new(Blurv1OrderCancelled)
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
func (it *Blurv1OrderCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1OrderCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1OrderCancelled represents a OrderCancelled event raised by the Blurv1 contract.
type Blurv1OrderCancelled struct {
	Hash [32]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOrderCancelled is a free log retrieval operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blurv1 *Blurv1Filterer) FilterOrderCancelled(opts *bind.FilterOpts) (*Blurv1OrderCancelledIterator, error) {

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return &Blurv1OrderCancelledIterator{contract: _Blurv1.contract, event: "OrderCancelled", logs: logs, sub: sub}, nil
}

// WatchOrderCancelled is a free log subscription operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blurv1 *Blurv1Filterer) WatchOrderCancelled(opts *bind.WatchOpts, sink chan<- *Blurv1OrderCancelled) (event.Subscription, error) {

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "OrderCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1OrderCancelled)
				if err := _Blurv1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
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

// ParseOrderCancelled is a log parse operation binding the contract event 0x5152abf959f6564662358c2e52b702259b78bac5ee7842a0f01937e670efcc7d.
//
// Solidity: event OrderCancelled(bytes32 hash)
func (_Blurv1 *Blurv1Filterer) ParseOrderCancelled(log types.Log) (*Blurv1OrderCancelled, error) {
	event := new(Blurv1OrderCancelled)
	if err := _Blurv1.contract.UnpackLog(event, "OrderCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1OrdersMatchedIterator is returned from FilterOrdersMatched and is used to iterate over the raw logs and unpacked data for OrdersMatched events raised by the Blurv1 contract.
type Blurv1OrdersMatchedIterator struct {
	Event *Blurv1OrdersMatched // Event containing the contract specifics and raw log

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
func (it *Blurv1OrdersMatchedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1OrdersMatched)
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
		it.Event = new(Blurv1OrdersMatched)
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
func (it *Blurv1OrdersMatchedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1OrdersMatchedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1OrdersMatched represents a OrdersMatched event raised by the Blurv1 contract.
type Blurv1OrdersMatched struct {
	Maker    common.Address
	Taker    common.Address
	Sell     Order
	SellHash [32]byte
	Buy      Order
	BuyHash  [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOrdersMatched is a free log retrieval operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blurv1 *Blurv1Filterer) FilterOrdersMatched(opts *bind.FilterOpts, maker []common.Address, taker []common.Address) (*Blurv1OrdersMatchedIterator, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1OrdersMatchedIterator{contract: _Blurv1.contract, event: "OrdersMatched", logs: logs, sub: sub}, nil
}

// WatchOrdersMatched is a free log subscription operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blurv1 *Blurv1Filterer) WatchOrdersMatched(opts *bind.WatchOpts, sink chan<- *Blurv1OrdersMatched, maker []common.Address, taker []common.Address) (event.Subscription, error) {

	var makerRule []interface{}
	for _, makerItem := range maker {
		makerRule = append(makerRule, makerItem)
	}
	var takerRule []interface{}
	for _, takerItem := range taker {
		takerRule = append(takerRule, takerItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "OrdersMatched", makerRule, takerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1OrdersMatched)
				if err := _Blurv1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
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

// ParseOrdersMatched is a log parse operation binding the contract event 0x61cbb2a3dee0b6064c2e681aadd61677fb4ef319f0b547508d495626f5a62f64.
//
// Solidity: event OrdersMatched(address indexed maker, address indexed taker, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) sell, bytes32 sellHash, (address,uint8,address,address,uint256,uint256,address,uint256,uint256,uint256,(uint16,address)[],uint256,bytes) buy, bytes32 buyHash)
func (_Blurv1 *Blurv1Filterer) ParseOrdersMatched(log types.Log) (*Blurv1OrdersMatched, error) {
	event := new(Blurv1OrdersMatched)
	if err := _Blurv1.contract.UnpackLog(event, "OrdersMatched", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Blurv1 contract.
type Blurv1OwnershipTransferredIterator struct {
	Event *Blurv1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *Blurv1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1OwnershipTransferred)
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
		it.Event = new(Blurv1OwnershipTransferred)
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
func (it *Blurv1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1OwnershipTransferred represents a OwnershipTransferred event raised by the Blurv1 contract.
type Blurv1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blurv1 *Blurv1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*Blurv1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1OwnershipTransferredIterator{contract: _Blurv1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Blurv1 *Blurv1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *Blurv1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1OwnershipTransferred)
				if err := _Blurv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Blurv1 *Blurv1Filterer) ParseOwnershipTransferred(log types.Log) (*Blurv1OwnershipTransferred, error) {
	event := new(Blurv1OwnershipTransferred)
	if err := _Blurv1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Blurv1UpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the Blurv1 contract.
type Blurv1UpgradedIterator struct {
	Event *Blurv1Upgraded // Event containing the contract specifics and raw log

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
func (it *Blurv1UpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Blurv1Upgraded)
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
		it.Event = new(Blurv1Upgraded)
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
func (it *Blurv1UpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Blurv1UpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Blurv1Upgraded represents a Upgraded event raised by the Blurv1 contract.
type Blurv1Upgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blurv1 *Blurv1Filterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*Blurv1UpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Blurv1.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &Blurv1UpgradedIterator{contract: _Blurv1.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blurv1 *Blurv1Filterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *Blurv1Upgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _Blurv1.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Blurv1Upgraded)
				if err := _Blurv1.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_Blurv1 *Blurv1Filterer) ParseUpgraded(log types.Log) (*Blurv1Upgraded, error) {
	event := new(Blurv1Upgraded)
	if err := _Blurv1.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
