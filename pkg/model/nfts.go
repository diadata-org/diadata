package models

import (
	"context"
	"fmt"
	"math/big"
	"strings"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// SetNFTClass stores @nftClass in postgres.
func (rdb *RelDB) SetNFTClass(nftClass dia.NFTClass) error {
	query := fmt.Sprintf("insert into %s (address,symbol,name,blockchain,contract_type,category) values ($1,$2,$3,$4,$5,NULLIF($6,''))", nftclassTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, nftClass.Address.String(), nftClass.Symbol, nftClass.Name, nftClass.Blockchain, nftClass.ContractType, nftClass.Category)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFTClassID(address common.Address, blockchain string) (ID string, err error) {
	query := fmt.Sprintf("select nftclass_id from %s where address=$1 and blockchain=$2", nftclassTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address.String(), blockchain).Scan(&ID)
	if err != nil {
		return
	}
	return ID, nil
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
		var category pgtype.Unknown
		err := rows.Scan(&address, &nftClass.Symbol, &nftClass.Name, &nftClass.Blockchain, &nftClass.ContractType, &category)
		if err != nil {
			log.Error(err)
		}
		nftClass.Address = common.HexToAddress(address)
		nftClass.Category = category.String
		nftClasses = append(nftClasses, nftClass)
	}
	return
}

// GetNFTClassPage returns @limit NFT classes with @offset.
func (rdb *RelDB) GetNFTClasses(limit, offset uint64) (nftClasses []dia.NFTClass, err error) {

	query := fmt.Sprintf("select address,symbol,name,blockchain,contract_type,category from %s LIMIT $1 OFFSET $2", nftclassTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query, limit, offset)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nftClass dia.NFTClass
		var category pgtype.Unknown
		var address string
		err := rows.Scan(&address, &nftClass.Symbol, &nftClass.Name, &nftClass.Blockchain, &nftClass.ContractType, &category)
		if err != nil {
			log.Error(err)
		}
		nftClass.Address = common.HexToAddress(address)
		nftClass.Category = category.String
		nftClasses = append(nftClasses, nftClass)

	}
	return
}

func (rdb *RelDB) UpdateNFTClassCategory(nftclassID string, category string) (bool, error) {
	query := fmt.Sprintf("update %s set category=$1 where nftclass_id=$2", nftclassTable)
	resp, err := rdb.postgresClient.Exec(context.Background(), query, category, nftclassID)
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

// GetNFTCategories returns all available NFT categories.
func (rdb *RelDB) GetNFTCategories() (categories []string, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf("select category from %s", nftcategoryTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var category string
		err := rows.Scan(&category)
		if err != nil {
			log.Error(err)
		}
		categories = append(categories, category)
	}
	return
}

func (rdb *RelDB) SetNFT(nft dia.NFT) error {
	// TO DO
	return nil
}

func (rdb *RelDB) GetNFT(address common.Address, tokenID uint64) (dia.NFT, error) {
	// TO DO
	return dia.NFT{}, nil
}

func (rdb *RelDB) SetNFTTrade(trade dia.NFTTrade) error {
	// TO DO
	return nil
}

func (rdb *RelDB) GetLastBlockNFTTrade(nft dia.NFT) (*big.Int, error) {
	// TO DO:
	// Return highest block number of recorded trades for @nft.
	// Returns 0 if no trade recorded.
	return big.NewInt(0), nil
}
