package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaOracleServiceV2 "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var diaBaseUrl string

var reethcontractabi = `[{"inputs":[{"internalType":"address","name":"_real","type":"address"},{"internalType":"address payable","name":"_vault","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"inputs":[],"name":"Minter__NotVault","type":"error"},{"inputs":[],"name":"Minter__ZeroAddress","type":"error"},{"inputs":[{"internalType":"address","name":"_from","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"burn","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"getTokenPrice","outputs":[{"internalType":"uint256","name":"price","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_to","type":"address"},{"internalType":"uint256","name":"_amount","type":"uint256"}],"name":"mint","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"real","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"address","name":"_vault","type":"address"}],"name":"setNewVault","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[],"name":"vault","outputs":[{"internalType":"address payable","name":"","type":"address"}],"stateMutability":"view","type":"function"}]`
var reethcontract *bind.BoundContract
var parsedAbi abi.ABI
var deployedContract string

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")

	deployedContract = utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	dataNode := utils.Getenv("DATA_NODE", "")

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

	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")

	baseaddress := "0x0000000000000000000000000000000000000000"
	baseblockchain := "Ethereum"

	oldPrice := 0.0

	parsedAbi, err = abi.JSON(strings.NewReader(reethcontractabi))
	if err != nil {
		log.Fatalf("Failed to parse parsedAbi: %v")
	}

	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		log.Fatalf("Failed to connect to the data client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Println(err)
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaOracleServiceV2.DIAOracleV2

	err = deployOrBindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	reethconn, err := NewMain(common.HexToAddress("0x655756824385F8903AC8cFDa17B656cc26f7C7da"), dataConn)
	if err != nil {
		log.Fatalf("Failed to get data contract conn: %v", err)
	}

	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("old price", oldPrice)
				oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, reethconn, conn, baseblockchain, baseaddress)
				if err != nil {
					log.Println(err)
				}
			}
		}
	}()

	select {}
}

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, reethconn *Main, conn *ethclient.Client, baseblockchain string, baseaddress string) (float64, error) {
	// Get asset quotation
	rawBaseQ, err := getAssetQuotationFromDia(baseblockchain, baseaddress)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", baseaddress, err)
		return oldPrice, err
	}
	rawBaseQ.Name = rawBaseQ.Symbol

	// Get reeth token price
	rawReethTokenRatio, err := reethconn.GetTokenPrice(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get reETH token ratio: %v", err)
		return oldPrice, err
	}

	floatReethTokenRatio := new(big.Float).SetInt(rawReethTokenRatio)
	reethNormalizationFactor := big.NewFloat(1e18)
	normalizedReethTokenRatio := new(big.Float).Quo(floatReethTokenRatio, reethNormalizationFactor)
	reethTokenRatio, _ := normalizedReethTokenRatio.Float64()
	log.Printf("reETH token ratio: %f", reethTokenRatio)

	newPrice := reethTokenRatio * rawBaseQ.Price
	log.Printf("newPrice: %f", newPrice)

	// Check for deviation
	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) ||
	  (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")

		err = updatePair(rawBaseQ, auth, contract, conn, newPrice)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, err
		}
		return newPrice, nil
	}

	return oldPrice, nil
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

func updatePair(baseQuotation *models.Quotation, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, reethTokenRatio float64) error {
	symbol := "re" + baseQuotation.Symbol
	price := reethTokenRatio
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp)
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
	log.Println("gasPrice:", gasPrice)
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)

	tx, err := contract.SetValue(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, key, big.NewInt(value), big.NewInt(timestamp))

	if err != nil {
		log.Println(err)
	}
	log.Printf("from: %s\n", auth.From.Hex())
	log.Printf("key: %s\n", key)
	log.Printf("value: %s\n", value)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getAssetQuotationFromDia(blockchain, address string) (*models.Quotation, error) {
	response, err := http.Get(diaBaseUrl + "/v1/assetQuotation/" + blockchain + "/" + address)
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
