package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

// -------------------------------------------------------------
// Postgres methods
// -------------------------------------------------------------

// 		-------------------------------------------------------------
// 		swap_relation TABLE methods
// 		-------------------------------------------------------------

// SetSwapRelation stores an swap relation into postgres.
func (rdb *RelDB) SetSwapRelation(relation dia.SwapRelation) error {
	query := fmt.Sprintf(
		"INSERT INTO %s (blockchain, parent_address,asset_address0, asset_address1) VALUES ($1,$2,$3,$4) ON CONFLICT (blockchain, parent_address) DO NOTHING",
		swapRelationTable,
	)
	_, err := rdb.postgresClient.Exec(
		context.Background(),
		query,
		relation.Blockchain,
		relation.ParentAddress,
		relation.AssetAddress0,
		relation.AssetAddress1,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetSwapRelationID returns the unique identifier of @swap_relation in postgres table asset, if the entry exists.
func (rdb *RelDB) GetSwapRelationID(relation dia.SwapRelation) (ID string, err error) {
	query := fmt.Sprintf("SELECT swap_relation_id FROM %s WHERE blockchain=$1 AND parent_address=$2", swapRelationTable)
	err = rdb.postgresClient.QueryRow(
		context.Background(),
		query,
		relation.Blockchain,
		relation.ParentAddress,
	).Scan(&ID)

	if err != nil {
		return
	}
	return
}

// GetSwapRelation is the standard method in order to uniquely retrieve an record from swap_relation table.
func (rdb *RelDB) GetSwapRelation(swapRelationID string) (relation dia.SwapRelation, err error) {
	query := fmt.Sprintf("SELECT blockchain, parent_address, asset_address0, asset_address1 FROM %s WHERE swap_relation_id=$1", swapRelationTable)
	err = rdb.postgresClient.QueryRow(context.Background(),
		query,
		swapRelationID,
	).Scan(
		&relation.Blockchain,
		&relation.ParentAddress,
		&relation.AssetAddress0,
		&relation.AssetAddress1,
	)
	if err != nil {
		return
	}

	return
}
