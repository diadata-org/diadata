package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	diaOracleV2MultiupdateService "github.com/diadata-org/diadata/pkg/dia/scraper/blockchain-scrapers/blockchains/ethereum/diaOracleV2MultiupdateService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/diadata-org/diadata/config/nftContracts/erc20"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/google/uuid"
)

const (
	baseURL      = "https://api.navconsulting.net"
)

type CashBalance struct {
	ReportDate           string  `json:"ReportDate"`
	Account              string  `json:"Account"`
	Currency             string  `json:"Currency"`
	CurrencyCode         string  `json:"CurrencyCode"`
	NativeValueEndBalance float64 `json:"NativeValueEndBalance"`
	BaseValueEndBalance   float64 `json:"BaseValueEndBalance"`
	ExchangeRate          float64 `json:"ExchangeRate"`
	Source                string  `json:"Source"`
}

// Since the response is an array of CashBalance, the result type for navFetchCashBalances will be a slice of CashBalance.
type CashBalancesResponse []CashBalance

func main() {
	key := utils.Getenv("PRIVATE_KEY", "")
	key_password := utils.Getenv("PRIVATE_KEY_PASSWORD", "")
	deployedContract := utils.Getenv("DEPLOYED_CONTRACT", "")
	blockchainNode := utils.Getenv("BLOCKCHAIN_NODE", "")
	backupNode := utils.Getenv("BACKUP_NODE", "")
	dataNode := utils.Getenv("DATA_NODE", "")
	apiKey := utils.Getenv("NAV_API_KEY", "")
	credential := utils.Getenv("NAV_CREDENTIAL", "")
	globalFundID := utils.Getenv("NAV_GLOBAL_FUND_ID", "")
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

	publishedValue := 0.0

	/*
	 * Setup connection to contract, deploy if necessary
	 */

	conn, err := ethclient.Dial(blockchainNode)
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	connBackup, err := ethclient.Dial(backupNode)
	if err != nil {
		log.Fatalf("Failed to connect to the backup Ethereum client: %v", err)
	}

	dataConn, err := ethclient.Dial(dataNode)
	if err != nil {
		log.Fatalf("Failed to connect to the data client: %v", err)
	}

	auth, err := bind.NewTransactorWithChainID(strings.NewReader(key), key_password, big.NewInt(chainId))
	if err != nil {
		log.Fatalf("Failed to create authorized transactor: %v", err)
	}

	var contract, contractBackup *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService
	err = deployOrBindContract(deployedContract, conn, connBackup, auth, &contract, &contractBackup)
	if err != nil {
		log.Fatalf("Failed to Deploy or Bind primary and backup contract: %v", err)
	}

	// Contract with USDLR and ERC20 credentials to get totalSupply
	usdlrconn, err := erc20.NewERC20(common.HexToAddress("0x68592c5c98c4f4a8a4bc6da2121e65da3d1c0917"), dataConn)
	if err != nil {
		log.Fatalf("Failed to get erc20 conn: %v", err)
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
					// Get cash balances from NAV API
  				cashBalances, navTimestamp := navFetchCashBalances(apiKey, credential, globalFundID)
				  log.Println(cashBalances)
				  if cashBalances == nil {
				  	log.Println("Error! NAV API return was empty")
				  	continue
					}
				  // Calculate value to be updated
					newValue, err := calculateNavUpdateValue(cashBalances, usdlrconn)
					if err != nil {
						log.Println(err)
						continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(publishedValue, newValue, navTimestamp, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(publishedValue, newValue, navTimestamp, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				case <-mandatoryticker.C:
					// Get cash balances from NAV API
  				cashBalances, navTimestamp := navFetchCashBalances(apiKey, credential, globalFundID)
				  log.Println(cashBalances)
				  if cashBalances == nil {
				  	log.Println("Error! NAV API return was empty")
				  	continue
					}
				  // Calculate value to be updated
					newValue, err := calculateNavUpdateValue(cashBalances, usdlrconn)
					if err != nil {
						log.Println(err)
						continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(0.0, newValue, navTimestamp, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(0.0, newValue, navTimestamp, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				}
			} else {
				select {
				case <-ticker.C:
					// Get cash balances from NAV API
  				cashBalances, navTimestamp := navFetchCashBalances(apiKey, credential, globalFundID)
				  log.Println(cashBalances)
				  if cashBalances == nil {
				  	log.Println("Error! NAV API return was empty")
				  	continue
					}
				  // Calculate value to be updated
					newValue, err := calculateNavUpdateValue(cashBalances, usdlrconn)
					if err != nil {
						log.Println(err)
						continue
					}
					// update value
					publishedValue, err = oracleUpdateExecutor(publishedValue, newValue, navTimestamp, deviationPermille, auth, contract, conn, chainId)
					if err != nil {
						log.Printf("Failed to execute oracle update using primary connection: %v. Retrying with backup connection...", err)

						// Attempt using the backup connection
						publishedValue, err = oracleUpdateExecutor(publishedValue, newValue, navTimestamp, deviationPermille, auth, contractBackup, connBackup, chainId)
						if err != nil {
							log.Fatalf("Failed to execute oracle update using backup connection: %v", err)
						}
					}
				}
			}
		}
	}()

	select {}
}

func oracleUpdateExecutor(
	publishedValue float64,
	newValue float64,
	navTimestamp time.Time,
	deviationPermille int,
	auth *bind.TransactOpts,
	contract *diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	conn *ethclient.Client,
	chainId int64) (float64, error) {

	// If a published price is 0, update in any case
	fmt.Println("new price", newValue)

	if (newValue > 1e-8 && ((newValue > (publishedValue * (1 + float64(deviationPermille)/1000))) || (newValue < (publishedValue * (1 - float64(deviationPermille)/1000))))) {
		log.Printf("Entering deviation based update zone for old value %.2f. New value: %.2f", publishedValue, newValue)
	} else {
		log.Println("No update necessary.")
		return publishedValue, nil
	}

	// Serialize keys and values into array for oracle update transaction
	var keys []string
	var values []int64

	keys = append(keys, "USDLR/USD")
	values = append(values, int64(newValue * 1000000))
	
	keys = append(keys, "TS-USDLR/USD")
	values = append(values, navTimestamp.Unix())

	// Update values in one transaction
	timestamp := time.Now().Unix()
	err := updateOracleMultiValues(conn, contract, auth, chainId, keys, values, timestamp)
	if err != nil {
		log.Printf("Failed to update Oracle: %v", err)
		return 0.0, err
	}

	return newValue, nil
}

func deployOrBindContract(
	deployedContract string,
	conn *ethclient.Client,
	connBackup *ethclient.Client,
	auth *bind.TransactOpts,
	contract **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService,
	contractBackup **diaOracleV2MultiupdateService.DiaOracleV2MultiupdateService) error {
	var err error
	if deployedContract != "" {
		// bind primary and backup
		*contract, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), conn)
		if err != nil {
			return err
		}
		*contractBackup, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(common.HexToAddress(deployedContract), connBackup)
		if err != nil {
			return err
		}
	} else {
		// deploy contract
		var addr common.Address
		var tx *types.Transaction
		addr, tx, *contract, err = diaOracleV2MultiupdateService.DeployDiaOracleV2MultiupdateService(auth, conn)
		if err != nil {
			log.Fatalf("could not deploy contract: %v", err)
			return err
		}
		log.Printf("Contract pending deploy: 0x%x\n", addr)
		log.Printf("Transaction waiting to be mined: 0x%x\n\n", tx.Hash())
		// bind backup
		*contractBackup, err = diaOracleV2MultiupdateService.NewDiaOracleV2MultiupdateService(addr, connBackup)
		if err != nil {
			return err
		}
		time.Sleep(180000 * time.Millisecond)
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

	// Get gas price suggestion
	gasPrice, err = client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Print(err)
		return err
	}

	// Get 110% of the gas price
	fGas := new(big.Float).SetInt(gasPrice)
	fGas.Mul(fGas, big.NewFloat(1.1))
	gasPrice, _ = fGas.Int(nil)

	for _, value := range values {
		// Create compressed argument with values/timestamps
		cValue := big.NewInt(value)
		cValue = cValue.Lsh(cValue, 128)
		cValue = cValue.Add(cValue, big.NewInt(timestamp))
		cValues = append(cValues, cValue)
	}

	// Write values to smart contract
	tx, err := contract.SetMultipleValues(&bind.TransactOpts{
		From:     auth.From,
		Signer:   auth.Signer,
		GasPrice: gasPrice,
	}, keys, cValues)
	// check if tx is sendable then fgo backup
	if err != nil {
		// backup in here
		return err
	}

	log.Printf("Gas price: %d\n", tx.GasPrice())
	log.Printf("Data: %x\n", tx.Data())
	log.Printf("Nonce: %d\n", tx.Nonce())
	log.Printf("Tx To: %s\n", tx.To().String())
	log.Printf("Tx Hash: 0x%x\n", tx.Hash())
	return nil
}

func calculateNavUpdateValue(apiResults *CashBalancesResponse, usdlrconn *erc20.ERC20) (float64, error) {
	var returnValue float64
	usdcPrice, err := getAssetQuotationFromDia("Ethereum", "0xA0b86991c6218b36c1d19D4a2e9Eb0cE3606eB48")
	if err != nil {
		return 0.0, err
	}
	// Get total supply on Ethereum for USDLR
	biUsdlrSupply, err := usdlrconn.TotalSupply(&bind.CallOpts{})
	if err != nil {
		log.Printf("Failed to get totalSupply of usdlr: %v", err)
		return 0.0, err
	}
	// Convert to float
	f := new(big.Float).SetInt(biUsdlrSupply)
	usdlrSupply, _ := f.Float64()
	usdlrSupply /= 1e6
	for _, result := range *apiResults {
		if result.CurrencyCode == "USD" {
			// BaseValueEndBalance of cash and T-bill reserves
			returnValue += result.BaseValueEndBalance
		} else if result.CurrencyCode == "USDC" {
			// BaseValueEndBalance of USDC * USDC price
			returnValue += result.BaseValueEndBalance * usdcPrice
		} else {
			// This shouldn't happen, unexpected API return value
			return 0.0, fmt.Errorf("Error! Unexpected return in NAV API response: %s", result)
		}
	}
	returnValue /= usdlrSupply
	return returnValue, nil
}

func getAssetQuotationFromDia(blockchain, address string) (float64, error) {
	// ibc special case: convert / to - in the query string
	if strings.Split(address, "/")[0] == "ibc" {
		address = strings.Split(address, "/")[0] + "-" + strings.Split(address, "/")[1]
	}

	// Execute the query
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

func navSignRequest(method, urlPath, body, apiKey, secret string) map[string]string {
	verb := strings.ToUpper(method)
	utcNow := time.Now().UTC().Format("Mon, 02 Jan 2006 15:04:05 ") + "GMT"
	nonce := uuid.NewString()

	hasher := sha256.New()
	hasher.Write([]byte(body))
	contentDigest := hasher.Sum(nil)
	contentHash := base64.StdEncoding.EncodeToString(contentDigest)

	stringToSign := apiKey + ";" + urlPath + ";" + verb + ";" + utcNow + ";" + nonce + ";" + contentHash

	decodedSecret := []byte(secret)
	hmacHasher := hmac.New(sha256.New, decodedSecret)
	hmacHasher.Write([]byte(stringToSign))
	digest := hmacHasher.Sum(nil)

	signature := base64.StdEncoding.EncodeToString(digest)

	return map[string]string{
		"x-date":              utcNow,
		"x-hmac256-signature": apiKey + ";" + nonce + ";" + signature,
		"x-content-sha256":    contentHash,
	}
}

// Modify the fetch function to accept a result pointer to unmarshal the JSON response directly into it.
func navFetch(slug string, body map[string]string, apiKey string, credential string, result interface{}) {
	params := url.Values{}
	for key, value := range body {
			params.Add(key, value)
	}
	slugWithParams := slug
	if len(body) > 0 {
			slugWithParams += "?" + params.Encode()
	}
	fullURL := baseURL + slugWithParams

	headers := navSignRequest("GET", slugWithParams, "", apiKey, credential)

	client := &http.Client{}
	request, err := http.NewRequest("GET", fullURL, nil)
	if err != nil {
			fmt.Println("Error creating request:", err)
			return
	}

	for key, value := range headers {
			request.Header.Add(key, value)
	}

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error fetching data from NAV API")
		return
	}
	defer response.Body.Close()

	if response.StatusCode != 200 {
		fmt.Println("Error fetching data:", response.Status)
		return
	}

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	err = json.Unmarshal(bodyBytes, result)
	if err != nil {
			fmt.Println("Error unmarshalling response body:", err)
			return
	}
}

func navFetchCashBalances(apiKey string, credential string, globalFundID string) (*CashBalancesResponse, time.Time) {
	var reportDate string //01-08-2024
	for offset := 1; offset < 31; offset++ {
		// Get report date (T-1) in the correct format
		requestDate := time.Now().AddDate(0, 0, -offset)
		reportDate = fmt.Sprintf(requestDate.Format("01-02-2006"))
		log.Printf("Requesting NAV data for date %s.", reportDate)
		body := map[string]string{
				"globalFundID": globalFundID,
				"reportDate":   reportDate,
		}
		slug := "/navapigateway/api/v1/FundReportData/GetCashBalancesForFund"
		result := &CashBalancesResponse{} // Using the newly defined type
		navFetch(slug, body, apiKey, credential, result)
		if len(*result) == 0 {
			continue
		}
		return result, requestDate.Truncate(24 * time.Hour)
	}
	return nil, time.Now().Truncate(24 * time.Hour)
}
