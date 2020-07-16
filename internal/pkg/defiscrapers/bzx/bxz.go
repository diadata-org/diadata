// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// BZxObjectsLoanOrder is an auto generated low-level Go binding around an user-defined struct.
type BZxObjectsLoanOrder struct {
	LoanTokenAddress            common.Address
	InterestTokenAddress        common.Address
	CollateralTokenAddress      common.Address
	OracleAddress               common.Address
	LoanTokenAmount             *big.Int
	InterestAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	LoanOrderHash               [32]byte
}

// BZxObjectsLoanPosition is an auto generated low-level Go binding around an user-defined struct.
type BZxObjectsLoanPosition struct {
	Trader                       common.Address
	CollateralTokenAddressFilled common.Address
	PositionTokenAddressFilled   common.Address
	LoanTokenAmountFilled        *big.Int
	LoanTokenAmountUsed          *big.Int
	CollateralTokenAmountFilled  *big.Int
	PositionTokenAmountFilled    *big.Int
	LoanStartUnixTimestampSec    *big.Int
	LoanEndUnixTimestampSec      *big.Int
	Active                       bool
	PositionId                   *big.Int
}

// LoanTokenStorageLoanData is an auto generated low-level Go binding around an user-defined struct.
type LoanTokenStorageLoanData struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}

// AdvancedTokenABI is the input ABI used to generate the binding from.
const AdvancedTokenABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenToFillAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawOnOpen\",\"type\":\"bool\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burntTokenReserveList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burntTokenReserveListIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burntTokenReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leverageList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loanOrderData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginPremiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loanOrderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spreadMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenizedRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAssetBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AdvancedTokenFuncSigs maps the 4-byte function signature to its string representation.
var AdvancedTokenFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"995363d3": "bZxContract()",
	"96c7871b": "bZxOracle()",
	"894ca308": "bZxVault()",
	"70a08231": "balanceOf(address)",
	"1f68f20a": "baseRate()",
	"7866c6c1": "burntTokenReserveList(uint256)",
	"fbd9574d": "burntTokenReserveListIndex(address)",
	"0c4925fd": "burntTokenReserved()",
	"7b7933b4": "checkpointSupply()",
	"313ce567": "decimals()",
	"1d0806ae": "initialPrice()",
	"9b3a54d1": "leverageList(uint256)",
	"2515aacd": "loanOrderData(bytes32)",
	"fe056342": "loanOrderHashes(uint256)",
	"797bf385": "loanTokenAddress()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"330691ac": "rateMultiplier()",
	"d84d2a47": "spreadMultiplier()",
	"95d89b41": "symbol()",
	"736ee3d3": "tokenizedRegistry()",
	"20f6d07c": "totalAssetBorrow()",
	"18160ddd": "totalSupply()",
	"f2fde38b": "transferOwnership(address)",
	"4780eac1": "wethContract()",
}

// AdvancedTokenBin is the compiled bytecode used for deploying new contracts.
var AdvancedTokenBin = "0x608060405260016000819055600a805460ff19169055670de0b6b3a7640000600b556801043561a882930000600c5580546001600160a01b031916331790556109f68061004d6000396000f3fe608060405234801561001057600080fd5b50600436106101a95760003560e01c80637866c6c1116100f9578063995363d311610097578063dd62ed3e11610071578063dd62ed3e14610309578063f2fde38b1461031c578063fbd9574d14610331578063fe05634214610352576101a9565b8063995363d3146102e65780639b3a54d1146102ee578063d84d2a4714610301576101a9565b8063894ca308116100d3578063894ca308146102c65780638da5cb5b146102ce57806395d89b41146102d657806396c7871b146102de576101a9565b80637866c6c114610295578063797bf385146102b65780637b7933b4146102be576101a9565b806320f6d07c11610166578063330691ac11610140578063330691ac1461025d5780634780eac11461026557806370a082311461027a578063736ee3d31461028d576101a9565b806320f6d07c146102195780632515aacd14610221578063313ce56714610248576101a9565b806306fdde03146101ae578063095ea7b3146101cc5780630c4925fd146101ec57806318160ddd146102015780631d0806ae146102095780631f68f20a14610211575b600080fd5b6101b6610365565b6040516101c391906108f0565b60405180910390f35b6101df6101da36600461077d565b6103f0565b6040516101c3919061085d565b6101f461045b565b6040516101c3919061086b565b6101f4610461565b6101f4610467565b6101f461046d565b6101f4610473565b61023461022f3660046107ad565b610479565b6040516101c3989796959493929190610879565b6102506104c5565b6040516101c3919061091c565b6101f46104ce565b61026d6104d4565b6040516101c3919061082d565b6101f461028836600461071d565b6104e3565b61026d6104fe565b6102a86102a33660046107ad565b610512565b6040516101c392919061083b565b61026d610547565b6101f4610556565b61026d61055c565b61026d61056b565b6101b661057a565b61026d6105d5565b61026d6105e4565b6101f46102fc3660046107ad565b6105f8565b6101f4610616565b6101f4610317366004610743565b61061c565b61032f61032a36600461071d565b610647565b005b61034461033f36600461071d565b61066a565b6040516101c3929190610901565b6101f46103603660046107ad565b610686565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103e85780601f106103bd576101008083540402835291602001916103e8565b820191906000526020600020905b8154815290600101906020018083116103cb57829003601f168201915b505050505081565b336000818152601a602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9259061044990869061086b565b60405180910390a35060015b92915050565b60135481565b601b5490565b60185481565b600b5481565b60155481565b600f60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007909701549596949593949293919290916001600160a01b031688565b60045460ff1681565b600c5481565b6007546001600160a01b031681565b6001600160a01b031660009081526019602052604090205490565b600a5461010090046001600160a01b031681565b6011818154811061051f57fe5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b6008546001600160a01b031681565b60165481565b6005546001600160a01b031681565b6001546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103e85780601f106103bd576101008083540402835291602001916103e8565b6006546001600160a01b031681565b60045461010090046001600160a01b031681565b6010818154811061060557fe5b600091825260209091200154905081565b600d5481565b6001600160a01b039182166000908152601a6020908152604080832093909416825291909152205490565b6001546001600160a01b0316331461065e57600080fd5b61066781610698565b50565b6012602052600090815260409020805460019091015460ff1682565b600e6020526000908152604090205481565b6001600160a01b0381166106ab57600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b803561045581610996565b8035610455816109aa565b60006020828403121561072f57600080fd5b600061073b8484610707565b949350505050565b6000806040838503121561075657600080fd5b60006107628585610707565b925050602061077385828601610707565b9150509250929050565b6000806040838503121561079057600080fd5b600061079c8585610707565b925050602061077385828601610712565b6000602082840312156107bf57600080fd5b600061073b8484610712565b6107d481610937565b82525050565b6107d481610942565b6107d481610947565b60006107f78261092a565b610801818561092e565b935061081181856020860161095c565b61081a8161098c565b9093019392505050565b6107d481610956565b6020810161045582846107cb565b6040810161084982856107cb565b61085660208301846107e3565b9392505050565b6020810161045582846107da565b6020810161045582846107e3565b6101008101610888828b6107e3565b610895602083018a6107e3565b6108a260408301896107e3565b6108af60608301886107e3565b6108bc60808301876107e3565b6108c960a08301866107e3565b6108d660c08301856107e3565b6108e360e08301846107cb565b9998505050505050505050565b6020808252810161085681846107ec565b6040810161090f82856107e3565b61085660208301846107da565b602081016104558284610824565b5190565b90815260200190565b60006104558261094a565b151590565b90565b6001600160a01b031690565b60ff1690565b60005b8381101561097757818101518382015260200161095f565b83811115610986576000848401525b50505050565b601f01601f191690565b61099f81610937565b811461066757600080fd5b61099f8161094756fea365627a7a72315820d97b1d79265865ca44f9727d590c9d07c59472320340c0ecb2a98a2ed915de516c6578706572696d656e74616cf564736f6c63430005110040"

// DeployAdvancedToken deploys a new Ethereum contract, binding an instance of AdvancedToken to it.
func DeployAdvancedToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdvancedToken, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvancedTokenABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdvancedTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdvancedToken{AdvancedTokenCaller: AdvancedTokenCaller{contract: contract}, AdvancedTokenTransactor: AdvancedTokenTransactor{contract: contract}, AdvancedTokenFilterer: AdvancedTokenFilterer{contract: contract}}, nil
}

// AdvancedToken is an auto generated Go binding around an Ethereum contract.
type AdvancedToken struct {
	AdvancedTokenCaller     // Read-only binding to the contract
	AdvancedTokenTransactor // Write-only binding to the contract
	AdvancedTokenFilterer   // Log filterer for contract events
}

// AdvancedTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdvancedTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdvancedTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdvancedTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdvancedTokenSession struct {
	Contract     *AdvancedToken    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AdvancedTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdvancedTokenCallerSession struct {
	Contract *AdvancedTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// AdvancedTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdvancedTokenTransactorSession struct {
	Contract     *AdvancedTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// AdvancedTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdvancedTokenRaw struct {
	Contract *AdvancedToken // Generic contract binding to access the raw methods on
}

// AdvancedTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdvancedTokenCallerRaw struct {
	Contract *AdvancedTokenCaller // Generic read-only contract binding to access the raw methods on
}

// AdvancedTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdvancedTokenTransactorRaw struct {
	Contract *AdvancedTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdvancedToken creates a new instance of AdvancedToken, bound to a specific deployed contract.
func NewAdvancedToken(address common.Address, backend bind.ContractBackend) (*AdvancedToken, error) {
	contract, err := bindAdvancedToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedToken{AdvancedTokenCaller: AdvancedTokenCaller{contract: contract}, AdvancedTokenTransactor: AdvancedTokenTransactor{contract: contract}, AdvancedTokenFilterer: AdvancedTokenFilterer{contract: contract}}, nil
}

// NewAdvancedTokenCaller creates a new read-only instance of AdvancedToken, bound to a specific deployed contract.
func NewAdvancedTokenCaller(address common.Address, caller bind.ContractCaller) (*AdvancedTokenCaller, error) {
	contract, err := bindAdvancedToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenCaller{contract: contract}, nil
}

// NewAdvancedTokenTransactor creates a new write-only instance of AdvancedToken, bound to a specific deployed contract.
func NewAdvancedTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedTokenTransactor, error) {
	contract, err := bindAdvancedToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenTransactor{contract: contract}, nil
}

// NewAdvancedTokenFilterer creates a new log filterer instance of AdvancedToken, bound to a specific deployed contract.
func NewAdvancedTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedTokenFilterer, error) {
	contract, err := bindAdvancedToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenFilterer{contract: contract}, nil
}

// bindAdvancedToken binds a generic wrapper to an already deployed contract.
func bindAdvancedToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvancedTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedToken *AdvancedTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdvancedToken.Contract.AdvancedTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedToken *AdvancedTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedToken.Contract.AdvancedTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedToken *AdvancedTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedToken.Contract.AdvancedTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedToken *AdvancedTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdvancedToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedToken *AdvancedTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedToken *AdvancedTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AdvancedToken.Contract.Allowance(&_AdvancedToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AdvancedToken.Contract.Allowance(&_AdvancedToken.CallOpts, _owner, _spender)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) BZxContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "bZxContract")
	return *ret0, err
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) BZxContract() (common.Address, error) {
	return _AdvancedToken.Contract.BZxContract(&_AdvancedToken.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) BZxContract() (common.Address, error) {
	return _AdvancedToken.Contract.BZxContract(&_AdvancedToken.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) BZxOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "bZxOracle")
	return *ret0, err
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) BZxOracle() (common.Address, error) {
	return _AdvancedToken.Contract.BZxOracle(&_AdvancedToken.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) BZxOracle() (common.Address, error) {
	return _AdvancedToken.Contract.BZxOracle(&_AdvancedToken.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) BZxVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "bZxVault")
	return *ret0, err
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) BZxVault() (common.Address, error) {
	return _AdvancedToken.Contract.BZxVault(&_AdvancedToken.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) BZxVault() (common.Address, error) {
	return _AdvancedToken.Contract.BZxVault(&_AdvancedToken.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AdvancedToken.Contract.BalanceOf(&_AdvancedToken.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AdvancedToken.Contract.BalanceOf(&_AdvancedToken.CallOpts, _owner)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "baseRate")
	return *ret0, err
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) BaseRate() (*big.Int, error) {
	return _AdvancedToken.Contract.BaseRate(&_AdvancedToken.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) BaseRate() (*big.Int, error) {
	return _AdvancedToken.Contract.BaseRate(&_AdvancedToken.CallOpts)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedToken *AdvancedTokenCaller) BurntTokenReserveList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Lender common.Address
		Amount *big.Int
	})
	out := ret
	err := _AdvancedToken.contract.Call(opts, out, "burntTokenReserveList", arg0)
	return *ret, err
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedToken *AdvancedTokenSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _AdvancedToken.Contract.BurntTokenReserveList(&_AdvancedToken.CallOpts, arg0)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedToken *AdvancedTokenCallerSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _AdvancedToken.Contract.BurntTokenReserveList(&_AdvancedToken.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedToken *AdvancedTokenCaller) BurntTokenReserveListIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	ret := new(struct {
		Index *big.Int
		IsSet bool
	})
	out := ret
	err := _AdvancedToken.contract.Call(opts, out, "burntTokenReserveListIndex", arg0)
	return *ret, err
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedToken *AdvancedTokenSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _AdvancedToken.Contract.BurntTokenReserveListIndex(&_AdvancedToken.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedToken *AdvancedTokenCallerSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _AdvancedToken.Contract.BurntTokenReserveListIndex(&_AdvancedToken.CallOpts, arg0)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) BurntTokenReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "burntTokenReserved")
	return *ret0, err
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) BurntTokenReserved() (*big.Int, error) {
	return _AdvancedToken.Contract.BurntTokenReserved(&_AdvancedToken.CallOpts)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) BurntTokenReserved() (*big.Int, error) {
	return _AdvancedToken.Contract.BurntTokenReserved(&_AdvancedToken.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) CheckpointSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "checkpointSupply")
	return *ret0, err
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) CheckpointSupply() (*big.Int, error) {
	return _AdvancedToken.Contract.CheckpointSupply(&_AdvancedToken.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) CheckpointSupply() (*big.Int, error) {
	return _AdvancedToken.Contract.CheckpointSupply(&_AdvancedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedToken *AdvancedTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedToken *AdvancedTokenSession) Decimals() (uint8, error) {
	return _AdvancedToken.Contract.Decimals(&_AdvancedToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedToken *AdvancedTokenCallerSession) Decimals() (uint8, error) {
	return _AdvancedToken.Contract.Decimals(&_AdvancedToken.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) InitialPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "initialPrice")
	return *ret0, err
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) InitialPrice() (*big.Int, error) {
	return _AdvancedToken.Contract.InitialPrice(&_AdvancedToken.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) InitialPrice() (*big.Int, error) {
	return _AdvancedToken.Contract.InitialPrice(&_AdvancedToken.CallOpts)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) LeverageList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "leverageList", arg0)
	return *ret0, err
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _AdvancedToken.Contract.LeverageList(&_AdvancedToken.CallOpts, arg0)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _AdvancedToken.Contract.LeverageList(&_AdvancedToken.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedToken *AdvancedTokenCaller) LoanOrderData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	ret := new(struct {
		LoanOrderHash               [32]byte
		LeverageAmount              *big.Int
		InitialMarginAmount         *big.Int
		MaintenanceMarginAmount     *big.Int
		MaxDurationUnixTimestampSec *big.Int
		Index                       *big.Int
		MarginPremiumAmount         *big.Int
		CollateralTokenAddress      common.Address
	})
	out := ret
	err := _AdvancedToken.contract.Call(opts, out, "loanOrderData", arg0)
	return *ret, err
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedToken *AdvancedTokenSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _AdvancedToken.Contract.LoanOrderData(&_AdvancedToken.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedToken *AdvancedTokenCallerSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _AdvancedToken.Contract.LoanOrderData(&_AdvancedToken.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedToken *AdvancedTokenCaller) LoanOrderHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "loanOrderHashes", arg0)
	return *ret0, err
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedToken *AdvancedTokenSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _AdvancedToken.Contract.LoanOrderHashes(&_AdvancedToken.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedToken *AdvancedTokenCallerSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _AdvancedToken.Contract.LoanOrderHashes(&_AdvancedToken.CallOpts, arg0)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) LoanTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "loanTokenAddress")
	return *ret0, err
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) LoanTokenAddress() (common.Address, error) {
	return _AdvancedToken.Contract.LoanTokenAddress(&_AdvancedToken.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) LoanTokenAddress() (common.Address, error) {
	return _AdvancedToken.Contract.LoanTokenAddress(&_AdvancedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedToken *AdvancedTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedToken *AdvancedTokenSession) Name() (string, error) {
	return _AdvancedToken.Contract.Name(&_AdvancedToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedToken *AdvancedTokenCallerSession) Name() (string, error) {
	return _AdvancedToken.Contract.Name(&_AdvancedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) Owner() (common.Address, error) {
	return _AdvancedToken.Contract.Owner(&_AdvancedToken.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) Owner() (common.Address, error) {
	return _AdvancedToken.Contract.Owner(&_AdvancedToken.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) RateMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "rateMultiplier")
	return *ret0, err
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) RateMultiplier() (*big.Int, error) {
	return _AdvancedToken.Contract.RateMultiplier(&_AdvancedToken.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) RateMultiplier() (*big.Int, error) {
	return _AdvancedToken.Contract.RateMultiplier(&_AdvancedToken.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) SpreadMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "spreadMultiplier")
	return *ret0, err
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) SpreadMultiplier() (*big.Int, error) {
	return _AdvancedToken.Contract.SpreadMultiplier(&_AdvancedToken.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) SpreadMultiplier() (*big.Int, error) {
	return _AdvancedToken.Contract.SpreadMultiplier(&_AdvancedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedToken *AdvancedTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedToken *AdvancedTokenSession) Symbol() (string, error) {
	return _AdvancedToken.Contract.Symbol(&_AdvancedToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedToken *AdvancedTokenCallerSession) Symbol() (string, error) {
	return _AdvancedToken.Contract.Symbol(&_AdvancedToken.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) TokenizedRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "tokenizedRegistry")
	return *ret0, err
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) TokenizedRegistry() (common.Address, error) {
	return _AdvancedToken.Contract.TokenizedRegistry(&_AdvancedToken.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) TokenizedRegistry() (common.Address, error) {
	return _AdvancedToken.Contract.TokenizedRegistry(&_AdvancedToken.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) TotalAssetBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "totalAssetBorrow")
	return *ret0, err
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) TotalAssetBorrow() (*big.Int, error) {
	return _AdvancedToken.Contract.TotalAssetBorrow(&_AdvancedToken.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) TotalAssetBorrow() (*big.Int, error) {
	return _AdvancedToken.Contract.TotalAssetBorrow(&_AdvancedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenSession) TotalSupply() (*big.Int, error) {
	return _AdvancedToken.Contract.TotalSupply(&_AdvancedToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedToken *AdvancedTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _AdvancedToken.Contract.TotalSupply(&_AdvancedToken.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedToken *AdvancedTokenCaller) WethContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedToken.contract.Call(opts, out, "wethContract")
	return *ret0, err
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedToken *AdvancedTokenSession) WethContract() (common.Address, error) {
	return _AdvancedToken.Contract.WethContract(&_AdvancedToken.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedToken *AdvancedTokenCallerSession) WethContract() (common.Address, error) {
	return _AdvancedToken.Contract.WethContract(&_AdvancedToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_AdvancedToken *AdvancedTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_AdvancedToken *AdvancedTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedToken.Contract.Approve(&_AdvancedToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_AdvancedToken *AdvancedTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _AdvancedToken.Contract.Approve(&_AdvancedToken.TransactOpts, _spender, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedToken *AdvancedTokenTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedToken.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedToken *AdvancedTokenSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedToken.Contract.TransferOwnership(&_AdvancedToken.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedToken *AdvancedTokenTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedToken.Contract.TransferOwnership(&_AdvancedToken.TransactOpts, _newOwner)
}

// AdvancedTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AdvancedToken contract.
type AdvancedTokenApprovalIterator struct {
	Event *AdvancedTokenApproval // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenApproval)
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
		it.Event = new(AdvancedTokenApproval)
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
func (it *AdvancedTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenApproval represents a Approval event raised by the AdvancedToken contract.
type AdvancedTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AdvancedTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenApprovalIterator{contract: _AdvancedToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AdvancedTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenApproval)
				if err := _AdvancedToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) ParseApproval(log types.Log) (*AdvancedTokenApproval, error) {
	event := new(AdvancedTokenApproval)
	if err := _AdvancedToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AdvancedToken contract.
type AdvancedTokenBorrowIterator struct {
	Event *AdvancedTokenBorrow // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenBorrow)
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
		it.Event = new(AdvancedTokenBorrow)
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
func (it *AdvancedTokenBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenBorrow represents a Borrow event raised by the AdvancedToken contract.
type AdvancedTokenBorrow struct {
	Borrower                common.Address
	BorrowAmount            *big.Int
	InterestRate            *big.Int
	CollateralTokenAddress  common.Address
	TradeTokenToFillAddress common.Address
	WithdrawOnOpen          bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedToken *AdvancedTokenFilterer) FilterBorrow(opts *bind.FilterOpts, borrower []common.Address) (*AdvancedTokenBorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenBorrowIterator{contract: _AdvancedToken.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedToken *AdvancedTokenFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AdvancedTokenBorrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenBorrow)
				if err := _AdvancedToken.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedToken *AdvancedTokenFilterer) ParseBorrow(log types.Log) (*AdvancedTokenBorrow, error) {
	event := new(AdvancedTokenBorrow)
	if err := _AdvancedToken.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AdvancedToken contract.
type AdvancedTokenBurnIterator struct {
	Event *AdvancedTokenBurn // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenBurn)
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
		it.Event = new(AdvancedTokenBurn)
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
func (it *AdvancedTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenBurn represents a Burn event raised by the AdvancedToken contract.
type AdvancedTokenBurn struct {
	Burner      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*AdvancedTokenBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenBurnIterator{contract: _AdvancedToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AdvancedTokenBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenBurn)
				if err := _AdvancedToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) ParseBurn(log types.Log) (*AdvancedTokenBurn, error) {
	event := new(AdvancedTokenBurn)
	if err := _AdvancedToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdvancedToken contract.
type AdvancedTokenClaimIterator struct {
	Event *AdvancedTokenClaim // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenClaim)
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
		it.Event = new(AdvancedTokenClaim)
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
func (it *AdvancedTokenClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenClaim represents a Claim event raised by the AdvancedToken contract.
type AdvancedTokenClaim struct {
	Claimant             common.Address
	TokenAmount          *big.Int
	AssetAmount          *big.Int
	RemainingTokenAmount *big.Int
	Price                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) FilterClaim(opts *bind.FilterOpts, claimant []common.Address) (*AdvancedTokenClaimIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenClaimIterator{contract: _AdvancedToken.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdvancedTokenClaim, claimant []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenClaim)
				if err := _AdvancedToken.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) ParseClaim(log types.Log) (*AdvancedTokenClaim, error) {
	event := new(AdvancedTokenClaim)
	if err := _AdvancedToken.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AdvancedToken contract.
type AdvancedTokenMintIterator struct {
	Event *AdvancedTokenMint // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenMint)
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
		it.Event = new(AdvancedTokenMint)
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
func (it *AdvancedTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenMint represents a Mint event raised by the AdvancedToken contract.
type AdvancedTokenMint struct {
	Minter      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) FilterMint(opts *bind.FilterOpts, minter []common.Address) (*AdvancedTokenMintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenMintIterator{contract: _AdvancedToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AdvancedTokenMint, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenMint)
				if err := _AdvancedToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedToken *AdvancedTokenFilterer) ParseMint(log types.Log) (*AdvancedTokenMint, error) {
	event := new(AdvancedTokenMint)
	if err := _AdvancedToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AdvancedToken contract.
type AdvancedTokenOwnershipTransferredIterator struct {
	Event *AdvancedTokenOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenOwnershipTransferred)
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
		it.Event = new(AdvancedTokenOwnershipTransferred)
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
func (it *AdvancedTokenOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenOwnershipTransferred represents a OwnershipTransferred event raised by the AdvancedToken contract.
type AdvancedTokenOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvancedToken *AdvancedTokenFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AdvancedTokenOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenOwnershipTransferredIterator{contract: _AdvancedToken.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvancedToken *AdvancedTokenFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedTokenOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenOwnershipTransferred)
				if err := _AdvancedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AdvancedToken *AdvancedTokenFilterer) ParseOwnershipTransferred(log types.Log) (*AdvancedTokenOwnershipTransferred, error) {
	event := new(AdvancedTokenOwnershipTransferred)
	if err := _AdvancedToken.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AdvancedToken contract.
type AdvancedTokenRepayIterator struct {
	Event *AdvancedTokenRepay // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenRepay)
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
		it.Event = new(AdvancedTokenRepay)
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
func (it *AdvancedTokenRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenRepay represents a Repay event raised by the AdvancedToken contract.
type AdvancedTokenRepay struct {
	LoanOrderHash [32]byte
	Borrower      common.Address
	Closer        common.Address
	Amount        *big.Int
	IsLiquidation bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedToken *AdvancedTokenFilterer) FilterRepay(opts *bind.FilterOpts, loanOrderHash [][32]byte, borrower []common.Address) (*AdvancedTokenRepayIterator, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenRepayIterator{contract: _AdvancedToken.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedToken *AdvancedTokenFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AdvancedTokenRepay, loanOrderHash [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenRepay)
				if err := _AdvancedToken.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedToken *AdvancedTokenFilterer) ParseRepay(log types.Log) (*AdvancedTokenRepay, error) {
	event := new(AdvancedTokenRepay)
	if err := _AdvancedToken.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AdvancedToken contract.
type AdvancedTokenTransferIterator struct {
	Event *AdvancedTokenTransfer // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenTransfer)
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
		it.Event = new(AdvancedTokenTransfer)
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
func (it *AdvancedTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenTransfer represents a Transfer event raised by the AdvancedToken contract.
type AdvancedTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenTransferIterator{contract: _AdvancedToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AdvancedTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenTransfer)
				if err := _AdvancedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedToken *AdvancedTokenFilterer) ParseTransfer(log types.Log) (*AdvancedTokenTransfer, error) {
	event := new(AdvancedTokenTransfer)
	if err := _AdvancedToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageABI is the input ABI used to generate the binding from.
const AdvancedTokenStorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenToFillAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawOnOpen\",\"type\":\"bool\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burntTokenReserveList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burntTokenReserveListIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burntTokenReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leverageList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loanOrderData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginPremiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loanOrderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spreadMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenizedRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAssetBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// AdvancedTokenStorageFuncSigs maps the 4-byte function signature to its string representation.
var AdvancedTokenStorageFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"995363d3": "bZxContract()",
	"96c7871b": "bZxOracle()",
	"894ca308": "bZxVault()",
	"70a08231": "balanceOf(address)",
	"1f68f20a": "baseRate()",
	"7866c6c1": "burntTokenReserveList(uint256)",
	"fbd9574d": "burntTokenReserveListIndex(address)",
	"0c4925fd": "burntTokenReserved()",
	"7b7933b4": "checkpointSupply()",
	"313ce567": "decimals()",
	"1d0806ae": "initialPrice()",
	"9b3a54d1": "leverageList(uint256)",
	"2515aacd": "loanOrderData(bytes32)",
	"fe056342": "loanOrderHashes(uint256)",
	"797bf385": "loanTokenAddress()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"330691ac": "rateMultiplier()",
	"d84d2a47": "spreadMultiplier()",
	"95d89b41": "symbol()",
	"736ee3d3": "tokenizedRegistry()",
	"20f6d07c": "totalAssetBorrow()",
	"18160ddd": "totalSupply()",
	"f2fde38b": "transferOwnership(address)",
	"4780eac1": "wethContract()",
}

// AdvancedTokenStorageBin is the compiled bytecode used for deploying new contracts.
var AdvancedTokenStorageBin = "0x608060405260016000819055600a805460ff19169055670de0b6b3a7640000600b556801043561a882930000600c5580546001600160a01b031916331790556109148061004d6000396000f3fe608060405234801561001057600080fd5b506004361061018e5760003560e01c8063797bf385116100de578063995363d311610097578063dd62ed3e11610071578063dd62ed3e146102ce578063f2fde38b146102e1578063fbd9574d146102f6578063fe056342146103175761018e565b8063995363d3146102ab5780639b3a54d1146102b3578063d84d2a47146102c65761018e565b8063797bf3851461027b5780637b7933b414610283578063894ca3081461028b5780638da5cb5b1461029357806395d89b411461029b57806396c7871b146102a35761018e565b80632515aacd1161014b5780634780eac1116101255780634780eac11461022a57806370a082311461023f578063736ee3d3146102525780637866c6c11461025a5761018e565b80632515aacd146101e6578063313ce5671461020d578063330691ac146102225761018e565b806306fdde03146101935780630c4925fd146101b157806318160ddd146101c65780631d0806ae146101ce5780631f68f20a146101d657806320f6d07c146101de575b600080fd5b61019b61032a565b6040516101a8919061080e565b60405180910390f35b6101b96103b5565b6040516101a89190610789565b6101b96103bb565b6101b96103c1565b6101b96103c7565b6101b96103cd565b6101f96101f43660046106d9565b6103d3565b6040516101a8989796959493929190610797565b61021561041f565b6040516101a8919061083a565b6101b9610428565b61023261042e565b6040516101a89190610759565b6101b961024d366004610679565b61043d565b610232610458565b61026d6102683660046106d9565b61046c565b6040516101a8929190610767565b6102326104a1565b6101b96104b0565b6102326104b6565b6102326104c5565b61019b6104d4565b61023261052f565b61023261053e565b6101b96102c13660046106d9565b610552565b6101b9610570565b6101b96102dc36600461069f565b610576565b6102f46102ef366004610679565b6105a3565b005b610309610304366004610679565b6105c6565b6040516101a892919061081f565b6101b96103253660046106d9565b6105e2565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156103ad5780601f10610382576101008083540402835291602001916103ad565b820191906000526020600020905b81548152906001019060200180831161039057829003601f168201915b505050505081565b60135481565b601b5490565b60185481565b600b5481565b60155481565b600f60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007909701549596949593949293919290916001600160a01b031688565b60045460ff1681565b600c5481565b6007546001600160a01b031681565b6001600160a01b031660009081526019602052604090205490565b600a5461010090046001600160a01b031681565b6011818154811061047957fe5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b6008546001600160a01b031681565b60165481565b6005546001600160a01b031681565b6001546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156103ad5780601f10610382576101008083540402835291602001916103ad565b6006546001600160a01b031681565b60045461010090046001600160a01b031681565b6010818154811061055f57fe5b600091825260209091200154905081565b600d5481565b6001600160a01b038083166000908152601a60209081526040808320938516835292905220545b92915050565b6001546001600160a01b031633146105ba57600080fd5b6105c3816105f4565b50565b6012602052600090815260409020805460019091015460ff1682565b600e6020526000908152604090205481565b6001600160a01b03811661060757600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b803561059d816108b4565b803561059d816108c8565b60006020828403121561068b57600080fd5b60006106978484610663565b949350505050565b600080604083850312156106b257600080fd5b60006106be8585610663565b92505060206106cf85828601610663565b9150509250929050565b6000602082840312156106eb57600080fd5b6000610697848461066e565b61070081610855565b82525050565b61070081610860565b61070081610865565b600061072382610848565b61072d818561084c565b935061073d81856020860161087a565b610746816108aa565b9093019392505050565b61070081610874565b6020810161059d82846106f7565b6040810161077582856106f7565b610782602083018461070f565b9392505050565b6020810161059d828461070f565b61010081016107a6828b61070f565b6107b3602083018a61070f565b6107c0604083018961070f565b6107cd606083018861070f565b6107da608083018761070f565b6107e760a083018661070f565b6107f460c083018561070f565b61080160e08301846106f7565b9998505050505050505050565b602080825281016107828184610718565b6040810161082d828561070f565b6107826020830184610706565b6020810161059d8284610750565b5190565b90815260200190565b600061059d82610868565b151590565b90565b6001600160a01b031690565b60ff1690565b60005b8381101561089557818101518382015260200161087d565b838111156108a4576000848401525b50505050565b601f01601f191690565b6108bd81610855565b81146105c357600080fd5b6108bd8161086556fea365627a7a72315820a7a06faa26e3c41dd95f4e0490a113f141f25320e97c12b380c378e057f1e84a6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployAdvancedTokenStorage deploys a new Ethereum contract, binding an instance of AdvancedTokenStorage to it.
func DeployAdvancedTokenStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AdvancedTokenStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvancedTokenStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(AdvancedTokenStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AdvancedTokenStorage{AdvancedTokenStorageCaller: AdvancedTokenStorageCaller{contract: contract}, AdvancedTokenStorageTransactor: AdvancedTokenStorageTransactor{contract: contract}, AdvancedTokenStorageFilterer: AdvancedTokenStorageFilterer{contract: contract}}, nil
}

// AdvancedTokenStorage is an auto generated Go binding around an Ethereum contract.
type AdvancedTokenStorage struct {
	AdvancedTokenStorageCaller     // Read-only binding to the contract
	AdvancedTokenStorageTransactor // Write-only binding to the contract
	AdvancedTokenStorageFilterer   // Log filterer for contract events
}

// AdvancedTokenStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type AdvancedTokenStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AdvancedTokenStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AdvancedTokenStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AdvancedTokenStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AdvancedTokenStorageSession struct {
	Contract     *AdvancedTokenStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// AdvancedTokenStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AdvancedTokenStorageCallerSession struct {
	Contract *AdvancedTokenStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// AdvancedTokenStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AdvancedTokenStorageTransactorSession struct {
	Contract     *AdvancedTokenStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// AdvancedTokenStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type AdvancedTokenStorageRaw struct {
	Contract *AdvancedTokenStorage // Generic contract binding to access the raw methods on
}

// AdvancedTokenStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AdvancedTokenStorageCallerRaw struct {
	Contract *AdvancedTokenStorageCaller // Generic read-only contract binding to access the raw methods on
}

// AdvancedTokenStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AdvancedTokenStorageTransactorRaw struct {
	Contract *AdvancedTokenStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAdvancedTokenStorage creates a new instance of AdvancedTokenStorage, bound to a specific deployed contract.
func NewAdvancedTokenStorage(address common.Address, backend bind.ContractBackend) (*AdvancedTokenStorage, error) {
	contract, err := bindAdvancedTokenStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorage{AdvancedTokenStorageCaller: AdvancedTokenStorageCaller{contract: contract}, AdvancedTokenStorageTransactor: AdvancedTokenStorageTransactor{contract: contract}, AdvancedTokenStorageFilterer: AdvancedTokenStorageFilterer{contract: contract}}, nil
}

// NewAdvancedTokenStorageCaller creates a new read-only instance of AdvancedTokenStorage, bound to a specific deployed contract.
func NewAdvancedTokenStorageCaller(address common.Address, caller bind.ContractCaller) (*AdvancedTokenStorageCaller, error) {
	contract, err := bindAdvancedTokenStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageCaller{contract: contract}, nil
}

// NewAdvancedTokenStorageTransactor creates a new write-only instance of AdvancedTokenStorage, bound to a specific deployed contract.
func NewAdvancedTokenStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*AdvancedTokenStorageTransactor, error) {
	contract, err := bindAdvancedTokenStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageTransactor{contract: contract}, nil
}

// NewAdvancedTokenStorageFilterer creates a new log filterer instance of AdvancedTokenStorage, bound to a specific deployed contract.
func NewAdvancedTokenStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*AdvancedTokenStorageFilterer, error) {
	contract, err := bindAdvancedTokenStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageFilterer{contract: contract}, nil
}

// bindAdvancedTokenStorage binds a generic wrapper to an already deployed contract.
func bindAdvancedTokenStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AdvancedTokenStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedTokenStorage *AdvancedTokenStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdvancedTokenStorage.Contract.AdvancedTokenStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedTokenStorage *AdvancedTokenStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.AdvancedTokenStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedTokenStorage *AdvancedTokenStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.AdvancedTokenStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _AdvancedTokenStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AdvancedTokenStorage *AdvancedTokenStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AdvancedTokenStorage *AdvancedTokenStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.Allowance(&_AdvancedTokenStorage.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.Allowance(&_AdvancedTokenStorage.CallOpts, _owner, _spender)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BZxContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "bZxContract")
	return *ret0, err
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BZxContract() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxContract(&_AdvancedTokenStorage.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BZxContract() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxContract(&_AdvancedTokenStorage.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BZxOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "bZxOracle")
	return *ret0, err
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BZxOracle() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxOracle(&_AdvancedTokenStorage.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BZxOracle() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxOracle(&_AdvancedTokenStorage.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BZxVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "bZxVault")
	return *ret0, err
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BZxVault() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxVault(&_AdvancedTokenStorage.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BZxVault() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.BZxVault(&_AdvancedTokenStorage.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BalanceOf(&_AdvancedTokenStorage.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BalanceOf(&_AdvancedTokenStorage.CallOpts, _owner)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "baseRate")
	return *ret0, err
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BaseRate() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BaseRate(&_AdvancedTokenStorage.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BaseRate() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BaseRate(&_AdvancedTokenStorage.CallOpts)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BurntTokenReserveList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Lender common.Address
		Amount *big.Int
	})
	out := ret
	err := _AdvancedTokenStorage.contract.Call(opts, out, "burntTokenReserveList", arg0)
	return *ret, err
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserveList(&_AdvancedTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserveList(&_AdvancedTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BurntTokenReserveListIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	ret := new(struct {
		Index *big.Int
		IsSet bool
	})
	out := ret
	err := _AdvancedTokenStorage.contract.Call(opts, out, "burntTokenReserveListIndex", arg0)
	return *ret, err
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserveListIndex(&_AdvancedTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserveListIndex(&_AdvancedTokenStorage.CallOpts, arg0)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) BurntTokenReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "burntTokenReserved")
	return *ret0, err
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) BurntTokenReserved() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserved(&_AdvancedTokenStorage.CallOpts)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) BurntTokenReserved() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.BurntTokenReserved(&_AdvancedTokenStorage.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) CheckpointSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "checkpointSupply")
	return *ret0, err
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) CheckpointSupply() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.CheckpointSupply(&_AdvancedTokenStorage.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) CheckpointSupply() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.CheckpointSupply(&_AdvancedTokenStorage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) Decimals() (uint8, error) {
	return _AdvancedTokenStorage.Contract.Decimals(&_AdvancedTokenStorage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) Decimals() (uint8, error) {
	return _AdvancedTokenStorage.Contract.Decimals(&_AdvancedTokenStorage.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) InitialPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "initialPrice")
	return *ret0, err
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) InitialPrice() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.InitialPrice(&_AdvancedTokenStorage.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) InitialPrice() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.InitialPrice(&_AdvancedTokenStorage.CallOpts)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) LeverageList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "leverageList", arg0)
	return *ret0, err
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.LeverageList(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.LeverageList(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) LoanOrderData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	ret := new(struct {
		LoanOrderHash               [32]byte
		LeverageAmount              *big.Int
		InitialMarginAmount         *big.Int
		MaintenanceMarginAmount     *big.Int
		MaxDurationUnixTimestampSec *big.Int
		Index                       *big.Int
		MarginPremiumAmount         *big.Int
		CollateralTokenAddress      common.Address
	})
	out := ret
	err := _AdvancedTokenStorage.contract.Call(opts, out, "loanOrderData", arg0)
	return *ret, err
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _AdvancedTokenStorage.Contract.LoanOrderData(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _AdvancedTokenStorage.Contract.LoanOrderData(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) LoanOrderHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "loanOrderHashes", arg0)
	return *ret0, err
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _AdvancedTokenStorage.Contract.LoanOrderHashes(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _AdvancedTokenStorage.Contract.LoanOrderHashes(&_AdvancedTokenStorage.CallOpts, arg0)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) LoanTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "loanTokenAddress")
	return *ret0, err
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) LoanTokenAddress() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.LoanTokenAddress(&_AdvancedTokenStorage.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) LoanTokenAddress() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.LoanTokenAddress(&_AdvancedTokenStorage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) Name() (string, error) {
	return _AdvancedTokenStorage.Contract.Name(&_AdvancedTokenStorage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) Name() (string, error) {
	return _AdvancedTokenStorage.Contract.Name(&_AdvancedTokenStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) Owner() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.Owner(&_AdvancedTokenStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) Owner() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.Owner(&_AdvancedTokenStorage.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) RateMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "rateMultiplier")
	return *ret0, err
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) RateMultiplier() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.RateMultiplier(&_AdvancedTokenStorage.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) RateMultiplier() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.RateMultiplier(&_AdvancedTokenStorage.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) SpreadMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "spreadMultiplier")
	return *ret0, err
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) SpreadMultiplier() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.SpreadMultiplier(&_AdvancedTokenStorage.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) SpreadMultiplier() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.SpreadMultiplier(&_AdvancedTokenStorage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) Symbol() (string, error) {
	return _AdvancedTokenStorage.Contract.Symbol(&_AdvancedTokenStorage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) Symbol() (string, error) {
	return _AdvancedTokenStorage.Contract.Symbol(&_AdvancedTokenStorage.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) TokenizedRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "tokenizedRegistry")
	return *ret0, err
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) TokenizedRegistry() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.TokenizedRegistry(&_AdvancedTokenStorage.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) TokenizedRegistry() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.TokenizedRegistry(&_AdvancedTokenStorage.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) TotalAssetBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "totalAssetBorrow")
	return *ret0, err
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) TotalAssetBorrow() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.TotalAssetBorrow(&_AdvancedTokenStorage.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) TotalAssetBorrow() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.TotalAssetBorrow(&_AdvancedTokenStorage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) TotalSupply() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.TotalSupply(&_AdvancedTokenStorage.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) TotalSupply() (*big.Int, error) {
	return _AdvancedTokenStorage.Contract.TotalSupply(&_AdvancedTokenStorage.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCaller) WethContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _AdvancedTokenStorage.contract.Call(opts, out, "wethContract")
	return *ret0, err
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) WethContract() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.WethContract(&_AdvancedTokenStorage.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_AdvancedTokenStorage *AdvancedTokenStorageCallerSession) WethContract() (common.Address, error) {
	return _AdvancedTokenStorage.Contract.WethContract(&_AdvancedTokenStorage.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedTokenStorage *AdvancedTokenStorageTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedTokenStorage.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedTokenStorage *AdvancedTokenStorageSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.TransferOwnership(&_AdvancedTokenStorage.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_AdvancedTokenStorage *AdvancedTokenStorageTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _AdvancedTokenStorage.Contract.TransferOwnership(&_AdvancedTokenStorage.TransactOpts, _newOwner)
}

// AdvancedTokenStorageApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageApprovalIterator struct {
	Event *AdvancedTokenStorageApproval // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageApproval)
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
		it.Event = new(AdvancedTokenStorageApproval)
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
func (it *AdvancedTokenStorageApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageApproval represents a Approval event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*AdvancedTokenStorageApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageApprovalIterator{contract: _AdvancedTokenStorage.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageApproval)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseApproval(log types.Log) (*AdvancedTokenStorageApproval, error) {
	event := new(AdvancedTokenStorageApproval)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageBorrowIterator struct {
	Event *AdvancedTokenStorageBorrow // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageBorrow)
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
		it.Event = new(AdvancedTokenStorageBorrow)
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
func (it *AdvancedTokenStorageBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageBorrow represents a Borrow event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageBorrow struct {
	Borrower                common.Address
	BorrowAmount            *big.Int
	InterestRate            *big.Int
	CollateralTokenAddress  common.Address
	TradeTokenToFillAddress common.Address
	WithdrawOnOpen          bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterBorrow(opts *bind.FilterOpts, borrower []common.Address) (*AdvancedTokenStorageBorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageBorrowIterator{contract: _AdvancedTokenStorage.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageBorrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageBorrow)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseBorrow(log types.Log) (*AdvancedTokenStorageBorrow, error) {
	event := new(AdvancedTokenStorageBorrow)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageBurnIterator struct {
	Event *AdvancedTokenStorageBurn // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageBurn)
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
		it.Event = new(AdvancedTokenStorageBurn)
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
func (it *AdvancedTokenStorageBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageBurn represents a Burn event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageBurn struct {
	Burner      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*AdvancedTokenStorageBurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageBurnIterator{contract: _AdvancedTokenStorage.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageBurn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageBurn)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseBurn(log types.Log) (*AdvancedTokenStorageBurn, error) {
	event := new(AdvancedTokenStorageBurn)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageClaimIterator struct {
	Event *AdvancedTokenStorageClaim // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageClaim)
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
		it.Event = new(AdvancedTokenStorageClaim)
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
func (it *AdvancedTokenStorageClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageClaim represents a Claim event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageClaim struct {
	Claimant             common.Address
	TokenAmount          *big.Int
	AssetAmount          *big.Int
	RemainingTokenAmount *big.Int
	Price                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterClaim(opts *bind.FilterOpts, claimant []common.Address) (*AdvancedTokenStorageClaimIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageClaimIterator{contract: _AdvancedTokenStorage.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageClaim, claimant []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageClaim)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseClaim(log types.Log) (*AdvancedTokenStorageClaim, error) {
	event := new(AdvancedTokenStorageClaim)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageMintIterator struct {
	Event *AdvancedTokenStorageMint // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageMint)
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
		it.Event = new(AdvancedTokenStorageMint)
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
func (it *AdvancedTokenStorageMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageMint represents a Mint event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageMint struct {
	Minter      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterMint(opts *bind.FilterOpts, minter []common.Address) (*AdvancedTokenStorageMintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageMintIterator{contract: _AdvancedTokenStorage.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageMint, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageMint)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseMint(log types.Log) (*AdvancedTokenStorageMint, error) {
	event := new(AdvancedTokenStorageMint)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageOwnershipTransferredIterator struct {
	Event *AdvancedTokenStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageOwnershipTransferred)
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
		it.Event = new(AdvancedTokenStorageOwnershipTransferred)
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
func (it *AdvancedTokenStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageOwnershipTransferred represents a OwnershipTransferred event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AdvancedTokenStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageOwnershipTransferredIterator{contract: _AdvancedTokenStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageOwnershipTransferred)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseOwnershipTransferred(log types.Log) (*AdvancedTokenStorageOwnershipTransferred, error) {
	event := new(AdvancedTokenStorageOwnershipTransferred)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageRepayIterator struct {
	Event *AdvancedTokenStorageRepay // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageRepay)
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
		it.Event = new(AdvancedTokenStorageRepay)
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
func (it *AdvancedTokenStorageRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageRepay represents a Repay event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageRepay struct {
	LoanOrderHash [32]byte
	Borrower      common.Address
	Closer        common.Address
	Amount        *big.Int
	IsLiquidation bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterRepay(opts *bind.FilterOpts, loanOrderHash [][32]byte, borrower []common.Address) (*AdvancedTokenStorageRepayIterator, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageRepayIterator{contract: _AdvancedTokenStorage.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageRepay, loanOrderHash [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageRepay)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseRepay(log types.Log) (*AdvancedTokenStorageRepay, error) {
	event := new(AdvancedTokenStorageRepay)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// AdvancedTokenStorageTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageTransferIterator struct {
	Event *AdvancedTokenStorageTransfer // Event containing the contract specifics and raw log

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
func (it *AdvancedTokenStorageTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AdvancedTokenStorageTransfer)
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
		it.Event = new(AdvancedTokenStorageTransfer)
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
func (it *AdvancedTokenStorageTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AdvancedTokenStorageTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AdvancedTokenStorageTransfer represents a Transfer event raised by the AdvancedTokenStorage contract.
type AdvancedTokenStorageTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*AdvancedTokenStorageTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &AdvancedTokenStorageTransferIterator{contract: _AdvancedTokenStorage.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *AdvancedTokenStorageTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _AdvancedTokenStorage.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AdvancedTokenStorageTransfer)
				if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_AdvancedTokenStorage *AdvancedTokenStorageFilterer) ParseTransfer(log types.Log) (*AdvancedTokenStorageTransfer, error) {
	event := new(AdvancedTokenStorageTransfer)
	if err := _AdvancedTokenStorage.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// BZxObjectsABI is the input ABI used to generate the binding from.
const BZxObjectsABI = "[]"

// BZxObjectsBin is the compiled bytecode used for deploying new contracts.
var BZxObjectsBin = "0x6080604052348015600f57600080fd5b50604c80601d6000396000f3fe6080604052600080fdfea365627a7a7231582045c07eb2a95aae47bec92d1b6e80a8330a4339a4e89401b686070aa11f79bbd26c6578706572696d656e74616cf564736f6c63430005110040"

// DeployBZxObjects deploys a new Ethereum contract, binding an instance of BZxObjects to it.
func DeployBZxObjects(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BZxObjects, error) {
	parsed, err := abi.JSON(strings.NewReader(BZxObjectsABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(BZxObjectsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BZxObjects{BZxObjectsCaller: BZxObjectsCaller{contract: contract}, BZxObjectsTransactor: BZxObjectsTransactor{contract: contract}, BZxObjectsFilterer: BZxObjectsFilterer{contract: contract}}, nil
}

// BZxObjects is an auto generated Go binding around an Ethereum contract.
type BZxObjects struct {
	BZxObjectsCaller     // Read-only binding to the contract
	BZxObjectsTransactor // Write-only binding to the contract
	BZxObjectsFilterer   // Log filterer for contract events
}

// BZxObjectsCaller is an auto generated read-only Go binding around an Ethereum contract.
type BZxObjectsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BZxObjectsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BZxObjectsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BZxObjectsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BZxObjectsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BZxObjectsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BZxObjectsSession struct {
	Contract     *BZxObjects       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BZxObjectsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BZxObjectsCallerSession struct {
	Contract *BZxObjectsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// BZxObjectsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BZxObjectsTransactorSession struct {
	Contract     *BZxObjectsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// BZxObjectsRaw is an auto generated low-level Go binding around an Ethereum contract.
type BZxObjectsRaw struct {
	Contract *BZxObjects // Generic contract binding to access the raw methods on
}

// BZxObjectsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BZxObjectsCallerRaw struct {
	Contract *BZxObjectsCaller // Generic read-only contract binding to access the raw methods on
}

// BZxObjectsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BZxObjectsTransactorRaw struct {
	Contract *BZxObjectsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBZxObjects creates a new instance of BZxObjects, bound to a specific deployed contract.
func NewBZxObjects(address common.Address, backend bind.ContractBackend) (*BZxObjects, error) {
	contract, err := bindBZxObjects(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BZxObjects{BZxObjectsCaller: BZxObjectsCaller{contract: contract}, BZxObjectsTransactor: BZxObjectsTransactor{contract: contract}, BZxObjectsFilterer: BZxObjectsFilterer{contract: contract}}, nil
}

// NewBZxObjectsCaller creates a new read-only instance of BZxObjects, bound to a specific deployed contract.
func NewBZxObjectsCaller(address common.Address, caller bind.ContractCaller) (*BZxObjectsCaller, error) {
	contract, err := bindBZxObjects(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BZxObjectsCaller{contract: contract}, nil
}

// NewBZxObjectsTransactor creates a new write-only instance of BZxObjects, bound to a specific deployed contract.
func NewBZxObjectsTransactor(address common.Address, transactor bind.ContractTransactor) (*BZxObjectsTransactor, error) {
	contract, err := bindBZxObjects(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BZxObjectsTransactor{contract: contract}, nil
}

// NewBZxObjectsFilterer creates a new log filterer instance of BZxObjects, bound to a specific deployed contract.
func NewBZxObjectsFilterer(address common.Address, filterer bind.ContractFilterer) (*BZxObjectsFilterer, error) {
	contract, err := bindBZxObjects(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BZxObjectsFilterer{contract: contract}, nil
}

// bindBZxObjects binds a generic wrapper to an already deployed contract.
func bindBZxObjects(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BZxObjectsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BZxObjects *BZxObjectsRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BZxObjects.Contract.BZxObjectsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BZxObjects *BZxObjectsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BZxObjects.Contract.BZxObjectsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BZxObjects *BZxObjectsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BZxObjects.Contract.BZxObjectsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BZxObjects *BZxObjectsCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _BZxObjects.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BZxObjects *BZxObjectsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BZxObjects.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BZxObjects *BZxObjectsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BZxObjects.Contract.contract.Transact(opts, method, params...)
}

// EIP20ABI is the input ABI used to generate the binding from.
const EIP20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// EIP20FuncSigs maps the 4-byte function signature to its string representation.
var EIP20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// EIP20 is an auto generated Go binding around an Ethereum contract.
type EIP20 struct {
	EIP20Caller     // Read-only binding to the contract
	EIP20Transactor // Write-only binding to the contract
	EIP20Filterer   // Log filterer for contract events
}

// EIP20Caller is an auto generated read-only Go binding around an Ethereum contract.
type EIP20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type EIP20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EIP20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EIP20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EIP20Session struct {
	Contract     *EIP20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EIP20CallerSession struct {
	Contract *EIP20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// EIP20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EIP20TransactorSession struct {
	Contract     *EIP20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// EIP20Raw is an auto generated low-level Go binding around an Ethereum contract.
type EIP20Raw struct {
	Contract *EIP20 // Generic contract binding to access the raw methods on
}

// EIP20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EIP20CallerRaw struct {
	Contract *EIP20Caller // Generic read-only contract binding to access the raw methods on
}

// EIP20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EIP20TransactorRaw struct {
	Contract *EIP20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewEIP20 creates a new instance of EIP20, bound to a specific deployed contract.
func NewEIP20(address common.Address, backend bind.ContractBackend) (*EIP20, error) {
	contract, err := bindEIP20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EIP20{EIP20Caller: EIP20Caller{contract: contract}, EIP20Transactor: EIP20Transactor{contract: contract}, EIP20Filterer: EIP20Filterer{contract: contract}}, nil
}

// NewEIP20Caller creates a new read-only instance of EIP20, bound to a specific deployed contract.
func NewEIP20Caller(address common.Address, caller bind.ContractCaller) (*EIP20Caller, error) {
	contract, err := bindEIP20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20Caller{contract: contract}, nil
}

// NewEIP20Transactor creates a new write-only instance of EIP20, bound to a specific deployed contract.
func NewEIP20Transactor(address common.Address, transactor bind.ContractTransactor) (*EIP20Transactor, error) {
	contract, err := bindEIP20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EIP20Transactor{contract: contract}, nil
}

// NewEIP20Filterer creates a new log filterer instance of EIP20, bound to a specific deployed contract.
func NewEIP20Filterer(address common.Address, filterer bind.ContractFilterer) (*EIP20Filterer, error) {
	contract, err := bindEIP20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EIP20Filterer{contract: contract}, nil
}

// bindEIP20 binds a generic wrapper to an already deployed contract.
func bindEIP20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EIP20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20 *EIP20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP20.Contract.EIP20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20 *EIP20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20.Contract.EIP20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20 *EIP20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20.Contract.EIP20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EIP20 *EIP20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _EIP20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EIP20 *EIP20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EIP20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EIP20 *EIP20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EIP20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EIP20 *EIP20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EIP20 *EIP20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EIP20.Contract.Allowance(&_EIP20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_EIP20 *EIP20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _EIP20.Contract.Allowance(&_EIP20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_EIP20 *EIP20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_EIP20 *EIP20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _EIP20.Contract.BalanceOf(&_EIP20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_EIP20 *EIP20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _EIP20.Contract.BalanceOf(&_EIP20.CallOpts, _who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EIP20 *EIP20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EIP20 *EIP20Session) Decimals() (uint8, error) {
	return _EIP20.Contract.Decimals(&_EIP20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_EIP20 *EIP20CallerSession) Decimals() (uint8, error) {
	return _EIP20.Contract.Decimals(&_EIP20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EIP20 *EIP20Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EIP20 *EIP20Session) Name() (string, error) {
	return _EIP20.Contract.Name(&_EIP20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_EIP20 *EIP20CallerSession) Name() (string, error) {
	return _EIP20.Contract.Name(&_EIP20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EIP20 *EIP20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EIP20 *EIP20Session) Symbol() (string, error) {
	return _EIP20.Contract.Symbol(&_EIP20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_EIP20 *EIP20CallerSession) Symbol() (string, error) {
	return _EIP20.Contract.Symbol(&_EIP20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20 *EIP20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _EIP20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20 *EIP20Session) TotalSupply() (*big.Int, error) {
	return _EIP20.Contract.TotalSupply(&_EIP20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_EIP20 *EIP20CallerSession) TotalSupply() (*big.Int, error) {
	return _EIP20.Contract.TotalSupply(&_EIP20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EIP20 *EIP20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EIP20 *EIP20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.Approve(&_EIP20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_EIP20 *EIP20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.Approve(&_EIP20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.Transfer(&_EIP20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.Transfer(&_EIP20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.TransferFrom(&_EIP20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_EIP20 *EIP20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _EIP20.Contract.TransferFrom(&_EIP20.TransactOpts, _from, _to, _value)
}

// EIP20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the EIP20 contract.
type EIP20ApprovalIterator struct {
	Event *EIP20Approval // Event containing the contract specifics and raw log

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
func (it *EIP20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20Approval)
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
		it.Event = new(EIP20Approval)
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
func (it *EIP20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20Approval represents a Approval event raised by the EIP20 contract.
type EIP20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EIP20 *EIP20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*EIP20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &EIP20ApprovalIterator{contract: _EIP20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EIP20 *EIP20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *EIP20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _EIP20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20Approval)
				if err := _EIP20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_EIP20 *EIP20Filterer) ParseApproval(log types.Log) (*EIP20Approval, error) {
	event := new(EIP20Approval)
	if err := _EIP20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// EIP20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the EIP20 contract.
type EIP20TransferIterator struct {
	Event *EIP20Transfer // Event containing the contract specifics and raw log

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
func (it *EIP20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EIP20Transfer)
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
		it.Event = new(EIP20Transfer)
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
func (it *EIP20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EIP20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EIP20Transfer represents a Transfer event raised by the EIP20 contract.
type EIP20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EIP20 *EIP20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*EIP20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &EIP20TransferIterator{contract: _EIP20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EIP20 *EIP20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *EIP20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _EIP20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EIP20Transfer)
				if err := _EIP20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_EIP20 *EIP20Filterer) ParseTransfer(log types.Log) (*EIP20Transfer, error) {
	event := new(EIP20Transfer)
	if err := _EIP20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20ABI is the input ABI used to generate the binding from.
const ERC20ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20FuncSigs maps the 4-byte function signature to its string representation.
var ERC20FuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC20 is an auto generated Go binding around an Ethereum contract.
type ERC20 struct {
	ERC20Caller     // Read-only binding to the contract
	ERC20Transactor // Write-only binding to the contract
	ERC20Filterer   // Log filterer for contract events
}

// ERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20Session struct {
	Contract     *ERC20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20CallerSession struct {
	Contract *ERC20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20TransactorSession struct {
	Contract     *ERC20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20Raw struct {
	Contract *ERC20 // Generic contract binding to access the raw methods on
}

// ERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20CallerRaw struct {
	Contract *ERC20Caller // Generic read-only contract binding to access the raw methods on
}

// ERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20TransactorRaw struct {
	Contract *ERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20 creates a new instance of ERC20, bound to a specific deployed contract.
func NewERC20(address common.Address, backend bind.ContractBackend) (*ERC20, error) {
	contract, err := bindERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20{ERC20Caller: ERC20Caller{contract: contract}, ERC20Transactor: ERC20Transactor{contract: contract}, ERC20Filterer: ERC20Filterer{contract: contract}}, nil
}

// NewERC20Caller creates a new read-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Caller(address common.Address, caller bind.ContractCaller) (*ERC20Caller, error) {
	contract, err := bindERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Caller{contract: contract}, nil
}

// NewERC20Transactor creates a new write-only instance of ERC20, bound to a specific deployed contract.
func NewERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC20Transactor, error) {
	contract, err := bindERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20Transactor{contract: contract}, nil
}

// NewERC20Filterer creates a new log filterer instance of ERC20, bound to a specific deployed contract.
func NewERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC20Filterer, error) {
	contract, err := bindERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20Filterer{contract: contract}, nil
}

// bindERC20 binds a generic wrapper to an already deployed contract.
func bindERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.ERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.ERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20 *ERC20CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20 *ERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20 *ERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20 *ERC20Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20 *ERC20Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_ERC20 *ERC20CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _ERC20.Contract.Allowance(&_ERC20.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20 *ERC20Caller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20 *ERC20Session) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20 *ERC20CallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20.Contract.BalanceOf(&_ERC20.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20Session) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20 *ERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20.Contract.TotalSupply(&_ERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Approve(&_ERC20.TransactOpts, _spender, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.Transfer(&_ERC20.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_ERC20 *ERC20TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20.Contract.TransferFrom(&_ERC20.TransactOpts, _from, _to, _value)
}

// ERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC20 contract.
type ERC20ApprovalIterator struct {
	Event *ERC20Approval // Event containing the contract specifics and raw log

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
func (it *ERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Approval)
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
		it.Event = new(ERC20Approval)
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
func (it *ERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Approval represents a Approval event raised by the ERC20 contract.
type ERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*ERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &ERC20ApprovalIterator{contract: _ERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Approval)
				if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_ERC20 *ERC20Filterer) ParseApproval(log types.Log) (*ERC20Approval, error) {
	event := new(ERC20Approval)
	if err := _ERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20 contract.
type ERC20TransferIterator struct {
	Event *ERC20Transfer // Event containing the contract specifics and raw log

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
func (it *ERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20Transfer)
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
		it.Event = new(ERC20Transfer)
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
func (it *ERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20Transfer represents a Transfer event raised by the ERC20 contract.
type ERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20TransferIterator{contract: _ERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20Transfer)
				if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20 *ERC20Filterer) ParseTransfer(log types.Log) (*ERC20Transfer, error) {
	event := new(ERC20Transfer)
	if err := _ERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ERC20BasicABI is the input ABI used to generate the binding from.
const ERC20BasicABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// ERC20BasicFuncSigs maps the 4-byte function signature to its string representation.
var ERC20BasicFuncSigs = map[string]string{
	"70a08231": "balanceOf(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
}

// ERC20Basic is an auto generated Go binding around an Ethereum contract.
type ERC20Basic struct {
	ERC20BasicCaller     // Read-only binding to the contract
	ERC20BasicTransactor // Write-only binding to the contract
	ERC20BasicFilterer   // Log filterer for contract events
}

// ERC20BasicCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC20BasicCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC20BasicTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC20BasicFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC20BasicSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC20BasicSession struct {
	Contract     *ERC20Basic       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC20BasicCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC20BasicCallerSession struct {
	Contract *ERC20BasicCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// ERC20BasicTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC20BasicTransactorSession struct {
	Contract     *ERC20BasicTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ERC20BasicRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC20BasicRaw struct {
	Contract *ERC20Basic // Generic contract binding to access the raw methods on
}

// ERC20BasicCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC20BasicCallerRaw struct {
	Contract *ERC20BasicCaller // Generic read-only contract binding to access the raw methods on
}

// ERC20BasicTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC20BasicTransactorRaw struct {
	Contract *ERC20BasicTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC20Basic creates a new instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20Basic(address common.Address, backend bind.ContractBackend) (*ERC20Basic, error) {
	contract, err := bindERC20Basic(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC20Basic{ERC20BasicCaller: ERC20BasicCaller{contract: contract}, ERC20BasicTransactor: ERC20BasicTransactor{contract: contract}, ERC20BasicFilterer: ERC20BasicFilterer{contract: contract}}, nil
}

// NewERC20BasicCaller creates a new read-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicCaller(address common.Address, caller bind.ContractCaller) (*ERC20BasicCaller, error) {
	contract, err := bindERC20Basic(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicCaller{contract: contract}, nil
}

// NewERC20BasicTransactor creates a new write-only instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC20BasicTransactor, error) {
	contract, err := bindERC20Basic(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransactor{contract: contract}, nil
}

// NewERC20BasicFilterer creates a new log filterer instance of ERC20Basic, bound to a specific deployed contract.
func NewERC20BasicFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC20BasicFilterer, error) {
	contract, err := bindERC20Basic(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicFilterer{contract: contract}, nil
}

// bindERC20Basic binds a generic wrapper to an already deployed contract.
func bindERC20Basic(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC20BasicABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.ERC20BasicCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.ERC20BasicTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC20Basic *ERC20BasicCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ERC20Basic.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC20Basic *ERC20BasicTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC20Basic.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _ERC20Basic.Contract.BalanceOf(&_ERC20Basic.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _ERC20Basic.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_ERC20Basic *ERC20BasicCallerSession) TotalSupply() (*big.Int, error) {
	return _ERC20Basic.Contract.TotalSupply(&_ERC20Basic.CallOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_ERC20Basic *ERC20BasicTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _ERC20Basic.Contract.Transfer(&_ERC20Basic.TransactOpts, _to, _value)
}

// ERC20BasicTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC20Basic contract.
type ERC20BasicTransferIterator struct {
	Event *ERC20BasicTransfer // Event containing the contract specifics and raw log

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
func (it *ERC20BasicTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC20BasicTransfer)
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
		it.Event = new(ERC20BasicTransfer)
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
func (it *ERC20BasicTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC20BasicTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC20BasicTransfer represents a Transfer event raised by the ERC20Basic contract.
type ERC20BasicTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*ERC20BasicTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ERC20BasicTransferIterator{contract: _ERC20Basic.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC20BasicTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ERC20Basic.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC20BasicTransfer)
				if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_ERC20Basic *ERC20BasicFilterer) ParseTransfer(log types.Log) (*ERC20BasicTransfer, error) {
	event := new(ERC20BasicTransfer)
	if err := _ERC20Basic.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// IBZxABI is the input ABI used to generate the binding from.
const IBZxABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"loanTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"collateralTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginAmount\",\"type\":\"uint256\"}],\"name\":\"getBorrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestTokenAddress\",\"type\":\"address\"}],\"name\":\"getLenderInterestForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"interestPaid\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestPaidDate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestOwedPerDay\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestUnPaid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"loanTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"newLoanAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginAmount\",\"type\":\"uint256\"}],\"name\":\"getRequiredCollateral\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"collateralTokenAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"}],\"name\":\"oracleAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestTokenAddress\",\"type\":\"address\"}],\"name\":\"payInterestForOracle\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"address[4]\",\"name\":\"sentAddresses\",\"type\":\"address[4]\"},{\"internalType\":\"uint256[7]\",\"name\":\"sentAmounts\",\"type\":\"uint256[7]\"},{\"internalType\":\"bytes\",\"name\":\"loanDataBytes\",\"type\":\"bytes\"}],\"name\":\"takeOrderFromiToken\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"}]"

// IBZxFuncSigs maps the 4-byte function signature to its string representation.
var IBZxFuncSigs = map[string]string{
	"f3d75a9c": "getBorrowAmount(address,address,address,uint256,uint256)",
	"0a90b578": "getLenderInterestForOracle(address,address,address)",
	"bc6cb1d9": "getRequiredCollateral(address,address,address,uint256,uint256)",
	"71eb125e": "oracleAddresses(address)",
	"327ab639": "payInterestForOracle(address,address)",
	"b1eac3ad": "takeOrderFromiToken(bytes32,address[4],uint256[7],bytes)",
}

// IBZx is an auto generated Go binding around an Ethereum contract.
type IBZx struct {
	IBZxCaller     // Read-only binding to the contract
	IBZxTransactor // Write-only binding to the contract
	IBZxFilterer   // Log filterer for contract events
}

// IBZxCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBZxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBZxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBZxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBZxSession struct {
	Contract     *IBZx             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBZxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBZxCallerSession struct {
	Contract *IBZxCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IBZxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBZxTransactorSession struct {
	Contract     *IBZxTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBZxRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBZxRaw struct {
	Contract *IBZx // Generic contract binding to access the raw methods on
}

// IBZxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBZxCallerRaw struct {
	Contract *IBZxCaller // Generic read-only contract binding to access the raw methods on
}

// IBZxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBZxTransactorRaw struct {
	Contract *IBZxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBZx creates a new instance of IBZx, bound to a specific deployed contract.
func NewIBZx(address common.Address, backend bind.ContractBackend) (*IBZx, error) {
	contract, err := bindIBZx(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBZx{IBZxCaller: IBZxCaller{contract: contract}, IBZxTransactor: IBZxTransactor{contract: contract}, IBZxFilterer: IBZxFilterer{contract: contract}}, nil
}

// NewIBZxCaller creates a new read-only instance of IBZx, bound to a specific deployed contract.
func NewIBZxCaller(address common.Address, caller bind.ContractCaller) (*IBZxCaller, error) {
	contract, err := bindIBZx(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBZxCaller{contract: contract}, nil
}

// NewIBZxTransactor creates a new write-only instance of IBZx, bound to a specific deployed contract.
func NewIBZxTransactor(address common.Address, transactor bind.ContractTransactor) (*IBZxTransactor, error) {
	contract, err := bindIBZx(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBZxTransactor{contract: contract}, nil
}

// NewIBZxFilterer creates a new log filterer instance of IBZx, bound to a specific deployed contract.
func NewIBZxFilterer(address common.Address, filterer bind.ContractFilterer) (*IBZxFilterer, error) {
	contract, err := bindIBZx(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBZxFilterer{contract: contract}, nil
}

// bindIBZx binds a generic wrapper to an already deployed contract.
func bindIBZx(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBZxABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBZx *IBZxRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBZx.Contract.IBZxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBZx *IBZxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBZx.Contract.IBZxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBZx *IBZxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBZx.Contract.IBZxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBZx *IBZxCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBZx.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBZx *IBZxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBZx.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBZx *IBZxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBZx.Contract.contract.Transact(opts, method, params...)
}

// GetBorrowAmount is a free data retrieval call binding the contract method 0xf3d75a9c.
//
// Solidity: function getBorrowAmount(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 collateralTokenAmount, uint256 marginAmount) view returns(uint256 borrowAmount)
func (_IBZx *IBZxCaller) GetBorrowAmount(opts *bind.CallOpts, loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, collateralTokenAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBZx.contract.Call(opts, out, "getBorrowAmount", loanTokenAddress, collateralTokenAddress, oracleAddress, collateralTokenAmount, marginAmount)
	return *ret0, err
}

// GetBorrowAmount is a free data retrieval call binding the contract method 0xf3d75a9c.
//
// Solidity: function getBorrowAmount(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 collateralTokenAmount, uint256 marginAmount) view returns(uint256 borrowAmount)
func (_IBZx *IBZxSession) GetBorrowAmount(loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, collateralTokenAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	return _IBZx.Contract.GetBorrowAmount(&_IBZx.CallOpts, loanTokenAddress, collateralTokenAddress, oracleAddress, collateralTokenAmount, marginAmount)
}

// GetBorrowAmount is a free data retrieval call binding the contract method 0xf3d75a9c.
//
// Solidity: function getBorrowAmount(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 collateralTokenAmount, uint256 marginAmount) view returns(uint256 borrowAmount)
func (_IBZx *IBZxCallerSession) GetBorrowAmount(loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, collateralTokenAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	return _IBZx.Contract.GetBorrowAmount(&_IBZx.CallOpts, loanTokenAddress, collateralTokenAddress, oracleAddress, collateralTokenAmount, marginAmount)
}

// GetLenderInterestForOracle is a free data retrieval call binding the contract method 0x0a90b578.
//
// Solidity: function getLenderInterestForOracle(address lender, address oracleAddress, address interestTokenAddress) view returns(uint256 interestPaid, uint256 interestPaidDate, uint256 interestOwedPerDay, uint256 interestUnPaid)
func (_IBZx *IBZxCaller) GetLenderInterestForOracle(opts *bind.CallOpts, lender common.Address, oracleAddress common.Address, interestTokenAddress common.Address) (struct {
	InterestPaid       *big.Int
	InterestPaidDate   *big.Int
	InterestOwedPerDay *big.Int
	InterestUnPaid     *big.Int
}, error) {
	ret := new(struct {
		InterestPaid       *big.Int
		InterestPaidDate   *big.Int
		InterestOwedPerDay *big.Int
		InterestUnPaid     *big.Int
	})
	out := ret
	err := _IBZx.contract.Call(opts, out, "getLenderInterestForOracle", lender, oracleAddress, interestTokenAddress)
	return *ret, err
}

// GetLenderInterestForOracle is a free data retrieval call binding the contract method 0x0a90b578.
//
// Solidity: function getLenderInterestForOracle(address lender, address oracleAddress, address interestTokenAddress) view returns(uint256 interestPaid, uint256 interestPaidDate, uint256 interestOwedPerDay, uint256 interestUnPaid)
func (_IBZx *IBZxSession) GetLenderInterestForOracle(lender common.Address, oracleAddress common.Address, interestTokenAddress common.Address) (struct {
	InterestPaid       *big.Int
	InterestPaidDate   *big.Int
	InterestOwedPerDay *big.Int
	InterestUnPaid     *big.Int
}, error) {
	return _IBZx.Contract.GetLenderInterestForOracle(&_IBZx.CallOpts, lender, oracleAddress, interestTokenAddress)
}

// GetLenderInterestForOracle is a free data retrieval call binding the contract method 0x0a90b578.
//
// Solidity: function getLenderInterestForOracle(address lender, address oracleAddress, address interestTokenAddress) view returns(uint256 interestPaid, uint256 interestPaidDate, uint256 interestOwedPerDay, uint256 interestUnPaid)
func (_IBZx *IBZxCallerSession) GetLenderInterestForOracle(lender common.Address, oracleAddress common.Address, interestTokenAddress common.Address) (struct {
	InterestPaid       *big.Int
	InterestPaidDate   *big.Int
	InterestOwedPerDay *big.Int
	InterestUnPaid     *big.Int
}, error) {
	return _IBZx.Contract.GetLenderInterestForOracle(&_IBZx.CallOpts, lender, oracleAddress, interestTokenAddress)
}

// GetRequiredCollateral is a free data retrieval call binding the contract method 0xbc6cb1d9.
//
// Solidity: function getRequiredCollateral(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 newLoanAmount, uint256 marginAmount) view returns(uint256 collateralTokenAmount)
func (_IBZx *IBZxCaller) GetRequiredCollateral(opts *bind.CallOpts, loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, newLoanAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _IBZx.contract.Call(opts, out, "getRequiredCollateral", loanTokenAddress, collateralTokenAddress, oracleAddress, newLoanAmount, marginAmount)
	return *ret0, err
}

// GetRequiredCollateral is a free data retrieval call binding the contract method 0xbc6cb1d9.
//
// Solidity: function getRequiredCollateral(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 newLoanAmount, uint256 marginAmount) view returns(uint256 collateralTokenAmount)
func (_IBZx *IBZxSession) GetRequiredCollateral(loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, newLoanAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	return _IBZx.Contract.GetRequiredCollateral(&_IBZx.CallOpts, loanTokenAddress, collateralTokenAddress, oracleAddress, newLoanAmount, marginAmount)
}

// GetRequiredCollateral is a free data retrieval call binding the contract method 0xbc6cb1d9.
//
// Solidity: function getRequiredCollateral(address loanTokenAddress, address collateralTokenAddress, address oracleAddress, uint256 newLoanAmount, uint256 marginAmount) view returns(uint256 collateralTokenAmount)
func (_IBZx *IBZxCallerSession) GetRequiredCollateral(loanTokenAddress common.Address, collateralTokenAddress common.Address, oracleAddress common.Address, newLoanAmount *big.Int, marginAmount *big.Int) (*big.Int, error) {
	return _IBZx.Contract.GetRequiredCollateral(&_IBZx.CallOpts, loanTokenAddress, collateralTokenAddress, oracleAddress, newLoanAmount, marginAmount)
}

// OracleAddresses is a free data retrieval call binding the contract method 0x71eb125e.
//
// Solidity: function oracleAddresses(address oracleAddress) view returns(address)
func (_IBZx *IBZxCaller) OracleAddresses(opts *bind.CallOpts, oracleAddress common.Address) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _IBZx.contract.Call(opts, out, "oracleAddresses", oracleAddress)
	return *ret0, err
}

// OracleAddresses is a free data retrieval call binding the contract method 0x71eb125e.
//
// Solidity: function oracleAddresses(address oracleAddress) view returns(address)
func (_IBZx *IBZxSession) OracleAddresses(oracleAddress common.Address) (common.Address, error) {
	return _IBZx.Contract.OracleAddresses(&_IBZx.CallOpts, oracleAddress)
}

// OracleAddresses is a free data retrieval call binding the contract method 0x71eb125e.
//
// Solidity: function oracleAddresses(address oracleAddress) view returns(address)
func (_IBZx *IBZxCallerSession) OracleAddresses(oracleAddress common.Address) (common.Address, error) {
	return _IBZx.Contract.OracleAddresses(&_IBZx.CallOpts, oracleAddress)
}

// PayInterestForOracle is a paid mutator transaction binding the contract method 0x327ab639.
//
// Solidity: function payInterestForOracle(address oracleAddress, address interestTokenAddress) returns(uint256)
func (_IBZx *IBZxTransactor) PayInterestForOracle(opts *bind.TransactOpts, oracleAddress common.Address, interestTokenAddress common.Address) (*types.Transaction, error) {
	return _IBZx.contract.Transact(opts, "payInterestForOracle", oracleAddress, interestTokenAddress)
}

// PayInterestForOracle is a paid mutator transaction binding the contract method 0x327ab639.
//
// Solidity: function payInterestForOracle(address oracleAddress, address interestTokenAddress) returns(uint256)
func (_IBZx *IBZxSession) PayInterestForOracle(oracleAddress common.Address, interestTokenAddress common.Address) (*types.Transaction, error) {
	return _IBZx.Contract.PayInterestForOracle(&_IBZx.TransactOpts, oracleAddress, interestTokenAddress)
}

// PayInterestForOracle is a paid mutator transaction binding the contract method 0x327ab639.
//
// Solidity: function payInterestForOracle(address oracleAddress, address interestTokenAddress) returns(uint256)
func (_IBZx *IBZxTransactorSession) PayInterestForOracle(oracleAddress common.Address, interestTokenAddress common.Address) (*types.Transaction, error) {
	return _IBZx.Contract.PayInterestForOracle(&_IBZx.TransactOpts, oracleAddress, interestTokenAddress)
}

// TakeOrderFromiToken is a paid mutator transaction binding the contract method 0xb1eac3ad.
//
// Solidity: function takeOrderFromiToken(bytes32 loanOrderHash, address[4] sentAddresses, uint256[7] sentAmounts, bytes loanDataBytes) payable returns(uint256)
func (_IBZx *IBZxTransactor) TakeOrderFromiToken(opts *bind.TransactOpts, loanOrderHash [32]byte, sentAddresses [4]common.Address, sentAmounts [7]*big.Int, loanDataBytes []byte) (*types.Transaction, error) {
	return _IBZx.contract.Transact(opts, "takeOrderFromiToken", loanOrderHash, sentAddresses, sentAmounts, loanDataBytes)
}

// TakeOrderFromiToken is a paid mutator transaction binding the contract method 0xb1eac3ad.
//
// Solidity: function takeOrderFromiToken(bytes32 loanOrderHash, address[4] sentAddresses, uint256[7] sentAmounts, bytes loanDataBytes) payable returns(uint256)
func (_IBZx *IBZxSession) TakeOrderFromiToken(loanOrderHash [32]byte, sentAddresses [4]common.Address, sentAmounts [7]*big.Int, loanDataBytes []byte) (*types.Transaction, error) {
	return _IBZx.Contract.TakeOrderFromiToken(&_IBZx.TransactOpts, loanOrderHash, sentAddresses, sentAmounts, loanDataBytes)
}

// TakeOrderFromiToken is a paid mutator transaction binding the contract method 0xb1eac3ad.
//
// Solidity: function takeOrderFromiToken(bytes32 loanOrderHash, address[4] sentAddresses, uint256[7] sentAmounts, bytes loanDataBytes) payable returns(uint256)
func (_IBZx *IBZxTransactorSession) TakeOrderFromiToken(loanOrderHash [32]byte, sentAddresses [4]common.Address, sentAmounts [7]*big.Int, loanDataBytes []byte) (*types.Transaction, error) {
	return _IBZx.Contract.TakeOrderFromiToken(&_IBZx.TransactOpts, loanOrderHash, sentAddresses, sentAmounts, loanDataBytes)
}

// IBZxOracleABI is the input ABI used to generate the binding from.
const IBZxOracleABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"sourceTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"destTokenAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"sourceTokenAmount\",\"type\":\"uint256\"}],\"name\":\"getTradeData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"sourceToDestRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"sourceToDestPrecision\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"destTokenAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// IBZxOracleFuncSigs maps the 4-byte function signature to its string representation.
var IBZxOracleFuncSigs = map[string]string{
	"06599aa0": "getTradeData(address,address,uint256)",
}

// IBZxOracle is an auto generated Go binding around an Ethereum contract.
type IBZxOracle struct {
	IBZxOracleCaller     // Read-only binding to the contract
	IBZxOracleTransactor // Write-only binding to the contract
	IBZxOracleFilterer   // Log filterer for contract events
}

// IBZxOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type IBZxOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IBZxOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IBZxOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IBZxOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IBZxOracleSession struct {
	Contract     *IBZxOracle       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IBZxOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IBZxOracleCallerSession struct {
	Contract *IBZxOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// IBZxOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IBZxOracleTransactorSession struct {
	Contract     *IBZxOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// IBZxOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type IBZxOracleRaw struct {
	Contract *IBZxOracle // Generic contract binding to access the raw methods on
}

// IBZxOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IBZxOracleCallerRaw struct {
	Contract *IBZxOracleCaller // Generic read-only contract binding to access the raw methods on
}

// IBZxOracleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IBZxOracleTransactorRaw struct {
	Contract *IBZxOracleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIBZxOracle creates a new instance of IBZxOracle, bound to a specific deployed contract.
func NewIBZxOracle(address common.Address, backend bind.ContractBackend) (*IBZxOracle, error) {
	contract, err := bindIBZxOracle(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IBZxOracle{IBZxOracleCaller: IBZxOracleCaller{contract: contract}, IBZxOracleTransactor: IBZxOracleTransactor{contract: contract}, IBZxOracleFilterer: IBZxOracleFilterer{contract: contract}}, nil
}

// NewIBZxOracleCaller creates a new read-only instance of IBZxOracle, bound to a specific deployed contract.
func NewIBZxOracleCaller(address common.Address, caller bind.ContractCaller) (*IBZxOracleCaller, error) {
	contract, err := bindIBZxOracle(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IBZxOracleCaller{contract: contract}, nil
}

// NewIBZxOracleTransactor creates a new write-only instance of IBZxOracle, bound to a specific deployed contract.
func NewIBZxOracleTransactor(address common.Address, transactor bind.ContractTransactor) (*IBZxOracleTransactor, error) {
	contract, err := bindIBZxOracle(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IBZxOracleTransactor{contract: contract}, nil
}

// NewIBZxOracleFilterer creates a new log filterer instance of IBZxOracle, bound to a specific deployed contract.
func NewIBZxOracleFilterer(address common.Address, filterer bind.ContractFilterer) (*IBZxOracleFilterer, error) {
	contract, err := bindIBZxOracle(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IBZxOracleFilterer{contract: contract}, nil
}

// bindIBZxOracle binds a generic wrapper to an already deployed contract.
func bindIBZxOracle(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IBZxOracleABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBZxOracle *IBZxOracleRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBZxOracle.Contract.IBZxOracleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBZxOracle *IBZxOracleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBZxOracle.Contract.IBZxOracleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBZxOracle *IBZxOracleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBZxOracle.Contract.IBZxOracleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IBZxOracle *IBZxOracleCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IBZxOracle.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IBZxOracle *IBZxOracleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IBZxOracle.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IBZxOracle *IBZxOracleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IBZxOracle.Contract.contract.Transact(opts, method, params...)
}

// GetTradeData is a free data retrieval call binding the contract method 0x06599aa0.
//
// Solidity: function getTradeData(address sourceTokenAddress, address destTokenAddress, uint256 sourceTokenAmount) view returns(uint256 sourceToDestRate, uint256 sourceToDestPrecision, uint256 destTokenAmount)
func (_IBZxOracle *IBZxOracleCaller) GetTradeData(opts *bind.CallOpts, sourceTokenAddress common.Address, destTokenAddress common.Address, sourceTokenAmount *big.Int) (struct {
	SourceToDestRate      *big.Int
	SourceToDestPrecision *big.Int
	DestTokenAmount       *big.Int
}, error) {
	ret := new(struct {
		SourceToDestRate      *big.Int
		SourceToDestPrecision *big.Int
		DestTokenAmount       *big.Int
	})
	out := ret
	err := _IBZxOracle.contract.Call(opts, out, "getTradeData", sourceTokenAddress, destTokenAddress, sourceTokenAmount)
	return *ret, err
}

// GetTradeData is a free data retrieval call binding the contract method 0x06599aa0.
//
// Solidity: function getTradeData(address sourceTokenAddress, address destTokenAddress, uint256 sourceTokenAmount) view returns(uint256 sourceToDestRate, uint256 sourceToDestPrecision, uint256 destTokenAmount)
func (_IBZxOracle *IBZxOracleSession) GetTradeData(sourceTokenAddress common.Address, destTokenAddress common.Address, sourceTokenAmount *big.Int) (struct {
	SourceToDestRate      *big.Int
	SourceToDestPrecision *big.Int
	DestTokenAmount       *big.Int
}, error) {
	return _IBZxOracle.Contract.GetTradeData(&_IBZxOracle.CallOpts, sourceTokenAddress, destTokenAddress, sourceTokenAmount)
}

// GetTradeData is a free data retrieval call binding the contract method 0x06599aa0.
//
// Solidity: function getTradeData(address sourceTokenAddress, address destTokenAddress, uint256 sourceTokenAmount) view returns(uint256 sourceToDestRate, uint256 sourceToDestPrecision, uint256 destTokenAmount)
func (_IBZxOracle *IBZxOracleCallerSession) GetTradeData(sourceTokenAddress common.Address, destTokenAddress common.Address, sourceTokenAmount *big.Int) (struct {
	SourceToDestRate      *big.Int
	SourceToDestPrecision *big.Int
	DestTokenAmount       *big.Int
}, error) {
	return _IBZxOracle.Contract.GetTradeData(&_IBZxOracle.CallOpts, sourceTokenAddress, destTokenAddress, sourceTokenAmount)
}

// IWethHelperABI is the input ABI used to generate the binding from.
const IWethHelperABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"claimEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"claimAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// IWethHelperFuncSigs maps the 4-byte function signature to its string representation.
var IWethHelperFuncSigs = map[string]string{
	"bfcf63b0": "claimEther(address,uint256)",
}

// IWethHelper is an auto generated Go binding around an Ethereum contract.
type IWethHelper struct {
	IWethHelperCaller     // Read-only binding to the contract
	IWethHelperTransactor // Write-only binding to the contract
	IWethHelperFilterer   // Log filterer for contract events
}

// IWethHelperCaller is an auto generated read-only Go binding around an Ethereum contract.
type IWethHelperCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethHelperTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IWethHelperTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethHelperFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IWethHelperFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IWethHelperSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IWethHelperSession struct {
	Contract     *IWethHelper      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IWethHelperCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IWethHelperCallerSession struct {
	Contract *IWethHelperCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// IWethHelperTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IWethHelperTransactorSession struct {
	Contract     *IWethHelperTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// IWethHelperRaw is an auto generated low-level Go binding around an Ethereum contract.
type IWethHelperRaw struct {
	Contract *IWethHelper // Generic contract binding to access the raw methods on
}

// IWethHelperCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IWethHelperCallerRaw struct {
	Contract *IWethHelperCaller // Generic read-only contract binding to access the raw methods on
}

// IWethHelperTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IWethHelperTransactorRaw struct {
	Contract *IWethHelperTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIWethHelper creates a new instance of IWethHelper, bound to a specific deployed contract.
func NewIWethHelper(address common.Address, backend bind.ContractBackend) (*IWethHelper, error) {
	contract, err := bindIWethHelper(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IWethHelper{IWethHelperCaller: IWethHelperCaller{contract: contract}, IWethHelperTransactor: IWethHelperTransactor{contract: contract}, IWethHelperFilterer: IWethHelperFilterer{contract: contract}}, nil
}

// NewIWethHelperCaller creates a new read-only instance of IWethHelper, bound to a specific deployed contract.
func NewIWethHelperCaller(address common.Address, caller bind.ContractCaller) (*IWethHelperCaller, error) {
	contract, err := bindIWethHelper(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IWethHelperCaller{contract: contract}, nil
}

// NewIWethHelperTransactor creates a new write-only instance of IWethHelper, bound to a specific deployed contract.
func NewIWethHelperTransactor(address common.Address, transactor bind.ContractTransactor) (*IWethHelperTransactor, error) {
	contract, err := bindIWethHelper(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IWethHelperTransactor{contract: contract}, nil
}

// NewIWethHelperFilterer creates a new log filterer instance of IWethHelper, bound to a specific deployed contract.
func NewIWethHelperFilterer(address common.Address, filterer bind.ContractFilterer) (*IWethHelperFilterer, error) {
	contract, err := bindIWethHelper(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IWethHelperFilterer{contract: contract}, nil
}

// bindIWethHelper binds a generic wrapper to an already deployed contract.
func bindIWethHelper(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IWethHelperABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWethHelper *IWethHelperRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWethHelper.Contract.IWethHelperCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWethHelper *IWethHelperRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWethHelper.Contract.IWethHelperTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWethHelper *IWethHelperRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWethHelper.Contract.IWethHelperTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IWethHelper *IWethHelperCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _IWethHelper.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IWethHelper *IWethHelperTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IWethHelper.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IWethHelper *IWethHelperTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IWethHelper.Contract.contract.Transact(opts, method, params...)
}

// ClaimEther is a paid mutator transaction binding the contract method 0xbfcf63b0.
//
// Solidity: function claimEther(address receiver, uint256 amount) returns(uint256 claimAmount)
func (_IWethHelper *IWethHelperTransactor) ClaimEther(opts *bind.TransactOpts, receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWethHelper.contract.Transact(opts, "claimEther", receiver, amount)
}

// ClaimEther is a paid mutator transaction binding the contract method 0xbfcf63b0.
//
// Solidity: function claimEther(address receiver, uint256 amount) returns(uint256 claimAmount)
func (_IWethHelper *IWethHelperSession) ClaimEther(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWethHelper.Contract.ClaimEther(&_IWethHelper.TransactOpts, receiver, amount)
}

// ClaimEther is a paid mutator transaction binding the contract method 0xbfcf63b0.
//
// Solidity: function claimEther(address receiver, uint256 amount) returns(uint256 claimAmount)
func (_IWethHelper *IWethHelperTransactorSession) ClaimEther(receiver common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IWethHelper.Contract.ClaimEther(&_IWethHelper.TransactOpts, receiver, amount)
}

// LoanTokenLogicV4ABI is the input ABI used to generate the binding from.
const LoanTokenLogicV4ABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenToFillAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawOnOpen\",\"type\":\"bool\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"fallback\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetBorrow\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"assetSupply\",\"type\":\"uint256\"}],\"name\":\"_supplyInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"assetBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"avgBorrowInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"borrowInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLoanDuration\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralTokenSent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"\",\"type\":\"bytes\"}],\"name\":\"borrowTokenFromDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanAmountPaid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"burnAmount\",\"type\":\"uint256\"}],\"name\":\"burnToEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"loanAmountPaid\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burntTokenReserveList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burntTokenReserveListIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burntTokenReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"checkpointPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBZxObjects.LoanOrder\",\"name\":\"loanOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddressFilled\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"positionTokenAddressFilled\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmountUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanStartUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanEndUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"}],\"internalType\":\"structBZxObjects.LoanPosition\",\"name\":\"loanPosition\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"loanCloser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"closeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"closeLoanNotifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLoanDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"name\":\"getBorrowAmountForDeposit\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialLoanDuration\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"name\":\"getDepositAmountForBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLeverageList\",\"outputs\":[{\"internalType\":\"uint256[]\",\"name\":\"\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"}],\"name\":\"getLoanData\",\"outputs\":[{\"components\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginPremiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"internalType\":\"structLoanTokenStorage.LoanData\",\"name\":\"\",\"type\":\"tuple\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"}],\"name\":\"getMaxEscrowAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leverageList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loanOrderData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginPremiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loanOrderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralTokenSent\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"tradeTokenSent\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"depositTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tradeTokenAddress\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"loanDataBytes\",\"type\":\"bytes\"}],\"name\":\"marginTradeFromDeposit\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"marketLiquidity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"depositAmount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiver\",\"type\":\"address\"}],\"name\":\"mintWithEther\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"mintAmount\",\"type\":\"uint256\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"}],\"name\":\"nextBorrowInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"useFixedInterestModel\",\"type\":\"bool\"}],\"name\":\"nextBorrowInterestRateWithOption\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"supplyAmount\",\"type\":\"uint256\"}],\"name\":\"nextSupplyInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"protocolInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spreadMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"supplyInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenizedRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAssetBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAssetSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"assetSupply\",\"type\":\"uint256\"}],\"name\":\"totalSupplyInterestRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"settingsTarget\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"callData\",\"type\":\"bytes\"}],\"name\":\"updateSettings\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LoanTokenLogicV4FuncSigs maps the 4-byte function signature to its string representation.
var LoanTokenLogicV4FuncSigs = map[string]string{
	"7288b344": "_supplyInterestRate(uint256,uint256)",
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"06b3efd6": "assetBalanceOf(address)",
	"44a4a003": "avgBorrowInterestRate()",
	"995363d3": "bZxContract()",
	"96c7871b": "bZxOracle()",
	"894ca308": "bZxVault()",
	"70a08231": "balanceOf(address)",
	"1f68f20a": "baseRate()",
	"8325a1c0": "borrowInterestRate()",
	"cfb65bb9": "borrowTokenFromDeposit(uint256,uint256,uint256,uint256,address,address,address,bytes)",
	"9dc29fac": "burn(address,uint256)",
	"81a6b250": "burnToEther(address,uint256)",
	"7866c6c1": "burntTokenReserveList(uint256)",
	"fbd9574d": "burntTokenReserveListIndex(address)",
	"0c4925fd": "burntTokenReserved()",
	"eebc5081": "checkpointPrice(address)",
	"7b7933b4": "checkpointSupply()",
	"cd4fa66d": "closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32),(address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256),address,uint256,bool)",
	"313ce567": "decimals()",
	"24d25f4a": "getBorrowAmountForDeposit(uint256,uint256,uint256,address)",
	"8423acd6": "getDepositAmountForBorrow(uint256,uint256,uint256,address)",
	"2ecae90a": "getLeverageList()",
	"c4d2b1b3": "getLoanData(bytes32)",
	"829b38f4": "getMaxEscrowAmount(uint256)",
	"1d0806ae": "initialPrice()",
	"9b3a54d1": "leverageList(uint256)",
	"2515aacd": "loanOrderData(bytes32)",
	"fe056342": "loanOrderHashes(uint256)",
	"797bf385": "loanTokenAddress()",
	"1c5d1da5": "marginTradeFromDeposit(uint256,uint256,uint256,uint256,uint256,address,address,address,address,bytes)",
	"612ef80b": "marketLiquidity()",
	"40c10f19": "mint(address,uint256)",
	"8f6ede1f": "mintWithEther(address)",
	"06fdde03": "name()",
	"b9fe1a8f": "nextBorrowInterestRate(uint256)",
	"7d90dcba": "nextBorrowInterestRateWithOption(uint256,bool)",
	"d65a5021": "nextSupplyInterestRate(uint256)",
	"8da5cb5b": "owner()",
	"fc3b72b1": "protocolInterestRate()",
	"330691ac": "rateMultiplier()",
	"d84d2a47": "spreadMultiplier()",
	"09ec6b6b": "supplyInterestRate()",
	"95d89b41": "symbol()",
	"7ff9b596": "tokenPrice()",
	"736ee3d3": "tokenizedRegistry()",
	"20f6d07c": "totalAssetBorrow()",
	"8fb807c5": "totalAssetSupply()",
	"18160ddd": "totalSupply()",
	"12416898": "totalSupplyInterestRate(uint256)",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"f2fde38b": "transferOwnership(address)",
	"284e2f56": "updateSettings(address,bytes)",
	"4780eac1": "wethContract()",
}

// LoanTokenLogicV4Bin is the compiled bytecode used for deploying new contracts.
var LoanTokenLogicV4Bin = "0x608060405260016000819055600a805460ff19169055670de0b6b3a7640000600b556801043561a882930000600c5580546001600160a01b03191633179055614d748061004d6000396000f3fe60806040526004361061036b5760003560e01c80637d90dcba116101c65780639dc29fac116100f7578063d84d2a4711610095578063f2fde38b1161006f578063f2fde38b14610948578063fbd9574d14610968578063fc3b72b114610996578063fe056342146109ab5761036b565b8063d84d2a47146108f3578063dd62ed3e14610908578063eebc5081146109285761036b565b8063c4d2b1b3116100d1578063c4d2b1b314610873578063cd4fa66d146108a0578063cfb65bb9146108c0578063d65a5021146108d35761036b565b80639dc29fac14610813578063a9059cbb14610833578063b9fe1a8f146108535761036b565b80638da5cb5b1161016457806395d89b411161013e57806395d89b41146107b457806396c7871b146107c9578063995363d3146107de5780639b3a54d1146107f35761036b565b80638da5cb5b146107775780638f6ede1f1461078c5780638fb807c51461079f5761036b565b8063829b38f4116101a0578063829b38f41461070d5780638325a1c01461072d5780638423acd614610742578063894ca308146107625761036b565b80637d90dcba146106b85780637ff9b596146106d857806381a6b250146106ed5761036b565b8063284e2f56116102a0578063612ef80b1161023e578063736ee3d311610218578063736ee3d31461064b5780637866c6c114610660578063797bf3851461068e5780637b7933b4146106a35761036b565b8063612ef80b146105f657806370a082311461060b5780637288b3441461062b5761036b565b8063330691ac1161027a578063330691ac1461058a57806340c10f191461059f57806344a4a003146105bf5780634780eac1146105d45761036b565b8063284e2f56146105245780632ecae90a14610546578063313ce567146105685761036b565b80631c5d1da51161030d57806320f6d07c116102e757806320f6d07c1461049b57806323b872dd146104b057806324d25f4a146104d05780632515aacd146104f05761036b565b80631c5d1da51461045e5780631d0806ae146104715780631f68f20a146104865761036b565b806309ec6b6b1161034957806309ec6b6b146103ff5780630c4925fd14610414578063124168981461042957806318160ddd146104495761036b565b806306b3efd61461037a57806306fdde03146103b0578063095ea7b3146103d2575b34801561037757600080fd5b50005b34801561038657600080fd5b5061039a610395366004613cf7565b6109cb565b6040516103a791906148ee565b60405180910390f35b3480156103bc57600080fd5b506103c5610a0d565b6040516103a791906149af565b3480156103de57600080fd5b506103f26103ed366004613e02565b610a98565b6040516103a791906148e0565b34801561040b57600080fd5b5061039a610b03565b34801561042057600080fd5b5061039a610b16565b34801561043557600080fd5b5061039a610444366004613e32565b610b1c565b34801561045557600080fd5b5061039a610b40565b61039a61046c3660046140f8565b610b46565b34801561047d57600080fd5b5061039a610d3a565b34801561049257600080fd5b5061039a610d40565b3480156104a757600080fd5b5061039a610d46565b3480156104bc57600080fd5b506103f26104cb366004613d6d565b610d4c565b3480156104dc57600080fd5b5061039a6104eb366004613f7a565b610f77565b3480156104fc57600080fd5b5061051061050b366004613e32565b610fbc565b6040516103a7989796959493929190614938565b34801561053057600080fd5b5061054461053f366004613dba565b611008565b005b34801561055257600080fd5b5061055b611140565b6040516103a791906148cf565b34801561057457600080fd5b5061057d611198565b6040516103a79190614be7565b34801561059657600080fd5b5061039a6111a1565b3480156105ab57600080fd5b5061039a6105ba366004613e02565b6111a7565b3480156105cb57600080fd5b5061039a6111e8565b3480156105e057600080fd5b506105e9611222565b6040516103a791906147c7565b34801561060257600080fd5b5061039a611231565b34801561061757600080fd5b5061039a610626366004613cf7565b61125c565b34801561063757600080fd5b5061039a610646366004613f18565b611277565b34801561065757600080fd5b506105e96112c5565b34801561066c57600080fd5b5061068061067b366004613e32565b6112d9565b6040516103a792919061488c565b34801561069a57600080fd5b506105e961130e565b3480156106af57600080fd5b5061039a61131d565b3480156106c457600080fd5b5061039a6106d3366004613ee8565b611323565b3480156106e457600080fd5b5061039a61132f565b3480156106f957600080fd5b5061039a610708366004613e02565b61135e565b34801561071957600080fd5b5061039a610728366004613e32565b6114bc565b34801561073957600080fd5b5061039a61157f565b34801561074e57600080fd5b5061039a61075d366004613f7a565b61158c565b34801561076e57600080fd5b506105e96117f3565b34801561078357600080fd5b506105e9611802565b61039a61079a366004613cf7565b611811565b3480156107ab57600080fd5b5061039a611874565b3480156107c057600080fd5b506103c5611895565b3480156107d557600080fd5b506105e96118f0565b3480156107ea57600080fd5b506105e96118ff565b3480156107ff57600080fd5b5061039a61080e366004613e32565b611913565b34801561081f57600080fd5b5061039a61082e366004613e02565b611931565b34801561083f57600080fd5b506103f261084e366004613e02565b61199d565b34801561085f57600080fd5b5061039a61086e366004613e32565b611b31565b34801561087f57600080fd5b5061089361088e366004613e32565b611b3e565b6040516103a79190614b60565b3480156108ac57600080fd5b506103f26108bb366004613e50565b611bbd565b61039a6108ce366004614030565b611db5565b3480156108df57600080fd5b5061039a6108ee366004613e32565b61207e565b3480156108ff57600080fd5b5061039a61208f565b34801561091457600080fd5b5061039a610923366004613d33565b612095565b34801561093457600080fd5b5061039a610943366004613cf7565b6120c0565b34801561095457600080fd5b50610544610963366004613cf7565b6120db565b34801561097457600080fd5b50610988610983366004613cf7565b6120fe565b6040516103a7929190614b6f565b3480156109a257600080fd5b5061039a61211a565b3480156109b757600080fd5b5061039a6109c6366004613e32565b612127565b6000610a05670de0b6b3a76400006109f96109e461132f565b6109ed8661125c565b9063ffffffff61213916565b9063ffffffff61215e16565b90505b919050565b6002805460408051602060018416156101000260001901909316849004601f81018490048402820184019092528181529291830182828015610a905780601f10610a6557610100808354040283529160200191610a90565b820191906000526020600020905b815481529060010190602001808311610a7357829003601f168201915b505050505081565b336000818152601a602090815260408083206001600160a01b038716808552925280832085905551919290917f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92590610af19086906148ee565b60405180910390a35060015b92915050565b6000610b10610444611874565b90505b90565b60135481565b6015546000908015610b3a57610b328184611277565b915050610a08565b50919050565b601b5490565b6001546000906001600160a01b03163314610b6057600080fd5b6001600160a01b03831615801590610b8657506008546001600160a01b03848116911614155b610bab5760405162461bcd60e51b8152600401610ba2906149e0565b60405180910390fd5b8a6001600160a01b038681169085161415610c52576006546008546040516232ccd560e51b81526001600160a01b03928316926306599aa092610bf8928992909116908690600401614864565b60606040518083038186803b158015610c1057600080fd5b505afa158015610c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250610c489190810190613f37565b9250610c7f915050565b6008546001600160a01b03878116911614610c7f5760405162461bcd60e51b8152600401610ba290614a70565b610d2a8b60405180608001604052808a6001600160a01b03166001600160a01b03168152602001886001600160a01b03166001600160a01b03168152602001876001600160a01b03166001600160a01b031681526020018a6001600160a01b03166001600160a01b03168152506040518060e0016040528060008152602001858152602001600081526020018e81526020018d81526020018c81526020016000815250600187612171565b9c9b505050505050505050505050565b60185481565b600b5481565b60155481565b6001600160a01b0383166000818152601a6020908152604080832033845282528083205493835260199091528120549091908311801590610d8d5750808311155b8015610da157506001600160a01b03841615155b610dbd5760405162461bcd60e51b8152600401610ba290614a40565b6001600160a01b038516600090815260196020526040902054610de6908463ffffffff6122eb16565b6001600160a01b038087166000908152601960205260408082209390935590861681522054610e1b908463ffffffff6122fd16565b6001600160a01b038516600090815260196020526040902055600019811015610e7357610e4e818463ffffffff6122eb16565b6001600160a01b0386166000908152601a602090815260408083203384529091529020555b6000610e7d61132f565b6001600160a01b03871660009081526019602052604090205490915015610ebe576001600160a01b0386166000908152600960205260409020819055610ed8565b6001600160a01b0386166000908152600960205260408120555b6001600160a01b03851660009081526019602052604090205415610f16576001600160a01b0385166000908152600960205260409020819055610f30565b6001600160a01b0385166000908152600960205260408120555b846001600160a01b0316866001600160a01b0316600080516020614d1283398151915286604051610f6191906148ee565b60405180910390a36001925050505b9392505050565b60008382604051602001610f8c9291906147a1565b6040516020818303038152906040528051906020012060001c9350610fb38585858561230a565b95945050505050565b600f60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007909701549596949593949293919290916001600160a01b031688565b6001546001600160a01b03163314611093577f7ad06df6a0af6bd602d90db766e0d5f253b45187c3717a0f9026ea8b10ff0d4b547f34b31cff1dbd8374124bd4505521fc29cab0f9554a5386ba7d784a4e611c7e3154336001600160a01b0383161480156110875750806001600160a01b0316846001600160a01b0316145b61109057600080fd5b50505b601c80546001600160a01b038481166001600160a01b031983161790925560405191169060009030906110c7908590614795565b6000604051808303816000865af19150503d8060008114611104576040519150601f19603f3d011682016040523d82523d6000602084013e611109565b606091505b50506040519091503d90816000823e82611121578181fd5b601c80546001600160a01b0319166001600160a01b0386161790558181f35b6060601080548060200260200160405190810160405280929190818152602001828054801561118e57602002820191906000526020600020905b81548152602001906001019080831161117a575b5050505050905090565b60045460ff1681565b600c5481565b60006001600054146111cb5760405162461bcd60e51b8152600401610ba290614b10565b60026000556111da838361252d565b90505b600160005592915050565b60155460009080156112165761120e6111ff611874565b6109f96016546109ed85612659565b915050610b13565b61120e612691565b5090565b6007546001600160a01b031681565b60008061123c611874565b905060155481111561121e5760155461120e90829063ffffffff6122eb16565b6001600160a01b031660009081526019602052604090205490565b600082158015906112885750828210155b15610afd576112be701d6329f1c35ca4bfabb9f56100000000006109f9600d546109ed6112b588886126b6565b6109ed89612659565b9050610afd565b600a5461010090046001600160a01b031681565b601181815481106112e657fe5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b6008546001600160a01b031681565b60165481565b6000610f7083836126e8565b600080426017541461134757611343612760565b9150505b61135861135382612827565b612884565b91505090565b60006001600054146113825760405162461bcd60e51b8152600401610ba290614b10565b60026000556007546008546001600160a01b039081169116146113b75760405162461bcd60e51b8152600401610ba290614a10565b6113c0826128b3565b905080156111dd576008546040805180820190915260018152600d60fa1b6020820152733b5bdccdfa2a0a1911984f203c19628eeb6036e091611412916001600160a01b039091169083908590612a21565b604051630bfcf63b60e41b81526001600160a01b0382169063bfcf63b090611440908790869060040161488c565b602060405180830381600087803b15801561145a57600080fd5b505af115801561146e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506114929190810190613eca565b82146114b05760405162461bcd60e51b8152600401610ba2906149c0565b50600160005592915050565b60006114c6613a0d565b506000828152600e60209081526040808320548352600f8252918290208251610100810184528154815260018201549281019290925260028101549282018390526003810154606083015260048101546080830152600581015460a0830152600681015460c0830152600701546001600160a01b031660e082015290611550576000915050610a08565b610f7061156f68056bc75e2d6310000083608001518460400151612aec565b6109f983604001516109ed611231565b6000610b106000806126e8565b600084156117eb5783826040516020016115a79291906147a1565b6040516020818303038152906040528051906020012060001c93506115ca613a0d565b506000848152600e60209081526040808320548352600f82528083208151610100810183528154815260018201549381019390935260028101549183018290526003810154606084015260048101546080840152600581015460a0840152600681015460c0840152600701546001600160a01b031660e08301529091906116609068056bc75e2d6310000063ffffffff6122fd16565b905061168969021e19e0c9bab24000006109f961167c88612b2d565b8a9063ffffffff61213916565b6008546040516370a0823160e01b81529198506001600160a01b0316906370a08231906116ba9030906004016147c7565b60206040518083038186803b1580156116d257600080fd5b505afa1580156116e6573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061170a9190810190613eca565b87116117e8576004546008546117df91600a916001600160a01b0361010090920482169163bc6cb1d991811690891661174e576007546001600160a01b0316611750565b885b6006546040516001600160e01b031960e086901b1681526117839392916001600160a01b0316908f908a90600401614818565b60206040518083038186803b15801561179b57600080fd5b505afa1580156117af573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506117d39190810190613eca565b9063ffffffff6122fd16565b925050506117eb565b50505b949350505050565b6005546001600160a01b031681565b6001546001600160a01b031681565b60006001600054146118355760405162461bcd60e51b8152600401610ba290614b10565b60026000556007546008546001600160a01b0390811691161461186a5760405162461bcd60e51b8152600401610ba290614ae0565b6111dd823461252d565b600080426017541461188c57611888612760565b9150505b61135881612827565b6003805460408051602060026001851615610100026000190190941693909304601f81018490048402820184019092528181529291830182828015610a905780601f10610a6557610100808354040283529160200191610a90565b6006546001600160a01b031681565b60045461010090046001600160a01b031681565b6010818154811061192057fe5b600091825260209091200154905081565b60006001600054146119555760405162461bcd60e51b8152600401610ba290614b10565b6002600055611963826128b3565b905080156111dd576008546040805180820190915260018152603560f81b60208201526111dd916001600160a01b03169085908490612a21565b3360009081526019602052604081205482118015906119c457506001600160a01b03831615155b6119e05760405162461bcd60e51b8152600401610ba290614a80565b33600090815260196020526040902054611a00908363ffffffff6122eb16565b33600090815260196020526040808220929092556001600160a01b03851681522054611a32908363ffffffff6122fd16565b6001600160a01b038416600090815260196020526040812091909155611a5661132f565b3360009081526019602052604090205490915015611a8557336000908152600960205260409020819055611a96565b336000908152600960205260408120555b6001600160a01b03841660009081526019602052604090205415611ad4576001600160a01b0384166000908152600960205260409020819055611aee565b6001600160a01b0384166000908152600960205260408120555b836001600160a01b0316336001600160a01b0316600080516020614d1283398151915285604051611b1f91906148ee565b60405180910390a35060019392505050565b6000610a058260006126e8565b611b46613a0d565b506000908152600f6020908152604091829020825161010081018452815481526001820154928101929092526002810154928201929092526003820154606082015260048201546080820152600582015460a0820152600682015460c08201526007909101546001600160a01b031660e082015290565b600480546006546040516338f5892f60e11b81526000936001600160a01b036101009094048416936371eb125e93611bf893911691016147c7565b60206040518083038186803b158015611c1057600080fd5b505afa158015611c24573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250611c489190810190613d15565b6001600160a01b0316336001600160a01b031614611c785760405162461bcd60e51b8152600401610ba290614af0565b611c80612b79565b611c88613a0d565b50610120860180516000908152600f602090815260409182902082516101008101845281548082526001830154938201939093526002820154938101939093526003810154606084015260048101546080840152600581015460a0840152600681015460c0840152600701546001600160a01b031660e0830152915190911415611da8578360155411611d1c576000611d2f565b601554611d2f908563ffffffff6122eb16565b60155585516101208801516040516001600160a01b03909216917f85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b7390611d7a908990899089906148a7565b60405180910390a383611d91576001915050610fb3565b611d9b6000612827565b6016555060019050610fb3565b5060009695505050505050565b6001546000906001600160a01b03163314611dcf57600080fd5b34158015611de557506001600160a01b03831615155b8015611df057508515155b80611e2d57503415801590611e2357506001600160a01b0383161580611e2357506007546001600160a01b038481169116145b8015611e2d575085155b611e495760405162461bcd60e51b8152600401610ba290614b20565b3415611e61576007543496506001600160a01b031692505b6040518990611e76908a9086906020016147a1565b60408051601f1981840301815291815281516020928301206000818152600e909352912054909950915081611ebd5760405162461bcd60e51b8152600401610ba290614a20565b611ec5612b79565b611ecd613a5e565b611ed5613a0d565b506000838152600f60209081526040918290208251610100810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015460808201819052600583015460a0830152600683015460c08301526007909201546001600160a01b031660e0820152901583611f8957611f5e8a8d8d8a61230a565b935083611f7d5760405162461bcd60e51b8152600401610ba290614a50565b60c08301849052611f91565b60c083018490525b611fa684611f9f6000612827565b8d84612c19565b60408681019290925291855280516080810182526001600160a01b038c811682528a8116602080840191909152600083850152908c166060830152825160e08101845287518152908101849052929650612049928892810187600260200201518152602001600081526020018e8152602001600081526020018760066007811061202c57fe5b602002015181525060405180602001604052806000815250612c6a565b60c08401819052841461206e5760405162461bcd60e51b8152600401610ba290614b30565b5050505098975050505050505050565b6000610a05610444836117d3611874565b600d5481565b6001600160a01b039182166000908152601a6020908152604080832093909416825291909152205490565b6001600160a01b031660009081526009602052604090205490565b6001546001600160a01b031633146120f257600080fd5b6120fb81612ec4565b50565b6012602052600090815260409020805460019091015460ff1682565b6000610b10601554612659565b600e6020526000908152604090205481565b60008261214857506000610afd565b508181028183828161215657fe5b0414610afd57fe5b600081838161216957fe5b049392505050565b60208301516000906121955760405162461bcd60e51b8152600401610ba290614b50565b506000858152600e6020526040902054806121c25760405162461bcd60e51b8152600401610ba290614b00565b6121ca612b79565b6121d2613a0d565b506000818152600f60209081526040918290208251610100810184528154815260018201549281019290925260028101549282019290925260038201546060820152600482015460808201819052600583015460a0830152600683015460c08301526007909201546001600160a01b031660e0820152901584156122755761226283876001602002015183612f33565b87526020870181905260c0870152612291565b602086015161228e906122886000612827565b8361304c565b86525b60408701516001600160a01b03166122ab57600060a08701525b60006122b984898988612c6a565b602088015190915081146122df5760405162461bcd60e51b8152600401610ba2906149d0565b50505095945050505050565b6000828211156122f757fe5b50900390565b81810182811015610afd57fe5b600084156117eb5761231a613a0d565b506000848152600e60209081526040808320548352600f82528083208151610100810183528154815260018201549381019390935260028101549183018290526003810154606084015260048101546080840152600581015460a0840152600681015460c0840152600701546001600160a01b031660e08301529091906123b09068056bc75e2d6310000063ffffffff6122fd16565b6004546008549192506001600160a01b0361010090910481169163f3d75a9c919081169087166123eb576007546001600160a01b03166123ed565b865b6006546040516001600160e01b031960e086901b1681526124209392916001600160a01b0316908d908890600401614818565b60206040518083038186803b15801561243857600080fd5b505afa15801561244c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506124709190810190613eca565b925061249861247e86612b2d565b6109f98569021e19e0c9bab240000063ffffffff61213916565b6008546040516370a0823160e01b81529194506001600160a01b0316906370a08231906124c99030906004016147c7565b60206040518083038186803b1580156124e157600080fd5b505afa1580156124f5573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506125199190810190613eca565b8311156117e8575060009695505050505050565b60008161254c5760405162461bcd60e51b8152600401610ba290614aa0565b612554612b79565b60006125636113536000612827565b9050612581816109f985670de0b6b3a764000063ffffffff61213916565b9150346125c257600854604080518082019091526002815261062760f31b60208201526125bd916001600160a01b031690339030908790613224565b61262c565b600760009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0846040518263ffffffff1660e01b81526004016000604051808303818588803b15801561261257600080fd5b505af1158015612626573d6000803e3d6000fd5b50505050505b612638848385846132f2565b6001600160a01b039093166000908152600960205260409020929092555090565b60008115610a0857600061266b612760565b509050610b3261016d6109ed856109f98568056bc75e2d6310000063ffffffff61213916565b7f3d82e958c891799f357c1316ae5543412952ae5c423336f8929ed7458039c9955490565b600082158015906126c657508115155b15610afd576112be826109f98568056bc75e2d6310000063ffffffff61213916565b600080831561274d57426017541461270657612702612760565b9150505b6008546040516370a0823160e01b815260009161273d9184916001600160a01b0316906370a08231906117839030906004016147c7565b90508085111561274b578094505b505b6117eb8461275a83612827565b8561304c565b6004805460065460085460405163015216af60e31b8152600094859461010090046001600160a01b0390811694630a90b578946127a79430949284169390911691016147f0565b60806040518083038186803b1580156127bf57600080fd5b505afa1580156127d3573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506127f79190810190613fdb565b600d549195509350612821925068056bc75e2d6310000091506109f990849063ffffffff61213916565b90509091565b6000601b54600014610a085760135480612874576015546008546040516370a0823160e01b815261287192916001600160a01b0316906370a08231906117839030906004016147c7565b90505b610b32818463ffffffff6122fd16565b601b546000908061289757601854610f70565b610f70816109f985670de0b6b3a764000063ffffffff61213916565b6000816128d25760405162461bcd60e51b8152600401610ba290614ac0565b6128db3361125c565b8211156128ee576128eb3361125c565b91505b6128f6612b79565b60006129056113536000612827565b90506000612925670de0b6b3a76400006109f9868563ffffffff61213916565b6008546040516370a0823160e01b81529192506000916001600160a01b03909116906370a082319061295b9030906004016147c7565b60206040518083038186803b15801561297357600080fd5b505afa158015612987573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052506129ab9190810190613eca565b9050819350808411156129d05760405162461bcd60e51b8152600401610ba290614a30565b6129dc338686866133f0565b3360009081526019602052604090205415612a0857336000908152600960205260409020839055612a19565b336000908152600960205260408120555b505050919050565b6000846001600160a01b031663a9059cbb8585604051602401612a4592919061488c565b6040516020818303038152906040529060e01b6020820180516001600160e01b038381831617835250505050604051612a7e9190614795565b6000604051808303816000865af19150503d8060008114612abb576040519150601f19603f3d011682016040523d82523d6000602084013e612ac0565b606091505b50509050808290612ae45760405162461bcd60e51b8152600401610ba291906149af565b505050505050565b600082612b025768056bc75e2d631000006117eb565b6117eb68056bc75e2d631000006117d3846109f9876109ed6301e13380838c8863ffffffff61213916565b6000610a0569021e19e0c9bab24000006117d36204cfe06109f9866109ed600b546117d368056bc75e2d631000006109f96804563918244f400000600c5461213990919063ffffffff16565b4260175414612c17576004805460065460085460405163327ab63960e01b81526001600160a01b0361010090940484169463327ab63994612bbf948116931691016147d5565b602060405180830381600087803b158015612bd957600080fd5b505af1158015612bed573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612c119190810190613eca565b50426017555b565b6000806000612c2a87878787613548565b9250612c499050612c41888463ffffffff6122fd16565b878787613548565b9093509150612c5e878363ffffffff6122fd16565b90509450945094915050565b6000612c74613587565b6008546040516370a0823160e01b81526001600160a01b03909116906370a0823190612ca49030906004016147c7565b60206040518083038186803b158015612cbc57600080fd5b505afa158015612cd0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250612cf49190810190613eca565b602084015111801590612d10575083516001600160a01b031615155b612d2c5760405162461bcd60e51b8152600401610ba290614a60565b60608401516001600160a01b0316612d4f5783516001600160a01b031660608501525b612d598484613607565b60208301516060840151612d6c916122fd565b606084015260003415612d8657504734811115612d865750345b6004805460405163b1eac3ad60e01b81526101009091046001600160a01b03169163b1eac3ad918491612dc1918b918b918b918b91016148fc565b6020604051808303818588803b158015612dda57600080fd5b505af1158015612dee573d6000803e3d6000fd5b50505050506040513d601f19601f82011682018060405250612e139190810190613eca565b60208501819052612e365760405162461bcd60e51b8152600401610ba290614a90565b6020840151601554612e47916122fd565b601555612e546000612827565b60165584516020858101518651918801516040808a015190516001600160a01b03958616957f86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e95612eaf959490939092909182161590614b8a565b60405180910390a25050506020015192915050565b6001600160a01b038116612ed757600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b600080612f3e613a0d565b506000858152600f60209081526040918290208251610100810184528154815260018201549281019290925260028101549282018390526003810154606083015260048101546080830152600581015460a0830152600681015460c0830152600701546001600160a01b031660e082015290612fcc5760405162461bcd60e51b8152600401610ba290614b40565b604081015161300090612ff2906109f98868056bc75e2d6310000063ffffffff61213916565b612ffa611874565b8661304c565b915061304181604001516109f96130208585608001518660400151612aec565b6109f989701d6329f1c35ca4bfabb9f561000000000063ffffffff61213916565b925050935093915050565b60008061306d613067866015546122fd90919063ffffffff16565b856126b6565b905060008060008086156130e4576804563918244f400000851015613099576804563918244f40000094505b50507f185a40c6b6d3f849f72c71ea950323d21149c27a9d90f7dc5e5ea2d332edcf7f547f9ff54bc0049f5eab56ca7cd14591be3f7ed6355b856d01e3770305c74a004ea254613131565b6802b5e3af16b1880000851015613128576130fd612691565b91507f2b4858b1bc9e2d14afab03340ce5f6c81b703c86a0c570653ae586534e095fb1549050613131565b5050600b54600c545b6804e1003b28d92800008511156131c25761315b856804e1003b28d928000063ffffffff6122eb16565b9450678ac7230489e8000085111561317957678ac7230489e8000094505b61319360646109f9605a6109ed858763ffffffff6122fd16565b92506131bb836117d3678ac7230489e800006109f961167c68056bc75e2d63100000896122eb565b9550613218565b6131e3826117d368056bc75e2d631000006109f9898663ffffffff61213916565b955090925082906131fa818363ffffffff6122fd16565b92508386101561320c57839550613218565b82861115613218578295505b50505050509392505050565b6000856001600160a01b03166323b872dd86868660405160240161324a93929190614864565b6040516020818303038152906040529060e01b6020820180516001600160e01b0383818316178352505050506040516132839190614795565b6000604051808303816000865af19150503d80600081146132c0576040519150601f19603f3d011682016040523d82523d6000602084013e6132c5565b606091505b505090508082906132e95760405162461bcd60e51b8152600401610ba291906149af565b50505050505050565b6001600160a01b0384166133185760405162461bcd60e51b8152600401610ba2906149f0565b601b5461332b908463ffffffff6122fd16565b601b556001600160a01b038416600090815260196020526040902054613357908463ffffffff6122fd16565b6001600160a01b038516600081815260196020526040908190209290925590517fb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb906133a890869086908690614bcc565b60405180910390a2836001600160a01b031660006001600160a01b0316600080516020614d12833981519152856040516133e291906148ee565b60405180910390a350505050565b6001600160a01b0384166000908152601960205260409020548311156134285760405162461bcd60e51b8152600401610ba290614a00565b6001600160a01b038416600090815260196020526040902054613451908463ffffffff6122eb16565b6001600160a01b0385166000908152601960205260409020819055600a106134b9576001600160a01b03841660009081526019602052604090205461349d90849063ffffffff6122fd16565b6001600160a01b03851660009081526019602052604081205592505b601b546134cc908463ffffffff6122eb16565b601b556040516001600160a01b038516907f743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b46449061350e90869086908690614bcc565b60405180910390a260006001600160a01b0316846001600160a01b0316600080516020614d12833981519152856040516133e291906148ee565b60008061355686868561304c565b915061357c6b0a3098c68eb9427db80000006109f9866109ed8a8763ffffffff61213916565b905094509492505050565b600080356001600160e01b0319167fd46a704bc285dbd6ff5ad3863506260b1df02812f4f857c8cc852317a6ac64f26040516020016135c792919061476f565b60405160208183030381529060405280519060200120905060008154905080156136035760405162461bcd60e51b8152600401610ba290614ab0565b5050565b60208083015160408401516060808601519385015190850151608086015160a087015160c08801519596949560006001600160a01b0388166137b7576007546008546001600160a01b039081169116141561372757600854604080516020810190915260008152733b5bdccdfa2a0a1911984f203c19628eeb6036e09161369d916001600160a01b039091169083908690612a21565b604051630bfcf63b60e41b81526001600160a01b0382169063bfcf63b0906136cb908b90879060040161488c565b602060405180830381600087803b1580156136e557600080fd5b505af11580156136f9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525061371d9190810190613eca565b8314915050613754565b600854604080516020810190915260008152613750916001600160a01b03169089908590612a21565b5060015b80801561376057508186115b1561379557600854600554604080516020810190915260008152613795926001600160a01b03908116921690858a0390612a21565b806137b25760405162461bcd60e51b8152600401610ba290614ad0565b6137ef565b600854600554604080518082019091526002815261191b60f11b60208201526137ef926001600160a01b039081169216908990612a21565b8315613957576007546001600160a01b038a8116911614801561381157503415155b801561381c57503484145b156138c457600760009054906101000a90046001600160a01b03166001600160a01b031663d0e30db0856040518263ffffffff1660e01b81526004016000604051808303818588803b15801561387157600080fd5b505af1158015613885573d6000803e3d6000fd5b5050600554604080518082019091526002815261323760f01b60208201526138bf94508d93506001600160a01b0390911691508790612a21565b613957565b6008546001600160a01b038a8116911614156138f1576138ea858563ffffffff6122fd16565b9450613957565b876001600160a01b0316896001600160a01b031614156139225761391b838563ffffffff6122fd16565b9250613957565b600554604080518082019091526002815261323760f01b6020820152613957918b9133916001600160a01b0316908890613224565b84156139c5576008546001600160a01b038981169116141561398a57613983838663ffffffff6122fd16565b92506139c5565b600854600554604080518082019091526002815261333160f01b60208201526139c5926001600160a01b039081169233929116908990613224565b8215613a0057600554604080518082019091526002815261199960f11b6020820152613a00918a9133916001600160a01b0316908790613224565b5050505050505050505050565b6040518061010001604052806000801916815260200160008152602001600081526020016000815260200160008152602001600081526020016000815260200160006001600160a01b031681525090565b6040518060e001604052806007906020820280388339509192915050565b8035610afd81614ceb565b8051610afd81614ceb565b8035610afd81614cff565b8035610afd81614d08565b600082601f830112613ab957600080fd5b8135613acc613ac782614c1c565b614bf5565b91508082526020830160208301858383011115613ae857600080fd5b613af3838284614c92565b50505092915050565b60006101408284031215613b0f57600080fd5b613b1a610140614bf5565b90506000613b288484613a7c565b8252506020613b3984848301613a7c565b6020830152506040613b4d84828501613a7c565b6040830152506060613b6184828501613a7c565b6060830152506080613b7584828501613a9d565b60808301525060a0613b8984828501613a9d565b60a08301525060c0613b9d84828501613a9d565b60c08301525060e0613bb184828501613a9d565b60e083015250610100613bc684828501613a9d565b61010083015250610120613bdc84828501613a9d565b6101208301525092915050565b60006101608284031215613bfc57600080fd5b613c07610160614bf5565b90506000613c158484613a7c565b8252506020613c2684848301613a7c565b6020830152506040613c3a84828501613a7c565b6040830152506060613c4e84828501613a9d565b6060830152506080613c6284828501613a9d565b60808301525060a0613c7684828501613a9d565b60a08301525060c0613c8a84828501613a9d565b60c08301525060e0613c9e84828501613a9d565b60e083015250610100613cb384828501613a9d565b61010083015250610120613cc984828501613a92565b61012083015250610140613cdf84828501613a9d565b6101408301525092915050565b8051610afd81614d08565b600060208284031215613d0957600080fd5b60006117eb8484613a7c565b600060208284031215613d2757600080fd5b60006117eb8484613a87565b60008060408385031215613d4657600080fd5b6000613d528585613a7c565b9250506020613d6385828601613a7c565b9150509250929050565b600080600060608486031215613d8257600080fd5b6000613d8e8686613a7c565b9350506020613d9f86828701613a7c565b9250506040613db086828701613a9d565b9150509250925092565b60008060408385031215613dcd57600080fd5b6000613dd98585613a7c565b925050602083013567ffffffffffffffff811115613df657600080fd5b613d6385828601613aa8565b60008060408385031215613e1557600080fd5b6000613e218585613a7c565b9250506020613d6385828601613a9d565b600060208284031215613e4457600080fd5b60006117eb8484613a9d565b60008060008060006103008688031215613e6957600080fd5b6000613e758888613afc565b955050610140613e8788828901613be9565b9450506102a0613e9988828901613a7c565b9350506102c0613eab88828901613a9d565b9250506102e0613ebd88828901613a92565b9150509295509295909350565b600060208284031215613edc57600080fd5b60006117eb8484613cec565b60008060408385031215613efb57600080fd5b6000613f078585613a9d565b9250506020613d6385828601613a92565b60008060408385031215613f2b57600080fd5b6000613e218585613a9d565b600080600060608486031215613f4c57600080fd5b6000613f588686613cec565b9350506020613f6986828701613cec565b9250506040613db086828701613cec565b60008060008060808587031215613f9057600080fd5b6000613f9c8787613a9d565b9450506020613fad87828801613a9d565b9350506040613fbe87828801613a9d565b9250506060613fcf87828801613a7c565b91505092959194509250565b60008060008060808587031215613ff157600080fd5b6000613ffd8787613cec565b945050602061400e87828801613cec565b935050604061401f87828801613cec565b9250506060613fcf87828801613cec565b600080600080600080600080610100898b03121561404d57600080fd5b60006140598b8b613a9d565b985050602061406a8b828c01613a9d565b975050604061407b8b828c01613a9d565b965050606061408c8b828c01613a9d565b955050608061409d8b828c01613a7c565b94505060a06140ae8b828c01613a7c565b93505060c06140bf8b828c01613a7c565b92505060e089013567ffffffffffffffff8111156140dc57600080fd5b6140e88b828c01613aa8565b9150509295985092959890939650565b6000806000806000806000806000806101408b8d03121561411857600080fd5b60006141248d8d613a9d565b9a505060206141358d828e01613a9d565b99505060406141468d828e01613a9d565b98505060606141578d828e01613a9d565b97505060806141688d828e01613a9d565b96505060a06141798d828e01613a7c565b95505060c061418a8d828e01613a7c565b94505060e061419b8d828e01613a7c565b9350506101006141ad8d828e01613a7c565b9250506101208b013567ffffffffffffffff8111156141cb57600080fd5b6141d78d828e01613aa8565b9150509295989b9194979a5092959850565b60006141f58383614209565b505060200190565b60006141f58383614323565b61421281614c63565b82525050565b61421261422482614c63565b614cca565b61423281614c4a565b61423c8184610a08565b925061424782610b13565b8060005b83811015612ae457815161425f87826141e9565b965061426a83614c44565b92505060010161424b565b61427e81614c50565b6142888184610a08565b925061429382610b13565b8060005b83811015612ae45781516142ab87826141fd565b96506142b683614c44565b925050600101614297565b60006142cc82614c56565b6142d68185614c5a565b93506142e183614c44565b8060005b8381101561430f5781516142f988826141fd565b975061430483614c44565b9250506001016142e5565b509495945050505050565b61421281614c6e565b61421281610b13565b61421261433882614c73565b610b13565b600061434882614c56565b6143528185614c5a565b9350614362818560208601614c9e565b61436b81614cdb565b9093019392505050565b600061438082614c56565b61438a8185610a08565b935061439a818560208601614c9e565b9290920192915050565b60006143b1600183614c5a565b600d60fa1b815260200192915050565b60006143ce600283614c5a565b61323360f01b815260200192915050565b60006143ec600283614c5a565b61031360f41b815260200192915050565b600061440a600283614c5a565b61313560f01b815260200192915050565b6000614428600283614c5a565b61189b60f11b815260200192915050565b6000614446600183614c5a565b603360f81b815260200192915050565b6000614463600183614c5a565b603760f81b815260200192915050565b6000614480600283614c5a565b61333760f01b815260200192915050565b600061449e600283614c5a565b610c4d60f21b815260200192915050565b60006144bc600283614c5a565b61333560f01b815260200192915050565b60006144da600283614c5a565b610c8d60f21b815260200192915050565b60006144f8600283614c5a565b61313160f01b815260200192915050565b6000614516600283614c5a565b61313360f01b815260200192915050565b6000614534600283614c5a565b61323560f01b815260200192915050565b6000614552600283614c5a565b61313760f01b815260200192915050565b6000614570600c83614c5a565b6b1d5b985d5d1a1bdc9a5e995960a21b815260200192915050565b6000614598600283614c5a565b61313960f01b815260200192915050565b60006145b6600283614c5a565b61191b60f11b815260200192915050565b60006145d4600183614c5a565b601960f91b815260200192915050565b60006145f1600183614c5a565b603160f81b815260200192915050565b600061460e600283614c5a565b61191960f11b815260200192915050565b600061462c600c83614c5a565b6b1b9bdb9499595b9d1c985b9d60a21b815260200192915050565b6000614654600183614c5a565b601b60f91b815260200192915050565b6000614671600183614c5a565b600760fb1b815260200192915050565b600061468e600283614c5a565b61333360f01b815260200192915050565b60006146ac600283614c5a565b61323160f01b815260200192915050565b80516101008301906146cf8482614323565b5060208201516146e26020850182614323565b5060408201516146f56040850182614323565b5060608201516147086060850182614323565b50608082015161471b6080850182614323565b5060a082015161472e60a0850182614323565b5060c082015161474160c0850182614323565b5060e082015161475460e0850182614209565b50505050565b61421261433882610b13565b61421281614c8c565b600061477b828561432c565b60048201915061478b828461475a565b5060200192915050565b6000610f708284614375565b60006147ad828561475a565b6020820191506147bd8284614218565b5060140192915050565b60208101610afd8284614209565b604081016147e38285614209565b610f706020830184614209565b606081016147fe8286614209565b61480b6020830185614209565b6117eb6040830184614209565b60a081016148268288614209565b6148336020830187614209565b6148406040830186614209565b61484d6060830185614323565b61485a6080830184614323565b9695505050505050565b606081016148728286614209565b61487f6020830185614209565b6117eb6040830184614323565b6040810161489a8285614209565b610f706020830184614323565b606081016148b58286614209565b6148c26020830185614323565b6117eb604083018461431a565b60208082528101610f7081846142c1565b60208101610afd828461431a565b60208101610afd8284614323565b6101a0810161490b8287614323565b6149186020830186614229565b61492560a0830185614275565b81810361018083015261485a818461433d565b6101008101614947828b614323565b614954602083018a614323565b6149616040830189614323565b61496e6060830188614323565b61497b6080830187614323565b61498860a0830186614323565b61499560c0830185614323565b6149a260e0830184614209565b9998505050505050505050565b60208082528101610f70818461433d565b60208082528101610a05816143a4565b60208082528101610a05816143c1565b60208082528101610a05816143df565b60208082528101610a05816143fd565b60208082528101610a058161441b565b60208082528101610a0581614439565b60208082528101610a0581614456565b60208082528101610a0581614473565b60208082528101610a0581614491565b60208082528101610a05816144af565b60208082528101610a05816144cd565b60208082528101610a05816144eb565b60208082528101610a0581614509565b60208082528101610a0581614527565b60208082528101610a0581614545565b60208082528101610a0581614563565b60208082528101610a058161458b565b60208082528101610a05816145a9565b60208082528101610a05816145c7565b60208082528101610a05816145e4565b60208082528101610a0581614601565b60208082528101610a058161461f565b60208082528101610a0581614647565b60208082528101610a0581614664565b60208082528101610a0581614681565b60208082528101610a058161469f565b6101008101610afd82846146bd565b60408101614b7d8285614323565b610f70602083018461431a565b60a08101614b988288614323565b614ba56020830187614323565b614bb26040830186614209565b614bbf6060830185614209565b61485a608083018461431a565b60608101614bda8286614323565b61487f6020830185614323565b60208101610afd8284614766565b60405181810167ffffffffffffffff81118282101715614c1457600080fd5b604052919050565b600067ffffffffffffffff821115614c3357600080fd5b506020601f91909101601f19160190565b60200190565b50600490565b50600790565b5190565b90815260200190565b6000610a0582614c80565b151590565b6001600160e01b03191690565b6001600160a01b031690565b60ff1690565b82818337506000910152565b60005b83811015614cb9578181015183820152602001614ca1565b838111156147545750506000910152565b6000610a05826000610a0582614ce5565b601f01601f191690565b60601b90565b614cf481614c63565b81146120fb57600080fd5b614cf481614c6e565b614cf481610b1356feddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa365627a7a723158202b80d18d52b494a55fc164e716e1df5fc6aaae24c069e582cbbaac841e7f475f6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployLoanTokenLogicV4 deploys a new Ethereum contract, binding an instance of LoanTokenLogicV4 to it.
func DeployLoanTokenLogicV4(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoanTokenLogicV4, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenLogicV4ABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoanTokenLogicV4Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoanTokenLogicV4{LoanTokenLogicV4Caller: LoanTokenLogicV4Caller{contract: contract}, LoanTokenLogicV4Transactor: LoanTokenLogicV4Transactor{contract: contract}, LoanTokenLogicV4Filterer: LoanTokenLogicV4Filterer{contract: contract}}, nil
}

// LoanTokenLogicV4 is an auto generated Go binding around an Ethereum contract.
type LoanTokenLogicV4 struct {
	LoanTokenLogicV4Caller     // Read-only binding to the contract
	LoanTokenLogicV4Transactor // Write-only binding to the contract
	LoanTokenLogicV4Filterer   // Log filterer for contract events
}

// LoanTokenLogicV4Caller is an auto generated read-only Go binding around an Ethereum contract.
type LoanTokenLogicV4Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenLogicV4Transactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanTokenLogicV4Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenLogicV4Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanTokenLogicV4Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenLogicV4Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanTokenLogicV4Session struct {
	Contract     *LoanTokenLogicV4 // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanTokenLogicV4CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanTokenLogicV4CallerSession struct {
	Contract *LoanTokenLogicV4Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LoanTokenLogicV4TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanTokenLogicV4TransactorSession struct {
	Contract     *LoanTokenLogicV4Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LoanTokenLogicV4Raw is an auto generated low-level Go binding around an Ethereum contract.
type LoanTokenLogicV4Raw struct {
	Contract *LoanTokenLogicV4 // Generic contract binding to access the raw methods on
}

// LoanTokenLogicV4CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanTokenLogicV4CallerRaw struct {
	Contract *LoanTokenLogicV4Caller // Generic read-only contract binding to access the raw methods on
}

// LoanTokenLogicV4TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanTokenLogicV4TransactorRaw struct {
	Contract *LoanTokenLogicV4Transactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanTokenLogicV4 creates a new instance of LoanTokenLogicV4, bound to a specific deployed contract.
func NewLoanTokenLogicV4(address common.Address, backend bind.ContractBackend) (*LoanTokenLogicV4, error) {
	contract, err := bindLoanTokenLogicV4(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4{LoanTokenLogicV4Caller: LoanTokenLogicV4Caller{contract: contract}, LoanTokenLogicV4Transactor: LoanTokenLogicV4Transactor{contract: contract}, LoanTokenLogicV4Filterer: LoanTokenLogicV4Filterer{contract: contract}}, nil
}

// NewLoanTokenLogicV4Caller creates a new read-only instance of LoanTokenLogicV4, bound to a specific deployed contract.
func NewLoanTokenLogicV4Caller(address common.Address, caller bind.ContractCaller) (*LoanTokenLogicV4Caller, error) {
	contract, err := bindLoanTokenLogicV4(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4Caller{contract: contract}, nil
}

// NewLoanTokenLogicV4Transactor creates a new write-only instance of LoanTokenLogicV4, bound to a specific deployed contract.
func NewLoanTokenLogicV4Transactor(address common.Address, transactor bind.ContractTransactor) (*LoanTokenLogicV4Transactor, error) {
	contract, err := bindLoanTokenLogicV4(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4Transactor{contract: contract}, nil
}

// NewLoanTokenLogicV4Filterer creates a new log filterer instance of LoanTokenLogicV4, bound to a specific deployed contract.
func NewLoanTokenLogicV4Filterer(address common.Address, filterer bind.ContractFilterer) (*LoanTokenLogicV4Filterer, error) {
	contract, err := bindLoanTokenLogicV4(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4Filterer{contract: contract}, nil
}

// bindLoanTokenLogicV4 binds a generic wrapper to an already deployed contract.
func bindLoanTokenLogicV4(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenLogicV4ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenLogicV4 *LoanTokenLogicV4Raw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenLogicV4.Contract.LoanTokenLogicV4Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenLogicV4 *LoanTokenLogicV4Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.LoanTokenLogicV4Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenLogicV4 *LoanTokenLogicV4Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.LoanTokenLogicV4Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenLogicV4.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.contract.Transact(opts, method, params...)
}

// SupplyInterestRate1 is a free data retrieval call binding the contract method 0x7288b344.
//
// Solidity: function _supplyInterestRate(uint256 assetBorrow, uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) SupplyInterestRate1(opts *bind.CallOpts, assetBorrow *big.Int, assetSupply *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "_supplyInterestRate", assetBorrow, assetSupply)
	return *ret0, err
}

// SupplyInterestRate1 is a free data retrieval call binding the contract method 0x7288b344.
//
// Solidity: function _supplyInterestRate(uint256 assetBorrow, uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) SupplyInterestRate1(assetBorrow *big.Int, assetSupply *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SupplyInterestRate1(&_LoanTokenLogicV4.CallOpts, assetBorrow, assetSupply)
}

// SupplyInterestRate1 is a free data retrieval call binding the contract method 0x7288b344.
//
// Solidity: function _supplyInterestRate(uint256 assetBorrow, uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) SupplyInterestRate1(assetBorrow *big.Int, assetSupply *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SupplyInterestRate1(&_LoanTokenLogicV4.CallOpts, assetBorrow, assetSupply)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.Allowance(&_LoanTokenLogicV4.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.Allowance(&_LoanTokenLogicV4.CallOpts, _owner, _spender)
}

// AssetBalanceOf is a free data retrieval call binding the contract method 0x06b3efd6.
//
// Solidity: function assetBalanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) AssetBalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "assetBalanceOf", _owner)
	return *ret0, err
}

// AssetBalanceOf is a free data retrieval call binding the contract method 0x06b3efd6.
//
// Solidity: function assetBalanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) AssetBalanceOf(_owner common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.AssetBalanceOf(&_LoanTokenLogicV4.CallOpts, _owner)
}

// AssetBalanceOf is a free data retrieval call binding the contract method 0x06b3efd6.
//
// Solidity: function assetBalanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) AssetBalanceOf(_owner common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.AssetBalanceOf(&_LoanTokenLogicV4.CallOpts, _owner)
}

// AvgBorrowInterestRate is a free data retrieval call binding the contract method 0x44a4a003.
//
// Solidity: function avgBorrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) AvgBorrowInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "avgBorrowInterestRate")
	return *ret0, err
}

// AvgBorrowInterestRate is a free data retrieval call binding the contract method 0x44a4a003.
//
// Solidity: function avgBorrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) AvgBorrowInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.AvgBorrowInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// AvgBorrowInterestRate is a free data retrieval call binding the contract method 0x44a4a003.
//
// Solidity: function avgBorrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) AvgBorrowInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.AvgBorrowInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BZxContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "bZxContract")
	return *ret0, err
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BZxContract() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxContract(&_LoanTokenLogicV4.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BZxContract() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxContract(&_LoanTokenLogicV4.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BZxOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "bZxOracle")
	return *ret0, err
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BZxOracle() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxOracle(&_LoanTokenLogicV4.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BZxOracle() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxOracle(&_LoanTokenLogicV4.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BZxVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "bZxVault")
	return *ret0, err
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BZxVault() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxVault(&_LoanTokenLogicV4.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BZxVault() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.BZxVault(&_LoanTokenLogicV4.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BalanceOf(&_LoanTokenLogicV4.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BalanceOf(&_LoanTokenLogicV4.CallOpts, _owner)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "baseRate")
	return *ret0, err
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BaseRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BaseRate(&_LoanTokenLogicV4.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BaseRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BaseRate(&_LoanTokenLogicV4.CallOpts)
}

// BorrowInterestRate is a free data retrieval call binding the contract method 0x8325a1c0.
//
// Solidity: function borrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BorrowInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "borrowInterestRate")
	return *ret0, err
}

// BorrowInterestRate is a free data retrieval call binding the contract method 0x8325a1c0.
//
// Solidity: function borrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BorrowInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BorrowInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// BorrowInterestRate is a free data retrieval call binding the contract method 0x8325a1c0.
//
// Solidity: function borrowInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BorrowInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BorrowInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BurntTokenReserveList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Lender common.Address
		Amount *big.Int
	})
	out := ret
	err := _LoanTokenLogicV4.contract.Call(opts, out, "burntTokenReserveList", arg0)
	return *ret, err
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserveList(&_LoanTokenLogicV4.CallOpts, arg0)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserveList(&_LoanTokenLogicV4.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BurntTokenReserveListIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	ret := new(struct {
		Index *big.Int
		IsSet bool
	})
	out := ret
	err := _LoanTokenLogicV4.contract.Call(opts, out, "burntTokenReserveListIndex", arg0)
	return *ret, err
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserveListIndex(&_LoanTokenLogicV4.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserveListIndex(&_LoanTokenLogicV4.CallOpts, arg0)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) BurntTokenReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "burntTokenReserved")
	return *ret0, err
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BurntTokenReserved() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserved(&_LoanTokenLogicV4.CallOpts)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) BurntTokenReserved() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.BurntTokenReserved(&_LoanTokenLogicV4.CallOpts)
}

// CheckpointPrice is a free data retrieval call binding the contract method 0xeebc5081.
//
// Solidity: function checkpointPrice(address _user) view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) CheckpointPrice(opts *bind.CallOpts, _user common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "checkpointPrice", _user)
	return *ret0, err
}

// CheckpointPrice is a free data retrieval call binding the contract method 0xeebc5081.
//
// Solidity: function checkpointPrice(address _user) view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) CheckpointPrice(_user common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.CheckpointPrice(&_LoanTokenLogicV4.CallOpts, _user)
}

// CheckpointPrice is a free data retrieval call binding the contract method 0xeebc5081.
//
// Solidity: function checkpointPrice(address _user) view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) CheckpointPrice(_user common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.CheckpointPrice(&_LoanTokenLogicV4.CallOpts, _user)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) CheckpointSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "checkpointSupply")
	return *ret0, err
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) CheckpointSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.CheckpointSupply(&_LoanTokenLogicV4.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) CheckpointSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.CheckpointSupply(&_LoanTokenLogicV4.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Decimals() (uint8, error) {
	return _LoanTokenLogicV4.Contract.Decimals(&_LoanTokenLogicV4.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) Decimals() (uint8, error) {
	return _LoanTokenLogicV4.Contract.Decimals(&_LoanTokenLogicV4.CallOpts)
}

// GetBorrowAmountForDeposit is a free data retrieval call binding the contract method 0x24d25f4a.
//
// Solidity: function getBorrowAmountForDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 borrowAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) GetBorrowAmountForDeposit(opts *bind.CallOpts, depositAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "getBorrowAmountForDeposit", depositAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
	return *ret0, err
}

// GetBorrowAmountForDeposit is a free data retrieval call binding the contract method 0x24d25f4a.
//
// Solidity: function getBorrowAmountForDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 borrowAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) GetBorrowAmountForDeposit(depositAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetBorrowAmountForDeposit(&_LoanTokenLogicV4.CallOpts, depositAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
}

// GetBorrowAmountForDeposit is a free data retrieval call binding the contract method 0x24d25f4a.
//
// Solidity: function getBorrowAmountForDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 borrowAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) GetBorrowAmountForDeposit(depositAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetBorrowAmountForDeposit(&_LoanTokenLogicV4.CallOpts, depositAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
}

// GetDepositAmountForBorrow is a free data retrieval call binding the contract method 0x8423acd6.
//
// Solidity: function getDepositAmountForBorrow(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 depositAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) GetDepositAmountForBorrow(opts *bind.CallOpts, borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "getDepositAmountForBorrow", borrowAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
	return *ret0, err
}

// GetDepositAmountForBorrow is a free data retrieval call binding the contract method 0x8423acd6.
//
// Solidity: function getDepositAmountForBorrow(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 depositAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) GetDepositAmountForBorrow(borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetDepositAmountForBorrow(&_LoanTokenLogicV4.CallOpts, borrowAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
}

// GetDepositAmountForBorrow is a free data retrieval call binding the contract method 0x8423acd6.
//
// Solidity: function getDepositAmountForBorrow(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, address collateralTokenAddress) view returns(uint256 depositAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) GetDepositAmountForBorrow(borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenAddress common.Address) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetDepositAmountForBorrow(&_LoanTokenLogicV4.CallOpts, borrowAmount, leverageAmount, initialLoanDuration, collateralTokenAddress)
}

// GetLeverageList is a free data retrieval call binding the contract method 0x2ecae90a.
//
// Solidity: function getLeverageList() view returns(uint256[])
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) GetLeverageList(opts *bind.CallOpts) ([]*big.Int, error) {
	var (
		ret0 = new([]*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "getLeverageList")
	return *ret0, err
}

// GetLeverageList is a free data retrieval call binding the contract method 0x2ecae90a.
//
// Solidity: function getLeverageList() view returns(uint256[])
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) GetLeverageList() ([]*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetLeverageList(&_LoanTokenLogicV4.CallOpts)
}

// GetLeverageList is a free data retrieval call binding the contract method 0x2ecae90a.
//
// Solidity: function getLeverageList() view returns(uint256[])
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) GetLeverageList() ([]*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetLeverageList(&_LoanTokenLogicV4.CallOpts)
}

// GetLoanData is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 loanOrderHash) view returns((bytes32,uint256,uint256,uint256,uint256,uint256,uint256,address))
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) GetLoanData(opts *bind.CallOpts, loanOrderHash [32]byte) (LoanTokenStorageLoanData, error) {
	var (
		ret0 = new(LoanTokenStorageLoanData)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "getLoanData", loanOrderHash)
	return *ret0, err
}

// GetLoanData is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 loanOrderHash) view returns((bytes32,uint256,uint256,uint256,uint256,uint256,uint256,address))
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) GetLoanData(loanOrderHash [32]byte) (LoanTokenStorageLoanData, error) {
	return _LoanTokenLogicV4.Contract.GetLoanData(&_LoanTokenLogicV4.CallOpts, loanOrderHash)
}

// GetLoanData is a free data retrieval call binding the contract method 0xc4d2b1b3.
//
// Solidity: function getLoanData(bytes32 loanOrderHash) view returns((bytes32,uint256,uint256,uint256,uint256,uint256,uint256,address))
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) GetLoanData(loanOrderHash [32]byte) (LoanTokenStorageLoanData, error) {
	return _LoanTokenLogicV4.Contract.GetLoanData(&_LoanTokenLogicV4.CallOpts, loanOrderHash)
}

// GetMaxEscrowAmount is a free data retrieval call binding the contract method 0x829b38f4.
//
// Solidity: function getMaxEscrowAmount(uint256 leverageAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) GetMaxEscrowAmount(opts *bind.CallOpts, leverageAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "getMaxEscrowAmount", leverageAmount)
	return *ret0, err
}

// GetMaxEscrowAmount is a free data retrieval call binding the contract method 0x829b38f4.
//
// Solidity: function getMaxEscrowAmount(uint256 leverageAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) GetMaxEscrowAmount(leverageAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetMaxEscrowAmount(&_LoanTokenLogicV4.CallOpts, leverageAmount)
}

// GetMaxEscrowAmount is a free data retrieval call binding the contract method 0x829b38f4.
//
// Solidity: function getMaxEscrowAmount(uint256 leverageAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) GetMaxEscrowAmount(leverageAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.GetMaxEscrowAmount(&_LoanTokenLogicV4.CallOpts, leverageAmount)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) InitialPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "initialPrice")
	return *ret0, err
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) InitialPrice() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.InitialPrice(&_LoanTokenLogicV4.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) InitialPrice() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.InitialPrice(&_LoanTokenLogicV4.CallOpts)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) LeverageList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "leverageList", arg0)
	return *ret0, err
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.LeverageList(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.LeverageList(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) LoanOrderData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	ret := new(struct {
		LoanOrderHash               [32]byte
		LeverageAmount              *big.Int
		InitialMarginAmount         *big.Int
		MaintenanceMarginAmount     *big.Int
		MaxDurationUnixTimestampSec *big.Int
		Index                       *big.Int
		MarginPremiumAmount         *big.Int
		CollateralTokenAddress      common.Address
	})
	out := ret
	err := _LoanTokenLogicV4.contract.Call(opts, out, "loanOrderData", arg0)
	return *ret, err
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _LoanTokenLogicV4.Contract.LoanOrderData(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _LoanTokenLogicV4.Contract.LoanOrderData(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) LoanOrderHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "loanOrderHashes", arg0)
	return *ret0, err
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _LoanTokenLogicV4.Contract.LoanOrderHashes(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _LoanTokenLogicV4.Contract.LoanOrderHashes(&_LoanTokenLogicV4.CallOpts, arg0)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) LoanTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "loanTokenAddress")
	return *ret0, err
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.LoanTokenAddress(&_LoanTokenLogicV4.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.LoanTokenAddress(&_LoanTokenLogicV4.CallOpts)
}

// MarketLiquidity is a free data retrieval call binding the contract method 0x612ef80b.
//
// Solidity: function marketLiquidity() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) MarketLiquidity(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "marketLiquidity")
	return *ret0, err
}

// MarketLiquidity is a free data retrieval call binding the contract method 0x612ef80b.
//
// Solidity: function marketLiquidity() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) MarketLiquidity() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.MarketLiquidity(&_LoanTokenLogicV4.CallOpts)
}

// MarketLiquidity is a free data retrieval call binding the contract method 0x612ef80b.
//
// Solidity: function marketLiquidity() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) MarketLiquidity() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.MarketLiquidity(&_LoanTokenLogicV4.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Name() (string, error) {
	return _LoanTokenLogicV4.Contract.Name(&_LoanTokenLogicV4.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) Name() (string, error) {
	return _LoanTokenLogicV4.Contract.Name(&_LoanTokenLogicV4.CallOpts)
}

// NextBorrowInterestRate is a free data retrieval call binding the contract method 0xb9fe1a8f.
//
// Solidity: function nextBorrowInterestRate(uint256 borrowAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) NextBorrowInterestRate(opts *bind.CallOpts, borrowAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "nextBorrowInterestRate", borrowAmount)
	return *ret0, err
}

// NextBorrowInterestRate is a free data retrieval call binding the contract method 0xb9fe1a8f.
//
// Solidity: function nextBorrowInterestRate(uint256 borrowAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) NextBorrowInterestRate(borrowAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextBorrowInterestRate(&_LoanTokenLogicV4.CallOpts, borrowAmount)
}

// NextBorrowInterestRate is a free data retrieval call binding the contract method 0xb9fe1a8f.
//
// Solidity: function nextBorrowInterestRate(uint256 borrowAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) NextBorrowInterestRate(borrowAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextBorrowInterestRate(&_LoanTokenLogicV4.CallOpts, borrowAmount)
}

// NextBorrowInterestRateWithOption is a free data retrieval call binding the contract method 0x7d90dcba.
//
// Solidity: function nextBorrowInterestRateWithOption(uint256 borrowAmount, bool useFixedInterestModel) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) NextBorrowInterestRateWithOption(opts *bind.CallOpts, borrowAmount *big.Int, useFixedInterestModel bool) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "nextBorrowInterestRateWithOption", borrowAmount, useFixedInterestModel)
	return *ret0, err
}

// NextBorrowInterestRateWithOption is a free data retrieval call binding the contract method 0x7d90dcba.
//
// Solidity: function nextBorrowInterestRateWithOption(uint256 borrowAmount, bool useFixedInterestModel) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) NextBorrowInterestRateWithOption(borrowAmount *big.Int, useFixedInterestModel bool) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextBorrowInterestRateWithOption(&_LoanTokenLogicV4.CallOpts, borrowAmount, useFixedInterestModel)
}

// NextBorrowInterestRateWithOption is a free data retrieval call binding the contract method 0x7d90dcba.
//
// Solidity: function nextBorrowInterestRateWithOption(uint256 borrowAmount, bool useFixedInterestModel) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) NextBorrowInterestRateWithOption(borrowAmount *big.Int, useFixedInterestModel bool) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextBorrowInterestRateWithOption(&_LoanTokenLogicV4.CallOpts, borrowAmount, useFixedInterestModel)
}

// NextSupplyInterestRate is a free data retrieval call binding the contract method 0xd65a5021.
//
// Solidity: function nextSupplyInterestRate(uint256 supplyAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) NextSupplyInterestRate(opts *bind.CallOpts, supplyAmount *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "nextSupplyInterestRate", supplyAmount)
	return *ret0, err
}

// NextSupplyInterestRate is a free data retrieval call binding the contract method 0xd65a5021.
//
// Solidity: function nextSupplyInterestRate(uint256 supplyAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) NextSupplyInterestRate(supplyAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextSupplyInterestRate(&_LoanTokenLogicV4.CallOpts, supplyAmount)
}

// NextSupplyInterestRate is a free data retrieval call binding the contract method 0xd65a5021.
//
// Solidity: function nextSupplyInterestRate(uint256 supplyAmount) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) NextSupplyInterestRate(supplyAmount *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.NextSupplyInterestRate(&_LoanTokenLogicV4.CallOpts, supplyAmount)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Owner() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.Owner(&_LoanTokenLogicV4.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) Owner() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.Owner(&_LoanTokenLogicV4.CallOpts)
}

// ProtocolInterestRate is a free data retrieval call binding the contract method 0xfc3b72b1.
//
// Solidity: function protocolInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) ProtocolInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "protocolInterestRate")
	return *ret0, err
}

// ProtocolInterestRate is a free data retrieval call binding the contract method 0xfc3b72b1.
//
// Solidity: function protocolInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) ProtocolInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.ProtocolInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// ProtocolInterestRate is a free data retrieval call binding the contract method 0xfc3b72b1.
//
// Solidity: function protocolInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) ProtocolInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.ProtocolInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) RateMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "rateMultiplier")
	return *ret0, err
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) RateMultiplier() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.RateMultiplier(&_LoanTokenLogicV4.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) RateMultiplier() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.RateMultiplier(&_LoanTokenLogicV4.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) SpreadMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "spreadMultiplier")
	return *ret0, err
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) SpreadMultiplier() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SpreadMultiplier(&_LoanTokenLogicV4.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) SpreadMultiplier() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SpreadMultiplier(&_LoanTokenLogicV4.CallOpts)
}

// SupplyInterestRate is a free data retrieval call binding the contract method 0x09ec6b6b.
//
// Solidity: function supplyInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) SupplyInterestRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "supplyInterestRate")
	return *ret0, err
}

// SupplyInterestRate is a free data retrieval call binding the contract method 0x09ec6b6b.
//
// Solidity: function supplyInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) SupplyInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SupplyInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// SupplyInterestRate is a free data retrieval call binding the contract method 0x09ec6b6b.
//
// Solidity: function supplyInterestRate() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) SupplyInterestRate() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.SupplyInterestRate(&_LoanTokenLogicV4.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Symbol() (string, error) {
	return _LoanTokenLogicV4.Contract.Symbol(&_LoanTokenLogicV4.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) Symbol() (string, error) {
	return _LoanTokenLogicV4.Contract.Symbol(&_LoanTokenLogicV4.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TokenPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "tokenPrice")
	return *ret0, err
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TokenPrice() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TokenPrice(&_LoanTokenLogicV4.CallOpts)
}

// TokenPrice is a free data retrieval call binding the contract method 0x7ff9b596.
//
// Solidity: function tokenPrice() view returns(uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TokenPrice() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TokenPrice(&_LoanTokenLogicV4.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TokenizedRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "tokenizedRegistry")
	return *ret0, err
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TokenizedRegistry() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.TokenizedRegistry(&_LoanTokenLogicV4.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TokenizedRegistry() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.TokenizedRegistry(&_LoanTokenLogicV4.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TotalAssetBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "totalAssetBorrow")
	return *ret0, err
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TotalAssetBorrow() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalAssetBorrow(&_LoanTokenLogicV4.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TotalAssetBorrow() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalAssetBorrow(&_LoanTokenLogicV4.CallOpts)
}

// TotalAssetSupply is a free data retrieval call binding the contract method 0x8fb807c5.
//
// Solidity: function totalAssetSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TotalAssetSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "totalAssetSupply")
	return *ret0, err
}

// TotalAssetSupply is a free data retrieval call binding the contract method 0x8fb807c5.
//
// Solidity: function totalAssetSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TotalAssetSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalAssetSupply(&_LoanTokenLogicV4.CallOpts)
}

// TotalAssetSupply is a free data retrieval call binding the contract method 0x8fb807c5.
//
// Solidity: function totalAssetSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TotalAssetSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalAssetSupply(&_LoanTokenLogicV4.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TotalSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalSupply(&_LoanTokenLogicV4.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TotalSupply() (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalSupply(&_LoanTokenLogicV4.CallOpts)
}

// TotalSupplyInterestRate is a free data retrieval call binding the contract method 0x12416898.
//
// Solidity: function totalSupplyInterestRate(uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) TotalSupplyInterestRate(opts *bind.CallOpts, assetSupply *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "totalSupplyInterestRate", assetSupply)
	return *ret0, err
}

// TotalSupplyInterestRate is a free data retrieval call binding the contract method 0x12416898.
//
// Solidity: function totalSupplyInterestRate(uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TotalSupplyInterestRate(assetSupply *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalSupplyInterestRate(&_LoanTokenLogicV4.CallOpts, assetSupply)
}

// TotalSupplyInterestRate is a free data retrieval call binding the contract method 0x12416898.
//
// Solidity: function totalSupplyInterestRate(uint256 assetSupply) view returns(uint256)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) TotalSupplyInterestRate(assetSupply *big.Int) (*big.Int, error) {
	return _LoanTokenLogicV4.Contract.TotalSupplyInterestRate(&_LoanTokenLogicV4.CallOpts, assetSupply)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Caller) WethContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenLogicV4.contract.Call(opts, out, "wethContract")
	return *ret0, err
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) WethContract() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.WethContract(&_LoanTokenLogicV4.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenLogicV4 *LoanTokenLogicV4CallerSession) WethContract() (common.Address, error) {
	return _LoanTokenLogicV4.Contract.WethContract(&_LoanTokenLogicV4.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Approve(&_LoanTokenLogicV4.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Approve(&_LoanTokenLogicV4.TransactOpts, _spender, _value)
}

// BorrowTokenFromDeposit is a paid mutator transaction binding the contract method 0xcfb65bb9.
//
// Solidity: function borrowTokenFromDeposit(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, uint256 collateralTokenSent, address borrower, address receiver, address collateralTokenAddress, bytes ) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) BorrowTokenFromDeposit(opts *bind.TransactOpts, borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenSent *big.Int, borrower common.Address, receiver common.Address, collateralTokenAddress common.Address, arg7 []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "borrowTokenFromDeposit", borrowAmount, leverageAmount, initialLoanDuration, collateralTokenSent, borrower, receiver, collateralTokenAddress, arg7)
}

// BorrowTokenFromDeposit is a paid mutator transaction binding the contract method 0xcfb65bb9.
//
// Solidity: function borrowTokenFromDeposit(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, uint256 collateralTokenSent, address borrower, address receiver, address collateralTokenAddress, bytes ) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BorrowTokenFromDeposit(borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenSent *big.Int, borrower common.Address, receiver common.Address, collateralTokenAddress common.Address, arg7 []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.BorrowTokenFromDeposit(&_LoanTokenLogicV4.TransactOpts, borrowAmount, leverageAmount, initialLoanDuration, collateralTokenSent, borrower, receiver, collateralTokenAddress, arg7)
}

// BorrowTokenFromDeposit is a paid mutator transaction binding the contract method 0xcfb65bb9.
//
// Solidity: function borrowTokenFromDeposit(uint256 borrowAmount, uint256 leverageAmount, uint256 initialLoanDuration, uint256 collateralTokenSent, address borrower, address receiver, address collateralTokenAddress, bytes ) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) BorrowTokenFromDeposit(borrowAmount *big.Int, leverageAmount *big.Int, initialLoanDuration *big.Int, collateralTokenSent *big.Int, borrower common.Address, receiver common.Address, collateralTokenAddress common.Address, arg7 []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.BorrowTokenFromDeposit(&_LoanTokenLogicV4.TransactOpts, borrowAmount, leverageAmount, initialLoanDuration, collateralTokenSent, borrower, receiver, collateralTokenAddress, arg7)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) Burn(opts *bind.TransactOpts, receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "burn", receiver, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Burn(receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Burn(&_LoanTokenLogicV4.TransactOpts, receiver, burnAmount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) Burn(receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Burn(&_LoanTokenLogicV4.TransactOpts, receiver, burnAmount)
}

// BurnToEther is a paid mutator transaction binding the contract method 0x81a6b250.
//
// Solidity: function burnToEther(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) BurnToEther(opts *bind.TransactOpts, receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "burnToEther", receiver, burnAmount)
}

// BurnToEther is a paid mutator transaction binding the contract method 0x81a6b250.
//
// Solidity: function burnToEther(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) BurnToEther(receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.BurnToEther(&_LoanTokenLogicV4.TransactOpts, receiver, burnAmount)
}

// BurnToEther is a paid mutator transaction binding the contract method 0x81a6b250.
//
// Solidity: function burnToEther(address receiver, uint256 burnAmount) returns(uint256 loanAmountPaid)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) BurnToEther(receiver common.Address, burnAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.BurnToEther(&_LoanTokenLogicV4.TransactOpts, receiver, burnAmount)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) CloseLoanNotifier(opts *bind.TransactOpts, loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "closeLoanNotifier", loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) CloseLoanNotifier(loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.CloseLoanNotifier(&_LoanTokenLogicV4.TransactOpts, loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) CloseLoanNotifier(loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.CloseLoanNotifier(&_LoanTokenLogicV4.TransactOpts, loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// MarginTradeFromDeposit is a paid mutator transaction binding the contract method 0x1c5d1da5.
//
// Solidity: function marginTradeFromDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 loanTokenSent, uint256 collateralTokenSent, uint256 tradeTokenSent, address trader, address depositTokenAddress, address collateralTokenAddress, address tradeTokenAddress, bytes loanDataBytes) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) MarginTradeFromDeposit(opts *bind.TransactOpts, depositAmount *big.Int, leverageAmount *big.Int, loanTokenSent *big.Int, collateralTokenSent *big.Int, tradeTokenSent *big.Int, trader common.Address, depositTokenAddress common.Address, collateralTokenAddress common.Address, tradeTokenAddress common.Address, loanDataBytes []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "marginTradeFromDeposit", depositAmount, leverageAmount, loanTokenSent, collateralTokenSent, tradeTokenSent, trader, depositTokenAddress, collateralTokenAddress, tradeTokenAddress, loanDataBytes)
}

// MarginTradeFromDeposit is a paid mutator transaction binding the contract method 0x1c5d1da5.
//
// Solidity: function marginTradeFromDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 loanTokenSent, uint256 collateralTokenSent, uint256 tradeTokenSent, address trader, address depositTokenAddress, address collateralTokenAddress, address tradeTokenAddress, bytes loanDataBytes) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) MarginTradeFromDeposit(depositAmount *big.Int, leverageAmount *big.Int, loanTokenSent *big.Int, collateralTokenSent *big.Int, tradeTokenSent *big.Int, trader common.Address, depositTokenAddress common.Address, collateralTokenAddress common.Address, tradeTokenAddress common.Address, loanDataBytes []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.MarginTradeFromDeposit(&_LoanTokenLogicV4.TransactOpts, depositAmount, leverageAmount, loanTokenSent, collateralTokenSent, tradeTokenSent, trader, depositTokenAddress, collateralTokenAddress, tradeTokenAddress, loanDataBytes)
}

// MarginTradeFromDeposit is a paid mutator transaction binding the contract method 0x1c5d1da5.
//
// Solidity: function marginTradeFromDeposit(uint256 depositAmount, uint256 leverageAmount, uint256 loanTokenSent, uint256 collateralTokenSent, uint256 tradeTokenSent, address trader, address depositTokenAddress, address collateralTokenAddress, address tradeTokenAddress, bytes loanDataBytes) payable returns(bytes32 loanOrderHash)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) MarginTradeFromDeposit(depositAmount *big.Int, leverageAmount *big.Int, loanTokenSent *big.Int, collateralTokenSent *big.Int, tradeTokenSent *big.Int, trader common.Address, depositTokenAddress common.Address, collateralTokenAddress common.Address, tradeTokenAddress common.Address, loanDataBytes []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.MarginTradeFromDeposit(&_LoanTokenLogicV4.TransactOpts, depositAmount, leverageAmount, loanTokenSent, collateralTokenSent, tradeTokenSent, trader, depositTokenAddress, collateralTokenAddress, tradeTokenAddress, loanDataBytes)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 depositAmount) returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) Mint(opts *bind.TransactOpts, receiver common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "mint", receiver, depositAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 depositAmount) returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Mint(receiver common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Mint(&_LoanTokenLogicV4.TransactOpts, receiver, depositAmount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address receiver, uint256 depositAmount) returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) Mint(receiver common.Address, depositAmount *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Mint(&_LoanTokenLogicV4.TransactOpts, receiver, depositAmount)
}

// MintWithEther is a paid mutator transaction binding the contract method 0x8f6ede1f.
//
// Solidity: function mintWithEther(address receiver) payable returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) MintWithEther(opts *bind.TransactOpts, receiver common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "mintWithEther", receiver)
}

// MintWithEther is a paid mutator transaction binding the contract method 0x8f6ede1f.
//
// Solidity: function mintWithEther(address receiver) payable returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) MintWithEther(receiver common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.MintWithEther(&_LoanTokenLogicV4.TransactOpts, receiver)
}

// MintWithEther is a paid mutator transaction binding the contract method 0x8f6ede1f.
//
// Solidity: function mintWithEther(address receiver) payable returns(uint256 mintAmount)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) MintWithEther(receiver common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.MintWithEther(&_LoanTokenLogicV4.TransactOpts, receiver)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Transfer(&_LoanTokenLogicV4.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Transfer(&_LoanTokenLogicV4.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.TransferFrom(&_LoanTokenLogicV4.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.TransferFrom(&_LoanTokenLogicV4.TransactOpts, _from, _to, _value)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.TransferOwnership(&_LoanTokenLogicV4.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.TransferOwnership(&_LoanTokenLogicV4.TransactOpts, _newOwner)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x284e2f56.
//
// Solidity: function updateSettings(address settingsTarget, bytes callData) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) UpdateSettings(opts *bind.TransactOpts, settingsTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.Transact(opts, "updateSettings", settingsTarget, callData)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x284e2f56.
//
// Solidity: function updateSettings(address settingsTarget, bytes callData) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) UpdateSettings(settingsTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.UpdateSettings(&_LoanTokenLogicV4.TransactOpts, settingsTarget, callData)
}

// UpdateSettings is a paid mutator transaction binding the contract method 0x284e2f56.
//
// Solidity: function updateSettings(address settingsTarget, bytes callData) returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) UpdateSettings(settingsTarget common.Address, callData []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.UpdateSettings(&_LoanTokenLogicV4.TransactOpts, settingsTarget, callData)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Transactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4Session) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Fallback(&_LoanTokenLogicV4.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() returns()
func (_LoanTokenLogicV4 *LoanTokenLogicV4TransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _LoanTokenLogicV4.Contract.Fallback(&_LoanTokenLogicV4.TransactOpts, calldata)
}

// LoanTokenLogicV4ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4ApprovalIterator struct {
	Event *LoanTokenLogicV4Approval // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Approval)
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
		it.Event = new(LoanTokenLogicV4Approval)
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
func (it *LoanTokenLogicV4ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Approval represents a Approval event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*LoanTokenLogicV4ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4ApprovalIterator{contract: _LoanTokenLogicV4.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Approval)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseApproval(log types.Log) (*LoanTokenLogicV4Approval, error) {
	event := new(LoanTokenLogicV4Approval)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4BorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4BorrowIterator struct {
	Event *LoanTokenLogicV4Borrow // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4BorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Borrow)
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
		it.Event = new(LoanTokenLogicV4Borrow)
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
func (it *LoanTokenLogicV4BorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4BorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Borrow represents a Borrow event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Borrow struct {
	Borrower                common.Address
	BorrowAmount            *big.Int
	InterestRate            *big.Int
	CollateralTokenAddress  common.Address
	TradeTokenToFillAddress common.Address
	WithdrawOnOpen          bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterBorrow(opts *bind.FilterOpts, borrower []common.Address) (*LoanTokenLogicV4BorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4BorrowIterator{contract: _LoanTokenLogicV4.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Borrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Borrow)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseBorrow(log types.Log) (*LoanTokenLogicV4Borrow, error) {
	event := new(LoanTokenLogicV4Borrow)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4BurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4BurnIterator struct {
	Event *LoanTokenLogicV4Burn // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4BurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Burn)
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
		it.Event = new(LoanTokenLogicV4Burn)
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
func (it *LoanTokenLogicV4BurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4BurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Burn represents a Burn event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Burn struct {
	Burner      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterBurn(opts *bind.FilterOpts, burner []common.Address) (*LoanTokenLogicV4BurnIterator, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4BurnIterator{contract: _LoanTokenLogicV4.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Burn, burner []common.Address) (event.Subscription, error) {

	var burnerRule []interface{}
	for _, burnerItem := range burner {
		burnerRule = append(burnerRule, burnerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Burn", burnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Burn)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x743033787f4738ff4d6a7225ce2bd0977ee5f86b91a902a58f5e4d0b297b4644.
//
// Solidity: event Burn(address indexed burner, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseBurn(log types.Log) (*LoanTokenLogicV4Burn, error) {
	event := new(LoanTokenLogicV4Burn)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4ClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4ClaimIterator struct {
	Event *LoanTokenLogicV4Claim // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4ClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Claim)
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
		it.Event = new(LoanTokenLogicV4Claim)
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
func (it *LoanTokenLogicV4ClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4ClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Claim represents a Claim event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Claim struct {
	Claimant             common.Address
	TokenAmount          *big.Int
	AssetAmount          *big.Int
	RemainingTokenAmount *big.Int
	Price                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterClaim(opts *bind.FilterOpts, claimant []common.Address) (*LoanTokenLogicV4ClaimIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4ClaimIterator{contract: _LoanTokenLogicV4.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Claim, claimant []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Claim)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseClaim(log types.Log) (*LoanTokenLogicV4Claim, error) {
	event := new(LoanTokenLogicV4Claim)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4MintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4MintIterator struct {
	Event *LoanTokenLogicV4Mint // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4MintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Mint)
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
		it.Event = new(LoanTokenLogicV4Mint)
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
func (it *LoanTokenLogicV4MintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4MintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Mint represents a Mint event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Mint struct {
	Minter      common.Address
	TokenAmount *big.Int
	AssetAmount *big.Int
	Price       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterMint(opts *bind.FilterOpts, minter []common.Address) (*LoanTokenLogicV4MintIterator, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4MintIterator{contract: _LoanTokenLogicV4.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchMint(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Mint, minter []common.Address) (event.Subscription, error) {

	var minterRule []interface{}
	for _, minterItem := range minter {
		minterRule = append(minterRule, minterItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Mint", minterRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Mint)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xb4c03061fb5b7fed76389d5af8f2e0ddb09f8c70d1333abbb62582835e10accb.
//
// Solidity: event Mint(address indexed minter, uint256 tokenAmount, uint256 assetAmount, uint256 price)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseMint(log types.Log) (*LoanTokenLogicV4Mint, error) {
	event := new(LoanTokenLogicV4Mint)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4OwnershipTransferredIterator struct {
	Event *LoanTokenLogicV4OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4OwnershipTransferred)
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
		it.Event = new(LoanTokenLogicV4OwnershipTransferred)
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
func (it *LoanTokenLogicV4OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4OwnershipTransferred represents a OwnershipTransferred event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoanTokenLogicV4OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4OwnershipTransferredIterator{contract: _LoanTokenLogicV4.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4OwnershipTransferred)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseOwnershipTransferred(log types.Log) (*LoanTokenLogicV4OwnershipTransferred, error) {
	event := new(LoanTokenLogicV4OwnershipTransferred)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4RepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4RepayIterator struct {
	Event *LoanTokenLogicV4Repay // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4RepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Repay)
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
		it.Event = new(LoanTokenLogicV4Repay)
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
func (it *LoanTokenLogicV4RepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4RepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Repay represents a Repay event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Repay struct {
	LoanOrderHash [32]byte
	Borrower      common.Address
	Closer        common.Address
	Amount        *big.Int
	IsLiquidation bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterRepay(opts *bind.FilterOpts, loanOrderHash [][32]byte, borrower []common.Address) (*LoanTokenLogicV4RepayIterator, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4RepayIterator{contract: _LoanTokenLogicV4.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Repay, loanOrderHash [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Repay)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseRepay(log types.Log) (*LoanTokenLogicV4Repay, error) {
	event := new(LoanTokenLogicV4Repay)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenLogicV4TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4TransferIterator struct {
	Event *LoanTokenLogicV4Transfer // Event containing the contract specifics and raw log

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
func (it *LoanTokenLogicV4TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenLogicV4Transfer)
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
		it.Event = new(LoanTokenLogicV4Transfer)
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
func (it *LoanTokenLogicV4TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenLogicV4TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenLogicV4Transfer represents a Transfer event raised by the LoanTokenLogicV4 contract.
type LoanTokenLogicV4Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*LoanTokenLogicV4TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenLogicV4TransferIterator{contract: _LoanTokenLogicV4.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *LoanTokenLogicV4Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _LoanTokenLogicV4.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenLogicV4Transfer)
				if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_LoanTokenLogicV4 *LoanTokenLogicV4Filterer) ParseTransfer(log types.Log) (*LoanTokenLogicV4Transfer, error) {
	event := new(LoanTokenLogicV4Transfer)
	if err := _LoanTokenLogicV4.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenStorageABI is the input ABI used to generate the binding from.
const LoanTokenStorageABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"interestRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tradeTokenToFillAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"withdrawOnOpen\",\"type\":\"bool\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"claimant\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"tokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"assetAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"remainingTokenAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"}],\"name\":\"Claim\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"borrower\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"closer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"baseRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"burntTokenReserveList\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"lender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"burntTokenReserveListIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isSet\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"burntTokenReserved\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"checkpointSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"initialPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"leverageList\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"loanOrderData\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"leverageAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"marginPremiumAmount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"loanOrderHashes\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"rateMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"spreadMultiplier\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"tokenizedRegistry\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalAssetBorrow\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LoanTokenStorageFuncSigs maps the 4-byte function signature to its string representation.
var LoanTokenStorageFuncSigs = map[string]string{
	"995363d3": "bZxContract()",
	"96c7871b": "bZxOracle()",
	"894ca308": "bZxVault()",
	"1f68f20a": "baseRate()",
	"7866c6c1": "burntTokenReserveList(uint256)",
	"fbd9574d": "burntTokenReserveListIndex(address)",
	"0c4925fd": "burntTokenReserved()",
	"7b7933b4": "checkpointSupply()",
	"313ce567": "decimals()",
	"1d0806ae": "initialPrice()",
	"9b3a54d1": "leverageList(uint256)",
	"2515aacd": "loanOrderData(bytes32)",
	"fe056342": "loanOrderHashes(uint256)",
	"797bf385": "loanTokenAddress()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"330691ac": "rateMultiplier()",
	"d84d2a47": "spreadMultiplier()",
	"95d89b41": "symbol()",
	"736ee3d3": "tokenizedRegistry()",
	"20f6d07c": "totalAssetBorrow()",
	"f2fde38b": "transferOwnership(address)",
	"4780eac1": "wethContract()",
}

// LoanTokenStorageBin is the compiled bytecode used for deploying new contracts.
var LoanTokenStorageBin = "0x608060405260016000819055600a805460ff19169055670de0b6b3a7640000600b556801043561a882930000600c5580546001600160a01b031916331790556108238061004d6000396000f3fe608060405234801561001057600080fd5b506004361061014d5760003560e01c8063797bf385116100c3578063995363d31161007c578063995363d31461024f5780639b3a54d114610257578063d84d2a471461026a578063f2fde38b14610272578063fbd9574d14610287578063fe056342146102a85761014d565b8063797bf3851461021f5780637b7933b414610227578063894ca3081461022f5780638da5cb5b1461023757806395d89b411461023f57806396c7871b146102475761014d565b80632515aacd116101155780632515aacd1461019d578063313ce567146101c4578063330691ac146101d95780634780eac1146101e1578063736ee3d3146101f65780637866c6c1146101fe5761014d565b806306fdde03146101525780630c4925fd146101705780631d0806ae146101855780631f68f20a1461018d57806320f6d07c14610195575b600080fd5b61015a6102bb565b604051610167919061071d565b60405180910390f35b610178610346565b6040516101679190610698565b61017861034c565b610178610352565b610178610358565b6101b06101ab3660046105e8565b61035e565b6040516101679897969594939291906106a6565b6101cc6103aa565b6040516101679190610749565b6101786103b3565b6101e96103b9565b6040516101679190610668565b6101e96103c8565b61021161020c3660046105e8565b6103dc565b604051610167929190610676565b6101e9610411565b610178610420565b6101e9610426565b6101e9610435565b61015a610444565b6101e961049f565b6101e96104ae565b6101786102653660046105e8565b6104c2565b6101786104e0565b6102856102803660046105c2565b6104e6565b005b61029a6102953660046105c2565b610509565b60405161016792919061072e565b6101786102b63660046105e8565b610525565b6002805460408051602060018416156101000260001901909316849004601f8101849004840282018401909252818152929183018282801561033e5780601f106103135761010080835404028352916020019161033e565b820191906000526020600020905b81548152906001019060200180831161032157829003601f168201915b505050505081565b60135481565b60185481565b600b5481565b60155481565b600f60205260009081526040902080546001820154600283015460038401546004850154600586015460068701546007909701549596949593949293919290916001600160a01b031688565b60045460ff1681565b600c5481565b6007546001600160a01b031681565b600a5461010090046001600160a01b031681565b601181815481106103e957fe5b6000918252602090912060029091020180546001909101546001600160a01b03909116915082565b6008546001600160a01b031681565b60165481565b6005546001600160a01b031681565b6001546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f8101849004840282018401909252818152929183018282801561033e5780601f106103135761010080835404028352916020019161033e565b6006546001600160a01b031681565b60045461010090046001600160a01b031681565b601081815481106104cf57fe5b600091825260209091200154905081565b600d5481565b6001546001600160a01b031633146104fd57600080fd5b61050681610537565b50565b6012602052600090815260409020805460019091015460ff1682565b600e6020526000908152604090205481565b6001600160a01b03811661054a57600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b80356105b1816107c3565b92915050565b80356105b1816107d7565b6000602082840312156105d457600080fd5b60006105e084846105a6565b949350505050565b6000602082840312156105fa57600080fd5b60006105e084846105b7565b61060f81610764565b82525050565b61060f8161076f565b61060f81610774565b600061063282610757565b61063c818561075b565b935061064c818560208601610789565b610655816107b9565b9093019392505050565b61060f81610783565b602081016105b18284610606565b604081016106848285610606565b610691602083018461061e565b9392505050565b602081016105b1828461061e565b61010081016106b5828b61061e565b6106c2602083018a61061e565b6106cf604083018961061e565b6106dc606083018861061e565b6106e9608083018761061e565b6106f660a083018661061e565b61070360c083018561061e565b61071060e0830184610606565b9998505050505050505050565b602080825281016106918184610627565b6040810161073c828561061e565b6106916020830184610615565b602081016105b1828461065f565b5190565b90815260200190565b60006105b182610777565b151590565b90565b6001600160a01b031690565b60ff1690565b60005b838110156107a457818101518382015260200161078c565b838111156107b3576000848401525b50505050565b601f01601f191690565b6107cc81610764565b811461050657600080fd5b6107cc8161077456fea365627a7a723158200853a9e06e23585f5a283e3e2599feb6bd1f04fd12235d4251ef78a6e5191f8b6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployLoanTokenStorage deploys a new Ethereum contract, binding an instance of LoanTokenStorage to it.
func DeployLoanTokenStorage(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoanTokenStorage, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenStorageABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoanTokenStorageBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoanTokenStorage{LoanTokenStorageCaller: LoanTokenStorageCaller{contract: contract}, LoanTokenStorageTransactor: LoanTokenStorageTransactor{contract: contract}, LoanTokenStorageFilterer: LoanTokenStorageFilterer{contract: contract}}, nil
}

// LoanTokenStorage is an auto generated Go binding around an Ethereum contract.
type LoanTokenStorage struct {
	LoanTokenStorageCaller     // Read-only binding to the contract
	LoanTokenStorageTransactor // Write-only binding to the contract
	LoanTokenStorageFilterer   // Log filterer for contract events
}

// LoanTokenStorageCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanTokenStorageCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenStorageTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanTokenStorageTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenStorageFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanTokenStorageFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenStorageSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanTokenStorageSession struct {
	Contract     *LoanTokenStorage // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanTokenStorageCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanTokenStorageCallerSession struct {
	Contract *LoanTokenStorageCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LoanTokenStorageTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanTokenStorageTransactorSession struct {
	Contract     *LoanTokenStorageTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LoanTokenStorageRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanTokenStorageRaw struct {
	Contract *LoanTokenStorage // Generic contract binding to access the raw methods on
}

// LoanTokenStorageCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanTokenStorageCallerRaw struct {
	Contract *LoanTokenStorageCaller // Generic read-only contract binding to access the raw methods on
}

// LoanTokenStorageTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanTokenStorageTransactorRaw struct {
	Contract *LoanTokenStorageTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanTokenStorage creates a new instance of LoanTokenStorage, bound to a specific deployed contract.
func NewLoanTokenStorage(address common.Address, backend bind.ContractBackend) (*LoanTokenStorage, error) {
	contract, err := bindLoanTokenStorage(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorage{LoanTokenStorageCaller: LoanTokenStorageCaller{contract: contract}, LoanTokenStorageTransactor: LoanTokenStorageTransactor{contract: contract}, LoanTokenStorageFilterer: LoanTokenStorageFilterer{contract: contract}}, nil
}

// NewLoanTokenStorageCaller creates a new read-only instance of LoanTokenStorage, bound to a specific deployed contract.
func NewLoanTokenStorageCaller(address common.Address, caller bind.ContractCaller) (*LoanTokenStorageCaller, error) {
	contract, err := bindLoanTokenStorage(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageCaller{contract: contract}, nil
}

// NewLoanTokenStorageTransactor creates a new write-only instance of LoanTokenStorage, bound to a specific deployed contract.
func NewLoanTokenStorageTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanTokenStorageTransactor, error) {
	contract, err := bindLoanTokenStorage(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageTransactor{contract: contract}, nil
}

// NewLoanTokenStorageFilterer creates a new log filterer instance of LoanTokenStorage, bound to a specific deployed contract.
func NewLoanTokenStorageFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanTokenStorageFilterer, error) {
	contract, err := bindLoanTokenStorage(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageFilterer{contract: contract}, nil
}

// bindLoanTokenStorage binds a generic wrapper to an already deployed contract.
func bindLoanTokenStorage(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenStorageABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenStorage *LoanTokenStorageRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenStorage.Contract.LoanTokenStorageCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenStorage *LoanTokenStorageRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.LoanTokenStorageTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenStorage *LoanTokenStorageRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.LoanTokenStorageTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenStorage *LoanTokenStorageCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenStorage.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenStorage *LoanTokenStorageTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenStorage *LoanTokenStorageTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.contract.Transact(opts, method, params...)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) BZxContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "bZxContract")
	return *ret0, err
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) BZxContract() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxContract(&_LoanTokenStorage.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BZxContract() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxContract(&_LoanTokenStorage.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) BZxOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "bZxOracle")
	return *ret0, err
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) BZxOracle() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxOracle(&_LoanTokenStorage.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BZxOracle() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxOracle(&_LoanTokenStorage.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) BZxVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "bZxVault")
	return *ret0, err
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) BZxVault() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxVault(&_LoanTokenStorage.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BZxVault() (common.Address, error) {
	return _LoanTokenStorage.Contract.BZxVault(&_LoanTokenStorage.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) BaseRate(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "baseRate")
	return *ret0, err
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) BaseRate() (*big.Int, error) {
	return _LoanTokenStorage.Contract.BaseRate(&_LoanTokenStorage.CallOpts)
}

// BaseRate is a free data retrieval call binding the contract method 0x1f68f20a.
//
// Solidity: function baseRate() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BaseRate() (*big.Int, error) {
	return _LoanTokenStorage.Contract.BaseRate(&_LoanTokenStorage.CallOpts)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenStorage *LoanTokenStorageCaller) BurntTokenReserveList(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	ret := new(struct {
		Lender common.Address
		Amount *big.Int
	})
	out := ret
	err := _LoanTokenStorage.contract.Call(opts, out, "burntTokenReserveList", arg0)
	return *ret, err
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenStorage *LoanTokenStorageSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserveList(&_LoanTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveList is a free data retrieval call binding the contract method 0x7866c6c1.
//
// Solidity: function burntTokenReserveList(uint256 ) view returns(address lender, uint256 amount)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BurntTokenReserveList(arg0 *big.Int) (struct {
	Lender common.Address
	Amount *big.Int
}, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserveList(&_LoanTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenStorage *LoanTokenStorageCaller) BurntTokenReserveListIndex(opts *bind.CallOpts, arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	ret := new(struct {
		Index *big.Int
		IsSet bool
	})
	out := ret
	err := _LoanTokenStorage.contract.Call(opts, out, "burntTokenReserveListIndex", arg0)
	return *ret, err
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenStorage *LoanTokenStorageSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserveListIndex(&_LoanTokenStorage.CallOpts, arg0)
}

// BurntTokenReserveListIndex is a free data retrieval call binding the contract method 0xfbd9574d.
//
// Solidity: function burntTokenReserveListIndex(address ) view returns(uint256 index, bool isSet)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BurntTokenReserveListIndex(arg0 common.Address) (struct {
	Index *big.Int
	IsSet bool
}, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserveListIndex(&_LoanTokenStorage.CallOpts, arg0)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) BurntTokenReserved(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "burntTokenReserved")
	return *ret0, err
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) BurntTokenReserved() (*big.Int, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserved(&_LoanTokenStorage.CallOpts)
}

// BurntTokenReserved is a free data retrieval call binding the contract method 0x0c4925fd.
//
// Solidity: function burntTokenReserved() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) BurntTokenReserved() (*big.Int, error) {
	return _LoanTokenStorage.Contract.BurntTokenReserved(&_LoanTokenStorage.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) CheckpointSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "checkpointSupply")
	return *ret0, err
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) CheckpointSupply() (*big.Int, error) {
	return _LoanTokenStorage.Contract.CheckpointSupply(&_LoanTokenStorage.CallOpts)
}

// CheckpointSupply is a free data retrieval call binding the contract method 0x7b7933b4.
//
// Solidity: function checkpointSupply() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) CheckpointSupply() (*big.Int, error) {
	return _LoanTokenStorage.Contract.CheckpointSupply(&_LoanTokenStorage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenStorage *LoanTokenStorageCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenStorage *LoanTokenStorageSession) Decimals() (uint8, error) {
	return _LoanTokenStorage.Contract.Decimals(&_LoanTokenStorage.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) Decimals() (uint8, error) {
	return _LoanTokenStorage.Contract.Decimals(&_LoanTokenStorage.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) InitialPrice(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "initialPrice")
	return *ret0, err
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) InitialPrice() (*big.Int, error) {
	return _LoanTokenStorage.Contract.InitialPrice(&_LoanTokenStorage.CallOpts)
}

// InitialPrice is a free data retrieval call binding the contract method 0x1d0806ae.
//
// Solidity: function initialPrice() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) InitialPrice() (*big.Int, error) {
	return _LoanTokenStorage.Contract.InitialPrice(&_LoanTokenStorage.CallOpts)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) LeverageList(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "leverageList", arg0)
	return *ret0, err
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _LoanTokenStorage.Contract.LeverageList(&_LoanTokenStorage.CallOpts, arg0)
}

// LeverageList is a free data retrieval call binding the contract method 0x9b3a54d1.
//
// Solidity: function leverageList(uint256 ) view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) LeverageList(arg0 *big.Int) (*big.Int, error) {
	return _LoanTokenStorage.Contract.LeverageList(&_LoanTokenStorage.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenStorage *LoanTokenStorageCaller) LoanOrderData(opts *bind.CallOpts, arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	ret := new(struct {
		LoanOrderHash               [32]byte
		LeverageAmount              *big.Int
		InitialMarginAmount         *big.Int
		MaintenanceMarginAmount     *big.Int
		MaxDurationUnixTimestampSec *big.Int
		Index                       *big.Int
		MarginPremiumAmount         *big.Int
		CollateralTokenAddress      common.Address
	})
	out := ret
	err := _LoanTokenStorage.contract.Call(opts, out, "loanOrderData", arg0)
	return *ret, err
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenStorage *LoanTokenStorageSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _LoanTokenStorage.Contract.LoanOrderData(&_LoanTokenStorage.CallOpts, arg0)
}

// LoanOrderData is a free data retrieval call binding the contract method 0x2515aacd.
//
// Solidity: function loanOrderData(bytes32 ) view returns(bytes32 loanOrderHash, uint256 leverageAmount, uint256 initialMarginAmount, uint256 maintenanceMarginAmount, uint256 maxDurationUnixTimestampSec, uint256 index, uint256 marginPremiumAmount, address collateralTokenAddress)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) LoanOrderData(arg0 [32]byte) (struct {
	LoanOrderHash               [32]byte
	LeverageAmount              *big.Int
	InitialMarginAmount         *big.Int
	MaintenanceMarginAmount     *big.Int
	MaxDurationUnixTimestampSec *big.Int
	Index                       *big.Int
	MarginPremiumAmount         *big.Int
	CollateralTokenAddress      common.Address
}, error) {
	return _LoanTokenStorage.Contract.LoanOrderData(&_LoanTokenStorage.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenStorage *LoanTokenStorageCaller) LoanOrderHashes(opts *bind.CallOpts, arg0 *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "loanOrderHashes", arg0)
	return *ret0, err
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenStorage *LoanTokenStorageSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _LoanTokenStorage.Contract.LoanOrderHashes(&_LoanTokenStorage.CallOpts, arg0)
}

// LoanOrderHashes is a free data retrieval call binding the contract method 0xfe056342.
//
// Solidity: function loanOrderHashes(uint256 ) view returns(bytes32)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) LoanOrderHashes(arg0 *big.Int) ([32]byte, error) {
	return _LoanTokenStorage.Contract.LoanOrderHashes(&_LoanTokenStorage.CallOpts, arg0)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) LoanTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "loanTokenAddress")
	return *ret0, err
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenStorage.Contract.LoanTokenAddress(&_LoanTokenStorage.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenStorage.Contract.LoanTokenAddress(&_LoanTokenStorage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageSession) Name() (string, error) {
	return _LoanTokenStorage.Contract.Name(&_LoanTokenStorage.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) Name() (string, error) {
	return _LoanTokenStorage.Contract.Name(&_LoanTokenStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) Owner() (common.Address, error) {
	return _LoanTokenStorage.Contract.Owner(&_LoanTokenStorage.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) Owner() (common.Address, error) {
	return _LoanTokenStorage.Contract.Owner(&_LoanTokenStorage.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) RateMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "rateMultiplier")
	return *ret0, err
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) RateMultiplier() (*big.Int, error) {
	return _LoanTokenStorage.Contract.RateMultiplier(&_LoanTokenStorage.CallOpts)
}

// RateMultiplier is a free data retrieval call binding the contract method 0x330691ac.
//
// Solidity: function rateMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) RateMultiplier() (*big.Int, error) {
	return _LoanTokenStorage.Contract.RateMultiplier(&_LoanTokenStorage.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) SpreadMultiplier(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "spreadMultiplier")
	return *ret0, err
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) SpreadMultiplier() (*big.Int, error) {
	return _LoanTokenStorage.Contract.SpreadMultiplier(&_LoanTokenStorage.CallOpts)
}

// SpreadMultiplier is a free data retrieval call binding the contract method 0xd84d2a47.
//
// Solidity: function spreadMultiplier() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) SpreadMultiplier() (*big.Int, error) {
	return _LoanTokenStorage.Contract.SpreadMultiplier(&_LoanTokenStorage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageSession) Symbol() (string, error) {
	return _LoanTokenStorage.Contract.Symbol(&_LoanTokenStorage.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) Symbol() (string, error) {
	return _LoanTokenStorage.Contract.Symbol(&_LoanTokenStorage.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) TokenizedRegistry(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "tokenizedRegistry")
	return *ret0, err
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) TokenizedRegistry() (common.Address, error) {
	return _LoanTokenStorage.Contract.TokenizedRegistry(&_LoanTokenStorage.CallOpts)
}

// TokenizedRegistry is a free data retrieval call binding the contract method 0x736ee3d3.
//
// Solidity: function tokenizedRegistry() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) TokenizedRegistry() (common.Address, error) {
	return _LoanTokenStorage.Contract.TokenizedRegistry(&_LoanTokenStorage.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCaller) TotalAssetBorrow(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "totalAssetBorrow")
	return *ret0, err
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageSession) TotalAssetBorrow() (*big.Int, error) {
	return _LoanTokenStorage.Contract.TotalAssetBorrow(&_LoanTokenStorage.CallOpts)
}

// TotalAssetBorrow is a free data retrieval call binding the contract method 0x20f6d07c.
//
// Solidity: function totalAssetBorrow() view returns(uint256)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) TotalAssetBorrow() (*big.Int, error) {
	return _LoanTokenStorage.Contract.TotalAssetBorrow(&_LoanTokenStorage.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCaller) WethContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenStorage.contract.Call(opts, out, "wethContract")
	return *ret0, err
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageSession) WethContract() (common.Address, error) {
	return _LoanTokenStorage.Contract.WethContract(&_LoanTokenStorage.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenStorage *LoanTokenStorageCallerSession) WethContract() (common.Address, error) {
	return _LoanTokenStorage.Contract.WethContract(&_LoanTokenStorage.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenStorage *LoanTokenStorageTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenStorage.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenStorage *LoanTokenStorageSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.TransferOwnership(&_LoanTokenStorage.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenStorage *LoanTokenStorageTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenStorage.Contract.TransferOwnership(&_LoanTokenStorage.TransactOpts, _newOwner)
}

// LoanTokenStorageBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the LoanTokenStorage contract.
type LoanTokenStorageBorrowIterator struct {
	Event *LoanTokenStorageBorrow // Event containing the contract specifics and raw log

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
func (it *LoanTokenStorageBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenStorageBorrow)
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
		it.Event = new(LoanTokenStorageBorrow)
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
func (it *LoanTokenStorageBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenStorageBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenStorageBorrow represents a Borrow event raised by the LoanTokenStorage contract.
type LoanTokenStorageBorrow struct {
	Borrower                common.Address
	BorrowAmount            *big.Int
	InterestRate            *big.Int
	CollateralTokenAddress  common.Address
	TradeTokenToFillAddress common.Address
	WithdrawOnOpen          bool
	Raw                     types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenStorage *LoanTokenStorageFilterer) FilterBorrow(opts *bind.FilterOpts, borrower []common.Address) (*LoanTokenStorageBorrowIterator, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.FilterLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageBorrowIterator{contract: _LoanTokenStorage.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenStorage *LoanTokenStorageFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *LoanTokenStorageBorrow, borrower []common.Address) (event.Subscription, error) {

	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.WatchLogs(opts, "Borrow", borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenStorageBorrow)
				if err := _LoanTokenStorage.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0x86e15dd78cd784ab7788bcf5b96b9395e86030e048e5faedcfe752c700f6157e.
//
// Solidity: event Borrow(address indexed borrower, uint256 borrowAmount, uint256 interestRate, address collateralTokenAddress, address tradeTokenToFillAddress, bool withdrawOnOpen)
func (_LoanTokenStorage *LoanTokenStorageFilterer) ParseBorrow(log types.Log) (*LoanTokenStorageBorrow, error) {
	event := new(LoanTokenStorageBorrow)
	if err := _LoanTokenStorage.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenStorageClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the LoanTokenStorage contract.
type LoanTokenStorageClaimIterator struct {
	Event *LoanTokenStorageClaim // Event containing the contract specifics and raw log

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
func (it *LoanTokenStorageClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenStorageClaim)
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
		it.Event = new(LoanTokenStorageClaim)
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
func (it *LoanTokenStorageClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenStorageClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenStorageClaim represents a Claim event raised by the LoanTokenStorage contract.
type LoanTokenStorageClaim struct {
	Claimant             common.Address
	TokenAmount          *big.Int
	AssetAmount          *big.Int
	RemainingTokenAmount *big.Int
	Price                *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenStorage *LoanTokenStorageFilterer) FilterClaim(opts *bind.FilterOpts, claimant []common.Address) (*LoanTokenStorageClaimIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.FilterLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageClaimIterator{contract: _LoanTokenStorage.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenStorage *LoanTokenStorageFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *LoanTokenStorageClaim, claimant []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.WatchLogs(opts, "Claim", claimantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenStorageClaim)
				if err := _LoanTokenStorage.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x68e1caf97c4c29c1ac46024e9590f80b7a1f690d393703879cf66eea4e1e8421.
//
// Solidity: event Claim(address indexed claimant, uint256 tokenAmount, uint256 assetAmount, uint256 remainingTokenAmount, uint256 price)
func (_LoanTokenStorage *LoanTokenStorageFilterer) ParseClaim(log types.Log) (*LoanTokenStorageClaim, error) {
	event := new(LoanTokenStorageClaim)
	if err := _LoanTokenStorage.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenStorageOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoanTokenStorage contract.
type LoanTokenStorageOwnershipTransferredIterator struct {
	Event *LoanTokenStorageOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoanTokenStorageOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenStorageOwnershipTransferred)
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
		it.Event = new(LoanTokenStorageOwnershipTransferred)
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
func (it *LoanTokenStorageOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenStorageOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenStorageOwnershipTransferred represents a OwnershipTransferred event raised by the LoanTokenStorage contract.
type LoanTokenStorageOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenStorage *LoanTokenStorageFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoanTokenStorageOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageOwnershipTransferredIterator{contract: _LoanTokenStorage.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenStorage *LoanTokenStorageFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoanTokenStorageOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenStorageOwnershipTransferred)
				if err := _LoanTokenStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LoanTokenStorage *LoanTokenStorageFilterer) ParseOwnershipTransferred(log types.Log) (*LoanTokenStorageOwnershipTransferred, error) {
	event := new(LoanTokenStorageOwnershipTransferred)
	if err := _LoanTokenStorage.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenStorageRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the LoanTokenStorage contract.
type LoanTokenStorageRepayIterator struct {
	Event *LoanTokenStorageRepay // Event containing the contract specifics and raw log

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
func (it *LoanTokenStorageRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenStorageRepay)
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
		it.Event = new(LoanTokenStorageRepay)
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
func (it *LoanTokenStorageRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenStorageRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenStorageRepay represents a Repay event raised by the LoanTokenStorage contract.
type LoanTokenStorageRepay struct {
	LoanOrderHash [32]byte
	Borrower      common.Address
	Closer        common.Address
	Amount        *big.Int
	IsLiquidation bool
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenStorage *LoanTokenStorageFilterer) FilterRepay(opts *bind.FilterOpts, loanOrderHash [][32]byte, borrower []common.Address) (*LoanTokenStorageRepayIterator, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.FilterLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenStorageRepayIterator{contract: _LoanTokenStorage.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenStorage *LoanTokenStorageFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *LoanTokenStorageRepay, loanOrderHash [][32]byte, borrower []common.Address) (event.Subscription, error) {

	var loanOrderHashRule []interface{}
	for _, loanOrderHashItem := range loanOrderHash {
		loanOrderHashRule = append(loanOrderHashRule, loanOrderHashItem)
	}
	var borrowerRule []interface{}
	for _, borrowerItem := range borrower {
		borrowerRule = append(borrowerRule, borrowerItem)
	}

	logs, sub, err := _LoanTokenStorage.contract.WatchLogs(opts, "Repay", loanOrderHashRule, borrowerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenStorageRepay)
				if err := _LoanTokenStorage.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x85dfc0033a3e5b3b9b3151bd779c1f9b855d66b83ff5bb79283b68d82e8e5b73.
//
// Solidity: event Repay(bytes32 indexed loanOrderHash, address indexed borrower, address closer, uint256 amount, bool isLiquidation)
func (_LoanTokenStorage *LoanTokenStorageFilterer) ParseRepay(log types.Log) (*LoanTokenStorageRepay, error) {
	event := new(LoanTokenStorageRepay)
	if err := _LoanTokenStorage.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	return event, nil
}

// LoanTokenizationABI is the input ABI used to generate the binding from.
const LoanTokenizationABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"bZxVault\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"loanTokenAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"wethContract\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// LoanTokenizationFuncSigs maps the 4-byte function signature to its string representation.
var LoanTokenizationFuncSigs = map[string]string{
	"995363d3": "bZxContract()",
	"96c7871b": "bZxOracle()",
	"894ca308": "bZxVault()",
	"313ce567": "decimals()",
	"797bf385": "loanTokenAddress()",
	"06fdde03": "name()",
	"8da5cb5b": "owner()",
	"95d89b41": "symbol()",
	"f2fde38b": "transferOwnership(address)",
	"4780eac1": "wethContract()",
}

// LoanTokenizationBin is the compiled bytecode used for deploying new contracts.
var LoanTokenizationBin = "0x60806040526001600081905580546001600160a01b031916331790556104808061002a6000396000f3fe608060405234801561001057600080fd5b506004361061009e5760003560e01c80638da5cb5b116100665780638da5cb5b146100fb57806395d89b411461010357806396c7871b1461010b578063995363d314610113578063f2fde38b1461011b5761009e565b806306fdde03146100a3578063313ce567146100c15780634780eac1146100d6578063797bf385146100eb578063894ca308146100f3575b600080fd5b6100ab610130565b6040516100b891906103a5565b60405180910390f35b6100c96101bb565b6040516100b891906103bd565b6100de6101c4565b6040516100b89190610397565b6100de6101d3565b6100de6101e2565b6100de6101f1565b6100ab610200565b6100de61025b565b6100de61026a565b61012e610129366004610321565b61027e565b005b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156101b35780601f10610188576101008083540402835291602001916101b3565b820191906000526020600020905b81548152906001019060200180831161019657829003601f168201915b505050505081565b60045460ff1681565b6007546001600160a01b031681565b6008546001600160a01b031681565b6005546001600160a01b031681565b6001546001600160a01b031681565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156101b35780601f10610188576101008083540402835291602001916101b3565b6006546001600160a01b031681565b60045461010090046001600160a01b031681565b6001546001600160a01b0316331461029557600080fd5b61029e816102a1565b50565b6001600160a01b0381166102b457600080fd5b6001546040516001600160a01b038084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3600180546001600160a01b0319166001600160a01b0392909216919091179055565b803561031b81610429565b92915050565b60006020828403121561033357600080fd5b600061033f8484610310565b949350505050565b610350816103d8565b82525050565b6000610361826103cb565b61036b81856103cf565b935061037b8185602086016103ef565b6103848161041f565b9093019392505050565b610350816103e9565b6020810161031b8284610347565b602080825281016103b68184610356565b9392505050565b6020810161031b828461038e565b5190565b90815260200190565b60006001600160a01b03821661031b565b60ff1690565b60005b8381101561040a5781810151838201526020016103f2565b83811115610419576000848401525b50505050565b601f01601f191690565b610432816103d8565b811461029e57600080fdfea365627a7a723158206e662627937e4502d07d9118632b14ffbad6272f0c78cb295ddb42760a4fd9e46c6578706572696d656e74616cf564736f6c63430005110040"

// DeployLoanTokenization deploys a new Ethereum contract, binding an instance of LoanTokenization to it.
func DeployLoanTokenization(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *LoanTokenization, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenizationABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(LoanTokenizationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &LoanTokenization{LoanTokenizationCaller: LoanTokenizationCaller{contract: contract}, LoanTokenizationTransactor: LoanTokenizationTransactor{contract: contract}, LoanTokenizationFilterer: LoanTokenizationFilterer{contract: contract}}, nil
}

// LoanTokenization is an auto generated Go binding around an Ethereum contract.
type LoanTokenization struct {
	LoanTokenizationCaller     // Read-only binding to the contract
	LoanTokenizationTransactor // Write-only binding to the contract
	LoanTokenizationFilterer   // Log filterer for contract events
}

// LoanTokenizationCaller is an auto generated read-only Go binding around an Ethereum contract.
type LoanTokenizationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenizationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type LoanTokenizationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenizationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type LoanTokenizationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// LoanTokenizationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type LoanTokenizationSession struct {
	Contract     *LoanTokenization // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// LoanTokenizationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type LoanTokenizationCallerSession struct {
	Contract *LoanTokenizationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// LoanTokenizationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type LoanTokenizationTransactorSession struct {
	Contract     *LoanTokenizationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// LoanTokenizationRaw is an auto generated low-level Go binding around an Ethereum contract.
type LoanTokenizationRaw struct {
	Contract *LoanTokenization // Generic contract binding to access the raw methods on
}

// LoanTokenizationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type LoanTokenizationCallerRaw struct {
	Contract *LoanTokenizationCaller // Generic read-only contract binding to access the raw methods on
}

// LoanTokenizationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type LoanTokenizationTransactorRaw struct {
	Contract *LoanTokenizationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewLoanTokenization creates a new instance of LoanTokenization, bound to a specific deployed contract.
func NewLoanTokenization(address common.Address, backend bind.ContractBackend) (*LoanTokenization, error) {
	contract, err := bindLoanTokenization(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &LoanTokenization{LoanTokenizationCaller: LoanTokenizationCaller{contract: contract}, LoanTokenizationTransactor: LoanTokenizationTransactor{contract: contract}, LoanTokenizationFilterer: LoanTokenizationFilterer{contract: contract}}, nil
}

// NewLoanTokenizationCaller creates a new read-only instance of LoanTokenization, bound to a specific deployed contract.
func NewLoanTokenizationCaller(address common.Address, caller bind.ContractCaller) (*LoanTokenizationCaller, error) {
	contract, err := bindLoanTokenization(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenizationCaller{contract: contract}, nil
}

// NewLoanTokenizationTransactor creates a new write-only instance of LoanTokenization, bound to a specific deployed contract.
func NewLoanTokenizationTransactor(address common.Address, transactor bind.ContractTransactor) (*LoanTokenizationTransactor, error) {
	contract, err := bindLoanTokenization(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &LoanTokenizationTransactor{contract: contract}, nil
}

// NewLoanTokenizationFilterer creates a new log filterer instance of LoanTokenization, bound to a specific deployed contract.
func NewLoanTokenizationFilterer(address common.Address, filterer bind.ContractFilterer) (*LoanTokenizationFilterer, error) {
	contract, err := bindLoanTokenization(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &LoanTokenizationFilterer{contract: contract}, nil
}

// bindLoanTokenization binds a generic wrapper to an already deployed contract.
func bindLoanTokenization(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(LoanTokenizationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenization *LoanTokenizationRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenization.Contract.LoanTokenizationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenization *LoanTokenizationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenization.Contract.LoanTokenizationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenization *LoanTokenizationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenization.Contract.LoanTokenizationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_LoanTokenization *LoanTokenizationCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _LoanTokenization.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_LoanTokenization *LoanTokenizationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _LoanTokenization.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_LoanTokenization *LoanTokenizationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _LoanTokenization.Contract.contract.Transact(opts, method, params...)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) BZxContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "bZxContract")
	return *ret0, err
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) BZxContract() (common.Address, error) {
	return _LoanTokenization.Contract.BZxContract(&_LoanTokenization.CallOpts)
}

// BZxContract is a free data retrieval call binding the contract method 0x995363d3.
//
// Solidity: function bZxContract() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) BZxContract() (common.Address, error) {
	return _LoanTokenization.Contract.BZxContract(&_LoanTokenization.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) BZxOracle(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "bZxOracle")
	return *ret0, err
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) BZxOracle() (common.Address, error) {
	return _LoanTokenization.Contract.BZxOracle(&_LoanTokenization.CallOpts)
}

// BZxOracle is a free data retrieval call binding the contract method 0x96c7871b.
//
// Solidity: function bZxOracle() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) BZxOracle() (common.Address, error) {
	return _LoanTokenization.Contract.BZxOracle(&_LoanTokenization.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) BZxVault(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "bZxVault")
	return *ret0, err
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) BZxVault() (common.Address, error) {
	return _LoanTokenization.Contract.BZxVault(&_LoanTokenization.CallOpts)
}

// BZxVault is a free data retrieval call binding the contract method 0x894ca308.
//
// Solidity: function bZxVault() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) BZxVault() (common.Address, error) {
	return _LoanTokenization.Contract.BZxVault(&_LoanTokenization.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenization *LoanTokenizationCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenization *LoanTokenizationSession) Decimals() (uint8, error) {
	return _LoanTokenization.Contract.Decimals(&_LoanTokenization.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_LoanTokenization *LoanTokenizationCallerSession) Decimals() (uint8, error) {
	return _LoanTokenization.Contract.Decimals(&_LoanTokenization.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) LoanTokenAddress(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "loanTokenAddress")
	return *ret0, err
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenization.Contract.LoanTokenAddress(&_LoanTokenization.CallOpts)
}

// LoanTokenAddress is a free data retrieval call binding the contract method 0x797bf385.
//
// Solidity: function loanTokenAddress() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) LoanTokenAddress() (common.Address, error) {
	return _LoanTokenization.Contract.LoanTokenAddress(&_LoanTokenization.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenization *LoanTokenizationCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenization *LoanTokenizationSession) Name() (string, error) {
	return _LoanTokenization.Contract.Name(&_LoanTokenization.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_LoanTokenization *LoanTokenizationCallerSession) Name() (string, error) {
	return _LoanTokenization.Contract.Name(&_LoanTokenization.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) Owner() (common.Address, error) {
	return _LoanTokenization.Contract.Owner(&_LoanTokenization.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) Owner() (common.Address, error) {
	return _LoanTokenization.Contract.Owner(&_LoanTokenization.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenization *LoanTokenizationCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenization *LoanTokenizationSession) Symbol() (string, error) {
	return _LoanTokenization.Contract.Symbol(&_LoanTokenization.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_LoanTokenization *LoanTokenizationCallerSession) Symbol() (string, error) {
	return _LoanTokenization.Contract.Symbol(&_LoanTokenization.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenization *LoanTokenizationCaller) WethContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _LoanTokenization.contract.Call(opts, out, "wethContract")
	return *ret0, err
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenization *LoanTokenizationSession) WethContract() (common.Address, error) {
	return _LoanTokenization.Contract.WethContract(&_LoanTokenization.CallOpts)
}

// WethContract is a free data retrieval call binding the contract method 0x4780eac1.
//
// Solidity: function wethContract() view returns(address)
func (_LoanTokenization *LoanTokenizationCallerSession) WethContract() (common.Address, error) {
	return _LoanTokenization.Contract.WethContract(&_LoanTokenization.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenization *LoanTokenizationTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenization.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenization *LoanTokenizationSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenization.Contract.TransferOwnership(&_LoanTokenization.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_LoanTokenization *LoanTokenizationTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _LoanTokenization.Contract.TransferOwnership(&_LoanTokenization.TransactOpts, _newOwner)
}

// LoanTokenizationOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the LoanTokenization contract.
type LoanTokenizationOwnershipTransferredIterator struct {
	Event *LoanTokenizationOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *LoanTokenizationOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(LoanTokenizationOwnershipTransferred)
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
		it.Event = new(LoanTokenizationOwnershipTransferred)
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
func (it *LoanTokenizationOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *LoanTokenizationOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// LoanTokenizationOwnershipTransferred represents a OwnershipTransferred event raised by the LoanTokenization contract.
type LoanTokenizationOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenization *LoanTokenizationFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*LoanTokenizationOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenization.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &LoanTokenizationOwnershipTransferredIterator{contract: _LoanTokenization.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_LoanTokenization *LoanTokenizationFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *LoanTokenizationOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _LoanTokenization.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(LoanTokenizationOwnershipTransferred)
				if err := _LoanTokenization.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_LoanTokenization *LoanTokenizationFilterer) ParseOwnershipTransferred(log types.Log) (*LoanTokenizationOwnershipTransferred, error) {
	event := new(LoanTokenizationOwnershipTransferred)
	if err := _LoanTokenization.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// OracleNotifierInterfaceABI is the input ABI used to generate the binding from.
const OracleNotifierInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"loanTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"oracleAddress\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"initialMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maintenanceMarginAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"maxDurationUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"bytes32\",\"name\":\"loanOrderHash\",\"type\":\"bytes32\"}],\"internalType\":\"structBZxObjects.LoanOrder\",\"name\":\"loanOrder\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"trader\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"collateralTokenAddressFilled\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"positionTokenAddressFilled\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanTokenAmountUsed\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"collateralTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"positionTokenAmountFilled\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanStartUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"loanEndUnixTimestampSec\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"active\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"positionId\",\"type\":\"uint256\"}],\"internalType\":\"structBZxObjects.LoanPosition\",\"name\":\"loanPosition\",\"type\":\"tuple\"},{\"internalType\":\"address\",\"name\":\"loanCloser\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"closeAmount\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"isLiquidation\",\"type\":\"bool\"}],\"name\":\"closeLoanNotifier\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OracleNotifierInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var OracleNotifierInterfaceFuncSigs = map[string]string{
	"cd4fa66d": "closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32),(address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256),address,uint256,bool)",
}

// OracleNotifierInterface is an auto generated Go binding around an Ethereum contract.
type OracleNotifierInterface struct {
	OracleNotifierInterfaceCaller     // Read-only binding to the contract
	OracleNotifierInterfaceTransactor // Write-only binding to the contract
	OracleNotifierInterfaceFilterer   // Log filterer for contract events
}

// OracleNotifierInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type OracleNotifierInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleNotifierInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OracleNotifierInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleNotifierInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OracleNotifierInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OracleNotifierInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OracleNotifierInterfaceSession struct {
	Contract     *OracleNotifierInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts            // Call options to use throughout this session
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// OracleNotifierInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OracleNotifierInterfaceCallerSession struct {
	Contract *OracleNotifierInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                  // Call options to use throughout this session
}

// OracleNotifierInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OracleNotifierInterfaceTransactorSession struct {
	Contract     *OracleNotifierInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                  // Transaction auth options to use throughout this session
}

// OracleNotifierInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type OracleNotifierInterfaceRaw struct {
	Contract *OracleNotifierInterface // Generic contract binding to access the raw methods on
}

// OracleNotifierInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OracleNotifierInterfaceCallerRaw struct {
	Contract *OracleNotifierInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// OracleNotifierInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OracleNotifierInterfaceTransactorRaw struct {
	Contract *OracleNotifierInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOracleNotifierInterface creates a new instance of OracleNotifierInterface, bound to a specific deployed contract.
func NewOracleNotifierInterface(address common.Address, backend bind.ContractBackend) (*OracleNotifierInterface, error) {
	contract, err := bindOracleNotifierInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OracleNotifierInterface{OracleNotifierInterfaceCaller: OracleNotifierInterfaceCaller{contract: contract}, OracleNotifierInterfaceTransactor: OracleNotifierInterfaceTransactor{contract: contract}, OracleNotifierInterfaceFilterer: OracleNotifierInterfaceFilterer{contract: contract}}, nil
}

// NewOracleNotifierInterfaceCaller creates a new read-only instance of OracleNotifierInterface, bound to a specific deployed contract.
func NewOracleNotifierInterfaceCaller(address common.Address, caller bind.ContractCaller) (*OracleNotifierInterfaceCaller, error) {
	contract, err := bindOracleNotifierInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OracleNotifierInterfaceCaller{contract: contract}, nil
}

// NewOracleNotifierInterfaceTransactor creates a new write-only instance of OracleNotifierInterface, bound to a specific deployed contract.
func NewOracleNotifierInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*OracleNotifierInterfaceTransactor, error) {
	contract, err := bindOracleNotifierInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OracleNotifierInterfaceTransactor{contract: contract}, nil
}

// NewOracleNotifierInterfaceFilterer creates a new log filterer instance of OracleNotifierInterface, bound to a specific deployed contract.
func NewOracleNotifierInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*OracleNotifierInterfaceFilterer, error) {
	contract, err := bindOracleNotifierInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OracleNotifierInterfaceFilterer{contract: contract}, nil
}

// bindOracleNotifierInterface binds a generic wrapper to an already deployed contract.
func bindOracleNotifierInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OracleNotifierInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleNotifierInterface *OracleNotifierInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleNotifierInterface.Contract.OracleNotifierInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleNotifierInterface *OracleNotifierInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.OracleNotifierInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleNotifierInterface *OracleNotifierInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.OracleNotifierInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OracleNotifierInterface *OracleNotifierInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _OracleNotifierInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OracleNotifierInterface *OracleNotifierInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OracleNotifierInterface *OracleNotifierInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.contract.Transact(opts, method, params...)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_OracleNotifierInterface *OracleNotifierInterfaceTransactor) CloseLoanNotifier(opts *bind.TransactOpts, loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _OracleNotifierInterface.contract.Transact(opts, "closeLoanNotifier", loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_OracleNotifierInterface *OracleNotifierInterfaceSession) CloseLoanNotifier(loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.CloseLoanNotifier(&_OracleNotifierInterface.TransactOpts, loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// CloseLoanNotifier is a paid mutator transaction binding the contract method 0xcd4fa66d.
//
// Solidity: function closeLoanNotifier((address,address,address,address,uint256,uint256,uint256,uint256,uint256,bytes32) loanOrder, (address,address,address,uint256,uint256,uint256,uint256,uint256,uint256,bool,uint256) loanPosition, address loanCloser, uint256 closeAmount, bool isLiquidation) returns(bool)
func (_OracleNotifierInterface *OracleNotifierInterfaceTransactorSession) CloseLoanNotifier(loanOrder BZxObjectsLoanOrder, loanPosition BZxObjectsLoanPosition, loanCloser common.Address, closeAmount *big.Int, isLiquidation bool) (*types.Transaction, error) {
	return _OracleNotifierInterface.Contract.CloseLoanNotifier(&_OracleNotifierInterface.TransactOpts, loanOrder, loanPosition, loanCloser, closeAmount, isLiquidation)
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x608060405234801561001057600080fd5b50600080546001600160a01b031916331790556101ca806100326000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c80638da5cb5b1461003b578063f2fde38b14610059575b600080fd5b61004361006e565b6040516100509190610154565b60405180910390f35b61006c61006736600461011f565b61007d565b005b6000546001600160a01b031681565b6000546001600160a01b0316331461009457600080fd5b61009d816100a0565b50565b6001600160a01b0381166100b357600080fd5b600080546040516001600160a01b03808516939216917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e091a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b803561011981610173565b92915050565b60006020828403121561013157600080fd5b600061013d848461010e565b949350505050565b61014e81610162565b82525050565b602081016101198284610145565b60006001600160a01b038216610119565b61017c81610162565b811461009d57600080fdfea365627a7a72315820712c1ccf6b5dbc9f1633f3d4781e1a00a4ae551c4be3fb57300722c76cfec4166c6578706572696d656e74616cf564736f6c63430005110040"

// DeployOwnable deploys a new Ethereum contract, binding an instance of Ownable to it.
func DeployOwnable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Ownable, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(OwnableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// Ownable is an auto generated Go binding around an Ethereum contract.
type Ownable struct {
	OwnableCaller     // Read-only binding to the contract
	OwnableTransactor // Write-only binding to the contract
	OwnableFilterer   // Log filterer for contract events
}

// OwnableCaller is an auto generated read-only Go binding around an Ethereum contract.
type OwnableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OwnableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OwnableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OwnableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OwnableSession struct {
	Contract     *Ownable          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OwnableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OwnableCallerSession struct {
	Contract *OwnableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// OwnableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OwnableTransactorSession struct {
	Contract     *OwnableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// OwnableRaw is an auto generated low-level Go binding around an Ethereum contract.
type OwnableRaw struct {
	Contract *Ownable // Generic contract binding to access the raw methods on
}

// OwnableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OwnableCallerRaw struct {
	Contract *OwnableCaller // Generic read-only contract binding to access the raw methods on
}

// OwnableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OwnableTransactorRaw struct {
	Contract *OwnableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOwnable creates a new instance of Ownable, bound to a specific deployed contract.
func NewOwnable(address common.Address, backend bind.ContractBackend) (*Ownable, error) {
	contract, err := bindOwnable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Ownable{OwnableCaller: OwnableCaller{contract: contract}, OwnableTransactor: OwnableTransactor{contract: contract}, OwnableFilterer: OwnableFilterer{contract: contract}}, nil
}

// NewOwnableCaller creates a new read-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableCaller(address common.Address, caller bind.ContractCaller) (*OwnableCaller, error) {
	contract, err := bindOwnable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableCaller{contract: contract}, nil
}

// NewOwnableTransactor creates a new write-only instance of Ownable, bound to a specific deployed contract.
func NewOwnableTransactor(address common.Address, transactor bind.ContractTransactor) (*OwnableTransactor, error) {
	contract, err := bindOwnable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OwnableTransactor{contract: contract}, nil
}

// NewOwnableFilterer creates a new log filterer instance of Ownable, bound to a specific deployed contract.
func NewOwnableFilterer(address common.Address, filterer bind.ContractFilterer) (*OwnableFilterer, error) {
	contract, err := bindOwnable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OwnableFilterer{contract: contract}, nil
}

// bindOwnable binds a generic wrapper to an already deployed contract.
func bindOwnable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(OwnableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.OwnableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.OwnableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Ownable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Ownable *OwnableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Ownable *OwnableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Ownable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Ownable.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Ownable *OwnableCallerSession) Owner() (common.Address, error) {
	return _Ownable.Contract.Owner(&_Ownable.CallOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, _newOwner)
}

// OwnableOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Ownable contract.
type OwnableOwnershipTransferredIterator struct {
	Event *OwnableOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OwnableOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OwnableOwnershipTransferred)
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
		it.Event = new(OwnableOwnershipTransferred)
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
func (it *OwnableOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OwnableOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OwnableOwnershipTransferred represents a OwnershipTransferred event raised by the Ownable contract.
type OwnableOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OwnableOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OwnableOwnershipTransferredIterator{contract: _Ownable.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Ownable *OwnableFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OwnableOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Ownable.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OwnableOwnershipTransferred)
				if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Ownable *OwnableFilterer) ParseOwnershipTransferred(log types.Log) (*OwnableOwnershipTransferred, error) {
	event := new(OwnableOwnershipTransferred)
	if err := _Ownable.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}

// ReentrancyGuardABI is the input ABI used to generate the binding from.
const ReentrancyGuardABI = "[]"

// ReentrancyGuardBin is the compiled bytecode used for deploying new contracts.
var ReentrancyGuardBin = "0x60806040526001600055348015601457600080fd5b50604c8060226000396000f3fe6080604052600080fdfea365627a7a72315820a95be2af28afd462188de41550c842bd6cbf8385519b4ea46231051c2b01745d6c6578706572696d656e74616cf564736f6c63430005110040"

// DeployReentrancyGuard deploys a new Ethereum contract, binding an instance of ReentrancyGuard to it.
func DeployReentrancyGuard(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReentrancyGuard, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ReentrancyGuardBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// ReentrancyGuard is an auto generated Go binding around an Ethereum contract.
type ReentrancyGuard struct {
	ReentrancyGuardCaller     // Read-only binding to the contract
	ReentrancyGuardTransactor // Write-only binding to the contract
	ReentrancyGuardFilterer   // Log filterer for contract events
}

// ReentrancyGuardCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReentrancyGuardCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReentrancyGuardFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReentrancyGuardSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReentrancyGuardSession struct {
	Contract     *ReentrancyGuard  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ReentrancyGuardCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReentrancyGuardCallerSession struct {
	Contract *ReentrancyGuardCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ReentrancyGuardTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReentrancyGuardTransactorSession struct {
	Contract     *ReentrancyGuardTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ReentrancyGuardRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReentrancyGuardRaw struct {
	Contract *ReentrancyGuard // Generic contract binding to access the raw methods on
}

// ReentrancyGuardCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReentrancyGuardCallerRaw struct {
	Contract *ReentrancyGuardCaller // Generic read-only contract binding to access the raw methods on
}

// ReentrancyGuardTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReentrancyGuardTransactorRaw struct {
	Contract *ReentrancyGuardTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReentrancyGuard creates a new instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuard(address common.Address, backend bind.ContractBackend) (*ReentrancyGuard, error) {
	contract, err := bindReentrancyGuard(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuard{ReentrancyGuardCaller: ReentrancyGuardCaller{contract: contract}, ReentrancyGuardTransactor: ReentrancyGuardTransactor{contract: contract}, ReentrancyGuardFilterer: ReentrancyGuardFilterer{contract: contract}}, nil
}

// NewReentrancyGuardCaller creates a new read-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardCaller(address common.Address, caller bind.ContractCaller) (*ReentrancyGuardCaller, error) {
	contract, err := bindReentrancyGuard(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardCaller{contract: contract}, nil
}

// NewReentrancyGuardTransactor creates a new write-only instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardTransactor(address common.Address, transactor bind.ContractTransactor) (*ReentrancyGuardTransactor, error) {
	contract, err := bindReentrancyGuard(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardTransactor{contract: contract}, nil
}

// NewReentrancyGuardFilterer creates a new log filterer instance of ReentrancyGuard, bound to a specific deployed contract.
func NewReentrancyGuardFilterer(address common.Address, filterer bind.ContractFilterer) (*ReentrancyGuardFilterer, error) {
	contract, err := bindReentrancyGuard(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReentrancyGuardFilterer{contract: contract}, nil
}

// bindReentrancyGuard binds a generic wrapper to an already deployed contract.
func bindReentrancyGuard(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReentrancyGuardABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.ReentrancyGuardCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.ReentrancyGuardTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReentrancyGuard *ReentrancyGuardCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _ReentrancyGuard.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReentrancyGuard *ReentrancyGuardTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReentrancyGuard.Contract.contract.Transact(opts, method, params...)
}

// SafeMathABI is the input ABI used to generate the binding from.
const SafeMathABI = "[]"

// SafeMathBin is the compiled bytecode used for deploying new contracts.
var SafeMathBin = "0x60636023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea365627a7a723158203e1713335a0f0ce53f3241858bb92cc5db58b82f012f56cc0a6ed2b952ee91d16c6578706572696d656e74616cf564736f6c63430005110040"

// DeploySafeMath deploys a new Ethereum contract, binding an instance of SafeMath to it.
func DeploySafeMath(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SafeMath, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SafeMathBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// SafeMath is an auto generated Go binding around an Ethereum contract.
type SafeMath struct {
	SafeMathCaller     // Read-only binding to the contract
	SafeMathTransactor // Write-only binding to the contract
	SafeMathFilterer   // Log filterer for contract events
}

// SafeMathCaller is an auto generated read-only Go binding around an Ethereum contract.
type SafeMathCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SafeMathTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SafeMathFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SafeMathSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SafeMathSession struct {
	Contract     *SafeMath         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SafeMathCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SafeMathCallerSession struct {
	Contract *SafeMathCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// SafeMathTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SafeMathTransactorSession struct {
	Contract     *SafeMathTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SafeMathRaw is an auto generated low-level Go binding around an Ethereum contract.
type SafeMathRaw struct {
	Contract *SafeMath // Generic contract binding to access the raw methods on
}

// SafeMathCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SafeMathCallerRaw struct {
	Contract *SafeMathCaller // Generic read-only contract binding to access the raw methods on
}

// SafeMathTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SafeMathTransactorRaw struct {
	Contract *SafeMathTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSafeMath creates a new instance of SafeMath, bound to a specific deployed contract.
func NewSafeMath(address common.Address, backend bind.ContractBackend) (*SafeMath, error) {
	contract, err := bindSafeMath(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SafeMath{SafeMathCaller: SafeMathCaller{contract: contract}, SafeMathTransactor: SafeMathTransactor{contract: contract}, SafeMathFilterer: SafeMathFilterer{contract: contract}}, nil
}

// NewSafeMathCaller creates a new read-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathCaller(address common.Address, caller bind.ContractCaller) (*SafeMathCaller, error) {
	contract, err := bindSafeMath(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathCaller{contract: contract}, nil
}

// NewSafeMathTransactor creates a new write-only instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathTransactor(address common.Address, transactor bind.ContractTransactor) (*SafeMathTransactor, error) {
	contract, err := bindSafeMath(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SafeMathTransactor{contract: contract}, nil
}

// NewSafeMathFilterer creates a new log filterer instance of SafeMath, bound to a specific deployed contract.
func NewSafeMathFilterer(address common.Address, filterer bind.ContractFilterer) (*SafeMathFilterer, error) {
	contract, err := bindSafeMath(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SafeMathFilterer{contract: contract}, nil
}

// bindSafeMath binds a generic wrapper to an already deployed contract.
func bindSafeMath(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SafeMathABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.SafeMathCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.SafeMathTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SafeMath *SafeMathCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _SafeMath.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SafeMath *SafeMathTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SafeMath *SafeMathTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SafeMath.Contract.contract.Transact(opts, method, params...)
}

// WETHInterfaceABI is the input ABI used to generate the binding from.
const WETHInterfaceABI = "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"wad\",\"type\":\"uint256\"}],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// WETHInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var WETHInterfaceFuncSigs = map[string]string{
	"dd62ed3e": "allowance(address,address)",
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"313ce567": "decimals()",
	"d0e30db0": "deposit()",
	"06fdde03": "name()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"2e1a7d4d": "withdraw(uint256)",
}

// WETHInterface is an auto generated Go binding around an Ethereum contract.
type WETHInterface struct {
	WETHInterfaceCaller     // Read-only binding to the contract
	WETHInterfaceTransactor // Write-only binding to the contract
	WETHInterfaceFilterer   // Log filterer for contract events
}

// WETHInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type WETHInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WETHInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WETHInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WETHInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WETHInterfaceSession struct {
	Contract     *WETHInterface    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WETHInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WETHInterfaceCallerSession struct {
	Contract *WETHInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// WETHInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WETHInterfaceTransactorSession struct {
	Contract     *WETHInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// WETHInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type WETHInterfaceRaw struct {
	Contract *WETHInterface // Generic contract binding to access the raw methods on
}

// WETHInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WETHInterfaceCallerRaw struct {
	Contract *WETHInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// WETHInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WETHInterfaceTransactorRaw struct {
	Contract *WETHInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWETHInterface creates a new instance of WETHInterface, bound to a specific deployed contract.
func NewWETHInterface(address common.Address, backend bind.ContractBackend) (*WETHInterface, error) {
	contract, err := bindWETHInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WETHInterface{WETHInterfaceCaller: WETHInterfaceCaller{contract: contract}, WETHInterfaceTransactor: WETHInterfaceTransactor{contract: contract}, WETHInterfaceFilterer: WETHInterfaceFilterer{contract: contract}}, nil
}

// NewWETHInterfaceCaller creates a new read-only instance of WETHInterface, bound to a specific deployed contract.
func NewWETHInterfaceCaller(address common.Address, caller bind.ContractCaller) (*WETHInterfaceCaller, error) {
	contract, err := bindWETHInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WETHInterfaceCaller{contract: contract}, nil
}

// NewWETHInterfaceTransactor creates a new write-only instance of WETHInterface, bound to a specific deployed contract.
func NewWETHInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*WETHInterfaceTransactor, error) {
	contract, err := bindWETHInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WETHInterfaceTransactor{contract: contract}, nil
}

// NewWETHInterfaceFilterer creates a new log filterer instance of WETHInterface, bound to a specific deployed contract.
func NewWETHInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*WETHInterfaceFilterer, error) {
	contract, err := bindWETHInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WETHInterfaceFilterer{contract: contract}, nil
}

// bindWETHInterface binds a generic wrapper to an already deployed contract.
func bindWETHInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(WETHInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHInterface *WETHInterfaceRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WETHInterface.Contract.WETHInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHInterface *WETHInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHInterface.Contract.WETHInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHInterface *WETHInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHInterface.Contract.WETHInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WETHInterface *WETHInterfaceCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _WETHInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WETHInterface *WETHInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WETHInterface *WETHInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WETHInterface.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WETHInterface *WETHInterfaceCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WETHInterface *WETHInterfaceSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WETHInterface.Contract.Allowance(&_WETHInterface.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address _owner, address _spender) view returns(uint256)
func (_WETHInterface *WETHInterfaceCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _WETHInterface.Contract.Allowance(&_WETHInterface.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_WETHInterface *WETHInterfaceCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_WETHInterface *WETHInterfaceSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _WETHInterface.Contract.BalanceOf(&_WETHInterface.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _who) view returns(uint256)
func (_WETHInterface *WETHInterfaceCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _WETHInterface.Contract.BalanceOf(&_WETHInterface.CallOpts, _who)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHInterface *WETHInterfaceCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHInterface *WETHInterfaceSession) Decimals() (uint8, error) {
	return _WETHInterface.Contract.Decimals(&_WETHInterface.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_WETHInterface *WETHInterfaceCallerSession) Decimals() (uint8, error) {
	return _WETHInterface.Contract.Decimals(&_WETHInterface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHInterface *WETHInterfaceCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHInterface *WETHInterfaceSession) Name() (string, error) {
	return _WETHInterface.Contract.Name(&_WETHInterface.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_WETHInterface *WETHInterfaceCallerSession) Name() (string, error) {
	return _WETHInterface.Contract.Name(&_WETHInterface.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHInterface *WETHInterfaceCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHInterface *WETHInterfaceSession) Symbol() (string, error) {
	return _WETHInterface.Contract.Symbol(&_WETHInterface.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_WETHInterface *WETHInterfaceCallerSession) Symbol() (string, error) {
	return _WETHInterface.Contract.Symbol(&_WETHInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHInterface *WETHInterfaceCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _WETHInterface.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHInterface *WETHInterfaceSession) TotalSupply() (*big.Int, error) {
	return _WETHInterface.Contract.TotalSupply(&_WETHInterface.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_WETHInterface *WETHInterfaceCallerSession) TotalSupply() (*big.Int, error) {
	return _WETHInterface.Contract.TotalSupply(&_WETHInterface.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Approve(&_WETHInterface.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _spender, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Approve(&_WETHInterface.TransactOpts, _spender, _value)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHInterface *WETHInterfaceTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WETHInterface.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHInterface *WETHInterfaceSession) Deposit() (*types.Transaction, error) {
	return _WETHInterface.Contract.Deposit(&_WETHInterface.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() payable returns()
func (_WETHInterface *WETHInterfaceTransactorSession) Deposit() (*types.Transaction, error) {
	return _WETHInterface.Contract.Deposit(&_WETHInterface.TransactOpts)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Transfer(&_WETHInterface.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Transfer(&_WETHInterface.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.TransferFrom(&_WETHInterface.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _value) returns(bool)
func (_WETHInterface *WETHInterfaceTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.TransferFrom(&_WETHInterface.TransactOpts, _from, _to, _value)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHInterface *WETHInterfaceTransactor) Withdraw(opts *bind.TransactOpts, wad *big.Int) (*types.Transaction, error) {
	return _WETHInterface.contract.Transact(opts, "withdraw", wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHInterface *WETHInterfaceSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Withdraw(&_WETHInterface.TransactOpts, wad)
}

// Withdraw is a paid mutator transaction binding the contract method 0x2e1a7d4d.
//
// Solidity: function withdraw(uint256 wad) returns()
func (_WETHInterface *WETHInterfaceTransactorSession) Withdraw(wad *big.Int) (*types.Transaction, error) {
	return _WETHInterface.Contract.Withdraw(&_WETHInterface.TransactOpts, wad)
}

// WETHInterfaceApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the WETHInterface contract.
type WETHInterfaceApprovalIterator struct {
	Event *WETHInterfaceApproval // Event containing the contract specifics and raw log

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
func (it *WETHInterfaceApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHInterfaceApproval)
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
		it.Event = new(WETHInterfaceApproval)
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
func (it *WETHInterfaceApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHInterfaceApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHInterfaceApproval represents a Approval event raised by the WETHInterface contract.
type WETHInterfaceApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*WETHInterfaceApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETHInterface.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &WETHInterfaceApprovalIterator{contract: _WETHInterface.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *WETHInterfaceApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _WETHInterface.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHInterfaceApproval)
				if err := _WETHInterface.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) ParseApproval(log types.Log) (*WETHInterfaceApproval, error) {
	event := new(WETHInterfaceApproval)
	if err := _WETHInterface.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	return event, nil
}

// WETHInterfaceTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the WETHInterface contract.
type WETHInterfaceTransferIterator struct {
	Event *WETHInterfaceTransfer // Event containing the contract specifics and raw log

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
func (it *WETHInterfaceTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WETHInterfaceTransfer)
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
		it.Event = new(WETHInterfaceTransfer)
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
func (it *WETHInterfaceTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WETHInterfaceTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WETHInterfaceTransfer represents a Transfer event raised by the WETHInterface contract.
type WETHInterfaceTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*WETHInterfaceTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETHInterface.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &WETHInterfaceTransferIterator{contract: _WETHInterface.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *WETHInterfaceTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _WETHInterface.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WETHInterfaceTransfer)
				if err := _WETHInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_WETHInterface *WETHInterfaceFilterer) ParseTransfer(log types.Log) (*WETHInterfaceTransfer, error) {
	event := new(WETHInterfaceTransfer)
	if err := _WETHInterface.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	return event, nil
}
