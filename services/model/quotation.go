package models

import (
	"encoding/json"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	WindowYesterday = 24 * 60 * 60
	Window2         = 24 * 60 * 60 * 8
	BufferTTL       = 60 * 60
	BiggestWindow   = Window2
	TimeOutRedis    = time.Duration(time.Second * (BiggestWindow + BufferTTL))
)

type Quotation struct {
	Symbol             string
	Name               string
	Price              float64
	PriceYesterday     *float64
	VolumeYesterdayUSD *float64
	Source             string
	Time               time.Time
}

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
		log.Printf("Error: %v on GetQuotation %v\n", err, symbol)
		return nil, err
	}
	// TOFIX
	v, err := db.getZSETValue(getKeyFilterZSET("MA"+strconv.Itoa(dia.BlockSizeSeconds)+"_"+symbol), time.Now().Unix()-WindowYesterday)
	if err == nil {
		value.PriceYesterday = &v
	}

	v, err2 := db.GetVolume(symbol)
	if err2 == nil {
		value.VolumeYesterdayUSD = &v
	}
	return value, err
}

func (db *DB) SetPriceZSET(key string, price float64) error {
	return db.setZSETValue(getKeyFilterZSET(key), price, time.Now().Unix(), BiggestWindow)
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
