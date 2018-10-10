package estimator

import (
	"math"
	"sort"
)

func computeMedian(samples []float64) (median float64) {
	var length = len(samples)
	if length > 0 {
		sort.Float64s(samples)
		if length%2 == 0 {
			median = (samples[length/2-1] + samples[length/2]) / 2
		} else {
			median = samples[(length+1)/2-1]
		}
	}
	return
}

func computeQuartiles(samples []float64) (Q1 float64, Q3 float64) {
	sort.Float64s(samples)
	var length = len(samples)
	if length > 0 {
		if length%2 == 0 {
			Q1 = computeMedian(samples[0 : length/2])
			Q3 = computeMedian(samples[length/2 : length])
		} else {
			Q1 = computeMedian(samples[0:int(math.Floor(float64(length/2)))])
			Q3 = computeMedian(samples[int(math.Floor(float64(length/2)))+1 : length])
		}
	}
	return
}

// RemoveOutliers Cleans a data set it accordance to the acceptable range within interquartile range.
func RemoveOutliers(samples []float64) (filteredSamples []float64) {
	if len(samples) == 0 {
		return
	}
	Q1, Q3 := computeQuartiles(samples)
	IQR := Q3 - Q1
	lowerBound := Q1 - 1.5*IQR
	upperBound := Q3 + 1.5*IQR
	lowerIndex := 0
	upperIndex := len(samples)
	for index, value := range samples {
		if value < lowerBound {
			lowerIndex = index + 1
		} else if value > upperBound {
			upperIndex = index
			break
		}
	}
	filteredSamples = samples[lowerIndex:upperIndex]
	return
}
