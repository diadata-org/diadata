package main

import (
	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/prometheus/common/log"
)

func main() {

	datastore, err := models.NewDataStore()
	if err != nil {
		log.Errorln("NewDataStore", err)
	}

	relDB, err := models.NewRelDataStore()
	if err != nil {
		log.Errorln("NewRelDataStore:", err)
	}

	totalAssets, err := relDB.GetActiveAssetCount()
	if err != nil {
		log.Errorln("Error getting Asset count", err)
	}
	log.Infoln("Total Assets :", totalAssets)

	pagecount := 0
	assetScanned := 0
	limit := 50
	skip := 0
	totalPage := totalAssets / limit
	for int(totalPage) >= pagecount {
		var assets []dia.Asset
		var assetIds []string

		log.Infoln("Asset Scanned ", assetScanned)
		log.Infoln("totalPage ", totalPage)
		assets, assetIds, err = relDB.GetActiveAsset(limit, skip)
		skip = skip + limit
		log.Infoln("page ", pagecount)
		if err != nil {
			log.Errorln("Error getting asssets", err)
		}
		pagecount++

		assetVolume := make(map[string]float64)

		for index, asset := range assets {
			assetScanned++
			log.Errorln(asset)
			log.Errorln(assetIds[index])
			volume, err := datastore.GetVolume(asset)
			if err != nil {
				log.Errorln("Error getting volume of asset", asset.Symbol)

			} else {
				assetVolume[assetIds[index]] = *volume

			}

		}

		if len(assetVolume) > 0 {
			err = relDB.SetAssetVolume(assetVolume)
			if err != nil {
				log.Errorln("Errorsaving asset volume", err)
			}

		}

	}

}
