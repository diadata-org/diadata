package models

import (
	"context"
	"fmt"
	"strconv"
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

func (rdb *RelDB) GetKeyPairID(publickey string) string {
	query := fmt.Sprintf(`SELECT id from   %s 
	WHERE publickey=$1`, keypairTable)
	row := rdb.postgresClient.QueryRow(context.Background(), query, publickey)
	var keypair_id string
	row.Scan(&keypair_id)

	return keypair_id
}

func (rdb *RelDB) SetOracleConfig(address, keypairID, creator, symbols, chainID string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(address,keypair_id,creator,symbols,chainID) VALUES ($1,$2,$3,$4,$5) on conflict(address)  
	do
	update set address=EXCLUDED.address`, oracleconfigTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, address, keypairID, creator, symbols, chainID)
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
	rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&feederidint)
	feederId = strconv.Itoa(feederidint)

	return
}

func (rdb *RelDB) GetOracleConfig(address string) (oracleconfigid string) {
	query := fmt.Sprintf(`SELECT id FROM %s 
	WHERE address=$1`, oracleconfigTable)
	log.Infoln("GetOracleConfig query", query)
	log.Infoln("GetOracleConfig address", address)

	rdb.postgresClient.QueryRow(context.Background(), query, address).Scan(&oracleconfigid)

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

func (rdb *RelDB) GetFeederAccessByID(id string) (owner, publickey string) {
	query := fmt.Sprintf(`SELECT owner,publickey from   %s 
	WHERE id=$1`, feederaccessTable)
	rdb.postgresClient.QueryRow(context.Background(), query, id).Scan(&owner, &publickey)
	return
}
