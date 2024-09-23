package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type CustomerInput struct {
	Email            string   `form:"email"`
	Name             string   `form:"name"`
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
}

type AddWalletInput struct {
	WalletPublicKeys []string `form:"wallet_public_keys[]"`
	Creator          string   `form:"creator"`
	AccessLevel      string   `form:"access_level" binding:"required,oneof=read read_write"`
	UserName         string   `form:"username"`
}
type ApproveWalletInput struct {
	WalletPublicKey string `form:"wallet_public_key"`
	Creator         string `form:"creator"`
	CustomerID      string `form:"customerID"`
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

func (ob *Env) viewAccount(context *gin.Context, publicKey string) (combined map[string]interface{}) {
	requestId := context.GetString(REQUEST_ID)

	customer, err := ob.RelDB.GetCustomerByPublicKey(publicKey)
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	totalFeeds, totalOracles, err := ob.billableResource(strconv.Itoa(customer.CustomerID))
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while  getting totalFeedusedByCustomer"})
		return
	}
	customer.NumberOfDataFeeds = totalFeeds

	customer.DeployedOracles = totalOracles

	plan, err := ob.RelDB.GetPlan(context, customer.CustomerPlan)
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount GetPlan err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while  getting plan"})
		return
	}
	pendingPk, err := ob.RelDB.GetPendingPublicKeyByCustomer(context, strconv.Itoa(customer.CustomerID))
	if err != nil {
		log.Errorf("Request ID: %s,  ViewAccount GetPlan err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error while  getting plan"})
		return
	}

	combined = map[string]interface{}{
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
		"pending_public_keys":   pendingPk,
	}
	return
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
	accountView := ob.viewAccount(context, input.Creator)

	context.JSON(http.StatusOK, accountView)

}

func (ob *Env) AddTempWallet(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input AddWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ob.RelDB.AddTempWalletKeys(input.Creator, input.UserName, input.AccessLevel, input.WalletPublicKeys)
	if err != nil {
		log.Errorf("Request ID: %s,  AddTempWalletKeys err %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": "error adding wallet"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Added successfully, waiting for approval from wallet"})

}
func (ob *Env) ApproveWallet(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	creator := context.GetString(CREATOR_ADDRESS)

	/*
		get request to add wallet
		verify if request is coming from same wallet
		move request to wallet key
	*/

	var input ApproveWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	keyId, accessLevel, username, err := ob.RelDB.GetTempWalletRequest(context, creator, input.CustomerID)
	if err != nil {
		log.Errorf("Request ID: %s,  GetTempWalletRequest err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "no request pending"})
		return
	}

	err = ob.RelDB.AddWalletKeys(creator, username, accessLevel, []string{creator}, input.CustomerID)
	if err != nil {
		log.Errorf("Request ID: %s,  AddWalletKeys err %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ob.RelDB.DeleteTempWalletRequest(context, strconv.Itoa(keyId))

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
	log.Infoln("Request ID: %s, UpdateAccess input %v ", requestId, input)
	input.UserName = " "

	err := ob.RelDB.UpdateAccessLevel(input.UserName, input.AccessLevel, input.WalletPublicKey)
	if err != nil {
		log.Errorf("Request ID: %s,  UpdateAccessLevel %v ", requestId, err)

		context.JSON(http.StatusInternalServerError, gin.H{"error": "error updating access level"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Access Level Updated successfully"})

}

func (ob *Env) RemoveWallet(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)

	var input RemoveWalletInput
	if err := context.ShouldBind(&input); err != nil {
		log.Errorf("Request ID: %s,  ShouldBind err %v ", requestId, err)

		context.JSON(http.StatusBadRequest, gin.H{"error": "error removing wallet"})
		return
	}
	//TODO check permission, remove self key
	err := ob.RelDB.RemoveWalletKeys(input.WalletPublicKeys)
	if err != nil {
		log.Errorf("Request ID: %s,  RemoveWalletKeys err %v ", requestId, err)
		context.JSON(http.StatusInternalServerError, gin.H{"error": "error removing wallet"})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "Wallet Removed successfully"})

}

func (ob *Env) CreateAccount(context *gin.Context) {
	requestId := context.GetString(REQUEST_ID)
	signer := context.GetString(CREATOR_ADDRESS)

	var input CustomerInput
	if err := context.ShouldBind(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": gin.H{"error": "error creating account"}})
		return
	}

	log.Infof("Request ID: %s,  CreateCustomer with wallet key   %s ", requestId, signer)

	customer, err := ob.RelDB.GetCustomerByPublicKey(signer)
	if err != nil && err.Error() == "no rows in result set" {
		log.Infof("Request ID: %s,  New customer create one   %s ", requestId, signer)
		err = ob.RelDB.CreateCustomer("", input.Name, 0, "", "", 2, []string{signer})
		if err != nil {
			log.Errorf("Request ID: %s,  CreateCustomer  %v ", requestId, err)

			context.JSON(http.StatusInternalServerError, gin.H{"error": gin.H{"error": "error creating account"}})
			return
		}

		context.JSON(http.StatusOK, gin.H{"message": "Customer created successfully"})
		return
	}

	if customer != nil {
		context.JSON(http.StatusOK, gin.H{"message": "Customer Already exists"})
		return
	}

}
