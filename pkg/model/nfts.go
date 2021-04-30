package models

import (
	"context"
	"fmt"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// SetNFTClass stores @nftClass in postgres.
func (rdb *RelDB) SetNFTClass(nftClass dia.NFTClass) error {
	query := fmt.Sprintf("insert into %s (address,symbol,name,blockchain,contract_type,category) values ($1,$2,$3,$4,$5,$6)", nftclassTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, nftClass.Address, nftClass.Symbol, nftClass.Name, nftClass.Blockchain, nftClass.ContractType, nftClass.Category)
	if err != nil {
		return err
	}
	return nil
}

// GetAllNFTClasses returns all NFT classes on @blockchain.
func (rdb *RelDB) GetAllNFTClasses(blockchain string) (nftClasses []dia.NFTClass, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf("select address,symbol,name,blockchain,contract_type,category from %s where blockchain=$1", nftclassTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, blockchain)
	if err != nil {
		return
	}
	defer rows.Close()

	var address string
	for rows.Next() {
		var nftClass dia.NFTClass
		err := rows.Scan(&address, &nftClass.Symbol, &nftClass.Name, &nftClass.Blockchain, &nftClass.ContractType, nftClass.Category)
		if err != nil {
			log.Error(err)
		}
		addressCommon := common.HexToAddress(address)
		nftClass.Address = addressCommon
		nftClasses = append(nftClasses, nftClass)
	}
	return
}
