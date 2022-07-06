package models

import (
	"context"
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
			`INSERT INTO %s (pool_id,asset_id,liquidity)
				VALUES ((SELECT pool_id from %s where address=$1 and blockchain=$2),(SELECT asset_id from %s where address=$3 and blockchain=$4),$5)
				ON CONFLICT (pool_id,asset_id) DO UPDATE SET liquidity=EXCLUDED.liquidity`,
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
		)
		if err != nil {
			return err
		}
	}

	return nil

}

// GetAllPoolAddrsExchange returns all pool addresses available for @exchange.
func (rdb *RelDB) GetAllPoolAddrsExchange(exchange string) (addresses []string, err error) {
	var rows pgx.Rows
	query := fmt.Sprintf("SELECT address FROM %s WHERE exchange=$1", poolTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, exchange)
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
