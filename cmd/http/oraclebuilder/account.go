package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerInput struct {
	Email            string   `form:"email"`
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
}

type AddWalletInput struct {
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
	Creator          string   `form:"creator"`
	AccessLevel      string   `form:"access_level" binding:"required,oneof=read read_write"`
	UserName         string   `form:"username"`
}
type RemoveWalletInput struct {
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
	Creator          string   `form:"creator"`
}
type UpdateAccessInput struct {
	WalletPublicKey string `form:"wallet_public_key"`
	AccessLevel     string `form:"access_level" binding:"required,oneof=read read_write"`
	UserName        string `form:"username"`
}

type ViewInput struct {
	Creator string `form:"creator"`
}

func (ob *Env) ViewAccount(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input ViewInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorln("ShouldBind", err)
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission
	customer, err := ob.RelDB.GetCustomerByPublicKey(input.Creator)
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalFeeds, err := ob.totalFeedsUsedByCustomer(strconv.Itoa(customer.CustomerID))
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while  getting totalFeedusedByCustomer"})
		return
	}
	customer.NumberOfDataFeeds = totalFeeds

	customer.DeployedOracles = ob.RelDB.GetTotalOracles(strconv.Itoa(customer.CustomerID))

	plan, err := ob.RelDB.GetPlan(context, customer.CustomerPlan)
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount GetPlan err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while  getting plan"})
		return
	}

	combined := map[string]interface{}{
		"plan":                  plan,
		"customer_id":           customer.CustomerID,
		"email":                 customer.Email,
		"account_creation_date": customer.AccountCreationDate,
		"customer_plan":         customer.CustomerPlan,
		"deployed_oracles":      customer.DeployedOracles,
		"payment_status":        customer.PaymentStatus,
		"payment_source":        customer.PaymentSource,
		"number_of_data_feeds":  customer.NumberOfDataFeeds,
		"active":                customer.Active,
		"public_keys":           customer.PublicKeys,
	}

	context.JSON(http.StatusOK, combined)

}

func (ob *Env) AddWallet(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input AddWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission
	err := ob.RelDB.AddWalletKeys(input.Creator, input.UserName, input.AccessLevel, input.WalletPublicKeys)
	if err != nil {
		log.Errorf("Request ID: %s,  AddWalletKeys err %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Added successfully"})

}

func (ob *Env) UpdateAccess(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input UpdateAccessInput
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission, remove self key, check should not be self, min one key has to me read_write
	log.Infoln("Request ID: %s, UpdateAccess input err %v ", requestId, input)

	err := ob.RelDB.UpdateAccessLevel(input.UserName, input.AccessLevel, input.WalletPublicKey)
	if err != nil {
		log.Errorf("Request ID: %s,  UpdateAccessLevel %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Access Level Updated successfully"})

}

func (ob *Env) RemoveWallet(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input RemoveWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	//TODO check permission, remove self key
	err := ob.RelDB.RemoveWalletKeys(input.WalletPublicKeys)
	if err != nil {
		log.Errorf("Request ID: %s,  RemoveWalletKeys err %v ", requestId, err)
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
