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

	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

const (
	baseUrl      = "https://price.denario.swiss/api"
)

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
	authKey := utils.Getenv("DENARIO_AUTH_KEY", "")
	denarioAssetsStr := utils.Getenv("DENARIO_ASSETS", "")

	publishedValues := make(map[string]float64)
	var denarioAssets []string

	denarioAssetsParsed := strings.Split(denarioAssetsStr, ";")

	for _, asset := range denarioAssetsParsed {
		denarioAssets = append(denarioAssets, strings.TrimSpace(asset))
	}

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
					for _, asset := range denarioAssets {
						newAssetPrice, err := retrieveDenarioAssetPrice(baseUrl, asset, authKey)
						if err != nil {
							log.Println("Denario API returned an error: ", err)
							continue
						}
						newAssetPrices[asset] = newAssetPrice
					}
					log.Println(newAssetPrices)
					// update value
					publishedValues, err = oracleUpdateExecutor(publishedValues, newAssetPrices, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValues, err = oracleUpdateExecutor(publishedValues, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				case <-mandatoryticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range denarioAssets {
						newAssetPrice, err := retrieveDenarioAssetPrice(baseUrl, asset, authKey)
						if err != nil {
							log.Println("Denario API returned an error: ", err)
							continue
						}
						newAssetPrices[asset] = newAssetPrice
					}
					log.Println(newAssetPrices)
					emptyMap := make(map[string]float64)
					// update value
					publishedValues, err = oracleUpdateExecutor(emptyMap, newAssetPrices, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValues, err = oracleUpdateExecutor(emptyMap, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId)
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
					for _, asset := range denarioAssets {
						newAssetPrice, err := retrieveDenarioAssetPrice(baseUrl, asset, authKey)
						if err != nil {
							log.Println("Denario API returned an error: ", err)
							continue
						}
						newAssetPrices[asset] = newAssetPrice
					}
					log.Println(newAssetPrices)
					// update value
					publishedValues, err = oracleUpdateExecutor(publishedValues, newAssetPrices, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValues, err = oracleUpdateExecutor(publishedValues, newAssetPrices, deviationPermille, auth, contractBackup, connBackup, chainId)
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
	publishedValues map[string]float64,
	newAssetPrices map[string]float64,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	chainId int64) (map[string]float64, error) {

	// Check for deviation and collect all new prices in a map
	// If a published price is 0, update in any case
	updateCollector := make(map[string]float64)
	priceCollector := make(map[string]float64)

	for asset, _ := range newAssetPrices {
		newPrice := newAssetPrices[asset]
		publishedValue := publishedValues[asset]

		if (newPrice > 1e-8 && ((newPrice > (publishedValue * (1 + float64(deviationPermille)/1000))) || (newPrice < (publishedValue * (1 - float64(deviationPermille)/1000))))) {
			log.Printf("Entering deviation based update zone for old value %.2f. New value: %.2f", publishedValue, newPrice)
			updateCollector[asset] = newPrice
			priceCollector[asset] = newPrice
		} else {
			priceCollector[asset] = publishedValue
		}
	}

	// Check if any update is to be executed
	if len(updateCollector) == 0 {
		log.Println("No update necessary.")
		return priceCollector, nil
	}

	// Serialize keys and values into array for oracle update transaction
	var keys []string
	var values []int64

	for key, price := range updateCollector {
		key = key
		keys = append(keys, key)
		integerPrice := int64(price * 100000000)
		values = append(values, integerPrice)
	}

	// Update values in one transaction
	timestamp := time.Now().Unix()
	err := updateOracleMultiValues(conn, contract, auth, chainId, keys, values, timestamp)
	if err != nil {
		log.Printf("Failed to update Oracle: %v", err)
		return nil, err
	}

	return priceCollector, nil
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

	// Get gas price suggestion
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Print(err)
		return err
	}

	// Get 110% of the gas price
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.5))
	gasPrice, _ = fGas.Int(nil)

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

func retrieveDenarioAssetPrice(baseUrl string, assetName string, authKey string) (float64, error) {
  // Execute the query
	client := http.Client{}
	request, err := http.NewRequest(http.MethodGet, baseUrl + "/" + assetName, nil)
  if err != nil {
    return 0.0, err
  }
  request.Header.Set("X-AUTH-TOKEN", authKey)
  request.Header.Set("Content-Type", "application/json")

  response, err := client.Do(request)
  if err != nil {
    return 0.0, err
  }

  defer response.Body.Close()
  if 200 != response.StatusCode {
    return 0.0, fmt.Errorf("Error on Denario api with return code %d", response.StatusCode)
  }
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return 0.0, err
  }

  exchangeRate := gjson.Get(string(contents), "priceAsk").Float()
  return exchangeRate, nil
}
