package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
	"strconv"
	"time"

	ratederivatives "github.com/diadata-org/diadata/internal/pkg/rateDerivatives"
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

	// Collect all possible keys in time range
	keys := []string{}
	auxDate := dateInit
	for dateFinal >= auxDate {
		keys = append(keys, "dia_quotation_"+symbol+"_"+auxDate+" 00:00:00 +0000 UTC")
		auxDate = utils.GetTomorrow(auxDate, "2006-01-02")
	}
	// Retrieve corresponding values from database
	result := db.redisClient.MGet(keys...).Val()
	allValues := []*InterestRate{}
	for _, val := range result {
		if val != nil {
			ir := &InterestRate{}
			err := json.Unmarshal([]byte(fmt.Sprint(val)), ir)
			if err != nil {
				log.Error("error parsing json")
				return []*InterestRate{}, err
			}
			allValues = append(allValues, ir)
		}
	}

	// Sort entries with respect to effective date
	sort.Slice(allValues, func(i, j int) bool {
		return (allValues[i].EffectiveDate).Before(allValues[j].EffectiveDate)
	})
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
		// Number of decimals
		decimals := 0
		switch symbol {
		case "SONIA":
			decimals = 4
		case "SOFR":
			decimals = 2
		case "SAFR":
			decimals = 8
		case "SOFR30":
			decimals = 5
		case "SOFR90":
			decimals = 5
		case "SOFR180":
			decimals = 5
		case "ESTER":
			decimals = 3
		default:
			decimals = 8
		}
		// Fill meta type
		newEntry := InterestRateMeta{symbol, newdate, decimals, issuer}
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
func (db *DB) GetCompoundedRate(symbol string, dateInit, date time.Time, daysPerYear int, rounding float64) (*InterestRate, error) {

	// Get first publication date for the rate with @symbol in order to check feasibility of dateInit
	firstPublication, err := db.GetFirstDate(symbol)
	if err != nil {
		return &InterestRate{}, err
	}
	if utils.AfterDay(firstPublication, dateInit) {
		log.Error("dateInit cannot be earlier than first publication date.")
		err = errors.New("dateInit cannot be earlier than first publication date")
		return &InterestRate{}, err
	}

	ratesAPI, err := db.GetInterestRateRange(symbol, dateInit.Format("2006-01-02"), date.Format("2006-01-02"))
	if err != nil {
		return &InterestRate{}, err
	}
	if len(ratesAPI) == 0 {
		err = errors.New("no rate information for this period")
		return &InterestRate{}, err
	}

	// Determine holidays through missing database entries
	existDates := []time.Time{}
	for _, entry := range ratesAPI {
		existDates = append(existDates, (*entry).EffectiveDate)
	}
	holidays := utils.GetHolidays(existDates, dateInit, date)

	// Sort ratesApi (type []*InterestRates) in increasing order according to date
	// and remove the data for the final date, as only past values are compounded.
	sort.Slice(ratesAPI, func(i, j int) bool {
		return (ratesAPI[i].EffectiveDate).Before(ratesAPI[j].EffectiveDate)
	})
	ratesAPI = ratesAPI[:len(ratesAPI)-1]

	// Remove holiday if after @date
	for i, day := range holidays {
		if day.After(date) {
			fmt.Println("removed holiday at the end")
			holidays = holidays[:i]
			break
		}
	}

	// Extract rates' values
	rates := []float64{}
	for i := range ratesAPI {
		rates = append(rates, ratesAPI[i].Value)
	}

	// Check, whether first day is a holiday or weekend. If so, prepend rate of
	// preceding business day (outside the considered time range!).
	if utils.ContainsDay(holidays, dateInit) || !utils.CheckWeekDay(dateInit) {
		firstRate, err := db.GetInterestRate(symbol, dateInit.Format("2006-01-02"))
		if err != nil {
			return &InterestRate{}, err
		}
		rates = append([]float64{firstRate.Value}, rates...)
	}

	// Get compounded rate
	compRate, err := ratedevs.CompoundedRate(rates, dateInit, date, holidays, daysPerYear, rounding)
	if err != nil {
		return &InterestRate{}, err
	}

	// Fill InterestRate type for return
	ir := &InterestRate{}
	ir.Symbol = symbol + "_compounded_by_DIA"
	ir.Value = compRate
	ir.EffectiveDate = date
	ir.Source = ratesAPI[0].Source

	return ir, nil
}

// GetCompoundedIndex returns the compounded index over the maximal period of existence of @symbol
func (db *DB) GetCompoundedIndex(symbol string, date time.Time, daysPerYear int, rounding float64) (*InterestRate, error) {
	// Get initial date for the rate with @symbol
	dateInit, err := db.GetFirstDate(symbol)
	if err != nil {
		return &InterestRate{}, err
	}
	return db.GetCompoundedRate(symbol, dateInit, date, daysPerYear, rounding)
}

// GetCompoundedIndexRange returns a range of compounded index values of @symbol
func (db *DB) GetCompoundedIndexRange(symbol string, dateInit, dateFinal time.Time, daysPerYear int, rounding float64) (values []*InterestRate, err error) {

	for utils.AfterDay(dateFinal, dateInit) {
		val, err := db.GetCompoundedIndex(symbol, dateInit, daysPerYear, rounding)
		if err != nil {
			dateInit = dateInit.AddDate(0, 0, 1)
		} else {
			values = append(values, val)
			dateInit = dateInit.AddDate(0, 0, 1)
		}
	}
	return
}

// GetCompoundedAvg returns the compounded average of the index @symbol over rolling @calDays calendar days.
func (db *DB) GetCompoundedAvg(symbol string, date time.Time, calDays, daysPerYear int, rounding float64) (*InterestRate, error) {

	dateInit := date.AddDate(0, 0, -calDays)

	index, err := db.GetCompoundedRate(symbol, dateInit, date, daysPerYear, rounding)
	if err != nil {
		return &InterestRate{}, err
	}

	// Fill return struct
	compAvg := &InterestRate{}
	compAvg.Symbol = symbol + strconv.Itoa(calDays) + "_compounded_by_DIA"
	compAvg.Value = 100 * (index.Value - 1) * float64(daysPerYear) / float64(calDays)
	compAvg.EffectiveDate = date
	compAvg.Source = index.Source

	return compAvg, nil
}

// --------------------------------------------------------------------------------------------
// Computation of compounded average range as done by FED and BOE, i.e. neglecting higher order
// terms accounting for holidays and weekends
// --------------------------------------------------------------------------------------------

// WeightedRates returns a map which maps a rate to each business day in the time period.
// Rates are weighted by the rate factor. intRates must be sorted by date in increasing order.
func WeightedRates(intRates []*InterestRate, dateInit, dateFinal time.Time, holidays []time.Time, startIndex int) (map[time.Time]float64, int) {

	rateMap := make(map[time.Time]float64)

	// Adjust rate if dateInit is not a business day
	initialFactor := 0
	auxDate := dateInit
	for !utils.CheckWeekDay(auxDate) || utils.ContainsDay(holidays, auxDate) {
		initialFactor++
		auxDate = auxDate.AddDate(0, 0, 1)
	}

	// Get index for first date inside global range
	for utils.AfterDay(dateInit, intRates[startIndex].EffectiveDate) {
		startIndex++
	}
	// Return index as cursor inside of entire range requested in API call
	index := startIndex

	// If first dateInit is non-business day, get previous rate
	if !utils.CheckWeekDay(dateInit) || utils.ContainsDay(holidays, dateInit) {
		startIndex = startIndex - 1
		rateMap[dateInit] = float64(initialFactor) * intRates[startIndex].Value
		startIndex++
	}

	// Compute compounded rate for period [dateInit, dateFinal]
	for utils.AfterDay(dateFinal, intRates[startIndex].EffectiveDate) && startIndex < len(intRates)-1 {
		ratefactor, _ := ratederivatives.RateFactor(intRates[startIndex].EffectiveDate, holidays)
		rateMap[intRates[startIndex].EffectiveDate] = float64(ratefactor) * intRates[startIndex].Value
		startIndex++
	}

	return rateMap, index
}

// GetCompoundedAvgRange returns the compounded average of the index @symbol over rolling @calDays calendar days.
func (db *DB) GetCompoundedAvgRange(symbol string, dateInit, dateFinal time.Time, calDays, daysPerYear int, rounding float64) (values []*InterestRate, err error) {

	dateStart := dateInit.AddDate(0, 0, -calDays)

	// Get first publication date for the rate with @symbol in order to check feasibility of dateInit
	firstPublication, err := db.GetFirstDate(symbol)
	if err != nil {
		return []*InterestRate{}, err
	}
	if utils.AfterDay(firstPublication, dateStart) {
		log.Error("dateStart cannot be earlier than first publication date.")
		err = errors.New("dateStart cannot be earlier than first publication date")
		return []*InterestRate{}, err
	}

	// Get rate data from database
	ratesAPI, err := db.GetInterestRateRange(symbol, dateStart.Format("2006-01-02"), dateFinal.Format("2006-01-02"))
	if err != nil {
		return []*InterestRate{}, err
	}
	if len(ratesAPI) == 0 {
		err = errors.New("no rate information for this period")
		return []*InterestRate{}, err
	}

	// Check, whether first day is a holiday or weekend. If so, prepend rate of
	// preceding business day (outside the considered time range!).
	// Determine holidays through missing database entries
	existDates := []time.Time{}
	for _, entry := range ratesAPI {
		existDates = append(existDates, (*entry).EffectiveDate)
	}
	holidays := utils.GetHolidays(existDates, dateStart, dateFinal)
	if utils.ContainsDay(holidays, dateStart) || !utils.CheckWeekDay(dateStart) {
		firstRate, err := db.GetInterestRate(symbol, dateStart.Format("2006-01-02"))
		if err != nil {
			return []*InterestRate{}, err
		}
		ratesAPI = append([]*InterestRate{firstRate}, ratesAPI...)
	}

	// Consider last business day if last given day is holiday or weekend
	for utils.ContainsDay(holidays, dateFinal) || !utils.CheckWeekDay(dateFinal) {
		dateFinal = dateFinal.AddDate(0, 0, -1)
	}

	// Sort ratesApi (type []*InterestRates) in increasing order according to date
	// and remove the data for the final date, as only past values are compounded.
	sort.Slice(ratesAPI, func(i, j int) bool {
		return (ratesAPI[i].EffectiveDate).Before(ratesAPI[j].EffectiveDate)
	})
	ratesAPI = ratesAPI[:len(ratesAPI)-1]

	// Iterate through interest periods of length @calDays
	cursor := 0
	for utils.AfterDay(dateFinal, dateInit) {

		// Get starting date of compounding period
		dateStart := dateInit.AddDate(0, 0, -calDays)

		// get a weighted rate for each business day in period of interest
		mapRates, index := WeightedRates(ratesAPI, dateStart, dateInit, holidays, cursor)
		cursor = index

		auxDate := dateStart
		ratesPeriod := []float64{}
		for utils.AfterDay(dateInit, auxDate) {
			val, ok := mapRates[auxDate]
			if ok {
				ratesPeriod = append(ratesPeriod, val)
				auxDate = auxDate.AddDate(0, 0, 1)
			} else {
				auxDate = auxDate.AddDate(0, 0, 1)
			}
		}

		compRate, err := ratedevs.CompoundedRateSimple(ratesPeriod, dateStart, dateInit, daysPerYear, rounding)

		if err != nil || utils.ContainsDay(holidays, dateInit) {
			dateInit = dateInit.AddDate(0, 0, 1)
		} else {

			// Fill return struct
			compAvg := &InterestRate{}
			compAvg.Symbol = symbol + strconv.Itoa(calDays) + "_compounded_by_DIA"
			compAvg.Value = 100 * (compRate - 1) * float64(daysPerYear) / float64(calDays)
			compAvg.EffectiveDate = dateInit
			compAvg.Source = ratesAPI[0].Source

			// Append data and increment initial date
			values = append(values, compAvg)
			dateInit = dateInit.AddDate(0, 0, 1)
		}

	}

	return values, nil
}

// ---------------------------------------------------------------------------------------------
// Computation of compounded averages in conservative way, i.e. including higher order terms
// ---------------------------------------------------------------------------------------------

// StraightRates returns a map which maps a rate to each day in the time period. This includes (artificial)
// rate values for non-business days. intRates must be sorted by date in increasing order.
func StraightRates(intRates []*InterestRate) map[time.Time]float64 {

	finalDay := intRates[len(intRates)-1].EffectiveDate
	count := 0
	day := intRates[count].EffectiveDate

	rateMap := make(map[time.Time]float64)
	rateMap[day] = intRates[count].Value
	day = day.AddDate(0, 0, 1)
	count++
	for utils.AfterDay(finalDay, day) {
		if utils.SameDays(day, intRates[count].EffectiveDate) {
			// For business day assign rate and increment day
			rateMap[day] = intRates[count].Value
			day = day.AddDate(0, 0, 1)
			count++
		} else {
			// holiday or weekend gets the previous rate
			rateMap[day] = intRates[count-1].Value
			day = day.AddDate(0, 0, 1)
		}
	}
	return rateMap
}

// GetCompoundedAvgDIARange returns the compounded average DIA index of @symbol over rolling @calDays calendar days.
func (db *DB) GetCompoundedAvgDIARange(symbol string, dateInit, dateFinal time.Time, calDays, daysPerYear int, rounding float64) (values []*InterestRate, err error) {

	dateStart := dateInit.AddDate(0, 0, -calDays)

	// Get first publication date for the rate with @symbol in order to check feasibility of dateInit
	firstPublication, err := db.GetFirstDate(symbol)
	if err != nil {
		return []*InterestRate{}, err
	}
	if utils.AfterDay(firstPublication, dateStart) {
		log.Error("dateStart cannot be earlier than first publication date.")
		err = errors.New("dateStart cannot be earlier than first publication date")
		return []*InterestRate{}, err
	}

	// Get rate data from database
	ratesAPI, err := db.GetInterestRateRange(symbol, dateStart.Format("2006-01-02"), dateFinal.Format("2006-01-02"))
	if err != nil {
		return []*InterestRate{}, err
	}
	if len(ratesAPI) == 0 {
		err = errors.New("no rate information for this period")
		return []*InterestRate{}, err
	}

	// Check, whether first day is a holiday or weekend. If so, prepend rate of
	// preceding business day (outside the considered time range!).
	// Determine holidays through missing database entries
	existDates := []time.Time{}
	for _, entry := range ratesAPI {
		existDates = append(existDates, (*entry).EffectiveDate)
	}
	holidays := utils.GetHolidays(existDates, dateStart, dateFinal)
	if utils.ContainsDay(holidays, dateStart) || !utils.CheckWeekDay(dateStart) {
		firstRate, err := db.GetInterestRate(symbol, dateStart.Format("2006-01-02"))
		if err != nil {
			return []*InterestRate{}, err
		}
		ratesAPI = append([]*InterestRate{firstRate}, ratesAPI...)
	}

	// Consider last business day if last given day is holiday or weekend
	for utils.ContainsDay(holidays, dateFinal) || !utils.CheckWeekDay(dateFinal) {
		dateFinal = dateFinal.AddDate(0, 0, -1)
	}

	// Sort ratesApi (type []*InterestRates) in increasing order according to date
	// and remove the data for the final date, as only past values are compounded.
	sort.Slice(ratesAPI, func(i, j int) bool {
		return (ratesAPI[i].EffectiveDate).Before(ratesAPI[j].EffectiveDate)
	})
	ratesAPI = ratesAPI[:len(ratesAPI)-1]

	// get a rate for each calendar day in period of interest
	mapRates := StraightRates(ratesAPI)

	for i := range ratesAPI {
		fmt.Println("ratesAPI[i]: ", ratesAPI[i])
	}
	fmt.Println("mapRates: ", mapRates)

	// Iterate through interest period
	count := 0
	for utils.AfterDay(dateFinal, dateInit) {
		// Get compounded rate
		dateStart := dateInit.AddDate(0, 0, -calDays)
		auxDate := dateStart
		ratesPeriod := []float64{}
		for i := count; i < count+calDays; i++ {
			ratesPeriod = append(ratesPeriod, mapRates[auxDate])
			auxDate = auxDate.AddDate(0, 0, 1)
		}
		compRate, err := ratedevs.CompoundedRateSimple(ratesPeriod, dateStart, dateInit, daysPerYear, rounding)

		if err != nil {
			dateInit = dateInit.AddDate(0, 0, 1)
		} else {

			// Fill return struct
			compAvg := &InterestRate{}
			compAvg.Symbol = symbol + strconv.Itoa(calDays) + "_compounded_by_DIA"
			compAvg.Value = 100 * (compRate - 1) * float64(daysPerYear) / float64(calDays)
			compAvg.EffectiveDate = dateInit
			compAvg.Source = ratesAPI[0].Source

			// Append data and increment initial date
			values = append(values, compAvg)
			dateInit = dateInit.AddDate(0, 0, 1)
			count++
		}

	}
	return values, nil
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
