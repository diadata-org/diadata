// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package cryptokitties

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
var ClockAuctionBin = "0x60606040526000805460a060020a60ff0219169055341561001f57600080fd5b604051604080610e66833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316178155909250905061271082111561006857600080fd5b50600281905581600160a060020a0381166301ffc9a77f9a20483d000000000000000000000000000000000000000000000000000000006040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401602060405180830381600087803b151561010f57600080fd5b5af1151561011c57600080fd5b50505060405180519050151561013157600080fd5b60018054600160a060020a03909216600160a060020a03199092169190911790555050610d03806101636000396000f3006060604052600436106100b65763ffffffff60e060020a60003504166327ebe40a81146100bb5780633f4ba83a146100e8578063454a2ab31461010f5780635c975abb1461011a5780635fd8c7101461012d57806378bd79351461014057806383b5ff8b146101915780638456cb59146101b6578063878eb368146101c95780638da5cb5b146101df57806396b5a7551461020e578063c55d0f5614610224578063dd1b7a0f1461023a578063f2fde38b1461024d575b600080fd5b34156100c657600080fd5b6100e6600435602435604435606435600160a060020a036084351661026c565b005b34156100f357600080fd5b6100fb610355565b604051901515815260200160405180910390f35b6100e66004356103d9565b341561012557600080fd5b6100fb610408565b341561013857600080fd5b6100e6610418565b341561014b57600080fd5b61015660043561048e565b604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390f35b341561019c57600080fd5b6101a461051b565b60405190815260200160405180910390f35b34156101c157600080fd5b6100fb610521565b34156101d457600080fd5b6100e66004356105aa565b34156101ea57600080fd5b6101f261061b565b604051600160a060020a03909116815260200160405180910390f35b341561021957600080fd5b6100e660043561062a565b341561022f57600080fd5b6101a4600435610678565b341561024557600080fd5b6101f26106aa565b341561025857600080fd5b6100e6600160a060020a03600435166106b9565b610274610ca9565b60005460a060020a900460ff161561028b57600080fd5b6001608060020a03851685146102a057600080fd5b6001608060020a03841684146102b557600080fd5b67ffffffffffffffff831683146102cb57600080fd5b6102d53387610710565b15156102e057600080fd5b6102ea3387610788565b60a06040519081016040528083600160a060020a03168152602001866001608060020a03168152602001856001608060020a031681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061034d86826107ff565b505050505050565b6000805433600160a060020a0390811691161461037157600080fd5b60005460a060020a900460ff16151561038957600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005460a060020a900460ff16156103f057600080fd5b6103fa813461099a565b506104053382610acb565b50565b60005460a060020a900460ff1681565b60015460008054600160a060020a03928316923381169116148061044d575081600160a060020a031633600160a060020a0316145b151561045857600080fd5b81600160a060020a03166108fc30600160a060020a0316319081150290604051600060405180830381858888f150505050505050565b600081815260036020526040812081908190819081906104ad81610b21565b15156104b857600080fd5b80546001820154600290920154600160a060020a03909116986001608060020a038084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b6000805433600160a060020a0390811691161461053d57600080fd5b60005460a060020a900460ff161561055457600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b6000805460a060020a900460ff1615156105c357600080fd5b60005433600160a060020a039081169116146105de57600080fd5b5060008181526003602052604090206105f681610b21565b151561060157600080fd5b8054610617908390600160a060020a0316610b42565b5050565b600054600160a060020a031681565b60008181526003602052604081209061064282610b21565b151561064d57600080fd5b508054600160a060020a03908116903316811461066957600080fd5b6106738382610b42565b505050565b600081815260036020526040812061068f81610b21565b151561069a57600080fd5b6106a381610b8c565b9392505050565b600154600160a060020a031681565b60005433600160a060020a039081169116146106d457600080fd5b600160a060020a038116156104055760008054600160a060020a03831673ffffffffffffffffffffffffffffffffffffffff1990911617905550565b600154600090600160a060020a038085169116636352211e8460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561075f57600080fd5b5af1151561076c57600080fd5b50505060405180519050600160a060020a031614905092915050565b600154600160a060020a03166323b872dd83308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b15156107eb57600080fd5b5af115156107f857600080fd5b5050505050565b603c816060015167ffffffffffffffff16101561081b57600080fd5b600082815260036020526040902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560408201516001820180546001608060020a03928316700100000000000000000000000000000000029216919091179055606082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560808201516002909101805467ffffffffffffffff9290921668010000000000000000026fffffffffffffffff000000000000000019909216919091179055507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78260208301516001608060020a031683604001516001608060020a0316846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b600082815260036020526040812081808080806109b686610b21565b15156109c157600080fd5b6109ca86610b8c565b9450848810156109d957600080fd5b8554600160a060020a031693506109ef89610c13565b6000851115610a3957610a0185610c60565b92508285039150600160a060020a03841682156108fc0283604051600060405180830381858888f193505050501515610a3957600080fd5b50838703600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610a6e57600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd28986336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a15092979650505050505050565b600154600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b15156107eb57600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610b4b82610c13565b610b558183610acb565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df8260405190815260200160405180910390a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610bd25750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106a3916001608060020a0380821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610c6c565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610c8057869350610c9e565b878703925085858402811515610c9257fe5b05915081880190508093505b505050949350505050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a7230582012c4316600eb344490ef44ce513857e05c89e9b2f102d2e58d9517941e7502680029"

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
var ClockAuctionBaseBin = "0x6060604052341561000f57600080fd5b60f68061001d6000396000f30060606040526004361060485763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166383b5ff8b8114604d578063dd1b7a0f14606f575b600080fd5b3415605757600080fd5b605d60a8565b60405190815260200160405180910390f35b3415607957600080fd5b607f60ae565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b60015481565b60005473ffffffffffffffffffffffffffffffffffffffff16815600a165627a7a72305820fe3cea5d82af0f85fe74428c3c5bda702554c4149ace518dbca409d272de68e10029"

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

// ERC721MetadataABI is the input ABI used to generate the binding from.
const ERC721MetadataABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"\",\"type\":\"string\"}],\"name\":\"getMetadata\",\"outputs\":[{\"name\":\"buffer\",\"type\":\"bytes32[4]\"},{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"}]"

// ERC721MetadataFuncSigs maps the 4-byte function signature to its string representation.
var ERC721MetadataFuncSigs = map[string]string{
	"cb4799f2": "getMetadata(uint256,string)",
}

// ERC721MetadataBin is the compiled bytecode used for deploying new contracts.
var ERC721MetadataBin = "0x6060604052341561000f57600080fd5b6102708061001e6000396000f3006060604052600436106100405763ffffffff7c0100000000000000000000000000000000000000000000000000000000600035041663cb4799f28114610045575b600080fd5b341561005057600080fd5b61009b600480359060446024803590810190830135806020601f820181900481020160405190810160405281815292919060208401838380828437509496506100da95505050505050565b6040518083608080838360005b838110156100c05780820151838201526020016100a8565b505050509050018281526020019250505060405180910390f35b6100e261021b565b6000836001141561011857507f48656c6c6f20576f726c6421203a4400000000000000000000000000000000008152600f610214565b836002141561017257507f4920776f756c6420646566696e6974656c792063686f6f73652061206d65646981527f756d206c656e67746820737472696e672e00000000000000000000000000000060208201526031610214565b836003141561021457507f4c6f72656d20697073756d20646f6c6f722073697420616d65742c206d69206581527f737420616363756d73616e2064617069627573206175677565206c6f72656d2c60208201527f2074726973746971756520766573746962756c756d2069642c206c696265726f60408201527f207375736369706974207661726975732073617069656e20616c697175616d2e606082015260805b9250929050565b60806040519081016040526004815b6000815260001991909101906020018161022a57905050905600a165627a7a72305820ecdf48182e19d599c4eac88072add5e16a742b4144bcf661005b44bdeb469f340029"

// DeployERC721Metadata deploys a new Ethereum contract, binding an instance of ERC721Metadata to it.
func DeployERC721Metadata(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *ERC721Metadata, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(ERC721MetadataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// ERC721Metadata is an auto generated Go binding around an Ethereum contract.
type ERC721Metadata struct {
	ERC721MetadataCaller     // Read-only binding to the contract
	ERC721MetadataTransactor // Write-only binding to the contract
	ERC721MetadataFilterer   // Log filterer for contract events
}

// ERC721MetadataCaller is an auto generated read-only Go binding around an Ethereum contract.
type ERC721MetadataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ERC721MetadataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ERC721MetadataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ERC721MetadataSession struct {
	Contract     *ERC721Metadata   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ERC721MetadataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ERC721MetadataCallerSession struct {
	Contract *ERC721MetadataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// ERC721MetadataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ERC721MetadataTransactorSession struct {
	Contract     *ERC721MetadataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// ERC721MetadataRaw is an auto generated low-level Go binding around an Ethereum contract.
type ERC721MetadataRaw struct {
	Contract *ERC721Metadata // Generic contract binding to access the raw methods on
}

// ERC721MetadataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ERC721MetadataCallerRaw struct {
	Contract *ERC721MetadataCaller // Generic read-only contract binding to access the raw methods on
}

// ERC721MetadataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ERC721MetadataTransactorRaw struct {
	Contract *ERC721MetadataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewERC721Metadata creates a new instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721Metadata(address common.Address, backend bind.ContractBackend) (*ERC721Metadata, error) {
	contract, err := bindERC721Metadata(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ERC721Metadata{ERC721MetadataCaller: ERC721MetadataCaller{contract: contract}, ERC721MetadataTransactor: ERC721MetadataTransactor{contract: contract}, ERC721MetadataFilterer: ERC721MetadataFilterer{contract: contract}}, nil
}

// NewERC721MetadataCaller creates a new read-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataCaller(address common.Address, caller bind.ContractCaller) (*ERC721MetadataCaller, error) {
	contract, err := bindERC721Metadata(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataCaller{contract: contract}, nil
}

// NewERC721MetadataTransactor creates a new write-only instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataTransactor(address common.Address, transactor bind.ContractTransactor) (*ERC721MetadataTransactor, error) {
	contract, err := bindERC721Metadata(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataTransactor{contract: contract}, nil
}

// NewERC721MetadataFilterer creates a new log filterer instance of ERC721Metadata, bound to a specific deployed contract.
func NewERC721MetadataFilterer(address common.Address, filterer bind.ContractFilterer) (*ERC721MetadataFilterer, error) {
	contract, err := bindERC721Metadata(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ERC721MetadataFilterer{contract: contract}, nil
}

// bindERC721Metadata binds a generic wrapper to an already deployed contract.
func bindERC721Metadata(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ERC721MetadataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.ERC721MetadataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.ERC721MetadataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ERC721Metadata *ERC721MetadataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ERC721Metadata.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ERC721Metadata *ERC721MetadataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ERC721Metadata.Contract.contract.Transact(opts, method, params...)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 _tokenId, string ) view returns(bytes32[4] buffer, uint256 count)
func (_ERC721Metadata *ERC721MetadataCaller) GetMetadata(opts *bind.CallOpts, _tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	var out []interface{}
	err := _ERC721Metadata.contract.Call(opts, &out, "getMetadata", _tokenId, arg1)

	outstruct := new(struct {
		Buffer [4][32]byte
		Count  *big.Int
	})

	outstruct.Buffer = out[0].([4][32]byte)
	outstruct.Count = out[1].(*big.Int)

	return *outstruct, err

}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 _tokenId, string ) view returns(bytes32[4] buffer, uint256 count)
func (_ERC721Metadata *ERC721MetadataSession) GetMetadata(_tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	return _ERC721Metadata.Contract.GetMetadata(&_ERC721Metadata.CallOpts, _tokenId, arg1)
}

// GetMetadata is a free data retrieval call binding the contract method 0xcb4799f2.
//
// Solidity: function getMetadata(uint256 _tokenId, string ) view returns(bytes32[4] buffer, uint256 count)
func (_ERC721Metadata *ERC721MetadataCallerSession) GetMetadata(_tokenId *big.Int, arg1 string) (struct {
	Buffer [4][32]byte
	Count  *big.Int
}, error) {
	return _ERC721Metadata.Contract.GetMetadata(&_ERC721Metadata.CallOpts, _tokenId, arg1)
}

// GeneScienceInterfaceABI is the input ABI used to generate the binding from.
const GeneScienceInterfaceABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"genes1\",\"type\":\"uint256\"},{\"name\":\"genes2\",\"type\":\"uint256\"},{\"name\":\"targetBlock\",\"type\":\"uint256\"}],\"name\":\"mixGenes\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isGeneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"}]"

// GeneScienceInterfaceFuncSigs maps the 4-byte function signature to its string representation.
var GeneScienceInterfaceFuncSigs = map[string]string{
	"54c15b82": "isGeneScience()",
	"0d9f5aed": "mixGenes(uint256,uint256,uint256)",
}

// GeneScienceInterface is an auto generated Go binding around an Ethereum contract.
type GeneScienceInterface struct {
	GeneScienceInterfaceCaller     // Read-only binding to the contract
	GeneScienceInterfaceTransactor // Write-only binding to the contract
	GeneScienceInterfaceFilterer   // Log filterer for contract events
}

// GeneScienceInterfaceCaller is an auto generated read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceTransactor is an auto generated write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type GeneScienceInterfaceFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// GeneScienceInterfaceSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type GeneScienceInterfaceSession struct {
	Contract     *GeneScienceInterface // Generic contract binding to set the session for
	CallOpts     bind.CallOpts         // Call options to use throughout this session
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type GeneScienceInterfaceCallerSession struct {
	Contract *GeneScienceInterfaceCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts               // Call options to use throughout this session
}

// GeneScienceInterfaceTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type GeneScienceInterfaceTransactorSession struct {
	Contract     *GeneScienceInterfaceTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts               // Transaction auth options to use throughout this session
}

// GeneScienceInterfaceRaw is an auto generated low-level Go binding around an Ethereum contract.
type GeneScienceInterfaceRaw struct {
	Contract *GeneScienceInterface // Generic contract binding to access the raw methods on
}

// GeneScienceInterfaceCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type GeneScienceInterfaceCallerRaw struct {
	Contract *GeneScienceInterfaceCaller // Generic read-only contract binding to access the raw methods on
}

// GeneScienceInterfaceTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type GeneScienceInterfaceTransactorRaw struct {
	Contract *GeneScienceInterfaceTransactor // Generic write-only contract binding to access the raw methods on
}

// NewGeneScienceInterface creates a new instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterface(address common.Address, backend bind.ContractBackend) (*GeneScienceInterface, error) {
	contract, err := bindGeneScienceInterface(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterface{GeneScienceInterfaceCaller: GeneScienceInterfaceCaller{contract: contract}, GeneScienceInterfaceTransactor: GeneScienceInterfaceTransactor{contract: contract}, GeneScienceInterfaceFilterer: GeneScienceInterfaceFilterer{contract: contract}}, nil
}

// NewGeneScienceInterfaceCaller creates a new read-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceCaller(address common.Address, caller bind.ContractCaller) (*GeneScienceInterfaceCaller, error) {
	contract, err := bindGeneScienceInterface(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceCaller{contract: contract}, nil
}

// NewGeneScienceInterfaceTransactor creates a new write-only instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceTransactor(address common.Address, transactor bind.ContractTransactor) (*GeneScienceInterfaceTransactor, error) {
	contract, err := bindGeneScienceInterface(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceTransactor{contract: contract}, nil
}

// NewGeneScienceInterfaceFilterer creates a new log filterer instance of GeneScienceInterface, bound to a specific deployed contract.
func NewGeneScienceInterfaceFilterer(address common.Address, filterer bind.ContractFilterer) (*GeneScienceInterfaceFilterer, error) {
	contract, err := bindGeneScienceInterface(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &GeneScienceInterfaceFilterer{contract: contract}, nil
}

// bindGeneScienceInterface binds a generic wrapper to an already deployed contract.
func bindGeneScienceInterface(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(GeneScienceInterfaceABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.GeneScienceInterfaceTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_GeneScienceInterface *GeneScienceInterfaceCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _GeneScienceInterface.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_GeneScienceInterface *GeneScienceInterfaceTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.contract.Transact(opts, method, params...)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() pure returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCaller) IsGeneScience(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _GeneScienceInterface.contract.Call(opts, &out, "isGeneScience")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() pure returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// IsGeneScience is a free data retrieval call binding the contract method 0x54c15b82.
//
// Solidity: function isGeneScience() pure returns(bool)
func (_GeneScienceInterface *GeneScienceInterfaceCallerSession) IsGeneScience() (bool, error) {
	return _GeneScienceInterface.Contract.IsGeneScience(&_GeneScienceInterface.CallOpts)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(uint256 genes1, uint256 genes2, uint256 targetBlock) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactor) MixGenes(opts *bind.TransactOpts, genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.contract.Transact(opts, "mixGenes", genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(uint256 genes1, uint256 genes2, uint256 targetBlock) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}

// MixGenes is a paid mutator transaction binding the contract method 0x0d9f5aed.
//
// Solidity: function mixGenes(uint256 genes1, uint256 genes2, uint256 targetBlock) returns(uint256)
func (_GeneScienceInterface *GeneScienceInterfaceTransactorSession) MixGenes(genes1 *big.Int, genes2 *big.Int, targetBlock *big.Int) (*types.Transaction, error) {
	return _GeneScienceInterface.Contract.MixGenes(&_GeneScienceInterface.TransactOpts, genes1, genes2, targetBlock)
}

// KittyAccessControlABI is the input ABI used to generate the binding from.
const KittyAccessControlABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyAccessControlFuncSigs maps the 4-byte function signature to its string representation.
var KittyAccessControlFuncSigs = map[string]string{
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"3f4ba83a": "unpause()",
}

// KittyAccessControlBin is the compiled bytecode used for deploying new contracts.
var KittyAccessControlBin = "0x60606040526002805460a060020a60ff0219169055341561001f57600080fd5b6104478061002e6000396000f3006060604052600436106100985763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce79811461009d5780630a0f8168146100cc57806327d7874c146100df5780632ba73c15146101005780633f4ba83a1461011f5780634e0a3379146101325780635c975abb146101515780638456cb5914610178578063b047fb501461018b575b600080fd5b34156100a857600080fd5b6100b061019e565b604051600160a060020a03909116815260200160405180910390f35b34156100d757600080fd5b6100b06101ad565b34156100ea57600080fd5b6100fe600160a060020a03600435166101bc565b005b341561010b57600080fd5b6100fe600160a060020a036004351661021b565b341561012a57600080fd5b6100fe61027a565b341561013d57600080fd5b6100fe600160a060020a03600435166102de565b341561015c57600080fd5b61016461033d565b604051901515815260200160405180910390f35b341561018357600080fd5b6100fe61035e565b341561019657600080fd5b6100b061040c565b600154600160a060020a031681565b600054600160a060020a031681565b60005433600160a060020a039081169116146101d757600080fd5b600160a060020a03811615156101ec57600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461023657600080fd5b600160a060020a038116151561024b57600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461029557600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156102be57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b60005433600160a060020a039081169116146102f957600080fd5b600160a060020a038116151561030e57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025474010000000000000000000000000000000000000000900460ff1681565b60025433600160a060020a0390811691161480610389575060005433600160a060020a039081169116145b806103a2575060015433600160a060020a039081169116145b15156103ad57600080fd5b60025474010000000000000000000000000000000000000000900460ff16156103d557600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600254600160a060020a0316815600a165627a7a723058201647cca0924c8f20a3927b34a49ceb35d9ce3e905520c6d772384991a2af79600029"

// DeployKittyAccessControl deploys a new Ethereum contract, binding an instance of KittyAccessControl to it.
func DeployKittyAccessControl(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyAccessControl, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAccessControlABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyAccessControlBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyAccessControl{KittyAccessControlCaller: KittyAccessControlCaller{contract: contract}, KittyAccessControlTransactor: KittyAccessControlTransactor{contract: contract}, KittyAccessControlFilterer: KittyAccessControlFilterer{contract: contract}}, nil
}

// KittyAccessControl is an auto generated Go binding around an Ethereum contract.
type KittyAccessControl struct {
	KittyAccessControlCaller     // Read-only binding to the contract
	KittyAccessControlTransactor // Write-only binding to the contract
	KittyAccessControlFilterer   // Log filterer for contract events
}

// KittyAccessControlCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyAccessControlCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyAccessControlTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyAccessControlFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAccessControlSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyAccessControlSession struct {
	Contract     *KittyAccessControl // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// KittyAccessControlCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyAccessControlCallerSession struct {
	Contract *KittyAccessControlCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// KittyAccessControlTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyAccessControlTransactorSession struct {
	Contract     *KittyAccessControlTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// KittyAccessControlRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyAccessControlRaw struct {
	Contract *KittyAccessControl // Generic contract binding to access the raw methods on
}

// KittyAccessControlCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyAccessControlCallerRaw struct {
	Contract *KittyAccessControlCaller // Generic read-only contract binding to access the raw methods on
}

// KittyAccessControlTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyAccessControlTransactorRaw struct {
	Contract *KittyAccessControlTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyAccessControl creates a new instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControl(address common.Address, backend bind.ContractBackend) (*KittyAccessControl, error) {
	contract, err := bindKittyAccessControl(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControl{KittyAccessControlCaller: KittyAccessControlCaller{contract: contract}, KittyAccessControlTransactor: KittyAccessControlTransactor{contract: contract}, KittyAccessControlFilterer: KittyAccessControlFilterer{contract: contract}}, nil
}

// NewKittyAccessControlCaller creates a new read-only instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlCaller(address common.Address, caller bind.ContractCaller) (*KittyAccessControlCaller, error) {
	contract, err := bindKittyAccessControl(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlCaller{contract: contract}, nil
}

// NewKittyAccessControlTransactor creates a new write-only instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyAccessControlTransactor, error) {
	contract, err := bindKittyAccessControl(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlTransactor{contract: contract}, nil
}

// NewKittyAccessControlFilterer creates a new log filterer instance of KittyAccessControl, bound to a specific deployed contract.
func NewKittyAccessControlFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyAccessControlFilterer, error) {
	contract, err := bindKittyAccessControl(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlFilterer{contract: contract}, nil
}

// bindKittyAccessControl binds a generic wrapper to an already deployed contract.
func bindKittyAccessControl(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAccessControlABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAccessControl *KittyAccessControlRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyAccessControl.Contract.KittyAccessControlCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAccessControl *KittyAccessControlRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.KittyAccessControlTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAccessControl *KittyAccessControlRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.KittyAccessControlTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAccessControl *KittyAccessControlCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyAccessControl.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAccessControl *KittyAccessControlTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAccessControl *KittyAccessControlTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAccessControl.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CeoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CeoAddress(&_KittyAccessControl.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CeoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CeoAddress(&_KittyAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAccessControl.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CfoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CfoAddress(&_KittyAccessControl.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CfoAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CfoAddress(&_KittyAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAccessControl.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlSession) CooAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CooAddress(&_KittyAccessControl.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAccessControl *KittyAccessControlCallerSession) CooAddress() (common.Address, error) {
	return _KittyAccessControl.Contract.CooAddress(&_KittyAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAccessControl *KittyAccessControlCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyAccessControl.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAccessControl *KittyAccessControlSession) Paused() (bool, error) {
	return _KittyAccessControl.Contract.Paused(&_KittyAccessControl.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAccessControl *KittyAccessControlCallerSession) Paused() (bool, error) {
	return _KittyAccessControl.Contract.Paused(&_KittyAccessControl.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlSession) Pause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Pause(&_KittyAccessControl.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Pause(&_KittyAccessControl.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCEO(&_KittyAccessControl.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCEO(&_KittyAccessControl.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCFO(&_KittyAccessControl.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCFO(&_KittyAccessControl.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAccessControl *KittyAccessControlTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAccessControl *KittyAccessControlSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCOO(&_KittyAccessControl.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAccessControl.Contract.SetCOO(&_KittyAccessControl.TransactOpts, _newCOO)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAccessControl.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlSession) Unpause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Unpause(&_KittyAccessControl.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAccessControl *KittyAccessControlTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyAccessControl.Contract.Unpause(&_KittyAccessControl.TransactOpts)
}

// KittyAccessControlContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyAccessControl contract.
type KittyAccessControlContractUpgradeIterator struct {
	Event *KittyAccessControlContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyAccessControlContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAccessControlContractUpgrade)
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
		it.Event = new(KittyAccessControlContractUpgrade)
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
func (it *KittyAccessControlContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAccessControlContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAccessControlContractUpgrade represents a ContractUpgrade event raised by the KittyAccessControl contract.
type KittyAccessControlContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAccessControl *KittyAccessControlFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyAccessControlContractUpgradeIterator, error) {

	logs, sub, err := _KittyAccessControl.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyAccessControlContractUpgradeIterator{contract: _KittyAccessControl.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAccessControl *KittyAccessControlFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyAccessControlContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyAccessControl.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAccessControlContractUpgrade)
				if err := _KittyAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAccessControl *KittyAccessControlFilterer) ParseContractUpgrade(log types.Log) (*KittyAccessControlContractUpgrade, error) {
	event := new(KittyAccessControlContractUpgrade)
	if err := _KittyAccessControl.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyAuctionABI is the input ABI used to generate the binding from.
const KittyAuctionABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyAuctionFuncSigs maps the 4-byte function signature to its string representation.
var KittyAuctionFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"4dfff04f": "approveSiring(address,uint256)",
	"b0c35c05": "autoBirthFee()",
	"70a08231": "balanceOf(address)",
	"ed60ade6": "bidOnSiringAuction(uint256,uint256)",
	"f7d8c883": "breedWithAuto(uint256,uint256)",
	"46d22c70": "canBreedWith(uint256,uint256)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"3d7d3f5a": "createSaleAuction(uint256,uint256,uint256,uint256)",
	"4ad8c938": "createSiringAuction(uint256,uint256,uint256,uint256)",
	"bc4006f5": "erc721Metadata()",
	"f2b47d52": "geneScience()",
	"88c2a0bf": "giveBirth(uint256)",
	"1940a936": "isPregnant(uint256)",
	"d3e6f49f": "isReadyToBreed(uint256)",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"183a7947": "pregnantKitties()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"4b85fd55": "setAutoBirthFee(uint256)",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"24e7a38a": "setGeneScienceAddress(address)",
	"e17b25af": "setMetadataAddress(address)",
	"6fbde40d": "setSaleAuctionAddress(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"14001f4c": "setSiringAuctionAddress(address)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"0560ff44": "tokenMetadata(uint256,string)",
	"8462151c": "tokensOfOwner(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
	"91876e57": "withdrawAuctionBalances()",
}

// KittyAuctionBin is the compiled bytecode used for deploying new contracts.
var KittyAuctionBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a0820152620000a790600390600e620000ca565b50600f60055566071afd498d0000600e553415620000c457600080fd5b62000194565b6002830191839082156200015b5791602002820160005b838211156200012757835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000e1565b8015620001595782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000127565b505b50620001699291506200016d565b5090565b6200019191905b808211156200016957805463ffffffff1916815560010162000174565b90565b6128ef80620001a46000396000f3006060604052600436106102215763ffffffff60e060020a60003504166301ffc9a781146102265780630519ce79146102725780630560ff44146102a157806306fdde031461033a578063095ea7b31461034d5780630a0f81681461037157806314001f4c1461038457806318160ddd146103a3578063183a7947146103c85780631940a936146103db57806321717ebf146103f157806323b872dd1461040457806324e7a38a1461042c57806327d7874c1461044b5780632ba73c151461046a5780633d7d3f5a146104895780633f4ba83a146104a857806346116e6f146104bb57806346d22c70146104d1578063481af3d3146104ea5780634ad8c938146105005780634b85fd551461051f5780634dfff04f146105355780634e0a3379146105575780635663896e146105765780635c975abb1461058c5780636352211e1461059f5780636fbde40d146105b557806370a08231146105d45780637a7d4937146105f35780638456cb59146106065780638462151c1461061957806388c2a0bf1461068b57806391876e57146106a157806395d89b41146106b45780639d6fac6f146106c7578063a45f4bfc146106f6578063a9059cbb1461070c578063b047fb501461072e578063b0c35c0514610741578063bc4006f514610754578063d3e6f49f14610767578063e17b25af1461077d578063e6cbe3511461079c578063ed60ade6146107af578063f2b47d52146107bd578063f7d8c883146107d0575b600080fd5b341561023157600080fd5b61025e7fffffffff00000000000000000000000000000000000000000000000000000000600435166107de565b604051901515815260200160405180910390f35b341561027d57600080fd5b610285610a65565b604051600160a060020a03909116815260200160405180910390f35b34156102ac57600080fd5b6102c3600480359060248035908101910135610a74565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102ff5780820151838201526020016102e7565b50505050905090810190601f16801561032c5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561034557600080fd5b6102c3610b44565b341561035857600080fd5b61036f600160a060020a0360043516602435610b7b565b005b341561037c57600080fd5b610285610c05565b341561038f57600080fd5b61036f600160a060020a0360043516610c14565b34156103ae57600080fd5b6103b6610cb4565b60405190815260200160405180910390f35b34156103d357600080fd5b6103b6610cbf565b34156103e657600080fd5b61025e600435610cc5565b34156103fc57600080fd5b610285610d0a565b341561040f57600080fd5b61036f600160a060020a0360043581169060243516604435610d19565b341561043757600080fd5b61036f600160a060020a0360043516610da0565b341561045657600080fd5b61036f600160a060020a0360043516610e40565b341561047557600080fd5b61036f600160a060020a0360043516610e92565b341561049457600080fd5b61036f600435602435604435606435610ee4565b34156104b357600080fd5b61036f610fbf565b34156104c657600080fd5b610285600435611012565b34156104dc57600080fd5b61025e60043560243561102d565b34156104f557600080fd5b6102856004356110ad565b341561050b57600080fd5b61036f6004356024356044356064356110c8565b341561052a57600080fd5b61036f60043561118e565b341561054057600080fd5b61036f600160a060020a03600435166024356111ae565b341561056257600080fd5b61036f600160a060020a0360043516611208565b341561058157600080fd5b61036f60043561125a565b341561059757600080fd5b61025e6112c2565b34156105aa57600080fd5b6102856004356112d2565b34156105c057600080fd5b61036f600160a060020a03600435166112f6565b34156105df57600080fd5b6103b6600160a060020a0360043516611396565b34156105fe57600080fd5b6103b66113b1565b341561061157600080fd5b61036f6113b7565b341561062457600080fd5b610638600160a060020a0360043516611443565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561067757808201518382015260200161065f565b505050509050019250505060405180910390f35b341561069657600080fd5b6103b6600435611524565b34156106ac57600080fd5b61036f6117e7565b34156106bf57600080fd5b6102c36118d2565b34156106d257600080fd5b6106dd600435611909565b60405163ffffffff909116815260200160405180910390f35b341561070157600080fd5b610285600435611936565b341561071757600080fd5b61036f600160a060020a0360043516602435611951565b341561073957600080fd5b6102856119f8565b341561074c57600080fd5b6103b6611a07565b341561075f57600080fd5b610285611a0d565b341561077257600080fd5b61025e600435611a1c565b341561078857600080fd5b61036f600160a060020a0360043516611ae5565b34156107a757600080fd5b610285611b22565b61036f600435602435611b31565b34156107c857600080fd5b610285611c73565b61036f600435602435611c82565b60006040517f737570706f727473496e7465726661636528627974657334290000000000000081526019016040518091039020600160e060020a03191682600160e060020a0319161480610a5d57506040517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152601d0160405180910390206040517f746f6b656e734f664f776e657228616464726573732900000000000000000000815260160160405180910390206040517f7472616e7366657246726f6d28616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390206040517f7472616e7366657228616464726573732c75696e743235362900000000000000815260190160405180910390206040517f617070726f766528616464726573732c75696e74323536290000000000000000815260180160405180910390206040517f6f776e65724f662875696e743235362900000000000000000000000000000000815260100160405180910390206040517f62616c616e63654f662861646472657373290000000000000000000000000000815260120160405180910390206040517f746f74616c537570706c792829000000000000000000000000000000000000008152600d0160405180910390206040517f73796d626f6c2829000000000000000000000000000000000000000000000000815260080160405180910390206040517f6e616d652829000000000000000000000000000000000000000000000000000081526006016040518091039020181818181818181818600160e060020a03191682600160e060020a031916145b90505b919050565b600154600160a060020a031681565b610a7c6127fc565b610a8461280e565b600d54600090600160a060020a03161515610a9e57600080fd5b600d54600160a060020a031663cb4799f287878760405160e060020a63ffffffff861602815260048101848152604060248301908152604483018490529091606401848480828437820191505094505050505060a060405180830381600087803b1515610b0a57600080fd5b5af11515610b1757600080fd5b50505060405180608001805160209091016040529092509050610b3a8282611e7a565b9695505050505050565b60408051908101604052600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610b9257600080fd5b610b9c3382611ecf565b1515610ba757600080fd5b610bb18183611eef565b7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600054600160a060020a031681565b6000805433600160a060020a03908116911614610c3057600080fd5b5080600160a060020a0381166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610c6f57600080fd5b5af11515610c7c57600080fd5b505050604051805190501515610c9157600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b600654600019015b90565b600f5481565b6000808211610cd357600080fd5b6006805483908110610ce157fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b600c54600160a060020a031681565b60025460a060020a900460ff1615610d3057600080fd5b600160a060020a0382161515610d4557600080fd5b30600160a060020a031682600160a060020a031614151515610d6657600080fd5b610d703382611f1d565b1515610d7b57600080fd5b610d858382611ecf565b1515610d9057600080fd5b610d9b838383611f3d565b505050565b6000805433600160a060020a03908116911614610dbc57600080fd5b5080600160a060020a0381166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610dfb57600080fd5b5af11515610e0857600080fd5b505050604051805190501515610e1d57600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b60005433600160a060020a03908116911614610e5b57600080fd5b600160a060020a0381161515610e7057600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610ead57600080fd5b600160a060020a0381161515610ec257600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff1615610efb57600080fd5b610f053385611ecf565b1515610f1057600080fd5b610f1984610cc5565b15610f2357600080fd5b600b54610f3a908590600160a060020a0316611eef565b600b54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b1515610fa957600080fd5b5af11515610fb657600080fd5b50505050505050565b60005433600160a060020a03908116911614610fda57600080fd5b60025460a060020a900460ff161515610ff257600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b6000808080851161103d57600080fd5b6000841161104a57600080fd5b600680548690811061105857fe5b9060005260206000209060020201915060068481548110151561107757fe5b9060005260206000209060020201905061109382868387612025565b80156110a457506110a484866121a5565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff16156110df57600080fd5b6110e93385611ecf565b15156110f457600080fd5b6110fd84611a1c565b151561110857600080fd5b600c5461111f908590600160a060020a0316611eef565b600c54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b1515610fa957600080fd5b60025433600160a060020a039081169116146111a957600080fd5b600e55565b60025460a060020a900460ff16156111c557600080fd5b6111cf3382611ecf565b15156111da57600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461122357600080fd5b600160a060020a038116151561123857600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60025433600160a060020a0390811691161480611285575060005433600160a060020a039081169116145b8061129e575060015433600160a060020a039081169116145b15156112a957600080fd5b60035463ffffffff1681106112bd57600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a0316801515610a6057600080fd5b6000805433600160a060020a0390811691161461131257600080fd5b5080600160a060020a0381166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561135157600080fd5b5af1151561135e57600080fd5b50505060405180519050151561137357600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b60055481565b60025433600160a060020a03908116911614806113e2575060005433600160a060020a039081169116145b806113fb575060015433600160a060020a039081169116145b151561140657600080fd5b60025460a060020a900460ff161561141d57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b61144b6127fc565b60006114556127fc565b600080600061146387611396565b945084151561149357600060405180591061147b5750595b9080825280602002602001820160405250955061151a565b846040518059106114a15750595b908082528060200260200182016040525093506114bc610cb4565b925060009150600190505b82811161151657600081815260076020526040902054600160a060020a038881169116141561150e57808483815181106114fd57fe5b602090810290910101526001909101905b6001016114c7565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff1615151561154c57600080fd5b600680548a90811061155a57fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561158757600080fd5b61161c8761010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526121fa565b151561162757600080fd5b60018701546006805460c060020a90920463ffffffff169750908790811061164b57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a9283900481169650919004168490111561169957600185015460f060020a900461ffff1693505b6010548754865460018a0154600160a060020a0390931692630d9f5aed92919068010000000000000000900467ffffffffffffffff166000190160405160e060020a63ffffffff86160281526004810193909352602483019190915267ffffffffffffffff166044820152606401602060405180830381600087803b151561172057600080fd5b5af1151561172d57600080fd5b505050604051805160008b81526007602052604090205460018a810154929650600160a060020a03909116945061177c92508b9160c060020a900463ffffffff1690870161ffff168686612232565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54909150600160a060020a0333169080156108fc0290604051600060405180830381858888f150939c9b505050505050505050505050565b60025433600160a060020a0390811691161480611812575060005433600160a060020a039081169116145b8061182b575060015433600160a060020a039081169116145b151561183657600080fd5b600b54600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b151561187557600080fd5b5af1151561188257600080fd5b5050600c54600160a060020a03169050635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b15156118c557600080fd5b5af11515610d9b57600080fd5b60408051908101604052600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e811061191657fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff161561196857600080fd5b600160a060020a038216151561197d57600080fd5b30600160a060020a031682600160a060020a03161415151561199e57600080fd5b600b54600160a060020a03838116911614156119b957600080fd5b600c54600160a060020a03838116911614156119d457600080fd5b6119de3382611ecf565b15156119e957600080fd5b6119f4338383611f3d565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b600080808311611a2b57600080fd5b6006805484908110611a3957fe5b90600052602060002090600202019050611ade8161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526124de565b9392505050565b60005433600160a060020a03908116911614611b0057600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b60025460009060a060020a900460ff1615611b4b57600080fd5b611b553383611ecf565b1515611b6057600080fd5b611b6982611a1c565b1515611b7457600080fd5b611b7e8284612515565b1515611b8957600080fd5b600c54600160a060020a031663c55d0f568460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515611bd157600080fd5b5af11515611bde57600080fd5b5050506040518051600e5490925082013410159050611bfc57600080fd5b600c54600e54600160a060020a039091169063454a2ab39034038560405160e060020a63ffffffff851602815260048101919091526024016000604051808303818588803b1515611c4c57600080fd5b5af11515611c5957600080fd5b50505050610d9b8263ffffffff168463ffffffff16612564565b601054600160a060020a031681565b600254600090819060a060020a900460ff1615611c9e57600080fd5b600e54341015611cad57600080fd5b611cb73385611ecf565b1515611cc257600080fd5b611ccc83856121a5565b1515611cd757600080fd5b6006805485908110611ce557fe5b90600052602060002090600202019150611d8a8261010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526124de565b1515611d9557600080fd5b6006805484908110611da357fe5b90600052602060002090600202019050611e488161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e08201526124de565b1515611e5357600080fd5b611e5f82858386612025565b1515611e6a57600080fd5b611e748484612564565b50505050565b611e826127fc565b611e8a6127fc565b60008084604051805910611e9b5750595b818152601f19601f8301168101602001604052905092505060208201905084611ec58282876126ce565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a0319169091179055831615611fd057600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b6000818414156120375750600061219d565b6001850154608060020a900463ffffffff168214806120665750600185015460a060020a900463ffffffff1682145b156120735750600061219d565b6001830154608060020a900463ffffffff168414806120a25750600183015460a060020a900463ffffffff1684145b156120af5750600061219d565b6001830154608060020a900463ffffffff1615806120dc57506001850154608060020a900463ffffffff16155b156120e95750600161219d565b60018581015490840154608060020a9182900463ffffffff90811692909104161480612134575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b156121415750600061219d565b6001808601549084015460a060020a900463ffffffff908116608060020a90920416148061218c57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b156121995750600061219d565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a039182169116808214806110a457506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b60008160a0015163ffffffff1615801590610a5d57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b60008061223d612837565b600063ffffffff8916891461225157600080fd5b63ffffffff8816881461226357600080fd5b61ffff8716871461227357600080fd5b600287049250600d8361ffff16111561228b57600d92505b610100604051908101604090815287825267ffffffffffffffff42166020830152600090820181905263ffffffff808c1660608401528a16608083015260a082015261ffff80851660c0830152881660e0820152600680549193506001918083016122f6838261287b565b6000928352602090922085916002020181518155602082015160018201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555060c082015181600101601c6101000a81548161ffff021916908361ffff16021790555060e08201516001909101805461ffff9290921660f060020a027dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555003905063ffffffff8116811461245157600080fd5b7f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad58582846060015163ffffffff16856080015163ffffffff168651604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390a16124d260008683611f3d565b98975050505050505050565b60008160a0015163ffffffff16158015610a5d57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600080600060068581548110151561252957fe5b9060005260206000209060020201915060068481548110151561254857fe5b906000526020600020906002020190506110a482868387612025565b60008060068381548110151561257657fe5b9060005260206000209060020201915060068481548110151561259557fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff87160217905590506125e582612713565b6125ee81612713565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f8054600190810190915587845260079092529182902054908301547f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b8092600160a060020a0390921691879187916801000000000000000090910467ffffffffffffffff1690518085600160a060020a0316600160a060020a031681526020018481526020018381526020018267ffffffffffffffff16815260200194505050505060405180910390a150505050565b60005b602082106126f457825184526020840193506020830192506020820391506126d1565b6001826020036101000a03905080198351168185511617909352505050565b600554600182015443919060039060e060020a900461ffff16600e811061273657fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff1681151561276157fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff1610156127f9576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b60206040519081016040526000815290565b60806040519081016040526004815b6000815260001991909101906020018161281d5790505090565b6101006040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c0820181905260e082015290565b815481835581811511610d9b57600083815260209020610d9b91610cbc9160029182028101918502015b808211156128bf57600080825560018201556002016128a5565b50905600a165627a7a72305820d1a173ff405af79ee52b6fcb9afc5d0fe61dc5f26d44135dac5db708d0672aaf0029"

// DeployKittyAuction deploys a new Ethereum contract, binding an instance of KittyAuction to it.
func DeployKittyAuction(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyAuctionBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyAuction{KittyAuctionCaller: KittyAuctionCaller{contract: contract}, KittyAuctionTransactor: KittyAuctionTransactor{contract: contract}, KittyAuctionFilterer: KittyAuctionFilterer{contract: contract}}, nil
}

// KittyAuction is an auto generated Go binding around an Ethereum contract.
type KittyAuction struct {
	KittyAuctionCaller     // Read-only binding to the contract
	KittyAuctionTransactor // Write-only binding to the contract
	KittyAuctionFilterer   // Log filterer for contract events
}

// KittyAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyAuctionSession struct {
	Contract     *KittyAuction     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyAuctionCallerSession struct {
	Contract *KittyAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KittyAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyAuctionTransactorSession struct {
	Contract     *KittyAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KittyAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyAuctionRaw struct {
	Contract *KittyAuction // Generic contract binding to access the raw methods on
}

// KittyAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyAuctionCallerRaw struct {
	Contract *KittyAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// KittyAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyAuctionTransactorRaw struct {
	Contract *KittyAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyAuction creates a new instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuction(address common.Address, backend bind.ContractBackend) (*KittyAuction, error) {
	contract, err := bindKittyAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyAuction{KittyAuctionCaller: KittyAuctionCaller{contract: contract}, KittyAuctionTransactor: KittyAuctionTransactor{contract: contract}, KittyAuctionFilterer: KittyAuctionFilterer{contract: contract}}, nil
}

// NewKittyAuctionCaller creates a new read-only instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionCaller(address common.Address, caller bind.ContractCaller) (*KittyAuctionCaller, error) {
	contract, err := bindKittyAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionCaller{contract: contract}, nil
}

// NewKittyAuctionTransactor creates a new write-only instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyAuctionTransactor, error) {
	contract, err := bindKittyAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionTransactor{contract: contract}, nil
}

// NewKittyAuctionFilterer creates a new log filterer instance of KittyAuction, bound to a specific deployed contract.
func NewKittyAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyAuctionFilterer, error) {
	contract, err := bindKittyAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyAuctionFilterer{contract: contract}, nil
}

// bindKittyAuction binds a generic wrapper to an already deployed contract.
func bindKittyAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAuction *KittyAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyAuction.Contract.KittyAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAuction *KittyAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.Contract.KittyAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAuction *KittyAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAuction.Contract.KittyAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyAuction *KittyAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyAuction *KittyAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyAuction *KittyAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyAuction.Contract.contract.Transact(opts, method, params...)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyAuction *KittyAuctionCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "autoBirthFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyAuction *KittyAuctionSession) AutoBirthFee() (*big.Int, error) {
	return _KittyAuction.Contract.AutoBirthFee(&_KittyAuction.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyAuction.Contract.AutoBirthFee(&_KittyAuction.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyAuction *KittyAuctionCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyAuction *KittyAuctionSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyAuction.Contract.BalanceOf(&_KittyAuction.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyAuction *KittyAuctionCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyAuction.Contract.BalanceOf(&_KittyAuction.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyAuction *KittyAuctionCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "canBreedWith", _matronId, _sireId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyAuction *KittyAuctionSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyAuction.Contract.CanBreedWith(&_KittyAuction.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyAuction.Contract.CanBreedWith(&_KittyAuction.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAuction *KittyAuctionCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAuction *KittyAuctionSession) CeoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CeoAddress(&_KittyAuction.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CeoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CeoAddress(&_KittyAuction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAuction *KittyAuctionCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAuction *KittyAuctionSession) CfoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CfoAddress(&_KittyAuction.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CfoAddress() (common.Address, error) {
	return _KittyAuction.Contract.CfoAddress(&_KittyAuction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAuction *KittyAuctionCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAuction *KittyAuctionSession) CooAddress() (common.Address, error) {
	return _KittyAuction.Contract.CooAddress(&_KittyAuction.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) CooAddress() (common.Address, error) {
	return _KittyAuction.Contract.CooAddress(&_KittyAuction.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyAuction *KittyAuctionCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyAuction *KittyAuctionSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyAuction.Contract.Cooldowns(&_KittyAuction.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyAuction *KittyAuctionCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyAuction.Contract.Cooldowns(&_KittyAuction.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyAuction *KittyAuctionCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyAuction *KittyAuctionSession) Erc721Metadata() (common.Address, error) {
	return _KittyAuction.Contract.Erc721Metadata(&_KittyAuction.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyAuction.Contract.Erc721Metadata(&_KittyAuction.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyAuction *KittyAuctionCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyAuction *KittyAuctionSession) GeneScience() (common.Address, error) {
	return _KittyAuction.Contract.GeneScience(&_KittyAuction.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) GeneScience() (common.Address, error) {
	return _KittyAuction.Contract.GeneScience(&_KittyAuction.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "isPregnant", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsPregnant(&_KittyAuction.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsPregnant(&_KittyAuction.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "isReadyToBreed", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsReadyToBreed(&_KittyAuction.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyAuction.Contract.IsReadyToBreed(&_KittyAuction.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToApproved(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToApproved(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToOwner(&_KittyAuction.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.KittyIndexToOwner(&_KittyAuction.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyAuction *KittyAuctionCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyAuction *KittyAuctionSession) Name() (string, error) {
	return _KittyAuction.Contract.Name(&_KittyAuction.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyAuction *KittyAuctionCallerSession) Name() (string, error) {
	return _KittyAuction.Contract.Name(&_KittyAuction.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyAuction *KittyAuctionCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyAuction *KittyAuctionSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.OwnerOf(&_KittyAuction.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyAuction *KittyAuctionCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.OwnerOf(&_KittyAuction.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAuction *KittyAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAuction *KittyAuctionSession) Paused() (bool, error) {
	return _KittyAuction.Contract.Paused(&_KittyAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) Paused() (bool, error) {
	return _KittyAuction.Contract.Paused(&_KittyAuction.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyAuction *KittyAuctionCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "pregnantKitties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyAuction *KittyAuctionSession) PregnantKitties() (*big.Int, error) {
	return _KittyAuction.Contract.PregnantKitties(&_KittyAuction.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyAuction.Contract.PregnantKitties(&_KittyAuction.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyAuction *KittyAuctionCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyAuction *KittyAuctionSession) SaleAuction() (common.Address, error) {
	return _KittyAuction.Contract.SaleAuction(&_KittyAuction.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SaleAuction() (common.Address, error) {
	return _KittyAuction.Contract.SaleAuction(&_KittyAuction.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyAuction *KittyAuctionCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyAuction *KittyAuctionSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyAuction.Contract.SecondsPerBlock(&_KittyAuction.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyAuction.Contract.SecondsPerBlock(&_KittyAuction.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.SireAllowedToAddress(&_KittyAuction.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyAuction.Contract.SireAllowedToAddress(&_KittyAuction.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyAuction *KittyAuctionCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyAuction *KittyAuctionSession) SiringAuction() (common.Address, error) {
	return _KittyAuction.Contract.SiringAuction(&_KittyAuction.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyAuction *KittyAuctionCallerSession) SiringAuction() (common.Address, error) {
	return _KittyAuction.Contract.SiringAuction(&_KittyAuction.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyAuction *KittyAuctionCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyAuction *KittyAuctionSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyAuction.Contract.SupportsInterface(&_KittyAuction.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyAuction *KittyAuctionCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyAuction.Contract.SupportsInterface(&_KittyAuction.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyAuction *KittyAuctionCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyAuction *KittyAuctionSession) Symbol() (string, error) {
	return _KittyAuction.Contract.Symbol(&_KittyAuction.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyAuction *KittyAuctionCallerSession) Symbol() (string, error) {
	return _KittyAuction.Contract.Symbol(&_KittyAuction.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyAuction *KittyAuctionCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyAuction *KittyAuctionSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyAuction.Contract.TokenMetadata(&_KittyAuction.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyAuction *KittyAuctionCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyAuction.Contract.TokenMetadata(&_KittyAuction.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyAuction *KittyAuctionCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyAuction *KittyAuctionSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyAuction.Contract.TokensOfOwner(&_KittyAuction.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyAuction *KittyAuctionCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyAuction.Contract.TokensOfOwner(&_KittyAuction.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyAuction *KittyAuctionCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyAuction.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyAuction *KittyAuctionSession) TotalSupply() (*big.Int, error) {
	return _KittyAuction.Contract.TotalSupply(&_KittyAuction.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyAuction *KittyAuctionCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyAuction.Contract.TotalSupply(&_KittyAuction.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Approve(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Approve(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyAuction *KittyAuctionTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyAuction *KittyAuctionSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.ApproveSiring(&_KittyAuction.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyAuction *KittyAuctionTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.ApproveSiring(&_KittyAuction.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyAuction *KittyAuctionTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyAuction *KittyAuctionSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BidOnSiringAuction(&_KittyAuction.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyAuction *KittyAuctionTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BidOnSiringAuction(&_KittyAuction.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyAuction *KittyAuctionTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyAuction *KittyAuctionSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BreedWithAuto(&_KittyAuction.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyAuction *KittyAuctionTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.BreedWithAuto(&_KittyAuction.TransactOpts, _matronId, _sireId)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSaleAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSaleAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSiringAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyAuction *KittyAuctionTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.CreateSiringAuction(&_KittyAuction.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyAuction *KittyAuctionTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyAuction *KittyAuctionSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.GiveBirth(&_KittyAuction.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyAuction *KittyAuctionTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.GiveBirth(&_KittyAuction.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionSession) Pause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Pause(&_KittyAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyAuction *KittyAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Pause(&_KittyAuction.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyAuction *KittyAuctionTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyAuction *KittyAuctionSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetAutoBirthFee(&_KittyAuction.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetAutoBirthFee(&_KittyAuction.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAuction *KittyAuctionSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCEO(&_KittyAuction.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCEO(&_KittyAuction.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAuction *KittyAuctionSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCFO(&_KittyAuction.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCFO(&_KittyAuction.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAuction *KittyAuctionTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAuction *KittyAuctionSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCOO(&_KittyAuction.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetCOO(&_KittyAuction.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyAuction *KittyAuctionSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetGeneScienceAddress(&_KittyAuction.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetGeneScienceAddress(&_KittyAuction.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyAuction *KittyAuctionTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyAuction *KittyAuctionSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetMetadataAddress(&_KittyAuction.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetMetadataAddress(&_KittyAuction.TransactOpts, _contractAddress)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSaleAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSaleAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyAuction *KittyAuctionSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSecondsPerBlock(&_KittyAuction.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSecondsPerBlock(&_KittyAuction.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSiringAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyAuction *KittyAuctionTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyAuction.Contract.SetSiringAuctionAddress(&_KittyAuction.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Transfer(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.Transfer(&_KittyAuction.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.TransferFrom(&_KittyAuction.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyAuction *KittyAuctionTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyAuction.Contract.TransferFrom(&_KittyAuction.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionSession) Unpause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Unpause(&_KittyAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyAuction *KittyAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyAuction.Contract.Unpause(&_KittyAuction.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyAuction.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyAuction.Contract.WithdrawAuctionBalances(&_KittyAuction.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyAuction *KittyAuctionTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyAuction.Contract.WithdrawAuctionBalances(&_KittyAuction.TransactOpts)
}

// KittyAuctionApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyAuction contract.
type KittyAuctionApprovalIterator struct {
	Event *KittyAuctionApproval // Event containing the contract specifics and raw log

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
func (it *KittyAuctionApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionApproval)
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
		it.Event = new(KittyAuctionApproval)
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
func (it *KittyAuctionApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionApproval represents a Approval event raised by the KittyAuction contract.
type KittyAuctionApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyAuction *KittyAuctionFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyAuctionApprovalIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionApprovalIterator{contract: _KittyAuction.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyAuction *KittyAuctionFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyAuctionApproval) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionApproval)
				if err := _KittyAuction.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KittyAuction *KittyAuctionFilterer) ParseApproval(log types.Log) (*KittyAuctionApproval, error) {
	event := new(KittyAuctionApproval)
	if err := _KittyAuction.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyAuctionBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyAuction contract.
type KittyAuctionBirthIterator struct {
	Event *KittyAuctionBirth // Event containing the contract specifics and raw log

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
func (it *KittyAuctionBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionBirth)
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
		it.Event = new(KittyAuctionBirth)
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
func (it *KittyAuctionBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionBirth represents a Birth event raised by the KittyAuction contract.
type KittyAuctionBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyAuction *KittyAuctionFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyAuctionBirthIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionBirthIterator{contract: _KittyAuction.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyAuction *KittyAuctionFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyAuctionBirth) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionBirth)
				if err := _KittyAuction.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyAuction *KittyAuctionFilterer) ParseBirth(log types.Log) (*KittyAuctionBirth, error) {
	event := new(KittyAuctionBirth)
	if err := _KittyAuction.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyAuctionContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyAuction contract.
type KittyAuctionContractUpgradeIterator struct {
	Event *KittyAuctionContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyAuctionContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionContractUpgrade)
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
		it.Event = new(KittyAuctionContractUpgrade)
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
func (it *KittyAuctionContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionContractUpgrade represents a ContractUpgrade event raised by the KittyAuction contract.
type KittyAuctionContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAuction *KittyAuctionFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyAuctionContractUpgradeIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionContractUpgradeIterator{contract: _KittyAuction.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAuction *KittyAuctionFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyAuctionContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionContractUpgrade)
				if err := _KittyAuction.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyAuction *KittyAuctionFilterer) ParseContractUpgrade(log types.Log) (*KittyAuctionContractUpgrade, error) {
	event := new(KittyAuctionContractUpgrade)
	if err := _KittyAuction.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyAuctionPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyAuction contract.
type KittyAuctionPregnantIterator struct {
	Event *KittyAuctionPregnant // Event containing the contract specifics and raw log

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
func (it *KittyAuctionPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionPregnant)
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
		it.Event = new(KittyAuctionPregnant)
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
func (it *KittyAuctionPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionPregnant represents a Pregnant event raised by the KittyAuction contract.
type KittyAuctionPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyAuction *KittyAuctionFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyAuctionPregnantIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionPregnantIterator{contract: _KittyAuction.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyAuction *KittyAuctionFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyAuctionPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionPregnant)
				if err := _KittyAuction.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// ParsePregnant is a log parse operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyAuction *KittyAuctionFilterer) ParsePregnant(log types.Log) (*KittyAuctionPregnant, error) {
	event := new(KittyAuctionPregnant)
	if err := _KittyAuction.contract.UnpackLog(event, "Pregnant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyAuctionTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyAuction contract.
type KittyAuctionTransferIterator struct {
	Event *KittyAuctionTransfer // Event containing the contract specifics and raw log

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
func (it *KittyAuctionTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyAuctionTransfer)
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
		it.Event = new(KittyAuctionTransfer)
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
func (it *KittyAuctionTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyAuctionTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyAuctionTransfer represents a Transfer event raised by the KittyAuction contract.
type KittyAuctionTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyAuction *KittyAuctionFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyAuctionTransferIterator, error) {

	logs, sub, err := _KittyAuction.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyAuctionTransferIterator{contract: _KittyAuction.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyAuction *KittyAuctionFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyAuctionTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyAuction.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyAuctionTransfer)
				if err := _KittyAuction.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyAuction *KittyAuctionFilterer) ParseTransfer(log types.Log) (*KittyAuctionTransfer, error) {
	event := new(KittyAuctionTransfer)
	if err := _KittyAuction.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBaseABI is the input ABI used to generate the binding from.
const KittyBaseABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyBaseFuncSigs maps the 4-byte function signature to its string representation.
var KittyBaseFuncSigs = map[string]string{
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"3f4ba83a": "unpause()",
}

// KittyBaseBin is the compiled bytecode used for deploying new contracts.
var KittyBaseBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a08201526100a590600390600e6100bb565b50600f60055534156100b657600080fd5b61017b565b6002830191839082156101475791602002820160005b8382111561011557835183826101000a81548163ffffffff021916908363ffffffff16021790555092602001926004016020816003010492830192600103026100d1565b80156101455782816101000a81549063ffffffff0219169055600401602081600301049283019260010302610115565b505b50610153929150610157565b5090565b61017891905b8082111561015357805463ffffffff1916815560010161015d565b90565b61067b8061018a6000396000f3006060604052600436106100f05763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416630519ce7981146100f55780630a0f81681461012457806321717ebf1461013757806327d7874c1461014a5780632ba73c151461016b5780633f4ba83a1461018a57806346116e6f1461019d578063481af3d3146101b35780634e0a3379146101c95780635663896e146101e85780635c975abb146101fe5780637a7d4937146102255780638456cb591461024a5780639d6fac6f1461025d578063a45f4bfc1461028c578063b047fb50146102a2578063e6cbe351146102b5575b600080fd5b341561010057600080fd5b6101086102c8565b604051600160a060020a03909116815260200160405180910390f35b341561012f57600080fd5b6101086102d7565b341561014257600080fd5b6101086102e6565b341561015557600080fd5b610169600160a060020a03600435166102f5565b005b341561017657600080fd5b610169600160a060020a0360043516610354565b341561019557600080fd5b6101696103b3565b34156101a857600080fd5b610108600435610417565b34156101be57600080fd5b610108600435610432565b34156101d457600080fd5b610169600160a060020a036004351661044d565b34156101f357600080fd5b6101696004356104ac565b341561020957600080fd5b610211610514565b604051901515815260200160405180910390f35b341561023057600080fd5b610238610535565b60405190815260200160405180910390f35b341561025557600080fd5b61016961053b565b341561026857600080fd5b6102736004356105e9565b60405163ffffffff909116815260200160405180910390f35b341561029757600080fd5b610108600435610616565b34156102ad57600080fd5b610108610631565b34156102c057600080fd5b610108610640565b600154600160a060020a031681565b600054600160a060020a031681565b600c54600160a060020a031681565b60005433600160a060020a0390811691161461031057600080fd5b600160a060020a038116151561032557600080fd5b6000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461036f57600080fd5b600160a060020a038116151561038457600080fd5b6002805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60005433600160a060020a039081169116146103ce57600080fd5b60025474010000000000000000000000000000000000000000900460ff1615156103f757600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600960205260009081526040902054600160a060020a031681565b60005433600160a060020a0390811691161461046857600080fd5b600160a060020a038116151561047d57600080fd5b6001805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0392909216919091179055565b60025433600160a060020a03908116911614806104d7575060005433600160a060020a039081169116145b806104f0575060015433600160a060020a039081169116145b15156104fb57600080fd5b60035463ffffffff16811061050f57600080fd5b600555565b60025474010000000000000000000000000000000000000000900460ff1681565b60055481565b60025433600160a060020a0390811691161480610566575060005433600160a060020a039081169116145b8061057f575060015433600160a060020a039081169116145b151561058a57600080fd5b60025474010000000000000000000000000000000000000000900460ff16156105b257600080fd5b6002805474ff0000000000000000000000000000000000000000191674010000000000000000000000000000000000000000179055565b600381600e81106105f657fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b600254600160a060020a031681565b600b54600160a060020a0316815600a165627a7a7230582023cb3506f6d7eb9939c1071621b212edf1f9ab18a8d979c26061380aa18e32080029"

// DeployKittyBase deploys a new Ethereum contract, binding an instance of KittyBase to it.
func DeployKittyBase(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyBase, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBaseABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyBaseBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyBase{KittyBaseCaller: KittyBaseCaller{contract: contract}, KittyBaseTransactor: KittyBaseTransactor{contract: contract}, KittyBaseFilterer: KittyBaseFilterer{contract: contract}}, nil
}

// KittyBase is an auto generated Go binding around an Ethereum contract.
type KittyBase struct {
	KittyBaseCaller     // Read-only binding to the contract
	KittyBaseTransactor // Write-only binding to the contract
	KittyBaseFilterer   // Log filterer for contract events
}

// KittyBaseCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyBaseCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyBaseTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyBaseFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBaseSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyBaseSession struct {
	Contract     *KittyBase        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyBaseCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyBaseCallerSession struct {
	Contract *KittyBaseCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KittyBaseTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyBaseTransactorSession struct {
	Contract     *KittyBaseTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KittyBaseRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyBaseRaw struct {
	Contract *KittyBase // Generic contract binding to access the raw methods on
}

// KittyBaseCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyBaseCallerRaw struct {
	Contract *KittyBaseCaller // Generic read-only contract binding to access the raw methods on
}

// KittyBaseTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyBaseTransactorRaw struct {
	Contract *KittyBaseTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyBase creates a new instance of KittyBase, bound to a specific deployed contract.
func NewKittyBase(address common.Address, backend bind.ContractBackend) (*KittyBase, error) {
	contract, err := bindKittyBase(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyBase{KittyBaseCaller: KittyBaseCaller{contract: contract}, KittyBaseTransactor: KittyBaseTransactor{contract: contract}, KittyBaseFilterer: KittyBaseFilterer{contract: contract}}, nil
}

// NewKittyBaseCaller creates a new read-only instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseCaller(address common.Address, caller bind.ContractCaller) (*KittyBaseCaller, error) {
	contract, err := bindKittyBase(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBaseCaller{contract: contract}, nil
}

// NewKittyBaseTransactor creates a new write-only instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyBaseTransactor, error) {
	contract, err := bindKittyBase(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBaseTransactor{contract: contract}, nil
}

// NewKittyBaseFilterer creates a new log filterer instance of KittyBase, bound to a specific deployed contract.
func NewKittyBaseFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyBaseFilterer, error) {
	contract, err := bindKittyBase(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyBaseFilterer{contract: contract}, nil
}

// bindKittyBase binds a generic wrapper to an already deployed contract.
func bindKittyBase(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBaseABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBase *KittyBaseRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyBase.Contract.KittyBaseCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBase *KittyBaseRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.Contract.KittyBaseTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBase *KittyBaseRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBase.Contract.KittyBaseTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBase *KittyBaseCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyBase.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBase *KittyBaseTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBase *KittyBaseTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBase.Contract.contract.Transact(opts, method, params...)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBase *KittyBaseCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBase *KittyBaseSession) CeoAddress() (common.Address, error) {
	return _KittyBase.Contract.CeoAddress(&_KittyBase.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBase *KittyBaseCallerSession) CeoAddress() (common.Address, error) {
	return _KittyBase.Contract.CeoAddress(&_KittyBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBase *KittyBaseCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBase *KittyBaseSession) CfoAddress() (common.Address, error) {
	return _KittyBase.Contract.CfoAddress(&_KittyBase.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBase *KittyBaseCallerSession) CfoAddress() (common.Address, error) {
	return _KittyBase.Contract.CfoAddress(&_KittyBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBase *KittyBaseCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBase *KittyBaseSession) CooAddress() (common.Address, error) {
	return _KittyBase.Contract.CooAddress(&_KittyBase.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBase *KittyBaseCallerSession) CooAddress() (common.Address, error) {
	return _KittyBase.Contract.CooAddress(&_KittyBase.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBase *KittyBaseCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBase *KittyBaseSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBase.Contract.Cooldowns(&_KittyBase.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBase *KittyBaseCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBase.Contract.Cooldowns(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBase *KittyBaseSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToApproved(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToApproved(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBase *KittyBaseSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToOwner(&_KittyBase.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.KittyIndexToOwner(&_KittyBase.CallOpts, arg0)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBase *KittyBaseCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBase *KittyBaseSession) Paused() (bool, error) {
	return _KittyBase.Contract.Paused(&_KittyBase.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBase *KittyBaseCallerSession) Paused() (bool, error) {
	return _KittyBase.Contract.Paused(&_KittyBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBase *KittyBaseCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBase *KittyBaseSession) SaleAuction() (common.Address, error) {
	return _KittyBase.Contract.SaleAuction(&_KittyBase.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBase *KittyBaseCallerSession) SaleAuction() (common.Address, error) {
	return _KittyBase.Contract.SaleAuction(&_KittyBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBase *KittyBaseCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBase *KittyBaseSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBase.Contract.SecondsPerBlock(&_KittyBase.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBase *KittyBaseCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBase.Contract.SecondsPerBlock(&_KittyBase.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBase *KittyBaseSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.SireAllowedToAddress(&_KittyBase.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBase *KittyBaseCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBase.Contract.SireAllowedToAddress(&_KittyBase.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBase *KittyBaseCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBase.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBase *KittyBaseSession) SiringAuction() (common.Address, error) {
	return _KittyBase.Contract.SiringAuction(&_KittyBase.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBase *KittyBaseCallerSession) SiringAuction() (common.Address, error) {
	return _KittyBase.Contract.SiringAuction(&_KittyBase.CallOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseSession) Pause() (*types.Transaction, error) {
	return _KittyBase.Contract.Pause(&_KittyBase.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBase *KittyBaseTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyBase.Contract.Pause(&_KittyBase.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBase *KittyBaseTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBase *KittyBaseSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCEO(&_KittyBase.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCEO(&_KittyBase.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBase *KittyBaseTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBase *KittyBaseSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCFO(&_KittyBase.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCFO(&_KittyBase.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBase *KittyBaseTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBase *KittyBaseSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCOO(&_KittyBase.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBase *KittyBaseTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBase.Contract.SetCOO(&_KittyBase.TransactOpts, _newCOO)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBase *KittyBaseTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBase *KittyBaseSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.Contract.SetSecondsPerBlock(&_KittyBase.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBase *KittyBaseTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBase.Contract.SetSecondsPerBlock(&_KittyBase.TransactOpts, secs)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBase.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseSession) Unpause() (*types.Transaction, error) {
	return _KittyBase.Contract.Unpause(&_KittyBase.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBase *KittyBaseTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyBase.Contract.Unpause(&_KittyBase.TransactOpts)
}

// KittyBaseBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyBase contract.
type KittyBaseBirthIterator struct {
	Event *KittyBaseBirth // Event containing the contract specifics and raw log

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
func (it *KittyBaseBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseBirth)
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
		it.Event = new(KittyBaseBirth)
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
func (it *KittyBaseBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseBirth represents a Birth event raised by the KittyBase contract.
type KittyBaseBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBase *KittyBaseFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyBaseBirthIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyBaseBirthIterator{contract: _KittyBase.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBase *KittyBaseFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyBaseBirth) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseBirth)
				if err := _KittyBase.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBase *KittyBaseFilterer) ParseBirth(log types.Log) (*KittyBaseBirth, error) {
	event := new(KittyBaseBirth)
	if err := _KittyBase.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBaseContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyBase contract.
type KittyBaseContractUpgradeIterator struct {
	Event *KittyBaseContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyBaseContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseContractUpgrade)
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
		it.Event = new(KittyBaseContractUpgrade)
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
func (it *KittyBaseContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseContractUpgrade represents a ContractUpgrade event raised by the KittyBase contract.
type KittyBaseContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBase *KittyBaseFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyBaseContractUpgradeIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyBaseContractUpgradeIterator{contract: _KittyBase.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBase *KittyBaseFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyBaseContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseContractUpgrade)
				if err := _KittyBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBase *KittyBaseFilterer) ParseContractUpgrade(log types.Log) (*KittyBaseContractUpgrade, error) {
	event := new(KittyBaseContractUpgrade)
	if err := _KittyBase.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBaseTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyBase contract.
type KittyBaseTransferIterator struct {
	Event *KittyBaseTransfer // Event containing the contract specifics and raw log

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
func (it *KittyBaseTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBaseTransfer)
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
		it.Event = new(KittyBaseTransfer)
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
func (it *KittyBaseTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBaseTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBaseTransfer represents a Transfer event raised by the KittyBase contract.
type KittyBaseTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyBase *KittyBaseFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyBaseTransferIterator, error) {

	logs, sub, err := _KittyBase.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyBaseTransferIterator{contract: _KittyBase.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyBase *KittyBaseFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyBaseTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyBase.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBaseTransfer)
				if err := _KittyBase.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyBase *KittyBaseFilterer) ParseTransfer(log types.Log) (*KittyBaseTransfer, error) {
	event := new(KittyBaseTransfer)
	if err := _KittyBase.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBreedingABI is the input ABI used to generate the binding from.
const KittyBreedingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyBreedingFuncSigs maps the 4-byte function signature to its string representation.
var KittyBreedingFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"4dfff04f": "approveSiring(address,uint256)",
	"b0c35c05": "autoBirthFee()",
	"70a08231": "balanceOf(address)",
	"f7d8c883": "breedWithAuto(uint256,uint256)",
	"46d22c70": "canBreedWith(uint256,uint256)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"bc4006f5": "erc721Metadata()",
	"f2b47d52": "geneScience()",
	"88c2a0bf": "giveBirth(uint256)",
	"1940a936": "isPregnant(uint256)",
	"d3e6f49f": "isReadyToBreed(uint256)",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"183a7947": "pregnantKitties()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"4b85fd55": "setAutoBirthFee(uint256)",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"24e7a38a": "setGeneScienceAddress(address)",
	"e17b25af": "setMetadataAddress(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"0560ff44": "tokenMetadata(uint256,string)",
	"8462151c": "tokensOfOwner(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
}

// KittyBreedingBin is the compiled bytecode used for deploying new contracts.
var KittyBreedingBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a0820152620000a790600390600e620000ca565b50600f60055566071afd498d0000600e553415620000c457600080fd5b62000194565b6002830191839082156200015b5791602002820160005b838211156200012757835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000e1565b8015620001595782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000127565b505b50620001699291506200016d565b5090565b6200019191905b808211156200016957805463ffffffff1916815560010162000174565b90565b6122b380620001a46000396000f3006060604052600436106101df5763ffffffff60e060020a60003504166301ffc9a781146101e45780630519ce79146102305780630560ff441461025f57806306fdde03146102f8578063095ea7b31461030b5780630a0f81681461032f57806318160ddd14610342578063183a7947146103675780631940a9361461037a57806321717ebf1461039057806323b872dd146103a357806324e7a38a146103cb57806327d7874c146103ea5780632ba73c15146104095780633f4ba83a1461042857806346116e6f1461043b57806346d22c7014610451578063481af3d31461046a5780634b85fd55146104805780634dfff04f146104965780634e0a3379146104b85780635663896e146104d75780635c975abb146104ed5780636352211e1461050057806370a08231146105165780637a7d4937146105355780638456cb59146105485780638462151c1461055b57806388c2a0bf146105cd57806395d89b41146105e35780639d6fac6f146105f6578063a45f4bfc14610625578063a9059cbb1461063b578063b047fb501461065d578063b0c35c0514610670578063bc4006f514610683578063d3e6f49f14610696578063e17b25af146106ac578063e6cbe351146106cb578063f2b47d52146106de578063f7d8c883146106f1575b600080fd5b34156101ef57600080fd5b61021c7fffffffff00000000000000000000000000000000000000000000000000000000600435166106ff565b604051901515815260200160405180910390f35b341561023b57600080fd5b610243610986565b604051600160a060020a03909116815260200160405180910390f35b341561026a57600080fd5b610281600480359060248035908101910135610995565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156102bd5780820151838201526020016102a5565b50505050905090810190601f1680156102ea5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561030357600080fd5b610281610a65565b341561031657600080fd5b61032d600160a060020a0360043516602435610a9c565b005b341561033a57600080fd5b610243610b26565b341561034d57600080fd5b610355610b35565b60405190815260200160405180910390f35b341561037257600080fd5b610355610b40565b341561038557600080fd5b61021c600435610b46565b341561039b57600080fd5b610243610b8b565b34156103ae57600080fd5b61032d600160a060020a0360043581169060243516604435610b9a565b34156103d657600080fd5b61032d600160a060020a0360043516610c21565b34156103f557600080fd5b61032d600160a060020a0360043516610cc1565b341561041457600080fd5b61032d600160a060020a0360043516610d13565b341561043357600080fd5b61032d610d65565b341561044657600080fd5b610243600435610db8565b341561045c57600080fd5b61021c600435602435610dd3565b341561047557600080fd5b610243600435610e53565b341561048b57600080fd5b61032d600435610e6e565b34156104a157600080fd5b61032d600160a060020a0360043516602435610e8e565b34156104c357600080fd5b61032d600160a060020a0360043516610ee8565b34156104e257600080fd5b61032d600435610f3a565b34156104f857600080fd5b61021c610fa2565b341561050b57600080fd5b610243600435610fb2565b341561052157600080fd5b610355600160a060020a0360043516610fd6565b341561054057600080fd5b610355610ff1565b341561055357600080fd5b61032d610ff7565b341561056657600080fd5b61057a600160a060020a0360043516611083565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156105b95780820151838201526020016105a1565b505050509050019250505060405180910390f35b34156105d857600080fd5b610355600435611164565b34156105ee57600080fd5b610281611427565b341561060157600080fd5b61060c60043561145e565b60405163ffffffff909116815260200160405180910390f35b341561063057600080fd5b61024360043561148b565b341561064657600080fd5b61032d600160a060020a03600435166024356114a6565b341561066857600080fd5b61024361154d565b341561067b57600080fd5b61035561155c565b341561068e57600080fd5b610243611562565b34156106a157600080fd5b61021c600435611571565b34156106b757600080fd5b61032d600160a060020a036004351661163a565b34156106d657600080fd5b610243611677565b34156106e957600080fd5b610243611686565b61032d600435602435611695565b60006040517f737570706f727473496e7465726661636528627974657334290000000000000081526019016040518091039020600160e060020a03191682600160e060020a031916148061097e57506040517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152601d0160405180910390206040517f746f6b656e734f664f776e657228616464726573732900000000000000000000815260160160405180910390206040517f7472616e7366657246726f6d28616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390206040517f7472616e7366657228616464726573732c75696e743235362900000000000000815260190160405180910390206040517f617070726f766528616464726573732c75696e74323536290000000000000000815260180160405180910390206040517f6f776e65724f662875696e743235362900000000000000000000000000000000815260100160405180910390206040517f62616c616e63654f662861646472657373290000000000000000000000000000815260120160405180910390206040517f746f74616c537570706c792829000000000000000000000000000000000000008152600d0160405180910390206040517f73796d626f6c2829000000000000000000000000000000000000000000000000815260080160405180910390206040517f6e616d652829000000000000000000000000000000000000000000000000000081526006016040518091039020181818181818181818600160e060020a03191682600160e060020a031916145b90505b919050565b600154600160a060020a031681565b61099d6121c0565b6109a56121d2565b600d54600090600160a060020a031615156109bf57600080fd5b600d54600160a060020a031663cb4799f287878760405160e060020a63ffffffff861602815260048101848152604060248301908152604483018490529091606401848480828437820191505094505050505060a060405180830381600087803b1515610a2b57600080fd5b5af11515610a3857600080fd5b50505060405180608001805160209091016040529092509050610a5b828261188d565b9695505050505050565b60408051908101604052600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610ab357600080fd5b610abd33826118e2565b1515610ac857600080fd5b610ad28183611902565b7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600054600160a060020a031681565b600654600019015b90565b600f5481565b6000808211610b5457600080fd5b6006805483908110610b6257fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b600c54600160a060020a031681565b60025460a060020a900460ff1615610bb157600080fd5b600160a060020a0382161515610bc657600080fd5b30600160a060020a031682600160a060020a031614151515610be757600080fd5b610bf13382611930565b1515610bfc57600080fd5b610c0683826118e2565b1515610c1157600080fd5b610c1c838383611950565b505050565b6000805433600160a060020a03908116911614610c3d57600080fd5b5080600160a060020a0381166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610c7c57600080fd5b5af11515610c8957600080fd5b505050604051805190501515610c9e57600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b60005433600160a060020a03908116911614610cdc57600080fd5b600160a060020a0381161515610cf157600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610d2e57600080fd5b600160a060020a0381161515610d4357600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610d8057600080fd5b60025460a060020a900460ff161515610d9857600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b60008080808511610de357600080fd5b60008411610df057600080fd5b6006805486908110610dfe57fe5b90600052602060002090600202019150600684815481101515610e1d57fe5b90600052602060002090600202019050610e3982868387611a38565b8015610e4a5750610e4a8486611bb8565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025433600160a060020a03908116911614610e8957600080fd5b600e55565b60025460a060020a900460ff1615610ea557600080fd5b610eaf33826118e2565b1515610eba57600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610f0357600080fd5b600160a060020a0381161515610f1857600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60025433600160a060020a0390811691161480610f65575060005433600160a060020a039081169116145b80610f7e575060015433600160a060020a039081169116145b1515610f8957600080fd5b60035463ffffffff168110610f9d57600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a031680151561098157600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b60025433600160a060020a0390811691161480611022575060005433600160a060020a039081169116145b8061103b575060015433600160a060020a039081169116145b151561104657600080fd5b60025460a060020a900460ff161561105d57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b61108b6121c0565b60006110956121c0565b60008060006110a387610fd6565b94508415156110d35760006040518059106110bb5750595b9080825280602002602001820160405250955061115a565b846040518059106110e15750595b908082528060200260200182016040525093506110fc610b35565b925060009150600190505b82811161115657600081815260076020526040902054600160a060020a038881169116141561114e578084838151811061113d57fe5b602090810290910101526001909101905b600101611107565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff1615151561118c57600080fd5b600680548a90811061119a57fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff1615156111c757600080fd5b61125c8761010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152611c0d565b151561126757600080fd5b60018701546006805460c060020a90920463ffffffff169750908790811061128b57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a928390048116965091900416849011156112d957600185015460f060020a900461ffff1693505b6010548754865460018a0154600160a060020a0390931692630d9f5aed92919068010000000000000000900467ffffffffffffffff166000190160405160e060020a63ffffffff86160281526004810193909352602483019190915267ffffffffffffffff166044820152606401602060405180830381600087803b151561136057600080fd5b5af1151561136d57600080fd5b505050604051805160008b81526007602052604090205460018a810154929650600160a060020a0390911694506113bc92508b9160c060020a900463ffffffff1690870161ffff168686611c45565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54909150600160a060020a0333169080156108fc0290604051600060405180830381858888f150939c9b505050505050505050505050565b60408051908101604052600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e811061146b57fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff16156114bd57600080fd5b600160a060020a03821615156114d257600080fd5b30600160a060020a031682600160a060020a0316141515156114f357600080fd5b600b54600160a060020a038381169116141561150e57600080fd5b600c54600160a060020a038381169116141561152957600080fd5b61153333826118e2565b151561153e57600080fd5b611549338383611950565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b60008080831161158057600080fd5b600680548490811061158e57fe5b906000526020600020906002020190506116338161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152611ef1565b9392505050565b60005433600160a060020a0390811691161461165557600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b601054600160a060020a031681565b600254600090819060a060020a900460ff16156116b157600080fd5b600e543410156116c057600080fd5b6116ca33856118e2565b15156116d557600080fd5b6116df8385611bb8565b15156116ea57600080fd5b60068054859081106116f857fe5b9060005260206000209060020201915061179d8261010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152611ef1565b15156117a857600080fd5b60068054849081106117b657fe5b9060005260206000209060020201905061185b8161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152611ef1565b151561186657600080fd5b61187282858386611a38565b151561187d57600080fd5b6118878484611f28565b50505050565b6118956121c0565b61189d6121c0565b600080846040518059106118ae5750595b818152601f19601f83011681016020016040529050925050602082019050846118d8828287612092565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a03191690911790558316156119e357600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b600081841415611a4a57506000611bb0565b6001850154608060020a900463ffffffff16821480611a795750600185015460a060020a900463ffffffff1682145b15611a8657506000611bb0565b6001830154608060020a900463ffffffff16841480611ab55750600183015460a060020a900463ffffffff1684145b15611ac257506000611bb0565b6001830154608060020a900463ffffffff161580611aef57506001850154608060020a900463ffffffff16155b15611afc57506001611bb0565b60018581015490840154608060020a9182900463ffffffff90811692909104161480611b47575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b15611b5457506000611bb0565b6001808601549084015460a060020a900463ffffffff908116608060020a909204161480611b9f57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b15611bac57506000611bb0565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a03918216911680821480610e4a57506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b60008160a0015163ffffffff161580159061097e57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600080611c506121fb565b600063ffffffff89168914611c6457600080fd5b63ffffffff88168814611c7657600080fd5b61ffff87168714611c8657600080fd5b600287049250600d8361ffff161115611c9e57600d92505b610100604051908101604090815287825267ffffffffffffffff42166020830152600090820181905263ffffffff808c1660608401528a16608083015260a082015261ffff80851660c0830152881660e082015260068054919350600191808301611d09838261223f565b6000928352602090922085916002020181518155602082015160018201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555060c082015181600101601c6101000a81548161ffff021916908361ffff16021790555060e08201516001909101805461ffff9290921660f060020a027dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555003905063ffffffff81168114611e6457600080fd5b7f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad58582846060015163ffffffff16856080015163ffffffff168651604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390a1611ee560008683611950565b98975050505050505050565b60008160a0015163ffffffff1615801561097e57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600080600683815481101515611f3a57fe5b90600052602060002090600202019150600684815481101515611f5957fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff8716021790559050611fa9826120d7565b611fb2816120d7565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f8054600190810190915587845260079092529182902054908301547f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b8092600160a060020a0390921691879187916801000000000000000090910467ffffffffffffffff1690518085600160a060020a0316600160a060020a031681526020018481526020018381526020018267ffffffffffffffff16815260200194505050505060405180910390a150505050565b60005b602082106120b85782518452602084019350602083019250602082039150612095565b6001826020036101000a03905080198351168185511617909352505050565b600554600182015443919060039060e060020a900461ffff16600e81106120fa57fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff1681151561212557fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff1610156121bd576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b60206040519081016040526000815290565b60806040519081016040526004815b600081526000199190910190602001816121e15790505090565b6101006040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c0820181905260e082015290565b815481835581811511610c1c57600083815260209020610c1c91610b3d9160029182028101918502015b808211156122835760008082556001820155600201612269565b50905600a165627a7a723058205d71dc99559f427637552a755156c88aa2618886db2bfe0b595eb2e57ab3a2c60029"

// DeployKittyBreeding deploys a new Ethereum contract, binding an instance of KittyBreeding to it.
func DeployKittyBreeding(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyBreeding, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBreedingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyBreedingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyBreeding{KittyBreedingCaller: KittyBreedingCaller{contract: contract}, KittyBreedingTransactor: KittyBreedingTransactor{contract: contract}, KittyBreedingFilterer: KittyBreedingFilterer{contract: contract}}, nil
}

// KittyBreeding is an auto generated Go binding around an Ethereum contract.
type KittyBreeding struct {
	KittyBreedingCaller     // Read-only binding to the contract
	KittyBreedingTransactor // Write-only binding to the contract
	KittyBreedingFilterer   // Log filterer for contract events
}

// KittyBreedingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyBreedingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyBreedingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyBreedingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyBreedingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyBreedingSession struct {
	Contract     *KittyBreeding    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyBreedingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyBreedingCallerSession struct {
	Contract *KittyBreedingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// KittyBreedingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyBreedingTransactorSession struct {
	Contract     *KittyBreedingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// KittyBreedingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyBreedingRaw struct {
	Contract *KittyBreeding // Generic contract binding to access the raw methods on
}

// KittyBreedingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyBreedingCallerRaw struct {
	Contract *KittyBreedingCaller // Generic read-only contract binding to access the raw methods on
}

// KittyBreedingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyBreedingTransactorRaw struct {
	Contract *KittyBreedingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyBreeding creates a new instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreeding(address common.Address, backend bind.ContractBackend) (*KittyBreeding, error) {
	contract, err := bindKittyBreeding(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyBreeding{KittyBreedingCaller: KittyBreedingCaller{contract: contract}, KittyBreedingTransactor: KittyBreedingTransactor{contract: contract}, KittyBreedingFilterer: KittyBreedingFilterer{contract: contract}}, nil
}

// NewKittyBreedingCaller creates a new read-only instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingCaller(address common.Address, caller bind.ContractCaller) (*KittyBreedingCaller, error) {
	contract, err := bindKittyBreeding(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingCaller{contract: contract}, nil
}

// NewKittyBreedingTransactor creates a new write-only instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyBreedingTransactor, error) {
	contract, err := bindKittyBreeding(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingTransactor{contract: contract}, nil
}

// NewKittyBreedingFilterer creates a new log filterer instance of KittyBreeding, bound to a specific deployed contract.
func NewKittyBreedingFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyBreedingFilterer, error) {
	contract, err := bindKittyBreeding(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyBreedingFilterer{contract: contract}, nil
}

// bindKittyBreeding binds a generic wrapper to an already deployed contract.
func bindKittyBreeding(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyBreedingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBreeding *KittyBreedingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyBreeding.Contract.KittyBreedingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBreeding *KittyBreedingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.Contract.KittyBreedingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBreeding *KittyBreedingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBreeding.Contract.KittyBreedingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyBreeding *KittyBreedingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyBreeding.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyBreeding *KittyBreedingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyBreeding *KittyBreedingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyBreeding.Contract.contract.Transact(opts, method, params...)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "autoBirthFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyBreeding *KittyBreedingSession) AutoBirthFee() (*big.Int, error) {
	return _KittyBreeding.Contract.AutoBirthFee(&_KittyBreeding.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyBreeding.Contract.AutoBirthFee(&_KittyBreeding.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyBreeding *KittyBreedingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyBreeding *KittyBreedingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyBreeding.Contract.BalanceOf(&_KittyBreeding.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyBreeding *KittyBreedingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyBreeding.Contract.BalanceOf(&_KittyBreeding.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyBreeding *KittyBreedingCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "canBreedWith", _matronId, _sireId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyBreeding *KittyBreedingSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.CanBreedWith(&_KittyBreeding.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.CanBreedWith(&_KittyBreeding.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingSession) CeoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CeoAddress(&_KittyBreeding.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CeoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CeoAddress(&_KittyBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingSession) CfoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CfoAddress(&_KittyBreeding.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CfoAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CfoAddress(&_KittyBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBreeding *KittyBreedingSession) CooAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CooAddress(&_KittyBreeding.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) CooAddress() (common.Address, error) {
	return _KittyBreeding.Contract.CooAddress(&_KittyBreeding.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBreeding *KittyBreedingCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBreeding *KittyBreedingSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBreeding.Contract.Cooldowns(&_KittyBreeding.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyBreeding *KittyBreedingCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyBreeding.Contract.Cooldowns(&_KittyBreeding.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyBreeding *KittyBreedingSession) Erc721Metadata() (common.Address, error) {
	return _KittyBreeding.Contract.Erc721Metadata(&_KittyBreeding.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyBreeding.Contract.Erc721Metadata(&_KittyBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyBreeding *KittyBreedingSession) GeneScience() (common.Address, error) {
	return _KittyBreeding.Contract.GeneScience(&_KittyBreeding.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) GeneScience() (common.Address, error) {
	return _KittyBreeding.Contract.GeneScience(&_KittyBreeding.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "isPregnant", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsPregnant(&_KittyBreeding.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsPregnant(&_KittyBreeding.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "isReadyToBreed", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsReadyToBreed(&_KittyBreeding.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyBreeding.Contract.IsReadyToBreed(&_KittyBreeding.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToApproved(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToApproved(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToOwner(&_KittyBreeding.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.KittyIndexToOwner(&_KittyBreeding.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyBreeding *KittyBreedingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyBreeding *KittyBreedingSession) Name() (string, error) {
	return _KittyBreeding.Contract.Name(&_KittyBreeding.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyBreeding *KittyBreedingCallerSession) Name() (string, error) {
	return _KittyBreeding.Contract.Name(&_KittyBreeding.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyBreeding *KittyBreedingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyBreeding *KittyBreedingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.OwnerOf(&_KittyBreeding.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyBreeding *KittyBreedingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.OwnerOf(&_KittyBreeding.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBreeding *KittyBreedingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBreeding *KittyBreedingSession) Paused() (bool, error) {
	return _KittyBreeding.Contract.Paused(&_KittyBreeding.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) Paused() (bool, error) {
	return _KittyBreeding.Contract.Paused(&_KittyBreeding.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "pregnantKitties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyBreeding *KittyBreedingSession) PregnantKitties() (*big.Int, error) {
	return _KittyBreeding.Contract.PregnantKitties(&_KittyBreeding.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyBreeding.Contract.PregnantKitties(&_KittyBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBreeding *KittyBreedingSession) SaleAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SaleAuction(&_KittyBreeding.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SaleAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SaleAuction(&_KittyBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBreeding *KittyBreedingSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBreeding.Contract.SecondsPerBlock(&_KittyBreeding.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyBreeding.Contract.SecondsPerBlock(&_KittyBreeding.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.SireAllowedToAddress(&_KittyBreeding.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyBreeding.Contract.SireAllowedToAddress(&_KittyBreeding.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBreeding *KittyBreedingCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBreeding *KittyBreedingSession) SiringAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SiringAuction(&_KittyBreeding.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyBreeding *KittyBreedingCallerSession) SiringAuction() (common.Address, error) {
	return _KittyBreeding.Contract.SiringAuction(&_KittyBreeding.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyBreeding *KittyBreedingCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyBreeding *KittyBreedingSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyBreeding.Contract.SupportsInterface(&_KittyBreeding.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyBreeding *KittyBreedingCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyBreeding.Contract.SupportsInterface(&_KittyBreeding.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyBreeding *KittyBreedingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyBreeding *KittyBreedingSession) Symbol() (string, error) {
	return _KittyBreeding.Contract.Symbol(&_KittyBreeding.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyBreeding *KittyBreedingCallerSession) Symbol() (string, error) {
	return _KittyBreeding.Contract.Symbol(&_KittyBreeding.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyBreeding *KittyBreedingCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyBreeding *KittyBreedingSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyBreeding.Contract.TokenMetadata(&_KittyBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyBreeding *KittyBreedingCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyBreeding.Contract.TokenMetadata(&_KittyBreeding.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyBreeding *KittyBreedingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyBreeding *KittyBreedingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyBreeding.Contract.TokensOfOwner(&_KittyBreeding.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyBreeding *KittyBreedingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyBreeding.Contract.TokensOfOwner(&_KittyBreeding.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyBreeding *KittyBreedingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyBreeding.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyBreeding *KittyBreedingSession) TotalSupply() (*big.Int, error) {
	return _KittyBreeding.Contract.TotalSupply(&_KittyBreeding.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyBreeding *KittyBreedingCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyBreeding.Contract.TotalSupply(&_KittyBreeding.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Approve(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Approve(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyBreeding *KittyBreedingTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyBreeding *KittyBreedingSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.ApproveSiring(&_KittyBreeding.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.ApproveSiring(&_KittyBreeding.TransactOpts, _addr, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyBreeding *KittyBreedingTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyBreeding *KittyBreedingSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.BreedWithAuto(&_KittyBreeding.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyBreeding *KittyBreedingTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.BreedWithAuto(&_KittyBreeding.TransactOpts, _matronId, _sireId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyBreeding *KittyBreedingTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyBreeding *KittyBreedingSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.GiveBirth(&_KittyBreeding.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyBreeding *KittyBreedingTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.GiveBirth(&_KittyBreeding.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingSession) Pause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Pause(&_KittyBreeding.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Pause(&_KittyBreeding.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyBreeding *KittyBreedingSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetAutoBirthFee(&_KittyBreeding.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetAutoBirthFee(&_KittyBreeding.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBreeding *KittyBreedingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCEO(&_KittyBreeding.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCEO(&_KittyBreeding.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBreeding *KittyBreedingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCFO(&_KittyBreeding.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCFO(&_KittyBreeding.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBreeding *KittyBreedingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCOO(&_KittyBreeding.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetCOO(&_KittyBreeding.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyBreeding *KittyBreedingSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetGeneScienceAddress(&_KittyBreeding.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetGeneScienceAddress(&_KittyBreeding.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyBreeding *KittyBreedingSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetMetadataAddress(&_KittyBreeding.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetMetadataAddress(&_KittyBreeding.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBreeding *KittyBreedingTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBreeding *KittyBreedingSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetSecondsPerBlock(&_KittyBreeding.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.SetSecondsPerBlock(&_KittyBreeding.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Transfer(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.Transfer(&_KittyBreeding.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.TransferFrom(&_KittyBreeding.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyBreeding *KittyBreedingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyBreeding.Contract.TransferFrom(&_KittyBreeding.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyBreeding.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingSession) Unpause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Unpause(&_KittyBreeding.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyBreeding *KittyBreedingTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyBreeding.Contract.Unpause(&_KittyBreeding.TransactOpts)
}

// KittyBreedingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyBreeding contract.
type KittyBreedingApprovalIterator struct {
	Event *KittyBreedingApproval // Event containing the contract specifics and raw log

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
func (it *KittyBreedingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingApproval)
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
		it.Event = new(KittyBreedingApproval)
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
func (it *KittyBreedingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingApproval represents a Approval event raised by the KittyBreeding contract.
type KittyBreedingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyBreeding *KittyBreedingFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyBreedingApprovalIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingApprovalIterator{contract: _KittyBreeding.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyBreeding *KittyBreedingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyBreedingApproval) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingApproval)
				if err := _KittyBreeding.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KittyBreeding *KittyBreedingFilterer) ParseApproval(log types.Log) (*KittyBreedingApproval, error) {
	event := new(KittyBreedingApproval)
	if err := _KittyBreeding.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBreedingBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyBreeding contract.
type KittyBreedingBirthIterator struct {
	Event *KittyBreedingBirth // Event containing the contract specifics and raw log

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
func (it *KittyBreedingBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingBirth)
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
		it.Event = new(KittyBreedingBirth)
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
func (it *KittyBreedingBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingBirth represents a Birth event raised by the KittyBreeding contract.
type KittyBreedingBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBreeding *KittyBreedingFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyBreedingBirthIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingBirthIterator{contract: _KittyBreeding.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBreeding *KittyBreedingFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyBreedingBirth) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingBirth)
				if err := _KittyBreeding.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyBreeding *KittyBreedingFilterer) ParseBirth(log types.Log) (*KittyBreedingBirth, error) {
	event := new(KittyBreedingBirth)
	if err := _KittyBreeding.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBreedingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyBreeding contract.
type KittyBreedingContractUpgradeIterator struct {
	Event *KittyBreedingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyBreedingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingContractUpgrade)
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
		it.Event = new(KittyBreedingContractUpgrade)
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
func (it *KittyBreedingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingContractUpgrade represents a ContractUpgrade event raised by the KittyBreeding contract.
type KittyBreedingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBreeding *KittyBreedingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyBreedingContractUpgradeIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingContractUpgradeIterator{contract: _KittyBreeding.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBreeding *KittyBreedingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyBreedingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingContractUpgrade)
				if err := _KittyBreeding.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyBreeding *KittyBreedingFilterer) ParseContractUpgrade(log types.Log) (*KittyBreedingContractUpgrade, error) {
	event := new(KittyBreedingContractUpgrade)
	if err := _KittyBreeding.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBreedingPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyBreeding contract.
type KittyBreedingPregnantIterator struct {
	Event *KittyBreedingPregnant // Event containing the contract specifics and raw log

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
func (it *KittyBreedingPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingPregnant)
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
		it.Event = new(KittyBreedingPregnant)
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
func (it *KittyBreedingPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingPregnant represents a Pregnant event raised by the KittyBreeding contract.
type KittyBreedingPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyBreeding *KittyBreedingFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyBreedingPregnantIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingPregnantIterator{contract: _KittyBreeding.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyBreeding *KittyBreedingFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyBreedingPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingPregnant)
				if err := _KittyBreeding.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// ParsePregnant is a log parse operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyBreeding *KittyBreedingFilterer) ParsePregnant(log types.Log) (*KittyBreedingPregnant, error) {
	event := new(KittyBreedingPregnant)
	if err := _KittyBreeding.contract.UnpackLog(event, "Pregnant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyBreedingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyBreeding contract.
type KittyBreedingTransferIterator struct {
	Event *KittyBreedingTransfer // Event containing the contract specifics and raw log

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
func (it *KittyBreedingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyBreedingTransfer)
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
		it.Event = new(KittyBreedingTransfer)
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
func (it *KittyBreedingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyBreedingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyBreedingTransfer represents a Transfer event raised by the KittyBreeding contract.
type KittyBreedingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyBreeding *KittyBreedingFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyBreedingTransferIterator, error) {

	logs, sub, err := _KittyBreeding.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyBreedingTransferIterator{contract: _KittyBreeding.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyBreeding *KittyBreedingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyBreedingTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyBreeding.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyBreedingTransfer)
				if err := _KittyBreeding.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyBreeding *KittyBreedingFilterer) ParseTransfer(log types.Log) (*KittyBreedingTransfer, error) {
	event := new(KittyBreedingTransfer)
	if err := _KittyBreeding.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyCoreABI is the input ABI used to generate the binding from.
const KittyCoreABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"newContractAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_v2Address\",\"type\":\"address\"}],\"name\":\"setNewAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_id\",\"type\":\"uint256\"}],\"name\":\"getKitty\",\"outputs\":[{\"name\":\"isGestating\",\"type\":\"bool\"},{\"name\":\"isReady\",\"type\":\"bool\"},{\"name\":\"cooldownIndex\",\"type\":\"uint256\"},{\"name\":\"nextActionAt\",\"type\":\"uint256\"},{\"name\":\"siringWithId\",\"type\":\"uint256\"},{\"name\":\"birthTime\",\"type\":\"uint256\"},{\"name\":\"matronId\",\"type\":\"uint256\"},{\"name\":\"sireId\",\"type\":\"uint256\"},{\"name\":\"generation\",\"type\":\"uint256\"},{\"name\":\"genes\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyCoreFuncSigs maps the 4-byte function signature to its string representation.
var KittyCoreFuncSigs = map[string]string{
	"19c2f201": "GEN0_AUCTION_DURATION()",
	"680eba27": "GEN0_CREATION_LIMIT()",
	"0e583df0": "GEN0_STARTING_PRICE()",
	"defb9584": "PROMO_CREATION_LIMIT()",
	"095ea7b3": "approve(address,uint256)",
	"4dfff04f": "approveSiring(address,uint256)",
	"b0c35c05": "autoBirthFee()",
	"70a08231": "balanceOf(address)",
	"ed60ade6": "bidOnSiringAuction(uint256,uint256)",
	"f7d8c883": "breedWithAuto(uint256,uint256)",
	"46d22c70": "canBreedWith(uint256,uint256)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"c3bea9af": "createGen0Auction(uint256)",
	"56129134": "createPromoKitty(uint256,address)",
	"3d7d3f5a": "createSaleAuction(uint256,uint256,uint256,uint256)",
	"4ad8c938": "createSiringAuction(uint256,uint256,uint256,uint256)",
	"bc4006f5": "erc721Metadata()",
	"f1ca9410": "gen0CreatedCount()",
	"f2b47d52": "geneScience()",
	"e98b7f4d": "getKitty(uint256)",
	"88c2a0bf": "giveBirth(uint256)",
	"1940a936": "isPregnant(uint256)",
	"d3e6f49f": "isReadyToBreed(uint256)",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"06fdde03": "name()",
	"6af04a57": "newContractAddress()",
	"6352211e": "ownerOf(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"183a7947": "pregnantKitties()",
	"05e45546": "promoCreatedCount()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"4b85fd55": "setAutoBirthFee(uint256)",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"24e7a38a": "setGeneScienceAddress(address)",
	"e17b25af": "setMetadataAddress(address)",
	"71587988": "setNewAddress(address)",
	"6fbde40d": "setSaleAuctionAddress(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"14001f4c": "setSiringAuctionAddress(address)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"0560ff44": "tokenMetadata(uint256,string)",
	"8462151c": "tokensOfOwner(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
	"91876e57": "withdrawAuctionBalances()",
	"5fd8c710": "withdrawBalance()",
}

// KittyCoreBin is the compiled bytecode used for deploying new contracts.
var KittyCoreBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a0820152620000a790600390600e620004e4565b50600f60055566071afd498d0000600e553415620000c457600080fd5b6002805460008054600160a060020a033316600160a060020a03199182168117835560a060020a60ff02199093167401000000000000000000000000000000000000000017169091179091556200012f908080600019816401000000006200288e6200013682021704565b5062000649565b6000806200014362000587565b600063ffffffff891689146200015857600080fd5b63ffffffff881688146200016b57600080fd5b61ffff871687146200017c57600080fd5b600287049250600d8361ffff1611156200019557600d92505b61010060405190810160409081528782526001604060020a0342166020830152600090820181905263ffffffff808c1660608401528a16608083015260a082015261ffff80851660c0830152881660e082015260068054919350600191808301620002018382620005cb565b6000928352602090922085916002020181518155602082015160018201805467ffffffffffffffff19166001604060020a039290921691909117905560408201518160010160086101000a8154816001604060020a0302191690836001604060020a0316021790555060608201518160010160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555060c082015181600101601c6101000a81548161ffff021916908361ffff16021790555060e08201516001909101805461ffff929092167e0100000000000000000000000000000000000000000000000000000000000002600160f060020a039092169190911790555003905063ffffffff811681146200035e57600080fd5b7f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad58582846060015163ffffffff16856080015163ffffffff168651604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390a1620003ef600086836401000000006200257e620003fb82021704565b98975050505050505050565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a03191690911790558316156200048f57600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b600283019183908215620005755791602002820160005b838211156200054157835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620004fb565b8015620005735782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000541565b505b5062000583929150620005ff565b5090565b6101006040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c0820181905260e082015290565b815481835581811511620005fa57600202816002028360005260206000209182019101620005fa919062000626565b505050565b6200062391905b808211156200058357805463ffffffff1916815560010162000606565b90565b6200062391905b808211156200058357600080825560018201556002016200062d565b61302580620006596000396000f3006060604052600436106102a55763ffffffff60e060020a60003504166301ffc9a781146102dd5780630519ce79146103295780630560ff441461035857806305e45546146103f157806306fdde0314610416578063095ea7b3146104295780630a0f81681461044b5780630e583df01461045e57806314001f4c1461047157806318160ddd14610490578063183a7947146104a35780631940a936146104b657806319c2f201146104cc57806321717ebf146104df57806323b872dd146104f257806324e7a38a1461051a57806327d7874c146105395780632ba73c15146105585780633d7d3f5a146105775780633f4ba83a1461059657806346116e6f146105a957806346d22c70146105bf578063481af3d3146105d85780634ad8c938146105ee5780634b85fd551461060d5780634dfff04f146106235780634e0a33791461064557806356129134146106645780635663896e146106865780635c975abb1461069c5780635fd8c710146106af5780636352211e146106c2578063680eba27146106d85780636af04a57146106eb5780636fbde40d146106fe57806370a082311461071d578063715879881461073c5780637a7d49371461075b5780638456cb591461076e5780638462151c1461078157806388c2a0bf146107f357806391876e571461080957806395d89b411461081c5780639d6fac6f1461082f578063a45f4bfc1461085e578063a9059cbb14610874578063b047fb5014610896578063b0c35c05146108a9578063bc4006f5146108bc578063c3bea9af146108cf578063d3e6f49f146108e5578063defb9584146108fb578063e17b25af1461090e578063e6cbe3511461092d578063e98b7f4d14610940578063ed60ade6146109ae578063f1ca9410146109bc578063f2b47d52146109cf578063f7d8c883146109e2575b600b5433600160a060020a03908116911614806102d05750600c5433600160a060020a039081169116145b15156102db57600080fd5b005b34156102e857600080fd5b6103157fffffffff00000000000000000000000000000000000000000000000000000000600435166109f0565b604051901515815260200160405180910390f35b341561033457600080fd5b61033c610c77565b604051600160a060020a03909116815260200160405180910390f35b341561036357600080fd5b61037a600480359060248035908101910135610c86565b60405160208082528190810183818151815260200191508051906020019080838360005b838110156103b657808201518382015260200161039e565b50505050905090810190601f1680156103e35780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156103fc57600080fd5b610404610d56565b60405190815260200160405180910390f35b341561042157600080fd5b61037a610d5c565b341561043457600080fd5b6102db600160a060020a0360043516602435610d93565b341561045657600080fd5b61033c610e1d565b341561046957600080fd5b610404610e2c565b341561047c57600080fd5b6102db600160a060020a0360043516610e37565b341561049b57600080fd5b610404610ed7565b34156104ae57600080fd5b610404610ee2565b34156104c157600080fd5b610315600435610ee8565b34156104d757600080fd5b610404610f2d565b34156104ea57600080fd5b61033c610f34565b34156104fd57600080fd5b6102db600160a060020a0360043581169060243516604435610f43565b341561052557600080fd5b6102db600160a060020a0360043516610fca565b341561054457600080fd5b6102db600160a060020a036004351661106a565b341561056357600080fd5b6102db600160a060020a03600435166110bc565b341561058257600080fd5b6102db60043560243560443560643561110e565b34156105a157600080fd5b6102db6111e9565b34156105b457600080fd5b61033c600435611281565b34156105ca57600080fd5b61031560043560243561129c565b34156105e357600080fd5b61033c60043561131c565b34156105f957600080fd5b6102db600435602435604435606435611337565b341561061857600080fd5b6102db6004356113fd565b341561062e57600080fd5b6102db600160a060020a036004351660243561141d565b341561065057600080fd5b6102db600160a060020a0360043516611477565b341561066f57600080fd5b6102db600435600160a060020a03602435166114c9565b341561069157600080fd5b6102db600435611535565b34156106a757600080fd5b61031561159d565b34156106ba57600080fd5b6102db6115ad565b34156106cd57600080fd5b61033c60043561161e565b34156106e357600080fd5b610404611642565b34156106f657600080fd5b61033c611648565b341561070957600080fd5b6102db600160a060020a0360043516611657565b341561072857600080fd5b610404600160a060020a03600435166116f7565b341561074757600080fd5b6102db600160a060020a0360043516611712565b341561076657600080fd5b6104046117a0565b341561077957600080fd5b6102db6117a6565b341561078c57600080fd5b6107a0600160a060020a0360043516611832565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156107df5780820151838201526020016107c7565b505050509050019250505060405180910390f35b34156107fe57600080fd5b610404600435611913565b341561081457600080fd5b6102db611bd6565b341561082757600080fd5b61037a611cc1565b341561083a57600080fd5b610845600435611cf8565b60405163ffffffff909116815260200160405180910390f35b341561086957600080fd5b61033c600435611d25565b341561087f57600080fd5b6102db600160a060020a0360043516602435611d40565b34156108a157600080fd5b61033c611de3565b34156108b457600080fd5b610404611df2565b34156108c757600080fd5b61033c611df8565b34156108da57600080fd5b6102db600435611e07565b34156108f057600080fd5b610315600435611ef6565b341561090657600080fd5b610404611fbf565b341561091957600080fd5b6102db600160a060020a0360043516611fc5565b341561093857600080fd5b61033c612002565b341561094b57600080fd5b610956600435612011565b6040519915158a5297151560208a01526040808a01979097526060890195909552608088019390935260a087019190915260c086015260e0850152610100840152610120830191909152610140909101905180910390f35b6102db600435602435612172565b34156109c757600080fd5b6104046122b4565b34156109da57600080fd5b61033c6122ba565b6102db6004356024356122c9565b60006040517f737570706f727473496e7465726661636528627974657334290000000000000081526019016040518091039020600160e060020a03191682600160e060020a0319161480610c6f57506040517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152601d0160405180910390206040517f746f6b656e734f664f776e657228616464726573732900000000000000000000815260160160405180910390206040517f7472616e7366657246726f6d28616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390206040517f7472616e7366657228616464726573732c75696e743235362900000000000000815260190160405180910390206040517f617070726f766528616464726573732c75696e74323536290000000000000000815260180160405180910390206040517f6f776e65724f662875696e743235362900000000000000000000000000000000815260100160405180910390206040517f62616c616e63654f662861646472657373290000000000000000000000000000815260120160405180910390206040517f746f74616c537570706c792829000000000000000000000000000000000000008152600d0160405180910390206040517f73796d626f6c2829000000000000000000000000000000000000000000000000815260080160405180910390206040517f6e616d652829000000000000000000000000000000000000000000000000000081526006016040518091039020181818181818181818600160e060020a03191682600160e060020a031916145b90505b919050565b600154600160a060020a031681565b610c8e612f32565b610c96612f44565b600d54600090600160a060020a03161515610cb057600080fd5b600d54600160a060020a031663cb4799f287878760405160e060020a63ffffffff861602815260048101848152604060248301908152604483018490529091606401848480828437820191505094505050505060a060405180830381600087803b1515610d1c57600080fd5b5af11515610d2957600080fd5b50505060405180608001805160209091016040529092509050610d4c82826124bb565b9695505050505050565b60115481565b60408051908101604052600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610daa57600080fd5b610db43382612510565b1515610dbf57600080fd5b610dc98183612530565b7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600054600160a060020a031681565b662386f26fc1000081565b6000805433600160a060020a03908116911614610e5357600080fd5b5080600160a060020a0381166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610e9257600080fd5b5af11515610e9f57600080fd5b505050604051805190501515610eb457600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b600654600019015b90565b600f5481565b6000808211610ef657600080fd5b6006805483908110610f0457fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b6201518081565b600c54600160a060020a031681565b60025460a060020a900460ff1615610f5a57600080fd5b600160a060020a0382161515610f6f57600080fd5b30600160a060020a031682600160a060020a031614151515610f9057600080fd5b610f9a338261255e565b1515610fa557600080fd5b610faf8382612510565b1515610fba57600080fd5b610fc583838361257e565b505050565b6000805433600160a060020a03908116911614610fe657600080fd5b5080600160a060020a0381166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b151561102557600080fd5b5af1151561103257600080fd5b50505060405180519050151561104757600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b60005433600160a060020a0390811691161461108557600080fd5b600160a060020a038116151561109a57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a039081169116146110d757600080fd5b600160a060020a03811615156110ec57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff161561112557600080fd5b61112f3385612510565b151561113a57600080fd5b61114384610ee8565b1561114d57600080fd5b600b54611164908590600160a060020a0316612530565b600b54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b15156111d357600080fd5b5af115156111e057600080fd5b50505050505050565b60005433600160a060020a0390811691161461120457600080fd5b60025460a060020a900460ff16151561121c57600080fd5b600b54600160a060020a0316151561123357600080fd5b600c54600160a060020a0316151561124a57600080fd5b601054600160a060020a0316151561126157600080fd5b601354600160a060020a03161561127757600080fd5b61127f612666565b565b600a60205260009081526040902054600160a060020a031681565b600080808085116112ac57600080fd5b600084116112b957600080fd5b60068054869081106112c757fe5b906000526020600020906002020191506006848154811015156112e657fe5b90600052602060002090600202019050611302828683876126b9565b801561131357506113138486612839565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff161561134e57600080fd5b6113583385612510565b151561136357600080fd5b61136c84611ef6565b151561137757600080fd5b600c5461138e908590600160a060020a0316612530565b600c54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b15156111d357600080fd5b60025433600160a060020a0390811691161461141857600080fd5b600e55565b60025460a060020a900460ff161561143457600080fd5b61143e3382612510565b151561144957600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461149257600080fd5b600160a060020a03811615156114a757600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60025460009033600160a060020a039081169116146114e757600080fd5b5080600160a060020a03811615156115075750600254600160a060020a03165b601154611388901061151857600080fd5b60118054600101905561152f60008080868561288e565b50505050565b60025433600160a060020a0390811691161480611560575060005433600160a060020a039081169116145b80611579575060015433600160a060020a039081169116145b151561158457600080fd5b60035463ffffffff16811061159857600080fd5b600555565b60025460a060020a900460ff1681565b600154600090819033600160a060020a039081169116146115cd57600080fd5b30600160a060020a0316319150600e54600f546001010290508082111561161a57600154600160a060020a031681830380156108fc0290604051600060405180830381858888f150505050505b5050565b600081815260076020526040902054600160a060020a0316801515610c7257600080fd5b61afc881565b601354600160a060020a031681565b6000805433600160a060020a0390811691161461167357600080fd5b5080600160a060020a0381166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156116b257600080fd5b5af115156116bf57600080fd5b5050506040518051905015156116d457600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b60005433600160a060020a0390811691161461172d57600080fd5b60025460a060020a900460ff16151561174557600080fd5b60138054600160a060020a031916600160a060020a0383161790557f450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa44619930581604051600160a060020a03909116815260200160405180910390a150565b60055481565b60025433600160a060020a03908116911614806117d1575060005433600160a060020a039081169116145b806117ea575060015433600160a060020a039081169116145b15156117f557600080fd5b60025460a060020a900460ff161561180c57600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b61183a612f32565b6000611844612f32565b6000806000611852876116f7565b945084151561188257600060405180591061186a5750595b90808252806020026020018201604052509550611909565b846040518059106118905750595b908082528060200260200182016040525093506118ab610ed7565b925060009150600190505b82811161190557600081815260076020526040902054600160a060020a03888116911614156118fd57808483815181106118ec57fe5b602090810290910101526001909101905b6001016118b6565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff1615151561193b57600080fd5b600680548a90811061194957fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561197657600080fd5b611a0b8761010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612b3a565b1515611a1657600080fd5b60018701546006805460c060020a90920463ffffffff1697509087908110611a3a57fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a92839004811696509190041684901115611a8857600185015460f060020a900461ffff1693505b6010548754865460018a0154600160a060020a0390931692630d9f5aed92919068010000000000000000900467ffffffffffffffff166000190160405160e060020a63ffffffff86160281526004810193909352602483019190915267ffffffffffffffff166044820152606401602060405180830381600087803b1515611b0f57600080fd5b5af11515611b1c57600080fd5b505050604051805160008b81526007602052604090205460018a810154929650600160a060020a039091169450611b6b92508b9160c060020a900463ffffffff1690870161ffff16868661288e565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54909150600160a060020a0333169080156108fc0290604051600060405180830381858888f150939c9b505050505050505050505050565b60025433600160a060020a0390811691161480611c01575060005433600160a060020a039081169116145b80611c1a575060015433600160a060020a039081169116145b1515611c2557600080fd5b600b54600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b1515611c6457600080fd5b5af11515611c7157600080fd5b5050600c54600160a060020a03169050635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b1515611cb457600080fd5b5af11515610fc557600080fd5b60408051908101604052600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110611d0557fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615611d5757600080fd5b600160a060020a0382161515611d6c57600080fd5b30600160a060020a031682600160a060020a031614151515611d8d57600080fd5b600b54600160a060020a0383811691161415611da857600080fd5b600c54600160a060020a0383811691161415611dc357600080fd5b611dcd3382612510565b1515611dd857600080fd5b61161a33838361257e565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b60025460009033600160a060020a03908116911614611e2557600080fd5b60125461afc89010611e3657600080fd5b611e456000806000853061288e565b600b54909150611e5f908290600160a060020a0316612530565b600b54600160a060020a03166327ebe40a82611e79612b72565b6000620151803060405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b1515611ed957600080fd5b5af11515611ee657600080fd5b5050601280546001019055505050565b600080808311611f0557600080fd5b6006805484908110611f1357fe5b90600052602060002090600202019050611fb88161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612c14565b9392505050565b61138881565b60005433600160a060020a03908116911614611fe057600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b600080600080600080600080600080600060068c81548110151561203157fe5b906000526020600020906002020190508060010160189054906101000a900463ffffffff1663ffffffff16600014159a50438160010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff161115995080600101601c9054906101000a900461ffff1661ffff1698508060010160089054906101000a900467ffffffffffffffff1667ffffffffffffffff1697508060010160189054906101000a900463ffffffff1663ffffffff1696508060010160009054906101000a900467ffffffffffffffff1667ffffffffffffffff1695508060010160109054906101000a900463ffffffff1663ffffffff1694508060010160149054906101000a900463ffffffff1663ffffffff16935080600101601e9054906101000a900461ffff1661ffff16925080600001549150509193959799509193959799565b60025460009060a060020a900460ff161561218c57600080fd5b6121963383612510565b15156121a157600080fd5b6121aa82611ef6565b15156121b557600080fd5b6121bf8284612c4b565b15156121ca57600080fd5b600c54600160a060020a031663c55d0f568460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b151561221257600080fd5b5af1151561221f57600080fd5b5050506040518051600e549092508201341015905061223d57600080fd5b600c54600e54600160a060020a039091169063454a2ab39034038560405160e060020a63ffffffff851602815260048101919091526024016000604051808303818588803b151561228d57600080fd5b5af1151561229a57600080fd5b50505050610fc58263ffffffff168463ffffffff16612c9a565b60125481565b601054600160a060020a031681565b600254600090819060a060020a900460ff16156122e557600080fd5b600e543410156122f457600080fd5b6122fe3385612510565b151561230957600080fd5b6123138385612839565b151561231e57600080fd5b600680548590811061232c57fe5b906000526020600020906002020191506123d18261010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612c14565b15156123dc57600080fd5b60068054849081106123ea57fe5b9060005260206000209060020201905061248f8161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612c14565b151561249a57600080fd5b6124a6828583866126b9565b15156124b157600080fd5b61152f8484612c9a565b6124c3612f32565b6124cb612f32565b600080846040518059106124dc5750595b818152601f19601f8301168101602001604052905092505060208201905084612506828287612e04565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561261157600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60005433600160a060020a0390811691161461268157600080fd5b60025460a060020a900460ff16151561269957600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b6000818414156126cb57506000612831565b6001850154608060020a900463ffffffff168214806126fa5750600185015460a060020a900463ffffffff1682145b1561270757506000612831565b6001830154608060020a900463ffffffff168414806127365750600183015460a060020a900463ffffffff1684145b1561274357506000612831565b6001830154608060020a900463ffffffff16158061277057506001850154608060020a900463ffffffff16155b1561277d57506001612831565b60018581015490840154608060020a9182900463ffffffff908116929091041614806127c8575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b156127d557506000612831565b6001808601549084015460a060020a900463ffffffff908116608060020a90920416148061282057506001858101549084015460a060020a9182900463ffffffff9081169290910416145b1561282d57506000612831565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a0391821691168082148061131357506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b600080612899612f6d565b600063ffffffff891689146128ad57600080fd5b63ffffffff881688146128bf57600080fd5b61ffff871687146128cf57600080fd5b600287049250600d8361ffff1611156128e757600d92505b610100604051908101604090815287825267ffffffffffffffff42166020830152600090820181905263ffffffff808c1660608401528a16608083015260a082015261ffff80851660c0830152881660e0820152600680549193506001918083016129528382612fb1565b6000928352602090922085916002020181518155602082015160018201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555060c082015181600101601c6101000a81548161ffff021916908361ffff16021790555060e08201516001909101805461ffff9290921660f060020a027dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555003905063ffffffff81168114612aad57600080fd5b7f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad58582846060015163ffffffff16856080015163ffffffff168651604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390a1612b2e6000868361257e565b98975050505050505050565b60008160a0015163ffffffff1615801590610c6f57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600b5460009081908190600160a060020a031663eac9d94c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515612bb857600080fd5b5af11515612bc557600080fd5b50505060405180519250506fffffffffffffffffffffffffffffffff82168214612bee57600080fd5b50600281048101662386f26fc10000811015612c0e5750662386f26fc100005b92915050565b60008160a0015163ffffffff16158015610c6f57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b6000806000600685815481101515612c5f57fe5b90600052602060002090600202019150600684815481101515612c7e57fe5b90600052602060002090600202019050611313828683876126b9565b600080600683815481101515612cac57fe5b90600052602060002090600202019150600684815481101515612ccb57fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff8716021790559050612d1b82612e49565b612d2481612e49565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f8054600190810190915587845260079092529182902054908301547f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b8092600160a060020a0390921691879187916801000000000000000090910467ffffffffffffffff1690518085600160a060020a0316600160a060020a031681526020018481526020018381526020018267ffffffffffffffff16815260200194505050505060405180910390a150505050565b60005b60208210612e2a5782518452602084019350602083019250602082039150612e07565b6001826020036101000a03905080198351168185511617909352505050565b600554600182015443919060039060e060020a900461ffff16600e8110612e6c57fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff16811515612e9757fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff161015612f2f576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b60206040519081016040526000815290565b60806040519081016040526004815b60008152600019919091019060200181612f535790505090565b6101006040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c0820181905260e082015290565b815481835581811511610fc557600083815260209020610fc591610edf9160029182028101918502015b80821115612ff55760008082556001820155600201612fdb565b50905600a165627a7a723058201377ba8df8c95c58e83203c4cb84f8d1add7d4e548e68be2bc144b1800439e2c0029"

// DeployKittyCore deploys a new Ethereum contract, binding an instance of KittyCore to it.
func DeployKittyCore(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyCore, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyCoreBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyCore{KittyCoreCaller: KittyCoreCaller{contract: contract}, KittyCoreTransactor: KittyCoreTransactor{contract: contract}, KittyCoreFilterer: KittyCoreFilterer{contract: contract}}, nil
}

// KittyCore is an auto generated Go binding around an Ethereum contract.
type KittyCore struct {
	KittyCoreCaller     // Read-only binding to the contract
	KittyCoreTransactor // Write-only binding to the contract
	KittyCoreFilterer   // Log filterer for contract events
}

// KittyCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyCoreSession struct {
	Contract     *KittyCore        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyCoreCallerSession struct {
	Contract *KittyCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// KittyCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyCoreTransactorSession struct {
	Contract     *KittyCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// KittyCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyCoreRaw struct {
	Contract *KittyCore // Generic contract binding to access the raw methods on
}

// KittyCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyCoreCallerRaw struct {
	Contract *KittyCoreCaller // Generic read-only contract binding to access the raw methods on
}

// KittyCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyCoreTransactorRaw struct {
	Contract *KittyCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyCore creates a new instance of KittyCore, bound to a specific deployed contract.
func NewKittyCore(address common.Address, backend bind.ContractBackend) (*KittyCore, error) {
	contract, err := bindKittyCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyCore{KittyCoreCaller: KittyCoreCaller{contract: contract}, KittyCoreTransactor: KittyCoreTransactor{contract: contract}, KittyCoreFilterer: KittyCoreFilterer{contract: contract}}, nil
}

// NewKittyCoreCaller creates a new read-only instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreCaller(address common.Address, caller bind.ContractCaller) (*KittyCoreCaller, error) {
	contract, err := bindKittyCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyCoreCaller{contract: contract}, nil
}

// NewKittyCoreTransactor creates a new write-only instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyCoreTransactor, error) {
	contract, err := bindKittyCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyCoreTransactor{contract: contract}, nil
}

// NewKittyCoreFilterer creates a new log filterer instance of KittyCore, bound to a specific deployed contract.
func NewKittyCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyCoreFilterer, error) {
	contract, err := bindKittyCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyCoreFilterer{contract: contract}, nil
}

// bindKittyCore binds a generic wrapper to an already deployed contract.
func bindKittyCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyCore *KittyCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyCore.Contract.KittyCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyCore *KittyCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.Contract.KittyCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyCore *KittyCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyCore.Contract.KittyCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyCore *KittyCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyCore *KittyCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyCore *KittyCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyCore.Contract.contract.Transact(opts, method, params...)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0AUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "GEN0_AUCTION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyCore.Contract.GEN0AUCTIONDURATION(&_KittyCore.CallOpts)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyCore.Contract.GEN0AUCTIONDURATION(&_KittyCore.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0CREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "GEN0_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.GEN0CREATIONLIMIT(&_KittyCore.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.GEN0CREATIONLIMIT(&_KittyCore.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyCore *KittyCoreCaller) GEN0STARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "GEN0_STARTING_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyCore *KittyCoreSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyCore.Contract.GEN0STARTINGPRICE(&_KittyCore.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyCore.Contract.GEN0STARTINGPRICE(&_KittyCore.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "PROMO_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.PROMOCREATIONLIMIT(&_KittyCore.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyCore.Contract.PROMOCREATIONLIMIT(&_KittyCore.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyCore *KittyCoreCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "autoBirthFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyCore *KittyCoreSession) AutoBirthFee() (*big.Int, error) {
	return _KittyCore.Contract.AutoBirthFee(&_KittyCore.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyCore.Contract.AutoBirthFee(&_KittyCore.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyCore *KittyCoreCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyCore *KittyCoreSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyCore.Contract.BalanceOf(&_KittyCore.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyCore *KittyCoreCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyCore.Contract.BalanceOf(&_KittyCore.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyCore *KittyCoreCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "canBreedWith", _matronId, _sireId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyCore *KittyCoreSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyCore.Contract.CanBreedWith(&_KittyCore.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyCore *KittyCoreCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyCore.Contract.CanBreedWith(&_KittyCore.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyCore *KittyCoreCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyCore *KittyCoreSession) CeoAddress() (common.Address, error) {
	return _KittyCore.Contract.CeoAddress(&_KittyCore.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyCore *KittyCoreCallerSession) CeoAddress() (common.Address, error) {
	return _KittyCore.Contract.CeoAddress(&_KittyCore.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyCore *KittyCoreCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyCore *KittyCoreSession) CfoAddress() (common.Address, error) {
	return _KittyCore.Contract.CfoAddress(&_KittyCore.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyCore *KittyCoreCallerSession) CfoAddress() (common.Address, error) {
	return _KittyCore.Contract.CfoAddress(&_KittyCore.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyCore *KittyCoreCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyCore *KittyCoreSession) CooAddress() (common.Address, error) {
	return _KittyCore.Contract.CooAddress(&_KittyCore.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyCore *KittyCoreCallerSession) CooAddress() (common.Address, error) {
	return _KittyCore.Contract.CooAddress(&_KittyCore.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyCore *KittyCoreCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyCore *KittyCoreSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyCore.Contract.Cooldowns(&_KittyCore.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyCore *KittyCoreCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyCore.Contract.Cooldowns(&_KittyCore.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyCore *KittyCoreCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyCore *KittyCoreSession) Erc721Metadata() (common.Address, error) {
	return _KittyCore.Contract.Erc721Metadata(&_KittyCore.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyCore *KittyCoreCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyCore.Contract.Erc721Metadata(&_KittyCore.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "gen0CreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.Gen0CreatedCount(&_KittyCore.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.Gen0CreatedCount(&_KittyCore.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyCore *KittyCoreCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyCore *KittyCoreSession) GeneScience() (common.Address, error) {
	return _KittyCore.Contract.GeneScience(&_KittyCore.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyCore *KittyCoreCallerSession) GeneScience() (common.Address, error) {
	return _KittyCore.Contract.GeneScience(&_KittyCore.CallOpts)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_KittyCore *KittyCoreCaller) GetKitty(opts *bind.CallOpts, _id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "getKitty", _id)

	outstruct := new(struct {
		IsGestating   bool
		IsReady       bool
		CooldownIndex *big.Int
		NextActionAt  *big.Int
		SiringWithId  *big.Int
		BirthTime     *big.Int
		MatronId      *big.Int
		SireId        *big.Int
		Generation    *big.Int
		Genes         *big.Int
	})

	outstruct.IsGestating = out[0].(bool)
	outstruct.IsReady = out[1].(bool)
	outstruct.CooldownIndex = out[2].(*big.Int)
	outstruct.NextActionAt = out[3].(*big.Int)
	outstruct.SiringWithId = out[4].(*big.Int)
	outstruct.BirthTime = out[5].(*big.Int)
	outstruct.MatronId = out[6].(*big.Int)
	outstruct.SireId = out[7].(*big.Int)
	outstruct.Generation = out[8].(*big.Int)
	outstruct.Genes = out[9].(*big.Int)

	return *outstruct, err

}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_KittyCore *KittyCoreSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _KittyCore.Contract.GetKitty(&_KittyCore.CallOpts, _id)
}

// GetKitty is a free data retrieval call binding the contract method 0xe98b7f4d.
//
// Solidity: function getKitty(uint256 _id) view returns(bool isGestating, bool isReady, uint256 cooldownIndex, uint256 nextActionAt, uint256 siringWithId, uint256 birthTime, uint256 matronId, uint256 sireId, uint256 generation, uint256 genes)
func (_KittyCore *KittyCoreCallerSession) GetKitty(_id *big.Int) (struct {
	IsGestating   bool
	IsReady       bool
	CooldownIndex *big.Int
	NextActionAt  *big.Int
	SiringWithId  *big.Int
	BirthTime     *big.Int
	MatronId      *big.Int
	SireId        *big.Int
	Generation    *big.Int
	Genes         *big.Int
}, error) {
	return _KittyCore.Contract.GetKitty(&_KittyCore.CallOpts, _id)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "isPregnant", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsPregnant(&_KittyCore.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsPregnant(&_KittyCore.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "isReadyToBreed", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsReadyToBreed(&_KittyCore.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyCore *KittyCoreCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyCore.Contract.IsReadyToBreed(&_KittyCore.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyCore *KittyCoreSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToApproved(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToApproved(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyCore *KittyCoreSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToOwner(&_KittyCore.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.KittyIndexToOwner(&_KittyCore.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyCore *KittyCoreCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyCore *KittyCoreSession) Name() (string, error) {
	return _KittyCore.Contract.Name(&_KittyCore.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyCore *KittyCoreCallerSession) Name() (string, error) {
	return _KittyCore.Contract.Name(&_KittyCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_KittyCore *KittyCoreCaller) NewContractAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "newContractAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_KittyCore *KittyCoreSession) NewContractAddress() (common.Address, error) {
	return _KittyCore.Contract.NewContractAddress(&_KittyCore.CallOpts)
}

// NewContractAddress is a free data retrieval call binding the contract method 0x6af04a57.
//
// Solidity: function newContractAddress() view returns(address)
func (_KittyCore *KittyCoreCallerSession) NewContractAddress() (common.Address, error) {
	return _KittyCore.Contract.NewContractAddress(&_KittyCore.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyCore *KittyCoreCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyCore *KittyCoreSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyCore.Contract.OwnerOf(&_KittyCore.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyCore *KittyCoreCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyCore.Contract.OwnerOf(&_KittyCore.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyCore *KittyCoreCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyCore *KittyCoreSession) Paused() (bool, error) {
	return _KittyCore.Contract.Paused(&_KittyCore.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyCore *KittyCoreCallerSession) Paused() (bool, error) {
	return _KittyCore.Contract.Paused(&_KittyCore.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyCore *KittyCoreCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "pregnantKitties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyCore *KittyCoreSession) PregnantKitties() (*big.Int, error) {
	return _KittyCore.Contract.PregnantKitties(&_KittyCore.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyCore.Contract.PregnantKitties(&_KittyCore.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "promoCreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.PromoCreatedCount(&_KittyCore.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyCore.Contract.PromoCreatedCount(&_KittyCore.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyCore *KittyCoreCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyCore *KittyCoreSession) SaleAuction() (common.Address, error) {
	return _KittyCore.Contract.SaleAuction(&_KittyCore.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyCore *KittyCoreCallerSession) SaleAuction() (common.Address, error) {
	return _KittyCore.Contract.SaleAuction(&_KittyCore.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyCore *KittyCoreCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyCore *KittyCoreSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyCore.Contract.SecondsPerBlock(&_KittyCore.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyCore.Contract.SecondsPerBlock(&_KittyCore.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyCore *KittyCoreSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.SireAllowedToAddress(&_KittyCore.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyCore *KittyCoreCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyCore.Contract.SireAllowedToAddress(&_KittyCore.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyCore *KittyCoreCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyCore *KittyCoreSession) SiringAuction() (common.Address, error) {
	return _KittyCore.Contract.SiringAuction(&_KittyCore.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyCore *KittyCoreCallerSession) SiringAuction() (common.Address, error) {
	return _KittyCore.Contract.SiringAuction(&_KittyCore.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyCore *KittyCoreCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyCore *KittyCoreSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyCore.Contract.SupportsInterface(&_KittyCore.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyCore *KittyCoreCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyCore.Contract.SupportsInterface(&_KittyCore.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyCore *KittyCoreCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyCore *KittyCoreSession) Symbol() (string, error) {
	return _KittyCore.Contract.Symbol(&_KittyCore.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyCore *KittyCoreCallerSession) Symbol() (string, error) {
	return _KittyCore.Contract.Symbol(&_KittyCore.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyCore *KittyCoreCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyCore *KittyCoreSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyCore.Contract.TokenMetadata(&_KittyCore.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyCore *KittyCoreCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyCore.Contract.TokenMetadata(&_KittyCore.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyCore *KittyCoreCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyCore *KittyCoreSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyCore.Contract.TokensOfOwner(&_KittyCore.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyCore *KittyCoreCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyCore.Contract.TokensOfOwner(&_KittyCore.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyCore *KittyCoreCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyCore.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyCore *KittyCoreSession) TotalSupply() (*big.Int, error) {
	return _KittyCore.Contract.TotalSupply(&_KittyCore.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyCore *KittyCoreCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyCore.Contract.TotalSupply(&_KittyCore.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Approve(&_KittyCore.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Approve(&_KittyCore.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyCore *KittyCoreTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyCore *KittyCoreSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.ApproveSiring(&_KittyCore.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyCore *KittyCoreTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.ApproveSiring(&_KittyCore.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyCore *KittyCoreTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyCore *KittyCoreSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BidOnSiringAuction(&_KittyCore.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyCore *KittyCoreTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BidOnSiringAuction(&_KittyCore.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyCore *KittyCoreTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyCore *KittyCoreSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BreedWithAuto(&_KittyCore.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyCore *KittyCoreTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.BreedWithAuto(&_KittyCore.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyCore *KittyCoreTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyCore *KittyCoreSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateGen0Auction(&_KittyCore.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateGen0Auction(&_KittyCore.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyCore *KittyCoreTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyCore *KittyCoreSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.CreatePromoKitty(&_KittyCore.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyCore *KittyCoreTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.CreatePromoKitty(&_KittyCore.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSaleAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSaleAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSiringAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyCore *KittyCoreTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.CreateSiringAuction(&_KittyCore.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyCore *KittyCoreTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyCore *KittyCoreSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.GiveBirth(&_KittyCore.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyCore *KittyCoreTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.GiveBirth(&_KittyCore.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreSession) Pause() (*types.Transaction, error) {
	return _KittyCore.Contract.Pause(&_KittyCore.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyCore *KittyCoreTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyCore.Contract.Pause(&_KittyCore.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyCore *KittyCoreTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyCore *KittyCoreSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetAutoBirthFee(&_KittyCore.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyCore *KittyCoreTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetAutoBirthFee(&_KittyCore.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyCore *KittyCoreTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyCore *KittyCoreSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCEO(&_KittyCore.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCEO(&_KittyCore.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyCore *KittyCoreTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyCore *KittyCoreSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCFO(&_KittyCore.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCFO(&_KittyCore.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyCore *KittyCoreTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyCore *KittyCoreSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCOO(&_KittyCore.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyCore *KittyCoreTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetCOO(&_KittyCore.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyCore *KittyCoreSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetGeneScienceAddress(&_KittyCore.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetGeneScienceAddress(&_KittyCore.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyCore *KittyCoreTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyCore *KittyCoreSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetMetadataAddress(&_KittyCore.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyCore *KittyCoreTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetMetadataAddress(&_KittyCore.TransactOpts, _contractAddress)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_KittyCore *KittyCoreTransactor) SetNewAddress(opts *bind.TransactOpts, _v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setNewAddress", _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_KittyCore *KittyCoreSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetNewAddress(&_KittyCore.TransactOpts, _v2Address)
}

// SetNewAddress is a paid mutator transaction binding the contract method 0x71587988.
//
// Solidity: function setNewAddress(address _v2Address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetNewAddress(_v2Address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetNewAddress(&_KittyCore.TransactOpts, _v2Address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSaleAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSaleAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyCore *KittyCoreTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyCore *KittyCoreSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSecondsPerBlock(&_KittyCore.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSecondsPerBlock(&_KittyCore.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSiringAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyCore *KittyCoreTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyCore.Contract.SetSiringAuctionAddress(&_KittyCore.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Transfer(&_KittyCore.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.Transfer(&_KittyCore.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.TransferFrom(&_KittyCore.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyCore *KittyCoreTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyCore.Contract.TransferFrom(&_KittyCore.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreSession) Unpause() (*types.Transaction, error) {
	return _KittyCore.Contract.Unpause(&_KittyCore.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyCore *KittyCoreTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyCore.Contract.Unpause(&_KittyCore.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawAuctionBalances(&_KittyCore.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyCore *KittyCoreTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawAuctionBalances(&_KittyCore.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyCore.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreSession) WithdrawBalance() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawBalance(&_KittyCore.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_KittyCore *KittyCoreTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _KittyCore.Contract.WithdrawBalance(&_KittyCore.TransactOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KittyCore *KittyCoreTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _KittyCore.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KittyCore *KittyCoreSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _KittyCore.Contract.Fallback(&_KittyCore.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_KittyCore *KittyCoreTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _KittyCore.Contract.Fallback(&_KittyCore.TransactOpts, calldata)
}

// KittyCoreApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyCore contract.
type KittyCoreApprovalIterator struct {
	Event *KittyCoreApproval // Event containing the contract specifics and raw log

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
func (it *KittyCoreApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreApproval)
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
		it.Event = new(KittyCoreApproval)
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
func (it *KittyCoreApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreApproval represents a Approval event raised by the KittyCore contract.
type KittyCoreApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyCore *KittyCoreFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyCoreApprovalIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyCoreApprovalIterator{contract: _KittyCore.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyCore *KittyCoreFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyCoreApproval) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreApproval)
				if err := _KittyCore.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KittyCore *KittyCoreFilterer) ParseApproval(log types.Log) (*KittyCoreApproval, error) {
	event := new(KittyCoreApproval)
	if err := _KittyCore.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyCoreBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyCore contract.
type KittyCoreBirthIterator struct {
	Event *KittyCoreBirth // Event containing the contract specifics and raw log

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
func (it *KittyCoreBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreBirth)
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
		it.Event = new(KittyCoreBirth)
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
func (it *KittyCoreBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreBirth represents a Birth event raised by the KittyCore contract.
type KittyCoreBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyCore *KittyCoreFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyCoreBirthIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyCoreBirthIterator{contract: _KittyCore.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyCore *KittyCoreFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyCoreBirth) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreBirth)
				if err := _KittyCore.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyCore *KittyCoreFilterer) ParseBirth(log types.Log) (*KittyCoreBirth, error) {
	event := new(KittyCoreBirth)
	if err := _KittyCore.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyCoreContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyCore contract.
type KittyCoreContractUpgradeIterator struct {
	Event *KittyCoreContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyCoreContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreContractUpgrade)
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
		it.Event = new(KittyCoreContractUpgrade)
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
func (it *KittyCoreContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreContractUpgrade represents a ContractUpgrade event raised by the KittyCore contract.
type KittyCoreContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyCore *KittyCoreFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyCoreContractUpgradeIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyCoreContractUpgradeIterator{contract: _KittyCore.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyCore *KittyCoreFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyCoreContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreContractUpgrade)
				if err := _KittyCore.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyCore *KittyCoreFilterer) ParseContractUpgrade(log types.Log) (*KittyCoreContractUpgrade, error) {
	event := new(KittyCoreContractUpgrade)
	if err := _KittyCore.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyCorePregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyCore contract.
type KittyCorePregnantIterator struct {
	Event *KittyCorePregnant // Event containing the contract specifics and raw log

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
func (it *KittyCorePregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCorePregnant)
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
		it.Event = new(KittyCorePregnant)
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
func (it *KittyCorePregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCorePregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCorePregnant represents a Pregnant event raised by the KittyCore contract.
type KittyCorePregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyCore *KittyCoreFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyCorePregnantIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyCorePregnantIterator{contract: _KittyCore.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyCore *KittyCoreFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyCorePregnant) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCorePregnant)
				if err := _KittyCore.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// ParsePregnant is a log parse operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyCore *KittyCoreFilterer) ParsePregnant(log types.Log) (*KittyCorePregnant, error) {
	event := new(KittyCorePregnant)
	if err := _KittyCore.contract.UnpackLog(event, "Pregnant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyCoreTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyCore contract.
type KittyCoreTransferIterator struct {
	Event *KittyCoreTransfer // Event containing the contract specifics and raw log

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
func (it *KittyCoreTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyCoreTransfer)
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
		it.Event = new(KittyCoreTransfer)
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
func (it *KittyCoreTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyCoreTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyCoreTransfer represents a Transfer event raised by the KittyCore contract.
type KittyCoreTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyCore *KittyCoreFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyCoreTransferIterator, error) {

	logs, sub, err := _KittyCore.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyCoreTransferIterator{contract: _KittyCore.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyCore *KittyCoreFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyCoreTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyCore.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyCoreTransfer)
				if err := _KittyCore.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyCore *KittyCoreFilterer) ParseTransfer(log types.Log) (*KittyCoreTransfer, error) {
	event := new(KittyCoreTransfer)
	if err := _KittyCore.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyMintingABI is the input ABI used to generate the binding from.
const KittyMintingABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"promoCreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_STARTING_PRICE\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSiringAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"pregnantKitties\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isPregnant\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_AUCTION_DURATION\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setGeneScienceAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSaleAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"canBreedWith\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"}],\"name\":\"createSiringAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"val\",\"type\":\"uint256\"}],\"name\":\"setAutoBirthFee\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_addr\",\"type\":\"address\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"approveSiring\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"},{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"createPromoKitty\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GEN0_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"setSaleAuctionAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"giveBirth\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawAuctionBalances\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"autoBirthFee\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_genes\",\"type\":\"uint256\"}],\"name\":\"createGen0Auction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_kittyId\",\"type\":\"uint256\"}],\"name\":\"isReadyToBreed\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"PROMO_CREATION_LIMIT\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_sireId\",\"type\":\"uint256\"},{\"name\":\"_matronId\",\"type\":\"uint256\"}],\"name\":\"bidOnSiringAuction\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"gen0CreatedCount\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"geneScience\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_matronId\",\"type\":\"uint256\"},{\"name\":\"_sireId\",\"type\":\"uint256\"}],\"name\":\"breedWithAuto\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"cooldownEndBlock\",\"type\":\"uint256\"}],\"name\":\"Pregnant\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyMintingFuncSigs maps the 4-byte function signature to its string representation.
var KittyMintingFuncSigs = map[string]string{
	"19c2f201": "GEN0_AUCTION_DURATION()",
	"680eba27": "GEN0_CREATION_LIMIT()",
	"0e583df0": "GEN0_STARTING_PRICE()",
	"defb9584": "PROMO_CREATION_LIMIT()",
	"095ea7b3": "approve(address,uint256)",
	"4dfff04f": "approveSiring(address,uint256)",
	"b0c35c05": "autoBirthFee()",
	"70a08231": "balanceOf(address)",
	"ed60ade6": "bidOnSiringAuction(uint256,uint256)",
	"f7d8c883": "breedWithAuto(uint256,uint256)",
	"46d22c70": "canBreedWith(uint256,uint256)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"c3bea9af": "createGen0Auction(uint256)",
	"56129134": "createPromoKitty(uint256,address)",
	"3d7d3f5a": "createSaleAuction(uint256,uint256,uint256,uint256)",
	"4ad8c938": "createSiringAuction(uint256,uint256,uint256,uint256)",
	"bc4006f5": "erc721Metadata()",
	"f1ca9410": "gen0CreatedCount()",
	"f2b47d52": "geneScience()",
	"88c2a0bf": "giveBirth(uint256)",
	"1940a936": "isPregnant(uint256)",
	"d3e6f49f": "isReadyToBreed(uint256)",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"183a7947": "pregnantKitties()",
	"05e45546": "promoCreatedCount()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"4b85fd55": "setAutoBirthFee(uint256)",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"24e7a38a": "setGeneScienceAddress(address)",
	"e17b25af": "setMetadataAddress(address)",
	"6fbde40d": "setSaleAuctionAddress(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"14001f4c": "setSiringAuctionAddress(address)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"0560ff44": "tokenMetadata(uint256,string)",
	"8462151c": "tokensOfOwner(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
	"91876e57": "withdrawAuctionBalances()",
}

// KittyMintingBin is the compiled bytecode used for deploying new contracts.
var KittyMintingBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a0820152620000a790600390600e620000ca565b50600f60055566071afd498d0000600e553415620000c457600080fd5b62000194565b6002830191839082156200015b5791602002820160005b838211156200012757835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000e1565b8015620001595782816101000a81549063ffffffff021916905560040160208160030104928301926001030262000127565b505b50620001699291506200016d565b5090565b6200019191905b808211156200016957805463ffffffff1916815560010162000174565b90565b612c1280620001a46000396000f3006060604052600436106102795763ffffffff60e060020a60003504166301ffc9a7811461027e5780630519ce79146102ca5780630560ff44146102f957806305e455461461039257806306fdde03146103b7578063095ea7b3146103ca5780630a0f8168146103ee5780630e583df01461040157806314001f4c1461041457806318160ddd14610433578063183a7947146104465780631940a9361461045957806319c2f2011461046f57806321717ebf1461048257806323b872dd1461049557806324e7a38a146104bd57806327d7874c146104dc5780632ba73c15146104fb5780633d7d3f5a1461051a5780633f4ba83a1461053957806346116e6f1461054c57806346d22c7014610562578063481af3d31461057b5780634ad8c938146105915780634b85fd55146105b05780634dfff04f146105c65780634e0a3379146105e857806356129134146106075780635663896e146106295780635c975abb1461063f5780636352211e14610652578063680eba27146106685780636fbde40d1461067b57806370a082311461069a5780637a7d4937146106b95780638456cb59146106cc5780638462151c146106df57806388c2a0bf1461075157806391876e571461076757806395d89b411461077a5780639d6fac6f1461078d578063a45f4bfc146107bc578063a9059cbb146107d2578063b047fb50146107f4578063b0c35c0514610807578063bc4006f51461081a578063c3bea9af1461082d578063d3e6f49f14610843578063defb958414610859578063e17b25af1461086c578063e6cbe3511461088b578063ed60ade61461089e578063f1ca9410146108ac578063f2b47d52146108bf578063f7d8c883146108d2575b600080fd5b341561028957600080fd5b6102b67fffffffff00000000000000000000000000000000000000000000000000000000600435166108e0565b604051901515815260200160405180910390f35b34156102d557600080fd5b6102dd610b67565b604051600160a060020a03909116815260200160405180910390f35b341561030457600080fd5b61031b600480359060248035908101910135610b76565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561035757808201518382015260200161033f565b50505050905090810190601f1680156103845780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b341561039d57600080fd5b6103a5610c46565b60405190815260200160405180910390f35b34156103c257600080fd5b61031b610c4c565b34156103d557600080fd5b6103ec600160a060020a0360043516602435610c83565b005b34156103f957600080fd5b6102dd610d0d565b341561040c57600080fd5b6103a5610d1c565b341561041f57600080fd5b6103ec600160a060020a0360043516610d27565b341561043e57600080fd5b6103a5610dc7565b341561045157600080fd5b6103a5610dd2565b341561046457600080fd5b6102b6600435610dd8565b341561047a57600080fd5b6103a5610e1d565b341561048d57600080fd5b6102dd610e24565b34156104a057600080fd5b6103ec600160a060020a0360043581169060243516604435610e33565b34156104c857600080fd5b6103ec600160a060020a0360043516610eba565b34156104e757600080fd5b6103ec600160a060020a0360043516610f5a565b341561050657600080fd5b6103ec600160a060020a0360043516610fac565b341561052557600080fd5b6103ec600435602435604435606435610ffe565b341561054457600080fd5b6103ec6110d9565b341561055757600080fd5b6102dd60043561112c565b341561056d57600080fd5b6102b6600435602435611147565b341561058657600080fd5b6102dd6004356111c7565b341561059c57600080fd5b6103ec6004356024356044356064356111e2565b34156105bb57600080fd5b6103ec6004356112a8565b34156105d157600080fd5b6103ec600160a060020a03600435166024356112c8565b34156105f357600080fd5b6103ec600160a060020a0360043516611322565b341561061257600080fd5b6103ec600435600160a060020a0360243516611374565b341561063457600080fd5b6103ec6004356113e0565b341561064a57600080fd5b6102b6611448565b341561065d57600080fd5b6102dd600435611458565b341561067357600080fd5b6103a561147c565b341561068657600080fd5b6103ec600160a060020a0360043516611482565b34156106a557600080fd5b6103a5600160a060020a0360043516611522565b34156106c457600080fd5b6103a561153d565b34156106d757600080fd5b6103ec611543565b34156106ea57600080fd5b6106fe600160a060020a03600435166115cf565b60405160208082528190810183818151815260200191508051906020019060200280838360005b8381101561073d578082015183820152602001610725565b505050509050019250505060405180910390f35b341561075c57600080fd5b6103a56004356116b0565b341561077257600080fd5b6103ec611973565b341561078557600080fd5b61031b611a5e565b341561079857600080fd5b6107a3600435611a95565b60405163ffffffff909116815260200160405180910390f35b34156107c757600080fd5b6102dd600435611ac2565b34156107dd57600080fd5b6103ec600160a060020a0360043516602435611add565b34156107ff57600080fd5b6102dd611b84565b341561081257600080fd5b6103a5611b93565b341561082557600080fd5b6102dd611b99565b341561083857600080fd5b6103ec600435611ba8565b341561084e57600080fd5b6102b6600435611c97565b341561086457600080fd5b6103a5611d60565b341561087757600080fd5b6103ec600160a060020a0360043516611d66565b341561089657600080fd5b6102dd611da3565b6103ec600435602435611db2565b34156108b757600080fd5b6103a5611ef4565b34156108ca57600080fd5b6102dd611efa565b6103ec600435602435611f09565b60006040517f737570706f727473496e7465726661636528627974657334290000000000000081526019016040518091039020600160e060020a03191682600160e060020a0319161480610b5f57506040517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152601d0160405180910390206040517f746f6b656e734f664f776e657228616464726573732900000000000000000000815260160160405180910390206040517f7472616e7366657246726f6d28616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390206040517f7472616e7366657228616464726573732c75696e743235362900000000000000815260190160405180910390206040517f617070726f766528616464726573732c75696e74323536290000000000000000815260180160405180910390206040517f6f776e65724f662875696e743235362900000000000000000000000000000000815260100160405180910390206040517f62616c616e63654f662861646472657373290000000000000000000000000000815260120160405180910390206040517f746f74616c537570706c792829000000000000000000000000000000000000008152600d0160405180910390206040517f73796d626f6c2829000000000000000000000000000000000000000000000000815260080160405180910390206040517f6e616d652829000000000000000000000000000000000000000000000000000081526006016040518091039020181818181818181818600160e060020a03191682600160e060020a031916145b90505b919050565b600154600160a060020a031681565b610b7e612b1f565b610b86612b31565b600d54600090600160a060020a03161515610ba057600080fd5b600d54600160a060020a031663cb4799f287878760405160e060020a63ffffffff861602815260048101848152604060248301908152604483018490529091606401848480828437820191505094505050505060a060405180830381600087803b1515610c0c57600080fd5b5af11515610c1957600080fd5b50505060405180608001805160209091016040529092509050610c3c82826120fb565b9695505050505050565b60115481565b60408051908101604052600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff1615610c9a57600080fd5b610ca43382612150565b1515610caf57600080fd5b610cb98183612170565b7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600054600160a060020a031681565b662386f26fc1000081565b6000805433600160a060020a03908116911614610d4357600080fd5b5080600160a060020a0381166376190f8f6040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610d8257600080fd5b5af11515610d8f57600080fd5b505050604051805190501515610da457600080fd5b600c8054600160a060020a031916600160a060020a039290921691909117905550565b600654600019015b90565b600f5481565b6000808211610de657600080fd5b6006805483908110610df457fe5b600091825260209091206002909102016001015460c060020a900463ffffffff16151592915050565b6201518081565b600c54600160a060020a031681565b60025460a060020a900460ff1615610e4a57600080fd5b600160a060020a0382161515610e5f57600080fd5b30600160a060020a031682600160a060020a031614151515610e8057600080fd5b610e8a338261219e565b1515610e9557600080fd5b610e9f8382612150565b1515610eaa57600080fd5b610eb58383836121be565b505050565b6000805433600160a060020a03908116911614610ed657600080fd5b5080600160a060020a0381166354c15b826040518163ffffffff1660e060020a028152600401602060405180830381600087803b1515610f1557600080fd5b5af11515610f2257600080fd5b505050604051805190501515610f3757600080fd5b60108054600160a060020a031916600160a060020a039290921691909117905550565b60005433600160a060020a03908116911614610f7557600080fd5b600160a060020a0381161515610f8a57600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610fc757600080fd5b600160a060020a0381161515610fdc57600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60025460a060020a900460ff161561101557600080fd5b61101f3385612150565b151561102a57600080fd5b61103384610dd8565b1561103d57600080fd5b600b54611054908590600160a060020a0316612170565b600b54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b15156110c357600080fd5b5af115156110d057600080fd5b50505050505050565b60005433600160a060020a039081169116146110f457600080fd5b60025460a060020a900460ff16151561110c57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b6000808080851161115757600080fd5b6000841161116457600080fd5b600680548690811061117257fe5b9060005260206000209060020201915060068481548110151561119157fe5b906000526020600020906002020190506111ad828683876122a6565b80156111be57506111be8486612426565b95945050505050565b600960205260009081526040902054600160a060020a031681565b60025460a060020a900460ff16156111f957600080fd5b6112033385612150565b151561120e57600080fd5b61121784611c97565b151561122257600080fd5b600c54611239908590600160a060020a0316612170565b600c54600160a060020a03166327ebe40a858585853360405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b15156110c357600080fd5b60025433600160a060020a039081169116146112c357600080fd5b600e55565b60025460a060020a900460ff16156112df57600080fd5b6112e93382612150565b15156112f457600080fd5b6000908152600a602052604090208054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a0390811691161461133d57600080fd5b600160a060020a038116151561135257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60025460009033600160a060020a0390811691161461139257600080fd5b5080600160a060020a03811615156113b25750600254600160a060020a03165b60115461138890106113c357600080fd5b6011805460010190556113da60008080868561247b565b50505050565b60025433600160a060020a039081169116148061140b575060005433600160a060020a039081169116145b80611424575060015433600160a060020a039081169116145b151561142f57600080fd5b60035463ffffffff16811061144357600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a0316801515610b6257600080fd5b61afc881565b6000805433600160a060020a0390811691161461149e57600080fd5b5080600160a060020a0381166385b861886040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156114dd57600080fd5b5af115156114ea57600080fd5b5050506040518051905015156114ff57600080fd5b600b8054600160a060020a031916600160a060020a039290921691909117905550565b600160a060020a031660009081526008602052604090205490565b60055481565b60025433600160a060020a039081169116148061156e575060005433600160a060020a039081169116145b80611587575060015433600160a060020a039081169116145b151561159257600080fd5b60025460a060020a900460ff16156115a957600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b6115d7612b1f565b60006115e1612b1f565b60008060006115ef87611522565b945084151561161f5760006040518059106116075750595b908082528060200260200182016040525095506116a6565b8460405180591061162d5750595b90808252806020026020018201604052509350611648610dc7565b925060009150600190505b8281116116a257600081815260076020526040902054600160a060020a038881169116141561169a578084838151811061168957fe5b602090810290910101526001909101905b600101611653565b8395505b5050505050919050565b600080600080600080600080600260149054906101000a900460ff161515156116d857600080fd5b600680548a9081106116e657fe5b60009182526020909120600290910201600181015490975067ffffffffffffffff16151561171357600080fd5b6117a88761010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612727565b15156117b357600080fd5b60018701546006805460c060020a90920463ffffffff16975090879081106117d757fe5b600091825260209091206001808a015460029093029091019081015490965061ffff60f060020a9283900481169650919004168490111561182557600185015460f060020a900461ffff1693505b6010548754865460018a0154600160a060020a0390931692630d9f5aed92919068010000000000000000900467ffffffffffffffff166000190160405160e060020a63ffffffff86160281526004810193909352602483019190915267ffffffffffffffff166044820152606401602060405180830381600087803b15156118ac57600080fd5b5af115156118b957600080fd5b505050604051805160008b81526007602052604090205460018a810154929650600160a060020a03909116945061190892508b9160c060020a900463ffffffff1690870161ffff16868661247b565b6001880180547bffffffff00000000000000000000000000000000000000000000000019169055600f8054600019019055600e54909150600160a060020a0333169080156108fc0290604051600060405180830381858888f150939c9b505050505050505050505050565b60025433600160a060020a039081169116148061199e575060005433600160a060020a039081169116145b806119b7575060015433600160a060020a039081169116145b15156119c257600080fd5b600b54600160a060020a0316635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b1515611a0157600080fd5b5af11515611a0e57600080fd5b5050600c54600160a060020a03169050635fd8c7106040518163ffffffff1660e060020a028152600401600060405180830381600087803b1515611a5157600080fd5b5af11515610eb557600080fd5b60408051908101604052600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110611aa257fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615611af457600080fd5b600160a060020a0382161515611b0957600080fd5b30600160a060020a031682600160a060020a031614151515611b2a57600080fd5b600b54600160a060020a0383811691161415611b4557600080fd5b600c54600160a060020a0383811691161415611b6057600080fd5b611b6a3382612150565b1515611b7557600080fd5b611b803383836121be565b5050565b600254600160a060020a031681565b600e5481565b600d54600160a060020a031681565b60025460009033600160a060020a03908116911614611bc657600080fd5b60125461afc89010611bd757600080fd5b611be66000806000853061247b565b600b54909150611c00908290600160a060020a0316612170565b600b54600160a060020a03166327ebe40a82611c1a61275f565b6000620151803060405160e060020a63ffffffff88160281526004810195909552602485019390935260448401919091526064830152600160a060020a0316608482015260a401600060405180830381600087803b1515611c7a57600080fd5b5af11515611c8757600080fd5b5050601280546001019055505050565b600080808311611ca657600080fd5b6006805484908110611cb457fe5b90600052602060002090600202019050611d598161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612801565b9392505050565b61138881565b60005433600160a060020a03908116911614611d8157600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b60025460009060a060020a900460ff1615611dcc57600080fd5b611dd63383612150565b1515611de157600080fd5b611dea82611c97565b1515611df557600080fd5b611dff8284612838565b1515611e0a57600080fd5b600c54600160a060020a031663c55d0f568460405160e060020a63ffffffff84160281526004810191909152602401602060405180830381600087803b1515611e5257600080fd5b5af11515611e5f57600080fd5b5050506040518051600e5490925082013410159050611e7d57600080fd5b600c54600e54600160a060020a039091169063454a2ab39034038560405160e060020a63ffffffff851602815260048101919091526024016000604051808303818588803b1515611ecd57600080fd5b5af11515611eda57600080fd5b50505050610eb58263ffffffff168463ffffffff16612887565b60125481565b601054600160a060020a031681565b600254600090819060a060020a900460ff1615611f2557600080fd5b600e54341015611f3457600080fd5b611f3e3385612150565b1515611f4957600080fd5b611f538385612426565b1515611f5e57600080fd5b6006805485908110611f6c57fe5b906000526020600020906002020191506120118261010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612801565b151561201c57600080fd5b600680548490811061202a57fe5b906000526020600020906002020190506120cf8161010060405190810160409081528254825260019092015467ffffffffffffffff8082166020840152680100000000000000008204169282019290925263ffffffff608060020a83048116606083015260a060020a83048116608083015260c060020a83041660a082015261ffff60e060020a8304811660c083015260f060020a90920490911660e0820152612801565b15156120da57600080fd5b6120e6828583866122a6565b15156120f157600080fd5b6113da8484612887565b612103612b1f565b61210b612b1f565b6000808460405180591061211c5750595b818152601f19601f83011681016020016040529050925050602082019050846121468282876129f1565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561225157600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b6000818414156122b85750600061241e565b6001850154608060020a900463ffffffff168214806122e75750600185015460a060020a900463ffffffff1682145b156122f45750600061241e565b6001830154608060020a900463ffffffff168414806123235750600183015460a060020a900463ffffffff1684145b156123305750600061241e565b6001830154608060020a900463ffffffff16158061235d57506001850154608060020a900463ffffffff16155b1561236a5750600161241e565b60018581015490840154608060020a9182900463ffffffff908116929091041614806123b5575060018086015490840154608060020a900463ffffffff90811660a060020a90920416145b156123c25750600061241e565b6001808601549084015460a060020a900463ffffffff908116608060020a90920416148061240d57506001858101549084015460a060020a9182900463ffffffff9081169290910416145b1561241a5750600061241e565b5060015b949350505050565b6000818152600760205260408082205484835290822054600160a060020a039182169116808214806111be57506000858152600a6020526040902054600160a060020a03908116908316149250505092915050565b600080612486612b5a565b600063ffffffff8916891461249a57600080fd5b63ffffffff881688146124ac57600080fd5b61ffff871687146124bc57600080fd5b600287049250600d8361ffff1611156124d457600d92505b610100604051908101604090815287825267ffffffffffffffff42166020830152600090820181905263ffffffff808c1660608401528a16608083015260a082015261ffff80851660c0830152881660e08201526006805491935060019180830161253f8382612b9e565b6000928352602090922085916002020181518155602082015160018201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560408201518160010160086101000a81548167ffffffffffffffff021916908367ffffffffffffffff16021790555060608201518160010160106101000a81548163ffffffff021916908363ffffffff16021790555060808201518160010160146101000a81548163ffffffff021916908363ffffffff16021790555060a08201518160010160186101000a81548163ffffffff021916908363ffffffff16021790555060c082015181600101601c6101000a81548161ffff021916908361ffff16021790555060e08201516001909101805461ffff9290921660f060020a027dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555003905063ffffffff8116811461269a57600080fd5b7f0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad58582846060015163ffffffff16856080015163ffffffff168651604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390a161271b600086836121be565b98975050505050505050565b60008160a0015163ffffffff1615801590610b5f57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600b5460009081908190600160a060020a031663eac9d94c6040518163ffffffff1660e060020a028152600401602060405180830381600087803b15156127a557600080fd5b5af115156127b257600080fd5b50505060405180519250506fffffffffffffffffffffffffffffffff821682146127db57600080fd5b50600281048101662386f26fc100008110156127fb5750662386f26fc100005b92915050565b60008160a0015163ffffffff16158015610b5f57504367ffffffffffffffff16826040015167ffffffffffffffff16111592915050565b600080600060068581548110151561284c57fe5b9060005260206000209060020201915060068481548110151561286b57fe5b906000526020600020906002020190506111be828683876122a6565b60008060068381548110151561289957fe5b906000526020600020906002020191506006848154811015156128b857fe5b600091825260209091206002909102016001810180547bffffffff000000000000000000000000000000000000000000000000191660c060020a63ffffffff871602179055905061290882612a36565b61291181612a36565b6000848152600a602090815260408083208054600160a060020a031990811690915586845281842080549091169055600f8054600190810190915587845260079092529182902054908301547f241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b8092600160a060020a0390921691879187916801000000000000000090910467ffffffffffffffff1690518085600160a060020a0316600160a060020a031681526020018481526020018381526020018267ffffffffffffffff16815260200194505050505060405180910390a150505050565b60005b60208210612a1757825184526020840193506020830192506020820391506129f4565b6001826020036101000a03905080198351168185511617909352505050565b600554600182015443919060039060e060020a900461ffff16600e8110612a5957fe5b600891828204019190066004029054906101000a900463ffffffff1663ffffffff16811515612a8457fe5b6001840180546fffffffffffffffff0000000000000000191668010000000000000000939092049390930167ffffffffffffffff16919091021790819055600d60e060020a90910461ffff161015612b1c576001818101805461ffff60e060020a8083048216909401169092027fffff0000ffffffffffffffffffffffffffffffffffffffffffffffffffffffff9092169190911790555b50565b60206040519081016040526000815290565b60806040519081016040526004815b60008152600019919091019060200181612b405790505090565b6101006040519081016040908152600080835260208301819052908201819052606082018190526080820181905260a0820181905260c0820181905260e082015290565b815481835581811511610eb557600083815260209020610eb591610dcf9160029182028101918502015b80821115612be25760008082556001820155600201612bc8565b50905600a165627a7a72305820fef72fab47944956ff684690f1315b29324e5f8cbf3949370f53d5138314580f0029"

// DeployKittyMinting deploys a new Ethereum contract, binding an instance of KittyMinting to it.
func DeployKittyMinting(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyMinting, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyMintingABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyMintingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyMinting{KittyMintingCaller: KittyMintingCaller{contract: contract}, KittyMintingTransactor: KittyMintingTransactor{contract: contract}, KittyMintingFilterer: KittyMintingFilterer{contract: contract}}, nil
}

// KittyMinting is an auto generated Go binding around an Ethereum contract.
type KittyMinting struct {
	KittyMintingCaller     // Read-only binding to the contract
	KittyMintingTransactor // Write-only binding to the contract
	KittyMintingFilterer   // Log filterer for contract events
}

// KittyMintingCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyMintingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyMintingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyMintingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyMintingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyMintingSession struct {
	Contract     *KittyMinting     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyMintingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyMintingCallerSession struct {
	Contract *KittyMintingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// KittyMintingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyMintingTransactorSession struct {
	Contract     *KittyMintingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// KittyMintingRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyMintingRaw struct {
	Contract *KittyMinting // Generic contract binding to access the raw methods on
}

// KittyMintingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyMintingCallerRaw struct {
	Contract *KittyMintingCaller // Generic read-only contract binding to access the raw methods on
}

// KittyMintingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyMintingTransactorRaw struct {
	Contract *KittyMintingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyMinting creates a new instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMinting(address common.Address, backend bind.ContractBackend) (*KittyMinting, error) {
	contract, err := bindKittyMinting(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyMinting{KittyMintingCaller: KittyMintingCaller{contract: contract}, KittyMintingTransactor: KittyMintingTransactor{contract: contract}, KittyMintingFilterer: KittyMintingFilterer{contract: contract}}, nil
}

// NewKittyMintingCaller creates a new read-only instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingCaller(address common.Address, caller bind.ContractCaller) (*KittyMintingCaller, error) {
	contract, err := bindKittyMinting(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyMintingCaller{contract: contract}, nil
}

// NewKittyMintingTransactor creates a new write-only instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyMintingTransactor, error) {
	contract, err := bindKittyMinting(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyMintingTransactor{contract: contract}, nil
}

// NewKittyMintingFilterer creates a new log filterer instance of KittyMinting, bound to a specific deployed contract.
func NewKittyMintingFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyMintingFilterer, error) {
	contract, err := bindKittyMinting(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyMintingFilterer{contract: contract}, nil
}

// bindKittyMinting binds a generic wrapper to an already deployed contract.
func bindKittyMinting(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyMintingABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyMinting *KittyMintingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyMinting.Contract.KittyMintingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyMinting *KittyMintingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.Contract.KittyMintingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyMinting *KittyMintingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyMinting.Contract.KittyMintingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyMinting *KittyMintingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyMinting.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyMinting *KittyMintingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyMinting *KittyMintingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyMinting.Contract.contract.Transact(opts, method, params...)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0AUCTIONDURATION(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "GEN0_AUCTION_DURATION")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0AUCTIONDURATION(&_KittyMinting.CallOpts)
}

// GEN0AUCTIONDURATION is a free data retrieval call binding the contract method 0x19c2f201.
//
// Solidity: function GEN0_AUCTION_DURATION() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0AUCTIONDURATION() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0AUCTIONDURATION(&_KittyMinting.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0CREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "GEN0_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0CREATIONLIMIT(&_KittyMinting.CallOpts)
}

// GEN0CREATIONLIMIT is a free data retrieval call binding the contract method 0x680eba27.
//
// Solidity: function GEN0_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0CREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0CREATIONLIMIT(&_KittyMinting.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) GEN0STARTINGPRICE(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "GEN0_STARTING_PRICE")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyMinting *KittyMintingSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0STARTINGPRICE(&_KittyMinting.CallOpts)
}

// GEN0STARTINGPRICE is a free data retrieval call binding the contract method 0x0e583df0.
//
// Solidity: function GEN0_STARTING_PRICE() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) GEN0STARTINGPRICE() (*big.Int, error) {
	return _KittyMinting.Contract.GEN0STARTINGPRICE(&_KittyMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) PROMOCREATIONLIMIT(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "PROMO_CREATION_LIMIT")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.PROMOCREATIONLIMIT(&_KittyMinting.CallOpts)
}

// PROMOCREATIONLIMIT is a free data retrieval call binding the contract method 0xdefb9584.
//
// Solidity: function PROMO_CREATION_LIMIT() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PROMOCREATIONLIMIT() (*big.Int, error) {
	return _KittyMinting.Contract.PROMOCREATIONLIMIT(&_KittyMinting.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) AutoBirthFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "autoBirthFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyMinting *KittyMintingSession) AutoBirthFee() (*big.Int, error) {
	return _KittyMinting.Contract.AutoBirthFee(&_KittyMinting.CallOpts)
}

// AutoBirthFee is a free data retrieval call binding the contract method 0xb0c35c05.
//
// Solidity: function autoBirthFee() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) AutoBirthFee() (*big.Int, error) {
	return _KittyMinting.Contract.AutoBirthFee(&_KittyMinting.CallOpts)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyMinting *KittyMintingCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyMinting *KittyMintingSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyMinting.Contract.BalanceOf(&_KittyMinting.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyMinting *KittyMintingCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyMinting.Contract.BalanceOf(&_KittyMinting.CallOpts, _owner)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyMinting *KittyMintingCaller) CanBreedWith(opts *bind.CallOpts, _matronId *big.Int, _sireId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "canBreedWith", _matronId, _sireId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyMinting *KittyMintingSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyMinting.Contract.CanBreedWith(&_KittyMinting.CallOpts, _matronId, _sireId)
}

// CanBreedWith is a free data retrieval call binding the contract method 0x46d22c70.
//
// Solidity: function canBreedWith(uint256 _matronId, uint256 _sireId) view returns(bool)
func (_KittyMinting *KittyMintingCallerSession) CanBreedWith(_matronId *big.Int, _sireId *big.Int) (bool, error) {
	return _KittyMinting.Contract.CanBreedWith(&_KittyMinting.CallOpts, _matronId, _sireId)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyMinting *KittyMintingCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyMinting *KittyMintingSession) CeoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CeoAddress(&_KittyMinting.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) CeoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CeoAddress(&_KittyMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyMinting *KittyMintingCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyMinting *KittyMintingSession) CfoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CfoAddress(&_KittyMinting.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) CfoAddress() (common.Address, error) {
	return _KittyMinting.Contract.CfoAddress(&_KittyMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyMinting *KittyMintingCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyMinting *KittyMintingSession) CooAddress() (common.Address, error) {
	return _KittyMinting.Contract.CooAddress(&_KittyMinting.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) CooAddress() (common.Address, error) {
	return _KittyMinting.Contract.CooAddress(&_KittyMinting.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyMinting *KittyMintingCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyMinting *KittyMintingSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyMinting.Contract.Cooldowns(&_KittyMinting.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyMinting *KittyMintingCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyMinting.Contract.Cooldowns(&_KittyMinting.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyMinting *KittyMintingCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyMinting *KittyMintingSession) Erc721Metadata() (common.Address, error) {
	return _KittyMinting.Contract.Erc721Metadata(&_KittyMinting.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyMinting.Contract.Erc721Metadata(&_KittyMinting.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) Gen0CreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "gen0CreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.Gen0CreatedCount(&_KittyMinting.CallOpts)
}

// Gen0CreatedCount is a free data retrieval call binding the contract method 0xf1ca9410.
//
// Solidity: function gen0CreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) Gen0CreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.Gen0CreatedCount(&_KittyMinting.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyMinting *KittyMintingCaller) GeneScience(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "geneScience")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyMinting *KittyMintingSession) GeneScience() (common.Address, error) {
	return _KittyMinting.Contract.GeneScience(&_KittyMinting.CallOpts)
}

// GeneScience is a free data retrieval call binding the contract method 0xf2b47d52.
//
// Solidity: function geneScience() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) GeneScience() (common.Address, error) {
	return _KittyMinting.Contract.GeneScience(&_KittyMinting.CallOpts)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingCaller) IsPregnant(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "isPregnant", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsPregnant(&_KittyMinting.CallOpts, _kittyId)
}

// IsPregnant is a free data retrieval call binding the contract method 0x1940a936.
//
// Solidity: function isPregnant(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingCallerSession) IsPregnant(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsPregnant(&_KittyMinting.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingCaller) IsReadyToBreed(opts *bind.CallOpts, _kittyId *big.Int) (bool, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "isReadyToBreed", _kittyId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsReadyToBreed(&_KittyMinting.CallOpts, _kittyId)
}

// IsReadyToBreed is a free data retrieval call binding the contract method 0xd3e6f49f.
//
// Solidity: function isReadyToBreed(uint256 _kittyId) view returns(bool)
func (_KittyMinting *KittyMintingCallerSession) IsReadyToBreed(_kittyId *big.Int) (bool, error) {
	return _KittyMinting.Contract.IsReadyToBreed(&_KittyMinting.CallOpts, _kittyId)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToApproved(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToApproved(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToOwner(&_KittyMinting.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.KittyIndexToOwner(&_KittyMinting.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyMinting *KittyMintingCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyMinting *KittyMintingSession) Name() (string, error) {
	return _KittyMinting.Contract.Name(&_KittyMinting.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyMinting *KittyMintingCallerSession) Name() (string, error) {
	return _KittyMinting.Contract.Name(&_KittyMinting.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyMinting *KittyMintingCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyMinting *KittyMintingSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.OwnerOf(&_KittyMinting.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyMinting *KittyMintingCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.OwnerOf(&_KittyMinting.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyMinting *KittyMintingCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyMinting *KittyMintingSession) Paused() (bool, error) {
	return _KittyMinting.Contract.Paused(&_KittyMinting.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyMinting *KittyMintingCallerSession) Paused() (bool, error) {
	return _KittyMinting.Contract.Paused(&_KittyMinting.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) PregnantKitties(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "pregnantKitties")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyMinting *KittyMintingSession) PregnantKitties() (*big.Int, error) {
	return _KittyMinting.Contract.PregnantKitties(&_KittyMinting.CallOpts)
}

// PregnantKitties is a free data retrieval call binding the contract method 0x183a7947.
//
// Solidity: function pregnantKitties() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PregnantKitties() (*big.Int, error) {
	return _KittyMinting.Contract.PregnantKitties(&_KittyMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) PromoCreatedCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "promoCreatedCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.PromoCreatedCount(&_KittyMinting.CallOpts)
}

// PromoCreatedCount is a free data retrieval call binding the contract method 0x05e45546.
//
// Solidity: function promoCreatedCount() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) PromoCreatedCount() (*big.Int, error) {
	return _KittyMinting.Contract.PromoCreatedCount(&_KittyMinting.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyMinting *KittyMintingCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyMinting *KittyMintingSession) SaleAuction() (common.Address, error) {
	return _KittyMinting.Contract.SaleAuction(&_KittyMinting.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) SaleAuction() (common.Address, error) {
	return _KittyMinting.Contract.SaleAuction(&_KittyMinting.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyMinting *KittyMintingSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyMinting.Contract.SecondsPerBlock(&_KittyMinting.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyMinting.Contract.SecondsPerBlock(&_KittyMinting.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.SireAllowedToAddress(&_KittyMinting.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyMinting *KittyMintingCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyMinting.Contract.SireAllowedToAddress(&_KittyMinting.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyMinting *KittyMintingCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyMinting *KittyMintingSession) SiringAuction() (common.Address, error) {
	return _KittyMinting.Contract.SiringAuction(&_KittyMinting.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyMinting *KittyMintingCallerSession) SiringAuction() (common.Address, error) {
	return _KittyMinting.Contract.SiringAuction(&_KittyMinting.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyMinting *KittyMintingCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyMinting *KittyMintingSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyMinting.Contract.SupportsInterface(&_KittyMinting.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyMinting *KittyMintingCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyMinting.Contract.SupportsInterface(&_KittyMinting.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyMinting *KittyMintingCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyMinting *KittyMintingSession) Symbol() (string, error) {
	return _KittyMinting.Contract.Symbol(&_KittyMinting.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyMinting *KittyMintingCallerSession) Symbol() (string, error) {
	return _KittyMinting.Contract.Symbol(&_KittyMinting.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyMinting *KittyMintingCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyMinting *KittyMintingSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyMinting.Contract.TokenMetadata(&_KittyMinting.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyMinting *KittyMintingCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyMinting.Contract.TokenMetadata(&_KittyMinting.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyMinting *KittyMintingCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyMinting *KittyMintingSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyMinting.Contract.TokensOfOwner(&_KittyMinting.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyMinting *KittyMintingCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyMinting.Contract.TokensOfOwner(&_KittyMinting.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyMinting *KittyMintingCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyMinting.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyMinting *KittyMintingSession) TotalSupply() (*big.Int, error) {
	return _KittyMinting.Contract.TotalSupply(&_KittyMinting.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyMinting *KittyMintingCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyMinting.Contract.TotalSupply(&_KittyMinting.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Approve(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Approve(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyMinting *KittyMintingTransactor) ApproveSiring(opts *bind.TransactOpts, _addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "approveSiring", _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyMinting *KittyMintingSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.ApproveSiring(&_KittyMinting.TransactOpts, _addr, _sireId)
}

// ApproveSiring is a paid mutator transaction binding the contract method 0x4dfff04f.
//
// Solidity: function approveSiring(address _addr, uint256 _sireId) returns()
func (_KittyMinting *KittyMintingTransactorSession) ApproveSiring(_addr common.Address, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.ApproveSiring(&_KittyMinting.TransactOpts, _addr, _sireId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyMinting *KittyMintingTransactor) BidOnSiringAuction(opts *bind.TransactOpts, _sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "bidOnSiringAuction", _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyMinting *KittyMintingSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BidOnSiringAuction(&_KittyMinting.TransactOpts, _sireId, _matronId)
}

// BidOnSiringAuction is a paid mutator transaction binding the contract method 0xed60ade6.
//
// Solidity: function bidOnSiringAuction(uint256 _sireId, uint256 _matronId) payable returns()
func (_KittyMinting *KittyMintingTransactorSession) BidOnSiringAuction(_sireId *big.Int, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BidOnSiringAuction(&_KittyMinting.TransactOpts, _sireId, _matronId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyMinting *KittyMintingTransactor) BreedWithAuto(opts *bind.TransactOpts, _matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "breedWithAuto", _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyMinting *KittyMintingSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BreedWithAuto(&_KittyMinting.TransactOpts, _matronId, _sireId)
}

// BreedWithAuto is a paid mutator transaction binding the contract method 0xf7d8c883.
//
// Solidity: function breedWithAuto(uint256 _matronId, uint256 _sireId) payable returns()
func (_KittyMinting *KittyMintingTransactorSession) BreedWithAuto(_matronId *big.Int, _sireId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.BreedWithAuto(&_KittyMinting.TransactOpts, _matronId, _sireId)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyMinting *KittyMintingTransactor) CreateGen0Auction(opts *bind.TransactOpts, _genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createGen0Auction", _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyMinting *KittyMintingSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateGen0Auction(&_KittyMinting.TransactOpts, _genes)
}

// CreateGen0Auction is a paid mutator transaction binding the contract method 0xc3bea9af.
//
// Solidity: function createGen0Auction(uint256 _genes) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateGen0Auction(_genes *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateGen0Auction(&_KittyMinting.TransactOpts, _genes)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyMinting *KittyMintingTransactor) CreatePromoKitty(opts *bind.TransactOpts, _genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createPromoKitty", _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyMinting *KittyMintingSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreatePromoKitty(&_KittyMinting.TransactOpts, _genes, _owner)
}

// CreatePromoKitty is a paid mutator transaction binding the contract method 0x56129134.
//
// Solidity: function createPromoKitty(uint256 _genes, address _owner) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreatePromoKitty(_genes *big.Int, _owner common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreatePromoKitty(&_KittyMinting.TransactOpts, _genes, _owner)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingTransactor) CreateSaleAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createSaleAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSaleAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSaleAuction is a paid mutator transaction binding the contract method 0x3d7d3f5a.
//
// Solidity: function createSaleAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateSaleAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSaleAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingTransactor) CreateSiringAuction(opts *bind.TransactOpts, _kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "createSiringAuction", _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSiringAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// CreateSiringAuction is a paid mutator transaction binding the contract method 0x4ad8c938.
//
// Solidity: function createSiringAuction(uint256 _kittyId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration) returns()
func (_KittyMinting *KittyMintingTransactorSession) CreateSiringAuction(_kittyId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.CreateSiringAuction(&_KittyMinting.TransactOpts, _kittyId, _startingPrice, _endingPrice, _duration)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyMinting *KittyMintingTransactor) GiveBirth(opts *bind.TransactOpts, _matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "giveBirth", _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyMinting *KittyMintingSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.GiveBirth(&_KittyMinting.TransactOpts, _matronId)
}

// GiveBirth is a paid mutator transaction binding the contract method 0x88c2a0bf.
//
// Solidity: function giveBirth(uint256 _matronId) returns(uint256)
func (_KittyMinting *KittyMintingTransactorSession) GiveBirth(_matronId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.GiveBirth(&_KittyMinting.TransactOpts, _matronId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingSession) Pause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Pause(&_KittyMinting.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyMinting *KittyMintingTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Pause(&_KittyMinting.TransactOpts)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyMinting *KittyMintingTransactor) SetAutoBirthFee(opts *bind.TransactOpts, val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setAutoBirthFee", val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyMinting *KittyMintingSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetAutoBirthFee(&_KittyMinting.TransactOpts, val)
}

// SetAutoBirthFee is a paid mutator transaction binding the contract method 0x4b85fd55.
//
// Solidity: function setAutoBirthFee(uint256 val) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetAutoBirthFee(val *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetAutoBirthFee(&_KittyMinting.TransactOpts, val)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyMinting *KittyMintingTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyMinting *KittyMintingSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCEO(&_KittyMinting.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCEO(&_KittyMinting.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyMinting *KittyMintingTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyMinting *KittyMintingSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCFO(&_KittyMinting.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCFO(&_KittyMinting.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyMinting *KittyMintingTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyMinting *KittyMintingSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCOO(&_KittyMinting.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetCOO(&_KittyMinting.TransactOpts, _newCOO)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactor) SetGeneScienceAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setGeneScienceAddress", _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyMinting *KittyMintingSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetGeneScienceAddress(&_KittyMinting.TransactOpts, _address)
}

// SetGeneScienceAddress is a paid mutator transaction binding the contract method 0x24e7a38a.
//
// Solidity: function setGeneScienceAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetGeneScienceAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetGeneScienceAddress(&_KittyMinting.TransactOpts, _address)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyMinting *KittyMintingTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyMinting *KittyMintingSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetMetadataAddress(&_KittyMinting.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetMetadataAddress(&_KittyMinting.TransactOpts, _contractAddress)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactor) SetSaleAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSaleAuctionAddress", _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSaleAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSaleAuctionAddress is a paid mutator transaction binding the contract method 0x6fbde40d.
//
// Solidity: function setSaleAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSaleAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSaleAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyMinting *KittyMintingTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyMinting *KittyMintingSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSecondsPerBlock(&_KittyMinting.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSecondsPerBlock(&_KittyMinting.TransactOpts, secs)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactor) SetSiringAuctionAddress(opts *bind.TransactOpts, _address common.Address) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "setSiringAuctionAddress", _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSiringAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// SetSiringAuctionAddress is a paid mutator transaction binding the contract method 0x14001f4c.
//
// Solidity: function setSiringAuctionAddress(address _address) returns()
func (_KittyMinting *KittyMintingTransactorSession) SetSiringAuctionAddress(_address common.Address) (*types.Transaction, error) {
	return _KittyMinting.Contract.SetSiringAuctionAddress(&_KittyMinting.TransactOpts, _address)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Transfer(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.Transfer(&_KittyMinting.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.TransferFrom(&_KittyMinting.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyMinting *KittyMintingTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyMinting.Contract.TransferFrom(&_KittyMinting.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingSession) Unpause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Unpause(&_KittyMinting.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyMinting *KittyMintingTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyMinting.Contract.Unpause(&_KittyMinting.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingTransactor) WithdrawAuctionBalances(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyMinting.contract.Transact(opts, "withdrawAuctionBalances")
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyMinting.Contract.WithdrawAuctionBalances(&_KittyMinting.TransactOpts)
}

// WithdrawAuctionBalances is a paid mutator transaction binding the contract method 0x91876e57.
//
// Solidity: function withdrawAuctionBalances() returns()
func (_KittyMinting *KittyMintingTransactorSession) WithdrawAuctionBalances() (*types.Transaction, error) {
	return _KittyMinting.Contract.WithdrawAuctionBalances(&_KittyMinting.TransactOpts)
}

// KittyMintingApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyMinting contract.
type KittyMintingApprovalIterator struct {
	Event *KittyMintingApproval // Event containing the contract specifics and raw log

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
func (it *KittyMintingApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingApproval)
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
		it.Event = new(KittyMintingApproval)
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
func (it *KittyMintingApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingApproval represents a Approval event raised by the KittyMinting contract.
type KittyMintingApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyMinting *KittyMintingFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyMintingApprovalIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyMintingApprovalIterator{contract: _KittyMinting.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyMinting *KittyMintingFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyMintingApproval) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingApproval)
				if err := _KittyMinting.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KittyMinting *KittyMintingFilterer) ParseApproval(log types.Log) (*KittyMintingApproval, error) {
	event := new(KittyMintingApproval)
	if err := _KittyMinting.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyMintingBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyMinting contract.
type KittyMintingBirthIterator struct {
	Event *KittyMintingBirth // Event containing the contract specifics and raw log

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
func (it *KittyMintingBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingBirth)
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
		it.Event = new(KittyMintingBirth)
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
func (it *KittyMintingBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingBirth represents a Birth event raised by the KittyMinting contract.
type KittyMintingBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyMinting *KittyMintingFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyMintingBirthIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyMintingBirthIterator{contract: _KittyMinting.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyMinting *KittyMintingFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyMintingBirth) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingBirth)
				if err := _KittyMinting.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyMinting *KittyMintingFilterer) ParseBirth(log types.Log) (*KittyMintingBirth, error) {
	event := new(KittyMintingBirth)
	if err := _KittyMinting.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyMintingContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyMinting contract.
type KittyMintingContractUpgradeIterator struct {
	Event *KittyMintingContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyMintingContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingContractUpgrade)
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
		it.Event = new(KittyMintingContractUpgrade)
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
func (it *KittyMintingContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingContractUpgrade represents a ContractUpgrade event raised by the KittyMinting contract.
type KittyMintingContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyMinting *KittyMintingFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyMintingContractUpgradeIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyMintingContractUpgradeIterator{contract: _KittyMinting.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyMinting *KittyMintingFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyMintingContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingContractUpgrade)
				if err := _KittyMinting.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyMinting *KittyMintingFilterer) ParseContractUpgrade(log types.Log) (*KittyMintingContractUpgrade, error) {
	event := new(KittyMintingContractUpgrade)
	if err := _KittyMinting.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyMintingPregnantIterator is returned from FilterPregnant and is used to iterate over the raw logs and unpacked data for Pregnant events raised by the KittyMinting contract.
type KittyMintingPregnantIterator struct {
	Event *KittyMintingPregnant // Event containing the contract specifics and raw log

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
func (it *KittyMintingPregnantIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingPregnant)
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
		it.Event = new(KittyMintingPregnant)
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
func (it *KittyMintingPregnantIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingPregnantIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingPregnant represents a Pregnant event raised by the KittyMinting contract.
type KittyMintingPregnant struct {
	Owner            common.Address
	MatronId         *big.Int
	SireId           *big.Int
	CooldownEndBlock *big.Int
	Raw              types.Log // Blockchain specific contextual infos
}

// FilterPregnant is a free log retrieval operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyMinting *KittyMintingFilterer) FilterPregnant(opts *bind.FilterOpts) (*KittyMintingPregnantIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return &KittyMintingPregnantIterator{contract: _KittyMinting.contract, event: "Pregnant", logs: logs, sub: sub}, nil
}

// WatchPregnant is a free log subscription operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyMinting *KittyMintingFilterer) WatchPregnant(opts *bind.WatchOpts, sink chan<- *KittyMintingPregnant) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Pregnant")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingPregnant)
				if err := _KittyMinting.contract.UnpackLog(event, "Pregnant", log); err != nil {
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

// ParsePregnant is a log parse operation binding the contract event 0x241ea03ca20251805084d27d4440371c34a0b85ff108f6bb5611248f73818b80.
//
// Solidity: event Pregnant(address owner, uint256 matronId, uint256 sireId, uint256 cooldownEndBlock)
func (_KittyMinting *KittyMintingFilterer) ParsePregnant(log types.Log) (*KittyMintingPregnant, error) {
	event := new(KittyMintingPregnant)
	if err := _KittyMinting.contract.UnpackLog(event, "Pregnant", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyMintingTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyMinting contract.
type KittyMintingTransferIterator struct {
	Event *KittyMintingTransfer // Event containing the contract specifics and raw log

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
func (it *KittyMintingTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyMintingTransfer)
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
		it.Event = new(KittyMintingTransfer)
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
func (it *KittyMintingTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyMintingTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyMintingTransfer represents a Transfer event raised by the KittyMinting contract.
type KittyMintingTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyMinting *KittyMintingFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyMintingTransferIterator, error) {

	logs, sub, err := _KittyMinting.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyMintingTransferIterator{contract: _KittyMinting.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyMinting *KittyMintingFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyMintingTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyMinting.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyMintingTransfer)
				if err := _KittyMinting.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyMinting *KittyMintingFilterer) ParseTransfer(log types.Log) (*KittyMintingTransfer, error) {
	event := new(KittyMintingTransfer)
	if err := _KittyMinting.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyOwnershipABI is the input ABI used to generate the binding from.
const KittyOwnershipABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"_interfaceID\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cfoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_preferredTransport\",\"type\":\"string\"}],\"name\":\"tokenMetadata\",\"outputs\":[{\"name\":\"infoUrl\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ceoAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"siringAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCEO\",\"type\":\"address\"}],\"name\":\"setCEO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCOO\",\"type\":\"address\"}],\"name\":\"setCOO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"sireAllowedToAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToApproved\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newCFO\",\"type\":\"address\"}],\"name\":\"setCFO\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"secs\",\"type\":\"uint256\"}],\"name\":\"setSecondsPerBlock\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"count\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"secondsPerBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"tokensOfOwner\",\"outputs\":[{\"name\":\"ownerTokens\",\"type\":\"uint256[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"cooldowns\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"kittyIndexToOwner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"cooAddress\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"erc721Metadata\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_contractAddress\",\"type\":\"address\"}],\"name\":\"setMetadataAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"saleAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"kittyId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"matronId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"sireId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"genes\",\"type\":\"uint256\"}],\"name\":\"Birth\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"newContract\",\"type\":\"address\"}],\"name\":\"ContractUpgrade\",\"type\":\"event\"}]"

// KittyOwnershipFuncSigs maps the 4-byte function signature to its string representation.
var KittyOwnershipFuncSigs = map[string]string{
	"095ea7b3": "approve(address,uint256)",
	"70a08231": "balanceOf(address)",
	"0a0f8168": "ceoAddress()",
	"0519ce79": "cfoAddress()",
	"b047fb50": "cooAddress()",
	"9d6fac6f": "cooldowns(uint256)",
	"bc4006f5": "erc721Metadata()",
	"481af3d3": "kittyIndexToApproved(uint256)",
	"a45f4bfc": "kittyIndexToOwner(uint256)",
	"06fdde03": "name()",
	"6352211e": "ownerOf(uint256)",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"e6cbe351": "saleAuction()",
	"7a7d4937": "secondsPerBlock()",
	"27d7874c": "setCEO(address)",
	"4e0a3379": "setCFO(address)",
	"2ba73c15": "setCOO(address)",
	"e17b25af": "setMetadataAddress(address)",
	"5663896e": "setSecondsPerBlock(uint256)",
	"46116e6f": "sireAllowedToAddress(uint256)",
	"21717ebf": "siringAuction()",
	"01ffc9a7": "supportsInterface(bytes4)",
	"95d89b41": "symbol()",
	"0560ff44": "tokenMetadata(uint256,string)",
	"8462151c": "tokensOfOwner(address)",
	"18160ddd": "totalSupply()",
	"a9059cbb": "transfer(address,uint256)",
	"23b872dd": "transferFrom(address,address,uint256)",
	"3f4ba83a": "unpause()",
}

// KittyOwnershipBin is the compiled bytecode used for deploying new contracts.
var KittyOwnershipBin = "0x606060409081526002805460a060020a60ff02191690556101c090519081016040908152603c82526078602083015261012c9082015261025860608201526107086080820152610e1060a0820152611c2060c082015261384060e082015261708061010082015261e100610120820152620151806101408201526202a3006101608201526205460061018082015262093a806101a0820152620000a790600390600e620000bf565b50600f6005553415620000b957600080fd5b62000189565b600283019183908215620001505791602002820160005b838211156200011c57835183826101000a81548163ffffffff021916908363ffffffff1602179055509260200192600401602081600301049283019260010302620000d6565b80156200014e5782816101000a81549063ffffffff02191690556004016020816003010492830192600103026200011c565b505b506200015e92915062000162565b5090565b6200018691905b808211156200015e57805463ffffffff1916815560010162000169565b90565b61122580620001996000396000f30060606040526004361061017f5763ffffffff7c010000000000000000000000000000000000000000000000000000000060003504166301ffc9a781146101845780630519ce79146101d05780630560ff44146101ff57806306fdde0314610298578063095ea7b3146102ab5780630a0f8168146102cf57806318160ddd146102e257806321717ebf1461030757806323b872dd1461031a57806327d7874c146103425780632ba73c15146103615780633f4ba83a1461038057806346116e6f14610393578063481af3d3146103a95780634e0a3379146103bf5780635663896e146103de5780635c975abb146103f45780636352211e1461040757806370a082311461041d5780637a7d49371461043c5780638456cb591461044f5780638462151c1461046257806395d89b41146104d45780639d6fac6f146104e7578063a45f4bfc14610516578063a9059cbb1461052c578063b047fb501461054e578063bc4006f514610561578063e17b25af14610574578063e6cbe35114610593575b600080fd5b341561018f57600080fd5b6101bc7fffffffff00000000000000000000000000000000000000000000000000000000600435166105a6565b604051901515815260200160405180910390f35b34156101db57600080fd5b6101e361082d565b604051600160a060020a03909116815260200160405180910390f35b341561020a57600080fd5b61022160048035906024803590810191013561083c565b60405160208082528190810183818151815260200191508051906020019080838360005b8381101561025d578082015183820152602001610245565b50505050905090810190601f16801561028a5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34156102a357600080fd5b610221610925565b34156102b657600080fd5b6102cd600160a060020a036004351660243561095c565b005b34156102da57600080fd5b6101e36109e6565b34156102ed57600080fd5b6102f56109f5565b60405190815260200160405180910390f35b341561031257600080fd5b6101e36109ff565b341561032557600080fd5b6102cd600160a060020a0360043581169060243516604435610a0e565b341561034d57600080fd5b6102cd600160a060020a0360043516610a95565b341561036c57600080fd5b6102cd600160a060020a0360043516610ae7565b341561038b57600080fd5b6102cd610b39565b341561039e57600080fd5b6101e3600435610b8c565b34156103b457600080fd5b6101e3600435610ba7565b34156103ca57600080fd5b6102cd600160a060020a0360043516610bc2565b34156103e957600080fd5b6102cd600435610c14565b34156103ff57600080fd5b6101bc610c7c565b341561041257600080fd5b6101e3600435610c8c565b341561042857600080fd5b6102f5600160a060020a0360043516610cb0565b341561044757600080fd5b6102f5610ccb565b341561045a57600080fd5b6102cd610cd1565b341561046d57600080fd5b610481600160a060020a0360043516610d5d565b60405160208082528190810183818151815260200191508051906020019060200280838360005b838110156104c05780820151838201526020016104a8565b505050509050019250505060405180910390f35b34156104df57600080fd5b610221610e3e565b34156104f257600080fd5b6104fd600435610e75565b60405163ffffffff909116815260200160405180910390f35b341561052157600080fd5b6101e3600435610ea2565b341561053757600080fd5b6102cd600160a060020a0360043516602435610ebd565b341561055957600080fd5b6101e3610f64565b341561056c57600080fd5b6101e3610f73565b341561057f57600080fd5b6102cd600160a060020a0360043516610f82565b341561059e57600080fd5b6101e3610fbf565b60006040517f737570706f727473496e7465726661636528627974657334290000000000000081526019016040518091039020600160e060020a03191682600160e060020a031916148061082557506040517f746f6b656e4d657461646174612875696e743235362c737472696e67290000008152601d0160405180910390206040517f746f6b656e734f664f776e657228616464726573732900000000000000000000815260160160405180910390206040517f7472616e7366657246726f6d28616464726573732c616464726573732c75696e81527f7432353629000000000000000000000000000000000000000000000000000000602082015260250160405180910390206040517f7472616e7366657228616464726573732c75696e743235362900000000000000815260190160405180910390206040517f617070726f766528616464726573732c75696e74323536290000000000000000815260180160405180910390206040517f6f776e65724f662875696e743235362900000000000000000000000000000000815260100160405180910390206040517f62616c616e63654f662861646472657373290000000000000000000000000000815260120160405180910390206040517f746f74616c537570706c792829000000000000000000000000000000000000008152600d0160405180910390206040517f73796d626f6c2829000000000000000000000000000000000000000000000000815260080160405180910390206040517f6e616d652829000000000000000000000000000000000000000000000000000081526006016040518091039020181818181818181818600160e060020a03191682600160e060020a031916145b90505b919050565b600154600160a060020a031681565b6108446111be565b61084c6111d0565b600d54600090600160a060020a0316151561086657600080fd5b600d54600160a060020a031663cb4799f28787876040517c010000000000000000000000000000000000000000000000000000000063ffffffff861602815260048101848152604060248301908152604483018490529091606401848480828437820191505094505050505060a060405180830381600087803b15156108eb57600080fd5b5af115156108f857600080fd5b5050506040518060800180516020909101604052909250905061091b8282610fce565b9695505050505050565b60408051908101604052600d81527f43727970746f4b69747469657300000000000000000000000000000000000000602082015281565b60025460a060020a900460ff161561097357600080fd5b61097d3382611023565b151561098857600080fd5b6109928183611043565b7f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925338383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a15050565b600054600160a060020a031681565b6006546000190190565b600c54600160a060020a031681565b60025460a060020a900460ff1615610a2557600080fd5b600160a060020a0382161515610a3a57600080fd5b30600160a060020a031682600160a060020a031614151515610a5b57600080fd5b610a653382611071565b1515610a7057600080fd5b610a7a8382611023565b1515610a8557600080fd5b610a90838383611091565b505050565b60005433600160a060020a03908116911614610ab057600080fd5b600160a060020a0381161515610ac557600080fd5b60008054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610b0257600080fd5b600160a060020a0381161515610b1757600080fd5b60028054600160a060020a031916600160a060020a0392909216919091179055565b60005433600160a060020a03908116911614610b5457600080fd5b60025460a060020a900460ff161515610b6c57600080fd5b6002805474ff000000000000000000000000000000000000000019169055565b600a60205260009081526040902054600160a060020a031681565b600960205260009081526040902054600160a060020a031681565b60005433600160a060020a03908116911614610bdd57600080fd5b600160a060020a0381161515610bf257600080fd5b60018054600160a060020a031916600160a060020a0392909216919091179055565b60025433600160a060020a0390811691161480610c3f575060005433600160a060020a039081169116145b80610c58575060015433600160a060020a039081169116145b1515610c6357600080fd5b60035463ffffffff168110610c7757600080fd5b600555565b60025460a060020a900460ff1681565b600081815260076020526040902054600160a060020a031680151561082857600080fd5b600160a060020a031660009081526008602052604090205490565b60055481565b60025433600160a060020a0390811691161480610cfc575060005433600160a060020a039081169116145b80610d15575060015433600160a060020a039081169116145b1515610d2057600080fd5b60025460a060020a900460ff1615610d3757600080fd5b6002805474ff0000000000000000000000000000000000000000191660a060020a179055565b610d656111be565b6000610d6f6111be565b6000806000610d7d87610cb0565b9450841515610dad576000604051805910610d955750595b90808252806020026020018201604052509550610e34565b84604051805910610dbb5750595b90808252806020026020018201604052509350610dd66109f5565b925060009150600190505b828111610e3057600081815260076020526040902054600160a060020a0388811691161415610e285780848381518110610e1757fe5b602090810290910101526001909101905b600101610de1565b8395505b5050505050919050565b60408051908101604052600281527f434b000000000000000000000000000000000000000000000000000000000000602082015281565b600381600e8110610e8257fe5b60089182820401919006600402915054906101000a900463ffffffff1681565b600760205260009081526040902054600160a060020a031681565b60025460a060020a900460ff1615610ed457600080fd5b600160a060020a0382161515610ee957600080fd5b30600160a060020a031682600160a060020a031614151515610f0a57600080fd5b600b54600160a060020a0383811691161415610f2557600080fd5b600c54600160a060020a0383811691161415610f4057600080fd5b610f4a3382611023565b1515610f5557600080fd5b610f60338383611091565b5050565b600254600160a060020a031681565b600d54600160a060020a031681565b60005433600160a060020a03908116911614610f9d57600080fd5b600d8054600160a060020a031916600160a060020a0392909216919091179055565b600b54600160a060020a031681565b610fd66111be565b610fde6111be565b60008084604051805910610fef5750595b818152601f19601f8301168101602001604052905092505060208201905084611019828287611179565b5090949350505050565b600090815260076020526040902054600160a060020a0391821691161490565b6000918252600960205260409091208054600160a060020a031916600160a060020a03909216919091179055565b600090815260096020526040902054600160a060020a0391821691161490565b600160a060020a03808316600081815260086020908152604080832080546001019055858352600790915290208054600160a060020a031916909117905583161561112457600160a060020a03831660009081526008602090815260408083208054600019019055838352600a82528083208054600160a060020a03199081169091556009909252909120805490911690555b7fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef838383604051600160a060020a039384168152919092166020820152604080820192909252606001905180910390a1505050565b60005b6020821061119f578251845260208401935060208301925060208203915061117c565b6001826020036101000a03905080198351168185511617909352505050565b60206040519081016040526000815290565b60806040519081016040526004815b600081526000199190910190602001816111df57905050905600a165627a7a7230582026b14365de8966a12508427d03efd535efc282255d45ebf502e50ef0b9eba80a0029"

// DeployKittyOwnership deploys a new Ethereum contract, binding an instance of KittyOwnership to it.
func DeployKittyOwnership(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *KittyOwnership, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyOwnershipABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(KittyOwnershipBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &KittyOwnership{KittyOwnershipCaller: KittyOwnershipCaller{contract: contract}, KittyOwnershipTransactor: KittyOwnershipTransactor{contract: contract}, KittyOwnershipFilterer: KittyOwnershipFilterer{contract: contract}}, nil
}

// KittyOwnership is an auto generated Go binding around an Ethereum contract.
type KittyOwnership struct {
	KittyOwnershipCaller     // Read-only binding to the contract
	KittyOwnershipTransactor // Write-only binding to the contract
	KittyOwnershipFilterer   // Log filterer for contract events
}

// KittyOwnershipCaller is an auto generated read-only Go binding around an Ethereum contract.
type KittyOwnershipCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipTransactor is an auto generated write-only Go binding around an Ethereum contract.
type KittyOwnershipTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type KittyOwnershipFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// KittyOwnershipSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type KittyOwnershipSession struct {
	Contract     *KittyOwnership   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// KittyOwnershipCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type KittyOwnershipCallerSession struct {
	Contract *KittyOwnershipCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// KittyOwnershipTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type KittyOwnershipTransactorSession struct {
	Contract     *KittyOwnershipTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// KittyOwnershipRaw is an auto generated low-level Go binding around an Ethereum contract.
type KittyOwnershipRaw struct {
	Contract *KittyOwnership // Generic contract binding to access the raw methods on
}

// KittyOwnershipCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type KittyOwnershipCallerRaw struct {
	Contract *KittyOwnershipCaller // Generic read-only contract binding to access the raw methods on
}

// KittyOwnershipTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type KittyOwnershipTransactorRaw struct {
	Contract *KittyOwnershipTransactor // Generic write-only contract binding to access the raw methods on
}

// NewKittyOwnership creates a new instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnership(address common.Address, backend bind.ContractBackend) (*KittyOwnership, error) {
	contract, err := bindKittyOwnership(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &KittyOwnership{KittyOwnershipCaller: KittyOwnershipCaller{contract: contract}, KittyOwnershipTransactor: KittyOwnershipTransactor{contract: contract}, KittyOwnershipFilterer: KittyOwnershipFilterer{contract: contract}}, nil
}

// NewKittyOwnershipCaller creates a new read-only instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipCaller(address common.Address, caller bind.ContractCaller) (*KittyOwnershipCaller, error) {
	contract, err := bindKittyOwnership(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipCaller{contract: contract}, nil
}

// NewKittyOwnershipTransactor creates a new write-only instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipTransactor(address common.Address, transactor bind.ContractTransactor) (*KittyOwnershipTransactor, error) {
	contract, err := bindKittyOwnership(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipTransactor{contract: contract}, nil
}

// NewKittyOwnershipFilterer creates a new log filterer instance of KittyOwnership, bound to a specific deployed contract.
func NewKittyOwnershipFilterer(address common.Address, filterer bind.ContractFilterer) (*KittyOwnershipFilterer, error) {
	contract, err := bindKittyOwnership(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipFilterer{contract: contract}, nil
}

// bindKittyOwnership binds a generic wrapper to an already deployed contract.
func bindKittyOwnership(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(KittyOwnershipABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyOwnership *KittyOwnershipRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyOwnership.Contract.KittyOwnershipCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyOwnership *KittyOwnershipRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.Contract.KittyOwnershipTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyOwnership *KittyOwnershipRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyOwnership.Contract.KittyOwnershipTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_KittyOwnership *KittyOwnershipCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _KittyOwnership.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_KittyOwnership *KittyOwnershipTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_KittyOwnership *KittyOwnershipTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _KittyOwnership.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyOwnership *KittyOwnershipCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "balanceOf", _owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyOwnership *KittyOwnershipSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyOwnership.Contract.BalanceOf(&_KittyOwnership.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address _owner) view returns(uint256 count)
func (_KittyOwnership *KittyOwnershipCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _KittyOwnership.Contract.BalanceOf(&_KittyOwnership.CallOpts, _owner)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CeoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "ceoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) CeoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CeoAddress(&_KittyOwnership.CallOpts)
}

// CeoAddress is a free data retrieval call binding the contract method 0x0a0f8168.
//
// Solidity: function ceoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CeoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CeoAddress(&_KittyOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CfoAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "cfoAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) CfoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CfoAddress(&_KittyOwnership.CallOpts)
}

// CfoAddress is a free data retrieval call binding the contract method 0x0519ce79.
//
// Solidity: function cfoAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CfoAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CfoAddress(&_KittyOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) CooAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "cooAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) CooAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CooAddress(&_KittyOwnership.CallOpts)
}

// CooAddress is a free data retrieval call binding the contract method 0xb047fb50.
//
// Solidity: function cooAddress() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) CooAddress() (common.Address, error) {
	return _KittyOwnership.Contract.CooAddress(&_KittyOwnership.CallOpts)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyOwnership *KittyOwnershipCaller) Cooldowns(opts *bind.CallOpts, arg0 *big.Int) (uint32, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "cooldowns", arg0)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyOwnership *KittyOwnershipSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyOwnership.Contract.Cooldowns(&_KittyOwnership.CallOpts, arg0)
}

// Cooldowns is a free data retrieval call binding the contract method 0x9d6fac6f.
//
// Solidity: function cooldowns(uint256 ) view returns(uint32)
func (_KittyOwnership *KittyOwnershipCallerSession) Cooldowns(arg0 *big.Int) (uint32, error) {
	return _KittyOwnership.Contract.Cooldowns(&_KittyOwnership.CallOpts, arg0)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) Erc721Metadata(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "erc721Metadata")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) Erc721Metadata() (common.Address, error) {
	return _KittyOwnership.Contract.Erc721Metadata(&_KittyOwnership.CallOpts)
}

// Erc721Metadata is a free data retrieval call binding the contract method 0xbc4006f5.
//
// Solidity: function erc721Metadata() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) Erc721Metadata() (common.Address, error) {
	return _KittyOwnership.Contract.Erc721Metadata(&_KittyOwnership.CallOpts)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) KittyIndexToApproved(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "kittyIndexToApproved", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToApproved(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToApproved is a free data retrieval call binding the contract method 0x481af3d3.
//
// Solidity: function kittyIndexToApproved(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) KittyIndexToApproved(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToApproved(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) KittyIndexToOwner(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "kittyIndexToOwner", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToOwner(&_KittyOwnership.CallOpts, arg0)
}

// KittyIndexToOwner is a free data retrieval call binding the contract method 0xa45f4bfc.
//
// Solidity: function kittyIndexToOwner(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) KittyIndexToOwner(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.KittyIndexToOwner(&_KittyOwnership.CallOpts, arg0)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyOwnership *KittyOwnershipCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyOwnership *KittyOwnershipSession) Name() (string, error) {
	return _KittyOwnership.Contract.Name(&_KittyOwnership.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_KittyOwnership *KittyOwnershipCallerSession) Name() (string, error) {
	return _KittyOwnership.Contract.Name(&_KittyOwnership.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyOwnership *KittyOwnershipCaller) OwnerOf(opts *bind.CallOpts, _tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "ownerOf", _tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyOwnership *KittyOwnershipSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.OwnerOf(&_KittyOwnership.CallOpts, _tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 _tokenId) view returns(address owner)
func (_KittyOwnership *KittyOwnershipCallerSession) OwnerOf(_tokenId *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.OwnerOf(&_KittyOwnership.CallOpts, _tokenId)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyOwnership *KittyOwnershipCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyOwnership *KittyOwnershipSession) Paused() (bool, error) {
	return _KittyOwnership.Contract.Paused(&_KittyOwnership.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_KittyOwnership *KittyOwnershipCallerSession) Paused() (bool, error) {
	return _KittyOwnership.Contract.Paused(&_KittyOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SaleAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "saleAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) SaleAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SaleAuction(&_KittyOwnership.CallOpts)
}

// SaleAuction is a free data retrieval call binding the contract method 0xe6cbe351.
//
// Solidity: function saleAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SaleAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SaleAuction(&_KittyOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyOwnership *KittyOwnershipCaller) SecondsPerBlock(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "secondsPerBlock")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyOwnership *KittyOwnershipSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyOwnership.Contract.SecondsPerBlock(&_KittyOwnership.CallOpts)
}

// SecondsPerBlock is a free data retrieval call binding the contract method 0x7a7d4937.
//
// Solidity: function secondsPerBlock() view returns(uint256)
func (_KittyOwnership *KittyOwnershipCallerSession) SecondsPerBlock() (*big.Int, error) {
	return _KittyOwnership.Contract.SecondsPerBlock(&_KittyOwnership.CallOpts)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SireAllowedToAddress(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "sireAllowedToAddress", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.SireAllowedToAddress(&_KittyOwnership.CallOpts, arg0)
}

// SireAllowedToAddress is a free data retrieval call binding the contract method 0x46116e6f.
//
// Solidity: function sireAllowedToAddress(uint256 ) view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SireAllowedToAddress(arg0 *big.Int) (common.Address, error) {
	return _KittyOwnership.Contract.SireAllowedToAddress(&_KittyOwnership.CallOpts, arg0)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipCaller) SiringAuction(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "siringAuction")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipSession) SiringAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SiringAuction(&_KittyOwnership.CallOpts)
}

// SiringAuction is a free data retrieval call binding the contract method 0x21717ebf.
//
// Solidity: function siringAuction() view returns(address)
func (_KittyOwnership *KittyOwnershipCallerSession) SiringAuction() (common.Address, error) {
	return _KittyOwnership.Contract.SiringAuction(&_KittyOwnership.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyOwnership *KittyOwnershipCaller) SupportsInterface(opts *bind.CallOpts, _interfaceID [4]byte) (bool, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "supportsInterface", _interfaceID)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyOwnership *KittyOwnershipSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyOwnership.Contract.SupportsInterface(&_KittyOwnership.CallOpts, _interfaceID)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 _interfaceID) view returns(bool)
func (_KittyOwnership *KittyOwnershipCallerSession) SupportsInterface(_interfaceID [4]byte) (bool, error) {
	return _KittyOwnership.Contract.SupportsInterface(&_KittyOwnership.CallOpts, _interfaceID)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyOwnership *KittyOwnershipCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyOwnership *KittyOwnershipSession) Symbol() (string, error) {
	return _KittyOwnership.Contract.Symbol(&_KittyOwnership.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_KittyOwnership *KittyOwnershipCallerSession) Symbol() (string, error) {
	return _KittyOwnership.Contract.Symbol(&_KittyOwnership.CallOpts)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyOwnership *KittyOwnershipCaller) TokenMetadata(opts *bind.CallOpts, _tokenId *big.Int, _preferredTransport string) (string, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "tokenMetadata", _tokenId, _preferredTransport)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyOwnership *KittyOwnershipSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyOwnership.Contract.TokenMetadata(&_KittyOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokenMetadata is a free data retrieval call binding the contract method 0x0560ff44.
//
// Solidity: function tokenMetadata(uint256 _tokenId, string _preferredTransport) view returns(string infoUrl)
func (_KittyOwnership *KittyOwnershipCallerSession) TokenMetadata(_tokenId *big.Int, _preferredTransport string) (string, error) {
	return _KittyOwnership.Contract.TokenMetadata(&_KittyOwnership.CallOpts, _tokenId, _preferredTransport)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyOwnership *KittyOwnershipCaller) TokensOfOwner(opts *bind.CallOpts, _owner common.Address) ([]*big.Int, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "tokensOfOwner", _owner)

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyOwnership *KittyOwnershipSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyOwnership.Contract.TokensOfOwner(&_KittyOwnership.CallOpts, _owner)
}

// TokensOfOwner is a free data retrieval call binding the contract method 0x8462151c.
//
// Solidity: function tokensOfOwner(address _owner) view returns(uint256[] ownerTokens)
func (_KittyOwnership *KittyOwnershipCallerSession) TokensOfOwner(_owner common.Address) ([]*big.Int, error) {
	return _KittyOwnership.Contract.TokensOfOwner(&_KittyOwnership.CallOpts, _owner)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyOwnership *KittyOwnershipCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _KittyOwnership.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyOwnership *KittyOwnershipSession) TotalSupply() (*big.Int, error) {
	return _KittyOwnership.Contract.TotalSupply(&_KittyOwnership.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_KittyOwnership *KittyOwnershipCallerSession) TotalSupply() (*big.Int, error) {
	return _KittyOwnership.Contract.TotalSupply(&_KittyOwnership.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactor) Approve(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "approve", _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Approve(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Approve(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Approve(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipSession) Pause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Pause(&_KittyOwnership.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Pause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Pause(&_KittyOwnership.TransactOpts)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCEO(opts *bind.TransactOpts, _newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCEO", _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCEO(&_KittyOwnership.TransactOpts, _newCEO)
}

// SetCEO is a paid mutator transaction binding the contract method 0x27d7874c.
//
// Solidity: function setCEO(address _newCEO) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCEO(_newCEO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCEO(&_KittyOwnership.TransactOpts, _newCEO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCFO(opts *bind.TransactOpts, _newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCFO", _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCFO(&_KittyOwnership.TransactOpts, _newCFO)
}

// SetCFO is a paid mutator transaction binding the contract method 0x4e0a3379.
//
// Solidity: function setCFO(address _newCFO) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCFO(_newCFO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCFO(&_KittyOwnership.TransactOpts, _newCFO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetCOO(opts *bind.TransactOpts, _newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setCOO", _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyOwnership *KittyOwnershipSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCOO(&_KittyOwnership.TransactOpts, _newCOO)
}

// SetCOO is a paid mutator transaction binding the contract method 0x2ba73c15.
//
// Solidity: function setCOO(address _newCOO) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetCOO(_newCOO common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetCOO(&_KittyOwnership.TransactOpts, _newCOO)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetMetadataAddress(opts *bind.TransactOpts, _contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setMetadataAddress", _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyOwnership *KittyOwnershipSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetMetadataAddress(&_KittyOwnership.TransactOpts, _contractAddress)
}

// SetMetadataAddress is a paid mutator transaction binding the contract method 0xe17b25af.
//
// Solidity: function setMetadataAddress(address _contractAddress) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetMetadataAddress(_contractAddress common.Address) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetMetadataAddress(&_KittyOwnership.TransactOpts, _contractAddress)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyOwnership *KittyOwnershipTransactor) SetSecondsPerBlock(opts *bind.TransactOpts, secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "setSecondsPerBlock", secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyOwnership *KittyOwnershipSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetSecondsPerBlock(&_KittyOwnership.TransactOpts, secs)
}

// SetSecondsPerBlock is a paid mutator transaction binding the contract method 0x5663896e.
//
// Solidity: function setSecondsPerBlock(uint256 secs) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) SetSecondsPerBlock(secs *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.SetSecondsPerBlock(&_KittyOwnership.TransactOpts, secs)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "transfer", _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Transfer(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Transfer(_to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.Transfer(&_KittyOwnership.TransactOpts, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "transferFrom", _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.TransferFrom(&_KittyOwnership.TransactOpts, _from, _to, _tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address _from, address _to, uint256 _tokenId) returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) TransferFrom(_from common.Address, _to common.Address, _tokenId *big.Int) (*types.Transaction, error) {
	return _KittyOwnership.Contract.TransferFrom(&_KittyOwnership.TransactOpts, _from, _to, _tokenId)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _KittyOwnership.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipSession) Unpause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Unpause(&_KittyOwnership.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_KittyOwnership *KittyOwnershipTransactorSession) Unpause() (*types.Transaction, error) {
	return _KittyOwnership.Contract.Unpause(&_KittyOwnership.TransactOpts)
}

// KittyOwnershipApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the KittyOwnership contract.
type KittyOwnershipApprovalIterator struct {
	Event *KittyOwnershipApproval // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipApproval)
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
		it.Event = new(KittyOwnershipApproval)
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
func (it *KittyOwnershipApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipApproval represents a Approval event raised by the KittyOwnership contract.
type KittyOwnershipApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyOwnership *KittyOwnershipFilterer) FilterApproval(opts *bind.FilterOpts) (*KittyOwnershipApprovalIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipApprovalIterator{contract: _KittyOwnership.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address owner, address approved, uint256 tokenId)
func (_KittyOwnership *KittyOwnershipFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *KittyOwnershipApproval) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Approval")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipApproval)
				if err := _KittyOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_KittyOwnership *KittyOwnershipFilterer) ParseApproval(log types.Log) (*KittyOwnershipApproval, error) {
	event := new(KittyOwnershipApproval)
	if err := _KittyOwnership.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyOwnershipBirthIterator is returned from FilterBirth and is used to iterate over the raw logs and unpacked data for Birth events raised by the KittyOwnership contract.
type KittyOwnershipBirthIterator struct {
	Event *KittyOwnershipBirth // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipBirthIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipBirth)
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
		it.Event = new(KittyOwnershipBirth)
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
func (it *KittyOwnershipBirthIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipBirthIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipBirth represents a Birth event raised by the KittyOwnership contract.
type KittyOwnershipBirth struct {
	Owner    common.Address
	KittyId  *big.Int
	MatronId *big.Int
	SireId   *big.Int
	Genes    *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBirth is a free log retrieval operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyOwnership *KittyOwnershipFilterer) FilterBirth(opts *bind.FilterOpts) (*KittyOwnershipBirthIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipBirthIterator{contract: _KittyOwnership.contract, event: "Birth", logs: logs, sub: sub}, nil
}

// WatchBirth is a free log subscription operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyOwnership *KittyOwnershipFilterer) WatchBirth(opts *bind.WatchOpts, sink chan<- *KittyOwnershipBirth) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Birth")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipBirth)
				if err := _KittyOwnership.contract.UnpackLog(event, "Birth", log); err != nil {
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

// ParseBirth is a log parse operation binding the contract event 0x0a5311bd2a6608f08a180df2ee7c5946819a649b204b554bb8e39825b2c50ad5.
//
// Solidity: event Birth(address owner, uint256 kittyId, uint256 matronId, uint256 sireId, uint256 genes)
func (_KittyOwnership *KittyOwnershipFilterer) ParseBirth(log types.Log) (*KittyOwnershipBirth, error) {
	event := new(KittyOwnershipBirth)
	if err := _KittyOwnership.contract.UnpackLog(event, "Birth", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyOwnershipContractUpgradeIterator is returned from FilterContractUpgrade and is used to iterate over the raw logs and unpacked data for ContractUpgrade events raised by the KittyOwnership contract.
type KittyOwnershipContractUpgradeIterator struct {
	Event *KittyOwnershipContractUpgrade // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipContractUpgradeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipContractUpgrade)
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
		it.Event = new(KittyOwnershipContractUpgrade)
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
func (it *KittyOwnershipContractUpgradeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipContractUpgradeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipContractUpgrade represents a ContractUpgrade event raised by the KittyOwnership contract.
type KittyOwnershipContractUpgrade struct {
	NewContract common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterContractUpgrade is a free log retrieval operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyOwnership *KittyOwnershipFilterer) FilterContractUpgrade(opts *bind.FilterOpts) (*KittyOwnershipContractUpgradeIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipContractUpgradeIterator{contract: _KittyOwnership.contract, event: "ContractUpgrade", logs: logs, sub: sub}, nil
}

// WatchContractUpgrade is a free log subscription operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyOwnership *KittyOwnershipFilterer) WatchContractUpgrade(opts *bind.WatchOpts, sink chan<- *KittyOwnershipContractUpgrade) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "ContractUpgrade")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipContractUpgrade)
				if err := _KittyOwnership.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
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

// ParseContractUpgrade is a log parse operation binding the contract event 0x450db8da6efbe9c22f2347f7c2021231df1fc58d3ae9a2fa75d39fa446199305.
//
// Solidity: event ContractUpgrade(address newContract)
func (_KittyOwnership *KittyOwnershipFilterer) ParseContractUpgrade(log types.Log) (*KittyOwnershipContractUpgrade, error) {
	event := new(KittyOwnershipContractUpgrade)
	if err := _KittyOwnership.contract.UnpackLog(event, "ContractUpgrade", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// KittyOwnershipTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the KittyOwnership contract.
type KittyOwnershipTransferIterator struct {
	Event *KittyOwnershipTransfer // Event containing the contract specifics and raw log

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
func (it *KittyOwnershipTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(KittyOwnershipTransfer)
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
		it.Event = new(KittyOwnershipTransfer)
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
func (it *KittyOwnershipTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *KittyOwnershipTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// KittyOwnershipTransfer represents a Transfer event raised by the KittyOwnership contract.
type KittyOwnershipTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyOwnership *KittyOwnershipFilterer) FilterTransfer(opts *bind.FilterOpts) (*KittyOwnershipTransferIterator, error) {

	logs, sub, err := _KittyOwnership.contract.FilterLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return &KittyOwnershipTransferIterator{contract: _KittyOwnership.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address from, address to, uint256 tokenId)
func (_KittyOwnership *KittyOwnershipFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *KittyOwnershipTransfer) (event.Subscription, error) {

	logs, sub, err := _KittyOwnership.contract.WatchLogs(opts, "Transfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(KittyOwnershipTransfer)
				if err := _KittyOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_KittyOwnership *KittyOwnershipFilterer) ParseTransfer(log types.Log) (*KittyOwnershipTransfer, error) {
	event := new(KittyOwnershipTransfer)
	if err := _KittyOwnership.contract.UnpackLog(event, "Transfer", log); err != nil {
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
var OwnableBin = "0x6060604052341561000f57600080fd5b60008054600160a060020a033316600160a060020a031990911617905561017f8061003b6000396000f30060606040526004361061004b5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416638da5cb5b8114610050578063f2fde38b1461008c575b600080fd5b341561005b57600080fd5b6100636100ba565b60405173ffffffffffffffffffffffffffffffffffffffff909116815260200160405180910390f35b341561009757600080fd5b6100b873ffffffffffffffffffffffffffffffffffffffff600435166100d6565b005b60005473ffffffffffffffffffffffffffffffffffffffff1681565b6000543373ffffffffffffffffffffffffffffffffffffffff9081169116146100fe57600080fd5b73ffffffffffffffffffffffffffffffffffffffff811615610150576000805473ffffffffffffffffffffffffffffffffffffffff191673ffffffffffffffffffffffffffffffffffffffff83161790555b505600a165627a7a7230582052a652a6ee6dca506179948d92db1ba0441bd3a562fd92c002c39e6d99920c5c0029"

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
var PausableBin = "0x606060405260008054600160a060020a033316600160a860020a0319909116179055610300806100306000396000f30060606040526004361061006c5763ffffffff7c01000000000000000000000000000000000000000000000000000000006000350416633f4ba83a81146100715780635c975abb146100985780638456cb59146100ab5780638da5cb5b146100be578063f2fde38b146100ed575b600080fd5b341561007c57600080fd5b61008461010e565b604051901515815260200160405180910390f35b34156100a357600080fd5b6100846101a3565b34156100b657600080fd5b6100846101c4565b34156100c957600080fd5b6100d161026f565b604051600160a060020a03909116815260200160405180910390f35b34156100f857600080fd5b61010c600160a060020a036004351661027e565b005b6000805433600160a060020a0390811691161461012a57600080fd5b60005474010000000000000000000000000000000000000000900460ff16151561015357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60005474010000000000000000000000000000000000000000900460ff1681565b6000805433600160a060020a039081169116146101e057600080fd5b60005474010000000000000000000000000000000000000000900460ff161561020857600080fd5b6000805474ff00000000000000000000000000000000000000001916740100000000000000000000000000000000000000001790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b600054600160a060020a031681565b60005433600160a060020a0390811691161461029957600080fd5b600160a060020a038116156102d1576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b505600a165627a7a72305820a1649640b194c33c1b6ef2ee29f5bc68a321d7a64d3266cf4b3c684434c1be700029"

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
var SaleClockAuctionBin = "0x60606040526000805460a060020a60ff02191690556004805460ff19166001179055341561002c57600080fd5b604051604080610ef6833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316178155909250839150829061271082111561007857600080fd5b50600281905581600160a060020a0381166301ffc9a77f9a20483d000000000000000000000000000000000000000000000000000000006040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401602060405180830381600087803b151561011f57600080fd5b5af1151561012c57600080fd5b50505060405180519050151561014157600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610d81806101756000396000f3006060604052600436106100e25763ffffffff60e060020a60003504166327ebe40a81146100e75780633f4ba83a14610114578063454a2ab31461013b578063484eccb4146101465780635c975abb1461016e5780635fd8c7101461018157806378bd79351461019457806383b5ff8b146101e55780638456cb59146101f857806385b861881461020b578063878eb3681461021e5780638a98a9cc146102345780638da5cb5b1461024757806396b5a75514610276578063c55d0f561461028c578063dd1b7a0f146102a2578063eac9d94c146102b5578063f2fde38b146102c8575b600080fd5b34156100f257600080fd5b610112600435602435604435606435600160a060020a03608435166102e7565b005b341561011f57600080fd5b6101276103bf565b604051901515815260200160405180910390f35b610112600435610443565b341561015157600080fd5b61015c6004356104ad565b60405190815260200160405180910390f35b341561017957600080fd5b6101276104c1565b341561018c57600080fd5b6101126104d1565b341561019f57600080fd5b6101aa600435610547565b604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390f35b34156101f057600080fd5b61015c6105d4565b341561020357600080fd5b6101276105da565b341561021657600080fd5b610127610663565b341561022957600080fd5b61011260043561066c565b341561023f57600080fd5b61015c6106dd565b341561025257600080fd5b61025a6106e3565b604051600160a060020a03909116815260200160405180910390f35b341561028157600080fd5b6101126004356106f2565b341561029757600080fd5b61015c60043561073b565b34156102ad57600080fd5b61025a61076d565b34156102c057600080fd5b61015c61077c565b34156102d357600080fd5b610112600160a060020a03600435166107b0565b6102ef610d27565b6001608060020a038516851461030457600080fd5b6001608060020a038416841461031957600080fd5b67ffffffffffffffff8316831461032f57600080fd5b60015433600160a060020a0390811691161461034a57600080fd5b6103548287610806565b60a06040519081016040528083600160a060020a03168152602001866001608060020a03168152602001856001608060020a031681526020018467ffffffffffffffff1681526020014267ffffffffffffffff1681525090506103b7868261087d565b505050505050565b6000805433600160a060020a039081169116146103db57600080fd5b60005460a060020a900460ff1615156103f357600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b600081815260036020526040812054600160a060020a0316906104668334610a18565b90506104723384610b49565b600154600160a060020a03838116911614156104a857600580548291600691066005811061049c57fe5b01556005805460010190555b505050565b600681600581106104ba57fe5b0154905081565b60005460a060020a900460ff1681565b60015460008054600160a060020a039283169233811691161480610506575081600160a060020a031633600160a060020a0316145b151561051157600080fd5b81600160a060020a03166108fc30600160a060020a0316319081150290604051600060405180830381858888f150505050505050565b6000818152600360205260408120819081908190819061056681610b9f565b151561057157600080fd5b80546001820154600290920154600160a060020a03909116986001608060020a038084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b6000805433600160a060020a039081169116146105f657600080fd5b60005460a060020a900460ff161561060d57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b60045460ff1681565b6000805460a060020a900460ff16151561068557600080fd5b60005433600160a060020a039081169116146106a057600080fd5b5060008181526003602052604090206106b881610b9f565b15156106c357600080fd5b80546106d9908390600160a060020a0316610bc0565b5050565b60055481565b600054600160a060020a031681565b60008181526003602052604081209061070a82610b9f565b151561071557600080fd5b508054600160a060020a03908116903316811461073157600080fd5b6104a88382610bc0565b600081815260036020526040812061075281610b9f565b151561075d57600080fd5b61076681610c0a565b9392505050565b600154600160a060020a031681565b600080805b60058110156107a6576006816005811061079757fe5b01549190910190600101610781565b5060059004919050565b60005433600160a060020a039081169116146107cb57600080fd5b600160a060020a03811615610803576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154600160a060020a03166323b872dd83308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b151561086957600080fd5b5af1151561087657600080fd5b5050505050565b603c816060015167ffffffffffffffff16101561089957600080fd5b600082815260036020526040902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560408201516001820180546001608060020a03928316700100000000000000000000000000000000029216919091179055606082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560808201516002909101805467ffffffffffffffff9290921668010000000000000000026fffffffffffffffff000000000000000019909216919091179055507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78260208301516001608060020a031683604001516001608060020a0316846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b60008281526003602052604081208180808080610a3486610b9f565b1515610a3f57600080fd5b610a4886610c0a565b945084881015610a5757600080fd5b8554600160a060020a03169350610a6d89610c91565b6000851115610ab757610a7f85610cde565b92508285039150600160a060020a03841682156108fc0283604051600060405180830381858888f193505050501515610ab757600080fd5b50838703600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610aec57600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd28986336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a15092979650505050505050565b600154600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b151561086957600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610bc982610c91565b610bd38183610b49565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df8260405190815260200160405180910390a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610c505750600282015468010000000000000000900467ffffffffffffffff1642035b60018301546002840154610766916001608060020a0380821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610cea565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610cfe57869350610d1c565b878703925085858402811515610d1057fe5b05915081880190508093505b505050949350505050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a723058203fd71c63cb1e9083a70435e9914e093cafab4a3e87520e971a195b1fd2f253010029"

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

// SiringClockAuctionABI is the input ABI used to generate the binding from.
const SiringClockAuctionABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"name\":\"_startingPrice\",\"type\":\"uint256\"},{\"name\":\"_endingPrice\",\"type\":\"uint256\"},{\"name\":\"_duration\",\"type\":\"uint256\"},{\"name\":\"_seller\",\"type\":\"address\"}],\"name\":\"createAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"bid\",\"outputs\":[],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"withdrawBalance\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isSiringClockAuction\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getAuction\",\"outputs\":[{\"name\":\"seller\",\"type\":\"address\"},{\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"name\":\"duration\",\"type\":\"uint256\"},{\"name\":\"startedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ownerCut\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuctionWhenPaused\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"cancelAuction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_tokenId\",\"type\":\"uint256\"}],\"name\":\"getCurrentPrice\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"nonFungibleContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_nftAddr\",\"type\":\"address\"},{\"name\":\"_cut\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"startingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"endingPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"duration\",\"type\":\"uint256\"}],\"name\":\"AuctionCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"totalPrice\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"winner\",\"type\":\"address\"}],\"name\":\"AuctionSuccessful\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"AuctionCancelled\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"}]"

// SiringClockAuctionFuncSigs maps the 4-byte function signature to its string representation.
var SiringClockAuctionFuncSigs = map[string]string{
	"454a2ab3": "bid(uint256)",
	"96b5a755": "cancelAuction(uint256)",
	"878eb368": "cancelAuctionWhenPaused(uint256)",
	"27ebe40a": "createAuction(uint256,uint256,uint256,uint256,address)",
	"78bd7935": "getAuction(uint256)",
	"c55d0f56": "getCurrentPrice(uint256)",
	"76190f8f": "isSiringClockAuction()",
	"dd1b7a0f": "nonFungibleContract()",
	"8da5cb5b": "owner()",
	"83b5ff8b": "ownerCut()",
	"8456cb59": "pause()",
	"5c975abb": "paused()",
	"f2fde38b": "transferOwnership(address)",
	"3f4ba83a": "unpause()",
	"5fd8c710": "withdrawBalance()",
}

// SiringClockAuctionBin is the compiled bytecode used for deploying new contracts.
var SiringClockAuctionBin = "0x60606040526000805460a060020a60ff02191690556004805460ff19166001179055341561002c57600080fd5b604051604080610e32833981016040528080519190602001805160008054600160a060020a03191633600160a060020a0316178155909250839150829061271082111561007857600080fd5b50600281905581600160a060020a0381166301ffc9a77f9a20483d000000000000000000000000000000000000000000000000000000006040517c010000000000000000000000000000000000000000000000000000000063ffffffff84160281527fffffffff000000000000000000000000000000000000000000000000000000009091166004820152602401602060405180830381600087803b151561011f57600080fd5b5af1151561012c57600080fd5b50505060405180519050151561014157600080fd5b60018054600160a060020a03909216600160a060020a031990921691909117905550505050610cbd806101756000396000f3006060604052600436106100c15763ffffffff60e060020a60003504166327ebe40a81146100c65780633f4ba83a146100f3578063454a2ab31461011a5780635c975abb146101255780635fd8c7101461013857806376190f8f1461014b57806378bd79351461015e57806383b5ff8b146101af5780638456cb59146101d4578063878eb368146101e75780638da5cb5b146101fd57806396b5a7551461022c578063c55d0f5614610242578063dd1b7a0f14610258578063f2fde38b1461026b575b600080fd5b34156100d157600080fd5b6100f1600435602435604435606435600160a060020a036084351661028a565b005b34156100fe57600080fd5b610106610362565b604051901515815260200160405180910390f35b6100f16004356103e6565b341561013057600080fd5b610106610436565b341561014357600080fd5b6100f1610446565b341561015657600080fd5b6101066104bc565b341561016957600080fd5b6101746004356104c5565b604051600160a060020a03909516855260208501939093526040808501929092526060840152608083019190915260a0909101905180910390f35b34156101ba57600080fd5b6101c2610552565b60405190815260200160405180910390f35b34156101df57600080fd5b610106610558565b34156101f257600080fd5b6100f16004356105e1565b341561020857600080fd5b61021061064e565b604051600160a060020a03909116815260200160405180910390f35b341561023757600080fd5b6100f160043561065d565b341561024d57600080fd5b6101c26004356106ab565b341561026357600080fd5b6102106106dd565b341561027657600080fd5b6100f1600160a060020a03600435166106ec565b610292610c63565b6001608060020a03851685146102a757600080fd5b6001608060020a03841684146102bc57600080fd5b67ffffffffffffffff831683146102d257600080fd5b60015433600160a060020a039081169116146102ed57600080fd5b6102f78287610742565b60a06040519081016040528083600160a060020a03168152602001866001608060020a03168152602001856001608060020a031681526020018467ffffffffffffffff1681526020014267ffffffffffffffff16815250905061035a86826107b9565b505050505050565b6000805433600160a060020a0390811691161461037e57600080fd5b60005460a060020a900460ff16151561039657600080fd5b6000805474ff0000000000000000000000000000000000000000191690557f7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b3360405160405180910390a150600190565b60015460009033600160a060020a0390811691161461040457600080fd5b50600081815260036020526040902054600160a060020a03166104278234610954565b506104328183610a85565b5050565b60005460a060020a900460ff1681565b60015460008054600160a060020a03928316923381169116148061047b575081600160a060020a031633600160a060020a0316145b151561048657600080fd5b81600160a060020a03166108fc30600160a060020a0316319081150290604051600060405180830381858888f150505050505050565b60045460ff1681565b600081815260036020526040812081908190819081906104e481610adb565b15156104ef57600080fd5b80546001820154600290920154600160a060020a03909116986001608060020a038084169950700100000000000000000000000000000000909304909216965067ffffffffffffffff808216965068010000000000000000909104169350915050565b60025481565b6000805433600160a060020a0390811691161461057457600080fd5b60005460a060020a900460ff161561058b57600080fd5b6000805474ff0000000000000000000000000000000000000000191660a060020a1790557f6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff62560405160405180910390a150600190565b6000805460a060020a900460ff1615156105fa57600080fd5b60005433600160a060020a0390811691161461061557600080fd5b50600081815260036020526040902061062d81610adb565b151561063857600080fd5b8054610432908390600160a060020a0316610afc565b600054600160a060020a031681565b60008181526003602052604081209061067582610adb565b151561068057600080fd5b508054600160a060020a03908116903316811461069c57600080fd5b6106a68382610afc565b505050565b60008181526003602052604081206106c281610adb565b15156106cd57600080fd5b6106d681610b46565b9392505050565b600154600160a060020a031681565b60005433600160a060020a0390811691161461070757600080fd5b600160a060020a0381161561073f576000805473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a0383161790555b50565b600154600160a060020a03166323b872dd83308460405160e060020a63ffffffff8616028152600160a060020a0393841660048201529190921660248201526044810191909152606401600060405180830381600087803b15156107a557600080fd5b5af115156107b257600080fd5b5050505050565b603c816060015167ffffffffffffffff1610156107d557600080fd5b600082815260036020526040902081908151815473ffffffffffffffffffffffffffffffffffffffff1916600160a060020a039190911617815560208201516001820180546fffffffffffffffffffffffffffffffff19166001608060020a039290921691909117905560408201516001820180546001608060020a03928316700100000000000000000000000000000000029216919091179055606082015160028201805467ffffffffffffffff191667ffffffffffffffff9290921691909117905560808201516002909101805467ffffffffffffffff9290921668010000000000000000026fffffffffffffffff000000000000000019909216919091179055507fa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba78260208301516001608060020a031683604001516001608060020a0316846060015167ffffffffffffffff166040518085815260200184815260200183815260200182815260200194505050505060405180910390a15050565b6000828152600360205260408120818080808061097086610adb565b151561097b57600080fd5b61098486610b46565b94508488101561099357600080fd5b8554600160a060020a031693506109a989610bcd565b60008511156109f3576109bb85610c1a565b92508285039150600160a060020a03841682156108fc0283604051600060405180830381858888f1935050505015156109f357600080fd5b50838703600160a060020a03331681156108fc0282604051600060405180830381858888f193505050501515610a2857600080fd5b7f4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd28986336040519283526020830191909152600160a060020a03166040808301919091526060909101905180910390a15092979650505050505050565b600154600160a060020a031663a9059cbb838360405160e060020a63ffffffff8516028152600160a060020a0390921660048301526024820152604401600060405180830381600087803b15156107a557600080fd5b6002015460006801000000000000000090910467ffffffffffffffff161190565b610b0582610bcd565b610b0f8183610a85565b7f2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df8260405190815260200160405180910390a15050565b6002810154600090819068010000000000000000900467ffffffffffffffff16421115610b8c5750600282015468010000000000000000900467ffffffffffffffff1642035b600183015460028401546106d6916001608060020a0380821692700100000000000000000000000000000000909204169067ffffffffffffffff1684610c26565b6000908152600360205260408120805473ffffffffffffffffffffffffffffffffffffffff19168155600181019190915560020180546fffffffffffffffffffffffffffffffff19169055565b60025461271091020490565b6000808080858510610c3a57869350610c58565b878703925085858402811515610c4c57fe5b05915081880190508093505b505050949350505050565b60a06040519081016040908152600080835260208301819052908201819052606082018190526080820152905600a165627a7a7230582014a23cbca84564759f995822c14037e8dc35b7daab68d7b1fa19dfcb36c920bb0029"

// DeploySiringClockAuction deploys a new Ethereum contract, binding an instance of SiringClockAuction to it.
func DeploySiringClockAuction(auth *bind.TransactOpts, backend bind.ContractBackend, _nftAddr common.Address, _cut *big.Int) (common.Address, *types.Transaction, *SiringClockAuction, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(SiringClockAuctionBin), backend, _nftAddr, _cut)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// SiringClockAuction is an auto generated Go binding around an Ethereum contract.
type SiringClockAuction struct {
	SiringClockAuctionCaller     // Read-only binding to the contract
	SiringClockAuctionTransactor // Write-only binding to the contract
	SiringClockAuctionFilterer   // Log filterer for contract events
}

// SiringClockAuctionCaller is an auto generated read-only Go binding around an Ethereum contract.
type SiringClockAuctionCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SiringClockAuctionFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SiringClockAuctionSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SiringClockAuctionSession struct {
	Contract     *SiringClockAuction // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// SiringClockAuctionCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SiringClockAuctionCallerSession struct {
	Contract *SiringClockAuctionCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// SiringClockAuctionTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SiringClockAuctionTransactorSession struct {
	Contract     *SiringClockAuctionTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// SiringClockAuctionRaw is an auto generated low-level Go binding around an Ethereum contract.
type SiringClockAuctionRaw struct {
	Contract *SiringClockAuction // Generic contract binding to access the raw methods on
}

// SiringClockAuctionCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SiringClockAuctionCallerRaw struct {
	Contract *SiringClockAuctionCaller // Generic read-only contract binding to access the raw methods on
}

// SiringClockAuctionTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SiringClockAuctionTransactorRaw struct {
	Contract *SiringClockAuctionTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSiringClockAuction creates a new instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuction(address common.Address, backend bind.ContractBackend) (*SiringClockAuction, error) {
	contract, err := bindSiringClockAuction(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuction{SiringClockAuctionCaller: SiringClockAuctionCaller{contract: contract}, SiringClockAuctionTransactor: SiringClockAuctionTransactor{contract: contract}, SiringClockAuctionFilterer: SiringClockAuctionFilterer{contract: contract}}, nil
}

// NewSiringClockAuctionCaller creates a new read-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionCaller(address common.Address, caller bind.ContractCaller) (*SiringClockAuctionCaller, error) {
	contract, err := bindSiringClockAuction(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionCaller{contract: contract}, nil
}

// NewSiringClockAuctionTransactor creates a new write-only instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionTransactor(address common.Address, transactor bind.ContractTransactor) (*SiringClockAuctionTransactor, error) {
	contract, err := bindSiringClockAuction(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionTransactor{contract: contract}, nil
}

// NewSiringClockAuctionFilterer creates a new log filterer instance of SiringClockAuction, bound to a specific deployed contract.
func NewSiringClockAuctionFilterer(address common.Address, filterer bind.ContractFilterer) (*SiringClockAuctionFilterer, error) {
	contract, err := bindSiringClockAuction(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionFilterer{contract: contract}, nil
}

// bindSiringClockAuction binds a generic wrapper to an already deployed contract.
func bindSiringClockAuction(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(SiringClockAuctionABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.SiringClockAuctionCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.SiringClockAuctionTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SiringClockAuction *SiringClockAuctionCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SiringClockAuction.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SiringClockAuction *SiringClockAuctionTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.contract.Transact(opts, method, params...)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_SiringClockAuction *SiringClockAuctionCaller) GetAuction(opts *bind.CallOpts, _tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "getAuction", _tokenId)

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
func (_SiringClockAuction *SiringClockAuctionSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetAuction is a free data retrieval call binding the contract method 0x78bd7935.
//
// Solidity: function getAuction(uint256 _tokenId) view returns(address seller, uint256 startingPrice, uint256 endingPrice, uint256 duration, uint256 startedAt)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetAuction(_tokenId *big.Int) (struct {
	Seller        common.Address
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	StartedAt     *big.Int
}, error) {
	return _SiringClockAuction.Contract.GetAuction(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) GetCurrentPrice(opts *bind.CallOpts, _tokenId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "getCurrentPrice", _tokenId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// GetCurrentPrice is a free data retrieval call binding the contract method 0xc55d0f56.
//
// Solidity: function getCurrentPrice(uint256 _tokenId) view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) GetCurrentPrice(_tokenId *big.Int) (*big.Int, error) {
	return _SiringClockAuction.Contract.GetCurrentPrice(&_SiringClockAuction.CallOpts, _tokenId)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) IsSiringClockAuction(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "isSiringClockAuction")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// IsSiringClockAuction is a free data retrieval call binding the contract method 0x76190f8f.
//
// Solidity: function isSiringClockAuction() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) IsSiringClockAuction() (bool, error) {
	return _SiringClockAuction.Contract.IsSiringClockAuction(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) NonFungibleContract(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "nonFungibleContract")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// NonFungibleContract is a free data retrieval call binding the contract method 0xdd1b7a0f.
//
// Solidity: function nonFungibleContract() view returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) NonFungibleContract() (common.Address, error) {
	return _SiringClockAuction.Contract.NonFungibleContract(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiringClockAuction *SiringClockAuctionCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiringClockAuction *SiringClockAuctionSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Owner() (common.Address, error) {
	return _SiringClockAuction.Contract.Owner(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCaller) OwnerCut(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "ownerCut")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// OwnerCut is a free data retrieval call binding the contract method 0x83b5ff8b.
//
// Solidity: function ownerCut() view returns(uint256)
func (_SiringClockAuction *SiringClockAuctionCallerSession) OwnerCut() (*big.Int, error) {
	return _SiringClockAuction.Contract.OwnerCut(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _SiringClockAuction.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_SiringClockAuction *SiringClockAuctionCallerSession) Paused() (bool, error) {
	return _SiringClockAuction.Contract.Paused(&_SiringClockAuction.CallOpts)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) Bid(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "bid", _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SiringClockAuction *SiringClockAuctionSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// Bid is a paid mutator transaction binding the contract method 0x454a2ab3.
//
// Solidity: function bid(uint256 _tokenId) payable returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Bid(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Bid(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuction(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuction", _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuction is a paid mutator transaction binding the contract method 0x96b5a755.
//
// Solidity: function cancelAuction(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuction(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuction(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CancelAuctionWhenPaused(opts *bind.TransactOpts, _tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "cancelAuctionWhenPaused", _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CancelAuctionWhenPaused is a paid mutator transaction binding the contract method 0x878eb368.
//
// Solidity: function cancelAuctionWhenPaused(uint256 _tokenId) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CancelAuctionWhenPaused(_tokenId *big.Int) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CancelAuctionWhenPaused(&_SiringClockAuction.TransactOpts, _tokenId)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) CreateAuction(opts *bind.TransactOpts, _tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "createAuction", _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SiringClockAuction *SiringClockAuctionSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// CreateAuction is a paid mutator transaction binding the contract method 0x27ebe40a.
//
// Solidity: function createAuction(uint256 _tokenId, uint256 _startingPrice, uint256 _endingPrice, uint256 _duration, address _seller) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) CreateAuction(_tokenId *big.Int, _startingPrice *big.Int, _endingPrice *big.Int, _duration *big.Int, _seller common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.CreateAuction(&_SiringClockAuction.TransactOpts, _tokenId, _startingPrice, _endingPrice, _duration, _seller)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Pause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Pause(&_SiringClockAuction.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiringClockAuction *SiringClockAuctionSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SiringClockAuction.Contract.TransferOwnership(&_SiringClockAuction.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns(bool)
func (_SiringClockAuction *SiringClockAuctionTransactorSession) Unpause() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.Unpause(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactor) WithdrawBalance(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SiringClockAuction.contract.Transact(opts, "withdrawBalance")
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// WithdrawBalance is a paid mutator transaction binding the contract method 0x5fd8c710.
//
// Solidity: function withdrawBalance() returns()
func (_SiringClockAuction *SiringClockAuctionTransactorSession) WithdrawBalance() (*types.Transaction, error) {
	return _SiringClockAuction.Contract.WithdrawBalance(&_SiringClockAuction.TransactOpts)
}

// SiringClockAuctionAuctionCancelledIterator is returned from FilterAuctionCancelled and is used to iterate over the raw logs and unpacked data for AuctionCancelled events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelledIterator struct {
	Event *SiringClockAuctionAuctionCancelled // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCancelledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCancelled)
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
		it.Event = new(SiringClockAuctionAuctionCancelled)
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
func (it *SiringClockAuctionAuctionCancelledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCancelledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCancelled represents a AuctionCancelled event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCancelled struct {
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAuctionCancelled is a free log retrieval operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCancelled(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCancelledIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCancelledIterator{contract: _SiringClockAuction.contract, event: "AuctionCancelled", logs: logs, sub: sub}, nil
}

// WatchAuctionCancelled is a free log subscription operation binding the contract event 0x2809c7e17bf978fbc7194c0a694b638c4215e9140cacc6c38ca36010b45697df.
//
// Solidity: event AuctionCancelled(uint256 tokenId)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCancelled(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCancelled) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCancelled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCancelled)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
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
func (_SiringClockAuction *SiringClockAuctionFilterer) ParseAuctionCancelled(log types.Log) (*SiringClockAuctionAuctionCancelled, error) {
	event := new(SiringClockAuctionAuctionCancelled)
	if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCancelled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiringClockAuctionAuctionCreatedIterator is returned from FilterAuctionCreated and is used to iterate over the raw logs and unpacked data for AuctionCreated events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreatedIterator struct {
	Event *SiringClockAuctionAuctionCreated // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionCreated)
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
		it.Event = new(SiringClockAuctionAuctionCreated)
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
func (it *SiringClockAuctionAuctionCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionCreated represents a AuctionCreated event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionCreated struct {
	TokenId       *big.Int
	StartingPrice *big.Int
	EndingPrice   *big.Int
	Duration      *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAuctionCreated is a free log retrieval operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionCreated(opts *bind.FilterOpts) (*SiringClockAuctionAuctionCreatedIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionCreatedIterator{contract: _SiringClockAuction.contract, event: "AuctionCreated", logs: logs, sub: sub}, nil
}

// WatchAuctionCreated is a free log subscription operation binding the contract event 0xa9c8dfcda5664a5a124c713e386da27de87432d5b668e79458501eb296389ba7.
//
// Solidity: event AuctionCreated(uint256 tokenId, uint256 startingPrice, uint256 endingPrice, uint256 duration)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionCreated(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionCreated) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionCreated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionCreated)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
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
func (_SiringClockAuction *SiringClockAuctionFilterer) ParseAuctionCreated(log types.Log) (*SiringClockAuctionAuctionCreated, error) {
	event := new(SiringClockAuctionAuctionCreated)
	if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiringClockAuctionAuctionSuccessfulIterator is returned from FilterAuctionSuccessful and is used to iterate over the raw logs and unpacked data for AuctionSuccessful events raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessfulIterator struct {
	Event *SiringClockAuctionAuctionSuccessful // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionAuctionSuccessful)
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
		it.Event = new(SiringClockAuctionAuctionSuccessful)
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
func (it *SiringClockAuctionAuctionSuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionAuctionSuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionAuctionSuccessful represents a AuctionSuccessful event raised by the SiringClockAuction contract.
type SiringClockAuctionAuctionSuccessful struct {
	TokenId    *big.Int
	TotalPrice *big.Int
	Winner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAuctionSuccessful is a free log retrieval operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterAuctionSuccessful(opts *bind.FilterOpts) (*SiringClockAuctionAuctionSuccessfulIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionAuctionSuccessfulIterator{contract: _SiringClockAuction.contract, event: "AuctionSuccessful", logs: logs, sub: sub}, nil
}

// WatchAuctionSuccessful is a free log subscription operation binding the contract event 0x4fcc30d90a842164dd58501ab874a101a3749c3d4747139cefe7c876f4ccebd2.
//
// Solidity: event AuctionSuccessful(uint256 tokenId, uint256 totalPrice, address winner)
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchAuctionSuccessful(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionAuctionSuccessful) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "AuctionSuccessful")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionAuctionSuccessful)
				if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
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
func (_SiringClockAuction *SiringClockAuctionFilterer) ParseAuctionSuccessful(log types.Log) (*SiringClockAuctionAuctionSuccessful, error) {
	event := new(SiringClockAuctionAuctionSuccessful)
	if err := _SiringClockAuction.contract.UnpackLog(event, "AuctionSuccessful", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiringClockAuctionPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the SiringClockAuction contract.
type SiringClockAuctionPauseIterator struct {
	Event *SiringClockAuctionPause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionPause)
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
		it.Event = new(SiringClockAuctionPause)
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
func (it *SiringClockAuctionPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionPause represents a Pause event raised by the SiringClockAuction contract.
type SiringClockAuctionPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterPause(opts *bind.FilterOpts) (*SiringClockAuctionPauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionPauseIterator{contract: _SiringClockAuction.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: event Pause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionPause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionPause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
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
func (_SiringClockAuction *SiringClockAuctionFilterer) ParsePause(log types.Log) (*SiringClockAuctionPause, error) {
	event := new(SiringClockAuctionPause)
	if err := _SiringClockAuction.contract.UnpackLog(event, "Pause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SiringClockAuctionUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the SiringClockAuction contract.
type SiringClockAuctionUnpauseIterator struct {
	Event *SiringClockAuctionUnpause // Event containing the contract specifics and raw log

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
func (it *SiringClockAuctionUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SiringClockAuctionUnpause)
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
		it.Event = new(SiringClockAuctionUnpause)
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
func (it *SiringClockAuctionUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SiringClockAuctionUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SiringClockAuctionUnpause represents a Unpause event raised by the SiringClockAuction contract.
type SiringClockAuctionUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) FilterUnpause(opts *bind.FilterOpts) (*SiringClockAuctionUnpauseIterator, error) {

	logs, sub, err := _SiringClockAuction.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &SiringClockAuctionUnpauseIterator{contract: _SiringClockAuction.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: event Unpause()
func (_SiringClockAuction *SiringClockAuctionFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *SiringClockAuctionUnpause) (event.Subscription, error) {

	logs, sub, err := _SiringClockAuction.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SiringClockAuctionUnpause)
				if err := _SiringClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
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
func (_SiringClockAuction *SiringClockAuctionFilterer) ParseUnpause(log types.Log) (*SiringClockAuctionUnpause, error) {
	event := new(SiringClockAuctionUnpause)
	if err := _SiringClockAuction.contract.UnpackLog(event, "Unpause", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
