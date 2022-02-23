package models

func (datastore *DB) SetCurrencyChange(cc *Change) error {
	key := "dia_currencyChange"
	log.Debug("setting ", key, cc)
	err := datastore.redisClient.Set(key, cc, 0).Err()
	if err != nil {
		log.Errorln("Error: on SetCurrencyChange", err)
	}
	return err
}

func (datastore *DB) GetCurrencyChange() (*Change, error) {
	key := "dia_currencyChange"
	value := &Change{}
	err := datastore.redisClient.Get(key).Scan(value)
	if err != nil {
		log.Errorln("Error: on GetCurrencyChange", err, key)
		return nil, err
	}
	return value, err
}
