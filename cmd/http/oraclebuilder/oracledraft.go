package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	kr "github.com/99designs/keyring"
	k8sbridge "github.com/diadata-org/diadata/pkg/dia/helpers/k8sbridge/protoc"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
)

func (ob *Env) CreateUpdateOracle(context *gin.Context) {
	requestId := GenerateRandString()
	context.Set(REQUEST_ID, requestId)

	oracleaddress := context.PostForm("oracleaddress")
	oracleaddress = common.HexToAddress(oracleaddress).Hex()
	chainID := context.PostForm("chainID")
	creator := context.PostForm("creator")
	oracleName := context.PostForm("name")
	feederID := context.PostForm("feederID")

	creator = common.HexToAddress(creator).Hex()

	symbols := context.PostForm("symbols")
	signedData := context.PostForm("signeddata")
	frequency := context.PostForm("frequency")
	sleepSeconds := context.PostForm("sleepseconds")
	deviationPermille := context.PostForm("deviationpermille")

	blockchainnode := context.PostForm("blockchainnode")
	feedSelection := context.PostForm("feedselection")

	mandatoryFrequency := context.PostForm("mandatoryfrequency")

	log.Infof("Request ID: %s Creating  draft oracle: oracleAddress: %s, ChainID: %s, Creator: %s, Symbols: %s, frequency: %s, sleepSeconds: %s blockchainnode: %s, feedSelection %s name %s", requestId, oracleaddress, chainID, creator, symbols, frequency, sleepSeconds, blockchainnode, feedSelection, oracleName)

	signer, err := utils.GetSigner(chainID, creator, oracleaddress, "Verify its your address to create and update oracle", signedData)
	if err != nil {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer")
		return

	}
	log.Infof("Request ID: %s Creating oracle: signer %v", requestId, signer.Hex())

	if signer.Hex() != creator {
		handleError(context, http.StatusUnauthorized, "sign err", "Creating oracle: invalid signer %v", signer)
	}

	// validations
	// check for  symbols and feedSelection

	if feedSelection == "" {
		handleError(context, http.StatusBadRequest, "no feedSelection", "Creating oracle: no symbols or feedSelection %v", symbols)
		log.Errorf("Request ID: %s, empty feedSelection %v ,", requestId, err)
		return

	}

	// get customer id
	customer, err := ob.RelDB.GetCustomerByPublicKey(signer.Hex())
	if err != nil {
		log.Errorf("Request ID: %s, GetCustomerByPublicKey %v ,", requestId, err)
		handleError(context, http.StatusNotFound, "error getting customer", "Creating oracle:account not found for this wallet")
		return

	}

	customerID := customer.CustomerID

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
		return
	}

	mandatoryFrequencyInt, err := strconv.Atoi(mandatoryFrequency)
	if err != nil {
		mandatoryFrequencyInt = 0
		// handleError(context, http.StatusBadRequest, "invalid mandatoryFrequencyInt", "Creating oracle: invalid mandatoryFrequencyInt", err)
	}

	if frequencyInt != 0 || mandatoryFrequencyInt == 0 {
		if frequencyInt < 120 || frequencyInt > 2630000 {
			handleError(context, http.StatusBadRequest, "invalid frequency, out of range 120 - 2630000", "Creating oracle: out of range frequency %v", err)
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

	feederAddress := common.HexToAddress("0x000000000000000000000000000000000000dead").String()

	isNewFeeder := false
	if feederID == "" {
		isNewFeeder = true
		feederID = "feeder-" + utils.GenerateAutoname("-")

	} else {

		// check if caller of Request really own this feeder id

		//TODO recheck logic
		oc, err := ob.RelDB.GetFeeder(feederID)

		if err != nil {
			log.Errorf("RequestId: %s, error GetOracleConfig %v ", requestId, err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "no access to feederID"})
		}

		if oc.CustomerID != strconv.Itoa(customerID) {
			log.Errorf("RequestId: %s, invalid feeder id %v ", requestId, err)
			context.JSON(http.StatusUnauthorized, gin.H{"error": "no access to feederID"})

		}

		feederAddress = oc.FeederAddress

	}

	// check if this exists, restart it

	err = ob.RelDB.SetOracleConfig(context, strconv.Itoa(customerID), oracleaddress, feederID, creator, feederAddress, symbols, feedSelection, chainID, frequency, sleepSeconds, deviationPermille, blockchainnode, mandatoryFrequency, oracleName, true, false)
	if err != nil {
		log.Errorf("RequestId: %s, error SetOracleConfig %v", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error creating/updating new oracle config"})
		return
	}

	log.Infof("Created oracle: oracleAddress: %s, ChainID: %s, Creator: %s, Symbols: %s, frequency: %s, sleepSeconds: %s, Feeder ID :%s,", oracleaddress, chainID, creator, symbols, frequency, sleepSeconds, feederID)

	if isNewFeeder {
		err = ob.RelDB.ChangeOracleState(feederID, false)
		if err != nil {
			log.Errorf("RequestId: %s, error ChangeOracleState %v", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error creating/updating new oracle config"})
			return
		}
	}
	response := make(map[string]string)

	response["oracleaddress"] = oracleaddress
	response["chainId"] = chainID
	response["creator"] = creator
	response["symbols"] = symbols
	response["feederID"] = feederID

	// response["publicKey"] = address

	context.JSON(http.StatusCreated, response)

}

func (ob *Env) isValidPlan(context *gin.Context, customer *models.Customer, requestId string) bool {
	plan, err := ob.RelDB.GetPlan(context, customer.CustomerPlan)

	if err != nil {
		log.Errorf("Request ID: %s, GetPlan %v ,", requestId, err)
		return false

	}

	log.Infof("Request ID: %s, Plan %v ,", requestId, plan)

	/* total Feeds used

	 */
	totalFeeds, totalOracles, err := ob.billableResource(strconv.Itoa(customer.CustomerID))
	if err != nil {
		log.Errorf("Request ID: %s, Total feeds err %v ,", requestId, err)

	}
	if totalFeeds >= plan.TotalFeeds {
		log.Errorf("Request ID: %s, totalFeeds exceeds plan Limit %v ,", requestId, err)
		// handleError(context, http.StatusPaymentRequired, "totalFeeds exceeds plan Limit", "Creating oracle:  totalFeeds exceeds plan Limit")
		return false
	}
	if totalOracles >= plan.TotalOracles {
		log.Errorf("Request ID: %s, totalOracles exceeds plan Limit %v ,", requestId, err)
		// handleError(context, http.StatusPaymentRequired, "totalFeeds exceeds plan Limit", "Creating oracle:  totalFeeds exceeds plan Limit")
		return false
	}

	log.Infof("Request ID: %s, Total feeds %d,", requestId, totalFeeds)
	return true
}

func (ob *Env) InitFeeder(context *gin.Context) {

	var (
		// address string
		err     error
		keypair *k8sbridge.KeyPair
	)
	requestId := context.GetString(REQUEST_ID)
	customerID := context.GetString(CUSTOMER_ID)
	creator := context.GetString(CREATOR_ADDRESS)

	feederID := context.PostForm("feederID")

	// chainID := context.Query("oracleChainID")

	// oracleaddress = common.HexToAddress(oracleaddress).Hex()

	log.Infof("Request ID: %s, deploy feeder, customerID=%s,   ", requestId, customerID)

	oracleconfig, err := ob.RelDB.GetFeeder(feederID)
	if err != nil {
		log.Errorf("Request ID: %s,  InitFeeder err GetOracleConfig %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "feeder with this id not found"})
		return
	}

	if oracleconfig.Deleted {
		log.Errorf("Request ID: %s,   deploy cannot be done one  deleted oracle, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "this is deleted oracle"})
		return
	}

	customer, err := ob.RelDB.GetCustomerByPublicKey(creator)
	if err != nil {
		log.Errorf("Request ID: %s, GetCustomerByPublicKey %v ,", requestId, err)
		handleError(context, http.StatusNotFound, "error getting customer", "Creating oracle:account not found for this wallet")
		return

	}

	if oracleconfig.CustomerID != strconv.Itoa(customer.CustomerID) {
		log.Errorf("Request ID: %s, not authorised to Initialise feeder, customerID=%s, oracleconfig.CustomerID=%s %v ", requestId, customerID, oracleconfig.CustomerID, err)
		context.JSON(http.StatusUnauthorized, gin.H{"error": "not part of your team oracle"})
		return
	}

	if !ob.isValidPlan(context, customer, requestId) {
		log.Errorf("Request ID: %s, isValidPlan  %v ,", requestId, err)
		handleError(context, http.StatusNotFound, "plan exhausted or not subscribed", "subscribe for plan")
		return
	}

	// create new feeder wallet

	// check if it was already deployed

	// item, err := ob.Keyring.Get(oracleconfig.FeederID)
	// if err != nil {
	// 	log.Errorf("Request ID: %s, ob.Keyring.Get feederId: %s err  %v ,", requestId, oracleconfig.FeederID, err)
	// }

	// log.Infof("Request ID: %s, creating keys ob.Keyring.Set item.Data: %v for feeder  %s ,", requestId, item.Key, oracleconfig.FeederID)

	isRestart := true

	// if item.Data == nil {
	// create new

	log.Infof("Request ID: %s, feederId: %s oracleconfig.Active  %v ,", requestId, oracleconfig.FeederID, oracleconfig.Active)

	if !oracleconfig.Active {

		isRestart = false

		err = ob.Keyring.Set(kr.Item{
			Key: oracleconfig.FeederID,
		})

	}

	if err != nil {
		log.Errorf("Request ID: %s, error generating key  %v ,", requestId, err)
		context.JSON(http.StatusInternalServerError, errors.New("error generating key"))
		return
	}
	log.Infof("Request ID: %s, creating keys ob.Keyring.Set feederId: %s err  %v ,", requestId, oracleconfig.FeederID, err)

	// Get public key

	item, err := ob.Keyring.Get(oracleconfig.FeederID)
	if err != nil {
		log.Errorf("Request ID: %s, error getting key  %v ,", requestId, err)
		context.JSON(http.StatusInternalServerError, errors.New("error getting key"))
		return
	}
	marshalErr := json.Unmarshal(item.Data, &keypair)
	if marshalErr != nil {
		log.Errorf("Request ID: %s, error marshalErr  %v ,", requestId, marshalErr)
		context.JSON(http.StatusInternalServerError, errors.New("error getting key"))
		return
	}
	log.Infof("Request ID: %s, feeder public key  %v ,", requestId, keypair.GetPublickey())

	publicKey := keypair.GetPublickey()

	fc := &k8sbridge.FeederConfig{
		FeederID:      oracleconfig.FeederID,
		Creator:       creator,
		FeederAddress: publicKey,
		Oracle:        oracleconfig.Address,
		ChainID:       oracleconfig.ChainID,
		// Symbols:            oracleconfig.Symbols,
		FeedSelection:      oracleconfig.FeederSelection,
		Blockchainnode:     oracleconfig.BlockchainNode,
		Frequency:          oracleconfig.Frequency,
		SleepSeconds:       oracleconfig.SleepSeconds,
		DeviationPermille:  oracleconfig.DeviationPermille,
		MandatoryFrequency: oracleconfig.MandatoryFrequency,
	}

	if isRestart {
		_, err = ob.k8sbridgeClient.RestartPod(context, fc)
		if err != nil {
			log.Errorln("error RestartPod ", err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error starting feeder"})
			return
		}

	} else {
		_, err = ob.k8sbridgeClient.CreatePod(context, fc)
		if err != nil {
			log.Errorf("Request ID: %s, error CreatePod  %v ,", requestId, err)
			context.JSON(http.StatusInternalServerError, gin.H{"error": "error starting feeder"})
			return
		}
	}

	err = ob.RelDB.SetOracleConfig(context, strconv.Itoa(customer.CustomerID), oracleconfig.Address, oracleconfig.FeederID, creator, publicKey, "", oracleconfig.FeederSelection, oracleconfig.ChainID, oracleconfig.Frequency, oracleconfig.SleepSeconds, oracleconfig.DeviationPermille, oracleconfig.BlockchainNode, oracleconfig.MandatoryFrequency, oracleconfig.Name, false, true)
	if err != nil {
		log.Errorln("requestId: %s, error SetOracleConfig ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error starting feeder"})
		return
	}

	err = ob.RelDB.ChangeOracleState(feederID, true)
	if err != nil {
		log.Errorf("RequestId: %s, error ChangeOracleState %v", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error starting feeder"})
		return
	}
	response := make(map[string]interface{})

	response["oracleaddress"] = oracleconfig.Address
	response["chainId"] = oracleconfig.ChainID
	response["creator"] = creator
	response["symbols"] = oracleconfig.Symbols
	response["publicKey"] = publicKey

	context.JSON(http.StatusCreated, response)

}
