package utils

import (
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
