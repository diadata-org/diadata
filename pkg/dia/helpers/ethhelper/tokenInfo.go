package ethhelper

import (
	"encoding/json"
	"io/ioutil"
	"math/big"
	"os"
	"os/user"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const tokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"},{\"name\":\"_releaseTime\",\"type\":\"uint256\"}],\"name\":\"mintTimelocked\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// tokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type tokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
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

// newTokenCaller creates a new read-only instance of token, bound to a specific deployed contract.
func newTokenCaller(address common.Address, caller bind.ContractCaller) (*tokenCaller, error) {
	contract, err := bindToken(address, caller, nil)
	if err != nil {
		return nil, err
	}
	return &tokenCaller{contract: contract}, nil
}

// bindToken binds a generic wrapper to an already deployed contract.
func bindToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(tokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, nil), nil
}

// GetDecimals returns the decimals of token with address @address
func GetDecimals(address common.Address) (int, error) {
	ethConn, err := NewETHClient()
	if err != nil {
		return 0, err
	}
	tc, err := newTokenCaller(address, ethConn)
	if err != nil {
		log.Error("error: ", err)
	}
	var decimals = new(*big.Int)
	err = tc.contract.Call(&bind.CallOpts{}, decimals, "decimals")
	if err != nil {
		return 0, err
	}
	return int((*decimals).Int64()), nil

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
func GetAddressesFromFile() (addresses []string, err error) {
	configPath := ConfigFilePath("tokens_list")
	jsonFile, err := os.Open(configPath)
	if err != nil {
		log.Errorln("Error opening file", err)
		return
	}
	defer jsonFile.Close()

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
	json.Unmarshal(byteData, &tokeninfo)
	for _, token := range tokeninfo.Tokens {
		addresses = append(addresses, token.Address)
	}
	return
}
