package main

import (
	"context"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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
	gql "github.com/machinebox/graphql"
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
	collateralDeviationPermille, err := strconv.Atoi(utils.Getenv("COLLATERAL_DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse collateralDeviationPermille: %v")
	}
	assetsStr := utils.Getenv("ASSETS", "")
	synthAssetsStr := utils.Getenv("SYNTHASSETS", "")

	// Parse Assets
	assetsParsed := strings.Split(assetsStr, ",")
	addresses := []string{}
	blockchains := []string{}

	for _, asset := range assetsParsed {
		entries := strings.Split(asset, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		addresses = append(addresses, strings.TrimSpace(entries[1]))
	}

	// Parse Synthassets
	synthAssetsParsed := strings.Split(synthAssetsStr, ",")

	synthAddresses := []string{}
	synthProtocols := []string{}
	synthBlockchains := []string{}
	for _, asset := range synthAssetsParsed {
		entries := strings.Split(asset, "/")
		synthBlockchains = append(synthBlockchains, strings.TrimSpace(entries[0]))
		synthProtocols = append(synthProtocols, strings.TrimSpace(entries[1]))
		synthAddresses = append(synthAddresses, strings.TrimSpace(entries[2]))
	}

	oldPrices := make(map[int]float64)
	oldCollaterals := make(map[int]float64)

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
					synthAddress := synthAddresses[i]
					synthProtocol := synthProtocols[i]
					synthBlockchain := synthBlockchains[i]
					oldPrice := oldPrices[i]
					oldCollateral := oldCollaterals[i]
					log.Println("old price", oldPrice)
					log.Println("old ratio", oldCollateral)
					oldPrice, oldCollateral, err = periodicOracleUpdateHelper(oldPrice, oldCollateral, deviationPermille, collateralDeviationPermille, auth, contract, conn, blockchain, address, synthBlockchain, synthProtocol, synthAddress, sleepSeconds)
					oldPrices[i] = oldPrice
					oldCollaterals[i] = oldCollateral
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

func periodicOracleUpdateHelper(oldPrice float64, oldCollateralRatio float64, deviationPermille int, collateralDeviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string, synthBlockchain string, synthProtocol string, synthAddress string, sleepSeconds int) (float64, float64, error) {

	// Get quotation
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	windowSize := 120
	
	// sAVAX: Get 300s window from GQL
	if address == "0x2b2C81e08f1Af8835a78Bb2A90AE924ACE0eA4bE" && blockchain == "Avalanche" {
		windowSize = 300
	}
	
	price, symbol, err := getGraphqlAssetQuotationFromDia(blockchain, address, windowSize)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from Graphql on DIA: %v", address, err)
		return oldPrice, oldCollateralRatio, err
	}
	rawQ.Symbol = symbol
	rawQ.Price = price
	rawQ.Name = rawQ.Symbol

	// Get collateral ratio
	collateralRatio, err := getSynthCollateralFromDia(synthBlockchain, synthProtocol, synthAddress)
	if err != nil {
		log.Printf("Failed to retrieve %s collateralRatio from DIA: %v", address, err)
		return oldPrice, oldCollateralRatio, err
	}

	// Check for deviation
	toUpdate := false
	newPrice := rawQ.Price

	if math.Abs(newPrice - oldPrice) > oldPrice * float64(deviationPermille) / 1000 {
		log.Println("Entering price deviation based update zone")
		toUpdate = true
	} else if math.Abs(oldCollateralRatio - collateralRatio) > oldCollateralRatio * float64(collateralDeviationPermille) / 1000 {
		log.Println("Entering ratio deviation based update zone")
		toUpdate = true
	}

	if toUpdate {
		err = updateQuotation(rawQ, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, oldCollateralRatio, err
		}

		time.Sleep(time.Duration(sleepSeconds) * time.Second)
		err = updateCollateralRatio(collateralRatio, rawQ.Symbol, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, oldCollateralRatio, err
		}
		return newPrice, collateralRatio, nil
	}

	return oldPrice, oldCollateralRatio, nil
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

func updateCollateralRatio(collateralRatio float64, assetSymbol string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	key := assetSymbol + "-CR"
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, key, int64(collateralRatio * 100000000), timestamp)
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

func getSynthCollateralFromDia(blockchain, protocol, address string) (float64, error) {
	response, err := http.Get("https://api.diadata.org/v1/synthasset/" + blockchain + "/" + protocol + "?address=" + address)
	if err != nil {
		return 0.0, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}

	collateralRatio := gjson.Get(string(contents), "0.CollateralRatio")
	return collateralRatio.Float(), nil
}

func getGraphqlAssetQuotationFromDia(blockchain, address string, blockDuration int) (float64, string, error) {
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
		 GetChart(
		 	filter: "vwap", 
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
		return 0.0, "", err
	}
	if len(r.GetChart) == 0 {
		return 0.0, "", errors.New("no results")
	}
	return r.GetChart[len(r.GetChart)-1].Value, r.GetChart[len(r.GetChart)-1].Symbol, nil
}
