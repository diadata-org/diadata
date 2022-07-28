package models

import (
	"encoding/json"

	"github.com/diadata-org/diadata/pkg/dia"
)

// SetDefiProtocol writes @protocol into redis
func (datastore *DB) SetDefiProtocol(protocol dia.DefiProtocol) error {
	keyProtocol := "dia_DefiProtocol_" + protocol.Name
	mProtocol, err := json.Marshal(protocol)
	if err != nil {
		return err
	}
	err = datastore.redisClient.Set(keyProtocol, mProtocol, TimeOutRedis).Err()
	if err != nil {
		return err
	}

	return nil
}

// GetDefiProtocol returns the die protocol struct by name
func (datastore *DB) GetDefiProtocol(name string) (dia.DefiProtocol, error) {
	protocol := dia.DefiProtocol{}
	key := "dia_DefiProtocol_" + name
	err := datastore.redisClient.Get(key).Scan(&protocol)
	if err != nil {
		return protocol, err
	}
	return protocol, nil
}

// GetDefiProtocols returns a slice of all available DeFi protocols
func (datastore *DB) GetDefiProtocols() ([]dia.DefiProtocol, error) {
	allProtocols := []dia.DefiProtocol{}
	pattern := "dia_DefiProtocol_*"
	allKeys := datastore.redisClient.Keys(pattern).Val()
	for _, key := range allKeys {
		protocol := dia.DefiProtocol{}
		err := datastore.redisClient.Get(key).Scan(&protocol)
		if err != nil {
			return []dia.DefiProtocol{}, err
		}
		allProtocols = append(allProtocols, protocol)
	}
	return allProtocols, nil
}
