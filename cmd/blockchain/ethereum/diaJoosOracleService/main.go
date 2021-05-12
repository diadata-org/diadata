package main

import (
	"bufio"
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"strings"
	"time"

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/diaCoingeckoOracleService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
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
	var secretsFile = flag.String("secretsFile", "/run/secrets/oracle_keys", "File with wallet secrets")
	var blockchainNode = flag.String("blockchainNode", "http://159.69.120.42:8545/", "Node address for blockchain connection")
	var sleepSeconds = flag.Int("sleepSeconds", 120, "Number of seconds to sleep between calls")
	var frequencySeconds = flag.Int("frequencySeconds", 86400, "Number of seconds to sleep between full oracle runs")
	var chainId = flag.Int64("chainId", 1, "Chain-ID of the network to connect to")
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

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(*chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaCoingeckoOracleService.DIACoingeckoOracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}
	periodicOracleUpdateHelper(numCoins, *sleepSeconds, auth, contract, conn)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(*frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(numCoins, *sleepSeconds, auth, contract, conn)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(numCoins *int, sleepSeconds int, auth *bind.TransactOpts, contract *diaCoingeckoOracleService.DIACoingeckoOracle, conn *ethclient.Client) error {

	// Get quotation for JOOS coin and update Oracle
	rawQuot, err := getForeignQuotationByAddress("0x05f9abf4b0c5661e83b92c056a8791d5ccd7ca52")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data for JOOS: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuot, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle for JOOS: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Get quotation for WBTC coin and update Oracle
	rawQuotWBTC, err := getForeignQuotationByAddress("0x2260fac5e5542a773aa44fbcfedf7c193bc2c599")
	if err != nil {
		log.Fatalf("Failed to retrieve Coingecko data for WBTC: %v", err)
		return err
	}
	err = updateForeignQuotation(rawQuotWBTC, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Coingecko Oracle for WBTC: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	return nil
}

func updateForeignQuotation(foreignQuotation *models.ForeignQuotation, auth *bind.TransactOpts, contract *diaCoingeckoOracleService.DIACoingeckoOracle, conn *ethclient.Client) error {
	symbol := foreignQuotation.Symbol
	price := foreignQuotation.Price
	timestamp := foreignQuotation.Time.Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
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
	client *ethclient.Client,
	contract *diaCoingeckoOracleService.DIACoingeckoOracle,
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

// ------------------------------------------------------------
// Methods for getting additional foreign quotations by address
// ------------------------------------------------------------

type CoinData struct {
	Prices     [][]float64 `json:"prices"`
	MarketCaps [][]float64 `json:"market_caps"`
	Volumes    [][]float64 `json:"total_volumes"`
}

type BasicTokenInfo struct {
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

func getCoinInfoByAddress(address string) (name, symbol string, err error) {
	// Pull and unmarshall data from coingecko API
	response, err := utils.GetRequest("https://api.coingecko.com/api/v3/coins/ethereum/contract/" + address)
	if err != nil {
		return
	}
	var tokenInfo BasicTokenInfo
	err = json.Unmarshal(response, &tokenInfo)
	if err != nil {
		return
	}
	name = tokenInfo.Name
	symbol = tokenInfo.Symbol
	return
}

func getForeignQuotationByAddress(address string) (*models.ForeignQuotation, error) {
	// Pull and unmarshall data from coingecko API
	response, err := utils.GetRequest("https://api.coingecko.com/api/v3/coins/ethereum/contract/" + address + "/market_chart/?vs_currency=usd&days=2")
	if err != nil {
		return &models.ForeignQuotation{}, err
	}
	var coin CoinData
	err = json.Unmarshal(response, &coin)
	if err != nil {
		return &models.ForeignQuotation{}, err
	}

	// Get index of most recent timestamp
	prices := coin.Prices
	latestTimestamp := time.Time{}
	indexLatestTimestamp := 0
	for i, price := range prices {
		tm := time.Unix(int64(price[0]/1000), 0)
		if latestTimestamp.Before(tm) {
			latestTimestamp = tm
			indexLatestTimestamp = i
		}
	}

	// Get index of timestamp closest to latestTimestamp-24h
	indexYesterday := 0
	yesterday := latestTimestamp.AddDate(0, 0, -1)
	minDuration := time.Duration(int64(1e16))
	for i, price := range prices {
		tm := time.Unix(int64(price[0]/1000), 0)
		diff := yesterday.Sub(tm)
		if diff < 0 {
			diff = tm.Sub(yesterday)
		}
		if diff < minDuration {
			minDuration = diff
			indexYesterday = i
		}
	}

	// Get name and symbol by address from coingecko
	name, symbol, err := getCoinInfoByAddress(address)
	if err != nil {
		return &models.ForeignQuotation{}, err
	}

	var fq models.ForeignQuotation
	fq.Symbol = strings.ToUpper(symbol)
	fq.Name = name
	fq.Price = prices[indexLatestTimestamp][1]
	fq.PriceYesterday = prices[indexYesterday][1]
	fq.VolumeYesterdayUSD = coin.Volumes[indexYesterday][1]
	fq.Source = "Coingecko"
	fq.Time = latestTimestamp
	return &fq, nil
}
