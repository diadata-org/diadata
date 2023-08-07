// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package pool

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

// IPoolAddLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolAddLiquidityParams struct {
	Kind    uint8
	Pos     int32
	IsDelta bool
	DeltaA  *big.Int
	DeltaB  *big.Int
}

// IPoolBinDelta is an auto generated low-level Go binding around an user-defined struct.
type IPoolBinDelta struct {
	DeltaA         *big.Int
	DeltaB         *big.Int
	DeltaLpBalance *big.Int
	BinId          *big.Int
	Kind           uint8
	LowerTick      int32
	IsActive       bool
}

// IPoolBinState is an auto generated low-level Go binding around an user-defined struct.
type IPoolBinState struct {
	ReserveA        *big.Int
	ReserveB        *big.Int
	MergeBinBalance *big.Int
	MergeId         *big.Int
	TotalSupply     *big.Int
	Kind            uint8
	LowerTick       int32
}

// IPoolRemoveLiquidityParams is an auto generated low-level Go binding around an user-defined struct.
type IPoolRemoveLiquidityParams struct {
	BinId  *big.Int
	Amount *big.Int
}

// IPoolState is an auto generated low-level Go binding around an user-defined struct.
type IPoolState struct {
	ActiveTick       int32
	Status           uint8
	BinCounter       *big.Int
	ProtocolFeeRatio uint64
}

// IPoolTwaState is an auto generated low-level Go binding around an user-defined struct.
type IPoolTwaState struct {
	Twa           *big.Int
	Value         *big.Int
	LastTimestamp uint64
}

// PoolMetaData contains all meta data concerning the Pool contract.
var PoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"name\":\"AddLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"reserveA\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"reserveB\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"mergeId\",\"type\":\"uint128\"}],\"name\":\"BinMerged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"previousTick\",\"type\":\"int128\"},{\"indexed\":false,\"internalType\":\"int128\",\"name\":\"newTick\",\"type\":\"int128\"}],\"name\":\"BinMoved\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"maxRecursion\",\"type\":\"uint32\"}],\"name\":\"MigrateBinsUpStack\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isTokenA\",\"type\":\"bool\"}],\"name\":\"ProtocolFeeCollected\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"indexed\":false,\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"name\":\"RemoveLiquidity\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"protocolFee\",\"type\":\"uint256\"}],\"name\":\"SetProtocolFeeRatio\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"tokenAIn\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"exactOutput\",\"type\":\"bool\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"int32\",\"name\":\"activeTick\",\"type\":\"int32\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"indexed\":false,\"internalType\":\"structIPool.RemoveLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"TransferLiquidity\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"pos\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isDelta\",\"type\":\"bool\"},{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.AddLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"addLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBAmount\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"lpToken\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"binBalanceA\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"binBalanceB\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"tick\",\"type\":\"int32\"}],\"name\":\"binMap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"int32\",\"name\":\"tick\",\"type\":\"int32\"},{\"internalType\":\"uint256\",\"name\":\"kind\",\"type\":\"uint256\"}],\"name\":\"binPositions\",\"outputs\":[{\"internalType\":\"uint128\",\"name\":\"\",\"type\":\"uint128\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"factory\",\"outputs\":[{\"internalType\":\"contractIFactory\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"fee\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"}],\"name\":\"getBin\",\"outputs\":[{\"components\":[{\"internalType\":\"uint128\",\"name\":\"reserveA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"reserveB\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"mergeBinBalance\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"mergeId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"totalSupply\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"}],\"internalType\":\"structIPool.BinState\",\"name\":\"bin\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getCurrentTwa\",\"outputs\":[{\"internalType\":\"int256\",\"name\":\"\",\"type\":\"int256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getState\",\"outputs\":[{\"components\":[{\"internalType\":\"int32\",\"name\":\"activeTick\",\"type\":\"int32\"},{\"internalType\":\"uint8\",\"name\":\"status\",\"type\":\"uint8\"},{\"internalType\":\"uint128\",\"name\":\"binCounter\",\"type\":\"uint128\"},{\"internalType\":\"uint64\",\"name\":\"protocolFeeRatio\",\"type\":\"uint64\"}],\"internalType\":\"structIPool.State\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTwa\",\"outputs\":[{\"components\":[{\"internalType\":\"int96\",\"name\":\"twa\",\"type\":\"int96\"},{\"internalType\":\"int96\",\"name\":\"value\",\"type\":\"int96\"},{\"internalType\":\"uint64\",\"name\":\"lastTimestamp\",\"type\":\"uint64\"}],\"internalType\":\"structIPool.TwaState\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint32\",\"name\":\"maxRecursion\",\"type\":\"uint32\"}],\"name\":\"migrateBinUpStack\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.RemoveLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"removeLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenAOut\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tokenBOut\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"deltaA\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"deltaB\",\"type\":\"uint128\"},{\"internalType\":\"uint256\",\"name\":\"deltaLpBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint8\",\"name\":\"kind\",\"type\":\"uint8\"},{\"internalType\":\"int32\",\"name\":\"lowerTick\",\"type\":\"int32\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"}],\"internalType\":\"structIPool.BinDelta[]\",\"name\":\"binDeltas\",\"type\":\"tuple[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"tokenAIn\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"exactOutput\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"sqrtPriceLimit\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"data\",\"type\":\"bytes\"}],\"name\":\"swap\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"amountIn\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"amountOut\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tickSpacing\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenA\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenAScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenB\",\"outputs\":[{\"internalType\":\"contractIERC20\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"tokenBScale\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"fromTokenId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"toTokenId\",\"type\":\"uint256\"},{\"components\":[{\"internalType\":\"uint128\",\"name\":\"binId\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"amount\",\"type\":\"uint128\"}],\"internalType\":\"structIPool.RemoveLiquidityParams[]\",\"name\":\"params\",\"type\":\"tuple[]\"}],\"name\":\"transferLiquidity\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// PoolABI is the input ABI used to generate the binding from.
// Deprecated: Use PoolMetaData.ABI instead.
var PoolABI = PoolMetaData.ABI

// Pool is an auto generated Go binding around an Ethereum contract.
type Pool struct {
	PoolCaller     // Read-only binding to the contract
	PoolTransactor // Write-only binding to the contract
	PoolFilterer   // Log filterer for contract events
}

// PoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type PoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PoolSession struct {
	Contract     *Pool             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PoolCallerSession struct {
	Contract *PoolCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// PoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PoolTransactorSession struct {
	Contract     *PoolTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type PoolRaw struct {
	Contract *Pool // Generic contract binding to access the raw methods on
}

// PoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PoolCallerRaw struct {
	Contract *PoolCaller // Generic read-only contract binding to access the raw methods on
}

// PoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PoolTransactorRaw struct {
	Contract *PoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPool creates a new instance of Pool, bound to a specific deployed contract.
func NewPool(address common.Address, backend bind.ContractBackend) (*Pool, error) {
	contract, err := bindPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pool{PoolCaller: PoolCaller{contract: contract}, PoolTransactor: PoolTransactor{contract: contract}, PoolFilterer: PoolFilterer{contract: contract}}, nil
}

// NewPoolCaller creates a new read-only instance of Pool, bound to a specific deployed contract.
func NewPoolCaller(address common.Address, caller bind.ContractCaller) (*PoolCaller, error) {
	contract, err := bindPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PoolCaller{contract: contract}, nil
}

// NewPoolTransactor creates a new write-only instance of Pool, bound to a specific deployed contract.
func NewPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*PoolTransactor, error) {
	contract, err := bindPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PoolTransactor{contract: contract}, nil
}

// NewPoolFilterer creates a new log filterer instance of Pool, bound to a specific deployed contract.
func NewPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*PoolFilterer, error) {
	contract, err := bindPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PoolFilterer{contract: contract}, nil
}

// bindPool binds a generic wrapper to an already deployed contract.
func bindPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PoolMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.PoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.PoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pool *PoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pool *PoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pool *PoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pool.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x6da3bf8b.
//
// Solidity: function balanceOf(uint256 tokenId, uint128 binId) view returns(uint256 lpToken)
func (_Pool *PoolCaller) BalanceOf(opts *bind.CallOpts, tokenId *big.Int, binId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "balanceOf", tokenId, binId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x6da3bf8b.
//
// Solidity: function balanceOf(uint256 tokenId, uint128 binId) view returns(uint256 lpToken)
func (_Pool *PoolSession) BalanceOf(tokenId *big.Int, binId *big.Int) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, tokenId, binId)
}

// BalanceOf is a free data retrieval call binding the contract method 0x6da3bf8b.
//
// Solidity: function balanceOf(uint256 tokenId, uint128 binId) view returns(uint256 lpToken)
func (_Pool *PoolCallerSession) BalanceOf(tokenId *big.Int, binId *big.Int) (*big.Int, error) {
	return _Pool.Contract.BalanceOf(&_Pool.CallOpts, tokenId, binId)
}

// BinBalanceA is a free data retrieval call binding the contract method 0x75bbbd73.
//
// Solidity: function binBalanceA() view returns(uint128)
func (_Pool *PoolCaller) BinBalanceA(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "binBalanceA")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BinBalanceA is a free data retrieval call binding the contract method 0x75bbbd73.
//
// Solidity: function binBalanceA() view returns(uint128)
func (_Pool *PoolSession) BinBalanceA() (*big.Int, error) {
	return _Pool.Contract.BinBalanceA(&_Pool.CallOpts)
}

// BinBalanceA is a free data retrieval call binding the contract method 0x75bbbd73.
//
// Solidity: function binBalanceA() view returns(uint128)
func (_Pool *PoolCallerSession) BinBalanceA() (*big.Int, error) {
	return _Pool.Contract.BinBalanceA(&_Pool.CallOpts)
}

// BinBalanceB is a free data retrieval call binding the contract method 0xfa158509.
//
// Solidity: function binBalanceB() view returns(uint128)
func (_Pool *PoolCaller) BinBalanceB(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "binBalanceB")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BinBalanceB is a free data retrieval call binding the contract method 0xfa158509.
//
// Solidity: function binBalanceB() view returns(uint128)
func (_Pool *PoolSession) BinBalanceB() (*big.Int, error) {
	return _Pool.Contract.BinBalanceB(&_Pool.CallOpts)
}

// BinBalanceB is a free data retrieval call binding the contract method 0xfa158509.
//
// Solidity: function binBalanceB() view returns(uint128)
func (_Pool *PoolCallerSession) BinBalanceB() (*big.Int, error) {
	return _Pool.Contract.BinBalanceB(&_Pool.CallOpts)
}

// BinMap is a free data retrieval call binding the contract method 0xa2ba172f.
//
// Solidity: function binMap(int32 tick) view returns(uint256)
func (_Pool *PoolCaller) BinMap(opts *bind.CallOpts, tick int32) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "binMap", tick)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BinMap is a free data retrieval call binding the contract method 0xa2ba172f.
//
// Solidity: function binMap(int32 tick) view returns(uint256)
func (_Pool *PoolSession) BinMap(tick int32) (*big.Int, error) {
	return _Pool.Contract.BinMap(&_Pool.CallOpts, tick)
}

// BinMap is a free data retrieval call binding the contract method 0xa2ba172f.
//
// Solidity: function binMap(int32 tick) view returns(uint256)
func (_Pool *PoolCallerSession) BinMap(tick int32) (*big.Int, error) {
	return _Pool.Contract.BinMap(&_Pool.CallOpts, tick)
}

// BinPositions is a free data retrieval call binding the contract method 0x83f9c632.
//
// Solidity: function binPositions(int32 tick, uint256 kind) view returns(uint128)
func (_Pool *PoolCaller) BinPositions(opts *bind.CallOpts, tick int32, kind *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "binPositions", tick, kind)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BinPositions is a free data retrieval call binding the contract method 0x83f9c632.
//
// Solidity: function binPositions(int32 tick, uint256 kind) view returns(uint128)
func (_Pool *PoolSession) BinPositions(tick int32, kind *big.Int) (*big.Int, error) {
	return _Pool.Contract.BinPositions(&_Pool.CallOpts, tick, kind)
}

// BinPositions is a free data retrieval call binding the contract method 0x83f9c632.
//
// Solidity: function binPositions(int32 tick, uint256 kind) view returns(uint128)
func (_Pool *PoolCallerSession) BinPositions(tick int32, kind *big.Int) (*big.Int, error) {
	return _Pool.Contract.BinPositions(&_Pool.CallOpts, tick, kind)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCaller) Factory(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "factory")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// Factory is a free data retrieval call binding the contract method 0xc45a0155.
//
// Solidity: function factory() view returns(address)
func (_Pool *PoolCallerSession) Factory() (common.Address, error) {
	return _Pool.Contract.Factory(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pool *PoolCaller) Fee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "fee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pool *PoolSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// Fee is a free data retrieval call binding the contract method 0xddca3f43.
//
// Solidity: function fee() view returns(uint256)
func (_Pool *PoolCallerSession) Fee() (*big.Int, error) {
	return _Pool.Contract.Fee(&_Pool.CallOpts)
}

// GetBin is a free data retrieval call binding the contract method 0x44a185bb.
//
// Solidity: function getBin(uint128 binId) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32) bin)
func (_Pool *PoolCaller) GetBin(opts *bind.CallOpts, binId *big.Int) (IPoolBinState, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getBin", binId)

	if err != nil {
		return *new(IPoolBinState), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolBinState)).(*IPoolBinState)

	return out0, err

}

// GetBin is a free data retrieval call binding the contract method 0x44a185bb.
//
// Solidity: function getBin(uint128 binId) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32) bin)
func (_Pool *PoolSession) GetBin(binId *big.Int) (IPoolBinState, error) {
	return _Pool.Contract.GetBin(&_Pool.CallOpts, binId)
}

// GetBin is a free data retrieval call binding the contract method 0x44a185bb.
//
// Solidity: function getBin(uint128 binId) view returns((uint128,uint128,uint128,uint128,uint128,uint8,int32) bin)
func (_Pool *PoolCallerSession) GetBin(binId *big.Int) (IPoolBinState, error) {
	return _Pool.Contract.GetBin(&_Pool.CallOpts, binId)
}

// GetCurrentTwa is a free data retrieval call binding the contract method 0xd3d3861a.
//
// Solidity: function getCurrentTwa() view returns(int256)
func (_Pool *PoolCaller) GetCurrentTwa(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getCurrentTwa")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentTwa is a free data retrieval call binding the contract method 0xd3d3861a.
//
// Solidity: function getCurrentTwa() view returns(int256)
func (_Pool *PoolSession) GetCurrentTwa() (*big.Int, error) {
	return _Pool.Contract.GetCurrentTwa(&_Pool.CallOpts)
}

// GetCurrentTwa is a free data retrieval call binding the contract method 0xd3d3861a.
//
// Solidity: function getCurrentTwa() view returns(int256)
func (_Pool *PoolCallerSession) GetCurrentTwa() (*big.Int, error) {
	return _Pool.Contract.GetCurrentTwa(&_Pool.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((int32,uint8,uint128,uint64))
func (_Pool *PoolCaller) GetState(opts *bind.CallOpts) (IPoolState, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getState")

	if err != nil {
		return *new(IPoolState), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolState)).(*IPoolState)

	return out0, err

}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((int32,uint8,uint128,uint64))
func (_Pool *PoolSession) GetState() (IPoolState, error) {
	return _Pool.Contract.GetState(&_Pool.CallOpts)
}

// GetState is a free data retrieval call binding the contract method 0x1865c57d.
//
// Solidity: function getState() view returns((int32,uint8,uint128,uint64))
func (_Pool *PoolCallerSession) GetState() (IPoolState, error) {
	return _Pool.Contract.GetState(&_Pool.CallOpts)
}

// GetTwa is a free data retrieval call binding the contract method 0xa4ed496a.
//
// Solidity: function getTwa() view returns((int96,int96,uint64))
func (_Pool *PoolCaller) GetTwa(opts *bind.CallOpts) (IPoolTwaState, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "getTwa")

	if err != nil {
		return *new(IPoolTwaState), err
	}

	out0 := *abi.ConvertType(out[0], new(IPoolTwaState)).(*IPoolTwaState)

	return out0, err

}

// GetTwa is a free data retrieval call binding the contract method 0xa4ed496a.
//
// Solidity: function getTwa() view returns((int96,int96,uint64))
func (_Pool *PoolSession) GetTwa() (IPoolTwaState, error) {
	return _Pool.Contract.GetTwa(&_Pool.CallOpts)
}

// GetTwa is a free data retrieval call binding the contract method 0xa4ed496a.
//
// Solidity: function getTwa() view returns((int96,int96,uint64))
func (_Pool *PoolCallerSession) GetTwa() (IPoolTwaState, error) {
	return _Pool.Contract.GetTwa(&_Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(uint256)
func (_Pool *PoolCaller) TickSpacing(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tickSpacing")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(uint256)
func (_Pool *PoolSession) TickSpacing() (*big.Int, error) {
	return _Pool.Contract.TickSpacing(&_Pool.CallOpts)
}

// TickSpacing is a free data retrieval call binding the contract method 0xd0c93a7c.
//
// Solidity: function tickSpacing() view returns(uint256)
func (_Pool *PoolCallerSession) TickSpacing() (*big.Int, error) {
	return _Pool.Contract.TickSpacing(&_Pool.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Pool *PoolCaller) TokenA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tokenA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Pool *PoolSession) TokenA() (common.Address, error) {
	return _Pool.Contract.TokenA(&_Pool.CallOpts)
}

// TokenA is a free data retrieval call binding the contract method 0x0fc63d10.
//
// Solidity: function tokenA() view returns(address)
func (_Pool *PoolCallerSession) TokenA() (common.Address, error) {
	return _Pool.Contract.TokenA(&_Pool.CallOpts)
}

// TokenAScale is a free data retrieval call binding the contract method 0x3ab72c10.
//
// Solidity: function tokenAScale() view returns(uint256)
func (_Pool *PoolCaller) TokenAScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tokenAScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenAScale is a free data retrieval call binding the contract method 0x3ab72c10.
//
// Solidity: function tokenAScale() view returns(uint256)
func (_Pool *PoolSession) TokenAScale() (*big.Int, error) {
	return _Pool.Contract.TokenAScale(&_Pool.CallOpts)
}

// TokenAScale is a free data retrieval call binding the contract method 0x3ab72c10.
//
// Solidity: function tokenAScale() view returns(uint256)
func (_Pool *PoolCallerSession) TokenAScale() (*big.Int, error) {
	return _Pool.Contract.TokenAScale(&_Pool.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Pool *PoolCaller) TokenB(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tokenB")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Pool *PoolSession) TokenB() (common.Address, error) {
	return _Pool.Contract.TokenB(&_Pool.CallOpts)
}

// TokenB is a free data retrieval call binding the contract method 0x5f64b55b.
//
// Solidity: function tokenB() view returns(address)
func (_Pool *PoolCallerSession) TokenB() (common.Address, error) {
	return _Pool.Contract.TokenB(&_Pool.CallOpts)
}

// TokenBScale is a free data retrieval call binding the contract method 0x21272d4c.
//
// Solidity: function tokenBScale() view returns(uint256)
func (_Pool *PoolCaller) TokenBScale(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Pool.contract.Call(opts, &out, "tokenBScale")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenBScale is a free data retrieval call binding the contract method 0x21272d4c.
//
// Solidity: function tokenBScale() view returns(uint256)
func (_Pool *PoolSession) TokenBScale() (*big.Int, error) {
	return _Pool.Contract.TokenBScale(&_Pool.CallOpts)
}

// TokenBScale is a free data retrieval call binding the contract method 0x21272d4c.
//
// Solidity: function tokenBScale() view returns(uint256)
func (_Pool *PoolCallerSession) TokenBScale() (*big.Int, error) {
	return _Pool.Contract.TokenBScale(&_Pool.CallOpts)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9d5f20bb.
//
// Solidity: function addLiquidity(uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, bytes data) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolTransactor) AddLiquidity(opts *bind.TransactOpts, tokenId *big.Int, params []IPoolAddLiquidityParams, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "addLiquidity", tokenId, params, data)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9d5f20bb.
//
// Solidity: function addLiquidity(uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, bytes data) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolSession) AddLiquidity(tokenId *big.Int, params []IPoolAddLiquidityParams, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, tokenId, params, data)
}

// AddLiquidity is a paid mutator transaction binding the contract method 0x9d5f20bb.
//
// Solidity: function addLiquidity(uint256 tokenId, (uint8,int32,bool,uint128,uint128)[] params, bytes data) returns(uint256 tokenAAmount, uint256 tokenBAmount, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolTransactorSession) AddLiquidity(tokenId *big.Int, params []IPoolAddLiquidityParams, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.AddLiquidity(&_Pool.TransactOpts, tokenId, params, data)
}

// MigrateBinUpStack is a paid mutator transaction binding the contract method 0xc0c5d7fb.
//
// Solidity: function migrateBinUpStack(uint128 binId, uint32 maxRecursion) returns()
func (_Pool *PoolTransactor) MigrateBinUpStack(opts *bind.TransactOpts, binId *big.Int, maxRecursion uint32) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "migrateBinUpStack", binId, maxRecursion)
}

// MigrateBinUpStack is a paid mutator transaction binding the contract method 0xc0c5d7fb.
//
// Solidity: function migrateBinUpStack(uint128 binId, uint32 maxRecursion) returns()
func (_Pool *PoolSession) MigrateBinUpStack(binId *big.Int, maxRecursion uint32) (*types.Transaction, error) {
	return _Pool.Contract.MigrateBinUpStack(&_Pool.TransactOpts, binId, maxRecursion)
}

// MigrateBinUpStack is a paid mutator transaction binding the contract method 0xc0c5d7fb.
//
// Solidity: function migrateBinUpStack(uint128 binId, uint32 maxRecursion) returns()
func (_Pool *PoolTransactorSession) MigrateBinUpStack(binId *big.Int, maxRecursion uint32) (*types.Transaction, error) {
	return _Pool.Contract.MigrateBinUpStack(&_Pool.TransactOpts, binId, maxRecursion)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x57c8c7b0.
//
// Solidity: function removeLiquidity(address recipient, uint256 tokenId, (uint128,uint128)[] params) returns(uint256 tokenAOut, uint256 tokenBOut, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolTransactor) RemoveLiquidity(opts *bind.TransactOpts, recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "removeLiquidity", recipient, tokenId, params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x57c8c7b0.
//
// Solidity: function removeLiquidity(address recipient, uint256 tokenId, (uint128,uint128)[] params) returns(uint256 tokenAOut, uint256 tokenBOut, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolSession) RemoveLiquidity(recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, recipient, tokenId, params)
}

// RemoveLiquidity is a paid mutator transaction binding the contract method 0x57c8c7b0.
//
// Solidity: function removeLiquidity(address recipient, uint256 tokenId, (uint128,uint128)[] params) returns(uint256 tokenAOut, uint256 tokenBOut, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolTransactorSession) RemoveLiquidity(recipient common.Address, tokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.Contract.RemoveLiquidity(&_Pool.TransactOpts, recipient, tokenId, params)
}

// Swap is a paid mutator transaction binding the contract method 0xc51c9029.
//
// Solidity: function swap(address recipient, uint256 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit, bytes data) returns(uint256 amountIn, uint256 amountOut)
func (_Pool *PoolTransactor) Swap(opts *bind.TransactOpts, recipient common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "swap", recipient, amount, tokenAIn, exactOutput, sqrtPriceLimit, data)
}

// Swap is a paid mutator transaction binding the contract method 0xc51c9029.
//
// Solidity: function swap(address recipient, uint256 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit, bytes data) returns(uint256 amountIn, uint256 amountOut)
func (_Pool *PoolSession) Swap(recipient common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, recipient, amount, tokenAIn, exactOutput, sqrtPriceLimit, data)
}

// Swap is a paid mutator transaction binding the contract method 0xc51c9029.
//
// Solidity: function swap(address recipient, uint256 amount, bool tokenAIn, bool exactOutput, uint256 sqrtPriceLimit, bytes data) returns(uint256 amountIn, uint256 amountOut)
func (_Pool *PoolTransactorSession) Swap(recipient common.Address, amount *big.Int, tokenAIn bool, exactOutput bool, sqrtPriceLimit *big.Int, data []byte) (*types.Transaction, error) {
	return _Pool.Contract.Swap(&_Pool.TransactOpts, recipient, amount, tokenAIn, exactOutput, sqrtPriceLimit, data)
}

// TransferLiquidity is a paid mutator transaction binding the contract method 0xd279735f.
//
// Solidity: function transferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params) returns()
func (_Pool *PoolTransactor) TransferLiquidity(opts *bind.TransactOpts, fromTokenId *big.Int, toTokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.contract.Transact(opts, "transferLiquidity", fromTokenId, toTokenId, params)
}

// TransferLiquidity is a paid mutator transaction binding the contract method 0xd279735f.
//
// Solidity: function transferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params) returns()
func (_Pool *PoolSession) TransferLiquidity(fromTokenId *big.Int, toTokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.Contract.TransferLiquidity(&_Pool.TransactOpts, fromTokenId, toTokenId, params)
}

// TransferLiquidity is a paid mutator transaction binding the contract method 0xd279735f.
//
// Solidity: function transferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params) returns()
func (_Pool *PoolTransactorSession) TransferLiquidity(fromTokenId *big.Int, toTokenId *big.Int, params []IPoolRemoveLiquidityParams) (*types.Transaction, error) {
	return _Pool.Contract.TransferLiquidity(&_Pool.TransactOpts, fromTokenId, toTokenId, params)
}

// PoolAddLiquidityIterator is returned from FilterAddLiquidity and is used to iterate over the raw logs and unpacked data for AddLiquidity events raised by the Pool contract.
type PoolAddLiquidityIterator struct {
	Event *PoolAddLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolAddLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolAddLiquidity)
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
		it.Event = new(PoolAddLiquidity)
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
func (it *PoolAddLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolAddLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolAddLiquidity represents a AddLiquidity event raised by the Pool contract.
type PoolAddLiquidity struct {
	Sender    common.Address
	TokenId   *big.Int
	BinDeltas []IPoolBinDelta
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAddLiquidity is a free log retrieval operation binding the contract event 0x133a027327582be2089f6ca47137e3d337be4ca2cd921e5f0b178c9c2d5b8364.
//
// Solidity: event AddLiquidity(address indexed sender, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) FilterAddLiquidity(opts *bind.FilterOpts, sender []common.Address, tokenId []*big.Int) (*PoolAddLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "AddLiquidity", senderRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolAddLiquidityIterator{contract: _Pool.contract, event: "AddLiquidity", logs: logs, sub: sub}, nil
}

// WatchAddLiquidity is a free log subscription operation binding the contract event 0x133a027327582be2089f6ca47137e3d337be4ca2cd921e5f0b178c9c2d5b8364.
//
// Solidity: event AddLiquidity(address indexed sender, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) WatchAddLiquidity(opts *bind.WatchOpts, sink chan<- *PoolAddLiquidity, sender []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "AddLiquidity", senderRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolAddLiquidity)
				if err := _Pool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
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

// ParseAddLiquidity is a log parse operation binding the contract event 0x133a027327582be2089f6ca47137e3d337be4ca2cd921e5f0b178c9c2d5b8364.
//
// Solidity: event AddLiquidity(address indexed sender, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) ParseAddLiquidity(log types.Log) (*PoolAddLiquidity, error) {
	event := new(PoolAddLiquidity)
	if err := _Pool.contract.UnpackLog(event, "AddLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBinMergedIterator is returned from FilterBinMerged and is used to iterate over the raw logs and unpacked data for BinMerged events raised by the Pool contract.
type PoolBinMergedIterator struct {
	Event *PoolBinMerged // Event containing the contract specifics and raw log

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
func (it *PoolBinMergedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBinMerged)
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
		it.Event = new(PoolBinMerged)
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
func (it *PoolBinMergedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBinMergedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBinMerged represents a BinMerged event raised by the Pool contract.
type PoolBinMerged struct {
	BinId    *big.Int
	ReserveA *big.Int
	ReserveB *big.Int
	MergeId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBinMerged is a free log retrieval operation binding the contract event 0x8ecf1f9da718dc4c174482cdb4e334113856b46a85e5694deeec06d512e8f772.
//
// Solidity: event BinMerged(uint128 indexed binId, uint128 reserveA, uint128 reserveB, uint128 mergeId)
func (_Pool *PoolFilterer) FilterBinMerged(opts *bind.FilterOpts, binId []*big.Int) (*PoolBinMergedIterator, error) {

	var binIdRule []interface{}
	for _, binIdItem := range binId {
		binIdRule = append(binIdRule, binIdItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BinMerged", binIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolBinMergedIterator{contract: _Pool.contract, event: "BinMerged", logs: logs, sub: sub}, nil
}

// WatchBinMerged is a free log subscription operation binding the contract event 0x8ecf1f9da718dc4c174482cdb4e334113856b46a85e5694deeec06d512e8f772.
//
// Solidity: event BinMerged(uint128 indexed binId, uint128 reserveA, uint128 reserveB, uint128 mergeId)
func (_Pool *PoolFilterer) WatchBinMerged(opts *bind.WatchOpts, sink chan<- *PoolBinMerged, binId []*big.Int) (event.Subscription, error) {

	var binIdRule []interface{}
	for _, binIdItem := range binId {
		binIdRule = append(binIdRule, binIdItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BinMerged", binIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBinMerged)
				if err := _Pool.contract.UnpackLog(event, "BinMerged", log); err != nil {
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

// ParseBinMerged is a log parse operation binding the contract event 0x8ecf1f9da718dc4c174482cdb4e334113856b46a85e5694deeec06d512e8f772.
//
// Solidity: event BinMerged(uint128 indexed binId, uint128 reserveA, uint128 reserveB, uint128 mergeId)
func (_Pool *PoolFilterer) ParseBinMerged(log types.Log) (*PoolBinMerged, error) {
	event := new(PoolBinMerged)
	if err := _Pool.contract.UnpackLog(event, "BinMerged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolBinMovedIterator is returned from FilterBinMoved and is used to iterate over the raw logs and unpacked data for BinMoved events raised by the Pool contract.
type PoolBinMovedIterator struct {
	Event *PoolBinMoved // Event containing the contract specifics and raw log

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
func (it *PoolBinMovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolBinMoved)
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
		it.Event = new(PoolBinMoved)
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
func (it *PoolBinMovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolBinMovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolBinMoved represents a BinMoved event raised by the Pool contract.
type PoolBinMoved struct {
	BinId        *big.Int
	PreviousTick *big.Int
	NewTick      *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBinMoved is a free log retrieval operation binding the contract event 0x42e51620e75096344ac889cc1d899ab619aedbe89a4f6b230ee3cecb849c7e2f.
//
// Solidity: event BinMoved(uint128 indexed binId, int128 previousTick, int128 newTick)
func (_Pool *PoolFilterer) FilterBinMoved(opts *bind.FilterOpts, binId []*big.Int) (*PoolBinMovedIterator, error) {

	var binIdRule []interface{}
	for _, binIdItem := range binId {
		binIdRule = append(binIdRule, binIdItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "BinMoved", binIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolBinMovedIterator{contract: _Pool.contract, event: "BinMoved", logs: logs, sub: sub}, nil
}

// WatchBinMoved is a free log subscription operation binding the contract event 0x42e51620e75096344ac889cc1d899ab619aedbe89a4f6b230ee3cecb849c7e2f.
//
// Solidity: event BinMoved(uint128 indexed binId, int128 previousTick, int128 newTick)
func (_Pool *PoolFilterer) WatchBinMoved(opts *bind.WatchOpts, sink chan<- *PoolBinMoved, binId []*big.Int) (event.Subscription, error) {

	var binIdRule []interface{}
	for _, binIdItem := range binId {
		binIdRule = append(binIdRule, binIdItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "BinMoved", binIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolBinMoved)
				if err := _Pool.contract.UnpackLog(event, "BinMoved", log); err != nil {
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

// ParseBinMoved is a log parse operation binding the contract event 0x42e51620e75096344ac889cc1d899ab619aedbe89a4f6b230ee3cecb849c7e2f.
//
// Solidity: event BinMoved(uint128 indexed binId, int128 previousTick, int128 newTick)
func (_Pool *PoolFilterer) ParseBinMoved(log types.Log) (*PoolBinMoved, error) {
	event := new(PoolBinMoved)
	if err := _Pool.contract.UnpackLog(event, "BinMoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolMigrateBinsUpStackIterator is returned from FilterMigrateBinsUpStack and is used to iterate over the raw logs and unpacked data for MigrateBinsUpStack events raised by the Pool contract.
type PoolMigrateBinsUpStackIterator struct {
	Event *PoolMigrateBinsUpStack // Event containing the contract specifics and raw log

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
func (it *PoolMigrateBinsUpStackIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolMigrateBinsUpStack)
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
		it.Event = new(PoolMigrateBinsUpStack)
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
func (it *PoolMigrateBinsUpStackIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolMigrateBinsUpStackIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolMigrateBinsUpStack represents a MigrateBinsUpStack event raised by the Pool contract.
type PoolMigrateBinsUpStack struct {
	Sender       common.Address
	BinId        *big.Int
	MaxRecursion uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterMigrateBinsUpStack is a free log retrieval operation binding the contract event 0x6deceb91de75f84acd021df8c6410377aa442257495a79a9e3bfc7eba745853e.
//
// Solidity: event MigrateBinsUpStack(address indexed sender, uint128 binId, uint32 maxRecursion)
func (_Pool *PoolFilterer) FilterMigrateBinsUpStack(opts *bind.FilterOpts, sender []common.Address) (*PoolMigrateBinsUpStackIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "MigrateBinsUpStack", senderRule)
	if err != nil {
		return nil, err
	}
	return &PoolMigrateBinsUpStackIterator{contract: _Pool.contract, event: "MigrateBinsUpStack", logs: logs, sub: sub}, nil
}

// WatchMigrateBinsUpStack is a free log subscription operation binding the contract event 0x6deceb91de75f84acd021df8c6410377aa442257495a79a9e3bfc7eba745853e.
//
// Solidity: event MigrateBinsUpStack(address indexed sender, uint128 binId, uint32 maxRecursion)
func (_Pool *PoolFilterer) WatchMigrateBinsUpStack(opts *bind.WatchOpts, sink chan<- *PoolMigrateBinsUpStack, sender []common.Address) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "MigrateBinsUpStack", senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolMigrateBinsUpStack)
				if err := _Pool.contract.UnpackLog(event, "MigrateBinsUpStack", log); err != nil {
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

// ParseMigrateBinsUpStack is a log parse operation binding the contract event 0x6deceb91de75f84acd021df8c6410377aa442257495a79a9e3bfc7eba745853e.
//
// Solidity: event MigrateBinsUpStack(address indexed sender, uint128 binId, uint32 maxRecursion)
func (_Pool *PoolFilterer) ParseMigrateBinsUpStack(log types.Log) (*PoolMigrateBinsUpStack, error) {
	event := new(PoolMigrateBinsUpStack)
	if err := _Pool.contract.UnpackLog(event, "MigrateBinsUpStack", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolProtocolFeeCollectedIterator is returned from FilterProtocolFeeCollected and is used to iterate over the raw logs and unpacked data for ProtocolFeeCollected events raised by the Pool contract.
type PoolProtocolFeeCollectedIterator struct {
	Event *PoolProtocolFeeCollected // Event containing the contract specifics and raw log

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
func (it *PoolProtocolFeeCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolProtocolFeeCollected)
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
		it.Event = new(PoolProtocolFeeCollected)
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
func (it *PoolProtocolFeeCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolProtocolFeeCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolProtocolFeeCollected represents a ProtocolFeeCollected event raised by the Pool contract.
type PoolProtocolFeeCollected struct {
	ProtocolFee *big.Int
	IsTokenA    bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeCollected is a free log retrieval operation binding the contract event 0x292394e5b7a6b75d01122bb2dc85341cefec10b852325db9d3658a452f5eb211.
//
// Solidity: event ProtocolFeeCollected(uint256 protocolFee, bool isTokenA)
func (_Pool *PoolFilterer) FilterProtocolFeeCollected(opts *bind.FilterOpts) (*PoolProtocolFeeCollectedIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "ProtocolFeeCollected")
	if err != nil {
		return nil, err
	}
	return &PoolProtocolFeeCollectedIterator{contract: _Pool.contract, event: "ProtocolFeeCollected", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeCollected is a free log subscription operation binding the contract event 0x292394e5b7a6b75d01122bb2dc85341cefec10b852325db9d3658a452f5eb211.
//
// Solidity: event ProtocolFeeCollected(uint256 protocolFee, bool isTokenA)
func (_Pool *PoolFilterer) WatchProtocolFeeCollected(opts *bind.WatchOpts, sink chan<- *PoolProtocolFeeCollected) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "ProtocolFeeCollected")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolProtocolFeeCollected)
				if err := _Pool.contract.UnpackLog(event, "ProtocolFeeCollected", log); err != nil {
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

// ParseProtocolFeeCollected is a log parse operation binding the contract event 0x292394e5b7a6b75d01122bb2dc85341cefec10b852325db9d3658a452f5eb211.
//
// Solidity: event ProtocolFeeCollected(uint256 protocolFee, bool isTokenA)
func (_Pool *PoolFilterer) ParseProtocolFeeCollected(log types.Log) (*PoolProtocolFeeCollected, error) {
	event := new(PoolProtocolFeeCollected)
	if err := _Pool.contract.UnpackLog(event, "ProtocolFeeCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolRemoveLiquidityIterator is returned from FilterRemoveLiquidity and is used to iterate over the raw logs and unpacked data for RemoveLiquidity events raised by the Pool contract.
type PoolRemoveLiquidityIterator struct {
	Event *PoolRemoveLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolRemoveLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolRemoveLiquidity)
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
		it.Event = new(PoolRemoveLiquidity)
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
func (it *PoolRemoveLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolRemoveLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolRemoveLiquidity represents a RemoveLiquidity event raised by the Pool contract.
type PoolRemoveLiquidity struct {
	Sender    common.Address
	Recipient common.Address
	TokenId   *big.Int
	BinDeltas []IPoolBinDelta
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterRemoveLiquidity is a free log retrieval operation binding the contract event 0x65da280c1e973a1c5884c38d63e2c2b3c2a3158a0761e76545b64035e2489dfe.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed recipient, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) FilterRemoveLiquidity(opts *bind.FilterOpts, sender []common.Address, recipient []common.Address, tokenId []*big.Int) (*PoolRemoveLiquidityIterator, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Pool.contract.FilterLogs(opts, "RemoveLiquidity", senderRule, recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &PoolRemoveLiquidityIterator{contract: _Pool.contract, event: "RemoveLiquidity", logs: logs, sub: sub}, nil
}

// WatchRemoveLiquidity is a free log subscription operation binding the contract event 0x65da280c1e973a1c5884c38d63e2c2b3c2a3158a0761e76545b64035e2489dfe.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed recipient, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) WatchRemoveLiquidity(opts *bind.WatchOpts, sink chan<- *PoolRemoveLiquidity, sender []common.Address, recipient []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}
	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _Pool.contract.WatchLogs(opts, "RemoveLiquidity", senderRule, recipientRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolRemoveLiquidity)
				if err := _Pool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
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

// ParseRemoveLiquidity is a log parse operation binding the contract event 0x65da280c1e973a1c5884c38d63e2c2b3c2a3158a0761e76545b64035e2489dfe.
//
// Solidity: event RemoveLiquidity(address indexed sender, address indexed recipient, uint256 indexed tokenId, (uint128,uint128,uint256,uint128,uint8,int32,bool)[] binDeltas)
func (_Pool *PoolFilterer) ParseRemoveLiquidity(log types.Log) (*PoolRemoveLiquidity, error) {
	event := new(PoolRemoveLiquidity)
	if err := _Pool.contract.UnpackLog(event, "RemoveLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSetProtocolFeeRatioIterator is returned from FilterSetProtocolFeeRatio and is used to iterate over the raw logs and unpacked data for SetProtocolFeeRatio events raised by the Pool contract.
type PoolSetProtocolFeeRatioIterator struct {
	Event *PoolSetProtocolFeeRatio // Event containing the contract specifics and raw log

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
func (it *PoolSetProtocolFeeRatioIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSetProtocolFeeRatio)
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
		it.Event = new(PoolSetProtocolFeeRatio)
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
func (it *PoolSetProtocolFeeRatioIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSetProtocolFeeRatioIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSetProtocolFeeRatio represents a SetProtocolFeeRatio event raised by the Pool contract.
type PoolSetProtocolFeeRatio struct {
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSetProtocolFeeRatio is a free log retrieval operation binding the contract event 0x06e6ba2b10970ecae3ab2c29feb60ab2503358820756ef14a9827b0fa5add30f.
//
// Solidity: event SetProtocolFeeRatio(uint256 protocolFee)
func (_Pool *PoolFilterer) FilterSetProtocolFeeRatio(opts *bind.FilterOpts) (*PoolSetProtocolFeeRatioIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "SetProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return &PoolSetProtocolFeeRatioIterator{contract: _Pool.contract, event: "SetProtocolFeeRatio", logs: logs, sub: sub}, nil
}

// WatchSetProtocolFeeRatio is a free log subscription operation binding the contract event 0x06e6ba2b10970ecae3ab2c29feb60ab2503358820756ef14a9827b0fa5add30f.
//
// Solidity: event SetProtocolFeeRatio(uint256 protocolFee)
func (_Pool *PoolFilterer) WatchSetProtocolFeeRatio(opts *bind.WatchOpts, sink chan<- *PoolSetProtocolFeeRatio) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "SetProtocolFeeRatio")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSetProtocolFeeRatio)
				if err := _Pool.contract.UnpackLog(event, "SetProtocolFeeRatio", log); err != nil {
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

// ParseSetProtocolFeeRatio is a log parse operation binding the contract event 0x06e6ba2b10970ecae3ab2c29feb60ab2503358820756ef14a9827b0fa5add30f.
//
// Solidity: event SetProtocolFeeRatio(uint256 protocolFee)
func (_Pool *PoolFilterer) ParseSetProtocolFeeRatio(log types.Log) (*PoolSetProtocolFeeRatio, error) {
	event := new(PoolSetProtocolFeeRatio)
	if err := _Pool.contract.UnpackLog(event, "SetProtocolFeeRatio", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the Pool contract.
type PoolSwapIterator struct {
	Event *PoolSwap // Event containing the contract specifics and raw log

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
func (it *PoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolSwap)
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
		it.Event = new(PoolSwap)
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
func (it *PoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolSwap represents a Swap event raised by the Pool contract.
type PoolSwap struct {
	Sender      common.Address
	Recipient   common.Address
	TokenAIn    bool
	ExactOutput bool
	AmountIn    *big.Int
	AmountOut   *big.Int
	ActiveTick  int32
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0x3b841dc9ab51e3104bda4f61b41e4271192d22cd19da5ee6e292dc8e2744f713.
//
// Solidity: event Swap(address sender, address recipient, bool tokenAIn, bool exactOutput, uint256 amountIn, uint256 amountOut, int32 activeTick)
func (_Pool *PoolFilterer) FilterSwap(opts *bind.FilterOpts) (*PoolSwapIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return &PoolSwapIterator{contract: _Pool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0x3b841dc9ab51e3104bda4f61b41e4271192d22cd19da5ee6e292dc8e2744f713.
//
// Solidity: event Swap(address sender, address recipient, bool tokenAIn, bool exactOutput, uint256 amountIn, uint256 amountOut, int32 activeTick)
func (_Pool *PoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *PoolSwap) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "Swap")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolSwap)
				if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0x3b841dc9ab51e3104bda4f61b41e4271192d22cd19da5ee6e292dc8e2744f713.
//
// Solidity: event Swap(address sender, address recipient, bool tokenAIn, bool exactOutput, uint256 amountIn, uint256 amountOut, int32 activeTick)
func (_Pool *PoolFilterer) ParseSwap(log types.Log) (*PoolSwap, error) {
	event := new(PoolSwap)
	if err := _Pool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PoolTransferLiquidityIterator is returned from FilterTransferLiquidity and is used to iterate over the raw logs and unpacked data for TransferLiquidity events raised by the Pool contract.
type PoolTransferLiquidityIterator struct {
	Event *PoolTransferLiquidity // Event containing the contract specifics and raw log

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
func (it *PoolTransferLiquidityIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PoolTransferLiquidity)
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
		it.Event = new(PoolTransferLiquidity)
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
func (it *PoolTransferLiquidityIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PoolTransferLiquidityIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PoolTransferLiquidity represents a TransferLiquidity event raised by the Pool contract.
type PoolTransferLiquidity struct {
	FromTokenId *big.Int
	ToTokenId   *big.Int
	Params      []IPoolRemoveLiquidityParams
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferLiquidity is a free log retrieval operation binding the contract event 0xd384edefdfebd0bb45d82f94aed5ff327fd6510cc6c53ddc78a3ef4a0e7c715c.
//
// Solidity: event TransferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params)
func (_Pool *PoolFilterer) FilterTransferLiquidity(opts *bind.FilterOpts) (*PoolTransferLiquidityIterator, error) {

	logs, sub, err := _Pool.contract.FilterLogs(opts, "TransferLiquidity")
	if err != nil {
		return nil, err
	}
	return &PoolTransferLiquidityIterator{contract: _Pool.contract, event: "TransferLiquidity", logs: logs, sub: sub}, nil
}

// WatchTransferLiquidity is a free log subscription operation binding the contract event 0xd384edefdfebd0bb45d82f94aed5ff327fd6510cc6c53ddc78a3ef4a0e7c715c.
//
// Solidity: event TransferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params)
func (_Pool *PoolFilterer) WatchTransferLiquidity(opts *bind.WatchOpts, sink chan<- *PoolTransferLiquidity) (event.Subscription, error) {

	logs, sub, err := _Pool.contract.WatchLogs(opts, "TransferLiquidity")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PoolTransferLiquidity)
				if err := _Pool.contract.UnpackLog(event, "TransferLiquidity", log); err != nil {
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

// ParseTransferLiquidity is a log parse operation binding the contract event 0xd384edefdfebd0bb45d82f94aed5ff327fd6510cc6c53ddc78a3ef4a0e7c715c.
//
// Solidity: event TransferLiquidity(uint256 fromTokenId, uint256 toTokenId, (uint128,uint128)[] params)
func (_Pool *PoolFilterer) ParseTransferLiquidity(log types.Log) (*PoolTransferLiquidity, error) {
	event := new(PoolTransferLiquidity)
	if err := _Pool.contract.UnpackLog(event, "TransferLiquidity", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
