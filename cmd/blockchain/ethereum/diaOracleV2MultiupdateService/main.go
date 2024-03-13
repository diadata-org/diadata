package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gql "github.com/machinebox/graphql"
	"github.com/tidwall/gjson"
)

type Asset struct {
	blockchain string
	address    string
	symbol     string
	gqlParams  GqlParameters
}

// Update asset1 only if asset0 is updated
type ConditionalPair struct {
	asset0 int
	asset1 int
}

type GqlParameters struct {
	FeedSelection []struct {
		Address            string  `json:"Address"`
		Blockchain         string  `json:"Blockchain"`
		LiquidityThreshold float64 `json:"LiquidityThreshold"`
		Exchangepairs      []struct {
			Exchange string   `json:"Exchange"`
			Pairs    []string `json:"Pairs"`
		} `json:"Exchangepairs"`
	} `json:"FeedSelection"`
}

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	backupNode := utils.Getenv("BACKUP_NODE", "")
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
	gqlWindowSize, err := strconv.Atoi(utils.Getenv("GQL_WINDOW_SIZE", "120"))
	if err != nil {
		log.Fatalf("Failed to parse gqlWindowSize: %v")
	}
	conditionalAssets := utils.Getenv("CONDITIONAL_ASSETS", "")
	gqlMethodology := utils.Getenv("GQL_METHODOLOGY", "vwap")
	assetsStr := utils.Getenv("ASSETS", "")
	gqlAssetsStr := utils.Getenv("GQL_ASSETS", "")

	assets := []Asset{}
	conditionalPairs := []ConditionalPair{}
	useGql := false
	var assetsToParse string

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
	assetsParsed := strings.Split(assetsToParse, ";")

	for _, asset := range assetsParsed {
		var currAsset Asset

		// parse asset from env
		entries := strings.Split(asset, "-")
		//TODO: check if len(array) > 0
		currAsset.blockchain = strings.TrimSpace(entries[0])
		currAsset.address = strings.TrimSpace(entries[1])
		currAsset.symbol = strings.TrimSpace(entries[2])

		// Find out is there are additional GQL parameters for this asset
		if len(entries) > 3 {
			// Join the rest of the line together, because the previous split might have affected substrings of the parameters
			gqlFeedSelectionQuery := strings.Join(entries[3:], "-")
			var currGqlParams GqlParameters
			if useGql && gqlFeedSelectionQuery != "" {
				err := json.Unmarshal([]byte(gqlFeedSelectionQuery), &currGqlParams)
				if err != nil {
					log.Println("Error while parsing GQL asset string:", err)
				}
			}
			currAsset.gqlParams = currGqlParams
		}

		assets = append(assets, currAsset)
	}

	// Get conditional pairs for assets where an asset should only be updated if an update for the other asset is available
	if conditionalAssets != "" {
		conditionalPairsToParse := strings.Split(conditionalAssets, ";")
		for _, conditionalPair := range conditionalPairsToParse {
			var currPair ConditionalPair
			var err error

			entries := strings.Split(conditionalPair, "-")
			currPair.asset0, err = strconv.Atoi(entries[0])
			if err != nil {
				log.Fatalf("Failed to parse conditional assets: %v", err)
			}
			currPair.asset1, err = strconv.Atoi(entries[1])
			if err != nil {
				log.Fatalf("Failed to parse conditional assets: %v", err)
			}

			conditionalPairs = append(conditionalPairs, currPair)
		}
	}

	publishedPrices := make(map[string]float64)

	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	connBackup, err := ethclient.Dial(backupNode)
	if err != nil {
		log.Fatalf("Failed to connect to the backup Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract, contractBackup *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService
	err = deployOrBindContract(deployedContract, conn, connBackup, auth, &contract, &contractBackup)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}

	/*
	 * Update Oracle periodically
	 */
	var mandatoryticker *time.Ticker
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	if mandatoryFrequencySeconds > 0 {
		mandatoryticker = time.NewTicker(time.Duration(mandatoryFrequencySeconds) * time.Second)
	}
	go func() {
		for {
			if mandatoryFrequencySeconds > 0 {
				select {
				case <-ticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range assets {
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					log.Println(newAssetPrices)
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets, conditionalPairs)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId, assets, conditionalPairs)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				case <-mandatoryticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
				OUTER:
					for i, asset := range assets {
						// Check if we need to skip any assets due to being in a conditional pair
						for _, conditionalPair := range conditionalPairs {
							if i == conditionalPair.asset0 || i == conditionalPair.asset1 {
								continue OUTER
							}
						}
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices, regardless of deviation
					emptyMap := make(map[string]float64)
					publishedPrices, err = oracleUpdateExecutor(emptyMap, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets, conditionalPairs)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(emptyMap, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId, assets, conditionalPairs)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				}
			} else {
				select {
				case <-ticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range assets {
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets, conditionalPairs)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId, assets, conditionalPairs)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				}
			}
		}
	}()

	select {}
}

func oracleUpdateExecutor(
	publishedPrices map[string]float64,
	newPrices map[string]float64,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	chainId int64,
	assets []Asset,
	conditionalAssets []ConditionalPair) (map[string]float64, error) {
	// Check for deviation and collect all new prices in a map
	// If a published price is 0, update in any case
	updateCollector := make(map[string]float64)
	priceCollector := make(map[string]float64)
	for i, asset := range assets {
		updateAssetConditional := false

		// check if this asset is conditional
		for j := range conditionalAssets {
			// find out if the asset which decides on the condition needs to be updated
			if conditionalAssets[j].asset1 == i {
				// Compare asset0 if it will receive an update
				asset0 := assets[conditionalAssets[j].asset0]
				asset0NewPrice := newPrices[asset0.symbol]
				asset0OldPrice := publishedPrices[asset0.symbol]

				// Flag asset for update if it is conditional
				if asset0NewPrice > 1e-8 && ((asset0NewPrice > (asset0OldPrice * (1 + float64(deviationPermille)/1000))) || (asset0NewPrice < (asset0OldPrice * (1 - float64(deviationPermille)/1000)))) {
					updateAssetConditional = true
					log.Printf("Asset %s flagged for update because conditional asset %s is updated as well.", asset.symbol, asset0.symbol)
				} else {
					updateAssetConditional = false
					log.Printf("Asset %s is not updated because the conditional asset %s is not deviating.", asset.symbol, asset0.symbol)
				}
			}
		}

		newPrice := newPrices[asset.symbol]
		fmt.Println("new price", newPrice)
		oldPrice := publishedPrices[asset.symbol]

		if updateAssetConditional || (newPrice > 1e-8 && ((newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))))) {
			log.Printf("Entering deviation based update zone for old price %.2f of asset %s. New price: %.2f", oldPrice, asset.symbol, newPrice)
			updateCollector[asset.symbol] = newPrice
			priceCollector[asset.symbol] = newPrice
		} else {
			priceCollector[asset.symbol] = oldPrice
		}
	}

	// Check if any update is to be executed
	if len(updateCollector) == 0 {
		log.Println("No update necessary.")
		return priceCollector, nil
	}
	// Serialize keys and values into array for oracle update transaction
	var keys []string
	var prices []int64
	for key, price := range updateCollector {
		key = key + "/USD"
		keys = append(keys, key)
		integerPrice := int64(price * 100000000)
		prices = append(prices, integerPrice)
	}

	// Update prices in one transaction
	timestamp := time.Now().Unix()
	err := updateOracleMultiValues(conn, contract, auth, chainId, keys, prices, timestamp)
	if err != nil {
		log.Printf("Failed to update Oracle: %v", err)
		return nil, err
	}

	return priceCollector, nil
}

func retrieveAssetPrice(asset Asset, useGql bool, gqlWindowSize int, gqlMethodology string, gqlLiquidityParameters GqlParameters) (float64, error) {
	var err error
	var price float64

	// Get quotation for token and update Oracle
	if useGql {
		price, err = getGraphqlAssetQuotationFromDia(asset.blockchain, asset.address, gqlWindowSize, gqlMethodology, gqlLiquidityParameters)
		if err != nil {
			log.Printf("Failed to retrieve %s (%s) quotation data from Graphql on DIA: %v", asset.address, asset.symbol, err)
			return 0.0, err
		}
	} else {
		price, err = getAssetQuotationFromDia(asset.blockchain, asset.address)
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", asset.address, err)
			return 0.0, err
		}
	}
	return price, nil
}

func deployOrBindContract(
	deployedContract string,
	conn *ethclient.Client,
	connBackup *ethclient.Client,
	auth *bind.TransactOpts,
	contract **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	contractBackup **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService) error {
	var err error
	if deployedContract != "" {
		// bind primary and backup
		*contract, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
		*contractBackup, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), connBackup)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaOracleV2MultiupdateService.DeployDiaOracleV2MultiupdateService(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		// bind backup
		*contractBackup, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(addr, connBackup)
		if err != nil {
			return err
		}
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

func updateOracleMultiValues(
	client *ethclient.Client,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	auth *bind.TransactOpts,
	chainId int64,
	keys []string,
	values []int64,
	timestamp int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int
	var err error

	// Get proper gas price depending on chainId
	switch chainId {
	/*case 288: //Boba
	gasPrice = big.NewInt(1000000000)*/
	case 592: //Astar
		response, err := http.Get("https://gas.astar.network/api/gasnow?network=astar")
		if err != nil {
			return err
		}

		defer response.Body.Close()
		if 200 != response.StatusCode {
			return err
		}
		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return err
		}

		gasSuggestion := gjson.Get(string(contents), "data.fast")
		gasPrice = big.NewInt(gasSuggestion.Int())
	default:
		// Get gas price suggestion
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Print(err)
			return err
		}

		// Get 110% of the gas price
		fGas := new(big.Float).SetInt(gasPrice)
		fGas.Mul(fGas, big.NewFloat(1.1))
		gasPrice, _ = fGas.Int(nil)
	}

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := big.NewInt(value)
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	// check if tx is sendable then fgo backup
	if err != nil {
		// backup in here
		return err
	}

	log.Printf("Gas price: %d\n", tx.GasPrice())
	log.Printf("Data: %x\n", tx.Data())
	log.Printf("Nonce: %d\n", tx.Nonce())
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getAssetQuotationFromDia(blockchain, address string) (float64, error) {
	// ibc special case: convert / to - in the query string
	if strings.Split(address, "/")[0] == "ibc" {
		address = strings.Split(address, "/")[0] + "-" + strings.Split(address, "/")[1]
	}

	// Execute the query
	response, err := http.Get("https://api.diadata.org/v1/assetQuotation/" + blockchain + "/" + address)
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
	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return 0.0, err
	}
	return quotation.Price, nil
}

func getGraphqlAssetQuotationFromDia(blockchain, address string, windowSize int, gqlMethodology string, gqlParameters GqlParameters) (float64, error) {
	// Decide whether Feedselection or simple Address/blockchain logic is used
	feedSelectionQuery := "FeedSelection: ["
	if len(gqlParameters.FeedSelection) > 0 {
		// Loop through all selected feeds (e.g. for crosschain feeds)
		for _, selectedFeed := range gqlParameters.FeedSelection {
			// generate strings for optional parameters for liquidity threshold/pool selection
			var lqThresholdString string
			if selectedFeed.LiquidityThreshold > 0 {
				lqThresholdString = "LiquidityThreshold:" + fmt.Sprintf("%.2f", selectedFeed.LiquidityThreshold) + ","
			} else {
				lqThresholdString = ""
			}
			var exchangePairsString string
			if len(selectedFeed.Exchangepairs) > 0 {
				exchangePairsString = "Exchangepairs:[\n"
				for _, exchangePair := range selectedFeed.Exchangepairs {
					exchangePairsString += `{
					Exchange: "` + exchangePair.Exchange + `",`
					if len(exchangePair.Pairs) > 0 {
						exchangePairsString += `Pairs: [`
						for _, pair := range exchangePair.Pairs {
							exchangePairsString += `"` + pair + `",`
						}
						exchangePairsString += `]`
					}
					exchangePairsString += `},`
				}
				exchangePairsString += "]"
			} else {
				exchangePairsString = ""
			}
			feedSelectionQuery += `{
				Address:"` + selectedFeed.Address + `",
				Blockchain:"` + selectedFeed.Blockchain + `",` +
				lqThresholdString +
				exchangePairsString +
				`},`
		}
	} else {
		feedSelectionQuery += `{
				Address: "` + address + `",
				Blockchain: "` + blockchain + `",
			}`
	}
	feedSelectionQuery += "]"

	// Get times for start/end
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize*2) * time.Second)

	type Response struct {
		GetFeed []struct {
			Name  string    `json:"Name"`
			Time  time.Time `json:"Time"`
			Value float64   `json:"Value"`
			Pools string    `json:"Pools"`
			Pairs string    `json:"Pairs"`
		} `json:"GetFeed"`
	}

	client := gql.NewClient("https://api.diadata.org/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetFeed(
		 	Filter: "` + gqlMethodology + `",
			BlockSizeSeconds: ` + strconv.Itoa(windowSize) + `,
			BlockShiftSeconds: ` + strconv.Itoa(windowSize) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `,
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `,` +
		feedSelectionQuery +
		`) {
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
	if len(r.GetFeed) == 0 {
		return 0.0, errors.New("no results")
	}
	return r.GetFeed[len(r.GetFeed)-1].Value, nil
}
