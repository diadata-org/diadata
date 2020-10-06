package main

import (
	"time"

	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

func main() {

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/251a25bd10b8460fa040bb7202e22571")
	if err != nil {
		log.Fatal(err)
	}
	// Fetch token contract addresses from json file
	adds, err := ethhelper.GetAddressesFromFile()
	if err != nil {
		log.Fatal(err)
	}
	// Continuously update supplies once every 24h
	for {
		timeInit := time.Now()
		for _, address := range adds {
			supp, err := supplyservice.GetTotalSupplyfromMainNet(address, ds, conn)
			if err != nil || len(supp.Symbol) < 2 || supp.Supply < 2 {
				continue
			}
			ds.SetSupply(&supp)
		}
		timeFinal := time.Now()
		timeElapsed := timeFinal.Sub(timeInit)
		time.Sleep(24*60*60*time.Second - timeElapsed)
	}

}
