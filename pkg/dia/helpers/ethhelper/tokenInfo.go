package ethhelper

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"math/big"
	"os"
	"os/user"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const tokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_releaseTime\",\"type\":\"uint256\"}],\"name\":\"mintTimelocked\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]" //nolint:gosec

// const tokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\": \"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\": \"_owner\",\"type\":\"address\"}],\"name\": \"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"}]"

// TokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenCaller struct {
	Contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NewETHClient returns an ethclient, using either our node for production or infura for testing
func NewETHClient() (*ethclient.Client, error) {
	executionMode := os.Getenv("EXEC_MODE")
	var connection *ethclient.Client
	var err error
	if executionMode == "production" {
		connection, err = ethclient.Dial("http://159.69.120.42:8545/")
		if err != nil {
			log.Error("Error connecting Eth Client")
		}
	} else {
		connection, err = ethclient.Dial("https://mainnet.infura.io/v3/806b0419b2d041869fc83727e0043236")
		if err != nil {
			log.Error("Error connecting Eth Client")
		}
	}
	return connection, err
}

// NewTokenCaller creates a new read-only instance of token, bound to a specific deployed contract.
func NewTokenCaller(address common.Address, caller bind.ContractCaller) (*TokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &TokenCaller{Contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

func GetBalanceOf(tokenAddress common.Address, walletAddress common.Address, client *ethclient.Client) (*big.Int, error) {
	var balance []interface{}
	tokenCaller, err := NewTokenCaller(tokenAddress, client)
	if err != nil {
		return big.NewInt(0), err
	}
	err = tokenCaller.Contract.Call(&bind.CallOpts{}, &balance, "balanceOf", walletAddress)
	if err != nil {
		return big.NewInt(0), err
	}
	if len(balance) > 0 {
		return balance[0].(*big.Int), nil
	}
	return big.NewInt(0), errors.New("Not enough data.")
}

// ETHAddressToAsset takes an Ethereum address and returns the underlying
// token as a dia.Asset.
func ETHAddressToAsset(address common.Address, client *ethclient.Client, blockchainName string) (dia.Asset, error) {
	var asset dia.Asset
	tc, err := NewTokenCaller(address, client)
	if err != nil {
		log.Error("error: ", err)
	}

	asset.Address = address.Hex()

	var symbol []interface{}
	err = tc.Contract.Call(&bind.CallOpts{}, &symbol, "symbol")
	if err != nil {
		return dia.Asset{}, err
	}
	asset.Symbol = symbol[0].(string)

	var name []interface{}
	err = tc.Contract.Call(&bind.CallOpts{}, &name, "name")
	if err != nil {
		return dia.Asset{}, err
	}
	asset.Name = name[0].(string)

	var decimals []interface{}
	err = tc.Contract.Call(&bind.CallOpts{}, &decimals, "decimals")
	if err != nil {
		return dia.Asset{}, err
	}
	aux := decimals[0].(*big.Int)
	asset.Decimals = uint8(aux.Int64())
	asset.Blockchain = blockchainName
	return asset, err
}

func ConfigFilePath(filename string) string {
	usr, _ := user.Current()
	dir := usr.HomeDir
	if dir == "/root" || dir == "/home" {
		return "/config/token_supply/" + filename + ".json" //hack for docker...
	}
	if dir == "/home/travis" {
		return "../config/token_supply/" + filename + ".json" //hack for travis
	}
	return os.Getenv("GOPATH") + "/src/github.com/diadata-org/diadata/config/token_supply/" + filename + ".json"
}

// GetAddressesFromFile fetches token addresses from a config file available here:
// https://etherscan.io/exportData?type=open-source-contract-codes
func GetAddressesFromFile(filename string) (addresses []string, err error) {
	jsonFile, err := os.Open(ConfigFilePath(filename))
	if err != nil {
		log.Errorln("Error opening file", err)
		return
	}
	defer func() {
		cerr := jsonFile.Close()
		if err == nil {
			err = cerr
		}
	}()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(err)
		return
	}
	// For now we only return token address. Name can be retrieved through contract.
	type token struct {
		Address string
		Name    string
	}
	type TokenInfo struct {
		Tokens []token `json:"Tokens"`
	}
	var tokeninfo TokenInfo
	err = json.Unmarshal(byteData, &tokeninfo)
	if err != nil {
		return
	}
	for _, token := range tokeninfo.Tokens {
		addresses = append(addresses, token.Address)
	}
	return
}
