package main

import (
	//"context"
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
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v")
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v")
	}
	assetsStr := utils.Getenv("ASSETS", "")
	assetsParsed := strings.Split(assetsStr, ",")

	addresses := []string{}
	blockchains := []string{}
	
	for _, asset := range assetsParsed {
		entries := strings.Split(asset, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		addresses = append(addresses, strings.TrimSpace(entries[1]))
	}

	oldPrices := make(map[int]float64)

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
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i, address := range addresses {
					blockchain := blockchains[i]
					oldPrice := oldPrices[i]
					log.Println("old price", oldPrice)
					oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, conn, blockchain, address, chainId)
					oldPrices[i] = oldPrice
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string, chainId int64) (float64, error) {

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
		err = updateQuotation(rawQ, auth, contract, conn, chainId)
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

func updateQuotation(quotation *models.Quotation, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, chainId int64) error {
	symbol := quotation.Symbol + "/USD"
	price := quotation.Price
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp, chainId)
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
	timestamp int64,
	chainId int64) error {

	gasPrice, err := getGasSuggestion(chainId)
	if err != nil {
		log.Fatal(err)
	}

	// Get 110% of the gas price
	fmt.Println(gasPrice)
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)
	fmt.Println(gasPrice)

	// Write values to smart contract
	tx, err := contract.SetValue(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
		GasLimit: 1000725,
	}, key, big.NewInt(value), big.NewInt(timestamp))
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Printf("Nonce: %d\n", tx.Nonce())
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getAssetQuotationFromDia(blockchain, address string) (*models.Quotation, error) {
	response, err := http.Get("https://api.diadata.org/v1/assetQuotation/" + blockchain + "/" + address)
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

func getGasSuggestion(chainId int64) (*big.Int, error) {
	chainName := "shiden"
	if chainId == 592 {
		chainName = "astar"
	}
	if chainId == 81 {
		chainName = "shibuya"
	}
	response, err := http.Get("http://astargasstation.dia-services:3000/api/" + chainName + "/gasnow")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on astar gasstation with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	gasSuggestion := gjson.Get(string(contents), "data.fast")
	retval := big.NewInt(gasSuggestion.Int())

	return retval, nil
}
