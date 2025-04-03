// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package poolmanager

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

// IPoolManagerModifyLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerModifyLiquidityParams struct {
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
}

// IPoolManagerSwapParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolManagerSwapParams struct {
	ZeroForOne        bool
	AmountSpecified   *big.Int
	SqrtPriceLimitX96 *big.Int
}

// PoolKey is an auto generated low-level Go binding around an user-defined struct.
type PoolKey struct {
	Currency0   common.Address
	Currency1   common.Address
	Fee         *big.Int
	TickSpacing *big.Int
	Hooks       common.Address
}

// PoolmanagerMetaData contains all meta data concerning the Poolmanager contract.
var PoolmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initialOwner\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AlreadyUnlocked\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"currency1\",\"type\":\"address\"}],\"name\":\"CurrenciesOutOfOrderOrEqual\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"CurrencyNotSettled\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DelegateCallNotAllowed\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidCaller\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ManagerLocked\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"MustClearExactPositiveDelta\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"NonzeroNativeValue\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"PoolNotInitialized\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"ProtocolFeeCurrencySynced\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"ProtocolFeeTooLarge\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SwapAmountCannotBeZero\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"TickSpacingTooLarge\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"}],\"name\":\"TickSpacingTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UnauthorizedDynamicLPFeeUpdate\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"}],\"name\":\"Donate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"name\":\"Initialize\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"int256\",\"name\":\"liquidityDelta\",\"type\":\"int256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"name\":\"ModifyLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"OperatorSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"protocolFeeController\",\"type\":\"address\"}],\"name\":\"ProtocolFeeControllerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"protocolFee\",\"type\":\"uint24\"}],\"name\":\"ProtocolFeeUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"PoolId\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"amount0\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"amount1\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"liquidity\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"},{\"indexed\":false,\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"caller\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"clear\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"collectProtocolFees\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountCollected\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint256\",\"name\":\"amount0\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount1\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"donate\",\"outputs\":[{\"internalType\":\"BalanceDelta\",\"name\":\"delta\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"extsload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"startSlot\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"nSlots\",\"type\":\"uint256\"}],\"name\":\"extsload\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"slots\",\"type\":\"bytes32[]\"}],\"name\":\"extsload\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"slots\",\"type\":\"bytes32[]\"}],\"name\":\"exttload\",\"outputs\":[{\"internalType\":\"bytes32[]\",\"name\":\"\",\"type\":\"bytes32[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"slot\",\"type\":\"bytes32\"}],\"name\":\"exttload\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceX96\",\"type\":\"uint160\"}],\"name\":\"initialize\",\"outputs\":[{\"internalType\":\"int24\",\"name\":\"tick\",\"type\":\"int24\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"isOperator\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"int24\",\"name\":\"tickLower\",\"type\":\"int24\"},{\"internalType\":\"int24\",\"name\":\"tickUpper\",\"type\":\"int24\"},{\"internalType\":\"int256\",\"name\":\"liquidityDelta\",\"type\":\"int256\"},{\"internalType\":\"bytes32\",\"name\":\"salt\",\"type\":\"bytes32\"}],\"internalType\":\"structIPoolManager.ModifyLiquidityParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"modifyLiquidity\",\"outputs\":[{\"internalType\":\"BalanceDelta\",\"name\":\"callerDelta\",\"type\":\"int256\"},{\"internalType\":\"BalanceDelta\",\"name\":\"feesAccrued\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"protocolFeeController\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"protocolFeesAccrued\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setOperator\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"newProtocolFee\",\"type\":\"uint24\"}],\"name\":\"setProtocolFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"controller\",\"type\":\"address\"}],\"name\":\"setProtocolFeeController\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"settle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"settleFor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"bool\",\"name\":\"zeroForOne\",\"type\":\"bool\"},{\"internalType\":\"int256\",\"name\":\"amountSpecified\",\"type\":\"int256\"},{\"internalType\":\"uint160\",\"name\":\"sqrtPriceLimitX96\",\"type\":\"uint160\"}],\"internalType\":\"structIPoolManager.SwapParams\",\"name\":\"params\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"hookData\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"BalanceDelta\",\"name\":\"swapDelta\",\"type\":\"int256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"}],\"name\":\"sync\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"Currency\",\"name\":\"currency\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"take\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"unlock\",\"outputs\":[{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"Currency\",\"name\":\"currency0\",\"type\":\"address\"},{\"internalType\":\"Currency\",\"name\":\"currency1\",\"type\":\"address\"},{\"internalType\":\"uint24\",\"name\":\"fee\",\"type\":\"uint24\"},{\"internalType\":\"int24\",\"name\":\"tickSpacing\",\"type\":\"int24\"},{\"internalType\":\"contractIHooks\",\"name\":\"hooks\",\"type\":\"address\"}],\"internalType\":\"structPoolKey\",\"name\":\"key\",\"type\":\"tuple\"},{\"internalType\":\"uint24\",\"name\":\"newDynamicLPFee\",\"type\":\"uint24\"}],\"name\":\"updateDynamicLPFee\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolmanagerMetaData.ABI instead.
var PoolmanagerABI = PoolmanagerMetaData.ABI

// Poolmanager is an auto generated Go binding around an Ethereum contract.
type Poolmanager struct {
	PoolmanagerCaller     // Read-only binding to the contract
	PoolmanagerTransactor // Write-only binding to the contract
	PoolmanagerFilterer   // Log filterer for contract events
}

// PoolmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolmanagerSession struct {
	Contract     *Poolmanager      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolmanagerCallerSession struct {
	Contract *PoolmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// PoolmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolmanagerTransactorSession struct {
	Contract     *PoolmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// PoolmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolmanagerRaw struct {
	Contract *Poolmanager // Generic contract binding to access the raw methods on
}

// PoolmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolmanagerCallerRaw struct {
	Contract *PoolmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// PoolmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolmanagerTransactorRaw struct {
	Contract *PoolmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPoolmanager creates a new instance of Poolmanager, bound to a specific deployed contract.
func NewPoolmanager(address common.Address, backend bind.ContractBackend) (*Poolmanager, error) {
	contract, err := bindPoolmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Poolmanager{PoolmanagerCaller: PoolmanagerCaller{contract: contract}, PoolmanagerTransactor: PoolmanagerTransactor{contract: contract}, PoolmanagerFilterer: PoolmanagerFilterer{contract: contract}}, nil
}

// NewPoolmanagerCaller creates a new read-only instance of Poolmanager, bound to a specific deployed contract.
func NewPoolmanagerCaller(address common.Address, caller bind.ContractCaller) (*PoolmanagerCaller, error) {
	contract, err := bindPoolmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerCaller{contract: contract}, nil
}

// NewPoolmanagerTransactor creates a new write-only instance of Poolmanager, bound to a specific deployed contract.
func NewPoolmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolmanagerTransactor, error) {
	contract, err := bindPoolmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerTransactor{contract: contract}, nil
}

// NewPoolmanagerFilterer creates a new log filterer instance of Poolmanager, bound to a specific deployed contract.
func NewPoolmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolmanagerFilterer, error) {
	contract, err := bindPoolmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerFilterer{contract: contract}, nil
}

// bindPoolmanager binds a generic wrapper to an already deployed contract.
func bindPoolmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolmanager *PoolmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolmanager.Contract.PoolmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolmanager *PoolmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolmanager.Contract.PoolmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolmanager *PoolmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolmanager.Contract.PoolmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Poolmanager *PoolmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Poolmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Poolmanager *PoolmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Poolmanager *PoolmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Poolmanager.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "allowance", owner, spender, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerSession) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _Poolmanager.Contract.Allowance(&_Poolmanager.CallOpts, owner, spender, id)
}

// Allowance is a free data retrieval call binding the contract method 0x598af9e7.
//
// Solidity: function allowance(address owner, address spender, uint256 id) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerCallerSession) Allowance(owner common.Address, spender common.Address, id *big.Int) (*big.Int, error) {
	return _Poolmanager.Contract.Allowance(&_Poolmanager.CallOpts, owner, spender, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 balance)
func (_Poolmanager *PoolmanagerCaller) BalanceOf(opts *bind.CallOpts, owner common.Address, id *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "balanceOf", owner, id)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 balance)
func (_Poolmanager *PoolmanagerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Poolmanager.Contract.BalanceOf(&_Poolmanager.CallOpts, owner, id)
}

// BalanceOf is a free data retrieval call binding the contract method 0x00fdd58e.
//
// Solidity: function balanceOf(address owner, uint256 id) view returns(uint256 balance)
func (_Poolmanager *PoolmanagerCallerSession) BalanceOf(owner common.Address, id *big.Int) (*big.Int, error) {
	return _Poolmanager.Contract.BalanceOf(&_Poolmanager.CallOpts, owner, id)
}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerCaller) Extsload(opts *bind.CallOpts, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "extsload", slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerSession) Extsload(slot [32]byte) ([32]byte, error) {
	return _Poolmanager.Contract.Extsload(&_Poolmanager.CallOpts, slot)
}

// Extsload is a free data retrieval call binding the contract method 0x1e2eaeaf.
//
// Solidity: function extsload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerCallerSession) Extsload(slot [32]byte) ([32]byte, error) {
	return _Poolmanager.Contract.Extsload(&_Poolmanager.CallOpts, slot)
}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCaller) Extsload0(opts *bind.CallOpts, startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "extsload0", startSlot, nSlots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerSession) Extsload0(startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	return _Poolmanager.Contract.Extsload0(&_Poolmanager.CallOpts, startSlot, nSlots)
}

// Extsload0 is a free data retrieval call binding the contract method 0x35fd631a.
//
// Solidity: function extsload(bytes32 startSlot, uint256 nSlots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCallerSession) Extsload0(startSlot [32]byte, nSlots *big.Int) ([][32]byte, error) {
	return _Poolmanager.Contract.Extsload0(&_Poolmanager.CallOpts, startSlot, nSlots)
}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCaller) Extsload1(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "extsload1", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerSession) Extsload1(slots [][32]byte) ([][32]byte, error) {
	return _Poolmanager.Contract.Extsload1(&_Poolmanager.CallOpts, slots)
}

// Extsload1 is a free data retrieval call binding the contract method 0xdbd035ff.
//
// Solidity: function extsload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCallerSession) Extsload1(slots [][32]byte) ([][32]byte, error) {
	return _Poolmanager.Contract.Extsload1(&_Poolmanager.CallOpts, slots)
}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCaller) Exttload(opts *bind.CallOpts, slots [][32]byte) ([][32]byte, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "exttload", slots)

	if err != nil {
		return *new([][32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([][32]byte)).(*[][32]byte)

	return out0, err

}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerSession) Exttload(slots [][32]byte) ([][32]byte, error) {
	return _Poolmanager.Contract.Exttload(&_Poolmanager.CallOpts, slots)
}

// Exttload is a free data retrieval call binding the contract method 0x9bf6645f.
//
// Solidity: function exttload(bytes32[] slots) view returns(bytes32[])
func (_Poolmanager *PoolmanagerCallerSession) Exttload(slots [][32]byte) ([][32]byte, error) {
	return _Poolmanager.Contract.Exttload(&_Poolmanager.CallOpts, slots)
}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerCaller) Exttload0(opts *bind.CallOpts, slot [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "exttload0", slot)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerSession) Exttload0(slot [32]byte) ([32]byte, error) {
	return _Poolmanager.Contract.Exttload0(&_Poolmanager.CallOpts, slot)
}

// Exttload0 is a free data retrieval call binding the contract method 0xf135baaa.
//
// Solidity: function exttload(bytes32 slot) view returns(bytes32)
func (_Poolmanager *PoolmanagerCallerSession) Exttload0(slot [32]byte) ([32]byte, error) {
	return _Poolmanager.Contract.Exttload0(&_Poolmanager.CallOpts, slot)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address operator) view returns(bool isOperator)
func (_Poolmanager *PoolmanagerCaller) IsOperator(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "isOperator", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address operator) view returns(bool isOperator)
func (_Poolmanager *PoolmanagerSession) IsOperator(owner common.Address, operator common.Address) (bool, error) {
	return _Poolmanager.Contract.IsOperator(&_Poolmanager.CallOpts, owner, operator)
}

// IsOperator is a free data retrieval call binding the contract method 0xb6363cf2.
//
// Solidity: function isOperator(address owner, address operator) view returns(bool isOperator)
func (_Poolmanager *PoolmanagerCallerSession) IsOperator(owner common.Address, operator common.Address) (bool, error) {
	return _Poolmanager.Contract.IsOperator(&_Poolmanager.CallOpts, owner, operator)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolmanager *PoolmanagerCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolmanager *PoolmanagerSession) Owner() (common.Address, error) {
	return _Poolmanager.Contract.Owner(&_Poolmanager.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Poolmanager *PoolmanagerCallerSession) Owner() (common.Address, error) {
	return _Poolmanager.Contract.Owner(&_Poolmanager.CallOpts)
}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_Poolmanager *PoolmanagerCaller) ProtocolFeeController(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "protocolFeeController")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_Poolmanager *PoolmanagerSession) ProtocolFeeController() (common.Address, error) {
	return _Poolmanager.Contract.ProtocolFeeController(&_Poolmanager.CallOpts)
}

// ProtocolFeeController is a free data retrieval call binding the contract method 0xf02de3b2.
//
// Solidity: function protocolFeeController() view returns(address)
func (_Poolmanager *PoolmanagerCallerSession) ProtocolFeeController() (common.Address, error) {
	return _Poolmanager.Contract.ProtocolFeeController(&_Poolmanager.CallOpts)
}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerCaller) ProtocolFeesAccrued(opts *bind.CallOpts, currency common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "protocolFeesAccrued", currency)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerSession) ProtocolFeesAccrued(currency common.Address) (*big.Int, error) {
	return _Poolmanager.Contract.ProtocolFeesAccrued(&_Poolmanager.CallOpts, currency)
}

// ProtocolFeesAccrued is a free data retrieval call binding the contract method 0x97e8cd4e.
//
// Solidity: function protocolFeesAccrued(address currency) view returns(uint256 amount)
func (_Poolmanager *PoolmanagerCallerSession) ProtocolFeesAccrued(currency common.Address) (*big.Int, error) {
	return _Poolmanager.Contract.ProtocolFeesAccrued(&_Poolmanager.CallOpts, currency)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Poolmanager *PoolmanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Poolmanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Poolmanager *PoolmanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Poolmanager.Contract.SupportsInterface(&_Poolmanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Poolmanager *PoolmanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Poolmanager.Contract.SupportsInterface(&_Poolmanager.CallOpts, interfaceId)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactor) Approve(opts *bind.TransactOpts, spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "approve", spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerSession) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Approve(&_Poolmanager.TransactOpts, spender, id, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x426a8493.
//
// Solidity: function approve(address spender, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactorSession) Approve(spender common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Approve(&_Poolmanager.TransactOpts, spender, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactor) Burn(opts *bind.TransactOpts, from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "burn", from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Burn(&_Poolmanager.TransactOpts, from, id, amount)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address from, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactorSession) Burn(from common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Burn(&_Poolmanager.TransactOpts, from, id, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactor) Clear(opts *bind.TransactOpts, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "clear", currency, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_Poolmanager *PoolmanagerSession) Clear(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Clear(&_Poolmanager.TransactOpts, currency, amount)
}

// Clear is a paid mutator transaction binding the contract method 0x80f0b44c.
//
// Solidity: function clear(address currency, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactorSession) Clear(currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Clear(&_Poolmanager.TransactOpts, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_Poolmanager *PoolmanagerTransactor) CollectProtocolFees(opts *bind.TransactOpts, recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "collectProtocolFees", recipient, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_Poolmanager *PoolmanagerSession) CollectProtocolFees(recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.CollectProtocolFees(&_Poolmanager.TransactOpts, recipient, currency, amount)
}

// CollectProtocolFees is a paid mutator transaction binding the contract method 0x8161b874.
//
// Solidity: function collectProtocolFees(address recipient, address currency, uint256 amount) returns(uint256 amountCollected)
func (_Poolmanager *PoolmanagerTransactorSession) CollectProtocolFees(recipient common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.CollectProtocolFees(&_Poolmanager.TransactOpts, recipient, currency, amount)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256 delta)
func (_Poolmanager *PoolmanagerTransactor) Donate(opts *bind.TransactOpts, key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "donate", key, amount0, amount1, hookData)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256 delta)
func (_Poolmanager *PoolmanagerSession) Donate(key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Donate(&_Poolmanager.TransactOpts, key, amount0, amount1, hookData)
}

// Donate is a paid mutator transaction binding the contract method 0x234266d7.
//
// Solidity: function donate((address,address,uint24,int24,address) key, uint256 amount0, uint256 amount1, bytes hookData) returns(int256 delta)
func (_Poolmanager *PoolmanagerTransactorSession) Donate(key PoolKey, amount0 *big.Int, amount1 *big.Int, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Donate(&_Poolmanager.TransactOpts, key, amount0, amount1, hookData)
}

// Initialize is a paid mutator transaction binding the contract method 0x6276cbbe.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96) returns(int24 tick)
func (_Poolmanager *PoolmanagerTransactor) Initialize(opts *bind.TransactOpts, key PoolKey, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "initialize", key, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0x6276cbbe.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96) returns(int24 tick)
func (_Poolmanager *PoolmanagerSession) Initialize(key PoolKey, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Initialize(&_Poolmanager.TransactOpts, key, sqrtPriceX96)
}

// Initialize is a paid mutator transaction binding the contract method 0x6276cbbe.
//
// Solidity: function initialize((address,address,uint24,int24,address) key, uint160 sqrtPriceX96) returns(int24 tick)
func (_Poolmanager *PoolmanagerTransactorSession) Initialize(key PoolKey, sqrtPriceX96 *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Initialize(&_Poolmanager.TransactOpts, key, sqrtPriceX96)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactor) Mint(opts *bind.TransactOpts, to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "mint", to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerSession) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Mint(&_Poolmanager.TransactOpts, to, id, amount)
}

// Mint is a paid mutator transaction binding the contract method 0x156e29f6.
//
// Solidity: function mint(address to, uint256 id, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactorSession) Mint(to common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Mint(&_Poolmanager.TransactOpts, to, id, amount)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_Poolmanager *PoolmanagerTransactor) ModifyLiquidity(opts *bind.TransactOpts, key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "modifyLiquidity", key, params, hookData)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_Poolmanager *PoolmanagerSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.ModifyLiquidity(&_Poolmanager.TransactOpts, key, params, hookData)
}

// ModifyLiquidity is a paid mutator transaction binding the contract method 0x5a6bcfda.
//
// Solidity: function modifyLiquidity((address,address,uint24,int24,address) key, (int24,int24,int256,bytes32) params, bytes hookData) returns(int256 callerDelta, int256 feesAccrued)
func (_Poolmanager *PoolmanagerTransactorSession) ModifyLiquidity(key PoolKey, params IPoolManagerModifyLiquidityParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.ModifyLiquidity(&_Poolmanager.TransactOpts, key, params, hookData)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_Poolmanager *PoolmanagerTransactor) SetOperator(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "setOperator", operator, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_Poolmanager *PoolmanagerSession) SetOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetOperator(&_Poolmanager.TransactOpts, operator, approved)
}

// SetOperator is a paid mutator transaction binding the contract method 0x558a7297.
//
// Solidity: function setOperator(address operator, bool approved) returns(bool)
func (_Poolmanager *PoolmanagerTransactorSession) SetOperator(operator common.Address, approved bool) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetOperator(&_Poolmanager.TransactOpts, operator, approved)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_Poolmanager *PoolmanagerTransactor) SetProtocolFee(opts *bind.TransactOpts, key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "setProtocolFee", key, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_Poolmanager *PoolmanagerSession) SetProtocolFee(key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetProtocolFee(&_Poolmanager.TransactOpts, key, newProtocolFee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x7e87ce7d.
//
// Solidity: function setProtocolFee((address,address,uint24,int24,address) key, uint24 newProtocolFee) returns()
func (_Poolmanager *PoolmanagerTransactorSession) SetProtocolFee(key PoolKey, newProtocolFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetProtocolFee(&_Poolmanager.TransactOpts, key, newProtocolFee)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_Poolmanager *PoolmanagerTransactor) SetProtocolFeeController(opts *bind.TransactOpts, controller common.Address) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "setProtocolFeeController", controller)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_Poolmanager *PoolmanagerSession) SetProtocolFeeController(controller common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetProtocolFeeController(&_Poolmanager.TransactOpts, controller)
}

// SetProtocolFeeController is a paid mutator transaction binding the contract method 0x2d771389.
//
// Solidity: function setProtocolFeeController(address controller) returns()
func (_Poolmanager *PoolmanagerTransactorSession) SetProtocolFeeController(controller common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.SetProtocolFeeController(&_Poolmanager.TransactOpts, controller)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256)
func (_Poolmanager *PoolmanagerTransactor) Settle(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "settle")
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256)
func (_Poolmanager *PoolmanagerSession) Settle() (*types.Transaction, error) {
	return _Poolmanager.Contract.Settle(&_Poolmanager.TransactOpts)
}

// Settle is a paid mutator transaction binding the contract method 0x11da60b4.
//
// Solidity: function settle() payable returns(uint256)
func (_Poolmanager *PoolmanagerTransactorSession) Settle() (*types.Transaction, error) {
	return _Poolmanager.Contract.Settle(&_Poolmanager.TransactOpts)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256)
func (_Poolmanager *PoolmanagerTransactor) SettleFor(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "settleFor", recipient)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256)
func (_Poolmanager *PoolmanagerSession) SettleFor(recipient common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.SettleFor(&_Poolmanager.TransactOpts, recipient)
}

// SettleFor is a paid mutator transaction binding the contract method 0x3dd45adb.
//
// Solidity: function settleFor(address recipient) payable returns(uint256)
func (_Poolmanager *PoolmanagerTransactorSession) SettleFor(recipient common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.SettleFor(&_Poolmanager.TransactOpts, recipient)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_Poolmanager *PoolmanagerTransactor) Swap(opts *bind.TransactOpts, key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "swap", key, params, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_Poolmanager *PoolmanagerSession) Swap(key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Swap(&_Poolmanager.TransactOpts, key, params, hookData)
}

// Swap is a paid mutator transaction binding the contract method 0xf3cd914c.
//
// Solidity: function swap((address,address,uint24,int24,address) key, (bool,int256,uint160) params, bytes hookData) returns(int256 swapDelta)
func (_Poolmanager *PoolmanagerTransactorSession) Swap(key PoolKey, params IPoolManagerSwapParams, hookData []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Swap(&_Poolmanager.TransactOpts, key, params, hookData)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_Poolmanager *PoolmanagerTransactor) Sync(opts *bind.TransactOpts, currency common.Address) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "sync", currency)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_Poolmanager *PoolmanagerSession) Sync(currency common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.Sync(&_Poolmanager.TransactOpts, currency)
}

// Sync is a paid mutator transaction binding the contract method 0xa5841194.
//
// Solidity: function sync(address currency) returns()
func (_Poolmanager *PoolmanagerTransactorSession) Sync(currency common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.Sync(&_Poolmanager.TransactOpts, currency)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactor) Take(opts *bind.TransactOpts, currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "take", currency, to, amount)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_Poolmanager *PoolmanagerSession) Take(currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Take(&_Poolmanager.TransactOpts, currency, to, amount)
}

// Take is a paid mutator transaction binding the contract method 0x0b0d9c09.
//
// Solidity: function take(address currency, address to, uint256 amount) returns()
func (_Poolmanager *PoolmanagerTransactorSession) Take(currency common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Take(&_Poolmanager.TransactOpts, currency, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactor) Transfer(opts *bind.TransactOpts, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "transfer", receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerSession) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Transfer(&_Poolmanager.TransactOpts, receiver, id, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0x095bcdb6.
//
// Solidity: function transfer(address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactorSession) Transfer(receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.Transfer(&_Poolmanager.TransactOpts, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "transferFrom", sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerSession) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.TransferFrom(&_Poolmanager.TransactOpts, sender, receiver, id, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0xfe99049a.
//
// Solidity: function transferFrom(address sender, address receiver, uint256 id, uint256 amount) returns(bool)
func (_Poolmanager *PoolmanagerTransactorSession) TransferFrom(sender common.Address, receiver common.Address, id *big.Int, amount *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.TransferFrom(&_Poolmanager.TransactOpts, sender, receiver, id, amount)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolmanager *PoolmanagerTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolmanager *PoolmanagerSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.TransferOwnership(&_Poolmanager.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Poolmanager *PoolmanagerTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Poolmanager.Contract.TransferOwnership(&_Poolmanager.TransactOpts, newOwner)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Poolmanager *PoolmanagerTransactor) Unlock(opts *bind.TransactOpts, data []byte) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "unlock", data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Poolmanager *PoolmanagerSession) Unlock(data []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Unlock(&_Poolmanager.TransactOpts, data)
}

// Unlock is a paid mutator transaction binding the contract method 0x48c89491.
//
// Solidity: function unlock(bytes data) returns(bytes result)
func (_Poolmanager *PoolmanagerTransactorSession) Unlock(data []byte) (*types.Transaction, error) {
	return _Poolmanager.Contract.Unlock(&_Poolmanager.TransactOpts, data)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_Poolmanager *PoolmanagerTransactor) UpdateDynamicLPFee(opts *bind.TransactOpts, key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.contract.Transact(opts, "updateDynamicLPFee", key, newDynamicLPFee)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_Poolmanager *PoolmanagerSession) UpdateDynamicLPFee(key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.UpdateDynamicLPFee(&_Poolmanager.TransactOpts, key, newDynamicLPFee)
}

// UpdateDynamicLPFee is a paid mutator transaction binding the contract method 0x52759651.
//
// Solidity: function updateDynamicLPFee((address,address,uint24,int24,address) key, uint24 newDynamicLPFee) returns()
func (_Poolmanager *PoolmanagerTransactorSession) UpdateDynamicLPFee(key PoolKey, newDynamicLPFee *big.Int) (*types.Transaction, error) {
	return _Poolmanager.Contract.UpdateDynamicLPFee(&_Poolmanager.TransactOpts, key, newDynamicLPFee)
}

// PoolmanagerApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Poolmanager contract.
type PoolmanagerApprovalIterator struct {
	Event *PoolmanagerApproval // Event containing the contract specifics and raw log

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
func (it *PoolmanagerApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerApproval)
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
		it.Event = new(PoolmanagerApproval)
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
func (it *PoolmanagerApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerApproval represents a Approval event raised by the Poolmanager contract.
type PoolmanagerApproval struct {
	Owner   common.Address
	Spender common.Address
	Id      *big.Int
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address, id []*big.Int) (*PoolmanagerApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerApprovalIterator{contract: _Poolmanager.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PoolmanagerApproval, owner []common.Address, spender []common.Address, id []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerApproval)
				if err := _Poolmanager.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0xb3fd5071835887567a0671151121894ddccc2842f1d10bedad13e0d17cace9a7.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) ParseApproval(log types.Log) (*PoolmanagerApproval, error) {
	event := new(PoolmanagerApproval)
	if err := _Poolmanager.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerDonateIterator is returned from FilterDonate and is used to iterate over the raw logs and unpacked data for Donate events raised by the Poolmanager contract.
type PoolmanagerDonateIterator struct {
	Event *PoolmanagerDonate // Event containing the contract specifics and raw log

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
func (it *PoolmanagerDonateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerDonate)
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
		it.Event = new(PoolmanagerDonate)
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
func (it *PoolmanagerDonateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerDonateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerDonate represents a Donate event raised by the Poolmanager contract.
type PoolmanagerDonate struct {
	Id      [32]byte
	Sender  common.Address
	Amount0 *big.Int
	Amount1 *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterDonate is a free log retrieval operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_Poolmanager *PoolmanagerFilterer) FilterDonate(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolmanagerDonateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "Donate", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerDonateIterator{contract: _Poolmanager.contract, event: "Donate", logs: logs, sub: sub}, nil
}

// WatchDonate is a free log subscription operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_Poolmanager *PoolmanagerFilterer) WatchDonate(opts *bind.WatchOpts, sink chan<- *PoolmanagerDonate, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "Donate", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerDonate)
				if err := _Poolmanager.contract.UnpackLog(event, "Donate", log); err != nil {
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

// ParseDonate is a log parse operation binding the contract event 0x29ef05caaff9404b7cb6d1c0e9bbae9eaa7ab2541feba1a9c4248594c08156cb.
//
// Solidity: event Donate(bytes32 indexed id, address indexed sender, uint256 amount0, uint256 amount1)
func (_Poolmanager *PoolmanagerFilterer) ParseDonate(log types.Log) (*PoolmanagerDonate, error) {
	event := new(PoolmanagerDonate)
	if err := _Poolmanager.contract.UnpackLog(event, "Donate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerInitializeIterator is returned from FilterInitialize and is used to iterate over the raw logs and unpacked data for Initialize events raised by the Poolmanager contract.
type PoolmanagerInitializeIterator struct {
	Event *PoolmanagerInitialize // Event containing the contract specifics and raw log

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
func (it *PoolmanagerInitializeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerInitialize)
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
		it.Event = new(PoolmanagerInitialize)
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
func (it *PoolmanagerInitializeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerInitializeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerInitialize represents a Initialize event raised by the Poolmanager contract.
type PoolmanagerInitialize struct {
	Id           [32]byte
	Currency0    common.Address
	Currency1    common.Address
	Fee          *big.Int
	TickSpacing  *big.Int
	Hooks        common.Address
	SqrtPriceX96 *big.Int
	Tick         *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterInitialize is a free log retrieval operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_Poolmanager *PoolmanagerFilterer) FilterInitialize(opts *bind.FilterOpts, id [][32]byte, currency0 []common.Address, currency1 []common.Address) (*PoolmanagerInitializeIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var currency0Rule []interface{}
	for _, currency0Item := range currency0 {
		currency0Rule = append(currency0Rule, currency0Item)
	}
	var currency1Rule []interface{}
	for _, currency1Item := range currency1 {
		currency1Rule = append(currency1Rule, currency1Item)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "Initialize", idRule, currency0Rule, currency1Rule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerInitializeIterator{contract: _Poolmanager.contract, event: "Initialize", logs: logs, sub: sub}, nil
}

// WatchInitialize is a free log subscription operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_Poolmanager *PoolmanagerFilterer) WatchInitialize(opts *bind.WatchOpts, sink chan<- *PoolmanagerInitialize, id [][32]byte, currency0 []common.Address, currency1 []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var currency0Rule []interface{}
	for _, currency0Item := range currency0 {
		currency0Rule = append(currency0Rule, currency0Item)
	}
	var currency1Rule []interface{}
	for _, currency1Item := range currency1 {
		currency1Rule = append(currency1Rule, currency1Item)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "Initialize", idRule, currency0Rule, currency1Rule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerInitialize)
				if err := _Poolmanager.contract.UnpackLog(event, "Initialize", log); err != nil {
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

// ParseInitialize is a log parse operation binding the contract event 0xdd466e674ea557f56295e2d0218a125ea4b4f0f6f3307b95f85e6110838d6438.
//
// Solidity: event Initialize(bytes32 indexed id, address indexed currency0, address indexed currency1, uint24 fee, int24 tickSpacing, address hooks, uint160 sqrtPriceX96, int24 tick)
func (_Poolmanager *PoolmanagerFilterer) ParseInitialize(log types.Log) (*PoolmanagerInitialize, error) {
	event := new(PoolmanagerInitialize)
	if err := _Poolmanager.contract.UnpackLog(event, "Initialize", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerModifyLiquidityIterator is returned from FilterModifyLiquidity and is used to iterate over the raw logs and unpacked data for ModifyLiquidity events raised by the Poolmanager contract.
type PoolmanagerModifyLiquidityIterator struct {
	Event *PoolmanagerModifyLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolmanagerModifyLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerModifyLiquidity)
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
		it.Event = new(PoolmanagerModifyLiquidity)
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
func (it *PoolmanagerModifyLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerModifyLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerModifyLiquidity represents a ModifyLiquidity event raised by the Poolmanager contract.
type PoolmanagerModifyLiquidity struct {
	Id             [32]byte
	Sender         common.Address
	TickLower      *big.Int
	TickUpper      *big.Int
	LiquidityDelta *big.Int
	Salt           [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterModifyLiquidity is a free log retrieval operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_Poolmanager *PoolmanagerFilterer) FilterModifyLiquidity(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolmanagerModifyLiquidityIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "ModifyLiquidity", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerModifyLiquidityIterator{contract: _Poolmanager.contract, event: "ModifyLiquidity", logs: logs, sub: sub}, nil
}

// WatchModifyLiquidity is a free log subscription operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_Poolmanager *PoolmanagerFilterer) WatchModifyLiquidity(opts *bind.WatchOpts, sink chan<- *PoolmanagerModifyLiquidity, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "ModifyLiquidity", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerModifyLiquidity)
				if err := _Poolmanager.contract.UnpackLog(event, "ModifyLiquidity", log); err != nil {
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

// ParseModifyLiquidity is a log parse operation binding the contract event 0xf208f4912782fd25c7f114ca3723a2d5dd6f3bcc3ac8db5af63baa85f711d5ec.
//
// Solidity: event ModifyLiquidity(bytes32 indexed id, address indexed sender, int24 tickLower, int24 tickUpper, int256 liquidityDelta, bytes32 salt)
func (_Poolmanager *PoolmanagerFilterer) ParseModifyLiquidity(log types.Log) (*PoolmanagerModifyLiquidity, error) {
	event := new(PoolmanagerModifyLiquidity)
	if err := _Poolmanager.contract.UnpackLog(event, "ModifyLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerOperatorSetIterator is returned from FilterOperatorSet and is used to iterate over the raw logs and unpacked data for OperatorSet events raised by the Poolmanager contract.
type PoolmanagerOperatorSetIterator struct {
	Event *PoolmanagerOperatorSet // Event containing the contract specifics and raw log

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
func (it *PoolmanagerOperatorSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerOperatorSet)
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
		it.Event = new(PoolmanagerOperatorSet)
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
func (it *PoolmanagerOperatorSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerOperatorSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerOperatorSet represents a OperatorSet event raised by the Poolmanager contract.
type PoolmanagerOperatorSet struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOperatorSet is a free log retrieval operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_Poolmanager *PoolmanagerFilterer) FilterOperatorSet(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*PoolmanagerOperatorSetIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "OperatorSet", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerOperatorSetIterator{contract: _Poolmanager.contract, event: "OperatorSet", logs: logs, sub: sub}, nil
}

// WatchOperatorSet is a free log subscription operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_Poolmanager *PoolmanagerFilterer) WatchOperatorSet(opts *bind.WatchOpts, sink chan<- *PoolmanagerOperatorSet, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "OperatorSet", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerOperatorSet)
				if err := _Poolmanager.contract.UnpackLog(event, "OperatorSet", log); err != nil {
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

// ParseOperatorSet is a log parse operation binding the contract event 0xceb576d9f15e4e200fdb5096d64d5dfd667e16def20c1eefd14256d8e3faa267.
//
// Solidity: event OperatorSet(address indexed owner, address indexed operator, bool approved)
func (_Poolmanager *PoolmanagerFilterer) ParseOperatorSet(log types.Log) (*PoolmanagerOperatorSet, error) {
	event := new(PoolmanagerOperatorSet)
	if err := _Poolmanager.contract.UnpackLog(event, "OperatorSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Poolmanager contract.
type PoolmanagerOwnershipTransferredIterator struct {
	Event *PoolmanagerOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PoolmanagerOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerOwnershipTransferred)
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
		it.Event = new(PoolmanagerOwnershipTransferred)
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
func (it *PoolmanagerOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerOwnershipTransferred represents a OwnershipTransferred event raised by the Poolmanager contract.
type PoolmanagerOwnershipTransferred struct {
	User     common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Poolmanager *PoolmanagerFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, user []common.Address, newOwner []common.Address) (*PoolmanagerOwnershipTransferredIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerOwnershipTransferredIterator{contract: _Poolmanager.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Poolmanager *PoolmanagerFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PoolmanagerOwnershipTransferred, user []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "OwnershipTransferred", userRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerOwnershipTransferred)
				if err := _Poolmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed user, address indexed newOwner)
func (_Poolmanager *PoolmanagerFilterer) ParseOwnershipTransferred(log types.Log) (*PoolmanagerOwnershipTransferred, error) {
	event := new(PoolmanagerOwnershipTransferred)
	if err := _Poolmanager.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerProtocolFeeControllerUpdatedIterator is returned from FilterProtocolFeeControllerUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeControllerUpdated events raised by the Poolmanager contract.
type PoolmanagerProtocolFeeControllerUpdatedIterator struct {
	Event *PoolmanagerProtocolFeeControllerUpdated // Event containing the contract specifics and raw log

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
func (it *PoolmanagerProtocolFeeControllerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerProtocolFeeControllerUpdated)
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
		it.Event = new(PoolmanagerProtocolFeeControllerUpdated)
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
func (it *PoolmanagerProtocolFeeControllerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerProtocolFeeControllerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerProtocolFeeControllerUpdated represents a ProtocolFeeControllerUpdated event raised by the Poolmanager contract.
type PoolmanagerProtocolFeeControllerUpdated struct {
	ProtocolFeeController common.Address
	Raw                   types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeControllerUpdated is a free log retrieval operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_Poolmanager *PoolmanagerFilterer) FilterProtocolFeeControllerUpdated(opts *bind.FilterOpts, protocolFeeController []common.Address) (*PoolmanagerProtocolFeeControllerUpdatedIterator, error) {

	var protocolFeeControllerRule []interface{}
	for _, protocolFeeControllerItem := range protocolFeeController {
		protocolFeeControllerRule = append(protocolFeeControllerRule, protocolFeeControllerItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "ProtocolFeeControllerUpdated", protocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerProtocolFeeControllerUpdatedIterator{contract: _Poolmanager.contract, event: "ProtocolFeeControllerUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeControllerUpdated is a free log subscription operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_Poolmanager *PoolmanagerFilterer) WatchProtocolFeeControllerUpdated(opts *bind.WatchOpts, sink chan<- *PoolmanagerProtocolFeeControllerUpdated, protocolFeeController []common.Address) (event.Subscription, error) {

	var protocolFeeControllerRule []interface{}
	for _, protocolFeeControllerItem := range protocolFeeController {
		protocolFeeControllerRule = append(protocolFeeControllerRule, protocolFeeControllerItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "ProtocolFeeControllerUpdated", protocolFeeControllerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerProtocolFeeControllerUpdated)
				if err := _Poolmanager.contract.UnpackLog(event, "ProtocolFeeControllerUpdated", log); err != nil {
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

// ParseProtocolFeeControllerUpdated is a log parse operation binding the contract event 0xb4bd8ef53df690b9943d3318996006dbb82a25f54719d8c8035b516a2a5b8acc.
//
// Solidity: event ProtocolFeeControllerUpdated(address indexed protocolFeeController)
func (_Poolmanager *PoolmanagerFilterer) ParseProtocolFeeControllerUpdated(log types.Log) (*PoolmanagerProtocolFeeControllerUpdated, error) {
	event := new(PoolmanagerProtocolFeeControllerUpdated)
	if err := _Poolmanager.contract.UnpackLog(event, "ProtocolFeeControllerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerProtocolFeeUpdatedIterator is returned from FilterProtocolFeeUpdated and is used to iterate over the raw logs and unpacked data for ProtocolFeeUpdated events raised by the Poolmanager contract.
type PoolmanagerProtocolFeeUpdatedIterator struct {
	Event *PoolmanagerProtocolFeeUpdated // Event containing the contract specifics and raw log

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
func (it *PoolmanagerProtocolFeeUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerProtocolFeeUpdated)
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
		it.Event = new(PoolmanagerProtocolFeeUpdated)
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
func (it *PoolmanagerProtocolFeeUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerProtocolFeeUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerProtocolFeeUpdated represents a ProtocolFeeUpdated event raised by the Poolmanager contract.
type PoolmanagerProtocolFeeUpdated struct {
	Id          [32]byte
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeUpdated is a free log retrieval operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_Poolmanager *PoolmanagerFilterer) FilterProtocolFeeUpdated(opts *bind.FilterOpts, id [][32]byte) (*PoolmanagerProtocolFeeUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "ProtocolFeeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerProtocolFeeUpdatedIterator{contract: _Poolmanager.contract, event: "ProtocolFeeUpdated", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeUpdated is a free log subscription operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_Poolmanager *PoolmanagerFilterer) WatchProtocolFeeUpdated(opts *bind.WatchOpts, sink chan<- *PoolmanagerProtocolFeeUpdated, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "ProtocolFeeUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerProtocolFeeUpdated)
				if err := _Poolmanager.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
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

// ParseProtocolFeeUpdated is a log parse operation binding the contract event 0xe9c42593e71f84403b84352cd168d693e2c9fcd1fdbcc3feb21d92b43e6696f9.
//
// Solidity: event ProtocolFeeUpdated(bytes32 indexed id, uint24 protocolFee)
func (_Poolmanager *PoolmanagerFilterer) ParseProtocolFeeUpdated(log types.Log) (*PoolmanagerProtocolFeeUpdated, error) {
	event := new(PoolmanagerProtocolFeeUpdated)
	if err := _Poolmanager.contract.UnpackLog(event, "ProtocolFeeUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Poolmanager contract.
type PoolmanagerSwapIterator struct {
	Event *PoolmanagerSwap // Event containing the contract specifics and raw log

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
func (it *PoolmanagerSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerSwap)
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
		it.Event = new(PoolmanagerSwap)
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
func (it *PoolmanagerSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerSwap represents a Swap event raised by the Poolmanager contract.
type PoolmanagerSwap struct {
	Id           [32]byte
	Sender       common.Address
	Amount0      *big.Int
	Amount1      *big.Int
	SqrtPriceX96 *big.Int
	Liquidity    *big.Int
	Tick         *big.Int
	Fee          *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_Poolmanager *PoolmanagerFilterer) FilterSwap(opts *bind.FilterOpts, id [][32]byte, sender []common.Address) (*PoolmanagerSwapIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "Swap", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerSwapIterator{contract: _Poolmanager.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_Poolmanager *PoolmanagerFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolmanagerSwap, id [][32]byte, sender []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "Swap", idRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerSwap)
				if err := _Poolmanager.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x40e9cecb9f5f1f1c5b9c97dec2917b7ee92e57ba5563708daca94dd84ad7112f.
//
// Solidity: event Swap(bytes32 indexed id, address indexed sender, int128 amount0, int128 amount1, uint160 sqrtPriceX96, uint128 liquidity, int24 tick, uint24 fee)
func (_Poolmanager *PoolmanagerFilterer) ParseSwap(log types.Log) (*PoolmanagerSwap, error) {
	event := new(PoolmanagerSwap)
	if err := _Poolmanager.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolmanagerTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Poolmanager contract.
type PoolmanagerTransferIterator struct {
	Event *PoolmanagerTransfer // Event containing the contract specifics and raw log

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
func (it *PoolmanagerTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolmanagerTransfer)
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
		it.Event = new(PoolmanagerTransfer)
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
func (it *PoolmanagerTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolmanagerTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolmanagerTransfer represents a Transfer event raised by the Poolmanager contract.
type PoolmanagerTransfer struct {
	Caller common.Address
	From   common.Address
	To     common.Address
	Id     *big.Int
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, id []*big.Int) (*PoolmanagerTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.FilterLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return &PoolmanagerTransferIterator{contract: _Poolmanager.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PoolmanagerTransfer, from []common.Address, to []common.Address, id []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Poolmanager.contract.WatchLogs(opts, "Transfer", fromRule, toRule, idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolmanagerTransfer)
				if err := _Poolmanager.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0x1b3d7edb2e9c0b0e7c525b20aaaef0f5940d2ed71663c7d39266ecafac728859.
//
// Solidity: event Transfer(address caller, address indexed from, address indexed to, uint256 indexed id, uint256 amount)
func (_Poolmanager *PoolmanagerFilterer) ParseTransfer(log types.Log) (*PoolmanagerTransfer, error) {
	event := new(PoolmanagerTransfer)
	if err := _Poolmanager.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
