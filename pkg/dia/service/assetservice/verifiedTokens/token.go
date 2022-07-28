package verifiedTokens

import (
	"encoding/json"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/ethereum/go-ethereum/common"
)

type TokenList struct {
	Name      string    `json:"name"`
	Timestamp time.Time `json:"timestamp"`
	Version   struct {
		Major int `json:"major"`
		Minor int `json:"minor"`
		Patch int `json:"patch"`
	} `json:"version"`
	Keywords []string `json:"keywords"`
	Tokens   []struct {
		Address  string `json:"address"`
		ChainID  int    `json:"chainId"`
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals int    `json:"decimals"`
		LogoURI  string `json:"logoURI"`
	} `json:"tokens"`
	LogoURI string `json:"logoURI"`
}

type VerifiedTokens struct {
	tokenList TokenList
	tokenMap  map[string]dia.Asset
}

func New() (*VerifiedTokens, error) {
	var tokenList TokenList
	b, _, err := utils.GetRequest("https://tokens.coingecko.com/uniswap/all.json")
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal(b, &tokenList)
	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string]dia.Asset)
	for _, token := range tokenList.Tokens {
		tokenMap[token.Symbol] = dia.Asset{Symbol: token.Symbol, Name: token.Name, Decimals: uint8(token.Decimals), Address: common.HexToAddress(token.Address).Hex()}
	}
	vt := &VerifiedTokens{tokenList: tokenList, tokenMap: tokenMap}
	return vt, nil

}

func (vt *VerifiedTokens) AppendVerifiedTokens(assets []dia.Asset) {
	for _, asset := range assets {
		vt.tokenMap[asset.Symbol] = asset

	}
}

func (vt *VerifiedTokens) Get() map[string]dia.Asset {
	return vt.tokenMap
}

func (vt *VerifiedTokens) IsExists(asset dia.Asset) bool {

	verifiedAsset, ok := vt.tokenMap[asset.Symbol]
	if ok {
		if verifiedAsset.Address == asset.Address {
			return true
		}
	}
	return false

}
