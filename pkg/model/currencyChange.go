package models

import (
	log "github.com/sirupsen/logrus"
)

func (db *DB) SetCurrencyChange(cc *Change) error {
	key := "dia_currencyChange"
	log.Debug("setting ", key, cc)
	err := db.redisClient.Set(key, cc, 0).Err()
	if err != nil {
		log.Errorln("Error: on SetCurrencyChange", err)
	}
	return err
}

func (db *DB) GetCurrencyChange() (*Change, error) {
	key := "dia_currencyChange"
	value := &Change{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		log.Errorln("Error: on GetCurrencyChange", err, key)
		return nil, err
	}
	return value, err
}
