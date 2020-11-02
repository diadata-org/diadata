package models

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetDefiProtocol writes @protocol into redis
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

// GetDefiProtocol returns the die protocol struct by name
func (db *DB) GetDefiProtocol(name string) (dia.DefiProtocol, error) {
	protocol := dia.DefiProtocol{}
	key := "dia_DefiProtocol_" + name
	err := db.redisClient.Get(key).Scan(&protocol)
	if err != nil {
		return protocol, err
	}
	return protocol, nil
}

// GetDefiProtocols returns a slice of all available DeFi protocols
func (db *DB) GetDefiProtocols() ([]dia.DefiProtocol, error) {
	allProtocols := []dia.DefiProtocol{}
	pattern := "dia_DefiProtocol_*"
	allKeys := db.redisClient.Keys(pattern).Val()
	for _, key := range allKeys {
		protocol := dia.DefiProtocol{}
		err := db.redisClient.Get(key).Scan(&protocol)
		if err != nil {
			return []dia.DefiProtocol{}, err
		}
		allProtocols = append(allProtocols, protocol)
	}
	return allProtocols, nil
}
