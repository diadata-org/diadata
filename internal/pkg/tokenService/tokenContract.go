package tokenService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// getWalletsFromConfig returns wallet addresses from config file
func getWalletsFromConfig(filename string) ([]string, error) {

	fileName := fmt.Sprintf("../../../config/token_supply/%s.json", filename)
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Errorln("Error opening file", err)
		return []string{}, err
	}
	defer jsonFile.Close()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(err)
	}

	var result map[string][]string
	json.Unmarshal(byteData, &result) //If result becomes struct the coversion with byte array is no more needed

	return result["wallets"], nil
}

// GetWalletBalance returns balance of token with address @tokenAddr in wallet with address @walletAddr
func GetWalletBalance(walletAddr string, tokenAddr string, c *ethclient.Client) (balance float64, err error) {
	instance, err := NewERC20(common.HexToAddress(tokenAddr), c)
	if err != nil {
		log.Fatal(err)
		return
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
		return
	}

	walletBal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddr))
	if err != nil {
		log.Fatalf("Failed to retrieve token owner balance: %v", err)
		return
	}

	fbal := new(big.Float)
	fbal.SetString(walletBal.String())

	balance, _ = new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals)))).Float64()
	return
}

// GetTotalSupplyfromMainNet return total supply minus wallets' balances from config file
func GetTotalSupplyfromMainNet(tokenAddress string, datastore models.Datastore, client *ethclient.Client) (totalSupp, circulatingSupply float64, err error) {

	instance, err := NewERC20(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Fatal(err)
		return
	}
	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
		return
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
		return
	}

	fbal := new(big.Float)
	fbal.SetString(totalSupply.String())

	valuei := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	totalSupp, _ = valuei.Float64()

	wallets, err := getWalletsFromConfig("wallets")
	if err != nil {
		return
	}

	circulatingSupply = totalSupp
	for _, walletAddress := range wallets {

		balance, err := GetWalletBalance(walletAddress, tokenAddress, client)
		if err != nil {
			log.Errorf("error getting wallet balance for wallet %s \n", walletAddress)
		}
		circulatingSupply = circulatingSupply - balance

	}

	return
}
