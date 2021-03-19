package main

import (
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"

	models "github.com/diadata-org/diadata/pkg/model"
)

func main() {
	relDB, err := models.NewRelDataStore()
	if err != nil {
		fmt.Println(err)
	}

	baseToken := dia.Asset{
		Symbol:     "USDT",
		Name:       "Tether USD",
		Decimals:   uint8(6),
		Blockchain: dia.BlockChain{Name: "Ethereum"},
		Address:    "0xdAC17F958D2ee523a2206206994597C13D831ec7",
	}
	quoteToken := dia.Asset{
		Symbol:     "YFI",
		Name:       "yearn.finance",
		Decimals:   uint8(18),
		Blockchain: dia.BlockChain{Name: "Ethereum"},
		Address:    "0x0bc529c00C6401aEF6D220BE8C6Ea1667F6Ad93e",
	}
	// relDB.SetAsset(baseToken)
	// relDB.SetAsset(quoteToken)

	pair := dia.ExchangePair{
		Symbol:      "YFI",
		ForeignName: "YFI-USDT",
		Verified:    true,
		UnderlyingPair: dia.Pair{
			BaseToken:  baseToken,
			QuoteToken: quoteToken,
		},
	}
	relDB.SetExchangePair("KuCoin", pair, false)

	resppair, err := relDB.GetExchangePair("KuCoin", "YFI-USDT")
	fmt.Println("resppair, err: ", resppair, err)

}
