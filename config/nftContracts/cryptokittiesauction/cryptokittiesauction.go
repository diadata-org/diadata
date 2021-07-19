// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptokittiesauction

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

// ClockAuctionABI is the input ABI used to generate the binding from.
const ClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddress\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// ClockAuctionFuncSigs maps the 4-byte function signature to its string representation.
var ClockAuctionFuncSigs = map[string]string{
	"454a2ab3": "bid(uint256)",
	"96b5a755": "cancelAuction(uint256)",
	"878eb368": "cancelAuctionWhenPaused(uint256)",
	"27ebe40a": "createAuction(uint256,uint256,uint256,uint256,address)",
	"78bd7935": "getAuction(uint256)",
	"c55d0f56": "getCurrentPrice(uint256)",
	"dd1b7a0f": "nonFungibleContract()",
	"8da5cb5b": "owner()",
	"83b5ff8b": "ownerCut()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"5fd8c710": "withdrawBalance()",
}

// ClockAuctionBin is the compiled bytecode used for deploying new contracts.
var ClockAuctionBin = "0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b604051604080610e66833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316178155909250905061271082111561006857600080fd5b50600281905581600160a060020a0381166301ffc9a77f9a20483d000000000000000000000000000000000000000000000000000000006040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401602060405180830381600087803b151561010f57600080fd5b5af1151561011c57600080fd5b50505060405180519050151561013157600080fd5b60018054600160a060020a03909216600160a060020a03199092169190911790555050610d03806101636000396000f3006060604052600436106100b65763ffffffff60e060020a60003504166327ebe40a81146100bb5780633f4ba83a146100e8578063454a2ab31461010f5780635c975abb1461011a5780635fd8c7101461012d57806378bd79351461014057806383b5ff8b146101915780638456cb59146101b6578063878eb368146101c95780638da5cb5b146101df57806396b5a7551461020e578063c55d0f5614610224578063dd1b7a0f1461023a578063f2fde38b1461024d575b600080fd5b34156100c657600080fd5b6100e6600435602435604435606435600160a060020a036084351661026c565b005b34156100f357600080fd5b6100fb610355565b604051901515815260200160405180910390f35b6100e66004356103d9565b341561012557600080fd5b6100fb610408565b341561013857600080fd5b6100e6610418565b341561014b57600080fd5b61015660043561048e565b604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390f35b341561019c57600080fd5b6101a461051b565b60405190815260200160405180910390f35b34156101c157600080fd5b6100fb610521565b34156101d457600080fd5b6100e66004356105aa565b34156101ea57600080fd5b6101f261061b565b604051600160a060020a03909116815260200160405180910390f35b341561021957600080fd5b6100e660043561062a565b341561022f57600080fd5b6101a4600435610678565b341561024557600080fd5b6101f26106aa565b341561025857600080fd5b6100e6600160a060020a03600435166106b9565b610274610ca9565b60005460a060020a900460ff161561028b57600080fd5b6001608060020a03851685146102a057600080fd5b6001608060020a03841684146102b557600080fd5b67ffffffffffffffff831683146102cb57600080fd5b6102d53387610710565b15156102e057600080fd5b6102ea3387610788565b60a06040519081016040528083600160a060020a03168152602001866001608060020a03168152602001856001608060020a031681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061034d86826107ff565b505050505050565b6000805433600160a060020a0390811691161461037157600080fd5b60005460a060020a900460ff16151561038957600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff16156103f057600080fd5b6103fa813461099a565b506104053382610acb565b50565b60005460a060020a900460ff1681565b60015460008054600160a060020a03928316923381169116148061044d575081600160a060020a031633600160a060020a0316145b151561045857600080fd5b81600160a060020a03166108fc30600160a060020a0316319081150290604051600060405180830381858888f150505050505050565b600081815260036020526040812081908190819081906104ad81610b21565b15156104b857600080fd5b80546001820154600290920154600160a060020a03909116986001608060020a038084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b6000805433600160a060020a0390811691161461053d57600080fd5b60005460a060020a900460ff161561055457600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b6000805460a060020a900460ff1615156105c357600080fd5b60005433600160a060020a039081169116146105de57600080fd5b5060008181526003602052604090206105f681610b21565b151561060157600080fd5b8054610617908390600160a060020a0316610b42565b5050565b600054600160a060020a031681565b60008181526003602052604081209061064282610b21565b151561064d57600080fd5b508054600160a060020a03908116903316811461066957600080fd5b6106738382610b42565b505050565b600081815260036020526040812061068f81610b21565b151561069a57600080fd5b6106a381610b8c565b9392505050565b600154600160a060020a031681565b60005433600160a060020a039081169116146106d457600080fd5b600160a060020a038116156104055760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b600154600090600160a060020a038085169116636352211e8460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561075f57600080fd5b5af1151561076c57600080fd5b50505060405180519050600160a060020a031614905092915050565b600154600160a060020a03166323b872dd83308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b15156107eb57600080fd5b5af115156107f857600080fd5b5050505050565b603c816060015167ffffffffffffffff16101561081b57600080fd5b600082815260036020526040902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560408201516001820180546001608060020a03928316700100000000000000000000000000000000029216919091179055606082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560808201516002909101805467ffffffffffffffff9290921668010000000000000000026fffffffffffffffff000000000000000019909216919091179055507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78260208301516001608060020a031683604001516001608060020a0316846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b600082815260036020526040812081808080806109b686610b21565b15156109c157600080fd5b6109ca86610b8c565b9450848810156109d957600080fd5b8554600160a060020a031693506109ef89610c13565b6000851115610a3957610a0185610c60565b92508285039150600160a060020a03841682156108fc0283604051600060405180830381858888f193505050501515610a3957600080fd5b50838703600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610a6e57600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd28986336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a15092979650505050505050565b600154600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b15156107eb57600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610b4b82610c13565b610b558183610acb565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df8260405190815260200160405180910390a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610bd25750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106a3916001608060020a0380821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610c6c565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610c8057869350610c9e565b878703925085858402811515610c9257fe5b05915081880190508093505b505050949350505050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a72305820677d4a684ab2167ef72b7640b96a95f22a492811ea7d2a486422586e3a5889d40029"

// DeployClockAuction deploys a new Ethereum contract, binding an instance of ClockAuction to it.
func DeployClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddress common.Address, _cut *big.Int) (common.Address, *types.Transaction, *ClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBin), backend, _nftAddress, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// ClockAuction is an auto generated Go binding around an Ethereum contract.
type ClockAuction struct {
	ClockAuctionCaller     // Read-only binding to the contract
	ClockAuctionTransactor // Write-only binding to the contract
	ClockAuctionFilterer   // Log filterer for contract events
}

// ClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionSession struct {
	Contract     *ClockAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionCallerSession struct {
	Contract *ClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// ClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionTransactorSession struct {
	Contract     *ClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// ClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionRaw struct {
	Contract *ClockAuction // Generic contract binding to access the raw methods on
}

// ClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionCallerRaw struct {
	Contract *ClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionTransactorRaw struct {
	Contract *ClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuction creates a new instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuction(address common.Address, backend bind.ContractBackend) (*ClockAuction, error) {
	contract, err := bindClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuction{ClockAuctionCaller: ClockAuctionCaller{contract: contract}, ClockAuctionTransactor: ClockAuctionTransactor{contract: contract}, ClockAuctionFilterer: ClockAuctionFilterer{contract: contract}}, nil
}

// NewClockAuctionCaller creates a new read-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionCaller, error) {
	contract, err := bindClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionCaller{contract: contract}, nil
}

// NewClockAuctionTransactor creates a new write-only instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionTransactor, error) {
	contract, err := bindClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionTransactor{contract: contract}, nil
}

// NewClockAuctionFilterer creates a new log filterer instance of ClockAuction, bound to a specific deployed contract.
func NewClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionFilterer, error) {
	contract, err := bindClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionFilterer{contract: contract}, nil
}

// bindClockAuction binds a generic wrapper to an already deployed contract.
func bindClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.ClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.ClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuction *ClockAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuction *ClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuction *ClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_ClockAuction *ClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "getAuction", _tokenId)

	outstruct := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})

	outstruct.Seller = out[0].(common.Address)
	outstruct.StartingPrice = out[1].(*big.Int)
	outstruct.EndingPrice = out[2].(*big.Int)
	outstruct.Duration = out[3].(*big.Int)
	outstruct.StartedAt = out[4].(*big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_ClockAuction *ClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_ClockAuction *ClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _ClockAuction.Contract.GetAuction(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_ClockAuction *ClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "getCurrentPrice", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_ClockAuction *ClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _ClockAuction.Contract.GetCurrentPrice(&_ClockAuction.CallOpts, _tokenId)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuction *ClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "nonFungibleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuction *ClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuction *ClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuction.Contract.NonFungibleContract(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClockAuction *ClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClockAuction *ClockAuctionSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ClockAuction *ClockAuctionCallerSession) Owner() (common.Address, error) {
	return _ClockAuction.Contract.Owner(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuction *ClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuction *ClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuction *ClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuction.Contract.OwnerCut(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClockAuction *ClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ClockAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClockAuction *ClockAuctionSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_ClockAuction *ClockAuctionCallerSession) Paused() (bool, error) {
	return _ClockAuction.Contract.Paused(&_ClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_ClockAuction *ClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_ClockAuction *ClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_ClockAuction *ClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.Bid(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuction(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _ClockAuction.Contract.CancelAuctionWhenPaused(&_ClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_ClockAuction *ClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_ClockAuction *ClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_ClockAuction *ClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.CreateAuction(&_ClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Pause(&_ClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClockAuction *ClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClockAuction *ClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ClockAuction *ClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ClockAuction.Contract.TransferOwnership(&_ClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_ClockAuction *ClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _ClockAuction.Contract.Unpause(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_ClockAuction *ClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _ClockAuction.Contract.WithdrawBalance(&_ClockAuction.TransactOpts)
}

// ClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuction contract.
type ClockAuctionAuctionCancelledIterator struct {
	Event *ClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCancelled)
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
		it.Event = new(ClockAuctionAuctionCancelled)
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
func (it *ClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the ClockAuction contract.
type ClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCancelledIterator{contract: _ClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCancelled)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuction *ClockAuctionFilterer) ParseAuctionCancelled(log types.Log) (*ClockAuctionAuctionCancelled, error) {
	event := new(ClockAuctionAuctionCancelled)
	if err := _ClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuction contract.
type ClockAuctionAuctionCreatedIterator struct {
	Event *ClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionCreated)
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
		it.Event = new(ClockAuctionAuctionCreated)
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
func (it *ClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionCreated represents a AuctionCreated event raised by the ClockAuction contract.
type ClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionCreatedIterator{contract: _ClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionCreated)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuction *ClockAuctionFilterer) ParseAuctionCreated(log types.Log) (*ClockAuctionAuctionCreated, error) {
	event := new(ClockAuctionAuctionCreated)
	if err := _ClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessfulIterator struct {
	Event *ClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionAuctionSuccessful)
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
		it.Event = new(ClockAuctionAuctionSuccessful)
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
func (it *ClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuction contract.
type ClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuction *ClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionAuctionSuccessfulIterator{contract: _ClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuction *ClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionAuctionSuccessful)
				if err := _ClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuction *ClockAuctionFilterer) ParseAuctionSuccessful(log types.Log) (*ClockAuctionAuctionSuccessful, error) {
	event := new(ClockAuctionAuctionSuccessful)
	if err := _ClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the ClockAuction contract.
type ClockAuctionPauseIterator struct {
	Event *ClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionPause)
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
		it.Event = new(ClockAuctionPause)
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
func (it *ClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionPause represents a Pause event raised by the ClockAuction contract.
type ClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ClockAuction *ClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*ClockAuctionPauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionPauseIterator{contract: _ClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ClockAuction *ClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *ClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionPause)
				if err := _ClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_ClockAuction *ClockAuctionFilterer) ParsePause(log types.Log) (*ClockAuctionPause, error) {
	event := new(ClockAuctionPause)
	if err := _ClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the ClockAuction contract.
type ClockAuctionUnpauseIterator struct {
	Event *ClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *ClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionUnpause)
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
		it.Event = new(ClockAuctionUnpause)
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
func (it *ClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionUnpause represents a Unpause event raised by the ClockAuction contract.
type ClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ClockAuction *ClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*ClockAuctionUnpauseIterator, error) {

	logs, sub, err := _ClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionUnpauseIterator{contract: _ClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ClockAuction *ClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *ClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _ClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionUnpause)
				if err := _ClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_ClockAuction *ClockAuctionFilterer) ParseUnpause(log types.Log) (*ClockAuctionUnpause, error) {
	event := new(ClockAuctionUnpause)
	if err := _ClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionBaseABI is the input ABI used to generate the binding from.
const ClockAuctionBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"}]"

// ClockAuctionBaseFuncSigs maps the 4-byte function signature to its string representation.
var ClockAuctionBaseFuncSigs = map[string]string{
	"dd1b7a0f": "nonFungibleContract()",
	"83b5ff8b": "ownerCut()",
}

// ClockAuctionBaseBin is the compiled bytecode used for deploying new contracts.
var ClockAuctionBaseBin = "0x6060604052341561000f57600080fd5b60f68061001d6000396000f30060606040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166383b5ff8b8114604d578063dd1b7a0f14606f575b600080fd5b3415605757600080fd5b605d60a8565b60405190815260200160405180910390f35b3415607957600080fd5b607f60ae565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b60015481565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820b7cb6235cfe7070a6706e879770b0f1ac48c203be7c14473f84af0b3e99634fc0029"

// DeployClockAuctionBase deploys a new Ethereum contract, binding an instance of ClockAuctionBase to it.
func DeployClockAuctionBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ClockAuctionBase, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ClockAuctionBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// ClockAuctionBase is an auto generated Go binding around an Ethereum contract.
type ClockAuctionBase struct {
	ClockAuctionBaseCaller     // Read-only binding to the contract
	ClockAuctionBaseTransactor // Write-only binding to the contract
	ClockAuctionBaseFilterer   // Log filterer for contract events
}

// ClockAuctionBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClockAuctionBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClockAuctionBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClockAuctionBaseSession struct {
	Contract     *ClockAuctionBase // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClockAuctionBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClockAuctionBaseCallerSession struct {
	Contract *ClockAuctionBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ClockAuctionBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClockAuctionBaseTransactorSession struct {
	Contract     *ClockAuctionBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ClockAuctionBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClockAuctionBaseRaw struct {
	Contract *ClockAuctionBase // Generic contract binding to access the raw methods on
}

// ClockAuctionBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClockAuctionBaseCallerRaw struct {
	Contract *ClockAuctionBaseCaller // Generic read-only contract binding to access the raw methods on
}

// ClockAuctionBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClockAuctionBaseTransactorRaw struct {
	Contract *ClockAuctionBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClockAuctionBase creates a new instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBase(address common.Address, backend bind.ContractBackend) (*ClockAuctionBase, error) {
	contract, err := bindClockAuctionBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBase{ClockAuctionBaseCaller: ClockAuctionBaseCaller{contract: contract}, ClockAuctionBaseTransactor: ClockAuctionBaseTransactor{contract: contract}, ClockAuctionBaseFilterer: ClockAuctionBaseFilterer{contract: contract}}, nil
}

// NewClockAuctionBaseCaller creates a new read-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseCaller(address common.Address, caller bind.ContractCaller) (*ClockAuctionBaseCaller, error) {
	contract, err := bindClockAuctionBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseCaller{contract: contract}, nil
}

// NewClockAuctionBaseTransactor creates a new write-only instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*ClockAuctionBaseTransactor, error) {
	contract, err := bindClockAuctionBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseTransactor{contract: contract}, nil
}

// NewClockAuctionBaseFilterer creates a new log filterer instance of ClockAuctionBase, bound to a specific deployed contract.
func NewClockAuctionBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*ClockAuctionBaseFilterer, error) {
	contract, err := bindClockAuctionBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseFilterer{contract: contract}, nil
}

// bindClockAuctionBase binds a generic wrapper to an already deployed contract.
func bindClockAuctionBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClockAuctionBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.ClockAuctionBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.ClockAuctionBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ClockAuctionBase *ClockAuctionBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ClockAuctionBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ClockAuctionBase *ClockAuctionBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ClockAuctionBase.Contract.contract.Transact(opts, method, params...)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ClockAuctionBase.contract.Call(opts, &out, "nonFungibleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuctionBase *ClockAuctionBaseSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) NonFungibleContract() (common.Address, error) {
	return _ClockAuctionBase.Contract.NonFungibleContract(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ClockAuctionBase.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_ClockAuctionBase *ClockAuctionBaseCallerSession) OwnerCut() (*big.Int, error) {
	return _ClockAuctionBase.Contract.OwnerCut(&_ClockAuctionBase.CallOpts)
}

// ClockAuctionBaseAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelledIterator struct {
	Event *ClockAuctionBaseAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCancelled)
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
		it.Event = new(ClockAuctionBaseAuctionCancelled)
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
func (it *ClockAuctionBaseAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCancelled represents a AuctionCancelled event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCancelledIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCancelledIterator{contract: _ClockAuctionBase.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCancelled)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) ParseAuctionCancelled(log types.Log) (*ClockAuctionBaseAuctionCancelled, error) {
	event := new(ClockAuctionBaseAuctionCancelled)
	if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionBaseAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreatedIterator struct {
	Event *ClockAuctionBaseAuctionCreated // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionCreated)
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
		it.Event = new(ClockAuctionBaseAuctionCreated)
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
func (it *ClockAuctionBaseAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionCreated represents a AuctionCreated event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionCreatedIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionCreatedIterator{contract: _ClockAuctionBase.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionCreated)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) ParseAuctionCreated(log types.Log) (*ClockAuctionBaseAuctionCreated, error) {
	event := new(ClockAuctionBaseAuctionCreated)
	if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ClockAuctionBaseAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessfulIterator struct {
	Event *ClockAuctionBaseAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
		it.Event = new(ClockAuctionBaseAuctionSuccessful)
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
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ClockAuctionBaseAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ClockAuctionBaseAuctionSuccessful represents a AuctionSuccessful event raised by the ClockAuctionBase contract.
type ClockAuctionBaseAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*ClockAuctionBaseAuctionSuccessfulIterator, error) {

	logs, sub, err := _ClockAuctionBase.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &ClockAuctionBaseAuctionSuccessfulIterator{contract: _ClockAuctionBase.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *ClockAuctionBaseAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _ClockAuctionBase.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ClockAuctionBaseAuctionSuccessful)
				if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_ClockAuctionBase *ClockAuctionBaseFilterer) ParseAuctionSuccessful(log types.Log) (*ClockAuctionBaseAuctionSuccessful, error) {
	event := new(ClockAuctionBaseAuctionSuccessful)
	if err := _ClockAuctionBase.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721ABI is the input ABI used to generate the binding from.
const ERC721ABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"total\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// ERC721FuncSigs maps the 4-byte function signature to its string representation.
var ERC721FuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"6352211e": "ownerOf(uint256)",
	"01ffc9a7": "supportsInterface(bytes4)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
}

// ERC721 is an auto generated Go binding around an Ethereum contract.
type ERC721 struct {
	ERC721Caller     // Read-only binding to the contract
	ERC721Transactor // Write-only binding to the contract
	ERC721Filterer   // Log filterer for contract events
}

// ERC721Caller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Transactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721Session struct {
	Contract     *ERC721           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721CallerSession struct {
	Contract *ERC721Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ERC721TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721TransactorSession struct {
	Contract     *ERC721Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721Raw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721Raw struct {
	Contract *ERC721 // Generic contract binding to access the raw methods on
}

// ERC721CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721CallerRaw struct {
	Contract *ERC721Caller // Generic read-only contract binding to access the raw methods on
}

// ERC721TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721TransactorRaw struct {
	Contract *ERC721Transactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721 creates a new instance of ERC721, bound to a specific deployed contract.
func NewERC721(address common.Address, backend bind.ContractBackend) (*ERC721, error) {
	contract, err := bindERC721(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721{ERC721Caller: ERC721Caller{contract: contract}, ERC721Transactor: ERC721Transactor{contract: contract}, ERC721Filterer: ERC721Filterer{contract: contract}}, nil
}

// NewERC721Caller creates a new read-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Caller(address common.Address, caller bind.ContractCaller) (*ERC721Caller, error) {
	contract, err := bindERC721(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Caller{contract: contract}, nil
}

// NewERC721Transactor creates a new write-only instance of ERC721, bound to a specific deployed contract.
func NewERC721Transactor(address common.Address, transactor bind.ContractTransactor) (*ERC721Transactor, error) {
	contract, err := bindERC721(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721Transactor{contract: contract}, nil
}

// NewERC721Filterer creates a new log filterer instance of ERC721, bound to a specific deployed contract.
func NewERC721Filterer(address common.Address, filterer bind.ContractFilterer) (*ERC721Filterer, error) {
	contract, err := bindERC721(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721Filterer{contract: contract}, nil
}

// bindERC721 binds a generic wrapper to an already deployed contract.
func bindERC721(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.ERC721Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.ERC721Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721 *ERC721CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721 *ERC721TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721 *ERC721TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ERC721 *ERC721Caller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ERC721 *ERC721Session) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 balance)
func (_ERC721 *ERC721CallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _ERC721.Contract.BalanceOf(&_ERC721.CallOpts, _owner)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_ERC721 *ERC721Caller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_ERC721 *ERC721Session) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_ERC721 *ERC721CallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _ERC721.Contract.OwnerOf(&_ERC721.CallOpts, _tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_ERC721 *ERC721Caller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_ERC721 *ERC721Session) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_ERC721 *ERC721CallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _ERC721.Contract.SupportsInterface(&_ERC721.CallOpts, _interfaceID)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 total)
func (_ERC721 *ERC721Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ERC721.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 total)
func (_ERC721 *ERC721Session) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 total)
func (_ERC721 *ERC721CallerSession) TotalSupply() (*big.Int, error) {
	return _ERC721.Contract.TotalSupply(&_ERC721.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Approve(&_ERC721.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Transfer(&_ERC721.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.Transfer(&_ERC721.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Transactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721Session) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_ERC721 *ERC721TransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _ERC721.Contract.TransferFrom(&_ERC721.TransactOpts, _from, _to, _tokenId)
}

// ERC721ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the ERC721 contract.
type ERC721ApprovalIterator struct {
	Event *ERC721Approval // Event containing the contract specifics and raw log

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
func (it *ERC721ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Approval)
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
		it.Event = new(ERC721Approval)
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
func (it *ERC721ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Approval represents a Approval event raised by the ERC721 contract.
type ERC721Approval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_ERC721 *ERC721Filterer) FilterApproval(opts *bind.FilterOpts) (*ERC721ApprovalIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &ERC721ApprovalIterator{contract: _ERC721.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_ERC721 *ERC721Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *ERC721Approval) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Approval)
				if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
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
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_ERC721 *ERC721Filterer) ParseApproval(log types.Log) (*ERC721Approval, error) {
	event := new(ERC721Approval)
	if err := _ERC721.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ERC721TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the ERC721 contract.
type ERC721TransferIterator struct {
	Event *ERC721Transfer // Event containing the contract specifics and raw log

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
func (it *ERC721TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ERC721Transfer)
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
		it.Event = new(ERC721Transfer)
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
func (it *ERC721TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ERC721TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ERC721Transfer represents a Transfer event raised by the ERC721 contract.
type ERC721Transfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_ERC721 *ERC721Filterer) FilterTransfer(opts *bind.FilterOpts) (*ERC721TransferIterator, error) {

	logs, sub, err := _ERC721.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &ERC721TransferIterator{contract: _ERC721.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_ERC721 *ERC721Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *ERC721Transfer) (event.Subscription, error) {

	logs, sub, err := _ERC721.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ERC721Transfer)
				if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
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
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_ERC721 *ERC721Filterer) ParseTransfer(log types.Log) (*ERC721Transfer, error) {
	event := new(ERC721Transfer)
	if err := _ERC721.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OwnableABI is the input ABI used to generate the binding from.
const OwnableABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// OwnableFuncSigs maps the 4-byte function signature to its string representation.
var OwnableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"f2fde38b": "transferOwnership(address)",
}

// OwnableBin is the compiled bytecode used for deploying new contracts.
var OwnableBin = "0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561017f8061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008c575b600080fd5b341561005b57600080fd5b6100636100ba565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561009757600080fd5b6100b873ffffffffffffffffffffffffffffffffffffffff600435166100d6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100fe57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610150576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a723058209c6071079fe4d3455176e6b30f735ca99e5b67d24eb7deef2648d7ea1b152c3e0029"

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
func (_Ownable *OwnableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
func (_Ownable *OwnableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
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
	var out []interface{}
	err := _Ownable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

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
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Ownable *OwnableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Ownable.Contract.TransferOwnership(&_Ownable.TransactOpts, newOwner)
}

// PausableABI is the input ABI used to generate the binding from.
const PausableABI = "[{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// PausableFuncSigs maps the 4-byte function signature to its string representation.
var PausableFuncSigs = map[string]string{
	"8da5cb5b": "owner()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
}

// PausableBin is the compiled bytecode used for deploying new contracts.
var PausableBin = "0x606060405260008054600160a060020a033316600160a860020a0319909116179055610300806100306000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100985780638456cb59146100ab5780638da5cb5b146100be578063f2fde38b146100ed575b600080fd5b341561007c57600080fd5b61008461010e565b604051901515815260200160405180910390f35b34156100a357600080fd5b6100846101a3565b34156100b657600080fd5b6100846101c4565b34156100c957600080fd5b6100d161026f565b604051600160a060020a03909116815260200160405180910390f35b34156100f857600080fd5b61010c600160a060020a036004351661027e565b005b6000805433600160a060020a0390811691161461012a57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561015357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005474010000000000000000000000000000000000000000900460ff1681565b6000805433600160a060020a039081169116146101e057600080fd5b60005474010000000000000000000000000000000000000000900460ff161561020857600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b60005433600160a060020a0390811691161461029957600080fd5b600160a060020a038116156102d1576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a723058208f4ca20f8ce8152a24279eb8d5af2f301cf0e4b10593703bf9d8493817a14b190029"

// DeployPausable deploys a new Ethereum contract, binding an instance of Pausable to it.
func DeployPausable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Pausable, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(PausableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// Pausable is an auto generated Go binding around an Ethereum contract.
type Pausable struct {
	PausableCaller     // Read-only binding to the contract
	PausableTransactor // Write-only binding to the contract
	PausableFilterer   // Log filterer for contract events
}

// PausableCaller is an auto generated read-only Go binding around an Ethereum contract.
type PausableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PausableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PausableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PausableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PausableSession struct {
	Contract     *Pausable         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PausableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PausableCallerSession struct {
	Contract *PausableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// PausableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PausableTransactorSession struct {
	Contract     *PausableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// PausableRaw is an auto generated low-level Go binding around an Ethereum contract.
type PausableRaw struct {
	Contract *Pausable // Generic contract binding to access the raw methods on
}

// PausableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PausableCallerRaw struct {
	Contract *PausableCaller // Generic read-only contract binding to access the raw methods on
}

// PausableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PausableTransactorRaw struct {
	Contract *PausableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPausable creates a new instance of Pausable, bound to a specific deployed contract.
func NewPausable(address common.Address, backend bind.ContractBackend) (*Pausable, error) {
	contract, err := bindPausable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Pausable{PausableCaller: PausableCaller{contract: contract}, PausableTransactor: PausableTransactor{contract: contract}, PausableFilterer: PausableFilterer{contract: contract}}, nil
}

// NewPausableCaller creates a new read-only instance of Pausable, bound to a specific deployed contract.
func NewPausableCaller(address common.Address, caller bind.ContractCaller) (*PausableCaller, error) {
	contract, err := bindPausable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PausableCaller{contract: contract}, nil
}

// NewPausableTransactor creates a new write-only instance of Pausable, bound to a specific deployed contract.
func NewPausableTransactor(address common.Address, transactor bind.ContractTransactor) (*PausableTransactor, error) {
	contract, err := bindPausable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PausableTransactor{contract: contract}, nil
}

// NewPausableFilterer creates a new log filterer instance of Pausable, bound to a specific deployed contract.
func NewPausableFilterer(address common.Address, filterer bind.ContractFilterer) (*PausableFilterer, error) {
	contract, err := bindPausable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PausableFilterer{contract: contract}, nil
}

// bindPausable binds a generic wrapper to an already deployed contract.
func bindPausable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PausableABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.PausableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.PausableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Pausable *PausableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Pausable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Pausable *PausableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Pausable *PausableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Pausable.Contract.contract.Transact(opts, method, params...)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Pausable *PausableCallerSession) Owner() (common.Address, error) {
	return _Pausable.Contract.Owner(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Pausable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Pausable *PausableCallerSession) Paused() (bool, error) {
	return _Pausable.Contract.Paused(&_Pausable.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_Pausable *PausableTransactorSession) Pause() (*types.Transaction, error) {
	return _Pausable.Contract.Pause(&_Pausable.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Pausable *PausableTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Pausable.Contract.TransferOwnership(&_Pausable.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Pausable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_Pausable *PausableTransactorSession) Unpause() (*types.Transaction, error) {
	return _Pausable.Contract.Unpause(&_Pausable.TransactOpts)
}

// PausablePauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the Pausable contract.
type PausablePauseIterator struct {
	Event *PausablePause // Event containing the contract specifics and raw log

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
func (it *PausablePauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausablePause)
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
		it.Event = new(PausablePause)
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
func (it *PausablePauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausablePauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausablePause represents a Pause event raised by the Pausable contract.
type PausablePause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) FilterPause(opts *bind.FilterOpts) (*PausablePauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PausablePauseIterator{contract: _Pausable.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PausablePause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausablePause)
				if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_Pausable *PausableFilterer) ParsePause(log types.Log) (*PausablePause, error) {
	event := new(PausablePause)
	if err := _Pausable.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PausableUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the Pausable contract.
type PausableUnpauseIterator struct {
	Event *PausableUnpause // Event containing the contract specifics and raw log

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
func (it *PausableUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PausableUnpause)
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
		it.Event = new(PausableUnpause)
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
func (it *PausableUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PausableUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PausableUnpause represents a Unpause event raised by the Pausable contract.
type PausableUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) FilterUnpause(opts *bind.FilterOpts) (*PausableUnpauseIterator, error) {

	logs, sub, err := _Pausable.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PausableUnpauseIterator{contract: _Pausable.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PausableUnpause) (event.Subscription, error) {

	logs, sub, err := _Pausable.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PausableUnpause)
				if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_Pausable *PausableFilterer) ParseUnpause(log types.Log) (*PausableUnpause, error) {
	event := new(PausableUnpause)
	if err := _Pausable.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleClockAuctionABI is the input ABI used to generate the binding from.
const SaleClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"lastGen0SalePrices\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSaleClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0SaleCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"averageGen0SalePrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SaleClockAuctionFuncSigs maps the 4-byte function signature to its string representation.
var SaleClockAuctionFuncSigs = map[string]string{
	"eac9d94c": "averageGen0SalePrice()",
	"454a2ab3": "bid(uint256)",
	"96b5a755": "cancelAuction(uint256)",
	"878eb368": "cancelAuctionWhenPaused(uint256)",
	"27ebe40a": "createAuction(uint256,uint256,uint256,uint256,address)",
	"8a98a9cc": "gen0SaleCount()",
	"78bd7935": "getAuction(uint256)",
	"c55d0f56": "getCurrentPrice(uint256)",
	"85b86188": "isSaleClockAuction()",
	"484eccb4": "lastGen0SalePrices(uint256)",
	"dd1b7a0f": "nonFungibleContract()",
	"8da5cb5b": "owner()",
	"83b5ff8b": "ownerCut()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"5fd8c710": "withdrawBalance()",
}

// SaleClockAuctionBin is the compiled bytecode used for deploying new contracts.
var SaleClockAuctionBin = "0x60606040526000805460a060020a60ff02191690556004805460ff19166001179055341561002c57600080fd5b604051604080610ef6833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316178155909250839150829061271082111561007857600080fd5b50600281905581600160a060020a0381166301ffc9a77f9a20483d000000000000000000000000000000000000000000000000000000006040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401602060405180830381600087803b151561011f57600080fd5b5af1151561012c57600080fd5b50505060405180519050151561014157600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610d81806101756000396000f3006060604052600436106100e25763ffffffff60e060020a60003504166327ebe40a81146100e75780633f4ba83a14610114578063454a2ab31461013b578063484eccb4146101465780635c975abb1461016e5780635fd8c7101461018157806378bd79351461019457806383b5ff8b146101e55780638456cb59146101f857806385b861881461020b578063878eb3681461021e5780638a98a9cc146102345780638da5cb5b1461024757806396b5a75514610276578063c55d0f561461028c578063dd1b7a0f146102a2578063eac9d94c146102b5578063f2fde38b146102c8575b600080fd5b34156100f257600080fd5b610112600435602435604435606435600160a060020a03608435166102e7565b005b341561011f57600080fd5b6101276103bf565b604051901515815260200160405180910390f35b610112600435610443565b341561015157600080fd5b61015c6004356104ad565b60405190815260200160405180910390f35b341561017957600080fd5b6101276104c1565b341561018c57600080fd5b6101126104d1565b341561019f57600080fd5b6101aa600435610547565b604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390f35b34156101f057600080fd5b61015c6105d4565b341561020357600080fd5b6101276105da565b341561021657600080fd5b610127610663565b341561022957600080fd5b61011260043561066c565b341561023f57600080fd5b61015c6106dd565b341561025257600080fd5b61025a6106e3565b604051600160a060020a03909116815260200160405180910390f35b341561028157600080fd5b6101126004356106f2565b341561029757600080fd5b61015c60043561073b565b34156102ad57600080fd5b61025a61076d565b34156102c057600080fd5b61015c61077c565b34156102d357600080fd5b610112600160a060020a03600435166107b0565b6102ef610d27565b6001608060020a038516851461030457600080fd5b6001608060020a038416841461031957600080fd5b67ffffffffffffffff8316831461032f57600080fd5b60015433600160a060020a0390811691161461034a57600080fd5b6103548287610806565b60a06040519081016040528083600160a060020a03168152602001866001608060020a03168152602001856001608060020a031681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681525090506103b7868261087d565b505050505050565b6000805433600160a060020a039081169116146103db57600080fd5b60005460a060020a900460ff1615156103f357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b600081815260036020526040812054600160a060020a0316906104668334610a18565b90506104723384610b49565b600154600160a060020a03838116911614156104a857600580548291600691066005811061049c57fe5b01556005805460010190555b505050565b600681600581106104ba57fe5b0154905081565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169233811691161480610506575081600160a060020a031633600160a060020a0316145b151561051157600080fd5b81600160a060020a03166108fc30600160a060020a0316319081150290604051600060405180830381858888f150505050505050565b6000818152600360205260408120819081908190819061056681610b9f565b151561057157600080fd5b80546001820154600290920154600160a060020a03909116986001608060020a038084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b6000805433600160a060020a039081169116146105f657600080fd5b60005460a060020a900460ff161561060d57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b60045460ff1681565b6000805460a060020a900460ff16151561068557600080fd5b60005433600160a060020a039081169116146106a057600080fd5b5060008181526003602052604090206106b881610b9f565b15156106c357600080fd5b80546106d9908390600160a060020a0316610bc0565b5050565b60055481565b600054600160a060020a031681565b60008181526003602052604081209061070a82610b9f565b151561071557600080fd5b508054600160a060020a03908116903316811461073157600080fd5b6104a88382610bc0565b600081815260036020526040812061075281610b9f565b151561075d57600080fd5b61076681610c0a565b9392505050565b600154600160a060020a031681565b600080805b60058110156107a6576006816005811061079757fe5b01549190910190600101610781565b5060059004919050565b60005433600160a060020a039081169116146107cb57600080fd5b600160a060020a03811615610803576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154600160a060020a03166323b872dd83308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b151561086957600080fd5b5af1151561087657600080fd5b5050505050565b603c816060015167ffffffffffffffff16101561089957600080fd5b600082815260036020526040902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560408201516001820180546001608060020a03928316700100000000000000000000000000000000029216919091179055606082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560808201516002909101805467ffffffffffffffff9290921668010000000000000000026fffffffffffffffff000000000000000019909216919091179055507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78260208301516001608060020a031683604001516001608060020a0316846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b60008281526003602052604081208180808080610a3486610b9f565b1515610a3f57600080fd5b610a4886610c0a565b945084881015610a5757600080fd5b8554600160a060020a03169350610a6d89610c91565b6000851115610ab757610a7f85610cde565b92508285039150600160a060020a03841682156108fc0283604051600060405180830381858888f193505050501515610ab757600080fd5b50838703600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610aec57600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd28986336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a15092979650505050505050565b600154600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561086957600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610bc982610c91565b610bd38183610b49565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df8260405190815260200160405180910390a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610c505750600282015468010000000000000000900467ffffffffffffffff1642035b60018301546002840154610766916001608060020a0380821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610cea565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610cfe57869350610d1c565b878703925085858402811515610d1057fe5b05915081880190508093505b505050949350505050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a7230582033b714da41785b626c601bfc003f79c83f61e8b703ffe45458694d13d1f622d70029"

// DeploySaleClockAuction deploys a new Ethereum contract, binding an instance of SaleClockAuction to it.
func DeploySaleClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SaleClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SaleClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// SaleClockAuction is an auto generated Go binding around an Ethereum contract.
type SaleClockAuction struct {
	SaleClockAuctionCaller     // Read-only binding to the contract
	SaleClockAuctionTransactor // Write-only binding to the contract
	SaleClockAuctionFilterer   // Log filterer for contract events
}

// SaleClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SaleClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SaleClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SaleClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SaleClockAuctionSession struct {
	Contract     *SaleClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SaleClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SaleClockAuctionCallerSession struct {
	Contract *SaleClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// SaleClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SaleClockAuctionTransactorSession struct {
	Contract     *SaleClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// SaleClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SaleClockAuctionRaw struct {
	Contract *SaleClockAuction // Generic contract binding to access the raw methods on
}

// SaleClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SaleClockAuctionCallerRaw struct {
	Contract *SaleClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SaleClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SaleClockAuctionTransactorRaw struct {
	Contract *SaleClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSaleClockAuction creates a new instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuction(address common.Address, backend bind.ContractBackend) (*SaleClockAuction, error) {
	contract, err := bindSaleClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuction{SaleClockAuctionCaller: SaleClockAuctionCaller{contract: contract}, SaleClockAuctionTransactor: SaleClockAuctionTransactor{contract: contract}, SaleClockAuctionFilterer: SaleClockAuctionFilterer{contract: contract}}, nil
}

// NewSaleClockAuctionCaller creates a new read-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SaleClockAuctionCaller, error) {
	contract, err := bindSaleClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionCaller{contract: contract}, nil
}

// NewSaleClockAuctionTransactor creates a new write-only instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SaleClockAuctionTransactor, error) {
	contract, err := bindSaleClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionTransactor{contract: contract}, nil
}

// NewSaleClockAuctionFilterer creates a new log filterer instance of SaleClockAuction, bound to a specific deployed contract.
func NewSaleClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SaleClockAuctionFilterer, error) {
	contract, err := bindSaleClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionFilterer{contract: contract}, nil
}

// bindSaleClockAuction binds a generic wrapper to an already deployed contract.
func bindSaleClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SaleClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.SaleClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.SaleClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SaleClockAuction *SaleClockAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SaleClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SaleClockAuction *SaleClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.contract.Transact(opts, method, params...)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) AverageGen0SalePrice(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "averageGen0SalePrice")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// AverageGen0SalePrice is a free data retrieval call binding the contract method 0xeac9d94c.
//
// Solidity: function averageGen0SalePrice() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) AverageGen0SalePrice() (*big.Int, error) {
	return _SaleClockAuction.Contract.AverageGen0SalePrice(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) Gen0SaleCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "gen0SaleCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// Gen0SaleCount is a free data retrieval call binding the contract method 0x8a98a9cc.
//
// Solidity: function gen0SaleCount() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Gen0SaleCount() (*big.Int, error) {
	return _SaleClockAuction.Contract.Gen0SaleCount(&_SaleClockAuction.CallOpts)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_SaleClockAuction *SaleClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "getAuction", _tokenId)

	outstruct := new(struct {
		Seller        common.Address
		StartingPrice *big.Int
		EndingPrice   *big.Int
		Duration      *big.Int
		StartedAt     *big.Int
	})

	outstruct.Seller = out[0].(common.Address)
	outstruct.StartingPrice = out[1].(*big.Int)
	outstruct.EndingPrice = out[2].(*big.Int)
	outstruct.Duration = out[3].(*big.Int)
	outstruct.StartedAt = out[4].(*big.Int)

	return *outstruct, err

}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_SaleClockAuction *SaleClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SaleClockAuction.Contract.GetAuction(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "getCurrentPrice", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.GetCurrentPrice(&_SaleClockAuction.CallOpts, _tokenId)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) IsSaleClockAuction(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "isSaleClockAuction")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// IsSaleClockAuction is a free data retrieval call binding the contract method 0x85b86188.
//
// Solidity: function isSaleClockAuction() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) IsSaleClockAuction() (bool, error) {
	return _SaleClockAuction.Contract.IsSaleClockAuction(&_SaleClockAuction.CallOpts)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices(uint256 ) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) LastGen0SalePrices(opts *bind.CallOpts, arg0 *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "lastGen0SalePrices", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices(uint256 ) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// LastGen0SalePrices is a free data retrieval call binding the contract method 0x484eccb4.
//
// Solidity: function lastGen0SalePrices(uint256 ) view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) LastGen0SalePrices(arg0 *big.Int) (*big.Int, error) {
	return _SaleClockAuction.Contract.LastGen0SalePrices(&_SaleClockAuction.CallOpts, arg0)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "nonFungibleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SaleClockAuction.Contract.NonFungibleContract(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SaleClockAuction *SaleClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SaleClockAuction *SaleClockAuctionSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SaleClockAuction.Contract.Owner(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SaleClockAuction *SaleClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SaleClockAuction.Contract.OwnerCut(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SaleClockAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SaleClockAuction *SaleClockAuctionCallerSession) Paused() (bool, error) {
	return _SaleClockAuction.Contract.Paused(&_SaleClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SaleClockAuction *SaleClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Bid(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuction(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CancelAuctionWhenPaused(&_SaleClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SaleClockAuction *SaleClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.CreateAuction(&_SaleClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Pause(&_SaleClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SaleClockAuction *SaleClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SaleClockAuction.Contract.TransferOwnership(&_SaleClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SaleClockAuction *SaleClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.Unpause(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SaleClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SaleClockAuction *SaleClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SaleClockAuction.Contract.WithdrawBalance(&_SaleClockAuction.TransactOpts)
}

// SaleClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelledIterator struct {
	Event *SaleClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCancelled)
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
		it.Event = new(SaleClockAuctionAuctionCancelled)
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
func (it *SaleClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCancelledIterator{contract: _SaleClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCancelled)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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

// ParseAuctionCancelled is a log parse operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_SaleClockAuction *SaleClockAuctionFilterer) ParseAuctionCancelled(log types.Log) (*SaleClockAuctionAuctionCancelled, error) {
	event := new(SaleClockAuctionAuctionCancelled)
	if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreatedIterator struct {
	Event *SaleClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionCreated)
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
		it.Event = new(SaleClockAuctionAuctionCreated)
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
func (it *SaleClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionCreated represents a AuctionCreated event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SaleClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionCreatedIterator{contract: _SaleClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionCreated)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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

// ParseAuctionCreated is a log parse operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_SaleClockAuction *SaleClockAuctionFilterer) ParseAuctionCreated(log types.Log) (*SaleClockAuctionAuctionCreated, error) {
	event := new(SaleClockAuctionAuctionCreated)
	if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessfulIterator struct {
	Event *SaleClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionAuctionSuccessful)
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
		it.Event = new(SaleClockAuctionAuctionSuccessful)
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
func (it *SaleClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SaleClockAuction contract.
type SaleClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SaleClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionAuctionSuccessfulIterator{contract: _SaleClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionAuctionSuccessful)
				if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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

// ParseAuctionSuccessful is a log parse operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_SaleClockAuction *SaleClockAuctionFilterer) ParseAuctionSuccessful(log types.Log) (*SaleClockAuctionAuctionSuccessful, error) {
	event := new(SaleClockAuctionAuctionSuccessful)
	if err := _SaleClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SaleClockAuction contract.
type SaleClockAuctionPauseIterator struct {
	Event *SaleClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionPause)
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
		it.Event = new(SaleClockAuctionPause)
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
func (it *SaleClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionPause represents a Pause event raised by the SaleClockAuction contract.
type SaleClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SaleClockAuctionPauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionPauseIterator{contract: _SaleClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionPause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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

// ParsePause is a log parse operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_SaleClockAuction *SaleClockAuctionFilterer) ParsePause(log types.Log) (*SaleClockAuctionPause, error) {
	event := new(SaleClockAuctionPause)
	if err := _SaleClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SaleClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SaleClockAuction contract.
type SaleClockAuctionUnpauseIterator struct {
	Event *SaleClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SaleClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SaleClockAuctionUnpause)
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
		it.Event = new(SaleClockAuctionUnpause)
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
func (it *SaleClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SaleClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SaleClockAuctionUnpause represents a Unpause event raised by the SaleClockAuction contract.
type SaleClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SaleClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SaleClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SaleClockAuctionUnpauseIterator{contract: _SaleClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SaleClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SaleClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SaleClockAuctionUnpause)
				if err := _SaleClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// ParseUnpause is a log parse operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_SaleClockAuction *SaleClockAuctionFilterer) ParseUnpause(log types.Log) (*SaleClockAuctionUnpause, error) {
	event := new(SaleClockAuctionUnpause)
	if err := _SaleClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
