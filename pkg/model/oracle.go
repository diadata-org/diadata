package models

import (
	"context"
	"fmt"
)

func (rdb *RelDB) SetKeyPair(publickey string, privatekey string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(publickey,privatekey) VALUES ($1,$2)`, keypairTable)
	exec, err := rdb.postgresClient.Exec(context.Background(), query, publickey, privatekey)

	log.Infoln("exec", exec)
	if err != nil {
		return err
	}
	return nil
}

func (rdb *RelDB) GetKeyPairID(publickey string) string {
	query := fmt.Sprintf(`SELECT keypair_id from   %s 
	WHERE publickey=$1`, keypairTable)
	row := rdb.postgresClient.QueryRow(context.Background(), query, publickey)
	var keypair_id string
	row.Scan(&keypair_id)

	return keypair_id
}

func (rdb *RelDB) SetUserOracle(address, keypairID, creator, chainID string) error {
	query := fmt.Sprintf(`INSERT INTO %s 
	(address,keypair_id,creator,chainID) VALUES ($1,$2,$3,$4)`, useroracleTable)
	_, err := rdb.postgresClient.Exec(context.Background(), query, address, keypairID, creator, chainID)
	if err != nil {
		return err
	}
	return nil
}
