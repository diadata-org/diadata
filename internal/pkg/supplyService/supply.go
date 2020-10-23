package supplyservice

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"os"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	utils "github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

// GetLockedWalletsFromConfig returns a map which maps an asset to the list of its locked wallets
func GetLockedWalletsFromConfig(filename string) (map[string][]string, error) {

	fileName := fmt.Sprintf("../../../config/token_supply/%s.json", filename)
	jsonFile, err := os.Open(fileName)
	if err != nil {
		log.Errorln("Error opening file", err)
		return map[string][]string{}, err
	}
	defer jsonFile.Close()

	byteData, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		log.Error(err)
		return map[string][]string{}, err
	}

	type lockedAsset struct {
		Address       string   `json:"asset"`
		LockedWallets []string `json:"wallets"`
	}
	type lockedAssetList struct {
		AllAssets []lockedAsset `json:"lockedWallets"`
	}
	var allAssets lockedAssetList
	json.Unmarshal(byteData, &allAssets)
	// make map[string][]string from allAssets. This accounts for erroneous addition of new entry
	// for already existing asset in config file.
	allAssetsMap := make(map[string][]string)
	var diff []string
	for _, asset := range allAssets.AllAssets {
		if _, ok := allAssetsMap[asset.Address]; !ok {
			allAssetsMap[asset.Address] = asset.LockedWallets
		} else {
			diff = utils.SliceDifference(asset.LockedWallets, allAssetsMap[asset.Address])
			allAssetsMap[asset.Address] = append(allAssetsMap[asset.Address], diff...)
		}
	}

	return allAssetsMap, nil
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
		log.Error(err)
		return
	}

	walletBal, err := instance.BalanceOf(&bind.CallOpts{}, common.HexToAddress(walletAddr))
	if err != nil {
		log.Errorf("Failed to retrieve token owner balance from wallet %s: %v \n", walletAddr, err)
		return
	}

	fbal := new(big.Float)
	fbal.SetString(walletBal.String())

	balance, _ = new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals)))).Float64()
	return
}

// GetTotalSupplyfromMainNet returns total supply minus wallets' balances from list of wallets
func GetTotalSupplyfromMainNet(tokenAddress string, lockedWallets []string, client *ethclient.Client) (supply dia.Supply, err error) {

	instance, err := NewERC20(common.HexToAddress(tokenAddress), client)
	if err != nil {
		log.Error("error getting token contract: ", err)
		return
	}

	totalSupply, err := instance.TotalSupply(&bind.CallOpts{})
	if err != nil {
		return
	}
	symbol, err := instance.Symbol(&bind.CallOpts{})
	if err != nil {
		return
	}
	name, err := instance.Name(&bind.CallOpts{})
	if err != nil {
		return
	}
	decimals, err := instance.Decimals(&bind.CallOpts{})
	if err != nil {
		return
	}

	// Get total supply
	fbal := new(big.Float)
	fbal.SetString(totalSupply.String())
	valuei := new(big.Float).Quo(fbal, big.NewFloat(math.Pow10(int(decimals))))
	totalSupp, _ := valuei.Float64()

	// Subtract locked wallets' balances from total supply for circulating supply
	circulatingSupply := totalSupp
	for _, walletAddress := range lockedWallets {
		balance, err := GetWalletBalance(walletAddress, tokenAddress, client)
		if err != nil {
			log.Errorf("error getting wallet balance for wallet %s \n", walletAddress)
		}
		circulatingSupply = circulatingSupply - balance
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	supply = dia.Supply{
		Symbol:            symbol,
		Name:              name,
		Supply:            totalSupp,
		CirculatingSupply: circulatingSupply,
		Source:            "diadata.org",
		Time:              time.Unix(int64(header.Time), 0),
		Block:             header.Number.Int64(),
	}

	return
}
