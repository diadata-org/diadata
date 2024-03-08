package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
)

// GetExchangePair returns the unique exchange pair given by @exchange and @foreignname from postgres.
// It also returns the underlying pair if existent.
// If @caseSensitive is false case of @foreignname is ignored.
func (rdb *RelDB) GetExchangePair(exchange string, foreignname string, caseSensitive bool) (dia.ExchangePair, error) {
	var (
		exchangepair       dia.ExchangePair
		verified           bool
		query              string
		decimalsQuoteAsset sql.NullInt64
		decimalsBaseAsset  sql.NullInt64
	)

	exchangepair.Exchange = exchange

	if caseSensitive {
		query = fmt.Sprintf(`
			SELECT ep.symbol,ep.foreignname,ep.verified,a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals
			FROM %s ep
			INNER JOIN %s a
			ON ep.id_quotetoken=a.asset_id
			INNER JOIN %s b
			ON ep.id_basetoken=b.asset_id
			WHERE exchange=$1 AND foreignname=$2`,
			exchangepairTable,
			assetTable,
			assetTable,
		)
	} else {
		query = fmt.Sprintf(`
			SELECT ep.symbol,ep.foreignname,ep.verified,a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals
			FROM %s ep
			INNER JOIN %s a
			ON ep.id_quotetoken=a.asset_id
			INNER JOIN %s b
			ON ep.id_basetoken=b.asset_id
			WHERE exchange=$1 AND foreignname ILIKE $2`,
			exchangepairTable,
			assetTable,
			assetTable,
		)
	}
	err := rdb.postgresClient.QueryRow(context.Background(), query, exchange, foreignname).Scan(
		&exchangepair.Symbol,
		&exchangepair.ForeignName,
		&verified,
		&exchangepair.UnderlyingPair.QuoteToken.Symbol,
		&exchangepair.UnderlyingPair.QuoteToken.Name,
		&exchangepair.UnderlyingPair.QuoteToken.Address,
		&exchangepair.UnderlyingPair.QuoteToken.Blockchain,
		&decimalsQuoteAsset,
		&exchangepair.UnderlyingPair.BaseToken.Symbol,
		&exchangepair.UnderlyingPair.BaseToken.Name,
		&exchangepair.UnderlyingPair.BaseToken.Address,
		&exchangepair.UnderlyingPair.BaseToken.Blockchain,
		&decimalsBaseAsset,
	)
	if err != nil {
		return dia.ExchangePair{}, err
	}
	if decimalsQuoteAsset.Valid {
		exchangepair.UnderlyingPair.QuoteToken.Decimals = uint8(decimalsQuoteAsset.Int64)
	}
	if decimalsQuoteAsset.Valid {
		exchangepair.UnderlyingPair.BaseToken.Decimals = uint8(decimalsBaseAsset.Int64)
	}

	exchangepair.Verified = verified
	return exchangepair, nil
}

// SetExchangePair adds @pair to exchangepair table.
// If cache==true, it is also cached into redis
func (rdb *RelDB) SetExchangePair(exchange string, pair dia.ExchangePair, cache bool) error {
	var query string
	query = fmt.Sprintf(`
		INSERT INTO %s (symbol,foreignname,exchange) 
		SELECT $1,$2,$3 
		WHERE NOT EXISTS (SELECT 1 FROM %s WHERE symbol=$1 AND foreignname=$2 AND exchange=$3)`,
		exchangepairTable,
		exchangepairTable,
	)
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

// GetExchangePairSeparator returns the separator that is used as notation for an exchange pair.
// Examples: BTC-USDT, BTCUSDT, BTC/USDT.
func (rdb *RelDB) GetExchangePairSeparator(exchange string) (string, error) {
	var (
		foreignname      string
		symbolQuoteAsset string
		symbolBaseAsset  string
	)
	query := fmt.Sprintf(`
		SELECT ep.symbol,a.symbol,ep.foreignname 
		FROM %s ep
		INNER JOIN %s a
		ON ep.id_basetoken=a.asset_id
		WHERE exchange=$1
		LIMIT 1`,
		exchangepairTable,
		assetTable,
	)
	err := rdb.postgresClient.QueryRow(context.Background(), query, exchange).Scan(&symbolQuoteAsset, &symbolBaseAsset, &foreignname)
	if err != nil {
		return "", err
	}
	if len(strings.Split(foreignname, symbolQuoteAsset)) < 2 {
		return "", errors.New("not enough data.")
	}
	if len(strings.Split(strings.Split(foreignname, symbolQuoteAsset)[1], symbolBaseAsset)) < 1 {
		return "", errors.New("not enough data.")
	}
	separator := strings.Split(strings.Split(foreignname, symbolQuoteAsset)[1], symbolBaseAsset)[0]
	return separator, nil
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
	var (
		pairs []dia.ExchangePair
		rows  pgx.Rows
		err   error
	)
	exchangeType := GetExchangeType(exchange)
	if exchangeType != "CEX" {
		err = errors.New("query only feasible for centralized exchanges.")
		return pairs, err
	}

	query := fmt.Sprintf(`
		SELECT  a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals,e.verified,e.foreignname
		FROM %s e 
		INNER JOIN %s a 
		ON e.id_quotetoken=a.asset_id 
		INNER JOIN %s b 
		ON e.id_basetoken=b.asset_id 
		WHERE e.exchange=$1`,
		exchangepairTable,
		assetTable,
		assetTable,
	)
	if filterVerified {
		query += " AND e.verified=$2"
	}

	if filterVerified {
		rows, err = rdb.postgresClient.Query(context.Background(), query, exchange.Name, verified)
	} else {
		rows, err = rdb.postgresClient.Query(context.Background(), query, exchange.Name)
	}
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
	var (
		pairs []dia.ExchangePair
		rows  pgx.Rows
		err   error
	)

	query := fmt.Sprintf(`
		SELECT  a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals,e.verified,e.foreignname,e.exchange
		FROM %s e 
		INNER JOIN %s a 
		ON e.id_quotetoken=a.asset_id 
		INNER JOIN %s b 
		ON e.id_basetoken=b.asset_id 
		WHERE ((a.address=$1 and a.blockchain=$2) OR (b.address=$3 and b.blockchain=$4))`,
		exchangepairTable,
		assetTable,
		assetTable,
	)
	if filterVerified {
		query += " AND e.verified=$5"
	}
	if filterVerified {
		rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, asset.Address, asset.Blockchain, verified)
	} else {
		rows, err = rdb.postgresClient.Query(context.Background(), query, asset.Address, asset.Blockchain, asset.Address, asset.Blockchain)
	}
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

// GetExchangepairsByAsset returns all exchangepairs on @exchange where @asset is quotetoken for
// @basetoken=false and basetoken otherwise.
func (rdb *RelDB) GetExchangepairsByAsset(asset dia.Asset, exchange string, basetoken bool) (exchangepairs []dia.ExchangePair, err error) {
	var (
		secondExchangeQuery string
		baseQuoteQuery      string
	)
	if exchange == dia.BinanceExchange {
		secondExchangeQuery = fmt.Sprintf(" OR es.exchange='%s'", dia.Binance2Exchange)
	}
	if exchange == dia.BKEXExchange {
		secondExchangeQuery = fmt.Sprintf(" OR es.exchange='%s'", dia.BKEX2Exchange)
	}
	if basetoken {
		baseQuoteQuery = " b.address=$2 AND b.blockchain=$3 "
	} else {
		baseQuoteQuery = " a.address=$2 AND a.blockchain=$3 "
	}
	query := fmt.Sprintf(`
		SELECT a.symbol,a.name,a.address,a.blockchain,a.decimals,b.symbol,b.name,b.address,b.blockchain,b.decimals,ep.foreignname,ep.exchange
		FROM %s ep
		INNER JOIN %s a
		ON ep.id_quotetoken=a.asset_id
		INNER JOIN %s b
		ON ep.id_basetoken=b.asset_id
		WHERE (ep.exchange=$1 %s)
		AND %s
		`,
		exchangepairTable,
		assetTable,
		assetTable,
		secondExchangeQuery,
		baseQuoteQuery,
	)
	rows, err := rdb.postgresClient.Query(context.Background(), query, exchange, asset.Address, asset.Blockchain)
	if err != nil {
		return
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
			&pair.ForeignName,
			&pair.Exchange,
		)
		if err != nil {
			return exchangepairs, err
		}
		if quoteDecimals.Valid {
			pair.UnderlyingPair.QuoteToken.Decimals = uint8(quoteDecimals.Int64)
		}
		if baseDecimals.Valid {
			pair.UnderlyingPair.BaseToken.Decimals = uint8(baseDecimals.Int64)
		}
		pair.Symbol = pair.UnderlyingPair.QuoteToken.Symbol

		exchangepairs = append(exchangepairs, pair)
	}

	return exchangepairs, nil
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
		WHERE ep.verified=$1
		`,
		assetTable,
		exchangepairTable,
	)
	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query, verified)
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
