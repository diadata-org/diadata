package main

import (
	"fmt"
	"log"

	"github.com/diadata-org/diadata/internal/pkg/tokenService"
	models "github.com/diadata-org/diadata/pkg/model"
)

func main() {
	fmt.Println("Token: Golem")

	ds, err := models.NewInfluxDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}

	address := "0xa74476443119A942dE498590Fe1f2454d7D4aC0d"

	tokenService.GetTotalSupplyfromMainNet(address, ds)

	address2 := "0x1f9840a85d5aF5bf1D1762F925BDADdC4201F984"

	fmt.Println("Token: UNI")

	tokenService.GetTotalSupplyfromMainNet(address2, ds)

	address3 := "0xdac17f958d2ee523a2206206994597c13d831ec7"

	fmt.Println("Token: Tether USD")

	tokenService.GetTotalSupplyfromMainNet(address3, ds)

	address4 := "0x6b175474e89094c44da98b954eedeac495271d0f"

	fmt.Println("Token: Dai Stablecoin(DAI)")

	tokenService.GetTotalSupplyfromMainNet(address4, ds)

	address5 := "0x05d3606d5c81eb9b7b18530995ec9b29da05faba"

	fmt.Println("Token: TomoChain(TOMOE)")

	tokenService.GetTotalSupplyfromMainNet(address5, ds)

}
