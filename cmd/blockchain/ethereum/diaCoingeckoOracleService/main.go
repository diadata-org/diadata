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

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaCoingeckoOracleService"
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
	var numCoins = flag.Int("numCoins", 100, "Number of coins to push with the oracle")
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

	var contract *diaCoingeckoOracleService.DIACoingeckoOracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}
	periodicOracleUpdateHelper(numCoins, auth, contract)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(numCoins, auth, contract)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(numCoins *int, auth *bind.TransactOpts, contract *diaCoingeckoOracleService.DIACoingeckoOracle) error {

	topCoins, err := getTopCoinsFromCoingecko(*numCoins)
	if err != nil {
		log.Fatalf("Failed to get top %d coins from Coingecko: %v", numCoins, err)
	}
	// Get quotation for topCoins and update Oracle
	for _, symbol := range topCoins {
		rawQuot, err := getForeignQuotationFromDia("Coingecko", symbol)
		if err != nil {
			log.Fatalf("Failed to retrieve Coingecko data from DIA: %v", err)
			return err
		}
		err = updateForeignQuotation(rawQuot, auth, contract)
		if err != nil {
			log.Fatalf("Failed to update Coingecko Oracle: %v", err)
			return err
		}
		time.Sleep(5 * time.Minute)
	}

	return nil
}

func updateForeignQuotation(foreignQuotation *models.ForeignQuotation, auth *bind.TransactOpts, contract *diaCoingeckoOracleService.DIACoingeckoOracle) error {
	symbol := foreignQuotation.Symbol
	price := foreignQuotation.Price
	timestamp := foreignQuotation.Time.Unix()
	err := updateOracle(contract, auth, symbol, int64(price*100000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

// getTopCoinsFromCoingecko returns the symbols of the top @numCoins assets from coingecko by market cap
func getTopCoinsFromCoingecko(numCoins int) ([]string, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/coins/markets?vs_currency=usd&order=market_cap_desc&per_page=" + strconv.Itoa(numCoins) + "&page=1&sparkline=false")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on coingecko api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	type aux struct {
		Symbol string `json:"symbol"`
	}
	var quotations []aux
	err = json.Unmarshal(contents, &quotations)
	if err != nil {
		return []string{}, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	var symbols []string
	for _, quotation := range quotations {
		symbols = append(symbols, strings.ToUpper(quotation.Symbol))
	}
	return symbols, nil
}

func getForeignQuotationFromDia(source, symbol string) (*models.ForeignQuotation, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/foreignQuotation/" + strings.Title(strings.ToLower(source)) + "/" + strings.ToUpper(symbol))
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
	var quotation models.ForeignQuotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **diaCoingeckoOracleService.DIACoingeckoOracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = diaCoingeckoOracleService.NewDIACoingeckoOracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaCoingeckoOracleService.DeployDIACoingeckoOracle(auth, conn)
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
	contract *diaCoingeckoOracleService.DIACoingeckoOracle,
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
