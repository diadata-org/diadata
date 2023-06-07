package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

// GetExchangePair returns the unique exchange pair given by @exchange and @foreignname from postgres.
// It also returns the underlying pair if existent.
func (rdb *RelDB) GetExchangePair(exchange string, foreignname string) (dia.ExchangePair, error) {
	var (
		exchangepair    dia.ExchangePair
		verified        bool
		uuid_quotetoken pgtype.UUID
		uuid_basetoken  pgtype.UUID
	)

	exchangepair.Exchange = exchange
	exchangepair.ForeignName = foreignname

	query := fmt.Sprintf("SELECT symbol,verified,id_quotetoken,id_basetoken FROM %s WHERE exchange=$1 AND foreignname=$2", exchangepairTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, exchange, foreignname).Scan(&exchangepair.Symbol, &verified, &uuid_quotetoken, &uuid_basetoken)
	if err != nil {
		return dia.ExchangePair{}, err
	}
	exchangepair.Verified = verified

	// Decode uuids and fetch corresponding assets
	val1, err := uuid_quotetoken.Value()
	if err != nil {
		log.Error(err)
	}
	if val1 != nil {
		var quotetoken dia.Asset
		quotetoken, err = rdb.GetAssetByID(val1.(string))
		if err != nil {
			return dia.ExchangePair{}, err
		}
		exchangepair.UnderlyingPair.QuoteToken = quotetoken
	}

	val2, err := uuid_basetoken.Value()
	if err != nil {
		log.Error(err)
	}
	if val2 != nil {
		basetoken, err := rdb.GetAssetByID(val2.(string))
		if err != nil {
			return dia.ExchangePair{}, err
		}
		exchangepair.UnderlyingPair.BaseToken = basetoken
	}

	return exchangepair, nil
}

// SetExchangePair adds @pair to exchangepair table.
// If cache==true, it is also cached into redis
func (rdb *RelDB) SetExchangePair(exchange string, pair dia.ExchangePair, cache bool) error {
	var query string
	query = fmt.Sprintf("INSERT INTO %s (symbol,foreignname,exchange) SELECT $1,$2,$3 WHERE NOT EXISTS (SELECT 1 FROM %s WHERE symbol=$1 AND foreignname=$2 AND exchange=$3)", exchangepairTable, exchangepairTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, pair.Symbol, pair.ForeignName, exchange)
	if err != nil {
		return err
	}
	basetokenID, err := rdb.GetAssetID(pair.UnderlyingPair.BaseToken)
	if err != nil {
		log.Error(err)
	}
	quotetokenID, err := rdb.GetAssetID(pair.UnderlyingPair.QuoteToken)
	if err != nil {
		log.Error(err)
	}
	if basetokenID != "" {
		query = fmt.Sprintf("UPDATE %s SET id_basetoken='%s' WHERE foreignname='%s' AND exchange='%s'", exchangepairTable, basetokenID, pair.ForeignName, exchange)
		_, err = rdb.postgresClient.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	if quotetokenID != "" {
		query = fmt.Sprintf("UPDATE %s SET id_quotetoken='%s' WHERE foreignname='%s' AND exchange='%s'", exchangepairTable, quotetokenID, pair.ForeignName, exchange)
		_, err = rdb.postgresClient.Exec(context.Background(), query)
		if err != nil {
			return err
		}
	}
	query = fmt.Sprintf("UPDATE %s SET verified='%v' WHERE foreignname='%s' AND exchange='%s'", exchangepairTable, pair.Verified, pair.ForeignName, exchange)
	_, err = rdb.postgresClient.Exec(context.Background(), query)
	if err != nil {
		return err
	}
	if cache {
		err = rdb.SetExchangePairCache(exchange, pair)
		if err != nil {
			log.Errorf("setting pair %s to redis for exchange %s: %v", pair.ForeignName, exchange, err)
		}
	}
	return nil
}

// GetExchangePairSymbols returns all foreign names on @exchange from exchangepair table.
func (rdb *RelDB) GetExchangePairSymbols(exchange string) (pairs []dia.ExchangePair, err error) {
	query := fmt.Sprintf("SELECT symbol,foreignname FROM %s WHERE exchange=$1", exchangepairTable)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, exchange)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		pair := dia.ExchangePair{Exchange: exchange}
		err = rows.Scan(&pair.Symbol, &pair.ForeignName)
		if err != nil {
			return
		}
		pairs = append(pairs, pair)
	}
	return
}

// GetExchangePairs returns all pairs on a (centralized) @exchange.
func (rdb *RelDB) GetPairsForExchange(exchange dia.Exchange, filterVerified bool, verified bool) ([]dia.ExchangePair, error) {
	var pairs []dia.ExchangePair
	exchangeType := GetExchangeType(exchange)
	if exchangeType != "CEX" {
		err := errors.New("query only feasible for centralized exchanges.")
		return pairs, err
	}

	query := fmt.Sprintf(`
		SELECT  a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals,e.verified,e.foreignname
		FROM %s e 
		INNER JOIN %s a 
		ON e.id_quotetoken=a.asset_id 
		INNER JOIN %s b 
		ON e.id_basetoken=b.asset_id 
		WHERE e.exchange='%s'`,
		exchangepairTable,
		assetTable,
		assetTable,
		exchange.Name,
	)
	if filterVerified {
		query += fmt.Sprintf(" AND e.verified='%v'", verified)
	}

	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return pairs, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			pair          dia.ExchangePair
			quoteDecimals sql.NullInt64
			baseDecimals  sql.NullInt64
		)

		err := rows.Scan(
			&pair.UnderlyingPair.QuoteToken.Symbol,
			&pair.UnderlyingPair.QuoteToken.Name,
			&pair.UnderlyingPair.QuoteToken.Address,
			&pair.UnderlyingPair.QuoteToken.Blockchain,
			&quoteDecimals,
			&pair.UnderlyingPair.BaseToken.Symbol,
			&pair.UnderlyingPair.BaseToken.Name,
			&pair.UnderlyingPair.BaseToken.Address,
			&pair.UnderlyingPair.BaseToken.Blockchain,
			&baseDecimals,
			&pair.Verified,
			&pair.ForeignName,
		)
		if err != nil {
			return pairs, err
		}
		if quoteDecimals.Valid {
			pair.UnderlyingPair.QuoteToken.Decimals = uint8(quoteDecimals.Int64)
		}
		if baseDecimals.Valid {
			pair.UnderlyingPair.BaseToken.Decimals = uint8(baseDecimals.Int64)
		}
		pair.Exchange = exchange.Name
		pair.Symbol = pair.UnderlyingPair.QuoteToken.Symbol

		pairs = append(pairs, pair)
	}

	return pairs, nil
}

func (rdb *RelDB) GetPairsForAsset(asset dia.Asset, filterVerified bool, verified bool) ([]dia.ExchangePair, error) {
	var pairs []dia.ExchangePair

	query := fmt.Sprintf(`
		SELECT  a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals,e.verified,e.foreignname,e.exchange
		FROM %s e 
		INNER JOIN %s a 
		ON e.id_quotetoken=a.asset_id 
		INNER JOIN %s b 
		ON e.id_basetoken=b.asset_id 
		WHERE ((a.address='%s' and a.blockchain='%s') OR (b.address='%s' and b.blockchain='%s'))`,
		exchangepairTable,
		assetTable,
		assetTable,
		asset.Address,
		asset.Blockchain,
		asset.Address,
		asset.Blockchain,
	)
	if filterVerified {
		query += fmt.Sprintf(" AND e.verified='%v'", verified)
	}

	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return pairs, err
	}
	defer rows.Close()

	for rows.Next() {
		var (
			pair          dia.ExchangePair
			quoteDecimals sql.NullInt64
			baseDecimals  sql.NullInt64
		)

		err := rows.Scan(
			&pair.UnderlyingPair.QuoteToken.Symbol,
			&pair.UnderlyingPair.QuoteToken.Name,
			&pair.UnderlyingPair.QuoteToken.Address,
			&pair.UnderlyingPair.QuoteToken.Blockchain,
			&quoteDecimals,
			&pair.UnderlyingPair.BaseToken.Symbol,
			&pair.UnderlyingPair.BaseToken.Name,
			&pair.UnderlyingPair.BaseToken.Address,
			&pair.UnderlyingPair.BaseToken.Blockchain,
			&baseDecimals,
			&pair.Verified,
			&pair.ForeignName,
			&pair.Exchange,
		)
		if err != nil {
			return pairs, err
		}
		if quoteDecimals.Valid {
			pair.UnderlyingPair.QuoteToken.Decimals = uint8(quoteDecimals.Int64)
		}
		if baseDecimals.Valid {
			pair.UnderlyingPair.BaseToken.Decimals = uint8(baseDecimals.Int64)
		}
		pair.Symbol = pair.UnderlyingPair.QuoteToken.Symbol

		pairs = append(pairs, pair)
	}

	return pairs, nil
}

// GetNumPairs returns the number of exchangepairs/pools on @exchange.
func (rdb *RelDB) GetNumPairs(exchange dia.Exchange) (numPairs int, err error) {

	exchangeType := GetExchangeType(exchange)
	switch exchangeType {
	case "CEX":
		pairs, err := rdb.GetExchangePairSymbols(exchange.Name)
		if err != nil {
			return len(pairs), err
		}
		numPairs = len(pairs)
	case "DEX":
		pools, err := rdb.GetAllPoolAddrsExchange(exchange.Name, float64(0))
		if err != nil {
			return len(pools), err
		}
		numPairs = len(pools)
	}

	return
}

// GetAllExchangeAssets returns all assets traded as quotetoken on a CEX.
func (rdb *RelDB) GetAllExchangeAssets(verified bool) (assets []dia.Asset, err error) {
	query := fmt.Sprintf(`
		SELECT DISTINCT (a.address,a.blockchain), a.symbol,a.name,a.decimals FROM %s a
		INNER JOIN %s ep
		ON a.asset_id=ep.id_quotetoken
		WHERE ep.verified=%v
		`, assetTable, exchangepairTable, verified)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			asset    dia.Asset
			tmp      []interface{}
			decimals sql.NullInt64
		)

		err = rows.Scan(&tmp, &asset.Symbol, &asset.Name, &decimals)
		if err != nil {
			return
		}
		if decimals.Valid {
			asset.Decimals = uint8(decimals.Int64)
		}
		if len(tmp) == 2 {
			asset.Address = tmp[0].(string)
			asset.Blockchain = tmp[1].(string)
		}
		assets = append(assets, asset)
	}
	return
}
