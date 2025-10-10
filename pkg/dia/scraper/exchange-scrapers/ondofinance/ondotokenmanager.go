// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package ondotokenmanager

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

// IGMTokenManagerQuote is an auto generated low-level Go binding around an user-defined struct.
type IGMTokenManagerQuote struct {
	ChainId        *big.Int
	AttestationId  *big.Int
	UserId         [32]byte
	Asset          common.Address
	Price          *big.Int
	Quantity       *big.Int
	Expiration     *big.Int
	Side           uint8
	AdditionalData [32]byte
}

// OndotokenmanagerMetaData contains all meta data concerning the Ondotokenmanager contract.
var OndotokenmanagerMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_defaultAdmin\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_usdon\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_minimumDepositUSD\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"_minimumRedemptionUSD\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"AttestationAlreadyExecuted\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"AttestationExpired\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"DepositAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GMTokenMintsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GMTokenNotRegistered\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GMTokenRedemptionsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalMintsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"GlobalRedemptionsPaused\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IDRegistryAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidAttestationSigner\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidChainId\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"InvalidQuoteSide\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"IssuanceHoursAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RateLimiterAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"RedemptionAmountTooSmall\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"}],\"name\":\"SafeERC20FailedOperation\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"SanityCheckOracleAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"TokenAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"USDonAddressCantBeZero\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"USDonManagerNotEnabled\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"expected\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"actual\",\"type\":\"bytes32\"}],\"name\":\"UserIdMismatch\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"UserNotRegistered\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"recipientId\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"rwaToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rwaAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"AdminMint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"GMTokenMintingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"GMTokenMintingUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"GMTokenRedeemingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"GMTokenRedeemingUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"bool\",\"name\":\"registered\",\"type\":\"bool\"}],\"name\":\"GMTokenRegistered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"GlobalMintingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"GlobalMintingUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"GlobalRedeemingPaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"GlobalRedeemingUnpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldIssuanceHours\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newIssuanceHours\",\"type\":\"address\"}],\"name\":\"IssuanceHoursSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinDepositAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinDepositAmount\",\"type\":\"uint256\"}],\"name\":\"MinimumDepositAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"oldMinRedemptionAmount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"newMinRedemptionAmount\",\"type\":\"uint256\"}],\"name\":\"MinimumRedemptionAmountSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOndoIDRegistry\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOndoIDRegistry\",\"type\":\"address\"}],\"name\":\"OndoIDRegistrySet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOndoRateLimiter\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOndoRateLimiter\",\"type\":\"address\"}],\"name\":\"OndoRateLimiterSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldOndoSanityCheckOracle\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOndoSanityCheckOracle\",\"type\":\"address\"}],\"name\":\"OndoSanityCheckOracleSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"previousAdminRole\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"newAdminRole\",\"type\":\"bytes32\"}],\"name\":\"RoleAdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleGranted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"}],\"name\":\"RoleRevoked\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"executionId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"attestationId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"userId\",\"type\":\"bytes32\"},{\"indexed\":false,\"internalType\":\"enumIGMTokenManager.QuoteSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"}],\"name\":\"TradeExecuted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"oldUSDonManager\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newUSDonManager\",\"type\":\"address\"}],\"name\":\"USDonManagerSet\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"ADMIN_MINT_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ATTESTATION_SIGNER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CONFIGURER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"DEFAULT_ADMIN_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"GM_NORMALIZER\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"PAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"QUOTE_TYPEHASH\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UNPAUSER_ROLE\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"gmTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"metadata\",\"type\":\"bytes32\"}],\"name\":\"adminProcessMint\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"executedAttestationIds\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"executionId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDomainSeparator\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleAdmin\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"getRoleMember\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"}],\"name\":\"getRoleMemberCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalMintingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"globalRedeemingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"gmIdentifier\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gmTokenAccepted\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gmTokenMintingPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"gmTokenRedemptionsPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"grantRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"hasRole\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"issuanceHours\",\"outputs\":[{\"internalType\":\"contractIIssuanceHours\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumDepositUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumRedemptionUSD\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attestationId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"enumIGMTokenManager.QuoteSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"}],\"internalType\":\"structIGMTokenManager.Quote\",\"name\":\"quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"depositToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositTokenAmount\",\"type\":\"uint256\"}],\"name\":\"mintWithAttestation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ondoIDRegistry\",\"outputs\":[{\"internalType\":\"contractIOndoIDRegistry\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"ondoRateLimiter\",\"outputs\":[{\"internalType\":\"contractIOndoRateLimiter\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"pauseGMTokenMints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"pauseGMTokenRedemptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGlobalMints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"pauseGlobalRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"chainId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"attestationId\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"userId\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"quantity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"expiration\",\"type\":\"uint256\"},{\"internalType\":\"enumIGMTokenManager.QuoteSide\",\"name\":\"side\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"additionalData\",\"type\":\"bytes32\"}],\"internalType\":\"structIGMTokenManager.Quote\",\"name\":\"quote\",\"type\":\"tuple\"},{\"internalType\":\"bytes\",\"name\":\"signature\",\"type\":\"bytes\"},{\"internalType\":\"address\",\"name\":\"receiveToken\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"minimumReceiveAmount\",\"type\":\"uint256\"}],\"name\":\"redeemWithAttestation\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"renounceRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"retrieveTokens\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"role\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"revokeRole\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"sanityCheckOracle\",\"outputs\":[{\"internalType\":\"contractIOndoSanityCheckOracle\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"accepted\",\"type\":\"bool\"}],\"name\":\"setGMTokenRegistrationStatus\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_issuanceHours\",\"type\":\"address\"}],\"name\":\"setIssuanceHours\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumDepositUSD\",\"type\":\"uint256\"}],\"name\":\"setMinimumDepositAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_minimumRedemptionUSD\",\"type\":\"uint256\"}],\"name\":\"setMinimumRedemptionAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ondoIDRegistry\",\"type\":\"address\"}],\"name\":\"setOndoIDRegistry\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_ondoRateLimiter\",\"type\":\"address\"}],\"name\":\"setOndoRateLimiter\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_sanityCheckOracle\",\"type\":\"address\"}],\"name\":\"setSanityCheckOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_usdonManager\",\"type\":\"address\"}],\"name\":\"setUSDonManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"unpauseGMTokenMints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"gmToken\",\"type\":\"address\"}],\"name\":\"unpauseGMTokenRedemptions\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseGlobalMints\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"unpauseGlobalRedeems\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdon\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"usdonManager\",\"outputs\":[{\"internalType\":\"contractIUSDonManager\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// OndotokenmanagerABI is the input ABI used to generate the binding from.
// Deprecated: Use OndotokenmanagerMetaData.ABI instead.
var OndotokenmanagerABI = OndotokenmanagerMetaData.ABI

// Ondotokenmanager is an auto generated Go binding around an Ethereum contract.
type Ondotokenmanager struct {
	OndotokenmanagerCaller     // Read-only binding to the contract
	OndotokenmanagerTransactor // Write-only binding to the contract
	OndotokenmanagerFilterer   // Log filterer for contract events
}

// OndotokenmanagerCaller is an auto generated read-only Go binding around an Ethereum contract.
type OndotokenmanagerCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OndotokenmanagerTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OndotokenmanagerTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OndotokenmanagerFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OndotokenmanagerFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OndotokenmanagerSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OndotokenmanagerSession struct {
	Contract     *Ondotokenmanager // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OndotokenmanagerCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OndotokenmanagerCallerSession struct {
	Contract *OndotokenmanagerCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// OndotokenmanagerTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OndotokenmanagerTransactorSession struct {
	Contract     *OndotokenmanagerTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// OndotokenmanagerRaw is an auto generated low-level Go binding around an Ethereum contract.
type OndotokenmanagerRaw struct {
	Contract *Ondotokenmanager // Generic contract binding to access the raw methods on
}

// OndotokenmanagerCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OndotokenmanagerCallerRaw struct {
	Contract *OndotokenmanagerCaller // Generic read-only contract binding to access the raw methods on
}

// OndotokenmanagerTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OndotokenmanagerTransactorRaw struct {
	Contract *OndotokenmanagerTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOndotokenmanager creates a new instance of Ondotokenmanager, bound to a specific deployed contract.
func NewOndotokenmanager(address common.Address, backend bind.ContractBackend) (*Ondotokenmanager, error) {
	contract, err := bindOndotokenmanager(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ondotokenmanager{OndotokenmanagerCaller: OndotokenmanagerCaller{contract: contract}, OndotokenmanagerTransactor: OndotokenmanagerTransactor{contract: contract}, OndotokenmanagerFilterer: OndotokenmanagerFilterer{contract: contract}}, nil
}

// NewOndotokenmanagerCaller creates a new read-only instance of Ondotokenmanager, bound to a specific deployed contract.
func NewOndotokenmanagerCaller(address common.Address, caller bind.ContractCaller) (*OndotokenmanagerCaller, error) {
	contract, err := bindOndotokenmanager(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerCaller{contract: contract}, nil
}

// NewOndotokenmanagerTransactor creates a new write-only instance of Ondotokenmanager, bound to a specific deployed contract.
func NewOndotokenmanagerTransactor(address common.Address, transactor bind.ContractTransactor) (*OndotokenmanagerTransactor, error) {
	contract, err := bindOndotokenmanager(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerTransactor{contract: contract}, nil
}

// NewOndotokenmanagerFilterer creates a new log filterer instance of Ondotokenmanager, bound to a specific deployed contract.
func NewOndotokenmanagerFilterer(address common.Address, filterer bind.ContractFilterer) (*OndotokenmanagerFilterer, error) {
	contract, err := bindOndotokenmanager(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerFilterer{contract: contract}, nil
}

// bindOndotokenmanager binds a generic wrapper to an already deployed contract.
func bindOndotokenmanager(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OndotokenmanagerMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ondotokenmanager *OndotokenmanagerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ondotokenmanager.Contract.OndotokenmanagerCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ondotokenmanager *OndotokenmanagerRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.OndotokenmanagerTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ondotokenmanager *OndotokenmanagerRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.OndotokenmanagerTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ondotokenmanager *OndotokenmanagerCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Ondotokenmanager.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ondotokenmanager *OndotokenmanagerTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ondotokenmanager *OndotokenmanagerTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.contract.Transact(opts, method, params...)
}

// ADMINMINTROLE is a free data retrieval call binding the contract method 0x0e16f0c8.
//
// Solidity: function ADMIN_MINT_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) ADMINMINTROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "ADMIN_MINT_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ADMINMINTROLE is a free data retrieval call binding the contract method 0x0e16f0c8.
//
// Solidity: function ADMIN_MINT_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) ADMINMINTROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.ADMINMINTROLE(&_Ondotokenmanager.CallOpts)
}

// ADMINMINTROLE is a free data retrieval call binding the contract method 0x0e16f0c8.
//
// Solidity: function ADMIN_MINT_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) ADMINMINTROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.ADMINMINTROLE(&_Ondotokenmanager.CallOpts)
}

// ATTESTATIONSIGNERROLE is a free data retrieval call binding the contract method 0x9c617c7e.
//
// Solidity: function ATTESTATION_SIGNER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) ATTESTATIONSIGNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "ATTESTATION_SIGNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ATTESTATIONSIGNERROLE is a free data retrieval call binding the contract method 0x9c617c7e.
//
// Solidity: function ATTESTATION_SIGNER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) ATTESTATIONSIGNERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.ATTESTATIONSIGNERROLE(&_Ondotokenmanager.CallOpts)
}

// ATTESTATIONSIGNERROLE is a free data retrieval call binding the contract method 0x9c617c7e.
//
// Solidity: function ATTESTATION_SIGNER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) ATTESTATIONSIGNERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.ATTESTATIONSIGNERROLE(&_Ondotokenmanager.CallOpts)
}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) CONFIGURERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "CONFIGURER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) CONFIGURERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.CONFIGURERROLE(&_Ondotokenmanager.CallOpts)
}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) CONFIGURERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.CONFIGURERROLE(&_Ondotokenmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.DEFAULTADMINROLE(&_Ondotokenmanager.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.DEFAULTADMINROLE(&_Ondotokenmanager.CallOpts)
}

// GMNORMALIZER is a free data retrieval call binding the contract method 0xc7f8e171.
//
// Solidity: function GM_NORMALIZER() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCaller) GMNORMALIZER(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "GM_NORMALIZER")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GMNORMALIZER is a free data retrieval call binding the contract method 0xc7f8e171.
//
// Solidity: function GM_NORMALIZER() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) GMNORMALIZER() (*big.Int, error) {
	return _Ondotokenmanager.Contract.GMNORMALIZER(&_Ondotokenmanager.CallOpts)
}

// GMNORMALIZER is a free data retrieval call binding the contract method 0xc7f8e171.
//
// Solidity: function GM_NORMALIZER() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GMNORMALIZER() (*big.Int, error) {
	return _Ondotokenmanager.Contract.GMNORMALIZER(&_Ondotokenmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) PAUSERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.PAUSERROLE(&_Ondotokenmanager.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.PAUSERROLE(&_Ondotokenmanager.CallOpts)
}

// QUOTETYPEHASH is a free data retrieval call binding the contract method 0xf7581615.
//
// Solidity: function QUOTE_TYPEHASH() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) QUOTETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "QUOTE_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// QUOTETYPEHASH is a free data retrieval call binding the contract method 0xf7581615.
//
// Solidity: function QUOTE_TYPEHASH() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) QUOTETYPEHASH() ([32]byte, error) {
	return _Ondotokenmanager.Contract.QUOTETYPEHASH(&_Ondotokenmanager.CallOpts)
}

// QUOTETYPEHASH is a free data retrieval call binding the contract method 0xf7581615.
//
// Solidity: function QUOTE_TYPEHASH() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) QUOTETYPEHASH() ([32]byte, error) {
	return _Ondotokenmanager.Contract.QUOTETYPEHASH(&_Ondotokenmanager.CallOpts)
}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "UNPAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) UNPAUSERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.UNPAUSERROLE(&_Ondotokenmanager.CallOpts)
}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) UNPAUSERROLE() ([32]byte, error) {
	return _Ondotokenmanager.Contract.UNPAUSERROLE(&_Ondotokenmanager.CallOpts)
}

// ExecutedAttestationIds is a free data retrieval call binding the contract method 0x50aaef5a.
//
// Solidity: function executedAttestationIds(uint256 ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) ExecutedAttestationIds(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "executedAttestationIds", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ExecutedAttestationIds is a free data retrieval call binding the contract method 0x50aaef5a.
//
// Solidity: function executedAttestationIds(uint256 ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) ExecutedAttestationIds(arg0 *big.Int) (bool, error) {
	return _Ondotokenmanager.Contract.ExecutedAttestationIds(&_Ondotokenmanager.CallOpts, arg0)
}

// ExecutedAttestationIds is a free data retrieval call binding the contract method 0x50aaef5a.
//
// Solidity: function executedAttestationIds(uint256 ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) ExecutedAttestationIds(arg0 *big.Int) (bool, error) {
	return _Ondotokenmanager.Contract.ExecutedAttestationIds(&_Ondotokenmanager.CallOpts, arg0)
}

// ExecutionId is a free data retrieval call binding the contract method 0x76ce2dad.
//
// Solidity: function executionId() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCaller) ExecutionId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "executionId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ExecutionId is a free data retrieval call binding the contract method 0x76ce2dad.
//
// Solidity: function executionId() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) ExecutionId() (*big.Int, error) {
	return _Ondotokenmanager.Contract.ExecutionId(&_Ondotokenmanager.CallOpts)
}

// ExecutionId is a free data retrieval call binding the contract method 0x76ce2dad.
//
// Solidity: function executionId() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) ExecutionId() (*big.Int, error) {
	return _Ondotokenmanager.Contract.ExecutionId(&_Ondotokenmanager.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) GetDomainSeparator(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "getDomainSeparator")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) GetDomainSeparator() ([32]byte, error) {
	return _Ondotokenmanager.Contract.GetDomainSeparator(&_Ondotokenmanager.CallOpts)
}

// GetDomainSeparator is a free data retrieval call binding the contract method 0xed24911d.
//
// Solidity: function getDomainSeparator() view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GetDomainSeparator() ([32]byte, error) {
	return _Ondotokenmanager.Contract.GetDomainSeparator(&_Ondotokenmanager.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ondotokenmanager.Contract.GetRoleAdmin(&_Ondotokenmanager.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Ondotokenmanager.Contract.GetRoleAdmin(&_Ondotokenmanager.CallOpts, role)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) GetRoleMember(opts *bind.CallOpts, role [32]byte, index *big.Int) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "getRoleMember", role, index)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ondotokenmanager.Contract.GetRoleMember(&_Ondotokenmanager.CallOpts, role, index)
}

// GetRoleMember is a free data retrieval call binding the contract method 0x9010d07c.
//
// Solidity: function getRoleMember(bytes32 role, uint256 index) view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GetRoleMember(role [32]byte, index *big.Int) (common.Address, error) {
	return _Ondotokenmanager.Contract.GetRoleMember(&_Ondotokenmanager.CallOpts, role, index)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCaller) GetRoleMemberCount(opts *bind.CallOpts, role [32]byte) (*big.Int, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "getRoleMemberCount", role)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ondotokenmanager.Contract.GetRoleMemberCount(&_Ondotokenmanager.CallOpts, role)
}

// GetRoleMemberCount is a free data retrieval call binding the contract method 0xca15c873.
//
// Solidity: function getRoleMemberCount(bytes32 role) view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GetRoleMemberCount(role [32]byte) (*big.Int, error) {
	return _Ondotokenmanager.Contract.GetRoleMemberCount(&_Ondotokenmanager.CallOpts, role)
}

// GlobalMintingPaused is a free data retrieval call binding the contract method 0xeb27fa53.
//
// Solidity: function globalMintingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) GlobalMintingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "globalMintingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalMintingPaused is a free data retrieval call binding the contract method 0xeb27fa53.
//
// Solidity: function globalMintingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) GlobalMintingPaused() (bool, error) {
	return _Ondotokenmanager.Contract.GlobalMintingPaused(&_Ondotokenmanager.CallOpts)
}

// GlobalMintingPaused is a free data retrieval call binding the contract method 0xeb27fa53.
//
// Solidity: function globalMintingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GlobalMintingPaused() (bool, error) {
	return _Ondotokenmanager.Contract.GlobalMintingPaused(&_Ondotokenmanager.CallOpts)
}

// GlobalRedeemingPaused is a free data retrieval call binding the contract method 0x8c6f4ab7.
//
// Solidity: function globalRedeemingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) GlobalRedeemingPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "globalRedeemingPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GlobalRedeemingPaused is a free data retrieval call binding the contract method 0x8c6f4ab7.
//
// Solidity: function globalRedeemingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) GlobalRedeemingPaused() (bool, error) {
	return _Ondotokenmanager.Contract.GlobalRedeemingPaused(&_Ondotokenmanager.CallOpts)
}

// GlobalRedeemingPaused is a free data retrieval call binding the contract method 0x8c6f4ab7.
//
// Solidity: function globalRedeemingPaused() view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GlobalRedeemingPaused() (bool, error) {
	return _Ondotokenmanager.Contract.GlobalRedeemingPaused(&_Ondotokenmanager.CallOpts)
}

// GmIdentifier is a free data retrieval call binding the contract method 0x03cad645.
//
// Solidity: function gmIdentifier() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) GmIdentifier(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "gmIdentifier")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GmIdentifier is a free data retrieval call binding the contract method 0x03cad645.
//
// Solidity: function gmIdentifier() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) GmIdentifier() (common.Address, error) {
	return _Ondotokenmanager.Contract.GmIdentifier(&_Ondotokenmanager.CallOpts)
}

// GmIdentifier is a free data retrieval call binding the contract method 0x03cad645.
//
// Solidity: function gmIdentifier() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GmIdentifier() (common.Address, error) {
	return _Ondotokenmanager.Contract.GmIdentifier(&_Ondotokenmanager.CallOpts)
}

// GmTokenAccepted is a free data retrieval call binding the contract method 0x5c0478d1.
//
// Solidity: function gmTokenAccepted(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) GmTokenAccepted(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "gmTokenAccepted", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GmTokenAccepted is a free data retrieval call binding the contract method 0x5c0478d1.
//
// Solidity: function gmTokenAccepted(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) GmTokenAccepted(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenAccepted(&_Ondotokenmanager.CallOpts, arg0)
}

// GmTokenAccepted is a free data retrieval call binding the contract method 0x5c0478d1.
//
// Solidity: function gmTokenAccepted(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GmTokenAccepted(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenAccepted(&_Ondotokenmanager.CallOpts, arg0)
}

// GmTokenMintingPaused is a free data retrieval call binding the contract method 0x6269ff8f.
//
// Solidity: function gmTokenMintingPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) GmTokenMintingPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "gmTokenMintingPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GmTokenMintingPaused is a free data retrieval call binding the contract method 0x6269ff8f.
//
// Solidity: function gmTokenMintingPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) GmTokenMintingPaused(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenMintingPaused(&_Ondotokenmanager.CallOpts, arg0)
}

// GmTokenMintingPaused is a free data retrieval call binding the contract method 0x6269ff8f.
//
// Solidity: function gmTokenMintingPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GmTokenMintingPaused(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenMintingPaused(&_Ondotokenmanager.CallOpts, arg0)
}

// GmTokenRedemptionsPaused is a free data retrieval call binding the contract method 0xb28996b5.
//
// Solidity: function gmTokenRedemptionsPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) GmTokenRedemptionsPaused(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "gmTokenRedemptionsPaused", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// GmTokenRedemptionsPaused is a free data retrieval call binding the contract method 0xb28996b5.
//
// Solidity: function gmTokenRedemptionsPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) GmTokenRedemptionsPaused(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenRedemptionsPaused(&_Ondotokenmanager.CallOpts, arg0)
}

// GmTokenRedemptionsPaused is a free data retrieval call binding the contract method 0xb28996b5.
//
// Solidity: function gmTokenRedemptionsPaused(address ) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) GmTokenRedemptionsPaused(arg0 common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.GmTokenRedemptionsPaused(&_Ondotokenmanager.CallOpts, arg0)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.HasRole(&_Ondotokenmanager.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Ondotokenmanager.Contract.HasRole(&_Ondotokenmanager.CallOpts, role, account)
}

// IssuanceHours is a free data retrieval call binding the contract method 0x55358afb.
//
// Solidity: function issuanceHours() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) IssuanceHours(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "issuanceHours")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// IssuanceHours is a free data retrieval call binding the contract method 0x55358afb.
//
// Solidity: function issuanceHours() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) IssuanceHours() (common.Address, error) {
	return _Ondotokenmanager.Contract.IssuanceHours(&_Ondotokenmanager.CallOpts)
}

// IssuanceHours is a free data retrieval call binding the contract method 0x55358afb.
//
// Solidity: function issuanceHours() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) IssuanceHours() (common.Address, error) {
	return _Ondotokenmanager.Contract.IssuanceHours(&_Ondotokenmanager.CallOpts)
}

// MinimumDepositUSD is a free data retrieval call binding the contract method 0x531acc13.
//
// Solidity: function minimumDepositUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCaller) MinimumDepositUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "minimumDepositUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumDepositUSD is a free data retrieval call binding the contract method 0x531acc13.
//
// Solidity: function minimumDepositUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) MinimumDepositUSD() (*big.Int, error) {
	return _Ondotokenmanager.Contract.MinimumDepositUSD(&_Ondotokenmanager.CallOpts)
}

// MinimumDepositUSD is a free data retrieval call binding the contract method 0x531acc13.
//
// Solidity: function minimumDepositUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) MinimumDepositUSD() (*big.Int, error) {
	return _Ondotokenmanager.Contract.MinimumDepositUSD(&_Ondotokenmanager.CallOpts)
}

// MinimumRedemptionUSD is a free data retrieval call binding the contract method 0x8f8eb812.
//
// Solidity: function minimumRedemptionUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCaller) MinimumRedemptionUSD(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "minimumRedemptionUSD")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumRedemptionUSD is a free data retrieval call binding the contract method 0x8f8eb812.
//
// Solidity: function minimumRedemptionUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) MinimumRedemptionUSD() (*big.Int, error) {
	return _Ondotokenmanager.Contract.MinimumRedemptionUSD(&_Ondotokenmanager.CallOpts)
}

// MinimumRedemptionUSD is a free data retrieval call binding the contract method 0x8f8eb812.
//
// Solidity: function minimumRedemptionUSD() view returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) MinimumRedemptionUSD() (*big.Int, error) {
	return _Ondotokenmanager.Contract.MinimumRedemptionUSD(&_Ondotokenmanager.CallOpts)
}

// OndoIDRegistry is a free data retrieval call binding the contract method 0xc768b145.
//
// Solidity: function ondoIDRegistry() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) OndoIDRegistry(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "ondoIDRegistry")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OndoIDRegistry is a free data retrieval call binding the contract method 0xc768b145.
//
// Solidity: function ondoIDRegistry() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) OndoIDRegistry() (common.Address, error) {
	return _Ondotokenmanager.Contract.OndoIDRegistry(&_Ondotokenmanager.CallOpts)
}

// OndoIDRegistry is a free data retrieval call binding the contract method 0xc768b145.
//
// Solidity: function ondoIDRegistry() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) OndoIDRegistry() (common.Address, error) {
	return _Ondotokenmanager.Contract.OndoIDRegistry(&_Ondotokenmanager.CallOpts)
}

// OndoRateLimiter is a free data retrieval call binding the contract method 0x811400fa.
//
// Solidity: function ondoRateLimiter() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) OndoRateLimiter(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "ondoRateLimiter")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OndoRateLimiter is a free data retrieval call binding the contract method 0x811400fa.
//
// Solidity: function ondoRateLimiter() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) OndoRateLimiter() (common.Address, error) {
	return _Ondotokenmanager.Contract.OndoRateLimiter(&_Ondotokenmanager.CallOpts)
}

// OndoRateLimiter is a free data retrieval call binding the contract method 0x811400fa.
//
// Solidity: function ondoRateLimiter() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) OndoRateLimiter() (common.Address, error) {
	return _Ondotokenmanager.Contract.OndoRateLimiter(&_Ondotokenmanager.CallOpts)
}

// SanityCheckOracle is a free data retrieval call binding the contract method 0x5010af4d.
//
// Solidity: function sanityCheckOracle() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) SanityCheckOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "sanityCheckOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SanityCheckOracle is a free data retrieval call binding the contract method 0x5010af4d.
//
// Solidity: function sanityCheckOracle() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) SanityCheckOracle() (common.Address, error) {
	return _Ondotokenmanager.Contract.SanityCheckOracle(&_Ondotokenmanager.CallOpts)
}

// SanityCheckOracle is a free data retrieval call binding the contract method 0x5010af4d.
//
// Solidity: function sanityCheckOracle() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) SanityCheckOracle() (common.Address, error) {
	return _Ondotokenmanager.Contract.SanityCheckOracle(&_Ondotokenmanager.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ondotokenmanager.Contract.SupportsInterface(&_Ondotokenmanager.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Ondotokenmanager.Contract.SupportsInterface(&_Ondotokenmanager.CallOpts, interfaceId)
}

// Usdon is a free data retrieval call binding the contract method 0xbd5d370c.
//
// Solidity: function usdon() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) Usdon(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "usdon")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Usdon is a free data retrieval call binding the contract method 0xbd5d370c.
//
// Solidity: function usdon() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) Usdon() (common.Address, error) {
	return _Ondotokenmanager.Contract.Usdon(&_Ondotokenmanager.CallOpts)
}

// Usdon is a free data retrieval call binding the contract method 0xbd5d370c.
//
// Solidity: function usdon() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) Usdon() (common.Address, error) {
	return _Ondotokenmanager.Contract.Usdon(&_Ondotokenmanager.CallOpts)
}

// UsdonManager is a free data retrieval call binding the contract method 0x4f23c343.
//
// Solidity: function usdonManager() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCaller) UsdonManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Ondotokenmanager.contract.Call(opts, &out, "usdonManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// UsdonManager is a free data retrieval call binding the contract method 0x4f23c343.
//
// Solidity: function usdonManager() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerSession) UsdonManager() (common.Address, error) {
	return _Ondotokenmanager.Contract.UsdonManager(&_Ondotokenmanager.CallOpts)
}

// UsdonManager is a free data retrieval call binding the contract method 0x4f23c343.
//
// Solidity: function usdonManager() view returns(address)
func (_Ondotokenmanager *OndotokenmanagerCallerSession) UsdonManager() (common.Address, error) {
	return _Ondotokenmanager.Contract.UsdonManager(&_Ondotokenmanager.CallOpts)
}

// AdminProcessMint is a paid mutator transaction binding the contract method 0xed2a3540.
//
// Solidity: function adminProcessMint(address gmToken, address recipient, uint256 gmTokenAmount, bytes32 metadata) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) AdminProcessMint(opts *bind.TransactOpts, gmToken common.Address, recipient common.Address, gmTokenAmount *big.Int, metadata [32]byte) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "adminProcessMint", gmToken, recipient, gmTokenAmount, metadata)
}

// AdminProcessMint is a paid mutator transaction binding the contract method 0xed2a3540.
//
// Solidity: function adminProcessMint(address gmToken, address recipient, uint256 gmTokenAmount, bytes32 metadata) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) AdminProcessMint(gmToken common.Address, recipient common.Address, gmTokenAmount *big.Int, metadata [32]byte) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.AdminProcessMint(&_Ondotokenmanager.TransactOpts, gmToken, recipient, gmTokenAmount, metadata)
}

// AdminProcessMint is a paid mutator transaction binding the contract method 0xed2a3540.
//
// Solidity: function adminProcessMint(address gmToken, address recipient, uint256 gmTokenAmount, bytes32 metadata) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) AdminProcessMint(gmToken common.Address, recipient common.Address, gmTokenAmount *big.Int, metadata [32]byte) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.AdminProcessMint(&_Ondotokenmanager.TransactOpts, gmToken, recipient, gmTokenAmount, metadata)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.GrantRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.GrantRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// MintWithAttestation is a paid mutator transaction binding the contract method 0x445df08b.
//
// Solidity: function mintWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address depositToken, uint256 depositTokenAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerTransactor) MintWithAttestation(opts *bind.TransactOpts, quote IGMTokenManagerQuote, signature []byte, depositToken common.Address, depositTokenAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "mintWithAttestation", quote, signature, depositToken, depositTokenAmount)
}

// MintWithAttestation is a paid mutator transaction binding the contract method 0x445df08b.
//
// Solidity: function mintWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address depositToken, uint256 depositTokenAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) MintWithAttestation(quote IGMTokenManagerQuote, signature []byte, depositToken common.Address, depositTokenAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.MintWithAttestation(&_Ondotokenmanager.TransactOpts, quote, signature, depositToken, depositTokenAmount)
}

// MintWithAttestation is a paid mutator transaction binding the contract method 0x445df08b.
//
// Solidity: function mintWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address depositToken, uint256 depositTokenAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) MintWithAttestation(quote IGMTokenManagerQuote, signature []byte, depositToken common.Address, depositTokenAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.MintWithAttestation(&_Ondotokenmanager.TransactOpts, quote, signature, depositToken, depositTokenAmount)
}

// PauseGMTokenMints is a paid mutator transaction binding the contract method 0xc875edf6.
//
// Solidity: function pauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) PauseGMTokenMints(opts *bind.TransactOpts, gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "pauseGMTokenMints", gmToken)
}

// PauseGMTokenMints is a paid mutator transaction binding the contract method 0xc875edf6.
//
// Solidity: function pauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) PauseGMTokenMints(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGMTokenMints(&_Ondotokenmanager.TransactOpts, gmToken)
}

// PauseGMTokenMints is a paid mutator transaction binding the contract method 0xc875edf6.
//
// Solidity: function pauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) PauseGMTokenMints(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGMTokenMints(&_Ondotokenmanager.TransactOpts, gmToken)
}

// PauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x2d50bb74.
//
// Solidity: function pauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) PauseGMTokenRedemptions(opts *bind.TransactOpts, gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "pauseGMTokenRedemptions", gmToken)
}

// PauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x2d50bb74.
//
// Solidity: function pauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) PauseGMTokenRedemptions(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGMTokenRedemptions(&_Ondotokenmanager.TransactOpts, gmToken)
}

// PauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x2d50bb74.
//
// Solidity: function pauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) PauseGMTokenRedemptions(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGMTokenRedemptions(&_Ondotokenmanager.TransactOpts, gmToken)
}

// PauseGlobalMints is a paid mutator transaction binding the contract method 0x2542dbfd.
//
// Solidity: function pauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) PauseGlobalMints(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "pauseGlobalMints")
}

// PauseGlobalMints is a paid mutator transaction binding the contract method 0x2542dbfd.
//
// Solidity: function pauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerSession) PauseGlobalMints() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGlobalMints(&_Ondotokenmanager.TransactOpts)
}

// PauseGlobalMints is a paid mutator transaction binding the contract method 0x2542dbfd.
//
// Solidity: function pauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) PauseGlobalMints() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGlobalMints(&_Ondotokenmanager.TransactOpts)
}

// PauseGlobalRedeems is a paid mutator transaction binding the contract method 0x69d3605c.
//
// Solidity: function pauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) PauseGlobalRedeems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "pauseGlobalRedeems")
}

// PauseGlobalRedeems is a paid mutator transaction binding the contract method 0x69d3605c.
//
// Solidity: function pauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerSession) PauseGlobalRedeems() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGlobalRedeems(&_Ondotokenmanager.TransactOpts)
}

// PauseGlobalRedeems is a paid mutator transaction binding the contract method 0x69d3605c.
//
// Solidity: function pauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) PauseGlobalRedeems() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.PauseGlobalRedeems(&_Ondotokenmanager.TransactOpts)
}

// RedeemWithAttestation is a paid mutator transaction binding the contract method 0x6eefea49.
//
// Solidity: function redeemWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address receiveToken, uint256 minimumReceiveAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerTransactor) RedeemWithAttestation(opts *bind.TransactOpts, quote IGMTokenManagerQuote, signature []byte, receiveToken common.Address, minimumReceiveAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "redeemWithAttestation", quote, signature, receiveToken, minimumReceiveAmount)
}

// RedeemWithAttestation is a paid mutator transaction binding the contract method 0x6eefea49.
//
// Solidity: function redeemWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address receiveToken, uint256 minimumReceiveAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerSession) RedeemWithAttestation(quote IGMTokenManagerQuote, signature []byte, receiveToken common.Address, minimumReceiveAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RedeemWithAttestation(&_Ondotokenmanager.TransactOpts, quote, signature, receiveToken, minimumReceiveAmount)
}

// RedeemWithAttestation is a paid mutator transaction binding the contract method 0x6eefea49.
//
// Solidity: function redeemWithAttestation((uint256,uint256,bytes32,address,uint256,uint256,uint256,uint8,bytes32) quote, bytes signature, address receiveToken, uint256 minimumReceiveAmount) returns(uint256)
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) RedeemWithAttestation(quote IGMTokenManagerQuote, signature []byte, receiveToken common.Address, minimumReceiveAmount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RedeemWithAttestation(&_Ondotokenmanager.TransactOpts, quote, signature, receiveToken, minimumReceiveAmount)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "renounceRole", role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RenounceRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) RenounceRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RenounceRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// RetrieveTokens is a paid mutator transaction binding the contract method 0x7b0e1c57.
//
// Solidity: function retrieveTokens(address token, address to, uint256 amount) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) RetrieveTokens(opts *bind.TransactOpts, token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "retrieveTokens", token, to, amount)
}

// RetrieveTokens is a paid mutator transaction binding the contract method 0x7b0e1c57.
//
// Solidity: function retrieveTokens(address token, address to, uint256 amount) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) RetrieveTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RetrieveTokens(&_Ondotokenmanager.TransactOpts, token, to, amount)
}

// RetrieveTokens is a paid mutator transaction binding the contract method 0x7b0e1c57.
//
// Solidity: function retrieveTokens(address token, address to, uint256 amount) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) RetrieveTokens(token common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RetrieveTokens(&_Ondotokenmanager.TransactOpts, token, to, amount)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RevokeRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.RevokeRole(&_Ondotokenmanager.TransactOpts, role, account)
}

// SetGMTokenRegistrationStatus is a paid mutator transaction binding the contract method 0x246b3502.
//
// Solidity: function setGMTokenRegistrationStatus(address token, bool accepted) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetGMTokenRegistrationStatus(opts *bind.TransactOpts, token common.Address, accepted bool) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setGMTokenRegistrationStatus", token, accepted)
}

// SetGMTokenRegistrationStatus is a paid mutator transaction binding the contract method 0x246b3502.
//
// Solidity: function setGMTokenRegistrationStatus(address token, bool accepted) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetGMTokenRegistrationStatus(token common.Address, accepted bool) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetGMTokenRegistrationStatus(&_Ondotokenmanager.TransactOpts, token, accepted)
}

// SetGMTokenRegistrationStatus is a paid mutator transaction binding the contract method 0x246b3502.
//
// Solidity: function setGMTokenRegistrationStatus(address token, bool accepted) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetGMTokenRegistrationStatus(token common.Address, accepted bool) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetGMTokenRegistrationStatus(&_Ondotokenmanager.TransactOpts, token, accepted)
}

// SetIssuanceHours is a paid mutator transaction binding the contract method 0xebaf3f21.
//
// Solidity: function setIssuanceHours(address _issuanceHours) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetIssuanceHours(opts *bind.TransactOpts, _issuanceHours common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setIssuanceHours", _issuanceHours)
}

// SetIssuanceHours is a paid mutator transaction binding the contract method 0xebaf3f21.
//
// Solidity: function setIssuanceHours(address _issuanceHours) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetIssuanceHours(_issuanceHours common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetIssuanceHours(&_Ondotokenmanager.TransactOpts, _issuanceHours)
}

// SetIssuanceHours is a paid mutator transaction binding the contract method 0xebaf3f21.
//
// Solidity: function setIssuanceHours(address _issuanceHours) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetIssuanceHours(_issuanceHours common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetIssuanceHours(&_Ondotokenmanager.TransactOpts, _issuanceHours)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 _minimumDepositUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetMinimumDepositAmount(opts *bind.TransactOpts, _minimumDepositUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setMinimumDepositAmount", _minimumDepositUSD)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 _minimumDepositUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetMinimumDepositAmount(_minimumDepositUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetMinimumDepositAmount(&_Ondotokenmanager.TransactOpts, _minimumDepositUSD)
}

// SetMinimumDepositAmount is a paid mutator transaction binding the contract method 0xaab483d6.
//
// Solidity: function setMinimumDepositAmount(uint256 _minimumDepositUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetMinimumDepositAmount(_minimumDepositUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetMinimumDepositAmount(&_Ondotokenmanager.TransactOpts, _minimumDepositUSD)
}

// SetMinimumRedemptionAmount is a paid mutator transaction binding the contract method 0x4ef1ccd1.
//
// Solidity: function setMinimumRedemptionAmount(uint256 _minimumRedemptionUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetMinimumRedemptionAmount(opts *bind.TransactOpts, _minimumRedemptionUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setMinimumRedemptionAmount", _minimumRedemptionUSD)
}

// SetMinimumRedemptionAmount is a paid mutator transaction binding the contract method 0x4ef1ccd1.
//
// Solidity: function setMinimumRedemptionAmount(uint256 _minimumRedemptionUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetMinimumRedemptionAmount(_minimumRedemptionUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetMinimumRedemptionAmount(&_Ondotokenmanager.TransactOpts, _minimumRedemptionUSD)
}

// SetMinimumRedemptionAmount is a paid mutator transaction binding the contract method 0x4ef1ccd1.
//
// Solidity: function setMinimumRedemptionAmount(uint256 _minimumRedemptionUSD) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetMinimumRedemptionAmount(_minimumRedemptionUSD *big.Int) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetMinimumRedemptionAmount(&_Ondotokenmanager.TransactOpts, _minimumRedemptionUSD)
}

// SetOndoIDRegistry is a paid mutator transaction binding the contract method 0x23991e4b.
//
// Solidity: function setOndoIDRegistry(address _ondoIDRegistry) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetOndoIDRegistry(opts *bind.TransactOpts, _ondoIDRegistry common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setOndoIDRegistry", _ondoIDRegistry)
}

// SetOndoIDRegistry is a paid mutator transaction binding the contract method 0x23991e4b.
//
// Solidity: function setOndoIDRegistry(address _ondoIDRegistry) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetOndoIDRegistry(_ondoIDRegistry common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetOndoIDRegistry(&_Ondotokenmanager.TransactOpts, _ondoIDRegistry)
}

// SetOndoIDRegistry is a paid mutator transaction binding the contract method 0x23991e4b.
//
// Solidity: function setOndoIDRegistry(address _ondoIDRegistry) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetOndoIDRegistry(_ondoIDRegistry common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetOndoIDRegistry(&_Ondotokenmanager.TransactOpts, _ondoIDRegistry)
}

// SetOndoRateLimiter is a paid mutator transaction binding the contract method 0xad129e77.
//
// Solidity: function setOndoRateLimiter(address _ondoRateLimiter) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetOndoRateLimiter(opts *bind.TransactOpts, _ondoRateLimiter common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setOndoRateLimiter", _ondoRateLimiter)
}

// SetOndoRateLimiter is a paid mutator transaction binding the contract method 0xad129e77.
//
// Solidity: function setOndoRateLimiter(address _ondoRateLimiter) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetOndoRateLimiter(_ondoRateLimiter common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetOndoRateLimiter(&_Ondotokenmanager.TransactOpts, _ondoRateLimiter)
}

// SetOndoRateLimiter is a paid mutator transaction binding the contract method 0xad129e77.
//
// Solidity: function setOndoRateLimiter(address _ondoRateLimiter) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetOndoRateLimiter(_ondoRateLimiter common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetOndoRateLimiter(&_Ondotokenmanager.TransactOpts, _ondoRateLimiter)
}

// SetSanityCheckOracle is a paid mutator transaction binding the contract method 0x84b78414.
//
// Solidity: function setSanityCheckOracle(address _sanityCheckOracle) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetSanityCheckOracle(opts *bind.TransactOpts, _sanityCheckOracle common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setSanityCheckOracle", _sanityCheckOracle)
}

// SetSanityCheckOracle is a paid mutator transaction binding the contract method 0x84b78414.
//
// Solidity: function setSanityCheckOracle(address _sanityCheckOracle) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetSanityCheckOracle(_sanityCheckOracle common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetSanityCheckOracle(&_Ondotokenmanager.TransactOpts, _sanityCheckOracle)
}

// SetSanityCheckOracle is a paid mutator transaction binding the contract method 0x84b78414.
//
// Solidity: function setSanityCheckOracle(address _sanityCheckOracle) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetSanityCheckOracle(_sanityCheckOracle common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetSanityCheckOracle(&_Ondotokenmanager.TransactOpts, _sanityCheckOracle)
}

// SetUSDonManager is a paid mutator transaction binding the contract method 0x3483369d.
//
// Solidity: function setUSDonManager(address _usdonManager) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) SetUSDonManager(opts *bind.TransactOpts, _usdonManager common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "setUSDonManager", _usdonManager)
}

// SetUSDonManager is a paid mutator transaction binding the contract method 0x3483369d.
//
// Solidity: function setUSDonManager(address _usdonManager) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) SetUSDonManager(_usdonManager common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetUSDonManager(&_Ondotokenmanager.TransactOpts, _usdonManager)
}

// SetUSDonManager is a paid mutator transaction binding the contract method 0x3483369d.
//
// Solidity: function setUSDonManager(address _usdonManager) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) SetUSDonManager(_usdonManager common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.SetUSDonManager(&_Ondotokenmanager.TransactOpts, _usdonManager)
}

// UnpauseGMTokenMints is a paid mutator transaction binding the contract method 0x48c6ca77.
//
// Solidity: function unpauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) UnpauseGMTokenMints(opts *bind.TransactOpts, gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "unpauseGMTokenMints", gmToken)
}

// UnpauseGMTokenMints is a paid mutator transaction binding the contract method 0x48c6ca77.
//
// Solidity: function unpauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) UnpauseGMTokenMints(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGMTokenMints(&_Ondotokenmanager.TransactOpts, gmToken)
}

// UnpauseGMTokenMints is a paid mutator transaction binding the contract method 0x48c6ca77.
//
// Solidity: function unpauseGMTokenMints(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) UnpauseGMTokenMints(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGMTokenMints(&_Ondotokenmanager.TransactOpts, gmToken)
}

// UnpauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x05ac4e0b.
//
// Solidity: function unpauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) UnpauseGMTokenRedemptions(opts *bind.TransactOpts, gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "unpauseGMTokenRedemptions", gmToken)
}

// UnpauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x05ac4e0b.
//
// Solidity: function unpauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerSession) UnpauseGMTokenRedemptions(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGMTokenRedemptions(&_Ondotokenmanager.TransactOpts, gmToken)
}

// UnpauseGMTokenRedemptions is a paid mutator transaction binding the contract method 0x05ac4e0b.
//
// Solidity: function unpauseGMTokenRedemptions(address gmToken) returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) UnpauseGMTokenRedemptions(gmToken common.Address) (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGMTokenRedemptions(&_Ondotokenmanager.TransactOpts, gmToken)
}

// UnpauseGlobalMints is a paid mutator transaction binding the contract method 0x81d69bba.
//
// Solidity: function unpauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) UnpauseGlobalMints(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "unpauseGlobalMints")
}

// UnpauseGlobalMints is a paid mutator transaction binding the contract method 0x81d69bba.
//
// Solidity: function unpauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerSession) UnpauseGlobalMints() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGlobalMints(&_Ondotokenmanager.TransactOpts)
}

// UnpauseGlobalMints is a paid mutator transaction binding the contract method 0x81d69bba.
//
// Solidity: function unpauseGlobalMints() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) UnpauseGlobalMints() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGlobalMints(&_Ondotokenmanager.TransactOpts)
}

// UnpauseGlobalRedeems is a paid mutator transaction binding the contract method 0x0250f9fe.
//
// Solidity: function unpauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactor) UnpauseGlobalRedeems(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ondotokenmanager.contract.Transact(opts, "unpauseGlobalRedeems")
}

// UnpauseGlobalRedeems is a paid mutator transaction binding the contract method 0x0250f9fe.
//
// Solidity: function unpauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerSession) UnpauseGlobalRedeems() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGlobalRedeems(&_Ondotokenmanager.TransactOpts)
}

// UnpauseGlobalRedeems is a paid mutator transaction binding the contract method 0x0250f9fe.
//
// Solidity: function unpauseGlobalRedeems() returns()
func (_Ondotokenmanager *OndotokenmanagerTransactorSession) UnpauseGlobalRedeems() (*types.Transaction, error) {
	return _Ondotokenmanager.Contract.UnpauseGlobalRedeems(&_Ondotokenmanager.TransactOpts)
}

// OndotokenmanagerAdminMintIterator is returned from FilterAdminMint and is used to iterate over the raw logs and unpacked data for AdminMint events raised by the Ondotokenmanager contract.
type OndotokenmanagerAdminMintIterator struct {
	Event *OndotokenmanagerAdminMint // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerAdminMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerAdminMint)
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
		it.Event = new(OndotokenmanagerAdminMint)
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
func (it *OndotokenmanagerAdminMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerAdminMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerAdminMint represents a AdminMint event raised by the Ondotokenmanager contract.
type OndotokenmanagerAdminMint struct {
	Recipient   common.Address
	RecipientId [32]byte
	RwaToken    common.Address
	RwaAmount   *big.Int
	Metadata    [32]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAdminMint is a free log retrieval operation binding the contract event 0x64246d088e96d6b1bbb49bd2ab458376867796a99add0e25a88c53bcee7ddbdf.
//
// Solidity: event AdminMint(address indexed recipient, bytes32 indexed recipientId, address indexed rwaToken, uint256 rwaAmount, bytes32 metadata)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterAdminMint(opts *bind.FilterOpts, recipient []common.Address, recipientId [][32]byte, rwaToken []common.Address) (*OndotokenmanagerAdminMintIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var recipientIdRule []interface{}
	for _, recipientIdItem := range recipientId {
		recipientIdRule = append(recipientIdRule, recipientIdItem)
	}
	var rwaTokenRule []interface{}
	for _, rwaTokenItem := range rwaToken {
		rwaTokenRule = append(rwaTokenRule, rwaTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "AdminMint", recipientRule, recipientIdRule, rwaTokenRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerAdminMintIterator{contract: _Ondotokenmanager.contract, event: "AdminMint", logs: logs, sub: sub}, nil
}

// WatchAdminMint is a free log subscription operation binding the contract event 0x64246d088e96d6b1bbb49bd2ab458376867796a99add0e25a88c53bcee7ddbdf.
//
// Solidity: event AdminMint(address indexed recipient, bytes32 indexed recipientId, address indexed rwaToken, uint256 rwaAmount, bytes32 metadata)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchAdminMint(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerAdminMint, recipient []common.Address, recipientId [][32]byte, rwaToken []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}
	var recipientIdRule []interface{}
	for _, recipientIdItem := range recipientId {
		recipientIdRule = append(recipientIdRule, recipientIdItem)
	}
	var rwaTokenRule []interface{}
	for _, rwaTokenItem := range rwaToken {
		rwaTokenRule = append(rwaTokenRule, rwaTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "AdminMint", recipientRule, recipientIdRule, rwaTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerAdminMint)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "AdminMint", log); err != nil {
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

// ParseAdminMint is a log parse operation binding the contract event 0x64246d088e96d6b1bbb49bd2ab458376867796a99add0e25a88c53bcee7ddbdf.
//
// Solidity: event AdminMint(address indexed recipient, bytes32 indexed recipientId, address indexed rwaToken, uint256 rwaAmount, bytes32 metadata)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseAdminMint(log types.Log) (*OndotokenmanagerAdminMint, error) {
	event := new(OndotokenmanagerAdminMint)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "AdminMint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGMTokenMintingPausedIterator is returned from FilterGMTokenMintingPaused and is used to iterate over the raw logs and unpacked data for GMTokenMintingPaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenMintingPausedIterator struct {
	Event *OndotokenmanagerGMTokenMintingPaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGMTokenMintingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGMTokenMintingPaused)
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
		it.Event = new(OndotokenmanagerGMTokenMintingPaused)
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
func (it *OndotokenmanagerGMTokenMintingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGMTokenMintingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGMTokenMintingPaused represents a GMTokenMintingPaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenMintingPaused struct {
	GmToken common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGMTokenMintingPaused is a free log retrieval operation binding the contract event 0x0de7e27328f4d5d9fc5797bf8cc1f3ca9c0d366a83aae68215e9e7dbf156a1f1.
//
// Solidity: event GMTokenMintingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGMTokenMintingPaused(opts *bind.FilterOpts, gmToken []common.Address) (*OndotokenmanagerGMTokenMintingPausedIterator, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GMTokenMintingPaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGMTokenMintingPausedIterator{contract: _Ondotokenmanager.contract, event: "GMTokenMintingPaused", logs: logs, sub: sub}, nil
}

// WatchGMTokenMintingPaused is a free log subscription operation binding the contract event 0x0de7e27328f4d5d9fc5797bf8cc1f3ca9c0d366a83aae68215e9e7dbf156a1f1.
//
// Solidity: event GMTokenMintingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGMTokenMintingPaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGMTokenMintingPaused, gmToken []common.Address) (event.Subscription, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GMTokenMintingPaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGMTokenMintingPaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenMintingPaused", log); err != nil {
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

// ParseGMTokenMintingPaused is a log parse operation binding the contract event 0x0de7e27328f4d5d9fc5797bf8cc1f3ca9c0d366a83aae68215e9e7dbf156a1f1.
//
// Solidity: event GMTokenMintingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGMTokenMintingPaused(log types.Log) (*OndotokenmanagerGMTokenMintingPaused, error) {
	event := new(OndotokenmanagerGMTokenMintingPaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenMintingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGMTokenMintingUnpausedIterator is returned from FilterGMTokenMintingUnpaused and is used to iterate over the raw logs and unpacked data for GMTokenMintingUnpaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenMintingUnpausedIterator struct {
	Event *OndotokenmanagerGMTokenMintingUnpaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGMTokenMintingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGMTokenMintingUnpaused)
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
		it.Event = new(OndotokenmanagerGMTokenMintingUnpaused)
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
func (it *OndotokenmanagerGMTokenMintingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGMTokenMintingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGMTokenMintingUnpaused represents a GMTokenMintingUnpaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenMintingUnpaused struct {
	GmToken common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGMTokenMintingUnpaused is a free log retrieval operation binding the contract event 0x8ea8fa511d23c1a22f92d6480c103dff2994b522f89b46b18c9493130d1121de.
//
// Solidity: event GMTokenMintingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGMTokenMintingUnpaused(opts *bind.FilterOpts, gmToken []common.Address) (*OndotokenmanagerGMTokenMintingUnpausedIterator, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GMTokenMintingUnpaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGMTokenMintingUnpausedIterator{contract: _Ondotokenmanager.contract, event: "GMTokenMintingUnpaused", logs: logs, sub: sub}, nil
}

// WatchGMTokenMintingUnpaused is a free log subscription operation binding the contract event 0x8ea8fa511d23c1a22f92d6480c103dff2994b522f89b46b18c9493130d1121de.
//
// Solidity: event GMTokenMintingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGMTokenMintingUnpaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGMTokenMintingUnpaused, gmToken []common.Address) (event.Subscription, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GMTokenMintingUnpaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGMTokenMintingUnpaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenMintingUnpaused", log); err != nil {
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

// ParseGMTokenMintingUnpaused is a log parse operation binding the contract event 0x8ea8fa511d23c1a22f92d6480c103dff2994b522f89b46b18c9493130d1121de.
//
// Solidity: event GMTokenMintingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGMTokenMintingUnpaused(log types.Log) (*OndotokenmanagerGMTokenMintingUnpaused, error) {
	event := new(OndotokenmanagerGMTokenMintingUnpaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenMintingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGMTokenRedeemingPausedIterator is returned from FilterGMTokenRedeemingPaused and is used to iterate over the raw logs and unpacked data for GMTokenRedeemingPaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRedeemingPausedIterator struct {
	Event *OndotokenmanagerGMTokenRedeemingPaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGMTokenRedeemingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGMTokenRedeemingPaused)
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
		it.Event = new(OndotokenmanagerGMTokenRedeemingPaused)
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
func (it *OndotokenmanagerGMTokenRedeemingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGMTokenRedeemingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGMTokenRedeemingPaused represents a GMTokenRedeemingPaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRedeemingPaused struct {
	GmToken common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGMTokenRedeemingPaused is a free log retrieval operation binding the contract event 0xcacc4dbd81af0cceec871b59f5155fa60dc3a52b970778ba54b2cf66f0b5deeb.
//
// Solidity: event GMTokenRedeemingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGMTokenRedeemingPaused(opts *bind.FilterOpts, gmToken []common.Address) (*OndotokenmanagerGMTokenRedeemingPausedIterator, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GMTokenRedeemingPaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGMTokenRedeemingPausedIterator{contract: _Ondotokenmanager.contract, event: "GMTokenRedeemingPaused", logs: logs, sub: sub}, nil
}

// WatchGMTokenRedeemingPaused is a free log subscription operation binding the contract event 0xcacc4dbd81af0cceec871b59f5155fa60dc3a52b970778ba54b2cf66f0b5deeb.
//
// Solidity: event GMTokenRedeemingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGMTokenRedeemingPaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGMTokenRedeemingPaused, gmToken []common.Address) (event.Subscription, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GMTokenRedeemingPaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGMTokenRedeemingPaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRedeemingPaused", log); err != nil {
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

// ParseGMTokenRedeemingPaused is a log parse operation binding the contract event 0xcacc4dbd81af0cceec871b59f5155fa60dc3a52b970778ba54b2cf66f0b5deeb.
//
// Solidity: event GMTokenRedeemingPaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGMTokenRedeemingPaused(log types.Log) (*OndotokenmanagerGMTokenRedeemingPaused, error) {
	event := new(OndotokenmanagerGMTokenRedeemingPaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRedeemingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGMTokenRedeemingUnpausedIterator is returned from FilterGMTokenRedeemingUnpaused and is used to iterate over the raw logs and unpacked data for GMTokenRedeemingUnpaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRedeemingUnpausedIterator struct {
	Event *OndotokenmanagerGMTokenRedeemingUnpaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGMTokenRedeemingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGMTokenRedeemingUnpaused)
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
		it.Event = new(OndotokenmanagerGMTokenRedeemingUnpaused)
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
func (it *OndotokenmanagerGMTokenRedeemingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGMTokenRedeemingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGMTokenRedeemingUnpaused represents a GMTokenRedeemingUnpaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRedeemingUnpaused struct {
	GmToken common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGMTokenRedeemingUnpaused is a free log retrieval operation binding the contract event 0xf7db4b816293b43fac22e1773acf5d3774a3d70ab993f04995406fadd92b546a.
//
// Solidity: event GMTokenRedeemingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGMTokenRedeemingUnpaused(opts *bind.FilterOpts, gmToken []common.Address) (*OndotokenmanagerGMTokenRedeemingUnpausedIterator, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GMTokenRedeemingUnpaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGMTokenRedeemingUnpausedIterator{contract: _Ondotokenmanager.contract, event: "GMTokenRedeemingUnpaused", logs: logs, sub: sub}, nil
}

// WatchGMTokenRedeemingUnpaused is a free log subscription operation binding the contract event 0xf7db4b816293b43fac22e1773acf5d3774a3d70ab993f04995406fadd92b546a.
//
// Solidity: event GMTokenRedeemingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGMTokenRedeemingUnpaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGMTokenRedeemingUnpaused, gmToken []common.Address) (event.Subscription, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GMTokenRedeemingUnpaused", gmTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGMTokenRedeemingUnpaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRedeemingUnpaused", log); err != nil {
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

// ParseGMTokenRedeemingUnpaused is a log parse operation binding the contract event 0xf7db4b816293b43fac22e1773acf5d3774a3d70ab993f04995406fadd92b546a.
//
// Solidity: event GMTokenRedeemingUnpaused(address indexed gmToken)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGMTokenRedeemingUnpaused(log types.Log) (*OndotokenmanagerGMTokenRedeemingUnpaused, error) {
	event := new(OndotokenmanagerGMTokenRedeemingUnpaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRedeemingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGMTokenRegisteredIterator is returned from FilterGMTokenRegistered and is used to iterate over the raw logs and unpacked data for GMTokenRegistered events raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRegisteredIterator struct {
	Event *OndotokenmanagerGMTokenRegistered // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGMTokenRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGMTokenRegistered)
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
		it.Event = new(OndotokenmanagerGMTokenRegistered)
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
func (it *OndotokenmanagerGMTokenRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGMTokenRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGMTokenRegistered represents a GMTokenRegistered event raised by the Ondotokenmanager contract.
type OndotokenmanagerGMTokenRegistered struct {
	GmToken    common.Address
	Registered bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGMTokenRegistered is a free log retrieval operation binding the contract event 0x8f1a189329487ea2f651b3797ac87d45216b42a18c01a7335169539b76df1338.
//
// Solidity: event GMTokenRegistered(address indexed gmToken, bool indexed registered)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGMTokenRegistered(opts *bind.FilterOpts, gmToken []common.Address, registered []bool) (*OndotokenmanagerGMTokenRegisteredIterator, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}
	var registeredRule []interface{}
	for _, registeredItem := range registered {
		registeredRule = append(registeredRule, registeredItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GMTokenRegistered", gmTokenRule, registeredRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGMTokenRegisteredIterator{contract: _Ondotokenmanager.contract, event: "GMTokenRegistered", logs: logs, sub: sub}, nil
}

// WatchGMTokenRegistered is a free log subscription operation binding the contract event 0x8f1a189329487ea2f651b3797ac87d45216b42a18c01a7335169539b76df1338.
//
// Solidity: event GMTokenRegistered(address indexed gmToken, bool indexed registered)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGMTokenRegistered(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGMTokenRegistered, gmToken []common.Address, registered []bool) (event.Subscription, error) {

	var gmTokenRule []interface{}
	for _, gmTokenItem := range gmToken {
		gmTokenRule = append(gmTokenRule, gmTokenItem)
	}
	var registeredRule []interface{}
	for _, registeredItem := range registered {
		registeredRule = append(registeredRule, registeredItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GMTokenRegistered", gmTokenRule, registeredRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGMTokenRegistered)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRegistered", log); err != nil {
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

// ParseGMTokenRegistered is a log parse operation binding the contract event 0x8f1a189329487ea2f651b3797ac87d45216b42a18c01a7335169539b76df1338.
//
// Solidity: event GMTokenRegistered(address indexed gmToken, bool indexed registered)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGMTokenRegistered(log types.Log) (*OndotokenmanagerGMTokenRegistered, error) {
	event := new(OndotokenmanagerGMTokenRegistered)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GMTokenRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGlobalMintingPausedIterator is returned from FilterGlobalMintingPaused and is used to iterate over the raw logs and unpacked data for GlobalMintingPaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalMintingPausedIterator struct {
	Event *OndotokenmanagerGlobalMintingPaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGlobalMintingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGlobalMintingPaused)
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
		it.Event = new(OndotokenmanagerGlobalMintingPaused)
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
func (it *OndotokenmanagerGlobalMintingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGlobalMintingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGlobalMintingPaused represents a GlobalMintingPaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalMintingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalMintingPaused is a free log retrieval operation binding the contract event 0x94eae7ff948a5ddfb0e5f291aa214ae9dd7d12ea33086a414143bc4a34087341.
//
// Solidity: event GlobalMintingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGlobalMintingPaused(opts *bind.FilterOpts) (*OndotokenmanagerGlobalMintingPausedIterator, error) {

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GlobalMintingPaused")
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGlobalMintingPausedIterator{contract: _Ondotokenmanager.contract, event: "GlobalMintingPaused", logs: logs, sub: sub}, nil
}

// WatchGlobalMintingPaused is a free log subscription operation binding the contract event 0x94eae7ff948a5ddfb0e5f291aa214ae9dd7d12ea33086a414143bc4a34087341.
//
// Solidity: event GlobalMintingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGlobalMintingPaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGlobalMintingPaused) (event.Subscription, error) {

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GlobalMintingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGlobalMintingPaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalMintingPaused", log); err != nil {
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

// ParseGlobalMintingPaused is a log parse operation binding the contract event 0x94eae7ff948a5ddfb0e5f291aa214ae9dd7d12ea33086a414143bc4a34087341.
//
// Solidity: event GlobalMintingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGlobalMintingPaused(log types.Log) (*OndotokenmanagerGlobalMintingPaused, error) {
	event := new(OndotokenmanagerGlobalMintingPaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalMintingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGlobalMintingUnpausedIterator is returned from FilterGlobalMintingUnpaused and is used to iterate over the raw logs and unpacked data for GlobalMintingUnpaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalMintingUnpausedIterator struct {
	Event *OndotokenmanagerGlobalMintingUnpaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGlobalMintingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGlobalMintingUnpaused)
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
		it.Event = new(OndotokenmanagerGlobalMintingUnpaused)
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
func (it *OndotokenmanagerGlobalMintingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGlobalMintingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGlobalMintingUnpaused represents a GlobalMintingUnpaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalMintingUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalMintingUnpaused is a free log retrieval operation binding the contract event 0x7c5bd2c6bf55553fb8379a172f2f9d0369fc18a7bf82a47486ac611a92150601.
//
// Solidity: event GlobalMintingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGlobalMintingUnpaused(opts *bind.FilterOpts) (*OndotokenmanagerGlobalMintingUnpausedIterator, error) {

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GlobalMintingUnpaused")
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGlobalMintingUnpausedIterator{contract: _Ondotokenmanager.contract, event: "GlobalMintingUnpaused", logs: logs, sub: sub}, nil
}

// WatchGlobalMintingUnpaused is a free log subscription operation binding the contract event 0x7c5bd2c6bf55553fb8379a172f2f9d0369fc18a7bf82a47486ac611a92150601.
//
// Solidity: event GlobalMintingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGlobalMintingUnpaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGlobalMintingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GlobalMintingUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGlobalMintingUnpaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalMintingUnpaused", log); err != nil {
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

// ParseGlobalMintingUnpaused is a log parse operation binding the contract event 0x7c5bd2c6bf55553fb8379a172f2f9d0369fc18a7bf82a47486ac611a92150601.
//
// Solidity: event GlobalMintingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGlobalMintingUnpaused(log types.Log) (*OndotokenmanagerGlobalMintingUnpaused, error) {
	event := new(OndotokenmanagerGlobalMintingUnpaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalMintingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGlobalRedeemingPausedIterator is returned from FilterGlobalRedeemingPaused and is used to iterate over the raw logs and unpacked data for GlobalRedeemingPaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalRedeemingPausedIterator struct {
	Event *OndotokenmanagerGlobalRedeemingPaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGlobalRedeemingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGlobalRedeemingPaused)
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
		it.Event = new(OndotokenmanagerGlobalRedeemingPaused)
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
func (it *OndotokenmanagerGlobalRedeemingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGlobalRedeemingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGlobalRedeemingPaused represents a GlobalRedeemingPaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalRedeemingPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalRedeemingPaused is a free log retrieval operation binding the contract event 0x694abd7c120a829fef1aac704885e405362fcd69c82f114ccec2959a402633c5.
//
// Solidity: event GlobalRedeemingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGlobalRedeemingPaused(opts *bind.FilterOpts) (*OndotokenmanagerGlobalRedeemingPausedIterator, error) {

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GlobalRedeemingPaused")
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGlobalRedeemingPausedIterator{contract: _Ondotokenmanager.contract, event: "GlobalRedeemingPaused", logs: logs, sub: sub}, nil
}

// WatchGlobalRedeemingPaused is a free log subscription operation binding the contract event 0x694abd7c120a829fef1aac704885e405362fcd69c82f114ccec2959a402633c5.
//
// Solidity: event GlobalRedeemingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGlobalRedeemingPaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGlobalRedeemingPaused) (event.Subscription, error) {

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GlobalRedeemingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGlobalRedeemingPaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalRedeemingPaused", log); err != nil {
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

// ParseGlobalRedeemingPaused is a log parse operation binding the contract event 0x694abd7c120a829fef1aac704885e405362fcd69c82f114ccec2959a402633c5.
//
// Solidity: event GlobalRedeemingPaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGlobalRedeemingPaused(log types.Log) (*OndotokenmanagerGlobalRedeemingPaused, error) {
	event := new(OndotokenmanagerGlobalRedeemingPaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalRedeemingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerGlobalRedeemingUnpausedIterator is returned from FilterGlobalRedeemingUnpaused and is used to iterate over the raw logs and unpacked data for GlobalRedeemingUnpaused events raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalRedeemingUnpausedIterator struct {
	Event *OndotokenmanagerGlobalRedeemingUnpaused // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerGlobalRedeemingUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerGlobalRedeemingUnpaused)
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
		it.Event = new(OndotokenmanagerGlobalRedeemingUnpaused)
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
func (it *OndotokenmanagerGlobalRedeemingUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerGlobalRedeemingUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerGlobalRedeemingUnpaused represents a GlobalRedeemingUnpaused event raised by the Ondotokenmanager contract.
type OndotokenmanagerGlobalRedeemingUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterGlobalRedeemingUnpaused is a free log retrieval operation binding the contract event 0x8cba7e15b16021395abffdc05ba74a88afd2b3dce03e075ddb041760a3890688.
//
// Solidity: event GlobalRedeemingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterGlobalRedeemingUnpaused(opts *bind.FilterOpts) (*OndotokenmanagerGlobalRedeemingUnpausedIterator, error) {

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "GlobalRedeemingUnpaused")
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerGlobalRedeemingUnpausedIterator{contract: _Ondotokenmanager.contract, event: "GlobalRedeemingUnpaused", logs: logs, sub: sub}, nil
}

// WatchGlobalRedeemingUnpaused is a free log subscription operation binding the contract event 0x8cba7e15b16021395abffdc05ba74a88afd2b3dce03e075ddb041760a3890688.
//
// Solidity: event GlobalRedeemingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchGlobalRedeemingUnpaused(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerGlobalRedeemingUnpaused) (event.Subscription, error) {

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "GlobalRedeemingUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerGlobalRedeemingUnpaused)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalRedeemingUnpaused", log); err != nil {
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

// ParseGlobalRedeemingUnpaused is a log parse operation binding the contract event 0x8cba7e15b16021395abffdc05ba74a88afd2b3dce03e075ddb041760a3890688.
//
// Solidity: event GlobalRedeemingUnpaused()
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseGlobalRedeemingUnpaused(log types.Log) (*OndotokenmanagerGlobalRedeemingUnpaused, error) {
	event := new(OndotokenmanagerGlobalRedeemingUnpaused)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "GlobalRedeemingUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerIssuanceHoursSetIterator is returned from FilterIssuanceHoursSet and is used to iterate over the raw logs and unpacked data for IssuanceHoursSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerIssuanceHoursSetIterator struct {
	Event *OndotokenmanagerIssuanceHoursSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerIssuanceHoursSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerIssuanceHoursSet)
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
		it.Event = new(OndotokenmanagerIssuanceHoursSet)
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
func (it *OndotokenmanagerIssuanceHoursSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerIssuanceHoursSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerIssuanceHoursSet represents a IssuanceHoursSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerIssuanceHoursSet struct {
	OldIssuanceHours common.Address
	NewIssuanceHours common.Address
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterIssuanceHoursSet is a free log retrieval operation binding the contract event 0xcaf02d44223f855a96856532b7f6b262b8c1e57639d6ea29ff00f49b65d6a0be.
//
// Solidity: event IssuanceHoursSet(address indexed oldIssuanceHours, address indexed newIssuanceHours)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterIssuanceHoursSet(opts *bind.FilterOpts, oldIssuanceHours []common.Address, newIssuanceHours []common.Address) (*OndotokenmanagerIssuanceHoursSetIterator, error) {

	var oldIssuanceHoursRule []interface{}
	for _, oldIssuanceHoursItem := range oldIssuanceHours {
		oldIssuanceHoursRule = append(oldIssuanceHoursRule, oldIssuanceHoursItem)
	}
	var newIssuanceHoursRule []interface{}
	for _, newIssuanceHoursItem := range newIssuanceHours {
		newIssuanceHoursRule = append(newIssuanceHoursRule, newIssuanceHoursItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "IssuanceHoursSet", oldIssuanceHoursRule, newIssuanceHoursRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerIssuanceHoursSetIterator{contract: _Ondotokenmanager.contract, event: "IssuanceHoursSet", logs: logs, sub: sub}, nil
}

// WatchIssuanceHoursSet is a free log subscription operation binding the contract event 0xcaf02d44223f855a96856532b7f6b262b8c1e57639d6ea29ff00f49b65d6a0be.
//
// Solidity: event IssuanceHoursSet(address indexed oldIssuanceHours, address indexed newIssuanceHours)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchIssuanceHoursSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerIssuanceHoursSet, oldIssuanceHours []common.Address, newIssuanceHours []common.Address) (event.Subscription, error) {

	var oldIssuanceHoursRule []interface{}
	for _, oldIssuanceHoursItem := range oldIssuanceHours {
		oldIssuanceHoursRule = append(oldIssuanceHoursRule, oldIssuanceHoursItem)
	}
	var newIssuanceHoursRule []interface{}
	for _, newIssuanceHoursItem := range newIssuanceHours {
		newIssuanceHoursRule = append(newIssuanceHoursRule, newIssuanceHoursItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "IssuanceHoursSet", oldIssuanceHoursRule, newIssuanceHoursRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerIssuanceHoursSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "IssuanceHoursSet", log); err != nil {
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

// ParseIssuanceHoursSet is a log parse operation binding the contract event 0xcaf02d44223f855a96856532b7f6b262b8c1e57639d6ea29ff00f49b65d6a0be.
//
// Solidity: event IssuanceHoursSet(address indexed oldIssuanceHours, address indexed newIssuanceHours)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseIssuanceHoursSet(log types.Log) (*OndotokenmanagerIssuanceHoursSet, error) {
	event := new(OndotokenmanagerIssuanceHoursSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "IssuanceHoursSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerMinimumDepositAmountSetIterator is returned from FilterMinimumDepositAmountSet and is used to iterate over the raw logs and unpacked data for MinimumDepositAmountSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerMinimumDepositAmountSetIterator struct {
	Event *OndotokenmanagerMinimumDepositAmountSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerMinimumDepositAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerMinimumDepositAmountSet)
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
		it.Event = new(OndotokenmanagerMinimumDepositAmountSet)
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
func (it *OndotokenmanagerMinimumDepositAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerMinimumDepositAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerMinimumDepositAmountSet represents a MinimumDepositAmountSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerMinimumDepositAmountSet struct {
	OldMinDepositAmount *big.Int
	NewMinDepositAmount *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterMinimumDepositAmountSet is a free log retrieval operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldMinDepositAmount, uint256 indexed newMinDepositAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterMinimumDepositAmountSet(opts *bind.FilterOpts, oldMinDepositAmount []*big.Int, newMinDepositAmount []*big.Int) (*OndotokenmanagerMinimumDepositAmountSetIterator, error) {

	var oldMinDepositAmountRule []interface{}
	for _, oldMinDepositAmountItem := range oldMinDepositAmount {
		oldMinDepositAmountRule = append(oldMinDepositAmountRule, oldMinDepositAmountItem)
	}
	var newMinDepositAmountRule []interface{}
	for _, newMinDepositAmountItem := range newMinDepositAmount {
		newMinDepositAmountRule = append(newMinDepositAmountRule, newMinDepositAmountItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "MinimumDepositAmountSet", oldMinDepositAmountRule, newMinDepositAmountRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerMinimumDepositAmountSetIterator{contract: _Ondotokenmanager.contract, event: "MinimumDepositAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinimumDepositAmountSet is a free log subscription operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldMinDepositAmount, uint256 indexed newMinDepositAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchMinimumDepositAmountSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerMinimumDepositAmountSet, oldMinDepositAmount []*big.Int, newMinDepositAmount []*big.Int) (event.Subscription, error) {

	var oldMinDepositAmountRule []interface{}
	for _, oldMinDepositAmountItem := range oldMinDepositAmount {
		oldMinDepositAmountRule = append(oldMinDepositAmountRule, oldMinDepositAmountItem)
	}
	var newMinDepositAmountRule []interface{}
	for _, newMinDepositAmountItem := range newMinDepositAmount {
		newMinDepositAmountRule = append(newMinDepositAmountRule, newMinDepositAmountItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "MinimumDepositAmountSet", oldMinDepositAmountRule, newMinDepositAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerMinimumDepositAmountSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "MinimumDepositAmountSet", log); err != nil {
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

// ParseMinimumDepositAmountSet is a log parse operation binding the contract event 0xe6e25add7363f8f8a40cbea9810d3115a33703b10972ef759104219b00657436.
//
// Solidity: event MinimumDepositAmountSet(uint256 indexed oldMinDepositAmount, uint256 indexed newMinDepositAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseMinimumDepositAmountSet(log types.Log) (*OndotokenmanagerMinimumDepositAmountSet, error) {
	event := new(OndotokenmanagerMinimumDepositAmountSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "MinimumDepositAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerMinimumRedemptionAmountSetIterator is returned from FilterMinimumRedemptionAmountSet and is used to iterate over the raw logs and unpacked data for MinimumRedemptionAmountSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerMinimumRedemptionAmountSetIterator struct {
	Event *OndotokenmanagerMinimumRedemptionAmountSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerMinimumRedemptionAmountSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerMinimumRedemptionAmountSet)
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
		it.Event = new(OndotokenmanagerMinimumRedemptionAmountSet)
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
func (it *OndotokenmanagerMinimumRedemptionAmountSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerMinimumRedemptionAmountSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerMinimumRedemptionAmountSet represents a MinimumRedemptionAmountSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerMinimumRedemptionAmountSet struct {
	OldMinRedemptionAmount *big.Int
	NewMinRedemptionAmount *big.Int
	Raw                    types.Log // Blockchain specific contextual infos
}

// FilterMinimumRedemptionAmountSet is a free log retrieval operation binding the contract event 0xfdaf6ed728cef208e62328a008209556f8281f3062b14dd08aaaa90fa1594211.
//
// Solidity: event MinimumRedemptionAmountSet(uint256 indexed oldMinRedemptionAmount, uint256 indexed newMinRedemptionAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterMinimumRedemptionAmountSet(opts *bind.FilterOpts, oldMinRedemptionAmount []*big.Int, newMinRedemptionAmount []*big.Int) (*OndotokenmanagerMinimumRedemptionAmountSetIterator, error) {

	var oldMinRedemptionAmountRule []interface{}
	for _, oldMinRedemptionAmountItem := range oldMinRedemptionAmount {
		oldMinRedemptionAmountRule = append(oldMinRedemptionAmountRule, oldMinRedemptionAmountItem)
	}
	var newMinRedemptionAmountRule []interface{}
	for _, newMinRedemptionAmountItem := range newMinRedemptionAmount {
		newMinRedemptionAmountRule = append(newMinRedemptionAmountRule, newMinRedemptionAmountItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "MinimumRedemptionAmountSet", oldMinRedemptionAmountRule, newMinRedemptionAmountRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerMinimumRedemptionAmountSetIterator{contract: _Ondotokenmanager.contract, event: "MinimumRedemptionAmountSet", logs: logs, sub: sub}, nil
}

// WatchMinimumRedemptionAmountSet is a free log subscription operation binding the contract event 0xfdaf6ed728cef208e62328a008209556f8281f3062b14dd08aaaa90fa1594211.
//
// Solidity: event MinimumRedemptionAmountSet(uint256 indexed oldMinRedemptionAmount, uint256 indexed newMinRedemptionAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchMinimumRedemptionAmountSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerMinimumRedemptionAmountSet, oldMinRedemptionAmount []*big.Int, newMinRedemptionAmount []*big.Int) (event.Subscription, error) {

	var oldMinRedemptionAmountRule []interface{}
	for _, oldMinRedemptionAmountItem := range oldMinRedemptionAmount {
		oldMinRedemptionAmountRule = append(oldMinRedemptionAmountRule, oldMinRedemptionAmountItem)
	}
	var newMinRedemptionAmountRule []interface{}
	for _, newMinRedemptionAmountItem := range newMinRedemptionAmount {
		newMinRedemptionAmountRule = append(newMinRedemptionAmountRule, newMinRedemptionAmountItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "MinimumRedemptionAmountSet", oldMinRedemptionAmountRule, newMinRedemptionAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerMinimumRedemptionAmountSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "MinimumRedemptionAmountSet", log); err != nil {
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

// ParseMinimumRedemptionAmountSet is a log parse operation binding the contract event 0xfdaf6ed728cef208e62328a008209556f8281f3062b14dd08aaaa90fa1594211.
//
// Solidity: event MinimumRedemptionAmountSet(uint256 indexed oldMinRedemptionAmount, uint256 indexed newMinRedemptionAmount)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseMinimumRedemptionAmountSet(log types.Log) (*OndotokenmanagerMinimumRedemptionAmountSet, error) {
	event := new(OndotokenmanagerMinimumRedemptionAmountSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "MinimumRedemptionAmountSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerOndoIDRegistrySetIterator is returned from FilterOndoIDRegistrySet and is used to iterate over the raw logs and unpacked data for OndoIDRegistrySet events raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoIDRegistrySetIterator struct {
	Event *OndotokenmanagerOndoIDRegistrySet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerOndoIDRegistrySetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerOndoIDRegistrySet)
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
		it.Event = new(OndotokenmanagerOndoIDRegistrySet)
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
func (it *OndotokenmanagerOndoIDRegistrySetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerOndoIDRegistrySetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerOndoIDRegistrySet represents a OndoIDRegistrySet event raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoIDRegistrySet struct {
	OldOndoIDRegistry common.Address
	NewOndoIDRegistry common.Address
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterOndoIDRegistrySet is a free log retrieval operation binding the contract event 0xf78e014ab86d7bb38135262d64726cef2e151dae47c0947a1d6bb970702c30d6.
//
// Solidity: event OndoIDRegistrySet(address indexed oldOndoIDRegistry, address indexed newOndoIDRegistry)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterOndoIDRegistrySet(opts *bind.FilterOpts, oldOndoIDRegistry []common.Address, newOndoIDRegistry []common.Address) (*OndotokenmanagerOndoIDRegistrySetIterator, error) {

	var oldOndoIDRegistryRule []interface{}
	for _, oldOndoIDRegistryItem := range oldOndoIDRegistry {
		oldOndoIDRegistryRule = append(oldOndoIDRegistryRule, oldOndoIDRegistryItem)
	}
	var newOndoIDRegistryRule []interface{}
	for _, newOndoIDRegistryItem := range newOndoIDRegistry {
		newOndoIDRegistryRule = append(newOndoIDRegistryRule, newOndoIDRegistryItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "OndoIDRegistrySet", oldOndoIDRegistryRule, newOndoIDRegistryRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerOndoIDRegistrySetIterator{contract: _Ondotokenmanager.contract, event: "OndoIDRegistrySet", logs: logs, sub: sub}, nil
}

// WatchOndoIDRegistrySet is a free log subscription operation binding the contract event 0xf78e014ab86d7bb38135262d64726cef2e151dae47c0947a1d6bb970702c30d6.
//
// Solidity: event OndoIDRegistrySet(address indexed oldOndoIDRegistry, address indexed newOndoIDRegistry)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchOndoIDRegistrySet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerOndoIDRegistrySet, oldOndoIDRegistry []common.Address, newOndoIDRegistry []common.Address) (event.Subscription, error) {

	var oldOndoIDRegistryRule []interface{}
	for _, oldOndoIDRegistryItem := range oldOndoIDRegistry {
		oldOndoIDRegistryRule = append(oldOndoIDRegistryRule, oldOndoIDRegistryItem)
	}
	var newOndoIDRegistryRule []interface{}
	for _, newOndoIDRegistryItem := range newOndoIDRegistry {
		newOndoIDRegistryRule = append(newOndoIDRegistryRule, newOndoIDRegistryItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "OndoIDRegistrySet", oldOndoIDRegistryRule, newOndoIDRegistryRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerOndoIDRegistrySet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoIDRegistrySet", log); err != nil {
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

// ParseOndoIDRegistrySet is a log parse operation binding the contract event 0xf78e014ab86d7bb38135262d64726cef2e151dae47c0947a1d6bb970702c30d6.
//
// Solidity: event OndoIDRegistrySet(address indexed oldOndoIDRegistry, address indexed newOndoIDRegistry)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseOndoIDRegistrySet(log types.Log) (*OndotokenmanagerOndoIDRegistrySet, error) {
	event := new(OndotokenmanagerOndoIDRegistrySet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoIDRegistrySet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerOndoRateLimiterSetIterator is returned from FilterOndoRateLimiterSet and is used to iterate over the raw logs and unpacked data for OndoRateLimiterSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoRateLimiterSetIterator struct {
	Event *OndotokenmanagerOndoRateLimiterSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerOndoRateLimiterSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerOndoRateLimiterSet)
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
		it.Event = new(OndotokenmanagerOndoRateLimiterSet)
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
func (it *OndotokenmanagerOndoRateLimiterSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerOndoRateLimiterSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerOndoRateLimiterSet represents a OndoRateLimiterSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoRateLimiterSet struct {
	OldOndoRateLimiter common.Address
	NewOndoRateLimiter common.Address
	Raw                types.Log // Blockchain specific contextual infos
}

// FilterOndoRateLimiterSet is a free log retrieval operation binding the contract event 0xc54bf4c3067c1d8f65e053dafb5dbb699615b1b527d2866bd0223102bb2e692d.
//
// Solidity: event OndoRateLimiterSet(address indexed oldOndoRateLimiter, address indexed newOndoRateLimiter)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterOndoRateLimiterSet(opts *bind.FilterOpts, oldOndoRateLimiter []common.Address, newOndoRateLimiter []common.Address) (*OndotokenmanagerOndoRateLimiterSetIterator, error) {

	var oldOndoRateLimiterRule []interface{}
	for _, oldOndoRateLimiterItem := range oldOndoRateLimiter {
		oldOndoRateLimiterRule = append(oldOndoRateLimiterRule, oldOndoRateLimiterItem)
	}
	var newOndoRateLimiterRule []interface{}
	for _, newOndoRateLimiterItem := range newOndoRateLimiter {
		newOndoRateLimiterRule = append(newOndoRateLimiterRule, newOndoRateLimiterItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "OndoRateLimiterSet", oldOndoRateLimiterRule, newOndoRateLimiterRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerOndoRateLimiterSetIterator{contract: _Ondotokenmanager.contract, event: "OndoRateLimiterSet", logs: logs, sub: sub}, nil
}

// WatchOndoRateLimiterSet is a free log subscription operation binding the contract event 0xc54bf4c3067c1d8f65e053dafb5dbb699615b1b527d2866bd0223102bb2e692d.
//
// Solidity: event OndoRateLimiterSet(address indexed oldOndoRateLimiter, address indexed newOndoRateLimiter)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchOndoRateLimiterSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerOndoRateLimiterSet, oldOndoRateLimiter []common.Address, newOndoRateLimiter []common.Address) (event.Subscription, error) {

	var oldOndoRateLimiterRule []interface{}
	for _, oldOndoRateLimiterItem := range oldOndoRateLimiter {
		oldOndoRateLimiterRule = append(oldOndoRateLimiterRule, oldOndoRateLimiterItem)
	}
	var newOndoRateLimiterRule []interface{}
	for _, newOndoRateLimiterItem := range newOndoRateLimiter {
		newOndoRateLimiterRule = append(newOndoRateLimiterRule, newOndoRateLimiterItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "OndoRateLimiterSet", oldOndoRateLimiterRule, newOndoRateLimiterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerOndoRateLimiterSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoRateLimiterSet", log); err != nil {
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

// ParseOndoRateLimiterSet is a log parse operation binding the contract event 0xc54bf4c3067c1d8f65e053dafb5dbb699615b1b527d2866bd0223102bb2e692d.
//
// Solidity: event OndoRateLimiterSet(address indexed oldOndoRateLimiter, address indexed newOndoRateLimiter)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseOndoRateLimiterSet(log types.Log) (*OndotokenmanagerOndoRateLimiterSet, error) {
	event := new(OndotokenmanagerOndoRateLimiterSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoRateLimiterSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerOndoSanityCheckOracleSetIterator is returned from FilterOndoSanityCheckOracleSet and is used to iterate over the raw logs and unpacked data for OndoSanityCheckOracleSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoSanityCheckOracleSetIterator struct {
	Event *OndotokenmanagerOndoSanityCheckOracleSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerOndoSanityCheckOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerOndoSanityCheckOracleSet)
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
		it.Event = new(OndotokenmanagerOndoSanityCheckOracleSet)
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
func (it *OndotokenmanagerOndoSanityCheckOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerOndoSanityCheckOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerOndoSanityCheckOracleSet represents a OndoSanityCheckOracleSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerOndoSanityCheckOracleSet struct {
	OldOndoSanityCheckOracle common.Address
	NewOndoSanityCheckOracle common.Address
	Raw                      types.Log // Blockchain specific contextual infos
}

// FilterOndoSanityCheckOracleSet is a free log retrieval operation binding the contract event 0x7c40bb6767b8c6caf208ae3b9558c37a91b5710f5157deb182aae429ac24d336.
//
// Solidity: event OndoSanityCheckOracleSet(address indexed oldOndoSanityCheckOracle, address indexed newOndoSanityCheckOracle)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterOndoSanityCheckOracleSet(opts *bind.FilterOpts, oldOndoSanityCheckOracle []common.Address, newOndoSanityCheckOracle []common.Address) (*OndotokenmanagerOndoSanityCheckOracleSetIterator, error) {

	var oldOndoSanityCheckOracleRule []interface{}
	for _, oldOndoSanityCheckOracleItem := range oldOndoSanityCheckOracle {
		oldOndoSanityCheckOracleRule = append(oldOndoSanityCheckOracleRule, oldOndoSanityCheckOracleItem)
	}
	var newOndoSanityCheckOracleRule []interface{}
	for _, newOndoSanityCheckOracleItem := range newOndoSanityCheckOracle {
		newOndoSanityCheckOracleRule = append(newOndoSanityCheckOracleRule, newOndoSanityCheckOracleItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "OndoSanityCheckOracleSet", oldOndoSanityCheckOracleRule, newOndoSanityCheckOracleRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerOndoSanityCheckOracleSetIterator{contract: _Ondotokenmanager.contract, event: "OndoSanityCheckOracleSet", logs: logs, sub: sub}, nil
}

// WatchOndoSanityCheckOracleSet is a free log subscription operation binding the contract event 0x7c40bb6767b8c6caf208ae3b9558c37a91b5710f5157deb182aae429ac24d336.
//
// Solidity: event OndoSanityCheckOracleSet(address indexed oldOndoSanityCheckOracle, address indexed newOndoSanityCheckOracle)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchOndoSanityCheckOracleSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerOndoSanityCheckOracleSet, oldOndoSanityCheckOracle []common.Address, newOndoSanityCheckOracle []common.Address) (event.Subscription, error) {

	var oldOndoSanityCheckOracleRule []interface{}
	for _, oldOndoSanityCheckOracleItem := range oldOndoSanityCheckOracle {
		oldOndoSanityCheckOracleRule = append(oldOndoSanityCheckOracleRule, oldOndoSanityCheckOracleItem)
	}
	var newOndoSanityCheckOracleRule []interface{}
	for _, newOndoSanityCheckOracleItem := range newOndoSanityCheckOracle {
		newOndoSanityCheckOracleRule = append(newOndoSanityCheckOracleRule, newOndoSanityCheckOracleItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "OndoSanityCheckOracleSet", oldOndoSanityCheckOracleRule, newOndoSanityCheckOracleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerOndoSanityCheckOracleSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoSanityCheckOracleSet", log); err != nil {
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

// ParseOndoSanityCheckOracleSet is a log parse operation binding the contract event 0x7c40bb6767b8c6caf208ae3b9558c37a91b5710f5157deb182aae429ac24d336.
//
// Solidity: event OndoSanityCheckOracleSet(address indexed oldOndoSanityCheckOracle, address indexed newOndoSanityCheckOracle)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseOndoSanityCheckOracleSet(log types.Log) (*OndotokenmanagerOndoSanityCheckOracleSet, error) {
	event := new(OndotokenmanagerOndoSanityCheckOracleSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "OndoSanityCheckOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleAdminChangedIterator struct {
	Event *OndotokenmanagerRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerRoleAdminChanged)
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
		it.Event = new(OndotokenmanagerRoleAdminChanged)
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
func (it *OndotokenmanagerRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerRoleAdminChanged represents a RoleAdminChanged event raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*OndotokenmanagerRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerRoleAdminChangedIterator{contract: _Ondotokenmanager.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerRoleAdminChanged)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseRoleAdminChanged(log types.Log) (*OndotokenmanagerRoleAdminChanged, error) {
	event := new(OndotokenmanagerRoleAdminChanged)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleGrantedIterator struct {
	Event *OndotokenmanagerRoleGranted // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerRoleGranted)
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
		it.Event = new(OndotokenmanagerRoleGranted)
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
func (it *OndotokenmanagerRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerRoleGranted represents a RoleGranted event raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OndotokenmanagerRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerRoleGrantedIterator{contract: _Ondotokenmanager.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerRoleGranted)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseRoleGranted(log types.Log) (*OndotokenmanagerRoleGranted, error) {
	event := new(OndotokenmanagerRoleGranted)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleRevokedIterator struct {
	Event *OndotokenmanagerRoleRevoked // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerRoleRevoked)
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
		it.Event = new(OndotokenmanagerRoleRevoked)
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
func (it *OndotokenmanagerRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerRoleRevoked represents a RoleRevoked event raised by the Ondotokenmanager contract.
type OndotokenmanagerRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*OndotokenmanagerRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerRoleRevokedIterator{contract: _Ondotokenmanager.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerRoleRevoked)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseRoleRevoked(log types.Log) (*OndotokenmanagerRoleRevoked, error) {
	event := new(OndotokenmanagerRoleRevoked)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerTradeExecutedIterator is returned from FilterTradeExecuted and is used to iterate over the raw logs and unpacked data for TradeExecuted events raised by the Ondotokenmanager contract.
type OndotokenmanagerTradeExecutedIterator struct {
	Event *OndotokenmanagerTradeExecuted // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerTradeExecutedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerTradeExecuted)
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
		it.Event = new(OndotokenmanagerTradeExecuted)
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
func (it *OndotokenmanagerTradeExecutedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerTradeExecutedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerTradeExecuted represents a TradeExecuted event raised by the Ondotokenmanager contract.
type OndotokenmanagerTradeExecuted struct {
	ExecutionId    *big.Int
	AttestationId  *big.Int
	ChainId        *big.Int
	UserId         [32]byte
	Side           uint8
	Asset          common.Address
	Price          *big.Int
	Quantity       *big.Int
	Expiration     *big.Int
	AdditionalData [32]byte
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterTradeExecuted is a free log retrieval operation binding the contract event 0xdfe4debbd7bc491e9db43820916e30a4ef9490863fc22b1e79e2f76013901035.
//
// Solidity: event TradeExecuted(uint256 executionId, uint256 attestationId, uint256 chainId, bytes32 userId, uint8 side, address asset, uint256 price, uint256 quantity, uint256 expiration, bytes32 additionalData)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterTradeExecuted(opts *bind.FilterOpts) (*OndotokenmanagerTradeExecutedIterator, error) {

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerTradeExecutedIterator{contract: _Ondotokenmanager.contract, event: "TradeExecuted", logs: logs, sub: sub}, nil
}

// WatchTradeExecuted is a free log subscription operation binding the contract event 0xdfe4debbd7bc491e9db43820916e30a4ef9490863fc22b1e79e2f76013901035.
//
// Solidity: event TradeExecuted(uint256 executionId, uint256 attestationId, uint256 chainId, bytes32 userId, uint8 side, address asset, uint256 price, uint256 quantity, uint256 expiration, bytes32 additionalData)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchTradeExecuted(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerTradeExecuted) (event.Subscription, error) {

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "TradeExecuted")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerTradeExecuted)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
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

// ParseTradeExecuted is a log parse operation binding the contract event 0xdfe4debbd7bc491e9db43820916e30a4ef9490863fc22b1e79e2f76013901035.
//
// Solidity: event TradeExecuted(uint256 executionId, uint256 attestationId, uint256 chainId, bytes32 userId, uint8 side, address asset, uint256 price, uint256 quantity, uint256 expiration, bytes32 additionalData)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseTradeExecuted(log types.Log) (*OndotokenmanagerTradeExecuted, error) {
	event := new(OndotokenmanagerTradeExecuted)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "TradeExecuted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OndotokenmanagerUSDonManagerSetIterator is returned from FilterUSDonManagerSet and is used to iterate over the raw logs and unpacked data for USDonManagerSet events raised by the Ondotokenmanager contract.
type OndotokenmanagerUSDonManagerSetIterator struct {
	Event *OndotokenmanagerUSDonManagerSet // Event containing the contract specifics and raw log

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
func (it *OndotokenmanagerUSDonManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OndotokenmanagerUSDonManagerSet)
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
		it.Event = new(OndotokenmanagerUSDonManagerSet)
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
func (it *OndotokenmanagerUSDonManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OndotokenmanagerUSDonManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OndotokenmanagerUSDonManagerSet represents a USDonManagerSet event raised by the Ondotokenmanager contract.
type OndotokenmanagerUSDonManagerSet struct {
	OldUSDonManager common.Address
	NewUSDonManager common.Address
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterUSDonManagerSet is a free log retrieval operation binding the contract event 0xad45eab3ddda99c8b095504a58e25095a54c31dc487bb3eed9c0a9112ee4eaa8.
//
// Solidity: event USDonManagerSet(address indexed oldUSDonManager, address indexed newUSDonManager)
func (_Ondotokenmanager *OndotokenmanagerFilterer) FilterUSDonManagerSet(opts *bind.FilterOpts, oldUSDonManager []common.Address, newUSDonManager []common.Address) (*OndotokenmanagerUSDonManagerSetIterator, error) {

	var oldUSDonManagerRule []interface{}
	for _, oldUSDonManagerItem := range oldUSDonManager {
		oldUSDonManagerRule = append(oldUSDonManagerRule, oldUSDonManagerItem)
	}
	var newUSDonManagerRule []interface{}
	for _, newUSDonManagerItem := range newUSDonManager {
		newUSDonManagerRule = append(newUSDonManagerRule, newUSDonManagerItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.FilterLogs(opts, "USDonManagerSet", oldUSDonManagerRule, newUSDonManagerRule)
	if err != nil {
		return nil, err
	}
	return &OndotokenmanagerUSDonManagerSetIterator{contract: _Ondotokenmanager.contract, event: "USDonManagerSet", logs: logs, sub: sub}, nil
}

// WatchUSDonManagerSet is a free log subscription operation binding the contract event 0xad45eab3ddda99c8b095504a58e25095a54c31dc487bb3eed9c0a9112ee4eaa8.
//
// Solidity: event USDonManagerSet(address indexed oldUSDonManager, address indexed newUSDonManager)
func (_Ondotokenmanager *OndotokenmanagerFilterer) WatchUSDonManagerSet(opts *bind.WatchOpts, sink chan<- *OndotokenmanagerUSDonManagerSet, oldUSDonManager []common.Address, newUSDonManager []common.Address) (event.Subscription, error) {

	var oldUSDonManagerRule []interface{}
	for _, oldUSDonManagerItem := range oldUSDonManager {
		oldUSDonManagerRule = append(oldUSDonManagerRule, oldUSDonManagerItem)
	}
	var newUSDonManagerRule []interface{}
	for _, newUSDonManagerItem := range newUSDonManager {
		newUSDonManagerRule = append(newUSDonManagerRule, newUSDonManagerItem)
	}

	logs, sub, err := _Ondotokenmanager.contract.WatchLogs(opts, "USDonManagerSet", oldUSDonManagerRule, newUSDonManagerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OndotokenmanagerUSDonManagerSet)
				if err := _Ondotokenmanager.contract.UnpackLog(event, "USDonManagerSet", log); err != nil {
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

// ParseUSDonManagerSet is a log parse operation binding the contract event 0xad45eab3ddda99c8b095504a58e25095a54c31dc487bb3eed9c0a9112ee4eaa8.
//
// Solidity: event USDonManagerSet(address indexed oldUSDonManager, address indexed newUSDonManager)
func (_Ondotokenmanager *OndotokenmanagerFilterer) ParseUSDonManagerSet(log types.Log) (*OndotokenmanagerUSDonManagerSet, error) {
	event := new(OndotokenmanagerUSDonManagerSet)
	if err := _Ondotokenmanager.contract.UnpackLog(event, "USDonManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
