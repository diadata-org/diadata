package main

import (
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha256"
	"database/sql"
	"encoding/base64"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"time"

	"strings"

	"github.com/ethereum/go-ethereum/common"

	kr "github.com/99designs/keyring"
	"github.com/diadata-org/diadata/pkg/dia"
	k8sbridge "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
	"github.com/diadata-org/diadata/pkg/http/restApi"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"

	"github.com/gin-gonic/gin"
)

/*
Auth using EIP712 spec
*/

//goland:noinspection ALL
type Env struct {
	DataStore               models.Datastore
	RelDB                   models.RelDatastore
	k8sbridgeClient         k8sbridge.K8SHelperClient
	Keyring                 kr.Keyring
	RateLimitOracleCreation int
	SupportedChain          []string
	LoopPaymentSecret       string
	RedirectURL             string
}

const (
	REQUEST_ID      = "REQUEST_ID"
	CUSTOMER_ID     = "CUSTOMER_ID"
	ACCESS_LEVEL    = "ACCESS_LEVEL"
	CUSTOMER_PLAN   = "CUSTOMER_PLAN"
	CREATOR_ADDRESS = "CREATOR_ADDRESS"
)

func NewEnv(relStore models.RelDatastore, ds models.Datastore, kc k8sbridge.K8SHelperClient, ring kr.Keyring, rateLimitOracleCreation int, lps string) *Env {

	ru := utils.Getenv("PAYMENT_SUCCESS_URL", "")
	return &Env{
		RedirectURL:             ru,
		DataStore:               ds,
		RelDB:                   relStore,
		Keyring:                 ring,
		k8sbridgeClient:         kc,
		RateLimitOracleCreation: rateLimitOracleCreation,
		LoopPaymentSecret:       lps,
		SupportedChain:          []string{"11155111", "421614", "123420111", "94204209", "88153591557", "80002"},
	}
}

func (ob *Env) ViewLimit(context *gin.Context) {

	response := make(map[string]interface{})
	response["Count"] = 1
	response["IsAllowed"] = true

	context.JSON(http.StatusOK, response)
}

func (ob *Env) billableResource(customerId string) (totalFeeds, totalOracles int, err error) {
	// get all oracleconfig
	oracleConfig, err := ob.RelDB.GetOraclesByCustomer(customerId)
	if err != nil {
		return 0, 0, err
	}

	for _, oracle := range oracleConfig {
		if !oracle.Deleted && oracle.Billable {
			var sf []SymbolFeed
			json.Unmarshal([]byte(oracle.FeederSelection), &sf)
			totalFeeds = totalFeeds + len(sf)
			totalOracles++
		}

	}
	return
}

// Create new oracle feeder if creator has resources
func (ob *Env) CreateOneStep(context *gin.Context) {

	log.Infoln("Create New Oracle")
	var (
		address  string
		err      error
		keypair  *k8sbridge.KeyPair
		isUpdate bool
		response map[string]string
	)
	response = make(map[string]string)

	requestId := GenerateRandString()
	context.Set(REQUEST_ID, requestId)

	isUpdate = false

	oracleaddress := context.PostForm("oracleaddress")
	oracleaddress = common.HexToAddress(oracleaddress).Hex()
	chainID := context.PostForm("chainID")
	creator := context.PostForm("creator")
	oracleName := context.PostForm("name")

	creator = common.HexToAddress(creator).Hex()

	symbols := context.PostForm("symbols")
	signedData := context.PostForm("signeddata")
	feederID := context.PostForm("feederID")
	frequency := context.PostForm("frequency")
	sleepSeconds := context.PostForm("sleepseconds")
	deviationPermille := context.PostForm("deviationpermille")

	blockchainnode := context.PostForm("blockchainnode")
	feedSelection := context.PostForm("feedselection")

	mandatoryFrequency := context.PostForm("mandatoryfrequency")

	log.Infof("Request ID: %s Creating oracle: oracleAddress: %s, ChainID: %s, Creator: %s, Symbols: %s, frequency: %s, sleepSeconds: %s blockchainnode: %s, feedSelection %s name %s", requestId, oracleaddress, chainID, creator, symbols, frequency, sleepSeconds, blockchainnode, feedSelection, oracleName)

	signer, err := utils.GetSigner(chainID, creator, oracleaddress, "Verify its your address to call oracle builder", signedData)
	if err != nil {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer")
		return

	}
	log.Infof("Request ID: %s Creating oracle: signer", requestId, signer)

	if signer.Hex() != creator {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer %v", signer)
	}

	// validations
	// check for  symbols and feedSelection

	if symbols == "" && feedSelection == "" {
		handleError(context, http.StatusBadRequest, "no symbols", "Creating oracle: no symbols or feedSelection %v", symbols)

	}

	// get customer id
	customer, err := ob.RelDB.GetCustomerByPublicKey(signer.Hex())
	if err != nil {
		log.Errorf("Request ID: %s, GetCustomerByPublicKey %v ,", requestId, err)
		handleError(context, http.StatusNotFound, "error getting customer", "Creating oracle:account not found for this wallet")
		return

	}

	customerID := customer.CustomerID

	// Get Plan and allowed feeds

	plan, err := ob.RelDB.GetPlan(context, customer.CustomerPlan)

	if err != nil {
		log.Errorf("Request ID: %s, GetPlan %v ,", requestId, err)

		handleError(context, http.StatusInternalServerError, "plan err", "Creating oracle: invalid plan")
		return

	}

	log.Infof("Request ID: %s, Plan %v ,", requestId, plan)

	/* total Feeds used

	 */
	totalFeeds, totalOracles, err := ob.billableResource(strconv.Itoa(customerID))
	if err != nil {
		log.Errorf("Request ID: %s, Total feeds err %v ,", requestId, err)

	}
	if totalFeeds >= plan.TotalFeeds {
		log.Errorf("Request ID: %s, totalFeeds exceeds plan Limit %v ,", requestId, err)
		handleError(context, http.StatusPaymentRequired, "totalFeeds exceeds plan Limit", "Creating oracle:  totalFeeds exceeds plan Limit")
		return
	}
	if totalOracles >= plan.TotalOracles {
		log.Errorf("Request ID: %s, totalOracles exceeds plan Limit %v ,", requestId, err)
		handleError(context, http.StatusPaymentRequired, "totalFeeds exceeds plan Limit", "Creating oracle:  totalFeeds exceeds plan Limit")
		return
	}

	log.Infof("Request ID: %s, Total feeds %d,", requestId, totalFeeds)

	var sf []SymbolFeed

	json.Unmarshal([]byte(feedSelection), &sf)

	log.Infof("Request ID: %s, CustomerID: %d Creating oracle: total feeds", requestId, customerID, len(sf))

	symbolsArray := strings.Split(symbols, ",")

	if len(symbolsArray) > 10 {
		handleError(context, http.StatusBadRequest, "max symbols exceed", "Creating oracle: max symbols exceed %d", len(symbols))
	}

	// check for duplicate symbol

	if utils.CheckDuplicates(symbolsArray) {
		handleError(context, http.StatusBadRequest, "duplicate symbols", "Creating oracle: duplicate symbols %v", symbols)

	}

	// check frequency limit

	frequencyInt, err := strconv.Atoi(frequency)
	if err != nil {
		handleError(context, http.StatusBadRequest, "invalid frequency", "Creating oracle: invalid frequency %v", err)
	}

	mandatoryFrequencyInt, err := strconv.Atoi(mandatoryFrequency)
	if err != nil {
		mandatoryFrequencyInt = 0
		// handleError(context, http.StatusBadRequest, "invalid mandatoryFrequencyInt", "Creating oracle: invalid mandatoryFrequencyInt", err)
	}

	if frequencyInt != 0 || mandatoryFrequencyInt == 0 {
		if frequencyInt < 120 || frequencyInt > 2630000 {
			context.JSON(http.StatusBadRequest, errors.New("invalid frequency, out of range"))
			log.Errorln("Creating oracle: invalid frequency, out of range", frequencyInt)
			return
		}
	}

	if frequencyInt == 0 || mandatoryFrequencyInt > 0 {
		if mandatoryFrequencyInt < 120 || mandatoryFrequencyInt > 2630000 {
			handleError(context, http.StatusBadRequest, "invalid mandatoryFrequencyInt, out of range", "Creating oracle: invalid mandatoryFrequencyInt, out of range %v", err)
		}

	}

	deviationPermilleFloat, err := strconv.ParseFloat(deviationPermille, 64)
	if err != nil {
		deviationPermilleFloat = 0.0

	}

	if deviationPermilleFloat > 0 {
		if deviationPermilleFloat < 0.1 && deviationPermilleFloat > 10000 {
			if err != nil {
				context.JSON(http.StatusBadRequest, errors.New("invalid deviationPermille"))
				log.Errorln("Creating oracle: invalid deviationPermille", err)
				return
			}

		}

		deviationPermilleFloat = deviationPermilleFloat * 10
		deviationPermille = fmt.Sprintf("%.2f", deviationPermilleFloat)

	}

	log.Infoln("feederId from creator", feederID)

	if feederID == "" {
		//Oracle Creation Action

		feederID = "feeder-" + utils.GenerateAutoname("-")

		err = ob.Keyring.Set(kr.Item{
			Key: feederID,
		})

		if err != nil {
			log.Errorln("error getting key", err)
			context.JSON(http.StatusUnauthorized, errors.New("need access to this feeder"))
			return
		}
	} else {

		owner := ob.RelDB.GetFeederByID(feederID)

		if owner != creator {
			log.Infoln("no access to feederID, owner is ", owner)
			context.JSON(http.StatusInternalServerError, errors.New("no access to feederID"))
			return
		}
		isUpdate = true

	}

	item, err := ob.Keyring.Get(feederID)
	if err != nil {
		log.Infoln("error getting key", err)
		context.JSON(http.StatusInternalServerError, errors.New("error getting key"))
		return
	}
	marshalErr := json.Unmarshal(item.Data, &keypair)
	if marshalErr != nil {
		return
	}
	log.Infoln("public key", keypair.GetPublickey())
	address = keypair.GetPublickey()

	var symbolFeeds []SymbolFeed

	if feedSelection != "" {
		if err := json.Unmarshal([]byte(feedSelection), &symbolFeeds); err != nil {
			return
		}

	}

	log.Println("total feeds selected", len(symbolFeeds))

	if !isUpdate {

		fc := &k8sbridge.FeederConfig{
			FeederID:           feederID,
			Creator:            creator,
			FeederAddress:      address,
			Oracle:             oracleaddress,
			ChainID:            chainID,
			Symbols:            symbols,
			FeedSelection:      feedSelection,
			Blockchainnode:     blockchainnode,
			Frequency:          frequency,
			SleepSeconds:       sleepSeconds,
			DeviationPermille:  deviationPermille,
			MandatoryFrequency: mandatoryFrequency,
		}
		_, err = ob.k8sbridgeClient.CreatePod(context, fc)
		if err != nil {
			log.Errorln("error CreateOracleFeeder ", err)
			context.JSON(http.StatusInternalServerError, errors.New("error creating oraclefeeder"))
			return
		}

	}

	err = ob.RelDB.SetOracleConfig(context, strconv.Itoa(customerID), oracleaddress, feederID, creator, address, symbols, feedSelection, chainID, frequency, sleepSeconds, deviationPermille, blockchainnode, mandatoryFrequency, oracleName, false, true)
	if err != nil {
		log.Errorln("requestId: %s, error SetOracleConfig ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error creating new oracle config"})
		return
	}

	if isUpdate {
		oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress, chainID)
		if err != nil {
			log.Errorln("error GetOracleConfig ", err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}

		fc := &k8sbridge.FeederConfig{
			FeederID:           feederID,
			Creator:            oracleconfig.Owner,
			FeederAddress:      oracleconfig.FeederAddress,
			Oracle:             oracleconfig.Address,
			ChainID:            oracleconfig.ChainID,
			Symbols:            strings.Join(oracleconfig.Symbols[:], ","),
			FeedSelection:      oracleconfig.FeederSelection,
			Blockchainnode:     oracleconfig.BlockchainNode,
			Frequency:          oracleconfig.Frequency,
			SleepSeconds:       oracleconfig.SleepSeconds,
			DeviationPermille:  oracleconfig.DeviationPermille,
			MandatoryFrequency: oracleconfig.MandatoryFrequency,
		}
		_, err = ob.k8sbridgeClient.CreatePod(context, fc)

		if err != nil {
			log.Errorln("error RestartOracleFeeder ", err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	log.Infof("Created oracle: oracleAddress: %s, ChainID: %s, Creator: %s, Symbols: %s, frequency: %s, sleepSeconds: %s, Feeder ID :%s,", oracleaddress, chainID, creator, symbols, frequency, sleepSeconds, feederID)

	response["oracleaddress"] = oracleaddress
	response["chainId"] = chainID
	response["creator"] = creator
	response["symbols"] = symbols
	response["publicKey"] = address

	context.JSON(http.StatusCreated, response)
}

// combim list and view account to save signing extra extra endpoint on dashboard
func (ob *Env) ListAndViewAccount(context *gin.Context) {
	creator := context.Query("creator")
	requestId := context.GetString(REQUEST_ID)

	var (
		oracles []dia.OracleConfig
		account map[string]interface{}
	)

	// get customer id, get oracles by customer

	customerId, err := ob.RelDB.GetCustomerIDByWalletPublicKey(creator)
	if err != nil && err.Error() == "no rows in result set" {
		// get pending requests

		pending, err := ob.RelDB.GetPendingInvites(context, creator)
		if err != nil {

			// new wallet
			log.Errorf("Request ID: %s,  err GetPendingInvites %v ", requestId, err)

		}

		account = map[string]interface{}{

			"pending_public_keys": pending,
		}

	} else {
		// check if 0 oracles
		oracles, err = ob.RelDB.GetOraclesByCustomer(strconv.Itoa(customerId))
		if err != nil {
			errorMsg := "error fetching oracles for this customer"
			logMsg := "List Oracles: error on getOraclesByOwner"
			handleError(context, http.StatusNotFound, errorMsg, logMsg, err)
			return
		}

		account = ob.viewAccount(context, creator)

	}

	r := map[string]interface{}{
		"account": account,
		"oracles": oracles,
	}
	context.JSON(http.StatusOK, r)
}

// List: list owner oracles
func (ob *Env) List(context *gin.Context) {
	creator := context.Query("creator")

	// get customer id, get oracles by customer

	customerId, err := ob.RelDB.GetCustomerIDByWalletPublicKey(creator)
	if err != nil {
		errorMsg := "error fetching customer by this address " + creator
		logMsg := "List Oracles: error on GetCustomerIDByWalletPublicKey"
		handleError(context, http.StatusNotFound, errorMsg, logMsg, err)
		return
	}

	oracles, err := ob.RelDB.GetOraclesByCustomer(strconv.Itoa(customerId))
	if err != nil {
		errorMsg := "error fetching oracles for this customer"
		logMsg := "List Oracles: error on getOraclesByOwner"
		handleError(context, http.StatusNotFound, errorMsg, logMsg, err)
		return
	}
	context.JSON(http.StatusOK, oracles)
}

// List whitelisted addresses
func (ob *Env) Whitelist(context *gin.Context) {
	log.Errorln("Whitelist ")

	addresses, err := ob.RelDB.GetFeederResources()
	if err != nil {
		log.Errorln("List Whitelist: error on GetFeederResources ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, addresses)
}

func (ob *Env) SupportedChains(context *gin.Context) {
	chains, err := ob.RelDB.GetAllChains()
	if err != nil {
		log.Errorln("List SupportedChains: error on GetAllChains ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	chainArray := []string{}

	for _, chainInfo := range chains {

		if contains(ob.SupportedChain, chainInfo.ChainID) {
			chainArray = append(chainArray, chainInfo.ChainID)

		}

	}

	context.JSON(http.StatusOK, chainArray)
}

func convertExchangePairs(exchangePairs []Exchangepair) string {
	// Create a slice to hold JSON objects
	var jsonObjects []string

	// Iterate through the exchangePairs
	for _, exchangePair := range exchangePairs {
		jsonObject := fmt.Sprintf(`{Exchange: "%s"`, exchangePair.Exchange)

		var quotedArr []string

		for _, item := range exchangePair.Pairs {
			quotedArr = append(quotedArr, `"`+item+`"`)
		}
		pairs := `[` + strings.Join(quotedArr, ",") + `]`

		jsonObject += `,Pairs:` + pairs

		jsonObject += "}"

		jsonObjects = append(jsonObjects, jsonObject)
	}

	jsonArray := "[" + strings.Join(jsonObjects, ",") + "]"

	return jsonArray
}

type Exchangepair struct {
	Exchange string   `json:"Exchange"`
	Pairs    []string `json:"Pairs"`
}

type FeedSelection struct {
	Address       string         `json:"Address"`
	Blockchain    string         `json:"Blockchain"`
	Exchangepairs []Exchangepair `json:"Exchangepairs"`
}

type SymbolFeed struct {
	LiquidityThreshold string          `json:"LiquidityThreshold"`
	Methodology        string          `json:"Methodology"`
	Symbol             string          `json:"Symbol"`
	FeedSelection      []FeedSelection `json:"FeedSelection"`
}

func generateFeedSelectionQuery(feedSelections []FeedSelection) string {
	var jsonObjects []string

	for _, feedSelection := range feedSelections {

		exchangepair := convertExchangePairs(feedSelection.Exchangepairs)
		jsonObject := fmt.Sprintf(`{
			Address: %q,
			Blockchain: %q,
			Exchangepairs: %s
		}`, feedSelection.Address, feedSelection.Blockchain, exchangepair)

		jsonObjects = append(jsonObjects, jsonObject)
	}

	jsonArray := "[" + strings.Join(jsonObjects, ",") + "]"

	return jsonArray
}

func (ob *Env) Dashboard(context *gin.Context) {
	address := context.Query("address")
	chainID := context.Query("chainID")

	symbol := context.Query("symbol")

	starttime, endtime, err := utils.MakeTimerange(context.Query("starttime"), context.Query("endtime"), time.Duration(30*24*time.Hour))
	if err != nil {
		endtime = time.Now()
		starttime = endtime.Add(-time.Duration(30 * 24 * time.Hour))

	}

	if ok := utils.ValidTimeRange(starttime, endtime, time.Duration(30*24*time.Hour)); !ok {
		restApi.SendError(context, http.StatusInternalServerError, fmt.Errorf("time-range too big. max duration is %v", 30*24*time.Hour))
		return
	}

	if strings.Contains(address, "0x") {
		address = common.HexToAddress(address).Hex()
	}

	//no pagination required
	offset := -1

	totalUpdates, err := ob.RelDB.GetOracleUpdateCount(address, chainID, symbol)
	if err != nil {
		errorMsg := "Error fetching oracle update count"
		logMsg := "Oracle Stats error GetOracleUpdateCount"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	updates, err := ob.RelDB.GetOracleUpdatesByTimeRange(address, chainID, symbol, offset, starttime, endtime)
	if err != nil {
		errorMsg := "Error fetching oracle updates GetOracleUpdatesByTimeRange"
		logMsg := "GetOracleUpdatesByTimeRange"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	type localAssetConfig struct {
		Symbol     string
		Address    string
		Blockchain string
		Deviation  string
		HeartBeat  string
		GQLQuery   string
		SymbolFeed interface{}

		Volume24h         float64
		LastReportedPrice string
		LastReportedTime  time.Time
	}
	var symbolFeeds []SymbolFeed

	oracleConfig, err := ob.DataStore.GetOracleConfigCache(address + "-" + chainID)
	if err != nil {
		log.Infoln("oracle not in redis might be oracle builder oracle", err)
	}

	if err == nil {

	} else {
		oracleConfig, err = ob.RelDB.GetOracleConfig(address, chainID)
		if err != nil {
			errorMsg := "Error fetching oracle config"
			logMsg := "Oracle Stats error"
			handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
			return
		}

	}

	if oracleConfig.FeederSelection != "" {
		if err := json.Unmarshal([]byte(oracleConfig.FeederSelection), &symbolFeeds); err != nil {
			return
		}

	}

	var (
		assets                                              []localAssetConfig
		avgGasSpend, gasSpend, gasRemaining, avgUpdateCount float64
		updateFeeds                                         []dia.FeedUpdates
	)

	if len(symbolFeeds) > 0 {

		for _, sf := range symbolFeeds {
			blockchain := strings.TrimSpace(strings.Split(sf.Symbol, "-")[0])

			for _, fs := range sf.FeedSelection {
				blockchain = fs.Blockchain
			}

			feedAddress := strings.TrimSpace(strings.Split(sf.Symbol, "-")[1])

			blockDuration := 120
			currentTime := time.Now()
			starttime := currentTime.Add(time.Duration(-blockDuration*2) * time.Second)
			query := `
    query  {
		GetFeed(
		 	Filter: "` + strings.ToLower(sf.Methodology) + `", 
			BlockSizeSeconds: ` + strconv.Itoa(blockDuration) + `, 
			BlockShiftSeconds: ` + strconv.Itoa(blockDuration) + `,
			StartTime: ` + strconv.FormatInt(starttime.Unix(), 10) + `, 
			EndTime: ` + strconv.FormatInt(currentTime.Unix(), 10) + `, 
			FeedSelection: ` + generateFeedSelectionQuery(sf.FeedSelection) + `
	  ) 
	{
				Name
				Symbol
				Time
				Value
	  	}
		}`

			var (
				LastReportedPrice string
				LastReportedTime  time.Time
			)

			var assetSymbol string

			assetDetail, err := ob.RelDB.GetAsset(feedAddress, blockchain)

			if err != nil {
				assetSymbol = blockchain
			} else {
				assetSymbol = assetDetail.Symbol
			}

			LastReportedTime, LastReportedPrice, _ = ob.RelDB.GetOracleLastUpdate(address, chainID, assetSymbol)

			volume24h, err := ob.RelDB.GetLastAssetVolume24H(dia.Asset{Blockchain: blockchain, Address: strings.Split(sf.Symbol, "-")[1]})
			if err != nil {
				volume24h = 0
			}
			asset := localAssetConfig{Symbol: assetSymbol, Address: feedAddress, Blockchain: blockchain, Deviation: oracleConfig.DeviationPermille, HeartBeat: oracleConfig.Frequency, GQLQuery: query, Volume24h: volume24h, LastReportedPrice: LastReportedPrice, LastReportedTime: LastReportedTime, SymbolFeed: sf}
			assets = append(assets, asset)

		}
	} else {

		if len(oracleConfig.Symbols) > 0 {
			for _, symbol := range oracleConfig.Symbols {

				LastReportedTime, LastReportedPrice, _ := ob.RelDB.GetOracleLastUpdate(address, chainID, symbol)

				if len(strings.Split(symbol, "-")) >= 2 {

					blockchain := strings.TrimSpace(strings.Split(symbol, "-")[0])
					address := strings.TrimSpace(strings.Split(symbol, "-")[1])

					var assetSymbol string

					assetDetail, err := ob.RelDB.GetAsset(address, blockchain)
					if err != nil {
						assetSymbol = ""
					} else {
						assetSymbol = assetDetail.Symbol
					}

					volume24h, err := ob.RelDB.GetLastAssetVolume24H(dia.Asset{Blockchain: blockchain, Address: address})
					if err != nil {
						volume24h = 0
					}
					asset := localAssetConfig{Symbol: assetSymbol, Address: address, Blockchain: blockchain, Deviation: oracleConfig.DeviationPermille, HeartBeat: oracleConfig.Frequency, Volume24h: volume24h, LastReportedPrice: LastReportedPrice, LastReportedTime: LastReportedTime}
					assets = append(assets, asset)
				}

			}
		}

	}

	gasSpend, err = ob.RelDB.GetTotalGasSpend(address, chainID)
	if err != nil {
		gasSpend = 0.0
	}

	gasRemaining, err = ob.RelDB.GetBalanceRemaining(address, chainID)
	if err != nil {
		gasRemaining = 0.0
	}
	updateFeeds, avgGasSpend, avgUpdateCount, err = ob.RelDB.GetDayWiseUpdates(address, chainID)
	if err != nil {
		avgGasSpend = 0.0
		avgUpdateCount = 0.0
	}

	lastOracleUpdate, _ := ob.RelDB.GetLastOracleUpdate(address, chainID)

	response := make(map[string]interface{})
	response["OracleAddress"] = address
	response["Chain"] = chainID
	response["OracleType"] = 0

	response["GasRemaining"] = gasRemaining
	response["GasSpend"] = gasSpend
	response["AvgGasSpend"] = avgGasSpend
	response["AvgDailyTxs"] = avgUpdateCount

	if len(lastOracleUpdate) > 0 {
		response["LastOracleUpdate"] = lastOracleUpdate[0].UpdateTime
	}

	response["DayWiseUpdates"] = updateFeeds

	response["Transactions"] = updates
	response["Assets"] = assets
	response["Count"] = totalUpdates
	context.JSON(http.StatusOK, response)
}

func (ob *Env) Stats(context *gin.Context) {
	address := context.Query("address")
	chainID := context.Query("chainID")
	page := context.Query("page")
	symbol := context.Query("symbol")

	var (
		startTime, endTime time.Time
		err                error
	)

	if strings.Contains(address, "0x") {
		address = common.HexToAddress(address).Hex()
	}
	if context.Query("starttime") == "" || context.Query("endTime") == "" {
		startTime = time.Unix(0, 0)
		endTime = time.Unix(0, 0)

	} else {
		startTime, endTime, err = utils.MakeTimerange(context.Query("starttime"), context.Query("endtime"), time.Duration(30*24*time.Hour))

		if err != nil {
			endTime = time.Now()
			startTime = endTime.Add(-time.Duration(30 * 24 * time.Hour))

		}

	}

	offset := 0
	if page != "" {
		pageInt, err := strconv.Atoi(page)
		if err != nil || pageInt < 1 {
			offset = 0
		} else {
			offset = (pageInt - 1) * 20
		}
	} else {
		offset = 0
	}

	totalUpdates, err := ob.RelDB.GetOracleUpdateCount(address, chainID, symbol)
	if err != nil {
		errorMsg := "Error fetching oracle update count"
		logMsg := "Oracle Stats error GetOracleUpdateCount"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	updates, err := ob.RelDB.GetOracleUpdatesByTimeRange(address, chainID, symbol, offset, startTime, endTime)
	if err != nil {
		errorMsg := "Error fetching oracle updates"
		logMsg := "Oracle Stats error"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	response := make(map[string]interface{})
	response["Count"] = totalUpdates
	response["Updates"] = updates

	context.JSON(http.StatusOK, response)
}

// List: list All feeders
func (ob *Env) ListAll(context *gin.Context) {
	oracles, err := ob.RelDB.GetAllFeeders(false, false)
	if err != nil {
		log.Errorln("List All Oracles: error on GetAllFeeders ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, oracles)
}

// view oracle config
func (ob *Env) View(context *gin.Context) {
	var (
		// address string
		err error
	)
	chainID := context.Query("chainID")
	creator := context.Query("creator")
	oracleaddress := context.Query("oracleaddress")

	signedData, err := getAuthToken(context.Request)

	if err != nil {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("missing auth token", err)
		return
	}

	signer, err := utils.GetSigner(chainID, creator, oracleaddress, "Verify its your address to delete oracle", signedData)
	if err != nil {
		log.Errorln("error identifying signer", err)
	}
	log.Infoln("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("invalid signer", signer)
		return
	}

	oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress, chainID)
	if err != nil {
		log.Errorln("error GetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, oracleconfig)

}
func (ob *Env) Pause(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	customerID := context.GetString(CUSTOMER_ID)
	var (
		// address string
		err          error
		oracleconfig dia.OracleConfig
	)
	oracleaddress := context.Query("oracleaddress")
	chainid := context.Query("oracleChainID")
	feederID := context.Query("feederID")

	if feederID != "" {
		oracleconfig, err = ob.RelDB.GetFeeder(feederID)
		if err != nil {
			log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on pause"})
			return
		}

	} else {
		oracleconfig, err = ob.RelDB.GetOracleConfig(oracleaddress, chainid)
		if err != nil {
			log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on pause"})
			return
		}
	}

	if oracleconfig.CustomerID != customerID {
		log.Errorf("Request ID: %s, not authorised to pause, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not part of your team oracle"})
		return
	}

	if oracleconfig.Deleted {
		log.Errorf("Request ID: %s, pause cannot be done one  deleted oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "this is deleted oracle"})
		return
	}
	if oracleconfig.Draft {
		log.Errorf("Request ID: %s, pause cannot be done one  draft oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "this is deleted oracle"})
		return
	}

	fc := &k8sbridge.FeederConfig{
		FeederID: oracleconfig.FeederID,
	}
	_, err = ob.k8sbridgeClient.DeletePod(context, fc)
	if err != nil {
		log.Errorf("Request ID: %s,  err DeleteOracleFeeder %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on pause"})
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, false)
	if err != nil {
		log.Errorf("Request ID: %s,  err ChangeOracleState %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on pause"})
		return
	}
	context.JSON(http.StatusOK, oracleconfig)
}

func (ob *Env) SetEcosystem(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	creator := context.GetString(CREATOR_ADDRESS)

	feederID := context.PostForm("feederID")
	enable := context.PostForm("enable")

	// customerID := context.GetString(CUSTOMER_ID)

	customer, err := ob.RelDB.GetCustomerByPublicKey(creator)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": "address not associated with any team"})
		context.Abort()
		return
	}

	log.Infof("Request ID: %s,  , customerID=%d, enable=%s %v ", requestId, customer.CustomerID, enable, feederID)

	enableBool, err := strconv.ParseBool(enable)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid boolean query parameter",
		})
		return
	}

	oracleconfig, err := ob.RelDB.GetFeeder(feederID)
	if err != nil {
		log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on pause"})
		return
	}

	if oracleconfig.CustomerID != strconv.Itoa(customer.CustomerID) {
		log.Errorf("Request ID: %s, not authorised to SetEcosystem, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customer.CustomerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not part of your team oracle"})
		return
	}

	err = ob.RelDB.ChangeEcosystemConfig(feederID, enableBool)
	if err != nil {
		log.Errorf("Request ID: %s,  ChangeEcosystemConfig err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error changing ecosystem state wallet"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "ecosystem state changed"})

}

func (ob *Env) Delete(context *gin.Context) {
	var (
		// address string
		err          error
		oracleconfig dia.OracleConfig
	)
	requestId := context.GetString(REQUEST_ID)
	customerID := context.GetString(CUSTOMER_ID)
	oracleaddress := context.Query("oracleaddress")
	chainID := context.Query("oracleChainID")

	oracleaddress = common.HexToAddress(oracleaddress).Hex()

	feederID := context.Query("feederID")

	if feederID != "" {
		oracleconfig, err = ob.RelDB.GetFeeder(feederID)
		if err != nil {
			log.Errorf("Request ID: %s,  Delete err GetFeeder %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on Delete"})
			return
		}

	} else {
		oracleconfig, err = ob.RelDB.GetOracleConfig(oracleaddress, chainID)
		if err != nil {
			log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on delete"})
			return
		}

	}

	log.Infof("Request ID: %s, delete oracle , deleted %s, isActive %s, draft %s  oracleconfig.CustomerID=%s %v ", requestId, oracleconfig.Deleted, oracleconfig.Active, oracleconfig.Draft, oracleconfig.CustomerID, err)

	if oracleconfig.CustomerID != customerID {
		log.Errorf("Request ID: %s, not authorised to delete, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not part of your team oracle"})
		return
	}

	if oracleconfig.Deleted {
		log.Errorf("Request ID: %s,already deleted oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "already deleted oracle"})
		return
	}
	// if oracleconfig.Draft {
	// 	log.Errorf("Request ID: %s, delete cannot be done on  draft oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
	// 	context.JSON(http.StatusUnauthorized, gin.H{"error": "this is draft oracle"})
	// 	return
	// }
	fc := &k8sbridge.FeederConfig{
		FeederID: oracleconfig.FeederID,
	}

	_, err = ob.k8sbridgeClient.DeletePod(context, fc)
	if err != nil {
		log.Errorf("Request ID: %s,  err DeleteOracleFeeder %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on delete"})
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, false)
	if err != nil {
		log.Errorf("Request ID: %s,  err ChangeOracleState %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on delete"})
		return
	}
	err = ob.RelDB.DeleteOracle(oracleconfig.FeederID)
	if err != nil {
		log.Errorf("Request ID: %s,  err DeleteOracle %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on delete"})
		return
	}
	context.JSON(http.StatusOK, oracleconfig)
}

func (ob *Env) Restart(context *gin.Context) {
	var (
		err          error
		oracleconfig dia.OracleConfig
	)
	requestId := context.GetString(REQUEST_ID)
	customerID := context.GetString(CUSTOMER_ID)

	oracleaddress := context.Query("oracleaddress")
	chainid := context.Query("oracleChainID")

	oracleaddress = common.HexToAddress(oracleaddress).Hex()
	feederID := context.Query("feederID")

	if feederID != "" {
		oracleconfig, err = ob.RelDB.GetFeeder(feederID)
		if err != nil {
			log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on restart"})
			return
		}

	} else {
		oracleconfig, err = ob.RelDB.GetOracleConfig(oracleaddress, chainid)
		if err != nil {
			log.Errorf("Request ID: %s,  err GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error on restart"})
			return
		}

	}

	if oracleconfig.CustomerID != customerID {
		log.Errorf("Request ID: %s, not authorised to Restart, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not part of your team oracle"})
		return
	}

	if oracleconfig.Deleted {
		log.Errorf("Request ID: %s, restart cannot be done one  deleted oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "this is deleted oracle"})
		return
	}

	fc := &k8sbridge.FeederConfig{
		FeederID:           oracleconfig.FeederID,
		Creator:            oracleconfig.Owner,
		FeederAddress:      oracleconfig.FeederAddress,
		Oracle:             oracleconfig.Address,
		ChainID:            oracleconfig.ChainID,
		Symbols:            strings.Join(oracleconfig.Symbols[:], ","),
		FeedSelection:      oracleconfig.FeederSelection,
		Blockchainnode:     oracleconfig.BlockchainNode,
		Frequency:          oracleconfig.Frequency,
		SleepSeconds:       oracleconfig.SleepSeconds,
		DeviationPermille:  oracleconfig.DeviationPermille,
		MandatoryFrequency: oracleconfig.MandatoryFrequency,
	}

	_, err = ob.k8sbridgeClient.RestartPod(context, fc)
	if err != nil {
		log.Errorf("Request ID: %s,  err RestartOracleFeeder %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on restart"})
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, true)
	if err != nil {
		log.Errorf("Request ID: %s,  err ChangeOracleState %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error on restart"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "feeder restarted"})

}

func getAuthToken(req *http.Request) (string, error) {
	authHeader := req.Header.Get("Authorization")
	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		return "", errors.New("bad authorization header")
	}
	token := authFields[1]
	return token, nil
}

func (ob *Env) CanRead(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	creator := context.Query("creator")
	if creator == "" {
		creator = context.PostForm("creator")
	}
	accessLevel, err := ob.RelDB.GetAccessLevel(creator)
	if err != nil {
		context.Abort()
		return
	}
	if accessLevel == "read" || accessLevel == "read_write" {

	} else {
		log.Errorf("Request ID: %s,  caller %s has no read access", requestId, creator)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "no read access"})
		context.Abort()
		return
	}

}
func (ob *Env) CanWrite(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	creator := context.Query("creator")
	if creator == "" {
		creator = context.PostForm("creator")
	}
	accessLevel, err := ob.RelDB.GetAccessLevel(creator)
	if err != nil {
		context.Abort()
		return
	}
	if accessLevel == "read_write" {

	} else {
		log.Errorf("Request ID: %s,  caller %s has no write access", requestId, creator)

		context.JSON(http.StatusUnauthorized, gin.H{"error": "no write access"})
		context.Abort()
		return
	}

}

func GenerateRandString() string {
	b := make([]byte, 10)
	rand.Read(b)
	return hex.EncodeToString(b)
}

func (ob *Env) Auth(isCustomer bool) gin.HandlerFunc {

	return func(context *gin.Context) {

		requestId := GenerateRandString()

		context.Set(REQUEST_ID, requestId)

		chainID := context.Query("chainID")

		creator := context.Query("creator")
		if creator == "" {
			creator = context.PostForm("creator")
		}

		if chainID == "" {
			chainID = context.PostForm("chainId")
			if chainID == "" {
				chainID = context.PostForm("chainID")

			}
		}
		oracleaddress := context.Query("oracleaddress")

		if oracleaddress == "" {
			oracleaddress = context.PostForm("oracleaddress")

		}

		if oracleaddress == "" {
			oracleaddress = creator
		} else {
			oracleaddress = common.HexToAddress(oracleaddress).Hex()
		}

		signedData, err := getAuthToken(context.Request)

		if err != nil {

			context.JSON(http.StatusUnauthorized, errors.New("sign err"))
			log.Errorln("missing auth token", err)
			context.Abort()
			return
		}
		actionMessage := context.GetString("message")
		log.Infof("Request ID: %s, ActionMessage: %s, chainID: %s, creator: %s, signedData: %s, oracleaddress: %s", requestId, actionMessage, chainID, creator, signedData, oracleaddress)

		signer, err := utils.GetSigner(chainID, creator, oracleaddress, actionMessage, signedData)

		if err != nil {
			log.Errorf("Request ID: %s, error GetSigner %v", requestId, err)
		}

		log.Infof("Request ID: %s signer: %s", requestId, signer)

		if signer.Hex() != creator {

			context.JSON(http.StatusUnauthorized, gin.H{"error": "invalid signer"})
			log.Errorf("Request ID: %s, invalid signer %s", requestId, signer)
			context.Abort()
			return

		}
		context.Set(CREATOR_ADDRESS, signer.Hex())

		if isCustomer {

			customer, err := ob.RelDB.GetCustomerByPublicKey(creator)
			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "address not associated with any team"})
				log.Errorf("Request ID: %s, address not associated with any team %s", requestId, signer)
				context.Abort()
				return
			}

			accessLevel, err := ob.RelDB.GetAccessLevel(creator)

			if err != nil {
				context.JSON(http.StatusUnauthorized, gin.H{"error": "address has invalid access level"})
				log.Errorf("Request ID: %s, address has invalid access level %s", requestId, signer)
				context.Abort()
				return

			}
			log.Infof("Request ID: %s, setting customer details %s customer id %d", requestId, CUSTOMER_ID, customer.CustomerID)

			context.Set(CUSTOMER_ID, customer.CustomerID)
			context.Set(CUSTOMER_PLAN, customer.CustomerPlan)
			context.Set(ACCESS_LEVEL, accessLevel)
			return
		}

	}

}

func contains(arr []string, str string) bool {
	for _, v := range arr {
		if v == str {
			return true
		}
	}
	return false
}

func (ob *Env) LoopPaymentStatus(context *gin.Context) {
	walletAddress := context.Query("wallet")

	status, err := ob.RelDB.GetLastPaymentByEndUser(walletAddress)
	if err != nil {
		context.JSON(http.StatusNotFound, nil)
		return
	}
	status.RedirectUrl = ob.RedirectURL
	context.JSON(http.StatusOK, status)

}

func (ob *Env) LoopWebHook(context *gin.Context) {
	requestId := GenerateRandString()

	signature := context.Request.Header.Get("loop-signature")

	body, err := io.ReadAll(context.Request.Body)
	if err != nil {
		return
	}
	ldr := models.LoopPaymentResponse{}
	err = json.Unmarshal(body, &ldr)
	if err != nil {
		context.JSON(http.StatusOK, nil)

	}

	l := verifyWebhook(body, ob.LoopPaymentSecret, signature)

	if l {
		log.Errorf("Request ID: %s,LoopWebHook verified ,", requestId)
	} else {
		log.Errorf("Request ID: %s,LoopWebHook not verified ,", requestId)

		return

	}

	switch ldr.Event {
	case "AgreementSignedUp":
		{
			// update user plan ldr.subscriber, plan name ldr.item
			// 1) find customer
			// 2) update plan
			// 3) update email
			//

			customer, err := ob.RelDB.GetCustomerByPublicKey(common.HexToAddress(ldr.Subscriber).String())
			if err != nil {
				log.Errorf("Request ID: %s,AgreementSignedUp GetCustomerByPublicKey %v ,", requestId, err)
				log.Errorf("Request ID: %s,sql.ErrNoRows %v ,", requestId, sql.ErrNoRows)
				log.Errorf("Request ID: %s,err %v ,", requestId, err)

				if err.Error() == "no rows in result set" {
					err := ob.RelDB.CreateCustomer(ldr.Email, "", 0, "", "", 2, []string{common.HexToAddress(ldr.Subscriber).String()})
					if err != nil {
						log.Errorf("Request ID: %s,customer err %v", err)

						context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
						return
					}
					log.Errorf("Request ID: %s,customer created ", requestId)
					customer, err := ob.RelDB.GetCustomerByPublicKey(common.HexToAddress(ldr.Subscriber).String())
					if err != nil {
						log.Errorf("Request ID: %s,AgreementCancelled GetCustomerByPublicKey %v ,", requestId, err)
					}
					log.Errorf("Request ID: %s,customer created %v", customer)

				} else {
					log.Errorf("Request ID: %s,AgreementSignedUp err GetCustomerByPublicKey %v ,", requestId, err)

				}

				//

			}
			log.Infof("Request ID: %s,AgreementSignedUp GetCustomerByPublicKey %v ,", requestId, customer)
			log.Infof("Request ID: %s,ldr  %v ,", requestId, ldr)

			ldr.Subscriber = common.HexToAddress(ldr.Subscriber).String()

			err = ob.RelDB.InsertLoopPaymentResponse(context, ldr)
			if err != nil {
				log.Errorf("Request ID: %s,AgreementSignedUp InsertLoopPaymentResponse %v ,", requestId, err)

			}

		}
	case "AgreementCancelled":

		{

			// cancel user plan
			customer, err := ob.RelDB.GetCustomerByPublicKey(common.HexToAddress(ldr.Subscriber).String())
			if err != nil {
				log.Errorf("Request ID: %s,AgreementCancelled GetCustomerByPublicKey %v ,", requestId, err)
			}
			log.Infof("Request ID: %s,AgreementCancelled customer %v ,", requestId, customer)
			log.Infof("Request ID: %s,AgreementCancelled ldr %v ,", requestId, ldr)

		}
	case "TransferCreated":
		{
			// create agreement
			log.Infof("Request ID: %s,TransferCreated ldr %v ,", requestId, ldr)

		}
	case "TransferProcessed":
		{

			var lptp models.LoopPaymentTransferProcessed

			err = json.Unmarshal(body, &lptp)
			if err != nil {
				context.JSON(http.StatusOK, nil)

			}

			err := ob.RelDB.InsertLoopPaymentTransferProcessed(context, lptp)
			if err != nil {
				log.Errorf("Request ID: %s, InsertLoopPaymentTransferProcessed %v ,", requestId, err)
				return
			}

			// Get customer
			customer, err := ob.RelDB.GetCustomerByPublicKey(common.HexToAddress(lptp.EndUser).String())
			if err != nil {
				log.Errorf("Request ID: %s, TransferProcessed GetCustomerByPublicKey %v ,", requestId, err)
				return
			}
			log.Infof("Request ID: %s,TransferProcessed GetCustomerByPublicKey %v ,", requestId, customer)
			log.Infof("Request ID: %s,ldr  %v ,", requestId, ldr)
			log.Infof("Request ID: %s,AgreementID  %s ,", requestId, ldr.AgreementID)

			// get agreement

			lpr, err := ob.RelDB.GetLoopPaymentResponseByAgreementID(context, ldr.AgreementID)
			if err != nil {
				log.Errorf("Request ID: %s, GetLoopPaymentResponseByAgreementID %v ,", requestId, err)
			}
			log.Infof("Request ID: %s,lpr  %v ", requestId, lpr)
			log.Infof("Request ID: %s,lpr ItemID  %v ", requestId, lpr.Item)

			var planId = 0

			switch lpr.Item {

			case "Product 1":
				planId = 1
			case "Product 2":
				planId = 2

			}

			log.Infof("Request ID: %s, customerID %d,", requestId, customer.CustomerID)
			log.Infof("Request ID: %s, planId %d,", requestId, planId)
			log.Infof("Request ID: %s, lpr.PaymentTokenSymbol %s,", requestId, lpr.PaymentTokenSymbol)
			log.Infof("Request ID: %s, lpr.EventDate %s,", requestId, lpr.EventDate)

			err = ob.RelDB.UpdateCustomerPlan(context, customer.CustomerID, planId, lpr.PaymentTokenSymbol, strconv.Itoa(lpr.EventDate))

			if err != nil {
				log.Errorf("Request ID: %s, UpdateCustomerPlan %v, customerID %d,", requestId, err, customer.CustomerID)

			} else {
				log.Infof("Request ID: %s, UpdateCustomerPlan %v customerID %d, Plan ID %d,", requestId, err, customer.CustomerID)

			}

		}
	}

	context.JSON(http.StatusOK, nil)

}

func verifyWebhook(body []byte, loopSecret, signature string) bool {
	// secret := "745c0448-6bb9-4927-9bcb-6556fb3bdd6e"
	h := hmac.New(sha256.New, []byte(loopSecret))
	h.Write(body)
	expectedSignature := base64.StdEncoding.EncodeToString(h.Sum(nil))
	return expectedSignature == signature
}
