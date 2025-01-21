package bitflowhelper

type WrapToken struct {
	TokenDecimals uint   `json:"tokenDecimals"`
	TokenContract string `json:"tokenContract"`
	Name          string `json:"tokenName"`
}

type TokenMetadata struct {
	TokenID       string               `json:"token-id"`
	Name          string               `json:"name"`
	Symbol        string               `json:"symbol"`
	TokenDecimals uint                 `json:"tokenDecimals"`
	TokenContract string               `json:"tokenContract"`
	WrapTokens    map[string]WrapToken `json:"wrapTokens"`
}

type GetAllTokensResponse struct {
	Tokens []TokenMetadata `json:"tokens"`
}
