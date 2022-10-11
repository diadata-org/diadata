package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/go-redis/redis"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
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
	query := fmt.Sprintf("INSERT INTO %s (symbol,name,address,decimals,blockchain) VALUES ($1,$2,$3,$4,$5) ON CONFLICT (address,blockchain) DO NOTHING", assetTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, asset.Symbol, asset.Name, asset.Address, strconv.Itoa(int(asset.Decimals)), asset.Blockchain)
	if err != nil {
		return err
	}
	return nil
}

// GetAssetID returns the unique identifier of @asset in postgres table asset, if the entry exists.
func (rdb *RelDB) GetAssetID(asset dia.Asset) (ID string, err error) {
	query := fmt.Sprintf("SELECT asset_id FROM %s WHERE address=$1 AND blockchain=$2", assetTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, asset.Address, asset.Blockchain).Scan(&ID)
	if err != nil {
		return
	}
	return
}

func (rdb *RelDB) GetAssetMap(asset_id string) (ID string, err error) {
	query := fmt.Sprintf("SELECT group_id FROM %s WHERE asset_id=$1", assetIdent)
	err = rdb.postgresClient.QueryRow(context.Background(), query, asset_id).Scan(&ID)
	if err != nil {
		return
	}
	return
}

func (rdb *RelDB) GetAssetByGroupID(group_id string) (assets []dia.Asset, err error) {
	var rows pgx.Rows

	query := fmt.Sprintf("SELECT symbol,name,address,blockchain,decimals FROM %s WHERE asset_id in (select asset_id from %s where group_id=$1)", assetTable, assetIdent)

	rows, err = rdb.postgresClient.Query(context.Background(), query, group_id)
	if err != nil {
		return
	}
	defer rows.Close()

	var decimals string
	for rows.Next() {
		var asset dia.Asset
		err := rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Blockchain, &decimals)
		if err != nil {
			log.Error(err)
		}
		decimalsInt, err := strconv.Atoi(decimals)
		if err != nil {
			continue
		}
		asset.Decimals = uint8(decimalsInt)
		// asset.Blockchain = blockchain
		assets = append(assets, asset)
	}
	return
}

// SetAsset stores an asset into postgres.
func (rdb *RelDB) InsertAssetMap(group_id string, asset_id string) error {
	query := fmt.Sprintf("INSERT INTO %s (group_id,asset_id) VALUES ($1,$2)", assetIdent)
	log.Println("query", query)

	_, err := rdb.postgresClient.Exec(context.Background(), query, group_id, asset_id)
	if err != nil {
		return err
	}
	return nil
}
func (rdb *RelDB) InsertNewAssetMap(asset_id string) error {
	query := fmt.Sprintf("INSERT INTO %s (asset_id) VALUES ($1)", assetIdent)
	log.Println("query", query)
	_, err := rdb.postgresClient.Exec(context.Background(), query, asset_id)
	if err != nil {
		return err
	}
	return nil
}

var assetCache = make(map[string]dia.Asset)

// GetAsset is the standard method in order to uniquely retrieve an asset from asset table.
func (rdb *RelDB) GetAsset(address, blockchain string) (asset dia.Asset, err error) {
	assetKey := "GetAsset_" + address + "_" + blockchain
	cachedAsset, found := assetCache[assetKey]
	if found {
		asset = cachedAsset
		return
	}
	var decimals string
	query := fmt.Sprintf("SELECT symbol,name,address,decimals,blockchain FROM %s WHERE address=$1 AND blockchain=$2", assetTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	assetCache[assetKey] = asset
	return
}

// GetAssetByID returns an asset by its uuid
func (rdb *RelDB) GetAssetByID(assetID string) (asset dia.Asset, err error) {
	var decimals string
	query := fmt.Sprintf("SELECT symbol,name,address,decimals,blockchain FROM %s WHERE asset_id=$1", assetTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, assetID).Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	return
}

// GetAllAssets returns all assets on @blockchain from asset table.
func (rdb *RelDB) GetAllAssets(blockchain string) (assets []dia.Asset, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf("SELECT symbol,name,address,decimals FROM %s WHERE blockchain=$1", assetTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, blockchain)
	if err != nil {
		return
	}
	defer rows.Close()

	var decimals string
	for rows.Next() {
		var asset dia.Asset
		err := rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals)
		if err != nil {
			log.Error(err)
		}
		decimalsInt, err := strconv.Atoi(decimals)
		if err != nil {
			continue
		}
		asset.Decimals = uint8(decimalsInt)
		asset.Blockchain = blockchain
		assets = append(assets, asset)
	}
	return
}

// GetAssetsBySymbolName returns a (possibly multiple) dia.Asset by its symbol and name from postgres.
// If @name is an empty string, it returns all assets with @symbol.
// If @symbol is an empty string, it returns all assets with @name.
func (rdb *RelDB) GetAssetsBySymbolName(symbol, name string) (assets []dia.Asset, err error) {
	var decimals string
	var rows pgx.Rows
	var query string
	if name == "" {
		query = fmt.Sprintf(`
		SELECT symbol,name,address,decimals,blockchain 
		FROM %s a
		INNER JOIN %s av
		ON av.asset_id=a.asset_id
		WHERE av.volume>0
		AND av.time_stamp IS NOT NULL
		AND symbol ILIKE '%s%%'
		ORDER BY av.volume DESC`,
			assetTable,
			assetVolumeTable,
			symbol,
		)
	} else if symbol == "" {
		query = fmt.Sprintf(`
		SELECT symbol,name,address,decimals,blockchain 
		FROM %s a
		INNER JOIN %s av
		ON av.asset_id=a.asset_id
		WHERE av.volume>0
		AND av.time_stamp IS NOT NULL
		AND name ILIKE '%s%%'
		ORDER BY av.volume DESC`,
			assetTable,
			assetVolumeTable,
			symbol,
		)
	} else {
		query = fmt.Sprintf(`
		SELECT symbol,name,address,decimals,blockchain 
		FROM %s a 
		INNER JOIN %s av 
		ON av.asset_id=a.asset_id 
		WHERE av.volume>0
		AND av.time_stamp IS NOT NULL
		AND (symbol ILIKE '%s%%' OR name ILIKE '%s%%')
		ORDER BY av.volume DESC`,
			assetTable,
			assetVolumeTable,
			name,
			symbol,
		)
	}
	if err != nil {
		return
	}
	rows, err = rdb.postgresClient.Query(context.Background(), query)

	log.Infoln("GetAssetsBySymbolName query", query)
	defer rows.Close()
	for rows.Next() {
		var decimalsInt int
		var asset dia.Asset
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assets = append(assets, asset)
	}
	return
}

// GetAssetsByAddress returns a (possibly multiple) dia.Asset by its address from postgres.
func (rdb *RelDB) GetAssetsByAddress(address string) (assets []dia.Asset, err error) {
	var decimals string
	var rows pgx.Rows
	query := fmt.Sprintf(`
	SELECT symbol,name,address,decimals,blockchain 
	FROM %s a 
	INNER JOIN %s av 
	ON a.asset_id=av.asset_id
	WHERE av.volume>0
	AND av.time_stamp IS NOT NULL
	AND address ILIKE '%s%%'
	ORDER BY av.volume DESC`,
		assetTable,
		assetVolumeTable,
		address,
	)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		var decimalsInt int
		var asset dia.Asset
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assets = append(assets, asset)
	}
	return
}

// GetFiatAssetBySymbol returns a fiat asset by its symbol. This is possible as
// fiat currencies are uniquely defined by their symbol.
func (rdb *RelDB) GetFiatAssetBySymbol(symbol string) (asset dia.Asset, err error) {
	var decimals string
	query := fmt.Sprintf("SELECT name,address,decimals FROM %s WHERE symbol=$1 AND blockchain='Fiat'", assetTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, symbol).Scan(&asset.Name, &asset.Address, &decimals)
	if err != nil {
		return
	}
	decimalsInt, err := strconv.Atoi(decimals)
	if err != nil {
		return
	}
	asset.Decimals = uint8(decimalsInt)
	asset.Symbol = symbol
	asset.Blockchain = "Fiat"
	// TO DO: Get Blockchain by name from postgres and add to asset
	return
}

// IdentifyAsset looks for all assets in postgres which match the non-null fields in @asset
// Comment 1: The only critical field is @Decimals, as this is initialized with 0, while an
// asset is allowed to have zero decimals as well (for instance sngls, trxc).
// Comment 2: Should we add a preprocessing step in which notation is corrected corresponding
// to the notation in the underlying contract on the blockchain?
// Comment 3: Can we improve this? How to treat cases like CoinBase emitting symbol name
// 'Wrapped Bitcoin' instead of the correct 'Wrapped BTC', or 'United States Dollar' instead
// of 'United States dollar'? On idea would be to add a table with alternative names for
// symbol tickers, so WBTC -> [Wrapped Bitcoin, Wrapped bitcoin, Wrapped BTC,...]
func (rdb *RelDB) IdentifyAsset(asset dia.Asset) (assets []dia.Asset, err error) {
	query := fmt.Sprintf("SELECT symbol,name,address,decimals,blockchain FROM %s WHERE ", assetTable)
	var and string
	if asset.Symbol != "" {
		query += fmt.Sprintf("symbol='%s'", asset.Symbol)
		and = " AND "
	}
	if asset.Name != "" {
		query += fmt.Sprintf(and+"name='%s'", asset.Name)
		and = " AND "
	}
	if asset.Address != "" {
		query += fmt.Sprintf(and+"address='%s'", common.HexToAddress(asset.Address).Hex())
		and = " AND "
	}
	if asset.Decimals != 0 {
		query += fmt.Sprintf(and+"decimals='%d'", asset.Decimals)
		and = " AND "
	}
	if asset.Blockchain != "" {
		query += fmt.Sprintf(and+"blockchain='%s'", asset.Blockchain)
	}
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	var decimals string
	for rows.Next() {
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
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

// SetExchangeSymbol writes unique data into exchangesymbol table if not yet in there.
func (rdb *RelDB) SetExchangeSymbol(exchange string, symbol string) error {
	query := fmt.Sprintf("INSERT INTO %s (symbol,exchange) SELECT $1,$2 WHERE NOT EXISTS (SELECT 1 FROM exchangesymbol WHERE symbol=$1 AND exchange=$2)", exchangesymbolTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, symbol, exchange)
	if err != nil {
		return err
	}
	return nil
}

// GetAssets returns all assets which share the symbol ticker @symbol.
func (rdb *RelDB) GetAssets(symbol string) (assets []dia.Asset, err error) {
	query := fmt.Sprintf("SELECT symbol,name,address,decimals,blockchain FROM %s WHERE symbol=$1 ", assetTable)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, symbol)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var decimals string
		var decimalsInt int
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assets = append(assets, asset)
	}
	return
}

// GetAssetExchnage returns all assets which share the symbol ticker @symbol.
func (rdb *RelDB) GetAssetExchange(symbol string) (exchanges []string, err error) {

	query := fmt.Sprintf("SELECT exchange FROM %s  INNER JOIN %s ON asset.asset_id = exchangesymbol.asset_id WHERE exchangesymbol.symbol = $1 ", exchangesymbolTable, assetTable)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, symbol)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var exchange string

		err = rows.Scan(&exchange)
		if err != nil {
			return
		}
		exchanges = append(exchanges, exchange)
	}
	return
}

// GetUnverifiedExchangeSymbols returns all symbols from @exchange which haven't been verified yet.
func (rdb *RelDB) GetUnverifiedExchangeSymbols(exchange string) (symbols []string, err error) {
	query := fmt.Sprintf("SELECT symbol FROM %s WHERE exchange=$1 AND verified=false ORDER BY symbol ASC", exchangesymbolTable)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, exchange)
	if err != nil {
		return
	}
	defer rows.Close()
	for rows.Next() {
		symbol := ""
		err = rows.Scan(&symbol)
		if err != nil {
			return []string{}, err
		}
		symbols = append(symbols, symbol)
	}
	return
}

// GetExchangeSymbols returns all symbols traded on @exchange.
// If @exchange is the empty string, all symbols are returned.
// If @substring is not the empty string, all symbols that begin with @substring (case insensitive) are returned.
func (rdb *RelDB) GetExchangeSymbols(exchange string, substring string) (symbols []string, err error) {
	var query string
	var rows pgx.Rows
	if exchange != "" {
		if substring != "" {
			query = fmt.Sprintf("SELECT symbol FROM %s WHERE exchange=$1 AND symbol ILIKE '%s%%'", exchangesymbolTable, substring)
			rows, err = rdb.postgresClient.Query(context.Background(), query, exchange)

		} else {
			query = fmt.Sprintf("SELECT symbol FROM %s WHERE exchange=$1", exchangesymbolTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query, exchange)
		}
	} else {
		if substring != "" {
			query = fmt.Sprintf("SELECT symbol FROM %s WHERE symbol ILIKE '%s%%'", exchangesymbolTable, substring)
			log.Info("query: ", query)
			rows, err = rdb.postgresClient.Query(context.Background(), query)
		} else {
			query = fmt.Sprintf("SELECT symbol FROM %s", exchangesymbolTable)
			rows, err = rdb.postgresClient.Query(context.Background(), query)
		}
	}
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		symbol := ""
		err = rows.Scan(&symbol)
		if err != nil {
			return []string{}, err
		}
		symbols = append(symbols, symbol)
	}
	return
}

// VerifyExchangeSymbol verifies @symbol on @exchange and maps it uniquely to @assetID in asset table.
// It returns true if symbol,exchange is present and succesfully updated.
func (rdb *RelDB) VerifyExchangeSymbol(exchange string, symbol string, assetID string) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET verified=true,asset_id=$1 WHERE symbol=$2 AND exchange=$3", exchangesymbolTable)
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

// GetExchangeSymbolAssetID returns the ID of the unique asset associated to @symbol on @exchange
// in case the symbol is verified. An empty string if not.
func (rdb *RelDB) GetExchangeSymbolAssetID(exchange string, symbol string) (assetID string, verified bool, err error) {
	var uuid pgtype.UUID
	query := fmt.Sprintf("SELECT asset_id, verified FROM %s WHERE symbol=$1 AND exchange=$2", exchangesymbolTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, symbol, exchange).Scan(&uuid, &verified)
	if err != nil {
		return
	}
	val, err := uuid.Value()
	if err != nil {
		log.Error(err)
	}
	if val != nil {
		assetID = val.(string)
	}
	return
}

// -------------------------------------------------------------
// Blockchain methods
// -------------------------------------------------------------

func (rdb *RelDB) SetBlockchain(blockchain dia.BlockChain) (err error) {
	fields := fmt.Sprintf("INSERT INTO %s (name,genesisdate,nativetoken_id,verificationmechanism,chain_id) VALUES ", blockchainTable)
	values := "($1,$2,(SELECT asset_id FROM asset WHERE address=$3 AND blockchain=$1),$4,NULLIF($5,''))"
	conflict := " ON CONFLICT (name) DO UPDATE SET genesisdate=$2,verificationmechanism=$4,chain_id=NULLIF($5,''),nativetoken_id=(SELECT asset_id FROM asset WHERE address=$3 AND blockchain=$1) "

	query := fields + values + conflict
	_, err = rdb.postgresClient.Exec(context.Background(), query,
		blockchain.Name,
		blockchain.GenesisDate,
		blockchain.NativeToken.Address,
		blockchain.VerificationMechanism,
		blockchain.ChainID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetBlockchain(name string) (blockchain dia.BlockChain, err error) {
	query := fmt.Sprintf("SELECT genesisdate,verificationmechanism,chain_id,address,symbol FROM %s INNER JOIN %s ON %s.nativetoken_id=%s.asset_id where %s.name=$1", blockchainTable, assetTable, blockchainTable, assetTable, blockchainTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, name).Scan(
		&blockchain.GenesisDate,
		&blockchain.VerificationMechanism,
		&blockchain.ChainID,
		&blockchain.NativeToken.Address,
		&blockchain.NativeToken.Symbol,
	)
	if err != nil {
		return
	}
	blockchain.Name = name
	return
}

// GetAllBlockchains returns all blockchains from the blockchain table.
// If fullAsset=true it returns the complete native token as asset, otherwise only its symbol string.
func (rdb *RelDB) GetAllBlockchains(fullAsset bool) ([]dia.BlockChain, error) {
	var blockchains []dia.BlockChain
	var query string
	if fullAsset {
		queryString := "SELECT b.name,b.genesisdate,a.Symbol,a.Name,a.Address,a.Decimals,b.verificationmechanism,b.chain_id FROM %s b LEFT JOIN %s a ON nativetoken_id = a.asset_id"
		query = fmt.Sprintf(queryString, blockchainTable, assetTable)
	} else {
		query = fmt.Sprintf("SELECT b.name,b.genesisdate,a.Symbol,b.verificationmechanism,b.chain_id FROM %s b LEFT JOIN %s a ON nativetoken_id = a.asset_id", blockchainTable, assetTable)
	}
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []dia.BlockChain{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			blockchain     dia.BlockChain
			genDate        sql.NullFloat64
			symbol         sql.NullString
			verifMechanism sql.NullString
			chainID        sql.NullString
			//  fullAsset
			name     sql.NullString
			address  sql.NullString
			decimals sql.NullInt64
		)

		if fullAsset {
			err = rows.Scan(
				&blockchain.Name,
				&genDate,
				&symbol,
				&name,
				&address,
				&decimals,
				&verifMechanism,
				&chainID,
			)
		} else {
			err = rows.Scan(
				&blockchain.Name,
				&genDate,
				&symbol,
				&verifMechanism,
				&chainID,
			)
		}
		if err != nil {
			return []dia.BlockChain{}, err
		}
		if genDate.Valid {
			blockchain.GenesisDate = int64(genDate.Float64)
		}
		if symbol.Valid {
			blockchain.NativeToken.Symbol = symbol.String
		}
		if verifMechanism.Valid {
			blockchain.VerificationMechanism = dia.VerificationMechanism(verifMechanism.String)
		}
		if chainID.Valid {
			blockchain.ChainID = chainID.String
		}
		if fullAsset {
			if name.Valid {
				blockchain.NativeToken.Name = name.String
			}
			if address.Valid {
				blockchain.NativeToken.Address = address.String
			}
			if decimals.Valid {
				blockchain.NativeToken.Decimals = uint8(decimals.Int64)
			}
			blockchain.NativeToken.Blockchain = blockchain.Name
		}
		blockchains = append(blockchains, blockchain)
	}

	return blockchains, nil
}

// GetAllAssetsBlockchains returns all blockchain names existent in the asset table.
func (rdb *RelDB) GetAllAssetsBlockchains() ([]string, error) {
	var blockchains []string
	query := fmt.Sprintf("SELECT DISTINCT blockchain FROM %s WHERE name!='' ORDER BY blockchain ASC", assetTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []string{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var blockchain string
		err := rows.Scan(&blockchain)
		if err != nil {
			return []string{}, err
		}
		blockchains = append(blockchains, blockchain)
	}

	return blockchains, nil
}

// -------------------------------------------------------------
// General methods
// -------------------------------------------------------------

// GetPage returns assets per page number. @hasNext is true iff there is a non-empty next page.
func (rdb *RelDB) GetPage(pageNumber uint32) (assets []dia.Asset, hasNextPage bool, err error) {

	pagesize := rdb.pagesize
	skip := pagesize * pageNumber
	rows, err := rdb.postgresClient.Query(context.Background(), "SELECT symbol,name,address,decimals,blockchain FROM asset LIMIT $1 OFFSET $2 ", pagesize, skip)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		fmt.Println("---")
		var asset dia.Asset
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &asset.Decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		assets = append(assets, asset)
	}
	// Last page (or empty page)
	if len(rows.RawValues()) < int(pagesize) {
		hasNextPage = false
		return
	}
	// No next page
	nextPageRows, err := rdb.postgresClient.Query(context.Background(), "SELECT symbol,name,address,decimals,blockchain FROM asset LIMIT $1 OFFSET $2 ", pagesize, skip+1)
	if len(nextPageRows.RawValues()) == 0 {
		hasNextPage = false
		return
	}
	defer nextPageRows.Close()
	hasNextPage = true
	return
}

// Count returns the number of assets stored in postgres
func (rdb *RelDB) Count() (count uint32, err error) {
	err = rdb.postgresClient.QueryRow(context.Background(), "SELECT COUNT(*) FROM asset").Scan(&count)
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
		if !errors.Is(err, redis.Nil) {
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
		if !errors.Is(err, redis.Nil) {
			log.Errorf("GetExchangePairCache on %s with foreign name %s: %v\n", exchange, foreignName, err)
		}
		return exchangePair, err
	}
	return exchangePair, nil
}

func (rdb *RelDB) SetAssetVolume24H(asset dia.Asset, volume float64, timestamp time.Time) error {

	initialStr := fmt.Sprintf("INSERT INTO %s (asset_id,volume,time_stamp) VALUES ", assetVolumeTable)
	substring := fmt.Sprintf(
		"((SELECT asset_id FROM asset WHERE address='%s' AND blockchain='%s'),%f,to_timestamp(%v))",
		asset.Address,
		asset.Blockchain,
		volume,
		timestamp.Unix(),
	)
	conflict := " ON CONFLICT (asset_id) DO UPDATE SET volume=EXCLUDED.volume,time_stamp=EXCLUDED.time_stamp"

	query := initialStr + substring + conflict
	_, err := rdb.postgresClient.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetAssetVolume24H(asset dia.Asset) (volume float64, err error) {
	query := fmt.Sprintf("SELECT volume FROM %s INNER JOIN %s ON assetvolume.asset_id = asset.asset_id WHERE address=$1 AND blockchain=$2", assetVolumeTable, assetTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, asset.Address, asset.Blockchain).Scan(&volume)
	return
}

func (rdb *RelDB) GetTopAssetByVolume(symbol string) (assets []dia.Asset, err error) {
	query := fmt.Sprintf("SELECT symbol,name,address,decimals,blockchain FROM %s INNER JOIN %s ON asset.asset_id = assetvolume.asset_id WHERE symbol=$1 ORDER BY volume DESC", assetTable, assetVolumeTable)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, symbol)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var decimals string
		var decimalsInt int
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assets = append(assets, asset)
	}
	return
}

func (rdb *RelDB) GetByLimit(limit, skip uint32) (assets []dia.Asset, assetIds []string, err error) {

	rows, err := rdb.postgresClient.Query(context.Background(), "SELECT asset_id,symbol,name,address,decimals,blockchain FROM asset LIMIT $1 OFFSET $2 ", limit, skip)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {

		var (
			decimals    string
			decimalsInt int
			assetID     string
			asset       dia.Asset
		)
		err = rows.Scan(&assetID, &asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)

		assets = append(assets, asset)
		assetIds = append(assetIds, assetID)
	}

	return
}

// GetAssetsWithVolByBlockchain returns all assets from assetvolume table that have a timestamp in the time-range (@starttime,@endtime].
// If blockchain is a non-empty string it only returns assets from @blockchain.
func (rdb *RelDB) GetAssetsWithVolByBlockchain(starttime time.Time, endtime time.Time, blockchain string) (assets []dia.AssetVolume, err error) {
	var (
		query string
		rows  pgx.Rows
	)

	query = fmt.Sprintf(`
	SELECT * FROM (
		SELECT DISTINCT ON (address,blockchain) symbol,name,address,decimals,blockchain,volume
		FROM %s 
		INNER JOIN %s
		ON (asset.asset_id = assetvolume.asset_id)
		WHERE time_stamp>to_timestamp(%v) and time_stamp<=to_timestamp(%v)`,
		assetTable,
		assetVolumeTable,
		starttime.Unix(),
		endtime.Unix(),
	)
	if blockchain != "" {
		query += fmt.Sprintf(" AND asset.blockchain='%s')", blockchain)
	} else {
		query += (")")
	}
	query += " sub ORDER BY volume DESC"

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			decimals sql.NullInt64
			volume   float64
		)
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain, &volume)
		if err != nil {
			return
		}
		if decimals.Valid {
			asset.Decimals = uint8(decimals.Int64)
		}
		assetvolume := dia.AssetVolume{Asset: asset, Volume: volume}
		assets = append(assets, assetvolume)
	}
	return
}

// GetSortedAssetSymbols search asstet by symbol
func (rdb *RelDB) GetSortedAssetSymbols(numAssets int64, skip int64, search string) (volumeSortedAssets []dia.AssetVolume, err error) {
	var (
		queryString string
		query       string
		rows        pgx.Rows
	)

	if numAssets == 0 {
		queryString = "SELECT symbol,name,address,decimals,blockchain,volume FROM %s INNER JOIN %s ON (asset.asset_id = assetvolume.asset_id) WHERE symbol ILIKE '%s%%' ORDER BY assetvolume.volume DESC LIMIT 100"
		query = fmt.Sprintf(queryString, assetTable, assetVolumeTable, search)
	} else {
		queryString = "SELECT DISTINCT ON (av.volume,av.asset_id)  a.symbol,a.name,a.address,a.decimals,a.blockchain,av.volume FROM %s  av INNER JOIN %s a ON av.asset_id=a.asset_id INNER JOIN %s es ON av.asset_id=es.asset_id INNER JOIN %s e ON es.exchange=e.name WHERE e.centralized=true  and a.symbol ILIKE '%s%%' ORDER BY av.volume DESC LIMIT %d OFFSET %d"
		query = fmt.Sprintf(queryString, assetVolumeTable, assetTable, exchangesymbolTable, exchangeTable, search, numAssets, skip)
	}
	log.Infoln("GetSortedAssetSymbols query", query)

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			decimals    string
			decimalsInt int
			volume      float64
		)
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain, &volume)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assetvolume := dia.AssetVolume{Asset: asset, Volume: volume}
		volumeSortedAssets = append(volumeSortedAssets, assetvolume)
	}
	return

}

// GetAssetsWithVOL returns the first @numAssets assets with entry in the assetvolume table, sorted by volume in descending order.
// If @numAssets==0, all assets are returned.
// If @substring is not the empty string, results are filtered by the first letters being @substring.
func (rdb *RelDB) GetAssetsWithVOL(numAssets int64, skip int64, onlycex bool, blockchain string) (volumeSortedAssets []dia.AssetVolume, err error) {
	var (
		queryString string
		query       string
		rows        pgx.Rows
	)
	if numAssets == 0 {
		numAssets = 100
	}

	if !onlycex {

		if blockchain == "" {
			queryString = `
			SELECT symbol,name,address,decimals,blockchain,volume 
			FROM %s INNER JOIN %s ON (asset.asset_id = assetvolume.asset_id) 
			ORDER BY assetvolume.volume 
			DESC LIMIT %d OFFSET %d`
			query = fmt.Sprintf(queryString, assetTable, assetVolumeTable, numAssets, skip)
		} else {
			queryString = `
			SELECT symbol,name,address,decimals,blockchain,volume 
			FROM %s INNER JOIN %s ON (asset.asset_id = assetvolume.asset_id) 
			WHERE blockchain= '%s' 
			ORDER BY assetvolume.volume 
			DESC LIMIT %d OFFSET %d`
			query = fmt.Sprintf(queryString, assetTable, assetVolumeTable, blockchain, numAssets, skip)
		}

	} else {
		if blockchain == "" {
			queryString = `
			SELECT DISTINCT ON (av.volume,av.asset_id)  a.symbol,a.name,
			a.address,a.decimals,a.blockchain,av.volume 
			FROM %s  av INNER JOIN %s a ON av.asset_id=a.asset_id 
			INNER JOIN %s es ON av.asset_id=es.asset_id INNER JOIN %s e 
			ON es.exchange=e.name 
			WHERE e.centralized=true 
			ORDER BY av.volume 
			DESC  LIMIT %d OFFSET %d`
			query = fmt.Sprintf(queryString, assetVolumeTable, assetTable, exchangesymbolTable, exchangeTable, numAssets, skip)
		} else {
			queryString = `
			SELECT DISTINCT ON (av.volume,av.asset_id) 
			a.symbol,a.name,a.address,a.decimals,a.blockchain,av.volume 
			FROM %s  av 
			INNER JOIN %s a  ON av.asset_id=a.asset_id 
			INNER JOIN %s es ON av.asset_id=es.asset_id 
			INNER JOIN %s e ON es.exchange=e.name 
			WHERE e.centralized=true AND a.blockchain = '%s' 
			ORDER BY av.volume 
			DESC  LIMIT %d OFFSET %d`
			query = fmt.Sprintf(queryString, assetVolumeTable, assetTable, exchangesymbolTable, exchangeTable, blockchain, numAssets, skip)
		}

	}

	log.Infoln("GetAssetsWithVOL query", query)

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			decimals    string
			decimalsInt int
			volume      float64
		)
		asset := dia.Asset{}
		err = rows.Scan(&asset.Symbol, &asset.Name, &asset.Address, &decimals, &asset.Blockchain, &volume)
		if err != nil {
			return
		}
		decimalsInt, err = strconv.Atoi(decimals)
		if err != nil {
			return
		}
		asset.Decimals = uint8(decimalsInt)
		assetvolume := dia.AssetVolume{Asset: asset, Volume: volume}
		volumeSortedAssets = append(volumeSortedAssets, assetvolume)
	}
	return
}

// GetAssetSource returns all exchanges @asset is traded on.
// For @cex true, only CEXes are returned. Otherwise only DEXes.
func (rdb *RelDB) GetAssetSource(asset dia.Asset, cex bool) (exchanges []string, err error) {
	var query string
	if cex {
		query = fmt.Sprintf(`
		SELECT DISTINCT ON (es.exchange) es.exchange 
		FROM %s es 
		INNER JOIN %s a ON es.asset_id = a.asset_id 
		WHERE a.blockchain='%s' AND a.address='%s'
		`, exchangesymbolTable, assetTable, asset.Blockchain, asset.Address)
	} else {
		query = fmt.Sprintf(`
		SELECT  DISTINCT ON (p.exchange) p.exchange
		FROM %s p 
		INNER JOIN %s pa ON p.pool_id=pa.pool_id 
		INNER JOIN %s a ON pa.asset_id=a.asset_id 
		WHERE a.blockchain='%s' AND a.address='%s'
		`, poolTable, poolassetTable, assetTable, asset.Blockchain, asset.Address)
	}

	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			exchange string
		)

		err = rows.Scan(&exchange)
		if err != nil {
			return
		}

		exchanges = append(exchanges, exchange)
	}
	return

}

// GetAssetsWithVOLInflux returns all assets that have an entry in Influx's volumes table and hence have been traded since @timeInit.
func (datastore *DB) GetAssetsWithVOLInflux(timeInit time.Time) ([]dia.Asset, error) {
	var quotedAssets []dia.Asset
	q := fmt.Sprintf("SELECT address,blockchain,value FROM %s WHERE filter='VOL120' AND exchange='' AND time>%d AND time<now()", influxDbFiltersTable, timeInit.UnixNano())
	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return quotedAssets, err
	}

	// Filter and store all unique assets from the filters table.
	uniqueMap := make(map[dia.Asset]struct{})
	if len(res) > 0 && len(res[0].Series) > 0 {
		if len(res[0].Series[0].Values) > 0 {
			var asset dia.Asset
			for _, val := range res[0].Series[0].Values {
				if val[1] == nil || val[2] == nil {
					continue
				}
				asset.Address = val[1].(string)
				asset.Blockchain = val[2].(string)
				if _, ok := uniqueMap[asset]; !ok {
					quotedAssets = append(quotedAssets, asset)
					uniqueMap[asset] = struct{}{}
				}
			}
		} else {
			return quotedAssets, errors.New("no recent assets with volume in influx")
		}
	} else {
		return quotedAssets, errors.New("no recent asset with volume in influx")
	}
	return quotedAssets, nil
}
