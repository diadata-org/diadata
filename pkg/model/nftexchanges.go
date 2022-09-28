package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (rdb *RelDB) SetNFTExchange(exchange dia.NFTExchange) (err error) {
	fields := fmt.Sprintf("INSERT INTO %s (name,centralized,contract,blockchain,rest_api,ws_api,watchdog_delay) VALUES ", nftExchangeTable)
	values := "($1,$2,NULLIF($3,''),$4,NULLIF($5,''),NULLIF($6,''),$7)"
	conflict := " ON CONFLICT (name) DO UPDATE SET contract=NULLIF($3,''),rest_api=$5,ws_api=$6,watchdog_delay=$7"

	query := fields + values + conflict
	_, err = rdb.postgresClient.Exec(context.Background(), query,
		exchange.Name,
		exchange.Centralized,
		exchange.Contract,
		exchange.BlockChain.Name,
		exchange.RestAPI,
		exchange.WsAPI,
		exchange.WatchdogDelay,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFTExchange(name string) (exchange dia.Exchange, err error) {
	query := fmt.Sprintf("SELECT centralized,contract,blockchain,rest_api,ws_api,watchdog_delay FROM %s WHERE name=$1", exchangeTable)
	var contract sql.NullString
	var blockchainName sql.NullString
	var restAPI sql.NullString
	var wsAPI sql.NullString
	err = rdb.postgresClient.QueryRow(context.Background(), query, name).Scan(
		&exchange.Centralized,
		&contract,
		&blockchainName,
		&restAPI,
		&wsAPI,
		&exchange.WatchdogDelay,
	)
	if err != nil {
		return
	}
	if contract.Valid {
		exchange.Contract = contract.String
	}
	if blockchainName.Valid {
		exchange.BlockChain.Name = blockchainName.String
	}
	if restAPI.Valid {
		exchange.RestAPI = restAPI.String
	}
	if wsAPI.Valid {
		exchange.WsAPI = wsAPI.String
	}

	exchange.Name = name
	return
}

// GetAllNFTExchanges returns all nft exchanges existent in the nftexchange table.
func (rdb *RelDB) GetAllNFTExchanges() (exchanges []dia.NFTExchange, err error) {
	query := fmt.Sprintf("SELECT name,contract, centralized,blockchain,rest_api,ws_api,watchdog_delay FROM %s", nftExchangeTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []dia.NFTExchange{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var exchange dia.NFTExchange
		var contract sql.NullString
		var blockchainName sql.NullString
		var restAPI sql.NullString
		var wsAPI sql.NullString
		err := rows.Scan(
			&exchange.Name,
			&contract,
			&exchange.Centralized,
			&blockchainName,
			&restAPI,
			&wsAPI,
			&exchange.WatchdogDelay,
		)
		if err != nil {
			return []dia.NFTExchange{}, err
		}
		if contract.Valid {
			exchange.Contract = contract.String
		}
		if blockchainName.Valid {
			exchange.BlockChain.Name = blockchainName.String
		}
		if restAPI.Valid {
			exchange.RestAPI = restAPI.String
		}
		if wsAPI.Valid {
			exchange.WsAPI = wsAPI.String
		}

		exchanges = append(exchanges, exchange)
	}

	return exchanges, nil
}

// Get24HoursNFTExchangeTrades returns the number of trades in last 24 hours
func (rdb *RelDB) Get24HoursNFTExchangeTrades(exchange string) (int64, error) {

	query := fmt.Sprintf(`
	SELECT count(*) 
	FROM %s  
	WHERE trade_time>now()- INTERVAL '1 days' 
	AND trade_time<=now()
	AND marketplace='%s'`,
		NfttradeCurrTable,
		exchange,
	)

	var numTrades sql.NullInt64
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&numTrades)
	if numTrades.Valid {
		return numTrades.Int64, nil
	}
	return 0, err
}

// Get24HoursNFTExchangeVolume returns the volume traded in last 24 hours
func (rdb *RelDB) Get24HoursNFTExchangeVolume(exchange string) (float64, error) {

	query := fmt.Sprintf(`
		SELECT SUM(price::numeric) 
		FROM %s  
		WHERE trade_time>now()- INTERVAL '1 days' 
		AND trade_time<=now()
        AND marketplace='%s'`,
		NfttradeCurrTable,
		exchange,
	)

	// TO DO: address currency issue.
	var volume sql.NullFloat64
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&volume)
	if volume.Valid {
		return volume.Float64 / 1e18, nil
	}
	return 0, err
}

// GetCollectionCountByExchange returns the  number of NFT collections traded on exchange
func (rdb *RelDB) GetCollectionCountByExchange(exchange string) (int64, error) {
	query := fmt.Sprintf(`
		SELECT COUNT (DISTINCT nftclass_id) 
		FROM %s  
        WHERE marketplace='%s'`,
		NfttradeCurrTable,
		exchange,
	)

	var collections sql.NullInt64
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&collections)
	if collections.Valid {
		return collections.Int64, nil
	}
	return 0, err
}
