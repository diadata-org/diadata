package models

import (
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// GetRates returns a (unique) list of all rates that have been written into the database
func (db *DB) GetRates() []string {
	// log.Info("Fetching set of available rates")
	allRates := db.redisClient.SMembers(keyAllRates).Val()
	return allRates
}

// GetRatesMeta returns a list of all available rate symbols along with their first
// timestamp in the database.
func (db *DB) GetRatesMeta() (RatesMeta []InterestRateMeta, err error) {
	allRates := db.GetRates()
	for _, symbol := range allRates {
		// Get first publication date
		newdate, err := db.GetFirstDate(symbol)
		if err != nil {
			return []InterestRateMeta{}, err
		}
		// Get issuing entity
		issuer, err := db.GetIssuer(symbol)
		if err != nil {
			return []InterestRateMeta{}, err
		}
		// Fill meta type
		newEntry := InterestRateMeta{symbol, newdate, issuer}
		RatesMeta = append(RatesMeta, newEntry)
	}
	return
}

// GetIssuer returns the issuing entity of the rate given by @symbol
func (db *DB) GetIssuer(symbol string) (string, error) {
	newdate, err := db.GetFirstDate(symbol)
	if err != nil {
		return "", err
	}
	layout := "2006-01-02 15:04:05 +0000 UTC"
	time, _ := time.Parse(layout, newdate)
	fmt.Println(time)
	key := getKeyInterestRate(symbol, time)
	ir := &InterestRate{}
	err = db.redisClient.Get(key).Scan(ir)
	if err != nil {
		return "", err
	}
	return ir.Source, nil
}

// GetFirstDate returns the oldest date written in the database for the rate with symbol @symbol
func (db *DB) GetFirstDate(symbol string) (string, error) {
	allSyms := db.GetRates()
	if !(utils.Contains(&allSyms, symbol)) {
		log.Errorf("The symbol %v does not exist in the database.", symbol)
		return "", errors.New("database error")
	}
	// Fetch all available keys for @symbol
	patt := "dia_quotation_" + symbol + "_*"
	// Comment: This could be improved. Should be when the database gets larger.
	allKeys := db.redisClient.Keys(patt).Val()
	oldestKey, _ := utils.MinString(allKeys)
	// Isolate the timestamp from the key
	oldestDate := strings.SplitN(oldestKey, "_", 4)[3]
	return oldestDate, nil
}
