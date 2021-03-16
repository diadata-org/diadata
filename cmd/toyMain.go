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

	// db, err := models.NewDataStore()
	// if err != nil {
	// 	fmt.Println(err)
	// }

	asset1 := dia.Asset{
		Symbol:     "BTC",
		Name:       "Bitcoin",
		Blockchain: dia.BlockChain{Name: "Bitcoin"},
	}
	asset2 := dia.Asset{
		Symbol:     "BTC",
		Name:       "Bitcoin",
		Blockchain: dia.BlockChain{Name: "Ethereum"},
	}
	relDB.SetAsset(asset1)
	relDB.SetAsset(asset2)

	// assets, err := relDB.GetAssetsBySymbolName("BTC", "")
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println("assets: ", assets)

	assets, err := relDB.GetAllAssets("Ethereum")
	fmt.Println("assets, err: ", assets, err)

}
