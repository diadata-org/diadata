package models

import (
	"encoding/json"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	WindowYesterday = 24 * 60 * 60
	Window1h        = 60 * 60
	Window7d        = 7 * 24 * 60 * 60
	Window14d       = 7 * 24 * 60 * 60
	Window30d       = 30 * 24 * 60 * 60
	Window2         = 24 * 60 * 60 * 8
	BufferTTL       = 60 * 60
	BiggestWindow   = Window2
	TimeOutRedis    = time.Duration(time.Second * (BiggestWindow + BufferTTL))
)

// MarshalBinary for quotations
func (e *Quotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// MarshalBinary for interest rates
func (e *InterestRate) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary for quotations
func (e *Quotation) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

// UnmarshalBinary for interest rates
func (e *InterestRate) UnmarshalBinary(data []byte) error {
	if err := json.Unmarshal(data, &e); err != nil {
		return err
	}
	return nil
}

func getKeyQuotation(value string) string {
	return "dia_quotation_USD_" + value
}

func getKeyQuotationEUR(value string) string {
	return "dia_quotation_EUR_" + value
}

// ------------------------------------------------------------------------------
// EXCHANGE RATES
// ------------------------------------------------------------------------------

func (db *DB) SetPriceUSD(symbol string, price float64) error {

	return db.SetQuotation(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   time.Now(),
	})
}

func (a *DB) SetPriceEUR(symbol string, price float64) error {
	return a.SetQuotationEUR(&Quotation{
		Symbol: symbol,
		Name:   helpers.NameForSymbol(symbol),
		Price:  price,
		Source: dia.Diadata,
		Time:   time.Now(),
	})
}

func (db *DB) GetPriceUSD(symbol string) (float64, error) {
	key := getKeyQuotation(symbol)
	value := &Quotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetPriceUSD %v\n", err, symbol)
		}
		return 0.0, err
	}
	return value.Price, nil
}

func (db *DB) GetQuotation(symbol string) (*Quotation, error) {
	key := getKeyQuotation(symbol)
	value := &Quotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetQuotation %v\n", err, key)
		}
		return nil, err
	}
	value.Name = helpers.NameForSymbol(symbol) // in case we updated the helper functions ;)
	v, err2 := db.GetPriceYesterday(symbol, "")
	if err2 == nil {
		value.PriceYesterday = &v
	}
	v2, err2 := db.GetVolume(symbol)
	value.VolumeYesterdayUSD = v2
	itin, err := db.GetItinBySymbol(symbol)
	if err != nil {
		value.ITIN = "undefined"
		log.Error(err)
	} else {
		value.ITIN = itin.Itin
	}
	return value, nil
}

func (db *DB) SetQuotation(quotation *Quotation) error {
	if db.redisClient == nil {
		return nil
	}
	key := getKeyQuotation(quotation.Symbol)
	log.Debug("setting ", key, quotation)
	err := db.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}

func (db *DB) SetQuotationEUR(quotation *Quotation) error {
	if db.redisClient == nil {
		return nil
	}
	key := getKeyQuotationEUR(quotation.Symbol)
	log.Debug("setting ", key, quotation)
	err := db.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}

func (db *DB) GetPaxgQuotationOunces() (*Quotation, error) {
	return db.GetQuotation("PAXG")
}

func (db *DB) GetPaxgQuotationGrams() (*Quotation, error) {
	q, err := db.GetQuotation("PAXG")
	if err != nil {
		return nil, err
	}
	q.Symbol = q.Symbol + "-gram"
	q.Name = q.Name + "-gram"
	q.Price = q.Price / 31.1034768
	return q, err
}
