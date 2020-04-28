package ratederivatives

import (
	"errors"
	"fmt"
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
	date = date.AddDate(0, 0, 1)
	// Remark: Caution with the check for holidays - formats have to be the same.
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
// 			saturday or sunday) in the respective period. Each day is in the format yyyy-mm-dd.
// @dateInit, @dateFinal determine the period of the loan. Both are given as yyyy-mm-dd.
func CompoundedRate(rate []float64, dateInit, dateFinal time.Time, holidays []time.Time, daysPerYear int) (float64, error) {

	if utils.AfterDay(dateInit, dateFinal) {
		log.Error("The final date cannot be before the initial date.")
		return float64(0), errors.New("date error")
	}
	calendarDays, _ := utils.CountDays(dateInit, dateFinal, false)
	businessDays, _ := utils.CountDays(dateInit, dateFinal, true)
	businessDays -= len(holidays)

	fmt.Println("No of business days: ", businessDays)
	fmt.Println("No of calendar days: ", calendarDays)

	// Check consistency of dates, holidays and rates
	if len(rate) != businessDays {
		log.Error("The given number of rate values and business days is not consistent.")
		return float64(0), errors.New("date error")
	}
	fmt.Println(rate)
	// Iterate through business days
	prod := float64(1)
	for i := 0; i < businessDays; i++ {
		n, _ := RateFactor(dateInit, holidays)
		if utils.SameDays(dateInit, dateFinal) {
			// If the last day is before a holiday/weekend, only count it once.
			n = 1
		}
		fmt.Println(n)
		factor := 1 + (rate[i]/100)*float64(n)/float64(daysPerYear)
		prod *= factor
		dateInit = dateInit.AddDate(0, 0, n)
	}
	// return (prod - float64(1)) * float64(daysPerYear) / float64(calendarDays), nil
	return math.Round(prod*1e8) * 1e-8, nil
}
