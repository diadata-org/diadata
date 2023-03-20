package models

import (
	"context"
	"fmt"
	"strconv"
	"strings"

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

func (rdb *RelDB) SetOracleConfig(address, feederID, owner, symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(address,feeder_id,owner,symbols,chainID,frequency,sleepseconds, deviationpermille,blockchainnode) VALUES ($1,$2,$3,$4,$5,$6,$7,$8, $9) on conflict(feeder_id)  
	DO UPDATE SET symbols=$4,frequency=$6,sleepseconds=$7, deviationpermille=$8, blockchainnode=$9`, oracleconfigTable)

	log.Infoln("SetOracleConfig Query", query)
	_, err := rdb.postgresClient.Exec(context.Background(), query, address, feederID, owner, symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode)
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
		rows pgx.Rows
	)

	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, active
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
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille, &oracleconfig.Active)
		if err != nil {
			log.Error(err)
		}

		oracleconfig.Symbols = strings.Split(symbols, " ")

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}

func (rdb *RelDB) GetOraclesByOwner(owner string) (oracleconfigs []dia.OracleConfig, err error) {
	var (
		rows pgx.Rows
	)

	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainID, frequency, sleepseconds, deviationpermille, blockchainnode
	FROM %s 
	WHERE owner=$1 and active=true`, oracleconfigTable)
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
		err := rows.Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.Frequency, &oracleconfig.SleepSeconds, &oracleconfig.DeviationPermille, &oracleconfig.BlockchainNode)
		if err != nil {
			log.Error(err)
		}

		oracleconfig.Symbols = strings.Split(symbols, " ")

		oracleconfigs = append(oracleconfigs, oracleconfig)
	}
	return
}

func (rdb *RelDB) GetOracleConfig(address string) (oracleconfig dia.OracleConfig, err error) {
	var (
		symbols string
	)
	query := fmt.Sprintf(`
	SELECT address, feeder_id, owner,symbols, chainid, deviationpermille, sleepseconds,frequency, blockchainnode
	FROM %s 
	WHERE address=$1`, oracleconfigTable)
	err = rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&oracleconfig.Address, &oracleconfig.FeederID, &oracleconfig.Owner, &symbols, &oracleconfig.ChainID, &oracleconfig.DeviationPermille, &oracleconfig.SleepSeconds, &oracleconfig.Frequency, &oracleconfig.BlockchainNode)
	if err != nil {
		return
	}
	oracleconfig.Symbols = strings.Split(symbols, " ")

	return
}

func (rdb *RelDB) ChangeOracleState(feederID string, active bool) (err error) {
	query := fmt.Sprintf(`
	UPDATE %s 
	SET active=$1
	WHERE feeder_id=$2`, oracleconfigTable)
	_, err = rdb.postgresClient.Exec(context.Background(), query, active, feederID)
	if err != nil {
		return
	}

	return
}
