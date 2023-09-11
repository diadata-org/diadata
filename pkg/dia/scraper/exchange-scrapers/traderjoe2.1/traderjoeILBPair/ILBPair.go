// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoeILBPair

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

// ILBPairMetaData contains all meta data concerning the ILBPair contract.
var ILBPairMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"LBPair__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__AlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__EmptyMarketConfigs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__FlashLoanCallbackFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__FlashLoanInsufficientAmount\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__InsufficientAmountIn\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__InsufficientAmountOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__InvalidInput\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__InvalidStaticFeeParameters\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__MaxTotalFeeExceeded\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__OnlyFactory\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__OnlyProtocolFeeRecipient\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__OutOfLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__TokenNotSupported\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"LBPair__ZeroAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"LBPair__ZeroAmountsOut\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBPair__ZeroBorrowAmount\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"LBPair__ZeroShares\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__AddressThisOrZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBToken__BurnExceedsBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBToken__InvalidLength\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"LBToken__SelfApproval\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"LBToken__SpenderNotApproved\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"LBToken__TransferExceedsBalance\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"protocolFees\",\"type\":\"bytes32\"}],\"name\":\"CollectedProtocolFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"totalFees\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"protocolFees\",\"type\":\"bytes32\"}],\"name\":\"CompositionFees\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"amounts\",\"type\":\"bytes32[]\"}],\"name\":\"DepositedToBins\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractILBFlashLoanCallback\",\"name\":\"receiver\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"amounts\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"totalFees\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"protocolFees\",\"type\":\"bytes32\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"idReference\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"volatilityReference\",\"type\":\"uint24\"}],\"name\":\"ForcedDecay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"oracleLength\",\"type\":\"uint16\"}],\"name\":\"OracleLengthIncreased\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"}],\"name\":\"StaticFeeParametersSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"amountsIn\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"amountsOut\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"volatilityAccumulator\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"totalFees\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"protocolFees\",\"type\":\"bytes32\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"TransferBatch\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"bytes32[]\",\"name\":\"amounts\",\"type\":\"bytes32[]\"}],\"name\":\"WithdrawnFromBins\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"approveForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"accounts\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"}],\"name\":\"balanceOfBatch\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"}],\"name\":\"batchTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"ids\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsToBurn\",\"type\":\"uint256[]\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"amounts\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"collectProtocolFees\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"collectedProtocolFees\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBFlashLoanCallback\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"amounts\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"forceDecay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveId\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"getBin\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"binReserveX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"binReserveY\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBinStep\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFactory\",\"outputs\":[{\"internalType\":\"contractILBFactory\",\"name\":\"factory\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"getIdFromPrice\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"},{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"getNextNonEmptyBin\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"nextId\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOracleParameters\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"sampleLifetime\",\"type\":\"uint8\"},{\"internalType\":\"uint16\",\"name\":\"size\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"activeSize\",\"type\":\"uint16\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"uint40\",\"name\":\"firstTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint40\",\"name\":\"lookupTimestamp\",\"type\":\"uint40\"}],\"name\":\"getOracleSampleAt\",\"outputs\":[{\"internalType\":\"uint64\",\"name\":\"cumulativeId\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cumulativeVolatility\",\"type\":\"uint64\"},{\"internalType\":\"uint64\",\"name\":\"cumulativeBinCrossed\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"id\",\"type\":\"uint24\"}],\"name\":\"getPriceFromId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFees\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"protocolFeeX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"protocolFeeY\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReserves\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"reserveX\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reserveY\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getStaticFeeParameters\",\"outputs\":[{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapIn\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOutLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"amountIn\",\"type\":\"uint128\"},{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"}],\"name\":\"getSwapOut\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"amountInLeft\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amountOut\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"fee\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenX\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTokenY\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVariableFeeParameters\",\"outputs\":[{\"internalType\":\"uint24\",\"name\":\"volatilityAccumulator\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"volatilityReference\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"idReference\",\"type\":\"uint24\"},{\"internalType\":\"uint40\",\"name\":\"timeOfLastUpdate\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"newLength\",\"type\":\"uint16\"}],\"name\":\"increaseOracleLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"},{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"bytes32[]\",\"name\":\"liquidityConfigs\",\"type\":\"bytes32[]\"},{\"internalType\":\"address\",\"name\":\"refundTo\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"amountsReceived\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"amountsLeft\",\"type\":\"bytes32\"},{\"internalType\":\"uint256[]\",\"name\":\"liquidityMinted\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"}],\"name\":\"setStaticFeeParameters\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"swapForY\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"amountsOut\",\"type\":\"bytes32\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// ILBPairABI is the input ABI used to generate the binding from.
// Deprecated: Use ILBPairMetaData.ABI instead.
var ILBPairABI = ILBPairMetaData.ABI

// ILBPair is an auto generated Go binding around an Ethereum contract.
type ILBPair struct {
	ILBPairCaller     // Read-only binding to the contract
	ILBPairTransactor // Write-only binding to the contract
	ILBPairFilterer   // Log filterer for contract events
}

// ILBPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILBPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILBPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILBPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILBPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILBPairSession struct {
	Contract     *ILBPair          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILBPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILBPairCallerSession struct {
	Contract *ILBPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// ILBPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILBPairTransactorSession struct {
	Contract     *ILBPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// ILBPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILBPairRaw struct {
	Contract *ILBPair // Generic contract binding to access the raw methods on
}

// ILBPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILBPairCallerRaw struct {
	Contract *ILBPairCaller // Generic read-only contract binding to access the raw methods on
}

// ILBPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILBPairTransactorRaw struct {
	Contract *ILBPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILBPair creates a new instance of ILBPair, bound to a specific deployed contract.
func NewILBPair(address common.Address, backend bind.ContractBackend) (*ILBPair, error) {
	contract, err := bindILBPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILBPair{ILBPairCaller: ILBPairCaller{contract: contract}, ILBPairTransactor: ILBPairTransactor{contract: contract}, ILBPairFilterer: ILBPairFilterer{contract: contract}}, nil
}

// NewILBPairCaller creates a new read-only instance of ILBPair, bound to a specific deployed contract.
func NewILBPairCaller(address common.Address, caller bind.ContractCaller) (*ILBPairCaller, error) {
	contract, err := bindILBPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILBPairCaller{contract: contract}, nil
}

// NewILBPairTransactor creates a new write-only instance of ILBPair, bound to a specific deployed contract.
func NewILBPairTransactor(address common.Address, transactor bind.ContractTransactor) (*ILBPairTransactor, error) {
	contract, err := bindILBPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILBPairTransactor{contract: contract}, nil
}

// NewILBPairFilterer creates a new log filterer instance of ILBPair, bound to a specific deployed contract.
func NewILBPairFilterer(address common.Address, filterer bind.ContractFilterer) (*ILBPairFilterer, error) {
	contract, err := bindILBPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILBPairFilterer{contract: contract}, nil
}

// bindILBPair binds a generic wrapper to an already deployed contract.
func bindILBPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ILBPairMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILBPair *ILBPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILBPair.Contract.ILBPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILBPair *ILBPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBPair.Contract.ILBPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILBPair *ILBPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILBPair.Contract.ILBPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILBPair *ILBPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILBPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILBPair *ILBPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILBPair *ILBPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILBPair.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ILBPair *ILBPairCaller) BalanceOf(opts *bind.CallOpts, account common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "balanceOf", account, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ILBPair *ILBPairSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.BalanceOf(&_ILBPair.CallOpts, account, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address account, uint256 id) view returns(uint256)
func (_ILBPair *ILBPairCallerSession) BalanceOf(account common.Address, id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.BalanceOf(&_ILBPair.CallOpts, account, id)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ILBPair *ILBPairCaller) BalanceOfBatch(opts *bind.CallOpts, accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "balanceOfBatch", accounts, ids)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ILBPair *ILBPairSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ILBPair.Contract.BalanceOfBatch(&_ILBPair.CallOpts, accounts, ids)
}

// BalanceOfBatch is a free data retrieval call binding the contract method 0x4e1273f4.
//
// Solidity: function balanceOfBatch(address[] accounts, uint256[] ids) view returns(uint256[])
func (_ILBPair *ILBPairCallerSession) BalanceOfBatch(accounts []common.Address, ids []*big.Int) ([]*big.Int, error) {
	return _ILBPair.Contract.BalanceOfBatch(&_ILBPair.CallOpts, accounts, ids)
}

// GetActiveId is a free data retrieval call binding the contract method 0xdbe65edc.
//
// Solidity: function getActiveId() view returns(uint24 activeId)
func (_ILBPair *ILBPairCaller) GetActiveId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getActiveId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetActiveId is a free data retrieval call binding the contract method 0xdbe65edc.
//
// Solidity: function getActiveId() view returns(uint24 activeId)
func (_ILBPair *ILBPairSession) GetActiveId() (*big.Int, error) {
	return _ILBPair.Contract.GetActiveId(&_ILBPair.CallOpts)
}

// GetActiveId is a free data retrieval call binding the contract method 0xdbe65edc.
//
// Solidity: function getActiveId() view returns(uint24 activeId)
func (_ILBPair *ILBPairCallerSession) GetActiveId() (*big.Int, error) {
	return _ILBPair.Contract.GetActiveId(&_ILBPair.CallOpts)
}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 id) view returns(uint128 binReserveX, uint128 binReserveY)
func (_ILBPair *ILBPairCaller) GetBin(opts *bind.CallOpts, id *big.Int) (struct {
	BinReserveX *big.Int
	BinReserveY *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getBin", id)

	outstruct := new(struct {
		BinReserveX *big.Int
		BinReserveY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BinReserveX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.BinReserveY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 id) view returns(uint128 binReserveX, uint128 binReserveY)
func (_ILBPair *ILBPairSession) GetBin(id *big.Int) (struct {
	BinReserveX *big.Int
	BinReserveY *big.Int
}, error) {
	return _ILBPair.Contract.GetBin(&_ILBPair.CallOpts, id)
}

// GetBin is a free data retrieval call binding the contract method 0x0abe9688.
//
// Solidity: function getBin(uint24 id) view returns(uint128 binReserveX, uint128 binReserveY)
func (_ILBPair *ILBPairCallerSession) GetBin(id *big.Int) (struct {
	BinReserveX *big.Int
	BinReserveY *big.Int
}, error) {
	return _ILBPair.Contract.GetBin(&_ILBPair.CallOpts, id)
}

// GetBinStep is a free data retrieval call binding the contract method 0x17f11ecc.
//
// Solidity: function getBinStep() view returns(uint16 binStep)
func (_ILBPair *ILBPairCaller) GetBinStep(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getBinStep")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetBinStep is a free data retrieval call binding the contract method 0x17f11ecc.
//
// Solidity: function getBinStep() view returns(uint16 binStep)
func (_ILBPair *ILBPairSession) GetBinStep() (uint16, error) {
	return _ILBPair.Contract.GetBinStep(&_ILBPair.CallOpts)
}

// GetBinStep is a free data retrieval call binding the contract method 0x17f11ecc.
//
// Solidity: function getBinStep() view returns(uint16 binStep)
func (_ILBPair *ILBPairCallerSession) GetBinStep() (uint16, error) {
	return _ILBPair.Contract.GetBinStep(&_ILBPair.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address factory)
func (_ILBPair *ILBPairCaller) GetFactory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getFactory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address factory)
func (_ILBPair *ILBPairSession) GetFactory() (common.Address, error) {
	return _ILBPair.Contract.GetFactory(&_ILBPair.CallOpts)
}

// GetFactory is a free data retrieval call binding the contract method 0x88cc58e4.
//
// Solidity: function getFactory() view returns(address factory)
func (_ILBPair *ILBPairCallerSession) GetFactory() (common.Address, error) {
	return _ILBPair.Contract.GetFactory(&_ILBPair.CallOpts)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf5e29329.
//
// Solidity: function getIdFromPrice(uint256 price) view returns(uint24 id)
func (_ILBPair *ILBPairCaller) GetIdFromPrice(opts *bind.CallOpts, price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getIdFromPrice", price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf5e29329.
//
// Solidity: function getIdFromPrice(uint256 price) view returns(uint24 id)
func (_ILBPair *ILBPairSession) GetIdFromPrice(price *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetIdFromPrice(&_ILBPair.CallOpts, price)
}

// GetIdFromPrice is a free data retrieval call binding the contract method 0xf5e29329.
//
// Solidity: function getIdFromPrice(uint256 price) view returns(uint24 id)
func (_ILBPair *ILBPairCallerSession) GetIdFromPrice(price *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetIdFromPrice(&_ILBPair.CallOpts, price)
}

// GetNextNonEmptyBin is a free data retrieval call binding the contract method 0xa41a01fb.
//
// Solidity: function getNextNonEmptyBin(bool swapForY, uint24 id) view returns(uint24 nextId)
func (_ILBPair *ILBPairCaller) GetNextNonEmptyBin(opts *bind.CallOpts, swapForY bool, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getNextNonEmptyBin", swapForY, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextNonEmptyBin is a free data retrieval call binding the contract method 0xa41a01fb.
//
// Solidity: function getNextNonEmptyBin(bool swapForY, uint24 id) view returns(uint24 nextId)
func (_ILBPair *ILBPairSession) GetNextNonEmptyBin(swapForY bool, id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetNextNonEmptyBin(&_ILBPair.CallOpts, swapForY, id)
}

// GetNextNonEmptyBin is a free data retrieval call binding the contract method 0xa41a01fb.
//
// Solidity: function getNextNonEmptyBin(bool swapForY, uint24 id) view returns(uint24 nextId)
func (_ILBPair *ILBPairCallerSession) GetNextNonEmptyBin(swapForY bool, id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetNextNonEmptyBin(&_ILBPair.CallOpts, swapForY, id)
}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint8 sampleLifetime, uint16 size, uint16 activeSize, uint40 lastUpdated, uint40 firstTimestamp)
func (_ILBPair *ILBPairCaller) GetOracleParameters(opts *bind.CallOpts) (struct {
	SampleLifetime uint8
	Size           uint16
	ActiveSize     uint16
	LastUpdated    *big.Int
	FirstTimestamp *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getOracleParameters")

	outstruct := new(struct {
		SampleLifetime uint8
		Size           uint16
		ActiveSize     uint16
		LastUpdated    *big.Int
		FirstTimestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SampleLifetime = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Size = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.ActiveSize = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.LastUpdated = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.FirstTimestamp = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint8 sampleLifetime, uint16 size, uint16 activeSize, uint40 lastUpdated, uint40 firstTimestamp)
func (_ILBPair *ILBPairSession) GetOracleParameters() (struct {
	SampleLifetime uint8
	Size           uint16
	ActiveSize     uint16
	LastUpdated    *big.Int
	FirstTimestamp *big.Int
}, error) {
	return _ILBPair.Contract.GetOracleParameters(&_ILBPair.CallOpts)
}

// GetOracleParameters is a free data retrieval call binding the contract method 0x55182894.
//
// Solidity: function getOracleParameters() view returns(uint8 sampleLifetime, uint16 size, uint16 activeSize, uint40 lastUpdated, uint40 firstTimestamp)
func (_ILBPair *ILBPairCallerSession) GetOracleParameters() (struct {
	SampleLifetime uint8
	Size           uint16
	ActiveSize     uint16
	LastUpdated    *big.Int
	FirstTimestamp *big.Int
}, error) {
	return _ILBPair.Contract.GetOracleParameters(&_ILBPair.CallOpts)
}

// GetOracleSampleAt is a free data retrieval call binding the contract method 0x8940a16a.
//
// Solidity: function getOracleSampleAt(uint40 lookupTimestamp) view returns(uint64 cumulativeId, uint64 cumulativeVolatility, uint64 cumulativeBinCrossed)
func (_ILBPair *ILBPairCaller) GetOracleSampleAt(opts *bind.CallOpts, lookupTimestamp *big.Int) (struct {
	CumulativeId         uint64
	CumulativeVolatility uint64
	CumulativeBinCrossed uint64
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getOracleSampleAt", lookupTimestamp)

	outstruct := new(struct {
		CumulativeId         uint64
		CumulativeVolatility uint64
		CumulativeBinCrossed uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CumulativeId = *abi.ConvertType(out[0], new(uint64)).(*uint64)
	outstruct.CumulativeVolatility = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.CumulativeBinCrossed = *abi.ConvertType(out[2], new(uint64)).(*uint64)

	return *outstruct, err

}

// GetOracleSampleAt is a free data retrieval call binding the contract method 0x8940a16a.
//
// Solidity: function getOracleSampleAt(uint40 lookupTimestamp) view returns(uint64 cumulativeId, uint64 cumulativeVolatility, uint64 cumulativeBinCrossed)
func (_ILBPair *ILBPairSession) GetOracleSampleAt(lookupTimestamp *big.Int) (struct {
	CumulativeId         uint64
	CumulativeVolatility uint64
	CumulativeBinCrossed uint64
}, error) {
	return _ILBPair.Contract.GetOracleSampleAt(&_ILBPair.CallOpts, lookupTimestamp)
}

// GetOracleSampleAt is a free data retrieval call binding the contract method 0x8940a16a.
//
// Solidity: function getOracleSampleAt(uint40 lookupTimestamp) view returns(uint64 cumulativeId, uint64 cumulativeVolatility, uint64 cumulativeBinCrossed)
func (_ILBPair *ILBPairCallerSession) GetOracleSampleAt(lookupTimestamp *big.Int) (struct {
	CumulativeId         uint64
	CumulativeVolatility uint64
	CumulativeBinCrossed uint64
}, error) {
	return _ILBPair.Contract.GetOracleSampleAt(&_ILBPair.CallOpts, lookupTimestamp)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0x4c7cffbd.
//
// Solidity: function getPriceFromId(uint24 id) view returns(uint256 price)
func (_ILBPair *ILBPairCaller) GetPriceFromId(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getPriceFromId", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetPriceFromId is a free data retrieval call binding the contract method 0x4c7cffbd.
//
// Solidity: function getPriceFromId(uint24 id) view returns(uint256 price)
func (_ILBPair *ILBPairSession) GetPriceFromId(id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetPriceFromId(&_ILBPair.CallOpts, id)
}

// GetPriceFromId is a free data retrieval call binding the contract method 0x4c7cffbd.
//
// Solidity: function getPriceFromId(uint24 id) view returns(uint256 price)
func (_ILBPair *ILBPairCallerSession) GetPriceFromId(id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.GetPriceFromId(&_ILBPair.CallOpts, id)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint128 protocolFeeX, uint128 protocolFeeY)
func (_ILBPair *ILBPairCaller) GetProtocolFees(opts *bind.CallOpts) (struct {
	ProtocolFeeX *big.Int
	ProtocolFeeY *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getProtocolFees")

	outstruct := new(struct {
		ProtocolFeeX *big.Int
		ProtocolFeeY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ProtocolFeeX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ProtocolFeeY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint128 protocolFeeX, uint128 protocolFeeY)
func (_ILBPair *ILBPairSession) GetProtocolFees() (struct {
	ProtocolFeeX *big.Int
	ProtocolFeeY *big.Int
}, error) {
	return _ILBPair.Contract.GetProtocolFees(&_ILBPair.CallOpts)
}

// GetProtocolFees is a free data retrieval call binding the contract method 0xd8dfcea0.
//
// Solidity: function getProtocolFees() view returns(uint128 protocolFeeX, uint128 protocolFeeY)
func (_ILBPair *ILBPairCallerSession) GetProtocolFees() (struct {
	ProtocolFeeX *big.Int
	ProtocolFeeY *big.Int
}, error) {
	return _ILBPair.Contract.GetProtocolFees(&_ILBPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint128 reserveX, uint128 reserveY)
func (_ILBPair *ILBPairCaller) GetReserves(opts *bind.CallOpts) (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getReserves")

	outstruct := new(struct {
		ReserveX *big.Int
		ReserveY *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ReserveX = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ReserveY = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint128 reserveX, uint128 reserveY)
func (_ILBPair *ILBPairSession) GetReserves() (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	return _ILBPair.Contract.GetReserves(&_ILBPair.CallOpts)
}

// GetReserves is a free data retrieval call binding the contract method 0x0902f1ac.
//
// Solidity: function getReserves() view returns(uint128 reserveX, uint128 reserveY)
func (_ILBPair *ILBPairCallerSession) GetReserves() (struct {
	ReserveX *big.Int
	ReserveY *big.Int
}, error) {
	return _ILBPair.Contract.GetReserves(&_ILBPair.CallOpts)
}

// GetStaticFeeParameters is a free data retrieval call binding the contract method 0x7ca0de30.
//
// Solidity: function getStaticFeeParameters() view returns(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairCaller) GetStaticFeeParameters(opts *bind.CallOpts) (struct {
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       *big.Int
	ProtocolShare            uint16
	MaxVolatilityAccumulator *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getStaticFeeParameters")

	outstruct := new(struct {
		BaseFactor               uint16
		FilterPeriod             uint16
		DecayPeriod              uint16
		ReductionFactor          uint16
		VariableFeeControl       *big.Int
		ProtocolShare            uint16
		MaxVolatilityAccumulator *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseFactor = *abi.ConvertType(out[0], new(uint16)).(*uint16)
	outstruct.FilterPeriod = *abi.ConvertType(out[1], new(uint16)).(*uint16)
	outstruct.DecayPeriod = *abi.ConvertType(out[2], new(uint16)).(*uint16)
	outstruct.ReductionFactor = *abi.ConvertType(out[3], new(uint16)).(*uint16)
	outstruct.VariableFeeControl = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ProtocolShare = *abi.ConvertType(out[5], new(uint16)).(*uint16)
	outstruct.MaxVolatilityAccumulator = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetStaticFeeParameters is a free data retrieval call binding the contract method 0x7ca0de30.
//
// Solidity: function getStaticFeeParameters() view returns(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairSession) GetStaticFeeParameters() (struct {
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       *big.Int
	ProtocolShare            uint16
	MaxVolatilityAccumulator *big.Int
}, error) {
	return _ILBPair.Contract.GetStaticFeeParameters(&_ILBPair.CallOpts)
}

// GetStaticFeeParameters is a free data retrieval call binding the contract method 0x7ca0de30.
//
// Solidity: function getStaticFeeParameters() view returns(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairCallerSession) GetStaticFeeParameters() (struct {
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       *big.Int
	ProtocolShare            uint16
	MaxVolatilityAccumulator *big.Int
}, error) {
	return _ILBPair.Contract.GetStaticFeeParameters(&_ILBPair.CallOpts)
}

// GetSwapIn is a free data retrieval call binding the contract method 0xabcd7830.
//
// Solidity: function getSwapIn(uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBPair *ILBPairCaller) GetSwapIn(opts *bind.CallOpts, amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getSwapIn", amountOut, swapForY)

	outstruct := new(struct {
		AmountIn      *big.Int
		AmountOutLeft *big.Int
		Fee           *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountIn = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOutLeft = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapIn is a free data retrieval call binding the contract method 0xabcd7830.
//
// Solidity: function getSwapIn(uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBPair *ILBPairSession) GetSwapIn(amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _ILBPair.Contract.GetSwapIn(&_ILBPair.CallOpts, amountOut, swapForY)
}

// GetSwapIn is a free data retrieval call binding the contract method 0xabcd7830.
//
// Solidity: function getSwapIn(uint128 amountOut, bool swapForY) view returns(uint128 amountIn, uint128 amountOutLeft, uint128 fee)
func (_ILBPair *ILBPairCallerSession) GetSwapIn(amountOut *big.Int, swapForY bool) (struct {
	AmountIn      *big.Int
	AmountOutLeft *big.Int
	Fee           *big.Int
}, error) {
	return _ILBPair.Contract.GetSwapIn(&_ILBPair.CallOpts, amountOut, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xe77366f8.
//
// Solidity: function getSwapOut(uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBPair *ILBPairCaller) GetSwapOut(opts *bind.CallOpts, amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getSwapOut", amountIn, swapForY)

	outstruct := new(struct {
		AmountInLeft *big.Int
		AmountOut    *big.Int
		Fee          *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AmountInLeft = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.AmountOut = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Fee = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetSwapOut is a free data retrieval call binding the contract method 0xe77366f8.
//
// Solidity: function getSwapOut(uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBPair *ILBPairSession) GetSwapOut(amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _ILBPair.Contract.GetSwapOut(&_ILBPair.CallOpts, amountIn, swapForY)
}

// GetSwapOut is a free data retrieval call binding the contract method 0xe77366f8.
//
// Solidity: function getSwapOut(uint128 amountIn, bool swapForY) view returns(uint128 amountInLeft, uint128 amountOut, uint128 fee)
func (_ILBPair *ILBPairCallerSession) GetSwapOut(amountIn *big.Int, swapForY bool) (struct {
	AmountInLeft *big.Int
	AmountOut    *big.Int
	Fee          *big.Int
}, error) {
	return _ILBPair.Contract.GetSwapOut(&_ILBPair.CallOpts, amountIn, swapForY)
}

// GetTokenX is a free data retrieval call binding the contract method 0x05e8746d.
//
// Solidity: function getTokenX() view returns(address tokenX)
func (_ILBPair *ILBPairCaller) GetTokenX(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getTokenX")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenX is a free data retrieval call binding the contract method 0x05e8746d.
//
// Solidity: function getTokenX() view returns(address tokenX)
func (_ILBPair *ILBPairSession) GetTokenX() (common.Address, error) {
	return _ILBPair.Contract.GetTokenX(&_ILBPair.CallOpts)
}

// GetTokenX is a free data retrieval call binding the contract method 0x05e8746d.
//
// Solidity: function getTokenX() view returns(address tokenX)
func (_ILBPair *ILBPairCallerSession) GetTokenX() (common.Address, error) {
	return _ILBPair.Contract.GetTokenX(&_ILBPair.CallOpts)
}

// GetTokenY is a free data retrieval call binding the contract method 0xda10610c.
//
// Solidity: function getTokenY() view returns(address tokenY)
func (_ILBPair *ILBPairCaller) GetTokenY(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getTokenY")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetTokenY is a free data retrieval call binding the contract method 0xda10610c.
//
// Solidity: function getTokenY() view returns(address tokenY)
func (_ILBPair *ILBPairSession) GetTokenY() (common.Address, error) {
	return _ILBPair.Contract.GetTokenY(&_ILBPair.CallOpts)
}

// GetTokenY is a free data retrieval call binding the contract method 0xda10610c.
//
// Solidity: function getTokenY() view returns(address tokenY)
func (_ILBPair *ILBPairCallerSession) GetTokenY() (common.Address, error) {
	return _ILBPair.Contract.GetTokenY(&_ILBPair.CallOpts)
}

// GetVariableFeeParameters is a free data retrieval call binding the contract method 0x8d7024e5.
//
// Solidity: function getVariableFeeParameters() view returns(uint24 volatilityAccumulator, uint24 volatilityReference, uint24 idReference, uint40 timeOfLastUpdate)
func (_ILBPair *ILBPairCaller) GetVariableFeeParameters(opts *bind.CallOpts) (struct {
	VolatilityAccumulator *big.Int
	VolatilityReference   *big.Int
	IdReference           *big.Int
	TimeOfLastUpdate      *big.Int
}, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "getVariableFeeParameters")

	outstruct := new(struct {
		VolatilityAccumulator *big.Int
		VolatilityReference   *big.Int
		IdReference           *big.Int
		TimeOfLastUpdate      *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.VolatilityAccumulator = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.VolatilityReference = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.IdReference = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.TimeOfLastUpdate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetVariableFeeParameters is a free data retrieval call binding the contract method 0x8d7024e5.
//
// Solidity: function getVariableFeeParameters() view returns(uint24 volatilityAccumulator, uint24 volatilityReference, uint24 idReference, uint40 timeOfLastUpdate)
func (_ILBPair *ILBPairSession) GetVariableFeeParameters() (struct {
	VolatilityAccumulator *big.Int
	VolatilityReference   *big.Int
	IdReference           *big.Int
	TimeOfLastUpdate      *big.Int
}, error) {
	return _ILBPair.Contract.GetVariableFeeParameters(&_ILBPair.CallOpts)
}

// GetVariableFeeParameters is a free data retrieval call binding the contract method 0x8d7024e5.
//
// Solidity: function getVariableFeeParameters() view returns(uint24 volatilityAccumulator, uint24 volatilityReference, uint24 idReference, uint40 timeOfLastUpdate)
func (_ILBPair *ILBPairCallerSession) GetVariableFeeParameters() (struct {
	VolatilityAccumulator *big.Int
	VolatilityReference   *big.Int
	IdReference           *big.Int
	TimeOfLastUpdate      *big.Int
}, error) {
	return _ILBPair.Contract.GetVariableFeeParameters(&_ILBPair.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address spender) view returns(bool)
func (_ILBPair *ILBPairCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, spender common.Address) (bool, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "isApprovedForAll", owner, spender)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address spender) view returns(bool)
func (_ILBPair *ILBPairSession) IsApprovedForAll(owner common.Address, spender common.Address) (bool, error) {
	return _ILBPair.Contract.IsApprovedForAll(&_ILBPair.CallOpts, owner, spender)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address spender) view returns(bool)
func (_ILBPair *ILBPairCallerSession) IsApprovedForAll(owner common.Address, spender common.Address) (bool, error) {
	return _ILBPair.Contract.IsApprovedForAll(&_ILBPair.CallOpts, owner, spender)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILBPair *ILBPairCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILBPair *ILBPairSession) Name() (string, error) {
	return _ILBPair.Contract.Name(&_ILBPair.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_ILBPair *ILBPairCallerSession) Name() (string, error) {
	return _ILBPair.Contract.Name(&_ILBPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILBPair *ILBPairCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILBPair *ILBPairSession) Symbol() (string, error) {
	return _ILBPair.Contract.Symbol(&_ILBPair.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_ILBPair *ILBPairCallerSession) Symbol() (string, error) {
	return _ILBPair.Contract.Symbol(&_ILBPair.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ILBPair *ILBPairCaller) TotalSupply(opts *bind.CallOpts, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ILBPair.contract.Call(opts, &out, "totalSupply", id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ILBPair *ILBPairSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.TotalSupply(&_ILBPair.CallOpts, id)
}

// TotalSupply is a free data retrieval call binding the contract method 0xbd85b039.
//
// Solidity: function totalSupply(uint256 id) view returns(uint256)
func (_ILBPair *ILBPairCallerSession) TotalSupply(id *big.Int) (*big.Int, error) {
	return _ILBPair.Contract.TotalSupply(&_ILBPair.CallOpts, id)
}

// ApproveForAll is a paid mutator transaction binding the contract method 0xe584b654.
//
// Solidity: function approveForAll(address spender, bool approved) returns()
func (_ILBPair *ILBPairTransactor) ApproveForAll(opts *bind.TransactOpts, spender common.Address, approved bool) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "approveForAll", spender, approved)
}

// ApproveForAll is a paid mutator transaction binding the contract method 0xe584b654.
//
// Solidity: function approveForAll(address spender, bool approved) returns()
func (_ILBPair *ILBPairSession) ApproveForAll(spender common.Address, approved bool) (*types.Transaction, error) {
	return _ILBPair.Contract.ApproveForAll(&_ILBPair.TransactOpts, spender, approved)
}

// ApproveForAll is a paid mutator transaction binding the contract method 0xe584b654.
//
// Solidity: function approveForAll(address spender, bool approved) returns()
func (_ILBPair *ILBPairTransactorSession) ApproveForAll(spender common.Address, approved bool) (*types.Transaction, error) {
	return _ILBPair.Contract.ApproveForAll(&_ILBPair.TransactOpts, spender, approved)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_ILBPair *ILBPairTransactor) BatchTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "batchTransferFrom", from, to, ids, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_ILBPair *ILBPairSession) BatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.BatchTransferFrom(&_ILBPair.TransactOpts, from, to, ids, amounts)
}

// BatchTransferFrom is a paid mutator transaction binding the contract method 0x17fad7fc.
//
// Solidity: function batchTransferFrom(address from, address to, uint256[] ids, uint256[] amounts) returns()
func (_ILBPair *ILBPairTransactorSession) BatchTransferFrom(from common.Address, to common.Address, ids []*big.Int, amounts []*big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.BatchTransferFrom(&_ILBPair.TransactOpts, from, to, ids, amounts)
}

// Burn is a paid mutator transaction binding the contract method 0xc9939f5e.
//
// Solidity: function burn(address from, address to, uint256[] ids, uint256[] amountsToBurn) returns(bytes32[] amounts)
func (_ILBPair *ILBPairTransactor) Burn(opts *bind.TransactOpts, from common.Address, to common.Address, ids []*big.Int, amountsToBurn []*big.Int) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "burn", from, to, ids, amountsToBurn)
}

// Burn is a paid mutator transaction binding the contract method 0xc9939f5e.
//
// Solidity: function burn(address from, address to, uint256[] ids, uint256[] amountsToBurn) returns(bytes32[] amounts)
func (_ILBPair *ILBPairSession) Burn(from common.Address, to common.Address, ids []*big.Int, amountsToBurn []*big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.Burn(&_ILBPair.TransactOpts, from, to, ids, amountsToBurn)
}

// Burn is a paid mutator transaction binding the contract method 0xc9939f5e.
//
// Solidity: function burn(address from, address to, uint256[] ids, uint256[] amountsToBurn) returns(bytes32[] amounts)
func (_ILBPair *ILBPairTransactorSession) Burn(from common.Address, to common.Address, ids []*big.Int, amountsToBurn []*big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.Burn(&_ILBPair.TransactOpts, from, to, ids, amountsToBurn)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(bytes32 collectedProtocolFees)
func (_ILBPair *ILBPairTransactor) CollectProtocolFees(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "collectProtocolFees")
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(bytes32 collectedProtocolFees)
func (_ILBPair *ILBPairSession) CollectProtocolFees() (*types.Transaction, error) {
	return _ILBPair.Contract.CollectProtocolFees(&_ILBPair.TransactOpts)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0xa1af5b9a.
//
// Solidity: function collectProtocolFees() returns(bytes32 collectedProtocolFees)
func (_ILBPair *ILBPairTransactorSession) CollectProtocolFees() (*types.Transaction, error) {
	return _ILBPair.Contract.CollectProtocolFees(&_ILBPair.TransactOpts)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xea3446bf.
//
// Solidity: function flashLoan(address receiver, bytes32 amounts, bytes data) returns()
func (_ILBPair *ILBPairTransactor) FlashLoan(opts *bind.TransactOpts, receiver common.Address, amounts [32]byte, data []byte) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "flashLoan", receiver, amounts, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xea3446bf.
//
// Solidity: function flashLoan(address receiver, bytes32 amounts, bytes data) returns()
func (_ILBPair *ILBPairSession) FlashLoan(receiver common.Address, amounts [32]byte, data []byte) (*types.Transaction, error) {
	return _ILBPair.Contract.FlashLoan(&_ILBPair.TransactOpts, receiver, amounts, data)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xea3446bf.
//
// Solidity: function flashLoan(address receiver, bytes32 amounts, bytes data) returns()
func (_ILBPair *ILBPairTransactorSession) FlashLoan(receiver common.Address, amounts [32]byte, data []byte) (*types.Transaction, error) {
	return _ILBPair.Contract.FlashLoan(&_ILBPair.TransactOpts, receiver, amounts, data)
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_ILBPair *ILBPairTransactor) ForceDecay(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "forceDecay")
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_ILBPair *ILBPairSession) ForceDecay() (*types.Transaction, error) {
	return _ILBPair.Contract.ForceDecay(&_ILBPair.TransactOpts)
}

// ForceDecay is a paid mutator transaction binding the contract method 0xd3b9fbe4.
//
// Solidity: function forceDecay() returns()
func (_ILBPair *ILBPairTransactorSession) ForceDecay() (*types.Transaction, error) {
	return _ILBPair.Contract.ForceDecay(&_ILBPair.TransactOpts)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 newLength) returns()
func (_ILBPair *ILBPairTransactor) IncreaseOracleLength(opts *bind.TransactOpts, newLength uint16) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "increaseOracleLength", newLength)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 newLength) returns()
func (_ILBPair *ILBPairSession) IncreaseOracleLength(newLength uint16) (*types.Transaction, error) {
	return _ILBPair.Contract.IncreaseOracleLength(&_ILBPair.TransactOpts, newLength)
}

// IncreaseOracleLength is a paid mutator transaction binding the contract method 0xc7bd6586.
//
// Solidity: function increaseOracleLength(uint16 newLength) returns()
func (_ILBPair *ILBPairTransactorSession) IncreaseOracleLength(newLength uint16) (*types.Transaction, error) {
	return _ILBPair.Contract.IncreaseOracleLength(&_ILBPair.TransactOpts, newLength)
}

// Initialize is a paid mutator transaction binding the contract method 0x47973bff.
//
// Solidity: function initialize(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, uint24 activeId) returns()
func (_ILBPair *ILBPairTransactor) Initialize(opts *bind.TransactOpts, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, activeId *big.Int) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "initialize", baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, activeId)
}

// Initialize is a paid mutator transaction binding the contract method 0x47973bff.
//
// Solidity: function initialize(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, uint24 activeId) returns()
func (_ILBPair *ILBPairSession) Initialize(baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, activeId *big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.Initialize(&_ILBPair.TransactOpts, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, activeId)
}

// Initialize is a paid mutator transaction binding the contract method 0x47973bff.
//
// Solidity: function initialize(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, uint24 activeId) returns()
func (_ILBPair *ILBPairTransactorSession) Initialize(baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, activeId *big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.Initialize(&_ILBPair.TransactOpts, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, activeId)
}

// Mint is a paid mutator transaction binding the contract method 0x383d15c5.
//
// Solidity: function mint(address to, bytes32[] liquidityConfigs, address refundTo) returns(bytes32 amountsReceived, bytes32 amountsLeft, uint256[] liquidityMinted)
func (_ILBPair *ILBPairTransactor) Mint(opts *bind.TransactOpts, to common.Address, liquidityConfigs [][32]byte, refundTo common.Address) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "mint", to, liquidityConfigs, refundTo)
}

// Mint is a paid mutator transaction binding the contract method 0x383d15c5.
//
// Solidity: function mint(address to, bytes32[] liquidityConfigs, address refundTo) returns(bytes32 amountsReceived, bytes32 amountsLeft, uint256[] liquidityMinted)
func (_ILBPair *ILBPairSession) Mint(to common.Address, liquidityConfigs [][32]byte, refundTo common.Address) (*types.Transaction, error) {
	return _ILBPair.Contract.Mint(&_ILBPair.TransactOpts, to, liquidityConfigs, refundTo)
}

// Mint is a paid mutator transaction binding the contract method 0x383d15c5.
//
// Solidity: function mint(address to, bytes32[] liquidityConfigs, address refundTo) returns(bytes32 amountsReceived, bytes32 amountsLeft, uint256[] liquidityMinted)
func (_ILBPair *ILBPairTransactorSession) Mint(to common.Address, liquidityConfigs [][32]byte, refundTo common.Address) (*types.Transaction, error) {
	return _ILBPair.Contract.Mint(&_ILBPair.TransactOpts, to, liquidityConfigs, refundTo)
}

// SetStaticFeeParameters is a paid mutator transaction binding the contract method 0x6653851a.
//
// Solidity: function setStaticFeeParameters(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_ILBPair *ILBPairTransactor) SetStaticFeeParameters(opts *bind.TransactOpts, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "setStaticFeeParameters", baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// SetStaticFeeParameters is a paid mutator transaction binding the contract method 0x6653851a.
//
// Solidity: function setStaticFeeParameters(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_ILBPair *ILBPairSession) SetStaticFeeParameters(baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.SetStaticFeeParameters(&_ILBPair.TransactOpts, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// SetStaticFeeParameters is a paid mutator transaction binding the contract method 0x6653851a.
//
// Solidity: function setStaticFeeParameters(uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_ILBPair *ILBPairTransactorSession) SetStaticFeeParameters(baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _ILBPair.Contract.SetStaticFeeParameters(&_ILBPair.TransactOpts, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool swapForY, address to) returns(bytes32 amountsOut)
func (_ILBPair *ILBPairTransactor) Swap(opts *bind.TransactOpts, swapForY bool, to common.Address) (*types.Transaction, error) {
	return _ILBPair.contract.Transact(opts, "swap", swapForY, to)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool swapForY, address to) returns(bytes32 amountsOut)
func (_ILBPair *ILBPairSession) Swap(swapForY bool, to common.Address) (*types.Transaction, error) {
	return _ILBPair.Contract.Swap(&_ILBPair.TransactOpts, swapForY, to)
}

// Swap is a paid mutator transaction binding the contract method 0x53c059a0.
//
// Solidity: function swap(bool swapForY, address to) returns(bytes32 amountsOut)
func (_ILBPair *ILBPairTransactorSession) Swap(swapForY bool, to common.Address) (*types.Transaction, error) {
	return _ILBPair.Contract.Swap(&_ILBPair.TransactOpts, swapForY, to)
}

// ILBPairApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the ILBPair contract.
type ILBPairApprovalForAllIterator struct {
	Event *ILBPairApprovalForAll // Event containing the contract specifics and raw log

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
func (it *ILBPairApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairApprovalForAll)
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
		it.Event = new(ILBPairApprovalForAll)
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
func (it *ILBPairApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairApprovalForAll represents a ApprovalForAll event raised by the ILBPair contract.
type ILBPairApprovalForAll struct {
	Account  common.Address
	Sender   common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_ILBPair *ILBPairFilterer) FilterApprovalForAll(opts *bind.FilterOpts, account []common.Address, sender []common.Address) (*ILBPairApprovalForAllIterator, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "ApprovalForAll", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairApprovalForAllIterator{contract: _ILBPair.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_ILBPair *ILBPairFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *ILBPairApprovalForAll, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "ApprovalForAll", accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairApprovalForAll)
				if err := _ILBPair.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
// Solidity: event ApprovalForAll(address indexed account, address indexed sender, bool approved)
func (_ILBPair *ILBPairFilterer) ParseApprovalForAll(log types.Log) (*ILBPairApprovalForAll, error) {
	event := new(ILBPairApprovalForAll)
	if err := _ILBPair.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairCollectedProtocolFeesIterator is returned from FilterCollectedProtocolFees and is used to iterate over the raw logs and unpacked data for CollectedProtocolFees events raised by the ILBPair contract.
type ILBPairCollectedProtocolFeesIterator struct {
	Event *ILBPairCollectedProtocolFees // Event containing the contract specifics and raw log

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
func (it *ILBPairCollectedProtocolFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairCollectedProtocolFees)
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
		it.Event = new(ILBPairCollectedProtocolFees)
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
func (it *ILBPairCollectedProtocolFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairCollectedProtocolFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairCollectedProtocolFees represents a CollectedProtocolFees event raised by the ILBPair contract.
type ILBPairCollectedProtocolFees struct {
	FeeRecipient common.Address
	ProtocolFees [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCollectedProtocolFees is a free log retrieval operation binding the contract event 0x3f41a5ddc53701cc7db577ade4f1fca9838a8ec0b5ea50b9f0f5d17bc4554e32.
//
// Solidity: event CollectedProtocolFees(address indexed feeRecipient, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) FilterCollectedProtocolFees(opts *bind.FilterOpts, feeRecipient []common.Address) (*ILBPairCollectedProtocolFeesIterator, error) {

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "CollectedProtocolFees", feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairCollectedProtocolFeesIterator{contract: _ILBPair.contract, event: "CollectedProtocolFees", logs: logs, sub: sub}, nil
}

// WatchCollectedProtocolFees is a free log subscription operation binding the contract event 0x3f41a5ddc53701cc7db577ade4f1fca9838a8ec0b5ea50b9f0f5d17bc4554e32.
//
// Solidity: event CollectedProtocolFees(address indexed feeRecipient, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) WatchCollectedProtocolFees(opts *bind.WatchOpts, sink chan<- *ILBPairCollectedProtocolFees, feeRecipient []common.Address) (event.Subscription, error) {

	var feeRecipientRule []interface{}
	for _, feeRecipientItem := range feeRecipient {
		feeRecipientRule = append(feeRecipientRule, feeRecipientItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "CollectedProtocolFees", feeRecipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairCollectedProtocolFees)
				if err := _ILBPair.contract.UnpackLog(event, "CollectedProtocolFees", log); err != nil {
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

// ParseCollectedProtocolFees is a log parse operation binding the contract event 0x3f41a5ddc53701cc7db577ade4f1fca9838a8ec0b5ea50b9f0f5d17bc4554e32.
//
// Solidity: event CollectedProtocolFees(address indexed feeRecipient, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) ParseCollectedProtocolFees(log types.Log) (*ILBPairCollectedProtocolFees, error) {
	event := new(ILBPairCollectedProtocolFees)
	if err := _ILBPair.contract.UnpackLog(event, "CollectedProtocolFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairCompositionFeesIterator is returned from FilterCompositionFees and is used to iterate over the raw logs and unpacked data for CompositionFees events raised by the ILBPair contract.
type ILBPairCompositionFeesIterator struct {
	Event *ILBPairCompositionFees // Event containing the contract specifics and raw log

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
func (it *ILBPairCompositionFeesIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairCompositionFees)
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
		it.Event = new(ILBPairCompositionFees)
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
func (it *ILBPairCompositionFeesIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairCompositionFeesIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairCompositionFees represents a CompositionFees event raised by the ILBPair contract.
type ILBPairCompositionFees struct {
	Sender       common.Address
	Id           *big.Int
	TotalFees    [32]byte
	ProtocolFees [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterCompositionFees is a free log retrieval operation binding the contract event 0x3f0b46725027bb418b2005f4683538eccdbcdf1de2b8649a29dbd9c507d16ff4.
//
// Solidity: event CompositionFees(address indexed sender, uint24 id, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) FilterCompositionFees(opts *bind.FilterOpts, sender []common.Address) (*ILBPairCompositionFeesIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "CompositionFees", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairCompositionFeesIterator{contract: _ILBPair.contract, event: "CompositionFees", logs: logs, sub: sub}, nil
}

// WatchCompositionFees is a free log subscription operation binding the contract event 0x3f0b46725027bb418b2005f4683538eccdbcdf1de2b8649a29dbd9c507d16ff4.
//
// Solidity: event CompositionFees(address indexed sender, uint24 id, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) WatchCompositionFees(opts *bind.WatchOpts, sink chan<- *ILBPairCompositionFees, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "CompositionFees", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairCompositionFees)
				if err := _ILBPair.contract.UnpackLog(event, "CompositionFees", log); err != nil {
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

// ParseCompositionFees is a log parse operation binding the contract event 0x3f0b46725027bb418b2005f4683538eccdbcdf1de2b8649a29dbd9c507d16ff4.
//
// Solidity: event CompositionFees(address indexed sender, uint24 id, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) ParseCompositionFees(log types.Log) (*ILBPairCompositionFees, error) {
	event := new(ILBPairCompositionFees)
	if err := _ILBPair.contract.UnpackLog(event, "CompositionFees", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairDepositedToBinsIterator is returned from FilterDepositedToBins and is used to iterate over the raw logs and unpacked data for DepositedToBins events raised by the ILBPair contract.
type ILBPairDepositedToBinsIterator struct {
	Event *ILBPairDepositedToBins // Event containing the contract specifics and raw log

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
func (it *ILBPairDepositedToBinsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairDepositedToBins)
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
		it.Event = new(ILBPairDepositedToBins)
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
func (it *ILBPairDepositedToBinsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairDepositedToBinsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairDepositedToBins represents a DepositedToBins event raised by the ILBPair contract.
type ILBPairDepositedToBins struct {
	Sender  common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts [][32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDepositedToBins is a free log retrieval operation binding the contract event 0x87f1f9dcf5e8089a3e00811b6a008d8f30293a3da878cb1fe8c90ca376402f8a.
//
// Solidity: event DepositedToBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) FilterDepositedToBins(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ILBPairDepositedToBinsIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "DepositedToBins", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairDepositedToBinsIterator{contract: _ILBPair.contract, event: "DepositedToBins", logs: logs, sub: sub}, nil
}

// WatchDepositedToBins is a free log subscription operation binding the contract event 0x87f1f9dcf5e8089a3e00811b6a008d8f30293a3da878cb1fe8c90ca376402f8a.
//
// Solidity: event DepositedToBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) WatchDepositedToBins(opts *bind.WatchOpts, sink chan<- *ILBPairDepositedToBins, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "DepositedToBins", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairDepositedToBins)
				if err := _ILBPair.contract.UnpackLog(event, "DepositedToBins", log); err != nil {
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

// ParseDepositedToBins is a log parse operation binding the contract event 0x87f1f9dcf5e8089a3e00811b6a008d8f30293a3da878cb1fe8c90ca376402f8a.
//
// Solidity: event DepositedToBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) ParseDepositedToBins(log types.Log) (*ILBPairDepositedToBins, error) {
	event := new(ILBPairDepositedToBins)
	if err := _ILBPair.contract.UnpackLog(event, "DepositedToBins", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the ILBPair contract.
type ILBPairFlashLoanIterator struct {
	Event *ILBPairFlashLoan // Event containing the contract specifics and raw log

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
func (it *ILBPairFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairFlashLoan)
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
		it.Event = new(ILBPairFlashLoan)
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
func (it *ILBPairFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairFlashLoan represents a FlashLoan event raised by the ILBPair contract.
type ILBPairFlashLoan struct {
	Sender       common.Address
	Receiver     common.Address
	ActiveId     *big.Int
	Amounts      [32]byte
	TotalFees    [32]byte
	ProtocolFees [32]byte
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0xd126bd9d94daca8e55ffd8283fac05394aec8326c6b1639e1e8a445fbe8bbc7d.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, uint24 activeId, bytes32 amounts, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) FilterFlashLoan(opts *bind.FilterOpts, sender []common.Address, receiver []common.Address) (*ILBPairFlashLoanIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "FlashLoan", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairFlashLoanIterator{contract: _ILBPair.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0xd126bd9d94daca8e55ffd8283fac05394aec8326c6b1639e1e8a445fbe8bbc7d.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, uint24 activeId, bytes32 amounts, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ILBPairFlashLoan, sender []common.Address, receiver []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var receiverRule []interface{}
	for _, receiverItem := range receiver {
		receiverRule = append(receiverRule, receiverItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "FlashLoan", senderRule, receiverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairFlashLoan)
				if err := _ILBPair.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0xd126bd9d94daca8e55ffd8283fac05394aec8326c6b1639e1e8a445fbe8bbc7d.
//
// Solidity: event FlashLoan(address indexed sender, address indexed receiver, uint24 activeId, bytes32 amounts, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) ParseFlashLoan(log types.Log) (*ILBPairFlashLoan, error) {
	event := new(ILBPairFlashLoan)
	if err := _ILBPair.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairForcedDecayIterator is returned from FilterForcedDecay and is used to iterate over the raw logs and unpacked data for ForcedDecay events raised by the ILBPair contract.
type ILBPairForcedDecayIterator struct {
	Event *ILBPairForcedDecay // Event containing the contract specifics and raw log

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
func (it *ILBPairForcedDecayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairForcedDecay)
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
		it.Event = new(ILBPairForcedDecay)
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
func (it *ILBPairForcedDecayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairForcedDecayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairForcedDecay represents a ForcedDecay event raised by the ILBPair contract.
type ILBPairForcedDecay struct {
	Sender              common.Address
	IdReference         *big.Int
	VolatilityReference *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterForcedDecay is a free log retrieval operation binding the contract event 0x282afaeeae84c1d85ad1424a3aa2ddbedaeefca3b1e53d889d15265fe44db7fc.
//
// Solidity: event ForcedDecay(address indexed sender, uint24 idReference, uint24 volatilityReference)
func (_ILBPair *ILBPairFilterer) FilterForcedDecay(opts *bind.FilterOpts, sender []common.Address) (*ILBPairForcedDecayIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "ForcedDecay", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairForcedDecayIterator{contract: _ILBPair.contract, event: "ForcedDecay", logs: logs, sub: sub}, nil
}

// WatchForcedDecay is a free log subscription operation binding the contract event 0x282afaeeae84c1d85ad1424a3aa2ddbedaeefca3b1e53d889d15265fe44db7fc.
//
// Solidity: event ForcedDecay(address indexed sender, uint24 idReference, uint24 volatilityReference)
func (_ILBPair *ILBPairFilterer) WatchForcedDecay(opts *bind.WatchOpts, sink chan<- *ILBPairForcedDecay, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "ForcedDecay", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairForcedDecay)
				if err := _ILBPair.contract.UnpackLog(event, "ForcedDecay", log); err != nil {
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

// ParseForcedDecay is a log parse operation binding the contract event 0x282afaeeae84c1d85ad1424a3aa2ddbedaeefca3b1e53d889d15265fe44db7fc.
//
// Solidity: event ForcedDecay(address indexed sender, uint24 idReference, uint24 volatilityReference)
func (_ILBPair *ILBPairFilterer) ParseForcedDecay(log types.Log) (*ILBPairForcedDecay, error) {
	event := new(ILBPairForcedDecay)
	if err := _ILBPair.contract.UnpackLog(event, "ForcedDecay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairOracleLengthIncreasedIterator is returned from FilterOracleLengthIncreased and is used to iterate over the raw logs and unpacked data for OracleLengthIncreased events raised by the ILBPair contract.
type ILBPairOracleLengthIncreasedIterator struct {
	Event *ILBPairOracleLengthIncreased // Event containing the contract specifics and raw log

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
func (it *ILBPairOracleLengthIncreasedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairOracleLengthIncreased)
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
		it.Event = new(ILBPairOracleLengthIncreased)
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
func (it *ILBPairOracleLengthIncreasedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairOracleLengthIncreasedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairOracleLengthIncreased represents a OracleLengthIncreased event raised by the ILBPair contract.
type ILBPairOracleLengthIncreased struct {
	Sender       common.Address
	OracleLength uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOracleLengthIncreased is a free log retrieval operation binding the contract event 0xc975541e72d695746a43ba65745d79963a23082637c8f4609354d9bcf70194d6.
//
// Solidity: event OracleLengthIncreased(address indexed sender, uint16 oracleLength)
func (_ILBPair *ILBPairFilterer) FilterOracleLengthIncreased(opts *bind.FilterOpts, sender []common.Address) (*ILBPairOracleLengthIncreasedIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "OracleLengthIncreased", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairOracleLengthIncreasedIterator{contract: _ILBPair.contract, event: "OracleLengthIncreased", logs: logs, sub: sub}, nil
}

// WatchOracleLengthIncreased is a free log subscription operation binding the contract event 0xc975541e72d695746a43ba65745d79963a23082637c8f4609354d9bcf70194d6.
//
// Solidity: event OracleLengthIncreased(address indexed sender, uint16 oracleLength)
func (_ILBPair *ILBPairFilterer) WatchOracleLengthIncreased(opts *bind.WatchOpts, sink chan<- *ILBPairOracleLengthIncreased, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "OracleLengthIncreased", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairOracleLengthIncreased)
				if err := _ILBPair.contract.UnpackLog(event, "OracleLengthIncreased", log); err != nil {
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

// ParseOracleLengthIncreased is a log parse operation binding the contract event 0xc975541e72d695746a43ba65745d79963a23082637c8f4609354d9bcf70194d6.
//
// Solidity: event OracleLengthIncreased(address indexed sender, uint16 oracleLength)
func (_ILBPair *ILBPairFilterer) ParseOracleLengthIncreased(log types.Log) (*ILBPairOracleLengthIncreased, error) {
	event := new(ILBPairOracleLengthIncreased)
	if err := _ILBPair.contract.UnpackLog(event, "OracleLengthIncreased", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairStaticFeeParametersSetIterator is returned from FilterStaticFeeParametersSet and is used to iterate over the raw logs and unpacked data for StaticFeeParametersSet events raised by the ILBPair contract.
type ILBPairStaticFeeParametersSetIterator struct {
	Event *ILBPairStaticFeeParametersSet // Event containing the contract specifics and raw log

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
func (it *ILBPairStaticFeeParametersSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairStaticFeeParametersSet)
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
		it.Event = new(ILBPairStaticFeeParametersSet)
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
func (it *ILBPairStaticFeeParametersSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairStaticFeeParametersSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairStaticFeeParametersSet represents a StaticFeeParametersSet event raised by the ILBPair contract.
type ILBPairStaticFeeParametersSet struct {
	Sender                   common.Address
	BaseFactor               uint16
	FilterPeriod             uint16
	DecayPeriod              uint16
	ReductionFactor          uint16
	VariableFeeControl       *big.Int
	ProtocolShare            uint16
	MaxVolatilityAccumulator *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterStaticFeeParametersSet is a free log retrieval operation binding the contract event 0xd09e5ddc721ff14c5c1e66a305cbba1fd70b82c5232bc391aad6f55e62e4b046.
//
// Solidity: event StaticFeeParametersSet(address indexed sender, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairFilterer) FilterStaticFeeParametersSet(opts *bind.FilterOpts, sender []common.Address) (*ILBPairStaticFeeParametersSetIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "StaticFeeParametersSet", senderRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairStaticFeeParametersSetIterator{contract: _ILBPair.contract, event: "StaticFeeParametersSet", logs: logs, sub: sub}, nil
}

// WatchStaticFeeParametersSet is a free log subscription operation binding the contract event 0xd09e5ddc721ff14c5c1e66a305cbba1fd70b82c5232bc391aad6f55e62e4b046.
//
// Solidity: event StaticFeeParametersSet(address indexed sender, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairFilterer) WatchStaticFeeParametersSet(opts *bind.WatchOpts, sink chan<- *ILBPairStaticFeeParametersSet, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "StaticFeeParametersSet", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairStaticFeeParametersSet)
				if err := _ILBPair.contract.UnpackLog(event, "StaticFeeParametersSet", log); err != nil {
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

// ParseStaticFeeParametersSet is a log parse operation binding the contract event 0xd09e5ddc721ff14c5c1e66a305cbba1fd70b82c5232bc391aad6f55e62e4b046.
//
// Solidity: event StaticFeeParametersSet(address indexed sender, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator)
func (_ILBPair *ILBPairFilterer) ParseStaticFeeParametersSet(log types.Log) (*ILBPairStaticFeeParametersSet, error) {
	event := new(ILBPairStaticFeeParametersSet)
	if err := _ILBPair.contract.UnpackLog(event, "StaticFeeParametersSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ILBPair contract.
type ILBPairSwapIterator struct {
	Event *ILBPairSwap // Event containing the contract specifics and raw log

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
func (it *ILBPairSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairSwap)
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
		it.Event = new(ILBPairSwap)
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
func (it *ILBPairSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairSwap represents a Swap event raised by the ILBPair contract.
type ILBPairSwap struct {
	Sender                common.Address
	To                    common.Address
	Id                    *big.Int
	AmountsIn             [32]byte
	AmountsOut            [32]byte
	VolatilityAccumulator *big.Int
	TotalFees             [32]byte
	ProtocolFees          [32]byte
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xad7d6f97abf51ce18e17a38f4d70e975be9c0708474987bb3e26ad21bd93ca70.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint24 id, bytes32 amountsIn, bytes32 amountsOut, uint24 volatilityAccumulator, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) FilterSwap(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ILBPairSwapIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairSwapIterator{contract: _ILBPair.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xad7d6f97abf51ce18e17a38f4d70e975be9c0708474987bb3e26ad21bd93ca70.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint24 id, bytes32 amountsIn, bytes32 amountsOut, uint24 volatilityAccumulator, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ILBPairSwap, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "Swap", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairSwap)
				if err := _ILBPair.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xad7d6f97abf51ce18e17a38f4d70e975be9c0708474987bb3e26ad21bd93ca70.
//
// Solidity: event Swap(address indexed sender, address indexed to, uint24 id, bytes32 amountsIn, bytes32 amountsOut, uint24 volatilityAccumulator, bytes32 totalFees, bytes32 protocolFees)
func (_ILBPair *ILBPairFilterer) ParseSwap(log types.Log) (*ILBPairSwap, error) {
	event := new(ILBPairSwap)
	if err := _ILBPair.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairTransferBatchIterator is returned from FilterTransferBatch and is used to iterate over the raw logs and unpacked data for TransferBatch events raised by the ILBPair contract.
type ILBPairTransferBatchIterator struct {
	Event *ILBPairTransferBatch // Event containing the contract specifics and raw log

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
func (it *ILBPairTransferBatchIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairTransferBatch)
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
		it.Event = new(ILBPairTransferBatch)
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
func (it *ILBPairTransferBatchIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairTransferBatchIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairTransferBatch represents a TransferBatch event raised by the ILBPair contract.
type ILBPairTransferBatch struct {
	Sender  common.Address
	From    common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts []*big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransferBatch is a free log retrieval operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_ILBPair *ILBPairFilterer) FilterTransferBatch(opts *bind.FilterOpts, sender []common.Address, from []common.Address, to []common.Address) (*ILBPairTransferBatchIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "TransferBatch", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairTransferBatchIterator{contract: _ILBPair.contract, event: "TransferBatch", logs: logs, sub: sub}, nil
}

// WatchTransferBatch is a free log subscription operation binding the contract event 0x4a39dc06d4c0dbc64b70af90fd698a233a518aa5d07e595d983b8c0526c8f7fb.
//
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_ILBPair *ILBPairFilterer) WatchTransferBatch(opts *bind.WatchOpts, sink chan<- *ILBPairTransferBatch, sender []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "TransferBatch", senderRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairTransferBatch)
				if err := _ILBPair.contract.UnpackLog(event, "TransferBatch", log); err != nil {
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
// Solidity: event TransferBatch(address indexed sender, address indexed from, address indexed to, uint256[] ids, uint256[] amounts)
func (_ILBPair *ILBPairFilterer) ParseTransferBatch(log types.Log) (*ILBPairTransferBatch, error) {
	event := new(ILBPairTransferBatch)
	if err := _ILBPair.contract.UnpackLog(event, "TransferBatch", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILBPairWithdrawnFromBinsIterator is returned from FilterWithdrawnFromBins and is used to iterate over the raw logs and unpacked data for WithdrawnFromBins events raised by the ILBPair contract.
type ILBPairWithdrawnFromBinsIterator struct {
	Event *ILBPairWithdrawnFromBins // Event containing the contract specifics and raw log

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
func (it *ILBPairWithdrawnFromBinsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILBPairWithdrawnFromBins)
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
		it.Event = new(ILBPairWithdrawnFromBins)
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
func (it *ILBPairWithdrawnFromBinsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILBPairWithdrawnFromBinsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILBPairWithdrawnFromBins represents a WithdrawnFromBins event raised by the ILBPair contract.
type ILBPairWithdrawnFromBins struct {
	Sender  common.Address
	To      common.Address
	Ids     []*big.Int
	Amounts [][32]byte
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdrawnFromBins is a free log retrieval operation binding the contract event 0xa32e146844d6144a22e94c586715a1317d58a8aa3581ec33d040113ddcb24350.
//
// Solidity: event WithdrawnFromBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) FilterWithdrawnFromBins(opts *bind.FilterOpts, sender []common.Address, to []common.Address) (*ILBPairWithdrawnFromBinsIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.FilterLogs(opts, "WithdrawnFromBins", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILBPairWithdrawnFromBinsIterator{contract: _ILBPair.contract, event: "WithdrawnFromBins", logs: logs, sub: sub}, nil
}

// WatchWithdrawnFromBins is a free log subscription operation binding the contract event 0xa32e146844d6144a22e94c586715a1317d58a8aa3581ec33d040113ddcb24350.
//
// Solidity: event WithdrawnFromBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) WatchWithdrawnFromBins(opts *bind.WatchOpts, sink chan<- *ILBPairWithdrawnFromBins, sender []common.Address, to []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILBPair.contract.WatchLogs(opts, "WithdrawnFromBins", senderRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILBPairWithdrawnFromBins)
				if err := _ILBPair.contract.UnpackLog(event, "WithdrawnFromBins", log); err != nil {
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

// ParseWithdrawnFromBins is a log parse operation binding the contract event 0xa32e146844d6144a22e94c586715a1317d58a8aa3581ec33d040113ddcb24350.
//
// Solidity: event WithdrawnFromBins(address indexed sender, address indexed to, uint256[] ids, bytes32[] amounts)
func (_ILBPair *ILBPairFilterer) ParseWithdrawnFromBins(log types.Log) (*ILBPairWithdrawnFromBins, error) {
	event := new(ILBPairWithdrawnFromBins)
	if err := _ILBPair.contract.UnpackLog(event, "WithdrawnFromBins", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
