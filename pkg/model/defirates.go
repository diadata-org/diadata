package models

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (db *DB) SetDefiProtocol(protocol dia.DefiProtocol) error {
	keyProtocol := "dia_DefiProtocol_" + protocol.Name
	mProtocol, err := json.Marshal(protocol)
	if err != nil {
		return err
	}
	err = db.redisClient.Set(keyProtocol, mProtocol, TimeOutRedis).Err()
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetDefiProtocol(name string) (dia.DefiProtocol, error) {
	protocol := dia.DefiProtocol{}
	key := "dia_DefiProtocol_" + name
	err := db.redisClient.Get(key).Scan(&protocol)
	if err != nil {
		return protocol, err
	}
	return protocol, nil
}
