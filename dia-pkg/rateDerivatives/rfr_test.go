package ratederivatives

import (
	"math"
	"testing"
	"time"
)

func TestRateFactor(t *testing.T) {
	date1, _ := time.Parse("2006-01-02", "2020-04-01")
	date2, _ := time.Parse("2006-01-02", "2020-04-02")
	date3, _ := time.Parse("2006-01-02", "2020-04-03")
	date4, _ := time.Parse("2006-01-02", "2020-04-06")
	tables := []struct {
		date       time.Time
		holidays   []time.Time
		ratefactor int
		err        error
	}{
		{date1, []time.Time{}, 1, nil},
		{date1, []time.Time{date2}, 2, nil},
		{date1, []time.Time{date2, date3}, 5, nil},
		{date1, []time.Time{date2, date4}, 2, nil},
	}
	for _, table := range tables {
		value, err := RateFactor(table.date, table.holidays)
		if value != table.ratefactor {
			t.Errorf("Rate factor is %v but should be %v.", value, table.ratefactor)
		}
		if err != nil {
			t.Errorf("Error should be %v but is %v", table.err, err)
		}
	}
}

func TestCompoundedRate(t *testing.T) {
	tol := 10e-8
	rates1 := []float64{0, 0, 0}
	rates2 := []float64{1.1, 1.2, 1.5}
	date1, _ := time.Parse("2006-01-02 15:04:05", "2020-04-01 14:22:55")
	date2, _ := time.Parse("2006-01-02", "2020-04-02")
	date3, _ := time.Parse("2006-01-02", "2020-04-03")
	date4, _ := time.Parse("2006-01-02", "2020-04-06")
	daysPerYear := 365

	tables := []struct {
		rates       []float64
		dateInit    time.Time
		dateFinal   time.Time
		holidays    []time.Time
		daysPerYear int
		rounding    int
		cumRate     float64
		err         error
	}{
		{rates1[:2], date1, date3, []time.Time{}, daysPerYear, 0, 0, nil},
		{rates2[:2], date1, date3, []time.Time{}, daysPerYear, 0, 1.151808219, nil},
		{rates2, date1, date4, []time.Time{}, daysPerYear, 0, 1.366403438, nil},
		{rates2[:2], date1, date4, []time.Time{date2}, daysPerYear, 0, 1.164339726, nil},
		{rates2[0:1], date1, date2, []time.Time{}, daysPerYear, 0, rates2[0], nil},
	}
	for _, table := range tables {
		value, err := CompoundedRate(table.rates, table.dateInit, table.dateFinal, table.holidays, table.daysPerYear, table.rounding)
		if math.Abs(value-table.cumRate) > tol {
			t.Errorf("Rate factor is %v but should be %v.", value, table.cumRate)
		}
		if err != nil {
			t.Errorf("Error should be %v but is %v", table.err, err)
		}
	}
}
