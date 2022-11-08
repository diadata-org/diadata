package oraclebuilder

import (
	"errors"
	"fmt"
	"log"
	"net/http"

	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/gin-gonic/gin"
)

/*

Auth using EIP712 spec
*/

type Env struct {
	DataStore models.Datastore
	RelDB     *models.RelDB
}

func (ob *Env) InitiateOracle(context *gin.Context) {

	var (
		address    string
		privatekey string
	)

	oracleaddress := context.PostForm("oracleaddress")
	chainId := context.PostForm("chainId")
	creator := context.PostForm("creator")
	symbols := context.PostForm("symbols")
	signedData := context.PostForm("signeddata")
	// feederId := context.PostForm("feederId")

	k := make(map[string]string)

	log.Println("oracleaddress", oracleaddress)
	log.Println("chainId", chainId)
	log.Println("creator", creator)
	log.Println("symbols", symbols)
	log.Println("signeddata", signedData)
	// log.Println("feederId", feederId)

	signer, err := utils.GetSigner(chainId, creator, oracleaddress, signedData)

	log.Println("signer", signer)

	if signer.Hex() != creator {
		context.JSON(http.StatusUnauthorized, errors.New("sign err"))
		return
	}

	feederId := ob.RelDB.GetFeederID(creator)

	log.Println("feederId from creator", feederId)

	if feederId == "" {
		// address, privatekey = utils.NewKeyPair()
		context.JSON(http.StatusUnauthorized, errors.New("need access to this feeder"))
		return
	} else {
		log.Println("getting GetFeederAccessByID", feederId)

		var owneraddress string
		owneraddress, address = ob.RelDB.GetFeederAccessByID(feederId)
		if owneraddress != creator {
			context.JSON(http.StatusUnauthorized, errors.New("need access to this feeder"))
			return
		}

	}
	log.Println("owneraddress GetFeederAccessByID", address)

	err = ob.RelDB.SetKeyPair(address, privatekey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	id := ob.RelDB.GetKeyPairID(address)

	err = ob.RelDB.SetOracleConfig(oracleaddress, id, creator, symbols, chainId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	if feederId != "" {
		oracleconfigid := ob.RelDB.GetOracleConfig(oracleaddress)
		fmt.Println("err oracleconfigid", oracleconfigid)

		err = ob.RelDB.SetFeederConfig(feederId, oracleconfigid)

		if err != nil {
			fmt.Println("err SetFeederConfig", err)
			context.JSON(http.StatusInternalServerError, err)
			return
		}
	}

	k["oracleaddress"] = oracleaddress
	k["chainId"] = chainId
	k["creator"] = creator
	k["symbols"] = symbols
	k["publicKey"] = address

	context.JSON(http.StatusCreated, k)
}
