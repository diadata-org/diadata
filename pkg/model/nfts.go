package models

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"math"
	"math/big"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/ethereum/go-ethereum/common"
	"github.com/jackc/pgtype"
	"github.com/jackc/pgx/v4"
)

var currencyCache = make(map[string]dia.Asset)

// SetNFTClass stores @nftClass in postgres.
func (rdb *RelDB) SetNFTClass(nftClass dia.NFTClass) error {
	query := fmt.Sprintf("INSERT INTO %s (address,symbol,name,blockchain,contract_type,category) VALUES ($1,$2,$3,$4,$5,NULLIF($6,''))", nftclassTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, nftClass.Address, nftClass.Symbol, nftClass.Name, nftClass.Blockchain, nftClass.ContractType, nftClass.Category)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFTClass(address string, blockchain string) (nftclass dia.NFTClass, err error) {
	query := fmt.Sprintf("SELECT symbol,name,contract_type,category FROM %s WHERE address=$1 AND blockchain=$2", nftclassTable)
	var category sql.NullString
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain).Scan(&nftclass.Symbol, &nftclass.Name, &nftclass.ContractType, &category)
	if err != nil {
		return
	}
	nftclass.Address = address
	nftclass.Blockchain = blockchain
	if category.Valid {
		nftclass.Category = category.String
	}
	return
}

func (rdb *RelDB) GetNFTClassID(address string, blockchain string) (ID string, err error) {
	query := fmt.Sprintf("SELECT nftclass_id FROM %s WHERE address=$1 AND blockchain=$2", nftclassTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain).Scan(&ID)
	if err != nil {
		return
	}
	return ID, nil
}

func (rdb *RelDB) GetNFTClassByID(id string) (nftclass dia.NFTClass, err error) {
	query := fmt.Sprintf("SELECT address,symbol,name,blockchain,contract_type,category FROM %s WHERE nftclass_id=$1", nftclassTable)
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
	query := fmt.Sprintf("SELECT address,symbol,name,blockchain,contract_type,category FROM %s WHERE blockchain=$1 ORDER BY name DESC", nftclassTable)
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

func (rdb *RelDB) GetTradedNFTClasses(starttime time.Time) (nftClasses []dia.NFTClass, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf(`
		SELECT DISTINCT nc.address, nc.blockchain,nc.symbol,nc.name
		FROM %s nc 
		WHERE EXISTS (
			SELECT 1 
			FROM %s ntc 
			WHERE ntc.nftclass_id=nc.nftclass_id
			AND ntc.trade_time>=to_timestamp(%v)
		)`, nftclassTable, NfttradeCurrTable, starttime.Unix())
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var nftClass dia.NFTClass
		var category pgtype.Unknown
		err := rows.Scan(&nftClass.Address, &nftClass.Blockchain, &nftClass.Symbol, &nftClass.Name)
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

	query := fmt.Sprintf("SELECT address,symbol,name,blockchain,contract_type,category FROM %s LIMIT $1 OFFSET $2", nftclassTable)
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
	query := fmt.Sprintf("UPDATE %s SET category=$1 WHERE nftclass_id=$2", nftclassTable)
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
	query := fmt.Sprintf("SELECT category FROM %s", nftcategoryTable)
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
	query := fmt.Sprintf("INSERT INTO %s (nftclass_id,token_id,creation_time,creator_address,uri,attributes) VALUES ($1,$2,$3,$4,$5,$6)", nftTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftClassID, nft.TokenID, nft.CreationTime, nft.CreatorAddress, nft.URI, nft.Attributes)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetNFT(address string, blockchain string, tokenID string) (dia.NFT, error) {
	nft := dia.NFT{}
	if blockchain == dia.ETHEREUM {
		address = common.HexToAddress(address).Hex()
	}

	query := fmt.Sprintf("SELECT c.address, c.symbol, c.name, c.blockchain, c.contract_type, c.category, n.token_id, n.creation_time, n.creator_address, n.uri, n.attributes FROM %s n INNER JOIN %s c ON(c.nftclass_id=n.nftclass_id AND c.address=$1 AND c.blockchain=$2) WHERE n.token_id=$3", nftTable, nftclassTable)

	var contractType sql.NullString
	var classCat sql.NullString

	err := rdb.postgresClient.QueryRow(context.Background(), query, address, blockchain, tokenID).Scan(
		&nft.NFTClass.Address,
		&nft.NFTClass.Symbol,
		&nft.NFTClass.Name,
		&nft.NFTClass.Blockchain,
		&contractType,
		&classCat,
		&nft.TokenID,
		&nft.CreationTime,
		&nft.CreatorAddress,
		&nft.URI,
		&nft.Attributes,
	)

	if contractType.Valid {
		nft.NFTClass.ContractType = contractType.String
	}
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
	query := fmt.Sprintf("SELECT nft_id FROM %s WHERE nftclass_id=$1 AND token_id=$2 ", nftTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, nftclassID, tokenID).Scan(&ID)
	if err != nil {
		return
	}
	return
}

// GetLastBlockheightTopshot returns the last block number before timestamp given by @upperBound.
func (rdb *RelDB) GetLastBlockheightTopshot(upperBound time.Time) (uint64, error) {
	query := fmt.Sprintf("SELECT attributes FROM %s WHERE nftclass_id=(select nftclass_id FROM %s WHERE address='0x0b2a3299cc857e29' AND blockchain='Flow') ORDER BY creation_time DESC LIMIT 1;", nftTable, nftclassTable)
	attributes := make(map[string]interface{})
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&attributes)
	if err != nil {
		return 0, err
	}
	currentBlock := uint64(attributes["blocknumber"].(float64))
	return currentBlock, nil
}

// SetNFTTTrade is a wrapper for SetNFTTradeToTable that stores @trade into the main nfttrade table.
func (rdb *RelDB) SetNFTTrade(trade dia.NFTTrade) error {
	return rdb.SetNFTTradeToTable(trade, NfttradeCurrTable)
}

// SetNFTTradeToTable  stores into @table.
func (rdb *RelDB) SetNFTTradeToTable(trade dia.NFTTrade, table string) error {
	nftclassID, err := rdb.GetNFTClassID(trade.NFT.NFTClass.Address, trade.NFT.NFTClass.Blockchain)
	if err != nil {
		return err
	}
	nftID, err := rdb.GetNFTID(trade.NFT.NFTClass.Address, trade.NFT.NFTClass.Blockchain, trade.NFT.TokenID)
	if err != nil {
		return err
	}
	currencyID, err := rdb.GetAssetID(trade.Currency)
	if err != nil {
		log.Error("get currency ID: ", err)
	}
	price := trade.Price.String()
	tradeVars := "nftclass_id,nft_id,price,price_usd,transfer_from,transfer_to,currency_id,bundle_sale,block_number,trade_time,tx_hash,marketplace"
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12)", table, tradeVars)
	_, err = rdb.postgresClient.Exec(context.Background(), query, nftclassID, nftID, price, trade.PriceUSD, trade.FromAddress, trade.ToAddress, currencyID, trade.BundleSale, trade.BlockNumber, trade.Timestamp, trade.TxHash, trade.Exchange)
	if err != nil {
		return err
	}
	return nil
}

// GetLastBlockNFTTtrade returns the last blocknumber that was scraped for trades in @nftclass.
func (rdb *RelDB) GetLastBlockNFTTrade(nftclass dia.NFTClass) (blocknumber uint64, err error) {
	query := fmt.Sprintf("SELECT block_number FROM %s WHERE nftclass_id=(SELECT nftclass_id FROM %s WHERE address='%s' AND blockchain='%s') ORDER BY block_number DESC LIMIT 1;", NfttradeCurrTable, nftclassTable, nftclass.Address, nftclass.Blockchain)
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&blocknumber)
	if err != nil {
		return
	}
	return
}

// GetNFTTradesCollection returns all trades done on the nft collection given by @address and @blockchain.
func (rdb *RelDB) GetNFTTradesCollection(address string, blockchain string, starttime time.Time, endtime time.Time) (trades []dia.NFTTrade, err error) {
	var rows pgx.Rows

	tradeVars := "price,price_usd,transfer_from,transfer_to,currency_id,bundle_sale,block_number,trade_time,tx_hash,marketplace,n.token_id"
	query := fmt.Sprintf(
		`SELECT %s FROM %s nt 
		INNER JOIN %s nc 
		ON nt.nftclass_id=nc.nftclass_id 
		INNER JOIN %s n
		ON nt.nft_id=n.nft_id
		WHERE nc.blockchain='%s' AND nc.address='%s'
		AND trade_time>to_timestamp(%v) AND trade_time<to_timestamp(%v) 
		ORDER BY trade_time DESC`,
		tradeVars,
		NfttradeCurrTable,
		nftclassTable,
		nftTable,
		blockchain,
		address,
		starttime.Unix(),
		endtime.Unix(),
	)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			trade      dia.NFTTrade
			price      string
			currencyID sql.NullString
			tokenID    sql.NullString
		)
		err := rows.Scan(
			&price,
			&trade.PriceUSD,
			&trade.FromAddress,
			&trade.ToAddress,
			&currencyID,
			&trade.BundleSale,
			&trade.BlockNumber,
			&trade.Timestamp,
			&trade.TxHash,
			&trade.Exchange,
			&tokenID,
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

		if currencyID.Valid {
			if asset, ok := currencyCache[currencyID.String]; ok {
				trade.Currency = asset
			} else {
				asset, err := rdb.GetAssetByID(currencyID.String)
				if err != nil {
					log.Errorf("cannot fetch asset with postgres id %s", currencyID.String)
				}
				trade.Currency = asset
				currencyCache[currencyID.String] = asset
			}
		}
		if tokenID.Valid {
			trade.NFT.TokenID = tokenID.String
		}

		trades = append(trades, trade)
	}
	return
}

// GetNFTTrades returns all trades done on the nft given by @address, @blockchain and @tokenID.
func (rdb *RelDB) GetNFTTrades(address string, blockchain string, tokenID string, starttime time.Time, endtime time.Time) (trades []dia.NFTTrade, err error) {
	var rows pgx.Rows
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	if err != nil {
		return
	}
	tradeVars := "price,price_usd,transfer_from,transfer_to,currency_id,bundle_sale,block_number,trade_time,tx_hash,marketplace"
	query := fmt.Sprintf(
		"SELECT %s FROM %s WHERE nft_id='%s' AND trade_time>to_timestamp(%v) AND trade_time<to_timestamp(%v) ORDER BY trade_time DESC",
		tradeVars,
		NfttradeCurrTable,
		nftID,
		starttime.Unix(),
		endtime.Unix(),
	)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	log.Info("query: ", query)

	for rows.Next() {
		var trade dia.NFTTrade
		var price string
		var currencyID sql.NullString
		err := rows.Scan(
			&price,
			&trade.PriceUSD,
			&trade.FromAddress,
			&trade.ToAddress,
			&currencyID,
			&trade.BundleSale,
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

		if currencyID.Valid {
			if asset, ok := currencyCache[currencyID.String]; ok {
				trade.Currency = asset
			} else {
				asset, err := rdb.GetAssetByID(currencyID.String)
				if err != nil {
					log.Errorf("cannot fetch asset with postgres id %s", currencyID.String)
				}
				trade.Currency = asset
				currencyCache[currencyID.String] = asset
			}
		}

		trades = append(trades, trade)
	}
	return
}

// GetAllLastTrades returns the last recorded trade for each NFT from the collection given by @nftclass.
// Caution: Currently, not all dia.NFTTrade variables are returned.
func (rdb *RelDB) GetAllLastTrades(nftclass dia.NFTClass) (trades []dia.NFTTrade, err error) {
	query := fmt.Sprintf(
		`SELECT s.price,s.token_id,s.trade_time,s.address,s.blockchain,s.decimals FROM (
			SELECT DISTINCT ON (col.nft_id) ntc.price,col.token_id,ntc.trade_time,a.address,a.blockchain,a.decimals
			FROM %s ntc 
			INNER JOIN 
			(
				SELECT nft_id, token_id
				FROM %s n 
				INNER JOIN %s nc 
				ON n.nftclass_id=nc.nftclass_id 
				WHERE nc.address='%s' AND nc.blockchain='%s'
			) AS col 
			ON ntc.nft_id=col.nft_id 
			INNER JOIN %s a 
			ON ntc.currency_id=a.asset_id 
			ORDER BY col.nft_id,ntc.trade_time ASC
			) s ORDER BY s.trade_time ASC`,
		NfttradeCurrTable,
		nftTable,
		nftclassTable,
		nftclass.Address,
		nftclass.Blockchain,
		assetTable,
	)

	var rows pgx.Rows
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			trade              dia.NFTTrade
			price              string
			tokenID            sql.NullString
			tradeTime          sql.NullTime
			currencyAddress    sql.NullString
			currencyBlockchain sql.NullString
			currencyDecimals   sql.NullInt64
		)
		err = rows.Scan(
			&price,
			&tokenID,
			&tradeTime,
			&currencyAddress,
			&currencyBlockchain,
			&currencyDecimals,
		)
		if err != nil {
			return
		}
		n := new(big.Int)
		n, ok := n.SetString(price, 10)
		if !ok {
			err = errors.New("Cannot parse trade price.")
			return
		}
		trade.Price = n
		if tokenID.Valid {
			trade.NFT.TokenID = tokenID.String
		}
		if tradeTime.Valid {
			if tradeTime.Time == (time.Time{}) {
				log.Warn("zero timestamp, continue.")
				continue
			}
			trade.Timestamp = tradeTime.Time
		}
		if currencyAddress.Valid {
			trade.Currency.Address = currencyAddress.String
		}
		if currencyBlockchain.Valid {
			trade.Currency.Blockchain = currencyBlockchain.String
		}
		if currencyDecimals.Valid {
			trade.Currency.Decimals = uint8(currencyDecimals.Int64)
		}
		trades = append(trades, trade)
	}
	return
}

// GetNFTFloorLevel returns the floor price of @nftclass w.r.t. the last 24h.
// Here, floor is w.r.t the lower bound @level.
// For Ethereum, only trades with @currencies are taken into account.
func (rdb *RelDB) GetNFTFloorLevel(
	nftclass dia.NFTClass,
	timestamp time.Time,
	floorWindowSeconds time.Duration,
	currencies []dia.Asset,
	level float64,
	noBundles bool,
	exchange string,
) (floor float64, err error) {

	query := fmt.Sprintf(`
	SELECT min(tr.price::numeric)
	FROM %s tr INNER JOIN %s n
	ON tr.nftclass_id=n.nftclass_id
	WHERE tr.trade_time<=to_timestamp(%d) AND tr.trade_time>to_timestamp(%d)
	AND tr.price::numeric>%v
	AND n.address='%s' AND n.blockchain='%s'`,
		NfttradeCurrTable,
		nftclassTable,
		timestamp.Unix(),
		timestamp.Add(-floorWindowSeconds).Unix(),
		level,
		nftclass.Address,
		nftclass.Blockchain,
	)
	if exchange != "" {
		query += fmt.Sprintf(" AND tr.marketplace='%s'", exchange)
	}

	// Only take into account selected currencies for payment.
	if nftclass.Blockchain == dia.ETHEREUM || nftclass.Blockchain == dia.ASTAR {
		for i, currency := range currencies {
			if i == 0 {
				query += " AND ("
			}
			query += fmt.Sprintf("  currency_id=(SELECT asset_id FROM %s WHERE blockchain='%s' AND address='%s') ",
				assetTable,
				currency.Blockchain,
				currency.Address,
			)
			if i < len(currencies)-1 {
				query += " OR "
			} else {
				query += " ) "
			}
		}
	}
	if noBundles {
		query += " AND tr.bundle_sale=false"
	}

	var floorFloat sql.NullFloat64
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&floorFloat)
	if err != nil {
		return
	}

	if floorFloat.Valid {
		floor, _ = new(big.Float).Quo(big.NewFloat(0).SetFloat64(floorFloat.Float64), new(big.Float).SetFloat64(math.Pow10(18))).Float64()
	} else {
		err = errors.New("no result in given time-range")
		return
	}

	return
}

// GetNFTFloor returns the floor price of @nftclass w.r.t. the last 24h.
func (rdb *RelDB) GetNFTFloor(
	nftclass dia.NFTClass,
	timestamp time.Time,
	floorWindowSeconds time.Duration,
	noBundles bool,
	exchange string,
) (floor float64, err error) {
	var paymentCurrencies []dia.Asset
	switch nftclass.Blockchain {
	case dia.ETHEREUM:
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.ETHEREUM, Address: "0x0000000000000000000000000000000000000000"})
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.ETHEREUM, Address: "0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2"})
	case dia.ASTAR:
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.ASTAR, Address: "0x0000000000000000000000000000000000000000"})
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.ASTAR, Address: "0x9dA4A3a345bf6371f8e47c63Cad2293e532022dE"})
	case dia.BINANCESMARTCHAIN:
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.BINANCESMARTCHAIN, Address: "0x0000000000000000000000000000000000000000"})
		paymentCurrencies = append(paymentCurrencies, dia.Asset{Blockchain: dia.BINANCESMARTCHAIN, Address: "0xbb4CdB9CBd36B01bD1cBaEBF2De08d9173bc095c"})
	}
	return rdb.GetNFTFloorLevel(nftclass, timestamp, floorWindowSeconds, paymentCurrencies, float64(0), noBundles, exchange)
}

// GetNFTFloorRecursive returns the floor price of @nftclass. If necessary, it iterates back in time until it finds a floor price.
func (rdb *RelDB) GetNFTFloorRecursive(
	nftClass dia.NFTClass,
	timestamp time.Time,
	floorWindowSeconds time.Duration,
	stepBackLimit int,
	noBundles bool,
	exchange string,
) (floor float64, err error) {
	var (
		count      int
		foundFloor bool
	)

	for !foundFloor && count < stepBackLimit {
		floor, err = rdb.GetNFTFloor(nftClass, timestamp, floorWindowSeconds, noBundles, exchange)
		if err != nil {
			if strings.Contains(err.Error(), "no result") {
				count++
				timestamp = timestamp.Add(-floorWindowSeconds)
				continue
			} else {
				return
			}
		} else {
			foundFloor = true
		}
	}
	return
}

// GetNFTFloorRange returns a slice of floor prices in the given time range @starttime -- @endtime.
func (rdb *RelDB) GetNFTFloorRange(
	nftClass dia.NFTClass,
	starttime time.Time,
	endtime time.Time,
	floorWindowSeconds time.Duration,
	stepBackLimit int,
	noBundles bool,
	exchange string,
) (floorPrices []float64, err error) {

	// Find initial floor price by going back in time if necessary.
	floor, err := rdb.GetNFTFloorRecursive(nftClass, starttime, floorWindowSeconds, stepBackLimit, noBundles, exchange)
	if err != nil {
		if strings.Contains(err.Error(), "no result") {
			log.Warn("could not find initial floor price.")
		} else {
			return
		}
	}
	floorPrices = append(floorPrices, floor)
	starttime = starttime.Add(floorWindowSeconds)

	// Continue filling floor prices. If none is found add the last one.
	for starttime.Before(endtime) {
		floor, err := rdb.GetNFTFloor(nftClass, starttime, floorWindowSeconds, noBundles, exchange)
		if err != nil {
			if len(floorPrices) > 0 {
				floorPrices = append(floorPrices, floorPrices[len(floorPrices)-1])
			}
		} else {
			floorPrices = append(floorPrices, floor)
		}
		starttime = starttime.Add(floorWindowSeconds)
	}

	return
}

// GetTopNFTsEth returns a list of @numCollections NFT collections sorted by trading volume in [@starttime, @endtime]
// in descending order. Only takes into account trades done with ETH.
func (rdb *RelDB) GetTopNFTsEth(numCollections int, offset int64, exchanges []string, starttime time.Time, endtime time.Time) (nftVolumes []struct {
	Name       string
	Address    string
	Blockchain string
	Volume     float64
}, err error) {

	var (
		rows          pgx.Rows
		exchangeQuery string
	)

	query := fmt.Sprintf(`
	SELECT nc.name,nc.address,nc.blockchain,SUM(price::numeric) 
	FROM %s INNER JOIN %s nc 
	ON nfttradecurrent.nftclass_id=nc.nftclass_id 
	WHERE trade_time>to_timestamp(%v) 
	AND trade_time<=to_timestamp(%v) 
	AND (currency_id=(SELECT asset_id FROM %s WHERE blockchain='%s' AND address='%s') 
	OR currency_id=(SELECT asset_id FROM %s WHERE blockchain='%s' AND address='%s') ) `,
		NfttradeCurrTable,
		nftclassTable,
		starttime.Unix(),
		endtime.Unix(),
		assetTable,
		dia.ETHEREUM,
		"0x0000000000000000000000000000000000000000",
		assetTable,
		dia.ETHEREUM,
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
	)

	for i, exchange := range exchanges {
		if i == 0 {
			exchangeQuery += fmt.Sprintf(" AND (marketplace='%s' ", exchange)
			continue
		}
		exchangeQuery += fmt.Sprintf(" OR marketplace='%s' ", exchange)
	}
	if len(exchanges) > 0 {
		exchangeQuery += ") "
		query += exchangeQuery
	}

	query += fmt.Sprintf(`
	GROUP BY nc.name,nc.address,nc.blockchain
	ORDER BY sum(price::numeric) DESC LIMIT %d
	OFFSET %d`,
		numCollections,
		offset,
	)

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			name       string
			address    string
			blockchain string
			volume     float64
		)
		err = rows.Scan(
			&name,
			&address,
			&blockchain,
			&volume,
		)
		if err != nil {
			return
		}

		nftVolumes = append(nftVolumes, struct {
			Name       string
			Address    string
			Blockchain string
			Volume     float64
		}{Name: name, Address: address, Blockchain: blockchain, Volume: volume * 1e-18})
	}
	return
}

// GetNFTVolume returns the trade volume of a collection in the time-range (@starttime, @endtime].
func (rdb *RelDB) GetNFTVolume(address, blockchain, exchange string, starttime time.Time, endtime time.Time) (float64, error) {
	var query string
	if exchange == "" {
		query = fmt.Sprintf(`
	SELECT SUM(price::numeric) 
	FROM %s INNER JOIN %s nc 
	ON nfttradecurrent.nftclass_id=nc.nftclass_id 
	WHERE trade_time>to_timestamp(%v) 
	AND trade_time<=to_timestamp(%v) 
	AND nc.address='%s' AND nc.blockchain='%s'`,
			NfttradeCurrTable,
			nftclassTable,
			starttime.Unix(),
			endtime.Unix(),
			address,
			blockchain,
		)
	} else {
		query = fmt.Sprintf(`
		SELECT SUM(price::numeric) 
		FROM %s INNER JOIN %s nc 
		ON nfttradecurrent.nftclass_id=nc.nftclass_id 
		WHERE trade_time>to_timestamp(%v) 
		AND trade_time<=to_timestamp(%v) 
		AND nc.address='%s' AND nc.blockchain='%s' AND marketplace='%s' `,
			NfttradeCurrTable,
			nftclassTable,
			starttime.Unix(),
			endtime.Unix(),
			address,
			blockchain,
			exchange,
		)

	}
	// TO DO: address currency issue.
	var volume sql.NullFloat64
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&volume)
	if volume.Valid {
		return volume.Float64 / 1e18, nil
	}
	return 0, err
}

// GetNFTExchanges returns the exchanges in which nft is traded
func (rdb *RelDB) GetNFTExchanges(address string, blockchain string) (exchanges []string, err error) {
	query := fmt.Sprintf(`
	SELECT DISTINCT marketplace
	FROM %s INNER JOIN %s nc 
	ON nfttradecurrent.nftclass_id=nc.nftclass_id 
	WHERE nc.address='%s' AND nc.blockchain='%s'`,
		NfttradeCurrTable,
		nftclassTable,
		address,
		blockchain,
	)

	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}

	defer rows.Close()

	for rows.Next() {
		var (
			exchange sql.NullString
		)
		errScan := rows.Scan(&exchange)
		if errScan != nil {
			log.Error("Error getting results from db ", errScan)
		}
		if exchange.Valid {
			exchanges = append(exchanges, exchange.String)

		}
	}

	return exchanges, err
}

// GetNumNFTTrades returns the number of trades recorded in [@starttime,@endtime] on the collection on @blockchain with @address.
func (rdb *RelDB) GetNumNFTTrades(address, blockchain, exchange string, starttime time.Time, endtime time.Time) (int, error) {
	var query string
	if exchange == "" {
		query = fmt.Sprintf(`
	SELECT count(*) 
	FROM %s INNER JOIN %s nc 
	ON nfttradecurrent.nftclass_id=nc.nftclass_id 
	WHERE trade_time>to_timestamp(%v) AND trade_time<to_timestamp(%v) 
	AND nc.address='%s' AND nc.blockchain='%s'`,
			NfttradeCurrTable,
			nftclassTable,
			starttime.Unix(),
			endtime.Unix(),
			address,
			blockchain,
		)
	} else {
		query = fmt.Sprintf(`
		SELECT count(*) 
		FROM %s INNER JOIN %s nc 
		ON nfttradecurrent.nftclass_id=nc.nftclass_id 
		WHERE trade_time>to_timestamp(%v) AND trade_time<to_timestamp(%v) 
		AND nc.address='%s' AND nc.blockchain='%s' and marketplace='%s'`,
			NfttradeCurrTable,
			nftclassTable,
			starttime.Unix(),
			endtime.Unix(),
			address,
			blockchain,
			exchange,
		)

	}
	var numTrades sql.NullInt64
	err := rdb.postgresClient.QueryRow(context.Background(), query).Scan(&numTrades)
	if numTrades.Valid {
		return int(numTrades.Int64), nil
	}
	return 0, err
}

// GetNFTOffers returns all offers done on the nft given by @address, @blockchain and @tokenID.
func (rdb *RelDB) GetNFTOffers(address string, blockchain string, tokenID string) (offers []dia.NFTOffer, err error) {
	var rows pgx.Rows
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	tradeVars := "start_value,end_value,duration,from_address,auction_type,currency_symbol,currency_address,currency_decimals,blocknumber,offer_time,tx_hash,marketplace"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE nft_id='%s' ORDER BY offer_time DESC", tradeVars, nftofferTable, nftID)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			offer      dia.NFTOffer
			startvalue string
			endvalue   string
			duration   int
		)
		err := rows.Scan(
			&startvalue,
			&endvalue,
			&duration,
			&offer.FromAddress,
			&offer.AuctionType,
			&offer.CurrencySymbol,
			&offer.CurrencyAddress,
			&offer.CurrencyDecimals,
			&offer.BlockNumber,
			&offer.Timestamp,
			&offer.TxHash,
			&offer.Exchange,
		)
		if err != nil {
			return []dia.NFTOffer{}, err
		}
		if startvalue != "<nil>" {
			s := new(big.Int)
			s, ok := s.SetString(startvalue, 10)
			if !ok {
				log.Error("err parse startvalue.")
				return []dia.NFTOffer{}, err
			}
			offer.StartValue = s
		}
		if endvalue != "<nil>" {
			e := new(big.Int)
			e, ok := e.SetString(endvalue, 10)
			if !ok {
				log.Error("err parse endvalue.")
				return []dia.NFTOffer{}, err
			}
			offer.EndValue = e
		}
		offer.Duration = time.Duration(duration) * time.Second

		offers = append(offers, offer)
	}
	return
}

// GetNFTBids returns all bids done on the nft given by @address, @blockchain and @tokenID.
func (rdb *RelDB) GetNFTBids(address string, blockchain string, tokenID string) (bids []dia.NFTBid, err error) {
	var rows pgx.Rows
	nftID, err := rdb.GetNFTID(address, blockchain, tokenID)
	tradeVars := "bid_value,from_address,currency_symbol,currency_address,currency_decimals,blocknumber,bid_time,tx_hash,marketplace"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE nft_id='%s' ORDER BY bid_time DESC", tradeVars, nftbidTable, nftID)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {

		var (
			bid      dia.NFTBid
			bidvalue string
		)
		err := rows.Scan(
			&bidvalue,
			&bid.FromAddress,
			&bid.CurrencySymbol,
			&bid.CurrencyAddress,
			&bid.CurrencyDecimals,
			&bid.BlockNumber,
			&bid.Timestamp,
			&bid.TxHash,
			&bid.Exchange,
		)
		if err != nil {
			return []dia.NFTBid{}, err
		}
		if bidvalue != "<nil>" {
			e := new(big.Int)
			e, ok := e.SetString(bidvalue, 10)
			if !ok {
				log.Error("err parse bidvalue.")
				return []dia.NFTBid{}, err
			}
			bid.Value = e
		}

		bids = append(bids, bid)
	}
	return
}

// SetNFTBid stores @bid.
func (rdb *RelDB) SetNFTBid(bid dia.NFTBid) error {
	nftID, err := rdb.GetNFTID(bid.NFT.NFTClass.Address, bid.NFT.NFTClass.Blockchain, bid.NFT.TokenID)
	if err != nil {
		return err
	}
	bidVars := "nft_id,bid_value,from_address,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,bid_time,tx_hash,marketplace"
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)", nftbidTable, bidVars)
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
	subquery := fmt.Sprintf("SELECT blocknumber FROM %s WHERE nft_id='%s' AND blocknumber<=%d ORDER BY blocknumber DESC LIMIT 1", nftbidTable, nftID, blockNumber)
	// Next, restrict to largest blockPosition in this block.
	returnVars := "bid_value,from_address,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,bid_time,tx_hash,marketplace"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE nft_id='%s' AND blocknumber=(%s) ORDER BY blockposition DESC LIMIT 1", returnVars, nftbidTable, nftID, subquery)
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

// GetLastBlockNFTBid returns the last blocknumber that was scraped for bids in @nftclass.
func (rdb *RelDB) GetLastBlockNFTBid(nftclass dia.NFTClass) (blocknumber uint64, err error) {
	query := fmt.Sprintf("SELECT b.blocknumber FROM %s b INNER JOIN %s n ON b.nft_id=n.nft_id INNER JOIN %s c ON(n.nftclass_id=c.nftclass_id AND c.address='%s' and c.blockchain='%s') ORDER BY b.blocknumber DESC LIMIT 1;", nftbidTable, nftTable, nftclassTable, nftclass.Address, nftclass.Blockchain)
	log.Info("query: ", query)
	err = rdb.postgresClient.QueryRow(context.Background(), query).Scan(&blocknumber)
	if err != nil {
		return
	}
	return
}

// GetLastBlockNFTOffer returns the last blocknumber that was scraped for offers in @nftclass.
func (rdb *RelDB) GetLastBlockNFTOffer(nftclass dia.NFTClass) (blocknumber uint64, err error) {
	query := fmt.Sprintf("SELECT b.blocknumber FROM %s b INNER JOIN %s n ON b.nft_id=n.nft_id INNER JOIN %s c ON(n.nftclass_id=c.nftclass_id AND c.address='%s' and c.blockchain='%s') ORDER BY b.blocknumber DESC LIMIT 1;", nftofferTable, nftTable, nftclassTable, nftclass.Address, nftclass.Blockchain)
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
	query := fmt.Sprintf("INSERT INTO %s (%s) VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14)", nftofferTable, bidVars)
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
	subquery := fmt.Sprintf("SELECT blocknumber FROM %s WHERE nft_id='%s' AND blocknumber<=%d ORDER BY blocknumber DESC LIMIT 1", nftofferTable, nftID, blockNumber)
	// Next, restrict to largest blockPosition in this block.
	returnVars := "start_value,end_value,duration,from_address,auction_type,currency_symbol,currency_address,currency_decimals,blocknumber,blockposition,offer_time,tx_hash,marketplace"
	query := fmt.Sprintf("SELECT %s FROM %s WHERE nft_id='%s' AND blocknumber=(%s) ORDER BY blockposition DESC LIMIT 1", returnVars, nftofferTable, nftID, subquery)
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

// GetNFTClassByNameSymbol returns all nft collections which have @searchstring
// in either its name or symbol. Search is case-insensitive.
func (rdb *RelDB) GetNFTClassesByNameSymbol(searchstring string) (collections []dia.NFTClass, err error) {
	var query string
	var rows pgx.Rows

	query = fmt.Sprintf(`
	SELECT nc.address,nc.symbol,nc.name,nc.blockchain,nc.contract_type,category
	FROM %s nc 
	INNER JOIN %s nt 
	ON nc.nftclass_id=nt.nftclass_id
	WHERE (symbol ILIKE '%s%%'  or name ILIKE '%s%%')
	AND nc.blockchain='%s'
	AND (
		currency_id=(SELECT currency_id FROM asset WHERE address='%s' AND blockchain='%s') 
		OR currency_id=(SELECT currency_id FROM asset WHERE address='%s' AND blockchain='%s')
	)
	GROUP BY nc.address,nc.symbol,nc.name,nc.blockchain,nc.contract_type,nc.category
	ORDER BY SUM(nt.price::numeric) DESC`,
		nftclassTable,
		NfttradeCurrTable,
		searchstring,
		searchstring,
		dia.ETHEREUM,
		"0x0000000000000000000000000000000000000000",
		dia.ETHEREUM,
		"0xC02aaA39b223FE8D0A0e5C4F27eAD9083C756Cc2",
		dia.ETHEREUM,
	)

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}

	defer rows.Close()
	for rows.Next() {
		var (
			collection   dia.NFTClass
			contractType sql.NullString
			category     sql.NullString
		)
		err = rows.Scan(&collection.Address, &collection.Symbol, &collection.Name, &collection.Blockchain, &contractType, &category)
		if err != nil {
			return
		}
		if contractType.Valid {
			collection.ContractType = contractType.String
		}
		if category.Valid {
			collection.Category = category.String
		}

		collections = append(collections, collection)
	}
	return
}
