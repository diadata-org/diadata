package alephiumhelper

type EventContract struct {
	EventIndex      int     `json:"eventIndex"`
	ContractAddress string  `json:"contractAddress"`
	TxHash          string  `json:"txHash"`
	Fields          []Field `json:"fields"`
}

type EventContractResponse struct {
	Events    []EventContract `json:"events"`
	NextStart int             `json:"nextStart"`
}
