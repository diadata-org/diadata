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
	var chainId = flag.Int64("chainId", 1, "Chain-ID of the network to connect to")
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
	conn, err := ethclient.Dial(*blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the EVM client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(*chainId))
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
	rawXRPS, err := getSupplyFromDia("XRP")
	if err != nil {
		log.Fatalf("Failed to retrieve XRP supply data from DIA: %v", err)
		return err
	}
	err = updateQuotation(rawXRPQ, rawXRPS, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update XRP Oracle: %v", err)
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

	// CREAM Rates
	rawCream, err := getDefiRatesFromDia("CREAM", "UNI")
	if err != nil {
		log.Fatalf("Failed to retrieve CREAM data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawCream, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update CREAM Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// fortube Rates
	rawFortube, err := getDefiRatesFromDia("FORTUBE", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve forTube data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawFortube, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Fortube Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// nuo Rates
	rawNuo, err := getDefiRatesFromDia("NUO", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Nuo data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawNuo, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Nuo Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// bZx Rates
	rawBzx, err := getDefiRatesFromDia("BZX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve bZx data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawBzx, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update bZx Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Compound Rates
	rawCompound, err := getDefiRatesFromDia("COMPOUND", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Compound data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawCompound, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Compound Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// DYDX Rates
	rawDydx, err := getDefiRatesFromDia("DYDX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve DYDX data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawDydx, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DYDX Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Aave Rates
	rawAave, err := getDefiRatesFromDia("AAVE", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Aave data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawAave, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Aave Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Aave Rates
	rawBitfinex, err := getDefiRatesFromDia("BITFINEX", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Bitfinex data from DIA: %v", err)
		return err
	}
	err = updateDefiRate(rawBitfinex, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Bitfinex Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// -----------------------------------------------------------------------
	// LENDING PROTOCOL STATES
	// -----------------------------------------------------------------------

	// MAKERDAO State Data
	rawMakerState, err := getDefiStateFromDia("MAKERDAO")
	if err != nil {
		log.Fatalf("Failed to retrieve Maker state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawMakerState, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Maker state Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

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

	// DYDX State Data
	rawDydxState, err := getDefiStateFromDia("DYDX")
	if err != nil {
		log.Fatalf("Failed to retrieve DYDX state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawDydxState, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DYDX state Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Compound State Data
	rawCompoundState, err := getDefiStateFromDia("COMPOUND")
	if err != nil {
		log.Fatalf("Failed to retrieve Compound state data from DIA: %v", err)
		return err
	}
	err = updateDefiState(rawCompoundState, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Compound state Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// --------------------------------------------------------
	// EXCHANGE CHART POINTS
	// --------------------------------------------------------

	// ECB Chart Point
	/*rawECB, err := getECBRatesFromDia("EUR")
	if err != nil {
		log.Fatalf("Failed to retrieve ECB from DIA: %v", err)
		return err
	}
	err = updateECBRate(rawECB, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update ECB Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)*/

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

	// CREX24 Chart Point
	rawCrex24, err := getDEXFromDia("CREX24", "CREX")
	if err != nil {
		log.Fatalf("Failed to retrieve CREX24 from DIA: %v", err)
		return err
	}

	err = updateDEX(rawCrex24, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update CREX24 Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Bitmax CEX Chart Point
	rawBitmax, err := getDEXFromDia("Bitmax", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Bitmax from DIA: %v", err)
		return err
	}

	err = updateDEX(rawBitmax, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Bitmax Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// // Maker DEX Chart Point
	// rawMaker, err := getDEXFromDia("Maker", "ETH")
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve Maker from DIA: %v", err)
	// 	return err
	// }

	// err = updateDEX(rawMaker, auth, contract)
	// if err != nil {
	// 	log.Fatalf("Failed to update Maker Oracle: %v", err)
	// 	return err
	// }
	// time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Curvefi DEX Chart Point
	rawCurvefi, err := getDEXFromDia("Curvefi", "DAI")
	if err != nil {
		log.Fatalf("Failed to retrieve Curvefi from DIA: %v", err)
		return err
	}

	err = updateDEX(rawCurvefi, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Curvefi Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Gnosis DEX Chart Point
	rawGnosis, err := getDEXFromDia("Gnosis", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Gnosis from DIA: %v", err)
		return err
	}

	err = updateDEX(rawGnosis, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Gnosis Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Uniswap Chart Point
	rawUniswap, err := getDEXFromDia("Uniswap", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Uniswap from DIA: %v", err)
		return err
	}

	err = updateDEX(rawUniswap, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Uniswap Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// // Loopring Chart Point
	// rawLoopring, err := getDEXFromDia("Loopring", "ETH")
	// if err != nil {
	// 	log.Fatalf("Failed to retrieve Loopring from DIA: %v", err)
	// 	return err
	// }

	// err = updateDEX(rawLoopring, auth, contract)
	// if err != nil {
	// 	log.Fatalf("Failed to update Loopring Oracle: %v", err)
	// 	return err
	// }
	// time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Bancor Chart Point
	rawBancor, err := getDEXFromDia("Bancor", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Bancor from DIA: %v", err)
		return err
	}

	err = updateDEX(rawBancor, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Bancor Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// 0x Chart Point
	raw0x, err := getDEXFromDia("0x", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve 0x from DIA: %v", err)
		return err
	}

	err = updateDEX(raw0x, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update 0x Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Kyber Chart Point
	rawKyber, err := getDEXFromDia("Kyber", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Kyber from DIA: %v", err)
		return err
	}

	err = updateDEX(rawKyber, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Kyber Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Sushi Chart Point
	rawSushi, err := getDEXFromDia("SushiSwap", "ETH")
	if err != nil {
		log.Fatalf("Failed to retrieve Sushi from DIA: %v", err)
		return err
	}

	err = updateDEX(rawSushi, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Sushi Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// STEX Chart Point
	rawSTEX, err := getDEXFromDia("STEX", "PLEX")
	if err != nil {
		log.Fatalf("Failed to retrieve STEX from DIA: %v", err)
		return err
	}

	err = updateDEX(rawSTEX, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update STEX Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// DIA token
	diaToken, err := getCoinDetailsFromDia("DIA")
	if err != nil {
		log.Fatalf("Failed to retrieve token DIA from DIA: %v", err)
		return err
	}
	err = updateCoin(*diaToken, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update DIA Token Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// --------------------------------------------------------
	// FARMING POOL RATES
	// --------------------------------------------------------

	// Balancer WETH/WBTC pool rate
	/*rawBalancer, err := getFarmingPoolFromDia("Balancer", "0x1efF8aF5D577060BA4ac8A29A13525bb0Ee2A3D5")
	if err != nil {
		log.Fatalf("Failed to retrieve Balancer pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawBalancer, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Balancer Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// CVAULT WETH pool rate
	rawCvault, err := getFarmingPoolFromDia("Cvault", "0")
	if err != nil {
		log.Fatalf("Failed to retrieve CVAULT pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawCvault, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update CVAULT Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)*/

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

	// SYNTHETIX sETH total debt
	rawSYNTHETIX, err := getFarmingPoolFromDia("Synthetix", "0xD0DC005d31C2979CC0d38718e23c82D1A50004C0")
	if err != nil {
		log.Fatalf("Failed to retrieve SYNTHETIX pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawSYNTHETIX, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update SYNTHETIX Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// LOOPRING total reward
	rawLRC, err := getFarmingPoolFromDia("Loopring", "LRC")
	if err != nil {
		log.Fatalf("Failed to retrieve LOOPRING pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawLRC, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update LOOPRING Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// CURVEFI virtual price
	rawCURVEFI, err := getFarmingPoolFromDia("Curvefi", "3")
	if err != nil {
		log.Fatalf("Failed to retrieve CURVEFI pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawCURVEFI, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update CURVEFI Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// BARNBRIDGE total reward
	rawBARNBRIDGE, err := getFarmingPoolFromDia("BARNBRIDGE", "STABLECOIN")
	if err != nil {
		log.Fatalf("Failed to retrieve BARNBRIDGE pool from DIA: %v", err)
		return err
	}

	err = updateFarmingPool(rawBARNBRIDGE, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update BARNBRIDGE Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)

	// Top 15 coins
	/*rawCoins, err := getToplistFromDia()
	if err != nil {
		log.Fatalf("Failed to retrieve toplist from DIA: %v", err)
		return err
	}

	cleanedCoins := []models.Coin{}

	for key := range rawCoins.Coins {
		if rawCoins.Coins[key].CirculatingSupply != nil {
			cleanedCoins = append(cleanedCoins, rawCoins.Coins[key])
		}
	}
	sort.Slice(cleanedCoins, func(i, j int) bool {
		return cleanedCoins[i].Price**cleanedCoins[i].CirculatingSupply > cleanedCoins[j].Price**cleanedCoins[j].CirculatingSupply
	})
	topCoinSlice := cleanedCoins[:*topCoins]

	err = updateTopCoins(topCoinSlice, auth, contract, conn)
	if err != nil {
		log.Fatalf("Failed to update Coins Oracle: %v", err)
		return err
	}
	time.Sleep(time.Duration(sleepSeconds) * time.Second)*/

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
