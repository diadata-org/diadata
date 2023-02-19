package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaOracleServiceV2 "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/tidwall/gjson"
)

type Floor struct {
	Value     float64   `json:"Floor_Price"`
	Timestamp time.Time `json:"Time"`
	Source    string    `json:"Source"`
}

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	sleepSeconds, err := strconv.Atoi(utils.Getenv("SLEEP_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse sleepSeconds: %v")
	}
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v")
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v")
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v")
	}
	assetsStr := utils.Getenv("ASSETS", "")

	addresses := []string{}
	blockchains := []string{}

	assetsParsed := strings.Split(assetsStr, ",")

	for _, asset := range assetsParsed {
		entries := strings.Split(asset, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		addresses = append(addresses, strings.TrimSpace(entries[1]))
	}

	oldValues := make(map[int]float64)

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
					oldValue := oldValues[i]
					log.Println("old value", oldValue)
					oldValue, err = periodicOracleUpdateHelper(oldValue, deviationPermille, auth, contract, conn, blockchain, address)
					oldValues[i] = oldValue
					if err != nil {
						log.Println(err)
					}
					time.Sleep(time.Duration(sleepSeconds) * time.Second)
				}
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(oldValue float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string) (float64, error) {

	// Empty quotation for our request
	var err error

	// Get quotation for NFT and update Oracle
	rawAvg, err := getNFTDistributionFromDia(blockchain, address)
	if err != nil {
		log.Fatalf("Failed to retrieve %s NFT Average data from DIA: %v", address, err)
		return oldValue, err
	}

	// Check for deviation
	newValue := rawAvg

	if (newValue > (oldValue * (1 + float64(deviationPermille)/1000))) || (newValue < (oldValue * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		// Get correspondin floor price
		floor, err := getFloor(blockchain, address)
		if err != nil {
			log.Fatalf("Failed to get floor price: %v", err)
			return oldValue, err
		}

		rawFloor := floor.Value

		err = updateNFTAverageAndFloor(rawAvg, rawFloor, blockchain, address, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldValue, err
		}
		return newValue, nil
	}

	return oldValue, nil
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

func updateNFTAverageAndFloor(average float64, floor float64, blockchain string, address string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	symbol := blockchain + "-" + address 
	price := average
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}
	err = updateOracle(conn, contract, auth, symbol + "-Floor", int64(floor*100000000), timestamp)
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
		From:     auth.From,
		Signer:   auth.Signer,
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

func getNFTDistributionFromDia(blockchain, address string) (float64, error) {
	response, err := http.Get("https://api.diadata.org/v1/NFTDistribution/" + blockchain + "/" + address)
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
	
	average := gjson.Get(string(contents), "Average").Float()
	return average, nil
}

func getFloor(blockchain, address string) (Floor, error) {
	response, err := http.Get("https://api.diadata.org/v1/NFTFloor/" + blockchain + "/" + address)
	if err != nil {
		return Floor{}, err
	}
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return Floor{}, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Floor{}, err
	}

	var resp Floor
	err = json.Unmarshal(contents, &resp)
	if err != nil {
		return Floor{}, err
	}

	return resp, err
}
