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

	"github.com/diadata-org/diadata/config/nftContracts/erc20"
	"github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleServiceV2"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

var damfinanceoracleabi = `[{"inputs":[{"internalType":"address","name":"_d2oTokenAddress","type":"address"}],"stateMutability":"nonpayable","type":"constructor"},{"anonymous":false,"inputs":[{"indexed":false,"internalType":"string","name":"key","type":"string"},{"indexed":false,"internalType":"uint128","name":"value","type":"uint128"},{"indexed":false,"internalType":"uint128","name":"timestamp","type":"uint128"}],"name":"OracleUpdate","type":"event"},{"anonymous":false,"inputs":[{"indexed":true,"internalType":"address","name":"previousOwner","type":"address"},{"indexed":true,"internalType":"address","name":"newOwner","type":"address"}],"name":"OwnershipTransferred","type":"event"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"collateralRatio","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"","type":"string"}],"name":"d20TotalSupply","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"d2oTokenAddress","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"getPrice","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"}],"name":"getValue","outputs":[{"internalType":"uint128","name":"","type":"uint128"},{"internalType":"uint128","name":"","type":"uint128"}],"stateMutability":"view","type":"function"},{"inputs":[],"name":"owner","outputs":[{"internalType":"address","name":"","type":"address"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"},{"internalType":"uint128","name":"value","type":"uint128"},{"internalType":"uint128","name":"timestamp","type":"uint128"}],"name":"setValue","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"key","type":"string"},{"internalType":"uint128","name":"value","type":"uint128"},{"internalType":"uint128","name":"timestamp","type":"uint128"},{"internalType":"uint256","name":"totalSupply","type":"uint256"},{"internalType":"uint256","name":"balance","type":"uint256"}],"name":"setValues","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"address","name":"newOwner","type":"address"}],"name":"transferOwnership","outputs":[],"stateMutability":"nonpayable","type":"function"},{"inputs":[{"internalType":"string","name":"","type":"string"}],"name":"usdcBalance","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"},{"inputs":[{"internalType":"string","name":"","type":"string"}],"name":"values","outputs":[{"internalType":"uint256","name":"","type":"uint256"}],"stateMutability":"view","type":"function"}]`
var damcontract *bind.BoundContract
var parsedAbi abi.ABI
var deployedContract string

var diaBaseUrl string

/*

1) DATA_RPC := rpc url for another chain from which
2) Get Balance of USDC from contract 0xB1fbcD7415F9177F5EBD3d9700eD5F15B476a5Fe
3) Get total supply of 0x2FdA8c6783Aa36BeD645baD28a4cDC8769dCD252 token
4) Supplu usdc price, balance and totalSupply to contract

*/

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")

	deployedContract = utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	dataNode := utils.Getenv("DATA_NODE", "")

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
	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")

	baseaddress := "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48"
	baseblockchain := "Ethereum"

	oldPrice := 0.0
	oldBalance := big.NewInt(0)
	oldTotalSupply := big.NewInt(0)

	parsedAbi, err = abi.JSON(strings.NewReader(damfinanceoracleabi))
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

	damfinancecontract := bind.NewBoundContract(common.HexToAddress(deployedContract), parsedAbi, nil, conn, nil)

	err = deployOrBindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	usdcconn, err := erc20.NewERC20(common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"), dataConn)
	if err != nil {
		log.Fatalf("Failed to get erc20 conn: %v", err)
	}

	d20tokenconn, err := erc20.NewERC20(common.HexToAddress("0x2FdA8c6783Aa36BeD645baD28a4cDC8769dCD252"), dataConn)
	if err != nil {
		log.Fatalf("Failed to get erc20 conn: %v", err)
	}

	totalSupply, _ := d20tokenconn.TotalSupply(&bind.CallOpts{})

	fmt.Println("-----------", totalSupply)

	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("old price", oldPrice)
				log.Println("old balance", oldBalance)
				log.Println("old totalSupply", oldTotalSupply)
				oldPrice, oldBalance, oldTotalSupply, err = periodicOracleUpdateHelper(oldPrice, oldBalance, oldTotalSupply, deviationPermille, auth, damfinancecontract, usdcconn, d20tokenconn, conn, baseblockchain, baseaddress)
				if err != nil {
					log.Println(err)
				}
				time.Sleep(time.Duration(sleepSeconds) * time.Second)
			}
		}
	}()

	select {}
}

func periodicOracleUpdateHelper(oldPrice float64, oldBalance, oldTotalSupply *big.Int, deviationPermille int, auth *bind.TransactOpts, contract *bind.BoundContract, usdcconn, d20tokenconn *erc20.ERC20, conn *ethclient.Client, baseblockchain string, baseaddress string) (float64, *big.Int, *big.Int, error) {
	var usdcbalance, totalSupply *big.Int

	// Get asset quotation
	rawBaseQ, err := getAssetQuotationFromDia(baseblockchain, baseaddress)
	if err != nil {
		log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", baseaddress, err)
		return oldPrice, oldBalance, oldTotalSupply, err
	}
	rawBaseQ.Name = rawBaseQ.Symbol

	// Get USDC balance
	usdcbalance, err = usdcconn.BalanceOf(&bind.CallOpts{}, common.HexToAddress("0xB1fbcD7415F9177F5EBD3d9700eD5F15B476a5Fe"))
	if err != nil {
		log.Fatalf("Failed to get balance of usdc: %v", err)
		return oldPrice, oldBalance, oldTotalSupply, err
	}

	// Get total supply
	totalSupply, err = d20tokenconn.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Fatalf("Failed to get totalSupply of d20: %v", err)
		return oldPrice, oldBalance, oldTotalSupply, err
	}

	newPrice := rawBaseQ.Price
	newBalance := usdcbalance
	newTotalSupply := totalSupply

	// Check for deviation
	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) ||
	  (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) ||
		newBalance.Cmp(oldBalance) != 0 ||
		newTotalSupply.Cmp(oldTotalSupply) != 0 {
		log.Println("Entering deviation based update zone")

		err = updatePair(rawBaseQ, auth, contract, conn, usdcbalance, totalSupply)
		if err != nil {
			log.Fatalf("Failed to update DIA Oracle: %v", err)
			return oldPrice, oldBalance, oldTotalSupply, err
		}
		return newPrice, newBalance, newTotalSupply, nil
	}

	return oldPrice, oldBalance, oldTotalSupply, nil
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

func updatePair(baseQuotation *models.Quotation, auth *bind.TransactOpts, contract *bind.BoundContract, conn *ethclient.Client, usdcbalance,
	totalsupply *big.Int) error {
	symbol := baseQuotation.Symbol
	price := baseQuotation.Price
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), usdcbalance, totalsupply, timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateOracle(
	client *ethclient.Client,
	contract *bind.BoundContract,
	auth *bind.TransactOpts,
	key string,
	value int64,
	usdcbalance,
	totalsupply *big.Int,
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

	tx, err := contract.Transact(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, "setValues", key, big.NewInt(value), big.NewInt(timestamp), totalsupply, usdcbalance)

	if err != nil {
		log.Println(err)
	}
	log.Printf("from: %s\n", auth.From.Hex())
	log.Printf("key: %s\n", key)
	log.Printf("totalsupply: %s\n", totalsupply)
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
