package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaScifiOracleService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	/*
	 * Read in Oracle address
	 */
	var deployedContract = flag.String("deployedContract", "", "Address of the deployed oracle contract")
	flag.Parse()

	/*
	 * Read secrets for unlocking the ETH account
	 */
	var lines []string
	file, err := os.Open("/run/secrets/oracle_keys") // Read in key information
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

	conn, err := ethclient.Dial("http://159.69.120.42:8545/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), key_password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaScifiOracleService.DIAScifiOracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}
	indexName := "SCIFI"
	periodicOracleUpdateHelper(indexName, auth, contract)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(indexName, auth, contract)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(indexName string, auth *bind.TransactOpts, contract *diaScifiOracleService.DIAScifiOracle) error {
	rawIndex, err := getIndexValueFromDia(indexName)
	if err != nil {
		log.Fatalf("Failed to retrieve crypto index data from DIA: %v", err)
		return err
	}
	err = updateIndexValue(rawIndex, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Coinmarketcap Oracle: %v", err)
		return err
	}

	return nil
}

func updateIndexValue(iv *models.CryptoIndex, auth *bind.TransactOpts, contract *diaScifiOracleService.DIAScifiOracle) error {
	symbol := iv.Name
	value := iv.Value
	timestamp := iv.CalculationTime.Unix()
	err := updateOracle(contract, auth, symbol, int64(value * 10000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func getIndexValueFromDia(symbol string) (*models.CryptoIndex, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/index/" + symbol + "?starttime=" + strconv.FormatInt(time.Now().Add(-200 * time.Second).Unix(), 10) + "&endtime=" + strconv.FormatInt(time.Now().Unix(), 10))
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
	var indices []models.CryptoIndex
	err = json.Unmarshal([]byte(contents), &indices)
	if err != nil {
		return nil, err
	}
	return &indices[0], nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **diaScifiOracleService.DIAScifiOracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = diaScifiOracleService.NewDIAScifiOracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaScifiOracleService.DeployDIAScifiOracle(auth, conn)
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

func updateOracle(
	contract *diaScifiOracleService.DIAScifiOracle,
	auth *bind.TransactOpts,
	key string,
	value int64,
	timestamp int64) error {
	// Write values to smart contract
	tx, err := contract.SetValue(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 800725,
		//	Nonce: big.NewInt(time.Now().Unix()),
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
