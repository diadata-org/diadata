package main

import (
	"context"
	"fmt"
	oracleService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/oracleService"
	"log"
	"math/big"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

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

	/*
	 * Setup connection to contract, deploy if necessary
	 */
	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the EVM client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *oracleService.DiaOracle
	err = deployOrBindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	err = periodicOracleUpdateHelper(sleepSeconds, auth, contract, conn)
	if err != nil {
		log.Fatalf("failed periodic update: %v", err)
	}
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	go func() {
		for range ticker.C {
			err = periodicOracleUpdateHelper(sleepSeconds, auth, contract, conn)
			if err != nil {
				log.Fatalf("failed periodic update: %v", err)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(sleepSeconds int, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {

	// --------------------------------------------------------
	// PRICE QUOTATIONS
	// --------------------------------------------------------

	// BTC quotation
	time.Sleep(time.Duration(sleepSeconds) * time.Second)
	rawBTCQ, err := getQuotationFromDia("BTC")
	if err != nil {
		log.Fatalf("Failed to retrieve BTC quotation data from DIA: %v", err)
		return err
	}
	rawBTCS, err := getSupplyFromDia("WBTC")
	if err != nil {
		log.Fatalf("Failed to retrieve BTC supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawBTCQ, rawBTCS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update BTC Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// BNB Quotation
	rawBNBQ, err := getQuotationFromDia("BNB")
	if err != nil {
		log.Fatalf("Failed to retrieve BNB quotation data from DIA: %v", err)
		return err
	}
	rawBNBS, err := getSupplyFromDia("BNB")
	if err != nil {
		log.Fatalf("Failed to retrieve BNB supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawBNBQ, rawBNBS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update BNB Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// ETH Quotation
	rawETHQ, err := getQuotationFromDia("ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve ETH quotation data from DIA: %v", err)
		return err
	}
	rawETHS, err := getSupplyFromDia("ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve ETH supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawETHQ, rawETHS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update ETH Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// USDT Quotation
	rawUSDTQ, err := getQuotationFromDia("USDT")
	if err != nil {
		log.Fatalf("Failed to retrieve USDT quotation data from DIA: %v", err)
		return err
	}
	rawUSDTS, err := getSupplyFromDia("USDT")
	if err != nil {
		log.Fatalf("Failed to retrieve USDT supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawUSDTQ, rawUSDTS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update USDT Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// XRP Quotation
	rawXRPQ, err := getQuotationFromDia("XRP")
	if err != nil {
		log.Fatalf("Failed to retrieve XRP quotation data from DIA: %v", err)
		return err
	}
	var rawXRPS dia.Supply
  rawXRPS.CirculatingSupply = 0.0
	err = updateQuotation(rawXRPQ, &rawXRPS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update XRP Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// USDC Quotation
	rawUSDCQ, err := getQuotationFromDia("USDC")
	if err != nil {
		log.Fatalf("Failed to retrieve USDC quotation data from DIA: %v", err)
		return err
	}
	rawUSDCS, err := getSupplyFromDia("USDC")
	if err != nil {
		log.Fatalf("Failed to retrieve USDC supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawUSDCQ, rawUSDCS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update USDC Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// DIA Quotation
	rawDIAQ, err := getQuotationFromDia("DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve DIA quotation data from DIA: %v", err)
		return err
	}
	rawDIAS, err := getSupplyFromDia("DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve DIA supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawDIAQ, rawDIAS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	return nil
}

// ------------------------------------------------------------------------------------------------
// Update methods
// ------------------------------------------------------------------------------------------------

func updateQuotation(quotation *models.Quotation, supply *dia.Supply, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	name := quotation.Name
	symbol := quotation.Symbol
	price := quotation.Price
	circSupply := supply.CirculatingSupply
	err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), int64(circSupply))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **oracleService.DiaOracle) error {
	var err error
	if deployedContract != "" {
		*contract, err = oracleService.NewDiaOracle(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
	} else {
		gasPrice, err := conn.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// Get 110% of the gas price
		fGas := new(big.Float).SetInt(gasPrice)
		fGas.Mul(fGas, big.NewFloat(1.1))
		gasPrice, _ = fGas.Int(nil)

		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		auth.GasPrice = gasPrice
		addr, tx, *contract, err = oracleService.DeployDiaOracle(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		fmt.Println(tx.GasPrice())
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n", tx.Hash())
		log.Printf("Tx Nonce: %d\n\n", tx.Nonce())
		time.Sleep(180000 * time.Millisecond)
	}
	return nil
}

// ------------------------------------------------------------------------------------------------
// Data retrieval from DIA API
// ------------------------------------------------------------------------------------------------

func getQuotationFromDia(symbol string) (*models.Quotation, error) {
	contents, statusCode, err := utils.GetRequest(dia.BaseUrl + "/v1/quotation/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("error on dia api with return code %d", statusCode)
	}

	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &quotation, nil
}

func getSupplyFromDia(symbol string) (*dia.Supply, error) {
	contents, statusCode, err := utils.GetRequest(dia.BaseUrl + "/v1/supply/" + symbol)
	if err != nil {
		return nil, err
	}
	if statusCode != 200 {
		return nil, fmt.Errorf("error on dia api with return code %d", statusCode)
	}

	var supply dia.Supply
	err = supply.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &supply, nil
}

func updateOracle(
	client *ethclient.Client,
	contract *oracleService.DiaOracle,
	auth *bind.TransactOpts,
	name string,
	symbol string,
	price int64,
	supply int64) error {

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
	tx, err := contract.UpdateCoinInfo(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 800725,
		GasPrice: gasPrice,
		//	Nonce: big.NewInt(time.Now().Unix()),
	}, name, symbol, big.NewInt(price), big.NewInt(supply), big.NewInt(time.Now().Unix()))
	// prices are with 5 digits after the comma
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("Symbol: %s\n", symbol)
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	log.Printf("Tx Nonce: %d\n", tx.Nonce())
	return nil
}
