package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (rdb *RelDB) SetBlockData(blockdata dia.BlockData) error {
	query := fmt.Sprintf("insert into %s (blockchain,block_number,block_data) values ($1,$2,$3)", blockdataTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, blockdata.BlockchainName, blockdata.Number, blockdata.Data)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetBlockData(blockchain string, blocknumber string) (dia.BlockData, error) {
	var blockdata dia.BlockData

	query := fmt.Sprintf("select block_data from %s where blockchain=$1 and block_number=$2", blockdataTable)

	err := rdb.postgresClient.QueryRow(context.Background(), query, blockchain, blocknumber).Scan(
		&blockdata.Data,
	)
	if err != nil {
		return blockdata, err
	}
	blockdata.BlockchainName = blockchain
	blockdata.Number = blocknumber
	return blockdata, nil
}
