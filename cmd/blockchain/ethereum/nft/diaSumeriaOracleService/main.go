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

type SumeriaReturn struct {
	Blockchain string
	Address    string
	Floor      float64
	FloorMA    float64
	Drawdown   Drawdown
	Exchange   string
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

type Drawdown struct {
	Drawdown  float64   `json:"Weekly_Drawdown"`
	Average   float64   `json:"Downday_Average"`
	Deviation float64   `json:"Downday_Deviation"`
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
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "1200"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}
	timeBasedUpdateSeconds, err := strconv.Atoi(utils.Getenv("TIME_BASED_UPDATE_SECONDS", "28800"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v", err)
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "3"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v", err)
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v", err)
	}

	addresses := []string{
		"0xBC4CA0EdA7647A8aB7C2061c2E118A18a936f13D", //BAYC
		"0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB", //Cryptopunks
		"0x60E4d786628Fea6478F785A6d7e704777c86a7c6", //Mutant Ape Yacht Club
		"0x34d85c9CDeB23FA97cb08333b511ac86E1C4E258", //Otherdeed for Otherside
		"0x23581767a106ae21c074b2276D25e5C3e136a68b", //Moonbirds
		"0x49cF6f5d44E70224e2E23fDcdd2C053F30aDA28B", //Clone X
		"0x8a90CAb2b38dba80c64b7734e58Ee1dB38B8992e", //Doodles
		"0x7Bd29408f11D2bFC23c34f18275bBf23bB716Bc7", //Meebits
		"0xa3AEe8BcE55BEeA1951EF834b99f3Ac60d1ABeeB", //Vee Friends
		"0xe785E82358879F061BC3dcAC6f0444462D4b5330", //World of Women
		"0x1A92f7381B9F03921564a437210bB9396471050C", //Cool Cats
		"0x59468516a8259058baD1cA5F8f4BFF190d30E066", //Invisible Friends
		"0x6dc6001535e15b9def7b0f6A20a2111dFA9454E2", //MetaHero Universe
		"0x79FCDEF22feeD20eDDacbB2587640e45491b757f", //Mfers
		"0x059EDD72Cd353dF5106D2B9cC5ab83a52287aC3a", //Chromie Squiggles by Snowfro
		"0xba30E5F9Bb24caa003E9f2f0497Ad287FDF95623", //Bored Ape Kennel Club
		"0xaaDc2D4261199ce24A4B0a57370c4FCf43BB60aa", //Damien Hirst - The Currency
		"0xED5AF388653567Af2F388E6224dC7C4b3241C544", //Azuki
		"0x9C8fF314C9Bc7F6e59A9d9225Fb22946427eDC03", //Nouns
		"0x08D7C0242953446436F34b4C78Fe9da38c73668d", //Proof Collective
		"0xd1258DB6Ac08eB0e625B75b371C023dA478E94A9", //DigiDaigaku
	}
	blockchains := []string{
		"Ethereum", //BAYC
		"Ethereum", //Cryptopunks
		"Ethereum", //Mutant Ape Y
		"Ethereum", //Otherdeed fo
		"Ethereum", //Moonbirds
		"Ethereum", //Clone X
		"Ethereum", //Doodles
		"Ethereum", //Meebits
		"Ethereum", //Vee Friends
		"Ethereum", //World of Wom
		"Ethereum", //Cool Cats
		"Ethereum", //Invisible Fr
		"Ethereum", //MetaHero Uni
		"Ethereum", //Mfers
		"Ethereum", //Chromie Squi
		"Ethereum", //Bored Ape Ke
		"Ethereum", //Damien Hirst
		"Ethereum", //Azuki
		"Ethereum", //Nouns
		"Ethereum", //Proof Collec
		"Ethereum", //DigiDaigaku
	}

	exchanges := []string{
		"LooksRare",
		"OpenSea",
		"X2Y2",
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
	 * exceeds threshold and update if so. Every @timeBasedUpdateSeconds update regardless of deviation.
	 */
	timeBasedUpdateTicker := time.NewTicker(time.Duration(timeBasedUpdateSeconds) * time.Second)
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	var timeBasedUpdate bool
	go func() {
		for {
			select {
			case <-ticker.C:
				timeBasedUpdate = false
			case <-timeBasedUpdateTicker.C:
				timeBasedUpdate = true
			}
			// Update all collections depending on @oldFloor and @timeBasedUpdate.
			for i, address := range addresses {
				for _, exchange := range exchanges {
					// Cryptopunk special case: Only traded on Cryptopunkmarket
					if address == "0xb47e3cd837dDF8e4c57F05d70Ab865de6e193BBB" {
						if exchange == "OpenSea" {
							exchange = "CryptopunkMarket"
						} else {
							continue
						}
					}
					blockchain := blockchains[i]
					oldFloor := oldFloors[address + "-" + exchange]
					log.Println("old price", oldFloor)
					newFloor, err := periodicOracleUpdateHelper(oldFloor, deviationPermille, timeBasedUpdate, auth, contract, conn, blockchain, address, exchange)
					oldFloors[address + "-" + exchange] = newFloor
					//Only sleep if an update was actually executed
					if newFloor != oldFloor {
						time.Sleep(time.Duration(sleepSeconds) * time.Second)
					}
					if err != nil {
						log.Println(err)
					}
				}
			}
		}
	}()
	select {}
}

// periodicOracleUpdateHelper updates a collection on either of the two conditions:
// 1. The difference of the (new) floor price and @oldFloor exceeds @deviationPermille.
// 2. @update is true.
func periodicOracleUpdateHelper(oldFloor float64, deviationPermille int, update bool, auth *bind.TransactOpts, contract *diaNFTOracleService.DIANFTOracle, conn *ethclient.Client, blockchain string, address string, exchange string) (float64, error) {
	var data SumeriaReturn
	data.Blockchain = blockchain
	data.Address = address
	data.Exchange = exchange

	// Get floor price
	floor, err := getFloorPerExchange(blockchain, address, exchange)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from DIA: %v", address, err)
		return oldFloor, err
	}
	data.Floor = floor.Value

	// Get MA of floor price
	floorMA, err := getFloorMA(blockchain, address)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from DIA: %v", address, err)
		return oldFloor, err
	}
	data.FloorMA = floorMA.Value

	// Get drawdown data
	drawdown, err := getDrawdown(blockchain, address)
	if err != nil {
		log.Printf("Failed to retrieve %s quotation data from DIA: %v", address, err)
		return oldFloor, err
	}
	data.Drawdown.Drawdown = drawdown.Drawdown
	data.Drawdown.Average = drawdown.Average
	data.Drawdown.Deviation = drawdown.Deviation

	// Check for deviation in floor price.
	// TO DO: I suggest to take MA of floor price here instead.
	newFloor := floor.Value

	if math.Abs(newFloor-oldFloor) > oldFloor*float64(deviationPermille)/1000 || update {
		log.Println("Entering deviation based update zone")
		err = updateNFTData(data, auth, contract, conn)
		if err != nil {
			log.Printf("Failed to update DIA Oracle: %v", err)
			return oldFloor, err
		}
		return newFloor, nil
	}

	return newFloor, nil
}

func updateNFTData(data SumeriaReturn, auth *bind.TransactOpts, contract *diaNFTOracleService.DIANFTOracle, conn *ethclient.Client) error {
	timestamp := uint64(time.Now().Unix())

	// Update floor
	symbol := data.Blockchain + "-" + data.Address + "-" + data.Exchange
	var values []uint64
	values = append(values, uint64(data.Floor*100000000))
	values = append(values, uint64(data.FloorMA*100000000))
	values = append(values, uint64(-data.Drawdown.Drawdown*100000000))
	values = append(values, uint64(-data.Drawdown.Average*100000000))
	values = append(values, uint64(data.Drawdown.Deviation*100000000))

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

func getFloorPerExchange(blockchain, address, exchange string) (Floor, error) {
	response, err := http.Get("https://api.diadata.org/v1/NFTFloor/" + blockchain + "/" + address + "?exchange=" + exchange)
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

func getDrawdown(blockchain, address string) (Drawdown, error) {
	response, err := http.Get("https://api.diadata.org/v1/NFTDownday/" + blockchain + "/" + address)
	if err != nil {
		log.Error(err)
		return Drawdown{}, err
	}
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return Drawdown{}, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return Drawdown{}, err
	}

	var resp Drawdown
	err = json.Unmarshal(contents, &resp)
	if err != nil {
		return Drawdown{}, err
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
