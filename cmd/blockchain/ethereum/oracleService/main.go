package main

import(
	"os"
	"bufio"
	"fmt"
	"flag"
	"log"
	"math/big"
	"time"
	"strings"
	"net/http"
	"io/ioutil"
	"sort"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/http/restServer/diaApi"
	"github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/internal/pkg/blockchain-scrapers/blockchains/ethereum/oracleService"
)


func main() {
	/*
	 * Read in Oracle address
	 */
	var deployedContract = flag.String("deployedContract", "", "Address of the deployed oracle contract")
	var topCoins = flag.Int("topCoins", 15, "Number of coins to push with the oracle")
	flag.Parse()

	/*
	 * Read secrets for unlocking the ETH account
	 */
	var lines []string
	file, err := os.Open("/run/secrets/oracle_keys") // Read in key information
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
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/ec6581408f09414b8e4446067cd3ba08")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
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

	periodicOracleUpdateHelper(topCoins, auth, contract)
	/*
	 * Update Oracle periodically with top coins
	 */
	ticker := time.NewTicker(4 * time.Hour)
	go func() {
		for {
			select {
			case <- ticker.C:
				periodicOracleUpdateHelper(topCoins, auth, contract)
			}
		}
	}()
	select{}
}

func periodicOracleUpdateHelper(topCoins *int, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	rawCoins, err := getToplistFromDia()
	if err != nil {
		log.Fatalf("Failed to retrieve toplist from DIA: %v", err)
		return err
	}

	sort.Slice(rawCoins.Coins, func(i, j int) bool {
		return rawCoins.Coins[i].Price * *rawCoins.Coins[i].CirculatingSupply > rawCoins.Coins[j].Price * *rawCoins.Coins[j].CirculatingSupply
	})
	topCoinSlice := rawCoins.Coins[:*topCoins]
	// Search for NEU tokens
	neumarkData, err := getCoinDetailsFromDia("NEU")
	if err != nil {
		log.Fatalf("Failed to retrieve NEU token from DIA: %v", err)
		return err
	}
	topCoinSlice = append(topCoinSlice, *neumarkData)
	err = updateTopCoins(topCoinSlice, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}


	// Get EUR and CAD exchange rates
	eurRate, err := getECBRatesFromDia("EUR")
	if err != nil {
		log.Fatalf("Failed to retrieve currency %s from DIA: %v", "EUR", err)
		return err
	}
	cadRate, err := getECBRatesFromDia("CAD")
	if err != nil {
		log.Fatalf("Failed to retrieve currency %s from DIA: %v", "CAD", err)
		return err
	}

	ecbRates := []models.Quotation{*eurRate, *cadRate}

	err = updateECBRates(ecbRates, auth, contract)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return err
	}
	return nil
}

func updateTopCoins(topCoins []diaApi.Coin, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	for _, element := range topCoins {
		symbol := strings.ToUpper(element.Symbol)
		name := element.Name
		supply := element.CirculatingSupply
		price := element.Price
		// Get 5 digits after the comma by multiplying price with 100000
		err := updateOracle(contract, auth, name, symbol, int64(price * 100000), int64(*supply));
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
		time.Sleep(10 * time.Minute)
	}
	return nil
}

func updateECBRates(ecbRates []models.Quotation, auth *bind.TransactOpts, contract *oracleService.DiaOracle) error {
	for _, element := range ecbRates {
		symbol := strings.ToUpper(element.Symbol)
		name := element.Name
		price := element.Price
		// Get 5 digits after the comma by multiplying price with 100000
		// Set supply to 0, as we don't have a supply for fiat currencies
		err := updateOracle(contract, auth, name, symbol, int64(price * 100000), 0);
		if err != nil {
			log.Fatalf("Failed to update Oracle: %v", err)
			return err
		}
		time.Sleep(10 * time.Minute)
	}
	return nil
}

func deployOrBindContract(deployedContract string, conn *ethclient.Client, auth *bind.TransactOpts, contract **oracleService.DiaOracle) (error) {
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

func getCoinDetailsFromDia(symbol string) (*diaApi.Coin, error) {
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

		var b diaApi.SymbolDetails
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b.Coin, nil
		}
		return nil, err
	}
}

func getToplistFromDia() (*diaApi.Coins, error) {
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

		var b diaApi.Coins
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

// Getting EUR vs XXX rate
func getECBRatesFromDia(symbol string) (*models.Quotation, error) {
	response, err := http.Get(dia.BaseUrl + "/v1/quotation/" + strings.ToUpper(symbol))
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

		var b models.Quotation
		err = b.UnmarshalBinary(contents)
		if err == nil {
			return &b, nil
		}
		return nil, err
	}
}

func updateOracle(
	contract *oracleService.DiaOracle,
	auth *bind.TransactOpts,
	name string,
	symbol string,
	price int64,
	supply int64) error {
	// Write values to smart contract
	tx, err := contract.UpdateCoinInfo(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasLimit: 800725,
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
	return nil
}
