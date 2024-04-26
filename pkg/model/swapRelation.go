package models

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
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

func (rdb *RelDB) GetSwapRelationsByBlockchain(blockchain string) (relations []dia.SwapRelationWithAssets, err error) {
	var (
		rows      pgx.Rows
		decimals0 sql.NullInt64
		decimals1 sql.NullInt64
	)

	query := fmt.Sprintf(`
		SELECT  
		sr.blockchain,
		sr.parent_address,
		a0.name as name0,
		a0.address as address0,
		a0.decimals as decimals0,
		a0.symbol as symbol0,
		a1.name as name1,
		a1.address as address1,
		a1.decimals as decimals1,
		a1.symbol as symbol1
		FROM %s sr 
		INNER JOIN %s a0 
		ON sr.asset_address0=a0.address 
		INNER JOIN %s a1
		ON sr.asset_address1=a1.address 
		WHERE sr.blockchain=$1`,
		swapRelationTable,
		assetTable,
		assetTable,
	)

	rows, err = rdb.postgresClient.Query(context.Background(), query, blockchain)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var rel dia.SwapRelationWithAssets
		var asset0 dia.Asset
		var asset1 dia.Asset
		err := rows.Scan(
			&rel.Blockchain,
			&rel.ParentAddress,
			&asset0.Name,
			&asset0.Address,
			&decimals0,
			&asset0.Symbol,
			&asset1.Name,
			&asset1.Address,
			&decimals1,
			&asset1.Symbol,
		)

		if err != nil {
			log.Error(err)
		}
		if decimals0.Valid {
			asset0.Decimals = uint8(decimals0.Int64)
		}
		if decimals1.Valid {
			asset1.Decimals = uint8(decimals1.Int64)
		}
		rel.Asset0 = asset0
		rel.Asset1 = asset1

		relations = append(relations, rel)
	}
	return
}
