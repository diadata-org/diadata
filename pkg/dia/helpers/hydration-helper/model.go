package hydrationhelper

const (
	DefaultRefreshDelay              = 400 // millisec
	DefaultSleepBetweenContractCalls = 300 // millisec
	DefaultEventsLimit               = 100
	DefaultSwapContractsLimit        = 100
)

type HydrationAssetMetadata struct {
	Id                 string `json:"id"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Decimals           uint   `json:"decimals"`
	Icon               string `json:"icon"`
	Type               string `json:"type"`
	IsSufficient       bool   `json:"isSufficient"`
	ExistentialDeposit string `json:"existentialDeposit"`
	Origin             int    `json:"origin"`
}

type HydrationPoolTokenMetada struct {
	ID                 string `json:"id"`
	Balance            string `json:"balance"`
	Name               string `json:"name"`
	Symbol             string `json:"symbol"`
	Decimals           int    `json:"decimals"`
	Icon               string `json:"icon"`
	Type               string `json:"type"`
	IsSufficient       bool   `json:"isSufficient"`
	ExistentialDeposit string `json:"existentialDeposit"`
	Origin             int    `json:"origin,omitempty"`
	UsdBalance         string `json:"usdBalance"`
	Index              int    `json:"index"`
}

type HydrationPoolUsdBalanceMetadata struct {
	Amount   string `json:"amount"`
	Decimals int    `json:"decimals"`
}

type HydrationPoolMetada struct {
	Address         string                     `json:"address"`
	Type            string                     `json:"type"`
	Tokens          []HydrationPoolTokenMetada `json:"tokens"`
	MaxInRatio      int                        `json:"maxInRatio"`
	MaxOutRatio     int                        `json:"maxOutRatio"`
	MinTradingLimit int                        `json:"minTradingLimit"`
}

type HydrationSwapEvent struct {
	TxID      string `json:"txId"`
	Timestamp int64  `json:"timestamp"`
	BlockHash string `json:"blockHash"`
	AssetIn   string `json:"assetIn"`
	AssetOut  string `json:"assetOut"`
	AmountIn  string `json:"amountIn"`
	AmountOut string `json:"amountOut"`
}
