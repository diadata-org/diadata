package main

import (
	"strings"
	"time"

	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"
	"github.com/diadata-org/diadata/pkg/dia/helpers/ethhelper"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/diadata-org/diadata/pkg/utils"
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
	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Fatal("relational datastore error: ", err)
	}
	conn, err := ethclient.Dial("http://159.69.120.42:8545/")
	// conn, err := ethclient.Dial("https://mainnet.infura.io/v3/806b0419b2d041869fc83727e0043236")
	if err != nil {
		log.Fatal(err)
	}
	// Fetch token contract addresses from json file
	tokenAddresses, err := ethhelper.GetAddressesFromFile(tokensListFilename)
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("got %v assets from json", len(tokenAddresses))
	// Fetch all assets from postgres asset table
	assetTableAssets, err := relDB.GetAllAssets("Ethereum")
	if err != nil {
		log.Fatal(err)
	}
	log.Infof("got %v assets from postgres", len(assetTableAssets))
	time.Sleep(1 * time.Minute)
	var moreAddresses []string
	for _, asset := range assetTableAssets {
		moreAddresses = append(moreAddresses, asset.Address)
	}
	tokenAddresses = append(tokenAddresses, moreAddresses...)
	tokenAddressesFinal := utils.UniqueStrings(tokenAddresses)
	log.Infof("got %v assets in total", len(tokenAddressesFinal))

	// Get map for locked wallets per asset
	lockedWalletsMap, err := supplyservice.GetLockedWalletsFromConfig(lockedWalletsFilename)
	if err != nil {
		log.Error(err)
	}
	// Initial run
	setSupplies(tokenAddressesFinal, lockedWalletsMap, ds, conn)

	// Continuously update supplies once every 24h
	ticker := time.NewTicker(24 * time.Hour)
	go func() {
		for range ticker.C {
			setSupplies(tokenAddressesFinal, lockedWalletsMap, ds, conn)
			if err != nil {
				log.Error(err)
			}
		}
	}()
	select {}

}

func setSupplies(tokenAddresses []string, lockedWalletsMap map[string][]string, ds models.Datastore, conn *ethclient.Client) {
	for _, address := range tokenAddresses {

		supp, err := supplyservice.GetTotalSupplyfromMainNet(address, lockedWalletsMap[address], conn)
		if err != nil || len(supp.Asset.Symbol) < 2 || supp.Supply < 2 {
			if strings.ToLower(address) == "0x9f8f72aa9304c8b593d555f12ef6589cc3a579a2" {
				// Comment: maker contract emits byte32 instead of string
				supp.Asset.Symbol = "MKR"
				supp.Asset.Name = "Maker"
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
		if supp.Asset.Symbol == "YAM" {
			supp.CirculatingSupply = float64(13907678)
		}
		if supp.Asset.Symbol == "CRO" {
			supp.CirculatingSupply = float64(20631963470)
		}
		if supp.Asset.Symbol == "DTA" {
			supp.CirculatingSupply = float64(21000000)
		}
		if supp.Asset.Symbol == "DIA" {
			supp.CirculatingSupply = float64(25549170)
		}
		if supp.Asset.Symbol == "SPICE" {
			supp.CirculatingSupply = float64(1945426.80)
		}
		if strings.ToLower(address) == "0xa1faa113cbe53436df28ff0aee54275c13b40975" {
			supp.CirculatingSupply = float64(174136442)
		}

		err = ds.SetSupply(&supp)
		if err != nil {
			log.Errorf("error setting supply for %s: %v\n", supp.Asset.Symbol, err)
		} else {
			log.Info("set supply: " + supp.Asset.Name + " - " + supp.Asset.Symbol)
		}
	}
}
