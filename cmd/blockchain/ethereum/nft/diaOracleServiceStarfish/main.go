package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaNFTOracleService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaNFTOracleService"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

type StarfishReturn struct {
	Blockchain string
	Address    string
	Floor      float64
	FloorMA    float64
}

type Floor struct {
	Value     float64   `json:"Floor_Price"`
	Timestamp time.Time `json:"Time"`
	Source    string    `json:"Source"`
}

type FloorMA struct {
	Value     float64   `json:"Moving_Average_Floor_Price"`
	Timestamp time.Time `json:"Time"`
	Source    string    `json:"Source"`
}

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	sleepSeconds, err := strconv.Atoi(utils.Getenv("SLEEP_SECONDS", "60"))
	if err != nil {
		log.Fatalf("Failed to parse sleepSeconds: %v", err)
	}
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "60"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}

	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "80001"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "30"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v", err)
	}

	addresses := []string{
		"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D", //BAYC
		"0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB", //Cryptopunks
		"0x60E4d786628Fea6478F785A6d7e704777c86a7c6", //Mutant Ape Yacht Club
		"0x49cF6f5d44E70224e2E23fDcdd2C053F30aDA28B", //Clone X
		"0x8a90CAb2b38dba80c64b7734e58Ee1dB38B8992e", //Doodles
		"0xED5AF388653567Af2F388E6224dC7C4b3241C544", //Azuki
		"0xd59fC6Bfd9732AB19b03664a45dC29B8421BDA9a", //Astar Degens
	}
	blockchains := []string{
		"Ethereum",
		"Ethereum",
		"Ethereum",
		"Ethereum",
		"Ethereum",
		"Ethereum",
		"Astar",
	}

	//
	oldFloors := make(map[string]float64)

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

	var contract *diaNFTOracleService.DIANFTOracle
	err = deployOrBindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	/*
	 * Update Oracle periodically with top collections. Check each @frequencySeconds whether deviation
	 * exceeds threshold and update if so.
	 */

	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for range ticker.C {
			// Update all collections depending on @oldFloor and @timeBasedUpdate.
			for i, address := range addresses {
				blockchain := blockchains[i]
				oldFloor := oldFloors[address]
				log.Println("old floor", oldFloor)
				newFloor, err := periodicOracleUpdateHelper(oldFloor, deviationPermille, auth, contract, conn, blockchain, address)
				oldFloors[address] = newFloor
				if err != nil {
					log.Println(err)
				}
				time.Sleep(time.Duration(sleepSeconds) * time.Second)
			}
		}
	}()
	select {}
}

// periodicOracleUpdateHelper updates a collection on the condition:
// The difference of the (new) floor price and @oldFloor exceeds @deviationPermille.
func periodicOracleUpdateHelper(oldFloor float64, deviationPermille int, auth *bind.TransactOpts, contract *diaNFTOracleService.DIANFTOracle, conn *ethclient.Client, blockchain string, address string) (float64, error) {
	var data StarfishReturn
	data.Blockchain = blockchain
	data.Address = address

	// Get floor price
	floor, err := getFloor(blockchain, address)
	if err != nil {
		log.Fatalf("Failed to retrieve %s floor data from DIA: %v", address, err)
		return oldFloor, err
	}
	data.Floor = floor.Value

	// Get MA of floor price
	floorMA, err := getFloorMA(blockchain, address)
	if err != nil {
		log.Fatalf("Failed to retrieve %s floorMA data from DIA: %v", address, err)
		return oldFloor, err
	}
	data.FloorMA = floorMA.Value

	// Check for deviation in floor price.
	newFloor := floor.Value

	if math.Abs(newFloor-oldFloor) > oldFloor*float64(deviationPermille)/1000 {
		log.Println("Entering deviation based update zone")
		err = updateNFTData(data, auth, contract, conn)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldFloor, err
		}
		return newFloor, nil
	}

	return newFloor, nil
}

func updateNFTData(data StarfishReturn, auth *bind.TransactOpts, contract *diaNFTOracleService.DIANFTOracle, conn *ethclient.Client) error {
	timestamp := uint64(time.Now().Unix())

	// Update floor
	symbol := data.Blockchain + "-" + data.Address
	var values []uint64
	values = append(values, uint64(data.Floor*100000000))
	values = append(values, uint64(data.FloorMA*100000000))
	values = append(values, []uint64{0, 0, 0}...)

	err := updateOracle(conn, contract, auth, symbol, values, timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateOracle(
	client *ethclient.Client,
	contract *diaNFTOracleService.DIANFTOracle,
	auth *bind.TransactOpts,
	key string,
	values []uint64,
	timestamp uint64) error {

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
	}, key, values[0], values[1], values[2], values[3], values[4], timestamp)
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
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

func getFloorMA(blockchain, address string) (FloorMA, error) {
	response, err := http.Get("https://api.diadata.org/v1/NFTFloorMA/" + blockchain + "/" + address)
	if err != nil {
		return FloorMA{}, err
	}
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return FloorMA{}, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return FloorMA{}, err
	}

	var resp FloorMA
	err = json.Unmarshal(contents, &resp)
	if err != nil {
		return FloorMA{}, err
	}

	return resp, err
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **diaNFTOracleService.DIANFTOracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = diaNFTOracleService.NewDIANFTOracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaNFTOracleService.DeployDIANFTOracle(auth, conn)
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
