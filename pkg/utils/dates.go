package utils

import (
	"errors"
	"time"

	cal "github.com/rickar/cal"
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
		if SameDays(a, date) {
			return true
		}
	}
	return false
}

// SameDays returns true if @date1 is the same date as @date2, independent of the daytime.
func SameDays(date1, date2 time.Time) bool {
	if date1.Year() == date2.Year() && date1.Month() == date2.Month() && date1.Day() == date2.Day() {
		return true
	}
	return false
}

// AfterDay returns true if date1 is a date after date2, irrespective of the daytime.
// The go method "After" respects daytime.
func AfterDay(date1, date2 time.Time) bool {
	date1Str := date1.Format("2006-01-02")
	date2Str := date2.Format("2006-01-02")
	if date1Str > date2Str {
		return true
	}
	return false
}

// CountDays returns the number of days between
// @dateInit and @dateFinal, both given as converted from a string in the format yyyy-mm-dd, excluding the last day.
// @bool If true only business days are counted.
func CountDays(dateInit, dateFinal time.Time, business bool) (days int, err error) {

	if SameDays(dateInit, dateFinal) {
		return 0, nil
	}
	days = 0
	if dateInit.After(dateFinal) {
		log.Error("The final date cannot be smaller than the initial date.")
		err = errors.New("date error")
		return
	}

	for {
		if SameDays(dateInit, dateFinal) {
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

// GetHolidays returns "holidays" as non-weekend complement of given days @workdays
func GetHolidays(workdays []time.Time, dateInit, date time.Time) []time.Time {

	if AfterDay(dateInit, date) {
		log.Error("The initial date must not be after the final date.")
		return []time.Time{}
	}
	auxDate := dateInit
	holidays := []time.Time{}
	for !SameDays(auxDate, date.AddDate(0, 0, 1)) {
		if CheckWeekDay(auxDate) && !ContainsDay(workdays, auxDate) {
			holidays = append(holidays, auxDate)
			auxDate = auxDate.AddDate(0, 0, 1)
		} else {
			auxDate = auxDate.AddDate(0, 0, 1)
		}
	}
	return holidays
}

// GetHolidaysZone returns a slice of dates which are holidays in @zone between @dateInit and @dateFinal
func GetHolidaysZone(zone string, dateInit, dateFinal time.Time) (holidays []time.Time, err error) {
	c := LocalCal(zone)
	if SameDays(dateInit, dateFinal) {
		if c.IsHoliday(dateInit) {
			holidays = []time.Time{dateInit}
			return
		}
		return []time.Time{}, nil
	}
	if dateInit.After(dateFinal) {
		log.Error("The final date cannot be smaller than the initial date.")
		err = errors.New("date error")
		return
	}
	for dateFinal.After(dateInit) {
		if c.IsHoliday(dateInit) {
			holidays = append(holidays, dateInit)
		}
		dateInit = dateInit.Add(time.Hour * 24)
	}
	return
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

// GetTomorrow returns the day before @date in the world of strings, formatted as @layout
func GetTomorrow(date, layout string) string {
	dateTime, err := time.Parse(layout, date)
	if err != nil {
		log.Printf("Error: %v on date format %s\n", err, date)
	}
	tomorrow := dateTime.AddDate(0, 0, 1)
	return tomorrow.Format(layout)
}

// LocalCal returns a (pointer to a) calendar in the country/zone given by the string @zone.
func LocalCal(zone string) *cal.Calendar {
	c := cal.NewCalendar()
	switch zone {
	case "Australia":
		cal.AddAustralianHolidays(c)
	case "Austria":
		cal.AddAustrianHolidays(c)
	case "Belgium":
		cal.AddBelgiumHolidays(c)
	case "Canada":
		cal.AddCanadianHolidays(c)
	case "Czech":
		cal.AddCzechHolidays(c)
	case "Danmark":
		cal.AddDanishHolidays(c)
	case "Ecb":
		cal.AddEcbHolidays(c)
	case "France":
		cal.AddFranceHolidays(c)
	case "Germany":
		cal.AddGermanHolidays(c)
	case "Italy":
		cal.AddItalianHolidays(c)
	case "Netherlands":
		cal.AddDutchHolidays(c)
	case "NewZealand":
		cal.AddNewZealandHoliday(c)
	case "Norway":
		cal.AddNorwegianHolidays(c)
	case "Poland":
		cal.AddPolandHolidays(c)
	case "Slowakia":
		cal.AddSlovakHolidays(c)
	case "SouthAfrica":
		cal.AddSouthAfricanHolidays(c)
	case "Spain":
		cal.AddSpainHolidays(c)
	case "Sweden":
		cal.AddSwedishHolidays(c)
	case "UK":
		cal.AddBritishHolidays(c)
	case "US":
		cal.AddUsHolidays(c)
	case "Ukrania":
		cal.AddUkraineHolidays(c)
	default:
		cal.AddEcbHolidays(c)
	}

	return c
}
