package models

import (
	"encoding/json"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers"
	log "github.com/sirupsen/logrus"
	"time"
)

const (
	WindowYesterday = 24 * 60 * 60
	Window2         = 24 * 60 * 60 * 8
	BufferTTL       = 60 * 60
	BiggestWindow   = Window2
	TimeOutRedis    = time.Duration(time.Second * (BiggestWindow + BufferTTL))
)

// MarshalBinary -
func (e *Quotation) MarshalBinary() ([]byte, error) {
	return json.Marshal(e)
}

// UnmarshalBinary -
func (e *Quotation) UnmarshalBinary(data []byte) error {
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

func (a *DB) SetPriceUSD(symbol string, price float64) error {

	return a.SetQuotation(&Quotation{
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
		log.Printf("Error: %v on GetPriceUSD %v\n", err, symbol)
		return 0.0, err
	}
	return value.Price, nil
}

func (db *DB) GetQuotation(symbol string) (*Quotation, error) {
	key := getKeyQuotation(symbol)
	value := &Quotation{}
	err := db.redisClient.Get(key).Scan(value)
	if err != nil {
		log.Error("Error: %v on GetQuotation %v\n", err, key)
		return nil, err
	}

	v, err2 := db.GetPriceYesterday(symbol, "")
	if err2 == nil {
		value.PriceYesterday = &v
	}

	v2, err2 := db.GetVolume(symbol)
	value.VolumeYesterdayUSD = v2
	return value, err
}

func (db *DB) SetQuotation(quotation *Quotation) error {
	key := getKeyQuotation(quotation.Symbol)
	log.Println("setting ", key, quotation)
	err := db.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}

func (a *DB) SetQuotationEUR(quotation *Quotation) error {
	key := getKeyQuotationEUR(quotation.Symbol)
	log.Println("setting ", key, quotation)
	err := a.redisClient.Set(key, quotation, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetQuotation %v\n", err, quotation.Symbol)
	}
	return err
}
