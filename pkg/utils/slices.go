package utils

import (
	"errors"
	"math"
	"sort"
	"time"
)

// ArgsortableSlice is a wrapper struct around the sort interface. It allows
// for implemetation of argsort for all sortable types.
type ArgsortableSlice struct {
	sort.Interface
	idx []int
}

type TimeBin struct {
	Starttime time.Time
	Endtime   time.Time
}

func (as ArgsortableSlice) Ind() []int {
	return as.idx
}

// Swap swaps the corresponding indices together with the values.
func (s ArgsortableSlice) Swap(i, j int) {
	s.Interface.Swap(i, j)
	s.idx[i], s.idx[j] = s.idx[j], s.idx[i]
}

// NewFloat64Slice returns the wrapped float64slice that can be argsorted.
func NewFloat64Slice(sf sort.Float64Slice) *ArgsortableSlice {
	s := &ArgsortableSlice{
		Interface: sf,
		idx:       make([]int, sf.Len()),
	}
	for i := range s.idx {
		s.idx[i] = i
	}
	return s
}

// TO DO: Switch to generics for these simple algebraic functions.

// Average returns the average of @samples.
func Average(series []float64) (average float64) {
	length := float64(len(series))
	if length == 0 {
		return
	}
	for _, s := range series {
		average += s

	}
	average /= length
	return
}

func Variance(series []float64) (variance float64) {
	length := float64(len(series))
	if length == 0 {
		return
	}
	if length == 1 {
		return 0
	}

	avg := Average(series)
	for _, item := range series {
		variance += math.Pow(item-avg, float64(2))
	}
	variance /= (length - 1)
	return
}

func StandardDeviation(series []float64) float64 {
	return math.Sqrt(Variance(series))
}

// MakeBins returns a slice of @TimeBin according to block sizes and time shifts.
func MakeBins(starttime time.Time, endtime time.Time, blockSizeSeconds int64, blockShiftSeconds int64) (bins []TimeBin) {
	timeInit := starttime
	blockDuration := time.Duration(blockSizeSeconds) * time.Second

	for timeInit.Add(blockDuration).Before(endtime) || timeInit.Add(blockDuration) == endtime {
		b := TimeBin{Starttime: timeInit, Endtime: timeInit.Add(time.Duration(blockSizeSeconds) * time.Second)}
		bins = append(bins, b)
		timeInit = timeInit.Add(time.Duration(blockShiftSeconds) * time.Second)
	}

	return
}

// IsInBin returns true in case @timestamp is in half-open interval @bin.
func IsInBin(timestamp time.Time, bin TimeBin) bool {
	if timestamp.After(bin.Starttime) && timestamp.Before(bin.Endtime) {
		return true
	}
	if timestamp == bin.Endtime {
		return true
	}
	return false
}

func Min(values []int64) (min int64, e error) {
	if len(values) == 0 {
		return 0, errors.New("cannot detect a minimum value in an empty slice")
	}

	min = values[0]
	for _, v := range values {
		if v < min {
			min = v
		}
	}

	return min, nil
}
