package models

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
	log "github.com/sirupsen/logrus"
)

// SetNFTClass stores @nftClass in postgres.
func (rdb *RelDB) SetNFTClass(nftClass dia.NFTClass) error {
	query := fmt.Sprintf("insert into %s (address,symbol,name,blockchain,contract_type,category) values ($1,$2,$3,$4,$5,NULLIF($6,''))", nftclassTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, nftClass.Address, nftClass.Symbol, nftClass.Name, nftClass.Blockchain, nftClass.ContractType, nftClass.Category)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFTClass(address string, blockchain string) (nftclass dia.NFTClass, err error) {
	query := fmt.Sprintf("select symbol,name,contract_type,category from %s where address=$1 and blockchain=$2", nftclassTable)
	var category interface{}
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain).Scan(&nftclass.Symbol, &nftclass.Name, &nftclass.ContractType, &category)
	if err != nil {
		return
	}
	nftclass.Address = address
	nftclass.Blockchain = blockchain
	if category != nil {
		nftclass.Category = category.(string)
	}
	return
}

func (rdb *RelDB) GetNFTClassID(address string, blockchain string) (ID string, err error) {
	query := fmt.Sprintf("select nftclass_id from %s where address=$1 and blockchain=$2", nftclassTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain).Scan(&ID)
	if err != nil {
		return
	}
	return ID, nil
}

func (rdb *RelDB) GetNFTClassByID(id string) (nftclass dia.NFTClass, err error) {
	query := fmt.Sprintf("select address,symbol,name,blockchain,contract_type,category from %s where nftclass_id=$1", nftclassTable)
	var category interface{}
	err = rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&nftclass.Address, &nftclass.Symbol, &nftclass.Name, &nftclass.Blockchain, &nftclass.ContractType, &category)
	if err != nil {
		return
	}
	if category != nil {
		nftclass.Category = category.(string)
	}
	return
}

// GetAllNFTClasses returns all NFT classes on @blockchain.
func (rdb *RelDB) GetAllNFTClasses(blockchain string) (nftClasses []dia.NFTClass, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf("select address,symbol,name,blockchain,contract_type,category from %s where blockchain=$1 order by name desc", nftclassTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, blockchain)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nftClass dia.NFTClass
		var category pgtype.Unknown
		err := rows.Scan(&nftClass.Address, &nftClass.Symbol, &nftClass.Name, &nftClass.Blockchain, &nftClass.ContractType, &category)
		if err != nil {
			log.Error(err)
		}
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
		err := rows.Scan(&nftClass.Address, &nftClass.Symbol, &nftClass.Name, &nftClass.Blockchain, &nftClass.ContractType, &category)
		if err != nil {
			log.Error(err)
		}
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
	nftClassID, err := rdb.GetNFTClassID(nft.NFTClass.Address, nft.NFTClass.Blockchain)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (nftclass_id,tokenID,creation_time,creator_address,uri,attributes) values ($1,$2,$3,$4,$5,$6)", nftTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftClassID, nft.TokenID, nft.CreationTime, nft.CreatorAddress, nft.URI, nft.Attributes)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error) {
	query := fmt.Sprintf("select nftclass_id,creation_time,creator_address,uri,attributes from %s where tokenid=$1", nftTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query, tokenID)
	if err != nil {
		return dia.NFT{}, err
	}
	defer rows.Close()

	var nfts []dia.NFT
	var classIDs []string
	for rows.Next() {
		var nft dia.NFT
		var nftClassID string
		err = rows.Scan(&nftClassID, &nft.CreationTime, &nft.CreatorAddress, &nft.URI, &nft.Attributes)
		if err != nil {
			log.Error(err)
		}
		nfts = append(nfts, nft)
		classIDs = append(classIDs, nftClassID)
	}
	// Find the correct underlying nft class separately due too "conn busy" error:
	// https://github.com/jackc/pgx/issues/635
	for i := range nfts {
		var refclass dia.NFTClass
		refclass, err = rdb.GetNFTClassByID(classIDs[i])
		if err != nil {
			log.Error("could not find underlying nft class: ", err)
		}
		if refclass.Address == address && refclass.Blockchain == blockchain {
			nft := nfts[i]
			nft.NFTClass = refclass
			return nft, nil
		}
	}
	return dia.NFT{}, err
}

func (rdb *RelDB) GetNFTID(address string, blockchain string, tokenID string) (ID string, err error) {
	nftclassID, err := rdb.GetNFTClassID(address, blockchain)
	if err != nil {
		return
	}
	query := fmt.Sprintf("select nft_id from %s where nftclass_id=$1 and tokenID=$2 ", nftTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, nftclassID, tokenID).Scan(&ID)
	if err != nil {
		return
	}
	return
}

// GetLastBlockheightTopshot returns the last block number before timestamp given by @upperBound.
func (rdb *RelDB) GetLastBlockheightTopshot(upperBound time.Time) (uint64, error) {
	query := fmt.Sprintf("select attributes from %s where nftclass_id=(select nftclass_id from %s where address='0x0b2a3299cc857e29' and blockchain='Flow') order by creation_time desc limit 1;", nftTable, nftclassTable)
	attributes := make(map[string]interface{})
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&attributes)
	if err != nil {
		return 0, err
	}
	currentBlock := uint64(attributes["blocknumber"].(float64))
	return currentBlock, nil
}

// SetNFTTTrade stores @trade.
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

// SetNFTBid stores @bid.
func (rdb *RelDB) SetNFTBid(bid dia.NFTBid) error {
	nftID, err := rdb.GetNFTID(bid.NFT.NFTClass.Address, bid.NFT.NFTClass.Blockchain, bid.NFT.TokenID)
	if err != nil {
		return err
	}
	query := fmt.Sprintf("insert into %s (nft_id,time,blocknumber,blockposition,bid_value,currency,from_address,marketplace) values ($1,$2,$3,$4,$5,$6,$7,$8)", nftbidTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftID, bid.Time, bid.BlockNumber, bid.BlockPosition, bid.Value, bid.Currency, bid.FromAddress, bid.Exchange)
	if err != nil {
		return err
	}
	return nil
}

// GetLastBid returns the last bid on the nft with @address and @tokenID.
// Here, 'last' refers to block number and block position smaller or equal
// (in the case of block number) than @blockNumber and @blockPosition resp.
func (rdb *RelDB) GetLastNFTBid(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (nftBid dia.NFTBid, err error) {
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	if err != nil {
		return
	}

	// First fetch biggest blocknumber<=@blockNumber for given nft.
	subquery := fmt.Sprintf("select blocknumber from nftbid where nft_id='%s' and blocknumber<=%d order by blocknumber desc limit 1", nftID, blockNumber)
	// Next, restrict to largest blockPosition in this block.
	returnVars := "time,blocknumber,blockposition,bid_value,currency,from_address,tx_hash,marketplace"
	query := fmt.Sprintf("select %s from nftbid where nft_id='%s' and blocknumber=(%s) order by blockposition desc limit 1", returnVars, nftID, subquery)

	var valueString string
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&nftBid.Time, &nftBid.BlockNumber, &nftBid.BlockPosition, &valueString, &nftBid.Currency, &nftBid.FromAddress, &nftBid.TxHash, &nftBid.Exchange)
	if err != nil {
		return
	}
	return
}

func (rdb *RelDB) GetLastBlockNFTBid(nftclass dia.NFTClass) (uint64, error) {
	// TO DO
	return 0, nil
}
