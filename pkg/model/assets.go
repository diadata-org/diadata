package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/go-redis/redis"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

const (
	keyAssetCache        = "dia_asset_"
	keyExchangePairCache = "dia_exchangepair_"
)

// GetKeyAsset returns an asset's key in the redis cache of the asset table.
// @assetID refers to the primary key asset_id in the asset table.
func (rdb *RelDB) GetKeyAsset(asset dia.Asset) (string, error) {
	ID, err := rdb.GetAssetID(asset)
	if err != nil {
		return "", err
	}
	return keyAssetCache + ID, nil
}

// -------------------------------------------------------------
// Postgres methods
// -------------------------------------------------------------

// 		-------------------------------------------------------------
// 		asset TABLE methods
// 		-------------------------------------------------------------

// SetAsset stores an asset into postgres.
func (rdb *RelDB) SetAsset(asset dia.Asset) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into asset(symbol,name,address,decimals,blockchain) values ($1,$2,$3,$4,$5)", asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain.Name)
	if err != nil {
		return err
	}
	return nil
}

// GetAssetID returns the unique identifier of @asset in postgres table asset, if the entry exists.
func (rdb *RelDB) GetAssetID(asset dia.Asset) (ID string, err error) {
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

// GetAssetByID returns an asset by its uuid
func (rdb *RelDB) GetAssetByID(assetID string) (asset dia.Asset, err error) {
	var decimals string
	err = rdb.postgresClient.QueryRow(context.Background(), "select symbol,name,address,decimals,blockchain from asset where asset_id=$1", assetID).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
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
// If @name is an empty string, it returns all assets with @symbol.
// If @symbol is an empty string, it returns all assets with @name.
func (rdb *RelDB) GetAssetsBySymbolName(symbol, name string) (assets []dia.Asset, err error) {
	var decimals string
	var rows pgx.Rows
	if name == "" {
		rows, err = rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1", symbol)
	} else if symbol == "" {
		rows, err = rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset where name=$1", name)
	} else {
		rows, err = rdb.postgresClient.Query(context.Background(), "select symbol,name,address,decimals,blockchain from asset where symbol=$1 and name=$2", symbol, name)
	}
	if err != nil {
		return
	}
	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
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

// IdentifyAsset looks for all assets in postgres which match the non-null fields in @asset
// Comment 1: The only critical field is @Decimals, as this is initialized with 0, while an
// asset is allowed to have zero decimals as well (for instance sngls, trxc).
// Comment 2: Should we add a preprocessing step in which notation is corrected corresponding
// to the notation in the underlying contract on the blockchain?
func (rdb *RelDB) IdentifyAsset(asset dia.Asset) (assets []dia.Asset, err error) {
	query := "select symbol,name,address,decimals,blockchain from asset where "
	var and string
	if asset.Symbol != "" {
		query += fmt.Sprintf("symbol='%s'", asset.Symbol)
		and = " and "
	}
	if asset.Name != "" {
		query += fmt.Sprintf(and+"name='%s'", asset.Name)
		and = " and "
	}
	if asset.Address != "" {
		query += fmt.Sprintf(and+"address='%s'", asset.Address)
		and = " and "
	}
	if asset.Decimals != 0 {
		query += fmt.Sprintf(and+"decimals='%d'", asset.Decimals)
		and = " and "
	}
	if asset.Blockchain.Name != "" {
		query += fmt.Sprintf(and+"blockchain='%s'", asset.Blockchain.Name)
	}
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	var decimals string
	for rows.Next() {
		asset := dia.Asset{}
		rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain.Name)
		intDecimals, err := strconv.Atoi(decimals)
		if err != nil {
			log.Error("error parsing decimals string")
			continue
		}
		asset.Decimals = uint8(intDecimals)
		assets = append(assets, asset)
	}

	return
}

// 		-------------------------------------------------------------
// 		exchangesymbol TABLE methods
// 		-------------------------------------------------------------

// SetExchangeSymbol writes unique data into exchangesymbol table.
func (rdb *RelDB) SetExchangeSymbol(exchange string, symbol string) error {
	query := "insert into exchangesymbol(symbol,exchange) values ($1,$2)"
	_, err := rdb.postgresClient.Exec(context.Background(), query, symbol, exchange)
	if err != nil {
		return err
	}
	return nil
}

// VerifyExchangeSymbol verifies @symbol on @exchange and maps it uniquely to @assetID in asset table.
// It returns true if symbol,exchange is present and succesfully updated.
func (rdb *RelDB) VerifyExchangeSymbol(exchange string, symbol string, assetID string) (bool, error) {
	query := "update exchangesymbol set verified=true,asset_id=$1 where symbol=$2 and exchange=$3"
	resp, err := rdb.postgresClient.Exec(context.Background(), query, assetID, symbol, exchange)
	if err != nil {
		return false, err
	}
	var success bool
	respSlice := strings.Split(string(resp), " ")
	numUpdates := respSlice[1]
	if numUpdates != "0" {
		success = true
	}
	return success, nil
}

// GetExchangeSymbolID returns the ID of the unique asset associated to @symbol on @exchange
// in case the symbol is verified. An empty string if not.
func (rdb *RelDB) GetExchangeSymbolID(exchange string, symbol string) (assetID string, verified bool, err error) {
	err = rdb.postgresClient.QueryRow(context.Background(), "select asset_id, verified from exchangesymbol where symbol=$1 and exchange=$2", symbol, exchange).Scan(&assetID, &verified)
	if err != nil {
		return
	}
	return
}

// 		-------------------------------------------------------------
// 		exchangepair TABLE methods
// 		-------------------------------------------------------------

// GetExchangePairs returns all trading pairs on @exchange from exchangepair table
func (rdb *RelDB) GetExchangePairs(exchange string) (pairs []dia.ExchangePair, err error) {

	rows, err := rdb.postgresClient.Query(context.Background(), "select symbol,foreignname from exchangepair where exchange=$1", exchange)
	for rows.Next() {
		pair := dia.ExchangePair{}
		rows.Scan(&pair.Symbol, &pair.ForeignName)
		pairs = append(pairs, pair)
	}

	return pairs, nil
}

// SetExchangePair adds @pair to exchangepair table
// TO DO: extend by fields verified, id_basetoken and id_quotetoken
func (rdb *RelDB) SetExchangePair(exchange string, pair dia.ExchangePair) error {
	_, err := rdb.postgresClient.Exec(context.Background(), "insert into exchangepair(symbol,foreignname,exchange) values ($1,$2,$3)", pair.Symbol, pair.ForeignName, exchange)
	if err != nil {
		return err
	}
	return nil
}

// -------------------------------------------------------------
// General methods
// -------------------------------------------------------------

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
			log.Errorf("Error: %v on GetAssetCache with postgres asset_id %s\n", err, assetID)
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

// -------------- Caching exchange pairs -------------------

// SetExchangePairCache stores @pairs in redis
func (rdb *RelDB) SetExchangePairCache(exchange string, pair dia.ExchangePair) error {
	key := keyExchangePairCache + exchange + "_" + pair.ForeignName
	return rdb.redisClient.Set(key, &pair, 0).Err()
}

// GetExchangePairCache returns an exchange pair by @exchange and @foreigName
func (rdb *RelDB) GetExchangePairCache(exchange string, foreignName string) (dia.ExchangePair, error) {
	exchangePair := dia.ExchangePair{}
	err := rdb.redisClient.Get(keyExchangePairCache + exchange + "_" + foreignName).Scan(&exchangePair)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("GetExchangePairCache on %s with foreign name %s: %v\n", exchange, foreignName, err)
		}
		return exchangePair, err
	}
	return exchangePair, nil
}
