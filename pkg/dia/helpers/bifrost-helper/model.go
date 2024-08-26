package bifrosthelper

const (
	DefaultRefreshDelay              = 400 // millisec
	DefaultSleepBetweenContractCalls = 300 // millisec
	DefaultEventsLimit               = 100
	DefaultSwapContractsLimit        = 100
)

type BifrostAssetMetadata struct {
	Name           string `json:"name"`
	Symbol         string `json:"symbol"`
	Decimals       string `json:"decimals"`
	MinimalBalance string `json:"minimalBalance"`
}
