package ratederivatives

import (
	"errors"
	"math"
	"time"

	"github.com/diadata-org/diadata/pkg/utils"
	log "github.com/sirupsen/logrus"
)

// RateFactor returns the integer a rate is multiplied by in the computation
// of (compounded) RFRs.
func RateFactor(date time.Time, holidays []time.Time) (int, error) {
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
// @dateInit, @dateFinal determine the period of the loan.
// @holidays is a slice of strings where each entry corresponds to a special holiday (i.e. not a
// 			saturday or sunday) in the respective period.
// @daysPerYear determines the total number of days per business year.
// @rounding is a float of type 1e-n which rounds the result to n digits. If @rounding == 0 no rounding
func CompoundedRate(rate []float64, dateInit, dateFinal time.Time, holidays []time.Time, daysPerYear int, rounding float64) (float64, error) {

	if !utils.CheckWeekDay(dateFinal) || utils.ContainsDay(holidays, dateFinal) {
		log.Error("The final date cannot be a holiday or weekend")
		return float64(0), errors.New("The final date cannot be a holiday or weekend")
	}
	if utils.AfterDay(dateInit, dateFinal) {
		log.Error("The final date cannot be before the initial date.")
		return float64(0), errors.New("The final date cannot be before the initial date")
	}
	if daysPerYear == 0 {
		log.Error("Days per year must be a positive integer.")
		return float64(0), errors.New("Days per year must be a positive integer")
	}

	// Check consistency of dates, holidays and rates
	NumBusinessDays, _ := utils.CountDays(dateInit, dateFinal, true)
	NumBusinessDays -= len(holidays)
	if !utils.CheckWeekDay(dateInit) || utils.ContainsDay(holidays, dateInit) {
		// When first day is holiday or weekend, there has to be an additional rate for the
		// previous working day which does not fall into the loan period.
		if len(rate) != NumBusinessDays+1 {
			log.Error("The given number of rate values and business days is not consistent.")
			return float64(0), errors.New("date error")
		}
	} else {
		if len(rate) != NumBusinessDays {
			log.Error("The given number of rate values and business days is not consistent.")
			return float64(0), errors.New("date error")
		}
	}

	// Iterate through business days to compute the compounded rate
	prod := float64(1)
	summe := 0
	for i := 0; i < len(rate); i++ {
		n, _ := RateFactor(dateInit, holidays)
		summe += n
		factor := 1 + (rate[i]/100)*float64(n)/float64(daysPerYear)
		prod *= factor
		dateInit = dateInit.AddDate(0, 0, n)
	}

	// In case of the SOFR Index, results are rounded to eight decimals
	if rounding != 0 {
		result := math.Round(prod/rounding) * rounding
		return result, nil
	}
	return prod, nil
}
