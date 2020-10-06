package main

import (
	"fmt"

	"github.com/diadata-org/diadata/internal/pkg/tokenService"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func main() {

	ds, err := models.NewInfluxDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/251a25bd10b8460fa040bb7202e22571")
	if err != nil {
		log.Error(err)
	}

	address := "0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48"

	supp, circSupp, _ := tokenService.GetTotalSupplyfromMainNet(address, ds, conn)
	fmt.Println("supply: ", supp, circSupp)

}
