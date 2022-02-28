package models

import (
	"encoding/json"
	"github.com/diadata-org/diadata/pkg/dia"
)

func (datastore *DB) SetItinData(token dia.ItinToken) error {
	key_itin := "dia_Itin_" + token.Itin
	mToken, err := json.Marshal(token)
	if err != nil {
		return err
	}
	err = datastore.redisClient.Set(key_itin, mToken, TimeOutRedis).Err()
	if err != nil {
		return err
	}

	key_itin_by_symbol := "dia_Itin_Symbol_" + token.Symbol
	err = datastore.redisClient.Set(key_itin_by_symbol, mToken, TimeOutRedis).Err()
	if err != nil {
		return err
	}
	return nil
}

func (datastore *DB) GetItinBySymbol(symbol string) (dia.ItinToken, error) {
	token := dia.ItinToken{}
	key := "dia_Itin_Symbol_" + symbol
	err := datastore.redisClient.Get(key).Scan(&token)
	if err != nil {
		return token, err
	}
	return token, nil
}
