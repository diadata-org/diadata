package filters

import (
	"math"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

// Filter defines a filter's methods processing trades from the tradesBlockService
type Filter interface {
	compute(trade dia.Trade)
	finalCompute(t time.Time) float64
	filterPointForBlock() *dia.FilterPoint
	save(ds models.Datastore) error
}

// RemoveOutliers Cleans a data set it accordance to the acceptable range within interquartile range.
func removeOutliers(samples []float64) []float64 {
	if len(samples) == 0 || len(samples) == 1 {
		return samples
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
	return samples[lowerIndex:upperIndex]
}

// ------------ Auxilliary function for removeQoutliers -------------

func computeMean(samples []float64) (mean float64) {
	var total float64
	length := float64(len(samples))
	if length == 0 {
		return
	}
	for _, s := range samples {
		total += s
	}
	mean = total / length
	return
}

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
