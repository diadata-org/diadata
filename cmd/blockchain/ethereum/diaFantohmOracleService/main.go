package main

import (
	"context"
	"errors"
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
	gql "github.com/machinebox/graphql"
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

	addresses := []string{
		"0x6Fc9383486c163fA48becdEC79d6058f984f62cA",//USDB (Fantom)
		"0x6B175474E89094C44Da98b954EedeAC495271d0F",//DAI
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48",//USDC
		"0xdAC17F958D2ee523a2206206994597C13D831ec7",//USDT
		"0x0000000000000000000000000000000000000000",//FTM
		"0x98878B06940aE243284CA214f92Bb71a2b032B8A",//WMOVR
	}
	blockchains := []string{
		"Fantom", //USDB (Fantom)
		"Ethereum",//DAI
		"Ethereum",//USDC
		"Ethereum",//USDT
		"Fantom",//FTM
		"Moonriver",//WMOVR
	}
	vwapAddresses := []string{
		"0xfa1FBb8Ef55A4855E5688C0eE13aC3f202486286",//FHM (Fantom)
		"0xfa1FBb8Ef55A4855E5688C0eE13aC3f202486286",//FHM (Moonriver)
	}
	vwapBlockchains := []string{
		"Fantom",
		"Moonriver",
	}

	oldPrices := make(map[int]float64)
	vwapOldPrice := 0.0

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
				log.Println("old vwap price", vwapOldPrice)
				vwapOldPrice, err = periodicVwapOracleUpdateHelper(vwapOldPrice, deviationPermille, auth, contract, conn, vwapBlockchains, vwapAddresses)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(time.Duration(sleepSeconds) * time.Second)
			}
		}
	}()
	select {}
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

func periodicVwapOracleUpdateHelper(oldVwapPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchains []string, addresses []string) (float64, error) {

	var prices []float64
	var volumes []float64
	// Get Quotations/Prices for all assets
	for i, address := range addresses {
		rawQ, err := getAssetQuotationFromDia(blockchains[i], addresses[i])
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
			return oldVwapPrice, err
		}
		prices = append(prices, rawQ.Price)

		rawV, err := getAssetVolumeFromDia(blockchains[i], addresses[i], 120)
		if err != nil {
			log.Printf("Failed to retrieve %s volume data from DIA: %v", blockchains[i], err)
		}
		volumes = append(volumes, rawV)
	}
	vwapValue, err := calculateVwap(prices, volumes)
	if err != nil {
		log.Printf("Error in VWAP calculation: %v", err)
		return oldVwapPrice, err
	}

	newVwapPrice := vwapValue
	// Check if we got a new price at all
	if newVwapPrice == 0.0 {
		return oldVwapPrice, nil
	}

	// Check for deviation
	if (newVwapPrice > (oldVwapPrice * (1 + float64(deviationPermille)/1000))) || (newVwapPrice < (oldVwapPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone for vwap")
		err = updateVwap(newVwapPrice, "FHM-VWAP", auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldVwapPrice, err
		}
		return newVwapPrice, nil
	}
	return oldVwapPrice, nil
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

func updateVwap(value float64, name string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, name, int64(value*100000000), timestamp)
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
	log.Printf("nonce: %d\n", tx.Nonce())
	log.Printf("gas price: %d\n", tx.GasPrice())
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

func getAssetVolumeFromDia(blockchain, address string, blockDuration int) (float64, error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-blockDuration * 2) * time.Second)
	type Response struct {
		GetChart []struct {
			Name   string    `json:"Name"`
			Symbol string    `json:"Symbol"`
			Time   time.Time `json:"Time"`
			Value  float64   `json:"Value"`
		} `json:"GetChart"`
	}
	client := gql.NewClient("https://api.diadata.org/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetChart(filter: "vol", 
			Symbol:"Asset",
			BlockDurationSeconds: ` + strconv.Itoa(blockDuration) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(blockDuration) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			Address: "` + address + `", 
			BlockChain: "` + blockchain + `") {
				Name
				Symbol
				Time
				Value
	  	}
		}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, err
	}
	if len(r.GetChart) == 0 {
		return 0.0, errors.New("no results")
	}
	return r.GetChart[len(r.GetChart)-1].Value, nil
}

func calculateVwap(prices, volumes []float64) (float64, error) {
	if len(prices) != len(volumes) {
		return 0.0, errors.New("VWAP needs to be calculated on same number of volumes and prices.")
	}

	accumulator := 0.0
	volumeSum := 0.0
	for i, price := range prices {
		accumulator += price * volumes[i]
		volumeSum += volumes[i]
	}
	return accumulator / volumeSum, nil
}
