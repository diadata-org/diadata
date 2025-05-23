package main

import (
	"context"
	"errors"
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
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	gql "github.com/machinebox/graphql"
	"github.com/tidwall/gjson"
)

var diaBaseUrl string

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	coingeckoApiKey := utils.Getenv("COINGECKO_API_KEY", "")
	diaBaseUrl = utils.Getenv("DIA_BASE_URL", "https://api.diadata.org")
	sleepSeconds, err := strconv.Atoi(utils.Getenv("SLEEP_SECONDS", "10"))
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

	addresses := []string{
		"0xdAC17F958D2ee523a2206206994597C13D831ec7", //USDT
		"0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48", //USDC
		"0x0000000000000000000000000000000000000000", //ETH
		"0x2260FAC5E5542a773Aa44fBCfeDf7C193bc2C599", //WBTC
		"0x0000000000000000000000000000000000000000", //ASTR
		"0x0000000000000000000000000000000000000000", //SDN
		"0x6B175474E89094C44Da98b954EedeAC495271d0F", //DAI
		"0xe9e7CEA3DedcA5984780Bafc599bD69ADd087D56", //BUSD
		"0x0000000000000000000000000000000000000000", //BNB
		"0x0000000000000000000000000000000000001010", //MATIC
		"0x0000000000000000000000000000000000000000", //DOT
		"0xE511ED88575C57767BAfb72BfD10775413E3F2b0", //nASTR
		"vDOT",                                       //vDOT
		"vASTR",                                      //vASTR
	}
	blockchains := []string{
		"Ethereum", //USDT
		"Ethereum", //USDC
		"Ethereum", //ETH
		"Ethereum", //WBTC
		"Astar", //ASTR
		"Shiden", //SDN
		"Ethereum", //DAI
		"BinanceSmartChain", //BUSD
		"BinanceSmartChain", //BNB
		"Polygon", //MATIC
		"Polkadot", //DOT
		"Astar", //nASTR
		"Bifrost", //vDOT
		"Bifrost", //vASTR
	}
	cgNames := []string{
		"tether",
		"usd-coin",
		"ethereum",
		"wrapped-bitcoin",
		"astar",
		"shiden",
		"dai",
		"binance-usd",
		"binancecoin",
		"matic-network",
		"polkadot",
		"liquid-astr",
		"voucher-dot",
		"bifrost-voucher-astr",
	}
	oldPrices := make(map[int]float64)

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

	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	/*
	 * Update Oracle periodically with top coins
	 */
	go func() {
		for {
			select {
			case <-ticker.C:
				for i, address := range addresses {
					blockchain := blockchains[i]
					oldPrice := oldPrices[i]
					coingeckoName := cgNames[i]
					log.Println("old price", oldPrice)
					oldPrice, err = periodicOracleUpdateHelper(oldPrice, deviationPermille, auth, contract, conn, blockchain, address, chainId, coingeckoName, coingeckoApiKey)
					oldPrices[i] = oldPrice
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

func periodicOracleUpdateHelper(oldPrice float64, deviationPermille int, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, blockchain string, address string, chainId int64, coingeckoName, coingeckoApiKey string) (float64, error) {

	// Empty quotation for our request
	var rawQ *models.Quotation
	rawQ = new(models.Quotation)
	var err error
	// special case: vDOT and vASTR: get fair price
	if (address == "vASTR" || address == "vDOT") && blockchain == "Bifrost" {
		fairPrice, err := getFairPriceFromDia(address)
		if err != nil {
			log.Fatalf("Failed to retrieve %s fair price data from DIA: %v", address, err)
			return oldPrice, err
		}
		rawQ.Price = fairPrice
		rawQ.Symbol = address
	} else {
	// Get quotation for token and update Oracle
		rawQ, err = getAssetQuotationFromDia(blockchain, address)
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
			return oldPrice, err
		}
	}
	rawQ.Name = rawQ.Symbol

	// Special case Polygon: Use old symbol MATIC
	if rawQ.Symbol == "POL" {
		rawQ.Name = "MATIC"
		rawQ.Symbol= "MATIC"
	}

	// Check for deviation
	newPrice := rawQ.Price

	// stablecoin gql DAI and USDC
	if address == "0x6B175474E89094C44Da98b954EedeAC495271d0F" || address == "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48" {
		log.Printf("Entered graphql mode")
		rawPrice, _, err := getGraphqlAssetQuotationFromDia(blockchain, address, 120)
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", address, err)
			return oldPrice, err
		}
		newPrice = rawPrice
	}


	if (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
		log.Println("Entering deviation based update zone")
		rawQ.Price = newPrice
		
		// only perform CG check for asset that is not AUSD
		// For vDOT and vASTR do the "deviation from previous update" check
		if address != "Token:AUSD" && address != "vDOT" && address != "vASTR" {
			// check coingecko before sending out an update transaction
			cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
			if err != nil {
				return oldPrice, err
			}
			if cgPrice == 0.0 {
				log.Printf("Error! Coingecko API returned price 0.0.")
				return oldPrice, nil
			}
			if (math.Abs(cgPrice - rawQ.Price) / cgPrice) > 0.2 {
				// Error case, stop transaction from happening
				log.Printf("Error! Price %f for asset %s-%s out of coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)
				return oldPrice, nil
			}
			log.Printf("Price %f for asset %s-%s in coingecko range %f.", rawQ.Price, blockchain, address, cgPrice)
	  } else if oldPrice != 0.0 && (address == "vDOT" || address == "vASTR") {
	  	// Check if newPrice is >= 20% away from oldPrice
	  	if (newPrice > oldPrice * 1.2 || newPrice < oldPrice * 0.8) {
	  		// Get coingecko price of the asset and use this to update the oracle
	  		log.Printf("Error! Asset %s was out of previous deviation check, using Coingecko price for the update", address)
				cgPrice, err := getCoingeckoPrice(coingeckoName, coingeckoApiKey)
				if err != nil {
					return oldPrice, err
				}
				if cgPrice == 0.0 {
					log.Printf("Error! Coingecko API returned price 0.0.")
					return oldPrice, nil
				}
				rawQ.Price = cgPrice
			}
		}
		
		err = updateQuotation(rawQ, auth, contract, conn, chainId)
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

func updateQuotation(quotation *models.Quotation, auth *bind.TransactOpts, contract *diaOracleServiceV2.DIAOracleV2, conn *ethclient.Client, chainId int64) error {
	symbol := quotation.Symbol + "/USD"
	price := quotation.Price
	timestamp := time.Now().Unix()
	err := updateOracle(conn, contract, auth, symbol, int64(price*100000000), timestamp, chainId)
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
	timestamp int64,
	chainId int64) error {

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if chainId == 592 || chainId == 336 { //Astar and Shiden need their own gas oracle
		gasPrice, err = getGasSuggestion(chainId)
	}
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
	}, key, big.NewInt(value), big.NewInt(timestamp))
	if err != nil {
		return err
	}
	fmt.Println(tx.GasPrice())
	log.Printf("key: %s\n", key)
	log.Printf("nonce: %d\n", tx.Nonce())
	log.Printf("gas price: %d\n", tx.GasPrice())
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

func getGraphqlAssetQuotationFromDia(blockchain, address string, blockDuration int) (float64, string, error) {
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-blockDuration * 2) * time.Second)
	type Response struct {
		GetChart []struct {
			Name   string    `json:"Name"`
			Symbol string    `json:"Symbol"`
			Time   time.Time `json:"Time"`
			Value  float64   `json:"Value"`
		} `json:"GetChart"`
	}
	client := gql.NewClient(diaBaseUrl + "/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetChart(
		 	filter: "mair", 
			Symbol:"Asset",
			BlockDurationSeconds: ` + strconv.Itoa(blockDuration) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(blockDuration) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			Address: "` + address + `", 
			BlockChain: "` + blockchain + `") {
				Name
				Symbol
				Time
				Value
	  	}
		}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, "", err
	}
	if len(r.GetChart) == 0 {
		return 0.0, "", errors.New("no results")
	}
	return r.GetChart[len(r.GetChart)-1].Value, r.GetChart[len(r.GetChart)-1].Symbol, nil
}

func getGasSuggestion(chainId int64) (*big.Int, error) {
	chainName := "shiden"
	if chainId == 592 {
		chainName = "astar"
	}
	response, err := http.Get("https://gas.astar.network/api/gasnow?network=" + chainName)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return nil, fmt.Errorf("Error on astar gasstation with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}
	
	gasSuggestion := gjson.Get(string(contents), "data.fast")
	retval := big.NewInt(gasSuggestion.Int())

	return retval, nil
}

func getCoingeckoPrice(assetName, coingeckoApiKey string) (float64, error) {
	url := "https://pro-api.coingecko.com/api/v3/simple/price?ids=" + assetName + "&vs_currencies=usd&x_cg_pro_api_key=" + coingeckoApiKey
	response, err := http.Get(url)
	if err != nil {
		return 0.0, err
	}
	
	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on coingecko API call with return code %d", response.StatusCode)
	}

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	price := gjson.Get(string(contents), assetName + ".usd").Float()
	return price, nil
}

func getFairPriceFromDia(symbol string) (float64, error) {
	response, err := http.Get(diaBaseUrl + "/xlsd/bifrost/" + symbol)
	if err != nil {
		return 0.0, err
	}

	defer response.Body.Close()
	if 200 != response.StatusCode {
		return 0.0, fmt.Errorf("Error on dia api with return code %d", response.StatusCode)
	}
	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0.0, err
	}
	fairPrice := gjson.Get(string(contents), "FairPrice").Float()

	return fairPrice, nil
}
