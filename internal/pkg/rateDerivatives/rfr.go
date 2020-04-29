package ratederivatives

import (
	"errors"
	"math"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// RateFactor returns the integer a rate is multiplied by in the computation
// of (compunded) RFRs.
func RateFactor(date time.Time, holidays []time.Time) (int, error) {

	if utils.CheckWeekDay(date) == false || utils.ContainsDay(holidays, date) {
		log.Error("The rate factor can only be computed for business days")
		err := errors.New("date error")
		return 0, err
	}
	rate := 1
	// Increment @rate for each holiday/weekend day after date
	date = date.AddDate(0, 0, 1)
	for !utils.CheckWeekDay(date) || utils.ContainsDay(holidays, date) {
		rate++
		date = date.AddDate(0, 0, 1)
	}
	return rate, nil
}

// CompoundedRate returns the compounded index for the rate values given by the slice @rate.
// @rate is a slice with daily rates for all business days in the respective period.
// @daysPerYear determines the total number of days per business year.
// @holidays is a slice of strings where each entry corresponds to a special holiday (i.e. not a
// 			saturday or sunday) in the respective period.
// @dateInit, @dateFinal determine the period of the loan.
func CompoundedRate(rate []float64, dateInit, dateFinal time.Time, holidays []time.Time, daysPerYear int) (float64, error) {

	if utils.AfterDay(dateInit, dateFinal) {
		log.Error("The final date cannot be before the initial date.")
		return float64(0), errors.New("date error")
	}

	// Check consistency of dates, holidays and rates
	NumBusinessDays, _ := utils.CountDays(dateInit, dateFinal, true)
	NumBusinessDays -= len(holidays)
	if len(rate) != NumBusinessDays {
		log.Error("The given number of rate values and business days is not consistent.")
		return float64(0), errors.New("date error")
	}

	// Iterate through business days to compute the compounded rate
	prod := float64(1)
	for i := 0; i < NumBusinessDays; i++ {
		n, _ := RateFactor(dateInit, holidays)
		factor := 1 + (rate[i]/100)*float64(n)/float64(daysPerYear)
		prod *= factor
		dateInit = dateInit.AddDate(0, 0, n)
	}

	// In case of the SOFR Index, results are rounded to eight decimals
	return math.Round(prod*1e8) * 1e-8, nil
}
