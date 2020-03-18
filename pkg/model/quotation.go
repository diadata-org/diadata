package models

import (
	"encoding/json"
	"errors"
	"strconv"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	"github.com/diadata-org/diadata/pkg/dia/helpers"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	WindowYesterday = 24 * 60 * 60
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

// getKeyInterestRate returns a string that is used as key for storing an interest
// rate in the Redis database.
// @symbol is the symbol of the interest rate (such as SFOR) set at time @date.
func getKeyInterestRate(symbol string, date time.Time) string {
	return "dia_quotation_" + symbol + "_" + date.String()
}

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
	return value, err
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

// SetInterestRate writes the interest rate struct ir into the Redis database
func (db *DB) SetInterestRate(ir *InterestRate) error {

	if db.redisClient == nil {
		return nil
	}
	// Prepare interest rate quantities for database
	key := getKeyInterestRate(ir.Symbol, ir.Time)
	// Write interest rate quantities into database
	log.Debug("setting", key, ir)
	err := db.redisClient.Set(key, ir, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetInterestRate %v\n", err, ir.Symbol)
	}
	return err
}

// GetInterestRate returns the interest rate value for the last time stamp before @date.
// If @date is an empty string it returns the rate at the latest time stamp.
// @symbol is the shorthand symbol for the requested interest rate.
// @date is a string in the format yyyy-mm-dd.
func (db *DB) GetInterestRate(symbol, date string) (*InterestRate, error) {

	if date == "" {
		date = time.Now().Format("2006-01-02")
	}
<<<<<<< HEAD
=======

>>>>>>> master
	key, _ := db.matchKeyInterestRate(symbol, date)

	// Run database querie with found key
	ir := &InterestRate{}
	err := db.redisClient.Get(key).Scan(ir)
	if err != nil {
		if err != redis.Nil {
			log.Errorf("Error: %v on GetInterestRate %v\n", err, symbol)
		}
		return ir, err
	}
	return ir, nil
}

// matchKeyInterestRate returns the key in the database db with the youngest timestamp
// younger than the date @date. Given as string formatted as "yyyy-mm-dd hh:mm:ss".
func (db *DB) matchKeyInterestRate(symbol, date string) (string, error) {

	exDate, err := db.findLastDay(symbol, date)
	if err != nil {

	}
	// Determine all database entries with given date
	pattern := "*" + symbol + "_" + exDate + "*"
	strSlice := db.redisClient.Keys(pattern).Val()

	var strSliceFormatted []string
	layout := "2006-01-02 15:04:05"
	for _, key := range strSlice {
		date, _ := time.Parse(layout, key)
		strSliceFormatted = append(strSliceFormatted, date.String())
	}
	_, index := maxString(strSliceFormatted)
	return strSlice[index], nil
}

<<<<<<< HEAD
// findLastDay returns the youngest date before @date that has an entry in the database.
// @date should be a substring of a string formatted as "yyyy-mm-dd hh:mm:ss"
func (db *DB) findLastDay(symbol, date string) (string, error) {
=======
func (db *DB) ExistInterestRate(symbol, date string) bool {
	// Returns true if a database entry with given date stamp exists, and false otherwise.
	// @date should be a substring of a string formatted as "yyyy-mm-dd hh:mm:ss".
	pattern := "*" + symbol + "_" + date + "*"
	strSlice := db.redisClient.Keys(pattern).Val()
	if len(strSlice) == 0 {
		return false
	}
	return true
}

func (db *DB) findLastDay(symbol, date string) (string, error) {
	// Return the oldest date before @date that has an entry in the database.
	// @date should be a substring of a string formatted as "yyyy-mm-dd hh:mm:ss"

>>>>>>> master
	maxDays := 30 // Remark: This could be a function parameter as well...
	for count := 0; count < maxDays; count++ {
		if db.ExistInterestRate(symbol, date) {
			return date, nil
		}
		// If date has no entry, look for one the day before
		date = getYesterday(date, "2006-01-02")
	}

	// If no entry found in the last @maxDays days return error
	err := errors.New("No database entry found in the last " + strconv.FormatInt(int64(maxDays), 10) + "days.")
	return "", err
}

// ExistInterestRate returns true if a database entry with given date stamp exists,
// and false otherwise.
// @date should be a substring of a string formatted as "yyyy-mm-dd hh:mm:ss".
func (db *DB) ExistInterestRate(symbol, date string) bool {
	pattern := "*" + symbol + "_" + date + "*"
	strSlice := db.redisClient.Keys(pattern).Val()
	if len(strSlice) == 0 {
		return false
	}
	return true
}

func getYesterday(date, layout string) string {
	// Returns the day before @date in the world of strings, formatted as @layout
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("Error: %v on date format %s\n", err, date)
	}
	yesterday := dateTime.AddDate(0, 0, -1)
	return yesterday.Format(layout)
}

func maxString(sl []string) (string, int64) {
	// Return the maximum of a slice of strings along with its index
	index := int64(0)
	max := sl[0]
	for k, entry := range sl {
		if entry > max {
			max = entry
			index = int64(k)
		}
	}
	return max, index
}
