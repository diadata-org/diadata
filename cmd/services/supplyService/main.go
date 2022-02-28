package main

import (
	"time"

	supplyservice "github.com/diadata-org/diadata/internal/pkg/supplyService"
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

}
