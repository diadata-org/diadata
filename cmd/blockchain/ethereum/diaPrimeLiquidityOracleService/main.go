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

var diaBaseUrl string

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
	liquidityRatioDeviationPermille, err := strconv.Atoi(utils.Getenv("LIQUIDITY_RATIO_DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse liquidityDeviationPermille: %v")
	}
	poolsStr := utils.Getenv("POOLS", "")
	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")

	// Parse Pools
	poolsParsed := strings.Split(poolsStr, ",")
	poolAddresses := []string{}
	blockchains := []string{}

	for _, pool := range poolsParsed {
		entries := strings.Split(pool, "-")
		blockchains = append(blockchains, strings.TrimSpace(entries[0]))
		poolAddresses = append(poolAddresses, strings.TrimSpace(entries[1]))
	}

	oldLiqAs := make(map[int]float64)
	oldLiqBs := make(map[int]float64)
	oldLiquidityRatios := make(map[int]float64)

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
	 * Update Oracle periodically with liquidity pools
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				for i, poolAddress := range poolAddresses {
					blockchain := blockchains[i]
					oldLiqA := oldLiqAs[i]
					oldLiqB := oldLiqBs[i]
					oldLiquidityRatio := oldLiquidityRatios[i]
					log.Println("old liqA", oldLiqA)
					log.Println("old liqB", oldLiqB)
					log.Println("old ratio", oldLiquidityRatio)
					oldLiqA, oldLiqB, oldLiquidityRatio, err = periodicOracleUpdateHelper(oldLiqA, oldLiqB, oldLiquidityRatio, deviationPermille, liquidityRatioDeviationPermille, auth, contract, conn, blockchain, poolAddress, sleepSeconds)
					oldLiqAs[i] = oldLiqA
					oldLiqBs[i] = oldLiqB
					oldLiquidityRatios[i] = oldLiquidityRatio
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

func periodicOracleUpdateHelper(oldLiqA float64, oldLiqB float64, oldLiquidityRatio float64, deviationPermille int, ratioDeviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, poolAddress string, sleepSeconds int) (float64, float64, float64, error) {

	// Get liquidity values
	liquidityValues, liquiditySymbols, err := getLiquidityValuesFromDia(blockchain, poolAddress)
	if err != nil {
		log.Printf("Failed to retrieve %s poolLiquidity from DIA: %v", poolAddress, err)
			return oldLiqA, oldLiqB, oldLiquidityRatio, err
	}

	// Check for deviation
	toUpdate := false
	newLiqA := liquidityValues[0]
	newLiqB := liquidityValues[1]
	newLiquidityRatio := newLiqA / newLiqB

	if math.Abs(newLiquidityRatio - oldLiquidityRatio) > oldLiquidityRatio * float64(ratioDeviationPermille) / 1000 {
		log.Println("Entering liquidity ratio deviation based update zone")
		toUpdate = true
	} else if math.Abs(oldLiqA - newLiqA) > oldLiqA * float64(deviationPermille) / 1000 {
		log.Println("Entering pool asset A deviation based update zone")
		toUpdate = true
	} else if math.Abs(oldLiqB - newLiqB) > oldLiqB * float64(deviationPermille) / 1000 {
		log.Println("Entering pool asset B deviation based update zone")
		toUpdate = true
	}

	if toUpdate {
		err = updateLiquidity(newLiqA, liquiditySymbols[0], poolAddress, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldLiqA, oldLiqB, oldLiquidityRatio, err
		}

		time.Sleep(time.Duration(sleepSeconds) * time.Second)
		err = updateLiquidity(newLiqB, liquiditySymbols[1], poolAddress, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldLiqA, oldLiqB, oldLiquidityRatio, err
		}
		return newLiqA, newLiqB, newLiquidityRatio, nil
	}

	return oldLiqA, oldLiqB, oldLiquidityRatio, nil
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

func updateLiquidity(liquidityValue float64, assetSymbol string, poolAddress string, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client) error {
	key := poolAddress + "-" + assetSymbol + "-LQ"
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, key, int64(liquidityValue * 100000000), timestamp)
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
		GasLimit: 1000725,
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

func getLiquidityValuesFromDia(blockchain, poolAddress string) ([]float64, []string, error) {
	response, err := http.Get(diaBaseUrl + "/v1/poolLiquidity/" + blockchain + "/" + poolAddress)
	if err != nil {
		return []float64{}, []string{}, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return []float64{}, []string{}, fmt.Errorf("Error on dia poolLiquidity api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return []float64{}, []string{}, err
	}

	liqA := gjson.Get(string(contents), "Liquidity.0.Liquidity")
	liqB := gjson.Get(string(contents), "Liquidity.1.Liquidity")
	symbolA := gjson.Get(string(contents), "Liquidity.0.Asset.Symbol")
	symbolB := gjson.Get(string(contents), "Liquidity.1.Asset.Symbol")
	return []float64{liqA.Float(), liqB.Float()}, []string{symbolA.String(), symbolB.String()}, nil
}
