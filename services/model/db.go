package models

import (
	"github.com/diadata-org/api-golang/dia"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

type Datastore interface {
	SetVolume(symbol string, volume float64) error
	GetVolume(symbol string) (float64, error)
	SymbolsWithASupply() ([]string, error)
	SetPriceUSD(symbol string, price float64) error
	SetPriceEUR(symbol string, price float64) error
	GetPriceUSD(symbol string) (float64, error)
	GetQuotation(symbol string) (*Quotation, error)
	SetQuotation(quotation *Quotation) error
	SetQuotationEUR(quotation *Quotation) error
	GetSupply(symbol string) (*dia.Supply, error)
	SetSupply(supply *dia.Supply) error
	SetPriceZSET(symbol string, price float64) error
}

type DB struct {
	redisClient *redis.Client
}

func NewDataStore() (*DB, error) {

	db := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong2, err := db.Ping().Result()
	if err != nil {
		log.Error("NewDataStore", err)
		//	return nil, err
	}
	log.Println("NewDB", pong2)

	return &DB{db}, nil
}
