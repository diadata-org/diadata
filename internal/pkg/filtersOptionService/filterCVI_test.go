package filters

import (
	"github.com/nazariyv/diadata-scrapers/pkg/scrapers"
	"testing"
	"time"
)

func TestMinutesInYear(t *testing.T) {
	const minutes float64 = 525_600
	const leapMinutes float64 = 527_040

	// 2020 - leap year
	actual, err := MinutesInYear(2020)
	if err != nil {
		t.Errorf("no error should be thrown")
	}
	if leapMinutes != actual {
		t.Errorf("expected %v minutes, got %v", leapMinutes, actual)
	}

	// 0
	actual, err = MinutesInYear(0)
	if err != nil {
		t.Errorf("no error should be thrown")
	}
	if leapMinutes != actual {
		t.Errorf("expected %v minutes, got %v", minutes, actual)
	}

	// 2019 - non-leap
	actual, err = MinutesInYear(1)
	if err != nil {
		t.Errorf("no error should be thrown")
	}
	if minutes != actual {
		t.Errorf("expected %v minutes, got %v", minutes, actual)
	}

	// -1 (invalid)
	actual, err = MinutesInYear(-1)
	if err == nil {
		t.Errorf("invalid year, should throw error")
	}
}

func TestMinutesBetweenTwoDays(t *testing.T) {
	const minutesInDay float64 = 60 * 24

	l, _ := time.LoadLocation("CET")
	t1 := time.Date(2019, 3, 21, 11, 10, 10, 0, l)
	t2 := time.Date(2019, 3, 23, 5, 5, 0, 0, l)
	minutesBetween, err := MinutesBetweenTwoDays(t1, t2)
	if err != nil {
		t.Errorf("problem computing minutes between two days, %v", err)
	}
	if minutesBetween != minutesInDay {
		t.Errorf("expected one day of minutes: %v, actual: %v", minutesInDay, minutesBetween)
	}

	// when no days in between
	t1 = time.Date(2019, 3, 21, 11, 5, 3, 0, l)
	t2 = time.Date(2019, 3, 22, 11, 3, 5, 2, l)
	minutesBetween, err = MinutesBetweenTwoDays(t1, t2)
	if err != nil {
		t.Errorf("problem computing minutes between two days, %v", err)
	}
	if minutesBetween != 0 {
		t.Errorf("expected no minutes between the two days, actual: %v", minutesBetween)
	}

	// when same day passed
	t1 = time.Date(2019, 3, 21, 11, 11, 11, 0, l)
	minutesBetween, err = MinutesBetweenTwoDays(t1, t1)
	if err != nil {
		t.Errorf("problem computing minutes between the same day, %v", err)
	}
	if minutesBetween != 0 {
		t.Errorf("expected no minutes between the two days, actual: %v", minutesBetween)
	}

	// when days are not passed in order
	t1 = time.Date(2019, 3, 21, 11, 5, 3, 0, l)
	t2 = time.Date(2019, 3, 18, 23, 1, 1, 0, l)
	minutesBetween, err = MinutesBetweenTwoDays(t1, t2)
	if err != nil {
		t.Errorf("problem computing minutes between two days, %v", err)
	}
	if minutesBetween != 2*minutesInDay {
		t.Errorf("expected two days of minutes: %v, actual: %v", 2*minutesInDay, minutesBetween)
	}
}

func TestMinutesUntilSettlement(t *testing.T) {
	timezone := "CET"
	const expectedRegularMinutes float64 = 60 * 9
	const expectedWeeklyMinutes float64 = 60 * 17.5

	m, err := MinutesUntilSettlement(scrapers.RegularOptionSettlement, timezone)
	if err != nil {
		t.Errorf("problem computing minutes until settlement, %v", err)
	}
	if m != expectedRegularMinutes {
		t.Errorf("expected %v minutes until settlement, actual: %v", expectedRegularMinutes, m)
	}

	m, err = MinutesUntilSettlement(scrapers.WeeklyOptionSettlement, timezone)
	if err != nil {
		t.Errorf("problem computing minutes until settlement, %v", err)
	}
	if m != expectedWeeklyMinutes {
		t.Errorf("expectd %v minutes until settlement, actual: %v", expectedWeeklyMinutes, m)
	}

	// unsuported timezone
	m, err = MinutesUntilSettlement(scrapers.RegularOptionSettlement, "")
	if err == nil {
		t.Errorf("expected an error, timezone is unsupported")
	}
}

func TestMinutesUntilMidnight(t *testing.T) {
	timezone := "CET"
	l, _ := time.LoadLocation(timezone)
	t1 := time.Date(2019, 3, 21, 23, 59, 0, 0, l)
	m, err := processMinutesUntilMidnight(t1, timezone)
	if err != nil {
		t.Errorf("problem processing minutes until midnight: %v", err)
	}
	if m != 1 {
		t.Errorf("expected one minute until midnight, actual: %v", m)
	}

	t1 = time.Date(2019, 3, 21, 23, 59, 30, 0, l)
	m, err = processMinutesUntilMidnight(t1, timezone)
	if err != nil {
		t.Errorf("problem processing minutes until midnight: %v", err)
	}
	if m != 0.5 {
		t.Errorf("expected half a minute until midnight, actual: %v", m)
	}

	t1 = time.Date(2019, 3, 22, 0, 0, 0, 0, l)
	m, err = processMinutesUntilMidnight(t1, timezone)
	if err != nil {
		t.Errorf("problem processing minutes until midnight: %v", err)
	}
	if m != 60*24 {
		t.Errorf("expected full day worth of minutes until midnight, actual: %v", m)
	}
}

func TestEod(t *testing.T) {
	timezone := "CET"
	l, _ := time.LoadLocation(timezone)
	t1 := time.Date(2019, 3, 21, 23, 59, 0, 0, l)
	expectedEod := time.Date(2019, 3, 22, 0, 0, 0, 0, l)

	eod, err := Eod(t1, timezone)
	if err != nil {
		t.Errorf("problem calculating the end of day: %v", err)
	}
	if !eod.Equal(expectedEod) {
		t.Errorf("expected eod: %v, actual: %v", expectedEod, eod)
	}

	t2 := time.Date(2019, 3, 22, 0, 0, 0, 0, l)
	expectedEod = time.Date(2019, 3, 23, 0, 0, 0, 0, l)

	eod, err = Eod(t2, timezone)
	if err != nil {
		t.Errorf("problem calculating the end of day: %v", err)
	}
	if !eod.Equal(expectedEod) {
		t.Errorf("expected eod: %v, actual: %v", expectedEod, eod)
	}
}

func TestBod(t *testing.T) {
	timezone := "CET"
	l, _ := time.LoadLocation(timezone)
	t1 := time.Date(2019, 3, 21, 23, 59, 0, 0, l)
	expectedBod := time.Date(2019, 3, 21, 0, 0, 0, 0, l)

	bod, err := Bod(t1, timezone)
	if err != nil {
		t.Errorf("problem calculating beignning of day: %v", err)
	}
	if !bod.Equal(expectedBod) {
		t.Errorf("expected bod: %v, actual: %v", expectedBod, bod)
	}

	t2 := time.Date(2019, 3, 21, 0, 0, 0, 0, l)
	expectedBod = time.Date(2019, 3, 21, 0, 0, 0, 0, l)

	bod, err = Bod(t2, timezone)
	if err != nil {
		t.Errorf("problem calculating beginning of day: %v", err)
	}
	if !bod.Equal(expectedBod) {
		t.Errorf("expected bod: %v, actual: %v", expectedBod, bod)
	}
}
