// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package traderjoe

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

// ILBFactoryLBPairInformation is an auto generated low-level Go binding around an user-defined struct.
type ILBFactoryLBPairInformation struct {
	BinStep           uint16
	LBPair            common.Address
	CreatedByOwner    bool
	IgnoredForRouting bool
}

// TraderjoeMetaData contains all meta data concerning the Traderjoe contract.
var TraderjoeMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"LBFactory__AddressZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepHasNoPreset\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__BinStepTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fees\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxFees\",\"type\":\"uint256\"}],\"name\":\"LBFactory__FlashLoanFeeAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"LBFactory__IdenticalAddresses\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__ImplementationNotSet\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairAlreadyExists\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairDoesNotExist\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__LBPairIgnoredIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__LBPairNotCreated\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__LBPairSafetyCheckFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"LBFactory__PresetIsLockedForUsers\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"LBFactory__PresetOpenStateIsAlreadyInTheSameState\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetAlreadyWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"LBFactory__QuoteAssetNotWhitelisted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"LBFactory__SameFeeRecipient\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"LBFactory__SameFlashLoanFee\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBFactory__SameImplementation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PairParametersHelper__InvalidParameter\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__AddressZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NoPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__NotPendingOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PendingOwnable__PendingOwnerAlreadySet\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SafeCast__Exceeds16Bits\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"x\",\"type\":\"uint256\"},{\"internalType\":\"int256\",\"name\":\"y\",\"type\":\"int256\"}],\"name\":\"Uint128x128Math__PowUnderflow\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldRecipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newRecipient\",\"type\":\"address\"}],\"name\":\"FeeRecipientSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"oldFlashLoanFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newFlashLoanFee\",\"type\":\"uint256\"}],\"name\":\"FlashLoanFeeSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"pid\",\"type\":\"uint256\"}],\"name\":\"LBPairCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"ignored\",\"type\":\"bool\"}],\"name\":\"LBPairIgnoredStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldLBPairImplementation\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"LBPairImplementation\",\"type\":\"address\"}],\"name\":\"LBPairImplementationSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pendingOwner\",\"type\":\"address\"}],\"name\":\"PendingOwnerSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"PresetOpenStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"PresetRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint256\"}],\"name\":\"PresetSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"QuoteAssetRemoved\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"addQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"becomeOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"activeId\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"createLBPair\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"pair\",\"type\":\"address\"}],\"name\":\"forceDecay\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllBinSteps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"binStepWithPreset\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"}],\"name\":\"getAllLBPairs\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation[]\",\"name\":\"lbPairsAvailable\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFeeRecipient\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFlashLoanFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getLBPairAtIndex\",\"outputs\":[{\"internalType\":\"contractILBPair\",\"name\":\"lbPair\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLBPairImplementation\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lbPairImplementation\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenA\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenB\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"getLBPairInformation\",\"outputs\":[{\"components\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"contractILBPair\",\"name\":\"LBPair\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"createdByOwner\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"ignoredForRouting\",\"type\":\"bool\"}],\"internalType\":\"structILBFactory.LBPairInformation\",\"name\":\"lbPairInformation\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMaxFlashLoanFee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"maxFee\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinBinStep\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"minBinStep\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfLBPairs\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lbPairNumber\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNumberOfQuoteAssets\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"numberOfQuoteAssets\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getOpenBinSteps\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"openBinStep\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"binStep\",\"type\":\"uint256\"}],\"name\":\"getPreset\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"baseFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"filterPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decayPeriod\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reductionFactor\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableFeeControl\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"protocolShare\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getQuoteAssetAtIndex\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"asset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"isQuoteAsset\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isQuote\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pendingOwner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"}],\"name\":\"removePreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"quoteAsset\",\"type\":\"address\"}],\"name\":\"removeQuoteAsset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"revokePendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"feeRecipient\",\"type\":\"address\"}],\"name\":\"setFeeRecipient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"}],\"name\":\"setFeesParametersOnPair\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"flashLoanFee\",\"type\":\"uint256\"}],\"name\":\"setFlashLoanFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenX\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenY\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"ignored\",\"type\":\"bool\"}],\"name\":\"setLBPairIgnored\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newLBPairImplementation\",\"type\":\"address\"}],\"name\":\"setLBPairImplementation\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pendingOwner_\",\"type\":\"address\"}],\"name\":\"setPendingOwner\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"baseFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"filterPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"decayPeriod\",\"type\":\"uint16\"},{\"internalType\":\"uint16\",\"name\":\"reductionFactor\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"variableFeeControl\",\"type\":\"uint24\"},{\"internalType\":\"uint16\",\"name\":\"protocolShare\",\"type\":\"uint16\"},{\"internalType\":\"uint24\",\"name\":\"maxVolatilityAccumulator\",\"type\":\"uint24\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"setPreset\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint16\",\"name\":\"binStep\",\"type\":\"uint16\"},{\"internalType\":\"bool\",\"name\":\"isOpen\",\"type\":\"bool\"}],\"name\":\"setPresetOpenState\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TraderjoeABI is the input ABI used to generate the binding from.
// Deprecated: Use TraderjoeMetaData.ABI instead.
var TraderjoeABI = TraderjoeMetaData.ABI

// Traderjoe is an auto generated Go binding around an Ethereum contract.
type Traderjoe struct {
	TraderjoeCaller     // Read-only binding to the contract
	TraderjoeTransactor // Write-only binding to the contract
	TraderjoeFilterer   // Log filterer for contract events
}

// TraderjoeCaller is an auto generated read-only Go binding around an Ethereum contract.
type TraderjoeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TraderjoeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TraderjoeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TraderjoeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TraderjoeSession struct {
	Contract     *Traderjoe        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TraderjoeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TraderjoeCallerSession struct {
	Contract *TraderjoeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TraderjoeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TraderjoeTransactorSession struct {
	Contract     *TraderjoeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TraderjoeRaw is an auto generated low-level Go binding around an Ethereum contract.
type TraderjoeRaw struct {
	Contract *Traderjoe // Generic contract binding to access the raw methods on
}

// TraderjoeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TraderjoeCallerRaw struct {
	Contract *TraderjoeCaller // Generic read-only contract binding to access the raw methods on
}

// TraderjoeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TraderjoeTransactorRaw struct {
	Contract *TraderjoeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTraderjoe creates a new instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoe(address common.Address, backend bind.ContractBackend) (*Traderjoe, error) {
	contract, err := bindTraderjoe(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Traderjoe{TraderjoeCaller: TraderjoeCaller{contract: contract}, TraderjoeTransactor: TraderjoeTransactor{contract: contract}, TraderjoeFilterer: TraderjoeFilterer{contract: contract}}, nil
}

// NewTraderjoeCaller creates a new read-only instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeCaller(address common.Address, caller bind.ContractCaller) (*TraderjoeCaller, error) {
	contract, err := bindTraderjoe(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeCaller{contract: contract}, nil
}

// NewTraderjoeTransactor creates a new write-only instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeTransactor(address common.Address, transactor bind.ContractTransactor) (*TraderjoeTransactor, error) {
	contract, err := bindTraderjoe(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TraderjoeTransactor{contract: contract}, nil
}

// NewTraderjoeFilterer creates a new log filterer instance of Traderjoe, bound to a specific deployed contract.
func NewTraderjoeFilterer(address common.Address, filterer bind.ContractFilterer) (*TraderjoeFilterer, error) {
	contract, err := bindTraderjoe(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TraderjoeFilterer{contract: contract}, nil
}

// bindTraderjoe binds a generic wrapper to an already deployed contract.
func bindTraderjoe(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TraderjoeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traderjoe *TraderjoeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traderjoe.Contract.TraderjoeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traderjoe *TraderjoeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.Contract.TraderjoeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traderjoe *TraderjoeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traderjoe.Contract.TraderjoeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Traderjoe *TraderjoeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Traderjoe.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Traderjoe *TraderjoeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Traderjoe *TraderjoeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Traderjoe.Contract.contract.Transact(opts, method, params...)
}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (_Traderjoe *TraderjoeCaller) GetAllBinSteps(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAllBinSteps")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (_Traderjoe *TraderjoeSession) GetAllBinSteps() ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAllBinSteps(&_Traderjoe.CallOpts)
}

// GetAllBinSteps is a free data retrieval call binding the contract method 0x5b35875c.
//
// Solidity: function getAllBinSteps() view returns(uint256[] binStepWithPreset)
func (_Traderjoe *TraderjoeCallerSession) GetAllBinSteps() ([]*big.Int, error) {
	return _Traderjoe.Contract.GetAllBinSteps(&_Traderjoe.CallOpts)
}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (_Traderjoe *TraderjoeCaller) GetAllLBPairs(opts *bind.CallOpts, tokenX common.Address, tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getAllLBPairs", tokenX, tokenY)

	if err != nil {
		return *new([]ILBFactoryLBPairInformation), err
	}

	out0 := *abi.ConvertType(out[0], new([]ILBFactoryLBPairInformation)).(*[]ILBFactoryLBPairInformation)

	return out0, err

}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (_Traderjoe *TraderjoeSession) GetAllLBPairs(tokenX common.Address, tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	return _Traderjoe.Contract.GetAllLBPairs(&_Traderjoe.CallOpts, tokenX, tokenY)
}

// GetAllLBPairs is a free data retrieval call binding the contract method 0x6622e0d7.
//
// Solidity: function getAllLBPairs(address tokenX, address tokenY) view returns((uint16,address,bool,bool)[] lbPairsAvailable)
func (_Traderjoe *TraderjoeCallerSession) GetAllLBPairs(tokenX common.Address, tokenY common.Address) ([]ILBFactoryLBPairInformation, error) {
	return _Traderjoe.Contract.GetAllLBPairs(&_Traderjoe.CallOpts, tokenX, tokenY)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (_Traderjoe *TraderjoeCaller) GetFeeRecipient(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getFeeRecipient")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (_Traderjoe *TraderjoeSession) GetFeeRecipient() (common.Address, error) {
	return _Traderjoe.Contract.GetFeeRecipient(&_Traderjoe.CallOpts)
}

// GetFeeRecipient is a free data retrieval call binding the contract method 0x4ccb20c0.
//
// Solidity: function getFeeRecipient() view returns(address feeRecipient)
func (_Traderjoe *TraderjoeCallerSession) GetFeeRecipient() (common.Address, error) {
	return _Traderjoe.Contract.GetFeeRecipient(&_Traderjoe.CallOpts)
}

// GetFlashLoanFee is a free data retrieval call binding the contract method 0xfd90c2be.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (_Traderjoe *TraderjoeCaller) GetFlashLoanFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getFlashLoanFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetFlashLoanFee is a free data retrieval call binding the contract method 0xfd90c2be.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (_Traderjoe *TraderjoeSession) GetFlashLoanFee() (*big.Int, error) {
	return _Traderjoe.Contract.GetFlashLoanFee(&_Traderjoe.CallOpts)
}

// GetFlashLoanFee is a free data retrieval call binding the contract method 0xfd90c2be.
//
// Solidity: function getFlashLoanFee() view returns(uint256 flashLoanFee)
func (_Traderjoe *TraderjoeCallerSession) GetFlashLoanFee() (*big.Int, error) {
	return _Traderjoe.Contract.GetFlashLoanFee(&_Traderjoe.CallOpts)
}

// GetLBPairAtIndex is a free data retrieval call binding the contract method 0x7daf5d66.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (_Traderjoe *TraderjoeCaller) GetLBPairAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getLBPairAtIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLBPairAtIndex is a free data retrieval call binding the contract method 0x7daf5d66.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (_Traderjoe *TraderjoeSession) GetLBPairAtIndex(index *big.Int) (common.Address, error) {
	return _Traderjoe.Contract.GetLBPairAtIndex(&_Traderjoe.CallOpts, index)
}

// GetLBPairAtIndex is a free data retrieval call binding the contract method 0x7daf5d66.
//
// Solidity: function getLBPairAtIndex(uint256 index) view returns(address lbPair)
func (_Traderjoe *TraderjoeCallerSession) GetLBPairAtIndex(index *big.Int) (common.Address, error) {
	return _Traderjoe.Contract.GetLBPairAtIndex(&_Traderjoe.CallOpts, index)
}

// GetLBPairImplementation is a free data retrieval call binding the contract method 0xaf371065.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (_Traderjoe *TraderjoeCaller) GetLBPairImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getLBPairImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLBPairImplementation is a free data retrieval call binding the contract method 0xaf371065.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (_Traderjoe *TraderjoeSession) GetLBPairImplementation() (common.Address, error) {
	return _Traderjoe.Contract.GetLBPairImplementation(&_Traderjoe.CallOpts)
}

// GetLBPairImplementation is a free data retrieval call binding the contract method 0xaf371065.
//
// Solidity: function getLBPairImplementation() view returns(address lbPairImplementation)
func (_Traderjoe *TraderjoeCallerSession) GetLBPairImplementation() (common.Address, error) {
	return _Traderjoe.Contract.GetLBPairImplementation(&_Traderjoe.CallOpts)
}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (_Traderjoe *TraderjoeCaller) GetLBPairInformation(opts *bind.CallOpts, tokenA common.Address, tokenB common.Address, binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getLBPairInformation", tokenA, tokenB, binStep)

	if err != nil {
		return *new(ILBFactoryLBPairInformation), err
	}

	out0 := *abi.ConvertType(out[0], new(ILBFactoryLBPairInformation)).(*ILBFactoryLBPairInformation)

	return out0, err

}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (_Traderjoe *TraderjoeSession) GetLBPairInformation(tokenA common.Address, tokenB common.Address, binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	return _Traderjoe.Contract.GetLBPairInformation(&_Traderjoe.CallOpts, tokenA, tokenB, binStep)
}

// GetLBPairInformation is a free data retrieval call binding the contract method 0x704037bd.
//
// Solidity: function getLBPairInformation(address tokenA, address tokenB, uint256 binStep) view returns((uint16,address,bool,bool) lbPairInformation)
func (_Traderjoe *TraderjoeCallerSession) GetLBPairInformation(tokenA common.Address, tokenB common.Address, binStep *big.Int) (ILBFactoryLBPairInformation, error) {
	return _Traderjoe.Contract.GetLBPairInformation(&_Traderjoe.CallOpts, tokenA, tokenB, binStep)
}

// GetMaxFlashLoanFee is a free data retrieval call binding the contract method 0x8ce9aa1c.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (_Traderjoe *TraderjoeCaller) GetMaxFlashLoanFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getMaxFlashLoanFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMaxFlashLoanFee is a free data retrieval call binding the contract method 0x8ce9aa1c.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (_Traderjoe *TraderjoeSession) GetMaxFlashLoanFee() (*big.Int, error) {
	return _Traderjoe.Contract.GetMaxFlashLoanFee(&_Traderjoe.CallOpts)
}

// GetMaxFlashLoanFee is a free data retrieval call binding the contract method 0x8ce9aa1c.
//
// Solidity: function getMaxFlashLoanFee() pure returns(uint256 maxFee)
func (_Traderjoe *TraderjoeCallerSession) GetMaxFlashLoanFee() (*big.Int, error) {
	return _Traderjoe.Contract.GetMaxFlashLoanFee(&_Traderjoe.CallOpts)
}

// GetMinBinStep is a free data retrieval call binding the contract method 0x701ab8c1.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (_Traderjoe *TraderjoeCaller) GetMinBinStep(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getMinBinStep")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinBinStep is a free data retrieval call binding the contract method 0x701ab8c1.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (_Traderjoe *TraderjoeSession) GetMinBinStep() (*big.Int, error) {
	return _Traderjoe.Contract.GetMinBinStep(&_Traderjoe.CallOpts)
}

// GetMinBinStep is a free data retrieval call binding the contract method 0x701ab8c1.
//
// Solidity: function getMinBinStep() pure returns(uint256 minBinStep)
func (_Traderjoe *TraderjoeCallerSession) GetMinBinStep() (*big.Int, error) {
	return _Traderjoe.Contract.GetMinBinStep(&_Traderjoe.CallOpts)
}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (_Traderjoe *TraderjoeCaller) GetNumberOfLBPairs(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getNumberOfLBPairs")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (_Traderjoe *TraderjoeSession) GetNumberOfLBPairs() (*big.Int, error) {
	return _Traderjoe.Contract.GetNumberOfLBPairs(&_Traderjoe.CallOpts)
}

// GetNumberOfLBPairs is a free data retrieval call binding the contract method 0x4e937c3a.
//
// Solidity: function getNumberOfLBPairs() view returns(uint256 lbPairNumber)
func (_Traderjoe *TraderjoeCallerSession) GetNumberOfLBPairs() (*big.Int, error) {
	return _Traderjoe.Contract.GetNumberOfLBPairs(&_Traderjoe.CallOpts)
}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (_Traderjoe *TraderjoeCaller) GetNumberOfQuoteAssets(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getNumberOfQuoteAssets")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (_Traderjoe *TraderjoeSession) GetNumberOfQuoteAssets() (*big.Int, error) {
	return _Traderjoe.Contract.GetNumberOfQuoteAssets(&_Traderjoe.CallOpts)
}

// GetNumberOfQuoteAssets is a free data retrieval call binding the contract method 0x80c5061e.
//
// Solidity: function getNumberOfQuoteAssets() view returns(uint256 numberOfQuoteAssets)
func (_Traderjoe *TraderjoeCallerSession) GetNumberOfQuoteAssets() (*big.Int, error) {
	return _Traderjoe.Contract.GetNumberOfQuoteAssets(&_Traderjoe.CallOpts)
}

// GetOpenBinSteps is a free data retrieval call binding the contract method 0x0282c9c1.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (_Traderjoe *TraderjoeCaller) GetOpenBinSteps(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getOpenBinSteps")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetOpenBinSteps is a free data retrieval call binding the contract method 0x0282c9c1.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (_Traderjoe *TraderjoeSession) GetOpenBinSteps() ([]*big.Int, error) {
	return _Traderjoe.Contract.GetOpenBinSteps(&_Traderjoe.CallOpts)
}

// GetOpenBinSteps is a free data retrieval call binding the contract method 0x0282c9c1.
//
// Solidity: function getOpenBinSteps() view returns(uint256[] openBinStep)
func (_Traderjoe *TraderjoeCallerSession) GetOpenBinSteps() ([]*big.Int, error) {
	return _Traderjoe.Contract.GetOpenBinSteps(&_Traderjoe.CallOpts)
}

// GetPreset is a free data retrieval call binding the contract method 0xaabc4b3c.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (_Traderjoe *TraderjoeCaller) GetPreset(opts *bind.CallOpts, binStep *big.Int) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	IsOpen                   bool
}, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getPreset", binStep)

	outstruct := new(struct {
		BaseFactor               *big.Int
		FilterPeriod             *big.Int
		DecayPeriod              *big.Int
		ReductionFactor          *big.Int
		VariableFeeControl       *big.Int
		ProtocolShare            *big.Int
		MaxVolatilityAccumulator *big.Int
		IsOpen                   bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.BaseFactor = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.FilterPeriod = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.DecayPeriod = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.ReductionFactor = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VariableFeeControl = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.ProtocolShare = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.MaxVolatilityAccumulator = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.IsOpen = *abi.ConvertType(out[7], new(bool)).(*bool)

	return *outstruct, err

}

// GetPreset is a free data retrieval call binding the contract method 0xaabc4b3c.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (_Traderjoe *TraderjoeSession) GetPreset(binStep *big.Int) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	IsOpen                   bool
}, error) {
	return _Traderjoe.Contract.GetPreset(&_Traderjoe.CallOpts, binStep)
}

// GetPreset is a free data retrieval call binding the contract method 0xaabc4b3c.
//
// Solidity: function getPreset(uint256 binStep) view returns(uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator, bool isOpen)
func (_Traderjoe *TraderjoeCallerSession) GetPreset(binStep *big.Int) (struct {
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	IsOpen                   bool
}, error) {
	return _Traderjoe.Contract.GetPreset(&_Traderjoe.CallOpts, binStep)
}

// GetQuoteAssetAtIndex is a free data retrieval call binding the contract method 0x0752092b.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (_Traderjoe *TraderjoeCaller) GetQuoteAssetAtIndex(opts *bind.CallOpts, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "getQuoteAssetAtIndex", index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetQuoteAssetAtIndex is a free data retrieval call binding the contract method 0x0752092b.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (_Traderjoe *TraderjoeSession) GetQuoteAssetAtIndex(index *big.Int) (common.Address, error) {
	return _Traderjoe.Contract.GetQuoteAssetAtIndex(&_Traderjoe.CallOpts, index)
}

// GetQuoteAssetAtIndex is a free data retrieval call binding the contract method 0x0752092b.
//
// Solidity: function getQuoteAssetAtIndex(uint256 index) view returns(address asset)
func (_Traderjoe *TraderjoeCallerSession) GetQuoteAssetAtIndex(index *big.Int) (common.Address, error) {
	return _Traderjoe.Contract.GetQuoteAssetAtIndex(&_Traderjoe.CallOpts, index)
}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (_Traderjoe *TraderjoeCaller) IsQuoteAsset(opts *bind.CallOpts, token common.Address) (bool, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "isQuoteAsset", token)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (_Traderjoe *TraderjoeSession) IsQuoteAsset(token common.Address) (bool, error) {
	return _Traderjoe.Contract.IsQuoteAsset(&_Traderjoe.CallOpts, token)
}

// IsQuoteAsset is a free data retrieval call binding the contract method 0x27721842.
//
// Solidity: function isQuoteAsset(address token) view returns(bool isQuote)
func (_Traderjoe *TraderjoeCallerSession) IsQuoteAsset(token common.Address) (bool, error) {
	return _Traderjoe.Contract.IsQuoteAsset(&_Traderjoe.CallOpts, token)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Traderjoe *TraderjoeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Traderjoe *TraderjoeSession) Owner() (common.Address, error) {
	return _Traderjoe.Contract.Owner(&_Traderjoe.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Traderjoe *TraderjoeCallerSession) Owner() (common.Address, error) {
	return _Traderjoe.Contract.Owner(&_Traderjoe.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Traderjoe *TraderjoeCaller) PendingOwner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Traderjoe.contract.Call(opts, &out, "pendingOwner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Traderjoe *TraderjoeSession) PendingOwner() (common.Address, error) {
	return _Traderjoe.Contract.PendingOwner(&_Traderjoe.CallOpts)
}

// PendingOwner is a free data retrieval call binding the contract method 0xe30c3978.
//
// Solidity: function pendingOwner() view returns(address)
func (_Traderjoe *TraderjoeCallerSession) PendingOwner() (common.Address, error) {
	return _Traderjoe.Contract.PendingOwner(&_Traderjoe.CallOpts)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeTransactor) AddQuoteAsset(opts *bind.TransactOpts, quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "addQuoteAsset", quoteAsset)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeSession) AddQuoteAsset(quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddQuoteAsset(&_Traderjoe.TransactOpts, quoteAsset)
}

// AddQuoteAsset is a paid mutator transaction binding the contract method 0x5a440923.
//
// Solidity: function addQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeTransactorSession) AddQuoteAsset(quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.AddQuoteAsset(&_Traderjoe.TransactOpts, quoteAsset)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Traderjoe *TraderjoeTransactor) BecomeOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "becomeOwner")
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Traderjoe *TraderjoeSession) BecomeOwner() (*types.Transaction, error) {
	return _Traderjoe.Contract.BecomeOwner(&_Traderjoe.TransactOpts)
}

// BecomeOwner is a paid mutator transaction binding the contract method 0xf9dca989.
//
// Solidity: function becomeOwner() returns()
func (_Traderjoe *TraderjoeTransactorSession) BecomeOwner() (*types.Transaction, error) {
	return _Traderjoe.Contract.BecomeOwner(&_Traderjoe.TransactOpts)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_Traderjoe *TraderjoeTransactor) CreateLBPair(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "createLBPair", tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_Traderjoe *TraderjoeSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.Contract.CreateLBPair(&_Traderjoe.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// CreateLBPair is a paid mutator transaction binding the contract method 0x659ac74b.
//
// Solidity: function createLBPair(address tokenX, address tokenY, uint24 activeId, uint16 binStep) returns(address pair)
func (_Traderjoe *TraderjoeTransactorSession) CreateLBPair(tokenX common.Address, tokenY common.Address, activeId *big.Int, binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.Contract.CreateLBPair(&_Traderjoe.TransactOpts, tokenX, tokenY, activeId, binStep)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address pair) returns()
func (_Traderjoe *TraderjoeTransactor) ForceDecay(opts *bind.TransactOpts, pair common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "forceDecay", pair)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address pair) returns()
func (_Traderjoe *TraderjoeSession) ForceDecay(pair common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.ForceDecay(&_Traderjoe.TransactOpts, pair)
}

// ForceDecay is a paid mutator transaction binding the contract method 0x3c78a941.
//
// Solidity: function forceDecay(address pair) returns()
func (_Traderjoe *TraderjoeTransactorSession) ForceDecay(pair common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.ForceDecay(&_Traderjoe.TransactOpts, pair)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 binStep) returns()
func (_Traderjoe *TraderjoeTransactor) RemovePreset(opts *bind.TransactOpts, binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removePreset", binStep)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 binStep) returns()
func (_Traderjoe *TraderjoeSession) RemovePreset(binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemovePreset(&_Traderjoe.TransactOpts, binStep)
}

// RemovePreset is a paid mutator transaction binding the contract method 0xe203a31f.
//
// Solidity: function removePreset(uint16 binStep) returns()
func (_Traderjoe *TraderjoeTransactorSession) RemovePreset(binStep uint16) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemovePreset(&_Traderjoe.TransactOpts, binStep)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeTransactor) RemoveQuoteAsset(opts *bind.TransactOpts, quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "removeQuoteAsset", quoteAsset)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeSession) RemoveQuoteAsset(quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveQuoteAsset(&_Traderjoe.TransactOpts, quoteAsset)
}

// RemoveQuoteAsset is a paid mutator transaction binding the contract method 0xddbfd941.
//
// Solidity: function removeQuoteAsset(address quoteAsset) returns()
func (_Traderjoe *TraderjoeTransactorSession) RemoveQuoteAsset(quoteAsset common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.RemoveQuoteAsset(&_Traderjoe.TransactOpts, quoteAsset)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Traderjoe *TraderjoeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Traderjoe *TraderjoeSession) RenounceOwnership() (*types.Transaction, error) {
	return _Traderjoe.Contract.RenounceOwnership(&_Traderjoe.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Traderjoe *TraderjoeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Traderjoe.Contract.RenounceOwnership(&_Traderjoe.TransactOpts)
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Traderjoe *TraderjoeTransactor) RevokePendingOwner(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "revokePendingOwner")
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Traderjoe *TraderjoeSession) RevokePendingOwner() (*types.Transaction, error) {
	return _Traderjoe.Contract.RevokePendingOwner(&_Traderjoe.TransactOpts)
}

// RevokePendingOwner is a paid mutator transaction binding the contract method 0x67ab8a4e.
//
// Solidity: function revokePendingOwner() returns()
func (_Traderjoe *TraderjoeTransactorSession) RevokePendingOwner() (*types.Transaction, error) {
	return _Traderjoe.Contract.RevokePendingOwner(&_Traderjoe.TransactOpts)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Traderjoe *TraderjoeTransactor) SetFeeRecipient(opts *bind.TransactOpts, feeRecipient common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setFeeRecipient", feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Traderjoe *TraderjoeSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFeeRecipient(&_Traderjoe.TransactOpts, feeRecipient)
}

// SetFeeRecipient is a paid mutator transaction binding the contract method 0xe74b981b.
//
// Solidity: function setFeeRecipient(address feeRecipient) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetFeeRecipient(feeRecipient common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFeeRecipient(&_Traderjoe.TransactOpts, feeRecipient)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address tokenX, address tokenY, uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_Traderjoe *TraderjoeTransactor) SetFeesParametersOnPair(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setFeesParametersOnPair", tokenX, tokenY, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address tokenX, address tokenY, uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_Traderjoe *TraderjoeSession) SetFeesParametersOnPair(tokenX common.Address, tokenY common.Address, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFeesParametersOnPair(&_Traderjoe.TransactOpts, tokenX, tokenY, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// SetFeesParametersOnPair is a paid mutator transaction binding the contract method 0x093ff769.
//
// Solidity: function setFeesParametersOnPair(address tokenX, address tokenY, uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetFeesParametersOnPair(tokenX common.Address, tokenY common.Address, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFeesParametersOnPair(&_Traderjoe.TransactOpts, tokenX, tokenY, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 flashLoanFee) returns()
func (_Traderjoe *TraderjoeTransactor) SetFlashLoanFee(opts *bind.TransactOpts, flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setFlashLoanFee", flashLoanFee)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 flashLoanFee) returns()
func (_Traderjoe *TraderjoeSession) SetFlashLoanFee(flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFlashLoanFee(&_Traderjoe.TransactOpts, flashLoanFee)
}

// SetFlashLoanFee is a paid mutator transaction binding the contract method 0xe92d0d5d.
//
// Solidity: function setFlashLoanFee(uint256 flashLoanFee) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetFlashLoanFee(flashLoanFee *big.Int) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetFlashLoanFee(&_Traderjoe.TransactOpts, flashLoanFee)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x69d56ea3.
//
// Solidity: function setLBPairIgnored(address tokenX, address tokenY, uint16 binStep, bool ignored) returns()
func (_Traderjoe *TraderjoeTransactor) SetLBPairIgnored(opts *bind.TransactOpts, tokenX common.Address, tokenY common.Address, binStep uint16, ignored bool) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setLBPairIgnored", tokenX, tokenY, binStep, ignored)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x69d56ea3.
//
// Solidity: function setLBPairIgnored(address tokenX, address tokenY, uint16 binStep, bool ignored) returns()
func (_Traderjoe *TraderjoeSession) SetLBPairIgnored(tokenX common.Address, tokenY common.Address, binStep uint16, ignored bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetLBPairIgnored(&_Traderjoe.TransactOpts, tokenX, tokenY, binStep, ignored)
}

// SetLBPairIgnored is a paid mutator transaction binding the contract method 0x69d56ea3.
//
// Solidity: function setLBPairIgnored(address tokenX, address tokenY, uint16 binStep, bool ignored) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetLBPairIgnored(tokenX common.Address, tokenY common.Address, binStep uint16, ignored bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetLBPairIgnored(&_Traderjoe.TransactOpts, tokenX, tokenY, binStep, ignored)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address newLBPairImplementation) returns()
func (_Traderjoe *TraderjoeTransactor) SetLBPairImplementation(opts *bind.TransactOpts, newLBPairImplementation common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setLBPairImplementation", newLBPairImplementation)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address newLBPairImplementation) returns()
func (_Traderjoe *TraderjoeSession) SetLBPairImplementation(newLBPairImplementation common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetLBPairImplementation(&_Traderjoe.TransactOpts, newLBPairImplementation)
}

// SetLBPairImplementation is a paid mutator transaction binding the contract method 0xb0384781.
//
// Solidity: function setLBPairImplementation(address newLBPairImplementation) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetLBPairImplementation(newLBPairImplementation common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetLBPairImplementation(&_Traderjoe.TransactOpts, newLBPairImplementation)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Traderjoe *TraderjoeTransactor) SetPendingOwner(opts *bind.TransactOpts, pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setPendingOwner", pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Traderjoe *TraderjoeSession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPendingOwner(&_Traderjoe.TransactOpts, pendingOwner_)
}

// SetPendingOwner is a paid mutator transaction binding the contract method 0xc42069ec.
//
// Solidity: function setPendingOwner(address pendingOwner_) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetPendingOwner(pendingOwner_ common.Address) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPendingOwner(&_Traderjoe.TransactOpts, pendingOwner_)
}

// SetPreset is a paid mutator transaction binding the contract method 0x379ee803.
//
// Solidity: function setPreset(uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, bool isOpen) returns()
func (_Traderjoe *TraderjoeTransactor) SetPreset(opts *bind.TransactOpts, binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setPreset", binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, isOpen)
}

// SetPreset is a paid mutator transaction binding the contract method 0x379ee803.
//
// Solidity: function setPreset(uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, bool isOpen) returns()
func (_Traderjoe *TraderjoeSession) SetPreset(binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPreset(&_Traderjoe.TransactOpts, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, isOpen)
}

// SetPreset is a paid mutator transaction binding the contract method 0x379ee803.
//
// Solidity: function setPreset(uint16 binStep, uint16 baseFactor, uint16 filterPeriod, uint16 decayPeriod, uint16 reductionFactor, uint24 variableFeeControl, uint16 protocolShare, uint24 maxVolatilityAccumulator, bool isOpen) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetPreset(binStep uint16, baseFactor uint16, filterPeriod uint16, decayPeriod uint16, reductionFactor uint16, variableFeeControl *big.Int, protocolShare uint16, maxVolatilityAccumulator *big.Int, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPreset(&_Traderjoe.TransactOpts, binStep, baseFactor, filterPeriod, decayPeriod, reductionFactor, variableFeeControl, protocolShare, maxVolatilityAccumulator, isOpen)
}

// SetPresetOpenState is a paid mutator transaction binding the contract method 0x4cd161d3.
//
// Solidity: function setPresetOpenState(uint16 binStep, bool isOpen) returns()
func (_Traderjoe *TraderjoeTransactor) SetPresetOpenState(opts *bind.TransactOpts, binStep uint16, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.contract.Transact(opts, "setPresetOpenState", binStep, isOpen)
}

// SetPresetOpenState is a paid mutator transaction binding the contract method 0x4cd161d3.
//
// Solidity: function setPresetOpenState(uint16 binStep, bool isOpen) returns()
func (_Traderjoe *TraderjoeSession) SetPresetOpenState(binStep uint16, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPresetOpenState(&_Traderjoe.TransactOpts, binStep, isOpen)
}

// SetPresetOpenState is a paid mutator transaction binding the contract method 0x4cd161d3.
//
// Solidity: function setPresetOpenState(uint16 binStep, bool isOpen) returns()
func (_Traderjoe *TraderjoeTransactorSession) SetPresetOpenState(binStep uint16, isOpen bool) (*types.Transaction, error) {
	return _Traderjoe.Contract.SetPresetOpenState(&_Traderjoe.TransactOpts, binStep, isOpen)
}

// TraderjoeFeeRecipientSetIterator is returned from FilterFeeRecipientSet and is used to iterate over the raw logs and unpacked data for FeeRecipientSet events raised by the Traderjoe contract.
type TraderjoeFeeRecipientSetIterator struct {
	Event *TraderjoeFeeRecipientSet // Event containing the contract specifics and raw log

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
func (it *TraderjoeFeeRecipientSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeFeeRecipientSet)
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
		it.Event = new(TraderjoeFeeRecipientSet)
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
func (it *TraderjoeFeeRecipientSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeFeeRecipientSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeFeeRecipientSet represents a FeeRecipientSet event raised by the Traderjoe contract.
type TraderjoeFeeRecipientSet struct {
	OldRecipient common.Address
	NewRecipient common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeRecipientSet is a free log retrieval operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Traderjoe *TraderjoeFilterer) FilterFeeRecipientSet(opts *bind.FilterOpts) (*TraderjoeFeeRecipientSetIterator, error) {

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return &TraderjoeFeeRecipientSetIterator{contract: _Traderjoe.contract, event: "FeeRecipientSet", logs: logs, sub: sub}, nil
}

// WatchFeeRecipientSet is a free log subscription operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Traderjoe *TraderjoeFilterer) WatchFeeRecipientSet(opts *bind.WatchOpts, sink chan<- *TraderjoeFeeRecipientSet) (event.Subscription, error) {

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "FeeRecipientSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeFeeRecipientSet)
				if err := _Traderjoe.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
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

// ParseFeeRecipientSet is a log parse operation binding the contract event 0x15d80a013f22151bc7246e3bc132e12828cde19de98870475e3fa70840152721.
//
// Solidity: event FeeRecipientSet(address oldRecipient, address newRecipient)
func (_Traderjoe *TraderjoeFilterer) ParseFeeRecipientSet(log types.Log) (*TraderjoeFeeRecipientSet, error) {
	event := new(TraderjoeFeeRecipientSet)
	if err := _Traderjoe.contract.UnpackLog(event, "FeeRecipientSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeFlashLoanFeeSetIterator is returned from FilterFlashLoanFeeSet and is used to iterate over the raw logs and unpacked data for FlashLoanFeeSet events raised by the Traderjoe contract.
type TraderjoeFlashLoanFeeSetIterator struct {
	Event *TraderjoeFlashLoanFeeSet // Event containing the contract specifics and raw log

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
func (it *TraderjoeFlashLoanFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeFlashLoanFeeSet)
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
		it.Event = new(TraderjoeFlashLoanFeeSet)
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
func (it *TraderjoeFlashLoanFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeFlashLoanFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeFlashLoanFeeSet represents a FlashLoanFeeSet event raised by the Traderjoe contract.
type TraderjoeFlashLoanFeeSet struct {
	OldFlashLoanFee *big.Int
	NewFlashLoanFee *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterFlashLoanFeeSet is a free log retrieval operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Traderjoe *TraderjoeFilterer) FilterFlashLoanFeeSet(opts *bind.FilterOpts) (*TraderjoeFlashLoanFeeSetIterator, error) {

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "FlashLoanFeeSet")
	if err != nil {
		return nil, err
	}
	return &TraderjoeFlashLoanFeeSetIterator{contract: _Traderjoe.contract, event: "FlashLoanFeeSet", logs: logs, sub: sub}, nil
}

// WatchFlashLoanFeeSet is a free log subscription operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Traderjoe *TraderjoeFilterer) WatchFlashLoanFeeSet(opts *bind.WatchOpts, sink chan<- *TraderjoeFlashLoanFeeSet) (event.Subscription, error) {

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "FlashLoanFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeFlashLoanFeeSet)
				if err := _Traderjoe.contract.UnpackLog(event, "FlashLoanFeeSet", log); err != nil {
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

// ParseFlashLoanFeeSet is a log parse operation binding the contract event 0x5c34e91c94c78b662a45d0bd4a25a4e32c584c54a45a76e4a4d43be27ba40e50.
//
// Solidity: event FlashLoanFeeSet(uint256 oldFlashLoanFee, uint256 newFlashLoanFee)
func (_Traderjoe *TraderjoeFilterer) ParseFlashLoanFeeSet(log types.Log) (*TraderjoeFlashLoanFeeSet, error) {
	event := new(TraderjoeFlashLoanFeeSet)
	if err := _Traderjoe.contract.UnpackLog(event, "FlashLoanFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeLBPairCreatedIterator is returned from FilterLBPairCreated and is used to iterate over the raw logs and unpacked data for LBPairCreated events raised by the Traderjoe contract.
type TraderjoeLBPairCreatedIterator struct {
	Event *TraderjoeLBPairCreated // Event containing the contract specifics and raw log

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
func (it *TraderjoeLBPairCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeLBPairCreated)
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
		it.Event = new(TraderjoeLBPairCreated)
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
func (it *TraderjoeLBPairCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeLBPairCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeLBPairCreated represents a LBPairCreated event raised by the Traderjoe contract.
type TraderjoeLBPairCreated struct {
	TokenX  common.Address
	TokenY  common.Address
	BinStep *big.Int
	LBPair  common.Address
	Pid     *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLBPairCreated is a free log retrieval operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Traderjoe *TraderjoeFilterer) FilterLBPairCreated(opts *bind.FilterOpts, tokenX []common.Address, tokenY []common.Address, binStep []*big.Int) (*TraderjoeLBPairCreatedIterator, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "LBPairCreated", tokenXRule, tokenYRule, binStepRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeLBPairCreatedIterator{contract: _Traderjoe.contract, event: "LBPairCreated", logs: logs, sub: sub}, nil
}

// WatchLBPairCreated is a free log subscription operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Traderjoe *TraderjoeFilterer) WatchLBPairCreated(opts *bind.WatchOpts, sink chan<- *TraderjoeLBPairCreated, tokenX []common.Address, tokenY []common.Address, binStep []*big.Int) (event.Subscription, error) {

	var tokenXRule []interface{}
	for _, tokenXItem := range tokenX {
		tokenXRule = append(tokenXRule, tokenXItem)
	}
	var tokenYRule []interface{}
	for _, tokenYItem := range tokenY {
		tokenYRule = append(tokenYRule, tokenYItem)
	}
	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "LBPairCreated", tokenXRule, tokenYRule, binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeLBPairCreated)
				if err := _Traderjoe.contract.UnpackLog(event, "LBPairCreated", log); err != nil {
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

// ParseLBPairCreated is a log parse operation binding the contract event 0x2c8d104b27c6b7f4492017a6f5cf3803043688934ebcaa6a03540beeaf976aff.
//
// Solidity: event LBPairCreated(address indexed tokenX, address indexed tokenY, uint256 indexed binStep, address LBPair, uint256 pid)
func (_Traderjoe *TraderjoeFilterer) ParseLBPairCreated(log types.Log) (*TraderjoeLBPairCreated, error) {
	event := new(TraderjoeLBPairCreated)
	if err := _Traderjoe.contract.UnpackLog(event, "LBPairCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeLBPairIgnoredStateChangedIterator is returned from FilterLBPairIgnoredStateChanged and is used to iterate over the raw logs and unpacked data for LBPairIgnoredStateChanged events raised by the Traderjoe contract.
type TraderjoeLBPairIgnoredStateChangedIterator struct {
	Event *TraderjoeLBPairIgnoredStateChanged // Event containing the contract specifics and raw log

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
func (it *TraderjoeLBPairIgnoredStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeLBPairIgnoredStateChanged)
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
		it.Event = new(TraderjoeLBPairIgnoredStateChanged)
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
func (it *TraderjoeLBPairIgnoredStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeLBPairIgnoredStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeLBPairIgnoredStateChanged represents a LBPairIgnoredStateChanged event raised by the Traderjoe contract.
type TraderjoeLBPairIgnoredStateChanged struct {
	LBPair  common.Address
	Ignored bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterLBPairIgnoredStateChanged is a free log retrieval operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Traderjoe *TraderjoeFilterer) FilterLBPairIgnoredStateChanged(opts *bind.FilterOpts, LBPair []common.Address) (*TraderjoeLBPairIgnoredStateChangedIterator, error) {

	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "LBPairIgnoredStateChanged", LBPairRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeLBPairIgnoredStateChangedIterator{contract: _Traderjoe.contract, event: "LBPairIgnoredStateChanged", logs: logs, sub: sub}, nil
}

// WatchLBPairIgnoredStateChanged is a free log subscription operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Traderjoe *TraderjoeFilterer) WatchLBPairIgnoredStateChanged(opts *bind.WatchOpts, sink chan<- *TraderjoeLBPairIgnoredStateChanged, LBPair []common.Address) (event.Subscription, error) {

	var LBPairRule []interface{}
	for _, LBPairItem := range LBPair {
		LBPairRule = append(LBPairRule, LBPairItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "LBPairIgnoredStateChanged", LBPairRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeLBPairIgnoredStateChanged)
				if err := _Traderjoe.contract.UnpackLog(event, "LBPairIgnoredStateChanged", log); err != nil {
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

// ParseLBPairIgnoredStateChanged is a log parse operation binding the contract event 0x44cf35361c9ff3c8c1397ec6410d5495cc481feaef35c9af11da1a637107de4f.
//
// Solidity: event LBPairIgnoredStateChanged(address indexed LBPair, bool ignored)
func (_Traderjoe *TraderjoeFilterer) ParseLBPairIgnoredStateChanged(log types.Log) (*TraderjoeLBPairIgnoredStateChanged, error) {
	event := new(TraderjoeLBPairIgnoredStateChanged)
	if err := _Traderjoe.contract.UnpackLog(event, "LBPairIgnoredStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeLBPairImplementationSetIterator is returned from FilterLBPairImplementationSet and is used to iterate over the raw logs and unpacked data for LBPairImplementationSet events raised by the Traderjoe contract.
type TraderjoeLBPairImplementationSetIterator struct {
	Event *TraderjoeLBPairImplementationSet // Event containing the contract specifics and raw log

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
func (it *TraderjoeLBPairImplementationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeLBPairImplementationSet)
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
		it.Event = new(TraderjoeLBPairImplementationSet)
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
func (it *TraderjoeLBPairImplementationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeLBPairImplementationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeLBPairImplementationSet represents a LBPairImplementationSet event raised by the Traderjoe contract.
type TraderjoeLBPairImplementationSet struct {
	OldLBPairImplementation common.Address
	LBPairImplementation    common.Address
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterLBPairImplementationSet is a free log retrieval operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Traderjoe *TraderjoeFilterer) FilterLBPairImplementationSet(opts *bind.FilterOpts) (*TraderjoeLBPairImplementationSetIterator, error) {

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "LBPairImplementationSet")
	if err != nil {
		return nil, err
	}
	return &TraderjoeLBPairImplementationSetIterator{contract: _Traderjoe.contract, event: "LBPairImplementationSet", logs: logs, sub: sub}, nil
}

// WatchLBPairImplementationSet is a free log subscription operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Traderjoe *TraderjoeFilterer) WatchLBPairImplementationSet(opts *bind.WatchOpts, sink chan<- *TraderjoeLBPairImplementationSet) (event.Subscription, error) {

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "LBPairImplementationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeLBPairImplementationSet)
				if err := _Traderjoe.contract.UnpackLog(event, "LBPairImplementationSet", log); err != nil {
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

// ParseLBPairImplementationSet is a log parse operation binding the contract event 0x900d0e3d359f50e4f923ecdc06b401e07dbb9f485e17b07bcfc91a13000b277e.
//
// Solidity: event LBPairImplementationSet(address oldLBPairImplementation, address LBPairImplementation)
func (_Traderjoe *TraderjoeFilterer) ParseLBPairImplementationSet(log types.Log) (*TraderjoeLBPairImplementationSet, error) {
	event := new(TraderjoeLBPairImplementationSet)
	if err := _Traderjoe.contract.UnpackLog(event, "LBPairImplementationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Traderjoe contract.
type TraderjoeOwnershipTransferredIterator struct {
	Event *TraderjoeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TraderjoeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeOwnershipTransferred)
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
		it.Event = new(TraderjoeOwnershipTransferred)
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
func (it *TraderjoeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeOwnershipTransferred represents a OwnershipTransferred event raised by the Traderjoe contract.
type TraderjoeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Traderjoe *TraderjoeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TraderjoeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeOwnershipTransferredIterator{contract: _Traderjoe.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Traderjoe *TraderjoeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TraderjoeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeOwnershipTransferred)
				if err := _Traderjoe.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Traderjoe *TraderjoeFilterer) ParseOwnershipTransferred(log types.Log) (*TraderjoeOwnershipTransferred, error) {
	event := new(TraderjoeOwnershipTransferred)
	if err := _Traderjoe.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoePendingOwnerSetIterator is returned from FilterPendingOwnerSet and is used to iterate over the raw logs and unpacked data for PendingOwnerSet events raised by the Traderjoe contract.
type TraderjoePendingOwnerSetIterator struct {
	Event *TraderjoePendingOwnerSet // Event containing the contract specifics and raw log

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
func (it *TraderjoePendingOwnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoePendingOwnerSet)
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
		it.Event = new(TraderjoePendingOwnerSet)
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
func (it *TraderjoePendingOwnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoePendingOwnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoePendingOwnerSet represents a PendingOwnerSet event raised by the Traderjoe contract.
type TraderjoePendingOwnerSet struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPendingOwnerSet is a free log retrieval operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Traderjoe *TraderjoeFilterer) FilterPendingOwnerSet(opts *bind.FilterOpts, pendingOwner []common.Address) (*TraderjoePendingOwnerSetIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "PendingOwnerSet", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoePendingOwnerSetIterator{contract: _Traderjoe.contract, event: "PendingOwnerSet", logs: logs, sub: sub}, nil
}

// WatchPendingOwnerSet is a free log subscription operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Traderjoe *TraderjoeFilterer) WatchPendingOwnerSet(opts *bind.WatchOpts, sink chan<- *TraderjoePendingOwnerSet, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "PendingOwnerSet", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoePendingOwnerSet)
				if err := _Traderjoe.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
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

// ParsePendingOwnerSet is a log parse operation binding the contract event 0x68f49b346b94582a8b5f9d10e3fe3365318fe8f191ff8dce7c59c6cad06b02f5.
//
// Solidity: event PendingOwnerSet(address indexed pendingOwner)
func (_Traderjoe *TraderjoeFilterer) ParsePendingOwnerSet(log types.Log) (*TraderjoePendingOwnerSet, error) {
	event := new(TraderjoePendingOwnerSet)
	if err := _Traderjoe.contract.UnpackLog(event, "PendingOwnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoePresetOpenStateChangedIterator is returned from FilterPresetOpenStateChanged and is used to iterate over the raw logs and unpacked data for PresetOpenStateChanged events raised by the Traderjoe contract.
type TraderjoePresetOpenStateChangedIterator struct {
	Event *TraderjoePresetOpenStateChanged // Event containing the contract specifics and raw log

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
func (it *TraderjoePresetOpenStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoePresetOpenStateChanged)
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
		it.Event = new(TraderjoePresetOpenStateChanged)
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
func (it *TraderjoePresetOpenStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoePresetOpenStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoePresetOpenStateChanged represents a PresetOpenStateChanged event raised by the Traderjoe contract.
type TraderjoePresetOpenStateChanged struct {
	BinStep *big.Int
	IsOpen  bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPresetOpenStateChanged is a free log retrieval operation binding the contract event 0x58a8b6a02b964cca2712e5a71d7b0d564a56b4a0f573b4c47f389341ade14cfd.
//
// Solidity: event PresetOpenStateChanged(uint256 indexed binStep, bool indexed isOpen)
func (_Traderjoe *TraderjoeFilterer) FilterPresetOpenStateChanged(opts *bind.FilterOpts, binStep []*big.Int, isOpen []bool) (*TraderjoePresetOpenStateChangedIterator, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}
	var isOpenRule []interface{}
	for _, isOpenItem := range isOpen {
		isOpenRule = append(isOpenRule, isOpenItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "PresetOpenStateChanged", binStepRule, isOpenRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoePresetOpenStateChangedIterator{contract: _Traderjoe.contract, event: "PresetOpenStateChanged", logs: logs, sub: sub}, nil
}

// WatchPresetOpenStateChanged is a free log subscription operation binding the contract event 0x58a8b6a02b964cca2712e5a71d7b0d564a56b4a0f573b4c47f389341ade14cfd.
//
// Solidity: event PresetOpenStateChanged(uint256 indexed binStep, bool indexed isOpen)
func (_Traderjoe *TraderjoeFilterer) WatchPresetOpenStateChanged(opts *bind.WatchOpts, sink chan<- *TraderjoePresetOpenStateChanged, binStep []*big.Int, isOpen []bool) (event.Subscription, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}
	var isOpenRule []interface{}
	for _, isOpenItem := range isOpen {
		isOpenRule = append(isOpenRule, isOpenItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "PresetOpenStateChanged", binStepRule, isOpenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoePresetOpenStateChanged)
				if err := _Traderjoe.contract.UnpackLog(event, "PresetOpenStateChanged", log); err != nil {
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

// ParsePresetOpenStateChanged is a log parse operation binding the contract event 0x58a8b6a02b964cca2712e5a71d7b0d564a56b4a0f573b4c47f389341ade14cfd.
//
// Solidity: event PresetOpenStateChanged(uint256 indexed binStep, bool indexed isOpen)
func (_Traderjoe *TraderjoeFilterer) ParsePresetOpenStateChanged(log types.Log) (*TraderjoePresetOpenStateChanged, error) {
	event := new(TraderjoePresetOpenStateChanged)
	if err := _Traderjoe.contract.UnpackLog(event, "PresetOpenStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoePresetRemovedIterator is returned from FilterPresetRemoved and is used to iterate over the raw logs and unpacked data for PresetRemoved events raised by the Traderjoe contract.
type TraderjoePresetRemovedIterator struct {
	Event *TraderjoePresetRemoved // Event containing the contract specifics and raw log

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
func (it *TraderjoePresetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoePresetRemoved)
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
		it.Event = new(TraderjoePresetRemoved)
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
func (it *TraderjoePresetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoePresetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoePresetRemoved represents a PresetRemoved event raised by the Traderjoe contract.
type TraderjoePresetRemoved struct {
	BinStep *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPresetRemoved is a free log retrieval operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Traderjoe *TraderjoeFilterer) FilterPresetRemoved(opts *bind.FilterOpts, binStep []*big.Int) (*TraderjoePresetRemovedIterator, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "PresetRemoved", binStepRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoePresetRemovedIterator{contract: _Traderjoe.contract, event: "PresetRemoved", logs: logs, sub: sub}, nil
}

// WatchPresetRemoved is a free log subscription operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Traderjoe *TraderjoeFilterer) WatchPresetRemoved(opts *bind.WatchOpts, sink chan<- *TraderjoePresetRemoved, binStep []*big.Int) (event.Subscription, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "PresetRemoved", binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoePresetRemoved)
				if err := _Traderjoe.contract.UnpackLog(event, "PresetRemoved", log); err != nil {
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

// ParsePresetRemoved is a log parse operation binding the contract event 0xdd86b848bb56ff540caa68683fa467d0e7eb5f8b2d44e4ee435742eeeae9be13.
//
// Solidity: event PresetRemoved(uint256 indexed binStep)
func (_Traderjoe *TraderjoeFilterer) ParsePresetRemoved(log types.Log) (*TraderjoePresetRemoved, error) {
	event := new(TraderjoePresetRemoved)
	if err := _Traderjoe.contract.UnpackLog(event, "PresetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoePresetSetIterator is returned from FilterPresetSet and is used to iterate over the raw logs and unpacked data for PresetSet events raised by the Traderjoe contract.
type TraderjoePresetSetIterator struct {
	Event *TraderjoePresetSet // Event containing the contract specifics and raw log

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
func (it *TraderjoePresetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoePresetSet)
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
		it.Event = new(TraderjoePresetSet)
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
func (it *TraderjoePresetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoePresetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoePresetSet represents a PresetSet event raised by the Traderjoe contract.
type TraderjoePresetSet struct {
	BinStep                  *big.Int
	BaseFactor               *big.Int
	FilterPeriod             *big.Int
	DecayPeriod              *big.Int
	ReductionFactor          *big.Int
	VariableFeeControl       *big.Int
	ProtocolShare            *big.Int
	MaxVolatilityAccumulator *big.Int
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterPresetSet is a free log retrieval operation binding the contract event 0x839844a256a87f87c9c835117d9a1c40be013954064c937072acb32d36db6a28.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator)
func (_Traderjoe *TraderjoeFilterer) FilterPresetSet(opts *bind.FilterOpts, binStep []*big.Int) (*TraderjoePresetSetIterator, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "PresetSet", binStepRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoePresetSetIterator{contract: _Traderjoe.contract, event: "PresetSet", logs: logs, sub: sub}, nil
}

// WatchPresetSet is a free log subscription operation binding the contract event 0x839844a256a87f87c9c835117d9a1c40be013954064c937072acb32d36db6a28.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator)
func (_Traderjoe *TraderjoeFilterer) WatchPresetSet(opts *bind.WatchOpts, sink chan<- *TraderjoePresetSet, binStep []*big.Int) (event.Subscription, error) {

	var binStepRule []interface{}
	for _, binStepItem := range binStep {
		binStepRule = append(binStepRule, binStepItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "PresetSet", binStepRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoePresetSet)
				if err := _Traderjoe.contract.UnpackLog(event, "PresetSet", log); err != nil {
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

// ParsePresetSet is a log parse operation binding the contract event 0x839844a256a87f87c9c835117d9a1c40be013954064c937072acb32d36db6a28.
//
// Solidity: event PresetSet(uint256 indexed binStep, uint256 baseFactor, uint256 filterPeriod, uint256 decayPeriod, uint256 reductionFactor, uint256 variableFeeControl, uint256 protocolShare, uint256 maxVolatilityAccumulator)
func (_Traderjoe *TraderjoeFilterer) ParsePresetSet(log types.Log) (*TraderjoePresetSet, error) {
	event := new(TraderjoePresetSet)
	if err := _Traderjoe.contract.UnpackLog(event, "PresetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeQuoteAssetAddedIterator is returned from FilterQuoteAssetAdded and is used to iterate over the raw logs and unpacked data for QuoteAssetAdded events raised by the Traderjoe contract.
type TraderjoeQuoteAssetAddedIterator struct {
	Event *TraderjoeQuoteAssetAdded // Event containing the contract specifics and raw log

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
func (it *TraderjoeQuoteAssetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeQuoteAssetAdded)
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
		it.Event = new(TraderjoeQuoteAssetAdded)
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
func (it *TraderjoeQuoteAssetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeQuoteAssetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeQuoteAssetAdded represents a QuoteAssetAdded event raised by the Traderjoe contract.
type TraderjoeQuoteAssetAdded struct {
	QuoteAsset common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuoteAssetAdded is a free log retrieval operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) FilterQuoteAssetAdded(opts *bind.FilterOpts, quoteAsset []common.Address) (*TraderjoeQuoteAssetAddedIterator, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "QuoteAssetAdded", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeQuoteAssetAddedIterator{contract: _Traderjoe.contract, event: "QuoteAssetAdded", logs: logs, sub: sub}, nil
}

// WatchQuoteAssetAdded is a free log subscription operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) WatchQuoteAssetAdded(opts *bind.WatchOpts, sink chan<- *TraderjoeQuoteAssetAdded, quoteAsset []common.Address) (event.Subscription, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "QuoteAssetAdded", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeQuoteAssetAdded)
				if err := _Traderjoe.contract.UnpackLog(event, "QuoteAssetAdded", log); err != nil {
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

// ParseQuoteAssetAdded is a log parse operation binding the contract event 0x84cc2115995684dcb0cd3d3a9565e3d32f075de81db70c8dc3a719b2a47af67e.
//
// Solidity: event QuoteAssetAdded(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) ParseQuoteAssetAdded(log types.Log) (*TraderjoeQuoteAssetAdded, error) {
	event := new(TraderjoeQuoteAssetAdded)
	if err := _Traderjoe.contract.UnpackLog(event, "QuoteAssetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TraderjoeQuoteAssetRemovedIterator is returned from FilterQuoteAssetRemoved and is used to iterate over the raw logs and unpacked data for QuoteAssetRemoved events raised by the Traderjoe contract.
type TraderjoeQuoteAssetRemovedIterator struct {
	Event *TraderjoeQuoteAssetRemoved // Event containing the contract specifics and raw log

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
func (it *TraderjoeQuoteAssetRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TraderjoeQuoteAssetRemoved)
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
		it.Event = new(TraderjoeQuoteAssetRemoved)
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
func (it *TraderjoeQuoteAssetRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TraderjoeQuoteAssetRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TraderjoeQuoteAssetRemoved represents a QuoteAssetRemoved event raised by the Traderjoe contract.
type TraderjoeQuoteAssetRemoved struct {
	QuoteAsset common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterQuoteAssetRemoved is a free log retrieval operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) FilterQuoteAssetRemoved(opts *bind.FilterOpts, quoteAsset []common.Address) (*TraderjoeQuoteAssetRemovedIterator, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Traderjoe.contract.FilterLogs(opts, "QuoteAssetRemoved", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return &TraderjoeQuoteAssetRemovedIterator{contract: _Traderjoe.contract, event: "QuoteAssetRemoved", logs: logs, sub: sub}, nil
}

// WatchQuoteAssetRemoved is a free log subscription operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) WatchQuoteAssetRemoved(opts *bind.WatchOpts, sink chan<- *TraderjoeQuoteAssetRemoved, quoteAsset []common.Address) (event.Subscription, error) {

	var quoteAssetRule []interface{}
	for _, quoteAssetItem := range quoteAsset {
		quoteAssetRule = append(quoteAssetRule, quoteAssetItem)
	}

	logs, sub, err := _Traderjoe.contract.WatchLogs(opts, "QuoteAssetRemoved", quoteAssetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TraderjoeQuoteAssetRemoved)
				if err := _Traderjoe.contract.UnpackLog(event, "QuoteAssetRemoved", log); err != nil {
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

// ParseQuoteAssetRemoved is a log parse operation binding the contract event 0x0b767739217755d8af5a2ba75b181a19fa1750f8bb701f09311cb19a90140cb3.
//
// Solidity: event QuoteAssetRemoved(address indexed quoteAsset)
func (_Traderjoe *TraderjoeFilterer) ParseQuoteAssetRemoved(log types.Log) (*TraderjoeQuoteAssetRemoved, error) {
	event := new(TraderjoeQuoteAssetRemoved)
	if err := _Traderjoe.contract.UnpackLog(event, "QuoteAssetRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
