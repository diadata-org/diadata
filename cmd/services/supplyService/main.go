package main

import (
	"strings"
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
	// conn, err := ethclient.Dial("http://159.69.120.42:8545/")
	conn, err := ethclient.Dial("https://mainnet.infura.io/v3/806b0419b2d041869fc83727e0043236")
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
	err = setSupplies(tokenAddresses, lockedWalletsMap, ds, conn)
	if err != nil {
		log.Error(err)
	}

	// Continuously update supplies once every 24h
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				err = setSupplies(tokenAddresses, lockedWalletsMap, ds, conn)
				if err != nil {
					log.Error(err)
				}
			}
		}
	}()
	select {}

}

func setSupplies(tokenAddresses []string, lockedWalletsMap map[string][]string, ds models.Datastore, conn *ethclient.Client) error {
	for _, address := range tokenAddresses {

		supp, err := supplyservice.GetTotalSupplyfromMainNet(address, lockedWalletsMap[address], conn)
		if err != nil || len(supp.Symbol) < 2 || supp.Supply < 2 {
			if strings.ToLower(address) == "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2" {
				// Comment: maker contract emits byte32 instead of string
				supp.Symbol = "MKR"
				supp.Name = "Maker"
				supp.CirculatingSupply = float64(902135)
				supp.Supply = float64(995691)
				supp.Source = "diadata.org"
				supp.Time = time.Now()
			} else {
				log.Error(err)
				continue
			}
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
		if strings.ToLower(address) == "0xa1faa113cbe53436df28ff0aee54275c13b40975" {
			supp.CirculatingSupply = float64(174136442)
		}

		err = ds.SetSupply(&supp)
		if err != nil {
			log.Errorf("error setting supply for %s: %v\n", supp.Symbol, err)
		} else {
			log.Info("set supply: " + supp.Name + " - " + supp.Symbol)
		}
	}
	return nil
}
