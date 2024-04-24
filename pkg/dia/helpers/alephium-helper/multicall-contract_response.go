package alephiumhelper

const ByteVecType = "ByteVec"
const U256Type = "U256"

type Field struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type Contract struct {
	Address string `json:"address"`
}

type CallContractResult struct {
	Type      string     `json:"type"`
	Error     *string    `json:"error"`
	Returns   []Field    `json:"returns"`
	Contracts []Contract `json:"contracts"`
}

type MulticallContractResponse struct {
	Results []CallContractResult `json:"results"`
}

type OutputField struct {
	Field
	ResponseResult string
}

type OutputResult struct {
	Address string
	Results []OutputField
}
