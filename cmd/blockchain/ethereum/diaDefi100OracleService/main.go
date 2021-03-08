package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"math/big"
	"net/http"
	"os"
	"strings"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaDefi100OracleService"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func init() {
	log = logrus.New()
}

func main() {
	/*
	 * Read in Oracle address
	 */
	var deployedContract = flag.String("deployedContract", "", "Address of the deployed oracle contract")
	var secretsFile = flag.String("secretsFile", "/run/secrets/oracle_keys_defi100_bsc", "File with wallet secrets")
	var blockchainNode = flag.String("blockchainNode", "https://bsc-dataseed.binance.org/", "Node address for blockchain connection")
	var sleepSeconds = flag.Int("sleepSeconds", 120, "Number of seconds to sleep between calls")
	var frequencySeconds = flag.Int("frequencySeconds", 86400, "Number of seconds to sleep between full oracle runs")
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

	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(*blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), key_password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaDefi100OracleService.DIADefi100Oracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}
	periodicOracleUpdateHelper(*sleepSeconds, auth, contract, conn)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(*frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(*sleepSeconds, auth, contract, conn)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(sleepSeconds int, auth *bind.TransactOpts, contract *diaDefi100OracleService.DIADefi100Oracle, conn *ethclient.Client) error {

	time.Sleep(time.Duration(sleepSeconds) * time.Second)
	// Get fresh market cap data and update Oracle
	marketcap, err := getDefiMCFromCoingecko()
	if err != nil {
		log.Fatalf("Failed to get data from Coingecko: %v", err)
	}

	err = updateMarketCap(marketcap, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Defi100 Oracle: %v", err)
		return err
	}

	return nil
}

// getDefiMCFromCoingecko returns the market cap of the top 100 defi tokens
func getDefiMCFromCoingecko() (float64, error) {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.coingecko.com/api/v3/global/decentralized_finance_defi", nil)
	if err != nil {
		log.Print(err)
		return 0.0, err
	}

	response, err := client.Do(req)
	if err != nil {
		log.Print("Error sending request to server: ", err)
		return 0.0, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on coingecko api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	type CoingeckoData struct {
		Data struct {
			DefiMarketCap string `json:"defi_market_cap"`
		} `json:"data"`
	}
	var rawdata CoingeckoData
	err = json.Unmarshal(contents, &rawdata)
	if err != nil {
		return 0.0, fmt.Errorf("Error on unmarshaling data from coingecko:", contents)
	}
	marketCap, err := strconv.ParseFloat(rawdata.Data.DefiMarketCap, 64)
	if err != nil {
		log.Error("Error parsing string %s to float", rawdata.Data.DefiMarketCap)
	}

	return marketCap, nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **diaDefi100OracleService.DIADefi100Oracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = diaDefi100OracleService.NewDIADefi100Oracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaDefi100OracleService.DeployDIADefi100Oracle(auth, conn)
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

func updateMarketCap(marketCap float64, auth *bind.TransactOpts, contract *diaDefi100OracleService.DIADefi100Oracle, conn *ethclient.Client) error {
	symbol := "Defi100"
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(marketCap*100000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateOracle(
	client *ethclient.Client,
	contract *diaDefi100OracleService.DIADefi100Oracle,
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
		GasLimit: 800725,
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
