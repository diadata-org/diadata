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

func getKeyAsset(symbol, name string) (key string) {
	key = keyAssetCache + symbol + "_" + name
	return
}

// -------------------------------------------------------------
// Postgres methods
// -------------------------------------------------------------

// SetAsset stores an asset into postgres
func (rdb *RelDB) SetAsset(asset dia.Asset) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetAsset returns a dia.Asset by its symbol and name from postgres
func (rdb *RelDB) GetAsset(symbol, name string) (asset dia.Asset, err error) {
	var decimals string
	err = rdb.postgresClient.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1 and name=$2", symbol, name).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
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
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals)
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

// SetAssetCache stores @asset in redis
func (rdb *RelDB) SetAssetCache(asset dia.Asset) error {
	return rdb.redisClient.Set(getKeyAsset(asset.Symbol, asset.Name), &asset, 0).Err()
}

// GetAssetCache returns an asset by name and symbol
func (rdb *RelDB) GetAssetCache(symbol, name string) (dia.Asset, error) {
	asset := dia.Asset{}
	err := rdb.redisClient.Get(getKeyAsset(symbol, name)).Scan(&asset)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetAsset %v\n", err, symbol)
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
func (rdb *RelDB) SetAvailablePairsCache(exchange string, pairs []dia.Pair) error {
	key := "dia_available_pairs_" + exchange
	var p dia.Pairs = pairs
	return rdb.redisClient.Set(key, &p, 0).Err()
}
