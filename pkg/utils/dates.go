package utils

import (
	"errors"
	"strconv"
	"time"
)

// StrToUnixtime converts a string corresponding to an int to Unix time
func StrToUnixtime(s string) (t time.Time, err error) {
	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return
	}
	t = time.Unix(i, 0)
	return
}

// CheckWeekDay returns true if @date is not weekend and false otherwise.
func CheckWeekDay(date time.Time) bool {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false
	}
	return true
}

// ContainsDay returns true if day @date is contained in slice @s, independent of the daytime.
// As a consequence, be cautious when comparing days in different timezones.
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
	return date1Str > date2Str
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
func GetHolidays(workdays []time.Time, dateInit, dateFinal time.Time) []time.Time {

	if AfterDay(dateInit, dateFinal) {
		log.Error("The initial date must not be after the final date.")
		return []time.Time{}
	}
	auxDate := dateInit
	holidays := []time.Time{}
	for !SameDays(auxDate, dateFinal.AddDate(0, 0, 1)) {
		if CheckWeekDay(auxDate) && !ContainsDay(workdays, auxDate) {
			holidays = append(holidays, auxDate)
			auxDate = auxDate.AddDate(0, 0, 1)
		} else {
			auxDate = auxDate.AddDate(0, 0, 1)
		}
	}
	return holidays
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

// MakeTimeRanges returns @numRanges start- and endtimes partitioning [@timeInit, @timeFinal] in intervals of identical size.
func MakeTimeRanges(timeInit, timeFinal time.Time, numRanges int) (starttimes, endtimes []time.Time) {
	a := timeInit
	b := timeFinal
	totalSize := b.Sub(a)
	sizeRange := totalSize / time.Duration(numRanges)
	starttime := timeInit
	for k := 0; k < numRanges; k++ {
		starttimes = append(starttimes, starttime)
		endtimes = append(endtimes, starttime.Add(sizeRange))
		starttime = starttime.Add(sizeRange)
	}
	return
}

// MakeTimerange parses Unix timestamps given as strings.
// In case one of the two is empty, it returns a time-range based on @timeRange.
// Default is a 24h window ending now.
func MakeTimerange(starttimeString string, endtimeString string, timeRange time.Duration) (starttime time.Time, endtime time.Time, err error) {

	var (
		starttimeInt int64
		endtimeInt   int64
	)

	if starttimeString == "" && endtimeString == "" {
		endtime = time.Now()
		starttime = endtime.Add(-timeRange)
	} else if starttimeString == "" && endtimeString != "" {
		// zero time if not given
		endtimeInt, err = strconv.ParseInt(endtimeString, 10, 64)
		if err != nil {
			return
		}
		endtime = time.Unix(endtimeInt, 0)
		starttime = endtime.Add(-timeRange)
	} else if starttimeString != "" && endtimeString == "" {
		starttimeInt, err = strconv.ParseInt(starttimeString, 10, 64)
		if err != nil {
			return
		}
		starttime = time.Unix(starttimeInt, 0)
		endtime = starttime.Add(timeRange)
	} else {
		starttimeInt, err = strconv.ParseInt(starttimeString, 10, 64)
		if err != nil {
			return
		}
		starttime = time.Unix(starttimeInt, 0)
		endtimeInt, err = strconv.ParseInt(endtimeString, 10, 64)
		if err != nil {
			return
		}
		endtime = time.Unix(endtimeInt, 0)
	}
	return starttime, endtime, nil
}

// ValidTimeRange returns true if the interval [@starttime, @endtime] is at most @maxDuration.
func ValidTimeRange(starttime time.Time, endtime time.Time, maxDuration time.Duration) (ok bool) {
	if endtime.Sub(starttime) <= maxDuration {
		ok = true
	}
	return
}
