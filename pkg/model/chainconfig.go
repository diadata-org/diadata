package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (rdb *RelDB) SetChainConfig(chainconfig dia.ChainConfig) (err error) {
	fields := fmt.Sprintf("INSERT INTO %s (rpcurl,wsurl,chainID) VALUES ", chainconfigTable)
	values := "($1,$2,$3)"

	query := fields + values
	_, err = rdb.postgresClient.Exec(context.Background(), query,
		chainconfig.RestURL,
		chainconfig.WSURL,
		chainconfig.ChainID,
	)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetAllChainConfig() (chainconfigs []dia.ChainConfig, err error) {
	query := fmt.Sprintf("SELECT rpcurl,wsurl,chainID FROM %s", chainconfigTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []dia.ChainConfig{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var chainconfig dia.ChainConfig
		var rpcurl string
		var wsurl string
		var chainID string

		err := rows.Scan(
			&rpcurl,
			&wsurl,
			&chainID,
		)
		if err != nil {
			return []dia.ChainConfig{}, err
		}
		chainconfig.RestURL = rpcurl
		chainconfig.WSURL = wsurl
		chainconfig.ChainID = chainID

		chainconfigs = append(chainconfigs, chainconfig)
	}

	return chainconfigs, nil
}

func (rdb *RelDB) GetAllChains() (chainconfigs []dia.ChainConfig, err error) {
	query := fmt.Sprintf("SELECT chainID FROM %s", chainconfigTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return []dia.ChainConfig{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var chainconfig dia.ChainConfig
		var chainID string

		err := rows.Scan(
			&chainID,
		)
		if err != nil {
			return []dia.ChainConfig{}, err
		}

		chainconfig.ChainID = chainID

		chainconfigs = append(chainconfigs, chainconfig)
	}

	return chainconfigs, nil
}
