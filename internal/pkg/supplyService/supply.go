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

var (
	numMaxRetryCG    = 4
	waitRetrySeconds = 60
	apiKey           = utils.Getenv("COINGECKO_SUPPLY_API_KEY", "")
)

type CGCoin struct {
	Address    string       `json:"contract_address"`
	MarketData CGMarketData `json:"market_data"`
	Platform   string       `json:"asset_platform_id"`
}

type CGMarketData struct {
	TotalSupply       float64 `json:"total_supply"`
	CirculatingSupply float64 `json:"circulating_supply"`
}

func GetETHSuppliesFromCG() (supplies []dia.Supply, err error) {
	IDs, err := getAllIDsCG()
	if err != nil {
		return
	}
	for _, ID := range IDs {
		retries := 0
		var coin CGCoin
		var err error
		var status int
		for retries < numMaxRetryCG {
			coin, status, err = getCGCoinInfo(ID)
			if err != nil {
				if status == 429 {
					time.Sleep(time.Duration(waitRetrySeconds) * time.Second)
					log.Info("rate limitation: sleep")
					retries++
				} else {
					log.Error("get coin info: ", err)
					break
				}
			} else {
				time.Sleep(1000 * time.Millisecond)
				break
			}
		}
		if coin.Address != "" && coin.Platform == "ethereum" {
			supply := dia.Supply{
				Asset: dia.Asset{
					Address:    common.HexToAddress(coin.Address).Hex(),
					Blockchain: dia.ETHEREUM,
				},
				Supply:            coin.MarketData.TotalSupply,
				CirculatingSupply: coin.MarketData.CirculatingSupply,
			}
			supplies = append(supplies, supply)
		}
	}
	return
}

func getCGCoinInfo(id string) (coin CGCoin, status int, err error) {
	var resp []byte
	resp, status, err = utils.GetRequest("https://pro-api.coingecko.com/api/v3/coins/" + id + "?x_cg_pro_api_key=" + apiKey)
	if err != nil {
		return
	}
	err = json.Unmarshal(resp, &coin)
	if err != nil {
		return
	}
	if id == "ethereum" {
		coin.Address = "0x0000000000000000000000000000000000000000"
		coin.Platform = "ethereum"
	}
	return
}

func getAllIDsCG() (IDs []string, err error) {
	resp, _, err := utils.GetRequest("https://pro-api.coingecko.com/api/v3/coins/list?x_cg_pro_api_key=" + apiKey)
	if err != nil {
		log.Error("getAllIDsCG: ", err)
		return
	}
	var data []interface{}
	err = json.Unmarshal(resp, &data)
	if err != nil {
		return
	}
	for _, item := range data {
		token := item.(map[string]interface{})
		IDs = append(IDs, token["id"].(string))
	}
	return
}

// GetLockedWalletsFromConfig returns a map which maps an asset to the list of its locked wallets
func GetLockedWalletsFromConfig(filename string) (allAssetsMap map[string][]string, err error) {

	var jsonFile *os.File

	executionMode := os.Getenv("EXEC_MODE")
	if executionMode == "production" {
		jsonFile, err = os.Open(fmt.Sprintf("/config/token_supply/%s.json", filename))
	} else {
		jsonFile, err = os.Open(fmt.Sprintf("../../../config/token_supply/%s.json", filename))
	}
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

	type lockedAsset struct {
		Address       string   `json:"asset"`
		LockedWallets []string `json:"wallets"`
	}
	type lockedAssetList struct {
		AllAssets []lockedAsset `json:"lockedWallets"`
	}
	var allAssets lockedAssetList
	err = json.Unmarshal(byteData, &allAssets)
	if err != nil {
		return
	}

	// make map[string][]string from allAssets. This accounts for erroneous addition of new entry
	// for already existing asset in config file.
	var diff []string
	for _, asset := range allAssets.AllAssets {
		if _, ok := allAssetsMap[asset.Address]; !ok {
			allAssetsMap[asset.Address] = asset.LockedWallets
		} else {
			diff = utils.SliceDifference(asset.LockedWallets, allAssetsMap[asset.Address])
			allAssetsMap[asset.Address] = append(allAssetsMap[asset.Address], diff...)
		}
	}

	return
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
		log.Errorf("failed to retrieve token owner balance from wallet %s: %v \n", walletAddr, err)
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
		var balance float64
		balance, err = GetWalletBalance(walletAddress, tokenAddress, client)
		if err != nil {
			log.Errorf("error getting wallet balance for wallet %s \n", walletAddress)
		}
		circulatingSupply = circulatingSupply - balance
	}

	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	asset := dia.Asset{
		Symbol:     symbol,
		Name:       name,
		Decimals:   decimals,
		Address:    common.HexToAddress(tokenAddress).Hex(),
		Blockchain: dia.ETHEREUM,
	}
	supply = dia.Supply{
		Asset:             asset,
		Supply:            totalSupp,
		CirculatingSupply: circulatingSupply,
		Source:            "diadata.org",
		Time:              time.Unix(int64(header.Time), 0),
	}

	return
}
