// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vaultextension

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

// HooksConfig is an auto generated low-level Go binding around an user-defined struct.
type HooksConfig struct {
	EnableHookAdjustedAmounts       bool
	ShouldCallBeforeInitialize      bool
	ShouldCallAfterInitialize       bool
	ShouldCallComputeDynamicSwapFee bool
	ShouldCallBeforeSwap            bool
	ShouldCallAfterSwap             bool
	ShouldCallBeforeAddLiquidity    bool
	ShouldCallAfterAddLiquidity     bool
	ShouldCallBeforeRemoveLiquidity bool
	ShouldCallAfterRemoveLiquidity  bool
	HooksContract                   common.Address
}

// LiquidityManagement is an auto generated low-level Go binding around an user-defined struct.
type LiquidityManagement struct {
	DisableUnbalancedLiquidity  bool
	EnableAddLiquidityCustom    bool
	EnableRemoveLiquidityCustom bool
	EnableDonation              bool
}

// PoolConfig is an auto generated low-level Go binding around an user-defined struct.
type PoolConfig struct {
	LiquidityManagement         LiquidityManagement
	StaticSwapFeePercentage     *big.Int
	AggregateSwapFeePercentage  *big.Int
	AggregateYieldFeePercentage *big.Int
	TokenDecimalDiffs           *big.Int
	PauseWindowEndTime          uint32
	IsPoolRegistered            bool
	IsPoolInitialized           bool
	IsPoolPaused                bool
	IsPoolInRecoveryMode        bool
}

// PoolData is an auto generated low-level Go binding around an user-defined struct.
type PoolData struct {
	PoolConfigBits        [32]byte
	Tokens                []common.Address
	TokenInfo             []TokenInfo
	BalancesRaw           []*big.Int
	BalancesLiveScaled18  []*big.Int
	TokenRates            []*big.Int
	DecimalScalingFactors []*big.Int
}

// PoolRoleAccounts is an auto generated low-level Go binding around an user-defined struct.
type PoolRoleAccounts struct {
	PauseManager   common.Address
	SwapFeeManager common.Address
	PoolCreator    common.Address
}

// PoolSwapParams is an auto generated low-level Go binding around an user-defined struct.
type PoolSwapParams struct {
	Kind                uint8
	AmountGivenScaled18 *big.Int
	BalancesScaled18    []*big.Int
	IndexIn             *big.Int
	IndexOut            *big.Int
	Router              common.Address
	UserData            []byte
}

// TokenConfig is an auto generated low-level Go binding around an user-defined struct.
type TokenConfig struct {
	Token         common.Address
	TokenType     uint8
	RateProvider  common.Address
	PaysYieldFees bool
}

// TokenInfo is an auto generated low-level Go binding around an user-defined struct.
type TokenInfo struct {
	TokenType     uint8
	RateProvider  common.Address
	PaysYieldFees bool
}

// VaultextensionMetaData contains all meta data concerning the Vaultextension contract.
var VaultextensionMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVault\",\"name\":\"mainVault\",\"type\":\"address\"},{\"internalType\":\"contractIVaultAdmin\",\"name\":\"vaultAdmin\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterAddLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterInitializeHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterRemoveLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterSwapHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountGivenZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"AmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"AmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeAddLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeInitializeHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeRemoveLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeSwapHookFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"BptAmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"BptAmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"BufferAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"BufferNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BufferSharesInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BufferSharesInvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"BufferTotalSupplyTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotReceiveEth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSwapSameToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CodecOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportAddLiquidityCustom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportDonation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportRemoveLiquidityCustom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportUnbalancedLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DynamicSwapFeeHookFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ErrorSelectorNotFound\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeePrecisionTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedAmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedAmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedSwapLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolHooksContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolFactory\",\"type\":\"address\"}],\"name\":\"HookRegistrationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddLiquidityKind\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRemoveLiquidityKind\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenConfiguration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"InvalidUnderlyingToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuedShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIssuedShares\",\"type\":\"uint256\"}],\"name\":\"IssuedSharesBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughBufferShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedUnderlyingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUnderlyingAmount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughUnderlying\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedWrappedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualWrappedAmount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughWrapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStaticCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVaultDelegateCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OutOfBounds\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseBufferPeriodDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PercentageAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolInRecoveryMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotInRecoveryMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolPauseWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"PoolTotalSupplyTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeesExceedTotalCollected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueriesDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueriesDisabledPermanently\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuoteResultSpoofed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"Result\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotTrusted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderIsNotVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFeePercentageTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFeePercentageTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"SwapLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actualToken\",\"type\":\"address\"}],\"name\":\"TokensMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokensNotSorted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultBuffersArePaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultIsNotUnlocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPauseWindowDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPauseWindowExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"WrapAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongProtocolFeeControllerDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"WrongUnderlyingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongVaultAdminDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongVaultExtensionDeployment\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aggregateSwapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"AggregateSwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aggregateYieldFeePercentage\",\"type\":\"uint256\"}],\"name\":\"AggregateYieldFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedShares\",\"type\":\"uint256\"}],\"name\":\"BufferSharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issuedShares\",\"type\":\"uint256\"}],\"name\":\"BufferSharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAddLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsAddedRaw\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"swapFeeAmountsRaw\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"LiquidityAddedToBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRemoveLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsRemovedRaw\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"swapFeeAmountsRaw\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"LiquidityRemovedFromBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PoolPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"recoveryMode\",\"type\":\"bool\"}],\"name\":\"PoolRecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paysYieldFees\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenConfig[]\",\"name\":\"tokenConfig\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"pauseManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolCreator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structPoolRoleAccounts\",\"name\":\"roleAccounts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"enableHookAdjustedAmounts\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallComputeDynamicSwapFee\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"hooksContract\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structHooksConfig\",\"name\":\"hooksConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"disableUnbalancedLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAddLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableRemoveLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableDonation\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structLiquidityManagement\",\"name\":\"liquidityManagement\",\"type\":\"tuple\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIProtocolFeeController\",\"name\":\"newProtocolFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeeAmount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"eventKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"eventData\",\"type\":\"bytes\"}],\"name\":\"VaultAuxiliary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"VaultBuffersPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"VaultPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VaultQueriesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VaultQueriesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"Wrap\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"amountGivenScaled18\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"balancesScaled18\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"indexIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"indexOut\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structPoolSwapParams\",\"name\":\"swapParams\",\"type\":\"tuple\"}],\"name\":\"computeDynamicSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"dynamicSwapFeePercentage\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"eventKey\",\"type\":\"bytes32\"},{\"internalType\":\"bytes\",\"name\":\"eventData\",\"type\":\"bytes\"}],\"name\":\"emitAuxiliaryEvent\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getAddLiquidityCalledFlag\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAggregateSwapFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAggregateYieldFeeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAuthorizer\",\"outputs\":[{\"internalType\":\"contractIAuthorizer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getBptRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getCurrentLiveBalances\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"balancesLiveScaled18\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"getERC4626BufferAsset\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getHooksConfig\",\"outputs\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"enableHookAdjustedAmounts\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallComputeDynamicSwapFee\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"hooksContract\",\"type\":\"address\"}],\"internalType\":\"structHooksConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNonzeroDeltaCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolConfig\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"bool\",\"name\":\"disableUnbalancedLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAddLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableRemoveLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableDonation\",\"type\":\"bool\"}],\"internalType\":\"structLiquidityManagement\",\"name\":\"liquidityManagement\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"staticSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregateSwapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"aggregateYieldFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"tokenDecimalDiffs\",\"type\":\"uint40\"},{\"internalType\":\"uint32\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"isPoolRegistered\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPoolInitialized\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPoolPaused\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isPoolInRecoveryMode\",\"type\":\"bool\"}],\"internalType\":\"structPoolConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolData\",\"outputs\":[{\"components\":[{\"internalType\":\"PoolConfigBits\",\"name\":\"poolConfigBits\",\"type\":\"bytes32\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paysYieldFees\",\"type\":\"bool\"}],\"internalType\":\"structTokenInfo[]\",\"name\":\"tokenInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balancesRaw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balancesLiveScaled18\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenRates\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"decimalScalingFactors\",\"type\":\"uint256[]\"}],\"internalType\":\"structPoolData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolPausedState\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolRoleAccounts\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pauseManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolCreator\",\"type\":\"address\"}],\"internalType\":\"structPoolRoleAccounts\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolTokenInfo\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"components\":[{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paysYieldFees\",\"type\":\"bool\"}],\"internalType\":\"structTokenInfo[]\",\"name\":\"tokenInfo\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256[]\",\"name\":\"balancesRaw\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"lastBalancesLiveScaled18\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolTokenRates\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"decimalScalingFactors\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"tokenRates\",\"type\":\"uint256[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getPoolTokens\",\"outputs\":[{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getProtocolFeeController\",\"outputs\":[{\"internalType\":\"contractIProtocolFeeController\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getReservesOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"getStaticSwapFeePercentage\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getTokenDelta\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"contractIERC20[]\",\"name\":\"tokens\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"exactAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minBptAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptAmountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"isERC4626BufferInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolInRecoveryMode\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"isPoolRegistered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isQueryDisabled\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isQueryDisabledPermanently\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isUnlocked\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"quote\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"quoteAndRevert\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reentrancyGuardEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paysYieldFees\",\"type\":\"bool\"}],\"internalType\":\"structTokenConfig[]\",\"name\":\"tokenConfig\",\"type\":\"tuple[]\"},{\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"internalType\":\"uint32\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint32\"},{\"internalType\":\"bool\",\"name\":\"protocolFeeExempt\",\"type\":\"bool\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"pauseManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolCreator\",\"type\":\"address\"}],\"internalType\":\"structPoolRoleAccounts\",\"name\":\"roleAccounts\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"poolHooksContract\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"disableUnbalancedLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAddLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableRemoveLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableDonation\",\"type\":\"bool\"}],\"internalType\":\"structLiquidityManagement\",\"name\":\"liquidityManagement\",\"type\":\"tuple\"}],\"name\":\"registerPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"exactBptAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"}],\"name\":\"removeLiquidityRecovery\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsOutRaw\",\"type\":\"uint256[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"vault\",\"outputs\":[{\"internalType\":\"contractIVault\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VaultextensionABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultextensionMetaData.ABI instead.
var VaultextensionABI = VaultextensionMetaData.ABI

// Vaultextension is an auto generated Go binding around an Ethereum contract.
type Vaultextension struct {
	VaultextensionCaller     // Read-only binding to the contract
	VaultextensionTransactor // Write-only binding to the contract
	VaultextensionFilterer   // Log filterer for contract events
}

// VaultextensionCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultextensionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultextensionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultextensionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultextensionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultextensionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultextensionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultextensionSession struct {
	Contract     *Vaultextension   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultextensionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultextensionCallerSession struct {
	Contract *VaultextensionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// VaultextensionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultextensionTransactorSession struct {
	Contract     *VaultextensionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// VaultextensionRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultextensionRaw struct {
	Contract *Vaultextension // Generic contract binding to access the raw methods on
}

// VaultextensionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultextensionCallerRaw struct {
	Contract *VaultextensionCaller // Generic read-only contract binding to access the raw methods on
}

// VaultextensionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultextensionTransactorRaw struct {
	Contract *VaultextensionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVaultextension creates a new instance of Vaultextension, bound to a specific deployed contract.
func NewVaultextension(address common.Address, backend bind.ContractBackend) (*Vaultextension, error) {
	contract, err := bindVaultextension(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vaultextension{VaultextensionCaller: VaultextensionCaller{contract: contract}, VaultextensionTransactor: VaultextensionTransactor{contract: contract}, VaultextensionFilterer: VaultextensionFilterer{contract: contract}}, nil
}

// NewVaultextensionCaller creates a new read-only instance of Vaultextension, bound to a specific deployed contract.
func NewVaultextensionCaller(address common.Address, caller bind.ContractCaller) (*VaultextensionCaller, error) {
	contract, err := bindVaultextension(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultextensionCaller{contract: contract}, nil
}

// NewVaultextensionTransactor creates a new write-only instance of Vaultextension, bound to a specific deployed contract.
func NewVaultextensionTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultextensionTransactor, error) {
	contract, err := bindVaultextension(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultextensionTransactor{contract: contract}, nil
}

// NewVaultextensionFilterer creates a new log filterer instance of Vaultextension, bound to a specific deployed contract.
func NewVaultextensionFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultextensionFilterer, error) {
	contract, err := bindVaultextension(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultextensionFilterer{contract: contract}, nil
}

// bindVaultextension binds a generic wrapper to an already deployed contract.
func bindVaultextension(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultextensionMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultextension *VaultextensionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vaultextension.Contract.VaultextensionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultextension *VaultextensionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultextension.Contract.VaultextensionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultextension *VaultextensionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultextension.Contract.VaultextensionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vaultextension *VaultextensionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vaultextension.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vaultextension *VaultextensionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultextension.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vaultextension *VaultextensionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vaultextension.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address token, address owner, address spender) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) Allowance(opts *bind.CallOpts, token common.Address, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "allowance", token, owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address token, address owner, address spender) view returns(uint256)
func (_Vaultextension *VaultextensionSession) Allowance(token common.Address, owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.Allowance(&_Vaultextension.CallOpts, token, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0x927da105.
//
// Solidity: function allowance(address token, address owner, address spender) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) Allowance(token common.Address, owner common.Address, spender common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.Allowance(&_Vaultextension.CallOpts, token, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) BalanceOf(opts *bind.CallOpts, token common.Address, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "balanceOf", token, account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256)
func (_Vaultextension *VaultextensionSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.BalanceOf(&_Vaultextension.CallOpts, token, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0xf7888aec.
//
// Solidity: function balanceOf(address token, address account) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) BalanceOf(token common.Address, account common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.BalanceOf(&_Vaultextension.CallOpts, token, account)
}

// ComputeDynamicSwapFeePercentage is a free data retrieval call binding the contract method 0x4d472bdd.
//
// Solidity: function computeDynamicSwapFeePercentage(address pool, (uint8,uint256,uint256[],uint256,uint256,address,bytes) swapParams) view returns(uint256 dynamicSwapFeePercentage)
func (_Vaultextension *VaultextensionCaller) ComputeDynamicSwapFeePercentage(opts *bind.CallOpts, pool common.Address, swapParams PoolSwapParams) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "computeDynamicSwapFeePercentage", pool, swapParams)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ComputeDynamicSwapFeePercentage is a free data retrieval call binding the contract method 0x4d472bdd.
//
// Solidity: function computeDynamicSwapFeePercentage(address pool, (uint8,uint256,uint256[],uint256,uint256,address,bytes) swapParams) view returns(uint256 dynamicSwapFeePercentage)
func (_Vaultextension *VaultextensionSession) ComputeDynamicSwapFeePercentage(pool common.Address, swapParams PoolSwapParams) (*big.Int, error) {
	return _Vaultextension.Contract.ComputeDynamicSwapFeePercentage(&_Vaultextension.CallOpts, pool, swapParams)
}

// ComputeDynamicSwapFeePercentage is a free data retrieval call binding the contract method 0x4d472bdd.
//
// Solidity: function computeDynamicSwapFeePercentage(address pool, (uint8,uint256,uint256[],uint256,uint256,address,bytes) swapParams) view returns(uint256 dynamicSwapFeePercentage)
func (_Vaultextension *VaultextensionCallerSession) ComputeDynamicSwapFeePercentage(pool common.Address, swapParams PoolSwapParams) (*big.Int, error) {
	return _Vaultextension.Contract.ComputeDynamicSwapFeePercentage(&_Vaultextension.CallOpts, pool, swapParams)
}

// GetAddLiquidityCalledFlag is a free data retrieval call binding the contract method 0xace9b89b.
//
// Solidity: function getAddLiquidityCalledFlag(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCaller) GetAddLiquidityCalledFlag(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getAddLiquidityCalledFlag", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GetAddLiquidityCalledFlag is a free data retrieval call binding the contract method 0xace9b89b.
//
// Solidity: function getAddLiquidityCalledFlag(address pool) view returns(bool)
func (_Vaultextension *VaultextensionSession) GetAddLiquidityCalledFlag(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.GetAddLiquidityCalledFlag(&_Vaultextension.CallOpts, pool)
}

// GetAddLiquidityCalledFlag is a free data retrieval call binding the contract method 0xace9b89b.
//
// Solidity: function getAddLiquidityCalledFlag(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) GetAddLiquidityCalledFlag(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.GetAddLiquidityCalledFlag(&_Vaultextension.CallOpts, pool)
}

// GetAggregateSwapFeeAmount is a free data retrieval call binding the contract method 0x85e0b999.
//
// Solidity: function getAggregateSwapFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) GetAggregateSwapFeeAmount(opts *bind.CallOpts, pool common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getAggregateSwapFeeAmount", pool, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggregateSwapFeeAmount is a free data retrieval call binding the contract method 0x85e0b999.
//
// Solidity: function getAggregateSwapFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionSession) GetAggregateSwapFeeAmount(pool common.Address, token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetAggregateSwapFeeAmount(&_Vaultextension.CallOpts, pool, token)
}

// GetAggregateSwapFeeAmount is a free data retrieval call binding the contract method 0x85e0b999.
//
// Solidity: function getAggregateSwapFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) GetAggregateSwapFeeAmount(pool common.Address, token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetAggregateSwapFeeAmount(&_Vaultextension.CallOpts, pool, token)
}

// GetAggregateYieldFeeAmount is a free data retrieval call binding the contract method 0x00fdfa13.
//
// Solidity: function getAggregateYieldFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) GetAggregateYieldFeeAmount(opts *bind.CallOpts, pool common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getAggregateYieldFeeAmount", pool, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAggregateYieldFeeAmount is a free data retrieval call binding the contract method 0x00fdfa13.
//
// Solidity: function getAggregateYieldFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionSession) GetAggregateYieldFeeAmount(pool common.Address, token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetAggregateYieldFeeAmount(&_Vaultextension.CallOpts, pool, token)
}

// GetAggregateYieldFeeAmount is a free data retrieval call binding the contract method 0x00fdfa13.
//
// Solidity: function getAggregateYieldFeeAmount(address pool, address token) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) GetAggregateYieldFeeAmount(pool common.Address, token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetAggregateYieldFeeAmount(&_Vaultextension.CallOpts, pool, token)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vaultextension *VaultextensionCaller) GetAuthorizer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getAuthorizer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vaultextension *VaultextensionSession) GetAuthorizer() (common.Address, error) {
	return _Vaultextension.Contract.GetAuthorizer(&_Vaultextension.CallOpts)
}

// GetAuthorizer is a free data retrieval call binding the contract method 0xaaabadc5.
//
// Solidity: function getAuthorizer() view returns(address)
func (_Vaultextension *VaultextensionCallerSession) GetAuthorizer() (common.Address, error) {
	return _Vaultextension.Contract.GetAuthorizer(&_Vaultextension.CallOpts)
}

// GetBptRate is a free data retrieval call binding the contract method 0x4f037ee7.
//
// Solidity: function getBptRate(address pool) view returns(uint256 rate)
func (_Vaultextension *VaultextensionCaller) GetBptRate(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getBptRate", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBptRate is a free data retrieval call binding the contract method 0x4f037ee7.
//
// Solidity: function getBptRate(address pool) view returns(uint256 rate)
func (_Vaultextension *VaultextensionSession) GetBptRate(pool common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetBptRate(&_Vaultextension.CallOpts, pool)
}

// GetBptRate is a free data retrieval call binding the contract method 0x4f037ee7.
//
// Solidity: function getBptRate(address pool) view returns(uint256 rate)
func (_Vaultextension *VaultextensionCallerSession) GetBptRate(pool common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetBptRate(&_Vaultextension.CallOpts, pool)
}

// GetCurrentLiveBalances is a free data retrieval call binding the contract method 0x535cfd8a.
//
// Solidity: function getCurrentLiveBalances(address pool) view returns(uint256[] balancesLiveScaled18)
func (_Vaultextension *VaultextensionCaller) GetCurrentLiveBalances(opts *bind.CallOpts, pool common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getCurrentLiveBalances", pool)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetCurrentLiveBalances is a free data retrieval call binding the contract method 0x535cfd8a.
//
// Solidity: function getCurrentLiveBalances(address pool) view returns(uint256[] balancesLiveScaled18)
func (_Vaultextension *VaultextensionSession) GetCurrentLiveBalances(pool common.Address) ([]*big.Int, error) {
	return _Vaultextension.Contract.GetCurrentLiveBalances(&_Vaultextension.CallOpts, pool)
}

// GetCurrentLiveBalances is a free data retrieval call binding the contract method 0x535cfd8a.
//
// Solidity: function getCurrentLiveBalances(address pool) view returns(uint256[] balancesLiveScaled18)
func (_Vaultextension *VaultextensionCallerSession) GetCurrentLiveBalances(pool common.Address) ([]*big.Int, error) {
	return _Vaultextension.Contract.GetCurrentLiveBalances(&_Vaultextension.CallOpts, pool)
}

// GetERC4626BufferAsset is a free data retrieval call binding the contract method 0x4afbaf5a.
//
// Solidity: function getERC4626BufferAsset(address wrappedToken) view returns(address asset)
func (_Vaultextension *VaultextensionCaller) GetERC4626BufferAsset(opts *bind.CallOpts, wrappedToken common.Address) (common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getERC4626BufferAsset", wrappedToken)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetERC4626BufferAsset is a free data retrieval call binding the contract method 0x4afbaf5a.
//
// Solidity: function getERC4626BufferAsset(address wrappedToken) view returns(address asset)
func (_Vaultextension *VaultextensionSession) GetERC4626BufferAsset(wrappedToken common.Address) (common.Address, error) {
	return _Vaultextension.Contract.GetERC4626BufferAsset(&_Vaultextension.CallOpts, wrappedToken)
}

// GetERC4626BufferAsset is a free data retrieval call binding the contract method 0x4afbaf5a.
//
// Solidity: function getERC4626BufferAsset(address wrappedToken) view returns(address asset)
func (_Vaultextension *VaultextensionCallerSession) GetERC4626BufferAsset(wrappedToken common.Address) (common.Address, error) {
	return _Vaultextension.Contract.GetERC4626BufferAsset(&_Vaultextension.CallOpts, wrappedToken)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xce8630d4.
//
// Solidity: function getHooksConfig(address pool) view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address))
func (_Vaultextension *VaultextensionCaller) GetHooksConfig(opts *bind.CallOpts, pool common.Address) (HooksConfig, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getHooksConfig", pool)

	if err != nil {
		return *new(HooksConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(HooksConfig)).(*HooksConfig)

	return out0, err

}

// GetHooksConfig is a free data retrieval call binding the contract method 0xce8630d4.
//
// Solidity: function getHooksConfig(address pool) view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address))
func (_Vaultextension *VaultextensionSession) GetHooksConfig(pool common.Address) (HooksConfig, error) {
	return _Vaultextension.Contract.GetHooksConfig(&_Vaultextension.CallOpts, pool)
}

// GetHooksConfig is a free data retrieval call binding the contract method 0xce8630d4.
//
// Solidity: function getHooksConfig(address pool) view returns((bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address))
func (_Vaultextension *VaultextensionCallerSession) GetHooksConfig(pool common.Address) (HooksConfig, error) {
	return _Vaultextension.Contract.GetHooksConfig(&_Vaultextension.CallOpts, pool)
}

// GetNonzeroDeltaCount is a free data retrieval call binding the contract method 0xdb817187.
//
// Solidity: function getNonzeroDeltaCount() view returns(uint256)
func (_Vaultextension *VaultextensionCaller) GetNonzeroDeltaCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getNonzeroDeltaCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNonzeroDeltaCount is a free data retrieval call binding the contract method 0xdb817187.
//
// Solidity: function getNonzeroDeltaCount() view returns(uint256)
func (_Vaultextension *VaultextensionSession) GetNonzeroDeltaCount() (*big.Int, error) {
	return _Vaultextension.Contract.GetNonzeroDeltaCount(&_Vaultextension.CallOpts)
}

// GetNonzeroDeltaCount is a free data retrieval call binding the contract method 0xdb817187.
//
// Solidity: function getNonzeroDeltaCount() view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) GetNonzeroDeltaCount() (*big.Int, error) {
	return _Vaultextension.Contract.GetNonzeroDeltaCount(&_Vaultextension.CallOpts)
}

// GetPoolConfig is a free data retrieval call binding the contract method 0xf29486a1.
//
// Solidity: function getPoolConfig(address pool) view returns(((bool,bool,bool,bool),uint256,uint256,uint256,uint40,uint32,bool,bool,bool,bool))
func (_Vaultextension *VaultextensionCaller) GetPoolConfig(opts *bind.CallOpts, pool common.Address) (PoolConfig, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolConfig", pool)

	if err != nil {
		return *new(PoolConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolConfig)).(*PoolConfig)

	return out0, err

}

// GetPoolConfig is a free data retrieval call binding the contract method 0xf29486a1.
//
// Solidity: function getPoolConfig(address pool) view returns(((bool,bool,bool,bool),uint256,uint256,uint256,uint40,uint32,bool,bool,bool,bool))
func (_Vaultextension *VaultextensionSession) GetPoolConfig(pool common.Address) (PoolConfig, error) {
	return _Vaultextension.Contract.GetPoolConfig(&_Vaultextension.CallOpts, pool)
}

// GetPoolConfig is a free data retrieval call binding the contract method 0xf29486a1.
//
// Solidity: function getPoolConfig(address pool) view returns(((bool,bool,bool,bool),uint256,uint256,uint256,uint40,uint32,bool,bool,bool,bool))
func (_Vaultextension *VaultextensionCallerSession) GetPoolConfig(pool common.Address) (PoolConfig, error) {
	return _Vaultextension.Contract.GetPoolConfig(&_Vaultextension.CallOpts, pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address pool) view returns((bytes32,address[],(uint8,address,bool)[],uint256[],uint256[],uint256[],uint256[]))
func (_Vaultextension *VaultextensionCaller) GetPoolData(opts *bind.CallOpts, pool common.Address) (PoolData, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolData", pool)

	if err != nil {
		return *new(PoolData), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolData)).(*PoolData)

	return out0, err

}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address pool) view returns((bytes32,address[],(uint8,address,bool)[],uint256[],uint256[],uint256[],uint256[]))
func (_Vaultextension *VaultextensionSession) GetPoolData(pool common.Address) (PoolData, error) {
	return _Vaultextension.Contract.GetPoolData(&_Vaultextension.CallOpts, pool)
}

// GetPoolData is a free data retrieval call binding the contract method 0x13d21cdf.
//
// Solidity: function getPoolData(address pool) view returns((bytes32,address[],(uint8,address,bool)[],uint256[],uint256[],uint256[],uint256[]))
func (_Vaultextension *VaultextensionCallerSession) GetPoolData(pool common.Address) (PoolData, error) {
	return _Vaultextension.Contract.GetPoolData(&_Vaultextension.CallOpts, pool)
}

// GetPoolPausedState is a free data retrieval call binding the contract method 0x15e32046.
//
// Solidity: function getPoolPausedState(address pool) view returns(bool, uint32, uint32, address)
func (_Vaultextension *VaultextensionCaller) GetPoolPausedState(opts *bind.CallOpts, pool common.Address) (bool, uint32, uint32, common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolPausedState", pool)

	if err != nil {
		return *new(bool), *new(uint32), *new(uint32), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)
	out1 := *abi.ConvertType(out[1], new(uint32)).(*uint32)
	out2 := *abi.ConvertType(out[2], new(uint32)).(*uint32)
	out3 := *abi.ConvertType(out[3], new(common.Address)).(*common.Address)

	return out0, out1, out2, out3, err

}

// GetPoolPausedState is a free data retrieval call binding the contract method 0x15e32046.
//
// Solidity: function getPoolPausedState(address pool) view returns(bool, uint32, uint32, address)
func (_Vaultextension *VaultextensionSession) GetPoolPausedState(pool common.Address) (bool, uint32, uint32, common.Address, error) {
	return _Vaultextension.Contract.GetPoolPausedState(&_Vaultextension.CallOpts, pool)
}

// GetPoolPausedState is a free data retrieval call binding the contract method 0x15e32046.
//
// Solidity: function getPoolPausedState(address pool) view returns(bool, uint32, uint32, address)
func (_Vaultextension *VaultextensionCallerSession) GetPoolPausedState(pool common.Address) (bool, uint32, uint32, common.Address, error) {
	return _Vaultextension.Contract.GetPoolPausedState(&_Vaultextension.CallOpts, pool)
}

// GetPoolRoleAccounts is a free data retrieval call binding the contract method 0xe9ddeb26.
//
// Solidity: function getPoolRoleAccounts(address pool) view returns((address,address,address))
func (_Vaultextension *VaultextensionCaller) GetPoolRoleAccounts(opts *bind.CallOpts, pool common.Address) (PoolRoleAccounts, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolRoleAccounts", pool)

	if err != nil {
		return *new(PoolRoleAccounts), err
	}

	out0 := *abi.ConvertType(out[0], new(PoolRoleAccounts)).(*PoolRoleAccounts)

	return out0, err

}

// GetPoolRoleAccounts is a free data retrieval call binding the contract method 0xe9ddeb26.
//
// Solidity: function getPoolRoleAccounts(address pool) view returns((address,address,address))
func (_Vaultextension *VaultextensionSession) GetPoolRoleAccounts(pool common.Address) (PoolRoleAccounts, error) {
	return _Vaultextension.Contract.GetPoolRoleAccounts(&_Vaultextension.CallOpts, pool)
}

// GetPoolRoleAccounts is a free data retrieval call binding the contract method 0xe9ddeb26.
//
// Solidity: function getPoolRoleAccounts(address pool) view returns((address,address,address))
func (_Vaultextension *VaultextensionCallerSession) GetPoolRoleAccounts(pool common.Address) (PoolRoleAccounts, error) {
	return _Vaultextension.Contract.GetPoolRoleAccounts(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0x67e0e076.
//
// Solidity: function getPoolTokenInfo(address pool) view returns(address[] tokens, (uint8,address,bool)[] tokenInfo, uint256[] balancesRaw, uint256[] lastBalancesLiveScaled18)
func (_Vaultextension *VaultextensionCaller) GetPoolTokenInfo(opts *bind.CallOpts, pool common.Address) (struct {
	Tokens                   []common.Address
	TokenInfo                []TokenInfo
	BalancesRaw              []*big.Int
	LastBalancesLiveScaled18 []*big.Int
}, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolTokenInfo", pool)

	outstruct := new(struct {
		Tokens                   []common.Address
		TokenInfo                []TokenInfo
		BalancesRaw              []*big.Int
		LastBalancesLiveScaled18 []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.TokenInfo = *abi.ConvertType(out[1], new([]TokenInfo)).(*[]TokenInfo)
	outstruct.BalancesRaw = *abi.ConvertType(out[2], new([]*big.Int)).(*[]*big.Int)
	outstruct.LastBalancesLiveScaled18 = *abi.ConvertType(out[3], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0x67e0e076.
//
// Solidity: function getPoolTokenInfo(address pool) view returns(address[] tokens, (uint8,address,bool)[] tokenInfo, uint256[] balancesRaw, uint256[] lastBalancesLiveScaled18)
func (_Vaultextension *VaultextensionSession) GetPoolTokenInfo(pool common.Address) (struct {
	Tokens                   []common.Address
	TokenInfo                []TokenInfo
	BalancesRaw              []*big.Int
	LastBalancesLiveScaled18 []*big.Int
}, error) {
	return _Vaultextension.Contract.GetPoolTokenInfo(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokenInfo is a free data retrieval call binding the contract method 0x67e0e076.
//
// Solidity: function getPoolTokenInfo(address pool) view returns(address[] tokens, (uint8,address,bool)[] tokenInfo, uint256[] balancesRaw, uint256[] lastBalancesLiveScaled18)
func (_Vaultextension *VaultextensionCallerSession) GetPoolTokenInfo(pool common.Address) (struct {
	Tokens                   []common.Address
	TokenInfo                []TokenInfo
	BalancesRaw              []*big.Int
	LastBalancesLiveScaled18 []*big.Int
}, error) {
	return _Vaultextension.Contract.GetPoolTokenInfo(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokenRates is a free data retrieval call binding the contract method 0x7e361bde.
//
// Solidity: function getPoolTokenRates(address pool) view returns(uint256[] decimalScalingFactors, uint256[] tokenRates)
func (_Vaultextension *VaultextensionCaller) GetPoolTokenRates(opts *bind.CallOpts, pool common.Address) (struct {
	DecimalScalingFactors []*big.Int
	TokenRates            []*big.Int
}, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolTokenRates", pool)

	outstruct := new(struct {
		DecimalScalingFactors []*big.Int
		TokenRates            []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.DecimalScalingFactors = *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)
	outstruct.TokenRates = *abi.ConvertType(out[1], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// GetPoolTokenRates is a free data retrieval call binding the contract method 0x7e361bde.
//
// Solidity: function getPoolTokenRates(address pool) view returns(uint256[] decimalScalingFactors, uint256[] tokenRates)
func (_Vaultextension *VaultextensionSession) GetPoolTokenRates(pool common.Address) (struct {
	DecimalScalingFactors []*big.Int
	TokenRates            []*big.Int
}, error) {
	return _Vaultextension.Contract.GetPoolTokenRates(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokenRates is a free data retrieval call binding the contract method 0x7e361bde.
//
// Solidity: function getPoolTokenRates(address pool) view returns(uint256[] decimalScalingFactors, uint256[] tokenRates)
func (_Vaultextension *VaultextensionCallerSession) GetPoolTokenRates(pool common.Address) (struct {
	DecimalScalingFactors []*big.Int
	TokenRates            []*big.Int
}, error) {
	return _Vaultextension.Contract.GetPoolTokenRates(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xca4f2803.
//
// Solidity: function getPoolTokens(address pool) view returns(address[] tokens)
func (_Vaultextension *VaultextensionCaller) GetPoolTokens(opts *bind.CallOpts, pool common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getPoolTokens", pool)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetPoolTokens is a free data retrieval call binding the contract method 0xca4f2803.
//
// Solidity: function getPoolTokens(address pool) view returns(address[] tokens)
func (_Vaultextension *VaultextensionSession) GetPoolTokens(pool common.Address) ([]common.Address, error) {
	return _Vaultextension.Contract.GetPoolTokens(&_Vaultextension.CallOpts, pool)
}

// GetPoolTokens is a free data retrieval call binding the contract method 0xca4f2803.
//
// Solidity: function getPoolTokens(address pool) view returns(address[] tokens)
func (_Vaultextension *VaultextensionCallerSession) GetPoolTokens(pool common.Address) ([]common.Address, error) {
	return _Vaultextension.Contract.GetPoolTokens(&_Vaultextension.CallOpts, pool)
}

// GetProtocolFeeController is a free data retrieval call binding the contract method 0x85f2dbd4.
//
// Solidity: function getProtocolFeeController() view returns(address)
func (_Vaultextension *VaultextensionCaller) GetProtocolFeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getProtocolFeeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetProtocolFeeController is a free data retrieval call binding the contract method 0x85f2dbd4.
//
// Solidity: function getProtocolFeeController() view returns(address)
func (_Vaultextension *VaultextensionSession) GetProtocolFeeController() (common.Address, error) {
	return _Vaultextension.Contract.GetProtocolFeeController(&_Vaultextension.CallOpts)
}

// GetProtocolFeeController is a free data retrieval call binding the contract method 0x85f2dbd4.
//
// Solidity: function getProtocolFeeController() view returns(address)
func (_Vaultextension *VaultextensionCallerSession) GetProtocolFeeController() (common.Address, error) {
	return _Vaultextension.Contract.GetProtocolFeeController(&_Vaultextension.CallOpts)
}

// GetReservesOf is a free data retrieval call binding the contract method 0x96787092.
//
// Solidity: function getReservesOf(address token) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) GetReservesOf(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getReservesOf", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReservesOf is a free data retrieval call binding the contract method 0x96787092.
//
// Solidity: function getReservesOf(address token) view returns(uint256)
func (_Vaultextension *VaultextensionSession) GetReservesOf(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetReservesOf(&_Vaultextension.CallOpts, token)
}

// GetReservesOf is a free data retrieval call binding the contract method 0x96787092.
//
// Solidity: function getReservesOf(address token) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) GetReservesOf(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetReservesOf(&_Vaultextension.CallOpts, token)
}

// GetStaticSwapFeePercentage is a free data retrieval call binding the contract method 0xb45090f9.
//
// Solidity: function getStaticSwapFeePercentage(address pool) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) GetStaticSwapFeePercentage(opts *bind.CallOpts, pool common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getStaticSwapFeePercentage", pool)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetStaticSwapFeePercentage is a free data retrieval call binding the contract method 0xb45090f9.
//
// Solidity: function getStaticSwapFeePercentage(address pool) view returns(uint256)
func (_Vaultextension *VaultextensionSession) GetStaticSwapFeePercentage(pool common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetStaticSwapFeePercentage(&_Vaultextension.CallOpts, pool)
}

// GetStaticSwapFeePercentage is a free data retrieval call binding the contract method 0xb45090f9.
//
// Solidity: function getStaticSwapFeePercentage(address pool) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) GetStaticSwapFeePercentage(pool common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetStaticSwapFeePercentage(&_Vaultextension.CallOpts, pool)
}

// GetTokenDelta is a free data retrieval call binding the contract method 0x9e825ff5.
//
// Solidity: function getTokenDelta(address token) view returns(int256)
func (_Vaultextension *VaultextensionCaller) GetTokenDelta(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getTokenDelta", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenDelta is a free data retrieval call binding the contract method 0x9e825ff5.
//
// Solidity: function getTokenDelta(address token) view returns(int256)
func (_Vaultextension *VaultextensionSession) GetTokenDelta(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetTokenDelta(&_Vaultextension.CallOpts, token)
}

// GetTokenDelta is a free data retrieval call binding the contract method 0x9e825ff5.
//
// Solidity: function getTokenDelta(address token) view returns(int256)
func (_Vaultextension *VaultextensionCallerSession) GetTokenDelta(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.GetTokenDelta(&_Vaultextension.CallOpts, token)
}

// GetVaultAdmin is a free data retrieval call binding the contract method 0x1ba0ae45.
//
// Solidity: function getVaultAdmin() view returns(address)
func (_Vaultextension *VaultextensionCaller) GetVaultAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "getVaultAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultAdmin is a free data retrieval call binding the contract method 0x1ba0ae45.
//
// Solidity: function getVaultAdmin() view returns(address)
func (_Vaultextension *VaultextensionSession) GetVaultAdmin() (common.Address, error) {
	return _Vaultextension.Contract.GetVaultAdmin(&_Vaultextension.CallOpts)
}

// GetVaultAdmin is a free data retrieval call binding the contract method 0x1ba0ae45.
//
// Solidity: function getVaultAdmin() view returns(address)
func (_Vaultextension *VaultextensionCallerSession) GetVaultAdmin() (common.Address, error) {
	return _Vaultextension.Contract.GetVaultAdmin(&_Vaultextension.CallOpts)
}

// IsERC4626BufferInitialized is a free data retrieval call binding the contract method 0x6844846b.
//
// Solidity: function isERC4626BufferInitialized(address wrappedToken) view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsERC4626BufferInitialized(opts *bind.CallOpts, wrappedToken common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isERC4626BufferInitialized", wrappedToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsERC4626BufferInitialized is a free data retrieval call binding the contract method 0x6844846b.
//
// Solidity: function isERC4626BufferInitialized(address wrappedToken) view returns(bool)
func (_Vaultextension *VaultextensionSession) IsERC4626BufferInitialized(wrappedToken common.Address) (bool, error) {
	return _Vaultextension.Contract.IsERC4626BufferInitialized(&_Vaultextension.CallOpts, wrappedToken)
}

// IsERC4626BufferInitialized is a free data retrieval call binding the contract method 0x6844846b.
//
// Solidity: function isERC4626BufferInitialized(address wrappedToken) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsERC4626BufferInitialized(wrappedToken common.Address) (bool, error) {
	return _Vaultextension.Contract.IsERC4626BufferInitialized(&_Vaultextension.CallOpts, wrappedToken)
}

// IsPoolInRecoveryMode is a free data retrieval call binding the contract method 0xbe7d628a.
//
// Solidity: function isPoolInRecoveryMode(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsPoolInRecoveryMode(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isPoolInRecoveryMode", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolInRecoveryMode is a free data retrieval call binding the contract method 0xbe7d628a.
//
// Solidity: function isPoolInRecoveryMode(address pool) view returns(bool)
func (_Vaultextension *VaultextensionSession) IsPoolInRecoveryMode(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolInRecoveryMode(&_Vaultextension.CallOpts, pool)
}

// IsPoolInRecoveryMode is a free data retrieval call binding the contract method 0xbe7d628a.
//
// Solidity: function isPoolInRecoveryMode(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsPoolInRecoveryMode(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolInRecoveryMode(&_Vaultextension.CallOpts, pool)
}

// IsPoolInitialized is a free data retrieval call binding the contract method 0x532cec7c.
//
// Solidity: function isPoolInitialized(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsPoolInitialized(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isPoolInitialized", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolInitialized is a free data retrieval call binding the contract method 0x532cec7c.
//
// Solidity: function isPoolInitialized(address pool) view returns(bool)
func (_Vaultextension *VaultextensionSession) IsPoolInitialized(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolInitialized(&_Vaultextension.CallOpts, pool)
}

// IsPoolInitialized is a free data retrieval call binding the contract method 0x532cec7c.
//
// Solidity: function isPoolInitialized(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsPoolInitialized(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolInitialized(&_Vaultextension.CallOpts, pool)
}

// IsPoolPaused is a free data retrieval call binding the contract method 0x6c9bc732.
//
// Solidity: function isPoolPaused(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsPoolPaused(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isPoolPaused", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolPaused is a free data retrieval call binding the contract method 0x6c9bc732.
//
// Solidity: function isPoolPaused(address pool) view returns(bool)
func (_Vaultextension *VaultextensionSession) IsPoolPaused(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolPaused(&_Vaultextension.CallOpts, pool)
}

// IsPoolPaused is a free data retrieval call binding the contract method 0x6c9bc732.
//
// Solidity: function isPoolPaused(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsPoolPaused(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolPaused(&_Vaultextension.CallOpts, pool)
}

// IsPoolRegistered is a free data retrieval call binding the contract method 0xc673bdaf.
//
// Solidity: function isPoolRegistered(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsPoolRegistered(opts *bind.CallOpts, pool common.Address) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isPoolRegistered", pool)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPoolRegistered is a free data retrieval call binding the contract method 0xc673bdaf.
//
// Solidity: function isPoolRegistered(address pool) view returns(bool)
func (_Vaultextension *VaultextensionSession) IsPoolRegistered(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolRegistered(&_Vaultextension.CallOpts, pool)
}

// IsPoolRegistered is a free data retrieval call binding the contract method 0xc673bdaf.
//
// Solidity: function isPoolRegistered(address pool) view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsPoolRegistered(pool common.Address) (bool, error) {
	return _Vaultextension.Contract.IsPoolRegistered(&_Vaultextension.CallOpts, pool)
}

// IsQueryDisabled is a free data retrieval call binding the contract method 0xb4aef0ab.
//
// Solidity: function isQueryDisabled() view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsQueryDisabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isQueryDisabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQueryDisabled is a free data retrieval call binding the contract method 0xb4aef0ab.
//
// Solidity: function isQueryDisabled() view returns(bool)
func (_Vaultextension *VaultextensionSession) IsQueryDisabled() (bool, error) {
	return _Vaultextension.Contract.IsQueryDisabled(&_Vaultextension.CallOpts)
}

// IsQueryDisabled is a free data retrieval call binding the contract method 0xb4aef0ab.
//
// Solidity: function isQueryDisabled() view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsQueryDisabled() (bool, error) {
	return _Vaultextension.Contract.IsQueryDisabled(&_Vaultextension.CallOpts)
}

// IsQueryDisabledPermanently is a free data retrieval call binding the contract method 0x13ef8a5d.
//
// Solidity: function isQueryDisabledPermanently() view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsQueryDisabledPermanently(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isQueryDisabledPermanently")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsQueryDisabledPermanently is a free data retrieval call binding the contract method 0x13ef8a5d.
//
// Solidity: function isQueryDisabledPermanently() view returns(bool)
func (_Vaultextension *VaultextensionSession) IsQueryDisabledPermanently() (bool, error) {
	return _Vaultextension.Contract.IsQueryDisabledPermanently(&_Vaultextension.CallOpts)
}

// IsQueryDisabledPermanently is a free data retrieval call binding the contract method 0x13ef8a5d.
//
// Solidity: function isQueryDisabledPermanently() view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsQueryDisabledPermanently() (bool, error) {
	return _Vaultextension.Contract.IsQueryDisabledPermanently(&_Vaultextension.CallOpts)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Vaultextension *VaultextensionCaller) IsUnlocked(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "isUnlocked")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Vaultextension *VaultextensionSession) IsUnlocked() (bool, error) {
	return _Vaultextension.Contract.IsUnlocked(&_Vaultextension.CallOpts)
}

// IsUnlocked is a free data retrieval call binding the contract method 0x8380edb7.
//
// Solidity: function isUnlocked() view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) IsUnlocked() (bool, error) {
	return _Vaultextension.Contract.IsUnlocked(&_Vaultextension.CallOpts)
}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vaultextension *VaultextensionCaller) ReentrancyGuardEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "reentrancyGuardEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vaultextension *VaultextensionSession) ReentrancyGuardEntered() (bool, error) {
	return _Vaultextension.Contract.ReentrancyGuardEntered(&_Vaultextension.CallOpts)
}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vaultextension *VaultextensionCallerSession) ReentrancyGuardEntered() (bool, error) {
	return _Vaultextension.Contract.ReentrancyGuardEntered(&_Vaultextension.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0xe4dc2aa4.
//
// Solidity: function totalSupply(address token) view returns(uint256)
func (_Vaultextension *VaultextensionCaller) TotalSupply(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "totalSupply", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0xe4dc2aa4.
//
// Solidity: function totalSupply(address token) view returns(uint256)
func (_Vaultextension *VaultextensionSession) TotalSupply(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.TotalSupply(&_Vaultextension.CallOpts, token)
}

// TotalSupply is a free data retrieval call binding the contract method 0xe4dc2aa4.
//
// Solidity: function totalSupply(address token) view returns(uint256)
func (_Vaultextension *VaultextensionCallerSession) TotalSupply(token common.Address) (*big.Int, error) {
	return _Vaultextension.Contract.TotalSupply(&_Vaultextension.CallOpts, token)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Vaultextension *VaultextensionCaller) Vault(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vaultextension.contract.Call(opts, &out, "vault")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Vaultextension *VaultextensionSession) Vault() (common.Address, error) {
	return _Vaultextension.Contract.Vault(&_Vaultextension.CallOpts)
}

// Vault is a free data retrieval call binding the contract method 0xfbfa77cf.
//
// Solidity: function vault() view returns(address)
func (_Vaultextension *VaultextensionCallerSession) Vault() (common.Address, error) {
	return _Vaultextension.Contract.Vault(&_Vaultextension.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address owner, address spender, uint256 amount) returns(bool)
func (_Vaultextension *VaultextensionTransactor) Approve(opts *bind.TransactOpts, owner common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "approve", owner, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address owner, address spender, uint256 amount) returns(bool)
func (_Vaultextension *VaultextensionSession) Approve(owner common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vaultextension.Contract.Approve(&_Vaultextension.TransactOpts, owner, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0xe1f21c67.
//
// Solidity: function approve(address owner, address spender, uint256 amount) returns(bool)
func (_Vaultextension *VaultextensionTransactorSession) Approve(owner common.Address, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vaultextension.Contract.Approve(&_Vaultextension.TransactOpts, owner, spender, amount)
}

// EmitAuxiliaryEvent is a paid mutator transaction binding the contract method 0xc8088247.
//
// Solidity: function emitAuxiliaryEvent(bytes32 eventKey, bytes eventData) returns()
func (_Vaultextension *VaultextensionTransactor) EmitAuxiliaryEvent(opts *bind.TransactOpts, eventKey [32]byte, eventData []byte) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "emitAuxiliaryEvent", eventKey, eventData)
}

// EmitAuxiliaryEvent is a paid mutator transaction binding the contract method 0xc8088247.
//
// Solidity: function emitAuxiliaryEvent(bytes32 eventKey, bytes eventData) returns()
func (_Vaultextension *VaultextensionSession) EmitAuxiliaryEvent(eventKey [32]byte, eventData []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.EmitAuxiliaryEvent(&_Vaultextension.TransactOpts, eventKey, eventData)
}

// EmitAuxiliaryEvent is a paid mutator transaction binding the contract method 0xc8088247.
//
// Solidity: function emitAuxiliaryEvent(bytes32 eventKey, bytes eventData) returns()
func (_Vaultextension *VaultextensionTransactorSession) EmitAuxiliaryEvent(eventKey [32]byte, eventData []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.EmitAuxiliaryEvent(&_Vaultextension.TransactOpts, eventKey, eventData)
}

// Initialize is a paid mutator transaction binding the contract method 0xba8a2be0.
//
// Solidity: function initialize(address pool, address to, address[] tokens, uint256[] exactAmountsIn, uint256 minBptAmountOut, bytes userData) returns(uint256 bptAmountOut)
func (_Vaultextension *VaultextensionTransactor) Initialize(opts *bind.TransactOpts, pool common.Address, to common.Address, tokens []common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "initialize", pool, to, tokens, exactAmountsIn, minBptAmountOut, userData)
}

// Initialize is a paid mutator transaction binding the contract method 0xba8a2be0.
//
// Solidity: function initialize(address pool, address to, address[] tokens, uint256[] exactAmountsIn, uint256 minBptAmountOut, bytes userData) returns(uint256 bptAmountOut)
func (_Vaultextension *VaultextensionSession) Initialize(pool common.Address, to common.Address, tokens []common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Initialize(&_Vaultextension.TransactOpts, pool, to, tokens, exactAmountsIn, minBptAmountOut, userData)
}

// Initialize is a paid mutator transaction binding the contract method 0xba8a2be0.
//
// Solidity: function initialize(address pool, address to, address[] tokens, uint256[] exactAmountsIn, uint256 minBptAmountOut, bytes userData) returns(uint256 bptAmountOut)
func (_Vaultextension *VaultextensionTransactorSession) Initialize(pool common.Address, to common.Address, tokens []common.Address, exactAmountsIn []*big.Int, minBptAmountOut *big.Int, userData []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Initialize(&_Vaultextension.TransactOpts, pool, to, tokens, exactAmountsIn, minBptAmountOut, userData)
}

// Quote is a paid mutator transaction binding the contract method 0xedfa3568.
//
// Solidity: function quote(bytes data) returns(bytes result)
func (_Vaultextension *VaultextensionTransactor) Quote(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "quote", data)
}

// Quote is a paid mutator transaction binding the contract method 0xedfa3568.
//
// Solidity: function quote(bytes data) returns(bytes result)
func (_Vaultextension *VaultextensionSession) Quote(data []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Quote(&_Vaultextension.TransactOpts, data)
}

// Quote is a paid mutator transaction binding the contract method 0xedfa3568.
//
// Solidity: function quote(bytes data) returns(bytes result)
func (_Vaultextension *VaultextensionTransactorSession) Quote(data []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Quote(&_Vaultextension.TransactOpts, data)
}

// QuoteAndRevert is a paid mutator transaction binding the contract method 0x757d64b3.
//
// Solidity: function quoteAndRevert(bytes data) returns()
func (_Vaultextension *VaultextensionTransactor) QuoteAndRevert(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "quoteAndRevert", data)
}

// QuoteAndRevert is a paid mutator transaction binding the contract method 0x757d64b3.
//
// Solidity: function quoteAndRevert(bytes data) returns()
func (_Vaultextension *VaultextensionSession) QuoteAndRevert(data []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.QuoteAndRevert(&_Vaultextension.TransactOpts, data)
}

// QuoteAndRevert is a paid mutator transaction binding the contract method 0x757d64b3.
//
// Solidity: function quoteAndRevert(bytes data) returns()
func (_Vaultextension *VaultextensionTransactorSession) QuoteAndRevert(data []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.QuoteAndRevert(&_Vaultextension.TransactOpts, data)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xeeec802f.
//
// Solidity: function registerPool(address pool, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, bool protocolFeeExempt, (address,address,address) roleAccounts, address poolHooksContract, (bool,bool,bool,bool) liquidityManagement) returns()
func (_Vaultextension *VaultextensionTransactor) RegisterPool(opts *bind.TransactOpts, pool common.Address, tokenConfig []TokenConfig, swapFeePercentage *big.Int, pauseWindowEndTime uint32, protocolFeeExempt bool, roleAccounts PoolRoleAccounts, poolHooksContract common.Address, liquidityManagement LiquidityManagement) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "registerPool", pool, tokenConfig, swapFeePercentage, pauseWindowEndTime, protocolFeeExempt, roleAccounts, poolHooksContract, liquidityManagement)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xeeec802f.
//
// Solidity: function registerPool(address pool, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, bool protocolFeeExempt, (address,address,address) roleAccounts, address poolHooksContract, (bool,bool,bool,bool) liquidityManagement) returns()
func (_Vaultextension *VaultextensionSession) RegisterPool(pool common.Address, tokenConfig []TokenConfig, swapFeePercentage *big.Int, pauseWindowEndTime uint32, protocolFeeExempt bool, roleAccounts PoolRoleAccounts, poolHooksContract common.Address, liquidityManagement LiquidityManagement) (*types.Transaction, error) {
	return _Vaultextension.Contract.RegisterPool(&_Vaultextension.TransactOpts, pool, tokenConfig, swapFeePercentage, pauseWindowEndTime, protocolFeeExempt, roleAccounts, poolHooksContract, liquidityManagement)
}

// RegisterPool is a paid mutator transaction binding the contract method 0xeeec802f.
//
// Solidity: function registerPool(address pool, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, bool protocolFeeExempt, (address,address,address) roleAccounts, address poolHooksContract, (bool,bool,bool,bool) liquidityManagement) returns()
func (_Vaultextension *VaultextensionTransactorSession) RegisterPool(pool common.Address, tokenConfig []TokenConfig, swapFeePercentage *big.Int, pauseWindowEndTime uint32, protocolFeeExempt bool, roleAccounts PoolRoleAccounts, poolHooksContract common.Address, liquidityManagement LiquidityManagement) (*types.Transaction, error) {
	return _Vaultextension.Contract.RegisterPool(&_Vaultextension.TransactOpts, pool, tokenConfig, swapFeePercentage, pauseWindowEndTime, protocolFeeExempt, roleAccounts, poolHooksContract, liquidityManagement)
}

// RemoveLiquidityRecovery is a paid mutator transaction binding the contract method 0xa07d6040.
//
// Solidity: function removeLiquidityRecovery(address pool, address from, uint256 exactBptAmountIn, uint256[] minAmountsOut) returns(uint256[] amountsOutRaw)
func (_Vaultextension *VaultextensionTransactor) RemoveLiquidityRecovery(opts *bind.TransactOpts, pool common.Address, from common.Address, exactBptAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Vaultextension.contract.Transact(opts, "removeLiquidityRecovery", pool, from, exactBptAmountIn, minAmountsOut)
}

// RemoveLiquidityRecovery is a paid mutator transaction binding the contract method 0xa07d6040.
//
// Solidity: function removeLiquidityRecovery(address pool, address from, uint256 exactBptAmountIn, uint256[] minAmountsOut) returns(uint256[] amountsOutRaw)
func (_Vaultextension *VaultextensionSession) RemoveLiquidityRecovery(pool common.Address, from common.Address, exactBptAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Vaultextension.Contract.RemoveLiquidityRecovery(&_Vaultextension.TransactOpts, pool, from, exactBptAmountIn, minAmountsOut)
}

// RemoveLiquidityRecovery is a paid mutator transaction binding the contract method 0xa07d6040.
//
// Solidity: function removeLiquidityRecovery(address pool, address from, uint256 exactBptAmountIn, uint256[] minAmountsOut) returns(uint256[] amountsOutRaw)
func (_Vaultextension *VaultextensionTransactorSession) RemoveLiquidityRecovery(pool common.Address, from common.Address, exactBptAmountIn *big.Int, minAmountsOut []*big.Int) (*types.Transaction, error) {
	return _Vaultextension.Contract.RemoveLiquidityRecovery(&_Vaultextension.TransactOpts, pool, from, exactBptAmountIn, minAmountsOut)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vaultextension *VaultextensionTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Vaultextension.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vaultextension *VaultextensionSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Fallback(&_Vaultextension.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vaultextension *VaultextensionTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vaultextension.Contract.Fallback(&_Vaultextension.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vaultextension *VaultextensionTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vaultextension.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vaultextension *VaultextensionSession) Receive() (*types.Transaction, error) {
	return _Vaultextension.Contract.Receive(&_Vaultextension.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vaultextension *VaultextensionTransactorSession) Receive() (*types.Transaction, error) {
	return _Vaultextension.Contract.Receive(&_Vaultextension.TransactOpts)
}

// VaultextensionAggregateSwapFeePercentageChangedIterator is returned from FilterAggregateSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for AggregateSwapFeePercentageChanged events raised by the Vaultextension contract.
type VaultextensionAggregateSwapFeePercentageChangedIterator struct {
	Event *VaultextensionAggregateSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionAggregateSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionAggregateSwapFeePercentageChanged)
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
		it.Event = new(VaultextensionAggregateSwapFeePercentageChanged)
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
func (it *VaultextensionAggregateSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionAggregateSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionAggregateSwapFeePercentageChanged represents a AggregateSwapFeePercentageChanged event raised by the Vaultextension contract.
type VaultextensionAggregateSwapFeePercentageChanged struct {
	Pool                       common.Address
	AggregateSwapFeePercentage *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAggregateSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xe4d371097beea42453a37406e2aef4c04f3c548f84ac50e72578662c0dcd7354.
//
// Solidity: event AggregateSwapFeePercentageChanged(address indexed pool, uint256 aggregateSwapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) FilterAggregateSwapFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionAggregateSwapFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "AggregateSwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionAggregateSwapFeePercentageChangedIterator{contract: _Vaultextension.contract, event: "AggregateSwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchAggregateSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xe4d371097beea42453a37406e2aef4c04f3c548f84ac50e72578662c0dcd7354.
//
// Solidity: event AggregateSwapFeePercentageChanged(address indexed pool, uint256 aggregateSwapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) WatchAggregateSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionAggregateSwapFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "AggregateSwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionAggregateSwapFeePercentageChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "AggregateSwapFeePercentageChanged", log); err != nil {
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

// ParseAggregateSwapFeePercentageChanged is a log parse operation binding the contract event 0xe4d371097beea42453a37406e2aef4c04f3c548f84ac50e72578662c0dcd7354.
//
// Solidity: event AggregateSwapFeePercentageChanged(address indexed pool, uint256 aggregateSwapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) ParseAggregateSwapFeePercentageChanged(log types.Log) (*VaultextensionAggregateSwapFeePercentageChanged, error) {
	event := new(VaultextensionAggregateSwapFeePercentageChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "AggregateSwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionAggregateYieldFeePercentageChangedIterator is returned from FilterAggregateYieldFeePercentageChanged and is used to iterate over the raw logs and unpacked data for AggregateYieldFeePercentageChanged events raised by the Vaultextension contract.
type VaultextensionAggregateYieldFeePercentageChangedIterator struct {
	Event *VaultextensionAggregateYieldFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionAggregateYieldFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionAggregateYieldFeePercentageChanged)
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
		it.Event = new(VaultextensionAggregateYieldFeePercentageChanged)
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
func (it *VaultextensionAggregateYieldFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionAggregateYieldFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionAggregateYieldFeePercentageChanged represents a AggregateYieldFeePercentageChanged event raised by the Vaultextension contract.
type VaultextensionAggregateYieldFeePercentageChanged struct {
	Pool                        common.Address
	AggregateYieldFeePercentage *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterAggregateYieldFeePercentageChanged is a free log retrieval operation binding the contract event 0x606eb97d83164bd6b200d638cd49c14c65d94d4f2c674cfd85e24e0e202c3ca5.
//
// Solidity: event AggregateYieldFeePercentageChanged(address indexed pool, uint256 aggregateYieldFeePercentage)
func (_Vaultextension *VaultextensionFilterer) FilterAggregateYieldFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionAggregateYieldFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "AggregateYieldFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionAggregateYieldFeePercentageChangedIterator{contract: _Vaultextension.contract, event: "AggregateYieldFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchAggregateYieldFeePercentageChanged is a free log subscription operation binding the contract event 0x606eb97d83164bd6b200d638cd49c14c65d94d4f2c674cfd85e24e0e202c3ca5.
//
// Solidity: event AggregateYieldFeePercentageChanged(address indexed pool, uint256 aggregateYieldFeePercentage)
func (_Vaultextension *VaultextensionFilterer) WatchAggregateYieldFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionAggregateYieldFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "AggregateYieldFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionAggregateYieldFeePercentageChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "AggregateYieldFeePercentageChanged", log); err != nil {
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

// ParseAggregateYieldFeePercentageChanged is a log parse operation binding the contract event 0x606eb97d83164bd6b200d638cd49c14c65d94d4f2c674cfd85e24e0e202c3ca5.
//
// Solidity: event AggregateYieldFeePercentageChanged(address indexed pool, uint256 aggregateYieldFeePercentage)
func (_Vaultextension *VaultextensionFilterer) ParseAggregateYieldFeePercentageChanged(log types.Log) (*VaultextensionAggregateYieldFeePercentageChanged, error) {
	event := new(VaultextensionAggregateYieldFeePercentageChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "AggregateYieldFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vaultextension contract.
type VaultextensionApprovalIterator struct {
	Event *VaultextensionApproval // Event containing the contract specifics and raw log

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
func (it *VaultextensionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionApproval)
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
		it.Event = new(VaultextensionApproval)
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
func (it *VaultextensionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionApproval represents a Approval event raised by the Vaultextension contract.
type VaultextensionApproval struct {
	Pool    common.Address
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xa0175360a15bca328baf7ea85c7b784d58b222a50d0ce760b10dba336d226a61.
//
// Solidity: event Approval(address indexed pool, address indexed owner, address indexed spender, uint256 value)
func (_Vaultextension *VaultextensionFilterer) FilterApproval(opts *bind.FilterOpts, pool []common.Address, owner []common.Address, spender []common.Address) (*VaultextensionApprovalIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "Approval", poolRule, ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionApprovalIterator{contract: _Vaultextension.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xa0175360a15bca328baf7ea85c7b784d58b222a50d0ce760b10dba336d226a61.
//
// Solidity: event Approval(address indexed pool, address indexed owner, address indexed spender, uint256 value)
func (_Vaultextension *VaultextensionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VaultextensionApproval, pool []common.Address, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "Approval", poolRule, ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionApproval)
				if err := _Vaultextension.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xa0175360a15bca328baf7ea85c7b784d58b222a50d0ce760b10dba336d226a61.
//
// Solidity: event Approval(address indexed pool, address indexed owner, address indexed spender, uint256 value)
func (_Vaultextension *VaultextensionFilterer) ParseApproval(log types.Log) (*VaultextensionApproval, error) {
	event := new(VaultextensionApproval)
	if err := _Vaultextension.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Vaultextension contract.
type VaultextensionAuthorizerChangedIterator struct {
	Event *VaultextensionAuthorizerChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionAuthorizerChanged)
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
		it.Event = new(VaultextensionAuthorizerChanged)
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
func (it *VaultextensionAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionAuthorizerChanged represents a AuthorizerChanged event raised by the Vaultextension contract.
type VaultextensionAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vaultextension *VaultextensionFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*VaultextensionAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionAuthorizerChangedIterator{contract: _Vaultextension.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vaultextension *VaultextensionFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionAuthorizerChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
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

// ParseAuthorizerChanged is a log parse operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vaultextension *VaultextensionFilterer) ParseAuthorizerChanged(log types.Log) (*VaultextensionAuthorizerChanged, error) {
	event := new(VaultextensionAuthorizerChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionBufferSharesBurnedIterator is returned from FilterBufferSharesBurned and is used to iterate over the raw logs and unpacked data for BufferSharesBurned events raised by the Vaultextension contract.
type VaultextensionBufferSharesBurnedIterator struct {
	Event *VaultextensionBufferSharesBurned // Event containing the contract specifics and raw log

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
func (it *VaultextensionBufferSharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionBufferSharesBurned)
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
		it.Event = new(VaultextensionBufferSharesBurned)
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
func (it *VaultextensionBufferSharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionBufferSharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionBufferSharesBurned represents a BufferSharesBurned event raised by the Vaultextension contract.
type VaultextensionBufferSharesBurned struct {
	WrappedToken common.Address
	From         common.Address
	BurnedShares *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferSharesBurned is a free log retrieval operation binding the contract event 0x4e09f7f7fc37ce2897800e2c2a9099565edb0a133d19d84a6871b3530af8846b.
//
// Solidity: event BufferSharesBurned(address indexed wrappedToken, address indexed from, uint256 burnedShares)
func (_Vaultextension *VaultextensionFilterer) FilterBufferSharesBurned(opts *bind.FilterOpts, wrappedToken []common.Address, from []common.Address) (*VaultextensionBufferSharesBurnedIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "BufferSharesBurned", wrappedTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionBufferSharesBurnedIterator{contract: _Vaultextension.contract, event: "BufferSharesBurned", logs: logs, sub: sub}, nil
}

// WatchBufferSharesBurned is a free log subscription operation binding the contract event 0x4e09f7f7fc37ce2897800e2c2a9099565edb0a133d19d84a6871b3530af8846b.
//
// Solidity: event BufferSharesBurned(address indexed wrappedToken, address indexed from, uint256 burnedShares)
func (_Vaultextension *VaultextensionFilterer) WatchBufferSharesBurned(opts *bind.WatchOpts, sink chan<- *VaultextensionBufferSharesBurned, wrappedToken []common.Address, from []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "BufferSharesBurned", wrappedTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionBufferSharesBurned)
				if err := _Vaultextension.contract.UnpackLog(event, "BufferSharesBurned", log); err != nil {
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

// ParseBufferSharesBurned is a log parse operation binding the contract event 0x4e09f7f7fc37ce2897800e2c2a9099565edb0a133d19d84a6871b3530af8846b.
//
// Solidity: event BufferSharesBurned(address indexed wrappedToken, address indexed from, uint256 burnedShares)
func (_Vaultextension *VaultextensionFilterer) ParseBufferSharesBurned(log types.Log) (*VaultextensionBufferSharesBurned, error) {
	event := new(VaultextensionBufferSharesBurned)
	if err := _Vaultextension.contract.UnpackLog(event, "BufferSharesBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionBufferSharesMintedIterator is returned from FilterBufferSharesMinted and is used to iterate over the raw logs and unpacked data for BufferSharesMinted events raised by the Vaultextension contract.
type VaultextensionBufferSharesMintedIterator struct {
	Event *VaultextensionBufferSharesMinted // Event containing the contract specifics and raw log

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
func (it *VaultextensionBufferSharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionBufferSharesMinted)
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
		it.Event = new(VaultextensionBufferSharesMinted)
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
func (it *VaultextensionBufferSharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionBufferSharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionBufferSharesMinted represents a BufferSharesMinted event raised by the Vaultextension contract.
type VaultextensionBufferSharesMinted struct {
	WrappedToken common.Address
	To           common.Address
	IssuedShares *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferSharesMinted is a free log retrieval operation binding the contract event 0xd66f031d33381c6408f0b32c884461e5de3df8808399b6f3a3d86b1368f8ec34.
//
// Solidity: event BufferSharesMinted(address indexed wrappedToken, address indexed to, uint256 issuedShares)
func (_Vaultextension *VaultextensionFilterer) FilterBufferSharesMinted(opts *bind.FilterOpts, wrappedToken []common.Address, to []common.Address) (*VaultextensionBufferSharesMintedIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "BufferSharesMinted", wrappedTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionBufferSharesMintedIterator{contract: _Vaultextension.contract, event: "BufferSharesMinted", logs: logs, sub: sub}, nil
}

// WatchBufferSharesMinted is a free log subscription operation binding the contract event 0xd66f031d33381c6408f0b32c884461e5de3df8808399b6f3a3d86b1368f8ec34.
//
// Solidity: event BufferSharesMinted(address indexed wrappedToken, address indexed to, uint256 issuedShares)
func (_Vaultextension *VaultextensionFilterer) WatchBufferSharesMinted(opts *bind.WatchOpts, sink chan<- *VaultextensionBufferSharesMinted, wrappedToken []common.Address, to []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "BufferSharesMinted", wrappedTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionBufferSharesMinted)
				if err := _Vaultextension.contract.UnpackLog(event, "BufferSharesMinted", log); err != nil {
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

// ParseBufferSharesMinted is a log parse operation binding the contract event 0xd66f031d33381c6408f0b32c884461e5de3df8808399b6f3a3d86b1368f8ec34.
//
// Solidity: event BufferSharesMinted(address indexed wrappedToken, address indexed to, uint256 issuedShares)
func (_Vaultextension *VaultextensionFilterer) ParseBufferSharesMinted(log types.Log) (*VaultextensionBufferSharesMinted, error) {
	event := new(VaultextensionBufferSharesMinted)
	if err := _Vaultextension.contract.UnpackLog(event, "BufferSharesMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Vaultextension contract.
type VaultextensionLiquidityAddedIterator struct {
	Event *VaultextensionLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *VaultextensionLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionLiquidityAdded)
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
		it.Event = new(VaultextensionLiquidityAdded)
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
func (it *VaultextensionLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionLiquidityAdded represents a LiquidityAdded event raised by the Vaultextension contract.
type VaultextensionLiquidityAdded struct {
	Pool              common.Address
	LiquidityProvider common.Address
	Kind              uint8
	TotalSupply       *big.Int
	AmountsAddedRaw   []*big.Int
	SwapFeeAmountsRaw []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAdded is a free log retrieval operation binding the contract event 0xa26a52d8d53702bba7f137907b8e1f99ff87f6d450144270ca25e72481cca871.
//
// Solidity: event LiquidityAdded(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsAddedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (*VaultextensionLiquidityAddedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "LiquidityAdded", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionLiquidityAddedIterator{contract: _Vaultextension.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xa26a52d8d53702bba7f137907b8e1f99ff87f6d450144270ca25e72481cca871.
//
// Solidity: event LiquidityAdded(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsAddedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *VaultextensionLiquidityAdded, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "LiquidityAdded", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionLiquidityAdded)
				if err := _Vaultextension.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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

// ParseLiquidityAdded is a log parse operation binding the contract event 0xa26a52d8d53702bba7f137907b8e1f99ff87f6d450144270ca25e72481cca871.
//
// Solidity: event LiquidityAdded(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsAddedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) ParseLiquidityAdded(log types.Log) (*VaultextensionLiquidityAdded, error) {
	event := new(VaultextensionLiquidityAdded)
	if err := _Vaultextension.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionLiquidityAddedToBufferIterator is returned from FilterLiquidityAddedToBuffer and is used to iterate over the raw logs and unpacked data for LiquidityAddedToBuffer events raised by the Vaultextension contract.
type VaultextensionLiquidityAddedToBufferIterator struct {
	Event *VaultextensionLiquidityAddedToBuffer // Event containing the contract specifics and raw log

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
func (it *VaultextensionLiquidityAddedToBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionLiquidityAddedToBuffer)
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
		it.Event = new(VaultextensionLiquidityAddedToBuffer)
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
func (it *VaultextensionLiquidityAddedToBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionLiquidityAddedToBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionLiquidityAddedToBuffer represents a LiquidityAddedToBuffer event raised by the Vaultextension contract.
type VaultextensionLiquidityAddedToBuffer struct {
	WrappedToken     common.Address
	AmountUnderlying *big.Int
	AmountWrapped    *big.Int
	BufferBalances   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAddedToBuffer is a free log retrieval operation binding the contract event 0x75c4dc5f23640eeba7d404d9165f515fc3d9e23a5c8b6e2d09b4b9da56ff00a9.
//
// Solidity: event LiquidityAddedToBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) FilterLiquidityAddedToBuffer(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultextensionLiquidityAddedToBufferIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "LiquidityAddedToBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionLiquidityAddedToBufferIterator{contract: _Vaultextension.contract, event: "LiquidityAddedToBuffer", logs: logs, sub: sub}, nil
}

// WatchLiquidityAddedToBuffer is a free log subscription operation binding the contract event 0x75c4dc5f23640eeba7d404d9165f515fc3d9e23a5c8b6e2d09b4b9da56ff00a9.
//
// Solidity: event LiquidityAddedToBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) WatchLiquidityAddedToBuffer(opts *bind.WatchOpts, sink chan<- *VaultextensionLiquidityAddedToBuffer, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "LiquidityAddedToBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionLiquidityAddedToBuffer)
				if err := _Vaultextension.contract.UnpackLog(event, "LiquidityAddedToBuffer", log); err != nil {
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

// ParseLiquidityAddedToBuffer is a log parse operation binding the contract event 0x75c4dc5f23640eeba7d404d9165f515fc3d9e23a5c8b6e2d09b4b9da56ff00a9.
//
// Solidity: event LiquidityAddedToBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) ParseLiquidityAddedToBuffer(log types.Log) (*VaultextensionLiquidityAddedToBuffer, error) {
	event := new(VaultextensionLiquidityAddedToBuffer)
	if err := _Vaultextension.contract.UnpackLog(event, "LiquidityAddedToBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the Vaultextension contract.
type VaultextensionLiquidityRemovedIterator struct {
	Event *VaultextensionLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *VaultextensionLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionLiquidityRemoved)
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
		it.Event = new(VaultextensionLiquidityRemoved)
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
func (it *VaultextensionLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionLiquidityRemoved represents a LiquidityRemoved event raised by the Vaultextension contract.
type VaultextensionLiquidityRemoved struct {
	Pool              common.Address
	LiquidityProvider common.Address
	Kind              uint8
	TotalSupply       *big.Int
	AmountsRemovedRaw []*big.Int
	SwapFeeAmountsRaw []*big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemoved is a free log retrieval operation binding the contract event 0xfbe5b0d79fb94f1e81c0a92bf86ae9d3a19e9d1bf6202c0d3e75120f65d5d8a5.
//
// Solidity: event LiquidityRemoved(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsRemovedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (*VaultextensionLiquidityRemovedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "LiquidityRemoved", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionLiquidityRemovedIterator{contract: _Vaultextension.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xfbe5b0d79fb94f1e81c0a92bf86ae9d3a19e9d1bf6202c0d3e75120f65d5d8a5.
//
// Solidity: event LiquidityRemoved(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsRemovedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *VaultextensionLiquidityRemoved, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var liquidityProviderRule []interface{}
	for _, liquidityProviderItem := range liquidityProvider {
		liquidityProviderRule = append(liquidityProviderRule, liquidityProviderItem)
	}
	var kindRule []interface{}
	for _, kindItem := range kind {
		kindRule = append(kindRule, kindItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "LiquidityRemoved", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionLiquidityRemoved)
				if err := _Vaultextension.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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

// ParseLiquidityRemoved is a log parse operation binding the contract event 0xfbe5b0d79fb94f1e81c0a92bf86ae9d3a19e9d1bf6202c0d3e75120f65d5d8a5.
//
// Solidity: event LiquidityRemoved(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsRemovedRaw, uint256[] swapFeeAmountsRaw)
func (_Vaultextension *VaultextensionFilterer) ParseLiquidityRemoved(log types.Log) (*VaultextensionLiquidityRemoved, error) {
	event := new(VaultextensionLiquidityRemoved)
	if err := _Vaultextension.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionLiquidityRemovedFromBufferIterator is returned from FilterLiquidityRemovedFromBuffer and is used to iterate over the raw logs and unpacked data for LiquidityRemovedFromBuffer events raised by the Vaultextension contract.
type VaultextensionLiquidityRemovedFromBufferIterator struct {
	Event *VaultextensionLiquidityRemovedFromBuffer // Event containing the contract specifics and raw log

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
func (it *VaultextensionLiquidityRemovedFromBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionLiquidityRemovedFromBuffer)
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
		it.Event = new(VaultextensionLiquidityRemovedFromBuffer)
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
func (it *VaultextensionLiquidityRemovedFromBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionLiquidityRemovedFromBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionLiquidityRemovedFromBuffer represents a LiquidityRemovedFromBuffer event raised by the Vaultextension contract.
type VaultextensionLiquidityRemovedFromBuffer struct {
	WrappedToken     common.Address
	AmountUnderlying *big.Int
	AmountWrapped    *big.Int
	BufferBalances   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemovedFromBuffer is a free log retrieval operation binding the contract event 0x44d97b36e99b590b3d2875aad3b167b1d7fb1e063f3f1325a1eeac76caee5113.
//
// Solidity: event LiquidityRemovedFromBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) FilterLiquidityRemovedFromBuffer(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultextensionLiquidityRemovedFromBufferIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "LiquidityRemovedFromBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionLiquidityRemovedFromBufferIterator{contract: _Vaultextension.contract, event: "LiquidityRemovedFromBuffer", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemovedFromBuffer is a free log subscription operation binding the contract event 0x44d97b36e99b590b3d2875aad3b167b1d7fb1e063f3f1325a1eeac76caee5113.
//
// Solidity: event LiquidityRemovedFromBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) WatchLiquidityRemovedFromBuffer(opts *bind.WatchOpts, sink chan<- *VaultextensionLiquidityRemovedFromBuffer, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "LiquidityRemovedFromBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionLiquidityRemovedFromBuffer)
				if err := _Vaultextension.contract.UnpackLog(event, "LiquidityRemovedFromBuffer", log); err != nil {
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

// ParseLiquidityRemovedFromBuffer is a log parse operation binding the contract event 0x44d97b36e99b590b3d2875aad3b167b1d7fb1e063f3f1325a1eeac76caee5113.
//
// Solidity: event LiquidityRemovedFromBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) ParseLiquidityRemovedFromBuffer(log types.Log) (*VaultextensionLiquidityRemovedFromBuffer, error) {
	event := new(VaultextensionLiquidityRemovedFromBuffer)
	if err := _Vaultextension.contract.UnpackLog(event, "LiquidityRemovedFromBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the Vaultextension contract.
type VaultextensionPoolInitializedIterator struct {
	Event *VaultextensionPoolInitialized // Event containing the contract specifics and raw log

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
func (it *VaultextensionPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionPoolInitialized)
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
		it.Event = new(VaultextensionPoolInitialized)
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
func (it *VaultextensionPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionPoolInitialized represents a PoolInitialized event raised by the Vaultextension contract.
type VaultextensionPoolInitialized struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed pool)
func (_Vaultextension *VaultextensionFilterer) FilterPoolInitialized(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionPoolInitializedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "PoolInitialized", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionPoolInitializedIterator{contract: _Vaultextension.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed pool)
func (_Vaultextension *VaultextensionFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *VaultextensionPoolInitialized, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "PoolInitialized", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionPoolInitialized)
				if err := _Vaultextension.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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

// ParsePoolInitialized is a log parse operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed pool)
func (_Vaultextension *VaultextensionFilterer) ParsePoolInitialized(log types.Log) (*VaultextensionPoolInitialized, error) {
	event := new(VaultextensionPoolInitialized)
	if err := _Vaultextension.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionPoolPausedStateChangedIterator is returned from FilterPoolPausedStateChanged and is used to iterate over the raw logs and unpacked data for PoolPausedStateChanged events raised by the Vaultextension contract.
type VaultextensionPoolPausedStateChangedIterator struct {
	Event *VaultextensionPoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionPoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionPoolPausedStateChanged)
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
		it.Event = new(VaultextensionPoolPausedStateChanged)
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
func (it *VaultextensionPoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionPoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionPoolPausedStateChanged represents a PoolPausedStateChanged event raised by the Vaultextension contract.
type VaultextensionPoolPausedStateChanged struct {
	Pool   common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolPausedStateChanged is a free log retrieval operation binding the contract event 0x57e20448028297190122571be7cb6c1b1ef85730c673f7c72f533c8662419aa7.
//
// Solidity: event PoolPausedStateChanged(address indexed pool, bool paused)
func (_Vaultextension *VaultextensionFilterer) FilterPoolPausedStateChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionPoolPausedStateChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "PoolPausedStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionPoolPausedStateChangedIterator{contract: _Vaultextension.contract, event: "PoolPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPoolPausedStateChanged is a free log subscription operation binding the contract event 0x57e20448028297190122571be7cb6c1b1ef85730c673f7c72f533c8662419aa7.
//
// Solidity: event PoolPausedStateChanged(address indexed pool, bool paused)
func (_Vaultextension *VaultextensionFilterer) WatchPoolPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionPoolPausedStateChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "PoolPausedStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionPoolPausedStateChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "PoolPausedStateChanged", log); err != nil {
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

// ParsePoolPausedStateChanged is a log parse operation binding the contract event 0x57e20448028297190122571be7cb6c1b1ef85730c673f7c72f533c8662419aa7.
//
// Solidity: event PoolPausedStateChanged(address indexed pool, bool paused)
func (_Vaultextension *VaultextensionFilterer) ParsePoolPausedStateChanged(log types.Log) (*VaultextensionPoolPausedStateChanged, error) {
	event := new(VaultextensionPoolPausedStateChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "PoolPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionPoolRecoveryModeStateChangedIterator is returned from FilterPoolRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for PoolRecoveryModeStateChanged events raised by the Vaultextension contract.
type VaultextensionPoolRecoveryModeStateChangedIterator struct {
	Event *VaultextensionPoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionPoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionPoolRecoveryModeStateChanged)
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
		it.Event = new(VaultextensionPoolRecoveryModeStateChanged)
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
func (it *VaultextensionPoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionPoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionPoolRecoveryModeStateChanged represents a PoolRecoveryModeStateChanged event raised by the Vaultextension contract.
type VaultextensionPoolRecoveryModeStateChanged struct {
	Pool         common.Address
	RecoveryMode bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xc2354cc2f78ea57777e55ddd43a7f22b112ce98868596880edaeb22b4f9c73a9.
//
// Solidity: event PoolRecoveryModeStateChanged(address indexed pool, bool recoveryMode)
func (_Vaultextension *VaultextensionFilterer) FilterPoolRecoveryModeStateChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionPoolRecoveryModeStateChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "PoolRecoveryModeStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionPoolRecoveryModeStateChangedIterator{contract: _Vaultextension.contract, event: "PoolRecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchPoolRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xc2354cc2f78ea57777e55ddd43a7f22b112ce98868596880edaeb22b4f9c73a9.
//
// Solidity: event PoolRecoveryModeStateChanged(address indexed pool, bool recoveryMode)
func (_Vaultextension *VaultextensionFilterer) WatchPoolRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionPoolRecoveryModeStateChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "PoolRecoveryModeStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionPoolRecoveryModeStateChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "PoolRecoveryModeStateChanged", log); err != nil {
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

// ParsePoolRecoveryModeStateChanged is a log parse operation binding the contract event 0xc2354cc2f78ea57777e55ddd43a7f22b112ce98868596880edaeb22b4f9c73a9.
//
// Solidity: event PoolRecoveryModeStateChanged(address indexed pool, bool recoveryMode)
func (_Vaultextension *VaultextensionFilterer) ParsePoolRecoveryModeStateChanged(log types.Log) (*VaultextensionPoolRecoveryModeStateChanged, error) {
	event := new(VaultextensionPoolRecoveryModeStateChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "PoolRecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Vaultextension contract.
type VaultextensionPoolRegisteredIterator struct {
	Event *VaultextensionPoolRegistered // Event containing the contract specifics and raw log

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
func (it *VaultextensionPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionPoolRegistered)
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
		it.Event = new(VaultextensionPoolRegistered)
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
func (it *VaultextensionPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionPoolRegistered represents a PoolRegistered event raised by the Vaultextension contract.
type VaultextensionPoolRegistered struct {
	Pool                common.Address
	Factory             common.Address
	TokenConfig         []TokenConfig
	SwapFeePercentage   *big.Int
	PauseWindowEndTime  uint32
	RoleAccounts        PoolRoleAccounts
	HooksConfig         HooksConfig
	LiquidityManagement LiquidityManagement
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterPoolRegistered is a free log retrieval operation binding the contract event 0xbc1561eeab9f40962e2fb827a7ff9c7cdb47a9d7c84caeefa4ed90e043842dad.
//
// Solidity: event PoolRegistered(address indexed pool, address indexed factory, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, (address,address,address) roleAccounts, (bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address) hooksConfig, (bool,bool,bool,bool) liquidityManagement)
func (_Vaultextension *VaultextensionFilterer) FilterPoolRegistered(opts *bind.FilterOpts, pool []common.Address, factory []common.Address) (*VaultextensionPoolRegisteredIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "PoolRegistered", poolRule, factoryRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionPoolRegisteredIterator{contract: _Vaultextension.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0xbc1561eeab9f40962e2fb827a7ff9c7cdb47a9d7c84caeefa4ed90e043842dad.
//
// Solidity: event PoolRegistered(address indexed pool, address indexed factory, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, (address,address,address) roleAccounts, (bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address) hooksConfig, (bool,bool,bool,bool) liquidityManagement)
func (_Vaultextension *VaultextensionFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *VaultextensionPoolRegistered, pool []common.Address, factory []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "PoolRegistered", poolRule, factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionPoolRegistered)
				if err := _Vaultextension.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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

// ParsePoolRegistered is a log parse operation binding the contract event 0xbc1561eeab9f40962e2fb827a7ff9c7cdb47a9d7c84caeefa4ed90e043842dad.
//
// Solidity: event PoolRegistered(address indexed pool, address indexed factory, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, (address,address,address) roleAccounts, (bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address) hooksConfig, (bool,bool,bool,bool) liquidityManagement)
func (_Vaultextension *VaultextensionFilterer) ParsePoolRegistered(log types.Log) (*VaultextensionPoolRegistered, error) {
	event := new(VaultextensionPoolRegistered)
	if err := _Vaultextension.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionProtocolFeeControllerChangedIterator is returned from FilterProtocolFeeControllerChanged and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerChanged events raised by the Vaultextension contract.
type VaultextensionProtocolFeeControllerChangedIterator struct {
	Event *VaultextensionProtocolFeeControllerChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionProtocolFeeControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionProtocolFeeControllerChanged)
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
		it.Event = new(VaultextensionProtocolFeeControllerChanged)
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
func (it *VaultextensionProtocolFeeControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionProtocolFeeControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionProtocolFeeControllerChanged represents a ProtocolFeeControllerChanged event raised by the Vaultextension contract.
type VaultextensionProtocolFeeControllerChanged struct {
	NewProtocolFeeController common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerChanged is a free log retrieval operation binding the contract event 0x280a60b1e63c1774d397d35cce80eb80e51408ead755fb446e6f744ce98e5df0.
//
// Solidity: event ProtocolFeeControllerChanged(address indexed newProtocolFeeController)
func (_Vaultextension *VaultextensionFilterer) FilterProtocolFeeControllerChanged(opts *bind.FilterOpts, newProtocolFeeController []common.Address) (*VaultextensionProtocolFeeControllerChangedIterator, error) {

	var newProtocolFeeControllerRule []interface{}
	for _, newProtocolFeeControllerItem := range newProtocolFeeController {
		newProtocolFeeControllerRule = append(newProtocolFeeControllerRule, newProtocolFeeControllerItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "ProtocolFeeControllerChanged", newProtocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionProtocolFeeControllerChangedIterator{contract: _Vaultextension.contract, event: "ProtocolFeeControllerChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerChanged is a free log subscription operation binding the contract event 0x280a60b1e63c1774d397d35cce80eb80e51408ead755fb446e6f744ce98e5df0.
//
// Solidity: event ProtocolFeeControllerChanged(address indexed newProtocolFeeController)
func (_Vaultextension *VaultextensionFilterer) WatchProtocolFeeControllerChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionProtocolFeeControllerChanged, newProtocolFeeController []common.Address) (event.Subscription, error) {

	var newProtocolFeeControllerRule []interface{}
	for _, newProtocolFeeControllerItem := range newProtocolFeeController {
		newProtocolFeeControllerRule = append(newProtocolFeeControllerRule, newProtocolFeeControllerItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "ProtocolFeeControllerChanged", newProtocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionProtocolFeeControllerChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "ProtocolFeeControllerChanged", log); err != nil {
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

// ParseProtocolFeeControllerChanged is a log parse operation binding the contract event 0x280a60b1e63c1774d397d35cce80eb80e51408ead755fb446e6f744ce98e5df0.
//
// Solidity: event ProtocolFeeControllerChanged(address indexed newProtocolFeeController)
func (_Vaultextension *VaultextensionFilterer) ParseProtocolFeeControllerChanged(log types.Log) (*VaultextensionProtocolFeeControllerChanged, error) {
	event := new(VaultextensionProtocolFeeControllerChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "ProtocolFeeControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Vaultextension contract.
type VaultextensionSwapIterator struct {
	Event *VaultextensionSwap // Event containing the contract specifics and raw log

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
func (it *VaultextensionSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionSwap)
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
		it.Event = new(VaultextensionSwap)
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
func (it *VaultextensionSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionSwap represents a Swap event raised by the Vaultextension contract.
type VaultextensionSwap struct {
	Pool              common.Address
	TokenIn           common.Address
	TokenOut          common.Address
	AmountIn          *big.Int
	AmountOut         *big.Int
	SwapFeePercentage *big.Int
	SwapFeeAmount     *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address indexed pool, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, uint256 swapFeePercentage, uint256 swapFeeAmount)
func (_Vaultextension *VaultextensionFilterer) FilterSwap(opts *bind.FilterOpts, pool []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*VaultextensionSwapIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "Swap", poolRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionSwapIterator{contract: _Vaultextension.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address indexed pool, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, uint256 swapFeePercentage, uint256 swapFeeAmount)
func (_Vaultextension *VaultextensionFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultextensionSwap, pool []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var tokenInRule []interface{}
	for _, tokenInItem := range tokenIn {
		tokenInRule = append(tokenInRule, tokenInItem)
	}
	var tokenOutRule []interface{}
	for _, tokenOutItem := range tokenOut {
		tokenOutRule = append(tokenOutRule, tokenOutItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "Swap", poolRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionSwap)
				if err := _Vaultextension.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address indexed pool, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, uint256 swapFeePercentage, uint256 swapFeeAmount)
func (_Vaultextension *VaultextensionFilterer) ParseSwap(log types.Log) (*VaultextensionSwap, error) {
	event := new(VaultextensionSwap)
	if err := _Vaultextension.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the Vaultextension contract.
type VaultextensionSwapFeePercentageChangedIterator struct {
	Event *VaultextensionSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionSwapFeePercentageChanged)
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
		it.Event = new(VaultextensionSwapFeePercentageChanged)
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
func (it *VaultextensionSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the Vaultextension contract.
type VaultextensionSwapFeePercentageChanged struct {
	Pool              common.Address
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0x89d41522342fabac1471ca6073a5623e5caf367b03ca6e9a001478d0cf8be4a1.
//
// Solidity: event SwapFeePercentageChanged(address indexed pool, uint256 swapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultextensionSwapFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "SwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionSwapFeePercentageChangedIterator{contract: _Vaultextension.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0x89d41522342fabac1471ca6073a5623e5caf367b03ca6e9a001478d0cf8be4a1.
//
// Solidity: event SwapFeePercentageChanged(address indexed pool, uint256 swapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionSwapFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "SwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionSwapFeePercentageChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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

// ParseSwapFeePercentageChanged is a log parse operation binding the contract event 0x89d41522342fabac1471ca6073a5623e5caf367b03ca6e9a001478d0cf8be4a1.
//
// Solidity: event SwapFeePercentageChanged(address indexed pool, uint256 swapFeePercentage)
func (_Vaultextension *VaultextensionFilterer) ParseSwapFeePercentageChanged(log types.Log) (*VaultextensionSwapFeePercentageChanged, error) {
	event := new(VaultextensionSwapFeePercentageChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vaultextension contract.
type VaultextensionTransferIterator struct {
	Event *VaultextensionTransfer // Event containing the contract specifics and raw log

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
func (it *VaultextensionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionTransfer)
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
		it.Event = new(VaultextensionTransfer)
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
func (it *VaultextensionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionTransfer represents a Transfer event raised by the Vaultextension contract.
type VaultextensionTransfer struct {
	Pool  common.Address
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed pool, address indexed from, address indexed to, uint256 value)
func (_Vaultextension *VaultextensionFilterer) FilterTransfer(opts *bind.FilterOpts, pool []common.Address, from []common.Address, to []common.Address) (*VaultextensionTransferIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "Transfer", poolRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionTransferIterator{contract: _Vaultextension.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed pool, address indexed from, address indexed to, uint256 value)
func (_Vaultextension *VaultextensionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaultextensionTransfer, pool []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "Transfer", poolRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionTransfer)
				if err := _Vaultextension.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed pool, address indexed from, address indexed to, uint256 value)
func (_Vaultextension *VaultextensionFilterer) ParseTransfer(log types.Log) (*VaultextensionTransfer, error) {
	event := new(VaultextensionTransfer)
	if err := _Vaultextension.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionUnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the Vaultextension contract.
type VaultextensionUnwrapIterator struct {
	Event *VaultextensionUnwrap // Event containing the contract specifics and raw log

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
func (it *VaultextensionUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionUnwrap)
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
		it.Event = new(VaultextensionUnwrap)
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
func (it *VaultextensionUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionUnwrap represents a Unwrap event raised by the Vaultextension contract.
type VaultextensionUnwrap struct {
	WrappedToken        common.Address
	BurnedShares        *big.Int
	WithdrawnUnderlying *big.Int
	BufferBalances      [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0xeeb740c90bf2b18c9532eb7d473137767036d893dff3e009f32718f821b2a4c0.
//
// Solidity: event Unwrap(address indexed wrappedToken, uint256 burnedShares, uint256 withdrawnUnderlying, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) FilterUnwrap(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultextensionUnwrapIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "Unwrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionUnwrapIterator{contract: _Vaultextension.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0xeeb740c90bf2b18c9532eb7d473137767036d893dff3e009f32718f821b2a4c0.
//
// Solidity: event Unwrap(address indexed wrappedToken, uint256 burnedShares, uint256 withdrawnUnderlying, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *VaultextensionUnwrap, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "Unwrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionUnwrap)
				if err := _Vaultextension.contract.UnpackLog(event, "Unwrap", log); err != nil {
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

// ParseUnwrap is a log parse operation binding the contract event 0xeeb740c90bf2b18c9532eb7d473137767036d893dff3e009f32718f821b2a4c0.
//
// Solidity: event Unwrap(address indexed wrappedToken, uint256 burnedShares, uint256 withdrawnUnderlying, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) ParseUnwrap(log types.Log) (*VaultextensionUnwrap, error) {
	event := new(VaultextensionUnwrap)
	if err := _Vaultextension.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionVaultAuxiliaryIterator is returned from FilterVaultAuxiliary and is used to iterate over the raw logs and unpacked data for VaultAuxiliary events raised by the Vaultextension contract.
type VaultextensionVaultAuxiliaryIterator struct {
	Event *VaultextensionVaultAuxiliary // Event containing the contract specifics and raw log

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
func (it *VaultextensionVaultAuxiliaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionVaultAuxiliary)
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
		it.Event = new(VaultextensionVaultAuxiliary)
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
func (it *VaultextensionVaultAuxiliaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionVaultAuxiliaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionVaultAuxiliary represents a VaultAuxiliary event raised by the Vaultextension contract.
type VaultextensionVaultAuxiliary struct {
	Pool      common.Address
	EventKey  [32]byte
	EventData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVaultAuxiliary is a free log retrieval operation binding the contract event 0x4bc4412e210115456903c65b5277d299a505e79f2eb852b92b1ca52d85856428.
//
// Solidity: event VaultAuxiliary(address indexed pool, bytes32 indexed eventKey, bytes eventData)
func (_Vaultextension *VaultextensionFilterer) FilterVaultAuxiliary(opts *bind.FilterOpts, pool []common.Address, eventKey [][32]byte) (*VaultextensionVaultAuxiliaryIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var eventKeyRule []interface{}
	for _, eventKeyItem := range eventKey {
		eventKeyRule = append(eventKeyRule, eventKeyItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "VaultAuxiliary", poolRule, eventKeyRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionVaultAuxiliaryIterator{contract: _Vaultextension.contract, event: "VaultAuxiliary", logs: logs, sub: sub}, nil
}

// WatchVaultAuxiliary is a free log subscription operation binding the contract event 0x4bc4412e210115456903c65b5277d299a505e79f2eb852b92b1ca52d85856428.
//
// Solidity: event VaultAuxiliary(address indexed pool, bytes32 indexed eventKey, bytes eventData)
func (_Vaultextension *VaultextensionFilterer) WatchVaultAuxiliary(opts *bind.WatchOpts, sink chan<- *VaultextensionVaultAuxiliary, pool []common.Address, eventKey [][32]byte) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var eventKeyRule []interface{}
	for _, eventKeyItem := range eventKey {
		eventKeyRule = append(eventKeyRule, eventKeyItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "VaultAuxiliary", poolRule, eventKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionVaultAuxiliary)
				if err := _Vaultextension.contract.UnpackLog(event, "VaultAuxiliary", log); err != nil {
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

// ParseVaultAuxiliary is a log parse operation binding the contract event 0x4bc4412e210115456903c65b5277d299a505e79f2eb852b92b1ca52d85856428.
//
// Solidity: event VaultAuxiliary(address indexed pool, bytes32 indexed eventKey, bytes eventData)
func (_Vaultextension *VaultextensionFilterer) ParseVaultAuxiliary(log types.Log) (*VaultextensionVaultAuxiliary, error) {
	event := new(VaultextensionVaultAuxiliary)
	if err := _Vaultextension.contract.UnpackLog(event, "VaultAuxiliary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionVaultBuffersPausedStateChangedIterator is returned from FilterVaultBuffersPausedStateChanged and is used to iterate over the raw logs and unpacked data for VaultBuffersPausedStateChanged events raised by the Vaultextension contract.
type VaultextensionVaultBuffersPausedStateChangedIterator struct {
	Event *VaultextensionVaultBuffersPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionVaultBuffersPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionVaultBuffersPausedStateChanged)
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
		it.Event = new(VaultextensionVaultBuffersPausedStateChanged)
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
func (it *VaultextensionVaultBuffersPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionVaultBuffersPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionVaultBuffersPausedStateChanged represents a VaultBuffersPausedStateChanged event raised by the Vaultextension contract.
type VaultextensionVaultBuffersPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultBuffersPausedStateChanged is a free log retrieval operation binding the contract event 0x300c7ca619eb846386aa0a6e5916ac2a41406448b0a2e99ba9ccafeb899015a5.
//
// Solidity: event VaultBuffersPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) FilterVaultBuffersPausedStateChanged(opts *bind.FilterOpts) (*VaultextensionVaultBuffersPausedStateChangedIterator, error) {

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "VaultBuffersPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &VaultextensionVaultBuffersPausedStateChangedIterator{contract: _Vaultextension.contract, event: "VaultBuffersPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchVaultBuffersPausedStateChanged is a free log subscription operation binding the contract event 0x300c7ca619eb846386aa0a6e5916ac2a41406448b0a2e99ba9ccafeb899015a5.
//
// Solidity: event VaultBuffersPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) WatchVaultBuffersPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionVaultBuffersPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "VaultBuffersPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionVaultBuffersPausedStateChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "VaultBuffersPausedStateChanged", log); err != nil {
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

// ParseVaultBuffersPausedStateChanged is a log parse operation binding the contract event 0x300c7ca619eb846386aa0a6e5916ac2a41406448b0a2e99ba9ccafeb899015a5.
//
// Solidity: event VaultBuffersPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) ParseVaultBuffersPausedStateChanged(log types.Log) (*VaultextensionVaultBuffersPausedStateChanged, error) {
	event := new(VaultextensionVaultBuffersPausedStateChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "VaultBuffersPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionVaultPausedStateChangedIterator is returned from FilterVaultPausedStateChanged and is used to iterate over the raw logs and unpacked data for VaultPausedStateChanged events raised by the Vaultextension contract.
type VaultextensionVaultPausedStateChangedIterator struct {
	Event *VaultextensionVaultPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultextensionVaultPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionVaultPausedStateChanged)
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
		it.Event = new(VaultextensionVaultPausedStateChanged)
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
func (it *VaultextensionVaultPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionVaultPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionVaultPausedStateChanged represents a VaultPausedStateChanged event raised by the Vaultextension contract.
type VaultextensionVaultPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultPausedStateChanged is a free log retrieval operation binding the contract event 0xe0629fe656e45ad7fd63a24b899da368690024c07043b88e57aee5095b1d3d02.
//
// Solidity: event VaultPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) FilterVaultPausedStateChanged(opts *bind.FilterOpts) (*VaultextensionVaultPausedStateChangedIterator, error) {

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "VaultPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &VaultextensionVaultPausedStateChangedIterator{contract: _Vaultextension.contract, event: "VaultPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchVaultPausedStateChanged is a free log subscription operation binding the contract event 0xe0629fe656e45ad7fd63a24b899da368690024c07043b88e57aee5095b1d3d02.
//
// Solidity: event VaultPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) WatchVaultPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultextensionVaultPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "VaultPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionVaultPausedStateChanged)
				if err := _Vaultextension.contract.UnpackLog(event, "VaultPausedStateChanged", log); err != nil {
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

// ParseVaultPausedStateChanged is a log parse operation binding the contract event 0xe0629fe656e45ad7fd63a24b899da368690024c07043b88e57aee5095b1d3d02.
//
// Solidity: event VaultPausedStateChanged(bool paused)
func (_Vaultextension *VaultextensionFilterer) ParseVaultPausedStateChanged(log types.Log) (*VaultextensionVaultPausedStateChanged, error) {
	event := new(VaultextensionVaultPausedStateChanged)
	if err := _Vaultextension.contract.UnpackLog(event, "VaultPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionVaultQueriesDisabledIterator is returned from FilterVaultQueriesDisabled and is used to iterate over the raw logs and unpacked data for VaultQueriesDisabled events raised by the Vaultextension contract.
type VaultextensionVaultQueriesDisabledIterator struct {
	Event *VaultextensionVaultQueriesDisabled // Event containing the contract specifics and raw log

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
func (it *VaultextensionVaultQueriesDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionVaultQueriesDisabled)
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
		it.Event = new(VaultextensionVaultQueriesDisabled)
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
func (it *VaultextensionVaultQueriesDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionVaultQueriesDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionVaultQueriesDisabled represents a VaultQueriesDisabled event raised by the Vaultextension contract.
type VaultextensionVaultQueriesDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVaultQueriesDisabled is a free log retrieval operation binding the contract event 0xbd204090fd387f08e3076528bf09b4fc99d8100d749eace96c06002d3fedc625.
//
// Solidity: event VaultQueriesDisabled()
func (_Vaultextension *VaultextensionFilterer) FilterVaultQueriesDisabled(opts *bind.FilterOpts) (*VaultextensionVaultQueriesDisabledIterator, error) {

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "VaultQueriesDisabled")
	if err != nil {
		return nil, err
	}
	return &VaultextensionVaultQueriesDisabledIterator{contract: _Vaultextension.contract, event: "VaultQueriesDisabled", logs: logs, sub: sub}, nil
}

// WatchVaultQueriesDisabled is a free log subscription operation binding the contract event 0xbd204090fd387f08e3076528bf09b4fc99d8100d749eace96c06002d3fedc625.
//
// Solidity: event VaultQueriesDisabled()
func (_Vaultextension *VaultextensionFilterer) WatchVaultQueriesDisabled(opts *bind.WatchOpts, sink chan<- *VaultextensionVaultQueriesDisabled) (event.Subscription, error) {

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "VaultQueriesDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionVaultQueriesDisabled)
				if err := _Vaultextension.contract.UnpackLog(event, "VaultQueriesDisabled", log); err != nil {
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

// ParseVaultQueriesDisabled is a log parse operation binding the contract event 0xbd204090fd387f08e3076528bf09b4fc99d8100d749eace96c06002d3fedc625.
//
// Solidity: event VaultQueriesDisabled()
func (_Vaultextension *VaultextensionFilterer) ParseVaultQueriesDisabled(log types.Log) (*VaultextensionVaultQueriesDisabled, error) {
	event := new(VaultextensionVaultQueriesDisabled)
	if err := _Vaultextension.contract.UnpackLog(event, "VaultQueriesDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionVaultQueriesEnabledIterator is returned from FilterVaultQueriesEnabled and is used to iterate over the raw logs and unpacked data for VaultQueriesEnabled events raised by the Vaultextension contract.
type VaultextensionVaultQueriesEnabledIterator struct {
	Event *VaultextensionVaultQueriesEnabled // Event containing the contract specifics and raw log

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
func (it *VaultextensionVaultQueriesEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionVaultQueriesEnabled)
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
		it.Event = new(VaultextensionVaultQueriesEnabled)
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
func (it *VaultextensionVaultQueriesEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionVaultQueriesEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionVaultQueriesEnabled represents a VaultQueriesEnabled event raised by the Vaultextension contract.
type VaultextensionVaultQueriesEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVaultQueriesEnabled is a free log retrieval operation binding the contract event 0x91d7478835f2b5adc315f5aad920f4a7f0a02f7fddf3042d17b2c80168ea17f5.
//
// Solidity: event VaultQueriesEnabled()
func (_Vaultextension *VaultextensionFilterer) FilterVaultQueriesEnabled(opts *bind.FilterOpts) (*VaultextensionVaultQueriesEnabledIterator, error) {

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "VaultQueriesEnabled")
	if err != nil {
		return nil, err
	}
	return &VaultextensionVaultQueriesEnabledIterator{contract: _Vaultextension.contract, event: "VaultQueriesEnabled", logs: logs, sub: sub}, nil
}

// WatchVaultQueriesEnabled is a free log subscription operation binding the contract event 0x91d7478835f2b5adc315f5aad920f4a7f0a02f7fddf3042d17b2c80168ea17f5.
//
// Solidity: event VaultQueriesEnabled()
func (_Vaultextension *VaultextensionFilterer) WatchVaultQueriesEnabled(opts *bind.WatchOpts, sink chan<- *VaultextensionVaultQueriesEnabled) (event.Subscription, error) {

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "VaultQueriesEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionVaultQueriesEnabled)
				if err := _Vaultextension.contract.UnpackLog(event, "VaultQueriesEnabled", log); err != nil {
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

// ParseVaultQueriesEnabled is a log parse operation binding the contract event 0x91d7478835f2b5adc315f5aad920f4a7f0a02f7fddf3042d17b2c80168ea17f5.
//
// Solidity: event VaultQueriesEnabled()
func (_Vaultextension *VaultextensionFilterer) ParseVaultQueriesEnabled(log types.Log) (*VaultextensionVaultQueriesEnabled, error) {
	event := new(VaultextensionVaultQueriesEnabled)
	if err := _Vaultextension.contract.UnpackLog(event, "VaultQueriesEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultextensionWrapIterator is returned from FilterWrap and is used to iterate over the raw logs and unpacked data for Wrap events raised by the Vaultextension contract.
type VaultextensionWrapIterator struct {
	Event *VaultextensionWrap // Event containing the contract specifics and raw log

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
func (it *VaultextensionWrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultextensionWrap)
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
		it.Event = new(VaultextensionWrap)
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
func (it *VaultextensionWrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultextensionWrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultextensionWrap represents a Wrap event raised by the Vaultextension contract.
type VaultextensionWrap struct {
	WrappedToken        common.Address
	DepositedUnderlying *big.Int
	MintedShares        *big.Int
	BufferBalances      [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterWrap is a free log retrieval operation binding the contract event 0x3771d13c67011e31e12031c54bb59b0bf544a80b81d280a3711e172aa8b7f47b.
//
// Solidity: event Wrap(address indexed wrappedToken, uint256 depositedUnderlying, uint256 mintedShares, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) FilterWrap(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultextensionWrapIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.FilterLogs(opts, "Wrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultextensionWrapIterator{contract: _Vaultextension.contract, event: "Wrap", logs: logs, sub: sub}, nil
}

// WatchWrap is a free log subscription operation binding the contract event 0x3771d13c67011e31e12031c54bb59b0bf544a80b81d280a3711e172aa8b7f47b.
//
// Solidity: event Wrap(address indexed wrappedToken, uint256 depositedUnderlying, uint256 mintedShares, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) WatchWrap(opts *bind.WatchOpts, sink chan<- *VaultextensionWrap, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vaultextension.contract.WatchLogs(opts, "Wrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultextensionWrap)
				if err := _Vaultextension.contract.UnpackLog(event, "Wrap", log); err != nil {
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

// ParseWrap is a log parse operation binding the contract event 0x3771d13c67011e31e12031c54bb59b0bf544a80b81d280a3711e172aa8b7f47b.
//
// Solidity: event Wrap(address indexed wrappedToken, uint256 depositedUnderlying, uint256 mintedShares, bytes32 bufferBalances)
func (_Vaultextension *VaultextensionFilterer) ParseWrap(log types.Log) (*VaultextensionWrap, error) {
	event := new(VaultextensionWrap)
	if err := _Vaultextension.contract.UnpackLog(event, "Wrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
