package models

import (
	"context"
	"fmt"
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	keyAssetCache = "dia_asset_"
)

// GetKeyAsset returns an asset's key in the redis cache of the asset table.
// @assetID refers to the primary key asset_id in the asset table.
func (rdb *RelDB) GetKeyAsset(asset dia.Asset) (string, error) {
	ID, err := rdb.GetAssetID(asset)
	if err != nil {
		return "", err
	}
	return keyAssetCache + strconv.Itoa(ID), nil
}

// -------------------------------------------------------------
// Postgres methods
// -------------------------------------------------------------

// SetAsset stores an asset into postgres.
func (rdb *RelDB) SetAsset(asset dia.Asset) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetAssetID returns the unique identifier of @asset in postgres table asset, if the entry exists.
func (rdb *RelDB) GetAssetID(asset dia.Asset) (ID int, err error) {
	err = rdb.postgresClient.QueryRow(context.Background(), "select asset_id from asset where address=$1 and blockchain=$2", asset.Address, asset.Blockchain.Name).Scan(&ID)
	if err != nil {
		return
	}
	return ID, nil
}

// GetAsset is the standard method in order to uniquely retrieve an asset from asset table.
func (rdb *RelDB) GetAsset(address, blockchain string) (asset dia.Asset, err error) {
	var decimals string
	err = rdb.postgresClient.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where address=$1 and blockchain=$2", address, blockchain).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	// TO DO: Get Blockchain by name from postgres and add to asset
	return
}

// GetAssetsBySymbolName returns a (possibly multiple) dia.Asset by its symbol and name from postgres.
func (rdb *RelDB) GetAssetsBySymbolName(symbol, name string) (assets []dia.Asset, err error) {
	var decimals string
	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1 and name=$2", symbol, name)
	if err != nil {
		return
	}
	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, decimals, &asset.Blockchain.Name)
		decimalsInt, err := strconv.Atoi(decimals)
		if err != nil {
			return []dia.Asset{}, err
		}
		asset.Decimals = uint8(decimalsInt)
		// TO DO: Get Blockchain by name from postgres and add to asset
		assets = append(assets, asset)
	}

	return
}

// GetPage returns assets per page number. @hasNext is true iff there is a non-empty next page.
func (rdb *RelDB) GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error) {

	pagesize := rdb.pagesize
	skip := pagesize * pageNumber
	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip)
	if err != nil {
		return
	}
	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals, &asset.Blockchain)
		assets = append(assets, asset)
	}
	// Last page (or empty page)
	if len(rows.RawValues()) < int(pagesize) {
		hasNextPage = false
		return
	}
	// No next page
	nextPageRows, err := rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset LIMIT $1 OFFSET $2 ", pagesize, skip+1)
	if len(nextPageRows.RawValues()) == 0 {
		hasNextPage = false
		return
	}
	hasNextPage = true
	return
}

// Count returns the number of assets stored in postgres
func (rdb *RelDB) Count() (count uint32, err error) {
	err = rdb.postgresClient.QueryRow(context.Background(), "select count(*) from asset").Scan(&count)
	if err != nil {
		return
	}
	return
}

// -------------------------------------------------------------
// Caching layer
// -------------------------------------------------------------

// SetAssetCache stores @asset in redis, using its primary key in postgres as key.
// As a consequence, @asset is only cached iff it exists in postgres.
func (rdb *RelDB) SetAssetCache(asset dia.Asset) error {
	key, err := rdb.GetKeyAsset(asset)
	fmt.Printf("cache asset %s with key %s\n ", asset.Symbol, key)
	if err != nil {
		return err
	}
	return rdb.redisClient.Set(key, &asset, 0).Err()
}

// GetAssetCache returns an asset by its asset_id as defined in asset table in postgres
func (rdb *RelDB) GetAssetCache(assetID string) (dia.Asset, error) {
	asset := dia.Asset{}
	err := rdb.redisClient.Get(keyAssetCache + assetID).Scan(&asset)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetAsset with postgres asset_id %s\n", err, assetID)
		}
		return asset, err
	}
	return asset, nil
}

// CountCache returns the number of assets in the cache
func (rdb *RelDB) CountCache() (uint32, error) {
	keysPattern := keyAssetCache + "*"
	allAssets := rdb.redisClient.Keys(keysPattern).Val()
	return uint32(len(allAssets)), nil
}

// SetAvailablePairsCache stores @pairs in redis
func (rdb *RelDB) SetAvailablePairsCache(exchange string, pairs []dia.ExchangePair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return rdb.redisClient.Set(key, &p, 0).Err()
}
