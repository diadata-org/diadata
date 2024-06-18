package alephiumhelper

type ContractEvent struct {
	TxID            string  `json:"txId"`
	ContractAddress string  `json:"contractAddress"`
	EventIndex      int     `json:"eventIndex"`
	Fields          []Field `json:"fields"`
}

type BlockEventsResponse struct {
	Events []ContractEvent `json:"events"`
}
