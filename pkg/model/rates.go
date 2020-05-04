package models

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	ratedevs "github.com/diadata-org/diadata/internal/pkg/rateDerivatives"
	"github.com/diadata-org/diadata/pkg/utils"
	"github.com/go-redis/redis"
	log "github.com/sirupsen/logrus"
)

const (
	keyAllRates     = "all_rates"
	TimeLayoutRedis = "2006-01-02 15:04:05 +0000 UTC"
)

// ---------------------------------------------------------------------------------------
// Setter and getter for interest rates
// ---------------------------------------------------------------------------------------

// getKeyInterestRate returns a string that is used as key for storing an interest
// rate in the Redis database.
// @symbol is the symbol of the interest rate (such as SOFR) set at time @date.
func getKeyInterestRate(symbol string, date time.Time) string {
	return "dia_quotation_" + symbol + "_" + date.String()
}

// SetInterestRate writes the interest rate struct ir into the Redis database
// and writes rate type into a set of all available rates (if not done yet).
func (db *DB) SetInterestRate(ir *InterestRate) error {

	if db.redisClient == nil {
		return nil
	}
	// Prepare interest rate quantities for database
	key := getKeyInterestRate(ir.Symbol, ir.EffectiveDate)
	// Write interest rate quantities into database
	log.Debug("setting", key, ir)
	err := db.redisClient.Set(key, ir, TimeOutRedis).Err()
	if err != nil {
		log.Printf("Error: %v on SetInterestRate %v\n", err, ir.Symbol)
	}

	// Write rate type into set of available rates
	err = db.redisClient.SAdd(keyAllRates, ir.Symbol).Err()
	if err != nil {
		log.Printf("Error: %v on writing rate %v into set of available rates\n", err, ir.Symbol)
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

// GetInterestRateRange returns the interest rate values for a range of timestamps.
// @symbol is the shorthand symbol for the requested interest rate.
// @dateInit and @dateFinal are strings in the format yyyy-mm-dd.
func (db *DB) GetInterestRateRange(symbol, dateInit, dateFinal string) ([]*InterestRate, error) {

	// dateInit = utils.GetTomorrow(dateInit, "2006-01-02")

	// Fetch all available keys for @symbol
	patt := "dia_quotation_" + symbol + "_*"
	// Comment: This could be improved. Should be when the database gets larger.
	allKeys := db.redisClient.Keys(patt).Val()

	// Set bounds on database's keys for the requested time range
	stampInit := "dia_quotation_" + symbol + "_" + dateInit + " 00:00:00 +0000 UTC"
	stampFinal := "dia_quotation_" + symbol + "_" + dateFinal + " 23:59:59 +0000 UTC"

	// Get value for each key
	allValues := []*InterestRate{}
	for _, key := range allKeys {
		if stampInit <= key && key <= stampFinal {
			// Run database querie with key
			ir := &InterestRate{}
			err := db.redisClient.Get(key).Scan(ir)
			if err != nil {
				if err != redis.Nil {
					log.Errorf("Error: %v on Symbol %v in redis database\n", err, symbol)
				}
				return allValues, err
			}
			allValues = append(allValues, ir)
		}
	}
	return allValues, nil
}

// ---------------------------------------------------------------------------------------
// Getter for rates' metadata
// ---------------------------------------------------------------------------------------

// GetRates returns a (unique) slice of all rates that have been written into the database
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
	key := getKeyInterestRate(symbol, newdate)
	ir := &InterestRate{}
	err = db.redisClient.Get(key).Scan(ir)
	if err != nil {
		return "", err
	}
	return ir.Source, nil
}

// GetFirstDate returns the oldest date written in the database for the rate with symbol @symbol
func (db *DB) GetFirstDate(symbol string) (time.Time, error) {
	allSyms := db.GetRates()
	if !(utils.Contains(&allSyms, symbol)) {
		log.Errorf("The symbol %v does not exist in the database.", symbol)
		return time.Time{}, errors.New("database error")
	}
	// Fetch all available keys for @symbol
	patt := "dia_quotation_" + symbol + "_*"
	// Comment: This could be improved. Should be when the database gets larger.
	allKeys := db.redisClient.Keys(patt).Val()
	oldestKey, _ := utils.MinString(allKeys)

	// Scan the struct corresponding to the oldest timestamp and fetch effective date.
	ir := &InterestRate{}
	err := db.redisClient.Get(oldestKey).Scan(ir)
	if err != nil {
		return time.Time{}, err
	}
	return ir.EffectiveDate, nil
}

// ---------------------------------------------------------------------------------------
// Risk-free rates methods
// ---------------------------------------------------------------------------------------

// GetCompoundedRate returns the compounded rate for the period @dateInit to @date. It computes the rate for all
// days for which an entry is present in the database. All other days are assumed to be holidays (or weekends).
func (db *DB) GetCompoundedRate(symbol string, dateInit, date time.Time, daysPerYear int) (*InterestRate, error) {

	// Get initial date for the rate with @symbol
	firstPublication, err := db.GetFirstDate(symbol)
	if err != nil {
		return &InterestRate{}, err
	}
	if utils.AfterDay(firstPublication, dateInit) {
		log.Error("Initial date cannot be earlier than first publication date.")
	}

	// Get rate data from database
	ratesAPI, err := db.GetInterestRateRange(symbol, dateInit.Format("2006-01-02"), date.Format("2006-01-02"))
	if err != nil {
		return &InterestRate{}, err
	}

	// Determine holidays through missing database entries
	existDates := []time.Time{}
	for _, entry := range ratesAPI {
		existDates = append(existDates, (*entry).EffectiveDate)
	}
	holidays := utils.GetHolidays(existDates, dateInit, date)

	// Check whether given (last) day is holiday or weekend. In case it is, consider last workday.
	for utils.ContainsDay(holidays, date) || !utils.CheckWeekDay(date) {
		date = date.AddDate(0, 0, -1)
	}

	// Sort ratesApi (type []*InterestRates) in increasing order according to date
	// and remove the data for the final date, as only past values are compounded.
	sort.Slice(ratesAPI, func(i, j int) bool {
		return (ratesAPI[i].EffectiveDate).Before(ratesAPI[j].EffectiveDate)
	})
	ratesAPI = ratesAPI[:len(ratesAPI)-1]

	// Remove holiday if after @date
	for i, day := range holidays {
		if day.After(date) {
			holidays = holidays[:i]
			break
		}
	}

	// Extract rates' values
	rates := []float64{}
	for i := range ratesAPI {
		rates = append(rates, ratesAPI[i].Value)
	}
	fmt.Println(rates)

	compRate, err := ratedevs.CompoundedRate(rates, dateInit, date, holidays, daysPerYear)
	if err != nil {
		return &InterestRate{}, err
	}

	// Fill InterestRate type for return
	ir := &InterestRate{}
	ir.Symbol = symbol + "_dia_image"
	ir.Value = compRate
	ir.EffectiveDate = date
	ir.Source = ratesAPI[0].Source

	return ir, nil
}

// GetCompoundedIndex returns the compounded index over the maximal period of existence of @symbol
func (db *DB) GetCompoundedIndex(symbol string, date time.Time, daysPerYear int) (*InterestRate, error) {
	// Get initial date for the rate with @symbol
	dateInit, err := db.GetFirstDate(symbol)
	if err != nil {
		return &InterestRate{}, err
	}
	return db.GetCompoundedRate(symbol, dateInit, date, daysPerYear)
}

// GetCompoundedIndexRange returns a range of compounded index values of @symbol
func (db *DB) GetCompoundedIndexRange(symbol string, dateInit, dateFinal time.Time, daysPerYear int) (values []*InterestRate, err error) {

	for utils.AfterDay(dateFinal, dateInit) {
		val, err := db.GetCompoundedIndex(symbol, dateInit, daysPerYear)
		if err != nil {
			log.Error("Error on compounded Index")
		}
		values = append(values, val)
	}
	return
}

// GetCompoundedAvg returns the compounded average of the index @symbol over rolling @calDays calendar days.
func (db *DB) GetCompoundedAvg(symbol string, date time.Time, calDays, daysPerYear int) (compAvg *InterestRate, err error) {
	dateInit := date.AddDate(0, 0, -calDays)
	fmt.Println(dateInit)
	index, err := db.GetCompoundedRate(symbol, dateInit, date, daysPerYear)
	if err != nil {
		return
	}
	compAvg.Value = 100 * (index.Value - 1) * float64(daysPerYear) / float64(calDays)
	return compAvg, nil
}

// ---------------------------------------------------------------------------------------
// Auxiliary functions
// ---------------------------------------------------------------------------------------

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

// matchKeyInterestRate returns the key in the database db with the youngest timestamp
// younger than the date @date, given as substring of a string formatted as "yyyy-mm-dd hh:mm:ss".
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
	_, index := utils.MaxString(strSliceFormatted)
	return strSlice[index], nil
}

// findLastDay returns the youngest date before @date that has an entry in the database.
// @date should be a substring of a string formatted as "yyyy-mm-dd hh:mm:ss"
func (db *DB) findLastDay(symbol, date string) (string, error) {
	maxDays := 30 // Remark: This could be a function parameter as well...
	for count := 0; count < maxDays; count++ {
		if db.ExistInterestRate(symbol, date) {
			return date, nil
		}
		// If date has no entry, look for one the day before
		date = utils.GetYesterday(date, "2006-01-02")
	}

	// If no entry found in the last @maxDays days return error
	err := errors.New("No database entry found in the last " + strconv.FormatInt(int64(maxDays), 10) + "days.")
	return "", err
}
