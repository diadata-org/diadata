package models

import (
	log "github.com/sirupsen/logrus"
)


func (db *DB) SetQuotation(q *Quotation) error {
	key := "dia_foreignQuotation"
	log.Debug("setting ", key, q)
	err := db.redisClient.Set(key, q, 0).Err()
	if err != nil {
		log.Errorln("Error: on SetQuotation", err)
	}
	return err
}

func (db *DB) GetQuotation() (*Quotation, error) {
	key := "dia_foreignQuotation"
	value := &ForeignQuotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		log.Errorln("Error: on GetQuotation", err, key)
		return nil, err
	}
	return value, err
}
