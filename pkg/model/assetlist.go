package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (rdb *RelDB) DeleteAssetList(sheetName string) error {
	result, err := rdb.postgresClient.Exec(context.Background(),
		"DELETE FROM asset_list WHERE list_name = $1", sheetName)
	if err != nil {
		return err
	}
	rowsAffected := result.RowsAffected()

	if rowsAffected == 0 {
		log.Printf("No asset found with the name: %s\n", sheetName)
	} else {
		log.Printf("Deleted asset and its related data with the name: %s\n", sheetName)
	}
	return nil
}

func (rdb *RelDB) SetAssetList(asset dia.AssetList) error {
	var assetID int
	err := rdb.postgresClient.QueryRow(context.Background(),
		`INSERT INTO asset_list (asset_name, custom_name, symbol, methodology,list_name) 
		VALUES ($1, $2, $3, $4, $5) RETURNING id`,
		asset.AssetName, asset.CustomName, asset.Symbol, asset.Methodology, asset.ListName,
	).Scan(&assetID)
	if err != nil {
		return err
	}

	for _, exchange := range asset.Exchanges {
		var exchangeID int
		err := rdb.postgresClient.QueryRow(context.Background(),
			`INSERT INTO exchange_list (name, asset_id) VALUES ($1, $2) RETURNING id`,
			exchange.Name, assetID,
		).Scan(&exchangeID)
		if err != nil {
			return err
		}

		for _, pair := range exchange.Pairs {
			_, err := rdb.postgresClient.Exec(context.Background(),
				"INSERT INTO exchange_pairs (exchange_id, pair) VALUES ($1, $2)",
				exchangeID, pair,
			)
			if err != nil {
				return err
			}
		}
	}
	fmt.Println("Asset and related exchanges and pairs inserted successfully!")
	return nil
}

func (rdb *RelDB) SearchAssetList(searchTerm string) ([]dia.AssetList, error) {
	var assets []dia.AssetList

	rows, err := rdb.postgresClient.Query(context.Background(),
		`SELECT id, assetName, customeName, symbol, methodology 
         FROM asset_list 
         WHERE assetName LIKE '%' || $1 || '%' OR symbol LIKE '%' || $1 || '%'`,
		searchTerm,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var asset dia.AssetList
		var assetID int
		if err := rows.Scan(&assetID, &asset.AssetName, &asset.CustomName, &asset.Symbol, &asset.Methodology); err != nil {
			return nil, err
		}

		exchanges, err := rdb.getExchangesByAssetID(assetID)
		if err != nil {
			return nil, err
		}
		asset.Exchanges = exchanges
		assets = append(assets, asset)
	}

	return assets, nil
}
func (rdb *RelDB) GetAssetListBySymbol(symbol string) ([]dia.AssetList, error) {
	var assets []dia.AssetList

	rows, err := rdb.postgresClient.Query(context.Background(),
		`SELECT id, asset_name, custom_name, symbol, methodology 
         FROM asset_list 
         WHERE customename LIKE   $1`,
		symbol,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var asset dia.AssetList
		var assetID int
		if err := rows.Scan(&assetID, &asset.AssetName, &asset.CustomName, &asset.Symbol, &asset.Methodology); err != nil {
			return nil, err
		}

		exchanges, err := rdb.getExchangesByAssetID(assetID)
		if err != nil {
			return nil, err
		}
		asset.Exchanges = exchanges
		assets = append(assets, asset)
	}

	return assets, nil
}

func (rdb *RelDB) getExchangesByAssetID(assetID int) ([]dia.ExchangeList, error) {
	rows, err := rdb.postgresClient.Query(context.Background(),
		`SELECT e.name, ep.pair 
         FROM exchange_list e
         JOIN exchange_pairs ep ON e.id = ep.exchange_id
         WHERE e.asset_id = $1`,
		assetID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	exchangeMap := make(map[string]*dia.ExchangeList)

	for rows.Next() {
		var exchangeName string
		var pair string
		if err := rows.Scan(&exchangeName, &pair); err != nil {
			return nil, err
		}

		if _, exists := exchangeMap[exchangeName]; !exists {
			exchangeMap[exchangeName] = &dia.ExchangeList{
				Name:  exchangeName,
				Pairs: []string{},
			}
		}
		exchangeMap[exchangeName].Pairs = append(exchangeMap[exchangeName].Pairs, pair)
	}

	var exchanges []dia.ExchangeList
	for _, exchange := range exchangeMap {
		exchanges = append(exchanges, *exchange)
	}

	return exchanges, nil
}
