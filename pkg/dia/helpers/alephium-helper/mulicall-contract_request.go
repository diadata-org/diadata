package alephiumhelper

type CallContractRequest struct {
	Group       int    `json:"group"`
	Address     string `json:"address"`
	MethodIndex int    `json:"methodIndex"`
}

type Calls struct {
	Calls []CallContractRequest `json:"calls"`
}
