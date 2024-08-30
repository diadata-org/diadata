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
	baseUrl      = "https://api.prosperavest.com/api/rate"
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

	publishedValue := 0.0

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
					// Get cash balances from prosperavet API
  				rate, err := getRateFromApi(baseUrl)
				  if err != nil {
				  	log.Println("Error! Prosperavest API returned an error", err)
				  	continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(publishedValue, rate, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(publishedValue, rate, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				case <-mandatoryticker.C:
					// Get cash balances from prosperavet API
  				rate, err := getRateFromApi(baseUrl)
				  if err != nil {
				  	log.Println("Error! Prosperavest API returned an error", err)
				  	continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(0.0, rate, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(0.0, rate, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				}
			} else {
				select {
				case <-ticker.C:
					// Get cash balances from prosperavet API
  				rate, err := getRateFromApi(baseUrl)
				  if err != nil {
				  	log.Println("Error! Prosperavest API returned an error", err)
				  	continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(publishedValue, rate, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(publishedValue, rate, deviationPermille, auth, contractBackup, connBackup, chainId)
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
	publishedValue float64,
	newValue float64,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	chainId int64) (float64, error) {

	// If a published price is 0, update in any case
	fmt.Println("new price", newValue)

	if (newValue > 1e-8 && ((newValue > (publishedValue * (1 + float64(deviationPermille)/1000))) || (newValue < (publishedValue * (1 - float64(deviationPermille)/1000))))) {
		log.Printf("Entering deviation based update zone for old value %.2f. New value: %.2f", publishedValue, newValue)
	} else {
		log.Println("No update necessary.")
		return publishedValue, nil
	}

	// Serialize keys and values into array for oracle update transaction
	var keys []string
	var values []int64

	keys = append(keys, "NGN/USD")
	values = append(values, int64(newValue * 100000000))

	// Update values in one transaction
	timestamp := time.Now().Unix()
	err := updateOracleMultiValues(conn, contract, auth, chainId, keys, values, timestamp)
	if err != nil {
		log.Printf("Failed to update Oracle: %v", err)
		return 0.0, err
	}

	return newValue, nil
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
	fGas.Mul(fGas, big.NewFloat(1.1))
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

func getRateFromApi(baseUrl string) (float64, error) {
  // Execute the query
  response, err := http.Get(baseUrl)
  if err != nil {
    return 0.0, err
  }

  defer response.Body.Close()
  if 200 != response.StatusCode {
    return 0.0, fmt.Errorf("Error on Prosperavest api with return code %d", response.StatusCode)
  }
  contents, err := ioutil.ReadAll(response.Body)
  if err != nil {
    return 0.0, err
  }

  exchangeRate := gjson.Get(string(contents), "values.0.centralrate").Float()
  return exchangeRate, nil
}
