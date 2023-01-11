package oraclebuilder

import (
	"encoding/json"
	"errors"
	"net/http"
	builderutils "oraclebuilder/utils"

	kr "github.com/99designs/keyring"
	k8sbridge "github.com/99designs/keyring/cmd/k8sbridge"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

/*

Auth using EIP712 spec
*/

type Env struct {
	DataStore models.Datastore
	RelDB     *models.RelDB
	PodHelper *builderutils.PodHelper
	Keyring   kr.Keyring
}

var log = logrus.New()

func (ob *Env) Create(context *gin.Context) {

	var (
		address string
		err     error
	)

	oracleaddress := context.PostForm("oracleaddress")
	chainID := context.PostForm("chainID")
	creator := context.PostForm("creator")
	symbols := context.PostForm("symbols")
	signedData := context.PostForm("signeddata")
	feederID := context.PostForm("feederID")
	frequency := context.PostForm("frequency")
	sleepSeconds := context.PostForm("sleepseconds")
	deviationPermille := context.PostForm("deviationpermille")

	blockchainnode := context.PostForm("blockchainnode")

	k := make(map[string]string)

	log.Infoln("oracleaddress", oracleaddress)
	log.Infoln("chainId", chainID)
	log.Infoln("creator", creator)
	log.Infoln("symbols", symbols)
	log.Infoln("signeddata", signedData)
	log.Infoln("feederId", feederID)
	log.Infoln("frequency", frequency)
	log.Infoln("sleepSeconds", sleepSeconds)

	signer, _ := utils.GetSigner(chainID, creator, oracleaddress, signedData)

	log.Infoln("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		return
	}

	log.Infoln("feederId from creator", feederID)

	if feederID == "" {
		limit := ob.RelDB.GetFeederLimit(creator)
		total := ob.RelDB.GetTotalFeeder(creator)

		log.Infoln("limit", limit)
		log.Infoln("total", total)

		if total >= limit {
			log.Errorln("not enought resource left ", creator)
			context.JSON(http.StatusUnauthorized, errors.New("limit over"))
			return
		}

		feederID = utils.GenerateAutoname("-")

		err = ob.Keyring.Set(kr.Item{
			Key: feederID,
		})

		if err != nil {
			log.Errorln("error getting key", err)
			context.JSON(http.StatusUnauthorized, errors.New("need access to this feeder"))
			return
		}

		var keypair *k8sbridge.KeyPair

		item, err := ob.Keyring.Get(feederID)
		if err != nil {
			log.Infoln("error getting key", err)
			context.JSON(http.StatusInternalServerError, errors.New("error getting key"))
			return
		}
		json.Unmarshal(item.Data, &keypair)
		log.Infoln("public key", keypair.GetPublickey())
		address = keypair.GetPublickey()

		err = ob.PodHelper.CreateOracleFeeder(feederID, address, oracleaddress, chainID, symbols, blockchainnode, frequency, sleepSeconds, deviationPermille)
		if err != nil {
			log.Errorln("error CreateOracleFeeder ", err)
			context.JSON(http.StatusInternalServerError, errors.New("error creating oraclefeeder"))
			return
		}

		err = ob.RelDB.SetOracleConfig(oracleaddress, feederID, creator, symbols, chainID)
		if err != nil {
			log.Errorln("error SetOracleConfig ", err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}

	}

	log.Infoln("owneraddress GetFeederAccessByID", address)

	k["oracleaddress"] = oracleaddress
	k["chainId"] = chainID
	k["creator"] = creator
	k["symbols"] = symbols
	k["publicKey"] = address

	context.JSON(http.StatusCreated, k)
}

// list owner oracles
func (ob *Env) List(context *gin.Context) {
	creator := context.Query("creator")
	oracles, err := ob.RelDB.GetOraclesByOwner(creator)
	if err != nil {
		log.Errorln("error GetOraclesByOwner ", err)
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
	oracleaddress := context.Query("oracleaddress")
	// creator := context.PostForm("creator")

	oracleconfig, err := ob.RelDB.GetOracleConfig(oracleaddress)
	if err != nil {
		log.Errorln("error GetOracleConfig ", err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	context.JSON(http.StatusOK, oracleconfig)

}
