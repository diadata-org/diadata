package oraclebuilder

import (
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

	oracleaddress := context.PostForm("oracleaddress")
	chainId := context.PostForm("chainId")
	creator := context.PostForm("creator")
	// signedData := context.PostForm("signeddata")

	k := make(map[string]string)

	address, privatekey := utils.NewKeyPair()

	err := ob.RelDB.SetKeyPair(address, privatekey)
	if err != nil {
		context.JSON(http.StatusInternalServerError, err)
		return
	}
	id := ob.RelDB.GetKeyPairID(address)

	err = ob.RelDB.SetUserOracle(oracleaddress, id, creator, chainId)
	if err != nil {
		context.JSON(http.StatusUnauthorized, err)
		return
	}

	k["oracleaddress"] = oracleaddress
	k["chainId"] = chainId
	k["creator"] = creator
	k["publicKey"] = address

	context.JSON(http.StatusCreated, k)
}
