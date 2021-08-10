package filters

import (
	"errors"
	"math"
	"sort"
	"time"

	"github.com/diadata-org/diadata/pkg/dia"
	models "github.com/diadata-org/diadata/pkg/model"
)

// Filter interface defines a filter's methods processing trades from the tradesBlockService.
type Filter interface {
	compute(trade dia.Trade)
	finalCompute(t time.Time) float64
	filterPointForBlock() *dia.FilterPoint
	save(ds models.Datastore) error
}

// RemoveOutliers Cleans a data set it accordance to the acceptable range within interquartile range.
// It returns the cleaned data slice plus a slice of lower and upper index bounds.
func removeOutliers(samples []float64) ([]float64, []int) {
	var indexBounds []int
	if len(samples) == 0 || len(samples) == 1 {
		return samples, indexBounds
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
	indexBounds = append(indexBounds, lowerIndex)
	indexBounds = append(indexBounds, upperIndex)
	return samples[lowerIndex:upperIndex], indexBounds
}

// computeMean returns the weighted mean of @samples with @weights.
// Special case of non-weighted mean is obtained by setting weights to constant 1-slice.
func computeMean(samples []float64, weights []float64) (mean float64, err error) {
	var totalPrice float64
	var totalVolume float64
	length := float64(len(samples))
	if length == 0 {
		return 0, nil
	}
	if length != float64(len(weights)) {
		return 0, errors.New("computeMean: samples and weights not of same size")
	}
	for index, s := range samples {
		totalPrice += s * math.Abs(weights[index])
		totalVolume += math.Abs(weights[index])
	}
	mean = totalPrice / totalVolume
	return
}

// ------------ Auxilliary functions for removeOutliers -------------

func computeQuartiles(samples []float64) (Q1 float64, Q3 float64) {
	sort.Float64s(samples)
	var length = len(samples)
	if length > 0 {
		if length%2 == 0 {
			Q1 = computeMedian(samples[0 : length/2])
			Q3 = computeMedian(samples[length/2 : length])
		} else {
			Q1 = computeMedian(samples[0:int(float64(length/2))])
			Q3 = computeMedian(samples[int(float64(length/2))+1 : length])
		}
	}
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
