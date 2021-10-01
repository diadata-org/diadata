// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package aave

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

// AaveProtocolDataProviderTokenData is an auto generated low-level Go binding around an user-defined struct.
type AaveProtocolDataProviderTokenData struct {
	Symbol       string
	TokenAddress common.Address
}

// DataTypesReserveConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveConfigurationMap struct {
	Data *big.Int
}

// DataTypesReserveData is an auto generated low-level Go binding around an user-defined struct.
type DataTypesReserveData struct {
	Configuration               DataTypesReserveConfigurationMap
	LiquidityIndex              *big.Int
	VariableBorrowIndex         *big.Int
	CurrentLiquidityRate        *big.Int
	CurrentVariableBorrowRate   *big.Int
	CurrentStableBorrowRate     *big.Int
	LastUpdateTimestamp         *big.Int
	ATokenAddress               common.Address
	StableDebtTokenAddress      common.Address
	VariableDebtTokenAddress    common.Address
	InterestRateStrategyAddress common.Address
	Id                          uint8
}

// DataTypesUserConfigurationMap is an auto generated low-level Go binding around an user-defined struct.
type DataTypesUserConfigurationMap struct {
	Data *big.Int
}

// AaveProtocolDataProviderMetaData contains all meta data concerning the AaveProtocolDataProvider contract.
var AaveProtocolDataProviderMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"addressesProvider\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"ADDRESSES_PROVIDER\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllATokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAllReservesTokens\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"internalType\":\"structAaveProtocolDataProvider.TokenData[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveConfigurationData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidationBonus\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"reserveFactor\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"borrowingEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"stableBorrowRateEnabled\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isActive\",\"type\":\"bool\"},{\"internalType\":\"bool\",\"name\":\"isFrozen\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"availableLiquidity\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"averageStableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveTokensAddresses\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserReserveData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"currentATokenBalance\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"principalStableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"scaledVariableDebt\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"stableRateLastUpdated\",\"type\":\"uint40\"},{\"internalType\":\"bool\",\"name\":\"usageAsCollateralEnabled\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60a060405234801561001057600080fd5b50604051611be3380380611be383398101604081905261002f91610044565b60601b6001600160601b031916608052610072565b600060208284031215610055578081fd5b81516001600160a01b038116811461006b578182fd5b9392505050565b60805160601c611b316100b26000398061015b528061019552806102ac52806107a75280610b2b5280610c7b5280610ff952806111295250611b316000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80633e1501411161005b5780633e150141146100f1578063b316ff891461011a578063d2493b6c1461012f578063f561ae41146101515761007d565b80630542975c1461008257806328dd2d01146100a057806335ea6a75146100c8575b600080fd5b61008a610159565b60405161009791906118e3565b60405180910390f35b6100b36100ae3660046115f5565b61017d565b60405161009799989796959493929190611a44565b6100db6100d63660046115b6565b61078e565b6040516100979a999897969594939291906119f8565b6101046100ff3660046115b6565b610b12565b6040516100979a999897969594939291906119a9565b610122610c75565b604051610097919061191a565b61014261013d3660046115b6565b610fea565b604051610097939291906118f7565b610122611123565b7f000000000000000000000000000000000000000000000000000000000000000081565b60008060008060008060008060006101936114b3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156101ec57600080fd5b505afa158015610200573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061022491906115d9565b6001600160a01b03166335ea6a758d6040518263ffffffff1660e01b815260040161024f91906118e3565b6101806040518083038186803b15801561026857600080fd5b505afa15801561027c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906102a0919061177f565b90506102aa61151e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561030357600080fd5b505afa158015610317573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061033b91906115d9565b6001600160a01b0316634417a5838d6040518263ffffffff1660e01b815260040161036691906118e3565b60206040518083038186803b15801561037e57600080fd5b505afa158015610392573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103b69190611764565b60e08301516040516370a0823160e01b81529192506001600160a01b0316906370a08231906103e9908f906004016118e3565b60206040518083038186803b15801561040157600080fd5b505afa158015610415573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610439919061187a565b6101208301516040516370a0823160e01b8152919c506001600160a01b0316906370a082319061046d908f906004016118e3565b60206040518083038186803b15801561048557600080fd5b505afa158015610499573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906104bd919061187a565b6101008301516040516370a0823160e01b8152919a506001600160a01b0316906370a08231906104f1908f906004016118e3565b60206040518083038186803b15801561050957600080fd5b505afa15801561051d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610541919061187a565b61010083015160405163631a6fd560e11b8152919b506001600160a01b03169063c634dfaa90610575908f906004016118e3565b60206040518083038186803b15801561058d57600080fd5b505afa1580156105a1573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105c5919061187a565b610120830151604051630ed1279f60e11b81529199506001600160a01b031690631da24f3e906105f9908f906004016118e3565b60206040518083038186803b15801561061157600080fd5b505afa158015610625573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610649919061187a565b965081606001516001600160801b031694508161010001516001600160a01b031663e78c9b3b8d6040518263ffffffff1660e01b815260040161068c91906118e3565b60206040518083038186803b1580156106a457600080fd5b505afa1580156106b8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106dc919061187a565b610100830151604051631e739ae360e21b81529197506001600160a01b0316906379ce6b8c90610710908f906004016118e3565b60206040518083038186803b15801561072857600080fd5b505afa15801561073c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107609190611892565b935061077d82610160015160ff16826113ea90919063ffffffff16565b925050509295985092959850929598565b6000806000806000806000806000806107a56114b3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b1580156107fe57600080fd5b505afa158015610812573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061083691906115d9565b6001600160a01b03166335ea6a758d6040518263ffffffff1660e01b815260040161086191906118e3565b6101806040518083038186803b15801561087a57600080fd5b505afa15801561088e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b2919061177f565b60e08101516040516370a0823160e01b81529192506001600160a01b038e16916370a08231916108e4916004016118e3565b60206040518083038186803b1580156108fc57600080fd5b505afa158015610910573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610934919061187a565b8161010001516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b15801561097257600080fd5b505afa158015610986573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906109aa919061187a565b8261012001516001600160a01b03166318160ddd6040518163ffffffff1660e01b815260040160206040518083038186803b1580156109e857600080fd5b505afa1580156109fc573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610a20919061187a565b836060015184608001518560a001518661010001516001600160a01b03166390f6fcf26040518163ffffffff1660e01b815260040160206040518083038186803b158015610a6d57600080fd5b505afa158015610a81573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610aa5919061187a565b876020015188604001518960c00151866001600160801b03169650856001600160801b03169550846001600160801b03169450826001600160801b03169250816001600160801b031691509a509a509a509a509a509a509a509a509a509a50509193959799509193959799565b600080600080600080600080600080610b2961151e565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610b8257600080fd5b505afa158015610b96573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610bba91906115d9565b6001600160a01b031663c44b11f78d6040518263ffffffff1660e01b8152600401610be591906118e3565b60206040518083038186803b158015610bfd57600080fd5b505afa158015610c11573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c359190611764565b9050610c408161144c565b909e50929c50909a5098509650610c5681611477565b9d9f9c9e509a9c999b989a8d15159a9099909850919650945092505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b158015610cd257600080fd5b505afa158015610ce6573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d0a91906115d9565b90506060816001600160a01b031663d1946dbc6040518163ffffffff1660e01b815260040160006040518083038186803b158015610d4757600080fd5b505afa158015610d5b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d83919081019061162d565b90506060815167ffffffffffffffff81118015610d9f57600080fd5b50604051908082528060200260200182016040528015610dd957816020015b610dc6611531565b815260200190600190039081610dbe5790505b50905060005b8251811015610fe257739f8f72aa9304c8b593d555f12ef6589cc3a579a26001600160a01b0316838281518110610e1257fe5b60200260200101516001600160a01b03161415610e915760405180604001604052806040518060400160405280600381526020016226a5a960e91b8152508152602001848381518110610e6157fe5b60200260200101516001600160a01b0316815250828281518110610e8157fe5b6020026020010181905250610fda565b73eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeeee6001600160a01b0316838281518110610ebb57fe5b60200260200101516001600160a01b03161415610f0a5760405180604001604052806040518060400160405280600381526020016208aa8960eb1b8152508152602001848381518110610e6157fe5b6040518060400160405280848381518110610f2157fe5b60200260200101516001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015610f6157600080fd5b505afa158015610f75573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f9d91908101906116d8565b8152602001848381518110610fae57fe5b60200260200101516001600160a01b0316815250828281518110610fce57fe5b60200260200101819052505b600101610ddf565b509250505090565b6000806000610ff76114b3565b7f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561105057600080fd5b505afa158015611064573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061108891906115d9565b6001600160a01b03166335ea6a75866040518263ffffffff1660e01b81526004016110b391906118e3565b6101806040518083038186803b1580156110cc57600080fd5b505afa1580156110e0573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611104919061177f565b60e0810151610100820151610120909201519097919650945092505050565b606060007f00000000000000000000000000000000000000000000000000000000000000006001600160a01b0316630261bf8b6040518163ffffffff1660e01b815260040160206040518083038186803b15801561118057600080fd5b505afa158015611194573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111b891906115d9565b90506060816001600160a01b031663d1946dbc6040518163ffffffff1660e01b815260040160006040518083038186803b1580156111f557600080fd5b505afa158015611209573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052611231919081019061162d565b90506060815167ffffffffffffffff8111801561124d57600080fd5b5060405190808252806020026020018201604052801561128757816020015b611274611531565b81526020019060019003908161126c5790505b50905060005b8251811015610fe25761129e6114b3565b846001600160a01b03166335ea6a758584815181106112b957fe5b60200260200101516040518263ffffffff1660e01b81526004016112dd91906118e3565b6101806040518083038186803b1580156112f657600080fd5b505afa15801561130a573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061132e919061177f565b905060405180604001604052808260e001516001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b15801561137857600080fd5b505afa15801561138c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526113b491908101906116d8565b81526020018260e001516001600160a01b03168152508383815181106113d657fe5b60209081029190910101525060010161128d565b60006080821060405180604001604052806002815260200161373760f01b815250906114325760405162461bcd60e51b81526004016114299190611996565b60405180910390fd5b5050815160016002830281019190911c1615155b92915050565b5161ffff80821692601083901c821692602081901c831692603082901c60ff169260409290921c1690565b51670100000000000000811615159167020000000000000082161515916704000000000000008116151591670800000000000000909116151590565b6040518061018001604052806114c761151e565b815260006020820181905260408201819052606082018190526080820181905260a0820181905260c0820181905260e082018190526101008201819052610120820181905261014082018190526101609091015290565b6040518060200160405280600081525090565b60408051808201909152606081526000602082015290565b805161144681611ae3565b600060208284031215611565578081fd5b61156f6020611a8c565b9151825250919050565b80516001600160801b038116811461144657600080fd5b805164ffffffffff8116811461144657600080fd5b805160ff8116811461144657600080fd5b6000602082840312156115c7578081fd5b81356115d281611ae3565b9392505050565b6000602082840312156115ea578081fd5b81516115d281611ae3565b60008060408385031215611607578081fd5b823561161281611ae3565b9150602083013561162281611ae3565b809150509250929050565b6000602080838503121561163f578182fd5b825167ffffffffffffffff80821115611656578384fd5b818501915085601f830112611669578384fd5b815181811115611677578485fd5b8381029150611687848301611a8c565b8181528481019084860184860187018a10156116a1578788fd5b8795505b838610156116cb576116b78a82611549565b8352600195909501949186019186016116a5565b5098975050505050505050565b6000602082840312156116e9578081fd5b815167ffffffffffffffff80821115611700578283fd5b818401915084601f830112611713578283fd5b815181811115611721578384fd5b611734601f8201601f1916602001611a8c565b915080825285602082850101111561174a578384fd5b61175b816020840160208601611ab3565b50949350505050565b600060208284031215611775578081fd5b6115d28383611554565b6000610180808385031215611792578182fd5b61179b81611a8c565b90506117a78484611554565b81526117b68460208501611579565b60208201526117c88460408501611579565b60408201526117da8460608501611579565b60608201526117ec8460808501611579565b60808201526117fe8460a08501611579565b60a08201526118108460c08501611590565b60c08201526118228460e08501611549565b60e082015261010061183685828601611549565b9082015261012061184985858301611549565b9082015261014061185c85858301611549565b9082015261016061186f858583016115a5565b908201529392505050565b60006020828403121561188b578081fd5b5051919050565b6000602082840312156118a3578081fd5b815164ffffffffff811681146115d2578182fd5b600081518084526118cf816020860160208601611ab3565b601f01601f19169290920160200192915050565b6001600160a01b0391909116815260200190565b6001600160a01b0393841681529183166020830152909116604082015260600190565b60208082528251828201819052600091906040908185019080840286018301878501865b8381101561198857888303603f1901855281518051878552611962888601826118b7565b918901516001600160a01b0316948901949094529487019492509086019060010161193e565b509098975050505050505050565b6000602082526115d260208301846118b7565b998a5260208a0198909852604089019690965260608801949094526080870192909252151560a0860152151560c0850152151560e0840152151561010083015215156101208201526101400190565b998a5260208a019890985260408901969096526060880194909452608087019290925260a086015260c085015260e084015261010083015264ffffffffff166101208201526101400190565b988952602089019790975260408801959095526060870193909352608086019190915260a085015260c084015264ffffffffff1660e083015215156101008201526101200190565b60405181810167ffffffffffffffff81118282101715611aab57600080fd5b604052919050565b60005b83811015611ace578181015183820152602001611ab6565b83811115611add576000848401525b50505050565b6001600160a01b0381168114611af857600080fd5b5056fea264697066735822122010f5e144b9567b0ac5e33edc0bdfe61fc0621c72f38d71e67464fcac21c85d9464736f6c634300060c0033",
}

// AaveProtocolDataProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use AaveProtocolDataProviderMetaData.ABI instead.
var AaveProtocolDataProviderABI = AaveProtocolDataProviderMetaData.ABI

// AaveProtocolDataProviderBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AaveProtocolDataProviderMetaData.Bin instead.
var AaveProtocolDataProviderBin = AaveProtocolDataProviderMetaData.Bin

// DeployAaveProtocolDataProvider deploys a new Ethereum contract, binding an instance of AaveProtocolDataProvider to it.
func DeployAaveProtocolDataProvider(auth *bind.TransactOpts, backend bind.ContractBackend, addressesProvider common.Address) (common.Address, *types.Transaction, *AaveProtocolDataProvider, error) {
	parsed, err := AaveProtocolDataProviderMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AaveProtocolDataProviderBin), backend, addressesProvider)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AaveProtocolDataProvider{AaveProtocolDataProviderCaller: AaveProtocolDataProviderCaller{contract: contract}, AaveProtocolDataProviderTransactor: AaveProtocolDataProviderTransactor{contract: contract}, AaveProtocolDataProviderFilterer: AaveProtocolDataProviderFilterer{contract: contract}}, nil
}

// AaveProtocolDataProvider is an auto generated Go binding around an Ethereum contract.
type AaveProtocolDataProvider struct {
	AaveProtocolDataProviderCaller     // Read-only binding to the contract
	AaveProtocolDataProviderTransactor // Write-only binding to the contract
	AaveProtocolDataProviderFilterer   // Log filterer for contract events
}

// AaveProtocolDataProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type AaveProtocolDataProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveProtocolDataProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AaveProtocolDataProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveProtocolDataProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AaveProtocolDataProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AaveProtocolDataProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AaveProtocolDataProviderSession struct {
	Contract     *AaveProtocolDataProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts             // Call options to use throughout this session
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// AaveProtocolDataProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AaveProtocolDataProviderCallerSession struct {
	Contract *AaveProtocolDataProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                   // Call options to use throughout this session
}

// AaveProtocolDataProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AaveProtocolDataProviderTransactorSession struct {
	Contract     *AaveProtocolDataProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                   // Transaction auth options to use throughout this session
}

// AaveProtocolDataProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type AaveProtocolDataProviderRaw struct {
	Contract *AaveProtocolDataProvider // Generic contract binding to access the raw methods on
}

// AaveProtocolDataProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AaveProtocolDataProviderCallerRaw struct {
	Contract *AaveProtocolDataProviderCaller // Generic read-only contract binding to access the raw methods on
}

// AaveProtocolDataProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AaveProtocolDataProviderTransactorRaw struct {
	Contract *AaveProtocolDataProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAaveProtocolDataProvider creates a new instance of AaveProtocolDataProvider, bound to a specific deployed contract.
func NewAaveProtocolDataProvider(address common.Address, backend bind.ContractBackend) (*AaveProtocolDataProvider, error) {
	contract, err := bindAaveProtocolDataProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AaveProtocolDataProvider{AaveProtocolDataProviderCaller: AaveProtocolDataProviderCaller{contract: contract}, AaveProtocolDataProviderTransactor: AaveProtocolDataProviderTransactor{contract: contract}, AaveProtocolDataProviderFilterer: AaveProtocolDataProviderFilterer{contract: contract}}, nil
}

// NewAaveProtocolDataProviderCaller creates a new read-only instance of AaveProtocolDataProvider, bound to a specific deployed contract.
func NewAaveProtocolDataProviderCaller(address common.Address, caller bind.ContractCaller) (*AaveProtocolDataProviderCaller, error) {
	contract, err := bindAaveProtocolDataProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AaveProtocolDataProviderCaller{contract: contract}, nil
}

// NewAaveProtocolDataProviderTransactor creates a new write-only instance of AaveProtocolDataProvider, bound to a specific deployed contract.
func NewAaveProtocolDataProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*AaveProtocolDataProviderTransactor, error) {
	contract, err := bindAaveProtocolDataProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AaveProtocolDataProviderTransactor{contract: contract}, nil
}

// NewAaveProtocolDataProviderFilterer creates a new log filterer instance of AaveProtocolDataProvider, bound to a specific deployed contract.
func NewAaveProtocolDataProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*AaveProtocolDataProviderFilterer, error) {
	contract, err := bindAaveProtocolDataProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AaveProtocolDataProviderFilterer{contract: contract}, nil
}

// bindAaveProtocolDataProvider binds a generic wrapper to an already deployed contract.
func bindAaveProtocolDataProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(AaveProtocolDataProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveProtocolDataProvider.Contract.AaveProtocolDataProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveProtocolDataProvider.Contract.AaveProtocolDataProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveProtocolDataProvider.Contract.AaveProtocolDataProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AaveProtocolDataProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AaveProtocolDataProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AaveProtocolDataProvider *AaveProtocolDataProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AaveProtocolDataProvider.Contract.contract.Transact(opts, method, params...)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) ADDRESSESPROVIDER(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "ADDRESSES_PROVIDER")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveProtocolDataProvider.CallOpts)
}

// ADDRESSESPROVIDER is a free data retrieval call binding the contract method 0x0542975c.
//
// Solidity: function ADDRESSES_PROVIDER() view returns(address)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) ADDRESSESPROVIDER() (common.Address, error) {
	return _AaveProtocolDataProvider.Contract.ADDRESSESPROVIDER(&_AaveProtocolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetAllATokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getAllATokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveProtocolDataProvider.Contract.GetAllATokens(&_AaveProtocolDataProvider.CallOpts)
}

// GetAllATokens is a free data retrieval call binding the contract method 0xf561ae41.
//
// Solidity: function getAllATokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetAllATokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveProtocolDataProvider.Contract.GetAllATokens(&_AaveProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetAllReservesTokens(opts *bind.CallOpts) ([]AaveProtocolDataProviderTokenData, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getAllReservesTokens")

	if err != nil {
		return *new([]AaveProtocolDataProviderTokenData), err
	}

	out0 := *abi.ConvertType(out[0], new([]AaveProtocolDataProviderTokenData)).(*[]AaveProtocolDataProviderTokenData)

	return out0, err

}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveProtocolDataProvider.CallOpts)
}

// GetAllReservesTokens is a free data retrieval call binding the contract method 0xb316ff89.
//
// Solidity: function getAllReservesTokens() view returns((string,address)[])
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetAllReservesTokens() ([]AaveProtocolDataProviderTokenData, error) {
	return _AaveProtocolDataProvider.Contract.GetAllReservesTokens(&_AaveProtocolDataProvider.CallOpts)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetReserveConfigurationData(opts *bind.CallOpts, asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getReserveConfigurationData", asset)

	outstruct := new(struct {
		Decimals                 *big.Int
		Ltv                      *big.Int
		LiquidationThreshold     *big.Int
		LiquidationBonus         *big.Int
		ReserveFactor            *big.Int
		UsageAsCollateralEnabled bool
		BorrowingEnabled         bool
		StableBorrowRateEnabled  bool
		IsActive                 bool
		IsFrozen                 bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Decimals = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.LiquidationThreshold = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidationBonus = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ReserveFactor = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[5], new(bool)).(*bool)
	outstruct.BorrowingEnabled = *abi.ConvertType(out[6], new(bool)).(*bool)
	outstruct.StableBorrowRateEnabled = *abi.ConvertType(out[7], new(bool)).(*bool)
	outstruct.IsActive = *abi.ConvertType(out[8], new(bool)).(*bool)
	outstruct.IsFrozen = *abi.ConvertType(out[9], new(bool)).(*bool)

	return *outstruct, err

}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetReserveConfigurationData is a free data retrieval call binding the contract method 0x3e150141.
//
// Solidity: function getReserveConfigurationData(address asset) view returns(uint256 decimals, uint256 ltv, uint256 liquidationThreshold, uint256 liquidationBonus, uint256 reserveFactor, bool usageAsCollateralEnabled, bool borrowingEnabled, bool stableBorrowRateEnabled, bool isActive, bool isFrozen)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetReserveConfigurationData(asset common.Address) (struct {
	Decimals                 *big.Int
	Ltv                      *big.Int
	LiquidationThreshold     *big.Int
	LiquidationBonus         *big.Int
	ReserveFactor            *big.Int
	UsageAsCollateralEnabled bool
	BorrowingEnabled         bool
	StableBorrowRateEnabled  bool
	IsActive                 bool
	IsFrozen                 bool
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveConfigurationData(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getReserveData", asset)

	outstruct := new(struct {
		AvailableLiquidity      *big.Int
		TotalStableDebt         *big.Int
		TotalVariableDebt       *big.Int
		LiquidityRate           *big.Int
		VariableBorrowRate      *big.Int
		StableBorrowRate        *big.Int
		AverageStableBorrowRate *big.Int
		LiquidityIndex          *big.Int
		VariableBorrowIndex     *big.Int
		LastUpdateTimestamp     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.AvailableLiquidity = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.TotalVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowRate = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.AverageStableBorrowRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.LiquidityIndex = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.VariableBorrowIndex = *abi.ConvertType(out[8], new(*big.Int)).(**big.Int)
	outstruct.LastUpdateTimestamp = *abi.ConvertType(out[9], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveData(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(uint256 availableLiquidity, uint256 totalStableDebt, uint256 totalVariableDebt, uint256 liquidityRate, uint256 variableBorrowRate, uint256 stableBorrowRate, uint256 averageStableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex, uint40 lastUpdateTimestamp)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetReserveData(asset common.Address) (struct {
	AvailableLiquidity      *big.Int
	TotalStableDebt         *big.Int
	TotalVariableDebt       *big.Int
	LiquidityRate           *big.Int
	VariableBorrowRate      *big.Int
	StableBorrowRate        *big.Int
	AverageStableBorrowRate *big.Int
	LiquidityIndex          *big.Int
	VariableBorrowIndex     *big.Int
	LastUpdateTimestamp     *big.Int
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveData(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetReserveTokensAddresses(opts *bind.CallOpts, asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getReserveTokensAddresses", asset)

	outstruct := new(struct {
		ATokenAddress            common.Address
		StableDebtTokenAddress   common.Address
		VariableDebtTokenAddress common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.ATokenAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.StableDebtTokenAddress = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.VariableDebtTokenAddress = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetReserveTokensAddresses is a free data retrieval call binding the contract method 0xd2493b6c.
//
// Solidity: function getReserveTokensAddresses(address asset) view returns(address aTokenAddress, address stableDebtTokenAddress, address variableDebtTokenAddress)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetReserveTokensAddresses(asset common.Address) (struct {
	ATokenAddress            common.Address
	StableDebtTokenAddress   common.Address
	VariableDebtTokenAddress common.Address
}, error) {
	return _AaveProtocolDataProvider.Contract.GetReserveTokensAddresses(&_AaveProtocolDataProvider.CallOpts, asset)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCaller) GetUserReserveData(opts *bind.CallOpts, asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	var out []interface{}
	err := _AaveProtocolDataProvider.contract.Call(opts, &out, "getUserReserveData", asset, user)

	outstruct := new(struct {
		CurrentATokenBalance     *big.Int
		CurrentStableDebt        *big.Int
		CurrentVariableDebt      *big.Int
		PrincipalStableDebt      *big.Int
		ScaledVariableDebt       *big.Int
		StableBorrowRate         *big.Int
		LiquidityRate            *big.Int
		StableRateLastUpdated    *big.Int
		UsageAsCollateralEnabled bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.CurrentATokenBalance = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.CurrentStableDebt = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.CurrentVariableDebt = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.PrincipalStableDebt = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ScaledVariableDebt = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.StableBorrowRate = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)
	outstruct.LiquidityRate = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.StableRateLastUpdated = *abi.ConvertType(out[7], new(*big.Int)).(**big.Int)
	outstruct.UsageAsCollateralEnabled = *abi.ConvertType(out[8], new(bool)).(*bool)

	return *outstruct, err

}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveProtocolDataProvider.Contract.GetUserReserveData(&_AaveProtocolDataProvider.CallOpts, asset, user)
}

// GetUserReserveData is a free data retrieval call binding the contract method 0x28dd2d01.
//
// Solidity: function getUserReserveData(address asset, address user) view returns(uint256 currentATokenBalance, uint256 currentStableDebt, uint256 currentVariableDebt, uint256 principalStableDebt, uint256 scaledVariableDebt, uint256 stableBorrowRate, uint256 liquidityRate, uint40 stableRateLastUpdated, bool usageAsCollateralEnabled)
func (_AaveProtocolDataProvider *AaveProtocolDataProviderCallerSession) GetUserReserveData(asset common.Address, user common.Address) (struct {
	CurrentATokenBalance     *big.Int
	CurrentStableDebt        *big.Int
	CurrentVariableDebt      *big.Int
	PrincipalStableDebt      *big.Int
	ScaledVariableDebt       *big.Int
	StableBorrowRate         *big.Int
	LiquidityRate            *big.Int
	StableRateLastUpdated    *big.Int
	UsageAsCollateralEnabled bool
}, error) {
	return _AaveProtocolDataProvider.Contract.GetUserReserveData(&_AaveProtocolDataProvider.CallOpts, asset, user)
}

// DataTypesMetaData contains all meta data concerning the DataTypes contract.
var DataTypesMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220bf39ba26992d049460ba35e85a867a609b1ac1e1b9b481a152e53e3413bd83dd64736f6c634300060c0033",
}

// DataTypesABI is the input ABI used to generate the binding from.
// Deprecated: Use DataTypesMetaData.ABI instead.
var DataTypesABI = DataTypesMetaData.ABI

// DataTypesBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use DataTypesMetaData.Bin instead.
var DataTypesBin = DataTypesMetaData.Bin

// DeployDataTypes deploys a new Ethereum contract, binding an instance of DataTypes to it.
func DeployDataTypes(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *DataTypes, error) {
	parsed, err := DataTypesMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(DataTypesBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &DataTypes{DataTypesCaller: DataTypesCaller{contract: contract}, DataTypesTransactor: DataTypesTransactor{contract: contract}, DataTypesFilterer: DataTypesFilterer{contract: contract}}, nil
}

// DataTypes is an auto generated Go binding around an Ethereum contract.
type DataTypes struct {
	DataTypesCaller     // Read-only binding to the contract
	DataTypesTransactor // Write-only binding to the contract
	DataTypesFilterer   // Log filterer for contract events
}

// DataTypesCaller is an auto generated read-only Go binding around an Ethereum contract.
type DataTypesCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesTransactor is an auto generated write-only Go binding around an Ethereum contract.
type DataTypesTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type DataTypesFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// DataTypesSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type DataTypesSession struct {
	Contract     *DataTypes        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// DataTypesCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type DataTypesCallerSession struct {
	Contract *DataTypesCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// DataTypesTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type DataTypesTransactorSession struct {
	Contract     *DataTypesTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// DataTypesRaw is an auto generated low-level Go binding around an Ethereum contract.
type DataTypesRaw struct {
	Contract *DataTypes // Generic contract binding to access the raw methods on
}

// DataTypesCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type DataTypesCallerRaw struct {
	Contract *DataTypesCaller // Generic read-only contract binding to access the raw methods on
}

// DataTypesTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type DataTypesTransactorRaw struct {
	Contract *DataTypesTransactor // Generic write-only contract binding to access the raw methods on
}

// NewDataTypes creates a new instance of DataTypes, bound to a specific deployed contract.
func NewDataTypes(address common.Address, backend bind.ContractBackend) (*DataTypes, error) {
	contract, err := bindDataTypes(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &DataTypes{DataTypesCaller: DataTypesCaller{contract: contract}, DataTypesTransactor: DataTypesTransactor{contract: contract}, DataTypesFilterer: DataTypesFilterer{contract: contract}}, nil
}

// NewDataTypesCaller creates a new read-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesCaller(address common.Address, caller bind.ContractCaller) (*DataTypesCaller, error) {
	contract, err := bindDataTypes(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesCaller{contract: contract}, nil
}

// NewDataTypesTransactor creates a new write-only instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesTransactor(address common.Address, transactor bind.ContractTransactor) (*DataTypesTransactor, error) {
	contract, err := bindDataTypes(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &DataTypesTransactor{contract: contract}, nil
}

// NewDataTypesFilterer creates a new log filterer instance of DataTypes, bound to a specific deployed contract.
func NewDataTypesFilterer(address common.Address, filterer bind.ContractFilterer) (*DataTypesFilterer, error) {
	contract, err := bindDataTypes(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &DataTypesFilterer{contract: contract}, nil
}

// bindDataTypes binds a generic wrapper to an already deployed contract.
func bindDataTypes(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(DataTypesABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.DataTypesCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.DataTypesTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_DataTypes *DataTypesCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _DataTypes.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_DataTypes *DataTypesTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_DataTypes *DataTypesTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _DataTypes.Contract.contract.Transact(opts, method, params...)
}

// ErrorsMetaData contains all meta data concerning the Errors contract.
var ErrorsMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"BORROW_ALLOWANCE_NOT_ENOUGH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CALLER_NOT_POOL_ADMIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CT_CALLER_MUST_BE_LENDING_POOL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CT_CANNOT_GIVE_ALLOWANCE_TO_HIMSELF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CT_INVALID_BURN_AMOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CT_INVALID_MINT_AMOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"CT_TRANSFER_AMOUNT_NOT_GT_0\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPAPR_INVALID_ADDRESSES_PROVIDER_ID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPAPR_PROVIDER_NOT_REGISTERED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPCM_COLLATERAL_CANNOT_BE_LIQUIDATED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPCM_HEALTH_FACTOR_NOT_BELOW_THRESHOLD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPCM_NOT_ENOUGH_LIQUIDITY_TO_LIQUIDATE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPCM_NO_ERRORS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPCM_SPECIFIED_CURRENCY_NOT_BORROWED_BY_USER\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_CALLER_NOT_EMERGENCY_ADMIN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_ADDRESSES_PROVIDER_ID\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_ATOKEN_POOL_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_CONFIGURATION\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_STABLE_DEBT_TOKEN_POOL_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_STABLE_DEBT_TOKEN_UNDERLYING_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_VARIABLE_DEBT_TOKEN_POOL_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_INVALID_VARIABLE_DEBT_TOKEN_UNDERLYING_ADDRESS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LPC_RESERVE_LIQUIDITY_NOT_0\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_CALLER_MUST_BE_AN_ATOKEN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_CALLER_NOT_LENDING_POOL_CONFIGURATOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_FAILED_COLLATERAL_SWAP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_FAILED_REPAY_WITH_COLLATERAL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INCONSISTENT_FLASHLOAN_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INCONSISTENT_PARAMS_LENGTH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INCONSISTENT_PROTOCOL_ACTUAL_BALANCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INTEREST_RATE_REBALANCE_CONDITIONS_NOT_MET\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INVALID_EQUAL_ASSETS_TO_SWAP\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INVALID_FLASHLOAN_MODE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_INVALID_FLASH_LOAN_EXECUTOR_RETURN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_IS_PAUSED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_LIQUIDATION_CALL_FAILED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_NOT_CONTRACT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_NOT_ENOUGH_LIQUIDITY_TO_BORROW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_NOT_ENOUGH_STABLE_BORROW_BALANCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_NO_MORE_RESERVES_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_REENTRANCY_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"LP_REQUESTED_AMOUNT_TOO_SMALL\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_ADDITION_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_DIVISION_BY_ZERO\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"MATH_MULTIPLICATION_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RC_INVALID_DECIMALS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RC_INVALID_LIQ_BONUS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RC_INVALID_LIQ_THRESHOLD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RC_INVALID_LTV\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RC_INVALID_RESERVE_FACTOR\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_LIQUIDITY_INDEX_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_LIQUIDITY_RATE_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_RESERVE_ALREADY_INITIALIZED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_STABLE_BORROW_RATE_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_VARIABLE_BORROW_INDEX_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"RL_VARIABLE_BORROW_RATE_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SDT_BURN_EXCEEDS_BALANCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"SDT_STABLE_DEBT_OVERFLOW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"UL_INVALID_INDEX\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_AMOUNT_BIGGER_THAN_MAX_LOAN_SIZE_STABLE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_BORROWING_NOT_ENABLED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_COLLATERAL_BALANCE_IS_0\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_COLLATERAL_CANNOT_COVER_NEW_BORROW\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_COLLATERAL_SAME_AS_BORROWING_CURRENCY\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_CURRENT_AVAILABLE_LIQUIDITY_NOT_ENOUGH\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_DEPOSIT_ALREADY_IN_USE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_HEALTH_FACTOR_LOWER_THAN_LIQUIDATION_THRESHOLD\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_INCONSISTENT_FLASHLOAN_PARAMS\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_INVALID_AMOUNT\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_INVALID_INTEREST_RATE_MODE_SELECTED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NOT_ENOUGH_AVAILABLE_USER_BALANCE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NO_ACTIVE_RESERVE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NO_DEBT_OF_SELECTED_TYPE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NO_EXPLICIT_AMOUNT_TO_REPAY_ON_BEHALF\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NO_STABLE_RATE_LOAN_IN_RESERVE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_NO_VARIABLE_RATE_LOAN_IN_RESERVE\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_RESERVE_FROZEN\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_STABLE_BORROWING_NOT_ENABLED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_TRANSFER_NOT_ALLOWED\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"VL_UNDERLYING_BALANCE_NOT_GREATER_THAN_0\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x611105610026600b82828239805160001a60731461001957fe5b30600052607381538281f3fe730000000000000000000000000000000000000000301460806040526004361061048a5760003560e01c80636ba4271f11610261578063cdad445a11610150578063e2c16d69116100cd578063f11c672011610091578063f11c6720146106fd578063f3d9cc1114610705578063f902735d1461070d578063fb681def14610715578063fe75fd261461071d5761048a565b8063e2c16d69146106d5578063e6632748146106dd578063e7bf91b3146106e5578063eca85d3a146106ed578063f0473259146106f55761048a565b8063d7510e0c11610114578063d7510e0c146106ad578063d7b079aa146106b5578063daf23547146106bd578063e0d7dfd7146106c5578063e29425dc146106cd5761048a565b8063cdad445a14610685578063d3e370ee1461068d578063d44e8e8814610695578063d57bb9641461069d578063d6f681b6146106a55761048a565b8063a39ed4ff116101de578063b89652cd116101a2578063b89652cd1461065d578063bd013f5b14610665578063c09e26181461066d578063c2d628df14610675578063cc5fc44c1461067d5761048a565b8063a39ed4ff14610635578063a84402411461063d578063ac75323614610645578063b36a2cf31461064d578063b72e40c7146106555761048a565b80637865a627116102255780637865a6271461060d578063871938a81461061557806391a9fb181461061d5780639be4f03a14610625578063a2fbc8ad1461062d5761048a565b80636ba4271f146105e55780636d422aa1146105ed578063708b8dd3146105f557806371a629da146105fd57806376f19030146106055761048a565b80633aa786a81161037d5780634a529f91116102fa578063614cf6a1116102be578063614cf6a1146105bd578063637a5a12146105c55780636422b257146105cd57806365344799146105d55780636ab5e615146105dd5761048a565b80634a529f91146105955780634fe4f1ab1461059d57806355bab12c146105a55780635a9786d4146105ad5780635e869ff1146105b55761048a565b80634349e3d8116103415780634349e3d81461056d578063449420041461057557806344dc4f701461057d57806347d25300146105855780634927c63a1461058d5761048a565b80633aa786a8146105455780633b5d25aa1461054d5780633f5d6ec814610555578063407374a41461055d57806341b40ba5146105655761048a565b806322a6f08e1161040b578063333e8ea8116103cf578063333e8ea81461051d57806335a9d21d1461052557806336565ab11461052d5780633872b0ad14610535578063390f34ba1461053d5761048a565b806322a6f08e146104f55780632ace698a146104fd5780632b34c349146105055780632b9c57f61461050d5780632ea347b0146105155761048a565b80631291a38b116104525780631291a38b146104cd578063179476c5146104d55780631befa78d146104dd5780631ea7c604146104e55780631ec68b1d146104ed5761048a565b806302454ad31461048f578063029d2344146104ad57806306f355ad146104b55780630b8fd588146104bd5780630f5ee482146104c5575b600080fd5b610497610725565b6040516104a4919061107c565b60405180910390f35b610497610743565b610497610761565b61049761077f565b61049761079d565b6104976107bb565b6104976107d9565b6104976107f6565b610497610814565b610497610832565b610497610850565b61049761086e565b61049761088c565b6104976108aa565b6104976108c8565b6104976108e6565b610497610904565b610497610922565b61049761093f565b61049761095d565b61049761097b565b610497610999565b6104976109b6565b6104976109d4565b6104976109f2565b610497610a10565b610497610a2e565b610497610a4c565b610497610a6a565b610497610a88565b610497610aa6565b610497610ac4565b610497610ae2565b610497610b00565b610497610b1e565b610497610b3c565b610497610b5a565b610497610b78565b610497610b96565b610497610bb4565b610497610bd2565b610497610bf0565b610497610c0e565b610497610c2b565b610497610c49565b610497610c67565b610497610c84565b610497610ca1565b610497610cbf565b610497610cdd565b610497610cfb565b610497610d19565b610497610d36565b610497610d54565b610497610d72565b610497610d90565b610497610dae565b610497610dcc565b610497610dea565b610497610e08565b610497610e26565b610497610e44565b610497610e62565b610497610e80565b610497610e9e565b610497610ebc565b610497610ed9565b610497610ef7565b610497610f15565b610497610f33565b610497610f51565b610497610f6f565b610497610f8d565b610497610fab565b610497610fc9565b610497610fe7565b610497611005565b610497611022565b610497611040565b61049761105e565b60405180604001604052806002815260200161373760f01b81525081565b60405180604001604052806002815260200161068760f31b81525081565b60405180604001604052806002815260200161033360f41b81525081565b60405180604001604052806002815260200161191b60f11b81525081565b60405180604001604052806002815260200161343960f01b81525081565b604051806040016040528060028152602001611a9b60f11b81525081565b604051806040016040528060018152602001600d60fa1b81525081565b60405180604001604052806002815260200161038360f41b81525081565b604051806040016040528060028152602001611a1b60f11b81525081565b60405180604001604052806002815260200161031360f41b81525081565b604051806040016040528060028152602001610c8d60f21b81525081565b60405180604001604052806002815260200161313160f01b81525081565b60405180604001604052806002815260200161064760f31b81525081565b6040518060400160405280600281526020016106a760f31b81525081565b604051806040016040528060028152602001610d4d60f21b81525081565b604051806040016040528060028152602001611b9960f11b81525081565b60405180604001604052806002815260200161313960f01b81525081565b604051806040016040528060018152602001603760f81b81525081565b60405180604001604052806002815260200161333960f01b81525081565b60405180604001604052806002815260200161323560f01b81525081565b604051806040016040528060028152602001610c4d60f21b81525081565b604051806040016040528060018152602001600760fb1b81525081565b60405180604001604052806002815260200161037360f41b81525081565b60405180604001604052806002815260200161343360f01b81525081565b60405180604001604052806002815260200161066760f31b81525081565b60405180604001604052806002815260200161035360f41b81525081565b604051806040016040528060028152602001611a9960f11b81525081565b60405180604001604052806002815260200161323160f01b81525081565b60405180604001604052806002815260200161373560f01b81525081565b60405180604001604052806002815260200161189960f11b81525081565b60405180604001604052806002815260200161323360f01b81525081565b60405180604001604052806002815260200161353160f01b81525081565b60405180604001604052806002815260200161036360f41b81525081565b60405180604001604052806002815260200161034360f41b81525081565b60405180604001604052806002815260200161363960f01b81525081565b60405180604001604052806002815260200161363760f01b81525081565b6040518060400160405280600281526020016106e760f31b81525081565b60405180604001604052806002815260200161313760f01b81525081565b604051806040016040528060028152602001610ccd60f21b81525081565b60405180604001604052806002815260200161062760f31b81525081565b60405180604001604052806002815260200161323960f01b81525081565b60405180604001604052806002815260200161353560f01b81525081565b604051806040016040528060018152602001603960f81b81525081565b604051806040016040528060028152602001610d0d60f21b81525081565b60405180604001604052806002815260200161363560f01b81525081565b604051806040016040528060018152602001601960f91b81525081565b604051806040016040528060018152602001603160f81b81525081565b60405180604001604052806002815260200161313560f01b81525081565b60405180604001604052806002815260200161373160f01b81525081565b60405180604001604052806002815260200161333160f01b81525081565b60405180604001604052806002815260200161313360f01b81525081565b604051806040016040528060018152602001603560f81b81525081565b60405180604001604052806002815260200161333360f01b81525081565b60405180604001604052806002815260200161323760f01b81525081565b604051806040016040528060028152602001610dcd60f21b81525081565b60405180604001604052806002815260200161191960f11b81525081565b6040518060400160405280600281526020016106c760f31b81525081565b60405180604001604052806002815260200161333760f01b81525081565b60405180604001604052806002815260200161363160f01b81525081565b60405180604001604052806002815260200161343560f01b81525081565b60405180604001604052806002815260200161373960f01b81525081565b604051806040016040528060028152602001611b9b60f11b81525081565b604051806040016040528060028152602001611b1b60f11b81525081565b604051806040016040528060028152602001610d8d60f21b81525081565b60405180604001604052806002815260200161343160f01b81525081565b604051806040016040528060018152602001603360f81b81525081565b60405180604001604052806002815260200161373360f01b81525081565b60405180604001604052806002815260200161189b60f11b81525081565b60405180604001604052806002815260200161199b60f11b81525081565b60405180604001604052806002815260200161032360f41b81525081565b60405180604001604052806002815260200161353960f01b81525081565b60405180604001604052806002815260200161353760f01b81525081565b60405180604001604052806002815260200161343760f01b81525081565b60405180604001604052806002815260200161363360f01b81525081565b60405180604001604052806002815260200161333560f01b81525081565b60405180604001604052806002815260200161353360f01b81525081565b604051806040016040528060018152602001601b60f91b81525081565b604051806040016040528060028152602001611b1960f11b81525081565b604051806040016040528060028152602001611a1960f11b81525081565b60405180604001604052806002815260200161199960f11b81525081565b6000602080835283518082850152825b818110156110a85785810183015185820160400152820161108c565b818111156110b95783604083870101525b50601f01601f191692909201604001939250505056fea2646970667358221220045465ede76c72ad0271d3fb613fc5c5258eee07900dc90dfb4def807a5e277c64736f6c634300060c0033",
}

// ErrorsABI is the input ABI used to generate the binding from.
// Deprecated: Use ErrorsMetaData.ABI instead.
var ErrorsABI = ErrorsMetaData.ABI

// ErrorsBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ErrorsMetaData.Bin instead.
var ErrorsBin = ErrorsMetaData.Bin

// DeployErrors deploys a new Ethereum contract, binding an instance of Errors to it.
func DeployErrors(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Errors, error) {
	parsed, err := ErrorsMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ErrorsBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// Errors is an auto generated Go binding around an Ethereum contract.
type Errors struct {
	ErrorsCaller     // Read-only binding to the contract
	ErrorsTransactor // Write-only binding to the contract
	ErrorsFilterer   // Log filterer for contract events
}

// ErrorsCaller is an auto generated read-only Go binding around an Ethereum contract.
type ErrorsCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ErrorsTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ErrorsFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ErrorsSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ErrorsSession struct {
	Contract     *Errors           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ErrorsCallerSession struct {
	Contract *ErrorsCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ErrorsTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ErrorsTransactorSession struct {
	Contract     *ErrorsTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ErrorsRaw is an auto generated low-level Go binding around an Ethereum contract.
type ErrorsRaw struct {
	Contract *Errors // Generic contract binding to access the raw methods on
}

// ErrorsCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ErrorsCallerRaw struct {
	Contract *ErrorsCaller // Generic read-only contract binding to access the raw methods on
}

// ErrorsTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ErrorsTransactorRaw struct {
	Contract *ErrorsTransactor // Generic write-only contract binding to access the raw methods on
}

// NewErrors creates a new instance of Errors, bound to a specific deployed contract.
func NewErrors(address common.Address, backend bind.ContractBackend) (*Errors, error) {
	contract, err := bindErrors(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Errors{ErrorsCaller: ErrorsCaller{contract: contract}, ErrorsTransactor: ErrorsTransactor{contract: contract}, ErrorsFilterer: ErrorsFilterer{contract: contract}}, nil
}

// NewErrorsCaller creates a new read-only instance of Errors, bound to a specific deployed contract.
func NewErrorsCaller(address common.Address, caller bind.ContractCaller) (*ErrorsCaller, error) {
	contract, err := bindErrors(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsCaller{contract: contract}, nil
}

// NewErrorsTransactor creates a new write-only instance of Errors, bound to a specific deployed contract.
func NewErrorsTransactor(address common.Address, transactor bind.ContractTransactor) (*ErrorsTransactor, error) {
	contract, err := bindErrors(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ErrorsTransactor{contract: contract}, nil
}

// NewErrorsFilterer creates a new log filterer instance of Errors, bound to a specific deployed contract.
func NewErrorsFilterer(address common.Address, filterer bind.ContractFilterer) (*ErrorsFilterer, error) {
	contract, err := bindErrors(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ErrorsFilterer{contract: contract}, nil
}

// bindErrors binds a generic wrapper to an already deployed contract.
func bindErrors(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ErrorsABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.ErrorsCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.ErrorsTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Errors *ErrorsCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Errors.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Errors *ErrorsTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Errors *ErrorsTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Errors.Contract.contract.Transact(opts, method, params...)
}

// BORROWALLOWANCENOTENOUGH is a free data retrieval call binding the contract method 0xe2c16d69.
//
// Solidity: function BORROW_ALLOWANCE_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsCaller) BORROWALLOWANCENOTENOUGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "BORROW_ALLOWANCE_NOT_ENOUGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// BORROWALLOWANCENOTENOUGH is a free data retrieval call binding the contract method 0xe2c16d69.
//
// Solidity: function BORROW_ALLOWANCE_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsSession) BORROWALLOWANCENOTENOUGH() (string, error) {
	return _Errors.Contract.BORROWALLOWANCENOTENOUGH(&_Errors.CallOpts)
}

// BORROWALLOWANCENOTENOUGH is a free data retrieval call binding the contract method 0xe2c16d69.
//
// Solidity: function BORROW_ALLOWANCE_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsCallerSession) BORROWALLOWANCENOTENOUGH() (string, error) {
	return _Errors.Contract.BORROWALLOWANCENOTENOUGH(&_Errors.CallOpts)
}

// CALLERNOTPOOLADMIN is a free data retrieval call binding the contract method 0xac753236.
//
// Solidity: function CALLER_NOT_POOL_ADMIN() view returns(string)
func (_Errors *ErrorsCaller) CALLERNOTPOOLADMIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CALLER_NOT_POOL_ADMIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CALLERNOTPOOLADMIN is a free data retrieval call binding the contract method 0xac753236.
//
// Solidity: function CALLER_NOT_POOL_ADMIN() view returns(string)
func (_Errors *ErrorsSession) CALLERNOTPOOLADMIN() (string, error) {
	return _Errors.Contract.CALLERNOTPOOLADMIN(&_Errors.CallOpts)
}

// CALLERNOTPOOLADMIN is a free data retrieval call binding the contract method 0xac753236.
//
// Solidity: function CALLER_NOT_POOL_ADMIN() view returns(string)
func (_Errors *ErrorsCallerSession) CALLERNOTPOOLADMIN() (string, error) {
	return _Errors.Contract.CALLERNOTPOOLADMIN(&_Errors.CallOpts)
}

// CTCALLERMUSTBELENDINGPOOL is a free data retrieval call binding the contract method 0x6ba4271f.
//
// Solidity: function CT_CALLER_MUST_BE_LENDING_POOL() view returns(string)
func (_Errors *ErrorsCaller) CTCALLERMUSTBELENDINGPOOL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CT_CALLER_MUST_BE_LENDING_POOL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CTCALLERMUSTBELENDINGPOOL is a free data retrieval call binding the contract method 0x6ba4271f.
//
// Solidity: function CT_CALLER_MUST_BE_LENDING_POOL() view returns(string)
func (_Errors *ErrorsSession) CTCALLERMUSTBELENDINGPOOL() (string, error) {
	return _Errors.Contract.CTCALLERMUSTBELENDINGPOOL(&_Errors.CallOpts)
}

// CTCALLERMUSTBELENDINGPOOL is a free data retrieval call binding the contract method 0x6ba4271f.
//
// Solidity: function CT_CALLER_MUST_BE_LENDING_POOL() view returns(string)
func (_Errors *ErrorsCallerSession) CTCALLERMUSTBELENDINGPOOL() (string, error) {
	return _Errors.Contract.CTCALLERMUSTBELENDINGPOOL(&_Errors.CallOpts)
}

// CTCANNOTGIVEALLOWANCETOHIMSELF is a free data retrieval call binding the contract method 0x06f355ad.
//
// Solidity: function CT_CANNOT_GIVE_ALLOWANCE_TO_HIMSELF() view returns(string)
func (_Errors *ErrorsCaller) CTCANNOTGIVEALLOWANCETOHIMSELF(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CT_CANNOT_GIVE_ALLOWANCE_TO_HIMSELF")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CTCANNOTGIVEALLOWANCETOHIMSELF is a free data retrieval call binding the contract method 0x06f355ad.
//
// Solidity: function CT_CANNOT_GIVE_ALLOWANCE_TO_HIMSELF() view returns(string)
func (_Errors *ErrorsSession) CTCANNOTGIVEALLOWANCETOHIMSELF() (string, error) {
	return _Errors.Contract.CTCANNOTGIVEALLOWANCETOHIMSELF(&_Errors.CallOpts)
}

// CTCANNOTGIVEALLOWANCETOHIMSELF is a free data retrieval call binding the contract method 0x06f355ad.
//
// Solidity: function CT_CANNOT_GIVE_ALLOWANCE_TO_HIMSELF() view returns(string)
func (_Errors *ErrorsCallerSession) CTCANNOTGIVEALLOWANCETOHIMSELF() (string, error) {
	return _Errors.Contract.CTCANNOTGIVEALLOWANCETOHIMSELF(&_Errors.CallOpts)
}

// CTINVALIDBURNAMOUNT is a free data retrieval call binding the contract method 0x2b9c57f6.
//
// Solidity: function CT_INVALID_BURN_AMOUNT() view returns(string)
func (_Errors *ErrorsCaller) CTINVALIDBURNAMOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CT_INVALID_BURN_AMOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CTINVALIDBURNAMOUNT is a free data retrieval call binding the contract method 0x2b9c57f6.
//
// Solidity: function CT_INVALID_BURN_AMOUNT() view returns(string)
func (_Errors *ErrorsSession) CTINVALIDBURNAMOUNT() (string, error) {
	return _Errors.Contract.CTINVALIDBURNAMOUNT(&_Errors.CallOpts)
}

// CTINVALIDBURNAMOUNT is a free data retrieval call binding the contract method 0x2b9c57f6.
//
// Solidity: function CT_INVALID_BURN_AMOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) CTINVALIDBURNAMOUNT() (string, error) {
	return _Errors.Contract.CTINVALIDBURNAMOUNT(&_Errors.CallOpts)
}

// CTINVALIDMINTAMOUNT is a free data retrieval call binding the contract method 0x1291a38b.
//
// Solidity: function CT_INVALID_MINT_AMOUNT() view returns(string)
func (_Errors *ErrorsCaller) CTINVALIDMINTAMOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CT_INVALID_MINT_AMOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CTINVALIDMINTAMOUNT is a free data retrieval call binding the contract method 0x1291a38b.
//
// Solidity: function CT_INVALID_MINT_AMOUNT() view returns(string)
func (_Errors *ErrorsSession) CTINVALIDMINTAMOUNT() (string, error) {
	return _Errors.Contract.CTINVALIDMINTAMOUNT(&_Errors.CallOpts)
}

// CTINVALIDMINTAMOUNT is a free data retrieval call binding the contract method 0x1291a38b.
//
// Solidity: function CT_INVALID_MINT_AMOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) CTINVALIDMINTAMOUNT() (string, error) {
	return _Errors.Contract.CTINVALIDMINTAMOUNT(&_Errors.CallOpts)
}

// CTTRANSFERAMOUNTNOTGT0 is a free data retrieval call binding the contract method 0xa2fbc8ad.
//
// Solidity: function CT_TRANSFER_AMOUNT_NOT_GT_0() view returns(string)
func (_Errors *ErrorsCaller) CTTRANSFERAMOUNTNOTGT0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "CT_TRANSFER_AMOUNT_NOT_GT_0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// CTTRANSFERAMOUNTNOTGT0 is a free data retrieval call binding the contract method 0xa2fbc8ad.
//
// Solidity: function CT_TRANSFER_AMOUNT_NOT_GT_0() view returns(string)
func (_Errors *ErrorsSession) CTTRANSFERAMOUNTNOTGT0() (string, error) {
	return _Errors.Contract.CTTRANSFERAMOUNTNOTGT0(&_Errors.CallOpts)
}

// CTTRANSFERAMOUNTNOTGT0 is a free data retrieval call binding the contract method 0xa2fbc8ad.
//
// Solidity: function CT_TRANSFER_AMOUNT_NOT_GT_0() view returns(string)
func (_Errors *ErrorsCallerSession) CTTRANSFERAMOUNTNOTGT0() (string, error) {
	return _Errors.Contract.CTTRANSFERAMOUNTNOTGT0(&_Errors.CallOpts)
}

// LPAPRINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x333e8ea8.
//
// Solidity: function LPAPR_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsCaller) LPAPRINVALIDADDRESSESPROVIDERID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPAPR_INVALID_ADDRESSES_PROVIDER_ID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPAPRINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x333e8ea8.
//
// Solidity: function LPAPR_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsSession) LPAPRINVALIDADDRESSESPROVIDERID() (string, error) {
	return _Errors.Contract.LPAPRINVALIDADDRESSESPROVIDERID(&_Errors.CallOpts)
}

// LPAPRINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x333e8ea8.
//
// Solidity: function LPAPR_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsCallerSession) LPAPRINVALIDADDRESSESPROVIDERID() (string, error) {
	return _Errors.Contract.LPAPRINVALIDADDRESSESPROVIDERID(&_Errors.CallOpts)
}

// LPAPRPROVIDERNOTREGISTERED is a free data retrieval call binding the contract method 0xd6f681b6.
//
// Solidity: function LPAPR_PROVIDER_NOT_REGISTERED() view returns(string)
func (_Errors *ErrorsCaller) LPAPRPROVIDERNOTREGISTERED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPAPR_PROVIDER_NOT_REGISTERED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPAPRPROVIDERNOTREGISTERED is a free data retrieval call binding the contract method 0xd6f681b6.
//
// Solidity: function LPAPR_PROVIDER_NOT_REGISTERED() view returns(string)
func (_Errors *ErrorsSession) LPAPRPROVIDERNOTREGISTERED() (string, error) {
	return _Errors.Contract.LPAPRPROVIDERNOTREGISTERED(&_Errors.CallOpts)
}

// LPAPRPROVIDERNOTREGISTERED is a free data retrieval call binding the contract method 0xd6f681b6.
//
// Solidity: function LPAPR_PROVIDER_NOT_REGISTERED() view returns(string)
func (_Errors *ErrorsCallerSession) LPAPRPROVIDERNOTREGISTERED() (string, error) {
	return _Errors.Contract.LPAPRPROVIDERNOTREGISTERED(&_Errors.CallOpts)
}

// LPCMCOLLATERALCANNOTBELIQUIDATED is a free data retrieval call binding the contract method 0x407374a4.
//
// Solidity: function LPCM_COLLATERAL_CANNOT_BE_LIQUIDATED() view returns(string)
func (_Errors *ErrorsCaller) LPCMCOLLATERALCANNOTBELIQUIDATED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPCM_COLLATERAL_CANNOT_BE_LIQUIDATED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCMCOLLATERALCANNOTBELIQUIDATED is a free data retrieval call binding the contract method 0x407374a4.
//
// Solidity: function LPCM_COLLATERAL_CANNOT_BE_LIQUIDATED() view returns(string)
func (_Errors *ErrorsSession) LPCMCOLLATERALCANNOTBELIQUIDATED() (string, error) {
	return _Errors.Contract.LPCMCOLLATERALCANNOTBELIQUIDATED(&_Errors.CallOpts)
}

// LPCMCOLLATERALCANNOTBELIQUIDATED is a free data retrieval call binding the contract method 0x407374a4.
//
// Solidity: function LPCM_COLLATERAL_CANNOT_BE_LIQUIDATED() view returns(string)
func (_Errors *ErrorsCallerSession) LPCMCOLLATERALCANNOTBELIQUIDATED() (string, error) {
	return _Errors.Contract.LPCMCOLLATERALCANNOTBELIQUIDATED(&_Errors.CallOpts)
}

// LPCMHEALTHFACTORNOTBELOWTHRESHOLD is a free data retrieval call binding the contract method 0xfb681def.
//
// Solidity: function LPCM_HEALTH_FACTOR_NOT_BELOW_THRESHOLD() view returns(string)
func (_Errors *ErrorsCaller) LPCMHEALTHFACTORNOTBELOWTHRESHOLD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPCM_HEALTH_FACTOR_NOT_BELOW_THRESHOLD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCMHEALTHFACTORNOTBELOWTHRESHOLD is a free data retrieval call binding the contract method 0xfb681def.
//
// Solidity: function LPCM_HEALTH_FACTOR_NOT_BELOW_THRESHOLD() view returns(string)
func (_Errors *ErrorsSession) LPCMHEALTHFACTORNOTBELOWTHRESHOLD() (string, error) {
	return _Errors.Contract.LPCMHEALTHFACTORNOTBELOWTHRESHOLD(&_Errors.CallOpts)
}

// LPCMHEALTHFACTORNOTBELOWTHRESHOLD is a free data retrieval call binding the contract method 0xfb681def.
//
// Solidity: function LPCM_HEALTH_FACTOR_NOT_BELOW_THRESHOLD() view returns(string)
func (_Errors *ErrorsCallerSession) LPCMHEALTHFACTORNOTBELOWTHRESHOLD() (string, error) {
	return _Errors.Contract.LPCMHEALTHFACTORNOTBELOWTHRESHOLD(&_Errors.CallOpts)
}

// LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE is a free data retrieval call binding the contract method 0xcc5fc44c.
//
// Solidity: function LPCM_NOT_ENOUGH_LIQUIDITY_TO_LIQUIDATE() view returns(string)
func (_Errors *ErrorsCaller) LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPCM_NOT_ENOUGH_LIQUIDITY_TO_LIQUIDATE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE is a free data retrieval call binding the contract method 0xcc5fc44c.
//
// Solidity: function LPCM_NOT_ENOUGH_LIQUIDITY_TO_LIQUIDATE() view returns(string)
func (_Errors *ErrorsSession) LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE() (string, error) {
	return _Errors.Contract.LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE(&_Errors.CallOpts)
}

// LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE is a free data retrieval call binding the contract method 0xcc5fc44c.
//
// Solidity: function LPCM_NOT_ENOUGH_LIQUIDITY_TO_LIQUIDATE() view returns(string)
func (_Errors *ErrorsCallerSession) LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE() (string, error) {
	return _Errors.Contract.LPCMNOTENOUGHLIQUIDITYTOLIQUIDATE(&_Errors.CallOpts)
}

// LPCMNOERRORS is a free data retrieval call binding the contract method 0x1ea7c604.
//
// Solidity: function LPCM_NO_ERRORS() view returns(string)
func (_Errors *ErrorsCaller) LPCMNOERRORS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPCM_NO_ERRORS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCMNOERRORS is a free data retrieval call binding the contract method 0x1ea7c604.
//
// Solidity: function LPCM_NO_ERRORS() view returns(string)
func (_Errors *ErrorsSession) LPCMNOERRORS() (string, error) {
	return _Errors.Contract.LPCMNOERRORS(&_Errors.CallOpts)
}

// LPCMNOERRORS is a free data retrieval call binding the contract method 0x1ea7c604.
//
// Solidity: function LPCM_NO_ERRORS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCMNOERRORS() (string, error) {
	return _Errors.Contract.LPCMNOERRORS(&_Errors.CallOpts)
}

// LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER is a free data retrieval call binding the contract method 0x71a629da.
//
// Solidity: function LPCM_SPECIFIED_CURRENCY_NOT_BORROWED_BY_USER() view returns(string)
func (_Errors *ErrorsCaller) LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPCM_SPECIFIED_CURRENCY_NOT_BORROWED_BY_USER")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER is a free data retrieval call binding the contract method 0x71a629da.
//
// Solidity: function LPCM_SPECIFIED_CURRENCY_NOT_BORROWED_BY_USER() view returns(string)
func (_Errors *ErrorsSession) LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER() (string, error) {
	return _Errors.Contract.LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER(&_Errors.CallOpts)
}

// LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER is a free data retrieval call binding the contract method 0x71a629da.
//
// Solidity: function LPCM_SPECIFIED_CURRENCY_NOT_BORROWED_BY_USER() view returns(string)
func (_Errors *ErrorsCallerSession) LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER() (string, error) {
	return _Errors.Contract.LPCMSPECIFIEDCURRENCYNOTBORROWEDBYUSER(&_Errors.CallOpts)
}

// LPCCALLERNOTEMERGENCYADMIN is a free data retrieval call binding the contract method 0xd3e370ee.
//
// Solidity: function LPC_CALLER_NOT_EMERGENCY_ADMIN() view returns(string)
func (_Errors *ErrorsCaller) LPCCALLERNOTEMERGENCYADMIN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_CALLER_NOT_EMERGENCY_ADMIN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCCALLERNOTEMERGENCYADMIN is a free data retrieval call binding the contract method 0xd3e370ee.
//
// Solidity: function LPC_CALLER_NOT_EMERGENCY_ADMIN() view returns(string)
func (_Errors *ErrorsSession) LPCCALLERNOTEMERGENCYADMIN() (string, error) {
	return _Errors.Contract.LPCCALLERNOTEMERGENCYADMIN(&_Errors.CallOpts)
}

// LPCCALLERNOTEMERGENCYADMIN is a free data retrieval call binding the contract method 0xd3e370ee.
//
// Solidity: function LPC_CALLER_NOT_EMERGENCY_ADMIN() view returns(string)
func (_Errors *ErrorsCallerSession) LPCCALLERNOTEMERGENCYADMIN() (string, error) {
	return _Errors.Contract.LPCCALLERNOTEMERGENCYADMIN(&_Errors.CallOpts)
}

// LPCINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x5a9786d4.
//
// Solidity: function LPC_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDADDRESSESPROVIDERID(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_ADDRESSES_PROVIDER_ID")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x5a9786d4.
//
// Solidity: function LPC_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDADDRESSESPROVIDERID() (string, error) {
	return _Errors.Contract.LPCINVALIDADDRESSESPROVIDERID(&_Errors.CallOpts)
}

// LPCINVALIDADDRESSESPROVIDERID is a free data retrieval call binding the contract method 0x5a9786d4.
//
// Solidity: function LPC_INVALID_ADDRESSES_PROVIDER_ID() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDADDRESSESPROVIDERID() (string, error) {
	return _Errors.Contract.LPCINVALIDADDRESSESPROVIDERID(&_Errors.CallOpts)
}

// LPCINVALIDATOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xf0473259.
//
// Solidity: function LPC_INVALID_ATOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDATOKENPOOLADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_ATOKEN_POOL_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDATOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xf0473259.
//
// Solidity: function LPC_INVALID_ATOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDATOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDATOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDATOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xf0473259.
//
// Solidity: function LPC_INVALID_ATOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDATOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDATOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDCONFIGURATION is a free data retrieval call binding the contract method 0x47d25300.
//
// Solidity: function LPC_INVALID_CONFIGURATION() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDCONFIGURATION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_CONFIGURATION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDCONFIGURATION is a free data retrieval call binding the contract method 0x47d25300.
//
// Solidity: function LPC_INVALID_CONFIGURATION() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDCONFIGURATION() (string, error) {
	return _Errors.Contract.LPCINVALIDCONFIGURATION(&_Errors.CallOpts)
}

// LPCINVALIDCONFIGURATION is a free data retrieval call binding the contract method 0x47d25300.
//
// Solidity: function LPC_INVALID_CONFIGURATION() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDCONFIGURATION() (string, error) {
	return _Errors.Contract.LPCINVALIDCONFIGURATION(&_Errors.CallOpts)
}

// LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xe0d7dfd7.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_STABLE_DEBT_TOKEN_POOL_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xe0d7dfd7.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xe0d7dfd7.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDSTABLEDEBTTOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x41b40ba5.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_STABLE_DEBT_TOKEN_UNDERLYING_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x41b40ba5.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x41b40ba5.
//
// Solidity: function LPC_INVALID_STABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDSTABLEDEBTTOKENUNDERLYINGADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xc09e2618.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_VARIABLE_DEBT_TOKEN_POOL_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xc09e2618.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS is a free data retrieval call binding the contract method 0xc09e2618.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_POOL_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDVARIABLEDEBTTOKENPOOLADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x3872b0ad.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsCaller) LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_INVALID_VARIABLE_DEBT_TOKEN_UNDERLYING_ADDRESS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x3872b0ad.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsSession) LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS(&_Errors.CallOpts)
}

// LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS is a free data retrieval call binding the contract method 0x3872b0ad.
//
// Solidity: function LPC_INVALID_VARIABLE_DEBT_TOKEN_UNDERLYING_ADDRESS() view returns(string)
func (_Errors *ErrorsCallerSession) LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS() (string, error) {
	return _Errors.Contract.LPCINVALIDVARIABLEDEBTTOKENUNDERLYINGADDRESS(&_Errors.CallOpts)
}

// LPCRESERVELIQUIDITYNOT0 is a free data retrieval call binding the contract method 0x65344799.
//
// Solidity: function LPC_RESERVE_LIQUIDITY_NOT_0() view returns(string)
func (_Errors *ErrorsCaller) LPCRESERVELIQUIDITYNOT0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LPC_RESERVE_LIQUIDITY_NOT_0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCRESERVELIQUIDITYNOT0 is a free data retrieval call binding the contract method 0x65344799.
//
// Solidity: function LPC_RESERVE_LIQUIDITY_NOT_0() view returns(string)
func (_Errors *ErrorsSession) LPCRESERVELIQUIDITYNOT0() (string, error) {
	return _Errors.Contract.LPCRESERVELIQUIDITYNOT0(&_Errors.CallOpts)
}

// LPCRESERVELIQUIDITYNOT0 is a free data retrieval call binding the contract method 0x65344799.
//
// Solidity: function LPC_RESERVE_LIQUIDITY_NOT_0() view returns(string)
func (_Errors *ErrorsCallerSession) LPCRESERVELIQUIDITYNOT0() (string, error) {
	return _Errors.Contract.LPCRESERVELIQUIDITYNOT0(&_Errors.CallOpts)
}

// LPCALLERMUSTBEANATOKEN is a free data retrieval call binding the contract method 0xeca85d3a.
//
// Solidity: function LP_CALLER_MUST_BE_AN_ATOKEN() view returns(string)
func (_Errors *ErrorsCaller) LPCALLERMUSTBEANATOKEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_CALLER_MUST_BE_AN_ATOKEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCALLERMUSTBEANATOKEN is a free data retrieval call binding the contract method 0xeca85d3a.
//
// Solidity: function LP_CALLER_MUST_BE_AN_ATOKEN() view returns(string)
func (_Errors *ErrorsSession) LPCALLERMUSTBEANATOKEN() (string, error) {
	return _Errors.Contract.LPCALLERMUSTBEANATOKEN(&_Errors.CallOpts)
}

// LPCALLERMUSTBEANATOKEN is a free data retrieval call binding the contract method 0xeca85d3a.
//
// Solidity: function LP_CALLER_MUST_BE_AN_ATOKEN() view returns(string)
func (_Errors *ErrorsCallerSession) LPCALLERMUSTBEANATOKEN() (string, error) {
	return _Errors.Contract.LPCALLERMUSTBEANATOKEN(&_Errors.CallOpts)
}

// LPCALLERNOTLENDINGPOOLCONFIGURATOR is a free data retrieval call binding the contract method 0xb36a2cf3.
//
// Solidity: function LP_CALLER_NOT_LENDING_POOL_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsCaller) LPCALLERNOTLENDINGPOOLCONFIGURATOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_CALLER_NOT_LENDING_POOL_CONFIGURATOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPCALLERNOTLENDINGPOOLCONFIGURATOR is a free data retrieval call binding the contract method 0xb36a2cf3.
//
// Solidity: function LP_CALLER_NOT_LENDING_POOL_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsSession) LPCALLERNOTLENDINGPOOLCONFIGURATOR() (string, error) {
	return _Errors.Contract.LPCALLERNOTLENDINGPOOLCONFIGURATOR(&_Errors.CallOpts)
}

// LPCALLERNOTLENDINGPOOLCONFIGURATOR is a free data retrieval call binding the contract method 0xb36a2cf3.
//
// Solidity: function LP_CALLER_NOT_LENDING_POOL_CONFIGURATOR() view returns(string)
func (_Errors *ErrorsCallerSession) LPCALLERNOTLENDINGPOOLCONFIGURATOR() (string, error) {
	return _Errors.Contract.LPCALLERNOTLENDINGPOOLCONFIGURATOR(&_Errors.CallOpts)
}

// LPFAILEDCOLLATERALSWAP is a free data retrieval call binding the contract method 0x55bab12c.
//
// Solidity: function LP_FAILED_COLLATERAL_SWAP() view returns(string)
func (_Errors *ErrorsCaller) LPFAILEDCOLLATERALSWAP(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_FAILED_COLLATERAL_SWAP")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPFAILEDCOLLATERALSWAP is a free data retrieval call binding the contract method 0x55bab12c.
//
// Solidity: function LP_FAILED_COLLATERAL_SWAP() view returns(string)
func (_Errors *ErrorsSession) LPFAILEDCOLLATERALSWAP() (string, error) {
	return _Errors.Contract.LPFAILEDCOLLATERALSWAP(&_Errors.CallOpts)
}

// LPFAILEDCOLLATERALSWAP is a free data retrieval call binding the contract method 0x55bab12c.
//
// Solidity: function LP_FAILED_COLLATERAL_SWAP() view returns(string)
func (_Errors *ErrorsCallerSession) LPFAILEDCOLLATERALSWAP() (string, error) {
	return _Errors.Contract.LPFAILEDCOLLATERALSWAP(&_Errors.CallOpts)
}

// LPFAILEDREPAYWITHCOLLATERAL is a free data retrieval call binding the contract method 0xe6632748.
//
// Solidity: function LP_FAILED_REPAY_WITH_COLLATERAL() view returns(string)
func (_Errors *ErrorsCaller) LPFAILEDREPAYWITHCOLLATERAL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_FAILED_REPAY_WITH_COLLATERAL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPFAILEDREPAYWITHCOLLATERAL is a free data retrieval call binding the contract method 0xe6632748.
//
// Solidity: function LP_FAILED_REPAY_WITH_COLLATERAL() view returns(string)
func (_Errors *ErrorsSession) LPFAILEDREPAYWITHCOLLATERAL() (string, error) {
	return _Errors.Contract.LPFAILEDREPAYWITHCOLLATERAL(&_Errors.CallOpts)
}

// LPFAILEDREPAYWITHCOLLATERAL is a free data retrieval call binding the contract method 0xe6632748.
//
// Solidity: function LP_FAILED_REPAY_WITH_COLLATERAL() view returns(string)
func (_Errors *ErrorsCallerSession) LPFAILEDREPAYWITHCOLLATERAL() (string, error) {
	return _Errors.Contract.LPFAILEDREPAYWITHCOLLATERAL(&_Errors.CallOpts)
}

// LPINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0x2b34c349.
//
// Solidity: function LP_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsCaller) LPINCONSISTENTFLASHLOANPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INCONSISTENT_FLASHLOAN_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0x2b34c349.
//
// Solidity: function LP_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsSession) LPINCONSISTENTFLASHLOANPARAMS() (string, error) {
	return _Errors.Contract.LPINCONSISTENTFLASHLOANPARAMS(&_Errors.CallOpts)
}

// LPINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0x2b34c349.
//
// Solidity: function LP_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsCallerSession) LPINCONSISTENTFLASHLOANPARAMS() (string, error) {
	return _Errors.Contract.LPINCONSISTENTFLASHLOANPARAMS(&_Errors.CallOpts)
}

// LPINCONSISTENTPARAMSLENGTH is a free data retrieval call binding the contract method 0xb72e40c7.
//
// Solidity: function LP_INCONSISTENT_PARAMS_LENGTH() view returns(string)
func (_Errors *ErrorsCaller) LPINCONSISTENTPARAMSLENGTH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INCONSISTENT_PARAMS_LENGTH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINCONSISTENTPARAMSLENGTH is a free data retrieval call binding the contract method 0xb72e40c7.
//
// Solidity: function LP_INCONSISTENT_PARAMS_LENGTH() view returns(string)
func (_Errors *ErrorsSession) LPINCONSISTENTPARAMSLENGTH() (string, error) {
	return _Errors.Contract.LPINCONSISTENTPARAMSLENGTH(&_Errors.CallOpts)
}

// LPINCONSISTENTPARAMSLENGTH is a free data retrieval call binding the contract method 0xb72e40c7.
//
// Solidity: function LP_INCONSISTENT_PARAMS_LENGTH() view returns(string)
func (_Errors *ErrorsCallerSession) LPINCONSISTENTPARAMSLENGTH() (string, error) {
	return _Errors.Contract.LPINCONSISTENTPARAMSLENGTH(&_Errors.CallOpts)
}

// LPINCONSISTENTPROTOCOLACTUALBALANCE is a free data retrieval call binding the contract method 0x0b8fd588.
//
// Solidity: function LP_INCONSISTENT_PROTOCOL_ACTUAL_BALANCE() view returns(string)
func (_Errors *ErrorsCaller) LPINCONSISTENTPROTOCOLACTUALBALANCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INCONSISTENT_PROTOCOL_ACTUAL_BALANCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINCONSISTENTPROTOCOLACTUALBALANCE is a free data retrieval call binding the contract method 0x0b8fd588.
//
// Solidity: function LP_INCONSISTENT_PROTOCOL_ACTUAL_BALANCE() view returns(string)
func (_Errors *ErrorsSession) LPINCONSISTENTPROTOCOLACTUALBALANCE() (string, error) {
	return _Errors.Contract.LPINCONSISTENTPROTOCOLACTUALBALANCE(&_Errors.CallOpts)
}

// LPINCONSISTENTPROTOCOLACTUALBALANCE is a free data retrieval call binding the contract method 0x0b8fd588.
//
// Solidity: function LP_INCONSISTENT_PROTOCOL_ACTUAL_BALANCE() view returns(string)
func (_Errors *ErrorsCallerSession) LPINCONSISTENTPROTOCOLACTUALBALANCE() (string, error) {
	return _Errors.Contract.LPINCONSISTENTPROTOCOLACTUALBALANCE(&_Errors.CallOpts)
}

// LPINTERESTRATEREBALANCECONDITIONSNOTMET is a free data retrieval call binding the contract method 0xb89652cd.
//
// Solidity: function LP_INTEREST_RATE_REBALANCE_CONDITIONS_NOT_MET() view returns(string)
func (_Errors *ErrorsCaller) LPINTERESTRATEREBALANCECONDITIONSNOTMET(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INTEREST_RATE_REBALANCE_CONDITIONS_NOT_MET")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINTERESTRATEREBALANCECONDITIONSNOTMET is a free data retrieval call binding the contract method 0xb89652cd.
//
// Solidity: function LP_INTEREST_RATE_REBALANCE_CONDITIONS_NOT_MET() view returns(string)
func (_Errors *ErrorsSession) LPINTERESTRATEREBALANCECONDITIONSNOTMET() (string, error) {
	return _Errors.Contract.LPINTERESTRATEREBALANCECONDITIONSNOTMET(&_Errors.CallOpts)
}

// LPINTERESTRATEREBALANCECONDITIONSNOTMET is a free data retrieval call binding the contract method 0xb89652cd.
//
// Solidity: function LP_INTEREST_RATE_REBALANCE_CONDITIONS_NOT_MET() view returns(string)
func (_Errors *ErrorsCallerSession) LPINTERESTRATEREBALANCECONDITIONSNOTMET() (string, error) {
	return _Errors.Contract.LPINTERESTRATEREBALANCECONDITIONSNOTMET(&_Errors.CallOpts)
}

// LPINVALIDEQUALASSETSTOSWAP is a free data retrieval call binding the contract method 0xc2d628df.
//
// Solidity: function LP_INVALID_EQUAL_ASSETS_TO_SWAP() view returns(string)
func (_Errors *ErrorsCaller) LPINVALIDEQUALASSETSTOSWAP(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INVALID_EQUAL_ASSETS_TO_SWAP")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINVALIDEQUALASSETSTOSWAP is a free data retrieval call binding the contract method 0xc2d628df.
//
// Solidity: function LP_INVALID_EQUAL_ASSETS_TO_SWAP() view returns(string)
func (_Errors *ErrorsSession) LPINVALIDEQUALASSETSTOSWAP() (string, error) {
	return _Errors.Contract.LPINVALIDEQUALASSETSTOSWAP(&_Errors.CallOpts)
}

// LPINVALIDEQUALASSETSTOSWAP is a free data retrieval call binding the contract method 0xc2d628df.
//
// Solidity: function LP_INVALID_EQUAL_ASSETS_TO_SWAP() view returns(string)
func (_Errors *ErrorsCallerSession) LPINVALIDEQUALASSETSTOSWAP() (string, error) {
	return _Errors.Contract.LPINVALIDEQUALASSETSTOSWAP(&_Errors.CallOpts)
}

// LPINVALIDFLASHLOANMODE is a free data retrieval call binding the contract method 0xe7bf91b3.
//
// Solidity: function LP_INVALID_FLASHLOAN_MODE() view returns(string)
func (_Errors *ErrorsCaller) LPINVALIDFLASHLOANMODE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INVALID_FLASHLOAN_MODE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINVALIDFLASHLOANMODE is a free data retrieval call binding the contract method 0xe7bf91b3.
//
// Solidity: function LP_INVALID_FLASHLOAN_MODE() view returns(string)
func (_Errors *ErrorsSession) LPINVALIDFLASHLOANMODE() (string, error) {
	return _Errors.Contract.LPINVALIDFLASHLOANMODE(&_Errors.CallOpts)
}

// LPINVALIDFLASHLOANMODE is a free data retrieval call binding the contract method 0xe7bf91b3.
//
// Solidity: function LP_INVALID_FLASHLOAN_MODE() view returns(string)
func (_Errors *ErrorsCallerSession) LPINVALIDFLASHLOANMODE() (string, error) {
	return _Errors.Contract.LPINVALIDFLASHLOANMODE(&_Errors.CallOpts)
}

// LPINVALIDFLASHLOANEXECUTORRETURN is a free data retrieval call binding the contract method 0xd44e8e88.
//
// Solidity: function LP_INVALID_FLASH_LOAN_EXECUTOR_RETURN() view returns(string)
func (_Errors *ErrorsCaller) LPINVALIDFLASHLOANEXECUTORRETURN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_INVALID_FLASH_LOAN_EXECUTOR_RETURN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPINVALIDFLASHLOANEXECUTORRETURN is a free data retrieval call binding the contract method 0xd44e8e88.
//
// Solidity: function LP_INVALID_FLASH_LOAN_EXECUTOR_RETURN() view returns(string)
func (_Errors *ErrorsSession) LPINVALIDFLASHLOANEXECUTORRETURN() (string, error) {
	return _Errors.Contract.LPINVALIDFLASHLOANEXECUTORRETURN(&_Errors.CallOpts)
}

// LPINVALIDFLASHLOANEXECUTORRETURN is a free data retrieval call binding the contract method 0xd44e8e88.
//
// Solidity: function LP_INVALID_FLASH_LOAN_EXECUTOR_RETURN() view returns(string)
func (_Errors *ErrorsCallerSession) LPINVALIDFLASHLOANEXECUTORRETURN() (string, error) {
	return _Errors.Contract.LPINVALIDFLASHLOANEXECUTORRETURN(&_Errors.CallOpts)
}

// LPISPAUSED is a free data retrieval call binding the contract method 0xd57bb964.
//
// Solidity: function LP_IS_PAUSED() view returns(string)
func (_Errors *ErrorsCaller) LPISPAUSED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_IS_PAUSED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPISPAUSED is a free data retrieval call binding the contract method 0xd57bb964.
//
// Solidity: function LP_IS_PAUSED() view returns(string)
func (_Errors *ErrorsSession) LPISPAUSED() (string, error) {
	return _Errors.Contract.LPISPAUSED(&_Errors.CallOpts)
}

// LPISPAUSED is a free data retrieval call binding the contract method 0xd57bb964.
//
// Solidity: function LP_IS_PAUSED() view returns(string)
func (_Errors *ErrorsCallerSession) LPISPAUSED() (string, error) {
	return _Errors.Contract.LPISPAUSED(&_Errors.CallOpts)
}

// LPLIQUIDATIONCALLFAILED is a free data retrieval call binding the contract method 0x4a529f91.
//
// Solidity: function LP_LIQUIDATION_CALL_FAILED() view returns(string)
func (_Errors *ErrorsCaller) LPLIQUIDATIONCALLFAILED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_LIQUIDATION_CALL_FAILED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPLIQUIDATIONCALLFAILED is a free data retrieval call binding the contract method 0x4a529f91.
//
// Solidity: function LP_LIQUIDATION_CALL_FAILED() view returns(string)
func (_Errors *ErrorsSession) LPLIQUIDATIONCALLFAILED() (string, error) {
	return _Errors.Contract.LPLIQUIDATIONCALLFAILED(&_Errors.CallOpts)
}

// LPLIQUIDATIONCALLFAILED is a free data retrieval call binding the contract method 0x4a529f91.
//
// Solidity: function LP_LIQUIDATION_CALL_FAILED() view returns(string)
func (_Errors *ErrorsCallerSession) LPLIQUIDATIONCALLFAILED() (string, error) {
	return _Errors.Contract.LPLIQUIDATIONCALLFAILED(&_Errors.CallOpts)
}

// LPNOTCONTRACT is a free data retrieval call binding the contract method 0x637a5a12.
//
// Solidity: function LP_NOT_CONTRACT() view returns(string)
func (_Errors *ErrorsCaller) LPNOTCONTRACT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_NOT_CONTRACT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPNOTCONTRACT is a free data retrieval call binding the contract method 0x637a5a12.
//
// Solidity: function LP_NOT_CONTRACT() view returns(string)
func (_Errors *ErrorsSession) LPNOTCONTRACT() (string, error) {
	return _Errors.Contract.LPNOTCONTRACT(&_Errors.CallOpts)
}

// LPNOTCONTRACT is a free data retrieval call binding the contract method 0x637a5a12.
//
// Solidity: function LP_NOT_CONTRACT() view returns(string)
func (_Errors *ErrorsCallerSession) LPNOTCONTRACT() (string, error) {
	return _Errors.Contract.LPNOTCONTRACT(&_Errors.CallOpts)
}

// LPNOTENOUGHLIQUIDITYTOBORROW is a free data retrieval call binding the contract method 0x22a6f08e.
//
// Solidity: function LP_NOT_ENOUGH_LIQUIDITY_TO_BORROW() view returns(string)
func (_Errors *ErrorsCaller) LPNOTENOUGHLIQUIDITYTOBORROW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_NOT_ENOUGH_LIQUIDITY_TO_BORROW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPNOTENOUGHLIQUIDITYTOBORROW is a free data retrieval call binding the contract method 0x22a6f08e.
//
// Solidity: function LP_NOT_ENOUGH_LIQUIDITY_TO_BORROW() view returns(string)
func (_Errors *ErrorsSession) LPNOTENOUGHLIQUIDITYTOBORROW() (string, error) {
	return _Errors.Contract.LPNOTENOUGHLIQUIDITYTOBORROW(&_Errors.CallOpts)
}

// LPNOTENOUGHLIQUIDITYTOBORROW is a free data retrieval call binding the contract method 0x22a6f08e.
//
// Solidity: function LP_NOT_ENOUGH_LIQUIDITY_TO_BORROW() view returns(string)
func (_Errors *ErrorsCallerSession) LPNOTENOUGHLIQUIDITYTOBORROW() (string, error) {
	return _Errors.Contract.LPNOTENOUGHLIQUIDITYTOBORROW(&_Errors.CallOpts)
}

// LPNOTENOUGHSTABLEBORROWBALANCE is a free data retrieval call binding the contract method 0x44dc4f70.
//
// Solidity: function LP_NOT_ENOUGH_STABLE_BORROW_BALANCE() view returns(string)
func (_Errors *ErrorsCaller) LPNOTENOUGHSTABLEBORROWBALANCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_NOT_ENOUGH_STABLE_BORROW_BALANCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPNOTENOUGHSTABLEBORROWBALANCE is a free data retrieval call binding the contract method 0x44dc4f70.
//
// Solidity: function LP_NOT_ENOUGH_STABLE_BORROW_BALANCE() view returns(string)
func (_Errors *ErrorsSession) LPNOTENOUGHSTABLEBORROWBALANCE() (string, error) {
	return _Errors.Contract.LPNOTENOUGHSTABLEBORROWBALANCE(&_Errors.CallOpts)
}

// LPNOTENOUGHSTABLEBORROWBALANCE is a free data retrieval call binding the contract method 0x44dc4f70.
//
// Solidity: function LP_NOT_ENOUGH_STABLE_BORROW_BALANCE() view returns(string)
func (_Errors *ErrorsCallerSession) LPNOTENOUGHSTABLEBORROWBALANCE() (string, error) {
	return _Errors.Contract.LPNOTENOUGHSTABLEBORROWBALANCE(&_Errors.CallOpts)
}

// LPNOMORERESERVESALLOWED is a free data retrieval call binding the contract method 0x76f19030.
//
// Solidity: function LP_NO_MORE_RESERVES_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) LPNOMORERESERVESALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_NO_MORE_RESERVES_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPNOMORERESERVESALLOWED is a free data retrieval call binding the contract method 0x76f19030.
//
// Solidity: function LP_NO_MORE_RESERVES_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) LPNOMORERESERVESALLOWED() (string, error) {
	return _Errors.Contract.LPNOMORERESERVESALLOWED(&_Errors.CallOpts)
}

// LPNOMORERESERVESALLOWED is a free data retrieval call binding the contract method 0x76f19030.
//
// Solidity: function LP_NO_MORE_RESERVES_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) LPNOMORERESERVESALLOWED() (string, error) {
	return _Errors.Contract.LPNOMORERESERVESALLOWED(&_Errors.CallOpts)
}

// LPREENTRANCYNOTALLOWED is a free data retrieval call binding the contract method 0xf902735d.
//
// Solidity: function LP_REENTRANCY_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) LPREENTRANCYNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_REENTRANCY_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPREENTRANCYNOTALLOWED is a free data retrieval call binding the contract method 0xf902735d.
//
// Solidity: function LP_REENTRANCY_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) LPREENTRANCYNOTALLOWED() (string, error) {
	return _Errors.Contract.LPREENTRANCYNOTALLOWED(&_Errors.CallOpts)
}

// LPREENTRANCYNOTALLOWED is a free data retrieval call binding the contract method 0xf902735d.
//
// Solidity: function LP_REENTRANCY_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) LPREENTRANCYNOTALLOWED() (string, error) {
	return _Errors.Contract.LPREENTRANCYNOTALLOWED(&_Errors.CallOpts)
}

// LPREQUESTEDAMOUNTTOOSMALL is a free data retrieval call binding the contract method 0x390f34ba.
//
// Solidity: function LP_REQUESTED_AMOUNT_TOO_SMALL() view returns(string)
func (_Errors *ErrorsCaller) LPREQUESTEDAMOUNTTOOSMALL(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "LP_REQUESTED_AMOUNT_TOO_SMALL")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// LPREQUESTEDAMOUNTTOOSMALL is a free data retrieval call binding the contract method 0x390f34ba.
//
// Solidity: function LP_REQUESTED_AMOUNT_TOO_SMALL() view returns(string)
func (_Errors *ErrorsSession) LPREQUESTEDAMOUNTTOOSMALL() (string, error) {
	return _Errors.Contract.LPREQUESTEDAMOUNTTOOSMALL(&_Errors.CallOpts)
}

// LPREQUESTEDAMOUNTTOOSMALL is a free data retrieval call binding the contract method 0x390f34ba.
//
// Solidity: function LP_REQUESTED_AMOUNT_TOO_SMALL() view returns(string)
func (_Errors *ErrorsCallerSession) LPREQUESTEDAMOUNTTOOSMALL() (string, error) {
	return _Errors.Contract.LPREQUESTEDAMOUNTTOOSMALL(&_Errors.CallOpts)
}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) MATHADDITIONOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_ADDITION_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) MATHADDITIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHADDITIONOVERFLOW(&_Errors.CallOpts)
}

// MATHADDITIONOVERFLOW is a free data retrieval call binding the contract method 0x0f5ee482.
//
// Solidity: function MATH_ADDITION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) MATHADDITIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHADDITIONOVERFLOW(&_Errors.CallOpts)
}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsCaller) MATHDIVISIONBYZERO(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_DIVISION_BY_ZERO")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsSession) MATHDIVISIONBYZERO() (string, error) {
	return _Errors.Contract.MATHDIVISIONBYZERO(&_Errors.CallOpts)
}

// MATHDIVISIONBYZERO is a free data retrieval call binding the contract method 0x4349e3d8.
//
// Solidity: function MATH_DIVISION_BY_ZERO() view returns(string)
func (_Errors *ErrorsCallerSession) MATHDIVISIONBYZERO() (string, error) {
	return _Errors.Contract.MATHDIVISIONBYZERO(&_Errors.CallOpts)
}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) MATHMULTIPLICATIONOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "MATH_MULTIPLICATION_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) MATHMULTIPLICATIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHMULTIPLICATIONOVERFLOW(&_Errors.CallOpts)
}

// MATHMULTIPLICATIONOVERFLOW is a free data retrieval call binding the contract method 0x029d2344.
//
// Solidity: function MATH_MULTIPLICATION_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) MATHMULTIPLICATIONOVERFLOW() (string, error) {
	return _Errors.Contract.MATHMULTIPLICATIONOVERFLOW(&_Errors.CallOpts)
}

// RCINVALIDDECIMALS is a free data retrieval call binding the contract method 0x3f5d6ec8.
//
// Solidity: function RC_INVALID_DECIMALS() view returns(string)
func (_Errors *ErrorsCaller) RCINVALIDDECIMALS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RC_INVALID_DECIMALS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RCINVALIDDECIMALS is a free data retrieval call binding the contract method 0x3f5d6ec8.
//
// Solidity: function RC_INVALID_DECIMALS() view returns(string)
func (_Errors *ErrorsSession) RCINVALIDDECIMALS() (string, error) {
	return _Errors.Contract.RCINVALIDDECIMALS(&_Errors.CallOpts)
}

// RCINVALIDDECIMALS is a free data retrieval call binding the contract method 0x3f5d6ec8.
//
// Solidity: function RC_INVALID_DECIMALS() view returns(string)
func (_Errors *ErrorsCallerSession) RCINVALIDDECIMALS() (string, error) {
	return _Errors.Contract.RCINVALIDDECIMALS(&_Errors.CallOpts)
}

// RCINVALIDLIQBONUS is a free data retrieval call binding the contract method 0x5e869ff1.
//
// Solidity: function RC_INVALID_LIQ_BONUS() view returns(string)
func (_Errors *ErrorsCaller) RCINVALIDLIQBONUS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RC_INVALID_LIQ_BONUS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RCINVALIDLIQBONUS is a free data retrieval call binding the contract method 0x5e869ff1.
//
// Solidity: function RC_INVALID_LIQ_BONUS() view returns(string)
func (_Errors *ErrorsSession) RCINVALIDLIQBONUS() (string, error) {
	return _Errors.Contract.RCINVALIDLIQBONUS(&_Errors.CallOpts)
}

// RCINVALIDLIQBONUS is a free data retrieval call binding the contract method 0x5e869ff1.
//
// Solidity: function RC_INVALID_LIQ_BONUS() view returns(string)
func (_Errors *ErrorsCallerSession) RCINVALIDLIQBONUS() (string, error) {
	return _Errors.Contract.RCINVALIDLIQBONUS(&_Errors.CallOpts)
}

// RCINVALIDLIQTHRESHOLD is a free data retrieval call binding the contract method 0xbd013f5b.
//
// Solidity: function RC_INVALID_LIQ_THRESHOLD() view returns(string)
func (_Errors *ErrorsCaller) RCINVALIDLIQTHRESHOLD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RC_INVALID_LIQ_THRESHOLD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RCINVALIDLIQTHRESHOLD is a free data retrieval call binding the contract method 0xbd013f5b.
//
// Solidity: function RC_INVALID_LIQ_THRESHOLD() view returns(string)
func (_Errors *ErrorsSession) RCINVALIDLIQTHRESHOLD() (string, error) {
	return _Errors.Contract.RCINVALIDLIQTHRESHOLD(&_Errors.CallOpts)
}

// RCINVALIDLIQTHRESHOLD is a free data retrieval call binding the contract method 0xbd013f5b.
//
// Solidity: function RC_INVALID_LIQ_THRESHOLD() view returns(string)
func (_Errors *ErrorsCallerSession) RCINVALIDLIQTHRESHOLD() (string, error) {
	return _Errors.Contract.RCINVALIDLIQTHRESHOLD(&_Errors.CallOpts)
}

// RCINVALIDLTV is a free data retrieval call binding the contract method 0x614cf6a1.
//
// Solidity: function RC_INVALID_LTV() view returns(string)
func (_Errors *ErrorsCaller) RCINVALIDLTV(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RC_INVALID_LTV")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RCINVALIDLTV is a free data retrieval call binding the contract method 0x614cf6a1.
//
// Solidity: function RC_INVALID_LTV() view returns(string)
func (_Errors *ErrorsSession) RCINVALIDLTV() (string, error) {
	return _Errors.Contract.RCINVALIDLTV(&_Errors.CallOpts)
}

// RCINVALIDLTV is a free data retrieval call binding the contract method 0x614cf6a1.
//
// Solidity: function RC_INVALID_LTV() view returns(string)
func (_Errors *ErrorsCallerSession) RCINVALIDLTV() (string, error) {
	return _Errors.Contract.RCINVALIDLTV(&_Errors.CallOpts)
}

// RCINVALIDRESERVEFACTOR is a free data retrieval call binding the contract method 0x9be4f03a.
//
// Solidity: function RC_INVALID_RESERVE_FACTOR() view returns(string)
func (_Errors *ErrorsCaller) RCINVALIDRESERVEFACTOR(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RC_INVALID_RESERVE_FACTOR")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RCINVALIDRESERVEFACTOR is a free data retrieval call binding the contract method 0x9be4f03a.
//
// Solidity: function RC_INVALID_RESERVE_FACTOR() view returns(string)
func (_Errors *ErrorsSession) RCINVALIDRESERVEFACTOR() (string, error) {
	return _Errors.Contract.RCINVALIDRESERVEFACTOR(&_Errors.CallOpts)
}

// RCINVALIDRESERVEFACTOR is a free data retrieval call binding the contract method 0x9be4f03a.
//
// Solidity: function RC_INVALID_RESERVE_FACTOR() view returns(string)
func (_Errors *ErrorsCallerSession) RCINVALIDRESERVEFACTOR() (string, error) {
	return _Errors.Contract.RCINVALIDRESERVEFACTOR(&_Errors.CallOpts)
}

// RLLIQUIDITYINDEXOVERFLOW is a free data retrieval call binding the contract method 0x4fe4f1ab.
//
// Solidity: function RL_LIQUIDITY_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) RLLIQUIDITYINDEXOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_LIQUIDITY_INDEX_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLLIQUIDITYINDEXOVERFLOW is a free data retrieval call binding the contract method 0x4fe4f1ab.
//
// Solidity: function RL_LIQUIDITY_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) RLLIQUIDITYINDEXOVERFLOW() (string, error) {
	return _Errors.Contract.RLLIQUIDITYINDEXOVERFLOW(&_Errors.CallOpts)
}

// RLLIQUIDITYINDEXOVERFLOW is a free data retrieval call binding the contract method 0x4fe4f1ab.
//
// Solidity: function RL_LIQUIDITY_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) RLLIQUIDITYINDEXOVERFLOW() (string, error) {
	return _Errors.Contract.RLLIQUIDITYINDEXOVERFLOW(&_Errors.CallOpts)
}

// RLLIQUIDITYRATEOVERFLOW is a free data retrieval call binding the contract method 0xf11c6720.
//
// Solidity: function RL_LIQUIDITY_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) RLLIQUIDITYRATEOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_LIQUIDITY_RATE_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLLIQUIDITYRATEOVERFLOW is a free data retrieval call binding the contract method 0xf11c6720.
//
// Solidity: function RL_LIQUIDITY_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) RLLIQUIDITYRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLLIQUIDITYRATEOVERFLOW(&_Errors.CallOpts)
}

// RLLIQUIDITYRATEOVERFLOW is a free data retrieval call binding the contract method 0xf11c6720.
//
// Solidity: function RL_LIQUIDITY_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) RLLIQUIDITYRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLLIQUIDITYRATEOVERFLOW(&_Errors.CallOpts)
}

// RLRESERVEALREADYINITIALIZED is a free data retrieval call binding the contract method 0xfe75fd26.
//
// Solidity: function RL_RESERVE_ALREADY_INITIALIZED() view returns(string)
func (_Errors *ErrorsCaller) RLRESERVEALREADYINITIALIZED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_RESERVE_ALREADY_INITIALIZED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLRESERVEALREADYINITIALIZED is a free data retrieval call binding the contract method 0xfe75fd26.
//
// Solidity: function RL_RESERVE_ALREADY_INITIALIZED() view returns(string)
func (_Errors *ErrorsSession) RLRESERVEALREADYINITIALIZED() (string, error) {
	return _Errors.Contract.RLRESERVEALREADYINITIALIZED(&_Errors.CallOpts)
}

// RLRESERVEALREADYINITIALIZED is a free data retrieval call binding the contract method 0xfe75fd26.
//
// Solidity: function RL_RESERVE_ALREADY_INITIALIZED() view returns(string)
func (_Errors *ErrorsCallerSession) RLRESERVEALREADYINITIALIZED() (string, error) {
	return _Errors.Contract.RLRESERVEALREADYINITIALIZED(&_Errors.CallOpts)
}

// RLSTABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x6d422aa1.
//
// Solidity: function RL_STABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) RLSTABLEBORROWRATEOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_STABLE_BORROW_RATE_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLSTABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x6d422aa1.
//
// Solidity: function RL_STABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) RLSTABLEBORROWRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLSTABLEBORROWRATEOVERFLOW(&_Errors.CallOpts)
}

// RLSTABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x6d422aa1.
//
// Solidity: function RL_STABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) RLSTABLEBORROWRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLSTABLEBORROWRATEOVERFLOW(&_Errors.CallOpts)
}

// RLVARIABLEBORROWINDEXOVERFLOW is a free data retrieval call binding the contract method 0x44942004.
//
// Solidity: function RL_VARIABLE_BORROW_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) RLVARIABLEBORROWINDEXOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_VARIABLE_BORROW_INDEX_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLVARIABLEBORROWINDEXOVERFLOW is a free data retrieval call binding the contract method 0x44942004.
//
// Solidity: function RL_VARIABLE_BORROW_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) RLVARIABLEBORROWINDEXOVERFLOW() (string, error) {
	return _Errors.Contract.RLVARIABLEBORROWINDEXOVERFLOW(&_Errors.CallOpts)
}

// RLVARIABLEBORROWINDEXOVERFLOW is a free data retrieval call binding the contract method 0x44942004.
//
// Solidity: function RL_VARIABLE_BORROW_INDEX_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) RLVARIABLEBORROWINDEXOVERFLOW() (string, error) {
	return _Errors.Contract.RLVARIABLEBORROWINDEXOVERFLOW(&_Errors.CallOpts)
}

// RLVARIABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x2ea347b0.
//
// Solidity: function RL_VARIABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) RLVARIABLEBORROWRATEOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "RL_VARIABLE_BORROW_RATE_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// RLVARIABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x2ea347b0.
//
// Solidity: function RL_VARIABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) RLVARIABLEBORROWRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLVARIABLEBORROWRATEOVERFLOW(&_Errors.CallOpts)
}

// RLVARIABLEBORROWRATEOVERFLOW is a free data retrieval call binding the contract method 0x2ea347b0.
//
// Solidity: function RL_VARIABLE_BORROW_RATE_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) RLVARIABLEBORROWRATEOVERFLOW() (string, error) {
	return _Errors.Contract.RLVARIABLEBORROWRATEOVERFLOW(&_Errors.CallOpts)
}

// SDTBURNEXCEEDSBALANCE is a free data retrieval call binding the contract method 0x1befa78d.
//
// Solidity: function SDT_BURN_EXCEEDS_BALANCE() view returns(string)
func (_Errors *ErrorsCaller) SDTBURNEXCEEDSBALANCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SDT_BURN_EXCEEDS_BALANCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SDTBURNEXCEEDSBALANCE is a free data retrieval call binding the contract method 0x1befa78d.
//
// Solidity: function SDT_BURN_EXCEEDS_BALANCE() view returns(string)
func (_Errors *ErrorsSession) SDTBURNEXCEEDSBALANCE() (string, error) {
	return _Errors.Contract.SDTBURNEXCEEDSBALANCE(&_Errors.CallOpts)
}

// SDTBURNEXCEEDSBALANCE is a free data retrieval call binding the contract method 0x1befa78d.
//
// Solidity: function SDT_BURN_EXCEEDS_BALANCE() view returns(string)
func (_Errors *ErrorsCallerSession) SDTBURNEXCEEDSBALANCE() (string, error) {
	return _Errors.Contract.SDTBURNEXCEEDSBALANCE(&_Errors.CallOpts)
}

// SDTSTABLEDEBTOVERFLOW is a free data retrieval call binding the contract method 0xcdad445a.
//
// Solidity: function SDT_STABLE_DEBT_OVERFLOW() view returns(string)
func (_Errors *ErrorsCaller) SDTSTABLEDEBTOVERFLOW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "SDT_STABLE_DEBT_OVERFLOW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// SDTSTABLEDEBTOVERFLOW is a free data retrieval call binding the contract method 0xcdad445a.
//
// Solidity: function SDT_STABLE_DEBT_OVERFLOW() view returns(string)
func (_Errors *ErrorsSession) SDTSTABLEDEBTOVERFLOW() (string, error) {
	return _Errors.Contract.SDTSTABLEDEBTOVERFLOW(&_Errors.CallOpts)
}

// SDTSTABLEDEBTOVERFLOW is a free data retrieval call binding the contract method 0xcdad445a.
//
// Solidity: function SDT_STABLE_DEBT_OVERFLOW() view returns(string)
func (_Errors *ErrorsCallerSession) SDTSTABLEDEBTOVERFLOW() (string, error) {
	return _Errors.Contract.SDTSTABLEDEBTOVERFLOW(&_Errors.CallOpts)
}

// ULINVALIDINDEX is a free data retrieval call binding the contract method 0x02454ad3.
//
// Solidity: function UL_INVALID_INDEX() view returns(string)
func (_Errors *ErrorsCaller) ULINVALIDINDEX(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "UL_INVALID_INDEX")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ULINVALIDINDEX is a free data retrieval call binding the contract method 0x02454ad3.
//
// Solidity: function UL_INVALID_INDEX() view returns(string)
func (_Errors *ErrorsSession) ULINVALIDINDEX() (string, error) {
	return _Errors.Contract.ULINVALIDINDEX(&_Errors.CallOpts)
}

// ULINVALIDINDEX is a free data retrieval call binding the contract method 0x02454ad3.
//
// Solidity: function UL_INVALID_INDEX() view returns(string)
func (_Errors *ErrorsCallerSession) ULINVALIDINDEX() (string, error) {
	return _Errors.Contract.ULINVALIDINDEX(&_Errors.CallOpts)
}

// VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE is a free data retrieval call binding the contract method 0x3aa786a8.
//
// Solidity: function VL_AMOUNT_BIGGER_THAN_MAX_LOAN_SIZE_STABLE() view returns(string)
func (_Errors *ErrorsCaller) VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_AMOUNT_BIGGER_THAN_MAX_LOAN_SIZE_STABLE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE is a free data retrieval call binding the contract method 0x3aa786a8.
//
// Solidity: function VL_AMOUNT_BIGGER_THAN_MAX_LOAN_SIZE_STABLE() view returns(string)
func (_Errors *ErrorsSession) VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE() (string, error) {
	return _Errors.Contract.VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE(&_Errors.CallOpts)
}

// VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE is a free data retrieval call binding the contract method 0x3aa786a8.
//
// Solidity: function VL_AMOUNT_BIGGER_THAN_MAX_LOAN_SIZE_STABLE() view returns(string)
func (_Errors *ErrorsCallerSession) VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE() (string, error) {
	return _Errors.Contract.VLAMOUNTBIGGERTHANMAXLOANSIZESTABLE(&_Errors.CallOpts)
}

// VLBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x36565ab1.
//
// Solidity: function VL_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsCaller) VLBORROWINGNOTENABLED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_BORROWING_NOT_ENABLED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x36565ab1.
//
// Solidity: function VL_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsSession) VLBORROWINGNOTENABLED() (string, error) {
	return _Errors.Contract.VLBORROWINGNOTENABLED(&_Errors.CallOpts)
}

// VLBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x36565ab1.
//
// Solidity: function VL_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsCallerSession) VLBORROWINGNOTENABLED() (string, error) {
	return _Errors.Contract.VLBORROWINGNOTENABLED(&_Errors.CallOpts)
}

// VLCOLLATERALBALANCEIS0 is a free data retrieval call binding the contract method 0x708b8dd3.
//
// Solidity: function VL_COLLATERAL_BALANCE_IS_0() view returns(string)
func (_Errors *ErrorsCaller) VLCOLLATERALBALANCEIS0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_COLLATERAL_BALANCE_IS_0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLCOLLATERALBALANCEIS0 is a free data retrieval call binding the contract method 0x708b8dd3.
//
// Solidity: function VL_COLLATERAL_BALANCE_IS_0() view returns(string)
func (_Errors *ErrorsSession) VLCOLLATERALBALANCEIS0() (string, error) {
	return _Errors.Contract.VLCOLLATERALBALANCEIS0(&_Errors.CallOpts)
}

// VLCOLLATERALBALANCEIS0 is a free data retrieval call binding the contract method 0x708b8dd3.
//
// Solidity: function VL_COLLATERAL_BALANCE_IS_0() view returns(string)
func (_Errors *ErrorsCallerSession) VLCOLLATERALBALANCEIS0() (string, error) {
	return _Errors.Contract.VLCOLLATERALBALANCEIS0(&_Errors.CallOpts)
}

// VLCOLLATERALCANNOTCOVERNEWBORROW is a free data retrieval call binding the contract method 0x2ace698a.
//
// Solidity: function VL_COLLATERAL_CANNOT_COVER_NEW_BORROW() view returns(string)
func (_Errors *ErrorsCaller) VLCOLLATERALCANNOTCOVERNEWBORROW(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_COLLATERAL_CANNOT_COVER_NEW_BORROW")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLCOLLATERALCANNOTCOVERNEWBORROW is a free data retrieval call binding the contract method 0x2ace698a.
//
// Solidity: function VL_COLLATERAL_CANNOT_COVER_NEW_BORROW() view returns(string)
func (_Errors *ErrorsSession) VLCOLLATERALCANNOTCOVERNEWBORROW() (string, error) {
	return _Errors.Contract.VLCOLLATERALCANNOTCOVERNEWBORROW(&_Errors.CallOpts)
}

// VLCOLLATERALCANNOTCOVERNEWBORROW is a free data retrieval call binding the contract method 0x2ace698a.
//
// Solidity: function VL_COLLATERAL_CANNOT_COVER_NEW_BORROW() view returns(string)
func (_Errors *ErrorsCallerSession) VLCOLLATERALCANNOTCOVERNEWBORROW() (string, error) {
	return _Errors.Contract.VLCOLLATERALCANNOTCOVERNEWBORROW(&_Errors.CallOpts)
}

// VLCOLLATERALSAMEASBORROWINGCURRENCY is a free data retrieval call binding the contract method 0xa39ed4ff.
//
// Solidity: function VL_COLLATERAL_SAME_AS_BORROWING_CURRENCY() view returns(string)
func (_Errors *ErrorsCaller) VLCOLLATERALSAMEASBORROWINGCURRENCY(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_COLLATERAL_SAME_AS_BORROWING_CURRENCY")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLCOLLATERALSAMEASBORROWINGCURRENCY is a free data retrieval call binding the contract method 0xa39ed4ff.
//
// Solidity: function VL_COLLATERAL_SAME_AS_BORROWING_CURRENCY() view returns(string)
func (_Errors *ErrorsSession) VLCOLLATERALSAMEASBORROWINGCURRENCY() (string, error) {
	return _Errors.Contract.VLCOLLATERALSAMEASBORROWINGCURRENCY(&_Errors.CallOpts)
}

// VLCOLLATERALSAMEASBORROWINGCURRENCY is a free data retrieval call binding the contract method 0xa39ed4ff.
//
// Solidity: function VL_COLLATERAL_SAME_AS_BORROWING_CURRENCY() view returns(string)
func (_Errors *ErrorsCallerSession) VLCOLLATERALSAMEASBORROWINGCURRENCY() (string, error) {
	return _Errors.Contract.VLCOLLATERALSAMEASBORROWINGCURRENCY(&_Errors.CallOpts)
}

// VLCURRENTAVAILABLELIQUIDITYNOTENOUGH is a free data retrieval call binding the contract method 0x179476c5.
//
// Solidity: function VL_CURRENT_AVAILABLE_LIQUIDITY_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsCaller) VLCURRENTAVAILABLELIQUIDITYNOTENOUGH(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_CURRENT_AVAILABLE_LIQUIDITY_NOT_ENOUGH")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLCURRENTAVAILABLELIQUIDITYNOTENOUGH is a free data retrieval call binding the contract method 0x179476c5.
//
// Solidity: function VL_CURRENT_AVAILABLE_LIQUIDITY_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsSession) VLCURRENTAVAILABLELIQUIDITYNOTENOUGH() (string, error) {
	return _Errors.Contract.VLCURRENTAVAILABLELIQUIDITYNOTENOUGH(&_Errors.CallOpts)
}

// VLCURRENTAVAILABLELIQUIDITYNOTENOUGH is a free data retrieval call binding the contract method 0x179476c5.
//
// Solidity: function VL_CURRENT_AVAILABLE_LIQUIDITY_NOT_ENOUGH() view returns(string)
func (_Errors *ErrorsCallerSession) VLCURRENTAVAILABLELIQUIDITYNOTENOUGH() (string, error) {
	return _Errors.Contract.VLCURRENTAVAILABLELIQUIDITYNOTENOUGH(&_Errors.CallOpts)
}

// VLDEPOSITALREADYINUSE is a free data retrieval call binding the contract method 0xe29425dc.
//
// Solidity: function VL_DEPOSIT_ALREADY_IN_USE() view returns(string)
func (_Errors *ErrorsCaller) VLDEPOSITALREADYINUSE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_DEPOSIT_ALREADY_IN_USE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLDEPOSITALREADYINUSE is a free data retrieval call binding the contract method 0xe29425dc.
//
// Solidity: function VL_DEPOSIT_ALREADY_IN_USE() view returns(string)
func (_Errors *ErrorsSession) VLDEPOSITALREADYINUSE() (string, error) {
	return _Errors.Contract.VLDEPOSITALREADYINUSE(&_Errors.CallOpts)
}

// VLDEPOSITALREADYINUSE is a free data retrieval call binding the contract method 0xe29425dc.
//
// Solidity: function VL_DEPOSIT_ALREADY_IN_USE() view returns(string)
func (_Errors *ErrorsCallerSession) VLDEPOSITALREADYINUSE() (string, error) {
	return _Errors.Contract.VLDEPOSITALREADYINUSE(&_Errors.CallOpts)
}

// VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x1ec68b1d.
//
// Solidity: function VL_HEALTH_FACTOR_LOWER_THAN_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsCaller) VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_HEALTH_FACTOR_LOWER_THAN_LIQUIDATION_THRESHOLD")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x1ec68b1d.
//
// Solidity: function VL_HEALTH_FACTOR_LOWER_THAN_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsSession) VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD() (string, error) {
	return _Errors.Contract.VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD(&_Errors.CallOpts)
}

// VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD is a free data retrieval call binding the contract method 0x1ec68b1d.
//
// Solidity: function VL_HEALTH_FACTOR_LOWER_THAN_LIQUIDATION_THRESHOLD() view returns(string)
func (_Errors *ErrorsCallerSession) VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD() (string, error) {
	return _Errors.Contract.VLHEALTHFACTORLOWERTHANLIQUIDATIONTHRESHOLD(&_Errors.CallOpts)
}

// VLINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0xd7b079aa.
//
// Solidity: function VL_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsCaller) VLINCONSISTENTFLASHLOANPARAMS(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_INCONSISTENT_FLASHLOAN_PARAMS")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0xd7b079aa.
//
// Solidity: function VL_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsSession) VLINCONSISTENTFLASHLOANPARAMS() (string, error) {
	return _Errors.Contract.VLINCONSISTENTFLASHLOANPARAMS(&_Errors.CallOpts)
}

// VLINCONSISTENTFLASHLOANPARAMS is a free data retrieval call binding the contract method 0xd7b079aa.
//
// Solidity: function VL_INCONSISTENT_FLASHLOAN_PARAMS() view returns(string)
func (_Errors *ErrorsCallerSession) VLINCONSISTENTFLASHLOANPARAMS() (string, error) {
	return _Errors.Contract.VLINCONSISTENTFLASHLOANPARAMS(&_Errors.CallOpts)
}

// VLINVALIDAMOUNT is a free data retrieval call binding the contract method 0x871938a8.
//
// Solidity: function VL_INVALID_AMOUNT() view returns(string)
func (_Errors *ErrorsCaller) VLINVALIDAMOUNT(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_INVALID_AMOUNT")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLINVALIDAMOUNT is a free data retrieval call binding the contract method 0x871938a8.
//
// Solidity: function VL_INVALID_AMOUNT() view returns(string)
func (_Errors *ErrorsSession) VLINVALIDAMOUNT() (string, error) {
	return _Errors.Contract.VLINVALIDAMOUNT(&_Errors.CallOpts)
}

// VLINVALIDAMOUNT is a free data retrieval call binding the contract method 0x871938a8.
//
// Solidity: function VL_INVALID_AMOUNT() view returns(string)
func (_Errors *ErrorsCallerSession) VLINVALIDAMOUNT() (string, error) {
	return _Errors.Contract.VLINVALIDAMOUNT(&_Errors.CallOpts)
}

// VLINVALIDINTERESTRATEMODESELECTED is a free data retrieval call binding the contract method 0x3b5d25aa.
//
// Solidity: function VL_INVALID_INTEREST_RATE_MODE_SELECTED() view returns(string)
func (_Errors *ErrorsCaller) VLINVALIDINTERESTRATEMODESELECTED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_INVALID_INTEREST_RATE_MODE_SELECTED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLINVALIDINTERESTRATEMODESELECTED is a free data retrieval call binding the contract method 0x3b5d25aa.
//
// Solidity: function VL_INVALID_INTEREST_RATE_MODE_SELECTED() view returns(string)
func (_Errors *ErrorsSession) VLINVALIDINTERESTRATEMODESELECTED() (string, error) {
	return _Errors.Contract.VLINVALIDINTERESTRATEMODESELECTED(&_Errors.CallOpts)
}

// VLINVALIDINTERESTRATEMODESELECTED is a free data retrieval call binding the contract method 0x3b5d25aa.
//
// Solidity: function VL_INVALID_INTEREST_RATE_MODE_SELECTED() view returns(string)
func (_Errors *ErrorsCallerSession) VLINVALIDINTERESTRATEMODESELECTED() (string, error) {
	return _Errors.Contract.VLINVALIDINTERESTRATEMODESELECTED(&_Errors.CallOpts)
}

// VLNOTENOUGHAVAILABLEUSERBALANCE is a free data retrieval call binding the contract method 0xa8440241.
//
// Solidity: function VL_NOT_ENOUGH_AVAILABLE_USER_BALANCE() view returns(string)
func (_Errors *ErrorsCaller) VLNOTENOUGHAVAILABLEUSERBALANCE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NOT_ENOUGH_AVAILABLE_USER_BALANCE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNOTENOUGHAVAILABLEUSERBALANCE is a free data retrieval call binding the contract method 0xa8440241.
//
// Solidity: function VL_NOT_ENOUGH_AVAILABLE_USER_BALANCE() view returns(string)
func (_Errors *ErrorsSession) VLNOTENOUGHAVAILABLEUSERBALANCE() (string, error) {
	return _Errors.Contract.VLNOTENOUGHAVAILABLEUSERBALANCE(&_Errors.CallOpts)
}

// VLNOTENOUGHAVAILABLEUSERBALANCE is a free data retrieval call binding the contract method 0xa8440241.
//
// Solidity: function VL_NOT_ENOUGH_AVAILABLE_USER_BALANCE() view returns(string)
func (_Errors *ErrorsCallerSession) VLNOTENOUGHAVAILABLEUSERBALANCE() (string, error) {
	return _Errors.Contract.VLNOTENOUGHAVAILABLEUSERBALANCE(&_Errors.CallOpts)
}

// VLNOACTIVERESERVE is a free data retrieval call binding the contract method 0x7865a627.
//
// Solidity: function VL_NO_ACTIVE_RESERVE() view returns(string)
func (_Errors *ErrorsCaller) VLNOACTIVERESERVE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NO_ACTIVE_RESERVE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNOACTIVERESERVE is a free data retrieval call binding the contract method 0x7865a627.
//
// Solidity: function VL_NO_ACTIVE_RESERVE() view returns(string)
func (_Errors *ErrorsSession) VLNOACTIVERESERVE() (string, error) {
	return _Errors.Contract.VLNOACTIVERESERVE(&_Errors.CallOpts)
}

// VLNOACTIVERESERVE is a free data retrieval call binding the contract method 0x7865a627.
//
// Solidity: function VL_NO_ACTIVE_RESERVE() view returns(string)
func (_Errors *ErrorsCallerSession) VLNOACTIVERESERVE() (string, error) {
	return _Errors.Contract.VLNOACTIVERESERVE(&_Errors.CallOpts)
}

// VLNODEBTOFSELECTEDTYPE is a free data retrieval call binding the contract method 0x91a9fb18.
//
// Solidity: function VL_NO_DEBT_OF_SELECTED_TYPE() view returns(string)
func (_Errors *ErrorsCaller) VLNODEBTOFSELECTEDTYPE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NO_DEBT_OF_SELECTED_TYPE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNODEBTOFSELECTEDTYPE is a free data retrieval call binding the contract method 0x91a9fb18.
//
// Solidity: function VL_NO_DEBT_OF_SELECTED_TYPE() view returns(string)
func (_Errors *ErrorsSession) VLNODEBTOFSELECTEDTYPE() (string, error) {
	return _Errors.Contract.VLNODEBTOFSELECTEDTYPE(&_Errors.CallOpts)
}

// VLNODEBTOFSELECTEDTYPE is a free data retrieval call binding the contract method 0x91a9fb18.
//
// Solidity: function VL_NO_DEBT_OF_SELECTED_TYPE() view returns(string)
func (_Errors *ErrorsCallerSession) VLNODEBTOFSELECTEDTYPE() (string, error) {
	return _Errors.Contract.VLNODEBTOFSELECTEDTYPE(&_Errors.CallOpts)
}

// VLNOEXPLICITAMOUNTTOREPAYONBEHALF is a free data retrieval call binding the contract method 0xdaf23547.
//
// Solidity: function VL_NO_EXPLICIT_AMOUNT_TO_REPAY_ON_BEHALF() view returns(string)
func (_Errors *ErrorsCaller) VLNOEXPLICITAMOUNTTOREPAYONBEHALF(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NO_EXPLICIT_AMOUNT_TO_REPAY_ON_BEHALF")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNOEXPLICITAMOUNTTOREPAYONBEHALF is a free data retrieval call binding the contract method 0xdaf23547.
//
// Solidity: function VL_NO_EXPLICIT_AMOUNT_TO_REPAY_ON_BEHALF() view returns(string)
func (_Errors *ErrorsSession) VLNOEXPLICITAMOUNTTOREPAYONBEHALF() (string, error) {
	return _Errors.Contract.VLNOEXPLICITAMOUNTTOREPAYONBEHALF(&_Errors.CallOpts)
}

// VLNOEXPLICITAMOUNTTOREPAYONBEHALF is a free data retrieval call binding the contract method 0xdaf23547.
//
// Solidity: function VL_NO_EXPLICIT_AMOUNT_TO_REPAY_ON_BEHALF() view returns(string)
func (_Errors *ErrorsCallerSession) VLNOEXPLICITAMOUNTTOREPAYONBEHALF() (string, error) {
	return _Errors.Contract.VLNOEXPLICITAMOUNTTOREPAYONBEHALF(&_Errors.CallOpts)
}

// VLNOSTABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6422b257.
//
// Solidity: function VL_NO_STABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsCaller) VLNOSTABLERATELOANINRESERVE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NO_STABLE_RATE_LOAN_IN_RESERVE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNOSTABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6422b257.
//
// Solidity: function VL_NO_STABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsSession) VLNOSTABLERATELOANINRESERVE() (string, error) {
	return _Errors.Contract.VLNOSTABLERATELOANINRESERVE(&_Errors.CallOpts)
}

// VLNOSTABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6422b257.
//
// Solidity: function VL_NO_STABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsCallerSession) VLNOSTABLERATELOANINRESERVE() (string, error) {
	return _Errors.Contract.VLNOSTABLERATELOANINRESERVE(&_Errors.CallOpts)
}

// VLNOVARIABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6ab5e615.
//
// Solidity: function VL_NO_VARIABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsCaller) VLNOVARIABLERATELOANINRESERVE(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_NO_VARIABLE_RATE_LOAN_IN_RESERVE")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLNOVARIABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6ab5e615.
//
// Solidity: function VL_NO_VARIABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsSession) VLNOVARIABLERATELOANINRESERVE() (string, error) {
	return _Errors.Contract.VLNOVARIABLERATELOANINRESERVE(&_Errors.CallOpts)
}

// VLNOVARIABLERATELOANINRESERVE is a free data retrieval call binding the contract method 0x6ab5e615.
//
// Solidity: function VL_NO_VARIABLE_RATE_LOAN_IN_RESERVE() view returns(string)
func (_Errors *ErrorsCallerSession) VLNOVARIABLERATELOANINRESERVE() (string, error) {
	return _Errors.Contract.VLNOVARIABLERATELOANINRESERVE(&_Errors.CallOpts)
}

// VLRESERVEFROZEN is a free data retrieval call binding the contract method 0xd7510e0c.
//
// Solidity: function VL_RESERVE_FROZEN() view returns(string)
func (_Errors *ErrorsCaller) VLRESERVEFROZEN(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_RESERVE_FROZEN")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLRESERVEFROZEN is a free data retrieval call binding the contract method 0xd7510e0c.
//
// Solidity: function VL_RESERVE_FROZEN() view returns(string)
func (_Errors *ErrorsSession) VLRESERVEFROZEN() (string, error) {
	return _Errors.Contract.VLRESERVEFROZEN(&_Errors.CallOpts)
}

// VLRESERVEFROZEN is a free data retrieval call binding the contract method 0xd7510e0c.
//
// Solidity: function VL_RESERVE_FROZEN() view returns(string)
func (_Errors *ErrorsCallerSession) VLRESERVEFROZEN() (string, error) {
	return _Errors.Contract.VLRESERVEFROZEN(&_Errors.CallOpts)
}

// VLSTABLEBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x4927c63a.
//
// Solidity: function VL_STABLE_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsCaller) VLSTABLEBORROWINGNOTENABLED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_STABLE_BORROWING_NOT_ENABLED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLSTABLEBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x4927c63a.
//
// Solidity: function VL_STABLE_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsSession) VLSTABLEBORROWINGNOTENABLED() (string, error) {
	return _Errors.Contract.VLSTABLEBORROWINGNOTENABLED(&_Errors.CallOpts)
}

// VLSTABLEBORROWINGNOTENABLED is a free data retrieval call binding the contract method 0x4927c63a.
//
// Solidity: function VL_STABLE_BORROWING_NOT_ENABLED() view returns(string)
func (_Errors *ErrorsCallerSession) VLSTABLEBORROWINGNOTENABLED() (string, error) {
	return _Errors.Contract.VLSTABLEBORROWINGNOTENABLED(&_Errors.CallOpts)
}

// VLTRANSFERNOTALLOWED is a free data retrieval call binding the contract method 0xf3d9cc11.
//
// Solidity: function VL_TRANSFER_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCaller) VLTRANSFERNOTALLOWED(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_TRANSFER_NOT_ALLOWED")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLTRANSFERNOTALLOWED is a free data retrieval call binding the contract method 0xf3d9cc11.
//
// Solidity: function VL_TRANSFER_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsSession) VLTRANSFERNOTALLOWED() (string, error) {
	return _Errors.Contract.VLTRANSFERNOTALLOWED(&_Errors.CallOpts)
}

// VLTRANSFERNOTALLOWED is a free data retrieval call binding the contract method 0xf3d9cc11.
//
// Solidity: function VL_TRANSFER_NOT_ALLOWED() view returns(string)
func (_Errors *ErrorsCallerSession) VLTRANSFERNOTALLOWED() (string, error) {
	return _Errors.Contract.VLTRANSFERNOTALLOWED(&_Errors.CallOpts)
}

// VLUNDERLYINGBALANCENOTGREATERTHAN0 is a free data retrieval call binding the contract method 0x35a9d21d.
//
// Solidity: function VL_UNDERLYING_BALANCE_NOT_GREATER_THAN_0() view returns(string)
func (_Errors *ErrorsCaller) VLUNDERLYINGBALANCENOTGREATERTHAN0(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Errors.contract.Call(opts, &out, "VL_UNDERLYING_BALANCE_NOT_GREATER_THAN_0")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// VLUNDERLYINGBALANCENOTGREATERTHAN0 is a free data retrieval call binding the contract method 0x35a9d21d.
//
// Solidity: function VL_UNDERLYING_BALANCE_NOT_GREATER_THAN_0() view returns(string)
func (_Errors *ErrorsSession) VLUNDERLYINGBALANCENOTGREATERTHAN0() (string, error) {
	return _Errors.Contract.VLUNDERLYINGBALANCENOTGREATERTHAN0(&_Errors.CallOpts)
}

// VLUNDERLYINGBALANCENOTGREATERTHAN0 is a free data retrieval call binding the contract method 0x35a9d21d.
//
// Solidity: function VL_UNDERLYING_BALANCE_NOT_GREATER_THAN_0() view returns(string)
func (_Errors *ErrorsCallerSession) VLUNDERLYINGBALANCENOTGREATERTHAN0() (string, error) {
	return _Errors.Contract.VLUNDERLYINGBALANCENOTGREATERTHAN0(&_Errors.CallOpts)
}

// IERC20MetaData contains all meta data concerning the IERC20 contract.
var IERC20MetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20ABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20MetaData.ABI instead.
var IERC20ABI = IERC20MetaData.ABI

// IERC20 is an auto generated Go binding around an Ethereum contract.
type IERC20 struct {
	IERC20Caller     // Read-only binding to the contract
	IERC20Transactor // Write-only binding to the contract
	IERC20Filterer   // Log filterer for contract events
}

// IERC20Caller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20Session struct {
	Contract     *IERC20           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20CallerSession struct {
	Contract *IERC20Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// IERC20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20TransactorSession struct {
	Contract     *IERC20Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20Raw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20Raw struct {
	Contract *IERC20 // Generic contract binding to access the raw methods on
}

// IERC20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20CallerRaw struct {
	Contract *IERC20Caller // Generic read-only contract binding to access the raw methods on
}

// IERC20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20TransactorRaw struct {
	Contract *IERC20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20 creates a new instance of IERC20, bound to a specific deployed contract.
func NewIERC20(address common.Address, backend bind.ContractBackend) (*IERC20, error) {
	contract, err := bindIERC20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20{IERC20Caller: IERC20Caller{contract: contract}, IERC20Transactor: IERC20Transactor{contract: contract}, IERC20Filterer: IERC20Filterer{contract: contract}}, nil
}

// NewIERC20Caller creates a new read-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Caller(address common.Address, caller bind.ContractCaller) (*IERC20Caller, error) {
	contract, err := bindIERC20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Caller{contract: contract}, nil
}

// NewIERC20Transactor creates a new write-only instance of IERC20, bound to a specific deployed contract.
func NewIERC20Transactor(address common.Address, transactor bind.ContractTransactor) (*IERC20Transactor, error) {
	contract, err := bindIERC20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20Transactor{contract: contract}, nil
}

// NewIERC20Filterer creates a new log filterer instance of IERC20, bound to a specific deployed contract.
func NewIERC20Filterer(address common.Address, filterer bind.ContractFilterer) (*IERC20Filterer, error) {
	contract, err := bindIERC20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20Filterer{contract: contract}, nil
}

// bindIERC20 binds a generic wrapper to an already deployed contract.
func bindIERC20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.IERC20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.IERC20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20 *IERC20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20 *IERC20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20 *IERC20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20 *IERC20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20.Contract.Allowance(&_IERC20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Caller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20Session) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20 *IERC20CallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20.Contract.BalanceOf(&_IERC20.CallOpts, account)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20Session) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20 *IERC20CallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20.Contract.TotalSupply(&_IERC20.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Approve(&_IERC20.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.Transfer(&_IERC20.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Transactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20Session) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20 *IERC20TransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20.Contract.TransferFrom(&_IERC20.TransactOpts, sender, recipient, amount)
}

// IERC20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20 contract.
type IERC20ApprovalIterator struct {
	Event *IERC20Approval // Event containing the contract specifics and raw log

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
func (it *IERC20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Approval)
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
		it.Event = new(IERC20Approval)
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
func (it *IERC20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Approval represents a Approval event raised by the IERC20 contract.
type IERC20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20ApprovalIterator{contract: _IERC20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20 *IERC20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Approval)
				if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseApproval(log types.Log) (*IERC20Approval, error) {
	event := new(IERC20Approval)
	if err := _IERC20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20 contract.
type IERC20TransferIterator struct {
	Event *IERC20Transfer // Event containing the contract specifics and raw log

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
func (it *IERC20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20Transfer)
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
		it.Event = new(IERC20Transfer)
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
func (it *IERC20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20Transfer represents a Transfer event raised by the IERC20 contract.
type IERC20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20TransferIterator{contract: _IERC20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20 *IERC20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20Transfer)
				if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20 *IERC20Filterer) ParseTransfer(log types.Log) (*IERC20Transfer, error) {
	event := new(IERC20Transfer)
	if err := _IERC20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20DetailedMetaData contains all meta data concerning the IERC20Detailed contract.
var IERC20DetailedMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"account\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"sender\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// IERC20DetailedABI is the input ABI used to generate the binding from.
// Deprecated: Use IERC20DetailedMetaData.ABI instead.
var IERC20DetailedABI = IERC20DetailedMetaData.ABI

// IERC20Detailed is an auto generated Go binding around an Ethereum contract.
type IERC20Detailed struct {
	IERC20DetailedCaller     // Read-only binding to the contract
	IERC20DetailedTransactor // Write-only binding to the contract
	IERC20DetailedFilterer   // Log filterer for contract events
}

// IERC20DetailedCaller is an auto generated read-only Go binding around an Ethereum contract.
type IERC20DetailedCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DetailedTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IERC20DetailedTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DetailedFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IERC20DetailedFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IERC20DetailedSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IERC20DetailedSession struct {
	Contract     *IERC20Detailed   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IERC20DetailedCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IERC20DetailedCallerSession struct {
	Contract *IERC20DetailedCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// IERC20DetailedTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IERC20DetailedTransactorSession struct {
	Contract     *IERC20DetailedTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// IERC20DetailedRaw is an auto generated low-level Go binding around an Ethereum contract.
type IERC20DetailedRaw struct {
	Contract *IERC20Detailed // Generic contract binding to access the raw methods on
}

// IERC20DetailedCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IERC20DetailedCallerRaw struct {
	Contract *IERC20DetailedCaller // Generic read-only contract binding to access the raw methods on
}

// IERC20DetailedTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IERC20DetailedTransactorRaw struct {
	Contract *IERC20DetailedTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIERC20Detailed creates a new instance of IERC20Detailed, bound to a specific deployed contract.
func NewIERC20Detailed(address common.Address, backend bind.ContractBackend) (*IERC20Detailed, error) {
	contract, err := bindIERC20Detailed(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IERC20Detailed{IERC20DetailedCaller: IERC20DetailedCaller{contract: contract}, IERC20DetailedTransactor: IERC20DetailedTransactor{contract: contract}, IERC20DetailedFilterer: IERC20DetailedFilterer{contract: contract}}, nil
}

// NewIERC20DetailedCaller creates a new read-only instance of IERC20Detailed, bound to a specific deployed contract.
func NewIERC20DetailedCaller(address common.Address, caller bind.ContractCaller) (*IERC20DetailedCaller, error) {
	contract, err := bindIERC20Detailed(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20DetailedCaller{contract: contract}, nil
}

// NewIERC20DetailedTransactor creates a new write-only instance of IERC20Detailed, bound to a specific deployed contract.
func NewIERC20DetailedTransactor(address common.Address, transactor bind.ContractTransactor) (*IERC20DetailedTransactor, error) {
	contract, err := bindIERC20Detailed(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IERC20DetailedTransactor{contract: contract}, nil
}

// NewIERC20DetailedFilterer creates a new log filterer instance of IERC20Detailed, bound to a specific deployed contract.
func NewIERC20DetailedFilterer(address common.Address, filterer bind.ContractFilterer) (*IERC20DetailedFilterer, error) {
	contract, err := bindIERC20Detailed(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IERC20DetailedFilterer{contract: contract}, nil
}

// bindIERC20Detailed binds a generic wrapper to an already deployed contract.
func bindIERC20Detailed(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IERC20DetailedABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Detailed *IERC20DetailedRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Detailed.Contract.IERC20DetailedCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Detailed *IERC20DetailedRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.IERC20DetailedTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Detailed *IERC20DetailedRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.IERC20DetailedTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IERC20Detailed *IERC20DetailedCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IERC20Detailed.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IERC20Detailed *IERC20DetailedTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IERC20Detailed *IERC20DetailedTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Detailed.Contract.Allowance(&_IERC20Detailed.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _IERC20Detailed.Contract.Allowance(&_IERC20Detailed.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Detailed.Contract.BalanceOf(&_IERC20Detailed.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _IERC20Detailed.Contract.BalanceOf(&_IERC20Detailed.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Detailed *IERC20DetailedCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Detailed *IERC20DetailedSession) Decimals() (uint8, error) {
	return _IERC20Detailed.Contract.Decimals(&_IERC20Detailed.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_IERC20Detailed *IERC20DetailedCallerSession) Decimals() (uint8, error) {
	return _IERC20Detailed.Contract.Decimals(&_IERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Detailed *IERC20DetailedCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Detailed *IERC20DetailedSession) Name() (string, error) {
	return _IERC20Detailed.Contract.Name(&_IERC20Detailed.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_IERC20Detailed *IERC20DetailedCallerSession) Name() (string, error) {
	return _IERC20Detailed.Contract.Name(&_IERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Detailed *IERC20DetailedCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Detailed *IERC20DetailedSession) Symbol() (string, error) {
	return _IERC20Detailed.Contract.Symbol(&_IERC20Detailed.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_IERC20Detailed *IERC20DetailedCallerSession) Symbol() (string, error) {
	return _IERC20Detailed.Contract.Symbol(&_IERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IERC20Detailed.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Detailed *IERC20DetailedSession) TotalSupply() (*big.Int, error) {
	return _IERC20Detailed.Contract.TotalSupply(&_IERC20Detailed.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_IERC20Detailed *IERC20DetailedCallerSession) TotalSupply() (*big.Int, error) {
	return _IERC20Detailed.Contract.TotalSupply(&_IERC20Detailed.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.Approve(&_IERC20Detailed.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.Approve(&_IERC20Detailed.TransactOpts, spender, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactor) Transfer(opts *bind.TransactOpts, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.contract.Transact(opts, "transfer", recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.Transfer(&_IERC20Detailed.TransactOpts, recipient, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactorSession) Transfer(recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.Transfer(&_IERC20Detailed.TransactOpts, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactor) TransferFrom(opts *bind.TransactOpts, sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.contract.Transact(opts, "transferFrom", sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.TransferFrom(&_IERC20Detailed.TransactOpts, sender, recipient, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address sender, address recipient, uint256 amount) returns(bool)
func (_IERC20Detailed *IERC20DetailedTransactorSession) TransferFrom(sender common.Address, recipient common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IERC20Detailed.Contract.TransferFrom(&_IERC20Detailed.TransactOpts, sender, recipient, amount)
}

// IERC20DetailedApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the IERC20Detailed contract.
type IERC20DetailedApprovalIterator struct {
	Event *IERC20DetailedApproval // Event containing the contract specifics and raw log

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
func (it *IERC20DetailedApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20DetailedApproval)
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
		it.Event = new(IERC20DetailedApproval)
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
func (it *IERC20DetailedApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20DetailedApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20DetailedApproval represents a Approval event raised by the IERC20Detailed contract.
type IERC20DetailedApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Detailed *IERC20DetailedFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*IERC20DetailedApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Detailed.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &IERC20DetailedApprovalIterator{contract: _IERC20Detailed.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_IERC20Detailed *IERC20DetailedFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *IERC20DetailedApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _IERC20Detailed.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20DetailedApproval)
				if err := _IERC20Detailed.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_IERC20Detailed *IERC20DetailedFilterer) ParseApproval(log types.Log) (*IERC20DetailedApproval, error) {
	event := new(IERC20DetailedApproval)
	if err := _IERC20Detailed.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IERC20DetailedTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the IERC20Detailed contract.
type IERC20DetailedTransferIterator struct {
	Event *IERC20DetailedTransfer // Event containing the contract specifics and raw log

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
func (it *IERC20DetailedTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IERC20DetailedTransfer)
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
		it.Event = new(IERC20DetailedTransfer)
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
func (it *IERC20DetailedTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IERC20DetailedTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IERC20DetailedTransfer represents a Transfer event raised by the IERC20Detailed contract.
type IERC20DetailedTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Detailed *IERC20DetailedFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*IERC20DetailedTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Detailed.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &IERC20DetailedTransferIterator{contract: _IERC20Detailed.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_IERC20Detailed *IERC20DetailedFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *IERC20DetailedTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _IERC20Detailed.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IERC20DetailedTransfer)
				if err := _IERC20Detailed.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_IERC20Detailed *IERC20DetailedFilterer) ParseTransfer(log types.Log) (*IERC20DetailedTransfer, error) {
	event := new(IERC20DetailedTransfer)
	if err := _IERC20Detailed.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolMetaData contains all meta data concerning the ILendingPool contract.
var ILendingPoolMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRateMode\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"borrowRate\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Borrow\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"uint16\",\"name\":\"referral\",\"type\":\"uint16\"}],\"name\":\"Deposit\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"target\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"premium\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"FlashLoan\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidatedCollateralAmount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"liquidator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"LiquidationCall\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Paused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"RebalanceStableBorrowRate\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"repayer\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Repay\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"stableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"liquidityIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"variableBorrowIndex\",\"type\":\"uint256\"}],\"name\":\"ReserveDataUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralDisabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"ReserveUsedAsCollateralEnabled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"}],\"name\":\"Swap\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpaused\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Withdraw\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"interestRateMode\",\"type\":\"uint256\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"borrow\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"deposit\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceFromAfter\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"balanceToBefore\",\"type\":\"uint256\"}],\"name\":\"finalizeTransfer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"receiverAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"assets\",\"type\":\"address[]\"},{\"internalType\":\"uint256[]\",\"name\":\"amounts\",\"type\":\"uint256[]\"},{\"internalType\":\"uint256[]\",\"name\":\"modes\",\"type\":\"uint256[]\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"params\",\"type\":\"bytes\"},{\"internalType\":\"uint16\",\"name\":\"referralCode\",\"type\":\"uint16\"}],\"name\":\"flashLoan\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAddressesProvider\",\"outputs\":[{\"internalType\":\"contractILendingPoolAddressesProvider\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveData\",\"outputs\":[{\"components\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.ReserveConfigurationMap\",\"name\":\"configuration\",\"type\":\"tuple\"},{\"internalType\":\"uint128\",\"name\":\"liquidityIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"variableBorrowIndex\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentLiquidityRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentVariableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint128\",\"name\":\"currentStableBorrowRate\",\"type\":\"uint128\"},{\"internalType\":\"uint40\",\"name\":\"lastUpdateTimestamp\",\"type\":\"uint40\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"id\",\"type\":\"uint8\"}],\"internalType\":\"structDataTypes.ReserveData\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedIncome\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"}],\"name\":\"getReserveNormalizedVariableDebt\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getReservesList\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserAccountData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"totalCollateralETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"totalDebtETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"availableBorrowsETH\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"currentLiquidationThreshold\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"ltv\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"healthFactor\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserConfiguration\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"data\",\"type\":\"uint256\"}],\"internalType\":\"structDataTypes.UserConfigurationMap\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"aTokenAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"stableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"variableDebtAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestRateStrategyAddress\",\"type\":\"address\"}],\"name\":\"initReserve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"collateralAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"debtAsset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"debtToCover\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"receiveAToken\",\"type\":\"bool\"}],\"name\":\"liquidationCall\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"rebalanceStableBorrowRate\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"}],\"name\":\"repay\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"configuration\",\"type\":\"uint256\"}],\"name\":\"setConfiguration\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bool\",\"name\":\"val\",\"type\":\"bool\"}],\"name\":\"setPause\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"reserve\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"rateStrategyAddress\",\"type\":\"address\"}],\"name\":\"setReserveInterestRateStrategyAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"useAsCollateral\",\"type\":\"bool\"}],\"name\":\"setUserUseReserveAsCollateral\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"rateMode\",\"type\":\"uint256\"}],\"name\":\"swapBorrowRateMode\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"asset\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"withdraw\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILendingPoolABI is the input ABI used to generate the binding from.
// Deprecated: Use ILendingPoolMetaData.ABI instead.
var ILendingPoolABI = ILendingPoolMetaData.ABI

// ILendingPool is an auto generated Go binding around an Ethereum contract.
type ILendingPool struct {
	ILendingPoolCaller     // Read-only binding to the contract
	ILendingPoolTransactor // Write-only binding to the contract
	ILendingPoolFilterer   // Log filterer for contract events
}

// ILendingPoolCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILendingPoolCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILendingPoolTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILendingPoolFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILendingPoolSession struct {
	Contract     *ILendingPool     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ILendingPoolCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILendingPoolCallerSession struct {
	Contract *ILendingPoolCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ILendingPoolTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILendingPoolTransactorSession struct {
	Contract     *ILendingPoolTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ILendingPoolRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILendingPoolRaw struct {
	Contract *ILendingPool // Generic contract binding to access the raw methods on
}

// ILendingPoolCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILendingPoolCallerRaw struct {
	Contract *ILendingPoolCaller // Generic read-only contract binding to access the raw methods on
}

// ILendingPoolTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILendingPoolTransactorRaw struct {
	Contract *ILendingPoolTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILendingPool creates a new instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPool(address common.Address, backend bind.ContractBackend) (*ILendingPool, error) {
	contract, err := bindILendingPool(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILendingPool{ILendingPoolCaller: ILendingPoolCaller{contract: contract}, ILendingPoolTransactor: ILendingPoolTransactor{contract: contract}, ILendingPoolFilterer: ILendingPoolFilterer{contract: contract}}, nil
}

// NewILendingPoolCaller creates a new read-only instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolCaller(address common.Address, caller bind.ContractCaller) (*ILendingPoolCaller, error) {
	contract, err := bindILendingPool(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolCaller{contract: contract}, nil
}

// NewILendingPoolTransactor creates a new write-only instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolTransactor(address common.Address, transactor bind.ContractTransactor) (*ILendingPoolTransactor, error) {
	contract, err := bindILendingPool(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolTransactor{contract: contract}, nil
}

// NewILendingPoolFilterer creates a new log filterer instance of ILendingPool, bound to a specific deployed contract.
func NewILendingPoolFilterer(address common.Address, filterer bind.ContractFilterer) (*ILendingPoolFilterer, error) {
	contract, err := bindILendingPool(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolFilterer{contract: contract}, nil
}

// bindILendingPool binds a generic wrapper to an already deployed contract.
func bindILendingPool(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILendingPoolABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPool *ILendingPoolRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPool.Contract.ILendingPoolCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPool *ILendingPoolRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPool.Contract.ILendingPoolTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPool *ILendingPoolRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPool.Contract.ILendingPoolTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPool *ILendingPoolCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPool.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPool *ILendingPoolTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPool.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPool *ILendingPoolTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPool.Contract.contract.Transact(opts, method, params...)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolCaller) GetAddressesProvider(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getAddressesProvider")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolSession) GetAddressesProvider() (common.Address, error) {
	return _ILendingPool.Contract.GetAddressesProvider(&_ILendingPool.CallOpts)
}

// GetAddressesProvider is a free data retrieval call binding the contract method 0xfe65acfe.
//
// Solidity: function getAddressesProvider() view returns(address)
func (_ILendingPool *ILendingPoolCallerSession) GetAddressesProvider() (common.Address, error) {
	return _ILendingPool.Contract.GetAddressesProvider(&_ILendingPool.CallOpts)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_ILendingPool *ILendingPoolCaller) GetConfiguration(opts *bind.CallOpts, asset common.Address) (DataTypesReserveConfigurationMap, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getConfiguration", asset)

	if err != nil {
		return *new(DataTypesReserveConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveConfigurationMap)).(*DataTypesReserveConfigurationMap)

	return out0, err

}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_ILendingPool *ILendingPoolSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _ILendingPool.Contract.GetConfiguration(&_ILendingPool.CallOpts, asset)
}

// GetConfiguration is a free data retrieval call binding the contract method 0xc44b11f7.
//
// Solidity: function getConfiguration(address asset) view returns((uint256))
func (_ILendingPool *ILendingPoolCallerSession) GetConfiguration(asset common.Address) (DataTypesReserveConfigurationMap, error) {
	return _ILendingPool.Contract.GetConfiguration(&_ILendingPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_ILendingPool *ILendingPoolCaller) GetReserveData(opts *bind.CallOpts, asset common.Address) (DataTypesReserveData, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserveData", asset)

	if err != nil {
		return *new(DataTypesReserveData), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesReserveData)).(*DataTypesReserveData)

	return out0, err

}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_ILendingPool *ILendingPoolSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _ILendingPool.Contract.GetReserveData(&_ILendingPool.CallOpts, asset)
}

// GetReserveData is a free data retrieval call binding the contract method 0x35ea6a75.
//
// Solidity: function getReserveData(address asset) view returns(((uint256),uint128,uint128,uint128,uint128,uint128,uint40,address,address,address,address,uint8))
func (_ILendingPool *ILendingPoolCallerSession) GetReserveData(asset common.Address) (DataTypesReserveData, error) {
	return _ILendingPool.Contract.GetReserveData(&_ILendingPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolCaller) GetReserveNormalizedIncome(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserveNormalizedIncome", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _ILendingPool.Contract.GetReserveNormalizedIncome(&_ILendingPool.CallOpts, asset)
}

// GetReserveNormalizedIncome is a free data retrieval call binding the contract method 0xd15e0053.
//
// Solidity: function getReserveNormalizedIncome(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolCallerSession) GetReserveNormalizedIncome(asset common.Address) (*big.Int, error) {
	return _ILendingPool.Contract.GetReserveNormalizedIncome(&_ILendingPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolCaller) GetReserveNormalizedVariableDebt(opts *bind.CallOpts, asset common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReserveNormalizedVariableDebt", asset)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _ILendingPool.Contract.GetReserveNormalizedVariableDebt(&_ILendingPool.CallOpts, asset)
}

// GetReserveNormalizedVariableDebt is a free data retrieval call binding the contract method 0x386497fd.
//
// Solidity: function getReserveNormalizedVariableDebt(address asset) view returns(uint256)
func (_ILendingPool *ILendingPoolCallerSession) GetReserveNormalizedVariableDebt(asset common.Address) (*big.Int, error) {
	return _ILendingPool.Contract.GetReserveNormalizedVariableDebt(&_ILendingPool.CallOpts, asset)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_ILendingPool *ILendingPoolCaller) GetReservesList(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getReservesList")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_ILendingPool *ILendingPoolSession) GetReservesList() ([]common.Address, error) {
	return _ILendingPool.Contract.GetReservesList(&_ILendingPool.CallOpts)
}

// GetReservesList is a free data retrieval call binding the contract method 0xd1946dbc.
//
// Solidity: function getReservesList() view returns(address[])
func (_ILendingPool *ILendingPoolCallerSession) GetReservesList() ([]common.Address, error) {
	return _ILendingPool.Contract.GetReservesList(&_ILendingPool.CallOpts)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolCaller) GetUserAccountData(opts *bind.CallOpts, user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getUserAccountData", user)

	outstruct := new(struct {
		TotalCollateralETH          *big.Int
		TotalDebtETH                *big.Int
		AvailableBorrowsETH         *big.Int
		CurrentLiquidationThreshold *big.Int
		Ltv                         *big.Int
		HealthFactor                *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.TotalCollateralETH = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.TotalDebtETH = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.AvailableBorrowsETH = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.CurrentLiquidationThreshold = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.Ltv = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.HealthFactor = *abi.ConvertType(out[5], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _ILendingPool.Contract.GetUserAccountData(&_ILendingPool.CallOpts, user)
}

// GetUserAccountData is a free data retrieval call binding the contract method 0xbf92857c.
//
// Solidity: function getUserAccountData(address user) view returns(uint256 totalCollateralETH, uint256 totalDebtETH, uint256 availableBorrowsETH, uint256 currentLiquidationThreshold, uint256 ltv, uint256 healthFactor)
func (_ILendingPool *ILendingPoolCallerSession) GetUserAccountData(user common.Address) (struct {
	TotalCollateralETH          *big.Int
	TotalDebtETH                *big.Int
	AvailableBorrowsETH         *big.Int
	CurrentLiquidationThreshold *big.Int
	Ltv                         *big.Int
	HealthFactor                *big.Int
}, error) {
	return _ILendingPool.Contract.GetUserAccountData(&_ILendingPool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_ILendingPool *ILendingPoolCaller) GetUserConfiguration(opts *bind.CallOpts, user common.Address) (DataTypesUserConfigurationMap, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "getUserConfiguration", user)

	if err != nil {
		return *new(DataTypesUserConfigurationMap), err
	}

	out0 := *abi.ConvertType(out[0], new(DataTypesUserConfigurationMap)).(*DataTypesUserConfigurationMap)

	return out0, err

}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_ILendingPool *ILendingPoolSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _ILendingPool.Contract.GetUserConfiguration(&_ILendingPool.CallOpts, user)
}

// GetUserConfiguration is a free data retrieval call binding the contract method 0x4417a583.
//
// Solidity: function getUserConfiguration(address user) view returns((uint256))
func (_ILendingPool *ILendingPoolCallerSession) GetUserConfiguration(user common.Address) (DataTypesUserConfigurationMap, error) {
	return _ILendingPool.Contract.GetUserConfiguration(&_ILendingPool.CallOpts, user)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ILendingPool *ILendingPoolCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ILendingPool.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ILendingPool *ILendingPoolSession) Paused() (bool, error) {
	return _ILendingPool.Contract.Paused(&_ILendingPool.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ILendingPool *ILendingPoolCallerSession) Paused() (bool, error) {
	return _ILendingPool.Contract.Paused(&_ILendingPool.CallOpts)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_ILendingPool *ILendingPoolTransactor) Borrow(opts *bind.TransactOpts, asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "borrow", asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_ILendingPool *ILendingPoolSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Borrow(&_ILendingPool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Borrow is a paid mutator transaction binding the contract method 0xa415bcad.
//
// Solidity: function borrow(address asset, uint256 amount, uint256 interestRateMode, uint16 referralCode, address onBehalfOf) returns()
func (_ILendingPool *ILendingPoolTransactorSession) Borrow(asset common.Address, amount *big.Int, interestRateMode *big.Int, referralCode uint16, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Borrow(&_ILendingPool.TransactOpts, asset, amount, interestRateMode, referralCode, onBehalfOf)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolTransactor) Deposit(opts *bind.TransactOpts, asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "deposit", asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Deposit(&_ILendingPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// Deposit is a paid mutator transaction binding the contract method 0xe8eda9df.
//
// Solidity: function deposit(address asset, uint256 amount, address onBehalfOf, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolTransactorSession) Deposit(asset common.Address, amount *big.Int, onBehalfOf common.Address, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.Deposit(&_ILendingPool.TransactOpts, asset, amount, onBehalfOf, referralCode)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromAfter, uint256 balanceToBefore) returns()
func (_ILendingPool *ILendingPoolTransactor) FinalizeTransfer(opts *bind.TransactOpts, asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromAfter *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "finalizeTransfer", asset, from, to, amount, balanceFromAfter, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromAfter, uint256 balanceToBefore) returns()
func (_ILendingPool *ILendingPoolSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromAfter *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.FinalizeTransfer(&_ILendingPool.TransactOpts, asset, from, to, amount, balanceFromAfter, balanceToBefore)
}

// FinalizeTransfer is a paid mutator transaction binding the contract method 0xd5ed3933.
//
// Solidity: function finalizeTransfer(address asset, address from, address to, uint256 amount, uint256 balanceFromAfter, uint256 balanceToBefore) returns()
func (_ILendingPool *ILendingPoolTransactorSession) FinalizeTransfer(asset common.Address, from common.Address, to common.Address, amount *big.Int, balanceFromAfter *big.Int, balanceToBefore *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.FinalizeTransfer(&_ILendingPool.TransactOpts, asset, from, to, amount, balanceFromAfter, balanceToBefore)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolTransactor) FlashLoan(opts *bind.TransactOpts, receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "flashLoan", receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.FlashLoan(&_ILendingPool.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// FlashLoan is a paid mutator transaction binding the contract method 0xab9c4b5d.
//
// Solidity: function flashLoan(address receiverAddress, address[] assets, uint256[] amounts, uint256[] modes, address onBehalfOf, bytes params, uint16 referralCode) returns()
func (_ILendingPool *ILendingPoolTransactorSession) FlashLoan(receiverAddress common.Address, assets []common.Address, amounts []*big.Int, modes []*big.Int, onBehalfOf common.Address, params []byte, referralCode uint16) (*types.Transaction, error) {
	return _ILendingPool.Contract.FlashLoan(&_ILendingPool.TransactOpts, receiverAddress, assets, amounts, modes, onBehalfOf, params, referralCode)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address reserve, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolTransactor) InitReserve(opts *bind.TransactOpts, reserve common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "initReserve", reserve, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address reserve, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolSession) InitReserve(reserve common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.InitReserve(&_ILendingPool.TransactOpts, reserve, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// InitReserve is a paid mutator transaction binding the contract method 0x7a708e92.
//
// Solidity: function initReserve(address reserve, address aTokenAddress, address stableDebtAddress, address variableDebtAddress, address interestRateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolTransactorSession) InitReserve(reserve common.Address, aTokenAddress common.Address, stableDebtAddress common.Address, variableDebtAddress common.Address, interestRateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.InitReserve(&_ILendingPool.TransactOpts, reserve, aTokenAddress, stableDebtAddress, variableDebtAddress, interestRateStrategyAddress)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_ILendingPool *ILendingPoolTransactor) LiquidationCall(opts *bind.TransactOpts, collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "liquidationCall", collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_ILendingPool *ILendingPoolSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.LiquidationCall(&_ILendingPool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// LiquidationCall is a paid mutator transaction binding the contract method 0x00a718a9.
//
// Solidity: function liquidationCall(address collateralAsset, address debtAsset, address user, uint256 debtToCover, bool receiveAToken) returns()
func (_ILendingPool *ILendingPoolTransactorSession) LiquidationCall(collateralAsset common.Address, debtAsset common.Address, user common.Address, debtToCover *big.Int, receiveAToken bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.LiquidationCall(&_ILendingPool.TransactOpts, collateralAsset, debtAsset, user, debtToCover, receiveAToken)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_ILendingPool *ILendingPoolTransactor) RebalanceStableBorrowRate(opts *bind.TransactOpts, asset common.Address, user common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "rebalanceStableBorrowRate", asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_ILendingPool *ILendingPoolSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.RebalanceStableBorrowRate(&_ILendingPool.TransactOpts, asset, user)
}

// RebalanceStableBorrowRate is a paid mutator transaction binding the contract method 0xcd112382.
//
// Solidity: function rebalanceStableBorrowRate(address asset, address user) returns()
func (_ILendingPool *ILendingPoolTransactorSession) RebalanceStableBorrowRate(asset common.Address, user common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.RebalanceStableBorrowRate(&_ILendingPool.TransactOpts, asset, user)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_ILendingPool *ILendingPoolTransactor) Repay(opts *bind.TransactOpts, asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "repay", asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_ILendingPool *ILendingPoolSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Repay(&_ILendingPool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}

// Repay is a paid mutator transaction binding the contract method 0x573ade81.
//
// Solidity: function repay(address asset, uint256 amount, uint256 rateMode, address onBehalfOf) returns(uint256)
func (_ILendingPool *ILendingPoolTransactorSession) Repay(asset common.Address, amount *big.Int, rateMode *big.Int, onBehalfOf common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Repay(&_ILendingPool.TransactOpts, asset, amount, rateMode, onBehalfOf)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address reserve, uint256 configuration) returns()
func (_ILendingPool *ILendingPoolTransactor) SetConfiguration(opts *bind.TransactOpts, reserve common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "setConfiguration", reserve, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address reserve, uint256 configuration) returns()
func (_ILendingPool *ILendingPoolSession) SetConfiguration(reserve common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetConfiguration(&_ILendingPool.TransactOpts, reserve, configuration)
}

// SetConfiguration is a paid mutator transaction binding the contract method 0xb8d29276.
//
// Solidity: function setConfiguration(address reserve, uint256 configuration) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SetConfiguration(reserve common.Address, configuration *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetConfiguration(&_ILendingPool.TransactOpts, reserve, configuration)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_ILendingPool *ILendingPoolTransactor) SetPause(opts *bind.TransactOpts, val bool) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "setPause", val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_ILendingPool *ILendingPoolSession) SetPause(val bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetPause(&_ILendingPool.TransactOpts, val)
}

// SetPause is a paid mutator transaction binding the contract method 0xbedb86fb.
//
// Solidity: function setPause(bool val) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SetPause(val bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetPause(&_ILendingPool.TransactOpts, val)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address reserve, address rateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolTransactor) SetReserveInterestRateStrategyAddress(opts *bind.TransactOpts, reserve common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "setReserveInterestRateStrategyAddress", reserve, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address reserve, address rateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolSession) SetReserveInterestRateStrategyAddress(reserve common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetReserveInterestRateStrategyAddress(&_ILendingPool.TransactOpts, reserve, rateStrategyAddress)
}

// SetReserveInterestRateStrategyAddress is a paid mutator transaction binding the contract method 0x1d2118f9.
//
// Solidity: function setReserveInterestRateStrategyAddress(address reserve, address rateStrategyAddress) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SetReserveInterestRateStrategyAddress(reserve common.Address, rateStrategyAddress common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetReserveInterestRateStrategyAddress(&_ILendingPool.TransactOpts, reserve, rateStrategyAddress)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_ILendingPool *ILendingPoolTransactor) SetUserUseReserveAsCollateral(opts *bind.TransactOpts, asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "setUserUseReserveAsCollateral", asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_ILendingPool *ILendingPoolSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetUserUseReserveAsCollateral(&_ILendingPool.TransactOpts, asset, useAsCollateral)
}

// SetUserUseReserveAsCollateral is a paid mutator transaction binding the contract method 0x5a3b74b9.
//
// Solidity: function setUserUseReserveAsCollateral(address asset, bool useAsCollateral) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SetUserUseReserveAsCollateral(asset common.Address, useAsCollateral bool) (*types.Transaction, error) {
	return _ILendingPool.Contract.SetUserUseReserveAsCollateral(&_ILendingPool.TransactOpts, asset, useAsCollateral)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_ILendingPool *ILendingPoolTransactor) SwapBorrowRateMode(opts *bind.TransactOpts, asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "swapBorrowRateMode", asset, rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_ILendingPool *ILendingPoolSession) SwapBorrowRateMode(asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.SwapBorrowRateMode(&_ILendingPool.TransactOpts, asset, rateMode)
}

// SwapBorrowRateMode is a paid mutator transaction binding the contract method 0x94ba89a2.
//
// Solidity: function swapBorrowRateMode(address asset, uint256 rateMode) returns()
func (_ILendingPool *ILendingPoolTransactorSession) SwapBorrowRateMode(asset common.Address, rateMode *big.Int) (*types.Transaction, error) {
	return _ILendingPool.Contract.SwapBorrowRateMode(&_ILendingPool.TransactOpts, asset, rateMode)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_ILendingPool *ILendingPoolTransactor) Withdraw(opts *bind.TransactOpts, asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ILendingPool.contract.Transact(opts, "withdraw", asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_ILendingPool *ILendingPoolSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Withdraw(&_ILendingPool.TransactOpts, asset, amount, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x69328dec.
//
// Solidity: function withdraw(address asset, uint256 amount, address to) returns(uint256)
func (_ILendingPool *ILendingPoolTransactorSession) Withdraw(asset common.Address, amount *big.Int, to common.Address) (*types.Transaction, error) {
	return _ILendingPool.Contract.Withdraw(&_ILendingPool.TransactOpts, asset, amount, to)
}

// ILendingPoolBorrowIterator is returned from FilterBorrow and is used to iterate over the raw logs and unpacked data for Borrow events raised by the ILendingPool contract.
type ILendingPoolBorrowIterator struct {
	Event *ILendingPoolBorrow // Event containing the contract specifics and raw log

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
func (it *ILendingPoolBorrowIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolBorrow)
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
		it.Event = new(ILendingPoolBorrow)
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
func (it *ILendingPoolBorrowIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolBorrowIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolBorrow represents a Borrow event raised by the ILendingPool contract.
type ILendingPoolBorrow struct {
	Reserve        common.Address
	User           common.Address
	OnBehalfOf     common.Address
	Amount         *big.Int
	BorrowRateMode *big.Int
	BorrowRate     *big.Int
	Referral       uint16
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBorrow is a free log retrieval operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) FilterBorrow(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*ILendingPoolBorrowIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolBorrowIterator{contract: _ILendingPool.contract, event: "Borrow", logs: logs, sub: sub}, nil
}

// WatchBorrow is a free log subscription operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) WatchBorrow(opts *bind.WatchOpts, sink chan<- *ILendingPoolBorrow, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Borrow", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolBorrow)
				if err := _ILendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
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

// ParseBorrow is a log parse operation binding the contract event 0xc6a898309e823ee50bac64e45ca8adba6690e99e7841c45d754e2a38e9019d9b.
//
// Solidity: event Borrow(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint256 borrowRateMode, uint256 borrowRate, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) ParseBorrow(log types.Log) (*ILendingPoolBorrow, error) {
	event := new(ILendingPoolBorrow)
	if err := _ILendingPool.contract.UnpackLog(event, "Borrow", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolDepositIterator is returned from FilterDeposit and is used to iterate over the raw logs and unpacked data for Deposit events raised by the ILendingPool contract.
type ILendingPoolDepositIterator struct {
	Event *ILendingPoolDeposit // Event containing the contract specifics and raw log

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
func (it *ILendingPoolDepositIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolDeposit)
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
		it.Event = new(ILendingPoolDeposit)
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
func (it *ILendingPoolDepositIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolDepositIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolDeposit represents a Deposit event raised by the ILendingPool contract.
type ILendingPoolDeposit struct {
	Reserve    common.Address
	User       common.Address
	OnBehalfOf common.Address
	Amount     *big.Int
	Referral   uint16
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDeposit is a free log retrieval operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) FilterDeposit(opts *bind.FilterOpts, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (*ILendingPoolDepositIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolDepositIterator{contract: _ILendingPool.contract, event: "Deposit", logs: logs, sub: sub}, nil
}

// WatchDeposit is a free log subscription operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) WatchDeposit(opts *bind.WatchOpts, sink chan<- *ILendingPoolDeposit, reserve []common.Address, onBehalfOf []common.Address, referral []uint16) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	var referralRule []interface{}
	for _, referralItem := range referral {
		referralRule = append(referralRule, referralItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Deposit", reserveRule, onBehalfOfRule, referralRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolDeposit)
				if err := _ILendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
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

// ParseDeposit is a log parse operation binding the contract event 0xde6857219544bb5b7746f48ed30be6386fefc61b2f864cacf559893bf50fd951.
//
// Solidity: event Deposit(address indexed reserve, address user, address indexed onBehalfOf, uint256 amount, uint16 indexed referral)
func (_ILendingPool *ILendingPoolFilterer) ParseDeposit(log types.Log) (*ILendingPoolDeposit, error) {
	event := new(ILendingPoolDeposit)
	if err := _ILendingPool.contract.UnpackLog(event, "Deposit", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolFlashLoanIterator is returned from FilterFlashLoan and is used to iterate over the raw logs and unpacked data for FlashLoan events raised by the ILendingPool contract.
type ILendingPoolFlashLoanIterator struct {
	Event *ILendingPoolFlashLoan // Event containing the contract specifics and raw log

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
func (it *ILendingPoolFlashLoanIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolFlashLoan)
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
		it.Event = new(ILendingPoolFlashLoan)
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
func (it *ILendingPoolFlashLoanIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolFlashLoanIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolFlashLoan represents a FlashLoan event raised by the ILendingPool contract.
type ILendingPoolFlashLoan struct {
	Target       common.Address
	Initiator    common.Address
	Asset        common.Address
	Amount       *big.Int
	Premium      *big.Int
	ReferralCode uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFlashLoan is a free log retrieval operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_ILendingPool *ILendingPoolFilterer) FilterFlashLoan(opts *bind.FilterOpts, target []common.Address, initiator []common.Address, asset []common.Address) (*ILendingPoolFlashLoanIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "FlashLoan", targetRule, initiatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolFlashLoanIterator{contract: _ILendingPool.contract, event: "FlashLoan", logs: logs, sub: sub}, nil
}

// WatchFlashLoan is a free log subscription operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_ILendingPool *ILendingPoolFilterer) WatchFlashLoan(opts *bind.WatchOpts, sink chan<- *ILendingPoolFlashLoan, target []common.Address, initiator []common.Address, asset []common.Address) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var initiatorRule []interface{}
	for _, initiatorItem := range initiator {
		initiatorRule = append(initiatorRule, initiatorItem)
	}
	var assetRule []interface{}
	for _, assetItem := range asset {
		assetRule = append(assetRule, assetItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "FlashLoan", targetRule, initiatorRule, assetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolFlashLoan)
				if err := _ILendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
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

// ParseFlashLoan is a log parse operation binding the contract event 0x631042c832b07452973831137f2d73e395028b44b250dedc5abb0ee766e168ac.
//
// Solidity: event FlashLoan(address indexed target, address indexed initiator, address indexed asset, uint256 amount, uint256 premium, uint16 referralCode)
func (_ILendingPool *ILendingPoolFilterer) ParseFlashLoan(log types.Log) (*ILendingPoolFlashLoan, error) {
	event := new(ILendingPoolFlashLoan)
	if err := _ILendingPool.contract.UnpackLog(event, "FlashLoan", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolLiquidationCallIterator is returned from FilterLiquidationCall and is used to iterate over the raw logs and unpacked data for LiquidationCall events raised by the ILendingPool contract.
type ILendingPoolLiquidationCallIterator struct {
	Event *ILendingPoolLiquidationCall // Event containing the contract specifics and raw log

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
func (it *ILendingPoolLiquidationCallIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolLiquidationCall)
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
		it.Event = new(ILendingPoolLiquidationCall)
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
func (it *ILendingPoolLiquidationCallIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolLiquidationCallIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolLiquidationCall represents a LiquidationCall event raised by the ILendingPool contract.
type ILendingPoolLiquidationCall struct {
	CollateralAsset            common.Address
	DebtAsset                  common.Address
	User                       common.Address
	DebtToCover                *big.Int
	LiquidatedCollateralAmount *big.Int
	Liquidator                 common.Address
	ReceiveAToken              bool
	Raw                        types.Log // Blockchain specific contextual infos
}

// FilterLiquidationCall is a free log retrieval operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_ILendingPool *ILendingPoolFilterer) FilterLiquidationCall(opts *bind.FilterOpts, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (*ILendingPoolLiquidationCallIterator, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolLiquidationCallIterator{contract: _ILendingPool.contract, event: "LiquidationCall", logs: logs, sub: sub}, nil
}

// WatchLiquidationCall is a free log subscription operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_ILendingPool *ILendingPoolFilterer) WatchLiquidationCall(opts *bind.WatchOpts, sink chan<- *ILendingPoolLiquidationCall, collateralAsset []common.Address, debtAsset []common.Address, user []common.Address) (event.Subscription, error) {

	var collateralAssetRule []interface{}
	for _, collateralAssetItem := range collateralAsset {
		collateralAssetRule = append(collateralAssetRule, collateralAssetItem)
	}
	var debtAssetRule []interface{}
	for _, debtAssetItem := range debtAsset {
		debtAssetRule = append(debtAssetRule, debtAssetItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "LiquidationCall", collateralAssetRule, debtAssetRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolLiquidationCall)
				if err := _ILendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
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

// ParseLiquidationCall is a log parse operation binding the contract event 0xe413a321e8681d831f4dbccbca790d2952b56f977908e45be37335533e005286.
//
// Solidity: event LiquidationCall(address indexed collateralAsset, address indexed debtAsset, address indexed user, uint256 debtToCover, uint256 liquidatedCollateralAmount, address liquidator, bool receiveAToken)
func (_ILendingPool *ILendingPoolFilterer) ParseLiquidationCall(log types.Log) (*ILendingPoolLiquidationCall, error) {
	event := new(ILendingPoolLiquidationCall)
	if err := _ILendingPool.contract.UnpackLog(event, "LiquidationCall", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the ILendingPool contract.
type ILendingPoolPausedIterator struct {
	Event *ILendingPoolPaused // Event containing the contract specifics and raw log

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
func (it *ILendingPoolPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolPaused)
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
		it.Event = new(ILendingPoolPaused)
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
func (it *ILendingPoolPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolPaused represents a Paused event raised by the ILendingPool contract.
type ILendingPoolPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ILendingPool *ILendingPoolFilterer) FilterPaused(opts *bind.FilterOpts) (*ILendingPoolPausedIterator, error) {

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &ILendingPoolPausedIterator{contract: _ILendingPool.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ILendingPool *ILendingPoolFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *ILendingPoolPaused) (event.Subscription, error) {

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolPaused)
				if err := _ILendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_ILendingPool *ILendingPoolFilterer) ParsePaused(log types.Log) (*ILendingPoolPaused, error) {
	event := new(ILendingPoolPaused)
	if err := _ILendingPool.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolRebalanceStableBorrowRateIterator is returned from FilterRebalanceStableBorrowRate and is used to iterate over the raw logs and unpacked data for RebalanceStableBorrowRate events raised by the ILendingPool contract.
type ILendingPoolRebalanceStableBorrowRateIterator struct {
	Event *ILendingPoolRebalanceStableBorrowRate // Event containing the contract specifics and raw log

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
func (it *ILendingPoolRebalanceStableBorrowRateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolRebalanceStableBorrowRate)
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
		it.Event = new(ILendingPoolRebalanceStableBorrowRate)
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
func (it *ILendingPoolRebalanceStableBorrowRateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolRebalanceStableBorrowRateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolRebalanceStableBorrowRate represents a RebalanceStableBorrowRate event raised by the ILendingPool contract.
type ILendingPoolRebalanceStableBorrowRate struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRebalanceStableBorrowRate is a free log retrieval operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) FilterRebalanceStableBorrowRate(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*ILendingPoolRebalanceStableBorrowRateIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolRebalanceStableBorrowRateIterator{contract: _ILendingPool.contract, event: "RebalanceStableBorrowRate", logs: logs, sub: sub}, nil
}

// WatchRebalanceStableBorrowRate is a free log subscription operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) WatchRebalanceStableBorrowRate(opts *bind.WatchOpts, sink chan<- *ILendingPoolRebalanceStableBorrowRate, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "RebalanceStableBorrowRate", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolRebalanceStableBorrowRate)
				if err := _ILendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
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

// ParseRebalanceStableBorrowRate is a log parse operation binding the contract event 0x9f439ae0c81e41a04d3fdfe07aed54e6a179fb0db15be7702eb66fa8ef6f5300.
//
// Solidity: event RebalanceStableBorrowRate(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) ParseRebalanceStableBorrowRate(log types.Log) (*ILendingPoolRebalanceStableBorrowRate, error) {
	event := new(ILendingPoolRebalanceStableBorrowRate)
	if err := _ILendingPool.contract.UnpackLog(event, "RebalanceStableBorrowRate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolRepayIterator is returned from FilterRepay and is used to iterate over the raw logs and unpacked data for Repay events raised by the ILendingPool contract.
type ILendingPoolRepayIterator struct {
	Event *ILendingPoolRepay // Event containing the contract specifics and raw log

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
func (it *ILendingPoolRepayIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolRepay)
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
		it.Event = new(ILendingPoolRepay)
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
func (it *ILendingPoolRepayIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolRepayIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolRepay represents a Repay event raised by the ILendingPool contract.
type ILendingPoolRepay struct {
	Reserve common.Address
	User    common.Address
	Repayer common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRepay is a free log retrieval operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) FilterRepay(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, repayer []common.Address) (*ILendingPoolRepayIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolRepayIterator{contract: _ILendingPool.contract, event: "Repay", logs: logs, sub: sub}, nil
}

// WatchRepay is a free log subscription operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) WatchRepay(opts *bind.WatchOpts, sink chan<- *ILendingPoolRepay, reserve []common.Address, user []common.Address, repayer []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var repayerRule []interface{}
	for _, repayerItem := range repayer {
		repayerRule = append(repayerRule, repayerItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Repay", reserveRule, userRule, repayerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolRepay)
				if err := _ILendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
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

// ParseRepay is a log parse operation binding the contract event 0x4cdde6e09bb755c9a5589ebaec640bbfedff1362d4b255ebf8339782b9942faa.
//
// Solidity: event Repay(address indexed reserve, address indexed user, address indexed repayer, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) ParseRepay(log types.Log) (*ILendingPoolRepay, error) {
	event := new(ILendingPoolRepay)
	if err := _ILendingPool.contract.UnpackLog(event, "Repay", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolReserveDataUpdatedIterator is returned from FilterReserveDataUpdated and is used to iterate over the raw logs and unpacked data for ReserveDataUpdated events raised by the ILendingPool contract.
type ILendingPoolReserveDataUpdatedIterator struct {
	Event *ILendingPoolReserveDataUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolReserveDataUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolReserveDataUpdated)
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
		it.Event = new(ILendingPoolReserveDataUpdated)
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
func (it *ILendingPoolReserveDataUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolReserveDataUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolReserveDataUpdated represents a ReserveDataUpdated event raised by the ILendingPool contract.
type ILendingPoolReserveDataUpdated struct {
	Reserve             common.Address
	LiquidityRate       *big.Int
	StableBorrowRate    *big.Int
	VariableBorrowRate  *big.Int
	LiquidityIndex      *big.Int
	VariableBorrowIndex *big.Int
	Raw                 types.Log // Blockchain specific contextual infos
}

// FilterReserveDataUpdated is a free log retrieval operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_ILendingPool *ILendingPoolFilterer) FilterReserveDataUpdated(opts *bind.FilterOpts, reserve []common.Address) (*ILendingPoolReserveDataUpdatedIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolReserveDataUpdatedIterator{contract: _ILendingPool.contract, event: "ReserveDataUpdated", logs: logs, sub: sub}, nil
}

// WatchReserveDataUpdated is a free log subscription operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_ILendingPool *ILendingPoolFilterer) WatchReserveDataUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolReserveDataUpdated, reserve []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "ReserveDataUpdated", reserveRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolReserveDataUpdated)
				if err := _ILendingPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
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

// ParseReserveDataUpdated is a log parse operation binding the contract event 0x804c9b842b2748a22bb64b345453a3de7ca54a6ca45ce00d415894979e22897a.
//
// Solidity: event ReserveDataUpdated(address indexed reserve, uint256 liquidityRate, uint256 stableBorrowRate, uint256 variableBorrowRate, uint256 liquidityIndex, uint256 variableBorrowIndex)
func (_ILendingPool *ILendingPoolFilterer) ParseReserveDataUpdated(log types.Log) (*ILendingPoolReserveDataUpdated, error) {
	event := new(ILendingPoolReserveDataUpdated)
	if err := _ILendingPool.contract.UnpackLog(event, "ReserveDataUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolReserveUsedAsCollateralDisabledIterator is returned from FilterReserveUsedAsCollateralDisabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralDisabled events raised by the ILendingPool contract.
type ILendingPoolReserveUsedAsCollateralDisabledIterator struct {
	Event *ILendingPoolReserveUsedAsCollateralDisabled // Event containing the contract specifics and raw log

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
func (it *ILendingPoolReserveUsedAsCollateralDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolReserveUsedAsCollateralDisabled)
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
		it.Event = new(ILendingPoolReserveUsedAsCollateralDisabled)
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
func (it *ILendingPoolReserveUsedAsCollateralDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolReserveUsedAsCollateralDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolReserveUsedAsCollateralDisabled represents a ReserveUsedAsCollateralDisabled event raised by the ILendingPool contract.
type ILendingPoolReserveUsedAsCollateralDisabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralDisabled is a free log retrieval operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) FilterReserveUsedAsCollateralDisabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*ILendingPoolReserveUsedAsCollateralDisabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolReserveUsedAsCollateralDisabledIterator{contract: _ILendingPool.contract, event: "ReserveUsedAsCollateralDisabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralDisabled is a free log subscription operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) WatchReserveUsedAsCollateralDisabled(opts *bind.WatchOpts, sink chan<- *ILendingPoolReserveUsedAsCollateralDisabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralDisabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolReserveUsedAsCollateralDisabled)
				if err := _ILendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
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

// ParseReserveUsedAsCollateralDisabled is a log parse operation binding the contract event 0x44c58d81365b66dd4b1a7f36c25aa97b8c71c361ee4937adc1a00000227db5dd.
//
// Solidity: event ReserveUsedAsCollateralDisabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) ParseReserveUsedAsCollateralDisabled(log types.Log) (*ILendingPoolReserveUsedAsCollateralDisabled, error) {
	event := new(ILendingPoolReserveUsedAsCollateralDisabled)
	if err := _ILendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolReserveUsedAsCollateralEnabledIterator is returned from FilterReserveUsedAsCollateralEnabled and is used to iterate over the raw logs and unpacked data for ReserveUsedAsCollateralEnabled events raised by the ILendingPool contract.
type ILendingPoolReserveUsedAsCollateralEnabledIterator struct {
	Event *ILendingPoolReserveUsedAsCollateralEnabled // Event containing the contract specifics and raw log

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
func (it *ILendingPoolReserveUsedAsCollateralEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolReserveUsedAsCollateralEnabled)
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
		it.Event = new(ILendingPoolReserveUsedAsCollateralEnabled)
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
func (it *ILendingPoolReserveUsedAsCollateralEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolReserveUsedAsCollateralEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolReserveUsedAsCollateralEnabled represents a ReserveUsedAsCollateralEnabled event raised by the ILendingPool contract.
type ILendingPoolReserveUsedAsCollateralEnabled struct {
	Reserve common.Address
	User    common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterReserveUsedAsCollateralEnabled is a free log retrieval operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) FilterReserveUsedAsCollateralEnabled(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*ILendingPoolReserveUsedAsCollateralEnabledIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolReserveUsedAsCollateralEnabledIterator{contract: _ILendingPool.contract, event: "ReserveUsedAsCollateralEnabled", logs: logs, sub: sub}, nil
}

// WatchReserveUsedAsCollateralEnabled is a free log subscription operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) WatchReserveUsedAsCollateralEnabled(opts *bind.WatchOpts, sink chan<- *ILendingPoolReserveUsedAsCollateralEnabled, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "ReserveUsedAsCollateralEnabled", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolReserveUsedAsCollateralEnabled)
				if err := _ILendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
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

// ParseReserveUsedAsCollateralEnabled is a log parse operation binding the contract event 0x00058a56ea94653cdf4f152d227ace22d4c00ad99e2a43f58cb7d9e3feb295f2.
//
// Solidity: event ReserveUsedAsCollateralEnabled(address indexed reserve, address indexed user)
func (_ILendingPool *ILendingPoolFilterer) ParseReserveUsedAsCollateralEnabled(log types.Log) (*ILendingPoolReserveUsedAsCollateralEnabled, error) {
	event := new(ILendingPoolReserveUsedAsCollateralEnabled)
	if err := _ILendingPool.contract.UnpackLog(event, "ReserveUsedAsCollateralEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolSwapIterator is returned from FilterSwap and is used to iterate over the raw logs and unpacked data for Swap events raised by the ILendingPool contract.
type ILendingPoolSwapIterator struct {
	Event *ILendingPoolSwap // Event containing the contract specifics and raw log

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
func (it *ILendingPoolSwapIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolSwap)
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
		it.Event = new(ILendingPoolSwap)
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
func (it *ILendingPoolSwapIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolSwapIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolSwap represents a Swap event raised by the ILendingPool contract.
type ILendingPoolSwap struct {
	Reserve  common.Address
	User     common.Address
	RateMode *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterSwap is a free log retrieval operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_ILendingPool *ILendingPoolFilterer) FilterSwap(opts *bind.FilterOpts, reserve []common.Address, user []common.Address) (*ILendingPoolSwapIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Swap", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolSwapIterator{contract: _ILendingPool.contract, event: "Swap", logs: logs, sub: sub}, nil
}

// WatchSwap is a free log subscription operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_ILendingPool *ILendingPoolFilterer) WatchSwap(opts *bind.WatchOpts, sink chan<- *ILendingPoolSwap, reserve []common.Address, user []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Swap", reserveRule, userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolSwap)
				if err := _ILendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
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

// ParseSwap is a log parse operation binding the contract event 0xea368a40e9570069bb8e6511d668293ad2e1f03b0d982431fd223de9f3b70ca6.
//
// Solidity: event Swap(address indexed reserve, address indexed user, uint256 rateMode)
func (_ILendingPool *ILendingPoolFilterer) ParseSwap(log types.Log) (*ILendingPoolSwap, error) {
	event := new(ILendingPoolSwap)
	if err := _ILendingPool.contract.UnpackLog(event, "Swap", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the ILendingPool contract.
type ILendingPoolUnpausedIterator struct {
	Event *ILendingPoolUnpaused // Event containing the contract specifics and raw log

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
func (it *ILendingPoolUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolUnpaused)
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
		it.Event = new(ILendingPoolUnpaused)
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
func (it *ILendingPoolUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolUnpaused represents a Unpaused event raised by the ILendingPool contract.
type ILendingPoolUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ILendingPool *ILendingPoolFilterer) FilterUnpaused(opts *bind.FilterOpts) (*ILendingPoolUnpausedIterator, error) {

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &ILendingPoolUnpausedIterator{contract: _ILendingPool.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ILendingPool *ILendingPoolFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *ILendingPoolUnpaused) (event.Subscription, error) {

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolUnpaused)
				if err := _ILendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_ILendingPool *ILendingPoolFilterer) ParseUnpaused(log types.Log) (*ILendingPoolUnpaused, error) {
	event := new(ILendingPoolUnpaused)
	if err := _ILendingPool.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the ILendingPool contract.
type ILendingPoolWithdrawIterator struct {
	Event *ILendingPoolWithdraw // Event containing the contract specifics and raw log

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
func (it *ILendingPoolWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolWithdraw)
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
		it.Event = new(ILendingPoolWithdraw)
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
func (it *ILendingPoolWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolWithdraw represents a Withdraw event raised by the ILendingPool contract.
type ILendingPoolWithdraw struct {
	Reserve common.Address
	User    common.Address
	To      common.Address
	Amount  *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) FilterWithdraw(opts *bind.FilterOpts, reserve []common.Address, user []common.Address, to []common.Address) (*ILendingPoolWithdrawIterator, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILendingPool.contract.FilterLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolWithdrawIterator{contract: _ILendingPool.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *ILendingPoolWithdraw, reserve []common.Address, user []common.Address, to []common.Address) (event.Subscription, error) {

	var reserveRule []interface{}
	for _, reserveItem := range reserve {
		reserveRule = append(reserveRule, reserveItem)
	}
	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _ILendingPool.contract.WatchLogs(opts, "Withdraw", reserveRule, userRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolWithdraw)
				if err := _ILendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x3115d1449a7b732c986cba18244e897a450f61e1bb8d589cd2e69e6c8924f9f7.
//
// Solidity: event Withdraw(address indexed reserve, address indexed user, address indexed to, uint256 amount)
func (_ILendingPool *ILendingPoolFilterer) ParseWithdraw(log types.Log) (*ILendingPoolWithdraw, error) {
	event := new(ILendingPoolWithdraw)
	if err := _ILendingPool.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderMetaData contains all meta data concerning the ILendingPoolAddressesProvider contract.
var ILendingPoolAddressesProviderMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"hasProxy\",\"type\":\"bool\"}],\"name\":\"AddressSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ConfigurationAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"EmergencyAdminUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolCollateralManagerUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolConfiguratorUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingPoolUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"LendingRateOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"newMarketId\",\"type\":\"string\"}],\"name\":\"MarketIdSet\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"PriceOracleUpdated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"ProxyCreated\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"}],\"name\":\"getAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEmergencyAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPool\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolCollateralManager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingPoolConfigurator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getLendingRateOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMarketId\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPoolAdmin\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPriceOracle\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"setAddress\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"id\",\"type\":\"bytes32\"},{\"internalType\":\"address\",\"name\":\"impl\",\"type\":\"address\"}],\"name\":\"setAddressAsProxy\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setEmergencyAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"manager\",\"type\":\"address\"}],\"name\":\"setLendingPoolCollateralManager\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"configurator\",\"type\":\"address\"}],\"name\":\"setLendingPoolConfiguratorImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"pool\",\"type\":\"address\"}],\"name\":\"setLendingPoolImpl\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lendingRateOracle\",\"type\":\"address\"}],\"name\":\"setLendingRateOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"marketId\",\"type\":\"string\"}],\"name\":\"setMarketId\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"setPoolAdmin\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"priceOracle\",\"type\":\"address\"}],\"name\":\"setPriceOracle\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ILendingPoolAddressesProviderABI is the input ABI used to generate the binding from.
// Deprecated: Use ILendingPoolAddressesProviderMetaData.ABI instead.
var ILendingPoolAddressesProviderABI = ILendingPoolAddressesProviderMetaData.ABI

// ILendingPoolAddressesProvider is an auto generated Go binding around an Ethereum contract.
type ILendingPoolAddressesProvider struct {
	ILendingPoolAddressesProviderCaller     // Read-only binding to the contract
	ILendingPoolAddressesProviderTransactor // Write-only binding to the contract
	ILendingPoolAddressesProviderFilterer   // Log filterer for contract events
}

// ILendingPoolAddressesProviderCaller is an auto generated read-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ILendingPoolAddressesProviderFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ILendingPoolAddressesProviderSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ILendingPoolAddressesProviderSession struct {
	Contract     *ILendingPoolAddressesProvider // Generic contract binding to set the session for
	CallOpts     bind.CallOpts                  // Call options to use throughout this session
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// ILendingPoolAddressesProviderCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ILendingPoolAddressesProviderCallerSession struct {
	Contract *ILendingPoolAddressesProviderCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                        // Call options to use throughout this session
}

// ILendingPoolAddressesProviderTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ILendingPoolAddressesProviderTransactorSession struct {
	Contract     *ILendingPoolAddressesProviderTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                        // Transaction auth options to use throughout this session
}

// ILendingPoolAddressesProviderRaw is an auto generated low-level Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderRaw struct {
	Contract *ILendingPoolAddressesProvider // Generic contract binding to access the raw methods on
}

// ILendingPoolAddressesProviderCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderCallerRaw struct {
	Contract *ILendingPoolAddressesProviderCaller // Generic read-only contract binding to access the raw methods on
}

// ILendingPoolAddressesProviderTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ILendingPoolAddressesProviderTransactorRaw struct {
	Contract *ILendingPoolAddressesProviderTransactor // Generic write-only contract binding to access the raw methods on
}

// NewILendingPoolAddressesProvider creates a new instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProvider(address common.Address, backend bind.ContractBackend) (*ILendingPoolAddressesProvider, error) {
	contract, err := bindILendingPoolAddressesProvider(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProvider{ILendingPoolAddressesProviderCaller: ILendingPoolAddressesProviderCaller{contract: contract}, ILendingPoolAddressesProviderTransactor: ILendingPoolAddressesProviderTransactor{contract: contract}, ILendingPoolAddressesProviderFilterer: ILendingPoolAddressesProviderFilterer{contract: contract}}, nil
}

// NewILendingPoolAddressesProviderCaller creates a new read-only instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderCaller(address common.Address, caller bind.ContractCaller) (*ILendingPoolAddressesProviderCaller, error) {
	contract, err := bindILendingPoolAddressesProvider(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderCaller{contract: contract}, nil
}

// NewILendingPoolAddressesProviderTransactor creates a new write-only instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderTransactor(address common.Address, transactor bind.ContractTransactor) (*ILendingPoolAddressesProviderTransactor, error) {
	contract, err := bindILendingPoolAddressesProvider(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderTransactor{contract: contract}, nil
}

// NewILendingPoolAddressesProviderFilterer creates a new log filterer instance of ILendingPoolAddressesProvider, bound to a specific deployed contract.
func NewILendingPoolAddressesProviderFilterer(address common.Address, filterer bind.ContractFilterer) (*ILendingPoolAddressesProviderFilterer, error) {
	contract, err := bindILendingPoolAddressesProvider(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderFilterer{contract: contract}, nil
}

// bindILendingPoolAddressesProvider binds a generic wrapper to an already deployed contract.
func bindILendingPoolAddressesProvider(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ILendingPoolAddressesProviderABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.ILendingPoolAddressesProviderTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ILendingPoolAddressesProvider.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.contract.Transact(opts, method, params...)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetAddress(opts *bind.CallOpts, id [32]byte) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getAddress", id)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetAddress(id [32]byte) (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetAddress(&_ILendingPoolAddressesProvider.CallOpts, id)
}

// GetAddress is a free data retrieval call binding the contract method 0x21f8a721.
//
// Solidity: function getAddress(bytes32 id) view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetAddress(id [32]byte) (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetAddress(&_ILendingPoolAddressesProvider.CallOpts, id)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetEmergencyAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getEmergencyAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetEmergencyAdmin() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetEmergencyAdmin(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetEmergencyAdmin is a free data retrieval call binding the contract method 0xddcaa9ea.
//
// Solidity: function getEmergencyAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetEmergencyAdmin() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetEmergencyAdmin(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingPool(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPool")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingPool() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPool(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPool is a free data retrieval call binding the contract method 0x0261bf8b.
//
// Solidity: function getLendingPool() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingPool() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPool(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingPoolCollateralManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolCollateralManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolCollateralManager(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolCollateralManager is a free data retrieval call binding the contract method 0x712d9171.
//
// Solidity: function getLendingPoolCollateralManager() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingPoolCollateralManager() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolCollateralManager(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingPoolConfigurator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingPoolConfigurator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingPoolConfigurator is a free data retrieval call binding the contract method 0x85c858b1.
//
// Solidity: function getLendingPoolConfigurator() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingPoolConfigurator() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingPoolConfigurator(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetLendingRateOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getLendingRateOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetLendingRateOracle() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetLendingRateOracle is a free data retrieval call binding the contract method 0x3618abba.
//
// Solidity: function getLendingRateOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetLendingRateOracle() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetLendingRateOracle(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetMarketId(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getMarketId")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetMarketId() (string, error) {
	return _ILendingPoolAddressesProvider.Contract.GetMarketId(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetMarketId is a free data retrieval call binding the contract method 0x568ef470.
//
// Solidity: function getMarketId() view returns(string)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetMarketId() (string, error) {
	return _ILendingPoolAddressesProvider.Contract.GetMarketId(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetPoolAdmin(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getPoolAdmin")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetPoolAdmin() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetPoolAdmin(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetPoolAdmin is a free data retrieval call binding the contract method 0xaecda378.
//
// Solidity: function getPoolAdmin() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetPoolAdmin() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetPoolAdmin(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCaller) GetPriceOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ILendingPoolAddressesProvider.contract.Call(opts, &out, "getPriceOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) GetPriceOracle() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetPriceOracle(&_ILendingPoolAddressesProvider.CallOpts)
}

// GetPriceOracle is a free data retrieval call binding the contract method 0xfca513a8.
//
// Solidity: function getPriceOracle() view returns(address)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderCallerSession) GetPriceOracle() (common.Address, error) {
	return _ILendingPoolAddressesProvider.Contract.GetPriceOracle(&_ILendingPoolAddressesProvider.CallOpts)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetAddress(opts *bind.TransactOpts, id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setAddress", id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetAddress(&_ILendingPoolAddressesProvider.TransactOpts, id, newAddress)
}

// SetAddress is a paid mutator transaction binding the contract method 0xca446dd9.
//
// Solidity: function setAddress(bytes32 id, address newAddress) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetAddress(id [32]byte, newAddress common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetAddress(&_ILendingPoolAddressesProvider.TransactOpts, id, newAddress)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetAddressAsProxy(opts *bind.TransactOpts, id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setAddressAsProxy", id, impl)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetAddressAsProxy(id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetAddressAsProxy(&_ILendingPoolAddressesProvider.TransactOpts, id, impl)
}

// SetAddressAsProxy is a paid mutator transaction binding the contract method 0x5dcc528c.
//
// Solidity: function setAddressAsProxy(bytes32 id, address impl) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetAddressAsProxy(id [32]byte, impl common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetAddressAsProxy(&_ILendingPoolAddressesProvider.TransactOpts, id, impl)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetEmergencyAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setEmergencyAdmin", admin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetEmergencyAdmin(admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetEmergencyAdmin(&_ILendingPoolAddressesProvider.TransactOpts, admin)
}

// SetEmergencyAdmin is a paid mutator transaction binding the contract method 0x35da3394.
//
// Solidity: function setEmergencyAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetEmergencyAdmin(admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetEmergencyAdmin(&_ILendingPoolAddressesProvider.TransactOpts, admin)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetLendingPoolCollateralManager(opts *bind.TransactOpts, manager common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolCollateralManager", manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolCollateralManager(&_ILendingPoolAddressesProvider.TransactOpts, manager)
}

// SetLendingPoolCollateralManager is a paid mutator transaction binding the contract method 0x398e5553.
//
// Solidity: function setLendingPoolCollateralManager(address manager) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetLendingPoolCollateralManager(manager common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolCollateralManager(&_ILendingPoolAddressesProvider.TransactOpts, manager)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetLendingPoolConfiguratorImpl(opts *bind.TransactOpts, configurator common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolConfiguratorImpl", configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_ILendingPoolAddressesProvider.TransactOpts, configurator)
}

// SetLendingPoolConfiguratorImpl is a paid mutator transaction binding the contract method 0xc12542df.
//
// Solidity: function setLendingPoolConfiguratorImpl(address configurator) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetLendingPoolConfiguratorImpl(configurator common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolConfiguratorImpl(&_ILendingPoolAddressesProvider.TransactOpts, configurator)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetLendingPoolImpl(opts *bind.TransactOpts, pool common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setLendingPoolImpl", pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_ILendingPoolAddressesProvider.TransactOpts, pool)
}

// SetLendingPoolImpl is a paid mutator transaction binding the contract method 0x5aef021f.
//
// Solidity: function setLendingPoolImpl(address pool) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetLendingPoolImpl(pool common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingPoolImpl(&_ILendingPoolAddressesProvider.TransactOpts, pool)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetLendingRateOracle(opts *bind.TransactOpts, lendingRateOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setLendingRateOracle", lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_ILendingPoolAddressesProvider.TransactOpts, lendingRateOracle)
}

// SetLendingRateOracle is a paid mutator transaction binding the contract method 0x820d1274.
//
// Solidity: function setLendingRateOracle(address lendingRateOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetLendingRateOracle(lendingRateOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetLendingRateOracle(&_ILendingPoolAddressesProvider.TransactOpts, lendingRateOracle)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetMarketId(opts *bind.TransactOpts, marketId string) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setMarketId", marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetMarketId(&_ILendingPoolAddressesProvider.TransactOpts, marketId)
}

// SetMarketId is a paid mutator transaction binding the contract method 0xf67b1847.
//
// Solidity: function setMarketId(string marketId) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetMarketId(marketId string) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetMarketId(&_ILendingPoolAddressesProvider.TransactOpts, marketId)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetPoolAdmin(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setPoolAdmin", admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetPoolAdmin(&_ILendingPoolAddressesProvider.TransactOpts, admin)
}

// SetPoolAdmin is a paid mutator transaction binding the contract method 0x283d62ad.
//
// Solidity: function setPoolAdmin(address admin) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetPoolAdmin(admin common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetPoolAdmin(&_ILendingPoolAddressesProvider.TransactOpts, admin)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactor) SetPriceOracle(opts *bind.TransactOpts, priceOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.contract.Transact(opts, "setPriceOracle", priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetPriceOracle(&_ILendingPoolAddressesProvider.TransactOpts, priceOracle)
}

// SetPriceOracle is a paid mutator transaction binding the contract method 0x530e784f.
//
// Solidity: function setPriceOracle(address priceOracle) returns()
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderTransactorSession) SetPriceOracle(priceOracle common.Address) (*types.Transaction, error) {
	return _ILendingPoolAddressesProvider.Contract.SetPriceOracle(&_ILendingPoolAddressesProvider.TransactOpts, priceOracle)
}

// ILendingPoolAddressesProviderAddressSetIterator is returned from FilterAddressSet and is used to iterate over the raw logs and unpacked data for AddressSet events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderAddressSetIterator struct {
	Event *ILendingPoolAddressesProviderAddressSet // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderAddressSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderAddressSet)
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
		it.Event = new(ILendingPoolAddressesProviderAddressSet)
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
func (it *ILendingPoolAddressesProviderAddressSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderAddressSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderAddressSet represents a AddressSet event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderAddressSet struct {
	Id         [32]byte
	NewAddress common.Address
	HasProxy   bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAddressSet is a free log retrieval operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterAddressSet(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderAddressSetIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderAddressSetIterator{contract: _ILendingPoolAddressesProvider.contract, event: "AddressSet", logs: logs, sub: sub}, nil
}

// WatchAddressSet is a free log subscription operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchAddressSet(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderAddressSet, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "AddressSet", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderAddressSet)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
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

// ParseAddressSet is a log parse operation binding the contract event 0xf2689d5d5cd0c639e137642cae5d40afced201a1a0327e7ac9358461dc9fff31.
//
// Solidity: event AddressSet(bytes32 id, address indexed newAddress, bool hasProxy)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseAddressSet(log types.Log) (*ILendingPoolAddressesProviderAddressSet, error) {
	event := new(ILendingPoolAddressesProviderAddressSet)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "AddressSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator is returned from FilterConfigurationAdminUpdated and is used to iterate over the raw logs and unpacked data for ConfigurationAdminUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderConfigurationAdminUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderConfigurationAdminUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderConfigurationAdminUpdated)
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
func (it *ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderConfigurationAdminUpdated represents a ConfigurationAdminUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderConfigurationAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterConfigurationAdminUpdated is a free log retrieval operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterConfigurationAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderConfigurationAdminUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "ConfigurationAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchConfigurationAdminUpdated is a free log subscription operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchConfigurationAdminUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderConfigurationAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "ConfigurationAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderConfigurationAdminUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
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

// ParseConfigurationAdminUpdated is a log parse operation binding the contract event 0xc20a317155a9e7d84e06b716b4b355d47742ab9f8c5d630e7f556553f582430d.
//
// Solidity: event ConfigurationAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseConfigurationAdminUpdated(log types.Log) (*ILendingPoolAddressesProviderConfigurationAdminUpdated, error) {
	event := new(ILendingPoolAddressesProviderConfigurationAdminUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "ConfigurationAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator is returned from FilterEmergencyAdminUpdated and is used to iterate over the raw logs and unpacked data for EmergencyAdminUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderEmergencyAdminUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderEmergencyAdminUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderEmergencyAdminUpdated)
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
func (it *ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderEmergencyAdminUpdated represents a EmergencyAdminUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderEmergencyAdminUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterEmergencyAdminUpdated is a free log retrieval operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterEmergencyAdminUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderEmergencyAdminUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "EmergencyAdminUpdated", logs: logs, sub: sub}, nil
}

// WatchEmergencyAdminUpdated is a free log subscription operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchEmergencyAdminUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderEmergencyAdminUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "EmergencyAdminUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderEmergencyAdminUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
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

// ParseEmergencyAdminUpdated is a log parse operation binding the contract event 0xe19673fc861bfeb894cf2d6b7662505497ef31c0f489b742db24ee3310826916.
//
// Solidity: event EmergencyAdminUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseEmergencyAdminUpdated(log types.Log) (*ILendingPoolAddressesProviderEmergencyAdminUpdated, error) {
	event := new(ILendingPoolAddressesProviderEmergencyAdminUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "EmergencyAdminUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator is returned from FilterLendingPoolCollateralManagerUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolCollateralManagerUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
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
func (it *ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated represents a LendingPoolCollateralManagerUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolCollateralManagerUpdated is a free log retrieval operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterLendingPoolCollateralManagerUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "LendingPoolCollateralManagerUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolCollateralManagerUpdated is a free log subscription operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchLendingPoolCollateralManagerUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolCollateralManagerUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
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

// ParseLendingPoolCollateralManagerUpdated is a log parse operation binding the contract event 0x991888326f0eab3df6084aadb82bee6781b5c9aa75379e8bc50ae86934541638.
//
// Solidity: event LendingPoolCollateralManagerUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseLendingPoolCollateralManagerUpdated(log types.Log) (*ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated, error) {
	event := new(ILendingPoolAddressesProviderLendingPoolCollateralManagerUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolCollateralManagerUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator is returned from FilterLendingPoolConfiguratorUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolConfiguratorUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
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
func (it *ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated represents a LendingPoolConfiguratorUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolConfiguratorUpdated is a free log retrieval operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterLendingPoolConfiguratorUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderLendingPoolConfiguratorUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "LendingPoolConfiguratorUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolConfiguratorUpdated is a free log subscription operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchLendingPoolConfiguratorUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolConfiguratorUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
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

// ParseLendingPoolConfiguratorUpdated is a log parse operation binding the contract event 0xdfabe479bad36782fb1e77fbfddd4e382671713527e4786cfc93a022ae763729.
//
// Solidity: event LendingPoolConfiguratorUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseLendingPoolConfiguratorUpdated(log types.Log) (*ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated, error) {
	event := new(ILendingPoolAddressesProviderLendingPoolConfiguratorUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolConfiguratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderLendingPoolUpdatedIterator is returned from FilterLendingPoolUpdated and is used to iterate over the raw logs and unpacked data for LendingPoolUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderLendingPoolUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderLendingPoolUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderLendingPoolUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderLendingPoolUpdated)
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
func (it *ILendingPoolAddressesProviderLendingPoolUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderLendingPoolUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderLendingPoolUpdated represents a LendingPoolUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingPoolUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingPoolUpdated is a free log retrieval operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterLendingPoolUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderLendingPoolUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderLendingPoolUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "LendingPoolUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingPoolUpdated is a free log subscription operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchLendingPoolUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderLendingPoolUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingPoolUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderLendingPoolUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
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

// ParseLendingPoolUpdated is a log parse operation binding the contract event 0xc4e6c6cdf28d0edbd8bcf071d724d33cc2e7a30be7d06443925656e9cb492aa4.
//
// Solidity: event LendingPoolUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseLendingPoolUpdated(log types.Log) (*ILendingPoolAddressesProviderLendingPoolUpdated, error) {
	event := new(ILendingPoolAddressesProviderLendingPoolUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingPoolUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator is returned from FilterLendingRateOracleUpdated and is used to iterate over the raw logs and unpacked data for LendingRateOracleUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderLendingRateOracleUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderLendingRateOracleUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderLendingRateOracleUpdated)
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
func (it *ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderLendingRateOracleUpdated represents a LendingRateOracleUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderLendingRateOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLendingRateOracleUpdated is a free log retrieval operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterLendingRateOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderLendingRateOracleUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "LendingRateOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchLendingRateOracleUpdated is a free log subscription operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchLendingRateOracleUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderLendingRateOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "LendingRateOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderLendingRateOracleUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
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

// ParseLendingRateOracleUpdated is a log parse operation binding the contract event 0x5c29179aba6942020a8a2d38f65de02fb6b7f784e7f049ed3a3cab97621859b5.
//
// Solidity: event LendingRateOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseLendingRateOracleUpdated(log types.Log) (*ILendingPoolAddressesProviderLendingRateOracleUpdated, error) {
	event := new(ILendingPoolAddressesProviderLendingRateOracleUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "LendingRateOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderMarketIdSetIterator is returned from FilterMarketIdSet and is used to iterate over the raw logs and unpacked data for MarketIdSet events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderMarketIdSetIterator struct {
	Event *ILendingPoolAddressesProviderMarketIdSet // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderMarketIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderMarketIdSet)
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
		it.Event = new(ILendingPoolAddressesProviderMarketIdSet)
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
func (it *ILendingPoolAddressesProviderMarketIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderMarketIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderMarketIdSet represents a MarketIdSet event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderMarketIdSet struct {
	NewMarketId string
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterMarketIdSet is a free log retrieval operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterMarketIdSet(opts *bind.FilterOpts) (*ILendingPoolAddressesProviderMarketIdSetIterator, error) {

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderMarketIdSetIterator{contract: _ILendingPoolAddressesProvider.contract, event: "MarketIdSet", logs: logs, sub: sub}, nil
}

// WatchMarketIdSet is a free log subscription operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchMarketIdSet(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderMarketIdSet) (event.Subscription, error) {

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "MarketIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderMarketIdSet)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
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

// ParseMarketIdSet is a log parse operation binding the contract event 0x5e667c32fd847cf8bce48ab3400175cbf107bdc82b2dea62e3364909dfaee799.
//
// Solidity: event MarketIdSet(string newMarketId)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseMarketIdSet(log types.Log) (*ILendingPoolAddressesProviderMarketIdSet, error) {
	event := new(ILendingPoolAddressesProviderMarketIdSet)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "MarketIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderPriceOracleUpdatedIterator is returned from FilterPriceOracleUpdated and is used to iterate over the raw logs and unpacked data for PriceOracleUpdated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderPriceOracleUpdatedIterator struct {
	Event *ILendingPoolAddressesProviderPriceOracleUpdated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderPriceOracleUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderPriceOracleUpdated)
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
		it.Event = new(ILendingPoolAddressesProviderPriceOracleUpdated)
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
func (it *ILendingPoolAddressesProviderPriceOracleUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderPriceOracleUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderPriceOracleUpdated represents a PriceOracleUpdated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderPriceOracleUpdated struct {
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPriceOracleUpdated is a free log retrieval operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterPriceOracleUpdated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderPriceOracleUpdatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderPriceOracleUpdatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "PriceOracleUpdated", logs: logs, sub: sub}, nil
}

// WatchPriceOracleUpdated is a free log subscription operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchPriceOracleUpdated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderPriceOracleUpdated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "PriceOracleUpdated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderPriceOracleUpdated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
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

// ParsePriceOracleUpdated is a log parse operation binding the contract event 0xefe8ab924ca486283a79dc604baa67add51afb82af1db8ac386ebbba643cdffd.
//
// Solidity: event PriceOracleUpdated(address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParsePriceOracleUpdated(log types.Log) (*ILendingPoolAddressesProviderPriceOracleUpdated, error) {
	event := new(ILendingPoolAddressesProviderPriceOracleUpdated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "PriceOracleUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ILendingPoolAddressesProviderProxyCreatedIterator is returned from FilterProxyCreated and is used to iterate over the raw logs and unpacked data for ProxyCreated events raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderProxyCreatedIterator struct {
	Event *ILendingPoolAddressesProviderProxyCreated // Event containing the contract specifics and raw log

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
func (it *ILendingPoolAddressesProviderProxyCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ILendingPoolAddressesProviderProxyCreated)
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
		it.Event = new(ILendingPoolAddressesProviderProxyCreated)
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
func (it *ILendingPoolAddressesProviderProxyCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ILendingPoolAddressesProviderProxyCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ILendingPoolAddressesProviderProxyCreated represents a ProxyCreated event raised by the ILendingPoolAddressesProvider contract.
type ILendingPoolAddressesProviderProxyCreated struct {
	Id         [32]byte
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterProxyCreated is a free log retrieval operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) FilterProxyCreated(opts *bind.FilterOpts, newAddress []common.Address) (*ILendingPoolAddressesProviderProxyCreatedIterator, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.FilterLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return &ILendingPoolAddressesProviderProxyCreatedIterator{contract: _ILendingPoolAddressesProvider.contract, event: "ProxyCreated", logs: logs, sub: sub}, nil
}

// WatchProxyCreated is a free log subscription operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) WatchProxyCreated(opts *bind.WatchOpts, sink chan<- *ILendingPoolAddressesProviderProxyCreated, newAddress []common.Address) (event.Subscription, error) {

	var newAddressRule []interface{}
	for _, newAddressItem := range newAddress {
		newAddressRule = append(newAddressRule, newAddressItem)
	}

	logs, sub, err := _ILendingPoolAddressesProvider.contract.WatchLogs(opts, "ProxyCreated", newAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ILendingPoolAddressesProviderProxyCreated)
				if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
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

// ParseProxyCreated is a log parse operation binding the contract event 0x1eb35cb4b5bbb23d152f3b4016a5a46c37a07ae930ed0956aba951e231142438.
//
// Solidity: event ProxyCreated(bytes32 id, address indexed newAddress)
func (_ILendingPoolAddressesProvider *ILendingPoolAddressesProviderFilterer) ParseProxyCreated(log types.Log) (*ILendingPoolAddressesProviderProxyCreated, error) {
	event := new(ILendingPoolAddressesProviderProxyCreated)
	if err := _ILendingPoolAddressesProvider.contract.UnpackLog(event, "ProxyCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IScaledBalanceTokenMetaData contains all meta data concerning the IScaledBalanceToken contract.
var IScaledBalanceTokenMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IScaledBalanceTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IScaledBalanceTokenMetaData.ABI instead.
var IScaledBalanceTokenABI = IScaledBalanceTokenMetaData.ABI

// IScaledBalanceToken is an auto generated Go binding around an Ethereum contract.
type IScaledBalanceToken struct {
	IScaledBalanceTokenCaller     // Read-only binding to the contract
	IScaledBalanceTokenTransactor // Write-only binding to the contract
	IScaledBalanceTokenFilterer   // Log filterer for contract events
}

// IScaledBalanceTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IScaledBalanceTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IScaledBalanceTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IScaledBalanceTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IScaledBalanceTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IScaledBalanceTokenSession struct {
	Contract     *IScaledBalanceToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// IScaledBalanceTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IScaledBalanceTokenCallerSession struct {
	Contract *IScaledBalanceTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// IScaledBalanceTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IScaledBalanceTokenTransactorSession struct {
	Contract     *IScaledBalanceTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// IScaledBalanceTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IScaledBalanceTokenRaw struct {
	Contract *IScaledBalanceToken // Generic contract binding to access the raw methods on
}

// IScaledBalanceTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IScaledBalanceTokenCallerRaw struct {
	Contract *IScaledBalanceTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IScaledBalanceTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IScaledBalanceTokenTransactorRaw struct {
	Contract *IScaledBalanceTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIScaledBalanceToken creates a new instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceToken(address common.Address, backend bind.ContractBackend) (*IScaledBalanceToken, error) {
	contract, err := bindIScaledBalanceToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceToken{IScaledBalanceTokenCaller: IScaledBalanceTokenCaller{contract: contract}, IScaledBalanceTokenTransactor: IScaledBalanceTokenTransactor{contract: contract}, IScaledBalanceTokenFilterer: IScaledBalanceTokenFilterer{contract: contract}}, nil
}

// NewIScaledBalanceTokenCaller creates a new read-only instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenCaller(address common.Address, caller bind.ContractCaller) (*IScaledBalanceTokenCaller, error) {
	contract, err := bindIScaledBalanceToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenCaller{contract: contract}, nil
}

// NewIScaledBalanceTokenTransactor creates a new write-only instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IScaledBalanceTokenTransactor, error) {
	contract, err := bindIScaledBalanceToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenTransactor{contract: contract}, nil
}

// NewIScaledBalanceTokenFilterer creates a new log filterer instance of IScaledBalanceToken, bound to a specific deployed contract.
func NewIScaledBalanceTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IScaledBalanceTokenFilterer, error) {
	contract, err := bindIScaledBalanceToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IScaledBalanceTokenFilterer{contract: contract}, nil
}

// bindIScaledBalanceToken binds a generic wrapper to an already deployed contract.
func bindIScaledBalanceToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IScaledBalanceTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScaledBalanceToken *IScaledBalanceTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.IScaledBalanceTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IScaledBalanceToken *IScaledBalanceTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IScaledBalanceToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IScaledBalanceToken *IScaledBalanceTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IScaledBalanceToken *IScaledBalanceTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IScaledBalanceToken.Contract.contract.Transact(opts, method, params...)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IScaledBalanceToken.Contract.GetScaledUserBalanceAndSupply(&_IScaledBalanceToken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IScaledBalanceToken.Contract.GetScaledUserBalanceAndSupply(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledBalanceOf(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledBalanceOf(&_IScaledBalanceToken.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IScaledBalanceToken.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenSession) ScaledTotalSupply() (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledTotalSupply(&_IScaledBalanceToken.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IScaledBalanceToken *IScaledBalanceTokenCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _IScaledBalanceToken.Contract.ScaledTotalSupply(&_IScaledBalanceToken.CallOpts)
}

// IStableDebtTokenMetaData contains all meta data concerning the IStableDebtToken contract.
var IStableDebtTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"currentBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"balanceIncrease\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"avgStableRate\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newTotalSupply\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getAverageStableRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getSupplyData\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupplyAndAvgRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTotalSupplyLastUpdated\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserLastUpdated\",\"outputs\":[{\"internalType\":\"uint40\",\"name\":\"\",\"type\":\"uint40\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getUserStableRate\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"rate\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"principalBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IStableDebtTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IStableDebtTokenMetaData.ABI instead.
var IStableDebtTokenABI = IStableDebtTokenMetaData.ABI

// IStableDebtToken is an auto generated Go binding around an Ethereum contract.
type IStableDebtToken struct {
	IStableDebtTokenCaller     // Read-only binding to the contract
	IStableDebtTokenTransactor // Write-only binding to the contract
	IStableDebtTokenFilterer   // Log filterer for contract events
}

// IStableDebtTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IStableDebtTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStableDebtTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IStableDebtTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStableDebtTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IStableDebtTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IStableDebtTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IStableDebtTokenSession struct {
	Contract     *IStableDebtToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IStableDebtTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IStableDebtTokenCallerSession struct {
	Contract *IStableDebtTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// IStableDebtTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IStableDebtTokenTransactorSession struct {
	Contract     *IStableDebtTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// IStableDebtTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IStableDebtTokenRaw struct {
	Contract *IStableDebtToken // Generic contract binding to access the raw methods on
}

// IStableDebtTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IStableDebtTokenCallerRaw struct {
	Contract *IStableDebtTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IStableDebtTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IStableDebtTokenTransactorRaw struct {
	Contract *IStableDebtTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIStableDebtToken creates a new instance of IStableDebtToken, bound to a specific deployed contract.
func NewIStableDebtToken(address common.Address, backend bind.ContractBackend) (*IStableDebtToken, error) {
	contract, err := bindIStableDebtToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IStableDebtToken{IStableDebtTokenCaller: IStableDebtTokenCaller{contract: contract}, IStableDebtTokenTransactor: IStableDebtTokenTransactor{contract: contract}, IStableDebtTokenFilterer: IStableDebtTokenFilterer{contract: contract}}, nil
}

// NewIStableDebtTokenCaller creates a new read-only instance of IStableDebtToken, bound to a specific deployed contract.
func NewIStableDebtTokenCaller(address common.Address, caller bind.ContractCaller) (*IStableDebtTokenCaller, error) {
	contract, err := bindIStableDebtToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IStableDebtTokenCaller{contract: contract}, nil
}

// NewIStableDebtTokenTransactor creates a new write-only instance of IStableDebtToken, bound to a specific deployed contract.
func NewIStableDebtTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IStableDebtTokenTransactor, error) {
	contract, err := bindIStableDebtToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IStableDebtTokenTransactor{contract: contract}, nil
}

// NewIStableDebtTokenFilterer creates a new log filterer instance of IStableDebtToken, bound to a specific deployed contract.
func NewIStableDebtTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IStableDebtTokenFilterer, error) {
	contract, err := bindIStableDebtToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IStableDebtTokenFilterer{contract: contract}, nil
}

// bindIStableDebtToken binds a generic wrapper to an already deployed contract.
func bindIStableDebtToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IStableDebtTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStableDebtToken *IStableDebtTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStableDebtToken.Contract.IStableDebtTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStableDebtToken *IStableDebtTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.IStableDebtTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStableDebtToken *IStableDebtTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.IStableDebtTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IStableDebtToken *IStableDebtTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IStableDebtToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IStableDebtToken *IStableDebtTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IStableDebtToken *IStableDebtTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.contract.Transact(opts, method, params...)
}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCaller) GetAverageStableRate(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getAverageStableRate")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenSession) GetAverageStableRate() (*big.Int, error) {
	return _IStableDebtToken.Contract.GetAverageStableRate(&_IStableDebtToken.CallOpts)
}

// GetAverageStableRate is a free data retrieval call binding the contract method 0x90f6fcf2.
//
// Solidity: function getAverageStableRate() view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetAverageStableRate() (*big.Int, error) {
	return _IStableDebtToken.Contract.GetAverageStableRate(&_IStableDebtToken.CallOpts)
}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_IStableDebtToken *IStableDebtTokenCaller) GetSupplyData(opts *bind.CallOpts) (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getSupplyData")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	out2 := *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	out3 := *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return out0, out1, out2, out3, err

}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_IStableDebtToken *IStableDebtTokenSession) GetSupplyData() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IStableDebtToken.Contract.GetSupplyData(&_IStableDebtToken.CallOpts)
}

// GetSupplyData is a free data retrieval call binding the contract method 0x79774338.
//
// Solidity: function getSupplyData() view returns(uint256, uint256, uint256, uint40)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetSupplyData() (*big.Int, *big.Int, *big.Int, *big.Int, error) {
	return _IStableDebtToken.Contract.GetSupplyData(&_IStableDebtToken.CallOpts)
}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_IStableDebtToken *IStableDebtTokenCaller) GetTotalSupplyAndAvgRate(opts *bind.CallOpts) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getTotalSupplyAndAvgRate")

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_IStableDebtToken *IStableDebtTokenSession) GetTotalSupplyAndAvgRate() (*big.Int, *big.Int, error) {
	return _IStableDebtToken.Contract.GetTotalSupplyAndAvgRate(&_IStableDebtToken.CallOpts)
}

// GetTotalSupplyAndAvgRate is a free data retrieval call binding the contract method 0xf731e9be.
//
// Solidity: function getTotalSupplyAndAvgRate() view returns(uint256, uint256)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetTotalSupplyAndAvgRate() (*big.Int, *big.Int, error) {
	return _IStableDebtToken.Contract.GetTotalSupplyAndAvgRate(&_IStableDebtToken.CallOpts)
}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenCaller) GetTotalSupplyLastUpdated(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getTotalSupplyLastUpdated")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenSession) GetTotalSupplyLastUpdated() (*big.Int, error) {
	return _IStableDebtToken.Contract.GetTotalSupplyLastUpdated(&_IStableDebtToken.CallOpts)
}

// GetTotalSupplyLastUpdated is a free data retrieval call binding the contract method 0xe7484890.
//
// Solidity: function getTotalSupplyLastUpdated() view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetTotalSupplyLastUpdated() (*big.Int, error) {
	return _IStableDebtToken.Contract.GetTotalSupplyLastUpdated(&_IStableDebtToken.CallOpts)
}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenCaller) GetUserLastUpdated(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getUserLastUpdated", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenSession) GetUserLastUpdated(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.GetUserLastUpdated(&_IStableDebtToken.CallOpts, user)
}

// GetUserLastUpdated is a free data retrieval call binding the contract method 0x79ce6b8c.
//
// Solidity: function getUserLastUpdated(address user) view returns(uint40)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetUserLastUpdated(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.GetUserLastUpdated(&_IStableDebtToken.CallOpts, user)
}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCaller) GetUserStableRate(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "getUserStableRate", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenSession) GetUserStableRate(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.GetUserStableRate(&_IStableDebtToken.CallOpts, user)
}

// GetUserStableRate is a free data retrieval call binding the contract method 0xe78c9b3b.
//
// Solidity: function getUserStableRate(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCallerSession) GetUserStableRate(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.GetUserStableRate(&_IStableDebtToken.CallOpts, user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCaller) PrincipalBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IStableDebtToken.contract.Call(opts, &out, "principalBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenSession) PrincipalBalanceOf(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.PrincipalBalanceOf(&_IStableDebtToken.CallOpts, user)
}

// PrincipalBalanceOf is a free data retrieval call binding the contract method 0xc634dfaa.
//
// Solidity: function principalBalanceOf(address user) view returns(uint256)
func (_IStableDebtToken *IStableDebtTokenCallerSession) PrincipalBalanceOf(user common.Address) (*big.Int, error) {
	return _IStableDebtToken.Contract.PrincipalBalanceOf(&_IStableDebtToken.CallOpts, user)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_IStableDebtToken *IStableDebtTokenTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.contract.Transact(opts, "burn", user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_IStableDebtToken *IStableDebtTokenSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.Burn(&_IStableDebtToken.TransactOpts, user, amount)
}

// Burn is a paid mutator transaction binding the contract method 0x9dc29fac.
//
// Solidity: function burn(address user, uint256 amount) returns()
func (_IStableDebtToken *IStableDebtTokenTransactorSession) Burn(user common.Address, amount *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.Burn(&_IStableDebtToken.TransactOpts, user, amount)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_IStableDebtToken *IStableDebtTokenTransactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.contract.Transact(opts, "mint", user, onBehalfOf, amount, rate)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_IStableDebtToken *IStableDebtTokenSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.Mint(&_IStableDebtToken.TransactOpts, user, onBehalfOf, amount, rate)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 rate) returns(bool)
func (_IStableDebtToken *IStableDebtTokenTransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, rate *big.Int) (*types.Transaction, error) {
	return _IStableDebtToken.Contract.Mint(&_IStableDebtToken.TransactOpts, user, onBehalfOf, amount, rate)
}

// IStableDebtTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IStableDebtToken contract.
type IStableDebtTokenBurnIterator struct {
	Event *IStableDebtTokenBurn // Event containing the contract specifics and raw log

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
func (it *IStableDebtTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStableDebtTokenBurn)
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
		it.Event = new(IStableDebtTokenBurn)
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
func (it *IStableDebtTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStableDebtTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStableDebtTokenBurn represents a Burn event raised by the IStableDebtToken contract.
type IStableDebtTokenBurn struct {
	User            common.Address
	Amount          *big.Int
	CurrentBalance  *big.Int
	BalanceIncrease *big.Int
	AvgStableRate   *big.Int
	NewTotalSupply  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*IStableDebtTokenBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IStableDebtToken.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &IStableDebtTokenBurnIterator{contract: _IStableDebtToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IStableDebtTokenBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IStableDebtToken.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStableDebtTokenBurn)
				if err := _IStableDebtToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x44bd20a79e993bdcc7cbedf54a3b4d19fb78490124b6b90d04fe3242eea579e8.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) ParseBurn(log types.Log) (*IStableDebtTokenBurn, error) {
	event := new(IStableDebtTokenBurn)
	if err := _IStableDebtToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IStableDebtTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IStableDebtToken contract.
type IStableDebtTokenMintIterator struct {
	Event *IStableDebtTokenMint // Event containing the contract specifics and raw log

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
func (it *IStableDebtTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IStableDebtTokenMint)
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
		it.Event = new(IStableDebtTokenMint)
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
func (it *IStableDebtTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IStableDebtTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IStableDebtTokenMint represents a Mint event raised by the IStableDebtToken contract.
type IStableDebtTokenMint struct {
	User            common.Address
	OnBehalfOf      common.Address
	Amount          *big.Int
	CurrentBalance  *big.Int
	BalanceIncrease *big.Int
	NewRate         *big.Int
	AvgStableRate   *big.Int
	NewTotalSupply  *big.Int
	Raw             types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) FilterMint(opts *bind.FilterOpts, user []common.Address, onBehalfOf []common.Address) (*IStableDebtTokenMintIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IStableDebtToken.contract.FilterLogs(opts, "Mint", userRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &IStableDebtTokenMintIterator{contract: _IStableDebtToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IStableDebtTokenMint, user []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IStableDebtToken.contract.WatchLogs(opts, "Mint", userRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IStableDebtTokenMint)
				if err := _IStableDebtToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0xc16f4e4ca34d790de4c656c72fd015c667d688f20be64eea360618545c4c530f.
//
// Solidity: event Mint(address indexed user, address indexed onBehalfOf, uint256 amount, uint256 currentBalance, uint256 balanceIncrease, uint256 newRate, uint256 avgStableRate, uint256 newTotalSupply)
func (_IStableDebtToken *IStableDebtTokenFilterer) ParseMint(log types.Log) (*IStableDebtTokenMint, error) {
	event := new(IStableDebtTokenMint)
	if err := _IStableDebtToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVariableDebtTokenMetaData contains all meta data concerning the IVariableDebtToken contract.
var IVariableDebtTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Burn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"getScaledUserBalanceAndSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"onBehalfOf\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"}],\"name\":\"scaledBalanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"scaledTotalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
}

// IVariableDebtTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use IVariableDebtTokenMetaData.ABI instead.
var IVariableDebtTokenABI = IVariableDebtTokenMetaData.ABI

// IVariableDebtToken is an auto generated Go binding around an Ethereum contract.
type IVariableDebtToken struct {
	IVariableDebtTokenCaller     // Read-only binding to the contract
	IVariableDebtTokenTransactor // Write-only binding to the contract
	IVariableDebtTokenFilterer   // Log filterer for contract events
}

// IVariableDebtTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type IVariableDebtTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVariableDebtTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IVariableDebtTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVariableDebtTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IVariableDebtTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IVariableDebtTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IVariableDebtTokenSession struct {
	Contract     *IVariableDebtToken // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IVariableDebtTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IVariableDebtTokenCallerSession struct {
	Contract *IVariableDebtTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// IVariableDebtTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IVariableDebtTokenTransactorSession struct {
	Contract     *IVariableDebtTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// IVariableDebtTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type IVariableDebtTokenRaw struct {
	Contract *IVariableDebtToken // Generic contract binding to access the raw methods on
}

// IVariableDebtTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IVariableDebtTokenCallerRaw struct {
	Contract *IVariableDebtTokenCaller // Generic read-only contract binding to access the raw methods on
}

// IVariableDebtTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IVariableDebtTokenTransactorRaw struct {
	Contract *IVariableDebtTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIVariableDebtToken creates a new instance of IVariableDebtToken, bound to a specific deployed contract.
func NewIVariableDebtToken(address common.Address, backend bind.ContractBackend) (*IVariableDebtToken, error) {
	contract, err := bindIVariableDebtToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtToken{IVariableDebtTokenCaller: IVariableDebtTokenCaller{contract: contract}, IVariableDebtTokenTransactor: IVariableDebtTokenTransactor{contract: contract}, IVariableDebtTokenFilterer: IVariableDebtTokenFilterer{contract: contract}}, nil
}

// NewIVariableDebtTokenCaller creates a new read-only instance of IVariableDebtToken, bound to a specific deployed contract.
func NewIVariableDebtTokenCaller(address common.Address, caller bind.ContractCaller) (*IVariableDebtTokenCaller, error) {
	contract, err := bindIVariableDebtToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtTokenCaller{contract: contract}, nil
}

// NewIVariableDebtTokenTransactor creates a new write-only instance of IVariableDebtToken, bound to a specific deployed contract.
func NewIVariableDebtTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*IVariableDebtTokenTransactor, error) {
	contract, err := bindIVariableDebtToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtTokenTransactor{contract: contract}, nil
}

// NewIVariableDebtTokenFilterer creates a new log filterer instance of IVariableDebtToken, bound to a specific deployed contract.
func NewIVariableDebtTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*IVariableDebtTokenFilterer, error) {
	contract, err := bindIVariableDebtToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtTokenFilterer{contract: contract}, nil
}

// bindIVariableDebtToken binds a generic wrapper to an already deployed contract.
func bindIVariableDebtToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IVariableDebtTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVariableDebtToken *IVariableDebtTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVariableDebtToken.Contract.IVariableDebtTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVariableDebtToken *IVariableDebtTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.IVariableDebtTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVariableDebtToken *IVariableDebtTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.IVariableDebtTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_IVariableDebtToken *IVariableDebtTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _IVariableDebtToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_IVariableDebtToken *IVariableDebtTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_IVariableDebtToken *IVariableDebtTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.contract.Transact(opts, method, params...)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IVariableDebtToken *IVariableDebtTokenCaller) GetScaledUserBalanceAndSupply(opts *bind.CallOpts, user common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _IVariableDebtToken.contract.Call(opts, &out, "getScaledUserBalanceAndSupply", user)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IVariableDebtToken *IVariableDebtTokenSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IVariableDebtToken.Contract.GetScaledUserBalanceAndSupply(&_IVariableDebtToken.CallOpts, user)
}

// GetScaledUserBalanceAndSupply is a free data retrieval call binding the contract method 0x0afbcdc9.
//
// Solidity: function getScaledUserBalanceAndSupply(address user) view returns(uint256, uint256)
func (_IVariableDebtToken *IVariableDebtTokenCallerSession) GetScaledUserBalanceAndSupply(user common.Address) (*big.Int, *big.Int, error) {
	return _IVariableDebtToken.Contract.GetScaledUserBalanceAndSupply(&_IVariableDebtToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenCaller) ScaledBalanceOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _IVariableDebtToken.contract.Call(opts, &out, "scaledBalanceOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IVariableDebtToken.Contract.ScaledBalanceOf(&_IVariableDebtToken.CallOpts, user)
}

// ScaledBalanceOf is a free data retrieval call binding the contract method 0x1da24f3e.
//
// Solidity: function scaledBalanceOf(address user) view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenCallerSession) ScaledBalanceOf(user common.Address) (*big.Int, error) {
	return _IVariableDebtToken.Contract.ScaledBalanceOf(&_IVariableDebtToken.CallOpts, user)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenCaller) ScaledTotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _IVariableDebtToken.contract.Call(opts, &out, "scaledTotalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenSession) ScaledTotalSupply() (*big.Int, error) {
	return _IVariableDebtToken.Contract.ScaledTotalSupply(&_IVariableDebtToken.CallOpts)
}

// ScaledTotalSupply is a free data retrieval call binding the contract method 0xb1bf962d.
//
// Solidity: function scaledTotalSupply() view returns(uint256)
func (_IVariableDebtToken *IVariableDebtTokenCallerSession) ScaledTotalSupply() (*big.Int, error) {
	return _IVariableDebtToken.Contract.ScaledTotalSupply(&_IVariableDebtToken.CallOpts)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_IVariableDebtToken *IVariableDebtTokenTransactor) Burn(opts *bind.TransactOpts, user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.contract.Transact(opts, "burn", user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_IVariableDebtToken *IVariableDebtTokenSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.Burn(&_IVariableDebtToken.TransactOpts, user, amount, index)
}

// Burn is a paid mutator transaction binding the contract method 0xf5298aca.
//
// Solidity: function burn(address user, uint256 amount, uint256 index) returns()
func (_IVariableDebtToken *IVariableDebtTokenTransactorSession) Burn(user common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.Burn(&_IVariableDebtToken.TransactOpts, user, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_IVariableDebtToken *IVariableDebtTokenTransactor) Mint(opts *bind.TransactOpts, user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.contract.Transact(opts, "mint", user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_IVariableDebtToken *IVariableDebtTokenSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.Mint(&_IVariableDebtToken.TransactOpts, user, onBehalfOf, amount, index)
}

// Mint is a paid mutator transaction binding the contract method 0xb3f1c93d.
//
// Solidity: function mint(address user, address onBehalfOf, uint256 amount, uint256 index) returns(bool)
func (_IVariableDebtToken *IVariableDebtTokenTransactorSession) Mint(user common.Address, onBehalfOf common.Address, amount *big.Int, index *big.Int) (*types.Transaction, error) {
	return _IVariableDebtToken.Contract.Mint(&_IVariableDebtToken.TransactOpts, user, onBehalfOf, amount, index)
}

// IVariableDebtTokenBurnIterator is returned from FilterBurn and is used to iterate over the raw logs and unpacked data for Burn events raised by the IVariableDebtToken contract.
type IVariableDebtTokenBurnIterator struct {
	Event *IVariableDebtTokenBurn // Event containing the contract specifics and raw log

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
func (it *IVariableDebtTokenBurnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVariableDebtTokenBurn)
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
		it.Event = new(IVariableDebtTokenBurn)
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
func (it *IVariableDebtTokenBurnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVariableDebtTokenBurnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVariableDebtTokenBurn represents a Burn event raised by the IVariableDebtToken contract.
type IVariableDebtTokenBurn struct {
	User   common.Address
	Amount *big.Int
	Index  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBurn is a free log retrieval operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) FilterBurn(opts *bind.FilterOpts, user []common.Address) (*IVariableDebtTokenBurnIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IVariableDebtToken.contract.FilterLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtTokenBurnIterator{contract: _IVariableDebtToken.contract, event: "Burn", logs: logs, sub: sub}, nil
}

// WatchBurn is a free log subscription operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) WatchBurn(opts *bind.WatchOpts, sink chan<- *IVariableDebtTokenBurn, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _IVariableDebtToken.contract.WatchLogs(opts, "Burn", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVariableDebtTokenBurn)
				if err := _IVariableDebtToken.contract.UnpackLog(event, "Burn", log); err != nil {
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

// ParseBurn is a log parse operation binding the contract event 0x49995e5dd6158cf69ad3e9777c46755a1a826a446c6416992167462dad033b2a.
//
// Solidity: event Burn(address indexed user, uint256 amount, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) ParseBurn(log types.Log) (*IVariableDebtTokenBurn, error) {
	event := new(IVariableDebtTokenBurn)
	if err := _IVariableDebtToken.contract.UnpackLog(event, "Burn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// IVariableDebtTokenMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the IVariableDebtToken contract.
type IVariableDebtTokenMintIterator struct {
	Event *IVariableDebtTokenMint // Event containing the contract specifics and raw log

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
func (it *IVariableDebtTokenMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(IVariableDebtTokenMint)
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
		it.Event = new(IVariableDebtTokenMint)
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
func (it *IVariableDebtTokenMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *IVariableDebtTokenMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// IVariableDebtTokenMint represents a Mint event raised by the IVariableDebtToken contract.
type IVariableDebtTokenMint struct {
	From       common.Address
	OnBehalfOf common.Address
	Value      *big.Int
	Index      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) FilterMint(opts *bind.FilterOpts, from []common.Address, onBehalfOf []common.Address) (*IVariableDebtTokenMintIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IVariableDebtToken.contract.FilterLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return &IVariableDebtTokenMintIterator{contract: _IVariableDebtToken.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *IVariableDebtTokenMint, from []common.Address, onBehalfOf []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var onBehalfOfRule []interface{}
	for _, onBehalfOfItem := range onBehalfOf {
		onBehalfOfRule = append(onBehalfOfRule, onBehalfOfItem)
	}

	logs, sub, err := _IVariableDebtToken.contract.WatchLogs(opts, "Mint", fromRule, onBehalfOfRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(IVariableDebtTokenMint)
				if err := _IVariableDebtToken.contract.UnpackLog(event, "Mint", log); err != nil {
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

// ParseMint is a log parse operation binding the contract event 0x2f00e3cdd69a77be7ed215ec7b2a36784dd158f921fca79ac29deffa353fe6ee.
//
// Solidity: event Mint(address indexed from, address indexed onBehalfOf, uint256 value, uint256 index)
func (_IVariableDebtToken *IVariableDebtTokenFilterer) ParseMint(log types.Log) (*IVariableDebtTokenMint, error) {
	event := new(IVariableDebtTokenMint)
	if err := _IVariableDebtToken.contract.UnpackLog(event, "Mint", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ReserveConfigurationMetaData contains all meta data concerning the ReserveConfiguration contract.
var ReserveConfigurationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea2646970667358221220d3c4ca4b53a5aa8577bc5cec1bafa7216cc6ff37f3255c779db4e1e1fad443cd64736f6c634300060c0033",
}

// ReserveConfigurationABI is the input ABI used to generate the binding from.
// Deprecated: Use ReserveConfigurationMetaData.ABI instead.
var ReserveConfigurationABI = ReserveConfigurationMetaData.ABI

// ReserveConfigurationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use ReserveConfigurationMetaData.Bin instead.
var ReserveConfigurationBin = ReserveConfigurationMetaData.Bin

// DeployReserveConfiguration deploys a new Ethereum contract, binding an instance of ReserveConfiguration to it.
func DeployReserveConfiguration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ReserveConfiguration, error) {
	parsed, err := ReserveConfigurationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(ReserveConfigurationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ReserveConfiguration{ReserveConfigurationCaller: ReserveConfigurationCaller{contract: contract}, ReserveConfigurationTransactor: ReserveConfigurationTransactor{contract: contract}, ReserveConfigurationFilterer: ReserveConfigurationFilterer{contract: contract}}, nil
}

// ReserveConfiguration is an auto generated Go binding around an Ethereum contract.
type ReserveConfiguration struct {
	ReserveConfigurationCaller     // Read-only binding to the contract
	ReserveConfigurationTransactor // Write-only binding to the contract
	ReserveConfigurationFilterer   // Log filterer for contract events
}

// ReserveConfigurationCaller is an auto generated read-only Go binding around an Ethereum contract.
type ReserveConfigurationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveConfigurationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ReserveConfigurationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveConfigurationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ReserveConfigurationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ReserveConfigurationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ReserveConfigurationSession struct {
	Contract     *ReserveConfiguration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// ReserveConfigurationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ReserveConfigurationCallerSession struct {
	Contract *ReserveConfigurationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// ReserveConfigurationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ReserveConfigurationTransactorSession struct {
	Contract     *ReserveConfigurationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// ReserveConfigurationRaw is an auto generated low-level Go binding around an Ethereum contract.
type ReserveConfigurationRaw struct {
	Contract *ReserveConfiguration // Generic contract binding to access the raw methods on
}

// ReserveConfigurationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ReserveConfigurationCallerRaw struct {
	Contract *ReserveConfigurationCaller // Generic read-only contract binding to access the raw methods on
}

// ReserveConfigurationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ReserveConfigurationTransactorRaw struct {
	Contract *ReserveConfigurationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewReserveConfiguration creates a new instance of ReserveConfiguration, bound to a specific deployed contract.
func NewReserveConfiguration(address common.Address, backend bind.ContractBackend) (*ReserveConfiguration, error) {
	contract, err := bindReserveConfiguration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ReserveConfiguration{ReserveConfigurationCaller: ReserveConfigurationCaller{contract: contract}, ReserveConfigurationTransactor: ReserveConfigurationTransactor{contract: contract}, ReserveConfigurationFilterer: ReserveConfigurationFilterer{contract: contract}}, nil
}

// NewReserveConfigurationCaller creates a new read-only instance of ReserveConfiguration, bound to a specific deployed contract.
func NewReserveConfigurationCaller(address common.Address, caller bind.ContractCaller) (*ReserveConfigurationCaller, error) {
	contract, err := bindReserveConfiguration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveConfigurationCaller{contract: contract}, nil
}

// NewReserveConfigurationTransactor creates a new write-only instance of ReserveConfiguration, bound to a specific deployed contract.
func NewReserveConfigurationTransactor(address common.Address, transactor bind.ContractTransactor) (*ReserveConfigurationTransactor, error) {
	contract, err := bindReserveConfiguration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ReserveConfigurationTransactor{contract: contract}, nil
}

// NewReserveConfigurationFilterer creates a new log filterer instance of ReserveConfiguration, bound to a specific deployed contract.
func NewReserveConfigurationFilterer(address common.Address, filterer bind.ContractFilterer) (*ReserveConfigurationFilterer, error) {
	contract, err := bindReserveConfiguration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ReserveConfigurationFilterer{contract: contract}, nil
}

// bindReserveConfiguration binds a generic wrapper to an already deployed contract.
func bindReserveConfiguration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ReserveConfigurationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveConfiguration *ReserveConfigurationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveConfiguration.Contract.ReserveConfigurationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveConfiguration *ReserveConfigurationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveConfiguration.Contract.ReserveConfigurationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveConfiguration *ReserveConfigurationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveConfiguration.Contract.ReserveConfigurationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ReserveConfiguration *ReserveConfigurationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ReserveConfiguration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ReserveConfiguration *ReserveConfigurationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ReserveConfiguration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ReserveConfiguration *ReserveConfigurationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ReserveConfiguration.Contract.contract.Transact(opts, method, params...)
}

// UserConfigurationMetaData contains all meta data concerning the UserConfiguration contract.
var UserConfigurationMetaData = &bind.MetaData{
	ABI: "[]",
	Bin: "0x60566023600b82828239805160001a607314601657fe5b30600052607381538281f3fe73000000000000000000000000000000000000000030146080604052600080fdfea26469706673582212207a644894502bc22d7eed8f90880025c284b6fb12cba8845a12c35cf52d7cb2c964736f6c634300060c0033",
}

// UserConfigurationABI is the input ABI used to generate the binding from.
// Deprecated: Use UserConfigurationMetaData.ABI instead.
var UserConfigurationABI = UserConfigurationMetaData.ABI

// UserConfigurationBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use UserConfigurationMetaData.Bin instead.
var UserConfigurationBin = UserConfigurationMetaData.Bin

// DeployUserConfiguration deploys a new Ethereum contract, binding an instance of UserConfiguration to it.
func DeployUserConfiguration(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UserConfiguration, error) {
	parsed, err := UserConfigurationMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(UserConfigurationBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UserConfiguration{UserConfigurationCaller: UserConfigurationCaller{contract: contract}, UserConfigurationTransactor: UserConfigurationTransactor{contract: contract}, UserConfigurationFilterer: UserConfigurationFilterer{contract: contract}}, nil
}

// UserConfiguration is an auto generated Go binding around an Ethereum contract.
type UserConfiguration struct {
	UserConfigurationCaller     // Read-only binding to the contract
	UserConfigurationTransactor // Write-only binding to the contract
	UserConfigurationFilterer   // Log filterer for contract events
}

// UserConfigurationCaller is an auto generated read-only Go binding around an Ethereum contract.
type UserConfigurationCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserConfigurationTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UserConfigurationTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserConfigurationFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UserConfigurationFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UserConfigurationSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UserConfigurationSession struct {
	Contract     *UserConfiguration // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// UserConfigurationCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UserConfigurationCallerSession struct {
	Contract *UserConfigurationCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// UserConfigurationTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UserConfigurationTransactorSession struct {
	Contract     *UserConfigurationTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// UserConfigurationRaw is an auto generated low-level Go binding around an Ethereum contract.
type UserConfigurationRaw struct {
	Contract *UserConfiguration // Generic contract binding to access the raw methods on
}

// UserConfigurationCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UserConfigurationCallerRaw struct {
	Contract *UserConfigurationCaller // Generic read-only contract binding to access the raw methods on
}

// UserConfigurationTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UserConfigurationTransactorRaw struct {
	Contract *UserConfigurationTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUserConfiguration creates a new instance of UserConfiguration, bound to a specific deployed contract.
func NewUserConfiguration(address common.Address, backend bind.ContractBackend) (*UserConfiguration, error) {
	contract, err := bindUserConfiguration(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UserConfiguration{UserConfigurationCaller: UserConfigurationCaller{contract: contract}, UserConfigurationTransactor: UserConfigurationTransactor{contract: contract}, UserConfigurationFilterer: UserConfigurationFilterer{contract: contract}}, nil
}

// NewUserConfigurationCaller creates a new read-only instance of UserConfiguration, bound to a specific deployed contract.
func NewUserConfigurationCaller(address common.Address, caller bind.ContractCaller) (*UserConfigurationCaller, error) {
	contract, err := bindUserConfiguration(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UserConfigurationCaller{contract: contract}, nil
}

// NewUserConfigurationTransactor creates a new write-only instance of UserConfiguration, bound to a specific deployed contract.
func NewUserConfigurationTransactor(address common.Address, transactor bind.ContractTransactor) (*UserConfigurationTransactor, error) {
	contract, err := bindUserConfiguration(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UserConfigurationTransactor{contract: contract}, nil
}

// NewUserConfigurationFilterer creates a new log filterer instance of UserConfiguration, bound to a specific deployed contract.
func NewUserConfigurationFilterer(address common.Address, filterer bind.ContractFilterer) (*UserConfigurationFilterer, error) {
	contract, err := bindUserConfiguration(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UserConfigurationFilterer{contract: contract}, nil
}

// bindUserConfiguration binds a generic wrapper to an already deployed contract.
func bindUserConfiguration(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UserConfigurationABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserConfiguration *UserConfigurationRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserConfiguration.Contract.UserConfigurationCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserConfiguration *UserConfigurationRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserConfiguration.Contract.UserConfigurationTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserConfiguration *UserConfigurationRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserConfiguration.Contract.UserConfigurationTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UserConfiguration *UserConfigurationCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UserConfiguration.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UserConfiguration *UserConfigurationTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UserConfiguration.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UserConfiguration *UserConfigurationTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UserConfiguration.Contract.contract.Transact(opts, method, params...)
}
