package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type CustomerInput struct {
	Email            []string `form:"email"`
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
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
