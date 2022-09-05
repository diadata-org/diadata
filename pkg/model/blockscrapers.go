package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetBlockData stores @blockdata in postgres.
func (rdb *RelDB) SetBlockData(blockdata dia.BlockData) error {
	query := fmt.Sprintf("insert into %s (blockchain,block_number,block_data) values ($1,$2,$3)", blockdataTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, blockdata.BlockchainName, blockdata.BlockNumber, blockdata.Data)
	if err != nil {
		return err
	}
	return nil
}

// GetBlockData returns information on the block with @blocknumber on @blockchain.
func (rdb *RelDB) GetBlockData(blockchain string, blocknumber int64) (dia.BlockData, error) {
	var blockdata dia.BlockData

	query := fmt.Sprintf("select block_data from %s where blockchain=$1 and block_number=$2", blockdataTable)

	err := rdb.postgresClient.QueryRow(context.Background(), query, blockchain, blocknumber).Scan(
		&blockdata.Data,
	)
	if err != nil {
		return blockdata, err
	}
	blockdata.BlockchainName = blockchain
	blockdata.BlockNumber = blocknumber
	return blockdata, nil
}

// FilterBlockData returns information on a range of blocks.
// Params:
// @jsonPathFilter: will be used to do the json-path filtering.
// @blockchain: the blockchain name.
// @fromBlocknumber, @toBlockNumber: Determine the block range for the filtering.
func (rdb *RelDB) FilterBlockData(blockchain string, fromBlocknumber, toBlockNumber int64, jsonPathFilter string) ([]dia.BlockData, error) {
	blockdata := make([]dia.BlockData, 0)

	blockUpperBoundCondition := ""
	if toBlockNumber > 0 {
		blockUpperBoundCondition = "and block_number<=$3"
	}
	query := fmt.Sprintf("select * from %s where blockchain=$1 and block_number=>$2 %s and block_data @@ $4 order by block_number asc",
		blockUpperBoundCondition, blockdataTable)

	err := rdb.postgresClient.QueryRow(context.Background(), query, blockchain, fromBlocknumber, toBlockNumber, jsonPathFilter).Scan(
		&blockdata,
	)
	if err != nil {
		return blockdata, err
	}
	return blockdata, nil
}

// GetLastBlockBlockscraper returns the last scraped block on @blockchain for block data scrapers.
func (rdb *RelDB) GetLastBlockBlockscraper(blockchain string) (blockNumber int64, err error) {
	query := fmt.Sprintf("select block_number from %s where blockchain=$1 order by block_number desc limit 1", blockdataTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, blockchain).Scan(
		&blockNumber,
	)
	if err != nil {
		return
	}
	return
}
