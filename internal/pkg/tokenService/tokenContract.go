package tokenService

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"time"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

//getWalletsTokenBalances
func getWalletsTokenBalances(address string) (map[string]interface{}, error) {

	fileName := fmt.Sprintf("../../../config/%s.json", address) // I used address to represent file name
	jsonFile, err := os.Open(fileName)
	var result map[string]interface{} // This may become a struct to properly parse JSON

	if err != nil {
		log.Errorln("Error opening file", err)
		return result, err
	}

	defer jsonFile.Close()

	byteValue, _ := ioutil.ReadAll(jsonFile)
	json.Unmarshal([]byte(byteValue), &result) //If result becomes struct the coversion with byte array is no more needed

	return result, nil

}

func GetTotalSupplyfromMainNet(tokenAddress string, datastore models.Datastore) float64 {
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/your-key-here")
	if err != nil {
		log.Error(err)
	}

	// Address "0xa74476443119A942dE498590Fe1f2454d7D4aC0d"
	tokenAddressInHex := common.HexToAddress(tokenAddress)
	instance, err := NewERC20(tokenAddressInHex, conn)
	if err != nil {
		log.Fatal(err)
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		log.Fatal(err)
	}

	fbal := new(big.Float)

	fbal.SetString(totalSupply.String())
	valuei := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	totalSup, _ := valuei.Float64()
	timeNow := time.Now()

	token := models.Token{
		Symbol:      symbol,
		Name:        name,
		TotalSupply: totalSup,
		Decimals:    int(decimals),
		Source:      "Ethereum Mainnet",
		Time:        timeNow,
	}

	datastore.SaveTokenDetailInflux(token)

	walletBalances, err := getWalletsTokenBalances(tokenAddress)
	if err != nil {
		return 0
	}

	for _, val := range walletBalances {
		// This option is used if we have to use balance value from wallet.
		/*totalSup = totalSup - val["Value"] */
		fmt.Println("val", val) // dummy  to be removed
		//if getting balanceOf from mainnet, val would the wallet address.

		/*
			stringToHex := common.HexToAddress(val["value"])
			ownerBal, err := instance.BalanceOf(&bind.CallOpts{}, stringToHex)
			if err != nil {
				log.Fatalf("Failed to retrieve token owner balance: %v", err)
			}

			fbal := new(big.Float)

			fbal.SetString(ownerBal.String())
			valuei := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
			accountBal, _ := valuei.Float64()

			totalSup = totalSup - accountBal
		*/

	}

	return totalSup
}
