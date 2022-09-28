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
