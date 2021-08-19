package models

import (
	"context"
	"database/sql"
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
	query := fmt.Sprintf("insert into %s (nftclass_id,token_id,creation_time,creator_address,uri,attributes) values ($1,$2,$3,$4,$5,$6)", nftTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftClassID, nft.TokenID, nft.CreationTime, nft.CreatorAddress, nft.URI, nft.Attributes)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error) {
	nft := dia.NFT{}

	query := fmt.Sprintf("select c.address, c.symbol, c.name, c.blockchain, c.contract_type, c.category, n.token_id, n.creation_time, n.creator_address, n.uri, n.attributes from %s n inner join %s c on(c.nftclass_id=n.nftclass_id and c.address=$1 and c.blockchain=$2) where n.token_id=$3", nftTable, nftclassTable)

	var classCat sql.NullString

	err := rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain, tokenID).Scan(
		&nft.NFTClass.Address,
		&nft.NFTClass.Symbol,
		&nft.NFTClass.Name,
		&nft.NFTClass.Blockchain,
		&nft.NFTClass.ContractType,
		&classCat,
		&nft.TokenID,
		&nft.CreationTime,
		&nft.CreatorAddress,
		&nft.URI,
		&nft.Attributes,
	)

	if classCat.Valid {
		nft.NFTClass.Category = classCat.String
	}

	return nft, err
}

func (rdb *RelDB) GetNFTID(address string, blockchain string, tokenID string) (ID string, err error) {
	nftclassID, err := rdb.GetNFTClassID(address, blockchain)
	if err != nil {
		return
	}
	query := fmt.Sprintf("select nft_id from %s where nftclass_id=$1 and token_id=$2 ", nftTable)
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
	nftclassID, err := rdb.GetNFTClassID(trade.NFT.NFTClass.Address, trade.NFT.NFTClass.Blockchain)
	if err != nil {
		return err
	}
	nftID, err := rdb.GetNFTID(trade.NFT.NFTClass.Address, trade.NFT.NFTClass.Blockchain, trade.NFT.TokenID)
	if err != nil {
		return err
	}
	price := trade.Price.String()
	tradeVars := "nftclass_id,nft_id,price,price_usd,transfer_from,transfer_to,currency_symbol,currency_address,currency_decimals,block_number,trade_time,tx_hash,marketplace"
	query := fmt.Sprintf("insert into %s (%s) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)", nfttradeTable, tradeVars)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftclassID, nftID, price, trade.PriceUSD, trade.FromAddress, trade.ToAddress, trade.CurrencySymbol, trade.CurrencyAddress, trade.CurrencyDecimals, trade.BlockNumber, trade.Timestamp, trade.TxHash, trade.Exchange)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetLastBlockNFTTradeScraper(nftclass dia.NFTClass) (blocknumber uint64, err error) {
	query := fmt.Sprintf("select block_number from %s where nftclass_id=(select nftclass_id from %s where address='%s' and blockchain='%s') order by block_number desc limit 1;", nfttradeTable, nftclassTable, nftclass.Address, nftclass.Blockchain)
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&blocknumber)
	if err != nil {
		return
	}
	return
}

// GetNFTTrades returns all trades done on @nft.
func (rdb *RelDB) GetNFTTrades(nft dia.NFT) (trades []dia.NFTTrade, err error) {
	var rows pgx.Rows
	nftID, err := rdb.GetNFTID(nft.NFTClass.Address, nft.NFTClass.Blockchain, nft.TokenID)
	tradeVars := "price,price_usd,transfer_from,transfer_to,currency_symbol,currency_address,currency_decimals,block_number,trade_time,tx_hash,marketplace"
	query := fmt.Sprintf("select %s from %s where nft_id='%s' order by trade_time desc", tradeVars, nfttradeTable, nftID)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var trade dia.NFTTrade
		var price string
		err := rows.Scan(
			&price,
			&trade.PriceUSD,
			&trade.FromAddress,
			&trade.ToAddress,
			&trade.CurrencySymbol,
			&trade.CurrencyAddress,
			&trade.CurrencyDecimals,
			&trade.BlockNumber,
			&trade.Timestamp,
			&trade.TxHash,
			&trade.Exchange,
		)
		if err != nil {
			return []dia.NFTTrade{}, err
		}
		n := new(big.Int)
		n, ok := n.SetString(price, 10)
		if !ok {
			return []dia.NFTTrade{}, err
		}
		trade.Price = n
		trades = append(trades, trade)
	}
	return
}

// GetNFTPrice30Days returns the average price of all NFTs in @nftclass over the last 30 days.
func (rdb *RelDB) GetNFTPrice30Days(nftclass dia.NFTClass) (float64, error) {
	// TO DO
	return 0, nil
}

// SetNFTBid stores @bid.
func (rdb *RelDB) SetNFTBid(bid dia.NFTBid) error {
	nftID, err := rdb.GetNFTID(bid.NFT.NFTClass.Address, bid.NFT.NFTClass.Blockchain, bid.NFT.TokenID)
	if err != nil {
		return err
	}
	bidVars := "nft_id,bid_value,from_address,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,bid_time,tx_hash,marketplace"
	query := fmt.Sprintf("insert into %s (%s) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", nftbidTable, bidVars)
	_, err = rdb.postgresClient.Exec(
		context.Background(),
		query,
		nftID,
		bid.Value.String(),
		bid.FromAddress,
		bid.CurrencySymbol,
		bid.CurrencyAddress,
		bid.CurrencyDecimals,
		bid.BlockNumber,
		bid.BlockPosition,
		bid.Timestamp,
		bid.TxHash,
		bid.Exchange,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetLastNFTBid returns the last bid on the nft with @address and @tokenID.
// Here, 'last' refers to block number and block position smaller or equal
// (in the case of block number) than @blockNumber and @blockPosition resp.
func (rdb *RelDB) GetLastNFTBid(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (nftBid dia.NFTBid, err error) {
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	if err != nil {
		return
	}
	nftBid.NFT.NFTClass.Address = address
	nftBid.NFT.NFTClass.Blockchain = blockchain
	nftBid.NFT.TokenID = tokenID

	// First fetch biggest blocknumber<=@blockNumber for given nft.
	subquery := fmt.Sprintf("select blocknumber from %s where nft_id='%s' and blocknumber<=%d order by blocknumber desc limit 1", nftbidTable, nftID, blockNumber)
	// Next, restrict to largest blockPosition in this block.
	returnVars := "bid_value,from_address,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,bid_time,tx_hash,marketplace"
	query := fmt.Sprintf("select %s from %s where nft_id='%s' and blocknumber=(%s) order by blockposition desc limit 1", returnVars, nftbidTable, nftID, subquery)
	var txHash sql.NullString
	var bidTime sql.NullTime
	var value string
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(
		&value,
		&nftBid.FromAddress,
		&nftBid.CurrencySymbol,
		&nftBid.CurrencyAddress,
		&nftBid.CurrencyDecimals,
		&nftBid.BlockNumber,
		&nftBid.BlockPosition,
		&bidTime,
		&txHash,
		&nftBid.Exchange,
	)
	if err != nil {
		return
	}
	if bidTime.Valid {
		nftBid.Timestamp = bidTime.Time
	}
	if txHash.Valid {
		nftBid.TxHash = txHash.String
	}
	n := new(big.Int)
	n, ok := n.SetString(value, 10)
	if !ok {
		return
	}
	nftBid.Value = n
	return
}

// GetLastBlockNFTBid returns the last blocknumber that was scraped in @nftclass.
func (rdb *RelDB) GetLastBlockNFTBid(nftclass dia.NFTClass) (blocknumber uint64, err error) {
	query := fmt.Sprintf("select block_number from %s where nftclass_id=(select nftclass_id from %s where address='%s' and blockchain='%s') order by block_number desc limit 1;", nftbidTable, nftclassTable, nftclass.Address, nftclass.Blockchain)
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&blocknumber)
	if err != nil {
		return
	}
	return
}

// SetNFTOffer stores @offer in postgres.
func (rdb *RelDB) SetNFTOffer(offer dia.NFTOffer) error {
	nftID, err := rdb.GetNFTID(offer.NFT.NFTClass.Address, offer.NFT.NFTClass.Blockchain, offer.NFT.TokenID)
	if err != nil {
		return err
	}
	bidVars := "nft_id,start_value,end_value,duration,from_address,auction_type,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,offer_time,tx_hash,marketplace"
	query := fmt.Sprintf("insert into %s (%s) values ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", nftofferTable, bidVars)
	_, err = rdb.postgresClient.Exec(
		context.Background(),
		query,
		nftID,
		offer.StartValue.String(),
		offer.EndValue.String(),
		offer.Duration,
		offer.FromAddress,
		offer.AuctionType,
		offer.CurrencySymbol,
		offer.CurrencyAddress,
		offer.CurrencyDecimals,
		offer.BlockNumber,
		offer.BlockPosition,
		offer.Timestamp,
		offer.TxHash,
		offer.Exchange,
	)
	if err != nil {
		return err
	}
	return nil
}

// GetLastNFTOffer returns the last offer on the nft with @address and @tokenID.
// Here, 'last' refers to block number and block position smaller or equal
// (in the case of block number) than @blockNumber and @blockPosition resp.
func (rdb *RelDB) GetLastNFTOffer(address string, blockchain string, tokenID string, blockNumber uint64, blockPosition uint) (offer dia.NFTOffer, err error) {
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	if err != nil {
		return
	}
	offer.NFT.NFTClass.Address = address
	offer.NFT.NFTClass.Blockchain = blockchain
	offer.NFT.TokenID = tokenID

	// First fetch biggest blocknumber<=@blockNumber for given nft.
	subquery := fmt.Sprintf("select blocknumber from %s where nft_id='%s' and blocknumber<=%d order by blocknumber desc limit 1", nftofferTable, nftID, blockNumber)
	// Next, restrict to largest blockPosition in this block.
	returnVars := "start_value,end_value,duration,from_address,auction_type,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,offer_time,tx_hash,marketplace"
	query := fmt.Sprintf("select %s from %s where nft_id='%s' and blocknumber=(%s) order by blockposition desc limit 1", returnVars, nftofferTable, nftID, subquery)
	var txHash sql.NullString
	var offerTime sql.NullTime
	var startValue string
	var endValue string
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(
		&startValue,
		&endValue,
		&offer.Duration,
		&offer.FromAddress,
		&offer.AuctionType,
		&offer.CurrencySymbol,
		&offer.CurrencyAddress,
		&offer.CurrencyDecimals,
		&offer.BlockNumber,
		&offer.BlockPosition,
		&offerTime,
		&txHash,
		&offer.Exchange,
	)
	if err != nil {
		return
	}
	if offerTime.Valid {
		offer.Timestamp = offerTime.Time
	}
	if txHash.Valid {
		offer.TxHash = txHash.String
	}
	n := new(big.Int)
	n, ok := n.SetString(startValue, 10)
	if !ok {
		return
	}
	offer.StartValue = n
	n = new(big.Int)
	n, ok = n.SetString(endValue, 10)
	if !ok {
		return
	}
	offer.EndValue = n
	return
}
