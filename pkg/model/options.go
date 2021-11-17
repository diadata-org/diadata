package models

import (
	"errors"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
)

func (datastore *DB) SetOptionMeta(optionMeta *dia.OptionMeta) error {
	if datastore.redisClient == nil {
		return errors.New("datastore has no redis client")
	}
	key := "dia_optionMeta_" + optionMeta.BaseCurrency
	log.Debug("setting ", key, optionMeta)
	err := datastore.redisClient.SAdd(key, optionMeta).Err()
	if err != nil {
		log.Printf("Error: %v on SetOptionMeta %v\n", err, key)
	}
	return err
}

func (datastore *DB) RemoveExpiredOptionMeta(baseCurrency string) error {
	if datastore.redisClient == nil {
		return errors.New("datastore has no redis client")
	}
	optionsMeta, err := datastore.GetOptionMeta(baseCurrency)
	if err != nil {
		return err
	}
	key := "dia_optionMeta_" + baseCurrency
	for _, optionMeta := range optionsMeta {
		if optionMeta.ExpirationTime.Before(time.Now()) {
			err = datastore.redisClient.SRem(key, optionMeta).Err()
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func (datastore *DB) GetOptionMeta(baseCurrency string) ([]dia.OptionMeta, error) {
	var result []dia.OptionMeta
	if datastore.redisClient == nil {
		return result, errors.New("datastore has no redis client")
	}
	key := "dia_optionMeta_" + baseCurrency
	resultStrings, err := datastore.redisClient.SMembers(key).Result()

	if err != nil {
		log.Error("GetOptionMeta: ", err)
	}

	for _, v := range resultStrings {
		currentOM := dia.OptionMeta{}
		err = currentOM.UnmarshalBinary([]byte(v))
		if err != nil {
			log.Error("GetOptionMeta: ", err)
		}
		result = append(result, currentOM)
	}
	return result, err
}
