package bitflowhelper

type TokenMetadata struct {
	TokenID       string `json:"token-id"`
	Name          string `json:"name"`
	Symbol        string `json:"symbol"`
	TokenDecimals uint   `json:"tokenDecimals"`
	TokenContract string `json:"tokenContract"`
}

type GetAllTokensResponse struct {
	Tokens []TokenMetadata `json:"tokens"`
}
