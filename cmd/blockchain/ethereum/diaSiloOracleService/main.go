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
	coingeckoApiKey := utils.Getenv("COINGECKO_API_KEY", "")
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
	xcDeviationPermille, err := strconv.Atoi(utils.Getenv("XC_DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse xcDeviationPermille: %v")
	}
	gqlWindowSize, err := strconv.Atoi(utils.Getenv("GQL_WINDOW_SIZE", "120"))
	if err != nil {
		log.Fatalf("Failed to parse gqlWindowSize: %v")
	}
	gqlMethodology := utils.Getenv("GQL_METHODOLOGY", "vwap")
	assetsStr := utils.Getenv("ASSETS", "")
	gqlAssetsStr := utils.Getenv("GQL_ASSETS", "")
	cgAssetsStr := utils.Getenv("CG_ASSETS", "")
	xcAssetsStr := utils.Getenv("XC_ASSETS", "")
	xcAssetSymbol := utils.Getenv("XC_ASSET_SYMBOL", "SILO")
	xcCoingeckoName := utils.Getenv("XC_COINGECKO_NAME", "silo-finance")

	addresses := []string{}
	blockchains := []string{}
	cgNames := []string{}
	useGql := false
	xcBlockchains := []string{}
	xcAddresses := []string{}
	var assetsToParse string
	var xcAssetsToParse string

	// Either Assets or GQL Assets must be non-empty
	if gqlAssetsStr != "" && assetsStr == "" {
		useGql = true
		assetsToParse = gqlAssetsStr
	} else if gqlAssetsStr == "" && assetsStr != "" {
		useGql = false
		assetsToParse = assetsStr
	} else {
		log.Fatalf("Use either ASSETS or GQL_ASSETS env variable")
	}
	xcAssetsToParse = xcAssetsStr
	assetsParsed := strings.Split(assetsToParse, ",")
	xcAssetsParsed := strings.Split(xcAssetsToParse, ",")

	cgAssetsParsed := strings.Split(cgAssetsStr, ",")

	for _, asset := range assetsParsed {
		entries := strings.Split(asset, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		addresses = append(addresses, strings.TrimSpace(entries[1]))
	}

	for _, asset := range xcAssetsParsed {
		entries := strings.Split(asset, "-")
		xcBlockchains = append(xcBlockchains, strings.TrimSpace(entries[0]))
		xcAddresses = append(xcAddresses, strings.TrimSpace(entries[1]))
	}

	for _, cgName := range cgAssetsParsed {
		cgNames = append(cgNames, strings.TrimSpace(cgName))
	}

	oldPrices := make(map[int]float64)
	oldXcPrice := 0.0

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
					coinGeckoName := cgNames[i]
					log.Println("old price", oldPrice)
					oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, conn, blockchain, address, useGql, gqlWindowSize, gqlMethodology, coinGeckoName, coingeckoApiKey)
					oldPrices[i] = oldPrice
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
				// xc updater
				log.Println("old xc price:", oldXcPrice)
				oldXcPrice, err = periodicXcOracleUpdateHelper(oldXcPrice, xcDeviationPermille, xcAssetSymbol, auth, contract, conn, xcBlockchains, xcAddresses, 3600, 120, "mair", xcCoingeckoName, coingeckoApiKey)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	if mandatoryFrequencySeconds > 0 {
		mandatoryticker := time.NewTicker(time.Duration(mandatoryFrequencySeconds) * time.Second)
		go func() {
			for {
				select {
				case <-mandatoryticker.C:
					for i, address := range addresses {
						blockchain := blockchains[i]
						oldPrice := oldPrices[i]
						coinGeckoName := cgNames[i]
						log.Println("old price", oldPrice)
						oldPrice, err = oracleUpdateHelper(oldPrice, auth, contract, conn, blockchain, address, useGql, gqlWindowSize, gqlMethodology, coinGeckoName, coingeckoApiKey)
						oldPrices[i] = oldPrice
						if err != nil {
							log.Println(err)
						}
						time.Sleep(time.Duration(sleepSeconds) * time.Second)
					}
					// xc updater
					log.Println("old xc price:", oldXcPrice)
					oldXcPrice, err = xcOracleUpdateHelper(oldXcPrice, xcAssetSymbol, auth, contract, conn, xcBlockchains, xcAddresses, 3600, 120, "mair", xcCoingeckoName, coingeckoApiKey)
					if err != nil {
						log.Println(err)
					}
				}
			}
		}()
	}

	select {}
}

func oracleUpdateHelper(oldPrice float64, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string, useGql bool, gqlWindowSize int, gqlMethodology string, coingeckoName, coingeckoApiKey string) (float64, error) {
	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	if useGql {
		price, symbol, err := getGraphqlAssetQuotationFromDia(blockchain, address, gqlWindowSize, gqlMethodology)
		if err != nil {
			log.Printf("Failed to retrieve %s quotation data from Graphql on DIA: %v", address, err)
			return oldPrice, err
		}
		rawQ.Symbol = symbol
		rawQ.Price = price
	} else {
		// Special case: RDPX
		if address == "0x32Eb7902D4134bf98A28b963D26de779AF92A212" && blockchain == "Arbitrum" {
			price, symbol, err := getRdpxGraphqlAssetQuotationFromDia(blockchain, address)
			if err != nil {
				log.Printf("Failed to retrieve %s (RDPX) quotation data from Graphql on DIA: %v", address, err)
				return oldPrice, err
			}
			rawQ.Symbol = symbol
			rawQ.Price = price
		} else {
			rawQ, err = getAssetQuotationFromDia(blockchain, address)
			if err != nil {
				log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
				return oldPrice, err
			}
		}
	}
	rawQ.Name = rawQ.Symbol

	newPrice := rawQ.Price

	// check coingecko before sending out an update transaction
	allowedCoingeckoDeviation := 0.075
	// Exception for RDPX: CG data is not super reliable, high deviations expected
	if address == "0x32Eb7902D4134bf98A28b963D26de779AF92A212" && blockchain == "Arbitrum" {
		allowedCoingeckoDeviation = 0.2
	}
	cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
	if err != nil {
		return oldPrice, err
	}
	if cgPrice == 0.0 {
		log.Printf("Error! Coingecko API returned price 0.0.")
		return oldPrice, nil
	}
	log.Printf("Deviation from coingecko: %f\n", math.Abs(cgPrice - rawQ.Price) / cgPrice)
	if (math.Abs(cgPrice - rawQ.Price) / cgPrice) > allowedCoingeckoDeviation {
		// Error case, stop transaction from happening
		log.Printf("Error! Price %f for asset %s-%s out of coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)
		return oldPrice, nil
	}
	log.Printf("Price %f for asset %s-%s in coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)

	err = updateQuotation(rawQ, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return oldPrice, err
	}
	return newPrice, nil

}

func xcOracleUpdateHelper(oldPrice float64, xcAssetSymbol string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchains []string, addresses []string, gqlWindowSize int, gqlWindowShift int, gqlMethodology string, coingeckoName, coingeckoApiKey string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	price, err := getGraphqlXcAssetQuotationFromDia(blockchains, addresses, gqlWindowSize, gqlWindowShift, gqlMethodology)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from xcGraphql on DIA: %v", addresses, err)
		return oldPrice, err
	}
	rawQ.Symbol = xcAssetSymbol
	rawQ.Price = price
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	// check coingecko before sending out an update transaction
	allowedCoingeckoDeviation := 0.075
	cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
	if err != nil {
		return oldPrice, err
	}
	if cgPrice == 0.0 {
		log.Printf("Error! Coingecko API returned price 0.0.")
		return oldPrice, nil
	}
	log.Printf("Deviation from coingecko: %f\n", math.Abs(cgPrice - rawQ.Price) / cgPrice)
	if (math.Abs(cgPrice - rawQ.Price) / cgPrice) > allowedCoingeckoDeviation {
		// Error case, stop transaction from happening
		log.Printf("Error! Price %f for asset %s out of coingecko range %f.", rawQ.Price, xcAssetSymbol, cgPrice)
		return oldPrice, nil
	}
	log.Printf("Price %f for asset %s in coingecko range %f.", rawQ.Price, xcAssetSymbol, cgPrice)

	err = updateQuotation(rawQ, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return oldPrice, err
	}
	return newPrice, nil
}

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string, useGql bool, gqlWindowSize int, gqlMethodology string, coingeckoName, coingeckoApiKey string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	if useGql {
		price, symbol, err := getGraphqlAssetQuotationFromDia(blockchain, address, gqlWindowSize, gqlMethodology)
		if err != nil {
			log.Printf("Failed to retrieve %s quotation data from Graphql on DIA: %v", address, err)
			return oldPrice, err
		}
		rawQ.Symbol = symbol
		rawQ.Price = price
	} else {
		// Special case: RDPX
		if address == "0x32Eb7902D4134bf98A28b963D26de779AF92A212" && blockchain == "Arbitrum" {
			price, symbol, err := getRdpxGraphqlAssetQuotationFromDia(blockchain, address)
			if err != nil {
				log.Printf("Failed to retrieve %s (RDPX) quotation data from Graphql on DIA: %v", address, err)
				return oldPrice, err
			}
			rawQ.Symbol = symbol
			rawQ.Price = price
		} else {
			rawQ, err = getAssetQuotationFromDia(blockchain, address)
			if err != nil {
				log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
				return oldPrice, err
			}
		}
	}
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	// Check if deviation should be 0.5% for ETH
	if blockchain == "Ethereum" && address == "0x0000000000000000000000000000000000000000" {
		deviationPermille = 5
	}

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		
		// check coingecko before sending out an update transaction
		allowedCoingeckoDeviation := 0.075
		// Exception for RDPX: CG data is not super reliable, high deviations expected
		if address == "0x32Eb7902D4134bf98A28b963D26de779AF92A212" && blockchain == "Arbitrum" {
			allowedCoingeckoDeviation = 0.2
		}
		cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
		if err != nil {
			return oldPrice, err
		}
		if cgPrice == 0.0 {
			log.Printf("Error! Coingecko API returned price 0.0.")
			return oldPrice, nil
		}
		log.Printf("Deviation from coingecko: %f\n", math.Abs(cgPrice - rawQ.Price) / cgPrice)
		if (math.Abs(cgPrice - rawQ.Price) / cgPrice) > allowedCoingeckoDeviation {
			// Error case, stop transaction from happening
			log.Printf("Error! Price %f for asset %s-%s out of coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)
			return oldPrice, nil
		}
		log.Printf("Price %f for asset %s-%s in coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)
		
		err = updateQuotation(rawQ, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
}

func periodicXcOracleUpdateHelper(oldPrice float64, deviationPermille int, xcAssetSymbol string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchains []string, addresses []string, gqlWindowSize int, gqlWindowShift int, gqlMethodology string, coingeckoName, coingeckoApiKey string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error

	// Get quotation for token and update Oracle
	price, err := getGraphqlXcAssetQuotationFromDia(blockchains, addresses, gqlWindowSize, gqlWindowShift, gqlMethodology)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from xcGraphql on DIA: %v", addresses, err)
		return oldPrice, err
	}
	rawQ.Symbol = xcAssetSymbol
	rawQ.Price = price
	rawQ.Name = rawQ.Symbol

	// Check for deviation
	newPrice := rawQ.Price

	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")

		// check coingecko before sending out an update transaction
		allowedCoingeckoDeviation := 0.075
		cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
		if err != nil {
			return oldPrice, err
		}
		if cgPrice == 0.0 {
			log.Printf("Error! Coingecko API returned price 0.0.")
			return oldPrice, nil
		}
		log.Printf("Deviation from coingecko: %f\n", math.Abs(cgPrice - rawQ.Price) / cgPrice)
		if (math.Abs(cgPrice - rawQ.Price) / cgPrice) > allowedCoingeckoDeviation {
			// Error case, stop transaction from happening
			log.Printf("Error! Price %f for asset %s out of coingecko range %f.", rawQ.Price, xcAssetSymbol, cgPrice)
			return oldPrice, nil
		}
		log.Printf("Price %f for asset %s in coingecko range %f.", rawQ.Price, xcAssetSymbol, cgPrice)

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

// Special case for RDPX: Only query with liquidity >500k
func getRdpxGraphqlAssetQuotationFromDia(blockchain, address string) (float64, string, error) {
	log.Println("Entering Rdpx special case: Get price with minimum liquidity of 500k")
	windowSize := 120
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize*2) * time.Second)
	type Response struct {
		GetFeed []struct {
			Name   string    `json:"Name"`
			Time   time.Time `json:"Time"`
			Value  float64   `json:"Value"`
		} `json:"GetFeed"`
	}
	client := gql.NewClient("https://api.diadata.org/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetFeed(
		 	Filter: "vwapir", 
			BlockSizeSeconds: ` + strconv.Itoa(windowSize) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(windowSize) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			FeedSelection: [
				{
					Address: "` + address + `", 
					Blockchain: "` + blockchain + `",
					LiquidityThreshold: 500000.0
				}
			]
		)
		{
			Name
			Time
			Value
  	}
	}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, "", err
	}
	if len(r.GetFeed) == 0 {
		return 0.0, "", errors.New("no results")
	}
	return r.GetFeed[len(r.GetFeed)-1].Value, "RDPX", nil
}

func getGraphqlAssetQuotationFromDia(blockchain, address string, windowSize int, gqlMethodology string) (float64, string, error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize*2) * time.Second)
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
		 	filter: "` + gqlMethodology + `", 
			Symbol:"Asset",
			BlockDurationSeconds: ` + strconv.Itoa(windowSize) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(windowSize) + `,
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

func getGraphqlXcAssetQuotationFromDia(blockchains, addresses []string, windowSize int, windowShift int, gqlMethodology string) (float64, error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize*2) * time.Second)
	quoteAssetsArrayStr := ""
	for i, address := range addresses {
		quoteAssetsArrayStr += "{Address:\"" + address + "\",BlockChain:\"" + blockchains[i] + "\"}"
		if i < len(addresses) - 1 {
			quoteAssetsArrayStr += ","
		}
		quoteAssetsArrayStr += "\n"
	}
	type Response struct {
		GetxcFeed []struct {
			Name   string    `json:"Name"`
			Time   time.Time `json:"Time"`
			Value  float64   `json:"Value"`
		} `json:"GetxcFeed"`
	}
	client := gql.NewClient("https://api.diadata.org/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetxcFeed(
		 	filter: "` + gqlMethodology + `", 
			BlockSizeSeconds: ` + strconv.Itoa(windowSize) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(windowShift) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			Exchanges: [],
			QuoteAssets: [` + quoteAssetsArrayStr +
			`])
			{
				Name
				Time
				Value
	  	}
		}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, err
	}
	if len(r.GetxcFeed) == 0 {
		return 0.0, errors.New("no results")
	}
	return r.GetxcFeed[len(r.GetxcFeed)-1].Value, nil
}

func getCoingeckoPrice(assetName, coingeckoApiKey string) (float64, error) {
	url := "https://pro-api.coingecko.com/api/v3/simple/price?ids=" + assetName + "&vs_currencies=usd&x_cg_pro_api_key=" + coingeckoApiKey
	response, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on coingecko API call with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	price := gjson.Get(string(contents), assetName + ".usd").Float()
	return price, nil
}
