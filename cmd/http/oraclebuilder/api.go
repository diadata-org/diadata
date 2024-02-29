package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"sync"
	"time"

	"strings"

	builderUtils "github.com/diadata-org/diadata/http/oraclebuilder/utils"
	"github.com/ethereum/go-ethereum/common"

	kr "github.com/99designs/keyring"
	"github.com/99designs/keyring/cmd/k8sbridge"
	"github.com/diadata-org/diadata/pkg/dia"
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
	PodHelper               *builderUtils.PodHelper
	Keyring                 kr.Keyring
	RateLimitOracleCreation int
	apiStats                APIStats
}

func NewEnv(relStore *models.RelDB, ds models.Datastore, ph *builderUtils.PodHelper, ring kr.Keyring, rateLimitOracleCreation int) *Env {
	apistats := *NewAPIStats()
	go apistats.ClearCachePeriodically(10 * time.Minute)

	return &Env{
		DataStore:               ds,
		RelDB:                   relStore,
		PodHelper:               ph,
		Keyring:                 ring,
		RateLimitOracleCreation: rateLimitOracleCreation,
		apiStats:                apistats,
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
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer", signer)
	}

	// validations
	// check for  symbols and feedSelection

	if symbols == "" && feedSelection == "" {
		handleError(context, http.StatusBadRequest, "no symbols", "Creating oracle: no symbols or feedSelection", symbols)

	}
	symbolsArray := strings.Split(symbols, ",")

	if len(symbolsArray) > 10 {
		handleError(context, http.StatusBadRequest, "max symbols exceed", "Creating oracle: max symbols exceed", symbols)
	}

	// check for duplicate symbol

	if utils.CheckDuplicates(symbolsArray) {
		handleError(context, http.StatusBadRequest, "duplicate symbols", "Creating oracle: duplicate symbols", symbols)

	}

	// check frequency limit

	frequencyInt, err := strconv.Atoi(frequency)
	if err != nil {
		handleError(context, http.StatusBadRequest, "invalid frequency", "Creating oracle: invalid frequency", err)
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
			handleError(context, http.StatusBadRequest, "invalid mandatoryFrequencyInt, out of range", "Creating oracle: invalid mandatoryFrequencyInt, out of range", err)
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
		count, isAllowed := ob.createOracleFeederLimiter(false)
		log.Infoln("Rate limit check, isAllowed", count, isAllowed)

		if !isAllowed {
			log.Errorln("Rate limit exceeded for createOracleFeeder", err)
			context.JSON(http.StatusTooManyRequests, "Rate limit exceeded for createOracleFeeder")
			return
		}
		// check if creator has resources to create new oracle feeder
		// limit := ob.RelDB.GetFeederLimit(creator)
		// total := ob.RelDB.GetTotalFeeder(creator)

		// log.Infof("Creating oracle: Feeders Limit %d, Total Feeders:%d, Creator: %s", limit, total, creator)
		// if total >= limit {
		// 	log.Errorln("not enought resource left ", creator)
		// 	context.JSON(http.StatusUnauthorized, errors.New("limit over"))
		// 	return
		// }

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
		err = ob.PodHelper.CreateOracleFeeder(context, feederID, creator, address, oracleaddress, chainID, symbols, feedSelection, blockchainnode, frequency, sleepSeconds, deviationPermille, mandatoryFrequency)
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

		err = ob.PodHelper.RestartOracleFeeder(context, feederID, oracleconfig)
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

func (ob *Env) createOracleFeederLimiter(isCheck bool) (int, bool) {

	count, exist := ob.apiStats.Get("createOracleFeeder")

	if !exist {
		ob.apiStats.Set("createOracleFeeder", 0)
		return count, true
	}
	if count >= ob.RateLimitOracleCreation {
		return count, false
	}
	if !isCheck {
		count++
		ob.apiStats.Set("createOracleFeeder", count)
	}

	return count, true

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
	page := context.Query("page")

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

	totalUpdates, err := ob.RelDB.GetOracleUpdateCount(address, chainID)
	if err != nil {
		errorMsg := "Error fetching oracle update count"
		logMsg := "Oracle Stats error GetOracleUpdateCount"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	updates, err := ob.RelDB.GetOracleUpdates(address, chainID, offset)
	if err != nil {
		errorMsg := "Error fetching oracle updates"
		logMsg := "Oracle Stats error"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	type localAssetConfig struct {
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
	fmt.Println("oracleConfigcache", oracleConfig)
	fmt.Println("err", err)

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
			fmt.Println("unmarshal :", err)
			return
		}

	}

	var (
		assets                []localAssetConfig
		avgGasSpend, gasSpend float64
		updateFeeds           []dia.FeedUpdates
	)

	if len(symbolFeeds) > 0 {

		for _, sf := range symbolFeeds {
			var blockchain string
			for _, feeds := range sf.FeedSelection {
				blockchain = blockchain + feeds.Blockchain

			}

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

			lastUpdate, err := ob.RelDB.GetOracleUpdates(address, chainID, offset)
			if err == nil && len(lastUpdate) > 0 {
				LastReportedPrice = lastUpdate[0].AssetPrice
				LastReportedTime = lastUpdate[0].UpdateTime

			}

			volume24h, err := ob.RelDB.GetLastAssetVolume24H(dia.Asset{Blockchain: blockchain, Address: strings.Split(sf.Symbol, "-")[1]})
			if err != nil {
				volume24h = 0
			}
			asset := localAssetConfig{Address: strings.Split(sf.Symbol, "-")[1], Blockchain: blockchain, Deviation: oracleConfig.DeviationPermille, HeartBeat: oracleConfig.Frequency, GQLQuery: query, Volume24h: volume24h, LastReportedPrice: LastReportedPrice, LastReportedTime: LastReportedTime, SymbolFeed: sf}
			assets = append(assets, asset)

		}
	} else {

		var (
			LastReportedPrice string
			LastReportedTime  time.Time
		)

		lastUpdate, err := ob.RelDB.GetOracleUpdates(address, chainID, offset)
		if err == nil && len(lastUpdate) > 0 {
			LastReportedPrice = lastUpdate[0].AssetPrice
			LastReportedTime = lastUpdate[0].UpdateTime

		}

		for _, symbol := range oracleConfig.Symbols {

			blockchain := strings.TrimSpace(strings.Split(symbol, "-")[0])
			address := strings.TrimSpace(strings.Split(symbol, "-")[1])

			volume24h, err := ob.RelDB.GetLastAssetVolume24H(dia.Asset{Blockchain: blockchain, Address: address})
			if err != nil {
				volume24h = 0
			}
			asset := localAssetConfig{Address: address, Blockchain: blockchain, Deviation: oracleConfig.DeviationPermille, HeartBeat: oracleConfig.Frequency, Volume24h: volume24h, LastReportedPrice: LastReportedPrice, LastReportedTime: LastReportedTime}
			assets = append(assets, asset)

		}

	}

	gasSpend, err = ob.RelDB.GetTotalGasSpend(address, chainID, 0)
	if err != nil {
		gasSpend = 0.0
	}
	updateFeeds, avgGasSpend, err = ob.RelDB.GetDayWiseUpdates(address, chainID)
	if err != nil {
		avgGasSpend = 0.0
	}

	response := make(map[string]interface{})
	response["OracleAddress"] = address
	response["Chain"] = chainID
	response["OracleType"] = 0
	response["GasRemaining"] = 0
	response["GasSpend"] = gasSpend
	response["AvgGasSpend"] = avgGasSpend
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

	totalUpdates, err := ob.RelDB.GetOracleUpdateCount(address, chainID)
	if err != nil {
		errorMsg := "Error fetching oracle update count"
		logMsg := "Oracle Stats error GetOracleUpdateCount"
		handleError(context, http.StatusInternalServerError, errorMsg, logMsg, err)
		return
	}

	updates, err := ob.RelDB.GetOracleUpdates(address, chainID, offset)
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
	log.Infoln("signedData", signedData)

	if err != nil {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("missing auth token", err)
		return
	}

	signer, _ := utils.GetSigner(chainID, creator, oracleaddress, "Verify its your address to delete oracle", signedData)

	log.Infoln("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("invalid signer", signer)
		return
	}
	// creator := context.PostForm("creator")

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
	err = ob.PodHelper.DeleteOracleFeeder(context, oracleconfig.FeederID)
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

func (ob *Env) ViewLimit(context *gin.Context) {

	response := make(map[string]interface{})
	count, isAllowed := ob.createOracleFeederLimiter(true)
	response["Count"] = count
	response["IsAllowed"] = isAllowed

	context.JSON(http.StatusOK, response)
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
	err = ob.PodHelper.DeleteOracleFeeder(context, oracleconfig.FeederID)
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

	err = ob.PodHelper.RestartOracleFeeder(context, oracleconfig.FeederID, oracleconfig)
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
	log.Println("authHeader", authHeader)
	authFields := strings.Fields(authHeader)
	if len(authFields) != 2 || strings.ToLower(authFields[0]) != "bearer" {
		return "", errors.New("bad authorization header")
	}
	token := authFields[1]
	return token, nil
}

func (ob *Env) Auth(context *gin.Context) {

	chainID := context.Query("chainID")
	creator := context.Query("creator")
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
		log.Error("error while signign %v", err)
	}

	log.Infoln("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		log.Errorln("invalid signer", signer)
		context.Abort()
		return

	}

}

type APIStats struct {
	data map[string]int
	mu   sync.RWMutex
}

func NewAPIStats() *APIStats {
	return &APIStats{
		data: make(map[string]int),
	}
}

func (c *APIStats) Set(key string, value int) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}

func (c *APIStats) Get(key string) (int, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	value, exists := c.data[key]
	return value, exists
}

func (c *APIStats) ClearCachePeriodically(interval time.Duration) {
	ticker := time.NewTicker(interval)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			c.mu.Lock()
			for k := range c.data {
				delete(c.data, k)
			}
			c.mu.Unlock()
		}
	}
}
