package main

import (
	"encoding/json"
	"errors"
	"fmt"
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
	RelDB                   *models.RelDB
	k8sbridgeClient         k8sbridge.K8SHelperClient
	Keyring                 kr.Keyring
	RateLimitOracleCreation int
	SupportedChains         []string
}

func NewEnv(relStore *models.RelDB, ds models.Datastore, kc k8sbridge.K8SHelperClient, ring kr.Keyring, rateLimitOracleCreation int) *Env {

	return &Env{
		DataStore:               ds,
		RelDB:                   relStore,
		Keyring:                 ring,
		k8sbridgeClient:         kc,
		RateLimitOracleCreation: rateLimitOracleCreation,
		SupportedChains:         []string{"11155111", "421614", "123420111", "94204209", "88153591557", "80002"},
	}
}

func handleError(context *gin.Context, status int, errorMsg, logMsg string, logArgs ...interface{}) {
	context.JSON(status, errors.New(errorMsg))
	log.Errorf(logMsg, logArgs...)
	context.Abort() // Prevent further handlers from being called
}

// Create new oracle feeder if creator has resources
func (ob *Env) Create(context *gin.Context) {

	var (
		address  string
		err      error
		keypair  *k8sbridge.KeyPair
		isUpdate bool
	)

	isUpdate = false

	oracleaddress := context.PostForm("oracleaddress")
	oracleaddress = common.HexToAddress(oracleaddress).Hex()
	chainID := context.PostForm("chainID")
	creator := context.PostForm("creator")
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

	k := make(map[string]string)

	log.Infof("Creating oracle: oracleAddress: %s, ChainID: %s, Creator: %s, Symbols: %s, frequency: %s, sleepSeconds: %s blockchainnode: %s, feedSelection %s", oracleaddress, chainID, creator, symbols, frequency, sleepSeconds, blockchainnode, feedSelection)

	signer, err := utils.GetSigner(chainID, creator, oracleaddress, "Verify its your address to call oracle builder", signedData)
	if err != nil {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer")
	}
	log.Infoln("Creating oracle: signer", signer)

	if signer.Hex() != creator {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer %v", signer)
	}

	// validations
	// check for  symbols and feedSelection

	if symbols == "" && feedSelection == "" {
		handleError(context, http.StatusBadRequest, "no symbols", "Creating oracle: no symbols or feedSelection %v", symbols)

	}
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

	err = ob.RelDB.SetOracleConfig(oracleaddress, feederID, creator, address, symbols, feedSelection, chainID, frequency, sleepSeconds, deviationPermille, blockchainnode, mandatoryFrequency)
	if err != nil {
		log.Errorln("error SetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
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

	k["oracleaddress"] = oracleaddress
	k["chainId"] = chainID
	k["creator"] = creator
	k["symbols"] = symbols
	k["publicKey"] = address

	context.JSON(http.StatusCreated, k)
}

// List: list owner oracles
func (ob *Env) List(context *gin.Context) {
	creator := context.Query("creator")

	oracles, err := ob.RelDB.GetOraclesByOwner(creator)
	if err != nil {
		errorMsg := "Error fetching oracles by owner"
		logMsg := "List Oracles: error on getOraclesByOwner"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
	}
	context.JSON(http.StatusOK, oracles)
}

// List whitelisted addresses
func (ob *Env) Whitelist(context *gin.Context) {
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

		if ob.SupportedChains == chainInfo.ChainID {
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
	var (
		// address string
		err error
	)
	oracleaddress := context.Query("oracleaddress")
	chainid := context.Query("oracleChainID")

	creator := context.Query("creator")

	oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress, chainid)
	if err != nil {
		log.Errorln("error GetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if oracleconfig.Owner != creator {
		log.Errorln("not authorised to delete  ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	fc := &k8sbridge.FeederConfig{
		FeederID: oracleconfig.FeederID,
	}
	_, err = ob.k8sbridgeClient.DeletePod(context, fc)
	if err != nil {
		log.Errorln("error DeleteOracleFeeder ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, false)
	if err != nil {
		log.Errorln("error ChangeOracleState ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, oracleconfig)
}

func (ob *Env) Delete(context *gin.Context) {
	var (
		// address string
		err error
	)
	oracleaddress := context.Query("oracleaddress")
	chainID := context.Query("oracleChainID")

	oracleaddress = common.HexToAddress(oracleaddress).Hex()

	creator := context.Query("creator")

	oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress, chainID)
	if err != nil {
		log.Errorln("error GetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if oracleconfig.Owner != creator {
		log.Errorln("not authorised to delete  ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	fc := &k8sbridge.FeederConfig{
		FeederID: oracleconfig.FeederID,
	}
	_, err = ob.k8sbridgeClient.DeletePod(context, fc)
	if err != nil {
		log.Errorln("error DeleteOracleFeeder ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, false)
	if err != nil {
		log.Errorln("error ChangeOracleState ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ob.RelDB.DeleteOracle(oracleconfig.FeederID)
	if err != nil {
		log.Errorln("error DeleteOracle ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, oracleconfig)
}

func (ob *Env) Restart(context *gin.Context) {
	var (
		err error
	)
	oracleaddress := context.Query("oracleaddress")
	chainid := context.Query("oracleChainID")

	oracleaddress = common.HexToAddress(oracleaddress).Hex()

	creator := context.Query("creator")

	oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress, chainid)
	if err != nil {
		log.Errorln("error GetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if oracleconfig.Owner != creator {
		log.Errorln("not authorised to delete  ", err)
		context.JSON(http.StatusInternalServerError, err)
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
		log.Errorln("error RestartOracleFeeder ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	err = ob.RelDB.ChangeOracleState(oracleconfig.FeederID, true)
	if err != nil {
		log.Errorln("error ChangeOracleState ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

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
		context.JSON(http.StatusUnauthorized, errors.New("no read access"))
		context.Abort()
		return
	}

}
func (ob *Env) CanWrite(context *gin.Context) {
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
		context.JSON(http.StatusUnauthorized, errors.New("no write access"))
		context.Abort()
		return
	}

}

func (ob *Env) Auth(context *gin.Context) {

	chainID := context.Query("chainID")
	creator := context.Query("creator")
	if creator == "" {
		creator = context.PostForm("creator")
	}

	if chainID == "" {
		chainID = context.PostForm("chainId")
	}
	oracleaddress := context.Query("oracleaddress")

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
	actionmessage := context.GetString("message")
	log.Infoln("actionmessage", actionmessage)
	log.Infoln("chainID", chainID)
	log.Infoln("creator", creator)
	log.Infoln("signedData", signedData)
	log.Infoln("oracleaddress", oracleaddress)

	signer, err := utils.GetSigner(chainID, creator, oracleaddress, actionmessage, signedData)

	if err != nil {
		log.Errorf("error while signign %v", err)
	}

	log.Infoln("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("invalid signer", signer)
		context.Abort()
		return

	}

}
