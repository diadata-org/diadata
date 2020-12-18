package main

import (
	"time"

	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
)

const (
	tokensListFilename    = "tokens_list"
	lockedWalletsFilename = "wallets"
)

func main() {

	ds, err := models.NewDataStore()
	if err != nil {
		log.Fatal("datastore error: ", err)
	}
	conn, err := ethclient.Dial("http://159.69.120.42:8545/")
	if err != nil {
		log.Fatal(err)
	}
	// Fetch token contract addresses from json file
	tokenAddresses, err := ethhelper.GetAddressesFromFile(tokensListFilename)
	if err != nil {
		log.Fatal(err)
	}

	// Get map for locked wallets per asset
	lockedWalletsMap, err := supplyservice.GetLockedWalletsFromConfig(lockedWalletsFilename)
	if err != nil {
		log.Error(err)
	}

	// Initial run
	for _, address := range tokenAddresses {
		supp, err := supplyservice.GetTotalSupplyfromMainNet(address, lockedWalletsMap[address], conn)
		if err != nil || len(supp.Symbol) < 2 || supp.Supply < 2 {
			continue
		}
		// Hardcoded hotfix for some supplies:
		if supp.Symbol == "YAM" {
			supp.CirculatingSupply = float64(13907678)
		}
		if supp.Symbol == "CRO" {
			supp.CirculatingSupply = float64(20631963470)
		}
		if supp.Symbol == "DTA" {
			supp.CirculatingSupply = float64(21000000)
		}

		ds.SetSupply(&supp)
		log.Info("set supply: ", supp)
	}

	// Continuously update supplies once every 24h
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				for _, address := range tokenAddresses {
					supp, err := supplyservice.GetTotalSupplyfromMainNet(address, lockedWalletsMap[address], conn)
					if err != nil || len(supp.Symbol) < 2 || supp.Supply < 2 {
						continue
					}
					// Hardcoded hotfix for some supplies:
					if supp.Symbol == "YAM" {
						supp.CirculatingSupply = float64(13907678)
					}
					if supp.Symbol == "CRO" {
						supp.CirculatingSupply = float64(20631963470)
					}
					if supp.Symbol == "DTA" {
						supp.CirculatingSupply = float64(21000000)
					}
					if supp.Symbol == "DIA" {
						supp.CirculatingSupply = float64(25549170)
					}
					if supp.Symbol == "SPICE" {
						supp.CirculatingSupply = float64(1945426.80)
					}

					ds.SetSupply(&supp)
					log.Info("set supply: ", supp)
				}

			}
		}
	}()
	select {}

}
