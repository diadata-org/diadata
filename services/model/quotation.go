package models

import (
	"encoding/json"
	"fmt"
	"github.com/diadata-org/api-golang/dia"
	"github.com/diadata-org/api-golang/dia/helpers"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
	"strconv"
	"time"
)

const (
	Window1       = 24 * 60 * 60
	Window2       = 24 * 60 * 60 * 7
	BufferTTL     = 60 * 60
	BiggestWindow = Window2
	TimeOutRedis  = time.Duration(time.Second * (BiggestWindow + BufferTTL))
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

func getKeyPriceZSET(value string) string {
	return "dia_price_" + value + "_ZSET"
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

	value.PriceYesterday = db.GetPriceZSET(symbol, Window1)

	v, err2 := db.GetVolume(symbol)
	if err2 == nil {
		value.VolumeYesterdayUSD = &v
	}

	return value, err
}

func (db *DB) GetPriceZSET(symbol string, secondsAgo int64) *float64 {
	key := getKeyPriceZSET(symbol)
	result := 0.0
	maxScore := strconv.FormatInt(time.Now().Unix()-secondsAgo, 10)

	vals, err := db.redisClient.ZRangeByScoreWithScores(key, redis.ZRangeBy{
		Min:    "-inf",
		Max:    maxScore,
		Offset: 0,
		Count:  2,
	}).Result()

	log.Printf("vals: %v on GetPriceZSET maxScore: %v", vals, maxScore)
	if err == nil {
		if len(vals) > 0 {
			fmt.Sscanf(vals[len(vals)-1].Member.(string), "%f", &result)
			log.Printf("returned value: %v", result)
			return &result
		} else {
			log.Error("cant find  price for", symbol)
		}
	}
	return nil
}

func (db *DB) SetPriceZSET(symbol string, price float64) error {

	key := getKeyPriceZSET(symbol)
	timeUnix := time.Now().Unix()
	member := strconv.FormatFloat(price, 'f', -1, 64) + " " + strconv.FormatInt(timeUnix, 10)
	err := db.redisClient.ZAdd(key, redis.Z{
		Score:  float64(timeUnix),
		Member: member,
	}).Err()
	if err != nil {
		log.Printf("Error: %v on SetPriceZSET %v\n", err, symbol)
	}
	log.Println("SetPriceZSET ", symbol, price, timeUnix)

	// purging old values
	timeUnix -= BiggestWindow

	log.Println("SetPriceZSET cut to", symbol, timeUnix, time.Now().Unix())

	err = db.redisClient.ZRemRangeByScore(key, "-inf", "("+strconv.FormatInt(timeUnix, 10)).Err()
	if err != nil {
		log.Printf("Error: %v on SetPriceZSET %v\n", err, symbol)
	}

	return err
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
