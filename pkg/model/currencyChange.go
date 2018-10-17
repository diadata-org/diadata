package models

import (
	//"fmt"
	//"github.com/diadata-org/diadata/pkg/dia"
	//clientInfluxdb "github.com/influxdata/influxdb/client/v2"
	log "github.com/sirupsen/logrus"
)

func (db *DB) SetCurrencyChange(cc *Change) error {
	key := "dia_currencyChange"
	log.Println("setting ", key, cc)
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

/*
func (db *DB) GetCurrencyChange() (*Change, error) {
	key := "dia_currencyChange"
	result := &Change{
		[]ChangeCurrency{},
	}

	[]Point{}
	vals, err := db.redisClient.ZRange(getKeyFilterZSET("dia_currencyChange"), 0, -1).Result()
	if err == nil {
		var p Point
		for _, v := range vals {
			fmt.Sscanf(v, "%f %d", &p.Value, &p.UnixTime)
			result = append(result, p)
		}
	}
	log.Println(symbol, "GetChartPoints:%v", result)

	return result, err
}
*/
