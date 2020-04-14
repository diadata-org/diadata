package utils

import (
	"errors"
	"time"

	log "github.com/sirupsen/logrus"
)

// CheckWeekDay returns true if @date is not weekend and false otherwise.
func CheckWeekDay(date time.Time) bool {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false
	}
	return true
}

// ContainsDay returns true if day @e is contained in slice @s, independent of the daytime.
// As a consequence, be cautious when comparing times in different timezones.
func ContainsDay(s []time.Time, date time.Time) bool {
	for _, a := range s {
		if CompareDays(a, date) {
			return true
		}
	}
	return false
}

// CompareDays returns true if @date1 is the same date as @date2, independent of the daytime.
func CompareDays(date1, date2 time.Time) bool {
	if date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day() {
		return true
	}
	return false
}

// CountDays returns the number of days between
// @dateInit and @dateFinal, both given as converted from a string in the format yyyy-mm-dd, excluding the last day.
// @bool If true only business days are counted.
func CountDays(dateInit, dateFinal time.Time, business bool) (days int, err error) {

	days = 0
	if dateInit.After(dateFinal) {
		log.Error("The final date cannot be smaller than the initial date.")
		err = errors.New("date error")
		return
	}

	for {
		if CompareDays(dateInit, dateFinal) {
			return days, nil
		}
		if business {
			if CheckWeekDay(dateInit) {
				days++
			}
			dateInit = dateInit.Add(time.Hour * 24)
		} else {
			days++
			dateInit = dateInit.Add(time.Hour * 24)
		}
	}
}

// GetYesterday returns the day before @date in the world of strings, formatted as @layout
func GetYesterday(date, layout string) string {
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("Error: %v on date format %s\n", err, date)
	}
	yesterday := dateTime.AddDate(0, 0, -1)
	return yesterday.Format(layout)
}
