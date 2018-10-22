package models

import (
	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

func getKeySupply(value string) string {
	return "dia_supply_" + value
}

func (db *DB) SymbolsWithASupply() ([]string, error) {
	result := []string{"MIOTA"}
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
			log.Debug("SymbolsWithASupply %v returns %v", key, result)
			return result, nil
		}
	}
}

func (a *DB) GetSupply(symbol string) (*dia.Supply, error) {
	switch symbol {
	case "MIOTA":
		return &dia.Supply{
			Symbol:            "MIOTA",
			CirculatingSupply: 2779530283.0,
			Name:              helpers.NameForSymbol("MIOTA"),
			Time:              time.Now(),
			Source:            dia.Diadata,
		}, nil
	default:
		key := getKeySupply(symbol)
		value := &dia.Supply{}
		err := a.redisClient.Get(key).Scan(value)
		if err != nil {
			log.Printf("Error: %v on GetSupply %v\n", err, symbol)
			return nil, err
		}
		return value, err
	}
}

func (a *DB) SetSupply(supply *dia.Supply) error {
	key := getKeySupply(supply.Symbol)
	log.Println("setting ", key, supply)
	err := a.redisClient.Set(key, supply, 0).Err()
	if err != nil {
		log.Printf("Error: %v on SetSupply %v\n", err, supply.Symbol)
	}
	return err
}
