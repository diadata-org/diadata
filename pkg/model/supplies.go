package models

import (
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
)

func getKeySupply(asset dia.Asset) string {
	return "dia_supply_" + asset.Blockchain + "_" + asset.Address
}

func getKeyDiaTotalSupply() string {
	return "dia_diaTotalSupply"
}

func getKeyDiaCirculatingSupply() string {
	return "dia_diaCirculatingSupply"
}

func (datastore *DB) GetLatestSupply(symbol string, relDB *RelDB) (*dia.Supply, error) {
	val, err := datastore.GetSupply(symbol, time.Time{}, time.Time{}, relDB)
	if err != nil {
		log.Error(err)
		return &dia.Supply{}, err
	}
	return &val[0], err
}

func (datastore *DB) GetSupply(symbol string, starttime, endtime time.Time, relDB *RelDB) ([]dia.Supply, error) {

	// First get asset with @symbol with largest market cap.
	topAsset, err := relDB.GetTopAssetByVolume(symbol)
	if err != nil || len(topAsset) < 1 {
		log.Error(err)
		return []dia.Supply{}, err
	}

	switch symbol {
	case "MIOTA":
		retArray := []dia.Supply{}
		s := dia.Supply{
			Asset:             dia.Asset{Name: helpers.NameForSymbol("MIOTA"), Symbol: "MIOTA"},
			CirculatingSupply: 2779530283.0,
			Time:              time.Now(),
			Source:            dia.Diadata,
		}
		retArray = append(retArray, s)
		return retArray, nil
	default:
		value, err := datastore.GetSupplyInflux(topAsset[0], starttime, endtime)
		if err != nil {
			log.Errorf("Error: %v on GetSupply %v\n", err, symbol)
			return []dia.Supply{}, err
		}
		return value, err
	}
}

func (datastore *DB) GetSupplyCache(asset dia.Asset) (supply dia.Supply, err error) {
	err = datastore.redisClient.Get(getKeySupply(asset)).Scan(&supply)
	if err != nil {
		return
	}
	return supply, nil
}

func (datastore *DB) SetSupply(supply *dia.Supply) error {
	key := getKeySupply(supply.Asset)
	log.Debug("setting ", key, supply)
	err := datastore.redisClient.Set(key, supply, 0).Err()
	if err != nil {
		log.Errorf("Error: %v on SetSupply (redis) %v\n", err, supply.Asset.Symbol)
	}
	err = datastore.SaveSupplyInflux(supply)
	if err != nil {
		log.Errorf("Error: %v on SetSupply (influx) %v\n", err, supply.Asset.Symbol)
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
