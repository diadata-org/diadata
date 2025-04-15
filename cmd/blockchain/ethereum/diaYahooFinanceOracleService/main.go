package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaOracleServiceV2 "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

var diaBaseUrl string

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	sleepSeconds, err := strconv.Atoi(utils.Getenv("SLEEP_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse sleepSeconds: %v")
	}
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v")
	}
	mandatoryFrequencySeconds, err := strconv.Atoi(utils.Getenv("MANDATORY_FREQUENCY_SECONDS", "0"))
	if err != nil {
		log.Fatalf("Failed to parse mandatoryFrequencySeconds: %v")
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v")
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v")
	}
	pairStr := utils.Getenv("PAIRS", "")
	assetsStr := utils.Getenv("ASSETS", "")
	cbAssetsStr := utils.Getenv("CB_ASSETS", "")
	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")

	// Parse assets
	assetsParsed := strings.Split(assetsStr, ",")

	addresses := []string{}
	blockchains := []string{}

	for _, asset := range assetsParsed {
		if asset == "" {
			continue
		}
		entries := strings.Split(asset, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		addresses = append(addresses, strings.TrimSpace(entries[1]))
	}

	// Parse pairs
	pairs := []string{}
	pairsParsed := strings.Split(pairStr, ",")

	for _, pair := range pairsParsed {
		if pair == "" {
			continue
		}
		pairs = append(pairs, strings.TrimSpace(pair))
	}

	// Parse Coinbase Assets
	cbAssets := []string{}
	cbAssetsParsed := strings.Split(cbAssetsStr, ",")

	for _, cbAsset := range cbAssetsParsed {
		cbAssets = append(cbAssets, strings.TrimSpace(cbAsset))
	}

	oldPrices := make(map[int]float64)
	oldAssetPrices := make(map[int]float64)
	oldCbAssetPrices := make(map[int]float64)

	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaOracleServiceV2.DIAOracleV2
	err = deployOrBindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	/*
	 * Update Oracle periodically
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i, pair := range pairs {
					oldPrice := oldPrices[i]
					log.Println("old price", oldPrice)
					oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, conn, pair)
					oldPrices[i] = oldPrice
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
				for i, cbAsset := range cbAssets {
					oldCbAssetPrice := oldCbAssetPrices[i]
					log.Println("old CB Asset price", oldCbAssetPrice)
					oldCbAssetPrice, err = periodicCbAssetOracleUpdateHelper(oldCbAssetPrice, deviationPermille, auth, contract, conn, cbAsset)
					oldCbAssetPrices[i] = oldCbAssetPrice
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
				for i, address := range addresses {
					oldAssetPrice := oldAssetPrices[i]
					blockchain := blockchains[i]
					log.Println("old asset price", oldAssetPrice)
					oldAssetPrice, err = periodicAssetOracleUpdateHelper(oldAssetPrice, deviationPermille, auth, contract, conn, blockchain, address)
					oldAssetPrices[i] = oldAssetPrice
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
			}
		}
	}()

	if mandatoryFrequencySeconds > 0 {
		mandatorytickerticker := time.NewTicker(time.Duration(mandatoryFrequencySeconds) * time.Second)
		go func() {
			for {
				select {
				case <-mandatorytickerticker.C:
					for i, pair := range pairs {
						oldPrice := oldPrices[i]
						log.Println("old price", oldPrice)
						oldPrice, err = oracleUpdateHelper(oldPrice, auth, contract, conn, pair)
						oldPrices[i] = oldPrice
						if err != nil {
							log.Println(err)
						}
						time.Sleep(time.Duration(sleepSeconds) * time.Second)
					}
					for i, cbAsset := range cbAssets {
						oldCbAssetPrice := oldCbAssetPrices[i]
						log.Println("old CB Asset price", oldCbAssetPrice)
						oldCbAssetPrice, err = cbAssetOracleUpdateHelper(oldCbAssetPrice, auth, contract, conn, cbAsset)
						oldCbAssetPrices[i] = oldCbAssetPrice
						if err != nil {
							log.Println(err)
						}
						time.Sleep(time.Duration(sleepSeconds) * time.Second)
					}
				}
			}
		}()
	}
	select {}
}

func oracleUpdateHelper(oldPrice float64, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, pair string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	rawQ, err = getForeignQuotationFromDia(pair)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", pair, err)
		return oldPrice, err
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	err = updateForeignQuotation(rawQ, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return oldPrice, err
	}

	return newPrice, nil
}

func cbAssetOracleUpdateHelper(oldPrice float64, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, asset string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	rawQ, err = getForeignQuotationFromCoinbase(asset)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from Coinbase: %v", asset, err)
		return oldPrice, err
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	err = updateQuotation(rawQ, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return oldPrice, err
	}

	return newPrice, nil
}

func periodicAssetOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string) (float64, error) {

	// Get quotation for token and update Oracle
	rawQ, err := getAssetQuotationFromDia(blockchain, address)
	if err != nil {
		log.Fatalf("Failed to retrieve %s - %s quotation data from DIA: %v", blockchain, address, err)
		return oldPrice, err
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		err = updateQuotation(rawQ, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
}

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, pair string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	rawQ, err = getForeignQuotationFromDia(pair)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", pair, err)
		return oldPrice, err
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		err = updateForeignQuotation(rawQ, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
}

func periodicCbAssetOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, asset string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	rawQ, err = getForeignQuotationFromCoinbase(asset)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from Coinbase: %v", asset, err)
		return oldPrice, err
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		err = updateQuotation(rawQ, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **diaOracleServiceV2.DIAOracleV2) error {
	var err error
	if deployedContract != "" {
		*contract, err = diaOracleServiceV2.NewDIAOracleV2(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaOracleServiceV2.DeployDIAOracleV2(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

func updateForeignQuotation(quotation *models.Quotation, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	symbol := quotation.Symbol
	price := quotation.Price
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateQuotation(quotation *models.Quotation, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	symbol := quotation.Symbol + "/USD"
	price := quotation.Price
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateOracle(
	client *ethclient.Client,
	contract *diaOracleServiceV2.DIAOracleV2,
	auth *bind.TransactOpts,
	key string,
	value int64,
	timestamp int64) error {

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	// Get 110% of the gas price
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)

	// Write values to smart contract
	tx, err := contract.SetValue(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		//GasLimit: 1000725,
		GasPrice: gasPrice,
	}, key, big.NewInt(value), big.NewInt(timestamp))
	if err != nil {
		return err
	}
	log.Printf("Gas price: %d\n", tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Printf("Data: %x\n", tx.Data())
	log.Printf("Nonce: %d\n", tx.Nonce())
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getForeignQuotationFromCoinbase(asset string) (*models.Quotation, error) {
	response, err := http.Get("https://api.coinbase.com/v2/exchange-rates")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on Coinbase API with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	var quotation models.Quotation
	invPrice := gjson.Get(string(contents), "data.rates." + asset).Float()
	price := 1/invPrice
	quotation.Price = price
	quotation.Symbol = asset

	return &quotation, nil
}

func getForeignQuotationFromDia(pair string) (*models.Quotation, error) {
	response, err := http.Get(diaBaseUrl + "/v1/foreignQuotation/YahooFinance/" + pair)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func getAssetQuotationFromDia(blockchain, address string) (*models.Quotation, error) {
	response, err := http.Get(diaBaseUrl + "/v1/assetQuotation/" + blockchain + "/" + address)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}
