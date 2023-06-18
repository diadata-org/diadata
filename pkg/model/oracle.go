package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/jackc/pgx/v4"
)

func (rdb *RelDB) SetKeyPair(publickey string, privatekey string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(publickey,privatekey) VALUES ($1,$2) 
	 on conflict(publickey)  
	do
	update set publickey=EXCLUDED.publickey`, keypairTable)
	exec, err := rdb.postgresClient.Exec(context.Background(), query, publickey, privatekey)

	log.Infoln("exec", exec)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetKeyPairID(publicKey string) string {
	query := fmt.Sprintf(`SELECT id from   %s 
	WHERE publickey=$1`, keypairTable)
	rows := rdb.postgresClient.QueryRow(context.Background(), query, publicKey)
	var keypairId string

	err := rows.Scan(&keypairId)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}

	return keypairId
}

func (rdb *RelDB) SetOracleConfig(address, feederID, owner, feederAddress, symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, mandatoryFrequency string) error {
	currentTime := time.Now()
	query := fmt.Sprintf(`INSERT INTO %s 
	(address,feeder_id,owner,symbols,chainID,frequency,sleepseconds, deviationpermille,blockchainnode, mandatory_frequency,feeder_address,createddate, lastupdate) 
	VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)
	on CONFLICT(feeder_id)  
	DO UPDATE SET symbols=$4,frequency=$6,sleepseconds=$7, deviationpermille=$8, blockchainnode=$9, mandatory_frequency=$10, feeder_address=$11, lastupdate=$13`, oracleconfigTable)

	log.Infoln("SetOracleConfig Query", query)
	_, err := rdb.postgresClient.Exec(context.Background(), query, address, feederID, owner, symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, mandatoryFrequency, feederAddress, currentTime, currentTime)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetFeederID(address string) (feederId string) {
	query := fmt.Sprintf(`SELECT id FROM %s 
	WHERE owner=$1`, feederaccessTable)
	log.Infoln("GetFeederID query", query)
	log.Infoln("address", address)
	var feederidint int
	err := rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&feederidint)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	feederId = strconv.Itoa(feederidint)

	return
}

func (rdb *RelDB) SetFeederConfig(feederid, oracleconfigid string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(id, oracleconfig_id) VALUES ($1,$2) on conflict(id)  
	do
	update set oracleconfig_id=EXCLUDED.oracleconfig_id`, feederconfigTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, feederid, oracleconfigid)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetFeederAccessByID(id string) (owner string) {
	query := fmt.Sprintf(`SELECT owner from   %s 
	WHERE feeder_id=$1`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&owner)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) GetFeederByID(id string) (owner string) {
	query := fmt.Sprintf(`SELECT owner from   %s 
	WHERE feeder_id=$1`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&owner)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}
func (rdb *RelDB) GetFeederLimit(owner string) (limit int) {
	query := fmt.Sprintf(`SELECT total from  %s 
	WHERE owner=$1`, feederResourceTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, owner).Scan(&limit)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) GetTotalFeeder(owner string) (total int) {
	query := fmt.Sprintf(`SELECT count(*) from  %s 
	WHERE owner=$1 and active=true`, oracleconfigTable)
	err := rdb.postgresClient.QueryRow(context.Background(), query, owner).Scan(&total)
	if err != nil {
		log.Error("Error getting results from db ", err)
	}
	return
}

func (rdb *RelDB) GetAllFeeders() (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows           pgx.Rows
		deviationFloat float64
	)

	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp),deleted
	FROM %s 
	`, oracleconfigTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig dia.OracleConfig
			symbols      string
		)
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress, &oracleconfig.CreatedDate, &oracleconfig.LastUpdate, &oracleconfig.Deleted)
		if err != nil {
			log.Error(err)
		}

		oracleconfig.Symbols = strings.Split(symbols, ",")
		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}
func (rdb *RelDB) GetFeederResources() (addresses []string, err error) {
	var (
		rows pgx.Rows
	)
	query := fmt.Sprintf(`
	SELECT owner
	FROM %s`, feederResourceTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			address string
		)
		err := rows.Scan(&address)
		if err != nil {
			log.Error(err)
		}
		addresses = append(addresses, address)
	}
	return

}

func (rdb *RelDB) GetOraclesByOwner(owner string) (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows           pgx.Rows
		deviationFloat float64
	)

	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode, active,mandatory_frequency, feeder_address, createddate, COALESCE(lastupdate, '0001-01-01 00:00:00'::timestamp)
	FROM %s 
	WHERE owner=$1 and deleted=false`, oracleconfigTable)
	rows, err = rdb.postgresClient.Query(context.Background(), query, owner)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var (
			oracleconfig dia.OracleConfig
			symbols      string
		)
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode, &oracleconfig.Active, &oracleconfig.MandatoryFrequency, &oracleconfig.FeederAddress, &oracleconfig.CreatedDate, &oracleconfig.LastUpdate)
		if err != nil {
			log.Error(err)
		}

		oracleconfig.Symbols = strings.Split(symbols, ",")

		if oracleconfig.DeviationPermille != "" {
			deviationFloat, err = strconv.ParseFloat(oracleconfig.DeviationPermille, 64)
			if err != nil {
				log.Error(err)

			}
			oracleconfig.DeviationPermille = fmt.Sprintf("%.2f", deviationFloat/10)
		}

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}

func (rdb *RelDB) GetOracleConfig(address string) (oracleconfig dia.OracleConfig, err error) {
	var (
		symbols string
	)
	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainid, deviationpermille, sleepseconds,frequency, blockchainnode, mandatory_frequency
	FROM %s 
	WHERE address=$1`, oracleconfigTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.DeviationPermille, &oracleconfig.SleepSeconds, &oracleconfig.Frequency, &oracleconfig.BlockchainNode, &oracleconfig.MandatoryFrequency)
	if err != nil {
		return
	}

	oracleconfig.Symbols = strings.Split(symbols, " ")

	return
}

func (rdb *RelDB) ChangeOracleState(feederID string, active bool) (err error) {
	currentTime := time.Now()

	query := fmt.Sprintf(`
	UPDATE %s 
	SET active=$1, lastupdate=$3
	WHERE feeder_id=$2`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, active, feederID, currentTime)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) DeleteOracle(feederID string) (err error) {
	currentTime := time.Now()
	query := fmt.Sprintf(`
	UPDATE %s 
	SET deleted=$1,lastupdate=$3
	WHERE feeder_id=$2`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, true, feederID, currentTime)
	if err != nil {
		return
	}

	return
}

func (rdb *RelDB) GetOracleUpdates(address string, chainid string) ([]dia.OracleUpdate, error) {
	query := fmt.Sprintf(`
		SELECT oracle_address, transaction_hash, transaction_cost, asset_key, asset_price, update_block, update_from, from_balance, gas_cost, gas_used, chain_id
		FROM %s
		WHERE oracle_address=$1 AND chain_id=$2
	`, feederupdatesTable)

	rows, err := rdb.postgresClient.Query(context.Background(), query, address, chainid)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var updates []dia.OracleUpdate

	for rows.Next() {

		var (
			update          dia.OracleUpdate
			assetprice      string
			block           uint64
			frombalance     string
			gascost         string
			gasused         string
			transactioncost string
		)
		err := rows.Scan(
			&update.OracleAddress,
			&update.TransactionHash,
			&transactioncost,
			&update.AssetKey,
			&assetprice,
			&block,
			&update.UpdateFrom,
			&frombalance,
			&gascost,
			&gasused,
			&update.ChainID,
		)
		if err != nil {
			return nil, err
		}

		updates = append(updates, update)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return updates, nil
}
