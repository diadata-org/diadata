package models

import (
	"strings"
	"time"
	"strconv"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

func getKeySupply(value string) string {
	return "dia_supply_" + value
}

func getKeyDiaTotalSupply() string {
	return "dia_diaTotalSupply"
}

func getKeyDiaCirculatingSupply() string {
	return "dia_diaCirculatingSupply"
}

func (db *DB) SymbolsWithASupply() ([]string, error) {
	result := []string{}
	var cursor uint64
	key := getKeySupply("")
	for {
		var keys []string
		var err error
		keys, cursor, err = db.redisClient.Scan(cursor, key+"*", 10).Result()
		if err != nil {
			log.Error("SymbolsWithASupply err", err)
			return result, err
		}
		for _, value := range keys {
			result = append(result, strings.Replace(value, key, "", 1))
		}
		if cursor == 0 {
			log.Debugf("SymbolsWithASupply %v returns %v", key, result)
			return result, nil
		}
	}
}

func (db *DB) GetLatestSupply(symbol string) (*dia.Supply, error) {
	val, err := db.GetSupply(symbol, time.Time{}, time.Time{})
	if err != nil {
		log.Error(err)
		return &dia.Supply{}, err
	}
	return &val[0], err
}

func (db *DB) GetSupply(symbol string, starttime, endtime time.Time) ([]dia.Supply, error) {
	switch symbol {
	case "MIOTA":
		retArray := []dia.Supply{}
		s := dia.Supply{
			Symbol:            "MIOTA",
			CirculatingSupply: 2779530283.0,
			Name:              helpers.NameForSymbol("MIOTA"),
			Time:              time.Now(),
			Source:            dia.Diadata,
		}
		retArray = append(retArray, s)
		return retArray, nil
	default:
		value, err := db.GetSupplyInflux(symbol, starttime, endtime)
		if err != nil {
			log.Errorf("Error: %v on GetSupply %v\n", err, symbol)
			return []dia.Supply{}, err
		}
		return value, err
	}
}

func (db *DB) SetSupply(supply *dia.Supply) error {
	key := getKeySupply(supply.Symbol)
	log.Debug("setting ", key, supply)
	err := db.redisClient.Set(key, supply, 0).Err()
	if err != nil {
		log.Errorf("Error: %v on SetSupply (redis) %v\n", err, supply.Symbol)
	}
	err = db.SaveSupplyInflux(supply)
	if err != nil {
		log.Errorf("Error: %v on SetSupply (influx) %v\n", err, supply.Symbol)
	}
	return err
}

func (db *DB) SetDiaTotalSupply(totalSupply float64) error {
	key := getKeyDiaTotalSupply()
	log.Debug("setting ", key, totalSupply)

	err := db.redisClient.Set(key, totalSupply, 0).Err()
	if err != nil {
		log.Errorf("Error: %v on SetDiaTotalSupply (redis) %v\n", err, totalSupply)
	}
	return err
}

func (db *DB) GetDiaTotalSupply() (float64, error) {
	key := getKeyDiaTotalSupply()
	value, err := db.redisClient.Get(key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetDiaTotalSupply\n", err)
		}
		return 0.0, err
	}
	retval, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Error("Cannot convert to float in GetDiaTotalSupply")
		return 0.0, err
	}
	return retval, nil
}

func (db *DB) SetDiaCirculatingSupply(circulatingSupply float64) error {
	key := getKeyDiaCirculatingSupply()
	log.Debug("setting ", key, circulatingSupply)

	err := db.redisClient.Set(key, circulatingSupply, 0).Err()
	if err != nil {
		log.Errorf("Error: %v on SetDiaCirculatingSupply (redis) %v\n", err, circulatingSupply)
	}
	return err
}

func (db *DB) GetDiaCirculatingSupply() (float64, error) {
	key := getKeyDiaCirculatingSupply()
	value, err := db.redisClient.Get(key).Result()
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetDiaCirculatingSupply\n", err)
		}
		return 0.0, err
	}
	retval, err := strconv.ParseFloat(value, 64)
	if err != nil {
		log.Error("Cannot convert to float in GetDiaCirculatingSupply")
		return 0.0, err
	}
	return retval, nil
}
