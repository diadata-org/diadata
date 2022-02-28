package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strings"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

func main() {
	/*
	 * Read in Oracle address
	 */

	var deployedContract = flag.String("deployedContract", "", "Address of the deployed oracle contract")
	var secretsFile = flag.String("secretsFile", "/run/secrets/oracle_keys", "File with wallet secrets")
	var blockchainNode = flag.String("blockchainNode", "https://matic-mainnet-full-rpc.bwarelabs.com", "Node address for blockchain connection")
	var sleepSeconds = flag.Int("sleepSeconds", 10, "Number of seconds to sleep between calls")
	var frequencySeconds = flag.Int("frequencySeconds", 120, "Number of seconds to sleep between checking oracle runs")
	var chainId = flag.Int64("chainId", 137, "Chain-ID of the network to connect to")
	var apiSecretsFile = flag.String("apiSecretsFile", "/run/secrets/apisecrets", "File with API secrets")
	flag.Parse()

	/*
	 * Read secrets for unlocking the ETH account
	 */
	var lines []string
	file, err := os.Open(*secretsFile) // Read in key information
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 2 {
		log.Fatal("Secrets file should have exactly two lines")
	}
	key := lines[0]
	key_password := lines[1]

	symbols := []string{"ETH", "SOL"}
	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(*blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(*chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaOracleServiceV2.DIAOracleV2
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(*frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, symbol := range symbols {
					err = periodicOracleUpdateHelper(auth, contract, conn, symbol, *apiSecretsFile)
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(*sleepSeconds) * time.Second)
				}
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, symbol string, apiSecretsFile string) (error) {

	// Get latest API value and update Oracle
	rawResult, err := getValueFromCfBenchmarks(symbol, apiSecretsFile)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from CFBenchmarks: %v", symbol, err)
		return err
	}

	err = updateFloatValue("DEFI_CF" + symbol, rawResult, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return err
	}
	return nil
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

func updateFloatValue(symbol string, value float64, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(value*100), timestamp)
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

func getValueFromCfBenchmarks(symbol string, apiSecretFile string) (float64, error) {
	// Get the API key
	var lines []string
	file, err := os.Open(apiSecretFile) // Read in key information
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	if len(lines) != 2 {
		log.Fatal("API secret file should have exactly two lines")
	}
	apiUsername := lines[0]
	apiPassword := lines[1]

	// HTTP Basic Auth
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://oracleprod1.cfbenchmarks.com/api/v1/summary?id=DEFI_CF" + symbol + "USD", nil)
	if err != nil {
		return 0, fmt.Errorf("Got error %s", err.Error())
	}
	req.SetBasicAuth(apiUsername, apiPassword)
	response, err := client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("Got error %s", err.Error())
	}
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, err
	}

	rawResult := gjson.Get(string(contents), "payload.value")
	retval, err := strconv.ParseFloat(rawResult.Str, 64)
	if err != nil {
		return 0, err
	}
	return retval, nil
}
