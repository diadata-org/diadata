package bifrosthelper

const (
	DefaultRefreshDelay              = 400 // millisec
	DefaultSleepBetweenContractCalls = 300 // millisec
	DefaultEventsLimit               = 100
	DefaultSwapContractsLimit        = 100
)

type BifrostAssetMetadata struct {
	AssetKey       string `json:"assetKey"`
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Decimals       string `json:"decimals"`
	MinimalBalance string `json:"minimalBalance"`
}

type BifrostPoolMetadata struct {
	PoolId         string            `json:"poolId"`
	PoolAsset      map[string]string `json:"poolAsset"`
	Assets         []string          `json:"assets"`
	Precisions     []string          `json:"precisions"`
	MintFee        string            `json:"mintFee"`
	SwapFee        string            `json:"swapFee"`
	RedeemFee      string            `json:"redeemFee"`
	TotalSupply    string            `json:"totalSupply"`
	Balances       []string          `json:"balances"`
	FeeRecipient   string            `json:"feeRecipient"`
	AccountId      string            `json:"accountId"`
	YieldRecipient string            `json:"yieldRecipient"`
	Precision      string            `json:"precision"`
}
