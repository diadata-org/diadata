// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptopunk

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

// CryptoPunksMarketABI is the input ABI used to generate the binding from.
const CryptoPunksMarketABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punksOfferedForSale\",\"outputs\":[{\"name\":\"isForSale\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"minValue\",\"type\":\"uint256\"},{\"name\":\"onlySellTo\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"enterBidForPunk\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minPrice\",\"type\":\"uint256\"}],\"name\":\"acceptBidForPunk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"addresses\",\"type\":\"address[]\"},{\"name\":\"indices\",\"type\":\"uint256[]\"}],\"name\":\"setInitialOwners\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdraw\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"imageHash\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nextPunkIndexToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkIndexToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"standard\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"punkBids\",\"outputs\":[{\"name\":\"hasBid\",\"type\":\"bool\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"bidder\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"allInitialOwnersAssigned\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"allPunksAssigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"buyPunk\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"transferPunk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"withdrawBidForPunk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"setInitialOwner\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"},{\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"offerPunkForSaleToAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"punksRemainingToAssign\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"name\":\"minSalePriceInWei\",\"type\":\"uint256\"}],\"name\":\"offerPunkForSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"getPunk\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"}],\"name\":\"pendingWithdrawals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"punkNoLongerForSale\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"Assign\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"minValue\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidEntered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"}],\"name\":\"PunkBidWithdrawn\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"fromAddress\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"toAddress\",\"type\":\"address\"}],\"name\":\"PunkBought\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"punkIndex\",\"type\":\"uint256\"}],\"name\":\"PunkNoLongerForSale\",\"type\":\"event\"}]"

// CryptoPunksMarketFuncSigs maps the 4-byte function signature to its string representation.
var CryptoPunksMarketFuncSigs = map[string]string{
	"23165b75": "acceptBidForPunk(uint256,uint256)",
	"7ecedac9": "allInitialOwnersAssigned()",
	"8126c38a": "allPunksAssigned()",
	"70a08231": "balanceOf(address)",
	"8264fe98": "buyPunk(uint256)",
	"313ce567": "decimals()",
	"091dbfd2": "enterBidForPunk(uint256)",
	"c81d1d5b": "getPunk(uint256)",
	"51605d80": "imageHash()",
	"06fdde03": "name()",
	"52f29a25": "nextPunkIndexToAssign()",
	"c44193c3": "offerPunkForSale(uint256,uint256)",
	"bf31196f": "offerPunkForSaleToAddress(uint256,uint256,address)",
	"f3f43703": "pendingWithdrawals(address)",
	"6e743fa9": "punkBids(uint256)",
	"58178168": "punkIndexToAddress(uint256)",
	"f6eeff1e": "punkNoLongerForSale(uint256)",
	"088f11f3": "punksOfferedForSale(uint256)",
	"c0d6ce63": "punksRemainingToAssign()",
	"a75a9049": "setInitialOwner(address,uint256)",
	"39c5dde6": "setInitialOwners(address[],uint256[])",
	"5a3b7e42": "standard()",
	"95d89b41": "symbol()",
	"18160ddd": "totalSupply()",
	"8b72a2ec": "transferPunk(address,uint256)",
	"3ccfd60b": "withdraw()",
	"979bc638": "withdrawBidForPunk(uint256)",
}

// CryptoPunksMarketBin is the compiled bytecode used for deploying new contracts.
var CryptoPunksMarketBin = "0x60e0604090815260808190527f616333396166343739333131396565343662626666333531643863623662356660a09081527f323364613630323232313236616464343236386532363131393961323932316260c05262000064916000919062000176565b5060408051808201909152600b8082527f43727970746f50756e6b730000000000000000000000000000000000000000006020909201918252620000ab9160029162000176565b5060006007556008805460ff1916905560018054600160a060020a03191633179055612710600681905560095560408051808201909152600b8082527f43525950544f50554e4b5300000000000000000000000000000000000000000060209092019182526200011e9160039162000176565b506040805180820190915260018082527f43000000000000000000000000000000000000000000000000000000000000006020909201918252620001659160049162000176565b506005805460ff191690556200021b565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10620001b957805160ff1916838001178555620001e9565b82800160010185558215620001e9579182015b82811115620001e9578251825591602001919060010190620001cc565b50620001f7929150620001fb565b5090565b6200021891905b80821115620001f7576000815560010162000202565b90565b611896806200022b6000396000f30060806040526004361061015e5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166306fdde038114610163578063088f11f3146101ed578063091dbfd21461023e57806318160ddd1461024b57806323165b7514610272578063313ce5671461028d57806339c5dde6146102b85780633ccfd60b1461034657806351605d801461035b57806352f29a251461037057806358178168146103855780635a3b7e42146103b95780636e743fa9146103ce57806370a08231146104165780637ecedac9146104375780638126c38a1461044c5780638264fe98146104755780638b72a2ec1461048057806395d89b41146104a4578063979bc638146104b9578063a75a9049146104d1578063bf31196f146104f5578063c0d6ce631461051c578063c44193c314610531578063c81d1d5b1461054c578063f3f4370314610564578063f6eeff1e14610585575b600080fd5b34801561016f57600080fd5b5061017861059d565b6040805160208082528351818301528351919283929083019185019080838360005b838110156101b257818101518382015260200161019a565b50505050905090810190601f1680156101df5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b3480156101f957600080fd5b5061020560043561062b565b6040805195151586526020860194909452600160a060020a03928316858501526060850191909152166080830152519081900360a00190f35b610249600435610669565b005b34801561025757600080fd5b506102606107e1565b60408051918252519081900360200190f35b34801561027e57600080fd5b506102496004356024356107e7565b34801561029957600080fd5b506102a2610b3d565b6040805160ff9092168252519081900360200190f35b3480156102c457600080fd5b506040805160206004803580820135838102808601850190965280855261024995369593946024949385019291829185019084908082843750506040805187358901803560208181028481018201909552818452989b9a998901989297509082019550935083925085019084908082843750949750610b469650505050505050565b34801561035257600080fd5b50610249610bb7565b34801561036757600080fd5b50610178610c11565b34801561037c57600080fd5b50610260610c6c565b34801561039157600080fd5b5061039d600435610c72565b60408051600160a060020a039092168252519081900360200190f35b3480156103c557600080fd5b50610178610c8d565b3480156103da57600080fd5b506103e6600435610ce5565b6040805194151585526020850193909352600160a060020a03909116838301526060830152519081900360800190f35b34801561042257600080fd5b50610260600160a060020a0360043516610d1c565b34801561044357600080fd5b50610249610d2e565b34801561045857600080fd5b50610461610d54565b604080519115158252519081900360200190f35b610249600435610d5d565b34801561048c57600080fd5b50610249600160a060020a0360043516602435610fb5565b3480156104b057600080fd5b50610178611192565b3480156104c557600080fd5b506102496004356111ed565b3480156104dd57600080fd5b50610249600160a060020a036004351660243561135a565b34801561050157600080fd5b50610249600435602435600160a060020a036044351661147e565b34801561052857600080fd5b5061026061158c565b34801561053d57600080fd5b50610249600435602435611592565b34801561055857600080fd5b5061024960043561169c565b34801561057057600080fd5b50610260600160a060020a0360043516611760565b34801561059157600080fd5b50610249600435611772565b6003805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106235780601f106105f857610100808354040283529160200191610623565b820191906000526020600020905b81548152906001019060200180831161060657829003601f168201915b505050505081565b600c602052600090815260409020805460018201546002830154600384015460049094015460ff909316939192600160a060020a0391821692911685565b6000612710821061067957600080fd5b60085460ff16151561068a57600080fd5b6000828152600a6020526040902054600160a060020a031615156106ad57600080fd5b6000828152600a6020526040902054600160a060020a03163314156106d157600080fd5b3415156106dd57600080fd5b506000818152600d60205260409020600381015434116106fc57600080fd5b6000816003015411156107325760038101546002820154600160a060020a03166000908152600e60205260409020805490910190555b604080516080810182526001808252602080830186815233848601818152346060870181815260008b8152600d87528990209751885460ff191690151517885593519587019590955551600286018054600160a060020a031916600160a060020a03909216919091179055905160039094019390935583519182529251919285927f5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a9281900390910190a35050565b60065481565b6000808061271085106107f957600080fd5b60085460ff16151561080a57600080fd5b6000858152600a6020526040902054600160a060020a0316331461082d57600080fd5b6000858152600d602052604090206003810154339450909250151561085157600080fd5b838260030154101561086257600080fd5b6002820180546000878152600a602090815260408083208054600160a060020a031916600160a060020a03958616179055878416808452600b83528184208054600019019055855485168452928190208054600190810190915594548151958652905193169391927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef9281900390910190a360a0604051908101604052806000151581526020018681526020018360020160009054906101000a9004600160a060020a0316600160a060020a03168152602001600081526020016000600160a060020a0316815250600c600087815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020160006101000a815481600160a060020a030219169083600160a060020a031602179055506060820151816003015560808201518160040160006101000a815481600160a060020a030219169083600160a060020a03160217905550905050816003015490506080604051908101604052806000151581526020018681526020016000600160a060020a031681526020016000815250600d600087815260200190815260200160002060008201518160000160006101000a81548160ff0219169083151502179055506020820151816001015560408201518160020160006101000a815481600160a060020a030219169083600160a060020a031602179055506060820151816003015590505080600e600085600160a060020a0316600160a060020a03168152602001908152602001600020600082825401925050819055508160020160009054906101000a9004600160a060020a0316600160a060020a031683600160a060020a0316867f58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e385600301546040518082815260200191505060405180910390a45050505050565b60055460ff1681565b6001546000908190600160a060020a03163314610b6257600080fd5b5050815160005b81811015610bb157610ba98482815181101515610b8257fe5b906020019060200201518483815181101515610b9a57fe5b9060200190602002015161135a565b600101610b69565b50505050565b60085460009060ff161515610bcb57600080fd5b50336000818152600e6020526040808220805490839055905190929183156108fc02918491818181858888f19350505050158015610c0d573d6000803e3d6000fd5b5050565b6000805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106235780601f106105f857610100808354040283529160200191610623565b60075481565b600a60205260009081526040902054600160a060020a031681565b6002805460408051602060018416156101000260001901909316849004601f810184900484028201840190925281815292918301828280156106235780601f106105f857610100808354040283529160200191610623565b600d60205260009081526040902080546001820154600283015460039093015460ff909216929091600160a060020a039091169084565b600b6020526000908152604090205481565b600154600160a060020a03163314610d4557600080fd5b6008805460ff19166001179055565b60085460ff1681565b6008546000908190819060ff161515610d7557600080fd5b6000848152600c6020526040902092506127108410610d9357600080fd5b825460ff161515610da357600080fd5b6004830154600160a060020a031615801590610dcc57506004830154600160a060020a03163314155b15610dd657600080fd5b8260030154341015610de757600080fd5b6000848152600a60205260409020546002840154600160a060020a03908116911614610e1257600080fd5b60028301546000858152600a602090815260408083208054600160a060020a03191633908117909155600160a060020a03909416808452600b8352818420805460001901905584845292819020805460019081019091558151908152905192955085927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef929181900390910190a3610ea984611772565b600160a060020a0382166000818152600e60209081526040918290208054349081019091558251908152915133939288927f58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e392918290030190a4506000838152600d602052604090206002810154600160a060020a0316331415610bb157600390810154336000908152600e602090815260408083208054909401909355825160808101845282815280820188815281850184815260608301858152998552600d909352939092209151825460ff1916901515178255915160018201559051600282018054600160a060020a031916600160a060020a0390921691909117905593519301929092555050565b60085460009060ff161515610fc957600080fd5b6000828152600a6020526040902054600160a060020a03163314610fec57600080fd5b6127108210610ffa57600080fd5b6000828152600c602052604090205460ff161561101a5761101a82611772565b6000828152600a602090815260408083208054600160a060020a031916600160a060020a03881690811790915533808552600b845282852080546000190190558185529382902080546001908101909155825190815291519093927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef928290030190a3604080518381529051600160a060020a0385169133917f05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d89181900360200190a3506000818152600d602052604090206002810154600160a060020a038481169116141561118d57600381810154600160a060020a038581166000908152600e6020908152604080832080549095019094558351608081018552828152808201888152818601848152606083018581528a8652600d909452959093209051815460ff1916901515178155915160018301559251600282018054600160a060020a031916919093161790915590519101555b505050565b6004805460408051602060026001851615610100026000190190941693909304601f810184900484028201840190925281815292918301828280156106235780601f106105f857610100808354040283529160200191610623565b60008061271083106111fe57600080fd5b60085460ff16151561120f57600080fd5b6000838152600a6020526040902054600160a060020a0316151561123257600080fd5b6000838152600a6020526040902054600160a060020a031633141561125657600080fd5b6000838152600d602052604090206002810154909250600160a060020a0316331461128057600080fd5b60038201546040805191825251339185917f6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c879329181900360200190a3506003818101546040805160808101825260008082526020808301888152838501838152606085018481528a8552600d9093528584209451855460ff19169015151785559051600185015551600284018054600160a060020a031916600160a060020a0390921691909117905551919094015551909133916108fc84150291849190818181858888f19350505050158015610bb1573d6000803e3d6000fd5b600154600160a060020a0316331461137157600080fd5b60085460ff161561138157600080fd5b612710811061138f57600080fd5b6000818152600a6020526040902054600160a060020a03838116911614610c0d576000818152600a6020526040902054600160a060020a0316156113fe576000818152600a6020908152604080832054600160a060020a03168352600b90915290208054600019019055611409565b600980546000190190555b6000818152600a602090815260408083208054600160a060020a031916600160a060020a038716908117909155808452600b83529281902080546001019055805184815290517f8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba929181900390910190a25050565b60085460ff16151561148f57600080fd5b6000838152600a6020526040902054600160a060020a031633146114b257600080fd5b61271083106114c057600080fd5b6040805160a081018252600180825260208083018781523384860190815260608501888152600160a060020a038881166080880181815260008d8152600c88528a90209851895460ff19169015151789559451968801969096559151600287018054600160a060020a03199081169285169290921790559051600387015591516004909501805490921694169390931790925582518581529251909286927f3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb92918290030190a3505050565b60095481565b60085460ff1615156115a357600080fd5b6000828152600a6020526040902054600160a060020a031633146115c657600080fd5b61271082106115d457600080fd5b6040805160a0810182526001808252602080830186815233848601908152606085018781526000608087018181528a8252600c86528882209751885460ff19169015151788559351958701959095559051600286018054600160a060020a0319908116600160a060020a039384161790915591516003870155915160049095018054909116949091169390931790925582518481529251909285927f3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb92918290030190a35050565b60085460ff1615156116ad57600080fd5b60095415156116bb57600080fd5b6000818152600a6020526040902054600160a060020a0316156116dd57600080fd5b61271081106116eb57600080fd5b6000818152600a602090815260408083208054600160a060020a03191633908117909155808452600b8352928190208054600101905560098054600019019055805184815290517f8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba929181900390910190a250565b600e6020526000908152604090205481565b60085460ff16151561178357600080fd5b6000818152600a6020526040902054600160a060020a031633146117a657600080fd5b61271081106117b457600080fd5b6040805160a08101825260008082526020808301858152338486019081526060850184815260808601858152888652600c9094528685209551865460ff19169015151786559151600186015551600285018054600160a060020a0319908116600160a060020a0393841617909155915160038601559151600490940180549091169390911692909217909155905182917fb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd491a2505600a165627a7a7230582025ce732791da45f970d857fe017e5178538adeea409482633ae0de87d1252f990029"

// DeployCryptoPunksMarket deploys a new Ethereum contract, binding an instance of CryptoPunksMarket to it.
func DeployCryptoPunksMarket(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CryptoPunksMarket, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoPunksMarketABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CryptoPunksMarketBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CryptoPunksMarket{CryptoPunksMarketCaller: CryptoPunksMarketCaller{contract: contract}, CryptoPunksMarketTransactor: CryptoPunksMarketTransactor{contract: contract}, CryptoPunksMarketFilterer: CryptoPunksMarketFilterer{contract: contract}}, nil
}

// CryptoPunksMarket is an auto generated Go binding around an Ethereum contract.
type CryptoPunksMarket struct {
	CryptoPunksMarketCaller     // Read-only binding to the contract
	CryptoPunksMarketTransactor // Write-only binding to the contract
	CryptoPunksMarketFilterer   // Log filterer for contract events
}

// CryptoPunksMarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type CryptoPunksMarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CryptoPunksMarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CryptoPunksMarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CryptoPunksMarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CryptoPunksMarketSession struct {
	Contract     *CryptoPunksMarket // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// CryptoPunksMarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CryptoPunksMarketCallerSession struct {
	Contract *CryptoPunksMarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// CryptoPunksMarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CryptoPunksMarketTransactorSession struct {
	Contract     *CryptoPunksMarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// CryptoPunksMarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type CryptoPunksMarketRaw struct {
	Contract *CryptoPunksMarket // Generic contract binding to access the raw methods on
}

// CryptoPunksMarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CryptoPunksMarketCallerRaw struct {
	Contract *CryptoPunksMarketCaller // Generic read-only contract binding to access the raw methods on
}

// CryptoPunksMarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CryptoPunksMarketTransactorRaw struct {
	Contract *CryptoPunksMarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCryptoPunksMarket creates a new instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarket(address common.Address, backend bind.ContractBackend) (*CryptoPunksMarket, error) {
	contract, err := bindCryptoPunksMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarket{CryptoPunksMarketCaller: CryptoPunksMarketCaller{contract: contract}, CryptoPunksMarketTransactor: CryptoPunksMarketTransactor{contract: contract}, CryptoPunksMarketFilterer: CryptoPunksMarketFilterer{contract: contract}}, nil
}

// NewCryptoPunksMarketCaller creates a new read-only instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketCaller(address common.Address, caller bind.ContractCaller) (*CryptoPunksMarketCaller, error) {
	contract, err := bindCryptoPunksMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketCaller{contract: contract}, nil
}

// NewCryptoPunksMarketTransactor creates a new write-only instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*CryptoPunksMarketTransactor, error) {
	contract, err := bindCryptoPunksMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketTransactor{contract: contract}, nil
}

// NewCryptoPunksMarketFilterer creates a new log filterer instance of CryptoPunksMarket, bound to a specific deployed contract.
func NewCryptoPunksMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*CryptoPunksMarketFilterer, error) {
	contract, err := bindCryptoPunksMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketFilterer{contract: contract}, nil
}

// bindCryptoPunksMarket binds a generic wrapper to an already deployed contract.
func bindCryptoPunksMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CryptoPunksMarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPunksMarket *CryptoPunksMarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.CryptoPunksMarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CryptoPunksMarket *CryptoPunksMarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CryptoPunksMarket.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CryptoPunksMarket *CryptoPunksMarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CryptoPunksMarket *CryptoPunksMarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.contract.Transact(opts, method, params...)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() view returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) AllPunksAssigned(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "allPunksAssigned")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() view returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketSession) AllPunksAssigned() (bool, error) {
	return _CryptoPunksMarket.Contract.AllPunksAssigned(&_CryptoPunksMarket.CallOpts)
}

// AllPunksAssigned is a free data retrieval call binding the contract method 0x8126c38a.
//
// Solidity: function allPunksAssigned() view returns(bool)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) AllPunksAssigned() (bool, error) {
	return _CryptoPunksMarket.Contract.AllPunksAssigned(&_CryptoPunksMarket.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) BalanceOf(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "balanceOf", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.BalanceOf(&_CryptoPunksMarket.CallOpts, arg0)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) BalanceOf(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.BalanceOf(&_CryptoPunksMarket.CallOpts, arg0)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Decimals() (uint8, error) {
	return _CryptoPunksMarket.Contract.Decimals(&_CryptoPunksMarket.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Decimals() (uint8, error) {
	return _CryptoPunksMarket.Contract.Decimals(&_CryptoPunksMarket.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) ImageHash(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "imageHash")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) ImageHash() (string, error) {
	return _CryptoPunksMarket.Contract.ImageHash(&_CryptoPunksMarket.CallOpts)
}

// ImageHash is a free data retrieval call binding the contract method 0x51605d80.
//
// Solidity: function imageHash() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) ImageHash() (string, error) {
	return _CryptoPunksMarket.Contract.ImageHash(&_CryptoPunksMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Name() (string, error) {
	return _CryptoPunksMarket.Contract.Name(&_CryptoPunksMarket.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Name() (string, error) {
	return _CryptoPunksMarket.Contract.Name(&_CryptoPunksMarket.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) NextPunkIndexToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "nextPunkIndexToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.NextPunkIndexToAssign(&_CryptoPunksMarket.CallOpts)
}

// NextPunkIndexToAssign is a free data retrieval call binding the contract method 0x52f29a25.
//
// Solidity: function nextPunkIndexToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) NextPunkIndexToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.NextPunkIndexToAssign(&_CryptoPunksMarket.CallOpts)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PendingWithdrawals(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "pendingWithdrawals", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PendingWithdrawals(&_CryptoPunksMarket.CallOpts, arg0)
}

// PendingWithdrawals is a free data retrieval call binding the contract method 0xf3f43703.
//
// Solidity: function pendingWithdrawals(address ) view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PendingWithdrawals(arg0 common.Address) (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PendingWithdrawals(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) view returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunkBids(opts *bind.CallOpts, arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punkBids", arg0)

	outstruct := new(struct {
		HasBid    bool
		PunkIndex *big.Int
		Bidder    common.Address
		Value     *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.HasBid = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PunkIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Bidder = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Value = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) view returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _CryptoPunksMarket.Contract.PunkBids(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkBids is a free data retrieval call binding the contract method 0x6e743fa9.
//
// Solidity: function punkBids(uint256 ) view returns(bool hasBid, uint256 punkIndex, address bidder, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunkBids(arg0 *big.Int) (struct {
	HasBid    bool
	PunkIndex *big.Int
	Bidder    common.Address
	Value     *big.Int
}, error) {
	return _CryptoPunksMarket.Contract.PunkBids(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) view returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunkIndexToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punkIndexToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) view returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoPunksMarket.Contract.PunkIndexToAddress(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunkIndexToAddress is a free data retrieval call binding the contract method 0x58178168.
//
// Solidity: function punkIndexToAddress(uint256 ) view returns(address)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunkIndexToAddress(arg0 *big.Int) (common.Address, error) {
	return _CryptoPunksMarket.Contract.PunkIndexToAddress(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) view returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunksOfferedForSale(opts *bind.CallOpts, arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punksOfferedForSale", arg0)

	outstruct := new(struct {
		IsForSale  bool
		PunkIndex  *big.Int
		Seller     common.Address
		MinValue   *big.Int
		OnlySellTo common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.IsForSale = *abi.ConvertType(out[0], new(bool)).(*bool)
	outstruct.PunkIndex = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Seller = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.MinValue = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.OnlySellTo = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)

	return *outstruct, err

}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) view returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _CryptoPunksMarket.Contract.PunksOfferedForSale(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksOfferedForSale is a free data retrieval call binding the contract method 0x088f11f3.
//
// Solidity: function punksOfferedForSale(uint256 ) view returns(bool isForSale, uint256 punkIndex, address seller, uint256 minValue, address onlySellTo)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunksOfferedForSale(arg0 *big.Int) (struct {
	IsForSale  bool
	PunkIndex  *big.Int
	Seller     common.Address
	MinValue   *big.Int
	OnlySellTo common.Address
}, error) {
	return _CryptoPunksMarket.Contract.PunksOfferedForSale(&_CryptoPunksMarket.CallOpts, arg0)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) PunksRemainingToAssign(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "punksRemainingToAssign")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunksRemainingToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PunksRemainingToAssign(&_CryptoPunksMarket.CallOpts)
}

// PunksRemainingToAssign is a free data retrieval call binding the contract method 0xc0d6ce63.
//
// Solidity: function punksRemainingToAssign() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) PunksRemainingToAssign() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.PunksRemainingToAssign(&_CryptoPunksMarket.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Standard(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "standard")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Standard() (string, error) {
	return _CryptoPunksMarket.Contract.Standard(&_CryptoPunksMarket.CallOpts)
}

// Standard is a free data retrieval call binding the contract method 0x5a3b7e42.
//
// Solidity: function standard() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Standard() (string, error) {
	return _CryptoPunksMarket.Contract.Standard(&_CryptoPunksMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketSession) Symbol() (string, error) {
	return _CryptoPunksMarket.Contract.Symbol(&_CryptoPunksMarket.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) Symbol() (string, error) {
	return _CryptoPunksMarket.Contract.Symbol(&_CryptoPunksMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _CryptoPunksMarket.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketSession) TotalSupply() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.TotalSupply(&_CryptoPunksMarket.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_CryptoPunksMarket *CryptoPunksMarketCallerSession) TotalSupply() (*big.Int, error) {
	return _CryptoPunksMarket.Contract.TotalSupply(&_CryptoPunksMarket.CallOpts)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) AcceptBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "acceptBidForPunk", punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AcceptBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex, minPrice)
}

// AcceptBidForPunk is a paid mutator transaction binding the contract method 0x23165b75.
//
// Solidity: function acceptBidForPunk(uint256 punkIndex, uint256 minPrice) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) AcceptBidForPunk(punkIndex *big.Int, minPrice *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AcceptBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex, minPrice)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) AllInitialOwnersAssigned(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "allInitialOwnersAssigned")
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AllInitialOwnersAssigned(&_CryptoPunksMarket.TransactOpts)
}

// AllInitialOwnersAssigned is a paid mutator transaction binding the contract method 0x7ecedac9.
//
// Solidity: function allInitialOwnersAssigned() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) AllInitialOwnersAssigned() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.AllInitialOwnersAssigned(&_CryptoPunksMarket.TransactOpts)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) BuyPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "buyPunk", punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.BuyPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// BuyPunk is a paid mutator transaction binding the contract method 0x8264fe98.
//
// Solidity: function buyPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) BuyPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.BuyPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) EnterBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "enterBidForPunk", punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.EnterBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// EnterBidForPunk is a paid mutator transaction binding the contract method 0x091dbfd2.
//
// Solidity: function enterBidForPunk(uint256 punkIndex) payable returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) EnterBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.EnterBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) GetPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "getPunk", punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.GetPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// GetPunk is a paid mutator transaction binding the contract method 0xc81d1d5b.
//
// Solidity: function getPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) GetPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.GetPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) OfferPunkForSale(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "offerPunkForSale", punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSale(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSale is a paid mutator transaction binding the contract method 0xc44193c3.
//
// Solidity: function offerPunkForSale(uint256 punkIndex, uint256 minSalePriceInWei) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) OfferPunkForSale(punkIndex *big.Int, minSalePriceInWei *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSale(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) OfferPunkForSaleToAddress(opts *bind.TransactOpts, punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "offerPunkForSaleToAddress", punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSaleToAddress(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// OfferPunkForSaleToAddress is a paid mutator transaction binding the contract method 0xbf31196f.
//
// Solidity: function offerPunkForSaleToAddress(uint256 punkIndex, uint256 minSalePriceInWei, address toAddress) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) OfferPunkForSaleToAddress(punkIndex *big.Int, minSalePriceInWei *big.Int, toAddress common.Address) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.OfferPunkForSaleToAddress(&_CryptoPunksMarket.TransactOpts, punkIndex, minSalePriceInWei, toAddress)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) PunkNoLongerForSale(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "punkNoLongerForSale", punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.PunkNoLongerForSale(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// PunkNoLongerForSale is a paid mutator transaction binding the contract method 0xf6eeff1e.
//
// Solidity: function punkNoLongerForSale(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) PunkNoLongerForSale(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.PunkNoLongerForSale(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) SetInitialOwner(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "setInitialOwner", to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwner(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// SetInitialOwner is a paid mutator transaction binding the contract method 0xa75a9049.
//
// Solidity: function setInitialOwner(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) SetInitialOwner(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwner(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) SetInitialOwners(opts *bind.TransactOpts, addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "setInitialOwners", addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwners(&_CryptoPunksMarket.TransactOpts, addresses, indices)
}

// SetInitialOwners is a paid mutator transaction binding the contract method 0x39c5dde6.
//
// Solidity: function setInitialOwners(address[] addresses, uint256[] indices) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) SetInitialOwners(addresses []common.Address, indices []*big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.SetInitialOwners(&_CryptoPunksMarket.TransactOpts, addresses, indices)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) TransferPunk(opts *bind.TransactOpts, to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "transferPunk", to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.TransferPunk(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// TransferPunk is a paid mutator transaction binding the contract method 0x8b72a2ec.
//
// Solidity: function transferPunk(address to, uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) TransferPunk(to common.Address, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.TransferPunk(&_CryptoPunksMarket.TransactOpts, to, punkIndex)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) Withdraw(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "withdraw")
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) Withdraw() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.Withdraw(&_CryptoPunksMarket.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x3ccfd60b.
//
// Solidity: function withdraw() returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) Withdraw() (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.Withdraw(&_CryptoPunksMarket.TransactOpts)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactor) WithdrawBidForPunk(opts *bind.TransactOpts, punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.contract.Transact(opts, "withdrawBidForPunk", punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.WithdrawBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// WithdrawBidForPunk is a paid mutator transaction binding the contract method 0x979bc638.
//
// Solidity: function withdrawBidForPunk(uint256 punkIndex) returns()
func (_CryptoPunksMarket *CryptoPunksMarketTransactorSession) WithdrawBidForPunk(punkIndex *big.Int) (*types.Transaction, error) {
	return _CryptoPunksMarket.Contract.WithdrawBidForPunk(&_CryptoPunksMarket.TransactOpts, punkIndex)
}

// CryptoPunksMarketAssignIterator is returned from FilterAssign and is used to iterate over the raw logs and unpacked data for Assign events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketAssignIterator struct {
	Event *CryptoPunksMarketAssign // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketAssignIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketAssign)
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
		it.Event = new(CryptoPunksMarketAssign)
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
func (it *CryptoPunksMarketAssignIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketAssignIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketAssign represents a Assign event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketAssign struct {
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterAssign is a free log retrieval operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterAssign(opts *bind.FilterOpts, to []common.Address) (*CryptoPunksMarketAssignIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketAssignIterator{contract: _CryptoPunksMarket.contract, event: "Assign", logs: logs, sub: sub}, nil
}

// WatchAssign is a free log subscription operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchAssign(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketAssign, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "Assign", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketAssign)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "Assign", log); err != nil {
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

// ParseAssign is a log parse operation binding the contract event 0x8a0e37b73a0d9c82e205d4d1a3ff3d0b57ce5f4d7bccf6bac03336dc101cb7ba.
//
// Solidity: event Assign(address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParseAssign(log types.Log) (*CryptoPunksMarketAssign, error) {
	event := new(CryptoPunksMarketAssign)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "Assign", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBidEnteredIterator is returned from FilterPunkBidEntered and is used to iterate over the raw logs and unpacked data for PunkBidEntered events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidEnteredIterator struct {
	Event *CryptoPunksMarketPunkBidEntered // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBidEnteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBidEntered)
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
		it.Event = new(CryptoPunksMarketPunkBidEntered)
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
func (it *CryptoPunksMarketPunkBidEnteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBidEnteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBidEntered represents a PunkBidEntered event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidEntered struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidEntered is a free log retrieval operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBidEntered(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*CryptoPunksMarketPunkBidEnteredIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBidEnteredIterator{contract: _CryptoPunksMarket.contract, event: "PunkBidEntered", logs: logs, sub: sub}, nil
}

// WatchPunkBidEntered is a free log subscription operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBidEntered(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBidEntered, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBidEntered", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBidEntered)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
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

// ParsePunkBidEntered is a log parse operation binding the contract event 0x5b859394fabae0c1ba88baffe67e751ab5248d2e879028b8c8d6897b0519f56a.
//
// Solidity: event PunkBidEntered(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBidEntered(log types.Log) (*CryptoPunksMarketPunkBidEntered, error) {
	event := new(CryptoPunksMarketPunkBidEntered)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidEntered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBidWithdrawnIterator is returned from FilterPunkBidWithdrawn and is used to iterate over the raw logs and unpacked data for PunkBidWithdrawn events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidWithdrawnIterator struct {
	Event *CryptoPunksMarketPunkBidWithdrawn // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBidWithdrawn)
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
		it.Event = new(CryptoPunksMarketPunkBidWithdrawn)
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
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBidWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBidWithdrawn represents a PunkBidWithdrawn event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBidWithdrawn struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBidWithdrawn is a free log retrieval operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBidWithdrawn(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address) (*CryptoPunksMarketPunkBidWithdrawnIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBidWithdrawnIterator{contract: _CryptoPunksMarket.contract, event: "PunkBidWithdrawn", logs: logs, sub: sub}, nil
}

// WatchPunkBidWithdrawn is a free log subscription operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBidWithdrawn(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBidWithdrawn, punkIndex []*big.Int, fromAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBidWithdrawn", punkIndexRule, fromAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBidWithdrawn)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
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

// ParsePunkBidWithdrawn is a log parse operation binding the contract event 0x6f30e1ee4d81dcc7a8a478577f65d2ed2edb120565960ac45fe7c50551c87932.
//
// Solidity: event PunkBidWithdrawn(uint256 indexed punkIndex, uint256 value, address indexed fromAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBidWithdrawn(log types.Log) (*CryptoPunksMarketPunkBidWithdrawn, error) {
	event := new(CryptoPunksMarketPunkBidWithdrawn)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBidWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkBoughtIterator is returned from FilterPunkBought and is used to iterate over the raw logs and unpacked data for PunkBought events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBoughtIterator struct {
	Event *CryptoPunksMarketPunkBought // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkBoughtIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkBought)
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
		it.Event = new(CryptoPunksMarketPunkBought)
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
func (it *CryptoPunksMarketPunkBoughtIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkBoughtIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkBought represents a PunkBought event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkBought struct {
	PunkIndex   *big.Int
	Value       *big.Int
	FromAddress common.Address
	ToAddress   common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterPunkBought is a free log retrieval operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkBought(opts *bind.FilterOpts, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (*CryptoPunksMarketPunkBoughtIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkBoughtIterator{contract: _CryptoPunksMarket.contract, event: "PunkBought", logs: logs, sub: sub}, nil
}

// WatchPunkBought is a free log subscription operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkBought(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkBought, punkIndex []*big.Int, fromAddress []common.Address, toAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var fromAddressRule []interface{}
	for _, fromAddressItem := range fromAddress {
		fromAddressRule = append(fromAddressRule, fromAddressItem)
	}
	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkBought", punkIndexRule, fromAddressRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkBought)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBought", log); err != nil {
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

// ParsePunkBought is a log parse operation binding the contract event 0x58e5d5a525e3b40bc15abaa38b5882678db1ee68befd2f60bafe3a7fd06db9e3.
//
// Solidity: event PunkBought(uint256 indexed punkIndex, uint256 value, address indexed fromAddress, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkBought(log types.Log) (*CryptoPunksMarketPunkBought, error) {
	event := new(CryptoPunksMarketPunkBought)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkBought", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkNoLongerForSaleIterator is returned from FilterPunkNoLongerForSale and is used to iterate over the raw logs and unpacked data for PunkNoLongerForSale events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkNoLongerForSaleIterator struct {
	Event *CryptoPunksMarketPunkNoLongerForSale // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkNoLongerForSale)
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
		it.Event = new(CryptoPunksMarketPunkNoLongerForSale)
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
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkNoLongerForSaleIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkNoLongerForSale represents a PunkNoLongerForSale event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkNoLongerForSale struct {
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkNoLongerForSale is a free log retrieval operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkNoLongerForSale(opts *bind.FilterOpts, punkIndex []*big.Int) (*CryptoPunksMarketPunkNoLongerForSaleIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkNoLongerForSaleIterator{contract: _CryptoPunksMarket.contract, event: "PunkNoLongerForSale", logs: logs, sub: sub}, nil
}

// WatchPunkNoLongerForSale is a free log subscription operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkNoLongerForSale(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkNoLongerForSale, punkIndex []*big.Int) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkNoLongerForSale", punkIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkNoLongerForSale)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
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

// ParsePunkNoLongerForSale is a log parse operation binding the contract event 0xb0e0a660b4e50f26f0b7ce75c24655fc76cc66e3334a54ff410277229fa10bd4.
//
// Solidity: event PunkNoLongerForSale(uint256 indexed punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkNoLongerForSale(log types.Log) (*CryptoPunksMarketPunkNoLongerForSale, error) {
	event := new(CryptoPunksMarketPunkNoLongerForSale)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkNoLongerForSale", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkOfferedIterator is returned from FilterPunkOffered and is used to iterate over the raw logs and unpacked data for PunkOffered events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkOfferedIterator struct {
	Event *CryptoPunksMarketPunkOffered // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkOffered)
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
		it.Event = new(CryptoPunksMarketPunkOffered)
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
func (it *CryptoPunksMarketPunkOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkOffered represents a PunkOffered event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkOffered struct {
	PunkIndex *big.Int
	MinValue  *big.Int
	ToAddress common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkOffered is a free log retrieval operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkOffered(opts *bind.FilterOpts, punkIndex []*big.Int, toAddress []common.Address) (*CryptoPunksMarketPunkOfferedIterator, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkOfferedIterator{contract: _CryptoPunksMarket.contract, event: "PunkOffered", logs: logs, sub: sub}, nil
}

// WatchPunkOffered is a free log subscription operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkOffered(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkOffered, punkIndex []*big.Int, toAddress []common.Address) (event.Subscription, error) {

	var punkIndexRule []interface{}
	for _, punkIndexItem := range punkIndex {
		punkIndexRule = append(punkIndexRule, punkIndexItem)
	}

	var toAddressRule []interface{}
	for _, toAddressItem := range toAddress {
		toAddressRule = append(toAddressRule, toAddressItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkOffered", punkIndexRule, toAddressRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkOffered)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkOffered", log); err != nil {
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

// ParsePunkOffered is a log parse operation binding the contract event 0x3c7b682d5da98001a9b8cbda6c647d2c63d698a4184fd1d55e2ce7b66f5d21eb.
//
// Solidity: event PunkOffered(uint256 indexed punkIndex, uint256 minValue, address indexed toAddress)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkOffered(log types.Log) (*CryptoPunksMarketPunkOffered, error) {
	event := new(CryptoPunksMarketPunkOffered)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkOffered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketPunkTransferIterator is returned from FilterPunkTransfer and is used to iterate over the raw logs and unpacked data for PunkTransfer events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkTransferIterator struct {
	Event *CryptoPunksMarketPunkTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketPunkTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketPunkTransfer)
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
		it.Event = new(CryptoPunksMarketPunkTransfer)
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
func (it *CryptoPunksMarketPunkTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketPunkTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketPunkTransfer represents a PunkTransfer event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketPunkTransfer struct {
	From      common.Address
	To        common.Address
	PunkIndex *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterPunkTransfer is a free log retrieval operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterPunkTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CryptoPunksMarketPunkTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketPunkTransferIterator{contract: _CryptoPunksMarket.contract, event: "PunkTransfer", logs: logs, sub: sub}, nil
}

// WatchPunkTransfer is a free log subscription operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchPunkTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketPunkTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "PunkTransfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketPunkTransfer)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
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

// ParsePunkTransfer is a log parse operation binding the contract event 0x05af636b70da6819000c49f85b21fa82081c632069bb626f30932034099107d8.
//
// Solidity: event PunkTransfer(address indexed from, address indexed to, uint256 punkIndex)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParsePunkTransfer(log types.Log) (*CryptoPunksMarketPunkTransfer, error) {
	event := new(CryptoPunksMarketPunkTransfer)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "PunkTransfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// CryptoPunksMarketTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the CryptoPunksMarket contract.
type CryptoPunksMarketTransferIterator struct {
	Event *CryptoPunksMarketTransfer // Event containing the contract specifics and raw log

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
func (it *CryptoPunksMarketTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(CryptoPunksMarketTransfer)
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
		it.Event = new(CryptoPunksMarketTransfer)
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
func (it *CryptoPunksMarketTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *CryptoPunksMarketTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// CryptoPunksMarketTransfer represents a Transfer event raised by the CryptoPunksMarket contract.
type CryptoPunksMarketTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*CryptoPunksMarketTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &CryptoPunksMarketTransferIterator{contract: _CryptoPunksMarket.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *CryptoPunksMarketTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _CryptoPunksMarket.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(CryptoPunksMarketTransfer)
				if err := _CryptoPunksMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_CryptoPunksMarket *CryptoPunksMarketFilterer) ParseTransfer(log types.Log) (*CryptoPunksMarketTransfer, error) {
	event := new(CryptoPunksMarketTransfer)
	if err := _CryptoPunksMarket.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
