package velarhelper

type TokenMetadata struct {
	ContractAddress string `json:"contractAddress"`
	Name            string `json:"name"`
	Symbol          string `json:"symbol"`
	Decimals        string `json:"decimal"`
	DecimalNum      int    `json:"tokenDecimalNum"`
}

type Ticker struct {
	ID             string  `json:"ticker_id"`
	PoolID         string  `json:"pool_id"`
	BaseCurrency   string  `json:"base_currency"`
	TargetCurrency string  `json:"target_currency"`
	BaseVolume     float64 `json:"base_volume"`
	TargetVolume   float64 `json:"target_volume"`
}

type SwapInfo struct {
	AmountIn    uint64
	AmountOut   uint64
	TokenIn     string
	TokenOut    string
	Symbol      string
	LpToken     string
	Token0      string
	Token1      string
	ProtocolFee Fee
	ShareFee    Fee
	SwapFee     Fee
}

type Fee struct {
	Denominator uint64
	Numerator   uint64
}