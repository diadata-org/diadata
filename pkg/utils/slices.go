package utils

import (
	"math"
	"sort"
)

// ArgsortableSlice is a wrapper struct around the sort interface. It allows
// for implemetation of argsort for all sortable types.
type ArgsortableSlice struct {
	sort.Interface
	idx []int
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
