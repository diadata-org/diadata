package main

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"strconv"
	"strings"
	"time"

	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	gql "github.com/machinebox/graphql"
)

type Asset struct {
	blockchain    string
	address       string
	symbol        string
	gqlParams     GqlParameters
}

type GqlParameters struct {
	FeedSelection []struct {
		Address            string  `json:"Address"`
		Blockchain         string  `json:"Blockchain"`
		LiquidityThreshold float64 `json:"LiquidityThreshold"`
		Exchangepairs      []struct {
			Exchange string   `json:"Exchange"`
			Pairs    []string `json:"Pairs"`
		} `json:"Exchangepairs"`
	} `json:"FeedSelection"`
}

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "0x0000000000000000000000000000000000000000")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	frequencySeconds, err := strconv.Atoi(utils.Getenv("FREQUENCY_SECONDS", "120"))
	if err != nil {
		log.Fatalf("Failed to parse frequencySeconds: %v")
	}
	mandatoryFrequencySeconds, err := strconv.Atoi(utils.Getenv("MANDATORY_FREQUENCY_SECONDS", "0"))
	if err != nil {
		log.Fatalf("Failed to parse mandatoryFrequencySeconds: %v")
	}
	chainId, err := strconv.ParseInt(utils.Getenv("CHAIN_ID", "1"), 10, 64)
	if err != nil {
		log.Fatalf("Failed to parse chainId: %v")
	}
	deviationPermille, err := strconv.Atoi(utils.Getenv("DEVIATION_PERMILLE", "10"))
	if err != nil {
		log.Fatalf("Failed to parse deviationPermille: %v")
	}
	gqlWindowSize, err := strconv.Atoi(utils.Getenv("GQL_WINDOW_SIZE", "120"))
	if err != nil {
		log.Fatalf("Failed to parse gqlWindowSize: %v")
	}
	gqlMethodology := utils.Getenv("GQL_METHODOLOGY", "vwap")
	assetsStr := utils.Getenv("ASSETS", "")
	gqlAssetsStr := utils.Getenv("GQL_ASSETS", "")

	assets := []Asset{}
	useGql := false
	var assetsToParse string

	// Either Assets or GQL Assets must be non-empty
	if gqlAssetsStr != "" && assetsStr == "" {
		useGql = true
		assetsToParse = gqlAssetsStr
	} else if gqlAssetsStr == "" && assetsStr != "" {
		useGql = false
		assetsToParse = assetsStr
	} else {
		log.Fatalf("Use either ASSETS or GQL_ASSETS env variable")
	}
	assetsParsed := strings.Split(assetsToParse, ";")

	for _, asset := range assetsParsed {
		var currAsset Asset

		// parse asset from env
		entries := strings.Split(asset, "-")
		currAsset.blockchain = strings.TrimSpace(entries[0])
		currAsset.address = strings.TrimSpace(entries[1])
		currAsset.symbol = strings.TrimSpace(entries[2])

		// Find out is there are additional GQL parameters for this asset
		if len(entries) > 3 {
			// Join the rest of the line together, because the previous split might have affected substrings of the parameters
			gqlFeedSelectionQuery := strings.Join(entries[3:], "-")
			var currGqlParams GqlParameters
			if useGql && gqlFeedSelectionQuery != "" {
				err := json.Unmarshal([]byte(gqlFeedSelectionQuery), &currGqlParams)
				if err != nil {
					log.Println("Error while parsing GQL asset string:", err)
				}
			}
			currAsset.gqlParams = currGqlParams
		}

		assets = append(assets, currAsset)
	}

	publishedPrices := make(map[string]float64)

	/*
	 * Setup connection to contract
	 */

	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService
	err = bindContract(deployedContract, conn, auth, &contract)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind contract: %v", err)
	}

	/*
	 * Update Oracle periodically
	 */
	var mandatoryticker *time.Ticker
	ticker := time.NewTicker(time.Duration(frequencySeconds) * time.Second)
	if mandatoryFrequencySeconds > 0 {
		mandatoryticker = time.NewTicker(time.Duration(mandatoryFrequencySeconds) * time.Second)
	}
	go func() {
		for {
			if mandatoryFrequencySeconds > 0 {
				select {
				case <-ticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range assets {
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					log.Println(newAssetPrices)
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets)
				case <-mandatoryticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range assets {
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices, regardless of deviation
					emptyMap := make(map[string]float64)
					publishedPrices, err = oracleUpdateExecutor(emptyMap, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets)
				}
			} else {
				select {
				case <-ticker.C:
					// Get prices for all assets from the API
					newAssetPrices := make(map[string]float64)
					for _, asset := range assets {
						newAssetPrice, err := retrieveAssetPrice(asset, useGql, gqlWindowSize, gqlMethodology, asset.gqlParams)
						if err != nil {
							log.Println(err)
							continue
						}
						newAssetPrices[asset.symbol] = newAssetPrice
					}
					// update all prices
					publishedPrices, err = oracleUpdateExecutor(publishedPrices, newAssetPrices, deviationPermille, auth, contract, conn, chainId, assets)
				}
			}
		}
	}()

	select {}
}

func oracleUpdateExecutor(
	publishedPrices map[string]float64,
	newPrices map[string]float64,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	chainId int64,
	assets []Asset) (map[string]float64, error) {
	// Check for deviation and collect all new prices in a map
	// If a published price is 0, update in any case
	updateCollector := make(map[string]float64)
	priceCollector := make(map[string]float64)
	for _, asset := range assets {
		newPrice := newPrices[asset.symbol]
		oldPrice := publishedPrices[asset.symbol]

		if newPrice > 0.0 && (newPrice > (oldPrice * (1 + float64(deviationPermille)/1000))) || (newPrice < (oldPrice * (1 - float64(deviationPermille)/1000))) {
			log.Printf("Entering deviation based update zone for old price %.2f of asset %s. New price: %.2f", oldPrice, asset.symbol, newPrice)
			updateCollector[asset.symbol] = newPrice
			priceCollector[asset.symbol] = newPrice
		} else {
			priceCollector[asset.symbol] = oldPrice
		}
	}

	// Check if any update is to be executed
	if len(updateCollector) == 0 {
		log.Println("No update necessary.")
		return priceCollector, nil
	}
	// Serialize keys and values into array for oracle update transaction
	var keys []string
	var prices []int64
	for key, price := range updateCollector {
		key = key + "/USD"
		keys = append(keys, key)
		integerPrice := int64(price * 100000000)
		prices = append(prices, integerPrice)
	}

	// Update prices in one transaction
	timestamp := time.Now().Unix()
	err := updateOracleMultiValues(conn, contract, auth, chainId, keys, prices, timestamp)
	if err != nil {
		log.Fatalf("Failed to update Oracle: %v", err)
		return nil, err
	}

	return priceCollector, nil
}

func retrieveAssetPrice(asset Asset, useGql bool, gqlWindowSize int, gqlMethodology string, gqlLiquidityParameters GqlParameters) (float64, error) {
	var err error
	var price float64

	// Get quotation for token and update Oracle
	if useGql {
		price, err = getGraphqlAssetQuotationFromDia(asset.blockchain, asset.address, gqlWindowSize, gqlMethodology, gqlLiquidityParameters)
		if err != nil {
			log.Printf("Failed to retrieve %s (%s) quotation data from Graphql on DIA: %v", asset.address, asset.symbol, err)
			return 0.0, err
		}
	} else {
		price, err = getAssetQuotationFromDia(asset.blockchain, asset.address)
		if err != nil {
			log.Fatalf("Failed to retrieve %s quotation data from DIA: %v", asset.address, err)
			return 0.0, err
		}
	}
	return price, nil
}

func bindContract(
	deployedContract string,
	conn *ethclient.Client,
	auth *bind.TransactOpts,
	contract **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService) error {
	var err error
	*contract, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), conn)
	if err != nil {
		return err
	}
	return nil
}

func updateOracleMultiValues(
	client *ethclient.Client,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	auth *bind.TransactOpts,
	chainId int64,
	keys []string,
	values []int64,
	timestamp int64) error {

	var cValues []*big.Int
	var gasPrice *big.Int
	var err error

	// Get proper gas price depending on chainId
	switch chainId {
	case 288:
		gasPrice = big.NewInt(1000000000)
	default:
		// Get gas price suggestion
		gasPrice, err = client.SuggestGasPrice(context.Background())
		if err != nil {
			log.Fatal(err)
		}

		// Get 110% of the gas price
		fGas := new(big.Float).SetInt(gasPrice)
		fGas.Mul(fGas, big.NewFloat(1.1))
		gasPrice, _ = fGas.Int(nil)
	}

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := big.NewInt(value)
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:   auth.From,
		Signer: auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	if err != nil {
		return err
	}

	log.Printf("Gas price: %d\n", tx.GasPrice())
	log.Printf("Data: %x\n", tx.Data())
	log.Printf("Nonce: %d\n", tx.Nonce())
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func getAssetQuotationFromDia(blockchain, address string) (float64, error) {
	response, err := http.Get("https://api.diadata.org/v1/assetQuotation/" + blockchain + "/" + address)
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
	var quotation models.Quotation
	err = quotation.UnmarshalBinary(contents)
	if err != nil {
		return 0.0, err
	}
	return quotation.Price, nil
}

func getGraphqlAssetQuotationFromDia(blockchain, address string, windowSize int, gqlMethodology string, gqlParameters GqlParameters) (float64, error) {
	// Decide whether Feedselection or simple Address/blockchain logic is used
	feedSelectionQuery := "FeedSelection: ["
	if len(gqlParameters.FeedSelection) > 0 {
		// Loop through all selected feeds (e.g. for crosschain feeds)
		for _, selectedFeed := range gqlParameters.FeedSelection {
			// generate strings for optional parameters for liquidity threshold/pool selection
			var lqThresholdString string
			if selectedFeed.LiquidityThreshold > 0 {
				lqThresholdString = "LiquidityThreshold:" + fmt.Sprintf("%.2f", gqlParameters.FeedSelection[0].LiquidityThreshold) + ","
			} else {
				lqThresholdString = ""
			}
			var exchangePairsString string
			if len(selectedFeed.Exchangepairs) > 0 {
				exchangePairsString = "Exchangepairs:[\n"
				for _, exchangePair := range selectedFeed.Exchangepairs {
					exchangePairsString += `{
					Exchange: "` + exchangePair.Exchange + `",
					Pairs: [`
					for _, pair := range exchangePair.Pairs {
						exchangePairsString += `"` + pair + `",`
					}
					exchangePairsString += `]},`
				}
				exchangePairsString += "]"
			} else {
				exchangePairsString = ""
			}
			feedSelectionQuery += `{
				Address:"` + selectedFeed.Address + `",
				Blockchain:"` + selectedFeed.Blockchain + `",` +
				lqThresholdString +
				exchangePairsString +
				`},`
		}
	}	else {
			feedSelectionQuery += `{
				Address: "` + address + `",
				Blockchain: "` + blockchain + `",
			}`
	}
	feedSelectionQuery += "]"

	// Get times for start/end
	currentTime := time.Now()
	starttime := currentTime.Add(time.Duration(-windowSize*2) * time.Second)

	type Response struct {
		GetFeed []struct {
			Name   string    `json:"Name"`
			Time   time.Time `json:"Time"`
			Value  float64   `json:"Value"`
			Pools  string    `json:"Pools"`
			Pairs  string    `json:"Pairs"`
		} `json:"GetFeed"`
	}

	client := gql.NewClient("https://api.diadata.org/graphql/query")
	req := gql.NewRequest(`
    query  {
		 GetFeed(
		 	Filter: "` + gqlMethodology + `",
			BlockSizeSeconds: ` + strconv.Itoa(windowSize) + `,
			BlockShiftSeconds: ` + strconv.Itoa(windowSize) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `,
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `,` +
			feedSelectionQuery +
		 `) {
				Name
				Time
				Value
	   }
		}`)

	ctx := context.Background()
	var r Response
	if err := client.Run(ctx, req, &r); err != nil {
		return 0.0, err
	}
	if len(r.GetFeed) == 0 {
		return 0.0, errors.New("no results")
	}
	return r.GetFeed[len(r.GetFeed)-1].Value, nil
}
