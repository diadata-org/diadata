// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package vault

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

// AddLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type AddLiquidityParams struct {
	Pool            common.Address
	To              common.Address
	MaxAmountsIn    []*big.Int
	MinBptAmountOut *big.Int
	Kind            uint8
	UserData        []byte
}

// BufferWrapOrUnwrapParams is an auto generated low-level Go binding around an user-defined struct.
type BufferWrapOrUnwrapParams struct {
	Kind           uint8
	Direction      uint8
	WrappedToken   common.Address
	AmountGivenRaw *big.Int
	LimitRaw       *big.Int
}

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

// PoolRoleAccounts is an auto generated low-level Go binding around an user-defined struct.
type PoolRoleAccounts struct {
	PauseManager   common.Address
	SwapFeeManager common.Address
	PoolCreator    common.Address
}

// RemoveLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type RemoveLiquidityParams struct {
	Pool           common.Address
	From           common.Address
	MaxBptAmountIn *big.Int
	MinAmountsOut  []*big.Int
	Kind           uint8
	UserData       []byte
}

// TokenConfig is an auto generated low-level Go binding around an user-defined struct.
type TokenConfig struct {
	Token         common.Address
	TokenType     uint8
	RateProvider  common.Address
	PaysYieldFees bool
}

// VaultSwapParams is an auto generated low-level Go binding around an user-defined struct.
type VaultSwapParams struct {
	Kind           uint8
	Pool           common.Address
	TokenIn        common.Address
	TokenOut       common.Address
	AmountGivenRaw *big.Int
	LimitRaw       *big.Int
	UserData       []byte
}

// VaultMetaData contains all meta data concerning the Vault contract.
var VaultMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIVaultExtension\",\"name\":\"vaultExtension\",\"type\":\"address\"},{\"internalType\":\"contractIAuthorizer\",\"name\":\"authorizer\",\"type\":\"address\"},{\"internalType\":\"contractIProtocolFeeController\",\"name\":\"protocolFeeController\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"}],\"name\":\"AddressEmptyCode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"AddressInsufficientBalance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterAddLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterInitializeHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterRemoveLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AfterSwapHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AllZeroInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AmountGivenZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"AmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"AmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BalanceOverflow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeAddLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeInitializeHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeRemoveLiquidityHookFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BeforeSwapHookFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"BptAmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"BptAmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"BufferAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"BufferNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BufferSharesInvalidOwner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"BufferSharesInvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"BufferTotalSupplyTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotReceiveEth\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CannotSwapSameToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportAddLiquidityCustom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportDonation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportRemoveLiquidityCustom\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DoesNotSupportUnbalancedLiquidity\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DynamicSwapFeeHookFailed\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"allowance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientAllowance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"needed\",\"type\":\"uint256\"}],\"name\":\"ERC20InsufficientBalance\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"approver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidApprover\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"ERC20InvalidReceiver\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSender\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"ERC20InvalidSpender\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FailedInnerCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"FeePrecisionTooHigh\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxAmountIn\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedAmountInAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minAmountOut\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedAmountOutBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"HookAdjustedSwapLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"poolHooksContract\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolFactory\",\"type\":\"address\"}],\"name\":\"HookRegistrationFailed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InputLengthMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAddLiquidityKind\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidRemoveLiquidityKind\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenConfiguration\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenDecimals\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidTokenType\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"InvalidUnderlyingToken\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"invariantRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxInvariantRatio\",\"type\":\"uint256\"}],\"name\":\"InvariantRatioAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"invariantRatio\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minInvariantRatio\",\"type\":\"uint256\"}],\"name\":\"InvariantRatioBelowMin\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"issuedShares\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minIssuedShares\",\"type\":\"uint256\"}],\"name\":\"IssuedSharesBelowMin\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MaxTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MinTokens\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MultipleNonZeroInputs\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotEnoughBufferShares\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedUnderlyingAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualUnderlyingAmount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughUnderlying\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"expectedWrappedAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"actualWrappedAmount\",\"type\":\"uint256\"}],\"name\":\"NotEnoughWrapped\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotStaticCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NotVaultDelegateCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PauseBufferPeriodDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PercentageAboveMax\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAlreadyInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolInRecoveryMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotInRecoveryMode\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotInitialized\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolPauseWindowExpired\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"}],\"name\":\"PoolTotalSupplyTooLow\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeesExceedTotalCollected\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueriesDisabled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QueriesDisabledPermanently\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"QuoteResultSpoofed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ReentrancyGuardReentrantCall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RouterNotTrusted\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int256\",\"name\":\"value\",\"type\":\"int256\"}],\"name\":\"SafeCastOverflowedIntToUint\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"SafeCastOverflowedUintToInt\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"SenderIsNotVault\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFeePercentageTooHigh\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapFeePercentageTooLow\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limit\",\"type\":\"uint256\"}],\"name\":\"SwapLimit\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenAlreadyRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"TokenNotRegistered\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"expectedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"actualToken\",\"type\":\"address\"}],\"name\":\"TokensMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TradeAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultBuffersArePaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultIsNotUnlocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultNotPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPauseWindowDurationTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPauseWindowExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"VaultPaused\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"}],\"name\":\"WrapAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongProtocolFeeControllerDeployment\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlyingToken\",\"type\":\"address\"}],\"name\":\"WrongUnderlyingToken\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongVaultAdminDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"WrongVaultExtensionDeployment\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ZeroDivision\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aggregateSwapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"AggregateSwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"aggregateYieldFeePercentage\",\"type\":\"uint256\"}],\"name\":\"AggregateYieldFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIAuthorizer\",\"name\":\"newAuthorizer\",\"type\":\"address\"}],\"name\":\"AuthorizerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedShares\",\"type\":\"uint256\"}],\"name\":\"BufferSharesBurned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"issuedShares\",\"type\":\"uint256\"}],\"name\":\"BufferSharesMinted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumAddLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsAddedRaw\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"swapFeeAmountsRaw\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"LiquidityAddedToBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityProvider\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"enumRemoveLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"totalSupply\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"amountsRemovedRaw\",\"type\":\"uint256[]\"},{\"indexed\":false,\"internalType\":\"uint256[]\",\"name\":\"swapFeeAmountsRaw\",\"type\":\"uint256[]\"}],\"name\":\"LiquidityRemoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountWrapped\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"LiquidityRemovedFromBuffer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"PoolInitialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"PoolPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"recoveryMode\",\"type\":\"bool\"}],\"name\":\"PoolRecoveryModeStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"factory\",\"type\":\"address\"},{\"components\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"enumTokenType\",\"name\":\"tokenType\",\"type\":\"uint8\"},{\"internalType\":\"contractIRateProvider\",\"name\":\"rateProvider\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"paysYieldFees\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structTokenConfig[]\",\"name\":\"tokenConfig\",\"type\":\"tuple[]\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"pauseWindowEndTime\",\"type\":\"uint32\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"pauseManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"swapFeeManager\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"poolCreator\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structPoolRoleAccounts\",\"name\":\"roleAccounts\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"enableHookAdjustedAmounts\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterInitialize\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallComputeDynamicSwapFee\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterSwap\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterAddLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallBeforeRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"shouldCallAfterRemoveLiquidity\",\"type\":\"bool\"},{\"internalType\":\"address\",\"name\":\"hooksContract\",\"type\":\"address\"}],\"indexed\":false,\"internalType\":\"structHooksConfig\",\"name\":\"hooksConfig\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"disableUnbalancedLiquidity\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableAddLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableRemoveLiquidityCustom\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"enableDonation\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structLiquidityManagement\",\"name\":\"liquidityManagement\",\"type\":\"tuple\"}],\"name\":\"PoolRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIProtocolFeeController\",\"name\":\"newProtocolFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeeAmount\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"swapFeePercentage\",\"type\":\"uint256\"}],\"name\":\"SwapFeePercentageChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"burnedShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"withdrawnUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"Unwrap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"eventKey\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"bytes\",\"name\":\"eventData\",\"type\":\"bytes\"}],\"name\":\"VaultAuxiliary\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"VaultBuffersPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"paused\",\"type\":\"bool\"}],\"name\":\"VaultPausedStateChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VaultQueriesDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"VaultQueriesEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"depositedUnderlying\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"mintedShares\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"bufferBalances\",\"type\":\"bytes32\"}],\"name\":\"Wrap\",\"type\":\"event\"},{\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256[]\",\"name\":\"maxAmountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"minBptAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"enumAddLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structAddLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"amountsIn\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256\",\"name\":\"bptAmountOut\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"enumWrappingDirection\",\"name\":\"direction\",\"type\":\"uint8\"},{\"internalType\":\"contractIERC4626\",\"name\":\"wrappedToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountGivenRaw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitRaw\",\"type\":\"uint256\"}],\"internalType\":\"structBufferWrapOrUnwrapParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"erc4626BufferWrapOrUnwrap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculatedRaw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountInRaw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOutRaw\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getPoolTokenCountAndIndexOfToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getVaultExtension\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"reentrancyGuardEntered\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"maxBptAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"minAmountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"enumRemoveLiquidityKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structRemoveLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"bptAmountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256[]\",\"name\":\"amountsOut\",\"type\":\"uint256[]\"},{\"internalType\":\"bytes\",\"name\":\"returnData\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"sendTo\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountHint\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"credit\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"enumSwapKind\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"contractIERC20\",\"name\":\"tokenOut\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amountGivenRaw\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"limitRaw\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"userData\",\"type\":\"bytes\"}],\"internalType\":\"structVaultSwapParams\",\"name\":\"vaultSwapParams\",\"type\":\"tuple\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCalculated\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"stateMutability\":\"payable\",\"type\":\"receive\"}]",
}

// VaultABI is the input ABI used to generate the binding from.
// Deprecated: Use VaultMetaData.ABI instead.
var VaultABI = VaultMetaData.ABI

// Vault is an auto generated Go binding around an Ethereum contract.
type Vault struct {
	VaultCaller     // Read-only binding to the contract
	VaultTransactor // Write-only binding to the contract
	VaultFilterer   // Log filterer for contract events
}

// VaultCaller is an auto generated read-only Go binding around an Ethereum contract.
type VaultCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultTransactor is an auto generated write-only Go binding around an Ethereum contract.
type VaultTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type VaultFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// VaultSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type VaultSession struct {
	Contract     *Vault            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type VaultCallerSession struct {
	Contract *VaultCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// VaultTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type VaultTransactorSession struct {
	Contract     *VaultTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// VaultRaw is an auto generated low-level Go binding around an Ethereum contract.
type VaultRaw struct {
	Contract *Vault // Generic contract binding to access the raw methods on
}

// VaultCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type VaultCallerRaw struct {
	Contract *VaultCaller // Generic read-only contract binding to access the raw methods on
}

// VaultTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type VaultTransactorRaw struct {
	Contract *VaultTransactor // Generic write-only contract binding to access the raw methods on
}

// NewVault creates a new instance of Vault, bound to a specific deployed contract.
func NewVault(address common.Address, backend bind.ContractBackend) (*Vault, error) {
	contract, err := bindVault(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Vault{VaultCaller: VaultCaller{contract: contract}, VaultTransactor: VaultTransactor{contract: contract}, VaultFilterer: VaultFilterer{contract: contract}}, nil
}

// NewVaultCaller creates a new read-only instance of Vault, bound to a specific deployed contract.
func NewVaultCaller(address common.Address, caller bind.ContractCaller) (*VaultCaller, error) {
	contract, err := bindVault(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &VaultCaller{contract: contract}, nil
}

// NewVaultTransactor creates a new write-only instance of Vault, bound to a specific deployed contract.
func NewVaultTransactor(address common.Address, transactor bind.ContractTransactor) (*VaultTransactor, error) {
	contract, err := bindVault(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &VaultTransactor{contract: contract}, nil
}

// NewVaultFilterer creates a new log filterer instance of Vault, bound to a specific deployed contract.
func NewVaultFilterer(address common.Address, filterer bind.ContractFilterer) (*VaultFilterer, error) {
	contract, err := bindVault(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &VaultFilterer{contract: contract}, nil
}

// bindVault binds a generic wrapper to an already deployed contract.
func bindVault(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := VaultMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.VaultCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.VaultTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Vault *VaultCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Vault.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Vault *VaultTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Vault *VaultTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Vault.Contract.contract.Transact(opts, method, params...)
}

// GetPoolTokenCountAndIndexOfToken is a free data retrieval call binding the contract method 0xc9c1661b.
//
// Solidity: function getPoolTokenCountAndIndexOfToken(address pool, address token) view returns(uint256, uint256)
func (_Vault *VaultCaller) GetPoolTokenCountAndIndexOfToken(opts *bind.CallOpts, pool common.Address, token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getPoolTokenCountAndIndexOfToken", pool, token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetPoolTokenCountAndIndexOfToken is a free data retrieval call binding the contract method 0xc9c1661b.
//
// Solidity: function getPoolTokenCountAndIndexOfToken(address pool, address token) view returns(uint256, uint256)
func (_Vault *VaultSession) GetPoolTokenCountAndIndexOfToken(pool common.Address, token common.Address) (*big.Int, *big.Int, error) {
	return _Vault.Contract.GetPoolTokenCountAndIndexOfToken(&_Vault.CallOpts, pool, token)
}

// GetPoolTokenCountAndIndexOfToken is a free data retrieval call binding the contract method 0xc9c1661b.
//
// Solidity: function getPoolTokenCountAndIndexOfToken(address pool, address token) view returns(uint256, uint256)
func (_Vault *VaultCallerSession) GetPoolTokenCountAndIndexOfToken(pool common.Address, token common.Address) (*big.Int, *big.Int, error) {
	return _Vault.Contract.GetPoolTokenCountAndIndexOfToken(&_Vault.CallOpts, pool, token)
}

// GetVaultExtension is a free data retrieval call binding the contract method 0xb9a8effa.
//
// Solidity: function getVaultExtension() view returns(address)
func (_Vault *VaultCaller) GetVaultExtension(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "getVaultExtension")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetVaultExtension is a free data retrieval call binding the contract method 0xb9a8effa.
//
// Solidity: function getVaultExtension() view returns(address)
func (_Vault *VaultSession) GetVaultExtension() (common.Address, error) {
	return _Vault.Contract.GetVaultExtension(&_Vault.CallOpts)
}

// GetVaultExtension is a free data retrieval call binding the contract method 0xb9a8effa.
//
// Solidity: function getVaultExtension() view returns(address)
func (_Vault *VaultCallerSession) GetVaultExtension() (common.Address, error) {
	return _Vault.Contract.GetVaultExtension(&_Vault.CallOpts)
}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vault *VaultCaller) ReentrancyGuardEntered(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Vault.contract.Call(opts, &out, "reentrancyGuardEntered")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vault *VaultSession) ReentrancyGuardEntered() (bool, error) {
	return _Vault.Contract.ReentrancyGuardEntered(&_Vault.CallOpts)
}

// ReentrancyGuardEntered is a free data retrieval call binding the contract method 0xd2c725e0.
//
// Solidity: function reentrancyGuardEntered() view returns(bool)
func (_Vault *VaultCallerSession) ReentrancyGuardEntered() (bool, error) {
	return _Vault.Contract.ReentrancyGuardEntered(&_Vault.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4af29ec4.
//
// Solidity: function addLiquidity((address,address,uint256[],uint256,uint8,bytes) params) returns(uint256[] amountsIn, uint256 bptAmountOut, bytes returnData)
func (_Vault *VaultTransactor) AddLiquidity(opts *bind.TransactOpts, params AddLiquidityParams) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "addLiquidity", params)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4af29ec4.
//
// Solidity: function addLiquidity((address,address,uint256[],uint256,uint8,bytes) params) returns(uint256[] amountsIn, uint256 bptAmountOut, bytes returnData)
func (_Vault *VaultSession) AddLiquidity(params AddLiquidityParams) (*types.Transaction, error) {
	return _Vault.Contract.AddLiquidity(&_Vault.TransactOpts, params)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x4af29ec4.
//
// Solidity: function addLiquidity((address,address,uint256[],uint256,uint8,bytes) params) returns(uint256[] amountsIn, uint256 bptAmountOut, bytes returnData)
func (_Vault *VaultTransactorSession) AddLiquidity(params AddLiquidityParams) (*types.Transaction, error) {
	return _Vault.Contract.AddLiquidity(&_Vault.TransactOpts, params)
}

// Erc4626BufferWrapOrUnwrap is a paid mutator transaction binding the contract method 0x43583be5.
//
// Solidity: function erc4626BufferWrapOrUnwrap((uint8,uint8,address,uint256,uint256) params) returns(uint256 amountCalculatedRaw, uint256 amountInRaw, uint256 amountOutRaw)
func (_Vault *VaultTransactor) Erc4626BufferWrapOrUnwrap(opts *bind.TransactOpts, params BufferWrapOrUnwrapParams) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "erc4626BufferWrapOrUnwrap", params)
}

// Erc4626BufferWrapOrUnwrap is a paid mutator transaction binding the contract method 0x43583be5.
//
// Solidity: function erc4626BufferWrapOrUnwrap((uint8,uint8,address,uint256,uint256) params) returns(uint256 amountCalculatedRaw, uint256 amountInRaw, uint256 amountOutRaw)
func (_Vault *VaultSession) Erc4626BufferWrapOrUnwrap(params BufferWrapOrUnwrapParams) (*types.Transaction, error) {
	return _Vault.Contract.Erc4626BufferWrapOrUnwrap(&_Vault.TransactOpts, params)
}

// Erc4626BufferWrapOrUnwrap is a paid mutator transaction binding the contract method 0x43583be5.
//
// Solidity: function erc4626BufferWrapOrUnwrap((uint8,uint8,address,uint256,uint256) params) returns(uint256 amountCalculatedRaw, uint256 amountInRaw, uint256 amountOutRaw)
func (_Vault *VaultTransactorSession) Erc4626BufferWrapOrUnwrap(params BufferWrapOrUnwrapParams) (*types.Transaction, error) {
	return _Vault.Contract.Erc4626BufferWrapOrUnwrap(&_Vault.TransactOpts, params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x21457897.
//
// Solidity: function removeLiquidity((address,address,uint256,uint256[],uint8,bytes) params) returns(uint256 bptAmountIn, uint256[] amountsOut, bytes returnData)
func (_Vault *VaultTransactor) RemoveLiquidity(opts *bind.TransactOpts, params RemoveLiquidityParams) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "removeLiquidity", params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x21457897.
//
// Solidity: function removeLiquidity((address,address,uint256,uint256[],uint8,bytes) params) returns(uint256 bptAmountIn, uint256[] amountsOut, bytes returnData)
func (_Vault *VaultSession) RemoveLiquidity(params RemoveLiquidityParams) (*types.Transaction, error) {
	return _Vault.Contract.RemoveLiquidity(&_Vault.TransactOpts, params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x21457897.
//
// Solidity: function removeLiquidity((address,address,uint256,uint256[],uint8,bytes) params) returns(uint256 bptAmountIn, uint256[] amountsOut, bytes returnData)
func (_Vault *VaultTransactorSession) RemoveLiquidity(params RemoveLiquidityParams) (*types.Transaction, error) {
	return _Vault.Contract.RemoveLiquidity(&_Vault.TransactOpts, params)
}

// SendTo is a paid mutator transaction binding the contract method 0xae639329.
//
// Solidity: function sendTo(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactor) SendTo(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "sendTo", token, to, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0xae639329.
//
// Solidity: function sendTo(address token, address to, uint256 amount) returns()
func (_Vault *VaultSession) SendTo(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SendTo(&_Vault.TransactOpts, token, to, amount)
}

// SendTo is a paid mutator transaction binding the contract method 0xae639329.
//
// Solidity: function sendTo(address token, address to, uint256 amount) returns()
func (_Vault *VaultTransactorSession) SendTo(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.SendTo(&_Vault.TransactOpts, token, to, amount)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address token, uint256 amountHint) returns(uint256 credit)
func (_Vault *VaultTransactor) Settle(opts *bind.TransactOpts, token common.Address, amountHint *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "settle", token, amountHint)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address token, uint256 amountHint) returns(uint256 credit)
func (_Vault *VaultSession) Settle(token common.Address, amountHint *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Settle(&_Vault.TransactOpts, token, amountHint)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(address token, uint256 amountHint) returns(uint256 credit)
func (_Vault *VaultTransactorSession) Settle(token common.Address, amountHint *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Settle(&_Vault.TransactOpts, token, amountHint)
}

// Swap is a paid mutator transaction binding the contract method 0x2bfb780c.
//
// Solidity: function swap((uint8,address,address,address,uint256,uint256,bytes) vaultSwapParams) returns(uint256 amountCalculated, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultTransactor) Swap(opts *bind.TransactOpts, vaultSwapParams VaultSwapParams) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "swap", vaultSwapParams)
}

// Swap is a paid mutator transaction binding the contract method 0x2bfb780c.
//
// Solidity: function swap((uint8,address,address,address,uint256,uint256,bytes) vaultSwapParams) returns(uint256 amountCalculated, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultSession) Swap(vaultSwapParams VaultSwapParams) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, vaultSwapParams)
}

// Swap is a paid mutator transaction binding the contract method 0x2bfb780c.
//
// Solidity: function swap((uint8,address,address,address,uint256,uint256,bytes) vaultSwapParams) returns(uint256 amountCalculated, uint256 amountIn, uint256 amountOut)
func (_Vault *VaultTransactorSession) Swap(vaultSwapParams VaultSwapParams) (*types.Transaction, error) {
	return _Vault.Contract.Swap(&_Vault.TransactOpts, vaultSwapParams)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address owner, address to, uint256 amount) returns(bool)
func (_Vault *VaultTransactor) Transfer(opts *bind.TransactOpts, owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transfer", owner, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address owner, address to, uint256 amount) returns(bool)
func (_Vault *VaultSession) Transfer(owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, owner, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xbeabacc8.
//
// Solidity: function transfer(address owner, address to, uint256 amount) returns(bool)
func (_Vault *VaultTransactorSession) Transfer(owner common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.Transfer(&_Vault.TransactOpts, owner, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address spender, address from, address to, uint256 amount) returns(bool)
func (_Vault *VaultTransactor) TransferFrom(opts *bind.TransactOpts, spender common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "transferFrom", spender, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address spender, address from, address to, uint256 amount) returns(bool)
func (_Vault *VaultSession) TransferFrom(spender common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferFrom(&_Vault.TransactOpts, spender, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x15dacbea.
//
// Solidity: function transferFrom(address spender, address from, address to, uint256 amount) returns(bool)
func (_Vault *VaultTransactorSession) TransferFrom(spender common.Address, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Vault.Contract.TransferFrom(&_Vault.TransactOpts, spender, from, to, amount)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Vault *VaultTransactor) Unlock(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Vault.contract.Transact(opts, "unlock", data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Vault *VaultSession) Unlock(data []byte) (*types.Transaction, error) {
	return _Vault.Contract.Unlock(&_Vault.TransactOpts, data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Vault *VaultTransactorSession) Unlock(data []byte) (*types.Transaction, error) {
	return _Vault.Contract.Unlock(&_Vault.TransactOpts, data)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vault *VaultTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vault *VaultSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vault.Contract.Fallback(&_Vault.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_Vault *VaultTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _Vault.Contract.Fallback(&_Vault.TransactOpts, calldata)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Vault.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Vault *VaultTransactorSession) Receive() (*types.Transaction, error) {
	return _Vault.Contract.Receive(&_Vault.TransactOpts)
}

// VaultAggregateSwapFeePercentageChangedIterator is returned from FilterAggregateSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for AggregateSwapFeePercentageChanged events raised by the Vault contract.
type VaultAggregateSwapFeePercentageChangedIterator struct {
	Event *VaultAggregateSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultAggregateSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAggregateSwapFeePercentageChanged)
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
		it.Event = new(VaultAggregateSwapFeePercentageChanged)
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
func (it *VaultAggregateSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAggregateSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAggregateSwapFeePercentageChanged represents a AggregateSwapFeePercentageChanged event raised by the Vault contract.
type VaultAggregateSwapFeePercentageChanged struct {
	Pool                       common.Address
	AggregateSwapFeePercentage *big.Int
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterAggregateSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0xe4d371097beea42453a37406e2aef4c04f3c548f84ac50e72578662c0dcd7354.
//
// Solidity: event AggregateSwapFeePercentageChanged(address indexed pool, uint256 aggregateSwapFeePercentage)
func (_Vault *VaultFilterer) FilterAggregateSwapFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultAggregateSwapFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AggregateSwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultAggregateSwapFeePercentageChangedIterator{contract: _Vault.contract, event: "AggregateSwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchAggregateSwapFeePercentageChanged is a free log subscription operation binding the contract event 0xe4d371097beea42453a37406e2aef4c04f3c548f84ac50e72578662c0dcd7354.
//
// Solidity: event AggregateSwapFeePercentageChanged(address indexed pool, uint256 aggregateSwapFeePercentage)
func (_Vault *VaultFilterer) WatchAggregateSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultAggregateSwapFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AggregateSwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAggregateSwapFeePercentageChanged)
				if err := _Vault.contract.UnpackLog(event, "AggregateSwapFeePercentageChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseAggregateSwapFeePercentageChanged(log types.Log) (*VaultAggregateSwapFeePercentageChanged, error) {
	event := new(VaultAggregateSwapFeePercentageChanged)
	if err := _Vault.contract.UnpackLog(event, "AggregateSwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultAggregateYieldFeePercentageChangedIterator is returned from FilterAggregateYieldFeePercentageChanged and is used to iterate over the raw logs and unpacked data for AggregateYieldFeePercentageChanged events raised by the Vault contract.
type VaultAggregateYieldFeePercentageChangedIterator struct {
	Event *VaultAggregateYieldFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultAggregateYieldFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAggregateYieldFeePercentageChanged)
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
		it.Event = new(VaultAggregateYieldFeePercentageChanged)
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
func (it *VaultAggregateYieldFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAggregateYieldFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAggregateYieldFeePercentageChanged represents a AggregateYieldFeePercentageChanged event raised by the Vault contract.
type VaultAggregateYieldFeePercentageChanged struct {
	Pool                        common.Address
	AggregateYieldFeePercentage *big.Int
	Raw                         types.Log // Blockchain specific contextual infos
}

// FilterAggregateYieldFeePercentageChanged is a free log retrieval operation binding the contract event 0x606eb97d83164bd6b200d638cd49c14c65d94d4f2c674cfd85e24e0e202c3ca5.
//
// Solidity: event AggregateYieldFeePercentageChanged(address indexed pool, uint256 aggregateYieldFeePercentage)
func (_Vault *VaultFilterer) FilterAggregateYieldFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultAggregateYieldFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AggregateYieldFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultAggregateYieldFeePercentageChangedIterator{contract: _Vault.contract, event: "AggregateYieldFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchAggregateYieldFeePercentageChanged is a free log subscription operation binding the contract event 0x606eb97d83164bd6b200d638cd49c14c65d94d4f2c674cfd85e24e0e202c3ca5.
//
// Solidity: event AggregateYieldFeePercentageChanged(address indexed pool, uint256 aggregateYieldFeePercentage)
func (_Vault *VaultFilterer) WatchAggregateYieldFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultAggregateYieldFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AggregateYieldFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAggregateYieldFeePercentageChanged)
				if err := _Vault.contract.UnpackLog(event, "AggregateYieldFeePercentageChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseAggregateYieldFeePercentageChanged(log types.Log) (*VaultAggregateYieldFeePercentageChanged, error) {
	event := new(VaultAggregateYieldFeePercentageChanged)
	if err := _Vault.contract.UnpackLog(event, "AggregateYieldFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Vault contract.
type VaultApprovalIterator struct {
	Event *VaultApproval // Event containing the contract specifics and raw log

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
func (it *VaultApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultApproval)
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
		it.Event = new(VaultApproval)
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
func (it *VaultApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultApproval represents a Approval event raised by the Vault contract.
type VaultApproval struct {
	Pool    common.Address
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xa0175360a15bca328baf7ea85c7b784d58b222a50d0ce760b10dba336d226a61.
//
// Solidity: event Approval(address indexed pool, address indexed owner, address indexed spender, uint256 value)
func (_Vault *VaultFilterer) FilterApproval(opts *bind.FilterOpts, pool []common.Address, owner []common.Address, spender []common.Address) (*VaultApprovalIterator, error) {

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

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Approval", poolRule, ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &VaultApprovalIterator{contract: _Vault.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xa0175360a15bca328baf7ea85c7b784d58b222a50d0ce760b10dba336d226a61.
//
// Solidity: event Approval(address indexed pool, address indexed owner, address indexed spender, uint256 value)
func (_Vault *VaultFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *VaultApproval, pool []common.Address, owner []common.Address, spender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Approval", poolRule, ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultApproval)
				if err := _Vault.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_Vault *VaultFilterer) ParseApproval(log types.Log) (*VaultApproval, error) {
	event := new(VaultApproval)
	if err := _Vault.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultAuthorizerChangedIterator is returned from FilterAuthorizerChanged and is used to iterate over the raw logs and unpacked data for AuthorizerChanged events raised by the Vault contract.
type VaultAuthorizerChangedIterator struct {
	Event *VaultAuthorizerChanged // Event containing the contract specifics and raw log

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
func (it *VaultAuthorizerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultAuthorizerChanged)
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
		it.Event = new(VaultAuthorizerChanged)
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
func (it *VaultAuthorizerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultAuthorizerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultAuthorizerChanged represents a AuthorizerChanged event raised by the Vault contract.
type VaultAuthorizerChanged struct {
	NewAuthorizer common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuthorizerChanged is a free log retrieval operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vault *VaultFilterer) FilterAuthorizerChanged(opts *bind.FilterOpts, newAuthorizer []common.Address) (*VaultAuthorizerChangedIterator, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return &VaultAuthorizerChangedIterator{contract: _Vault.contract, event: "AuthorizerChanged", logs: logs, sub: sub}, nil
}

// WatchAuthorizerChanged is a free log subscription operation binding the contract event 0x94b979b6831a51293e2641426f97747feed46f17779fed9cd18d1ecefcfe92ef.
//
// Solidity: event AuthorizerChanged(address indexed newAuthorizer)
func (_Vault *VaultFilterer) WatchAuthorizerChanged(opts *bind.WatchOpts, sink chan<- *VaultAuthorizerChanged, newAuthorizer []common.Address) (event.Subscription, error) {

	var newAuthorizerRule []interface{}
	for _, newAuthorizerItem := range newAuthorizer {
		newAuthorizerRule = append(newAuthorizerRule, newAuthorizerItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "AuthorizerChanged", newAuthorizerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultAuthorizerChanged)
				if err := _Vault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseAuthorizerChanged(log types.Log) (*VaultAuthorizerChanged, error) {
	event := new(VaultAuthorizerChanged)
	if err := _Vault.contract.UnpackLog(event, "AuthorizerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultBufferSharesBurnedIterator is returned from FilterBufferSharesBurned and is used to iterate over the raw logs and unpacked data for BufferSharesBurned events raised by the Vault contract.
type VaultBufferSharesBurnedIterator struct {
	Event *VaultBufferSharesBurned // Event containing the contract specifics and raw log

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
func (it *VaultBufferSharesBurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultBufferSharesBurned)
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
		it.Event = new(VaultBufferSharesBurned)
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
func (it *VaultBufferSharesBurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultBufferSharesBurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultBufferSharesBurned represents a BufferSharesBurned event raised by the Vault contract.
type VaultBufferSharesBurned struct {
	WrappedToken common.Address
	From         common.Address
	BurnedShares *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferSharesBurned is a free log retrieval operation binding the contract event 0x4e09f7f7fc37ce2897800e2c2a9099565edb0a133d19d84a6871b3530af8846b.
//
// Solidity: event BufferSharesBurned(address indexed wrappedToken, address indexed from, uint256 burnedShares)
func (_Vault *VaultFilterer) FilterBufferSharesBurned(opts *bind.FilterOpts, wrappedToken []common.Address, from []common.Address) (*VaultBufferSharesBurnedIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "BufferSharesBurned", wrappedTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &VaultBufferSharesBurnedIterator{contract: _Vault.contract, event: "BufferSharesBurned", logs: logs, sub: sub}, nil
}

// WatchBufferSharesBurned is a free log subscription operation binding the contract event 0x4e09f7f7fc37ce2897800e2c2a9099565edb0a133d19d84a6871b3530af8846b.
//
// Solidity: event BufferSharesBurned(address indexed wrappedToken, address indexed from, uint256 burnedShares)
func (_Vault *VaultFilterer) WatchBufferSharesBurned(opts *bind.WatchOpts, sink chan<- *VaultBufferSharesBurned, wrappedToken []common.Address, from []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "BufferSharesBurned", wrappedTokenRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultBufferSharesBurned)
				if err := _Vault.contract.UnpackLog(event, "BufferSharesBurned", log); err != nil {
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
func (_Vault *VaultFilterer) ParseBufferSharesBurned(log types.Log) (*VaultBufferSharesBurned, error) {
	event := new(VaultBufferSharesBurned)
	if err := _Vault.contract.UnpackLog(event, "BufferSharesBurned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultBufferSharesMintedIterator is returned from FilterBufferSharesMinted and is used to iterate over the raw logs and unpacked data for BufferSharesMinted events raised by the Vault contract.
type VaultBufferSharesMintedIterator struct {
	Event *VaultBufferSharesMinted // Event containing the contract specifics and raw log

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
func (it *VaultBufferSharesMintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultBufferSharesMinted)
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
		it.Event = new(VaultBufferSharesMinted)
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
func (it *VaultBufferSharesMintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultBufferSharesMintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultBufferSharesMinted represents a BufferSharesMinted event raised by the Vault contract.
type VaultBufferSharesMinted struct {
	WrappedToken common.Address
	To           common.Address
	IssuedShares *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBufferSharesMinted is a free log retrieval operation binding the contract event 0xd66f031d33381c6408f0b32c884461e5de3df8808399b6f3a3d86b1368f8ec34.
//
// Solidity: event BufferSharesMinted(address indexed wrappedToken, address indexed to, uint256 issuedShares)
func (_Vault *VaultFilterer) FilterBufferSharesMinted(opts *bind.FilterOpts, wrappedToken []common.Address, to []common.Address) (*VaultBufferSharesMintedIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "BufferSharesMinted", wrappedTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultBufferSharesMintedIterator{contract: _Vault.contract, event: "BufferSharesMinted", logs: logs, sub: sub}, nil
}

// WatchBufferSharesMinted is a free log subscription operation binding the contract event 0xd66f031d33381c6408f0b32c884461e5de3df8808399b6f3a3d86b1368f8ec34.
//
// Solidity: event BufferSharesMinted(address indexed wrappedToken, address indexed to, uint256 issuedShares)
func (_Vault *VaultFilterer) WatchBufferSharesMinted(opts *bind.WatchOpts, sink chan<- *VaultBufferSharesMinted, wrappedToken []common.Address, to []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "BufferSharesMinted", wrappedTokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultBufferSharesMinted)
				if err := _Vault.contract.UnpackLog(event, "BufferSharesMinted", log); err != nil {
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
func (_Vault *VaultFilterer) ParseBufferSharesMinted(log types.Log) (*VaultBufferSharesMinted, error) {
	event := new(VaultBufferSharesMinted)
	if err := _Vault.contract.UnpackLog(event, "BufferSharesMinted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidityAddedIterator is returned from FilterLiquidityAdded and is used to iterate over the raw logs and unpacked data for LiquidityAdded events raised by the Vault contract.
type VaultLiquidityAddedIterator struct {
	Event *VaultLiquidityAdded // Event containing the contract specifics and raw log

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
func (it *VaultLiquidityAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidityAdded)
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
		it.Event = new(VaultLiquidityAdded)
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
func (it *VaultLiquidityAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidityAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidityAdded represents a LiquidityAdded event raised by the Vault contract.
type VaultLiquidityAdded struct {
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
func (_Vault *VaultFilterer) FilterLiquidityAdded(opts *bind.FilterOpts, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (*VaultLiquidityAddedIterator, error) {

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

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidityAdded", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return &VaultLiquidityAddedIterator{contract: _Vault.contract, event: "LiquidityAdded", logs: logs, sub: sub}, nil
}

// WatchLiquidityAdded is a free log subscription operation binding the contract event 0xa26a52d8d53702bba7f137907b8e1f99ff87f6d450144270ca25e72481cca871.
//
// Solidity: event LiquidityAdded(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsAddedRaw, uint256[] swapFeeAmountsRaw)
func (_Vault *VaultFilterer) WatchLiquidityAdded(opts *bind.WatchOpts, sink chan<- *VaultLiquidityAdded, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (event.Subscription, error) {

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

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidityAdded", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidityAdded)
				if err := _Vault.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
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
func (_Vault *VaultFilterer) ParseLiquidityAdded(log types.Log) (*VaultLiquidityAdded, error) {
	event := new(VaultLiquidityAdded)
	if err := _Vault.contract.UnpackLog(event, "LiquidityAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidityAddedToBufferIterator is returned from FilterLiquidityAddedToBuffer and is used to iterate over the raw logs and unpacked data for LiquidityAddedToBuffer events raised by the Vault contract.
type VaultLiquidityAddedToBufferIterator struct {
	Event *VaultLiquidityAddedToBuffer // Event containing the contract specifics and raw log

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
func (it *VaultLiquidityAddedToBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidityAddedToBuffer)
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
		it.Event = new(VaultLiquidityAddedToBuffer)
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
func (it *VaultLiquidityAddedToBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidityAddedToBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidityAddedToBuffer represents a LiquidityAddedToBuffer event raised by the Vault contract.
type VaultLiquidityAddedToBuffer struct {
	WrappedToken     common.Address
	AmountUnderlying *big.Int
	AmountWrapped    *big.Int
	BufferBalances   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityAddedToBuffer is a free log retrieval operation binding the contract event 0x75c4dc5f23640eeba7d404d9165f515fc3d9e23a5c8b6e2d09b4b9da56ff00a9.
//
// Solidity: event LiquidityAddedToBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vault *VaultFilterer) FilterLiquidityAddedToBuffer(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultLiquidityAddedToBufferIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidityAddedToBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultLiquidityAddedToBufferIterator{contract: _Vault.contract, event: "LiquidityAddedToBuffer", logs: logs, sub: sub}, nil
}

// WatchLiquidityAddedToBuffer is a free log subscription operation binding the contract event 0x75c4dc5f23640eeba7d404d9165f515fc3d9e23a5c8b6e2d09b4b9da56ff00a9.
//
// Solidity: event LiquidityAddedToBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vault *VaultFilterer) WatchLiquidityAddedToBuffer(opts *bind.WatchOpts, sink chan<- *VaultLiquidityAddedToBuffer, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidityAddedToBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidityAddedToBuffer)
				if err := _Vault.contract.UnpackLog(event, "LiquidityAddedToBuffer", log); err != nil {
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
func (_Vault *VaultFilterer) ParseLiquidityAddedToBuffer(log types.Log) (*VaultLiquidityAddedToBuffer, error) {
	event := new(VaultLiquidityAddedToBuffer)
	if err := _Vault.contract.UnpackLog(event, "LiquidityAddedToBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidityRemovedIterator is returned from FilterLiquidityRemoved and is used to iterate over the raw logs and unpacked data for LiquidityRemoved events raised by the Vault contract.
type VaultLiquidityRemovedIterator struct {
	Event *VaultLiquidityRemoved // Event containing the contract specifics and raw log

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
func (it *VaultLiquidityRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidityRemoved)
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
		it.Event = new(VaultLiquidityRemoved)
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
func (it *VaultLiquidityRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidityRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidityRemoved represents a LiquidityRemoved event raised by the Vault contract.
type VaultLiquidityRemoved struct {
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
func (_Vault *VaultFilterer) FilterLiquidityRemoved(opts *bind.FilterOpts, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (*VaultLiquidityRemovedIterator, error) {

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

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidityRemoved", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return &VaultLiquidityRemovedIterator{contract: _Vault.contract, event: "LiquidityRemoved", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemoved is a free log subscription operation binding the contract event 0xfbe5b0d79fb94f1e81c0a92bf86ae9d3a19e9d1bf6202c0d3e75120f65d5d8a5.
//
// Solidity: event LiquidityRemoved(address indexed pool, address indexed liquidityProvider, uint8 indexed kind, uint256 totalSupply, uint256[] amountsRemovedRaw, uint256[] swapFeeAmountsRaw)
func (_Vault *VaultFilterer) WatchLiquidityRemoved(opts *bind.WatchOpts, sink chan<- *VaultLiquidityRemoved, pool []common.Address, liquidityProvider []common.Address, kind []uint8) (event.Subscription, error) {

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

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidityRemoved", poolRule, liquidityProviderRule, kindRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidityRemoved)
				if err := _Vault.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
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
func (_Vault *VaultFilterer) ParseLiquidityRemoved(log types.Log) (*VaultLiquidityRemoved, error) {
	event := new(VaultLiquidityRemoved)
	if err := _Vault.contract.UnpackLog(event, "LiquidityRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultLiquidityRemovedFromBufferIterator is returned from FilterLiquidityRemovedFromBuffer and is used to iterate over the raw logs and unpacked data for LiquidityRemovedFromBuffer events raised by the Vault contract.
type VaultLiquidityRemovedFromBufferIterator struct {
	Event *VaultLiquidityRemovedFromBuffer // Event containing the contract specifics and raw log

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
func (it *VaultLiquidityRemovedFromBufferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultLiquidityRemovedFromBuffer)
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
		it.Event = new(VaultLiquidityRemovedFromBuffer)
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
func (it *VaultLiquidityRemovedFromBufferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultLiquidityRemovedFromBufferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultLiquidityRemovedFromBuffer represents a LiquidityRemovedFromBuffer event raised by the Vault contract.
type VaultLiquidityRemovedFromBuffer struct {
	WrappedToken     common.Address
	AmountUnderlying *big.Int
	AmountWrapped    *big.Int
	BufferBalances   [32]byte
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterLiquidityRemovedFromBuffer is a free log retrieval operation binding the contract event 0x44d97b36e99b590b3d2875aad3b167b1d7fb1e063f3f1325a1eeac76caee5113.
//
// Solidity: event LiquidityRemovedFromBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vault *VaultFilterer) FilterLiquidityRemovedFromBuffer(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultLiquidityRemovedFromBufferIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "LiquidityRemovedFromBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultLiquidityRemovedFromBufferIterator{contract: _Vault.contract, event: "LiquidityRemovedFromBuffer", logs: logs, sub: sub}, nil
}

// WatchLiquidityRemovedFromBuffer is a free log subscription operation binding the contract event 0x44d97b36e99b590b3d2875aad3b167b1d7fb1e063f3f1325a1eeac76caee5113.
//
// Solidity: event LiquidityRemovedFromBuffer(address indexed wrappedToken, uint256 amountUnderlying, uint256 amountWrapped, bytes32 bufferBalances)
func (_Vault *VaultFilterer) WatchLiquidityRemovedFromBuffer(opts *bind.WatchOpts, sink chan<- *VaultLiquidityRemovedFromBuffer, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "LiquidityRemovedFromBuffer", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultLiquidityRemovedFromBuffer)
				if err := _Vault.contract.UnpackLog(event, "LiquidityRemovedFromBuffer", log); err != nil {
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
func (_Vault *VaultFilterer) ParseLiquidityRemovedFromBuffer(log types.Log) (*VaultLiquidityRemovedFromBuffer, error) {
	event := new(VaultLiquidityRemovedFromBuffer)
	if err := _Vault.contract.UnpackLog(event, "LiquidityRemovedFromBuffer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolInitializedIterator is returned from FilterPoolInitialized and is used to iterate over the raw logs and unpacked data for PoolInitialized events raised by the Vault contract.
type VaultPoolInitializedIterator struct {
	Event *VaultPoolInitialized // Event containing the contract specifics and raw log

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
func (it *VaultPoolInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolInitialized)
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
		it.Event = new(VaultPoolInitialized)
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
func (it *VaultPoolInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolInitialized represents a PoolInitialized event raised by the Vault contract.
type VaultPoolInitialized struct {
	Pool common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterPoolInitialized is a free log retrieval operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed pool)
func (_Vault *VaultFilterer) FilterPoolInitialized(opts *bind.FilterOpts, pool []common.Address) (*VaultPoolInitializedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolInitialized", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolInitializedIterator{contract: _Vault.contract, event: "PoolInitialized", logs: logs, sub: sub}, nil
}

// WatchPoolInitialized is a free log subscription operation binding the contract event 0xcad8c9d32507393b6508ca4a888b81979919b477510585bde8488f153072d6f3.
//
// Solidity: event PoolInitialized(address indexed pool)
func (_Vault *VaultFilterer) WatchPoolInitialized(opts *bind.WatchOpts, sink chan<- *VaultPoolInitialized, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolInitialized", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolInitialized)
				if err := _Vault.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePoolInitialized(log types.Log) (*VaultPoolInitialized, error) {
	event := new(VaultPoolInitialized)
	if err := _Vault.contract.UnpackLog(event, "PoolInitialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolPausedStateChangedIterator is returned from FilterPoolPausedStateChanged and is used to iterate over the raw logs and unpacked data for PoolPausedStateChanged events raised by the Vault contract.
type VaultPoolPausedStateChangedIterator struct {
	Event *VaultPoolPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultPoolPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolPausedStateChanged)
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
		it.Event = new(VaultPoolPausedStateChanged)
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
func (it *VaultPoolPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolPausedStateChanged represents a PoolPausedStateChanged event raised by the Vault contract.
type VaultPoolPausedStateChanged struct {
	Pool   common.Address
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPoolPausedStateChanged is a free log retrieval operation binding the contract event 0x57e20448028297190122571be7cb6c1b1ef85730c673f7c72f533c8662419aa7.
//
// Solidity: event PoolPausedStateChanged(address indexed pool, bool paused)
func (_Vault *VaultFilterer) FilterPoolPausedStateChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultPoolPausedStateChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolPausedStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolPausedStateChangedIterator{contract: _Vault.contract, event: "PoolPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchPoolPausedStateChanged is a free log subscription operation binding the contract event 0x57e20448028297190122571be7cb6c1b1ef85730c673f7c72f533c8662419aa7.
//
// Solidity: event PoolPausedStateChanged(address indexed pool, bool paused)
func (_Vault *VaultFilterer) WatchPoolPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultPoolPausedStateChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolPausedStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolPausedStateChanged)
				if err := _Vault.contract.UnpackLog(event, "PoolPausedStateChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePoolPausedStateChanged(log types.Log) (*VaultPoolPausedStateChanged, error) {
	event := new(VaultPoolPausedStateChanged)
	if err := _Vault.contract.UnpackLog(event, "PoolPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolRecoveryModeStateChangedIterator is returned from FilterPoolRecoveryModeStateChanged and is used to iterate over the raw logs and unpacked data for PoolRecoveryModeStateChanged events raised by the Vault contract.
type VaultPoolRecoveryModeStateChangedIterator struct {
	Event *VaultPoolRecoveryModeStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultPoolRecoveryModeStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolRecoveryModeStateChanged)
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
		it.Event = new(VaultPoolRecoveryModeStateChanged)
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
func (it *VaultPoolRecoveryModeStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolRecoveryModeStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolRecoveryModeStateChanged represents a PoolRecoveryModeStateChanged event raised by the Vault contract.
type VaultPoolRecoveryModeStateChanged struct {
	Pool         common.Address
	RecoveryMode bool
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterPoolRecoveryModeStateChanged is a free log retrieval operation binding the contract event 0xc2354cc2f78ea57777e55ddd43a7f22b112ce98868596880edaeb22b4f9c73a9.
//
// Solidity: event PoolRecoveryModeStateChanged(address indexed pool, bool recoveryMode)
func (_Vault *VaultFilterer) FilterPoolRecoveryModeStateChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultPoolRecoveryModeStateChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolRecoveryModeStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolRecoveryModeStateChangedIterator{contract: _Vault.contract, event: "PoolRecoveryModeStateChanged", logs: logs, sub: sub}, nil
}

// WatchPoolRecoveryModeStateChanged is a free log subscription operation binding the contract event 0xc2354cc2f78ea57777e55ddd43a7f22b112ce98868596880edaeb22b4f9c73a9.
//
// Solidity: event PoolRecoveryModeStateChanged(address indexed pool, bool recoveryMode)
func (_Vault *VaultFilterer) WatchPoolRecoveryModeStateChanged(opts *bind.WatchOpts, sink chan<- *VaultPoolRecoveryModeStateChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolRecoveryModeStateChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolRecoveryModeStateChanged)
				if err := _Vault.contract.UnpackLog(event, "PoolRecoveryModeStateChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePoolRecoveryModeStateChanged(log types.Log) (*VaultPoolRecoveryModeStateChanged, error) {
	event := new(VaultPoolRecoveryModeStateChanged)
	if err := _Vault.contract.UnpackLog(event, "PoolRecoveryModeStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultPoolRegisteredIterator is returned from FilterPoolRegistered and is used to iterate over the raw logs and unpacked data for PoolRegistered events raised by the Vault contract.
type VaultPoolRegisteredIterator struct {
	Event *VaultPoolRegistered // Event containing the contract specifics and raw log

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
func (it *VaultPoolRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultPoolRegistered)
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
		it.Event = new(VaultPoolRegistered)
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
func (it *VaultPoolRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultPoolRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultPoolRegistered represents a PoolRegistered event raised by the Vault contract.
type VaultPoolRegistered struct {
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
func (_Vault *VaultFilterer) FilterPoolRegistered(opts *bind.FilterOpts, pool []common.Address, factory []common.Address) (*VaultPoolRegisteredIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "PoolRegistered", poolRule, factoryRule)
	if err != nil {
		return nil, err
	}
	return &VaultPoolRegisteredIterator{contract: _Vault.contract, event: "PoolRegistered", logs: logs, sub: sub}, nil
}

// WatchPoolRegistered is a free log subscription operation binding the contract event 0xbc1561eeab9f40962e2fb827a7ff9c7cdb47a9d7c84caeefa4ed90e043842dad.
//
// Solidity: event PoolRegistered(address indexed pool, address indexed factory, (address,uint8,address,bool)[] tokenConfig, uint256 swapFeePercentage, uint32 pauseWindowEndTime, (address,address,address) roleAccounts, (bool,bool,bool,bool,bool,bool,bool,bool,bool,bool,address) hooksConfig, (bool,bool,bool,bool) liquidityManagement)
func (_Vault *VaultFilterer) WatchPoolRegistered(opts *bind.WatchOpts, sink chan<- *VaultPoolRegistered, pool []common.Address, factory []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var factoryRule []interface{}
	for _, factoryItem := range factory {
		factoryRule = append(factoryRule, factoryItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "PoolRegistered", poolRule, factoryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultPoolRegistered)
				if err := _Vault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
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
func (_Vault *VaultFilterer) ParsePoolRegistered(log types.Log) (*VaultPoolRegistered, error) {
	event := new(VaultPoolRegistered)
	if err := _Vault.contract.UnpackLog(event, "PoolRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultProtocolFeeControllerChangedIterator is returned from FilterProtocolFeeControllerChanged and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerChanged events raised by the Vault contract.
type VaultProtocolFeeControllerChangedIterator struct {
	Event *VaultProtocolFeeControllerChanged // Event containing the contract specifics and raw log

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
func (it *VaultProtocolFeeControllerChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultProtocolFeeControllerChanged)
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
		it.Event = new(VaultProtocolFeeControllerChanged)
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
func (it *VaultProtocolFeeControllerChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultProtocolFeeControllerChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultProtocolFeeControllerChanged represents a ProtocolFeeControllerChanged event raised by the Vault contract.
type VaultProtocolFeeControllerChanged struct {
	NewProtocolFeeController common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerChanged is a free log retrieval operation binding the contract event 0x280a60b1e63c1774d397d35cce80eb80e51408ead755fb446e6f744ce98e5df0.
//
// Solidity: event ProtocolFeeControllerChanged(address indexed newProtocolFeeController)
func (_Vault *VaultFilterer) FilterProtocolFeeControllerChanged(opts *bind.FilterOpts, newProtocolFeeController []common.Address) (*VaultProtocolFeeControllerChangedIterator, error) {

	var newProtocolFeeControllerRule []interface{}
	for _, newProtocolFeeControllerItem := range newProtocolFeeController {
		newProtocolFeeControllerRule = append(newProtocolFeeControllerRule, newProtocolFeeControllerItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "ProtocolFeeControllerChanged", newProtocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return &VaultProtocolFeeControllerChangedIterator{contract: _Vault.contract, event: "ProtocolFeeControllerChanged", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerChanged is a free log subscription operation binding the contract event 0x280a60b1e63c1774d397d35cce80eb80e51408ead755fb446e6f744ce98e5df0.
//
// Solidity: event ProtocolFeeControllerChanged(address indexed newProtocolFeeController)
func (_Vault *VaultFilterer) WatchProtocolFeeControllerChanged(opts *bind.WatchOpts, sink chan<- *VaultProtocolFeeControllerChanged, newProtocolFeeController []common.Address) (event.Subscription, error) {

	var newProtocolFeeControllerRule []interface{}
	for _, newProtocolFeeControllerItem := range newProtocolFeeController {
		newProtocolFeeControllerRule = append(newProtocolFeeControllerRule, newProtocolFeeControllerItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "ProtocolFeeControllerChanged", newProtocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultProtocolFeeControllerChanged)
				if err := _Vault.contract.UnpackLog(event, "ProtocolFeeControllerChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseProtocolFeeControllerChanged(log types.Log) (*VaultProtocolFeeControllerChanged, error) {
	event := new(VaultProtocolFeeControllerChanged)
	if err := _Vault.contract.UnpackLog(event, "ProtocolFeeControllerChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Vault contract.
type VaultSwapIterator struct {
	Event *VaultSwap // Event containing the contract specifics and raw log

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
func (it *VaultSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSwap)
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
		it.Event = new(VaultSwap)
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
func (it *VaultSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSwap represents a Swap event raised by the Vault contract.
type VaultSwap struct {
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
func (_Vault *VaultFilterer) FilterSwap(opts *bind.FilterOpts, pool []common.Address, tokenIn []common.Address, tokenOut []common.Address) (*VaultSwapIterator, error) {

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

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Swap", poolRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return &VaultSwapIterator{contract: _Vault.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x0874b2d545cb271cdbda4e093020c452328b24af12382ed62c4d00f5c26709db.
//
// Solidity: event Swap(address indexed pool, address indexed tokenIn, address indexed tokenOut, uint256 amountIn, uint256 amountOut, uint256 swapFeePercentage, uint256 swapFeeAmount)
func (_Vault *VaultFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *VaultSwap, pool []common.Address, tokenIn []common.Address, tokenOut []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Swap", poolRule, tokenInRule, tokenOutRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSwap)
				if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
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
func (_Vault *VaultFilterer) ParseSwap(log types.Log) (*VaultSwap, error) {
	event := new(VaultSwap)
	if err := _Vault.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultSwapFeePercentageChangedIterator is returned from FilterSwapFeePercentageChanged and is used to iterate over the raw logs and unpacked data for SwapFeePercentageChanged events raised by the Vault contract.
type VaultSwapFeePercentageChangedIterator struct {
	Event *VaultSwapFeePercentageChanged // Event containing the contract specifics and raw log

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
func (it *VaultSwapFeePercentageChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultSwapFeePercentageChanged)
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
		it.Event = new(VaultSwapFeePercentageChanged)
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
func (it *VaultSwapFeePercentageChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultSwapFeePercentageChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultSwapFeePercentageChanged represents a SwapFeePercentageChanged event raised by the Vault contract.
type VaultSwapFeePercentageChanged struct {
	Pool              common.Address
	SwapFeePercentage *big.Int
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterSwapFeePercentageChanged is a free log retrieval operation binding the contract event 0x89d41522342fabac1471ca6073a5623e5caf367b03ca6e9a001478d0cf8be4a1.
//
// Solidity: event SwapFeePercentageChanged(address indexed pool, uint256 swapFeePercentage)
func (_Vault *VaultFilterer) FilterSwapFeePercentageChanged(opts *bind.FilterOpts, pool []common.Address) (*VaultSwapFeePercentageChangedIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "SwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return &VaultSwapFeePercentageChangedIterator{contract: _Vault.contract, event: "SwapFeePercentageChanged", logs: logs, sub: sub}, nil
}

// WatchSwapFeePercentageChanged is a free log subscription operation binding the contract event 0x89d41522342fabac1471ca6073a5623e5caf367b03ca6e9a001478d0cf8be4a1.
//
// Solidity: event SwapFeePercentageChanged(address indexed pool, uint256 swapFeePercentage)
func (_Vault *VaultFilterer) WatchSwapFeePercentageChanged(opts *bind.WatchOpts, sink chan<- *VaultSwapFeePercentageChanged, pool []common.Address) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "SwapFeePercentageChanged", poolRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultSwapFeePercentageChanged)
				if err := _Vault.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseSwapFeePercentageChanged(log types.Log) (*VaultSwapFeePercentageChanged, error) {
	event := new(VaultSwapFeePercentageChanged)
	if err := _Vault.contract.UnpackLog(event, "SwapFeePercentageChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Vault contract.
type VaultTransferIterator struct {
	Event *VaultTransfer // Event containing the contract specifics and raw log

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
func (it *VaultTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultTransfer)
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
		it.Event = new(VaultTransfer)
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
func (it *VaultTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultTransfer represents a Transfer event raised by the Vault contract.
type VaultTransfer struct {
	Pool  common.Address
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed pool, address indexed from, address indexed to, uint256 value)
func (_Vault *VaultFilterer) FilterTransfer(opts *bind.FilterOpts, pool []common.Address, from []common.Address, to []common.Address) (*VaultTransferIterator, error) {

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

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Transfer", poolRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &VaultTransferIterator{contract: _Vault.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xd1398bee19313d6bf672ccb116e51f4a1a947e91c757907f51fbb5b5e56c698f.
//
// Solidity: event Transfer(address indexed pool, address indexed from, address indexed to, uint256 value)
func (_Vault *VaultFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *VaultTransfer, pool []common.Address, from []common.Address, to []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Transfer", poolRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultTransfer)
				if err := _Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_Vault *VaultFilterer) ParseTransfer(log types.Log) (*VaultTransfer, error) {
	event := new(VaultTransfer)
	if err := _Vault.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultUnwrapIterator is returned from FilterUnwrap and is used to iterate over the raw logs and unpacked data for Unwrap events raised by the Vault contract.
type VaultUnwrapIterator struct {
	Event *VaultUnwrap // Event containing the contract specifics and raw log

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
func (it *VaultUnwrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultUnwrap)
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
		it.Event = new(VaultUnwrap)
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
func (it *VaultUnwrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultUnwrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultUnwrap represents a Unwrap event raised by the Vault contract.
type VaultUnwrap struct {
	WrappedToken        common.Address
	BurnedShares        *big.Int
	WithdrawnUnderlying *big.Int
	BufferBalances      [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterUnwrap is a free log retrieval operation binding the contract event 0xeeb740c90bf2b18c9532eb7d473137767036d893dff3e009f32718f821b2a4c0.
//
// Solidity: event Unwrap(address indexed wrappedToken, uint256 burnedShares, uint256 withdrawnUnderlying, bytes32 bufferBalances)
func (_Vault *VaultFilterer) FilterUnwrap(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultUnwrapIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Unwrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultUnwrapIterator{contract: _Vault.contract, event: "Unwrap", logs: logs, sub: sub}, nil
}

// WatchUnwrap is a free log subscription operation binding the contract event 0xeeb740c90bf2b18c9532eb7d473137767036d893dff3e009f32718f821b2a4c0.
//
// Solidity: event Unwrap(address indexed wrappedToken, uint256 burnedShares, uint256 withdrawnUnderlying, bytes32 bufferBalances)
func (_Vault *VaultFilterer) WatchUnwrap(opts *bind.WatchOpts, sink chan<- *VaultUnwrap, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Unwrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultUnwrap)
				if err := _Vault.contract.UnpackLog(event, "Unwrap", log); err != nil {
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
func (_Vault *VaultFilterer) ParseUnwrap(log types.Log) (*VaultUnwrap, error) {
	event := new(VaultUnwrap)
	if err := _Vault.contract.UnpackLog(event, "Unwrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultVaultAuxiliaryIterator is returned from FilterVaultAuxiliary and is used to iterate over the raw logs and unpacked data for VaultAuxiliary events raised by the Vault contract.
type VaultVaultAuxiliaryIterator struct {
	Event *VaultVaultAuxiliary // Event containing the contract specifics and raw log

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
func (it *VaultVaultAuxiliaryIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultVaultAuxiliary)
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
		it.Event = new(VaultVaultAuxiliary)
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
func (it *VaultVaultAuxiliaryIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultVaultAuxiliaryIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultVaultAuxiliary represents a VaultAuxiliary event raised by the Vault contract.
type VaultVaultAuxiliary struct {
	Pool      common.Address
	EventKey  [32]byte
	EventData []byte
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterVaultAuxiliary is a free log retrieval operation binding the contract event 0x4bc4412e210115456903c65b5277d299a505e79f2eb852b92b1ca52d85856428.
//
// Solidity: event VaultAuxiliary(address indexed pool, bytes32 indexed eventKey, bytes eventData)
func (_Vault *VaultFilterer) FilterVaultAuxiliary(opts *bind.FilterOpts, pool []common.Address, eventKey [][32]byte) (*VaultVaultAuxiliaryIterator, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var eventKeyRule []interface{}
	for _, eventKeyItem := range eventKey {
		eventKeyRule = append(eventKeyRule, eventKeyItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "VaultAuxiliary", poolRule, eventKeyRule)
	if err != nil {
		return nil, err
	}
	return &VaultVaultAuxiliaryIterator{contract: _Vault.contract, event: "VaultAuxiliary", logs: logs, sub: sub}, nil
}

// WatchVaultAuxiliary is a free log subscription operation binding the contract event 0x4bc4412e210115456903c65b5277d299a505e79f2eb852b92b1ca52d85856428.
//
// Solidity: event VaultAuxiliary(address indexed pool, bytes32 indexed eventKey, bytes eventData)
func (_Vault *VaultFilterer) WatchVaultAuxiliary(opts *bind.WatchOpts, sink chan<- *VaultVaultAuxiliary, pool []common.Address, eventKey [][32]byte) (event.Subscription, error) {

	var poolRule []interface{}
	for _, poolItem := range pool {
		poolRule = append(poolRule, poolItem)
	}
	var eventKeyRule []interface{}
	for _, eventKeyItem := range eventKey {
		eventKeyRule = append(eventKeyRule, eventKeyItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "VaultAuxiliary", poolRule, eventKeyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultVaultAuxiliary)
				if err := _Vault.contract.UnpackLog(event, "VaultAuxiliary", log); err != nil {
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
func (_Vault *VaultFilterer) ParseVaultAuxiliary(log types.Log) (*VaultVaultAuxiliary, error) {
	event := new(VaultVaultAuxiliary)
	if err := _Vault.contract.UnpackLog(event, "VaultAuxiliary", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultVaultBuffersPausedStateChangedIterator is returned from FilterVaultBuffersPausedStateChanged and is used to iterate over the raw logs and unpacked data for VaultBuffersPausedStateChanged events raised by the Vault contract.
type VaultVaultBuffersPausedStateChangedIterator struct {
	Event *VaultVaultBuffersPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultVaultBuffersPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultVaultBuffersPausedStateChanged)
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
		it.Event = new(VaultVaultBuffersPausedStateChanged)
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
func (it *VaultVaultBuffersPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultVaultBuffersPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultVaultBuffersPausedStateChanged represents a VaultBuffersPausedStateChanged event raised by the Vault contract.
type VaultVaultBuffersPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultBuffersPausedStateChanged is a free log retrieval operation binding the contract event 0x300c7ca619eb846386aa0a6e5916ac2a41406448b0a2e99ba9ccafeb899015a5.
//
// Solidity: event VaultBuffersPausedStateChanged(bool paused)
func (_Vault *VaultFilterer) FilterVaultBuffersPausedStateChanged(opts *bind.FilterOpts) (*VaultVaultBuffersPausedStateChangedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "VaultBuffersPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &VaultVaultBuffersPausedStateChangedIterator{contract: _Vault.contract, event: "VaultBuffersPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchVaultBuffersPausedStateChanged is a free log subscription operation binding the contract event 0x300c7ca619eb846386aa0a6e5916ac2a41406448b0a2e99ba9ccafeb899015a5.
//
// Solidity: event VaultBuffersPausedStateChanged(bool paused)
func (_Vault *VaultFilterer) WatchVaultBuffersPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultVaultBuffersPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "VaultBuffersPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultVaultBuffersPausedStateChanged)
				if err := _Vault.contract.UnpackLog(event, "VaultBuffersPausedStateChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseVaultBuffersPausedStateChanged(log types.Log) (*VaultVaultBuffersPausedStateChanged, error) {
	event := new(VaultVaultBuffersPausedStateChanged)
	if err := _Vault.contract.UnpackLog(event, "VaultBuffersPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultVaultPausedStateChangedIterator is returned from FilterVaultPausedStateChanged and is used to iterate over the raw logs and unpacked data for VaultPausedStateChanged events raised by the Vault contract.
type VaultVaultPausedStateChangedIterator struct {
	Event *VaultVaultPausedStateChanged // Event containing the contract specifics and raw log

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
func (it *VaultVaultPausedStateChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultVaultPausedStateChanged)
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
		it.Event = new(VaultVaultPausedStateChanged)
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
func (it *VaultVaultPausedStateChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultVaultPausedStateChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultVaultPausedStateChanged represents a VaultPausedStateChanged event raised by the Vault contract.
type VaultVaultPausedStateChanged struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterVaultPausedStateChanged is a free log retrieval operation binding the contract event 0xe0629fe656e45ad7fd63a24b899da368690024c07043b88e57aee5095b1d3d02.
//
// Solidity: event VaultPausedStateChanged(bool paused)
func (_Vault *VaultFilterer) FilterVaultPausedStateChanged(opts *bind.FilterOpts) (*VaultVaultPausedStateChangedIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "VaultPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return &VaultVaultPausedStateChangedIterator{contract: _Vault.contract, event: "VaultPausedStateChanged", logs: logs, sub: sub}, nil
}

// WatchVaultPausedStateChanged is a free log subscription operation binding the contract event 0xe0629fe656e45ad7fd63a24b899da368690024c07043b88e57aee5095b1d3d02.
//
// Solidity: event VaultPausedStateChanged(bool paused)
func (_Vault *VaultFilterer) WatchVaultPausedStateChanged(opts *bind.WatchOpts, sink chan<- *VaultVaultPausedStateChanged) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "VaultPausedStateChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultVaultPausedStateChanged)
				if err := _Vault.contract.UnpackLog(event, "VaultPausedStateChanged", log); err != nil {
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
func (_Vault *VaultFilterer) ParseVaultPausedStateChanged(log types.Log) (*VaultVaultPausedStateChanged, error) {
	event := new(VaultVaultPausedStateChanged)
	if err := _Vault.contract.UnpackLog(event, "VaultPausedStateChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultVaultQueriesDisabledIterator is returned from FilterVaultQueriesDisabled and is used to iterate over the raw logs and unpacked data for VaultQueriesDisabled events raised by the Vault contract.
type VaultVaultQueriesDisabledIterator struct {
	Event *VaultVaultQueriesDisabled // Event containing the contract specifics and raw log

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
func (it *VaultVaultQueriesDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultVaultQueriesDisabled)
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
		it.Event = new(VaultVaultQueriesDisabled)
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
func (it *VaultVaultQueriesDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultVaultQueriesDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultVaultQueriesDisabled represents a VaultQueriesDisabled event raised by the Vault contract.
type VaultVaultQueriesDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVaultQueriesDisabled is a free log retrieval operation binding the contract event 0xbd204090fd387f08e3076528bf09b4fc99d8100d749eace96c06002d3fedc625.
//
// Solidity: event VaultQueriesDisabled()
func (_Vault *VaultFilterer) FilterVaultQueriesDisabled(opts *bind.FilterOpts) (*VaultVaultQueriesDisabledIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "VaultQueriesDisabled")
	if err != nil {
		return nil, err
	}
	return &VaultVaultQueriesDisabledIterator{contract: _Vault.contract, event: "VaultQueriesDisabled", logs: logs, sub: sub}, nil
}

// WatchVaultQueriesDisabled is a free log subscription operation binding the contract event 0xbd204090fd387f08e3076528bf09b4fc99d8100d749eace96c06002d3fedc625.
//
// Solidity: event VaultQueriesDisabled()
func (_Vault *VaultFilterer) WatchVaultQueriesDisabled(opts *bind.WatchOpts, sink chan<- *VaultVaultQueriesDisabled) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "VaultQueriesDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultVaultQueriesDisabled)
				if err := _Vault.contract.UnpackLog(event, "VaultQueriesDisabled", log); err != nil {
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
func (_Vault *VaultFilterer) ParseVaultQueriesDisabled(log types.Log) (*VaultVaultQueriesDisabled, error) {
	event := new(VaultVaultQueriesDisabled)
	if err := _Vault.contract.UnpackLog(event, "VaultQueriesDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultVaultQueriesEnabledIterator is returned from FilterVaultQueriesEnabled and is used to iterate over the raw logs and unpacked data for VaultQueriesEnabled events raised by the Vault contract.
type VaultVaultQueriesEnabledIterator struct {
	Event *VaultVaultQueriesEnabled // Event containing the contract specifics and raw log

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
func (it *VaultVaultQueriesEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultVaultQueriesEnabled)
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
		it.Event = new(VaultVaultQueriesEnabled)
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
func (it *VaultVaultQueriesEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultVaultQueriesEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultVaultQueriesEnabled represents a VaultQueriesEnabled event raised by the Vault contract.
type VaultVaultQueriesEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterVaultQueriesEnabled is a free log retrieval operation binding the contract event 0x91d7478835f2b5adc315f5aad920f4a7f0a02f7fddf3042d17b2c80168ea17f5.
//
// Solidity: event VaultQueriesEnabled()
func (_Vault *VaultFilterer) FilterVaultQueriesEnabled(opts *bind.FilterOpts) (*VaultVaultQueriesEnabledIterator, error) {

	logs, sub, err := _Vault.contract.FilterLogs(opts, "VaultQueriesEnabled")
	if err != nil {
		return nil, err
	}
	return &VaultVaultQueriesEnabledIterator{contract: _Vault.contract, event: "VaultQueriesEnabled", logs: logs, sub: sub}, nil
}

// WatchVaultQueriesEnabled is a free log subscription operation binding the contract event 0x91d7478835f2b5adc315f5aad920f4a7f0a02f7fddf3042d17b2c80168ea17f5.
//
// Solidity: event VaultQueriesEnabled()
func (_Vault *VaultFilterer) WatchVaultQueriesEnabled(opts *bind.WatchOpts, sink chan<- *VaultVaultQueriesEnabled) (event.Subscription, error) {

	logs, sub, err := _Vault.contract.WatchLogs(opts, "VaultQueriesEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultVaultQueriesEnabled)
				if err := _Vault.contract.UnpackLog(event, "VaultQueriesEnabled", log); err != nil {
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
func (_Vault *VaultFilterer) ParseVaultQueriesEnabled(log types.Log) (*VaultVaultQueriesEnabled, error) {
	event := new(VaultVaultQueriesEnabled)
	if err := _Vault.contract.UnpackLog(event, "VaultQueriesEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// VaultWrapIterator is returned from FilterWrap and is used to iterate over the raw logs and unpacked data for Wrap events raised by the Vault contract.
type VaultWrapIterator struct {
	Event *VaultWrap // Event containing the contract specifics and raw log

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
func (it *VaultWrapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(VaultWrap)
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
		it.Event = new(VaultWrap)
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
func (it *VaultWrapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *VaultWrapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// VaultWrap represents a Wrap event raised by the Vault contract.
type VaultWrap struct {
	WrappedToken        common.Address
	DepositedUnderlying *big.Int
	MintedShares        *big.Int
	BufferBalances      [32]byte
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterWrap is a free log retrieval operation binding the contract event 0x3771d13c67011e31e12031c54bb59b0bf544a80b81d280a3711e172aa8b7f47b.
//
// Solidity: event Wrap(address indexed wrappedToken, uint256 depositedUnderlying, uint256 mintedShares, bytes32 bufferBalances)
func (_Vault *VaultFilterer) FilterWrap(opts *bind.FilterOpts, wrappedToken []common.Address) (*VaultWrapIterator, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.FilterLogs(opts, "Wrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return &VaultWrapIterator{contract: _Vault.contract, event: "Wrap", logs: logs, sub: sub}, nil
}

// WatchWrap is a free log subscription operation binding the contract event 0x3771d13c67011e31e12031c54bb59b0bf544a80b81d280a3711e172aa8b7f47b.
//
// Solidity: event Wrap(address indexed wrappedToken, uint256 depositedUnderlying, uint256 mintedShares, bytes32 bufferBalances)
func (_Vault *VaultFilterer) WatchWrap(opts *bind.WatchOpts, sink chan<- *VaultWrap, wrappedToken []common.Address) (event.Subscription, error) {

	var wrappedTokenRule []interface{}
	for _, wrappedTokenItem := range wrappedToken {
		wrappedTokenRule = append(wrappedTokenRule, wrappedTokenItem)
	}

	logs, sub, err := _Vault.contract.WatchLogs(opts, "Wrap", wrappedTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(VaultWrap)
				if err := _Vault.contract.UnpackLog(event, "Wrap", log); err != nil {
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
func (_Vault *VaultFilterer) ParseWrap(log types.Log) (*VaultWrap, error) {
	event := new(VaultWrap)
	if err := _Vault.contract.UnpackLog(event, "Wrap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
