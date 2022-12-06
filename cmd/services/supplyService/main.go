package main

import (
	"time"

	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
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

	suppliesCG, err := supplyservice.GetETHSuppliesFromCG()
	if err != nil {
		log.Error("get supplies from coingecko: ", err)
	}
	for i := range suppliesCG {
		asset, err := relDB.GetAsset(suppliesCG[i].Asset.Address, suppliesCG[i].Asset.Blockchain)
		if err != nil {
			// Only add supply for assets we have in our database.
			continue
		}
		suppliesCG[i].Asset = asset
		suppliesCG[i].Time = time.Now()
		suppliesCG[i].Source = "Coingecko"
		err = ds.SetSupply(&suppliesCG[i])
		if err != nil {
			log.Errorf("error setting supply for %s: %v\n", suppliesCG[i].Asset.Symbol, err)
		} else {
			log.Info("set supply: " + suppliesCG[i].Asset.Name + " - " + suppliesCG[i].Asset.Symbol)
		}
	}

	// Save old "circulating" supply as total supply (i.e. #DIA without the burnt tokens)
	err = ds.SetDiaTotalSupply(float64(173296236.5011769))
	if err != nil {
		log.Errorf("error setting total supply for %s: %v\n", "DIA", err)
	} else {
		log.Info("set total supply: DIA")
	}
	// Set circulating supply
	err = ds.SetDiaCirculatingSupply(float64(88163785))
	if err != nil {
		log.Errorf("error setting circulating supply for %s: %v\n", "DIA", err)
	} else {
		log.Info("set circulating supply: DIA")
	}

	xrp, err := relDB.GetAsset("0x0000000000000000000000000000000000000000", "Ripple")
	if err != nil {
		log.Error("get xrp: ", err)
	}
	supplyXRP := dia.Supply{
		Asset:             xrp,
		Supply:            float64(100000000000),
		CirculatingSupply: float64(47949281138),
		Source:            "Coingecko",
		Time:              time.Now(),
	}
	err = ds.SetSupply(&supplyXRP)
	if err != nil {
		log.Errorf("error setting supply for %s: %v\n", supplyXRP.Asset.Symbol, err)
	} else {
		log.Info("set supply: " + supplyXRP.Asset.Name + " - " + supplyXRP.Asset.Symbol)
	}

	btc, err := relDB.GetAsset("0x0000000000000000000000000000000000000000", "Bitcoin")
	if err != nil {
		log.Error("get bitcoin: ", err)
	}
	supplyBTC := dia.Supply{
		Asset:             btc,
		Supply:            float64(21000000),
		CirculatingSupply: float64(18970462),
		Source:            "Coingecko",
		Time:              time.Now(),
	}
	err = ds.SetSupply(&supplyBTC)
	if err != nil {
		log.Errorf("error setting supply for %s: %v\n", supplyBTC.Asset.Symbol, err)
	} else {
		log.Info("set supply: " + supplyBTC.Asset.Name + " - " + supplyBTC.Asset.Symbol)
	}
}
