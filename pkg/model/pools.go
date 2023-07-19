package models

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	clientInfluxdb "github.com/influxdata/influxdb1-client/v2"
	"github.com/jackc/pgx/v4"
)

// SavePoolInflux stores a DEX pool in influx.
func (datastore *DB) SavePoolInflux(p dia.Pool) error {

	assetvolumesEncoded, err := json.Marshal(p.Assetvolumes)
	if err != nil {
		log.Error("marshal volumes: ", err)
	}

	// Create a point and add to batch
	tags := map[string]string{
		"exchange":   p.Exchange.Name,
		"blockchain": p.Blockchain.Name,
		"address":    p.Address,
	}
	fields := map[string]interface{}{
		"volumes": string(assetvolumesEncoded),
	}

	pt, err := clientInfluxdb.NewPoint(influxDbDEXPoolTable, tags, fields, p.Time)
	if err != nil {
		log.Errorln("NewTradeInflux:", err)
	} else {
		datastore.addPoint(pt)
	}

	err = datastore.WriteBatchInflux()
	if err != nil {
		log.Errorln("Write influx batch: ", err)
	}

	return err
}

// GetPoolInflux returns all info/liquidities of pool with @poolAddress in the time-range [starttime, endtime).
func (datastore *DB) GetPoolInflux(poolAddress string, starttime time.Time, endtime time.Time) ([]dia.Pool, error) {

	pools := []dia.Pool{}
	queryString := "SELECT \"exchange\",\"blockchain\",volumes FROM %s WHERE address='%s' AND time >= %d AND time < %d ORDER BY DESC"
	q := fmt.Sprintf(queryString, influxDbDEXPoolTable, poolAddress, starttime.UnixNano(), endtime.UnixNano())

	res, err := queryInfluxDB(datastore.influxClient, q)
	if err != nil {
		return pools, err
	}
	if len(res) > 0 && len(res[0].Series) > 0 {
		for i := 0; i < len(res[0].Series[0].Values); i++ {
			var pool dia.Pool
			pool.Time, err = time.Parse(time.RFC3339, res[0].Series[0].Values[i][0].(string))
			if err != nil {
				return pools, err
			}
			pool.Exchange.Name = res[0].Series[0].Values[i][1].(string)
			if err != nil {
				return pools, err
			}
			pool.Blockchain.Name = res[0].Series[0].Values[i][2].(string)
			stat := res[0].Series[0].Values[i][3].(string)
			if err := json.Unmarshal([]byte(stat), &pool.Assetvolumes); err != nil {
				log.Error("unmarshal: ", err)
			}
			pool.Address = poolAddress
			pools = append(pools, pool)
		}
	} else {
		return pools, errors.New("parsing pool from database")
	}
	return pools, nil
}

// SetPool writes pool data into pool table and the underlying asset and liquidity data into the poolasset table.
func (rdb *RelDB) SetPool(pool dia.Pool) error {
	if len(pool.Assetvolumes) < 2 {
		return errors.New("not enough asset data on pool")
	}

	query0 := fmt.Sprintf(
		`INSERT INTO %s (exchange,blockchain,address) VALUES ($1,$2,$3)`,
		poolTable,
	)
	_, err := rdb.postgresClient.Exec(
		context.Background(),
		query0,
		pool.Exchange.Name,
		pool.Blockchain.Name,
		pool.Address,
	)
	if err != nil {
		if !strings.Contains(err.Error(), "duplicate") {
			return err
		} else {
			log.Warn("pool already exists, update liquidity")
		}
	}

	// Add assets and liquidity to the underlying poolasset table.
	var query1 string
	for i := 0; i < len(pool.Assetvolumes); i++ {
		query1 = fmt.Sprintf(
			`INSERT INTO %s (pool_id,asset_id,liquidity,liquidity_usd,time_stamp,token_index)
				VALUES ((SELECT pool_id from %s where address=$1 and blockchain=$2),(SELECT asset_id from %s where address=$3 and blockchain=$4),$5,$6,$7,$8)
				ON CONFLICT (pool_id,asset_id) 
				DO UPDATE SET liquidity=EXCLUDED.liquidity, liquidity_usd=EXCLUDED.liquidity_usd, time_stamp=EXCLUDED.time_stamp, token_index=EXCLUDED.token_index`,
			poolassetTable,
			poolTable,
			assetTable,
		)

		_, err := rdb.postgresClient.Exec(
			context.Background(),
			query1,
			pool.Address,
			pool.Blockchain.Name,
			pool.Assetvolumes[i].Asset.Address,
			pool.Assetvolumes[i].Asset.Blockchain,
			pool.Assetvolumes[i].Volume,
			pool.Assetvolumes[i].VolumeUSD,
			pool.Time,
			pool.Assetvolumes[i].Index,
		)
		if err != nil {
			return err
		}
	}
	return nil
}

// GetPoolByAddress returns the most recent pool data, i.e. liquidity.
func (rdb *RelDB) GetPoolByAddress(blockchain string, address string) (pool dia.Pool, err error) {

	var rows pgx.Rows
	query := fmt.Sprintf(`
		SELECT pa.liquidity,pa.liquidity_usd,a.symbol,a.name,a.address,a.decimals,p.exchange,pa.time_stamp,pa.token_index 
		FROM %s pa 
		INNER JOIN %s p 
		ON p.pool_id=pa.pool_id 
		INNER JOIN %s a
		ON pa.asset_id=a.asset_id 
		WHERE p.blockchain='%s'
		AND p.address='%s'`,
		poolassetTable,
		poolTable,
		assetTable,
		blockchain,
		address,
	)

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			decimals     sql.NullInt64
			index        sql.NullInt64
			timestamp    sql.NullTime
			liquidity    sql.NullFloat64
			liquidityUSD sql.NullFloat64
			assetvolume  dia.AssetVolume
		)
		err = rows.Scan(
			&liquidity,
			&liquidityUSD,
			&assetvolume.Asset.Symbol,
			&assetvolume.Asset.Name,
			&assetvolume.Asset.Address,
			&decimals,
			&pool.Exchange.Name,
			&timestamp,
			&index,
		)
		if err != nil {
			return
		}
		if decimals.Valid {
			assetvolume.Asset.Decimals = uint8(decimals.Int64)
		}
		if index.Valid {
			assetvolume.Index = uint8(index.Int64)
		}
		if timestamp.Valid {
			pool.Time = timestamp.Time
		}
		if liquidity.Valid {
			assetvolume.Volume = liquidity.Float64
		}
		if liquidityUSD.Valid {
			assetvolume.VolumeUSD = liquidityUSD.Float64
		}
		assetvolume.Asset.Blockchain = blockchain
		pool.Assetvolumes = append(pool.Assetvolumes, assetvolume)
	}

	pool.Blockchain.Name = blockchain
	pool.Address = address

	return
}

// GetAllPoolAddrsExchange returns all pool addresses available for @exchange.
func (rdb *RelDB) GetAllPoolAddrsExchange(exchange string, liquiThreshold float64) (addresses []string, err error) {
	var (
		rows  pgx.Rows
		query string
	)
	if liquiThreshold == float64(0) {
		query = fmt.Sprintf("SELECT address FROM %s WHERE exchange='%s'", poolTable, exchange)
	} else {
		query = fmt.Sprintf(`
		SELECT DISTINCT p.address 
		FROM %s p 
		INNER JOIN %s pa 
		ON p.pool_id=pa.pool_id 
		WHERE p.exchange='%s' 
		AND pa.liquidity>=%v
		`, poolTable, poolassetTable, exchange, liquiThreshold)
	}

	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var poolAddr string
		err := rows.Scan(&poolAddr)
		if err != nil {
			log.Error(err)
		}
		addresses = append(addresses, poolAddr)
	}
	return
}

// GetAllPoolsExchange returns all pool addresses available for @exchange.
// Remark that it returns each pool n times where n is the number of assets in the pool.
func (rdb *RelDB) GetAllPoolsExchange(exchange string, liquiThreshold float64) (pools []dia.Pool, err error) {
	var (
		rows  pgx.Rows
		query string
	)

	query = fmt.Sprintf(`
		SELECT exch_pools.address,a.address,a.blockchain,a.decimals,a.symbol,a.name,pa.token_index,pa.liquidity,pa.liquidity_usd
		FROM (
			SELECT p.pool_id,p.address, SUM(CASE WHEN pa.liquidity<%v THEN 1 ELSE 0 END) AS no_liqui 
			FROM %s p 
			INNER JOIN %s pa 
			ON p.pool_id=pa.pool_id 
			WHERE p.exchange='%s' 
			GROUP BY p.pool_id,p.address
			) exch_pools 
		INNER JOIN %s pa 
		ON exch_pools.pool_id=pa.pool_id 
		INNER JOIN %s a 
		ON pa.asset_id=a.asset_id 
		WHERE exch_pools.no_liqui=0;
	`, liquiThreshold, poolTable, poolassetTable, exchange, poolassetTable, assetTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	poolIndexMap := make(map[string]int)

	for rows.Next() {
		var (
			poolAddress  string
			av           dia.AssetVolume
			decimals     sql.NullInt64
			index        sql.NullInt64
			liquidity    sql.NullFloat64
			liquidityUSD sql.NullFloat64
		)
		err := rows.Scan(
			&poolAddress,
			&av.Asset.Address,
			&av.Asset.Blockchain,
			&decimals,
			&av.Asset.Symbol,
			&av.Asset.Name,
			&index,
			&liquidity,
			&liquidityUSD,
		)
		if err != nil {
			log.Error(err)
		}
		if decimals.Valid {
			av.Asset.Decimals = uint8(decimals.Int64)
		}
		if index.Valid {
			av.Index = uint8(index.Int64)
		}
		if liquidity.Valid {
			av.Volume = liquidity.Float64
		}
		if liquidityUSD.Valid {
			av.VolumeUSD = liquidityUSD.Float64
		}

		// map poolasset to pool if pool address already exists.
		if _, ok := poolIndexMap[poolAddress]; !ok {
			// Pool does not exist yet, so initialize.
			pool := dia.Pool{Exchange: dia.Exchange{Name: exchange}, Address: poolAddress, Blockchain: dia.BlockChain{Name: av.Asset.Blockchain}}
			pool.Assetvolumes = append(pool.Assetvolumes, av)
			pools = append(pools, pool)
			poolIndexMap[poolAddress] = len(pools) - 1
		} else {
			// Pool already exists, just add pool asset.
			pools[poolIndexMap[poolAddress]].Assetvolumes = append(pools[poolIndexMap[poolAddress]].Assetvolumes, av)
		}

	}
	return
}

// GetPoolsByAsset returns all pools with @asset as a pool asset and both assets have liquidity above @liquiThreshold.
// If @liquidityThresholdUSD>0 AND @liquiThreshold=0, only pools where total liquidity is available
// AND above @liquidityThresholdUSD are returned.
func (rdb *RelDB) GetPoolsByAsset(asset dia.Asset, liquidityThreshold float64, liquidityThresholdUSD float64) ([]dia.Pool, error) {
	var (
		query string
		pools []dia.Pool
	)

	query = fmt.Sprintf(`
		SELECT exch_pools.exchange,exch_pools.address,a.address,a.blockchain,a.decimals,a.symbol,a.name,pa.token_index,pa.liquidity,pa.liquidity_usd,pa.time_stamp
		FROM (
			SELECT p.exchange,p.pool_id,p.address, SUM(CASE WHEN pa.liquidity>=%v THEN 0 ELSE 1 END) AS no_liqui, SUM(CASE WHEN a.address='%s' THEN 1 ELSE 0 END) AS correct_asset 
			FROM %s p 
			INNER JOIN %s pa 
			ON p.pool_id=pa.pool_id 
			INNER JOIN %s a 
			ON pa.asset_id=a.asset_id 
			WHERE p.blockchain='%s' 
			GROUP BY p.exchange,p.pool_id,p.address
			) exch_pools 
		INNER JOIN %s pa 
		ON exch_pools.pool_id=pa.pool_id 
		INNER JOIN %s a ON pa.asset_id=a.asset_id 
		WHERE exch_pools.no_liqui=0 
		AND exch_pools.correct_asset=1
		AND pa.time_stamp IS NOT NULL;
	`, liquidityThreshold, asset.Address, poolTable, poolassetTable, assetTable, asset.Blockchain, poolassetTable, assetTable)
	rows, err := rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return pools, err
	}
	defer rows.Close()

	poolIndexMap := make(map[string]int)

	for rows.Next() {
		var (
			exchange     string
			poolAddress  string
			av           dia.AssetVolume
			decimals     sql.NullInt64
			index        sql.NullInt64
			liquidity    sql.NullFloat64
			liquidityUSD sql.NullFloat64
			timestamp    sql.NullTime
		)
		err := rows.Scan(
			&exchange,
			&poolAddress,
			&av.Asset.Address,
			&av.Asset.Blockchain,
			&decimals,
			&av.Asset.Symbol,
			&av.Asset.Name,
			&index,
			&liquidity,
			&liquidityUSD,
			&timestamp,
		)
		if err != nil {
			log.Error(err)
		}
		if decimals.Valid {
			av.Asset.Decimals = uint8(decimals.Int64)
		}
		if index.Valid {
			av.Index = uint8(index.Int64)
		}
		if liquidity.Valid {
			av.Volume = liquidity.Float64
		}
		if liquidityUSD.Valid {
			av.VolumeUSD = liquidityUSD.Float64
		}

		// map poolasset to pool if pool address already exists.
		if _, ok := poolIndexMap[poolAddress]; !ok {
			// Pool does not exist yet, so initialize.
			pool := dia.Pool{Exchange: dia.Exchange{Name: exchange}, Address: poolAddress, Blockchain: dia.BlockChain{Name: av.Asset.Blockchain}}
			if timestamp.Valid {
				pool.Time = timestamp.Time
			}
			pool.Assetvolumes = append(pool.Assetvolumes, av)
			pools = append(pools, pool)
			poolIndexMap[poolAddress] = len(pools) - 1
		} else {
			// Pool already exists, just add pool asset.
			pools[poolIndexMap[poolAddress]].Assetvolumes = append(pools[poolIndexMap[poolAddress]].Assetvolumes, av)
		}

	}

	if liquidityThresholdUSD > 0 {
		var filteredPools []dia.Pool
		for _, pool := range pools {
			totalLiquidity, lowerBound := pool.GetPoolLiquidityUSD()
			if totalLiquidity > liquidityThresholdUSD && !lowerBound {
				filteredPools = append(filteredPools, pool)
			}
		}
		return filteredPools, nil
	}

	return pools, nil
}

// GetPoolLiquiditiesUSD attempts to fill the field @VolumeUSD by fetching the price
// of the corresponding asset.
// @priceCache acts as a poor man's cache for repeated requests.
func (datastore *DB) GetPoolLiquiditiesUSD(p *dia.Pool, priceCache map[string]float64) {
	for i, av := range p.Assetvolumes {
		var price float64
		var err error
		if _, ok := priceCache[av.Asset.Identifier()]; !ok {
			price, err = datastore.GetAssetPriceUSDCache(av.Asset)
			if err != nil {
				continue
			}
			priceCache[av.Asset.Identifier()] = price
		} else {
			price = priceCache[av.Asset.Identifier()]
		}
		p.Assetvolumes[i].VolumeUSD = price * p.Assetvolumes[i].Volume
	}
}
