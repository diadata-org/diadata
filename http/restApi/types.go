package restApi

import (
		"github.com/diadata-org/api-golang/services/model"
)

type APIError struct {
	ErrorCode    int    `json:"errorcode"`
	ErrorMessage string `json:"errormessage"`
}


type Coins struct {
	Coins []models.Quotation
}

