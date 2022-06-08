// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tofunft

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
)

// MarketNGDetail is an auto generated low-level Go binding around an user-defined struct.
type MarketNGDetail struct {
	IntentionHash [32]byte
	Signer        common.Address
	TxDeadline    *big.Int
	Salt          [32]byte
	Id            *big.Int
	Opcode        uint8
	Caller        common.Address
	Currency      common.Address
	Price         *big.Int
	IncentiveRate *big.Int
	Settlement    MarketNGSettlement
	Bundle        []MarketNGTokenPair
	Deadline      *big.Int
}

// MarketNGIntention is an auto generated low-level Go binding around an user-defined struct.
type MarketNGIntention struct {
	User     common.Address
	Bundle   []MarketNGTokenPair
	Currency common.Address
	Price    *big.Int
	Deadline *big.Int
	Salt     [32]byte
	Kind     uint8
}

// MarketNGInventory is an auto generated low-level Go binding around an user-defined struct.
type MarketNGInventory struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}

// MarketNGPair721 is an auto generated low-level Go binding around an user-defined struct.
type MarketNGPair721 struct {
	Token   common.Address
	TokenId *big.Int
}

// MarketNGSettlement is an auto generated low-level Go binding around an user-defined struct.
type MarketNGSettlement struct {
	Coupons           []*big.Int
	FeeRate           *big.Int
	RoyaltyRate       *big.Int
	BuyerCashbackRate *big.Int
	FeeAddress        common.Address
	RoyaltyAddress    common.Address
}

// MarketNGSwap is an auto generated low-level Go binding around an user-defined struct.
type MarketNGSwap struct {
	Salt     [32]byte
	Creator  common.Address
	Deadline *big.Int
	Has      []MarketNGPair721
	Wants    []MarketNGPair721
}

// MarketNGTokenPair is an auto generated low-level Go binding around an user-defined struct.
type MarketNGTokenPair struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}

// TofunftMetaData contains all meta data concerning the Tofunft contract.
var TofunftMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"stateMutability\":\"nonpayable\",\"inputs\":[{\"type\":\"address\",\"name\":\"weth_\",\"internalType\":\"contractIWETH\"}]},{\"type\":\"event\",\"name\":\"EvAuctionRefund\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"address\",\"name\":\"bidder\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"uint256\",\"name\":\"refund\",\"internalType\":\"uint256\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EvCouponSpent\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"uint256\",\"name\":\"couponId\",\"internalType\":\"uint256\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EvInventoryUpdate\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\",\"indexed\":true},{\"type\":\"tuple\",\"name\":\"inventory\",\"internalType\":\"structMarketNG.Inventory\",\"indexed\":false,\"components\":[{\"type\":\"address\",\"name\":\"seller\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"buyer\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"netPrice\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"},{\"type\":\"uint8\",\"name\":\"status\",\"internalType\":\"uint8\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EvMarketSignerUpdate\",\"inputs\":[{\"type\":\"address\",\"name\":\"addr\",\"internalType\":\"address\",\"indexed\":false},{\"type\":\"bool\",\"name\":\"isRemoval\",\"internalType\":\"bool\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EvSettingsUpdated\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EvSwapped\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"req\",\"internalType\":\"structMarketNG.Swap\",\"indexed\":false,\"components\":[{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"creator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"tuple[]\",\"name\":\"has\",\"internalType\":\"structMarketNG.Pair721[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC721\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple[]\",\"name\":\"wants\",\"internalType\":\"structMarketNG.Pair721[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC721\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]}]},{\"type\":\"bytes\",\"name\":\"signature\",\"internalType\":\"bytes\",\"indexed\":false},{\"type\":\"address\",\"name\":\"swapper\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"type\":\"address\",\"name\":\"previousOwner\",\"internalType\":\"address\",\"indexed\":true},{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\",\"indexed\":true}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"type\":\"address\",\"name\":\"account\",\"internalType\":\"address\",\"indexed\":false}],\"anonymous\":false},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"KIND_AUCTION\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"KIND_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"KIND_SELL\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_ACCEPT_AUCTION\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_ACCEPT_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_BID\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_CANCEL_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_COMPLETE_AUCTION\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_COMPLETE_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_COMPLETE_SELL\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_MAX\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_MIN\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"OP_REJECT_BUY\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"RATE_BASE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"STATUS_CANCELLED\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"STATUS_DONE\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"STATUS_OPEN\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"TOKEN_1155\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"TOKEN_721\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint8\",\"name\":\"\",\"internalType\":\"uint8\"}],\"name\":\"TOKEN_MINT\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"cancelBuys\",\"inputs\":[{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"couponSpent\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"emergencyCancelAuction\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"bool\",\"name\":\"noBundle\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"hasInv\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"hasSignedIntention\",\"inputs\":[{\"type\":\"uint8\",\"name\":\"op\",\"internalType\":\"uint8\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"inCaseMoneyGetsStuck\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"seller\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"buyer\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"netPrice\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"},{\"type\":\"uint8\",\"name\":\"status\",\"internalType\":\"uint8\"}],\"name\":\"inventories\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"inventoryTokenCounts\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"mintData\",\"internalType\":\"bytes\"}],\"name\":\"inventoryTokens\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isAuction\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isAuctionOpen\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isBundleApproved\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"invId\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"owner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isBuy\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isBuyOpen\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isExpired\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isSell\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isSignatureValid\",\"inputs\":[{\"type\":\"bytes\",\"name\":\"signature\",\"internalType\":\"bytes\"},{\"type\":\"bytes32\",\"name\":\"hash\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"signer\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"isStatusOpen\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"marketSigners\",\"inputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"minAuctionDuration\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"uint256\",\"name\":\"\",\"internalType\":\"uint256\"}],\"name\":\"minAuctionIncrement\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}],\"name\":\"onERC1155BatchReceived\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"uint256[]\",\"name\":\"ids\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256[]\",\"name\":\"values\",\"internalType\":\"uint256[]\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}],\"name\":\"onERC1155Received\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"value\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bytes4\",\"name\":\"\",\"internalType\":\"bytes4\"}],\"name\":\"onERC721Received\",\"inputs\":[{\"type\":\"address\",\"name\":\"operator\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"from\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"bytes\",\"name\":\"data\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"address\"}],\"name\":\"owner\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"pause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"paused\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"renounceOwnership\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"payable\",\"outputs\":[],\"name\":\"run\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"intent\",\"internalType\":\"structMarketNG.Intention\",\"components\":[{\"type\":\"address\",\"name\":\"user\",\"internalType\":\"address\"},{\"type\":\"tuple[]\",\"name\":\"bundle\",\"internalType\":\"structMarketNG.TokenPair[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"mintData\",\"internalType\":\"bytes\"}]},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"}]},{\"type\":\"tuple\",\"name\":\"detail\",\"internalType\":\"structMarketNG.Detail\",\"components\":[{\"type\":\"bytes32\",\"name\":\"intentionHash\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"signer\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"txDeadline\",\"internalType\":\"uint256\"},{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"uint256\",\"name\":\"id\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"opcode\",\"internalType\":\"uint8\"},{\"type\":\"address\",\"name\":\"caller\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"currency\",\"internalType\":\"contractIERC20\"},{\"type\":\"uint256\",\"name\":\"price\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"incentiveRate\",\"internalType\":\"uint256\"},{\"type\":\"tuple\",\"name\":\"settlement\",\"internalType\":\"structMarketNG.Settlement\",\"components\":[{\"type\":\"uint256[]\",\"name\":\"coupons\",\"internalType\":\"uint256[]\"},{\"type\":\"uint256\",\"name\":\"feeRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"royaltyRate\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"buyerCashbackRate\",\"internalType\":\"uint256\"},{\"type\":\"address\",\"name\":\"feeAddress\",\"internalType\":\"address\"},{\"type\":\"address\",\"name\":\"royaltyAddress\",\"internalType\":\"address\"}]},{\"type\":\"tuple[]\",\"name\":\"bundle\",\"internalType\":\"structMarketNG.TokenPair[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"amount\",\"internalType\":\"uint256\"},{\"type\":\"uint8\",\"name\":\"kind\",\"internalType\":\"uint8\"},{\"type\":\"bytes\",\"name\":\"mintData\",\"internalType\":\"bytes\"}]},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"}]},{\"type\":\"bytes\",\"name\":\"sigIntent\",\"internalType\":\"bytes\"},{\"type\":\"bytes\",\"name\":\"sigDetail\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"send\",\"inputs\":[{\"type\":\"address\",\"name\":\"to\",\"internalType\":\"address\"},{\"type\":\"tuple[]\",\"name\":\"tokens\",\"internalType\":\"structMarketNG.Pair721[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC721\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]}]},{\"type\":\"function\",\"stateMutability\":\"pure\",\"outputs\":[{\"type\":\"bool\",\"name\":\"\",\"internalType\":\"bool\"}],\"name\":\"supportsInterface\",\"inputs\":[{\"type\":\"bytes4\",\"name\":\"interfaceId\",\"internalType\":\"bytes4\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"swap\",\"inputs\":[{\"type\":\"tuple\",\"name\":\"req\",\"internalType\":\"structMarketNG.Swap\",\"components\":[{\"type\":\"bytes32\",\"name\":\"salt\",\"internalType\":\"bytes32\"},{\"type\":\"address\",\"name\":\"creator\",\"internalType\":\"address\"},{\"type\":\"uint256\",\"name\":\"deadline\",\"internalType\":\"uint256\"},{\"type\":\"tuple[]\",\"name\":\"has\",\"internalType\":\"structMarketNG.Pair721[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC721\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]},{\"type\":\"tuple[]\",\"name\":\"wants\",\"internalType\":\"structMarketNG.Pair721[]\",\"components\":[{\"type\":\"address\",\"name\":\"token\",\"internalType\":\"contractIERC721\"},{\"type\":\"uint256\",\"name\":\"tokenId\",\"internalType\":\"uint256\"}]}]},{\"type\":\"bytes\",\"name\":\"signature\",\"internalType\":\"bytes\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"transferOwnership\",\"inputs\":[{\"type\":\"address\",\"name\":\"newOwner\",\"internalType\":\"address\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"unpause\",\"inputs\":[]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateSettings\",\"inputs\":[{\"type\":\"uint256\",\"name\":\"minAuctionIncrement_\",\"internalType\":\"uint256\"},{\"type\":\"uint256\",\"name\":\"minAuctionDuration_\",\"internalType\":\"uint256\"}]},{\"type\":\"function\",\"stateMutability\":\"nonpayable\",\"outputs\":[],\"name\":\"updateSigner\",\"inputs\":[{\"type\":\"address\",\"name\":\"addr\",\"internalType\":\"address\"},{\"type\":\"bool\",\"name\":\"remove\",\"internalType\":\"bool\"}]},{\"type\":\"function\",\"stateMutability\":\"view\",\"outputs\":[{\"type\":\"address\",\"name\":\"\",\"internalType\":\"contractIWETH\"}],\"name\":\"weth\",\"inputs\":[]},{\"type\":\"receive\",\"stateMutability\":\"payable\"}]",
}

// TofunftABI is the input ABI used to generate the binding from.
// Deprecated: Use TofunftMetaData.ABI instead.
var TofunftABI = TofunftMetaData.ABI

// Tofunft is an auto generated Go binding around an Ethereum contract.
type Tofunft struct {
	TofunftCaller     // Read-only binding to the contract
	TofunftTransactor // Write-only binding to the contract
	TofunftFilterer   // Log filterer for contract events
}

// TofunftCaller is an auto generated read-only Go binding around an Ethereum contract.
type TofunftCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TofunftTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TofunftTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TofunftFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TofunftFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TofunftSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TofunftSession struct {
	Contract     *Tofunft          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TofunftCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TofunftCallerSession struct {
	Contract *TofunftCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// TofunftTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TofunftTransactorSession struct {
	Contract     *TofunftTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// TofunftRaw is an auto generated low-level Go binding around an Ethereum contract.
type TofunftRaw struct {
	Contract *Tofunft // Generic contract binding to access the raw methods on
}

// TofunftCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TofunftCallerRaw struct {
	Contract *TofunftCaller // Generic read-only contract binding to access the raw methods on
}

// TofunftTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TofunftTransactorRaw struct {
	Contract *TofunftTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTofunft creates a new instance of Tofunft, bound to a specific deployed contract.
func NewTofunft(address common.Address, backend bind.ContractBackend) (*Tofunft, error) {
	contract, err := bindTofunft(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Tofunft{TofunftCaller: TofunftCaller{contract: contract}, TofunftTransactor: TofunftTransactor{contract: contract}, TofunftFilterer: TofunftFilterer{contract: contract}}, nil
}

// NewTofunftCaller creates a new read-only instance of Tofunft, bound to a specific deployed contract.
func NewTofunftCaller(address common.Address, caller bind.ContractCaller) (*TofunftCaller, error) {
	contract, err := bindTofunft(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TofunftCaller{contract: contract}, nil
}

// NewTofunftTransactor creates a new write-only instance of Tofunft, bound to a specific deployed contract.
func NewTofunftTransactor(address common.Address, transactor bind.ContractTransactor) (*TofunftTransactor, error) {
	contract, err := bindTofunft(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TofunftTransactor{contract: contract}, nil
}

// NewTofunftFilterer creates a new log filterer instance of Tofunft, bound to a specific deployed contract.
func NewTofunftFilterer(address common.Address, filterer bind.ContractFilterer) (*TofunftFilterer, error) {
	contract, err := bindTofunft(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TofunftFilterer{contract: contract}, nil
}

// bindTofunft binds a generic wrapper to an already deployed contract.
func bindTofunft(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TofunftABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tofunft *TofunftRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tofunft.Contract.TofunftCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tofunft *TofunftRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.Contract.TofunftTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tofunft *TofunftRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tofunft.Contract.TofunftTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Tofunft *TofunftCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Tofunft.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Tofunft *TofunftTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Tofunft *TofunftTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Tofunft.Contract.contract.Transact(opts, method, params...)
}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCaller) KINDAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "KIND_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Tofunft *TofunftSession) KINDAUCTION() (uint8, error) {
	return _Tofunft.Contract.KINDAUCTION(&_Tofunft.CallOpts)
}

// KINDAUCTION is a free data retrieval call binding the contract method 0x7234d8f2.
//
// Solidity: function KIND_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCallerSession) KINDAUCTION() (uint8, error) {
	return _Tofunft.Contract.KINDAUCTION(&_Tofunft.CallOpts)
}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) KINDBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "KIND_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) KINDBUY() (uint8, error) {
	return _Tofunft.Contract.KINDBUY(&_Tofunft.CallOpts)
}

// KINDBUY is a free data retrieval call binding the contract method 0xe1784a02.
//
// Solidity: function KIND_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) KINDBUY() (uint8, error) {
	return _Tofunft.Contract.KINDBUY(&_Tofunft.CallOpts)
}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Tofunft *TofunftCaller) KINDSELL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "KIND_SELL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Tofunft *TofunftSession) KINDSELL() (uint8, error) {
	return _Tofunft.Contract.KINDSELL(&_Tofunft.CallOpts)
}

// KINDSELL is a free data retrieval call binding the contract method 0x25593ac2.
//
// Solidity: function KIND_SELL() view returns(uint8)
func (_Tofunft *TofunftCallerSession) KINDSELL() (uint8, error) {
	return _Tofunft.Contract.KINDSELL(&_Tofunft.CallOpts)
}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCaller) OPACCEPTAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_ACCEPT_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Tofunft *TofunftSession) OPACCEPTAUCTION() (uint8, error) {
	return _Tofunft.Contract.OPACCEPTAUCTION(&_Tofunft.CallOpts)
}

// OPACCEPTAUCTION is a free data retrieval call binding the contract method 0x7ae1ace0.
//
// Solidity: function OP_ACCEPT_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPACCEPTAUCTION() (uint8, error) {
	return _Tofunft.Contract.OPACCEPTAUCTION(&_Tofunft.CallOpts)
}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) OPACCEPTBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_ACCEPT_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) OPACCEPTBUY() (uint8, error) {
	return _Tofunft.Contract.OPACCEPTBUY(&_Tofunft.CallOpts)
}

// OPACCEPTBUY is a free data retrieval call binding the contract method 0x11f0794c.
//
// Solidity: function OP_ACCEPT_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPACCEPTBUY() (uint8, error) {
	return _Tofunft.Contract.OPACCEPTBUY(&_Tofunft.CallOpts)
}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Tofunft *TofunftCaller) OPBID(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_BID")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Tofunft *TofunftSession) OPBID() (uint8, error) {
	return _Tofunft.Contract.OPBID(&_Tofunft.CallOpts)
}

// OPBID is a free data retrieval call binding the contract method 0x81787a85.
//
// Solidity: function OP_BID() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPBID() (uint8, error) {
	return _Tofunft.Contract.OPBID(&_Tofunft.CallOpts)
}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) OPBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) OPBUY() (uint8, error) {
	return _Tofunft.Contract.OPBUY(&_Tofunft.CallOpts)
}

// OPBUY is a free data retrieval call binding the contract method 0xeb374261.
//
// Solidity: function OP_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPBUY() (uint8, error) {
	return _Tofunft.Contract.OPBUY(&_Tofunft.CallOpts)
}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) OPCANCELBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_CANCEL_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) OPCANCELBUY() (uint8, error) {
	return _Tofunft.Contract.OPCANCELBUY(&_Tofunft.CallOpts)
}

// OPCANCELBUY is a free data retrieval call binding the contract method 0x9e57feb5.
//
// Solidity: function OP_CANCEL_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPCANCELBUY() (uint8, error) {
	return _Tofunft.Contract.OPCANCELBUY(&_Tofunft.CallOpts)
}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCaller) OPCOMPLETEAUCTION(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_COMPLETE_AUCTION")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Tofunft *TofunftSession) OPCOMPLETEAUCTION() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETEAUCTION(&_Tofunft.CallOpts)
}

// OPCOMPLETEAUCTION is a free data retrieval call binding the contract method 0x6acc65db.
//
// Solidity: function OP_COMPLETE_AUCTION() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPCOMPLETEAUCTION() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETEAUCTION(&_Tofunft.CallOpts)
}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) OPCOMPLETEBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_COMPLETE_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) OPCOMPLETEBUY() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETEBUY(&_Tofunft.CallOpts)
}

// OPCOMPLETEBUY is a free data retrieval call binding the contract method 0xb50a2a55.
//
// Solidity: function OP_COMPLETE_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPCOMPLETEBUY() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETEBUY(&_Tofunft.CallOpts)
}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Tofunft *TofunftCaller) OPCOMPLETESELL(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_COMPLETE_SELL")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Tofunft *TofunftSession) OPCOMPLETESELL() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETESELL(&_Tofunft.CallOpts)
}

// OPCOMPLETESELL is a free data retrieval call binding the contract method 0x8f18439e.
//
// Solidity: function OP_COMPLETE_SELL() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPCOMPLETESELL() (uint8, error) {
	return _Tofunft.Contract.OPCOMPLETESELL(&_Tofunft.CallOpts)
}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Tofunft *TofunftCaller) OPMAX(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_MAX")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Tofunft *TofunftSession) OPMAX() (uint8, error) {
	return _Tofunft.Contract.OPMAX(&_Tofunft.CallOpts)
}

// OPMAX is a free data retrieval call binding the contract method 0xf0954160.
//
// Solidity: function OP_MAX() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPMAX() (uint8, error) {
	return _Tofunft.Contract.OPMAX(&_Tofunft.CallOpts)
}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Tofunft *TofunftCaller) OPMIN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_MIN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Tofunft *TofunftSession) OPMIN() (uint8, error) {
	return _Tofunft.Contract.OPMIN(&_Tofunft.CallOpts)
}

// OPMIN is a free data retrieval call binding the contract method 0x90c2b10e.
//
// Solidity: function OP_MIN() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPMIN() (uint8, error) {
	return _Tofunft.Contract.OPMIN(&_Tofunft.CallOpts)
}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Tofunft *TofunftCaller) OPREJECTBUY(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "OP_REJECT_BUY")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Tofunft *TofunftSession) OPREJECTBUY() (uint8, error) {
	return _Tofunft.Contract.OPREJECTBUY(&_Tofunft.CallOpts)
}

// OPREJECTBUY is a free data retrieval call binding the contract method 0x1bb03ca9.
//
// Solidity: function OP_REJECT_BUY() view returns(uint8)
func (_Tofunft *TofunftCallerSession) OPREJECTBUY() (uint8, error) {
	return _Tofunft.Contract.OPREJECTBUY(&_Tofunft.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Tofunft *TofunftCaller) RATEBASE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "RATE_BASE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Tofunft *TofunftSession) RATEBASE() (*big.Int, error) {
	return _Tofunft.Contract.RATEBASE(&_Tofunft.CallOpts)
}

// RATEBASE is a free data retrieval call binding the contract method 0x0873c6ec.
//
// Solidity: function RATE_BASE() view returns(uint256)
func (_Tofunft *TofunftCallerSession) RATEBASE() (*big.Int, error) {
	return _Tofunft.Contract.RATEBASE(&_Tofunft.CallOpts)
}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Tofunft *TofunftCaller) STATUSCANCELLED(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "STATUS_CANCELLED")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Tofunft *TofunftSession) STATUSCANCELLED() (uint8, error) {
	return _Tofunft.Contract.STATUSCANCELLED(&_Tofunft.CallOpts)
}

// STATUSCANCELLED is a free data retrieval call binding the contract method 0x5a4e5a15.
//
// Solidity: function STATUS_CANCELLED() view returns(uint8)
func (_Tofunft *TofunftCallerSession) STATUSCANCELLED() (uint8, error) {
	return _Tofunft.Contract.STATUSCANCELLED(&_Tofunft.CallOpts)
}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Tofunft *TofunftCaller) STATUSDONE(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "STATUS_DONE")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Tofunft *TofunftSession) STATUSDONE() (uint8, error) {
	return _Tofunft.Contract.STATUSDONE(&_Tofunft.CallOpts)
}

// STATUSDONE is a free data retrieval call binding the contract method 0x740db280.
//
// Solidity: function STATUS_DONE() view returns(uint8)
func (_Tofunft *TofunftCallerSession) STATUSDONE() (uint8, error) {
	return _Tofunft.Contract.STATUSDONE(&_Tofunft.CallOpts)
}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Tofunft *TofunftCaller) STATUSOPEN(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "STATUS_OPEN")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Tofunft *TofunftSession) STATUSOPEN() (uint8, error) {
	return _Tofunft.Contract.STATUSOPEN(&_Tofunft.CallOpts)
}

// STATUSOPEN is a free data retrieval call binding the contract method 0x24f8515b.
//
// Solidity: function STATUS_OPEN() view returns(uint8)
func (_Tofunft *TofunftCallerSession) STATUSOPEN() (uint8, error) {
	return _Tofunft.Contract.STATUSOPEN(&_Tofunft.CallOpts)
}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Tofunft *TofunftCaller) TOKEN1155(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "TOKEN_1155")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Tofunft *TofunftSession) TOKEN1155() (uint8, error) {
	return _Tofunft.Contract.TOKEN1155(&_Tofunft.CallOpts)
}

// TOKEN1155 is a free data retrieval call binding the contract method 0xf0d250ba.
//
// Solidity: function TOKEN_1155() view returns(uint8)
func (_Tofunft *TofunftCallerSession) TOKEN1155() (uint8, error) {
	return _Tofunft.Contract.TOKEN1155(&_Tofunft.CallOpts)
}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Tofunft *TofunftCaller) TOKEN721(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "TOKEN_721")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Tofunft *TofunftSession) TOKEN721() (uint8, error) {
	return _Tofunft.Contract.TOKEN721(&_Tofunft.CallOpts)
}

// TOKEN721 is a free data retrieval call binding the contract method 0xc477be20.
//
// Solidity: function TOKEN_721() view returns(uint8)
func (_Tofunft *TofunftCallerSession) TOKEN721() (uint8, error) {
	return _Tofunft.Contract.TOKEN721(&_Tofunft.CallOpts)
}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Tofunft *TofunftCaller) TOKENMINT(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "TOKEN_MINT")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Tofunft *TofunftSession) TOKENMINT() (uint8, error) {
	return _Tofunft.Contract.TOKENMINT(&_Tofunft.CallOpts)
}

// TOKENMINT is a free data retrieval call binding the contract method 0x853ca41a.
//
// Solidity: function TOKEN_MINT() view returns(uint8)
func (_Tofunft *TofunftCallerSession) TOKENMINT() (uint8, error) {
	return _Tofunft.Contract.TOKENMINT(&_Tofunft.CallOpts)
}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Tofunft *TofunftCaller) CouponSpent(opts *bind.CallOpts, arg0 *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "couponSpent", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Tofunft *TofunftSession) CouponSpent(arg0 *big.Int) (bool, error) {
	return _Tofunft.Contract.CouponSpent(&_Tofunft.CallOpts, arg0)
}

// CouponSpent is a free data retrieval call binding the contract method 0x3ed9ffb7.
//
// Solidity: function couponSpent(uint256 ) view returns(bool)
func (_Tofunft *TofunftCallerSession) CouponSpent(arg0 *big.Int) (bool, error) {
	return _Tofunft.Contract.CouponSpent(&_Tofunft.CallOpts, arg0)
}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) HasInv(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "hasInv", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) HasInv(id *big.Int) (bool, error) {
	return _Tofunft.Contract.HasInv(&_Tofunft.CallOpts, id)
}

// HasInv is a free data retrieval call binding the contract method 0xf5116bc9.
//
// Solidity: function hasInv(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) HasInv(id *big.Int) (bool, error) {
	return _Tofunft.Contract.HasInv(&_Tofunft.CallOpts, id)
}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Tofunft *TofunftCaller) HasSignedIntention(opts *bind.CallOpts, op uint8) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "hasSignedIntention", op)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Tofunft *TofunftSession) HasSignedIntention(op uint8) (bool, error) {
	return _Tofunft.Contract.HasSignedIntention(&_Tofunft.CallOpts, op)
}

// HasSignedIntention is a free data retrieval call binding the contract method 0xac5e2cb1.
//
// Solidity: function hasSignedIntention(uint8 op) pure returns(bool)
func (_Tofunft *TofunftCallerSession) HasSignedIntention(op uint8) (bool, error) {
	return _Tofunft.Contract.HasSignedIntention(&_Tofunft.CallOpts, op)
}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Tofunft *TofunftCaller) Inventories(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "inventories", arg0)

	outstruct := new(struct {
		Seller   common.Address
		Buyer    common.Address
		Currency common.Address
		Price    *big.Int
		NetPrice *big.Int
		Deadline *big.Int
		Kind     uint8
		Status   uint8
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Seller = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Buyer = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.Currency = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Price = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.NetPrice = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Deadline = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Status = *abi.ConvertType(out[7], new(uint8)).(*uint8)

	return *outstruct, err

}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Tofunft *TofunftSession) Inventories(arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	return _Tofunft.Contract.Inventories(&_Tofunft.CallOpts, arg0)
}

// Inventories is a free data retrieval call binding the contract method 0xcd78ba01.
//
// Solidity: function inventories(uint256 ) view returns(address seller, address buyer, address currency, uint256 price, uint256 netPrice, uint256 deadline, uint8 kind, uint8 status)
func (_Tofunft *TofunftCallerSession) Inventories(arg0 *big.Int) (struct {
	Seller   common.Address
	Buyer    common.Address
	Currency common.Address
	Price    *big.Int
	NetPrice *big.Int
	Deadline *big.Int
	Kind     uint8
	Status   uint8
}, error) {
	return _Tofunft.Contract.Inventories(&_Tofunft.CallOpts, arg0)
}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Tofunft *TofunftCaller) InventoryTokenCounts(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "inventoryTokenCounts", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Tofunft *TofunftSession) InventoryTokenCounts(arg0 *big.Int) (*big.Int, error) {
	return _Tofunft.Contract.InventoryTokenCounts(&_Tofunft.CallOpts, arg0)
}

// InventoryTokenCounts is a free data retrieval call binding the contract method 0x5fd34298.
//
// Solidity: function inventoryTokenCounts(uint256 ) view returns(uint256)
func (_Tofunft *TofunftCallerSession) InventoryTokenCounts(arg0 *big.Int) (*big.Int, error) {
	return _Tofunft.Contract.InventoryTokenCounts(&_Tofunft.CallOpts, arg0)
}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Tofunft *TofunftCaller) InventoryTokens(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "inventoryTokens", arg0, arg1)

	outstruct := new(struct {
		Token    common.Address
		TokenId  *big.Int
		Amount   *big.Int
		Kind     uint8
		MintData []byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Token = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.TokenId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Amount = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Kind = *abi.ConvertType(out[3], new(uint8)).(*uint8)
	outstruct.MintData = *abi.ConvertType(out[4], new([]byte)).(*[]byte)

	return *outstruct, err

}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Tofunft *TofunftSession) InventoryTokens(arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	return _Tofunft.Contract.InventoryTokens(&_Tofunft.CallOpts, arg0, arg1)
}

// InventoryTokens is a free data retrieval call binding the contract method 0xb4533aad.
//
// Solidity: function inventoryTokens(uint256 , uint256 ) view returns(address token, uint256 tokenId, uint256 amount, uint8 kind, bytes mintData)
func (_Tofunft *TofunftCallerSession) InventoryTokens(arg0 *big.Int, arg1 *big.Int) (struct {
	Token    common.Address
	TokenId  *big.Int
	Amount   *big.Int
	Kind     uint8
	MintData []byte
}, error) {
	return _Tofunft.Contract.InventoryTokens(&_Tofunft.CallOpts, arg0, arg1)
}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsAuction(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isAuction", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsAuction(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsAuction(&_Tofunft.CallOpts, id)
}

// IsAuction is a free data retrieval call binding the contract method 0x8704f2a3.
//
// Solidity: function isAuction(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsAuction(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsAuction(&_Tofunft.CallOpts, id)
}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsAuctionOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isAuctionOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsAuctionOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsAuctionOpen(&_Tofunft.CallOpts, id)
}

// IsAuctionOpen is a free data retrieval call binding the contract method 0x0ad48628.
//
// Solidity: function isAuctionOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsAuctionOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsAuctionOpen(&_Tofunft.CallOpts, id)
}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Tofunft *TofunftCaller) IsBundleApproved(opts *bind.CallOpts, invId *big.Int, owner common.Address) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isBundleApproved", invId, owner)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Tofunft *TofunftSession) IsBundleApproved(invId *big.Int, owner common.Address) (bool, error) {
	return _Tofunft.Contract.IsBundleApproved(&_Tofunft.CallOpts, invId, owner)
}

// IsBundleApproved is a free data retrieval call binding the contract method 0xf4a33e0d.
//
// Solidity: function isBundleApproved(uint256 invId, address owner) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsBundleApproved(invId *big.Int, owner common.Address) (bool, error) {
	return _Tofunft.Contract.IsBundleApproved(&_Tofunft.CallOpts, invId, owner)
}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsBuy(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isBuy", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsBuy(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsBuy(&_Tofunft.CallOpts, id)
}

// IsBuy is a free data retrieval call binding the contract method 0xa80d33fb.
//
// Solidity: function isBuy(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsBuy(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsBuy(&_Tofunft.CallOpts, id)
}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsBuyOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isBuyOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsBuyOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsBuyOpen(&_Tofunft.CallOpts, id)
}

// IsBuyOpen is a free data retrieval call binding the contract method 0xbdf52b45.
//
// Solidity: function isBuyOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsBuyOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsBuyOpen(&_Tofunft.CallOpts, id)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsExpired(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isExpired", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsExpired(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsExpired(&_Tofunft.CallOpts, id)
}

// IsExpired is a free data retrieval call binding the contract method 0xd9548e53.
//
// Solidity: function isExpired(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsExpired(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsExpired(&_Tofunft.CallOpts, id)
}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsSell(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isSell", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsSell(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsSell(&_Tofunft.CallOpts, id)
}

// IsSell is a free data retrieval call binding the contract method 0x1b01e72c.
//
// Solidity: function isSell(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsSell(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsSell(&_Tofunft.CallOpts, id)
}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Tofunft *TofunftCaller) IsSignatureValid(opts *bind.CallOpts, signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isSignatureValid", signature, hash, signer)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Tofunft *TofunftSession) IsSignatureValid(signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	return _Tofunft.Contract.IsSignatureValid(&_Tofunft.CallOpts, signature, hash, signer)
}

// IsSignatureValid is a free data retrieval call binding the contract method 0x781dc70a.
//
// Solidity: function isSignatureValid(bytes signature, bytes32 hash, address signer) pure returns(bool)
func (_Tofunft *TofunftCallerSession) IsSignatureValid(signature []byte, hash [32]byte, signer common.Address) (bool, error) {
	return _Tofunft.Contract.IsSignatureValid(&_Tofunft.CallOpts, signature, hash, signer)
}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCaller) IsStatusOpen(opts *bind.CallOpts, id *big.Int) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "isStatusOpen", id)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftSession) IsStatusOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsStatusOpen(&_Tofunft.CallOpts, id)
}

// IsStatusOpen is a free data retrieval call binding the contract method 0xee98ce91.
//
// Solidity: function isStatusOpen(uint256 id) view returns(bool)
func (_Tofunft *TofunftCallerSession) IsStatusOpen(id *big.Int) (bool, error) {
	return _Tofunft.Contract.IsStatusOpen(&_Tofunft.CallOpts, id)
}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Tofunft *TofunftCaller) MarketSigners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "marketSigners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Tofunft *TofunftSession) MarketSigners(arg0 common.Address) (bool, error) {
	return _Tofunft.Contract.MarketSigners(&_Tofunft.CallOpts, arg0)
}

// MarketSigners is a free data retrieval call binding the contract method 0x2bcd27df.
//
// Solidity: function marketSigners(address ) view returns(bool)
func (_Tofunft *TofunftCallerSession) MarketSigners(arg0 common.Address) (bool, error) {
	return _Tofunft.Contract.MarketSigners(&_Tofunft.CallOpts, arg0)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Tofunft *TofunftCaller) MinAuctionDuration(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "minAuctionDuration")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Tofunft *TofunftSession) MinAuctionDuration() (*big.Int, error) {
	return _Tofunft.Contract.MinAuctionDuration(&_Tofunft.CallOpts)
}

// MinAuctionDuration is a free data retrieval call binding the contract method 0x54134876.
//
// Solidity: function minAuctionDuration() view returns(uint256)
func (_Tofunft *TofunftCallerSession) MinAuctionDuration() (*big.Int, error) {
	return _Tofunft.Contract.MinAuctionDuration(&_Tofunft.CallOpts)
}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Tofunft *TofunftCaller) MinAuctionIncrement(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "minAuctionIncrement")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Tofunft *TofunftSession) MinAuctionIncrement() (*big.Int, error) {
	return _Tofunft.Contract.MinAuctionIncrement(&_Tofunft.CallOpts)
}

// MinAuctionIncrement is a free data retrieval call binding the contract method 0x708d4d35.
//
// Solidity: function minAuctionIncrement() view returns(uint256)
func (_Tofunft *TofunftCallerSession) MinAuctionIncrement() (*big.Int, error) {
	return _Tofunft.Contract.MinAuctionIncrement(&_Tofunft.CallOpts)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCaller) OnERC1155BatchReceived(opts *bind.CallOpts, operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "onERC1155BatchReceived", operator, from, ids, values, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Tofunft *TofunftSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC1155BatchReceived(&_Tofunft.CallOpts, operator, from, ids, values, data)
}

// OnERC1155BatchReceived is a free data retrieval call binding the contract method 0xbc197c81.
//
// Solidity: function onERC1155BatchReceived(address operator, address from, uint256[] ids, uint256[] values, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCallerSession) OnERC1155BatchReceived(operator common.Address, from common.Address, ids []*big.Int, values []*big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC1155BatchReceived(&_Tofunft.CallOpts, operator, from, ids, values, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCaller) OnERC1155Received(opts *bind.CallOpts, operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "onERC1155Received", operator, from, id, value, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Tofunft *TofunftSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC1155Received(&_Tofunft.CallOpts, operator, from, id, value, data)
}

// OnERC1155Received is a free data retrieval call binding the contract method 0xf23a6e61.
//
// Solidity: function onERC1155Received(address operator, address from, uint256 id, uint256 value, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCallerSession) OnERC1155Received(operator common.Address, from common.Address, id *big.Int, value *big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC1155Received(&_Tofunft.CallOpts, operator, from, id, value, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCaller) OnERC721Received(opts *bind.CallOpts, operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "onERC721Received", operator, from, tokenId, data)

	if err != nil {
		return *new([4]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([4]byte)).(*[4]byte)

	return out0, err

}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Tofunft *TofunftSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC721Received(&_Tofunft.CallOpts, operator, from, tokenId, data)
}

// OnERC721Received is a free data retrieval call binding the contract method 0x150b7a02.
//
// Solidity: function onERC721Received(address operator, address from, uint256 tokenId, bytes data) view returns(bytes4)
func (_Tofunft *TofunftCallerSession) OnERC721Received(operator common.Address, from common.Address, tokenId *big.Int, data []byte) ([4]byte, error) {
	return _Tofunft.Contract.OnERC721Received(&_Tofunft.CallOpts, operator, from, tokenId, data)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tofunft *TofunftCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tofunft *TofunftSession) Owner() (common.Address, error) {
	return _Tofunft.Contract.Owner(&_Tofunft.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Tofunft *TofunftCallerSession) Owner() (common.Address, error) {
	return _Tofunft.Contract.Owner(&_Tofunft.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tofunft *TofunftCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tofunft *TofunftSession) Paused() (bool, error) {
	return _Tofunft.Contract.Paused(&_Tofunft.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Tofunft *TofunftCallerSession) Paused() (bool, error) {
	return _Tofunft.Contract.Paused(&_Tofunft.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Tofunft *TofunftCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Tofunft *TofunftSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tofunft.Contract.SupportsInterface(&_Tofunft.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) pure returns(bool)
func (_Tofunft *TofunftCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Tofunft.Contract.SupportsInterface(&_Tofunft.CallOpts, interfaceId)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Tofunft *TofunftCaller) Weth(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Tofunft.contract.Call(opts, &out, "weth")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Tofunft *TofunftSession) Weth() (common.Address, error) {
	return _Tofunft.Contract.Weth(&_Tofunft.CallOpts)
}

// Weth is a free data retrieval call binding the contract method 0x3fc8cef3.
//
// Solidity: function weth() view returns(address)
func (_Tofunft *TofunftCallerSession) Weth() (common.Address, error) {
	return _Tofunft.Contract.Weth(&_Tofunft.CallOpts)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Tofunft *TofunftTransactor) CancelBuys(opts *bind.TransactOpts, ids []*big.Int) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "cancelBuys", ids)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Tofunft *TofunftSession) CancelBuys(ids []*big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.CancelBuys(&_Tofunft.TransactOpts, ids)
}

// CancelBuys is a paid mutator transaction binding the contract method 0xc1c30e80.
//
// Solidity: function cancelBuys(uint256[] ids) returns()
func (_Tofunft *TofunftTransactorSession) CancelBuys(ids []*big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.CancelBuys(&_Tofunft.TransactOpts, ids)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Tofunft *TofunftTransactor) EmergencyCancelAuction(opts *bind.TransactOpts, id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "emergencyCancelAuction", id, noBundle)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Tofunft *TofunftSession) EmergencyCancelAuction(id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Tofunft.Contract.EmergencyCancelAuction(&_Tofunft.TransactOpts, id, noBundle)
}

// EmergencyCancelAuction is a paid mutator transaction binding the contract method 0xe7d4a999.
//
// Solidity: function emergencyCancelAuction(uint256 id, bool noBundle) returns()
func (_Tofunft *TofunftTransactorSession) EmergencyCancelAuction(id *big.Int, noBundle bool) (*types.Transaction, error) {
	return _Tofunft.Contract.EmergencyCancelAuction(&_Tofunft.TransactOpts, id, noBundle)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Tofunft *TofunftTransactor) InCaseMoneyGetsStuck(opts *bind.TransactOpts, to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "inCaseMoneyGetsStuck", to, currency, amount)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Tofunft *TofunftSession) InCaseMoneyGetsStuck(to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.InCaseMoneyGetsStuck(&_Tofunft.TransactOpts, to, currency, amount)
}

// InCaseMoneyGetsStuck is a paid mutator transaction binding the contract method 0x80bc688f.
//
// Solidity: function inCaseMoneyGetsStuck(address to, address currency, uint256 amount) returns()
func (_Tofunft *TofunftTransactorSession) InCaseMoneyGetsStuck(to common.Address, currency common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.InCaseMoneyGetsStuck(&_Tofunft.TransactOpts, to, currency, amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tofunft *TofunftTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tofunft *TofunftSession) Pause() (*types.Transaction, error) {
	return _Tofunft.Contract.Pause(&_Tofunft.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Tofunft *TofunftTransactorSession) Pause() (*types.Transaction, error) {
	return _Tofunft.Contract.Pause(&_Tofunft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tofunft *TofunftTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tofunft *TofunftSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tofunft.Contract.RenounceOwnership(&_Tofunft.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Tofunft *TofunftTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Tofunft.Contract.RenounceOwnership(&_Tofunft.TransactOpts)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Tofunft *TofunftTransactor) Run(opts *bind.TransactOpts, intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "run", intent, detail, sigIntent, sigDetail)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Tofunft *TofunftSession) Run(intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Tofunft.Contract.Run(&_Tofunft.TransactOpts, intent, detail, sigIntent, sigDetail)
}

// Run is a paid mutator transaction binding the contract method 0xba847759.
//
// Solidity: function run((address,(address,uint256,uint256,uint8,bytes)[],address,uint256,uint256,bytes32,uint8) intent, (bytes32,address,uint256,bytes32,uint256,uint8,address,address,uint256,uint256,(uint256[],uint256,uint256,uint256,address,address),(address,uint256,uint256,uint8,bytes)[],uint256) detail, bytes sigIntent, bytes sigDetail) payable returns()
func (_Tofunft *TofunftTransactorSession) Run(intent MarketNGIntention, detail MarketNGDetail, sigIntent []byte, sigDetail []byte) (*types.Transaction, error) {
	return _Tofunft.Contract.Run(&_Tofunft.TransactOpts, intent, detail, sigIntent, sigDetail)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Tofunft *TofunftTransactor) Send(opts *bind.TransactOpts, to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "send", to, tokens)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Tofunft *TofunftSession) Send(to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Tofunft.Contract.Send(&_Tofunft.TransactOpts, to, tokens)
}

// Send is a paid mutator transaction binding the contract method 0xafd76a0b.
//
// Solidity: function send(address to, (address,uint256)[] tokens) returns()
func (_Tofunft *TofunftTransactorSession) Send(to common.Address, tokens []MarketNGPair721) (*types.Transaction, error) {
	return _Tofunft.Contract.Send(&_Tofunft.TransactOpts, to, tokens)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Tofunft *TofunftTransactor) Swap(opts *bind.TransactOpts, req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "swap", req, signature)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Tofunft *TofunftSession) Swap(req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Tofunft.Contract.Swap(&_Tofunft.TransactOpts, req, signature)
}

// Swap is a paid mutator transaction binding the contract method 0xe91274f3.
//
// Solidity: function swap((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature) returns()
func (_Tofunft *TofunftTransactorSession) Swap(req MarketNGSwap, signature []byte) (*types.Transaction, error) {
	return _Tofunft.Contract.Swap(&_Tofunft.TransactOpts, req, signature)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tofunft *TofunftTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tofunft *TofunftSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tofunft.Contract.TransferOwnership(&_Tofunft.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Tofunft *TofunftTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Tofunft.Contract.TransferOwnership(&_Tofunft.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tofunft *TofunftTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tofunft *TofunftSession) Unpause() (*types.Transaction, error) {
	return _Tofunft.Contract.Unpause(&_Tofunft.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Tofunft *TofunftTransactorSession) Unpause() (*types.Transaction, error) {
	return _Tofunft.Contract.Unpause(&_Tofunft.TransactOpts)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Tofunft *TofunftTransactor) UpdateSettings(opts *bind.TransactOpts, minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "updateSettings", minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Tofunft *TofunftSession) UpdateSettings(minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.UpdateSettings(&_Tofunft.TransactOpts, minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x015af8ee.
//
// Solidity: function updateSettings(uint256 minAuctionIncrement_, uint256 minAuctionDuration_) returns()
func (_Tofunft *TofunftTransactorSession) UpdateSettings(minAuctionIncrement_ *big.Int, minAuctionDuration_ *big.Int) (*types.Transaction, error) {
	return _Tofunft.Contract.UpdateSettings(&_Tofunft.TransactOpts, minAuctionIncrement_, minAuctionDuration_)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Tofunft *TofunftTransactor) UpdateSigner(opts *bind.TransactOpts, addr common.Address, remove bool) (*types.Transaction, error) {
	return _Tofunft.contract.Transact(opts, "updateSigner", addr, remove)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Tofunft *TofunftSession) UpdateSigner(addr common.Address, remove bool) (*types.Transaction, error) {
	return _Tofunft.Contract.UpdateSigner(&_Tofunft.TransactOpts, addr, remove)
}

// UpdateSigner is a paid mutator transaction binding the contract method 0xf460590b.
//
// Solidity: function updateSigner(address addr, bool remove) returns()
func (_Tofunft *TofunftTransactorSession) UpdateSigner(addr common.Address, remove bool) (*types.Transaction, error) {
	return _Tofunft.Contract.UpdateSigner(&_Tofunft.TransactOpts, addr, remove)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tofunft *TofunftTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Tofunft.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tofunft *TofunftSession) Receive() (*types.Transaction, error) {
	return _Tofunft.Contract.Receive(&_Tofunft.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_Tofunft *TofunftTransactorSession) Receive() (*types.Transaction, error) {
	return _Tofunft.Contract.Receive(&_Tofunft.TransactOpts)
}

// TofunftEvAuctionRefundIterator is returned from FilterEvAuctionRefund and is used to iterate over the raw logs and unpacked data for EvAuctionRefund events raised by the Tofunft contract.
type TofunftEvAuctionRefundIterator struct {
	Event *TofunftEvAuctionRefund // Event containing the contract specifics and raw log

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
func (it *TofunftEvAuctionRefundIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvAuctionRefund)
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
		it.Event = new(TofunftEvAuctionRefund)
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
func (it *TofunftEvAuctionRefundIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvAuctionRefundIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvAuctionRefund represents a EvAuctionRefund event raised by the Tofunft contract.
type TofunftEvAuctionRefund struct {
	Id     *big.Int
	Bidder common.Address
	Refund *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterEvAuctionRefund is a free log retrieval operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Tofunft *TofunftFilterer) FilterEvAuctionRefund(opts *bind.FilterOpts, id []*big.Int) (*TofunftEvAuctionRefundIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvAuctionRefund", idRule)
	if err != nil {
		return nil, err
	}
	return &TofunftEvAuctionRefundIterator{contract: _Tofunft.contract, event: "EvAuctionRefund", logs: logs, sub: sub}, nil
}

// WatchEvAuctionRefund is a free log subscription operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Tofunft *TofunftFilterer) WatchEvAuctionRefund(opts *bind.WatchOpts, sink chan<- *TofunftEvAuctionRefund, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvAuctionRefund", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvAuctionRefund)
				if err := _Tofunft.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
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

// ParseEvAuctionRefund is a log parse operation binding the contract event 0xa48bcf3362c21033397c03b92fb367d1962ba13b5bde0dfe491f9d88abb59e3f.
//
// Solidity: event EvAuctionRefund(uint256 indexed id, address bidder, uint256 refund)
func (_Tofunft *TofunftFilterer) ParseEvAuctionRefund(log types.Log) (*TofunftEvAuctionRefund, error) {
	event := new(TofunftEvAuctionRefund)
	if err := _Tofunft.contract.UnpackLog(event, "EvAuctionRefund", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftEvCouponSpentIterator is returned from FilterEvCouponSpent and is used to iterate over the raw logs and unpacked data for EvCouponSpent events raised by the Tofunft contract.
type TofunftEvCouponSpentIterator struct {
	Event *TofunftEvCouponSpent // Event containing the contract specifics and raw log

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
func (it *TofunftEvCouponSpentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvCouponSpent)
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
		it.Event = new(TofunftEvCouponSpent)
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
func (it *TofunftEvCouponSpentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvCouponSpentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvCouponSpent represents a EvCouponSpent event raised by the Tofunft contract.
type TofunftEvCouponSpent struct {
	Id       *big.Int
	CouponId *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterEvCouponSpent is a free log retrieval operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Tofunft *TofunftFilterer) FilterEvCouponSpent(opts *bind.FilterOpts, id []*big.Int, couponId []*big.Int) (*TofunftEvCouponSpentIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var couponIdRule []interface{}
	for _, couponIdItem := range couponId {
		couponIdRule = append(couponIdRule, couponIdItem)
	}

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvCouponSpent", idRule, couponIdRule)
	if err != nil {
		return nil, err
	}
	return &TofunftEvCouponSpentIterator{contract: _Tofunft.contract, event: "EvCouponSpent", logs: logs, sub: sub}, nil
}

// WatchEvCouponSpent is a free log subscription operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Tofunft *TofunftFilterer) WatchEvCouponSpent(opts *bind.WatchOpts, sink chan<- *TofunftEvCouponSpent, id []*big.Int, couponId []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var couponIdRule []interface{}
	for _, couponIdItem := range couponId {
		couponIdRule = append(couponIdRule, couponIdItem)
	}

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvCouponSpent", idRule, couponIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvCouponSpent)
				if err := _Tofunft.contract.UnpackLog(event, "EvCouponSpent", log); err != nil {
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

// ParseEvCouponSpent is a log parse operation binding the contract event 0x6aa71aa6b7aa6036ace4e4ceefbab7d89c4afb7fcfa1a3680499d7b37d32c82f.
//
// Solidity: event EvCouponSpent(uint256 indexed id, uint256 indexed couponId)
func (_Tofunft *TofunftFilterer) ParseEvCouponSpent(log types.Log) (*TofunftEvCouponSpent, error) {
	event := new(TofunftEvCouponSpent)
	if err := _Tofunft.contract.UnpackLog(event, "EvCouponSpent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftEvInventoryUpdateIterator is returned from FilterEvInventoryUpdate and is used to iterate over the raw logs and unpacked data for EvInventoryUpdate events raised by the Tofunft contract.
type TofunftEvInventoryUpdateIterator struct {
	Event *TofunftEvInventoryUpdate // Event containing the contract specifics and raw log

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
func (it *TofunftEvInventoryUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvInventoryUpdate)
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
		it.Event = new(TofunftEvInventoryUpdate)
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
func (it *TofunftEvInventoryUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvInventoryUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvInventoryUpdate represents a EvInventoryUpdate event raised by the Tofunft contract.
type TofunftEvInventoryUpdate struct {
	Id        *big.Int
	Inventory MarketNGInventory
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvInventoryUpdate is a free log retrieval operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Tofunft *TofunftFilterer) FilterEvInventoryUpdate(opts *bind.FilterOpts, id []*big.Int) (*TofunftEvInventoryUpdateIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvInventoryUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return &TofunftEvInventoryUpdateIterator{contract: _Tofunft.contract, event: "EvInventoryUpdate", logs: logs, sub: sub}, nil
}

// WatchEvInventoryUpdate is a free log subscription operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Tofunft *TofunftFilterer) WatchEvInventoryUpdate(opts *bind.WatchOpts, sink chan<- *TofunftEvInventoryUpdate, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvInventoryUpdate", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvInventoryUpdate)
				if err := _Tofunft.contract.UnpackLog(event, "EvInventoryUpdate", log); err != nil {
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

// ParseEvInventoryUpdate is a log parse operation binding the contract event 0x5beea7b3b87c573953fec05007114d17712e5775d364acc106d8da9e74849033.
//
// Solidity: event EvInventoryUpdate(uint256 indexed id, (address,address,address,uint256,uint256,uint256,uint8,uint8) inventory)
func (_Tofunft *TofunftFilterer) ParseEvInventoryUpdate(log types.Log) (*TofunftEvInventoryUpdate, error) {
	event := new(TofunftEvInventoryUpdate)
	if err := _Tofunft.contract.UnpackLog(event, "EvInventoryUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftEvMarketSignerUpdateIterator is returned from FilterEvMarketSignerUpdate and is used to iterate over the raw logs and unpacked data for EvMarketSignerUpdate events raised by the Tofunft contract.
type TofunftEvMarketSignerUpdateIterator struct {
	Event *TofunftEvMarketSignerUpdate // Event containing the contract specifics and raw log

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
func (it *TofunftEvMarketSignerUpdateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvMarketSignerUpdate)
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
		it.Event = new(TofunftEvMarketSignerUpdate)
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
func (it *TofunftEvMarketSignerUpdateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvMarketSignerUpdateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvMarketSignerUpdate represents a EvMarketSignerUpdate event raised by the Tofunft contract.
type TofunftEvMarketSignerUpdate struct {
	Addr      common.Address
	IsRemoval bool
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvMarketSignerUpdate is a free log retrieval operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Tofunft *TofunftFilterer) FilterEvMarketSignerUpdate(opts *bind.FilterOpts) (*TofunftEvMarketSignerUpdateIterator, error) {

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvMarketSignerUpdate")
	if err != nil {
		return nil, err
	}
	return &TofunftEvMarketSignerUpdateIterator{contract: _Tofunft.contract, event: "EvMarketSignerUpdate", logs: logs, sub: sub}, nil
}

// WatchEvMarketSignerUpdate is a free log subscription operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Tofunft *TofunftFilterer) WatchEvMarketSignerUpdate(opts *bind.WatchOpts, sink chan<- *TofunftEvMarketSignerUpdate) (event.Subscription, error) {

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvMarketSignerUpdate")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvMarketSignerUpdate)
				if err := _Tofunft.contract.UnpackLog(event, "EvMarketSignerUpdate", log); err != nil {
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

// ParseEvMarketSignerUpdate is a log parse operation binding the contract event 0x90d56af4745c314d9b45054b55dc973378c558c1ad1554bccc70d39aa63a2cc5.
//
// Solidity: event EvMarketSignerUpdate(address addr, bool isRemoval)
func (_Tofunft *TofunftFilterer) ParseEvMarketSignerUpdate(log types.Log) (*TofunftEvMarketSignerUpdate, error) {
	event := new(TofunftEvMarketSignerUpdate)
	if err := _Tofunft.contract.UnpackLog(event, "EvMarketSignerUpdate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftEvSettingsUpdatedIterator is returned from FilterEvSettingsUpdated and is used to iterate over the raw logs and unpacked data for EvSettingsUpdated events raised by the Tofunft contract.
type TofunftEvSettingsUpdatedIterator struct {
	Event *TofunftEvSettingsUpdated // Event containing the contract specifics and raw log

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
func (it *TofunftEvSettingsUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvSettingsUpdated)
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
		it.Event = new(TofunftEvSettingsUpdated)
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
func (it *TofunftEvSettingsUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvSettingsUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvSettingsUpdated represents a EvSettingsUpdated event raised by the Tofunft contract.
type TofunftEvSettingsUpdated struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEvSettingsUpdated is a free log retrieval operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Tofunft *TofunftFilterer) FilterEvSettingsUpdated(opts *bind.FilterOpts) (*TofunftEvSettingsUpdatedIterator, error) {

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return &TofunftEvSettingsUpdatedIterator{contract: _Tofunft.contract, event: "EvSettingsUpdated", logs: logs, sub: sub}, nil
}

// WatchEvSettingsUpdated is a free log subscription operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Tofunft *TofunftFilterer) WatchEvSettingsUpdated(opts *bind.WatchOpts, sink chan<- *TofunftEvSettingsUpdated) (event.Subscription, error) {

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvSettingsUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvSettingsUpdated)
				if err := _Tofunft.contract.UnpackLog(event, "EvSettingsUpdated", log); err != nil {
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

// ParseEvSettingsUpdated is a log parse operation binding the contract event 0x6c06ac894de6b71964f14d152b6674a4465a9b5d3f9cf9f216b8e7ea61467519.
//
// Solidity: event EvSettingsUpdated()
func (_Tofunft *TofunftFilterer) ParseEvSettingsUpdated(log types.Log) (*TofunftEvSettingsUpdated, error) {
	event := new(TofunftEvSettingsUpdated)
	if err := _Tofunft.contract.UnpackLog(event, "EvSettingsUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftEvSwappedIterator is returned from FilterEvSwapped and is used to iterate over the raw logs and unpacked data for EvSwapped events raised by the Tofunft contract.
type TofunftEvSwappedIterator struct {
	Event *TofunftEvSwapped // Event containing the contract specifics and raw log

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
func (it *TofunftEvSwappedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftEvSwapped)
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
		it.Event = new(TofunftEvSwapped)
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
func (it *TofunftEvSwappedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftEvSwappedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftEvSwapped represents a EvSwapped event raised by the Tofunft contract.
type TofunftEvSwapped struct {
	Req       MarketNGSwap
	Signature []byte
	Swapper   common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEvSwapped is a free log retrieval operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Tofunft *TofunftFilterer) FilterEvSwapped(opts *bind.FilterOpts) (*TofunftEvSwappedIterator, error) {

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "EvSwapped")
	if err != nil {
		return nil, err
	}
	return &TofunftEvSwappedIterator{contract: _Tofunft.contract, event: "EvSwapped", logs: logs, sub: sub}, nil
}

// WatchEvSwapped is a free log subscription operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Tofunft *TofunftFilterer) WatchEvSwapped(opts *bind.WatchOpts, sink chan<- *TofunftEvSwapped) (event.Subscription, error) {

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "EvSwapped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftEvSwapped)
				if err := _Tofunft.contract.UnpackLog(event, "EvSwapped", log); err != nil {
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

// ParseEvSwapped is a log parse operation binding the contract event 0x92060d15ec9a14885865b744d2efb1fff3cab53411058a530f51d480288a864c.
//
// Solidity: event EvSwapped((bytes32,address,uint256,(address,uint256)[],(address,uint256)[]) req, bytes signature, address swapper)
func (_Tofunft *TofunftFilterer) ParseEvSwapped(log types.Log) (*TofunftEvSwapped, error) {
	event := new(TofunftEvSwapped)
	if err := _Tofunft.contract.UnpackLog(event, "EvSwapped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Tofunft contract.
type TofunftOwnershipTransferredIterator struct {
	Event *TofunftOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *TofunftOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftOwnershipTransferred)
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
		it.Event = new(TofunftOwnershipTransferred)
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
func (it *TofunftOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftOwnershipTransferred represents a OwnershipTransferred event raised by the Tofunft contract.
type TofunftOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tofunft *TofunftFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TofunftOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TofunftOwnershipTransferredIterator{contract: _Tofunft.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Tofunft *TofunftFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TofunftOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftOwnershipTransferred)
				if err := _Tofunft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Tofunft *TofunftFilterer) ParseOwnershipTransferred(log types.Log) (*TofunftOwnershipTransferred, error) {
	event := new(TofunftOwnershipTransferred)
	if err := _Tofunft.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Tofunft contract.
type TofunftPausedIterator struct {
	Event *TofunftPaused // Event containing the contract specifics and raw log

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
func (it *TofunftPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftPaused)
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
		it.Event = new(TofunftPaused)
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
func (it *TofunftPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftPaused represents a Paused event raised by the Tofunft contract.
type TofunftPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tofunft *TofunftFilterer) FilterPaused(opts *bind.FilterOpts) (*TofunftPausedIterator, error) {

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TofunftPausedIterator{contract: _Tofunft.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tofunft *TofunftFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TofunftPaused) (event.Subscription, error) {

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftPaused)
				if err := _Tofunft.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Tofunft *TofunftFilterer) ParsePaused(log types.Log) (*TofunftPaused, error) {
	event := new(TofunftPaused)
	if err := _Tofunft.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TofunftUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Tofunft contract.
type TofunftUnpausedIterator struct {
	Event *TofunftUnpaused // Event containing the contract specifics and raw log

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
func (it *TofunftUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TofunftUnpaused)
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
		it.Event = new(TofunftUnpaused)
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
func (it *TofunftUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TofunftUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TofunftUnpaused represents a Unpaused event raised by the Tofunft contract.
type TofunftUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tofunft *TofunftFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TofunftUnpausedIterator, error) {

	logs, sub, err := _Tofunft.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TofunftUnpausedIterator{contract: _Tofunft.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tofunft *TofunftFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TofunftUnpaused) (event.Subscription, error) {

	logs, sub, err := _Tofunft.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TofunftUnpaused)
				if err := _Tofunft.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Tofunft *TofunftFilterer) ParseUnpaused(log types.Log) (*TofunftUnpaused, error) {
	event := new(TofunftUnpaused)
	if err := _Tofunft.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
