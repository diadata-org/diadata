package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"sync"
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
	blockchain               string
	address                  string
	symbol                   string
	coingeckoName            string
	cmcName                  string
	allowedGuardianDeviation float64
	deviation                int64
	gqlParams                GqlParameters
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
	// global window override (seconds)
	GqlWindow int `json:"GqlWindow"`
}

var diaBaseUrl string

func main() {
	var oracleUpdateMutex sync.Mutex

	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	backupNode := utils.Getenv("BACKUP_NODE", "")
	dataNodesStr := utils.Getenv("DATA_NODES", "")
	usdpDataAddressesStr := utils.Getenv("USDP_DATA_ADDRESSES", "")
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}
	mandatoryFrequencySeconds, err := strconv.Atoi(utils.Getenv("MANDATORY_FREQUENCY_SECONDS", "0"))
	if err != nil {
		log.Fatalf("Failed to parse mandatoryFrequencySeconds: %v", err)
	}
	mutexSeconds, err := strconv.Atoi(utils.Getenv("MUTEX_SECONDS", "10"))
	if err != nil {
		log.Fatalf("Failed to parse mutexSeconds: %v", err)
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v", err)
	}
	gqlWindowSize, err := strconv.Atoi(utils.Getenv("GQL_WINDOW_SIZE", "120"))
	if err != nil {
		log.Fatalf("Failed to parse gqlWindowSize: %v", err)
	}
	oracleDecimals, err := strconv.Atoi(utils.Getenv("ORACLE_DECIMALS", "8"))
	if err != nil {
		log.Fatalf("Failed to parse oracleDecimals: %v", err)
	}
	conditionalAssets := utils.Getenv("CONDITIONAL_ASSETS", "")
	gqlMethodology := utils.Getenv("GQL_METHODOLOGY", "vwapir")
	coingeckoApiKey := utils.Getenv("COINGECKO_API_KEY", "")
	cmcApiKey := utils.Getenv("CMC_API_KEY", "")
	assetsStr := utils.Getenv("ASSETS", "")
	gqlAssetsStr := utils.Getenv("GQL_ASSETS", "")
	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")
	compatibilityMode, err := strconv.ParseBool(utils.Getenv("COMPATIBILITY_MODE", "false"))
	if err != nil {
		log.Fatalf("Failed to parse compatibilityMode: %v", err)
	}
	deviationPerAssetMode, err := strconv.ParseBool(utils.Getenv("DEVIATION_PER_ASSET_MODE", "false"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPerAssetMode: %v", err)
	}
	gasMultiplier, err := strconv.ParseFloat(utils.Getenv("GAS_MULTIPLIER", "1.1"), 64)
	if err != nil {
		log.Fatalf("Failed to parse gasMultiplier: %v", err)
	}

	dataNodes := []string{}
	dataNodes = strings.Split(dataNodesStr, ";")

	usdpDataAddresses := []string{}
	usdpDataAddresses = strings.Split(usdpDataAddressesStr, ";")

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
		variableParserOffset := 0

		// parse asset from env
		entries := strings.Split(asset, "§")
		//check if len(array) > 0
		if len(entries) == 0 {
			log.Fatalf("No asset entries specified")
		}

		currAsset.blockchain = strings.TrimSpace(entries[0])
		currAsset.address = strings.TrimSpace(entries[1])
		currAsset.symbol = strings.TrimSpace(entries[2])

		// Check if deviation needs to be parsed per asset
		if deviationPerAssetMode {
			currAsset.deviation, err = strconv.ParseInt(strings.TrimSpace(entries[3]), 10, 64)
			variableParserOffset += 1
		} else {
			currAsset.deviation = int64(deviationPermille)
		}
		if len(entries) > (4 + variableParserOffset) {
			currAsset.coingeckoName = strings.TrimSpace(entries[3+variableParserOffset])
			currAsset.cmcName = strings.TrimSpace(entries[4+variableParserOffset])
			if currAsset.coingeckoName != "" || currAsset.cmcName != "" {
				allowedGuardianDeviation, err := strconv.ParseFloat(strings.TrimSpace(entries[5+variableParserOffset]), 64)
				if err != nil {
					log.Fatalf("Error converting guardian Deviation float on parsing %s-%s!", currAsset.blockchain, currAsset.address)
				}
				currAsset.allowedGuardianDeviation = allowedGuardianDeviation
			}
		}

		// Find out is there are additional GQL parameters for this asset
		if len(entries) > (6 + variableParserOffset) {
			// Join the rest of the line together, because the previous split might have affected substrings of the parameters
			gqlFeedSelectionQuery := strings.Join(entries[6+variableParserOffset:], "-")
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
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams, dataNodes, usdpDataAddresses, oracleDecimals)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice

					}
					log.Println(newAssetPrices)
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contract, conn, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contractBackup, connBackup, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
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
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams, dataNodes, usdpDataAddresses, oracleDecimals)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices, regardless of deviation
					emptyMap := make(map[string]float64)
					publishedPrices, err = oracleUpdateExecutor(emptyMap, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contract, conn, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(emptyMap, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contractBackup, connBackup, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
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
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams, dataNodes, usdpDataAddresses, oracleDecimals)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contract, conn, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, coingeckoApiKey, cmcApiKey, auth, contractBackup, connBackup, gasMultiplier, chainId, compatibilityMode, assets, conditionalPairs, oracleDecimals, mutexSeconds, oracleUpdateMutex)
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
	coingeckoApiKey string,
	cmcApiKey string,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	gasMultiplier float64,
	chainId int64,
	compatibilityMode bool,
	assets []Asset,
	conditionalAssets []ConditionalPair,
	oracleDecimals int,
	mutexSeconds int,
	oracleUpdateMutex sync.Mutex) (map[string]float64, error) {
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
				if asset0NewPrice > 1e-8 && ((asset0NewPrice > (asset0OldPrice * (1 + float64(asset0.deviation)/1000))) || (asset0NewPrice < (asset0OldPrice * (1 - float64(asset0.deviation)/1000)))) {
					updateAssetConditional = true
					log.Printf("Asset %s flagged for update because conditional asset %s is updated as well.", asset.symbol, asset0.symbol)
				} else {
					updateAssetConditional = false
					log.Printf("Asset %s is not updated because the conditional asset %s is not deviating.", asset.symbol, asset0.symbol)
				}
			}
		}

		newPrice := newPrices[asset.symbol]
		oldPrice := publishedPrices[asset.symbol]

		if updateAssetConditional || (newPrice > 1e-8 && ((newPrice > (oldPrice * (1 + float64(asset.deviation)/1000))) || (newPrice < (oldPrice * (1 - float64(asset.deviation)/1000))))) {
			var externalPrices []float64

			if asset.coingeckoName != "" {
				// Check coingecko for price deviation
				coingeckoPrice, err := getCoingeckoPrice(asset.coingeckoName, coingeckoApiKey)
				if err != nil {
					log.Printf("Error retrieving coingecko information for %s: %s", asset.symbol, err)
				}
				log.Printf("Coingecko price for %s: %f", asset.symbol, coingeckoPrice)
				externalPrices = append(externalPrices, coingeckoPrice)
			}
			if asset.cmcName != "" {
				cmcPrice, err := getCmcPrice(asset.cmcName, cmcApiKey)
				if err != nil {
					log.Printf("Error retrieving CMC information for %s: %s", asset.symbol, err)
				}
				log.Printf("CMC price for %s: %f", asset.symbol, cmcPrice)
				externalPrices = append(externalPrices, cmcPrice)
			}

			// Check if we have any external price available for the guardian
			if len(externalPrices) > 0 {
				numGuardianMatches := 0
				// Check for deviation
				for _, guardianPrice := range externalPrices {
					if math.Abs(guardianPrice-newPrice)/guardianPrice <= asset.allowedGuardianDeviation {
						numGuardianMatches += 1
					}
				}
				if numGuardianMatches == 0 && (asset.cmcName != "" || asset.coingeckoName != "") {
					log.Printf("Error: No guardian match found for asset %s with new price %f!", asset.symbol, newPrice)
					priceCollector[asset.symbol] = oldPrice
					continue
				}
			} else if asset.cmcName != "" || asset.coingeckoName != "" {
				// None of the external price providers works, but we expect at least one
				log.Printf("Error: None of the guardians returned a valid result")
				priceCollector[asset.symbol] = oldPrice
				continue
			}

			log.Printf("Entering deviation based update zone for old price %.2f of asset %s. New price: %.2f. Required deviation: %d permille", oldPrice, asset.symbol, newPrice, asset.deviation)
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
		integerPrice := int64(price * math.Pow(10, float64(oracleDecimals)))
		prices = append(prices, integerPrice)
	}

	// Update prices
	// check if we can do the multiupdate or use compatibility mode
	if compatibilityMode || len(keys) == 1 {
		for keyIndex := range keys {
			oracleUpdateMutex.Lock()
			timestamp := time.Now().Unix()
			err := updateOracleCompatibilityMode(conn, contract, auth, gasMultiplier, chainId, keys[keyIndex], prices[keyIndex], timestamp)
			time.Sleep(time.Duration(mutexSeconds) * time.Second)
			oracleUpdateMutex.Unlock()
			if err != nil {
				log.Printf("Failed to update Oracle: %v", err)
				return nil, err
			}
		}
	} else {
		oracleUpdateMutex.Lock()
		timestamp := time.Now().Unix()
		err := updateOracleMultiValues(conn, contract, auth, gasMultiplier, chainId, keys, prices, timestamp)
		time.Sleep(time.Duration(mutexSeconds) * time.Second)
		oracleUpdateMutex.Unlock()
		if err != nil {
			log.Printf("Failed to update Oracle: %v", err)
			return nil, err
		}
	}

	return priceCollector, nil
}

func retrieveAssetPrice(asset Asset, useGql bool, gqlWindowSize int, gqlMethodology string, gqlLiquidityParameters GqlParameters, dataNodes []string, usdpDataAddresses []string, oracleDecimals int) (float64, error) {
	var err error
	var price float64

	// Check if the asset is an RWA
	if strings.ToLower(strings.TrimSpace(asset.blockchain)) == "rwa-equity" {
		price, err = getRwaEquityPriceFromDia(asset.address)
		if err != nil {
			log.Printf("Failed to retrieve %s rwa equity price from DIA: %v", asset.address, err)
		}
		return price, nil
	} else if strings.ToLower(strings.TrimSpace(asset.blockchain)) == "rwa-commodity" {
		price, err = getRwaCommodityPriceFromDia(asset.address)
		if err != nil {
			log.Printf("Failed to retrieve %s rwa commodity price from DIA: %v", asset.address, err)
		}
		return price, nil
	}
	// Check if the asset is USDp
	if strings.ToLower(strings.TrimSpace(asset.blockchain)) == "fairvalue" && (strings.ToLower(strings.TrimSpace(asset.address)) == "usdp" || strings.ToLower(strings.TrimSpace(asset.address)) == "usdr") {
		var issuanceAccumulator float64
		var rawPriceAggregator float64
		for i, dataNode := range dataNodes {
			usdpDataAddress := usdpDataAddresses[i]
			usdpIssued, err := getChainSpecificUsdpIssued(dataNode, usdpDataAddress)
			if err != nil {
				log.Printf("Error getting chain-specific USDP issuance for address %s: %v",usdpDataAddress, err)
				continue
			}
			usdpPrice, err := getChainSpecificUsdpPrice(dataNode, usdpDataAddress)
			if err != nil {
				log.Printf("Error getting chain-specific USDP price for address %s: %v",usdpDataAddress, err)
				continue
			}
			log.Printf("USDp query for address %s returns price %f and tokens minted %f.", usdpDataAddress, usdpPrice, usdpIssued)

			issuanceAccumulator += usdpIssued
			rawPriceAggregator += usdpIssued * usdpPrice
	  }
	  if issuanceAccumulator == 0 {
	  	return 0.0, errors.New("No valid USDp data found") 
		}
		return float64(rawPriceAggregator / issuanceAccumulator), nil
	}
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

func updateOracleCompatibilityMode(
	client *ethclient.Client,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	auth *bind.TransactOpts,
	gasMultiplier float64,
	chainId int64,
	key string,
	value int64,
	timestamp int64) error {

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
		fGas.Mul(fGas, big.NewFloat(gasMultiplier))
		gasPrice, _ = fGas.Int(nil)
	}

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

func updateOracleMultiValues(
	client *ethclient.Client,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	auth *bind.TransactOpts,
	gasMultiplier float64,
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
		fGas.Mul(fGas, big.NewFloat(gasMultiplier))
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
	response, err := http.Get(diaBaseUrl + "/v1/assetQuotation/" + blockchain + "/" + address)
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
	// Workaround: Check for ETH and BTC, if yes use assetQuotation instead
	if address == "0x0000000000000000000000000000000000000000" && (blockchain == "Bitcoin" || blockchain == "Ethereum") && windowSize == 120 && len(gqlParameters.FeedSelection) == 0 {
		log.Printf("Enter assetQuotationRetrieval modus for BTC/ETH")
		return getAssetQuotationFromDia(blockchain, address)
	}
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

	if gqlParameters.GqlWindow > 0 {
		windowSize = gqlParameters.GqlWindow
	}

	// Get times for start/end
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize) * time.Second)

	type Response struct {
		GetFeed []struct {
			Name  string    `json:"Name"`
			Time  time.Time `json:"Time"`
			Value float64   `json:"Value"`
			Pools string    `json:"Pools"`
			Pairs string    `json:"Pairs"`
		} `json:"GetFeed"`
	}

	client := gql.NewClient(diaBaseUrl + "/graphql/query")
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

func getRwaEquityPriceFromDia(address string) (float64, error) {
	// Execute the query
	response, err := http.Get(diaBaseUrl + "/v1/rwa/Equities/" + address)
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
	var quotation models.ForeignQuotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return 0.0, err
	}
	return quotation.Price, nil
}

func getRwaCommodityPriceFromDia(address string) (float64, error) {
	// Execute the query
	response, err := http.Get(diaBaseUrl + "/v1/rwa/Commodities/" + address)
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
	var quotation models.ForeignQuotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return 0.0, err
	}
	return quotation.Price, nil
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
	price := gjson.Get(string(contents), assetName+".usd").Float()
	return price, nil
}

func getCmcPrice(assetName, cmcApiKey string) (float64, error) {
	url := "https://pro-api.coinmarketcap.com/v1/cryptocurrency/quotes/latest?id=" + assetName + "&CMC_PRO_API_KEY=" + cmcApiKey
	response, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on CMC API call with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	price := gjson.Get(string(contents), "data.*.quote.USD.price").Float()
	return price, nil
}

func getChainSpecificUsdpPrice(dataNode string, diaOracleAddress string) (float64, error) {
	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		log.Printf("Failed to connect to the data client: %v", err)
		return 0.0, err
	}
	diaOracleConn, err := NewMain(common.HexToAddress(diaOracleAddress), dataConn)
	if err != nil {
		log.Printf("Failed to instantiate DIAParallelOracle instance: %v", err)
		return 0.0, err
	}
	chainSpecificUsdpPrice, err := diaOracleConn.GetChainSpecificUsdpPrice(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to call getChainSpecificUsdpPrice in smart contract: %v", err)
		return 0.0, err
	}
	outputDecimalsBI, err := diaOracleConn.OutputDecimals(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to call outputDecimals in smart contract: %v", err)
		return 0.0, err
	}

	outputDecimals := outputDecimalsBI.Int64()
	price, _ := chainSpecificUsdpPrice.Float64()

	return price / math.Pow(10, float64(outputDecimals)), nil
}

func getChainSpecificUsdpIssued(dataNode string, diaOracleAddress string) (float64, error) {
	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		log.Printf("Failed to connect to the data client: %v", err)
		return 0.0, err
	}
	diaOracleConn, err := NewMain(common.HexToAddress(diaOracleAddress), dataConn)
	if err != nil {
		log.Printf("Failed to instantiate DIAParallelOracle instance: %v", err)
		return 0.0, err
	}
	chainSpecificUsdpIssued, err := diaOracleConn.GetChainSpecificUsdpIssued(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to call getChainSpecificUsdpIssued in smart contract: %v", err)
		return 0.0, err
	}
	retval, _ := chainSpecificUsdpIssued.Float64()
	return retval / math.Pow(10, 18), nil
}
