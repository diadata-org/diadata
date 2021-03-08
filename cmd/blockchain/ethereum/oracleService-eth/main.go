package main

import (
	"bufio"
	"context"
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

	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/oracleService"
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
	var topCoins = flag.Int("topCoins", 15, "Number of coins to push with the oracle")
	var secretsFile = flag.String("secretsFile", "/run/secrets/oracle_keys", "File with wallet secrets")
	var blockchainNode = flag.String("blockchainNode", "http://159.69.120.42:8545/", "Node address for blockchain connection")
	var sleepSeconds = flag.Int("sleepSeconds", 120, "Number of seconds to sleep between calls")
	var frequencySeconds = flag.Int("frequencySeconds", 86400, "Number of seconds to sleep between full oracle runs")
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
	//conn, err := ethclient.Dial("https://rpc-mainnet.matic.network")
	//conn, err := ethclient.Dial("https://data-seed-prebsc-1-s1.binance.org:8545/")
	conn, err := ethclient.Dial(*blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the EVM client: %v", err)
	}

	auth, err := bind.NewTransactor(strings.NewReader(key), key_password)
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *oracleService.DiaOracle
	err = deployOrBindContract(*deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	periodicOracleUpdateHelper(topCoins, *sleepSeconds, auth, contract, conn)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(time.Duration(*frequencySeconds) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				periodicOracleUpdateHelper(topCoins, *sleepSeconds, auth, contract, conn)
			}
		}
	}()
	select {}
}

func periodicOracleUpdateHelper(topCoins *int, sleepSeconds int, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {

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
	rawBTCS, err := getSupplyFromDia("BTC")
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

	// --------------------------------------------------------
	// LENDING/BORROWING RATES
	// --------------------------------------------------------

	// Maker Rate
	rawMaker, err := getDefiRatesFromDia("MAKERDAO", "ETH-A")
	if err != nil {
		log.Fatalf("Failed to retrieve Makerdao data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawMaker, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Makerdao Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// -----------------------------------------------------------------------
	// LENDING PROTOCOL STATES
	// -----------------------------------------------------------------------

	// CREAM State Data
	rawCreamState, err := getDefiStateFromDia("CREAM")
	if err != nil {
		log.Fatalf("Failed to retrieve CREAM state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawCreamState, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update CREAM state Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// --------------------------------------------------------
	// EXCHANGE CHART POINTS
	// --------------------------------------------------------

	// Pancakeswap Chart Point
	rawPancake, err := getDEXFromDia("PanCakeSwap", "WBNB")
	if err != nil {
		log.Fatalf("Failed to retrieve PanCakeSwap from DIA: %v", err)
		return err
	}

	err = updateDEX(rawPancake, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update PanCakeSwap Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// --------------------------------------------------------
	// FARMING POOL RATES
	// --------------------------------------------------------

	// YFI WETH pool rate
	rawYFI, err := getFarmingPoolFromDia("yfi", "WETH")
	if err != nil {
		log.Fatalf("Failed to retrieve YFI pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawYFI, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update YFI Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	return nil
}

// ------------------------------------------------------------------------------------------------
// Update methods
// ------------------------------------------------------------------------------------------------

func updateCoin(coin models.Coin, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	symbol := strings.ToUpper(coin.Symbol)
	name := coin.Name
	supply := coin.CirculatingSupply
	price := coin.Price
	// Get 5 digits after the comma by multiplying price with 100000
	err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), int64(*supply))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}
	return nil
}

func updateTopCoins(topCoins []models.Coin, sleepSeconds int, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	for _, element := range topCoins {
		symbol := strings.ToUpper(element.Symbol)
		name := element.Name
		supply := element.CirculatingSupply
		price := element.Price
		// Get 5 digits after the comma by multiplying price with 100000
		err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), int64(*supply))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
		time.Sleep(time.Duration(sleepSeconds) * time.Second)
	}
	return nil
}

func updateDEX(dexData *models.Points, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	if len(dexData.DataPoints[0].Series) > 0 && len(dexData.DataPoints[0].Series[0].Values) > 0 {
		symbol := strings.ToUpper(dexData.DataPoints[0].Series[0].Values[0][3].(string))
		name := dexData.DataPoints[0].Series[0].Values[0][1].(string)
		supply := 0
		price := dexData.DataPoints[0].Series[0].Values[0][4].(float64)
		// Get 5 digits after the comma by multiplying price with 100000
		// Set supply to 0, as we don't have a supply for one exchange
		err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), int64(supply))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
	} else {
		err := updateOracle(conn, contract, auth, "", "", int64(0), int64(0))
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
	}
	return nil
}

func updateECBRate(ecbRate *models.CurrencyChange, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	symbol := strings.ToUpper(ecbRate.Symbol)
	name := strings.ToUpper(ecbRate.Symbol)
	price := ecbRate.Rate
	// Get 5 digits after the comma by multiplying price with 100000
	// Set supply to 0, as we don't have a supply for fiat currencies
	err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateDefiRate(defiRate *dia.DefiRate, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	symbol := strings.ToUpper(defiRate.Asset)
	name := strings.ToUpper(defiRate.Protocol)
	lendingRate := defiRate.LendingRate
	borrowingRate := defiRate.BorrowingRate
	// Get 5 digits after the comma by multiplying price with 100000
	err := updateOracle(conn, contract, auth, name, symbol, int64(lendingRate*100000), int64(borrowingRate*100000))
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateDefiState(defiState *dia.DefiProtocolState, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	symbol := ""
	name := strings.ToUpper(defiState.Protocol.Name) + "-state"
	price := defiState.TotalUSD
	// Get 5 digits after the comma by multiplying price with 100000
	err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

func updateForeignQuotation(foreignQuotation *models.ForeignQuotation, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	name := foreignQuotation.Source + "-" + foreignQuotation.Name
	symbol := foreignQuotation.Symbol
	price := foreignQuotation.Price
	err := updateOracle(conn, contract, auth, name, symbol, int64(price*100000), 0)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}

	return nil
}

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

func updateFarmingPool(poolData *models.FarmingPool, auth *bind.TransactOpts, contract *oracleService.DiaOracle, conn *ethclient.Client) error {
	protocolName := poolData.ProtocolName
	poolID := poolData.PoolID
	rate := poolData.Rate
	balance := poolData.Balance
	err := updateOracle(conn, contract, auth, protocolName, poolID, int64(rate*100000), int64(balance*100000))
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
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = oracleService.DeployDiaOracle(auth, conn)
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

// ------------------------------------------------------------------------------------------------
// Data retrieval from DIA API
// ------------------------------------------------------------------------------------------------

func getCoinDetailsFromDia(symbol string) (*models.Coin, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/symbol/" + symbol)
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.SymbolDetails
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b.Coin, nil
		}
		return nil, err
	}
}

func getToplistFromDia() (*models.Coins, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/coins")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Coins
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

// Getting EUR vs XXX rate
func getECBRatesFromDia(symbol string) (*models.CurrencyChange, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/coins")
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Coins
		err = b.UnmarshalBinary(contents)
		if err != nil {
			return nil, err
		}

		for _, change := range b.Change.USD {
			if strings.ToUpper(change.Symbol) == strings.ToUpper(symbol) {
				return &change, nil
			}
		}
		return nil, nil
	}
}

// Getting defi rate
func getDefiRatesFromDia(protocol string, symbol string) (*dia.DefiRate, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingRate/" + strings.ToUpper(protocol) + "/" + strings.ToUpper(symbol) + "/" + strconv.FormatInt(time.Now().Unix(), 10))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiRate
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

// Getting defi state
func getDefiStateFromDia(protocol string) (*dia.DefiProtocolState, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/defiLendingState/" + strings.ToUpper(protocol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b dia.DefiProtocolState
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func getDEXFromDia(dexname string, symbol string) (*models.Points, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/chartPoints/MAIR120/" + dexname + "/" + strings.ToUpper(symbol))
	if err != nil {
		return nil, err
	} else {
		defer response.Body.Close()

		if 200 != response.StatusCode {
			return nil, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
		}

		contents, err := ioutil.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}

		var b models.Points
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
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

func getQuotationFromDia(symbol string) (*models.Quotation, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/quotation/" + strings.ToUpper(symbol))
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

func getSupplyFromDia(symbol string) (*dia.Supply, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/supply/" + symbol)
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
	var supply dia.Supply
	err = supply.UnmarshalBinary(contents)
	if err != nil {
		return nil, err
	}
	return &supply, nil
}

func getFarmingPoolFromDia(protocol string, poolID string) (*models.FarmingPool, error) {
	response, err := http.Get(dia.BaseUrl + "v1/FarmingPoolData/" + strings.ToUpper(protocol) + "/" + poolID)

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

	var fp models.FarmingPool
	err = fp.UnmarshalBinary(contents)
	if err == nil {
		return &fp, nil
	}
	return nil, err
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
