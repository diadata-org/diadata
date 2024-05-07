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
		`INSERT INTO %s (contract_address,blockchain,page) 
				VALUES ($1,$2,$3) 
				ON CONFLICT (blockchain, contract_address) DO NOTHING`,
		pollingTable,
	)
	_, err := rdb.postgresClient.Exec(
		context.Background(),
		query,
		polling.ContractAddress,
		polling.Blockchain,
		polling.Page,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetPolling is the standard method in order to uniquely retrieve an record from polling table.
func (rdb *RelDB) GetPolling(contractAddress, blockchain string) (polling dia.Polling, err error) {
	query := fmt.Sprintf(
		`SELECT contract_address,blockchain, page 
				FROM %s WHERE contract_address=$1 AND blockchain=$2 `,
		pollingTable,
	)
	err = rdb.postgresClient.QueryRow(context.Background(),
		query,
		contractAddress,
		blockchain,
	).Scan(
		&polling.ContractAddress,
		&polling.Blockchain,
		&polling.Page,
	)
	if err != nil {
		return
	}

	return
}

// UpdateNextStartInPolling updates next_start in @polling table
// It returns true if next_start succesfully updated.
func (rdb *RelDB) UpdateNextStartInPolling(polling dia.Polling) (bool, error) {
	query := fmt.Sprintf("UPDATE %s SET page=$3 WHERE blockchain=$1 AND contract_address=$2", pollingTable)
	resp, err := rdb.postgresClient.Exec(
		context.Background(),
		query,
		polling.Blockchain,
		polling.ContractAddress,
		polling.Page,
	)
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
