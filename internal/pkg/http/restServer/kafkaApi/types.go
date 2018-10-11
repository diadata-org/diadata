package kafkaApi

import (
	"github.com/diadata-org/diadata/pkg/dia"
)

type APIError struct {
	ErrorCode    int    `json:"errorcode"`
	ErrorMessage string `json:"errormessage"`
}

type TradesBlock struct {
	Offset   int `json:"offset"`
	Messages []dia.TradesBlock
}

type FiltersBlock struct {
	Offset   int `json:"offset"`
	Messages []dia.FiltersBlock
}

type Trades struct {
	Offset   int `json:"offset"`
	Messages []dia.Trade
}
