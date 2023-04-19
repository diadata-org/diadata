package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math"
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

type Dominance struct {
	Name string
	Percentage float64
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
	timeBasedUpdateSeconds, err := strconv.Atoi(utils.Getenv("TIME_BASED_UPDATE_SECONDS", "1200"))
	if err != nil {
		log.Fatalf("Failed to parse timeBasedUpdateSeconds: %v", err)
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v")
	}
	deviationBasisPoints, err := strconv.Atoi(utils.Getenv("DEVIATION_BASIS_POINTS", "100"))
	if err != nil {
		log.Fatalf("Failed to parse deviationBasisPoints: %v")
	}

	oldDominances := make(map[string]float64)

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
	timeBasedUpdateTicker := time.NewTicker(time.Duration(timeBasedUpdateSeconds) * time.Second)
	var timeBasedUpdate bool
	go func() {
		for {
			select {
			case <-ticker.C:
				timeBasedUpdate = false
			case <-timeBasedUpdateTicker.C:
				timeBasedUpdate = true
			}
			err = periodicOracleUpdateHelper(deviationBasisPoints, timeBasedUpdate, sleepSeconds, auth, contract, conn, oldDominances)
			if err != nil {
				log.Println(err)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(deviationBasisPoints int, update bool, sleepSeconds int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, oldDominances map[string]float64) (error) {

	// Get dominance metrics and update Oracle
	dominances, err := getDominancesFromCoingecko()
	if err != nil {
		log.Fatalf("Failed to retrieve dominance data from Coingecko: %v", err)
		return err
	}

	for _, dominance := range dominances {
		oldDominance := oldDominances[dominance.Name]
		newDominance := dominance.Percentage

		// Check for update condition
		if math.Abs(newDominance - oldDominance) > oldDominance * float64(deviationBasisPoints)/10000 || update {
			log.Println("Entering update zone")
			err = updateDominance(dominance, auth, contract, conn)
			if err != nil {
				log.Fatalf("Failed to update DIA Oracle: %v", err)
				return err
			}
			oldDominances[dominance.Name] = newDominance
			time.Sleep(time.Duration(sleepSeconds) * time.Second)
		}
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

func updateDominance(dominance Dominance, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	symbol := dominance.Name + "-DOM"
	percentage := dominance.Percentage
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(percentage*100000000), timestamp)
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

// getDominanceFromCoingecko returns the dominations of the top @numCoins assets from coingecko
func getDominancesFromCoingecko() ([]Dominance, error) {
	response, err := http.Get("https://api.coingecko.com/api/v3/global/coin_dominance")
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on GET for Coingecko API with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	symbols := gjson.Get(string(contents), "data.#.name")
	percentages := gjson.Get(string(contents), "data.#.dominance_percentage")

	var dominances []Dominance
	for i, symbol := range symbols.Array() {
		newDominance := Dominance{symbol.String(), percentages.Array()[i].Float()}
		dominances = append(dominances, newDominance)
	}
	return dominances, nil
}
