package utils

import "errors"

// MaxIntSlice returns the maximum of a slice of ints
func MaxIntSlice(sl []int) (int, error) {
	if len(sl) == 0 {
		return 0, errors.New("empty slice")
	}
	max := sl[0]
	for _, number := range sl {
		if number > max {
			max = number
		}
	}
	return max, nil
}
