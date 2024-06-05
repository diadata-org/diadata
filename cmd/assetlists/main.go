package main

import (
	"fmt"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
	"github.com/sirupsen/logrus"

	"github.com/xuri/excelize/v2"
)

func main() {
	log := logrus.New()

	reldb, err := models.NewPostgresDataStore()
	if err != nil {
		log.Errorln("Error connecting to asset DB: ", err)
		return
	}
	f, err := excelize.OpenFile("oracle_builder_asset_markets.xlsx")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func() {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}()

	sheets := f.GetSheetMap()
	if len(sheets) == 0 {
		fmt.Println("No sheets found in the Excel file")
		return
	}

	for _, sheetName := range sheets {
		fmt.Println("sheetname", sheetName)
		reldb.DeleteAssetList(sheetName)
		scrapeSheet(sheetName, f, reldb)
	}

}

func scrapeSheet(sheetname string, f *excelize.File, reldb *models.RelDB) {
	rows, err := f.GetRows(sheetname)
	if err != nil {
		fmt.Println(err)
		return
	}

	assetsMap := make(map[string]*dia.AssetList)
	var currentAsset *dia.AssetList

	for rc, row := range rows {
		if rc == 0 {
			continue
		}
		var e *dia.ExchangeList
		for i, colCell := range row {
			switch i {
			case 0:
				if assetVal, exists := assetsMap[colCell]; exists {
					currentAsset = assetVal
				} else {
					currentAsset = &dia.AssetList{AssetName: colCell}
					assetsMap[colCell] = currentAsset
				}
			case 1:
				if currentAsset != nil {
					currentAsset.CustomName = colCell
				}
			case 2:
				if currentAsset != nil {
					currentAsset.Symbol = colCell
				}
			case 3:
				e = &dia.ExchangeList{Name: colCell}
			case 4:
				pairs := parsePairs(colCell)
				e.Pairs = pairs
				if currentAsset != nil {
					currentAsset.Exchanges = append(currentAsset.Exchanges, *e)
				}
			}
		}
	}

	printAssets(reldb, assetsMap, sheetname)
}

func parsePairs(input string) []string {
	input = strings.Trim(input, "[]")
	input = strings.ReplaceAll(input, "'", "")
	return strings.Split(input, " ")
}

func printAssets(reldb *models.RelDB, assetsMap map[string]*dia.AssetList, sheetName string) {
	row := 0
	for _, assetVal := range assetsMap {
		assetVal.ListName = sheetName
		err := reldb.SetAssetList(*assetVal)
		if err != nil {
			fmt.Println("-------:", err)

		}
		fmt.Printf("\n Sheetname %s  row %d", sheetName, row)
		row++

	}
}
