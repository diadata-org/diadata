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

	"github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	sleepSeconds, err := strconv.Atoi(utils.Getenv("SLEEP_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse sleepSeconds: %v", err)
	}
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v", err)
	}

	addresses := []string{
		"0x3d9907F9a368ad0a51Be60f7Da3b97cf940982D8", //GRAIL
		"0x7Fc66500c84A76Ad7e9c93437bFc5Ac33E2DDaE9", //AAVE
		"0xba100000625a3754423978a60c9317c58a424e3D", //BAL
		"0xAf5191B0De278C7286d6C7CC6ab6BB8A73bA2Cd6", //STG
		"0x18c11FD286C5EC11c3b683Caa813B77f5163A122", //GNS
		"0x10393c20975cF177a3513071bC110f7962CD67da", //JONES
		"0x4e3FBD56CD56c3e72c1403e103b45Db9da5B9D2B", //CVX
		"0x090185f2135308BaD17527004364eBcC2D37e5F6", //SPELL
		"0xdeFA4e8a7bcBA345F687a2f1456F5Edd9CE97202", //KNC
		"0xc5102fE9359FD9a28f877a67E36B0F050d81a3CC", //HOP
		"0x6e84a6216eA6dACC71eE8E6b0a5B7322EEbC0fDd", //JOE
		"0x4945970EfeEc98D393b4b979b9bE265A3aE28A8B", //GMD
		"0x088cd8f5eF3652623c22D48b1605DCfE860Cd704", //VELA
		"0xfc5A1A6EB076a2C7aD06eD22C90d7E710E35ad0a", //GMX
		"0x1A5B0aaF478bf1FDA7b934c76E7692D722982a6D", //BFR
		"0x0C4681e6C0235179ec3D4F4fc4DF3d14FDD96017", //RDNT
		"0xC47D9753F3b32aA9548a7C3F30b6aEc3B2d2798C", //TND
		"0x51318B7D00db7ACc4026C88c3952B66278B6A67F", //PLS
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", //USDC
	}
	blockchains := []string{
		"Arbitrum",  //GRAIL
		"Ethereum",  //AAVE
		"Ethereum",  //ETH
		"Ethereum",  //STG
		"Arbitrum",  //GNS
		"Arbitrum",  //JONES
		"Ethereum",  //CVX
		"Ethereum",  //SPELL
		"Ethereum",  //KNC
		"Ethereum",  //HOP
		"Avalanche", //JOE
		"Arbitrum",  //GMD
		"Arbitrum",  //VELA
		"Arbitrum",  //GMX
    "Arbitrum",  //BFR
    "Arbitrum",  //RDNT
    "Arbitrum",  //TND
    "Arbitrum",  //PLS
    "Ethereum",  //USDC
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
					oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, conn, blockchain, address)
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

func periodicXChainOracleUpdateHelper(
	oldPrice float64,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleServiceV2.DIAOracleV2,
	conn *ethclient.Client,
	symbol string,
	blockchains []string,
	addresses []string,
) (newPrice float64, err error) {

	// Get quotations for tokens
	var (
		weightedPrice       float64
		totalVolume         float64
		rawQ                *models.Quotation
		aggregatedQuotation models.Quotation
	)
	for i := range addresses {
		rawQ, err = getAssetQuotationFromDia(blockchains[i], addresses[i])
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", addresses[i], err)
			return oldPrice, err
		}
		if *rawQ.VolumeYesterdayUSD > 0 {
			weightedPrice += rawQ.Price * (*rawQ.VolumeYesterdayUSD)
			totalVolume += *rawQ.VolumeYesterdayUSD
		}
	}
	aggregatedQuotation.Name = symbol
	aggregatedQuotation.Symbol = symbol

	// Check for deviation
	if totalVolume > 0 {
		newPrice = weightedPrice / totalVolume
	} else {
		newPrice = oldPrice
	}
	aggregatedQuotation.Price = newPrice

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		// Check for "too steep" update
		if oldPrice != 0 && (newPrice > (oldPrice*2)) || (newPrice < (oldPrice * 0.5)) {
			log.Println("New price %d out of bounds comparted to old price %d!", newPrice, oldPrice)
			return oldPrice, nil
		}
		log.Println("Entering deviation based update zone")
		err = updateQuotation(&aggregatedQuotation, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
}

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string) (float64, error) {

	// Get quotation for token and update Oracle
	rawQ, err := getAssetQuotationFromDia(blockchain, address)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
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
	fmt.Println(gasPrice)
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)
	fmt.Println(gasPrice)
	// Write values to smart contract
	tx, err := contract.SetValue(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 1000725,
		GasPrice: gasPrice,
	}, key, big.NewInt(value), big.NewInt(timestamp))
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getAssetQuotationFromDia(blockchain, address string) (*models.Quotation, error) {
	response, err := http.Get("https://rest.diadata.org/v1/assetQuotation/" + blockchain + "/" + address)
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
