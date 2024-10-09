package stackshelper

type Block struct {
	Height         int    `json:"height"`
	Hash           string `json:"hash"`
	BlockTime      int    `json:"block_time"`
	BlockTimeISO   string `json:"block_time_iso"`
	IndexBlockHash string `json:"index_block_hash"`
	TxCount        int    `json:"tx_count"`
}

type ClarityValue struct {
	Hex  string `json:"hex"`
	Repr string `json:"repr"`
}

type FunctionArg struct {
	ClarityValue
	Name string `json:"name"`
	Type string `json:"type"`
}

type ContractCall struct {
	ContractID   string        `json:"contract_id"`
	FunctionName string        `json:"function_name"`
	FunctionArgs []FunctionArg `json:"function_args"`
}

type ContractLog struct {
	ContractID string       `json:"contract_id"`
	Topic      string       `json:"topic"`
	Value      ClarityValue `json:"value"`
}

type Event struct {
	Index       int         `json:"event_index"`
	Type        string      `json:"event_type"`
	ContractLog ContractLog `json:"contract_log"`
}

type Transaction struct {
	TxID          string       `json:"tx_id"`
	SenderAddress string       `json:"sender_address"`
	BlockHash     string       `json:"block_hash"`
	BlockHeight   int          `json:"block_height"`
	BlockTime     int          `json:"block_time"`
	TxStatus      string       `json:"tx_status"`
	TxType        string       `json:"tx_type"`
	TxResult      ClarityValue `json:"tx_result"`
	ContractCall  ContractCall `json:"contract_call"`
	Events        []Event      `json:"events"`
}

type GetBlockTransactionsResponse struct {
	Limit   int           `json:"limit"`
	Offset  int           `json:"offset"`
	Total   int           `json:"total"`
	Results []Transaction `json:"results"`
}

type AddressTransaction struct {
	Tx Transaction `json:"tx"`
}

type GetAddressTransactionsResponse struct {
	Limit   int                  `json:"limit"`
	Offset  int                  `json:"offset"`
	Total   int                  `json:"total"`
	Results []AddressTransaction `json:"results"`
}

type DataMapEntry struct {
	Data  string `json:"data"`
	Proof string `json:"proof"`
}
