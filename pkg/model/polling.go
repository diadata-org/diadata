package models

import (
	"context"
	"fmt"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
)

// -------------------------------------------------------------
// Postgres methods
// -------------------------------------------------------------

// 		-------------------------------------------------------------
// 		polling TABLE methods
// 		-------------------------------------------------------------

// SetPolling stores an pollings into postgres.
func (rdb *RelDB) SetPolling(polling dia.Polling) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (contract_address,blockchain,next_start) VALUES ($1,$2,$3) ON CONFLICT (blockchain, contract_address) DO NOTHING",
		pollingTable,
	)
	_, err := rdb.postgresClient.Exec(
		context.Background(),
		query,
		polling.ContractAddress,
		polling.Blockchain,
		polling.NextStart,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetPollingID returns the unique identifier of @polling in postgres table asset, if the entry exists.
func (rdb *RelDB) GetPollingID(polling dia.Polling) (ID string, err error) {
	query := fmt.Sprintf("SELECT polling_id FROM %s WHERE contract_address=$1 AND blockchain=$2", pollingTable)
	err = rdb.postgresClient.QueryRow(
		context.Background(),
		query,
		polling.ContractAddress,
		polling.Blockchain,
	).Scan(&ID)

	if err != nil {
		return
	}
	return
}

// GetPolling is the standard method in order to uniquely retrieve an record from polling table.
func (rdb *RelDB) GetPolling(pollingID string) (polling dia.Polling, err error) {
	query := fmt.Sprintf("SELECT contract_address,blockchain, next_start FROM %s WHERE polling_id=$1", pollingTable)
	err = rdb.postgresClient.QueryRow(context.Background(),
		query,
		pollingID,
	).Scan(
		&polling.ContractAddress,
		&polling.Blockchain,
		&polling.NextStart,
	)
	if err != nil {
		return
	}

	return
}

// UpdateNextStartInPolling updates next_start in @polling table
// It returns true if next_start succesfully updated.
func (rdb *RelDB) UpdateNextStartInPolling(pollingID string, nextStart int) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET next_start=$1 WHERE [p;;ing+ig]=$2", pollingTable)
	resp, err := rdb.postgresClient.Exec(context.Background(), query, nextStart, pollingID)
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
