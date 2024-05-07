package alephiumhelper

type MutField struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type ContractStateResponse struct {
	Address   string     `json:"address"`
	MutFields []MutField `json:"mutFields"`
}
