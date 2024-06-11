package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerInput struct {
	Email            string   `form:"email"`
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
}

type AddWalletInput struct {
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
	Creator          string   `form:"creator"`
}
type UpdateAccessInput struct {
	WalletPublicKey string `form:"wallet_public_key"`
	AccessLevel     string `form:"access_level" binding:"required,oneof=read read_write"`
}

type ViewInput struct {
	Creator string `form:"creator"`
}

func (ob *Env) ViewAccount(context *gin.Context) {
	var input ViewInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorln("ShouldBind", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission
	customer, err := ob.RelDB.GetCustomerByPublicKey(input.Creator)
	if err != nil {
		log.Errorln("AddWalletKeys", err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, customer)

}

func (ob *Env) AddWallet(context *gin.Context) {
	var input AddWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorln("ShouldBind", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission
	err := ob.RelDB.AddWalletKeys(input.Creator, input.WalletPublicKeys)
	if err != nil {
		log.Errorln("AddWalletKeys", err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Added successfully"})

}

func (ob *Env) UpdateAccess(context *gin.Context) {
	var input UpdateAccessInput
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission, remove self key, check should not be self, min one key has to me read_write
	log.Errorln("input", input)

	err := ob.RelDB.UpdateAccessLevel(input.AccessLevel, input.WalletPublicKey)
	if err != nil {
		log.Errorln("UpdateAccessLevel", err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Access Level Updated successfully"})

}

func (ob *Env) RemoveWallet(context *gin.Context) {
	var input AddWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorln("ShouldBind", err)
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission, remove self key
	err := ob.RelDB.RemoveWalletKeys(input.WalletPublicKeys)
	if err != nil {
		log.Errorln("RemoveWalletKeys", err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Removed successfully"})

}

func (ob *Env) CreateAccount(context *gin.Context) {
	var input CustomerInput
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ob.RelDB.CreateCustomer("", 0, "", "", 2, input.WalletPublicKeys)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Customer created successfully"})

}
