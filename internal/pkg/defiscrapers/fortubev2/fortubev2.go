// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package fortubev2

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

// Fortubev2ABI is the input ABI used to generate the binding from.
const Fortubev2ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"}],\"name\":\"AddTokenToMarket\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newCollateralAbility\",\"type\":\"uint256\"}],\"name\":\"_setCollateralAbility\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationIncentive\",\"type\":\"uint256\"}],\"name\":\"_setLiquidationIncentive\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIFToken\",\"name\":\"fToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_collateralAbility\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_liquidationIncentive\",\"type\":\"uint256\"}],\"name\":\"_supportMarket\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"accountAssets\",\"outputs\":[{\"internalType\":\"contractIFToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"addExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"addAmount\",\"type\":\"uint256\"}],\"name\":\"addReserves\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"admin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allMarkets\",\"outputs\":[{\"internalType\":\"contractIFToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allUnderlyingMarkets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"bankEntryAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"borrowCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"}],\"name\":\"calcExchangeUnit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calcMaxBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calcMaxBorrowAmountWithRatio\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calcMaxCashOutAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"calcMaxWithdrawAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasSpend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"}],\"name\":\"calcRewardAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"gasSpend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_for\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"}],\"name\":\"calcRewardAmountByFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIFToken\",\"name\":\"fToken\",\"type\":\"address\"}],\"name\":\"checkAccountsIn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"claimAdministration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"closeFactor\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"divExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"divScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"divScalarByExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"divisor\",\"type\":\"uint256\"}],\"name\":\"divScalarByExpTruncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"fetchAssetPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAccountLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"liquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"shortfall\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"getAccountLiquidityExcludeDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllMarkets\",\"outputs\":[{\"internalType\":\"contractIFToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getAssetsIn\",\"outputs\":[{\"internalType\":\"contractIFToken[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferInAmount\",\"type\":\"uint256\"}],\"name\":\"getCashAfter\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"getCashPrior\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"}],\"name\":\"getDiv\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rational\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"num\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"denom\",\"type\":\"uint256\"}],\"name\":\"getExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"rational\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"getFTokeAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"getTotalDepositAndBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIFToken\",\"name\":\"fTokenNow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getUserLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"contractIFToken\",\"name\":\"excludeToken\",\"type\":\"address\"},{\"internalType\":\"contractIFToken\",\"name\":\"fTokenNow\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTokens\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"getUserLiquidityExcludeToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_mulsig\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"}],\"name\":\"isFTokenValid\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"repayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateBorrowCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fTokenBorrowed\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"fTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"actualRepayAmount\",\"type\":\"uint256\"}],\"name\":\"liquidateTokens\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"markets\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"fTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"isValid\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"collateralAbility\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationIncentive\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"}],\"name\":\"marketsContains\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"mintCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"mulExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"c\",\"type\":\"uint256\"}],\"name\":\"mulExp3\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"mulScalar\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"scaled\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"}],\"name\":\"mulScalarTruncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scalar\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"addend\",\"type\":\"uint256\"}],\"name\":\"mulScalarTruncateAddUInt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"mulsig\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"oracle\",\"outputs\":[{\"internalType\":\"contractIOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin_\",\"type\":\"address\"}],\"name\":\"proposeNewAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"proposedAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"reduceAmount\",\"type\":\"uint256\"}],\"name\":\"reduceReserves\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"}],\"name\":\"repayCheck\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"rewardFactors\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gasSpend\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"gasPrice\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rewardType\",\"type\":\"uint256\"}],\"name\":\"rewardForByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"rewardPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"cTokenCollateral\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"cTokenBorrowed\",\"type\":\"address\"}],\"name\":\"seizeCheck\",\"outputs\":[],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newBank\",\"type\":\"address\"}],\"name\":\"setBankEntryAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_closeFactor\",\"type\":\"uint256\"}],\"name\":\"setCloseFactor\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_oracle\",\"type\":\"address\"}],\"name\":\"setOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"rewaradType\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"factor\",\"type\":\"uint256\"}],\"name\":\"setRewardFactorByType\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_rewardPool\",\"type\":\"address\"}],\"name\":\"setRewardPool\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_theForceToken\",\"type\":\"address\"}],\"name\":\"setTheForceToken\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_transferEthGasCost\",\"type\":\"uint256\"}],\"name\":\"setTransferEthGasCost\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"a\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"b\",\"type\":\"uint256\"}],\"name\":\"subExp\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"result\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"theForceToken\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"tokenDecimals\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"src\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"dst\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"transferTokens\",\"type\":\"uint256\"}],\"name\":\"transferCheck\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"transferEthGasCost\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferIn\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"underlying\",\"type\":\"address\"},{\"internalType\":\"addresspayable\",\"name\":\"account\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferToUser\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"exp\",\"type\":\"uint256\"}],\"name\":\"truncate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"pure\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"fToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"withdrawer\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"withdrawTokens\",\"type\":\"uint256\"}],\"name\":\"withdrawCheck\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Fortubev2 is an auto generated Go binding around an Ethereum contract.
type Fortubev2 struct {
	Fortubev2Caller     // Read-only binding to the contract
	Fortubev2Transactor // Write-only binding to the contract
	Fortubev2Filterer   // Log filterer for contract events
}

// Fortubev2Caller is an auto generated read-only Go binding around an Ethereum contract.
type Fortubev2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fortubev2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Fortubev2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fortubev2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Fortubev2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Fortubev2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Fortubev2Session struct {
	Contract     *Fortubev2        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Fortubev2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Fortubev2CallerSession struct {
	Contract *Fortubev2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// Fortubev2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Fortubev2TransactorSession struct {
	Contract     *Fortubev2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// Fortubev2Raw is an auto generated low-level Go binding around an Ethereum contract.
type Fortubev2Raw struct {
	Contract *Fortubev2 // Generic contract binding to access the raw methods on
}

// Fortubev2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Fortubev2CallerRaw struct {
	Contract *Fortubev2Caller // Generic read-only contract binding to access the raw methods on
}

// Fortubev2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Fortubev2TransactorRaw struct {
	Contract *Fortubev2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFortubev2 creates a new instance of Fortubev2, bound to a specific deployed contract.
func NewFortubev2(address common.Address, backend bind.ContractBackend) (*Fortubev2, error) {
	contract, err := bindFortubev2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Fortubev2{Fortubev2Caller: Fortubev2Caller{contract: contract}, Fortubev2Transactor: Fortubev2Transactor{contract: contract}, Fortubev2Filterer: Fortubev2Filterer{contract: contract}}, nil
}

// NewFortubev2Caller creates a new read-only instance of Fortubev2, bound to a specific deployed contract.
func NewFortubev2Caller(address common.Address, caller bind.ContractCaller) (*Fortubev2Caller, error) {
	contract, err := bindFortubev2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Fortubev2Caller{contract: contract}, nil
}

// NewFortubev2Transactor creates a new write-only instance of Fortubev2, bound to a specific deployed contract.
func NewFortubev2Transactor(address common.Address, transactor bind.ContractTransactor) (*Fortubev2Transactor, error) {
	contract, err := bindFortubev2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Fortubev2Transactor{contract: contract}, nil
}

// NewFortubev2Filterer creates a new log filterer instance of Fortubev2, bound to a specific deployed contract.
func NewFortubev2Filterer(address common.Address, filterer bind.ContractFilterer) (*Fortubev2Filterer, error) {
	contract, err := bindFortubev2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Fortubev2Filterer{contract: contract}, nil
}

// bindFortubev2 binds a generic wrapper to an already deployed contract.
func bindFortubev2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Fortubev2ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fortubev2 *Fortubev2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fortubev2.Contract.Fortubev2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fortubev2 *Fortubev2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fortubev2.Contract.Fortubev2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fortubev2 *Fortubev2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fortubev2.Contract.Fortubev2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Fortubev2 *Fortubev2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Fortubev2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Fortubev2 *Fortubev2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fortubev2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Fortubev2 *Fortubev2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Fortubev2.Contract.contract.Transact(opts, method, params...)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Caller) AccountAssets(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "accountAssets", arg0, arg1)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Session) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AccountAssets(&_Fortubev2.CallOpts, arg0, arg1)
}

// AccountAssets is a free data retrieval call binding the contract method 0xdce15449.
//
// Solidity: function accountAssets(address , uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) AccountAssets(arg0 common.Address, arg1 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AccountAssets(&_Fortubev2.CallOpts, arg0, arg1)
}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2Caller) AddExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "addExp", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2Session) AddExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.AddExp(&_Fortubev2.CallOpts, a, b)
}

// AddExp is a free data retrieval call binding the contract method 0x4147da13.
//
// Solidity: function addExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2CallerSession) AddExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.AddExp(&_Fortubev2.CallOpts, a, b)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Fortubev2 *Fortubev2Caller) Admin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "admin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Fortubev2 *Fortubev2Session) Admin() (common.Address, error) {
	return _Fortubev2.Contract.Admin(&_Fortubev2.CallOpts)
}

// Admin is a free data retrieval call binding the contract method 0xf851a440.
//
// Solidity: function admin() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) Admin() (common.Address, error) {
	return _Fortubev2.Contract.Admin(&_Fortubev2.CallOpts)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Caller) AllMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "allMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Session) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AllMarkets(&_Fortubev2.CallOpts, arg0)
}

// AllMarkets is a free data retrieval call binding the contract method 0x52d84d1e.
//
// Solidity: function allMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) AllMarkets(arg0 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AllMarkets(&_Fortubev2.CallOpts, arg0)
}

// AllUnderlyingMarkets is a free data retrieval call binding the contract method 0x941a9b11.
//
// Solidity: function allUnderlyingMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Caller) AllUnderlyingMarkets(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "allUnderlyingMarkets", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AllUnderlyingMarkets is a free data retrieval call binding the contract method 0x941a9b11.
//
// Solidity: function allUnderlyingMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2Session) AllUnderlyingMarkets(arg0 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AllUnderlyingMarkets(&_Fortubev2.CallOpts, arg0)
}

// AllUnderlyingMarkets is a free data retrieval call binding the contract method 0x941a9b11.
//
// Solidity: function allUnderlyingMarkets(uint256 ) view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) AllUnderlyingMarkets(arg0 *big.Int) (common.Address, error) {
	return _Fortubev2.Contract.AllUnderlyingMarkets(&_Fortubev2.CallOpts, arg0)
}

// BankEntryAddress is a free data retrieval call binding the contract method 0x94df6bc4.
//
// Solidity: function bankEntryAddress() view returns(address)
func (_Fortubev2 *Fortubev2Caller) BankEntryAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "bankEntryAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BankEntryAddress is a free data retrieval call binding the contract method 0x94df6bc4.
//
// Solidity: function bankEntryAddress() view returns(address)
func (_Fortubev2 *Fortubev2Session) BankEntryAddress() (common.Address, error) {
	return _Fortubev2.Contract.BankEntryAddress(&_Fortubev2.CallOpts)
}

// BankEntryAddress is a free data retrieval call binding the contract method 0x94df6bc4.
//
// Solidity: function bankEntryAddress() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) BankEntryAddress() (common.Address, error) {
	return _Fortubev2.Contract.BankEntryAddress(&_Fortubev2.CallOpts)
}

// CalcExchangeUnit is a free data retrieval call binding the contract method 0x4d7ab633.
//
// Solidity: function calcExchangeUnit(address fToken) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcExchangeUnit(opts *bind.CallOpts, fToken common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcExchangeUnit", fToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcExchangeUnit is a free data retrieval call binding the contract method 0x4d7ab633.
//
// Solidity: function calcExchangeUnit(address fToken) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcExchangeUnit(fToken common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcExchangeUnit(&_Fortubev2.CallOpts, fToken)
}

// CalcExchangeUnit is a free data retrieval call binding the contract method 0x4d7ab633.
//
// Solidity: function calcExchangeUnit(address fToken) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcExchangeUnit(fToken common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcExchangeUnit(&_Fortubev2.CallOpts, fToken)
}

// CalcMaxBorrowAmount is a free data retrieval call binding the contract method 0xec63aaa3.
//
// Solidity: function calcMaxBorrowAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcMaxBorrowAmount(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcMaxBorrowAmount", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxBorrowAmount is a free data retrieval call binding the contract method 0xec63aaa3.
//
// Solidity: function calcMaxBorrowAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcMaxBorrowAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxBorrowAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxBorrowAmount is a free data retrieval call binding the contract method 0xec63aaa3.
//
// Solidity: function calcMaxBorrowAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcMaxBorrowAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxBorrowAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxBorrowAmountWithRatio is a free data retrieval call binding the contract method 0x75861c6f.
//
// Solidity: function calcMaxBorrowAmountWithRatio(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcMaxBorrowAmountWithRatio(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcMaxBorrowAmountWithRatio", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxBorrowAmountWithRatio is a free data retrieval call binding the contract method 0x75861c6f.
//
// Solidity: function calcMaxBorrowAmountWithRatio(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcMaxBorrowAmountWithRatio(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxBorrowAmountWithRatio(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxBorrowAmountWithRatio is a free data retrieval call binding the contract method 0x75861c6f.
//
// Solidity: function calcMaxBorrowAmountWithRatio(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcMaxBorrowAmountWithRatio(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxBorrowAmountWithRatio(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxCashOutAmount is a free data retrieval call binding the contract method 0xa9dc394c.
//
// Solidity: function calcMaxCashOutAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcMaxCashOutAmount(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcMaxCashOutAmount", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxCashOutAmount is a free data retrieval call binding the contract method 0xa9dc394c.
//
// Solidity: function calcMaxCashOutAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcMaxCashOutAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxCashOutAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxCashOutAmount is a free data retrieval call binding the contract method 0xa9dc394c.
//
// Solidity: function calcMaxCashOutAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcMaxCashOutAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxCashOutAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxWithdrawAmount is a free data retrieval call binding the contract method 0x2b91acc0.
//
// Solidity: function calcMaxWithdrawAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcMaxWithdrawAmount(opts *bind.CallOpts, user common.Address, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcMaxWithdrawAmount", user, token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcMaxWithdrawAmount is a free data retrieval call binding the contract method 0x2b91acc0.
//
// Solidity: function calcMaxWithdrawAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcMaxWithdrawAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxWithdrawAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcMaxWithdrawAmount is a free data retrieval call binding the contract method 0x2b91acc0.
//
// Solidity: function calcMaxWithdrawAmount(address user, address token) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcMaxWithdrawAmount(user common.Address, token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcMaxWithdrawAmount(&_Fortubev2.CallOpts, user, token)
}

// CalcRewardAmount is a free data retrieval call binding the contract method 0x2ddb60a3.
//
// Solidity: function calcRewardAmount(uint256 gasSpend, uint256 gasPrice, address _for) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcRewardAmount(opts *bind.CallOpts, gasSpend *big.Int, gasPrice *big.Int, _for common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcRewardAmount", gasSpend, gasPrice, _for)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRewardAmount is a free data retrieval call binding the contract method 0x2ddb60a3.
//
// Solidity: function calcRewardAmount(uint256 gasSpend, uint256 gasPrice, address _for) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcRewardAmount(gasSpend *big.Int, gasPrice *big.Int, _for common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcRewardAmount(&_Fortubev2.CallOpts, gasSpend, gasPrice, _for)
}

// CalcRewardAmount is a free data retrieval call binding the contract method 0x2ddb60a3.
//
// Solidity: function calcRewardAmount(uint256 gasSpend, uint256 gasPrice, address _for) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcRewardAmount(gasSpend *big.Int, gasPrice *big.Int, _for common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.CalcRewardAmount(&_Fortubev2.CallOpts, gasSpend, gasPrice, _for)
}

// CalcRewardAmountByFactor is a free data retrieval call binding the contract method 0x9ee30be8.
//
// Solidity: function calcRewardAmountByFactor(uint256 gasSpend, uint256 gasPrice, address _for, uint256 factor) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CalcRewardAmountByFactor(opts *bind.CallOpts, gasSpend *big.Int, gasPrice *big.Int, _for common.Address, factor *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "calcRewardAmountByFactor", gasSpend, gasPrice, _for, factor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CalcRewardAmountByFactor is a free data retrieval call binding the contract method 0x9ee30be8.
//
// Solidity: function calcRewardAmountByFactor(uint256 gasSpend, uint256 gasPrice, address _for, uint256 factor) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CalcRewardAmountByFactor(gasSpend *big.Int, gasPrice *big.Int, _for common.Address, factor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.CalcRewardAmountByFactor(&_Fortubev2.CallOpts, gasSpend, gasPrice, _for, factor)
}

// CalcRewardAmountByFactor is a free data retrieval call binding the contract method 0x9ee30be8.
//
// Solidity: function calcRewardAmountByFactor(uint256 gasSpend, uint256 gasPrice, address _for, uint256 factor) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CalcRewardAmountByFactor(gasSpend *big.Int, gasPrice *big.Int, _for common.Address, factor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.CalcRewardAmountByFactor(&_Fortubev2.CallOpts, gasSpend, gasPrice, _for, factor)
}

// CheckAccountsIn is a free data retrieval call binding the contract method 0xa53caa16.
//
// Solidity: function checkAccountsIn(address account, address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Caller) CheckAccountsIn(opts *bind.CallOpts, account common.Address, fToken common.Address) (bool, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "checkAccountsIn", account, fToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CheckAccountsIn is a free data retrieval call binding the contract method 0xa53caa16.
//
// Solidity: function checkAccountsIn(address account, address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Session) CheckAccountsIn(account common.Address, fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.CheckAccountsIn(&_Fortubev2.CallOpts, account, fToken)
}

// CheckAccountsIn is a free data retrieval call binding the contract method 0xa53caa16.
//
// Solidity: function checkAccountsIn(address account, address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2CallerSession) CheckAccountsIn(account common.Address, fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.CheckAccountsIn(&_Fortubev2.CallOpts, account, fToken)
}

// CloseFactor is a free data retrieval call binding the contract method 0x05308b9f.
//
// Solidity: function closeFactor() view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) CloseFactor(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "closeFactor")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CloseFactor is a free data retrieval call binding the contract method 0x05308b9f.
//
// Solidity: function closeFactor() view returns(uint256)
func (_Fortubev2 *Fortubev2Session) CloseFactor() (*big.Int, error) {
	return _Fortubev2.Contract.CloseFactor(&_Fortubev2.CallOpts)
}

// CloseFactor is a free data retrieval call binding the contract method 0x05308b9f.
//
// Solidity: function closeFactor() view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) CloseFactor() (*big.Int, error) {
	return _Fortubev2.Contract.CloseFactor(&_Fortubev2.CallOpts)
}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) DivExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "divExp", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) DivExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivExp(&_Fortubev2.CallOpts, a, b)
}

// DivExp is a free data retrieval call binding the contract method 0xb507f9f9.
//
// Solidity: function divExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) DivExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivExp(&_Fortubev2.CallOpts, a, b)
}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) DivScalar(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "divScalar", a, scalar)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) DivScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalar(&_Fortubev2.CallOpts, a, scalar)
}

// DivScalar is a free data retrieval call binding the contract method 0x4751b79c.
//
// Solidity: function divScalar(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) DivScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalar(&_Fortubev2.CallOpts, a, scalar)
}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) DivScalarByExp(opts *bind.CallOpts, scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "divScalarByExp", scalar, divisor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) DivScalarByExp(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalarByExp(&_Fortubev2.CallOpts, scalar, divisor)
}

// DivScalarByExp is a free data retrieval call binding the contract method 0xce603aad.
//
// Solidity: function divScalarByExp(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) DivScalarByExp(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalarByExp(&_Fortubev2.CallOpts, scalar, divisor)
}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) DivScalarByExpTruncate(opts *bind.CallOpts, scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "divScalarByExpTruncate", scalar, divisor)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) DivScalarByExpTruncate(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalarByExpTruncate(&_Fortubev2.CallOpts, scalar, divisor)
}

// DivScalarByExpTruncate is a free data retrieval call binding the contract method 0xd6079cc6.
//
// Solidity: function divScalarByExpTruncate(uint256 scalar, uint256 divisor) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) DivScalarByExpTruncate(scalar *big.Int, divisor *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.DivScalarByExpTruncate(&_Fortubev2.CallOpts, scalar, divisor)
}

// FetchAssetPrice is a free data retrieval call binding the contract method 0xddec280e.
//
// Solidity: function fetchAssetPrice(address token) view returns(uint256, bool)
func (_Fortubev2 *Fortubev2Caller) FetchAssetPrice(opts *bind.CallOpts, token common.Address) (*big.Int, bool, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "fetchAssetPrice", token)

	if err != nil {
		return *new(*big.Int), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)

	return out0, out1, err

}

// FetchAssetPrice is a free data retrieval call binding the contract method 0xddec280e.
//
// Solidity: function fetchAssetPrice(address token) view returns(uint256, bool)
func (_Fortubev2 *Fortubev2Session) FetchAssetPrice(token common.Address) (*big.Int, bool, error) {
	return _Fortubev2.Contract.FetchAssetPrice(&_Fortubev2.CallOpts, token)
}

// FetchAssetPrice is a free data retrieval call binding the contract method 0xddec280e.
//
// Solidity: function fetchAssetPrice(address token) view returns(uint256, bool)
func (_Fortubev2 *Fortubev2CallerSession) FetchAssetPrice(token common.Address) (*big.Int, bool, error) {
	return _Fortubev2.Contract.FetchAssetPrice(&_Fortubev2.CallOpts, token)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256 liquidity, uint256 shortfall)
func (_Fortubev2 *Fortubev2Caller) GetAccountLiquidity(opts *bind.CallOpts, account common.Address) (struct {
	Liquidity *big.Int
	Shortfall *big.Int
}, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getAccountLiquidity", account)

	outstruct := new(struct {
		Liquidity *big.Int
		Shortfall *big.Int
	})

	outstruct.Liquidity = out[0].(*big.Int)
	outstruct.Shortfall = out[1].(*big.Int)

	return *outstruct, err

}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256 liquidity, uint256 shortfall)
func (_Fortubev2 *Fortubev2Session) GetAccountLiquidity(account common.Address) (struct {
	Liquidity *big.Int
	Shortfall *big.Int
}, error) {
	return _Fortubev2.Contract.GetAccountLiquidity(&_Fortubev2.CallOpts, account)
}

// GetAccountLiquidity is a free data retrieval call binding the contract method 0x5ec88c79.
//
// Solidity: function getAccountLiquidity(address account) view returns(uint256 liquidity, uint256 shortfall)
func (_Fortubev2 *Fortubev2CallerSession) GetAccountLiquidity(account common.Address) (struct {
	Liquidity *big.Int
	Shortfall *big.Int
}, error) {
	return _Fortubev2.Contract.GetAccountLiquidity(&_Fortubev2.CallOpts, account)
}

// GetAccountLiquidityExcludeDeposit is a free data retrieval call binding the contract method 0x70de3062.
//
// Solidity: function getAccountLiquidityExcludeDeposit(address account, address token) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Caller) GetAccountLiquidityExcludeDeposit(opts *bind.CallOpts, account common.Address, token common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getAccountLiquidityExcludeDeposit", account, token)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetAccountLiquidityExcludeDeposit is a free data retrieval call binding the contract method 0x70de3062.
//
// Solidity: function getAccountLiquidityExcludeDeposit(address account, address token) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Session) GetAccountLiquidityExcludeDeposit(account common.Address, token common.Address) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetAccountLiquidityExcludeDeposit(&_Fortubev2.CallOpts, account, token)
}

// GetAccountLiquidityExcludeDeposit is a free data retrieval call binding the contract method 0x70de3062.
//
// Solidity: function getAccountLiquidityExcludeDeposit(address account, address token) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetAccountLiquidityExcludeDeposit(account common.Address, token common.Address) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetAccountLiquidityExcludeDeposit(&_Fortubev2.CallOpts, account, token)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Fortubev2 *Fortubev2Caller) GetAllMarkets(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getAllMarkets")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Fortubev2 *Fortubev2Session) GetAllMarkets() ([]common.Address, error) {
	return _Fortubev2.Contract.GetAllMarkets(&_Fortubev2.CallOpts)
}

// GetAllMarkets is a free data retrieval call binding the contract method 0xb0772d0b.
//
// Solidity: function getAllMarkets() view returns(address[])
func (_Fortubev2 *Fortubev2CallerSession) GetAllMarkets() ([]common.Address, error) {
	return _Fortubev2.Contract.GetAllMarkets(&_Fortubev2.CallOpts)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Fortubev2 *Fortubev2Caller) GetAssetsIn(opts *bind.CallOpts, account common.Address) ([]common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getAssetsIn", account)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Fortubev2 *Fortubev2Session) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Fortubev2.Contract.GetAssetsIn(&_Fortubev2.CallOpts, account)
}

// GetAssetsIn is a free data retrieval call binding the contract method 0xabfceffc.
//
// Solidity: function getAssetsIn(address account) view returns(address[])
func (_Fortubev2 *Fortubev2CallerSession) GetAssetsIn(account common.Address) ([]common.Address, error) {
	return _Fortubev2.Contract.GetAssetsIn(&_Fortubev2.CallOpts, account)
}

// GetCashAfter is a free data retrieval call binding the contract method 0xa3f0fa20.
//
// Solidity: function getCashAfter(address underlying, uint256 transferInAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) GetCashAfter(opts *bind.CallOpts, underlying common.Address, transferInAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getCashAfter", underlying, transferInAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCashAfter is a free data retrieval call binding the contract method 0xa3f0fa20.
//
// Solidity: function getCashAfter(address underlying, uint256 transferInAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) GetCashAfter(underlying common.Address, transferInAmount *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetCashAfter(&_Fortubev2.CallOpts, underlying, transferInAmount)
}

// GetCashAfter is a free data retrieval call binding the contract method 0xa3f0fa20.
//
// Solidity: function getCashAfter(address underlying, uint256 transferInAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetCashAfter(underlying common.Address, transferInAmount *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetCashAfter(&_Fortubev2.CallOpts, underlying, transferInAmount)
}

// GetCashPrior is a free data retrieval call binding the contract method 0x9f77437b.
//
// Solidity: function getCashPrior(address underlying) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) GetCashPrior(opts *bind.CallOpts, underlying common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getCashPrior", underlying)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCashPrior is a free data retrieval call binding the contract method 0x9f77437b.
//
// Solidity: function getCashPrior(address underlying) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) GetCashPrior(underlying common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.GetCashPrior(&_Fortubev2.CallOpts, underlying)
}

// GetCashPrior is a free data retrieval call binding the contract method 0x9f77437b.
//
// Solidity: function getCashPrior(address underlying) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetCashPrior(underlying common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.GetCashPrior(&_Fortubev2.CallOpts, underlying)
}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2Caller) GetDiv(opts *bind.CallOpts, num *big.Int, denom *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getDiv", num, denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2Session) GetDiv(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetDiv(&_Fortubev2.CallOpts, num, denom)
}

// GetDiv is a free data retrieval call binding the contract method 0x5293ff31.
//
// Solidity: function getDiv(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2CallerSession) GetDiv(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetDiv(&_Fortubev2.CallOpts, num, denom)
}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2Caller) GetExp(opts *bind.CallOpts, num *big.Int, denom *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getExp", num, denom)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2Session) GetExp(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetExp(&_Fortubev2.CallOpts, num, denom)
}

// GetExp is a free data retrieval call binding the contract method 0xba9316b7.
//
// Solidity: function getExp(uint256 num, uint256 denom) pure returns(uint256 rational)
func (_Fortubev2 *Fortubev2CallerSession) GetExp(num *big.Int, denom *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.GetExp(&_Fortubev2.CallOpts, num, denom)
}

// GetFTokeAddress is a free data retrieval call binding the contract method 0xfc1ce708.
//
// Solidity: function getFTokeAddress(address underlying) view returns(address)
func (_Fortubev2 *Fortubev2Caller) GetFTokeAddress(opts *bind.CallOpts, underlying common.Address) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getFTokeAddress", underlying)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetFTokeAddress is a free data retrieval call binding the contract method 0xfc1ce708.
//
// Solidity: function getFTokeAddress(address underlying) view returns(address)
func (_Fortubev2 *Fortubev2Session) GetFTokeAddress(underlying common.Address) (common.Address, error) {
	return _Fortubev2.Contract.GetFTokeAddress(&_Fortubev2.CallOpts, underlying)
}

// GetFTokeAddress is a free data retrieval call binding the contract method 0xfc1ce708.
//
// Solidity: function getFTokeAddress(address underlying) view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) GetFTokeAddress(underlying common.Address) (common.Address, error) {
	return _Fortubev2.Contract.GetFTokeAddress(&_Fortubev2.CallOpts, underlying)
}

// GetTotalDepositAndBorrow is a free data retrieval call binding the contract method 0x8e93aa73.
//
// Solidity: function getTotalDepositAndBorrow(address account) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Caller) GetTotalDepositAndBorrow(opts *bind.CallOpts, account common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getTotalDepositAndBorrow", account)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalDepositAndBorrow is a free data retrieval call binding the contract method 0x8e93aa73.
//
// Solidity: function getTotalDepositAndBorrow(address account) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Session) GetTotalDepositAndBorrow(account common.Address) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetTotalDepositAndBorrow(&_Fortubev2.CallOpts, account)
}

// GetTotalDepositAndBorrow is a free data retrieval call binding the contract method 0x8e93aa73.
//
// Solidity: function getTotalDepositAndBorrow(address account) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetTotalDepositAndBorrow(account common.Address) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetTotalDepositAndBorrow(&_Fortubev2.CallOpts, account)
}

// GetUserLiquidity is a free data retrieval call binding the contract method 0x87dac0cc.
//
// Solidity: function getUserLiquidity(address account, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Caller) GetUserLiquidity(opts *bind.CallOpts, account common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getUserLiquidity", account, fTokenNow, withdrawTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserLiquidity is a free data retrieval call binding the contract method 0x87dac0cc.
//
// Solidity: function getUserLiquidity(address account, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Session) GetUserLiquidity(account common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetUserLiquidity(&_Fortubev2.CallOpts, account, fTokenNow, withdrawTokens, borrowAmount)
}

// GetUserLiquidity is a free data retrieval call binding the contract method 0x87dac0cc.
//
// Solidity: function getUserLiquidity(address account, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetUserLiquidity(account common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetUserLiquidity(&_Fortubev2.CallOpts, account, fTokenNow, withdrawTokens, borrowAmount)
}

// GetUserLiquidityExcludeToken is a free data retrieval call binding the contract method 0x8ac4ddc8.
//
// Solidity: function getUserLiquidityExcludeToken(address account, address excludeToken, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Caller) GetUserLiquidityExcludeToken(opts *bind.CallOpts, account common.Address, excludeToken common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "getUserLiquidityExcludeToken", account, excludeToken, fTokenNow, withdrawTokens, borrowAmount)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetUserLiquidityExcludeToken is a free data retrieval call binding the contract method 0x8ac4ddc8.
//
// Solidity: function getUserLiquidityExcludeToken(address account, address excludeToken, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2Session) GetUserLiquidityExcludeToken(account common.Address, excludeToken common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetUserLiquidityExcludeToken(&_Fortubev2.CallOpts, account, excludeToken, fTokenNow, withdrawTokens, borrowAmount)
}

// GetUserLiquidityExcludeToken is a free data retrieval call binding the contract method 0x8ac4ddc8.
//
// Solidity: function getUserLiquidityExcludeToken(address account, address excludeToken, address fTokenNow, uint256 withdrawTokens, uint256 borrowAmount) view returns(uint256, uint256)
func (_Fortubev2 *Fortubev2CallerSession) GetUserLiquidityExcludeToken(account common.Address, excludeToken common.Address, fTokenNow common.Address, withdrawTokens *big.Int, borrowAmount *big.Int) (*big.Int, *big.Int, error) {
	return _Fortubev2.Contract.GetUserLiquidityExcludeToken(&_Fortubev2.CallOpts, account, excludeToken, fTokenNow, withdrawTokens, borrowAmount)
}

// IsFTokenValid is a free data retrieval call binding the contract method 0x2a4d98cf.
//
// Solidity: function isFTokenValid(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Caller) IsFTokenValid(opts *bind.CallOpts, fToken common.Address) (bool, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "isFTokenValid", fToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsFTokenValid is a free data retrieval call binding the contract method 0x2a4d98cf.
//
// Solidity: function isFTokenValid(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Session) IsFTokenValid(fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.IsFTokenValid(&_Fortubev2.CallOpts, fToken)
}

// IsFTokenValid is a free data retrieval call binding the contract method 0x2a4d98cf.
//
// Solidity: function isFTokenValid(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2CallerSession) IsFTokenValid(fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.IsFTokenValid(&_Fortubev2.CallOpts, fToken)
}

// LiquidateTokens is a free data retrieval call binding the contract method 0x0c7fa6e0.
//
// Solidity: function liquidateTokens(address fTokenBorrowed, address fTokenCollateral, uint256 actualRepayAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) LiquidateTokens(opts *bind.CallOpts, fTokenBorrowed common.Address, fTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "liquidateTokens", fTokenBorrowed, fTokenCollateral, actualRepayAmount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LiquidateTokens is a free data retrieval call binding the contract method 0x0c7fa6e0.
//
// Solidity: function liquidateTokens(address fTokenBorrowed, address fTokenCollateral, uint256 actualRepayAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) LiquidateTokens(fTokenBorrowed common.Address, fTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.LiquidateTokens(&_Fortubev2.CallOpts, fTokenBorrowed, fTokenCollateral, actualRepayAmount)
}

// LiquidateTokens is a free data retrieval call binding the contract method 0x0c7fa6e0.
//
// Solidity: function liquidateTokens(address fTokenBorrowed, address fTokenCollateral, uint256 actualRepayAmount) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) LiquidateTokens(fTokenBorrowed common.Address, fTokenCollateral common.Address, actualRepayAmount *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.LiquidateTokens(&_Fortubev2.CallOpts, fTokenBorrowed, fTokenCollateral, actualRepayAmount)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(address fTokenAddress, bool isValid, uint256 collateralAbility, uint256 liquidationIncentive)
func (_Fortubev2 *Fortubev2Caller) Markets(opts *bind.CallOpts, arg0 common.Address) (struct {
	FTokenAddress        common.Address
	IsValid              bool
	CollateralAbility    *big.Int
	LiquidationIncentive *big.Int
}, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "markets", arg0)

	outstruct := new(struct {
		FTokenAddress        common.Address
		IsValid              bool
		CollateralAbility    *big.Int
		LiquidationIncentive *big.Int
	})

	outstruct.FTokenAddress = out[0].(common.Address)
	outstruct.IsValid = out[1].(bool)
	outstruct.CollateralAbility = out[2].(*big.Int)
	outstruct.LiquidationIncentive = out[3].(*big.Int)

	return *outstruct, err

}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(address fTokenAddress, bool isValid, uint256 collateralAbility, uint256 liquidationIncentive)
func (_Fortubev2 *Fortubev2Session) Markets(arg0 common.Address) (struct {
	FTokenAddress        common.Address
	IsValid              bool
	CollateralAbility    *big.Int
	LiquidationIncentive *big.Int
}, error) {
	return _Fortubev2.Contract.Markets(&_Fortubev2.CallOpts, arg0)
}

// Markets is a free data retrieval call binding the contract method 0x8e8f294b.
//
// Solidity: function markets(address ) view returns(address fTokenAddress, bool isValid, uint256 collateralAbility, uint256 liquidationIncentive)
func (_Fortubev2 *Fortubev2CallerSession) Markets(arg0 common.Address) (struct {
	FTokenAddress        common.Address
	IsValid              bool
	CollateralAbility    *big.Int
	LiquidationIncentive *big.Int
}, error) {
	return _Fortubev2.Contract.Markets(&_Fortubev2.CallOpts, arg0)
}

// MarketsContains is a free data retrieval call binding the contract method 0x9deec7cb.
//
// Solidity: function marketsContains(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Caller) MarketsContains(opts *bind.CallOpts, fToken common.Address) (bool, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "marketsContains", fToken)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketsContains is a free data retrieval call binding the contract method 0x9deec7cb.
//
// Solidity: function marketsContains(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2Session) MarketsContains(fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.MarketsContains(&_Fortubev2.CallOpts, fToken)
}

// MarketsContains is a free data retrieval call binding the contract method 0x9deec7cb.
//
// Solidity: function marketsContains(address fToken) view returns(bool)
func (_Fortubev2 *Fortubev2CallerSession) MarketsContains(fToken common.Address) (bool, error) {
	return _Fortubev2.Contract.MarketsContains(&_Fortubev2.CallOpts, fToken)
}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) MulExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulExp", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) MulExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulExp(&_Fortubev2.CallOpts, a, b)
}

// MulExp is a free data retrieval call binding the contract method 0xde32abd1.
//
// Solidity: function mulExp(uint256 a, uint256 b) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) MulExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulExp(&_Fortubev2.CallOpts, a, b)
}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) MulExp3(opts *bind.CallOpts, a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulExp3", a, b, c)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) MulExp3(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulExp3(&_Fortubev2.CallOpts, a, b, c)
}

// MulExp3 is a free data retrieval call binding the contract method 0x8de46362.
//
// Solidity: function mulExp3(uint256 a, uint256 b, uint256 c) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) MulExp3(a *big.Int, b *big.Int, c *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulExp3(&_Fortubev2.CallOpts, a, b, c)
}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Fortubev2 *Fortubev2Caller) MulScalar(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulScalar", a, scalar)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Fortubev2 *Fortubev2Session) MulScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalar(&_Fortubev2.CallOpts, a, scalar)
}

// MulScalar is a free data retrieval call binding the contract method 0x2985fa31.
//
// Solidity: function mulScalar(uint256 a, uint256 scalar) pure returns(uint256 scaled)
func (_Fortubev2 *Fortubev2CallerSession) MulScalar(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalar(&_Fortubev2.CallOpts, a, scalar)
}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) MulScalarTruncate(opts *bind.CallOpts, a *big.Int, scalar *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulScalarTruncate", a, scalar)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) MulScalarTruncate(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalarTruncate(&_Fortubev2.CallOpts, a, scalar)
}

// MulScalarTruncate is a free data retrieval call binding the contract method 0xb4ab15e7.
//
// Solidity: function mulScalarTruncate(uint256 a, uint256 scalar) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) MulScalarTruncate(a *big.Int, scalar *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalarTruncate(&_Fortubev2.CallOpts, a, scalar)
}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) MulScalarTruncateAddUInt(opts *bind.CallOpts, a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulScalarTruncateAddUInt", a, scalar, addend)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) MulScalarTruncateAddUInt(a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalarTruncateAddUInt(&_Fortubev2.CallOpts, a, scalar, addend)
}

// MulScalarTruncateAddUInt is a free data retrieval call binding the contract method 0x6208fc41.
//
// Solidity: function mulScalarTruncateAddUInt(uint256 a, uint256 scalar, uint256 addend) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) MulScalarTruncateAddUInt(a *big.Int, scalar *big.Int, addend *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.MulScalarTruncateAddUInt(&_Fortubev2.CallOpts, a, scalar, addend)
}

// Mulsig is a free data retrieval call binding the contract method 0x1a435b55.
//
// Solidity: function mulsig() view returns(address)
func (_Fortubev2 *Fortubev2Caller) Mulsig(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "mulsig")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Mulsig is a free data retrieval call binding the contract method 0x1a435b55.
//
// Solidity: function mulsig() view returns(address)
func (_Fortubev2 *Fortubev2Session) Mulsig() (common.Address, error) {
	return _Fortubev2.Contract.Mulsig(&_Fortubev2.CallOpts)
}

// Mulsig is a free data retrieval call binding the contract method 0x1a435b55.
//
// Solidity: function mulsig() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) Mulsig() (common.Address, error) {
	return _Fortubev2.Contract.Mulsig(&_Fortubev2.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Fortubev2 *Fortubev2Caller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Fortubev2 *Fortubev2Session) Oracle() (common.Address, error) {
	return _Fortubev2.Contract.Oracle(&_Fortubev2.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) Oracle() (common.Address, error) {
	return _Fortubev2.Contract.Oracle(&_Fortubev2.CallOpts)
}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Fortubev2 *Fortubev2Caller) ProposedAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "proposedAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Fortubev2 *Fortubev2Session) ProposedAdmin() (common.Address, error) {
	return _Fortubev2.Contract.ProposedAdmin(&_Fortubev2.CallOpts)
}

// ProposedAdmin is a free data retrieval call binding the contract method 0x32f751ec.
//
// Solidity: function proposedAdmin() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) ProposedAdmin() (common.Address, error) {
	return _Fortubev2.Contract.ProposedAdmin(&_Fortubev2.CallOpts)
}

// RepayCheck is a free data retrieval call binding the contract method 0x44f4b506.
//
// Solidity: function repayCheck(address underlying) view returns()
func (_Fortubev2 *Fortubev2Caller) RepayCheck(opts *bind.CallOpts, underlying common.Address) error {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "repayCheck", underlying)

	if err != nil {
		return err
	}

	return err

}

// RepayCheck is a free data retrieval call binding the contract method 0x44f4b506.
//
// Solidity: function repayCheck(address underlying) view returns()
func (_Fortubev2 *Fortubev2Session) RepayCheck(underlying common.Address) error {
	return _Fortubev2.Contract.RepayCheck(&_Fortubev2.CallOpts, underlying)
}

// RepayCheck is a free data retrieval call binding the contract method 0x44f4b506.
//
// Solidity: function repayCheck(address underlying) view returns()
func (_Fortubev2 *Fortubev2CallerSession) RepayCheck(underlying common.Address) error {
	return _Fortubev2.Contract.RepayCheck(&_Fortubev2.CallOpts, underlying)
}

// RewardFactors is a free data retrieval call binding the contract method 0xdd4b8b7d.
//
// Solidity: function rewardFactors(uint256 ) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) RewardFactors(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "rewardFactors", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RewardFactors is a free data retrieval call binding the contract method 0xdd4b8b7d.
//
// Solidity: function rewardFactors(uint256 ) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) RewardFactors(arg0 *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.RewardFactors(&_Fortubev2.CallOpts, arg0)
}

// RewardFactors is a free data retrieval call binding the contract method 0xdd4b8b7d.
//
// Solidity: function rewardFactors(uint256 ) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) RewardFactors(arg0 *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.RewardFactors(&_Fortubev2.CallOpts, arg0)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_Fortubev2 *Fortubev2Caller) RewardPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "rewardPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_Fortubev2 *Fortubev2Session) RewardPool() (common.Address, error) {
	return _Fortubev2.Contract.RewardPool(&_Fortubev2.CallOpts)
}

// RewardPool is a free data retrieval call binding the contract method 0x66666aa9.
//
// Solidity: function rewardPool() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) RewardPool() (common.Address, error) {
	return _Fortubev2.Contract.RewardPool(&_Fortubev2.CallOpts)
}

// SeizeCheck is a free data retrieval call binding the contract method 0x7ceb4c97.
//
// Solidity: function seizeCheck(address cTokenCollateral, address cTokenBorrowed) view returns()
func (_Fortubev2 *Fortubev2Caller) SeizeCheck(opts *bind.CallOpts, cTokenCollateral common.Address, cTokenBorrowed common.Address) error {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "seizeCheck", cTokenCollateral, cTokenBorrowed)

	if err != nil {
		return err
	}

	return err

}

// SeizeCheck is a free data retrieval call binding the contract method 0x7ceb4c97.
//
// Solidity: function seizeCheck(address cTokenCollateral, address cTokenBorrowed) view returns()
func (_Fortubev2 *Fortubev2Session) SeizeCheck(cTokenCollateral common.Address, cTokenBorrowed common.Address) error {
	return _Fortubev2.Contract.SeizeCheck(&_Fortubev2.CallOpts, cTokenCollateral, cTokenBorrowed)
}

// SeizeCheck is a free data retrieval call binding the contract method 0x7ceb4c97.
//
// Solidity: function seizeCheck(address cTokenCollateral, address cTokenBorrowed) view returns()
func (_Fortubev2 *Fortubev2CallerSession) SeizeCheck(cTokenCollateral common.Address, cTokenBorrowed common.Address) error {
	return _Fortubev2.Contract.SeizeCheck(&_Fortubev2.CallOpts, cTokenCollateral, cTokenBorrowed)
}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2Caller) SubExp(opts *bind.CallOpts, a *big.Int, b *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "subExp", a, b)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2Session) SubExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.SubExp(&_Fortubev2.CallOpts, a, b)
}

// SubExp is a free data retrieval call binding the contract method 0x396a98cf.
//
// Solidity: function subExp(uint256 a, uint256 b) pure returns(uint256 result)
func (_Fortubev2 *Fortubev2CallerSession) SubExp(a *big.Int, b *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.SubExp(&_Fortubev2.CallOpts, a, b)
}

// TheForceToken is a free data retrieval call binding the contract method 0x96a614d2.
//
// Solidity: function theForceToken() view returns(address)
func (_Fortubev2 *Fortubev2Caller) TheForceToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "theForceToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// TheForceToken is a free data retrieval call binding the contract method 0x96a614d2.
//
// Solidity: function theForceToken() view returns(address)
func (_Fortubev2 *Fortubev2Session) TheForceToken() (common.Address, error) {
	return _Fortubev2.Contract.TheForceToken(&_Fortubev2.CallOpts)
}

// TheForceToken is a free data retrieval call binding the contract method 0x96a614d2.
//
// Solidity: function theForceToken() view returns(address)
func (_Fortubev2 *Fortubev2CallerSession) TheForceToken() (common.Address, error) {
	return _Fortubev2.Contract.TheForceToken(&_Fortubev2.CallOpts)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) TokenDecimals(opts *bind.CallOpts, token common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "tokenDecimals", token)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address token) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) TokenDecimals(token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.TokenDecimals(&_Fortubev2.CallOpts, token)
}

// TokenDecimals is a free data retrieval call binding the contract method 0x8ee573ac.
//
// Solidity: function tokenDecimals(address token) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) TokenDecimals(token common.Address) (*big.Int, error) {
	return _Fortubev2.Contract.TokenDecimals(&_Fortubev2.CallOpts, token)
}

// TransferEthGasCost is a free data retrieval call binding the contract method 0x4f9c751e.
//
// Solidity: function transferEthGasCost() view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) TransferEthGasCost(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "transferEthGasCost")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TransferEthGasCost is a free data retrieval call binding the contract method 0x4f9c751e.
//
// Solidity: function transferEthGasCost() view returns(uint256)
func (_Fortubev2 *Fortubev2Session) TransferEthGasCost() (*big.Int, error) {
	return _Fortubev2.Contract.TransferEthGasCost(&_Fortubev2.CallOpts)
}

// TransferEthGasCost is a free data retrieval call binding the contract method 0x4f9c751e.
//
// Solidity: function transferEthGasCost() view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) TransferEthGasCost() (*big.Int, error) {
	return _Fortubev2.Contract.TransferEthGasCost(&_Fortubev2.CallOpts)
}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Fortubev2 *Fortubev2Caller) Truncate(opts *bind.CallOpts, exp *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "truncate", exp)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Fortubev2 *Fortubev2Session) Truncate(exp *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.Truncate(&_Fortubev2.CallOpts, exp)
}

// Truncate is a free data retrieval call binding the contract method 0x7b94aaac.
//
// Solidity: function truncate(uint256 exp) pure returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) Truncate(exp *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.Truncate(&_Fortubev2.CallOpts, exp)
}

// WithdrawCheck is a free data retrieval call binding the contract method 0x57986129.
//
// Solidity: function withdrawCheck(address fToken, address withdrawer, uint256 withdrawTokens) view returns(uint256)
func (_Fortubev2 *Fortubev2Caller) WithdrawCheck(opts *bind.CallOpts, fToken common.Address, withdrawer common.Address, withdrawTokens *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Fortubev2.contract.Call(opts, &out, "withdrawCheck", fToken, withdrawer, withdrawTokens)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// WithdrawCheck is a free data retrieval call binding the contract method 0x57986129.
//
// Solidity: function withdrawCheck(address fToken, address withdrawer, uint256 withdrawTokens) view returns(uint256)
func (_Fortubev2 *Fortubev2Session) WithdrawCheck(fToken common.Address, withdrawer common.Address, withdrawTokens *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.WithdrawCheck(&_Fortubev2.CallOpts, fToken, withdrawer, withdrawTokens)
}

// WithdrawCheck is a free data retrieval call binding the contract method 0x57986129.
//
// Solidity: function withdrawCheck(address fToken, address withdrawer, uint256 withdrawTokens) view returns(uint256)
func (_Fortubev2 *Fortubev2CallerSession) WithdrawCheck(fToken common.Address, withdrawer common.Address, withdrawTokens *big.Int) (*big.Int, error) {
	return _Fortubev2.Contract.WithdrawCheck(&_Fortubev2.CallOpts, fToken, withdrawer, withdrawTokens)
}

// SetCollateralAbility is a paid mutator transaction binding the contract method 0x4b973e9a.
//
// Solidity: function _setCollateralAbility(address underlying, uint256 newCollateralAbility) returns()
func (_Fortubev2 *Fortubev2Transactor) SetCollateralAbility(opts *bind.TransactOpts, underlying common.Address, newCollateralAbility *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "_setCollateralAbility", underlying, newCollateralAbility)
}

// SetCollateralAbility is a paid mutator transaction binding the contract method 0x4b973e9a.
//
// Solidity: function _setCollateralAbility(address underlying, uint256 newCollateralAbility) returns()
func (_Fortubev2 *Fortubev2Session) SetCollateralAbility(underlying common.Address, newCollateralAbility *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetCollateralAbility(&_Fortubev2.TransactOpts, underlying, newCollateralAbility)
}

// SetCollateralAbility is a paid mutator transaction binding the contract method 0x4b973e9a.
//
// Solidity: function _setCollateralAbility(address underlying, uint256 newCollateralAbility) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetCollateralAbility(underlying common.Address, newCollateralAbility *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetCollateralAbility(&_Fortubev2.TransactOpts, underlying, newCollateralAbility)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x903e679c.
//
// Solidity: function _setLiquidationIncentive(address underlying, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2Transactor) SetLiquidationIncentive(opts *bind.TransactOpts, underlying common.Address, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "_setLiquidationIncentive", underlying, _liquidationIncentive)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x903e679c.
//
// Solidity: function _setLiquidationIncentive(address underlying, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2Session) SetLiquidationIncentive(underlying common.Address, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetLiquidationIncentive(&_Fortubev2.TransactOpts, underlying, _liquidationIncentive)
}

// SetLiquidationIncentive is a paid mutator transaction binding the contract method 0x903e679c.
//
// Solidity: function _setLiquidationIncentive(address underlying, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetLiquidationIncentive(underlying common.Address, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetLiquidationIncentive(&_Fortubev2.TransactOpts, underlying, _liquidationIncentive)
}

// SupportMarket is a paid mutator transaction binding the contract method 0x3cfa4d04.
//
// Solidity: function _supportMarket(address fToken, uint256 _collateralAbility, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2Transactor) SupportMarket(opts *bind.TransactOpts, fToken common.Address, _collateralAbility *big.Int, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "_supportMarket", fToken, _collateralAbility, _liquidationIncentive)
}

// SupportMarket is a paid mutator transaction binding the contract method 0x3cfa4d04.
//
// Solidity: function _supportMarket(address fToken, uint256 _collateralAbility, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2Session) SupportMarket(fToken common.Address, _collateralAbility *big.Int, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SupportMarket(&_Fortubev2.TransactOpts, fToken, _collateralAbility, _liquidationIncentive)
}

// SupportMarket is a paid mutator transaction binding the contract method 0x3cfa4d04.
//
// Solidity: function _supportMarket(address fToken, uint256 _collateralAbility, uint256 _liquidationIncentive) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SupportMarket(fToken common.Address, _collateralAbility *big.Int, _liquidationIncentive *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SupportMarket(&_Fortubev2.TransactOpts, fToken, _collateralAbility, _liquidationIncentive)
}

// AddReserves is a paid mutator transaction binding the contract method 0xfc99bd9a.
//
// Solidity: function addReserves(address underlying, uint256 addAmount) payable returns()
func (_Fortubev2 *Fortubev2Transactor) AddReserves(opts *bind.TransactOpts, underlying common.Address, addAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "addReserves", underlying, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0xfc99bd9a.
//
// Solidity: function addReserves(address underlying, uint256 addAmount) payable returns()
func (_Fortubev2 *Fortubev2Session) AddReserves(underlying common.Address, addAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.AddReserves(&_Fortubev2.TransactOpts, underlying, addAmount)
}

// AddReserves is a paid mutator transaction binding the contract method 0xfc99bd9a.
//
// Solidity: function addReserves(address underlying, uint256 addAmount) payable returns()
func (_Fortubev2 *Fortubev2TransactorSession) AddReserves(underlying common.Address, addAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.AddReserves(&_Fortubev2.TransactOpts, underlying, addAmount)
}

// BorrowCheck is a paid mutator transaction binding the contract method 0x96a13b32.
//
// Solidity: function borrowCheck(address account, address underlying, address fToken, uint256 borrowAmount) returns()
func (_Fortubev2 *Fortubev2Transactor) BorrowCheck(opts *bind.TransactOpts, account common.Address, underlying common.Address, fToken common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "borrowCheck", account, underlying, fToken, borrowAmount)
}

// BorrowCheck is a paid mutator transaction binding the contract method 0x96a13b32.
//
// Solidity: function borrowCheck(address account, address underlying, address fToken, uint256 borrowAmount) returns()
func (_Fortubev2 *Fortubev2Session) BorrowCheck(account common.Address, underlying common.Address, fToken common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.BorrowCheck(&_Fortubev2.TransactOpts, account, underlying, fToken, borrowAmount)
}

// BorrowCheck is a paid mutator transaction binding the contract method 0x96a13b32.
//
// Solidity: function borrowCheck(address account, address underlying, address fToken, uint256 borrowAmount) returns()
func (_Fortubev2 *Fortubev2TransactorSession) BorrowCheck(account common.Address, underlying common.Address, fToken common.Address, borrowAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.BorrowCheck(&_Fortubev2.TransactOpts, account, underlying, fToken, borrowAmount)
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Fortubev2 *Fortubev2Transactor) ClaimAdministration(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "claimAdministration")
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Fortubev2 *Fortubev2Session) ClaimAdministration() (*types.Transaction, error) {
	return _Fortubev2.Contract.ClaimAdministration(&_Fortubev2.TransactOpts)
}

// ClaimAdministration is a paid mutator transaction binding the contract method 0xad2cb239.
//
// Solidity: function claimAdministration() returns()
func (_Fortubev2 *Fortubev2TransactorSession) ClaimAdministration() (*types.Transaction, error) {
	return _Fortubev2.Contract.ClaimAdministration(&_Fortubev2.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mulsig) returns()
func (_Fortubev2 *Fortubev2Transactor) Initialize(opts *bind.TransactOpts, _mulsig common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "initialize", _mulsig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mulsig) returns()
func (_Fortubev2 *Fortubev2Session) Initialize(_mulsig common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.Initialize(&_Fortubev2.TransactOpts, _mulsig)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address _mulsig) returns()
func (_Fortubev2 *Fortubev2TransactorSession) Initialize(_mulsig common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.Initialize(&_Fortubev2.TransactOpts, _mulsig)
}

// LiquidateBorrowCheck is a paid mutator transaction binding the contract method 0x7cb9ac42.
//
// Solidity: function liquidateBorrowCheck(address fTokenBorrowed, address fTokenCollateral, address borrower, address liquidator, uint256 repayAmount) returns()
func (_Fortubev2 *Fortubev2Transactor) LiquidateBorrowCheck(opts *bind.TransactOpts, fTokenBorrowed common.Address, fTokenCollateral common.Address, borrower common.Address, liquidator common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "liquidateBorrowCheck", fTokenBorrowed, fTokenCollateral, borrower, liquidator, repayAmount)
}

// LiquidateBorrowCheck is a paid mutator transaction binding the contract method 0x7cb9ac42.
//
// Solidity: function liquidateBorrowCheck(address fTokenBorrowed, address fTokenCollateral, address borrower, address liquidator, uint256 repayAmount) returns()
func (_Fortubev2 *Fortubev2Session) LiquidateBorrowCheck(fTokenBorrowed common.Address, fTokenCollateral common.Address, borrower common.Address, liquidator common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.LiquidateBorrowCheck(&_Fortubev2.TransactOpts, fTokenBorrowed, fTokenCollateral, borrower, liquidator, repayAmount)
}

// LiquidateBorrowCheck is a paid mutator transaction binding the contract method 0x7cb9ac42.
//
// Solidity: function liquidateBorrowCheck(address fTokenBorrowed, address fTokenCollateral, address borrower, address liquidator, uint256 repayAmount) returns()
func (_Fortubev2 *Fortubev2TransactorSession) LiquidateBorrowCheck(fTokenBorrowed common.Address, fTokenCollateral common.Address, borrower common.Address, liquidator common.Address, repayAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.LiquidateBorrowCheck(&_Fortubev2.TransactOpts, fTokenBorrowed, fTokenCollateral, borrower, liquidator, repayAmount)
}

// MintCheck is a paid mutator transaction binding the contract method 0x71fdbaf5.
//
// Solidity: function mintCheck(address underlying, address minter) returns()
func (_Fortubev2 *Fortubev2Transactor) MintCheck(opts *bind.TransactOpts, underlying common.Address, minter common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "mintCheck", underlying, minter)
}

// MintCheck is a paid mutator transaction binding the contract method 0x71fdbaf5.
//
// Solidity: function mintCheck(address underlying, address minter) returns()
func (_Fortubev2 *Fortubev2Session) MintCheck(underlying common.Address, minter common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.MintCheck(&_Fortubev2.TransactOpts, underlying, minter)
}

// MintCheck is a paid mutator transaction binding the contract method 0x71fdbaf5.
//
// Solidity: function mintCheck(address underlying, address minter) returns()
func (_Fortubev2 *Fortubev2TransactorSession) MintCheck(underlying common.Address, minter common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.MintCheck(&_Fortubev2.TransactOpts, underlying, minter)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Fortubev2 *Fortubev2Transactor) ProposeNewAdmin(opts *bind.TransactOpts, admin_ common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "proposeNewAdmin", admin_)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Fortubev2 *Fortubev2Session) ProposeNewAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.ProposeNewAdmin(&_Fortubev2.TransactOpts, admin_)
}

// ProposeNewAdmin is a paid mutator transaction binding the contract method 0xa6376746.
//
// Solidity: function proposeNewAdmin(address admin_) returns()
func (_Fortubev2 *Fortubev2TransactorSession) ProposeNewAdmin(admin_ common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.ProposeNewAdmin(&_Fortubev2.TransactOpts, admin_)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0xe47e65a3.
//
// Solidity: function reduceReserves(address underlying, address account, uint256 reduceAmount) returns()
func (_Fortubev2 *Fortubev2Transactor) ReduceReserves(opts *bind.TransactOpts, underlying common.Address, account common.Address, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "reduceReserves", underlying, account, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0xe47e65a3.
//
// Solidity: function reduceReserves(address underlying, address account, uint256 reduceAmount) returns()
func (_Fortubev2 *Fortubev2Session) ReduceReserves(underlying common.Address, account common.Address, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.ReduceReserves(&_Fortubev2.TransactOpts, underlying, account, reduceAmount)
}

// ReduceReserves is a paid mutator transaction binding the contract method 0xe47e65a3.
//
// Solidity: function reduceReserves(address underlying, address account, uint256 reduceAmount) returns()
func (_Fortubev2 *Fortubev2TransactorSession) ReduceReserves(underlying common.Address, account common.Address, reduceAmount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.ReduceReserves(&_Fortubev2.TransactOpts, underlying, account, reduceAmount)
}

// RewardForByType is a paid mutator transaction binding the contract method 0x7f5cd256.
//
// Solidity: function rewardForByType(address account, uint256 gasSpend, uint256 gasPrice, uint256 rewardType) returns()
func (_Fortubev2 *Fortubev2Transactor) RewardForByType(opts *bind.TransactOpts, account common.Address, gasSpend *big.Int, gasPrice *big.Int, rewardType *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "rewardForByType", account, gasSpend, gasPrice, rewardType)
}

// RewardForByType is a paid mutator transaction binding the contract method 0x7f5cd256.
//
// Solidity: function rewardForByType(address account, uint256 gasSpend, uint256 gasPrice, uint256 rewardType) returns()
func (_Fortubev2 *Fortubev2Session) RewardForByType(account common.Address, gasSpend *big.Int, gasPrice *big.Int, rewardType *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.RewardForByType(&_Fortubev2.TransactOpts, account, gasSpend, gasPrice, rewardType)
}

// RewardForByType is a paid mutator transaction binding the contract method 0x7f5cd256.
//
// Solidity: function rewardForByType(address account, uint256 gasSpend, uint256 gasPrice, uint256 rewardType) returns()
func (_Fortubev2 *Fortubev2TransactorSession) RewardForByType(account common.Address, gasSpend *big.Int, gasPrice *big.Int, rewardType *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.RewardForByType(&_Fortubev2.TransactOpts, account, gasSpend, gasPrice, rewardType)
}

// SetBankEntryAddress is a paid mutator transaction binding the contract method 0x9945b8b0.
//
// Solidity: function setBankEntryAddress(address _newBank) returns()
func (_Fortubev2 *Fortubev2Transactor) SetBankEntryAddress(opts *bind.TransactOpts, _newBank common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setBankEntryAddress", _newBank)
}

// SetBankEntryAddress is a paid mutator transaction binding the contract method 0x9945b8b0.
//
// Solidity: function setBankEntryAddress(address _newBank) returns()
func (_Fortubev2 *Fortubev2Session) SetBankEntryAddress(_newBank common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetBankEntryAddress(&_Fortubev2.TransactOpts, _newBank)
}

// SetBankEntryAddress is a paid mutator transaction binding the contract method 0x9945b8b0.
//
// Solidity: function setBankEntryAddress(address _newBank) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetBankEntryAddress(_newBank common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetBankEntryAddress(&_Fortubev2.TransactOpts, _newBank)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x12348e96.
//
// Solidity: function setCloseFactor(uint256 _closeFactor) returns()
func (_Fortubev2 *Fortubev2Transactor) SetCloseFactor(opts *bind.TransactOpts, _closeFactor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setCloseFactor", _closeFactor)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x12348e96.
//
// Solidity: function setCloseFactor(uint256 _closeFactor) returns()
func (_Fortubev2 *Fortubev2Session) SetCloseFactor(_closeFactor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetCloseFactor(&_Fortubev2.TransactOpts, _closeFactor)
}

// SetCloseFactor is a paid mutator transaction binding the contract method 0x12348e96.
//
// Solidity: function setCloseFactor(uint256 _closeFactor) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetCloseFactor(_closeFactor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetCloseFactor(&_Fortubev2.TransactOpts, _closeFactor)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Fortubev2 *Fortubev2Transactor) SetOracle(opts *bind.TransactOpts, _oracle common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setOracle", _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Fortubev2 *Fortubev2Session) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetOracle(&_Fortubev2.TransactOpts, _oracle)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address _oracle) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetOracle(_oracle common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetOracle(&_Fortubev2.TransactOpts, _oracle)
}

// SetRewardFactorByType is a paid mutator transaction binding the contract method 0xd3a32bb0.
//
// Solidity: function setRewardFactorByType(uint256 rewaradType, uint256 factor) returns()
func (_Fortubev2 *Fortubev2Transactor) SetRewardFactorByType(opts *bind.TransactOpts, rewaradType *big.Int, factor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setRewardFactorByType", rewaradType, factor)
}

// SetRewardFactorByType is a paid mutator transaction binding the contract method 0xd3a32bb0.
//
// Solidity: function setRewardFactorByType(uint256 rewaradType, uint256 factor) returns()
func (_Fortubev2 *Fortubev2Session) SetRewardFactorByType(rewaradType *big.Int, factor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetRewardFactorByType(&_Fortubev2.TransactOpts, rewaradType, factor)
}

// SetRewardFactorByType is a paid mutator transaction binding the contract method 0xd3a32bb0.
//
// Solidity: function setRewardFactorByType(uint256 rewaradType, uint256 factor) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetRewardFactorByType(rewaradType *big.Int, factor *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetRewardFactorByType(&_Fortubev2.TransactOpts, rewaradType, factor)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_Fortubev2 *Fortubev2Transactor) SetRewardPool(opts *bind.TransactOpts, _rewardPool common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setRewardPool", _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_Fortubev2 *Fortubev2Session) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetRewardPool(&_Fortubev2.TransactOpts, _rewardPool)
}

// SetRewardPool is a paid mutator transaction binding the contract method 0x78238c37.
//
// Solidity: function setRewardPool(address _rewardPool) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetRewardPool(_rewardPool common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetRewardPool(&_Fortubev2.TransactOpts, _rewardPool)
}

// SetTheForceToken is a paid mutator transaction binding the contract method 0x04edc6fb.
//
// Solidity: function setTheForceToken(address _theForceToken) returns()
func (_Fortubev2 *Fortubev2Transactor) SetTheForceToken(opts *bind.TransactOpts, _theForceToken common.Address) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setTheForceToken", _theForceToken)
}

// SetTheForceToken is a paid mutator transaction binding the contract method 0x04edc6fb.
//
// Solidity: function setTheForceToken(address _theForceToken) returns()
func (_Fortubev2 *Fortubev2Session) SetTheForceToken(_theForceToken common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetTheForceToken(&_Fortubev2.TransactOpts, _theForceToken)
}

// SetTheForceToken is a paid mutator transaction binding the contract method 0x04edc6fb.
//
// Solidity: function setTheForceToken(address _theForceToken) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetTheForceToken(_theForceToken common.Address) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetTheForceToken(&_Fortubev2.TransactOpts, _theForceToken)
}

// SetTransferEthGasCost is a paid mutator transaction binding the contract method 0xc0627004.
//
// Solidity: function setTransferEthGasCost(uint256 _transferEthGasCost) returns()
func (_Fortubev2 *Fortubev2Transactor) SetTransferEthGasCost(opts *bind.TransactOpts, _transferEthGasCost *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "setTransferEthGasCost", _transferEthGasCost)
}

// SetTransferEthGasCost is a paid mutator transaction binding the contract method 0xc0627004.
//
// Solidity: function setTransferEthGasCost(uint256 _transferEthGasCost) returns()
func (_Fortubev2 *Fortubev2Session) SetTransferEthGasCost(_transferEthGasCost *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetTransferEthGasCost(&_Fortubev2.TransactOpts, _transferEthGasCost)
}

// SetTransferEthGasCost is a paid mutator transaction binding the contract method 0xc0627004.
//
// Solidity: function setTransferEthGasCost(uint256 _transferEthGasCost) returns()
func (_Fortubev2 *Fortubev2TransactorSession) SetTransferEthGasCost(_transferEthGasCost *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.SetTransferEthGasCost(&_Fortubev2.TransactOpts, _transferEthGasCost)
}

// TransferCheck is a paid mutator transaction binding the contract method 0xd38120c9.
//
// Solidity: function transferCheck(address fToken, address src, address dst, uint256 transferTokens) returns()
func (_Fortubev2 *Fortubev2Transactor) TransferCheck(opts *bind.TransactOpts, fToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "transferCheck", fToken, src, dst, transferTokens)
}

// TransferCheck is a paid mutator transaction binding the contract method 0xd38120c9.
//
// Solidity: function transferCheck(address fToken, address src, address dst, uint256 transferTokens) returns()
func (_Fortubev2 *Fortubev2Session) TransferCheck(fToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferCheck(&_Fortubev2.TransactOpts, fToken, src, dst, transferTokens)
}

// TransferCheck is a paid mutator transaction binding the contract method 0xd38120c9.
//
// Solidity: function transferCheck(address fToken, address src, address dst, uint256 transferTokens) returns()
func (_Fortubev2 *Fortubev2TransactorSession) TransferCheck(fToken common.Address, src common.Address, dst common.Address, transferTokens *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferCheck(&_Fortubev2.TransactOpts, fToken, src, dst, transferTokens)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address account, address underlying, uint256 amount) payable returns()
func (_Fortubev2 *Fortubev2Transactor) TransferIn(opts *bind.TransactOpts, account common.Address, underlying common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "transferIn", account, underlying, amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address account, address underlying, uint256 amount) payable returns()
func (_Fortubev2 *Fortubev2Session) TransferIn(account common.Address, underlying common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferIn(&_Fortubev2.TransactOpts, account, underlying, amount)
}

// TransferIn is a paid mutator transaction binding the contract method 0xe4652f49.
//
// Solidity: function transferIn(address account, address underlying, uint256 amount) payable returns()
func (_Fortubev2 *Fortubev2TransactorSession) TransferIn(account common.Address, underlying common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferIn(&_Fortubev2.TransactOpts, account, underlying, amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address underlying, address account, uint256 amount) returns()
func (_Fortubev2 *Fortubev2Transactor) TransferToUser(opts *bind.TransactOpts, underlying common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.contract.Transact(opts, "transferToUser", underlying, account, amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address underlying, address account, uint256 amount) returns()
func (_Fortubev2 *Fortubev2Session) TransferToUser(underlying common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferToUser(&_Fortubev2.TransactOpts, underlying, account, amount)
}

// TransferToUser is a paid mutator transaction binding the contract method 0xfa93b2a5.
//
// Solidity: function transferToUser(address underlying, address account, uint256 amount) returns()
func (_Fortubev2 *Fortubev2TransactorSession) TransferToUser(underlying common.Address, account common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Fortubev2.Contract.TransferToUser(&_Fortubev2.TransactOpts, underlying, account, amount)
}

// Fortubev2AddTokenToMarketIterator is returned from FilterAddTokenToMarket and is used to iterate over the raw logs and unpacked data for AddTokenToMarket events raised by the Fortubev2 contract.
type Fortubev2AddTokenToMarketIterator struct {
	Event *Fortubev2AddTokenToMarket // Event containing the contract specifics and raw log

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
func (it *Fortubev2AddTokenToMarketIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Fortubev2AddTokenToMarket)
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
		it.Event = new(Fortubev2AddTokenToMarket)
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
func (it *Fortubev2AddTokenToMarketIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Fortubev2AddTokenToMarketIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Fortubev2AddTokenToMarket represents a AddTokenToMarket event raised by the Fortubev2 contract.
type Fortubev2AddTokenToMarket struct {
	Underlying common.Address
	FToken     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddTokenToMarket is a free log retrieval operation binding the contract event 0x3f7e38f5830b709d15de3d0a45066834f763a6fe34ef0e4a25dde26a8fa8399b.
//
// Solidity: event AddTokenToMarket(address underlying, address fToken)
func (_Fortubev2 *Fortubev2Filterer) FilterAddTokenToMarket(opts *bind.FilterOpts) (*Fortubev2AddTokenToMarketIterator, error) {

	logs, sub, err := _Fortubev2.contract.FilterLogs(opts, "AddTokenToMarket")
	if err != nil {
		return nil, err
	}
	return &Fortubev2AddTokenToMarketIterator{contract: _Fortubev2.contract, event: "AddTokenToMarket", logs: logs, sub: sub}, nil
}

// WatchAddTokenToMarket is a free log subscription operation binding the contract event 0x3f7e38f5830b709d15de3d0a45066834f763a6fe34ef0e4a25dde26a8fa8399b.
//
// Solidity: event AddTokenToMarket(address underlying, address fToken)
func (_Fortubev2 *Fortubev2Filterer) WatchAddTokenToMarket(opts *bind.WatchOpts, sink chan<- *Fortubev2AddTokenToMarket) (event.Subscription, error) {

	logs, sub, err := _Fortubev2.contract.WatchLogs(opts, "AddTokenToMarket")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Fortubev2AddTokenToMarket)
				if err := _Fortubev2.contract.UnpackLog(event, "AddTokenToMarket", log); err != nil {
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

// ParseAddTokenToMarket is a log parse operation binding the contract event 0x3f7e38f5830b709d15de3d0a45066834f763a6fe34ef0e4a25dde26a8fa8399b.
//
// Solidity: event AddTokenToMarket(address underlying, address fToken)
func (_Fortubev2 *Fortubev2Filterer) ParseAddTokenToMarket(log types.Log) (*Fortubev2AddTokenToMarket, error) {
	event := new(Fortubev2AddTokenToMarket)
	if err := _Fortubev2.contract.UnpackLog(event, "AddTokenToMarket", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
